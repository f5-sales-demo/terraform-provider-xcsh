---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3231313122013121-2311023230023202-3030203231102013-2232010210223031-2122202021201100-2101203112321322-3213221000101002-2012223322330032"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 233123010113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1313302310211211-0113000002100212-1230302020211203-1330131333101033-1202032031212013-1220122010102001-1322323201222003-3212101223201010"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120211322211213-0113220313132213-0232321011302020-2032103221003110-3310122203000211-3201100321001103-1032221212331111-2203112111001032"></a>

## Direct properties — interface_ip_map / 233123010113 / 3

<a id="canonical-0030300301030203-1120100213002013-3032021333122330-1311013303013220-1201322231023021-3011301320320303-2301333013302011-1131113201103001"></a>

<a id="canonical-0123220033113122-3122312211322110-0321022123103111-3210333331020123-1011210103033201-0230112032211300-1020210230332030-3012030221033213"></a>

## interface_ip_map property — interface_ip_map / 233123010113 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-3120210022100111-0223130130301103-1003322322100111-1223023330203131-1332312220222031-0001102113302232-0010233012222032-1323033212333113"></a>

## Next pages — interface_ip_map / 233123010113 / 5

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2211332222002230-0122303122321113-3201122320313333-3310131113000011-1101321120213101-1203213233003121-2033310231011313-1012230033113111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203021110000210-3013032221122322-1302013030001012-2000001223120232-3312021321131301-1102021033101102-3212030312313021-3120113013313132"></a>

## gcp.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 031203133100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0023133312311013-1020102120132013-2031120023102022-1211021223300330-1023121303100112-2223203122111221-2102322312220112-2111001012331101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103003321201033-0110312332323131-1011222001213001-1233001133222303-2023111323321231-0211020010311031-0130301210133332-3223213311203203"></a>

## Direct properties — ethernet_interface / 031203133100 / 3

<a id="canonical-2202131113031322-2202201233333312-2100212212322300-0120200110203310-0210133313210132-2120111033323012-1232132323300323-3332212003132231"></a>

<a id="canonical-0310201321020301-3130102222332233-1211301012131013-3231213103001010-1303033132320133-3130023111003100-3113201232110100-3101320113312203"></a>

## device property — ethernet_interface / 031203133100 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2033232212010332-3213011001213120-3320310112330011-1022111033111320-1112031131320012-1300001230032031-3100221130230213-3122323200300212"></a>

<a id="canonical-0212102332012112-2131023313121022-0212133010020221-0221131310200132-1323332201033210-3012201120130011-2330121133322232-2121321330210313"></a>

## mac property — ethernet_interface / 031203133100 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-1321311030023021-1211201010133102-3110323032133000-1010322132011313-2221031031232102-0320133032230131-3223322022011130-0330122102333222"></a>

## Next pages — ethernet_interface / 031203133100 / 6

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310330212312213-1020322321000203-2221131002021013-0203111133111211-2333010020211010-2323200230311221-0023310120021313-2221132321123133"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 202300111301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3121303103203100-1101323212230321-3230012332132301-0232101222130000-3122012312020000-2133013130121123-3013331110201202-2233232321003333"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212303223332020-1230100030022222-0001033322003221-2320332021020310-2103301230200203-3023001002300312-3331133212101311-0112132223121210"></a>

## Direct properties — ipv6_auto_config / 202300111301 / 3

- [host](resources--securemesh_site_v2--reference--group-010.md#canonical-2112021120030020-3231333020020100-1131200101210103-0002113133023111-0231212221103103-2232011230301021-2320301133223222-0030010301313122): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020): complete subsection reference.

<a id="canonical-1233332221033313-1232232011320201-1203203023101220-0011012212003001-1010312023003330-3011333331120113-2320300132032102-3123300222221320"></a>

## Next pages — ipv6_auto_config / 202300111301 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-010.md#canonical-2112021120030020-3231333020020100-1131200101210103-0002113133023111-0231212221103103-2232011230301021-2320301133223222-0030010301313122)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2112021120030020-3231333020020100-1131200101210103-0002113133023111-0231212221103103-2232011230301021-2320301133223222-0030010301313122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313300201011222-0222202332320311-3130131312013133-1330103323331323-3021022232323023-1131121012332003-0033200013320023-1200033031022202"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 122023131003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0302330002233201-0301231330211202-3230300313203330-3131120002000011-3130030003000102-3300230131112320-1023333003302222-0312220133101033"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-3222012001313201-2010033221223303-0230203131021112-2313121020332203-1303102213233233-2302022211303230-2222230100103311-0010020023311333"></a>

## Direct properties — host / 122023131003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203032233202220-2231101312231001-3323311101201231-2223202033010022-1231021031221131-3000032222211332-1321122021003110-3123103312121312"></a>

## Next pages — host / 122023131003 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130123211210333-1013001111102000-1032200332203301-2020133211011322-1030231310122023-3211232101022022-2203030322113113-2010010221122332"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 021310313030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-3210112001333003-0330021231033233-3213023112222300-1031231133322111-2310133000031301-1112133102213312-1022130100321013-2131012312110121"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113011213030021-3201113011230222-3123113232231210-0002212332003301-0110223132021111-1233130133101210-1212023333201030-0113111130320030"></a>

## Direct properties — router / 021310313030 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011): complete subsection reference.

<a id="canonical-2301212332222121-1302223311302010-0301321012231103-3302110312321200-1120323333210321-0221013211111323-2333130012210221-0012220300110130"></a>

<a id="canonical-0132100111232222-2031120321111011-3121330103012013-0222311211301311-0133032201033123-1100002323102233-2013112300011300-1312031110103231"></a>

## network_prefix property — router / 021310313030 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130): complete subsection reference.

<a id="canonical-1212123003333112-3111103321002010-0313323002121002-0303313101311032-3310213201313330-1212331330123131-2322231123222102-3120012001203331"></a>

## Next pages — router / 021310313030 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102213311000311-3000002320301321-2312312132302000-3023202032222003-1200331222331111-2310223201023111-3220133200133031-2013030302313012"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 021131001300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2213011013201101-2003230021030321-3323100213110131-3123221023310300-1001112002020010-3022032332301110-2221103300221213-3312003012121320"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001001130101230-0223102000101223-3121102112333110-1020321130012323-2200101203210321-3330000200031122-3312112121333322-0222113202332210"></a>

## Direct properties — dns_config / 021131001300 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-010.md#canonical-0102032310030203-1222102212030202-3102110031010021-1011113132210110-1231001002302130-1110122112302011-3013300020220032-2102312101013033): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010): complete subsection reference.

<a id="canonical-3330101000113103-1212112322321322-2013201012313220-3201113310020303-3211122303331121-1332120332122210-2231101233120003-1232332031211213"></a>

## Next pages — dns_config / 021131001300 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-010.md#canonical-0102032310030203-1222102212030202-3102110031010021-1011113132210110-1231001002302130-1110122112302011-3013300020220032-2102312101013033)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0102032310030203-1222102212030202-3102110031010021-1011113132210110-1231001002302130-1110122112302011-3013300020220032-2102312101013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013023121113213-1020012023010223-1112212213002303-3023201203123231-0333013222223301-0312230233222202-2031013100221133-0020103303102201"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 330031322332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3123210121001102-3023303213112331-0331313333133230-3113031202111230-1312021301120101-0233022232000230-1322322110222010-1130023212320222"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132021200033202-1021103333233102-3111333030220100-1200223012313032-1323011013130330-1011011003220130-1110030011220312-3003302223013013"></a>

## Direct properties — configured_list / 330031322332 / 3

<a id="canonical-1112231303131122-2032000212231133-2121033310320230-3310312332230020-2333321310312222-2113122210130123-3212333231103033-1030020132210223"></a>

<a id="canonical-2223223012130102-0130201003202122-1200323000101021-3220130001021030-1233233011233020-1103233123312223-1111121200023302-1332032323301203"></a>

## dns_list property — configured_list / 330031322332 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030221123031220-1111102231213201-2123332013221203-1112112300003221-2022033010133000-2223133002030130-2210021302130331-3330202101302233"></a>

## Next pages — configured_list / 330031322332 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013002322021312-1031102021011121-2200020120031221-0300230133123001-3211032321000121-1033123003312313-2301331223200222-3323213303312221"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 203022013223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1333202111011101-2322313311320030-3322221213033032-1100320130213130-0022201020232123-0313311032011212-3230323103223112-0303212322210322"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300201000111113-2211211223220000-0312130023031211-2131313001331132-0023122320311211-2032230012021330-1200232310320213-1321033111211020"></a>

## Direct properties — local_dns / 203022013223 / 3

<a id="canonical-0230000303331021-3330000101003213-1132221100232322-3032102013110103-1221303232312113-3120233223023003-3211303212132202-0030020013000122"></a>

<a id="canonical-3310322210110203-1021120012021030-1032131302022112-0001032310231010-1033103102221010-2223000133303222-1301001101200302-3102003300213131"></a>

## configured_address property — local_dns / 203022013223 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1232122223032023-2013112221302031-0133111212023311-1012220213120000-1200313130123323-2313031311232022-3012133131113323-1220233303023232): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2210013013203202-3031130011301312-0023012332231020-2231032133323111-0220213130000200-1232011331021120-2111021033331100-0321031223002303): complete subsection reference.

<a id="canonical-0210121002322003-3120110022233133-1010130012303303-2120132122131002-0202112132221331-0222330223130323-2301303103023102-2301312002323023"></a>

## Next pages — local_dns / 203022013223 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1232122223032023-2013112221302031-0133111212023311-1012220213120000-1200313130123323-2313031311232022-3012133131113323-1220233303023232)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2210013013203202-3031130011301312-0023012332231020-2231032133323111-0220213130000200-1232011331021120-2111021033331100-0321031223002303)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1232122223032023-2013112221302031-0133111212023311-1012220213120000-1200313130123323-2313031311232022-3012133131113323-1220233303023232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130021201121113-2330020130211100-3033103320012302-3300101213123331-3333311202011003-0302321002331212-3310001312001000-2032331020210120"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 231022112100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2323120301123201-0031101110322311-3221210102333331-2312131030022202-2001333002123330-3113110300112102-1121313110201123-2101111230022110"></a>

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
first_address = {}
```

<a id="canonical-3131111000012000-1022313132231022-0011301213302103-1110300303311031-1120112122000000-2103133310111211-3310203233022333-2313030003031221"></a>

## Direct properties — first_address / 231022112100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032300120032213-2103120132333032-3200121210033230-1010312001330112-0020122223010222-2100231112121000-0111030310002111-3001101230111211"></a>

## Next pages — first_address / 231022112100 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2210013013203202-3031130011301312-0023012332231020-2231032133323111-0220213130000200-1232011331021120-2111021033331100-0321031223002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021133321233121-1311212302100302-1002022101303030-3012303233111130-3002033002130131-3332132211100221-1201313023231311-0101202021000120"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 221111200112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2221011102121103-0302332000001202-3101322231233000-1120031101131011-3333031313320232-3310000233231301-3321223023323322-1222202130210102"></a>

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
last_address = {}
```

<a id="canonical-3111033320100002-1020120123012302-1320122113033101-3333231322021023-2031011132001221-1331213300002120-2121011310222103-2212010133322002"></a>

## Direct properties — last_address / 221111200112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332022322203322-1123102112212103-0301011303310221-2200111122031132-2022211022111033-1123202001120301-1203030233003310-3201320233210102"></a>

## Next pages — last_address / 221111200112 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113001123101303-1112321112012010-2033300300012103-1112133022222212-1031120321303113-1132121000330202-0222103001000121-1223010012130331"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 020211213230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3001330023321311-3203102220022223-0311033020031020-2231211301120031-3110222311202121-2111123011323230-1021332132220201-3320331210022322"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330302113001023-1321212001100313-1011021211312101-0221213331133113-1232100233001331-3232321303030302-2020222131232110-1221130301002300"></a>

## Direct properties — stateful / 020211213230 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-3212320310102033-0301112300311120-1023311303202320-3233102100033320-2301032121021000-2230222231131023-0302320031301132-3110201000111213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-2333333223310132-3110110313231231-1312033123112312-2332001202002112-2332300300203331-2300212013213232-3331122003212013-2213303000300330): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133): complete subsection reference.

<a id="canonical-0021332011031012-3110302121130110-0123223222003313-0331200121321200-0011020120202203-0012002201231322-1321333323002300-1312231302122220"></a>

<a id="canonical-2122211130331213-2203113021201211-0030022031230310-1110131032010112-1030200332120322-3000211302311210-3303120130013220-2011300123002013"></a>

## fixed_ip_map property — stateful / 020211213230 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-2123230302111331-0223331032011030-3331322202310202-2202033132112023-0103133102323022-3233330203303323-1200202112333200-2102203333031113): complete subsection reference.

<a id="canonical-3120303311222110-3303302203010000-3323012103120233-3010302032013110-3211332223133101-0330303202231301-2131132321301223-2300303012030301"></a>

## Next pages — stateful / 020211213230 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-3212320310102033-0301112300311120-1023311303202320-3233102100033320-2301032121021000-2230222231131023-0302320031301132-3110201000111213)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-2333333223310132-3110110313231231-1312033123112312-2332001202002112-2332300300203331-2300212013213232-3331122003212013-2213303000300330)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-2123230302111331-0223331032011030-3331322202310202-2202033132112023-0103133102323022-3233330203303323-1200202112333200-2102203333031113)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3212320310102033-0301112300311120-1023311303202320-3233102100033320-2301032121021000-2230222231131023-0302320031301132-3110201000111213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113030303201222-1300121013112313-2010013332231031-2130000233112332-2130313011020013-0123130211030132-3321030031123211-0333131323131003"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 100320302123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0212130303131131-1102021233313300-2023333013021232-1211011113311032-2120123130300011-3210320202010121-0121323331332112-3011302001300021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-2010133221113210-0312100113002231-0121130133232331-2101012131211023-3312222100310323-2130100203000201-0331310231301031-3323022320210211"></a>

## Direct properties — automatic_from_end / 100320302123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303120020030332-3022333323202322-0302023320001203-3032021232320321-0302231010031022-3213313310032030-2322102111301313-0101232120303120"></a>

## Next pages — automatic_from_end / 100320302123 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2333333223310132-3110110313231231-1312033123112312-2332001202002112-2332300300203331-2300212013213232-3331122003212013-2213303000300330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331112232133111-0202020222002033-0132310220333133-0012033021123111-3221300303011121-3011123031002013-2100023010121300-3020301333202021"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 003313312021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2303201111302100-3301232333111023-3301301331002202-3013011322310302-3021013121132031-3323223133323001-3030303313322303-2001102320332130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-1130300110330231-3221021130331022-1312030330012321-3313122020223103-2223110032020230-1212300110332331-0122221112220121-3210333122020102"></a>

## Direct properties — automatic_from_start / 003313312021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033122323222022-2203101211101113-2013322233133233-0202231032322130-2110322101131310-0030011332231311-0210301002132032-0112331201120302"></a>

## Next pages — automatic_from_start / 003313312021 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013323000100211-3322031013120032-1220311321031201-0010100231010231-2333112313111331-1300000233133231-3210123220221333-3021210010202121"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 003231222220 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-3032131303303123-3031021131012011-1303332333121322-2021100231011123-2132132301300023-1322020110231131-0200112331213101-0213322023030112"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012110220121203-2002031103311112-3101010330302022-1030233312231000-2232101023110302-3001030111201021-3033102012333212-0213301101213312"></a>

## Direct properties — dhcp_networks / 003231222220 / 3

<a id="canonical-2302110133322310-1302122313132313-1131232013211321-3001303121201112-3010302230131303-0312310133311013-1333203102302111-2023012300200222"></a>

<a id="canonical-0132100222333121-0021200122231313-1200232001123201-2202012211230212-2313231002113001-3303123332333300-3120222311021330-2021301023132330"></a>

## network_prefix property — dhcp_networks / 003231222220 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-0222320203033033-0033111031133331-0220110232320102-3232101133010200-3302130213212232-3033112211103130-1300220100020332-0100103222232202"></a>

<a id="canonical-0001311001120003-2210230332332230-2302101120022012-3233302323012132-1101212302232220-3200331011000132-1302132133023123-0002032300130310"></a>

## pool_settings property — dhcp_networks / 003231222220 / 5

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-1230002211003033-1030231113023001-3001220103002231-1203301211030111-0212030230320320-2203213023322302-0333001233001230-1323332133221101): complete subsection reference.

<a id="canonical-2310022110211003-3213102303102231-3120232223003000-3210231330023122-2020022032033111-2321221110021322-1022331000112013-3223103300202223"></a>

## Next pages — dhcp_networks / 003231222220 / 6

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-010.md#canonical-1230002211003033-1030231113023001-3001220103002231-1203301211030111-0212030230320320-2203213023322302-0333001233001230-1323332133221101)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1230002211003033-1030231113023001-3001220103002231-1203301211030111-0212030230320320-2203213023322302-0333001233001230-1323332133221101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132010212120013-3032303322110031-1131232102323020-2032112211220322-2120030123000102-3002112000001220-0033112010031211-2310033200100213"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 201332212132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1301331111221302-2231132131002011-3202012230122031-3022201120301320-0303030221002112-3111233031010002-0300233320210100-0212203002230111"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230220112302332-2033213011101010-0131213313301333-1203110012320003-2100202112012021-3220112302200021-3301112103322302-0112022131021032"></a>

## Direct properties — pools / 201332212132 / 3

<a id="canonical-3031012131130022-2331211101022312-2301100110131311-3332120113130111-2121030013012213-1033120212301120-3120333333323131-3233000311323123"></a>

<a id="canonical-3121200122102103-0303211232210101-0003200322130033-0111023032132222-0102332002303032-0231122031230021-3300033012022112-2221313203122102"></a>

## end_ip property — pools / 201332212132 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0300021023312213-1313021321023300-2200233100033202-0233202100310330-2203200020121020-1202131310100023-1120011211123323-3030131323232310"></a>

<a id="canonical-2023301010220310-2133120221121012-2230322300333033-3002220201311220-1000112220013231-0120330101233132-2020022312220233-3303210123113333"></a>

## start_ip property — pools / 201332212132 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3031201132010021-3110222210031223-3101320013001013-1112023030302233-2102320322002002-3302211123120012-0212032231231201-2200301201022010"></a>

## Next pages — pools / 201332212132 / 6

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2123230302111331-0223331032011030-3331322202310202-2202033132112023-0103133102323022-3233330203303323-1200202112333200-2102203333031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112120003102123-3320003022101133-0012312330120133-0303300110230230-1010113011301001-1310033303331013-2221310220333000-3200331111203011"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 212131030002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0013010012301321-1123311311233212-3001101303321311-1010131111012030-2303032020003212-0202203120213213-2322331112301301-2000023020000300"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102011230231113-0313102010320102-0313201112031131-3202313010230233-3320133230023232-2001010330110302-3000122233323122-3213303100010001"></a>

## Direct properties — interface_ip_map / 212131030002 / 3

<a id="canonical-1332001203121113-2131303203010000-1230133333000223-0213111001100201-3322223303331023-3002002301110021-3033323021010333-3312120001030001"></a>

<a id="canonical-2000210013220201-3133230223032131-2320103000001013-0330203200023133-2303021132232221-1031022020111201-1212120003133233-3000322122032032"></a>

## interface_ip_map property — interface_ip_map / 212131030002 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-0032312313021220-0103203110212230-2100130223223103-3323033110031012-1033313313232100-2220121330101230-3022301232330030-0120303003113001"></a>

## Next pages — interface_ip_map / 212131030002 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1211223311200301-2113213233010230-1203012301302120-1022300331103323-0211120332111200-1030031000001303-2222221212332011-1213312211122110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010323023130001-1302311030000302-0112313012022112-0312211003202222-1120302121103203-2211121110231102-2022300202200322-2202103011012201"></a>

## gcp.not_managed.node_list.interface_list.monitor — monitor / 001222212200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.monitor

<a id="canonical-3011133112101111-1322133303013213-1220021130330333-2013232133102032-2211133323113012-1303332000003312-1233312330202120-2001001230310112"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

<a id="canonical-1101123300021310-1110230231332323-0020023211302023-3110301330032012-2100213122333221-0120331210123123-2113301322301301-3300211010102202"></a>

## Direct properties — monitor / 001222212200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100201023320231-1011110230303320-3322021322232111-2302012002000232-0233111231323121-1300022133323022-1112223031132120-2311201302000001"></a>

## Next pages — monitor / 001222212200 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2110013100330201-0203313102212200-0032133203312013-0101212301101020-2313122133233302-2220111322101030-0222330312032021-2111000113010023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200200130121012-0131021022233301-1212021232100033-1023022223030233-1123103213120023-1320230033313133-0010313103213333-0202132102001011"></a>

## gcp.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 330312212332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3122133323111102-0022203232223232-0101211033330010-3123332022111121-0213332222223023-2031020102303122-1013330110203130-0101312331313132"></a>

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
monitor_disabled = {}
```

<a id="canonical-3230023300303233-3232102123333132-1021213111110130-0120133221023301-2111101310033110-1231011000330113-2002021101033302-3230110133310201"></a>

## Direct properties — monitor_disabled / 330312212332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123020221231311-2002203202311131-0310303220312103-1230212002001200-1310020313033020-2012203230331130-0331332130110120-2101332010230100"></a>

## Next pages — monitor_disabled / 330312212332 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210001110232132-2023233123002131-1023120103302102-3000012331100321-2032213022213202-3213301122320300-1001230032212003-2311132303000120"></a>

## gcp.not_managed.node_list.interface_list.network_option — network_option / 231230101032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.network_option

<a id="canonical-1330310330110230-0303011333132301-2230311210222301-1331313130303233-3030211121111233-0031031131020001-2232101330303303-3113130033321111"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002021013131032-2110232200331332-0203010203212023-3302000202031200-0120333100202100-2331110011312110-0320031301311330-2233001333103010"></a>

## Direct properties — network_option / 231230101032 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-010.md#canonical-3333320202012202-0002100001111310-0103330113000112-3112313330312312-2001110311222301-3203002120212201-0201300202312003-0120323202330110): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-010.md#canonical-1002102000211210-2202011122321031-3100013213211323-2200032023330011-0013000130200301-0012233320202130-2320203003312011-3021003110001201): complete subsection reference.

<a id="canonical-3031013012312000-2031213022313011-2321330222301200-1230010000312220-0231310211332030-2111132333011301-1123123232331021-2103210231102203"></a>

## Next pages — network_option / 231230101032 / 4

- [gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-010.md#canonical-3333320202012202-0002100001111310-0103330113000112-3112313330312312-2001110311222301-3203002120212201-0201300202312003-0120323202330110)
- [gcp.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-010.md#canonical-1002102000211210-2202011122321031-3100013213211323-2200032023330011-0013000130200301-0012233320202130-2320203003312011-3021003110001201)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3333320202012202-0002100001111310-0103330113000112-3112313330312312-2001110311222301-3203002120212201-0201300202312003-0120323202330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301231321002123-1213033120213323-0001133301322130-0221220010133111-3121311123333301-2021013323033003-1302120223020012-2032103230021332"></a>

## gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 123022010011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3232031223301220-1213100132311120-1112330021032320-3112320303030321-1022311322333030-1023231022220202-1321322333013331-2003312132323331"></a>

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
site_local_inside_network = {}
```

<a id="canonical-1103321322200210-3132233313102330-3031213030203321-1102312101120132-0102111200112011-2310311302322122-0312223312230332-0221322311031322"></a>

## Direct properties — site_local_inside_network / 123022010011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232202313033113-0311130313302322-1322312020303111-2200330112113012-3120302322123210-0223313301010321-1113303002331211-1210200100112012"></a>

## Next pages — site_local_inside_network / 123022010011 / 4

- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1002102000211210-2202011122321031-3100013213211323-2200032023330011-0013000130200301-0012233320202130-2320203003312011-3021003110001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101030000301020-2110000301322100-3333100330321011-0001010020323313-2120011012332321-3320230131210021-0231100110112100-0021211233332031"></a>

## gcp.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 011300022131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- gcp.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1101130122001330-3200313132311102-0012333132101320-1012003123203130-1021002313023211-3101101032010201-0211013312312312-1132332022032300"></a>

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
site_local_network = {}
```

<a id="canonical-1100020221330311-0101311113031011-3321203012031013-3122011120111303-0002230230203331-3101333111220200-2303020021303201-3010121113221133"></a>

## Direct properties — site_local_network / 011300022131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032103131132101-3003102220020310-2321330013320102-3232112013303310-3223122102201012-3000111210232101-0303223302010122-1320222001233200"></a>

## Next pages — site_local_network / 011300022131 / 4

- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1203323220122012-2201322220300230-0230002310100030-3330031323321311-0300313103011131-2300131311231110-3013203313301020-2320101230133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101032332000211-2102333220120032-0010222302000320-0221212001333020-2311121111322023-1112310032302333-2102031220122231-3203231313331222"></a>

## gcp.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 130031321223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-1221003113232322-3112200022301210-1213101331310110-2002021133031101-0001100213131323-2103311331231201-0303222121302212-2321322032021313"></a>

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
no_ipv4_address = {}
```

<a id="canonical-1110223313031123-2221023330113130-1211300133000001-2321022312223211-3112211002110133-3202113310100031-3113330110123202-0133113322213120"></a>

## Direct properties — no_ipv4_address / 130031321223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101113031223012-3323001132101203-1030233211102020-1102202021221201-1003330332230022-2110102233221102-2122210010311133-0303103112310222"></a>

## Next pages — no_ipv4_address / 130031321223 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3000103023231312-2301010323330333-0333003302000002-1322000323123232-2002220330300031-1331112001201102-2321333303303203-2001023303022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210000113331333-0012130210103122-1121032003221032-3223330222213221-1103113003111220-1313222033002010-2222333020323201-2011013022210010"></a>

## gcp.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 023222010321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2232231013323133-3223302223010013-2303030223033313-0000223120321231-3320113130310310-3021130112031220-1333020210212230-3233012201232202"></a>

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
no_ipv6_address = {}
```

<a id="canonical-2000331020033313-2000220321303011-1323221000010210-1032020030200220-2202333101332122-1013302223313033-0201113103312002-1032021223220103"></a>

## Direct properties — no_ipv6_address / 023222010321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013321002110222-0033021103223033-1231203203310101-3303003113030320-2201000303311001-0112101003201330-2221223131221310-2330332111331012"></a>

## Next pages — no_ipv6_address / 023222010321 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1133201233031230-3333012130033031-0221332312123131-3003211133222313-1013122333012331-0100231220131122-1212222211103002-0131010321033230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131311020200101-2010022010212313-3132220020303212-2002011211003222-1030201122122203-0201200133300003-0103202003213130-0103301302032131"></a>

## gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 001103321200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3201230020221132-3322132103310220-0113330033100031-2030100300232321-0303012020132333-0212222101002330-2100330100100002-1110313022130022"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-2121331022203312-2232010303211210-3030010132332112-3212213211213132-3003120312033122-3133212312100012-0133213132013323-0011300313001001"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 001103321200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120001100133200-2330212300011310-3102211333021230-0013323123332320-1011302121230102-0222222132102110-2031201200110322-2210030331012220"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 001103321200 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220100020332132-3221320131203333-1233013032031323-3222113210012223-1121032113121201-1001013201031031-1210302012100032-0323200312201112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030301121130011-2031233312222010-2010020330312213-1301230103122330-2300003331131210-1012110201301131-2120100220131132-2130121222003201"></a>

## gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 133202021303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0333100100202200-1032122232320002-2320331101223032-3010301203011122-1333021033220321-3302321030121220-0133000323223100-2301223332112220"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-2110101212310303-2033230331223300-3213102301000233-2021330200102213-1321121111203330-3111230302220321-1331331223302220-0212312101122213"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 133202021303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233112322010002-0003321311223123-2221332213301010-3320031233111310-2200000000003102-3323011212103330-2313330211330221-0223321203102222"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 133202021303 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1030120100013303-3022221311322001-3212222021021102-1131222030210221-2220333320001121-0123333131111212-3121032230202313-3133313121120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220212000312130-1030123323222113-2313033010330311-1301011322232123-0000220021023113-3103013310312111-1011200300022133-3102033303302203"></a>

## gcp.not_managed.node_list.interface_list.static_ip — static_ip / 113020221011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.static_ip

<a id="canonical-3322232032012023-3111223102312021-2220310300031211-1123112210221212-3131221102230120-2022320323303220-1331013100101100-3011021002121002"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302210312103131-1202101301031001-1320322332211223-2021213211301313-0011222001201132-2322111120021011-3232221232133110-1122030112323332"></a>

## Direct properties — static_ip / 113020221011 / 3

<a id="canonical-0012032321231030-2231232231130333-3012020210002013-1202201203330030-3102101311221332-3100032101010022-0021032303111100-3011002033201330"></a>

<a id="canonical-0201310103313313-1012223000122123-0021320130332231-3333312020030201-2323311113330331-1032222302010331-1001030113313332-3212300230022203"></a>

## default_gw property — static_ip / 113020221011 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3020120111232210-2102110103323202-0010231223121032-1203213213222112-0012120310230320-3331021020120323-1112112321231121-1230333110220110"></a>

<a id="canonical-1313002200212032-3013100032132022-0100022112300132-1032123223231323-3101330111311003-1032021201013323-2001222000303210-2321020101011220"></a>

## dns_server property — static_ip / 113020221011 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3113320003112323-2023200210222112-0231103320303012-3003333000330332-2120201113112323-1123130120312120-1111033303223320-0310223122221311"></a>

<a id="canonical-1212121330223332-1101133301312210-2103312323032000-0201301211130001-0103233123103231-3102112101023110-3313223020133222-2131231200121113"></a>

## ip_address property — static_ip / 113020221011 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1122002321023211-2103002032133023-0211003311103232-0031012122000202-2000120000300220-0302312033300121-1200303023202001-1331230222101011"></a>

## Next pages — static_ip / 113020221011 / 7

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010303301000221-3220301322311313-1312021001323010-2220322233230222-3203013131020201-2112003322003012-3001110323323332-2110322210032311"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 211013000002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3203201112332013-2123213103300133-3111100312201222-0122232323111013-1131322330103333-3311010222120313-3031010030320213-1010212220012232"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202132233232220-2122223022012122-3301033231330133-1331120311303320-2312331221100110-1102003202203113-0110020232110310-2100221131121011"></a>

## Direct properties — static_ipv6_address / 211013000002 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3020020320210103-3302022223103100-3202322112032131-1311112321131031-0313211102210332-2300213000303300-1200033322223220-1232222222212002): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3212312131231313-1103210201121200-2232230110113111-3022100023003203-2030230112233301-3002203022132033-1021303122001320-2212221313021301): complete subsection reference.

<a id="canonical-1333100210232110-3203233230120333-1310131000130212-3023201233100312-1221133320310332-0320233313321301-0322323320331023-2133231013202303"></a>

## Next pages — static_ipv6_address / 211013000002 / 4

- [gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3020020320210103-3302022223103100-3202322112032131-1311112321131031-0313211102210332-2300213000303300-1200033322223220-1232222222212002)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3212312131231313-1103210201121200-2232230110113111-3022100023003203-2030230112233301-3002203022132033-1021303122001320-2212221313021301)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3020020320210103-3302022223103100-3202322112032131-1311112321131031-0313211102210332-2300213000303300-1200033322223220-1232222222212002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200023132232023-2232021233201212-3310310000003010-2032123031011220-2013303200311303-3001310231000022-2130201131133021-0031101022320102"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 021111310102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2332222231303130-3110133130001320-3122102203023232-0212000332330203-0311310133321311-2320010022302101-1320320312012210-1221101321321103"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323131312003122-1220303102001021-2111222200012001-0011222112223102-3320100122032013-0002030303021121-2223202000303020-0012132131000002"></a>

## Direct properties — cluster_static_ip / 021111310102 / 3

<a id="canonical-0011332330010000-3333112302132303-3333132311113310-1331333002020330-2110303131112111-1323113132101032-2033002001332203-3323030223130321"></a>

<a id="canonical-2013112101300313-1111313113331122-1120210003323320-0112121331312030-3023032133201133-2131323311131203-1013021322120302-3221310023010012"></a>

## interface_ip_map property — cluster_static_ip / 021111310102 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-3122012331220201-3333011000333013-0123001233022102-2331310113232300-0321001233211113-0211111112321320-1332001010330223-3322210302333033"></a>

## Next pages — cluster_static_ip / 021111310102 / 5

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3212312131231313-1103210201121200-2232230110113111-3022100023003203-2030230112233301-3002203022132033-1021303122001320-2212221313021301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022201003100222-0331131310001122-2311220223013333-0302003333231101-0201231011322323-0010122020302100-3031311301300131-2011000303322202"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 332333203133 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-0322132333223300-0101202020231122-1322010031300301-1133112212123112-2021031023222133-1013210110123212-0202021020323201-1332033310120130"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113210131011231-0311122321120010-3231221223102320-0321223131000122-1121123331130203-1020211223000130-1303221030010331-3313020201000012"></a>

## Direct properties — node_static_ip / 332333203133 / 3

<a id="canonical-2133113323021033-2021231212310330-2020002122300310-3031321323220102-0023033201300221-2132231123123030-1313031123001330-0313303110233000"></a>

<a id="canonical-3322103321303002-0230012221321010-2121010202110122-0000130131112312-0301331013210213-1331231110000023-0232100310331220-2203021330131232"></a>

## default_gw property — node_static_ip / 332333203133 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2203010120021321-0012122132132023-0311001223032011-1023302201001000-2323220312032130-3330023121101312-3010121130000332-0203020210231200"></a>

<a id="canonical-0112010032023221-3012201102113200-3103313002212302-3132002011032333-2002233210203323-2313130321320221-1212202033321032-0201120122302003"></a>

## dns_server property — node_static_ip / 332333203133 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3111001121103231-3221330003020002-1120021320332011-0020302211002113-3121123113212232-3233210321312112-3302012203100121-1110200113332023"></a>

<a id="canonical-3003020221222003-1210101203323013-3110020301011022-1001213101210213-0000331310120210-3330111002103012-3002100002003330-0001132301031332"></a>

## ip_address property — node_static_ip / 332333203133 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-0330101331220030-0322133313313303-2132113121320002-3011221311032103-1031130332230030-3303233202032010-1312113101031133-1320100101321003"></a>

## Next pages — node_static_ip / 332333203133 / 7

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2322033023330322-3002023133132222-1111333320023023-1223323000130112-0000221001032331-3310323132320101-1112120210112123-1112101101200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101331012201120-0111322330113331-0132300221232210-3012102001011010-2120020011230001-3223012022301003-2203301022000222-1310232000010112"></a>

## gcp.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 031133132003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-1332210301322023-2303020332221120-1002032331033200-2001110113233121-2223002310212110-1320221130123100-2300113003123123-2211213003330310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320123210123331-2211113231003023-1022220111320332-0223202102030121-1222230112120231-3023021230022122-2032130133232330-0221113232323313"></a>

## Direct properties — vlan_interface / 031133132003 / 3

<a id="canonical-1302030203310223-2020122133303310-0000133203122212-1033112012332322-1333222100231230-1211200120031033-2111301203131110-1222220120323133"></a>

<a id="canonical-0020102010112031-1303232233000333-1301213201030021-2001323002121313-1213032033313112-2221310121201112-3333203112230031-0223103032311113"></a>

## device property — vlan_interface / 031133132003 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1232101330033202-3023013122202030-2211121022020220-1300213320323300-2322021013031013-3223313021233310-3020330010320010-2131211103331023"></a>

<a id="canonical-2300201200110202-2120100130230313-0122013222101021-1211001032300110-1303010113133211-2032011230202022-0331030130113102-0211000232133021"></a>

## vlan_id property — vlan_interface / 031133132003 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-3021113031232333-1032200110011321-1232211322300033-3121223330203131-0223222321212321-3111033222231202-2001020133220131-3333020132010312"></a>

## Next pages — vlan_interface / 031133132003 / 6

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222132123022222-3131320000113211-1310102203111333-0231011113211110-3133211011101210-0300223232103112-3201232130202002-3201021112330010"></a>

## kvm — kvm / 233002101233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- kvm

<a id="canonical-3200011220002023-3002321213101032-3320012101030001-0002211010233312-1133321102002011-0013220032121103-1330120002111201-1130031202103111"></a>

Type: `"object"`. single nested block, Optional.

KVM Provider Type. KVM Provider Type.

Upstream description:

KVM Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
kvm {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220033122013212-3213231020331322-1333211213201022-0200012110101200-0121011020323320-3120103021122121-1003203120310023-1333032101002212"></a>

## Direct properties — kvm / 233002101233 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230): complete subsection reference.

<a id="canonical-1213201221011113-2230123101132013-3233301311301011-2130110220202012-2323321332301333-0122032313022310-0320201332302122-0133312132321123"></a>

## Next pages — kvm / 233002101233 / 4

- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222103031110130-0021020320102031-3030223002321221-0330002232120321-2313121110021110-1120303223102031-3102020113032303-2003331020312202"></a>

## kvm.not_managed — not_managed / 000222020203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- kvm.not_managed

<a id="canonical-1220031302201113-0112001003002033-0111323113131021-0122330230130101-1121310000313202-1333130033213203-2111131211010322-2231022321321213"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223302111031213-3330013121301221-0212020220100231-2020130102233321-2302223321211031-1011321301323322-0331211121123211-3123330320220032"></a>

## Direct properties — not_managed / 000222020203 / 3

- [node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112): complete subsection reference.

<a id="canonical-2132333210202221-2213200022020000-1213212321222321-0001203012331001-3330221322131301-3102323222333301-0023110211302033-2032210122130132"></a>

## Next pages — not_managed / 000222020203 / 4

- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030311312133123-3333301122213333-1231321022123333-2323030221211212-2313001212112001-3102203022233311-0100111330202120-3201333001223202"></a>

## kvm.not_managed.node_list — node_list / 121211201311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- kvm.not_managed.node_list

<a id="canonical-2001222030202010-3013102220100331-2030012022032322-1032103311000332-1311200020021031-1130222311112122-1102022002301110-0021030320001133"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311313321313133-2022302203222312-0213003310302333-3011103120030003-1103101112020212-3213013332120012-1222230320310323-1130010033233302"></a>

## Direct properties — node_list / 121211201311 / 3

<a id="canonical-0320232333031230-1002302023133212-3001131212020110-2013112121323030-3003222131012323-1332100233320233-1101322103332221-2111233132211321"></a>

<a id="canonical-2023300132300022-0102303302001331-2031310320131320-1221032220012002-2323213233000201-3221213122322003-0021123320031323-0110332231233110"></a>

## hostname property — node_list / 121211201311 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203): complete subsection reference.

<a id="canonical-0123223130332203-0302210031001133-2232010300331030-0301303300103330-0300222201312331-3002023002322103-1001011310322123-1100200331001033"></a>

<a id="canonical-0133233212313212-2203031310023231-0212031201230110-3210311022221000-1020022113131301-2133220133320230-0300101312203103-2310333201332303"></a>

## public_ip property — node_list / 121211201311 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2231132103032013-1203201302301021-3122012021310302-2010213110131002-3200331000121030-0011122021231303-0021100021032201-3123331011022323"></a>

<a id="canonical-3231320301033121-0333310023102201-1233320212333321-1201312222100000-3022030003202222-2303100310110213-0333201022031311-2200130220213130"></a>

## type property — node_list / 121211201311 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3130312313223110-0002111121032333-0022323003203110-3021200220321220-1303200010022032-1013103033223030-1013311221033311-3322232020122000"></a>

## Next pages — node_list / 121211201311 / 7

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031331333013013-0311133120003132-3233200122030233-0302312021010203-0322313321033011-0213320322110113-0332303132211000-1330003010130200"></a>

## kvm.not_managed.node_list.interface_list — interface_list / 230000333200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- kvm.not_managed.node_list.interface_list

<a id="canonical-3231220113131302-1202000222223103-0320221211232311-3100212200010030-0112232120220031-3013210133001212-3311302130123011-3121121030030033"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023120213033103-0033313013133032-1121103031203310-0311123221031331-3010013002012222-2120101122002101-3220102223011133-1321010111030311"></a>

## Direct properties — interface_list / 230000333200 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302): complete subsection reference.

<a id="canonical-1013113010203211-3222013230322010-2023320003222111-0300112111321010-1302101321033131-2310220001102031-0302103030313333-3112330330310032"></a>

<a id="canonical-2212103123132033-3313311011112011-2223222100023003-1332020333323122-1102020300332200-2032213213021322-0012332023130131-2132013121132122"></a>

## description_spec property — interface_list / 230000333200 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-010.md#canonical-2120300022110211-0201330330303132-3312130313332333-3232322022112121-1121032202220011-0230013130201002-1102333120120132-0330321020121012): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3230303333012202-1021300033312101-1113102213332132-0123230310002231-3301112312020310-0003210212213203-3321200333100110-1110210321322322): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311): complete subsection reference.

<a id="canonical-2223310020032302-0101112133221002-1332003221200302-0021231221212023-1032002323123223-3000303110233031-1123222220212102-0131130020312211"></a>

<a id="canonical-2303231303233123-3310013322131230-2231210200333221-3011331223103111-0321321111003110-0112122133211003-3303302213230313-3002121012221223"></a>

## is_management property — interface_list / 230000333200 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2031133333322020-2201212133023222-0203113212121312-2302010223012101-2230321130203103-1102233102003303-3120330322301333-1123022001132201"></a>

<a id="canonical-3120113122020230-2032230030221313-0013001320000313-3123133222131110-3121200320002301-2221331331102113-1200212020222312-3333312131020103"></a>

## is_primary property — interface_list / 230000333200 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3331120301013222-3113001233123322-0331002112320330-0020112311030321-0022120311131032-3323221022313123-2221202322031033-0211200331231011"></a>

<a id="canonical-1323230010111133-3111210202012320-3230302022303223-2100001221131332-3211210022221021-1211311321213311-2222332333012323-3133300122201102"></a>

## labels property — interface_list / 230000333200 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-011.md#canonical-2012311020303313-1030231232300200-2221030123233323-1122220201031103-2332220103012331-0300232332201133-2300231200212201-2020333320021213): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-1100202122033020-2002013131322332-3123310302210202-3133301323002113-0313101011311302-3121103222332202-0200231131113331-2121223300012122): complete subsection reference.

<a id="canonical-1003103030031012-0121022032000311-3111010121030003-2003332320213011-1111231022122000-2332121031303211-3311301220130230-0300222330113212"></a>

<a id="canonical-0213101210203231-3323320121311202-1332120000230333-3210230330000203-2301223232033303-2002131320012203-2012020231203030-3120020223320000"></a>

## mtu property — interface_list / 230000333200 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-0133120023022023-3011322131222123-0001301110132312-0131013311223320-3301010200031302-2032311031222131-3003200133130211-1200012220031202"></a>

<a id="canonical-0230002111300331-2203133301233001-2313222223032212-2221201221330313-2330213021000333-2201302113012211-0333101200111332-2322223203211200"></a>

## name property — interface_list / 230000333200 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-011.md#canonical-3123131131203320-0311213232102301-1111221021223032-0003223303123123-3221233231103221-0223301101112102-2323310000121122-0302333222323201): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-0311010103223322-1303203103333303-3322330303011033-3111120302320230-1100013130210201-0011001111033233-1112101120100212-2202332313302101): complete subsection reference.

<a id="canonical-2233230112022301-2331123302321111-0021333020130020-1301220010211322-3033220220310122-3211320122320001-0003001232230232-0321002111000221"></a>

<a id="canonical-3000011132313312-1223212312210312-2221130102330302-3000232011213110-1322320330031313-1020233020020203-3121121003123132-0132022310310302"></a>

## priority property — interface_list / 230000333200 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-3332312113013203-3301201000000220-2221200231232212-2203012012230202-3321023213020032-0233201012101203-0132330002212211-3013111113220233): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-011.md#canonical-3013032133102011-2021301032001100-1323130231102113-1023120032112202-1012023001233332-3220030033113122-3220021230120310-3031131022131111): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-0131210010111310-2321312120022211-1112020000212131-3230212132003223-0200222021213212-0100033002013220-1222130322202101-1321121211213033): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-1203031212100311-0311310330301021-3222021132312202-2201122021003312-0323010310321021-1223212321312130-0322333300122112-1313302312011123): complete subsection reference.

<a id="canonical-2202213000002300-1023333022032113-2210132321112111-3203213331312331-1222010033301301-2033103321110231-0312100323300122-1112300103221321"></a>

## Next pages — interface_list / 230000333200 / 11

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- [kvm.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-010.md#canonical-2120300022110211-0201330330303132-3312130313332333-3232322022112121-1121032202220011-0230013130201002-1102333120120132-0330321020121012)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [kvm.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3230303333012202-1021300033312101-1113102213332132-0123230310002231-3301112312020310-0003210212213203-3321200333100110-1110210321322322)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-011.md#canonical-2012311020303313-1030231232300200-2221030123233323-1122220201031103-2332220103012331-0300232332201133-2300231200212201-2020333320021213)
- [kvm.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-1100202122033020-2002013131322332-3123310302210202-3133301323002113-0313101011311302-3121103222332202-0200231131113331-2121223300012122)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113)
- [kvm.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-011.md#canonical-3123131131203320-0311213232102301-1111221021223032-0003223303123123-3221233231103221-0223301101112102-2323310000121122-0302333222323201)
- [kvm.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-0311010103223322-1303203103333303-3322330303011033-3111120302320230-1100013130210201-0011001111033233-1112101120100212-2202332313302101)
- [kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-3332312113013203-3301201000000220-2221200231232212-2203012012230202-3321023213020032-0233201012101203-0132330002212211-3013111113220233)
- [kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-011.md#canonical-3013032133102011-2021301032001100-1323130231102113-1023120032112202-1012023001233332-3220030033113122-3220021230120310-3031131022131111)
- [kvm.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-0131210010111310-2321312120022211-1112020000212131-3230212132003223-0200222021213212-0100033002013220-1222130322202101-1321121211213033)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321)
- [kvm.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-1203031212100311-0311310330301021-3222021132312202-2201122021003312-0323010310321021-1223212321312130-0322333300122112-1313302312011123)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331300312013312-0201102330121221-1133030132032102-1033211101022320-1213200121022112-3232112001001133-0122002112210133-2100020102313210"></a>

## kvm.not_managed.node_list.interface_list.bond_interface — bond_interface / 011030011222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3201312210120212-2101330120000123-0331032013302120-3323221223132131-2011202120100023-1013103122032233-1311331013100011-2310301133220333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103033200110113-2113111202300030-3221101131130111-2322022220002100-1131113111013313-3131223132320202-0121231231220233-2001322311133012"></a>

## Direct properties — bond_interface / 011030011222 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-010.md#canonical-0332101012330232-1031102131202332-3133121301210002-3231023203210331-0200020121313032-0012213210320030-1003220131133320-1111330221001100): complete subsection reference.

<a id="canonical-3223120311203201-3001303200231121-0332001002213232-3102321333122322-2012000231202033-3103100121331201-3031011210020001-0212323232031201"></a>

<a id="canonical-2012123012023023-0220123223301122-1013020222121130-3012011320221001-0112222312222113-0032200011312203-2013000130233102-2121032112202020"></a>

## devices property — bond_interface / 011030011222 / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-010.md#canonical-2030322113302030-2222013010030213-2223312220000212-1022113023332023-2313233111032021-0323121312021332-1233221013003220-2013210131130100): complete subsection reference.

<a id="canonical-2011003312203302-1212303331310303-2221131002032331-3231023131333333-2313030020123130-0321313223131301-3010300000100221-3132020330131003"></a>

<a id="canonical-1033322333122323-1232023102303010-0100120232332002-1210231331122200-0121311002022223-3121111200123302-1133312212312120-1101111020222121"></a>

## link_polling_interval property — bond_interface / 011030011222 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1321220213123023-3302321010311101-2122033301233031-3321202203000300-0321120130013002-1211300031303131-2101022330033032-1011320323121032"></a>

<a id="canonical-3203310303132101-3103033133320131-0313100311332222-1131111200121133-2033331112030222-0020110121212102-2210020230132100-1221310203333211"></a>

## link_up_delay property — bond_interface / 011030011222 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3122313223222000-3130210002122233-3003213033022213-0203132112311001-3211133032012131-2112103003020321-2123003330022023-3111200230201002"></a>

<a id="canonical-3010310102232122-2230322100212012-0220020103232023-3233320130302132-3312110321033100-3000223312322311-3103021302023111-1122313003101010"></a>

## name property — bond_interface / 011030011222 / 7

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0223011220033131-3222230332113011-0032001233111000-2132012012113230-2000100003300031-2103210001121133-1222310220221002-0200300312010121"></a>

## Next pages — bond_interface / 011030011222 / 8

- [kvm.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-010.md#canonical-0332101012330232-1031102131202332-3133121301210002-3231023203210331-0200020121313032-0012213210320030-1003220131133320-1111330221001100)
- [kvm.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-010.md#canonical-2030322113302030-2222013010030213-2223312220000212-1022113023332023-2313233111032021-0323121312021332-1233221013003220-2013210131130100)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0332101012330232-1031102131202332-3133121301210002-3231023203210331-0200020121313032-0012213210320030-1003220131133320-1111330221001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102302233002321-0312211232032321-1332003013103300-1221303033302310-2331330332311102-0223302310130312-3300033010220231-1332130221021000"></a>

## kvm.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 313123120311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- kvm.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1030301102223111-2131000102132012-3112320011021101-1021202330332223-3220332123113302-2113131311303021-1113221323031301-0231221133131130"></a>

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

<a id="canonical-2122220121332203-2030123000021020-0103202330221130-1121220321233233-0111332333310313-1211123232022210-2331223311022310-1011120303130032"></a>

## Direct properties — active_backup / 313123120311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000222122100233-2120100322300110-2203122230002200-2303032202203212-1300200303001222-2102302230310100-0102322311000112-2112030233021032"></a>

## Next pages — active_backup / 313123120311 / 4

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2030322113302030-2222013010030213-2223312220000212-1022113023332023-2313233111032021-0323121312021332-1233221013003220-2013210131130100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200202122113321-1232001310203123-2330223311111233-2003110121220011-2032133210332111-1112331123032222-0330220312320313-3311223311003030"></a>

## kvm.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 001321020212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- kvm.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1123033111000220-3333212131033010-2012321323032222-2002221300022020-0213112313200133-1111013230231210-2013130330132223-2310021233203012"></a>

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

<a id="canonical-1010210132201233-0303013011031021-2031031321112221-0330232132112023-2203212101231120-3212012201233213-3022110031212331-0301233032032301"></a>

## Direct properties — lacp / 001321020212 / 3

<a id="canonical-2131011301221230-1133202131203301-1103100301211011-3232100313031021-3023331233013302-1212103103002102-2011002203332212-1213031003121000"></a>

<a id="canonical-1103233133213231-0012033123210111-0003021121123212-1323221111030311-0303203203201303-2023213203102121-3121121033311032-2310102133322000"></a>

## rate property — lacp / 001321020212 / 4

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-3012203131311303-1331222311303203-1213223012101332-0333011201333020-2122201210110302-2121301103000121-0221003333331220-3120313213103020"></a>

## Next pages — lacp / 001321020212 / 5

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2120300022110211-0201330330303132-3312130313332333-3232322022112121-1121032202220011-0230013130201002-1102333120120132-0330321020121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332113133011122-2030320213333113-3103110102021231-0330211331200100-0020101331302013-0033331231023102-1320231002012020-3103131003232201"></a>

## kvm.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 000023001000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0323330113201303-1033023113023003-2333031322000130-1211112221033022-1022331021322101-1103313113203111-0013311012223032-3303003303333133"></a>

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
dhcp_client = {}
```

<a id="canonical-3202022003100330-1232201322200121-1022321021232210-2222120110022122-3223002301323302-2200130030213013-3101310011130201-0332101032321212"></a>

## Direct properties — dhcp_client / 000023001000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233010211123332-2301133301332220-1300232303133310-0122133131200032-2310221333200211-1201332200101220-3301210023220311-2132122202220101"></a>

## Next pages — dhcp_client / 000023001000 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321321003022223-3203211020002331-2110200130101312-0302220313000330-2333303323222313-0012310232012302-3102302200202310-0002102233233131"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 223012211211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2100033213100331-0301233002121230-2301223311323231-2300023300212303-2312233212013003-1303320020201203-1331003110302111-3022311203210131"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123011202311233-0013021230303001-0010231213330031-2001312302232212-3112202333000103-2211223012332311-3000102021123233-1110202232001112"></a>

## Direct properties — dhcp_server / 223012211211 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-0233202102020033-1212330233223200-2020231002031320-3201230001200133-3130333300122331-2021002303222132-1301133322003002-0131321231011012): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-0233312221032103-0330200223321020-3012102221120133-0203220111022302-1222123221112023-1023301210020321-2230020010130022-3203110210300232): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303): complete subsection reference.

<a id="canonical-0321021311031303-0232303120012121-1331322102011013-2020230310011322-2102120200233123-3033101122132021-1323012203222123-1101231003033130"></a>

<a id="canonical-0322021313202222-1323331233021101-3112310022022303-3332213013030112-0323210013220033-1010133223030132-0033032102220103-3312000020200221"></a>

## dhcp_option82_tag property — dhcp_server / 223012211211 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3223202122121111-3313331220001012-3110013110023102-0031131102010032-2201033133133110-3202012221000303-2300200020013000-0010022221320320"></a>

<a id="canonical-1030222310120120-2301031031230003-1113032133233013-1312300301002002-0003130303301213-1223022320113203-3222032112232032-2322310213122220"></a>

## fixed_ip_map property — dhcp_server / 223012211211 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-011.md#canonical-1002201022001012-3122312230232322-1221131303323001-1130322310223220-2022221102320120-2301322300212300-3332131001022330-0321100200303303): complete subsection reference.

<a id="canonical-0221331000011031-0000021132211323-2202310200102233-0031222213213213-1320113002030133-2312112032112111-2112320021120132-2212212111032102"></a>

## Next pages — dhcp_server / 223012211211 / 6

- [kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-0233202102020033-1212330233223200-2020231002031320-3201230001200133-3130333300122331-2021002303222132-1301133322003002-0131321231011012)
- [kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-0233312221032103-0330200223321020-3012102221120133-0203220111022302-1222123221112023-1023301210020321-2230020010130022-3203110210300232)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303)
- [kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-011.md#canonical-1002201022001012-3122312230232322-1221131303323001-1130322310223220-2022221102320120-2301322300212300-3332131001022330-0321100200303303)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0233202102020033-1212330233223200-2020231002031320-3201230001200133-3130333300122331-2021002303222132-1301133322003002-0131321231011012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110223301033102-1300201133001102-0020332303223223-0213011033220220-1111201223002231-0022301220033031-2130030102321230-1022203201322321"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 221120103130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-2303302312013222-0232311030120233-2323120300103103-2331021200312332-2003130023231031-1211123131120230-1320102303201020-1032311023023103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-3123103031221033-0233122031233121-0231332132122003-0003012020003211-2003321302131322-3023001313210203-2233333323231210-1001021332220032"></a>

## Direct properties — automatic_from_end / 221120103130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003330233131133-3221312113032210-3333222333331301-1101112201120200-2211123030111212-1313120120100120-1132333001232201-3312221111000231"></a>

## Next pages — automatic_from_end / 221120103130 / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0233312221032103-0330200223321020-3012102221120133-0203220111022302-1222123221112023-1023301210020321-2230020010130022-3203110210300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031310212000231-3301100223233222-1100103122333010-1022233133113332-1211122112310202-3232220223103322-2312120301311222-1130002100203303"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 000231001020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-0103201302301301-2202212331213013-1032330113210332-1110232131123322-0031123132312123-0332023303311202-2022121023233103-3023100331222302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-0312023220232101-3033133202230033-0033020330012011-0021233301022101-0210213123032003-3222133333232022-1113322331210212-3032103331030001"></a>

## Direct properties — automatic_from_start / 000231001020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221111033332012-1312321133202331-2033213022301233-0023000031302333-1033310310030103-0000023201302330-0112023333111003-3300133013233323"></a>

## Next pages — automatic_from_start / 000231001020 / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123121303112222-0011310302100202-1212012020131203-1320022110201122-2303333001101133-0301113033132113-1332302332102023-3301013113003011"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 321331012312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2321102210201002-2322033221222111-0130213010330320-2302022332223133-1000321220222101-2220311033323132-3033101120312132-2302030000023131"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003200311030023-1113122002302222-1201011231021000-1333313233122030-0220201233103133-1133133302213023-1133323101333221-2013131100102300"></a>

## Direct properties — dhcp_networks / 321331012312 / 3

<a id="canonical-2231032301331111-0213112113120322-0200112123003230-3001210221223211-1322230130230010-2202222302122020-2130010123330202-1210022323313023"></a>

<a id="canonical-0210203300030030-3010332103111230-1230130012133012-1312211003121211-3030210102230011-3131130000101113-0031322103023310-0220210330311230"></a>

## dgw_address property — dhcp_networks / 321331012312 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0313022201302210-1131313303302321-0122322333130100-3022332010000302-3332323332202232-1323010110102331-1023330333011101-1221322220033110"></a>

<a id="canonical-2303011323010111-0313121002222031-3321021213332023-3123331203312022-3013332223320000-1110130031302301-0033111330323230-2130021302022221"></a>

## dns_address property — dhcp_networks / 321331012312 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-011.md#canonical-0130010302110212-0131103231112310-2322101003132013-2330210103200131-1111332231021010-0133011323303031-0302021032321011-0223211100220103): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-011.md#canonical-0002331300232323-0211130123223120-2222032213223230-1230103002113023-3011321101201103-1003031330000023-0103331013232000-3320111011123123): complete subsection reference.

<a id="canonical-1230323023030111-1122231201323301-2132000133312113-1203333201122112-0310030330200323-2202332222100033-3021111221100031-0110203333130121"></a>

<a id="canonical-1132112212211303-1022020130330220-2211312320330300-0233222031213222-0001322201131230-2332002330302203-1310212111302203-0032322323320122"></a>

## network_prefix property — dhcp_networks / 321331012312 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2130101101331120-0023223013033121-2132302111222303-2221030231323122-0003330313121320-2002133300322223-2010220000031103-0111120120003121"></a>

<a id="canonical-2221011132333113-0211231322210222-0300300232233213-0112113312001103-0233322300313001-1233101030332211-0022201100133322-1200332013113311"></a>

## pool_settings property — dhcp_networks / 321331012312 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-011.md#canonical-2103231133112020-1221332201203232-2031023033311002-3112320023131232-0231002302021110-1000311231121110-2303003311332213-3202311003332212): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-011.md#canonical-0313020210031222-2021210110201020-0120313322030010-0132100303232222-3103020303200013-0201100300201132-3000120303131031-1133113320332301): complete subsection reference.
