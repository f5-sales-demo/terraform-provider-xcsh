---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323003003223103-1000001112132122-1013210131311220-0133201002312021-3222230130230213-2113330001211013-2003331013013103-3323201300332220"></a>

## Property reference — Property reference / 102202330212 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- Property reference

<a id="canonical-2311203203310023-1332320100201322-2300113033101112-2133320321022223-3110131132022300-3300002121113310-2002133231003203-1302300221222223"></a>

## Direct properties — Property reference / 102202330212 / 3

<a id="canonical-1130103102301012-1301230332202313-2332001013303000-3323222321002220-2010212020101311-2110002001213332-2021221322320312-3300102220110213"></a>

<a id="canonical-3123330111131033-0200030301201011-0321320311101112-0010012023000133-1233021222311303-2300001313013332-2100031333223013-1022101030121112"></a>

## annotations property — Property reference / 102202330212 / 4

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

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203): complete subsection reference.

- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232): complete subsection reference.

<a id="canonical-3200033302201211-3012200230333030-1130123210300130-3202320013220113-0121322312013302-0023002111011110-3203021212032210-1333202202103322"></a>

<a id="canonical-1223223011121001-0022211101010203-0213100331102002-1102211112100010-2303131002321111-2113301120323323-0311111012032322-3121020112300321"></a>

## description property — Property reference / 102202330212 / 5

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

<a id="canonical-1323010020030233-1220102010010032-0030330011033230-0323113020031130-0122300320103322-1232230312001312-3202122022023022-1122130120322012"></a>

<a id="canonical-3033122203013133-2230213021030322-3133323113010010-3211323120120003-1113011201311303-3231000011232233-0032213213130230-0313011323021031"></a>

## disable property — Property reference / 102202330212 / 6

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

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221): complete subsection reference.

<a id="canonical-0021013133230030-2310020213031203-1230310223200233-1302101232310320-0101131123132101-1221312133211133-1020230321131030-1121011010100020"></a>

<a id="canonical-1310101133021133-1113113201232020-0211011211223301-3220130032322100-2012212330021210-3022020333102231-2031303010231003-3333011330013323"></a>

## ID property — Property reference / 102202330212 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1012021230230032-1022101113231133-1021310013313112-0122203302101002-0322021330111123-3011032200121110-3122100310300222-1331000001231120"></a>

<a id="canonical-1032020100312323-3220223220130020-1330020302311302-0302033022011003-3313222032330012-2012020220222201-3103303312330323-3210333122000203"></a>

## labels property — Property reference / 102202330212 / 8

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

- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-0130112110231113-3130032011200330-3231321010032002-1110233231103021-2000032102333133-0312112112133231-3001223102103322-3003100130233100): complete subsection reference.

<a id="canonical-0213322111123022-1010312233233221-1111222301231013-3312310102333221-3032332300232201-0111222333213213-2323230103212103-3332000302123111"></a>

<a id="canonical-0201222113301323-3203111133122301-1331332103122112-2103231212322000-0323212112332132-3101111212133333-1333232331000021-2130002112022320"></a>

## name property — Property reference / 102202330212 / 9

Type: `"string"`. Required.

Name of the Network Interface. Must be unique within the namespace.

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

<a id="canonical-1313012322001222-2313200320020013-3100100131221202-3200012312013312-1013033020131303-1032131301011112-1031021203210310-2310110332212111"></a>

<a id="canonical-3021201213302122-3332202103221310-0130322200021123-1032212011132330-1200310121322231-2202302133313232-1331200311031221-3100122202031113"></a>

## namespace property — Property reference / 102202330212 / 10

Type: `"string"`. Required.

Namespace where the Network Interface is created.

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

- [timeouts](resources--network_interface--reference--group-002.md#canonical-1303132311333210-1221330012020213-1122223010233002-0030303311010230-0011322330320102-3132110020230113-0221222223121302-3210013332102313): complete subsection reference.

- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0031231223100101-0221202300220330-0013311301301203-3301232310130302-2330031103323032-2221301010223301-3302220320030311-3202002232312200): complete subsection reference.

<a id="canonical-1123132101300321-3113321002112201-1011312213013131-1313332011012103-1003222231013331-1123233111212012-3223310113031121-0130222123112201"></a>

## All schema paths — Property reference / 102202330212 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_interface--reference--group-001.md#canonical-1130103102301012-1301230332202313-2332001013303000-3323222321002220-2010212020101311-2110002001213332-2021221322320312-3300102220110213) |
| `dedicated_interface` | [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2311032303210230-2330133102111313-3210111232213100-2232200012231120-1323033232111001-3121231123311220-0021230331113323-2202100231033300) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0003202121112032-3010023322200202-3321312200013310-2313020311113100-1132000210100032-1012131023220101-3202312210030001-3133023210012133) |
| `dedicated_interface.device` | [dedicated_interface.device](resources--network_interface--reference--group-001.md#canonical-0300333102223013-1331220132200320-2221110321020021-0201232313321102-0123102301330130-1001323211013013-2020112320230012-2213120201221323) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-2201000020013231-2331113333332231-1110021022211331-0302031012223031-1233202003330102-2211031303221311-1102310033320220-0311132330231213) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](resources--network_interface--reference--group-001.md#canonical-0021033300032030-2102032000221322-3211330220233233-3032202123222121-3132121201320213-1322210100010131-0230222333211131-3021022023103333) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-1021301221000123-3021200130012300-3301231231131121-2230103231333101-2321013121203210-3123223330120203-2331212211231003-0112303133101230) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](resources--network_interface--reference--group-001.md#canonical-2203101213231221-0131310311311313-1320230000023212-1010200100332013-3132332233123223-3230220233130303-0112322303123312-3220021233302300) |
| `dedicated_interface.node` | [dedicated_interface.node](resources--network_interface--reference--group-001.md#canonical-0332101001120130-1020213012032331-3123001011010002-2221301233310322-3012133303000030-1000313300032033-1030123333023010-3121111022102011) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-1021311023200232-3321211320311330-2110330111201322-3222333213303323-2231231211313023-3030210210011011-2133321112121203-1001131321033132) |
| `dedicated_interface.priority` | [dedicated_interface.priority](resources--network_interface--reference--group-001.md#canonical-3300120202123033-3000233023102222-3001032032230302-0020230312300330-0023300131101120-3331121003330333-3103120122300120-1322002233313030) |
| `dedicated_management_interface` | [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-0201213233101103-1110200133202221-3330012220023323-0320031113102213-3112012333201221-2210023121113331-2030121013020030-2030030203100020) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0003122232203111-0013221331330010-2301130030312300-3320231022333231-2122220032132302-1221112201103321-3200032020001011-1012323121103230) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](resources--network_interface--reference--group-001.md#canonical-3330011233013321-2123220121303332-0023211232303333-3210232232323233-1300130312223223-0111303303311332-0123300021010013-3023330202232333) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](resources--network_interface--reference--group-001.md#canonical-0113002330223230-3031111130212332-1000312220212133-2113002232312111-1333311220212001-2031122103322110-2300013001030120-0020113333111213) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](resources--network_interface--reference--group-001.md#canonical-0232313333031200-0123110213221103-0120131312312003-2330112330121130-1020320030210303-2223131331220311-2020232111303000-3220101210233231) |
| `description` | [description](resources--network_interface--reference--group-001.md#canonical-3200033302201211-3012200230333030-1130123210300130-3202320013220113-0121322312013302-0023002111011110-3203021212032210-1333202202103322) |
| `disable` | [disable](resources--network_interface--reference--group-001.md#canonical-1323010020030233-1220102010010032-0030330011033230-0323113020031130-0122300320103322-1232230312001312-3202122022023022-1122130120322012) |
| `ethernet_interface` | [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-1232001211130103-2011021231102322-3202111110022202-1023232032132322-1313123232120110-0231131213313002-3101200112031101-1233122010322021) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0220123220130333-1012102313200310-3330102320132212-2022321301032322-0210311010333112-2111012102103133-2330133020101211-1320230003230321) |
| `ethernet_interface.device` | [ethernet_interface.device](resources--network_interface--reference--group-001.md#canonical-2231202301012103-0030231203033012-3333310030313313-2323112033020330-2121211203303212-0123221130102033-0322201003103320-3131123302203013) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](resources--network_interface--reference--group-001.md#canonical-0302012113320211-2332113111223000-1212201120220231-3231001203001133-1020223112311100-3233021311313023-3023020110222332-3232010300003103) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1321011133201010-3213023122122102-3132231032020003-3000021231021313-2321032000233101-0030232122001313-3112002110020011-0120223232233320) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-2201032021002012-3200330211033033-3112120101302012-2132323002202111-3132002320320210-0130130111112232-2323001210323201-2012022011132202) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-0321120323122101-3032202223111330-1002201001210201-2013310112213313-0203132132012231-1113201021113311-2020100223122113-3330021013131213) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-3133300213321232-3010111202313011-3003030200203313-2311221032323331-0032202230011312-0120321302122113-2121330110202131-1031300013223320) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](resources--network_interface--reference--group-001.md#canonical-0030010231031022-3100200302020122-3312332020130310-3300312213133111-2231131223102221-2230230122213301-1021111211332113-2112210232123033) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](resources--network_interface--reference--group-001.md#canonical-2312332002303201-3030333202000333-0203123100301233-1132320110022233-0223023013103330-0110133112020322-0101012202223302-0202200132303030) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--network_interface--reference--group-001.md#canonical-3311013102102330-1310033122331201-3120330333200032-0201112331031033-2002222123201110-3001030331302131-1110023013033020-2330233231103121) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--network_interface--reference--group-001.md#canonical-2132331222111013-0311020021300330-2011222213033133-3320013300031302-1030031113201313-0231220322320011-0001100020020030-0122211101232323) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](resources--network_interface--reference--group-001.md#canonical-3222112103312003-3032200001330003-1121121130311321-3122000131200003-1003001331002120-2012003022220213-3120230122001312-3310233233320033) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](resources--network_interface--reference--group-001.md#canonical-1033320213023000-3023330012000312-0213123302000110-1120121033330233-1200121230222323-3233322001001032-1232132213122230-1013103033203322) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-2223202202331203-3330223231103122-1130331113212002-0022221122310333-1131013310321002-0100031302232313-1212133030030201-3103001210003011) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](resources--network_interface--reference--group-001.md#canonical-0302321113203030-0210230031312113-3132111120012233-1103031021323110-2123310330022020-3101131202333203-0212121221202110-3332303303233010) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](resources--network_interface--reference--group-001.md#canonical-2022332220232330-0020001123112101-2201033122011013-0212003321110202-1132130221112223-2223212312200120-1130120201310113-3231120211121120) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](resources--network_interface--reference--group-001.md#canonical-1012230212031230-3303030113131033-3011331032003322-2113220313002011-3012121302220003-0311003030011200-0122232302121323-3210000112322123) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--network_interface--reference--group-001.md#canonical-2222310032012312-3030103003133122-0103030231333021-3312032300333300-1103212221022331-0233121003211302-3303212322101203-1312102310213120) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](resources--network_interface--reference--group-001.md#canonical-0032111201223221-2101312003023111-2011310233022112-2032132001202030-3332130333203200-0012311121213211-0223022102301222-3022202003230002) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](resources--network_interface--reference--group-001.md#canonical-1230002001213013-3112212013133203-1211123303120002-3313321212100111-3331110102121033-0103220222212003-0102023223012331-1021210320010002) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-2222101101101123-1101002022233301-0132313311110220-3201100222031202-0121132031231001-3223213021120103-3313300300021132-1023110322122023) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-3331002322130133-3113321120233111-2321122322322000-3323003000103130-3101011022330112-3120313133311310-3133111002012201-1123200131100113) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-0003200332000001-2222121003003222-0011010201323213-1001233230310310-3231031012011123-3330110202333131-2312213323300000-0031203302233330) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](resources--network_interface--reference--group-001.md#canonical-0231120323200230-3012101212131132-2123120213132010-2003212212000023-1233013201301110-0011210322022132-2100203321323121-0320200332311212) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-1120213113131012-0303211103323100-1332321311203032-1122010002103233-1121120020231103-3030213020102331-2103020021202312-0101002210313030) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-3130110331133221-0212302031333103-1100012112233101-3222121030133220-0311322231300213-2302021012120331-3231321022212311-1212011222221321) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--network_interface--reference--group-001.md#canonical-3012211213002333-3310231101202331-0211032123001113-1321213311202223-3310323320230230-2033100331232203-0012121312300023-0203223131301210) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](resources--network_interface--reference--group-001.md#canonical-0033300023120332-3202030212120322-2301033220130112-2302013213201132-1213022010332302-1001120313203312-2212123220323203-0122132120231013) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-2133212032322311-3201131202120223-0230011022001130-2013013011300032-1020011211210021-1323210233230233-3121032101113323-1301301030010312) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](resources--network_interface--reference--group-001.md#canonical-0201032022210101-1211220031123321-2022000133323233-1211112113033211-0303210010213112-2220311122331222-2102211023320300-3020130323111323) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--network_interface--reference--group-001.md#canonical-0002201233031130-2103033202331230-0023322000132230-0221131110101332-0312021002232212-3203312100011231-2300323002022221-3333010223030022) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--network_interface--reference--group-001.md#canonical-2110233021223032-2123023311322013-1312311300010320-3103233222131103-3221301312320030-3320310300323120-2333211212002312-1211003131323001) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](resources--network_interface--reference--group-001.md#canonical-0112210312131011-2230222012100303-2010233001311011-1321222110210333-2211133022021010-1031003301210320-3301303212333122-2023130323200100) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-0113300320232133-0200303131001233-0011303331001330-3221132011110100-0230122321003230-3211233103301320-2000003320333230-1123110310221122) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-1230221233010222-3132312222120202-2100100011133110-0133303301200323-3231032302020033-1311330100023202-1320011012200122-0012022301111103) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-2023331031332113-2320320210333121-0211333111203111-1221310200222111-2111321111120021-2321032032200320-0130121333122322-1001112321321100) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-2023003303132312-0030112320201211-3132230331011122-3321233220131103-2222102013220313-1031132103031322-1022221311322031-2311311200133332) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](resources--network_interface--reference--group-001.md#canonical-0202011311333310-2121231020100001-3231020131023120-3020103123011030-3202011012031130-1012030112011023-2212120202110000-0332011130013121) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](resources--network_interface--reference--group-001.md#canonical-0022131013323311-1010302310032300-1321110300010012-3120200012313322-1023012030212123-2132202310202033-0103333232300003-2323123333313021) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-0013020213320311-0222020330030300-0320302003131030-0323222212022211-3223111001322200-0000303320003200-0220230003321200-0020211122132033) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](resources--network_interface--reference--group-001.md#canonical-1233000101101012-3130321110102231-1230133132301211-1231003013123230-2211100132201113-0301133302013121-0331213010312312-3012310002010213) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](resources--network_interface--reference--group-001.md#canonical-1301020033322311-1132022121001201-2121133010003102-2201120031003311-2012112000100121-1231303032013232-1023200202210231-3212210232103123) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](resources--network_interface--reference--group-001.md#canonical-0031103130122233-1030012002020030-1312300231121033-2323031120110023-3130100320132312-3310002202330222-2000000102000221-1303003032130330) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-2031011232113112-3302330330322032-2232133020302222-0200002110203012-3333222012202010-3312301103200000-3033001322201320-2030311313031001) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-0332213212132233-1210301013132220-1331210031220212-3300212330020301-2200112121233202-1021232302010120-0003132211211211-3221223023230012) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-0222021130202303-0012130030002002-2312303031202011-3332232313113021-0213321223213131-1123101111121222-3011013120111331-2111223130232222) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](resources--network_interface--reference--group-001.md#canonical-0230023222211001-1001311133113012-0301131201202113-2310201320103332-3322121211130303-3012031321310000-2213221032112312-1012331221302133) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-0000230123130100-1020033121000221-0212002311321130-1211301011321031-1230313023010213-3321310120210202-3301021330121321-1202020133321323) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](resources--network_interface--reference--group-001.md#canonical-0103333211203321-3222311312203233-1330130203322013-0101021033110013-0303201203131102-2302200133200021-0133023222133232-1132312130031202) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-2001322023313200-3011100222320010-1220210030303113-2312302221132330-2133103012133230-1033000003002110-2323333202332011-1201003111200123) |
| `ethernet_interface.node` | [ethernet_interface.node](resources--network_interface--reference--group-001.md#canonical-0030020301233212-3023311302300223-1330212002102102-1010033222001332-1110322213132231-1011101313322333-1031120200131121-3201102200223323) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](resources--network_interface--reference--group-002.md#canonical-3112333223012231-1022003312312220-0033011131101320-0030032022210330-0311103120233113-1002311311301223-2023132331313032-1313023230333211) |
| `ethernet_interface.priority` | [ethernet_interface.priority](resources--network_interface--reference--group-001.md#canonical-2310220202133120-3101221230123030-0230311003103210-2203330211033220-0122230101112221-0300232203032223-1013220012021201-2002302131322023) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-2011331231311211-2121232201111231-3211301013222202-2222103230202010-1030120312132323-1110031101201030-1212132020131221-0100013302231212) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](resources--network_interface--reference--group-002.md#canonical-2003210120221131-1130300003032303-3110323103022031-2130011101130202-2010222302031032-1210111200010232-1012133110201101-3002311120220310) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-2323210322110110-0321001132311030-0333232221212332-2320322122031303-0132312120112212-3200332022202210-2033000121310203-1232021003022100) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-1222210321331322-1322113121332313-2103103002113103-2122110022221012-2001212223201121-0323020011223000-0231121223012003-2310231201302111) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-1113031331221323-1200323010300302-1330200002123132-1302133321231321-3120311013222102-1031211132223302-2110012310310212-0113212323313003) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-3113003132222213-0000112302223122-3010113131102300-1221121033211230-3202000112031223-1023002132303001-3331202303222001-0303210333132031) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-0231232233201013-1000023230103033-3002133232331313-3010233030011001-1233102111311310-3110022030322202-1133023130033232-0211330023033222) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-1033202003332123-3300033110030332-1131100130212031-2233011132322222-0321322310320111-2312101302111202-0122001020221021-1302203310231122) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-1323313032001322-1013100031021220-1301233230303030-3111312212011232-2120322010223110-3200321023010313-0001010032033031-1322120322111220) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-0300013323021220-3102320302120100-2320111322222030-0111010230311130-1223030223012202-2212213331301111-0331313231321212-2111230213300313) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-3101110300011233-3031310120213302-0112013032301300-2023320021022212-0001310331100231-0203121031323230-0011130111103312-3010113031133332) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-1300120130101020-3133221223121211-1311100010113010-3102232131030321-1113202300233023-2321001212010301-2321111010320230-3031200223001212) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](resources--network_interface--reference--group-002.md#canonical-3102001000022101-2132010031320031-2111003311322220-3320212133113112-1302210132020032-2320122020311210-2110133312113120-2323321130132302) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-0111023321013131-3013322232212002-0233311011122102-3332220023332002-2102122130023022-1102302223200030-3102333011323302-3102300123320122) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-0300311310230022-2200210333132122-3332303031331320-0003231200202001-0020031202331130-0303103330333133-3200111322003210-1032302232130110) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-1321130100030022-2122301102003212-3232313230300121-0003102222133022-3222300201133323-0310220103010320-2100013132312212-0231301311200033) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](resources--network_interface--reference--group-002.md#canonical-1333330013321223-3221000101133203-3220232301222231-1111001001302101-2002202312212223-1323000311223131-2223322231221020-0120022231010322) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](resources--network_interface--reference--group-002.md#canonical-2220132211130001-2323011033303331-2211110131210023-3231331212211232-1233221023012223-0022222013231103-0023110312202130-2322001323000300) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](resources--network_interface--reference--group-001.md#canonical-3030301320210110-0222330320212311-0121111212033303-0303200122003023-3031321033200230-0011130222212220-1310203122212023-3211133221322122) |
| `id` | [ID](resources--network_interface--reference--group-001.md#canonical-0021013133230030-2310020213031203-1230310223200233-1302101232310320-0101131123132101-1221312133211133-1020230321131030-1121011010100020) |
| `labels` | [labels](resources--network_interface--reference--group-001.md#canonical-1012021230230032-1022101113231133-1021310013313112-0122203302101002-0322021330111123-3011032200121110-3122100310300222-1331000001231120) |
| `layer2_interface` | [layer2_interface](resources--network_interface--reference--group-002.md#canonical-3220033332033110-2100231323012321-1321133303313311-3132103333020223-3020313223122032-2210020232312003-1320302312223012-3321332301202321) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-0222322132221230-2013221322201123-2330212100301123-1202230100122100-2111123133000121-1110231333222231-2033121333100113-1011320210222203) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](resources--network_interface--reference--group-002.md#canonical-1000103001321223-2213123200200320-1330213113133133-2103100232021031-2102331201130231-2122000030301300-2323100330333301-0122102000323300) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](resources--network_interface--reference--group-002.md#canonical-0223120210002003-1030022313233220-2233100103022031-3300020311010000-3133123303222103-0322222133131023-2113230032103110-2030100233101231) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-3100003011132031-0321321313203033-3112000320133211-2331330333102112-3223200302020130-3000101123322321-3310102311013201-0212030023232230) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](resources--network_interface--reference--group-002.md#canonical-1223122023132011-2003231022230123-2333121113321320-0120100110112230-3021223310033302-0301121221323011-1222021320213112-3111213130001222) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](resources--network_interface--reference--group-002.md#canonical-2233030003300033-1022323230301200-2003203030123130-0022111033201111-2002002022003323-1323022132030312-2310323212302230-2033323001313200) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-0023332122231233-3123023022210023-1113013001222023-0020332233232323-2322122212111320-2222110213312231-3013133000103200-0133012030302011) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](resources--network_interface--reference--group-002.md#canonical-0132120003302212-2311001212031123-1300102321013323-0102130313220222-0123112303130332-3113012331132033-2012330320313022-3102031031320032) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-1213032013101213-1100023103011012-1313322310201112-1101101013300121-0100321303031222-3103212310333111-3332020002233211-3022300220130110) |
| `name` | [name](resources--network_interface--reference--group-001.md#canonical-0213322111123022-1010312233233221-1111222301231013-3312310102333221-3032332300232201-0111222333213213-2323230103212103-3332000302123111) |
| `namespace` | [namespace](resources--network_interface--reference--group-001.md#canonical-1313012322001222-2313200320020013-3100100131221202-3200012312013312-1013033020131303-1032131301011112-1031021203210310-2310110332212111) |
| `timeouts` | [timeouts](resources--network_interface--reference--group-002.md#canonical-0033012210131322-3223013013133223-3302233210332100-2321212333223033-0123111001130122-0003310300102300-3012212210331230-0203320101103213) |
| `timeouts.create` | [timeouts.create](resources--network_interface--reference--group-002.md#canonical-3233233213120012-1111122130201320-3121023132100330-3103312203002001-0322230303211111-2003032311033131-0113212010223212-2201102001312022) |
| `timeouts.delete` | [timeouts.delete](resources--network_interface--reference--group-002.md#canonical-1220331200230222-2002311311003213-0302312102031232-0203200321122100-2322132311033130-1210221000012223-2113331001200020-3012330233203202) |
| `timeouts.read` | [timeouts.read](resources--network_interface--reference--group-002.md#canonical-0120020030011323-0331301112012320-2002230221121311-1232030021332023-1030201030320312-0312322311113103-3122030002313300-0330301111313022) |
| `timeouts.update` | [timeouts.update](resources--network_interface--reference--group-002.md#canonical-3001222021123212-2332103133320132-3023313330233331-2123030231221321-0011231030320133-3232300001132233-2102013220203102-0233020212221213) |
| `tunnel_interface` | [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-1123011233310210-3220002312121121-1312211122213111-2012111232333333-0203322012120121-0223022000333232-2011012300102201-1113320001331012) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](resources--network_interface--reference--group-002.md#canonical-0000101212330200-3202030023110110-3323113213121112-0302220200011110-1322320012030132-3230311221003032-0021003311202210-1032113332302103) |
| `tunnel_interface.node` | [tunnel_interface.node](resources--network_interface--reference--group-002.md#canonical-2321133201003031-2312331303220121-2232001131310031-3121232222303001-2221321230133333-3100311031031102-0003130122212310-3122003110122121) |
| `tunnel_interface.priority` | [tunnel_interface.priority](resources--network_interface--reference--group-002.md#canonical-2000001002020131-3322303103021230-1222031130321223-2330100310120312-2330332202333232-3213020211203032-3133212213002230-3103111113202020) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-3003222031010203-2000232033210101-3101301111102311-3213310011011220-3211021122212222-2102303111003020-3031012122232133-2033323210011102) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](resources--network_interface--reference--group-002.md#canonical-1321111200212133-0322001111210223-1022222001213231-1021003333213320-3011133131123001-1012111010223312-2103222123313220-3332331020003223) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-3012123133100113-1300122302223311-1221132032022133-0021300001330311-3332010133031103-2222303230222011-3120332120120110-2321220203230110) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-3103023112313303-1330310000000202-2213101302200113-1310001112300330-0210102031302011-3103313101122121-0312332132313030-2331003130002220) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-3021112021230211-2011132303113103-1132300102111101-2113300222131230-2232221113031211-2322211022021220-3220331230030101-2121332323103210) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-3013331022112231-2323212002202010-3130033020132332-2321321200021333-2023202301230232-0233303311333101-0231300033023230-3313103103220122) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-1121033311100333-2110301111120302-3323210303221231-1333231220012123-0302222102100010-3230211010302110-1022323132100012-1123233111020313) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-2111132203200303-0013330110230223-2120031331021202-2132332131002123-2122220232021013-2310213122320220-3332102110331110-2111020231333203) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-1323032322302331-1321121231032011-3322211203011033-1121112001010120-2013220322233223-3220323233100031-0132223212311112-2311213310033302) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](resources--network_interface--reference--group-002.md#canonical-3212213203211123-2000113103201332-0001201121300110-1233233100320031-0200021100332330-0130022202223103-1023222331130212-0002330320010200) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](resources--network_interface--reference--group-002.md#canonical-3212013302223121-3223222331233320-0333211023011121-1201100122200301-2321321130323023-3310231130210320-2200012333232101-2322222231301303) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](resources--network_interface--reference--group-002.md#canonical-1200323130131130-3112002002101213-2032300101230022-3132210323111102-3002013002002312-2202333120300300-3013213331100321-2122101313122302) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](resources--network_interface--reference--group-002.md#canonical-1321330323202020-0023302321230213-0330223031000031-0323333302100233-0122012310131130-2021222121122333-1021102201231313-2321012102111101) |

<a id="canonical-1213103312302232-1232333230332202-2023110102123112-2200231021001231-1213233101031103-3011320333010333-2221003133323333-0011130013112332"></a>

## Next pages — Property reference / 102202330212 / 12

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-0130112110231113-3130032011200330-3231321010032002-1110233231103021-2000032102333133-0312112112133231-3001223102103322-3003100130233100)
- [timeouts](resources--network_interface--reference--group-002.md#canonical-1303132311333210-1221330012020213-1122223010233002-0030303311010230-0011322330320102-3132110020230113-0221222223121302-3210013332102313)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0031231223100101-0221202300220330-0013311301301203-3301232310130302-2330031103323032-2221301010223301-3302220320030311-3202002232312200)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332131323133001-0023201003100132-1201001323001220-1311313002020021-3332102322021113-1303002113110121-2330303333200213-1021133130101001"></a>

## dedicated_interface — dedicated_interface / 230021123331 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- dedicated_interface

<a id="canonical-2311032303210230-2330133102111313-3210111232213100-2232200012231120-1323033232111001-3121231123311220-0021230331113323-2202100231033300"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled")}
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
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]"
}
```

OneOf alternatives in this subsection:

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2311032303210230-2330133102111313-3210111232213100-2232200012231120-1323033232111001-3121231123311220-0021230331113323-2202100231033300)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-0201213233101103-1110200133202221-3330012220023323-0320031113102213-3112012333201221-2210023121113331-2030121013020030-2030030203100020)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-1232001211130103-2011021231102322-3202111110022202-1023232032132322-1313123232120110-0231131213313002-3101200112031101-1233122010322021)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-3220033332033110-2100231323012321-1321133303313311-3132103333020223-3020313223122032-2210020232312003-1320302312223012-3321332301202321)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-1123011233310210-3220002312121121-1312211122213111-2012111232333333-0203322012120121-0223022000333232-2011012300102201-1113320001331012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dedicated_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232122201300013-1223220221021333-3121320210233211-2311200221320120-3113033101033100-1311113230310003-1030032133300103-0330133002021010"></a>

## Direct properties — dedicated_interface / 230021123331 / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-0213100202031002-3003330112113330-2022302121333132-3233311103010020-2020133223123011-2303121203030301-3210330331100101-0203022231111013): complete subsection reference.

<a id="canonical-0300333102223013-1331220132200320-2221110321020021-0201232313321102-0123102301330130-1001323211013013-2020112320230012-2213120201221323"></a>

<a id="canonical-1013132303122302-2333331302130210-0022303000022102-2201023111122212-2231103231301102-0003322132322221-1230202223320133-3220333230110212"></a>

## device property — dedicated_interface / 230021123331 / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

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

- [is_primary](resources--network_interface--reference--group-001.md#canonical-1120221121310330-2213103300200020-1013221212311133-2101131002332002-1123110000310113-1301032200222221-1130313112302310-1001101112232133): complete subsection reference.

- [monitor](resources--network_interface--reference--group-001.md#canonical-3230321221323212-3020031231321210-2321212001020302-3332321332202330-0202313030101032-3113030003102211-1301123123303333-1031203002101112): complete subsection reference.

- [monitor_disabled](resources--network_interface--reference--group-001.md#canonical-3330002131233003-1000331012011032-2002302130301223-3301120302100133-2301003002333133-0323222010313302-3011323213210133-2301112322030313): complete subsection reference.

<a id="canonical-2203101213231221-0131310311311313-1320230000023212-1010200100332013-3132332233123223-3230220233130303-0112322303123312-3220021233302300"></a>

<a id="canonical-3123132121022002-3330131013323332-1021202333113323-3133210313331030-3231023321110001-2110301002300131-0030100111103011-1301113113121003"></a>

## mtu property — dedicated_interface / 230021123331 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
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
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-0332101001120130-1020213012032331-3123001011010002-2221301233310322-3012133303000030-1000313300032033-1030123333023010-3121111022102011"></a>

<a id="canonical-3113210333131011-0132132121201032-0121322110213203-0321322021131303-2130312121210100-1000213312011230-1313031001130123-3133011003212110"></a>

## node property — dedicated_interface / 230021123331 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--network_interface--reference--group-001.md#canonical-1330232323323033-2213231220232120-1023303001313011-3022023100012313-3333211230300113-3301023101220201-2310300220101032-1112200313200311): complete subsection reference.

<a id="canonical-3300120202123033-3000233023102222-3001032032230302-0020230312300330-0023300131101120-3331121003330333-3103120122300120-1322002233313030"></a>

<a id="canonical-2200222101011213-0023203012223220-2231010120303331-2022230022311123-2023030112032210-3301330320030022-1322010301312223-0120102202222310"></a>

## priority property — dedicated_interface / 230021123331 / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

<a id="canonical-1310231003033333-2110312330320110-3331213130313000-1232131130330012-2220122200310231-2001022322121203-1131103203310113-1003133112030221"></a>

## Next pages — dedicated_interface / 230021123331 / 8

- [dedicated_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0213100202031002-3003330112113330-2022302121333132-3233311103010020-2020133223123011-2303121203030301-3210330331100101-0203022231111013)
- [dedicated_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-1120221121310330-2213103300200020-1013221212311133-2101131002332002-1123110000310113-1301032200222221-1130313112302310-1001101112232133)
- [dedicated_interface.monitor](resources--network_interface--reference--group-001.md#canonical-3230321221323212-3020031231321210-2321212001020302-3332321332202330-0202313030101032-3113030003102211-1301123123303333-1031203002101112)
- [dedicated_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-3330002131233003-1000331012011032-2002302130301223-3301120302100133-2301003002333133-0323222010313302-3011323213210133-2301112322030313)
- [dedicated_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-1330232323323033-2213231220232120-1023303001313011-3022023100012313-3333211230300113-3301023101220201-2310300220101032-1112200313200311)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0213100202031002-3003330112113330-2022302121333132-3233311103010020-2020133223123011-2303121203030301-3210330331100101-0203022231111013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312303332011213-0233222002220112-0101211322002110-3112201031321313-3301213012120223-1220001330300122-1121311233300303-3202303301230033"></a>

## dedicated_interface.cluster — cluster / 012120133120 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- dedicated_interface.cluster

<a id="canonical-0003202121112032-3010023322200202-3321312200013310-2313020311113100-1132000210100032-1012131023220101-3202312210030001-3133023210012133"></a>

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
cluster = {}
```

<a id="canonical-3120333213032202-1313020031232301-1230113031100331-0331302023212113-2323022110003030-1131113332201311-3133103310222002-3001202232023132"></a>

## Direct properties — cluster / 012120133120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223213231123031-0003200023020032-2002230233033221-2001121131302213-0013133303110201-1320221012031313-2033130131312012-1212301020130133"></a>

## Next pages — cluster / 012120133120 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1120221121310330-2213103300200020-1013221212311133-2101131002332002-1123110000310113-1301032200222221-1130313112302310-1001101112232133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122230111111222-2213123002233203-1223012132001313-1331011201200030-3123002203302220-2013103113110323-1320323333310231-3013000331302232"></a>

## dedicated_interface.is_primary — is_primary / 302023113030 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- dedicated_interface.is_primary

<a id="canonical-2201000020013231-2331113333332231-1110021022211331-0302031012223031-1233202003330102-2211031303221311-1102310033320220-0311132330231213"></a>

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
is_primary = {}
```

<a id="canonical-2232233222021301-3131212021010030-0202331111223220-1011001110333002-1332132022120111-3203210213113123-1221202021331331-1130330013003123"></a>

## Direct properties — is_primary / 302023113030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103122231320011-3120202213003310-0220023003333203-1023002133300002-3203111112330312-1011313313023332-1131003101111031-0200203030232330"></a>

## Next pages — is_primary / 302023113030 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3230321221323212-3020031231321210-2321212001020302-3332321332202330-0202313030101032-3113030003102211-1301123123303333-1031203002101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111000230033000-3002031102220003-3201123122000131-2232131013331020-3101131002302322-2323100301021022-1320020322103220-3212102230231110"></a>

## dedicated_interface.monitor — monitor / 010202212112 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- dedicated_interface.monitor

<a id="canonical-0021033300032030-2102032000221322-3211330220233233-3032202123222121-3132121201320213-1322210100010131-0230222333211131-3021022023103333"></a>

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

<a id="canonical-3321210312012203-2312100023323132-1032330032110011-3021031022023232-0121331110223332-0223211332033032-3303120112101233-3133001121033111"></a>

## Direct properties — monitor / 010202212112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210130103022030-3330102123202101-2330120313310222-1022322321102132-2120121312311103-3323223021302210-3012111121120100-3233203211013013"></a>

## Next pages — monitor / 010202212112 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3330002131233003-1000331012011032-2002302130301223-3301120302100133-2301003002333133-0323222010313302-3011323213210133-2301112322030313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223223320212121-1011020323313012-0201213123213312-3232232310201202-2311120033303221-2211203333010222-1000010120310001-0022121133210233"></a>

## dedicated_interface.monitor_disabled — monitor_disabled / 100011202320 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- dedicated_interface.monitor_disabled

<a id="canonical-1021301221000123-3021200130012300-3301231231131121-2230103231333101-2321013121203210-3123223330120203-2331212211231003-0112303133101230"></a>

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

<a id="canonical-0313201102111011-1323001000300022-3313212311011022-2110012112102120-0221103210313230-3203302020113130-2120333223021212-2203121030113103"></a>

## Direct properties — monitor_disabled / 100011202320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102313321330120-2123321130131033-0101012103230203-0212222123311310-2133000123312231-0113333223032113-3103222032222302-1122323330212311"></a>

## Next pages — monitor_disabled / 100011202320 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1330232323323033-2213231220232120-1023303001313011-3022023100012313-3333211230300113-3301023101220201-2310300220101032-1112200313200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111022301202111-2022002030223121-2220020323021222-3321230112023100-3333101220112213-0312012001030031-2302301022210331-1331103233323121"></a>

## dedicated_interface.not_primary — not_primary / 232203303020 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- dedicated_interface.not_primary

<a id="canonical-1021311023200232-3321211320311330-2110330111201322-3222333213303323-2231231211313023-3030210210011011-2133321112121203-1001131321033132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for not primary.

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
not_primary = {}
```

<a id="canonical-3203213202100020-1320013112103021-2112320021200201-0323123032202322-1121110110200011-2021100212230223-1021122120033133-2003030200032113"></a>

## Direct properties — not_primary / 232203303020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032100300133312-1131210123110212-2333331220013131-0110333121032111-2031003013021031-0320203100032110-3100110133003201-2222211213132231"></a>

## Next pages — not_primary / 232203303020 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-2202312132221001-2113331212010103-2130332010130032-2110133310110321-3231313332220130-1300110303021031-1333223321003031-2113330102231203)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132022003111303-1313331030222123-2233131000010330-0010232103030211-1330212012032132-3233331202300112-2123113003301113-0032330303333013"></a>

## dedicated_management_interface — dedicated_management_interface / 123213131201 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- dedicated_management_interface

<a id="canonical-0201213233101103-1110200133202221-3330012220023323-0320031113102213-3112012333201221-2210023121113331-2030121013020030-2030030203100020"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node")}
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
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

Terraform syntax:

```terraform
dedicated_management_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013332122311223-1000212200132121-3111030332302233-1333201230100012-2210023332000232-3032303213123311-3223231003010330-3121131201001122"></a>

## Direct properties — dedicated_management_interface / 123213131201 / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-3212120332223331-1122103323313100-1000321313130030-3303102022002020-0102311301223332-3031302220001013-2320131303320200-2203011232003120): complete subsection reference.

<a id="canonical-3330011233013321-2123220121303332-0023211232303333-3210232232323233-1300130312223223-0111303303311332-0123300021010013-3023330202232333"></a>

<a id="canonical-2230201222030230-1103303120133221-3222201103013111-3111013010132330-1120000231103211-3103202233232132-2310022002231031-2333200003300320"></a>

## device property — dedicated_management_interface / 123213131201 / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured.

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

<a id="canonical-0113002330223230-3031111130212332-1000312220212133-2113002232312111-1333311220212001-2031122103322110-2300013001030120-0020113333111213"></a>

<a id="canonical-0230220300131133-0010311200333123-1033132113000022-3011213101223221-3323301300022030-3330121313020210-1111202030203130-0323023301102010"></a>

## mtu property — dedicated_management_interface / 123213131201 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
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
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-0232313333031200-0123110213221103-0120131312312003-2330112330121130-1020320030210303-2223131331220311-2020232111303000-3220101210233231"></a>

<a id="canonical-1133113130332231-0202311201013112-0112033103012003-3300111203330322-1020033120130111-0223330122001100-0012323210022301-0231202202230022"></a>

## node property — dedicated_management_interface / 123213131201 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3032111133203313-2213221111233333-0132033020131331-0032221231101033-3130101122332003-1223313221112032-1023230020100110-2313003101320021"></a>

## Next pages — dedicated_management_interface / 123213131201 / 7

- [dedicated_management_interface.cluster](resources--network_interface--reference--group-001.md#canonical-3212120332223331-1122103323313100-1000321313130030-3303102022002020-0102311301223332-3031302220001013-2320131303320200-2203011232003120)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3212120332223331-1122103323313100-1000321313130030-3303102022002020-0102311301223332-3031302220001013-2320131303320200-2203011232003120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333133221121032-1322000301101131-2203021121230322-2230023032110220-2332020233120221-2231113201111301-3111332302100202-3021323221213202"></a>

## dedicated_management_interface.cluster — cluster / 132123013102 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232)
- dedicated_management_interface.cluster

<a id="canonical-0003122232203111-0013221331330010-2301130030312300-3320231022333231-2122220032132302-1221112201103321-3200032020001011-1012323121103230"></a>

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
cluster = {}
```

<a id="canonical-0303022122210022-0032123311021012-0332002111333033-1100200013023103-2323222010223000-2000332031200232-3222122031302222-0230322331210312"></a>

## Direct properties — cluster / 132123013102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101122112002211-3012323230332303-2320312200103301-3333021100333200-0022000033021011-0301020312011112-0202232032200102-2023111122111211"></a>

## Next pages — cluster / 132123013102 / 4

- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313210212222212-2012113331002222-0032111232211311-2223330212121310-0210221002120123-1210301230302033-0110110012231331-3032103330010103"></a>

## ethernet_interface — ethernet_interface / 213332312330 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- ethernet_interface

<a id="canonical-1232001211130103-2011021231102322-3202111110022202-1023232032132322-1313123232120110-0231131213313002-3101200112031101-1233122010322021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
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
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110300311331232-0130102313332233-3121312112113131-3233202103120133-3320132332012310-2213311103203020-2313123320202032-1113230110312210"></a>

## Direct properties — ethernet_interface / 213332312330 / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-0331123212002123-2021300312302231-0112021200213130-1333323300223102-0032121012222130-0223013313120221-1332033033003030-1000221030102001): complete subsection reference.

<a id="canonical-2231202301012103-0030231203033012-3333310030313313-2323112033020330-2121211203303212-0123221130102033-0322201003103320-3131123302203013"></a>

<a id="canonical-1310202222202212-0333230121332101-0021322203011131-2020333030320311-1022123212223001-0010321222010311-1020311013320123-3300333012130230"></a>

## device property — ethernet_interface / 213332312330 / 4

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

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

- [dhcp_client](resources--network_interface--reference--group-001.md#canonical-2130311230132221-3133220132001002-0000132231101303-3130210122232231-0123012021023021-1032303001111323-2002123222031132-1031113331320001): complete subsection reference.

- [dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122): complete subsection reference.

- [ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030): complete subsection reference.

- [is_primary](resources--network_interface--reference--group-001.md#canonical-2112312133010331-0302023100100333-3322231323213110-3302123302312303-3320110310302222-0110233222321212-3002203300130323-0022002111303332): complete subsection reference.

- [monitor](resources--network_interface--reference--group-001.md#canonical-1202200001221233-2202210301333102-2313310311302020-0113332023001130-0331031123021032-2313232302313220-1030023031111110-2103113103111213): complete subsection reference.

- [monitor_disabled](resources--network_interface--reference--group-001.md#canonical-1012000310230201-3031310213013132-3001301202120013-0102310113031200-1300230302030222-3033021303212131-1012210313123122-3103102222102113): complete subsection reference.

<a id="canonical-0103333211203321-3222311312203233-1330130203322013-0101021033110013-0303201203131102-2302200133200021-0133023222133232-1132312130031202"></a>

<a id="canonical-3321331300011110-0300122113313303-0222230031121013-2213100201010211-0103101330112020-2122122201303021-0331232103123220-1211233101121210"></a>

## mtu property — ethernet_interface / 213332312330 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
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
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-1333032111011330-0031133113021320-3311010213022003-0330032310232013-0203001200201102-0021320311032100-3220302232131331-2200200132331213): complete subsection reference.

<a id="canonical-0030020301233212-3023311302300223-1330212002102102-1010033222001332-1110322213132231-1011101313322333-1031120200131121-3201102200223323"></a>

<a id="canonical-0003320113231020-0103121302111300-1010321133300103-2130012133023013-3103301012301121-0111133133003021-2122001102003131-3102331330101110"></a>

## node property — ethernet_interface / 213332312330 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--network_interface--reference--group-001.md#canonical-1200321013310102-0032232013131303-0313231101032121-1032013113332222-3302320311120331-1320212223013321-0302331230222102-0102032230203100): complete subsection reference.

<a id="canonical-2310220202133120-3101221230123030-0230311003103210-2203330211033220-0122230101112221-0300232203032223-1013220012021201-2002302131322023"></a>

<a id="canonical-2011100101221311-2332023023012330-2100213131003022-3120031113222013-3123123232012310-2131301201221003-1000021001311000-3132132332100213"></a>

## priority property — ethernet_interface / 213332312330 / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

- [site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-0130021031001123-3103100331301201-3001310300213201-0220222112333033-1032212013313102-1032013030013102-0323300222121123-0313002332320133): complete subsection reference.

- [site_local_network](resources--network_interface--reference--group-002.md#canonical-1202003031332121-0320100021220100-0021231331001321-2313003211300332-3131230330132130-1100122212021011-3213120032122123-3021232101032213): complete subsection reference.

- [static_ip](resources--network_interface--reference--group-002.md#canonical-3022232102130132-0113030032123321-1012123313020210-2230003031231301-1032311210332033-2222203030010301-2030212000100303-0203232233101020): complete subsection reference.

- [static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-2002001211323102-2100132012330031-2302323310311332-2101323303113311-1132013232230332-3223112303200311-3323303331131101-3303332012102212): complete subsection reference.

- [storage_network](resources--network_interface--reference--group-002.md#canonical-1211221102021022-2110220001020321-2332331213201332-0002030102311330-0010010200013330-0330313231112202-2213212221310021-2102020220330011): complete subsection reference.

- [untagged](resources--network_interface--reference--group-002.md#canonical-2013330120312021-1032032031131103-2211211331111323-0310100111303230-0000211230300201-0103011311101311-0100022032120312-3310012312123032): complete subsection reference.

<a id="canonical-3030301320210110-0222330320212311-0121111212033303-0303200122003023-3031321033200230-0011130222212220-1310203122212023-3211133221322122"></a>

<a id="canonical-1011031320200332-2020230011201032-3233023012120203-0223201131232113-2230032120232013-1130320122201302-2102210213333032-1222233031300130"></a>

## vlan_id property — ethernet_interface / 213332312330 / 8

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

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
    "create": false,
    "minimum_config": false,
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

<a id="canonical-0211023220221001-0202130210132002-3122332213123303-1132201130133311-2210200030302221-2023102231212233-3233203033203130-2212300033132131"></a>

## Next pages — ethernet_interface / 213332312330 / 9

- [ethernet_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0331123212002123-2021300312302231-0112021200213130-1333323300223102-0032121012222130-0223013313120221-1332033033003030-1000221030102001)
- [ethernet_interface.dhcp_client](resources--network_interface--reference--group-001.md#canonical-2130311230132221-3133220132001002-0000132231101303-3130210122232231-0123012021023021-1032303001111323-2002123222031132-1031113331320001)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-2112312133010331-0302023100100333-3322231323213110-3302123302312303-3320110310302222-0110233222321212-3002203300130323-0022002111303332)
- [ethernet_interface.monitor](resources--network_interface--reference--group-001.md#canonical-1202200001221233-2202210301333102-2313310311302020-0113332023001130-0331031123021032-2313232302313220-1030023031111110-2103113103111213)
- [ethernet_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-1012000310230201-3031310213013132-3001301202120013-0102310113031200-1300230302030222-3033021303212131-1012210313123122-3103102222102113)
- [ethernet_interface.no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-1333032111011330-0031133113021320-3311010213022003-0330032310232013-0203001200201102-0021320311032100-3220302232131331-2200200132331213)
- [ethernet_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-1200321013310102-0032232013131303-0313231101032121-1032013113332222-3302320311120331-1320212223013321-0302331230222102-0102032230203100)
- [ethernet_interface.site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-0130021031001123-3103100331301201-3001310300213201-0220222112333033-1032212013313102-1032013030013102-0323300222121123-0313002332320133)
- [ethernet_interface.site_local_network](resources--network_interface--reference--group-002.md#canonical-1202003031332121-0320100021220100-0021231331001321-2313003211300332-3131230330132130-1100122212021011-3213120032122123-3021232101032213)
- [ethernet_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-3022232102130132-0113030032123321-1012123313020210-2230003031231301-1032311210332033-2222203030010301-2030212000100303-0203232233101020)
- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-2002001211323102-2100132012330031-2302323310311332-2101323303113311-1132013232230332-3223112303200311-3323303331131101-3303332012102212)
- [ethernet_interface.storage_network](resources--network_interface--reference--group-002.md#canonical-1211221102021022-2110220001020321-2332331213201332-0002030102311330-0010010200013330-0330313231112202-2213212221310021-2102020220330011)
- [ethernet_interface.untagged](resources--network_interface--reference--group-002.md#canonical-2013330120312021-1032032031131103-2211211331111323-0310100111303230-0000211230300201-0103011311101311-0100022032120312-3310012312123032)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0331123212002123-2021300312302231-0112021200213130-1333323300223102-0032121012222130-0223013313120221-1332033033003030-1000221030102001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103133233132232-1220313000212222-3321112323030111-1211333000003131-0233223132032010-3213031101223110-2120332031222030-0233320002320210"></a>

## ethernet_interface.cluster — cluster / 132211111013 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.cluster

<a id="canonical-0220123220130333-1012102313200310-3330102320132212-2022321301032322-0210311010333112-2111012102103133-2330133020101211-1320230003230321"></a>

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
cluster = {}
```

<a id="canonical-3000200131000311-2323001021231330-3113312103333023-2023110132120200-1033131112231222-3220211201311300-3223011111232112-3330300002010010"></a>

## Direct properties — cluster / 132211111013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300113202111222-2211130122231223-2213230022210102-0020123230113021-2112100130001121-0300002333101333-3333102132022313-0211121132100203"></a>

## Next pages — cluster / 132211111013 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2130311230132221-3133220132001002-0000132231101303-3130210122232231-0123012021023021-1032303001111323-2002123222031132-1031113331320001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033110002203320-2003002232311112-0222101311230200-1123002022220012-1112212222223323-2303313302033012-0022321303103310-2112101103121220"></a>

## ethernet_interface.dhcp_client — dhcp_client / 220022310122 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.dhcp_client

<a id="canonical-0302012113320211-2332113111223000-1212201120220231-3231001203001133-1020223112311100-3233021311313023-3023020110222332-3232010300003103"></a>

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

<a id="canonical-1231120110012021-1030013223010223-2012113303103121-2102200220320200-2313112123332202-1221310013010030-0001000300333211-2113200223220123"></a>

## Direct properties — dhcp_client / 220022310122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223310012222121-3310102223003112-1211210031320300-0012103303202323-1122223201223223-2313022120321210-0231211212121122-0201130233022313"></a>

## Next pages — dhcp_client / 220022310122 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122113233101202-1113033123330032-2311000332021211-2301123303331030-0230331102201120-2000032331331203-1213223212121312-3203310322211213"></a>

## ethernet_interface.dhcp_server — dhcp_server / 032321021321 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.dhcp_server

<a id="canonical-1321011133201010-3213023122122102-3132231032020003-3000021231021313-2321032000233101-0030232122001313-3112002110020011-0120223232233320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

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
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201222010313211-3320001112023202-1313003221312133-2303322230201133-3303322302032321-0030231231221303-1003201102131310-3102000123303011"></a>

## Direct properties — dhcp_server / 032321021321 / 3

- [automatic_from_end](resources--network_interface--reference--group-001.md#canonical-3002320333322321-2220003322203320-3112302231330330-2312012310202313-0022120020200230-3213113332100110-3231302232213302-1330203221022303): complete subsection reference.

- [automatic_from_start](resources--network_interface--reference--group-001.md#canonical-1322333201032002-2130012010212213-3310330320200133-3101121033211321-1013130030332232-3213110032002031-2033000203011202-0233022033000020): complete subsection reference.

- [dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322): complete subsection reference.

<a id="canonical-0032111201223221-2101312003023111-2011310233022112-2032132001202030-3332130333203200-0012311121213211-0223022102301222-3022202003230002"></a>

<a id="canonical-0301132132232002-0221020302130100-1000223101321032-2212003123321122-3132312200030012-0302101211103011-0111211333103131-0010300212113333"></a>

## dhcp_option82_tag property — dhcp_server / 032321021321 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-1230002001213013-3112212013133203-1211123303120002-3313321212100111-3331110102121033-0103220222212003-0102023223012331-1021210320010002"></a>

<a id="canonical-2213221302302030-1122013031022103-2313012203033301-2121033311112013-1303112130030220-1113103101312033-2210032110200013-1321103232003313"></a>

## fixed_ip_map property — dhcp_server / 032321021321 / 5

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

- [interface_ip_map](resources--network_interface--reference--group-001.md#canonical-1032110213302010-1223013131211001-3222001111333003-0230130123012331-1230230111222213-1023111332010313-1301322111322232-3120011033031332): complete subsection reference.

<a id="canonical-2212001012022032-1221103222221100-1121312220213233-0222200023032113-2213321303130010-3101322321011032-3333120011231312-0131103330313302"></a>

## Next pages — dhcp_server / 032321021321 / 6

- [ethernet_interface.dhcp_server.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-3002320333322321-2220003322203320-3112302231330330-2312012310202313-0022120020200230-3213113332100110-3231302232213302-1330203221022303)
- [ethernet_interface.dhcp_server.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-1322333201032002-2130012010212213-3310330320200133-3101121033211321-1013130030332232-3213110032002031-2033000203011202-0233022033000020)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- [ethernet_interface.dhcp_server.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-1032110213302010-1223013131211001-3222001111333003-0230130123012331-1230230111222213-1023111332010313-1301322111322232-3120011033031332)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3002320333322321-2220003322203320-3112302231330330-2312012310202313-0022120020200230-3213113332100110-3231302232213302-1330203221022303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131301112203001-0011230203232232-1211122131023331-3331300001223113-0133222210233202-2022230103011122-0302331111330320-2220103022220210"></a>

## ethernet_interface.dhcp_server.automatic_from_end — automatic_from_end / 210013200310 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-2201032021002012-3200330211033033-3112120101302012-2132323002202111-3132002320320210-0130130111112232-2323001210323201-2012022011132202"></a>

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

<a id="canonical-1330031012223321-2020331232112021-2221312000112312-1000320022222231-0221320101200020-3332022111131003-0130330232320001-0130031310331223"></a>

## Direct properties — automatic_from_end / 210013200310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233021233231323-1231222200322100-3223023212130003-0002112000310313-0210120100333211-2011133223110010-0221311311101320-3010013232033021"></a>

## Next pages — automatic_from_end / 210013200310 / 4

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1322333201032002-2130012010212213-3310330320200133-3101121033211321-1013130030332232-3213110032002031-2033000203011202-0233022033000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001333221201333-2123130031332020-1123211103330001-2232300301222322-2123032130021102-2020223330320222-1223031002222020-1001211212312221"></a>

## ethernet_interface.dhcp_server.automatic_from_start — automatic_from_start / 200332323310 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-0321120323122101-3032202223111330-1002201001210201-2013310112213313-0203132132012231-1113201021113311-2020100223122113-3330021013131213"></a>

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

<a id="canonical-1030020133320130-2212321023123102-1300321023010232-2132223011002031-2033021103113301-0012313223222213-0003210333021032-1232123121023013"></a>

## Direct properties — automatic_from_start / 200332323310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133313313322301-0030212223001132-3331000023231022-2032012021032302-3103013033123332-2200202021311212-0311322333003033-2203122030332113"></a>

## Next pages — automatic_from_start / 200332323310 / 4

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301003330023201-3321132202312300-1212122010311211-3302211102320322-1021022021202333-3212200321122202-0102130230230233-2032221222322320"></a>

## ethernet_interface.dhcp_server.dhcp_networks — dhcp_networks / 322201312031 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-3133300213321232-3010111202313011-3003030200203313-2311221032323331-0032202230011312-0120321302122113-2121330110202131-1031300013223320"></a>

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

<a id="canonical-2320310211310203-2020222112223100-1310011231222311-0302320033130330-3231231323110213-1232230301233211-1012103122300302-2102101102312211"></a>

## Direct properties — dhcp_networks / 322201312031 / 3

<a id="canonical-0030010231031022-3100200302020122-3312332020130310-3300312213133111-2231131223102221-2230230122213301-1021111211332113-2112210232123033"></a>

<a id="canonical-0300223310311313-1103223302113201-2102101121311203-2013301021101212-0131111221320303-0133021030031202-2121320121030322-2212112100203001"></a>

## dgw_address property — dhcp_networks / 322201312031 / 4

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

<a id="canonical-2312332002303201-3030333202000333-0203123100301233-1132320110022233-0223023013103330-0110133112020322-0101012202223302-0202200132303030"></a>

<a id="canonical-1031123302300312-0021133033331232-1312331221112021-3323322000232032-2132013113211021-2011330313212020-0103132222311320-1322003311122132"></a>

## dns_address property — dhcp_networks / 322201312031 / 5

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

- [first_address](resources--network_interface--reference--group-001.md#canonical-2031022032003101-2020131013112220-3212133333103023-0102133203201100-3313332032213210-3301101122011321-0222102131031012-2100012213102022): complete subsection reference.

- [last_address](resources--network_interface--reference--group-001.md#canonical-1320103312023013-1333131110302200-3203333131321000-3001013020301113-2332133101132330-1021122332032122-0333213011033212-3320210321033130): complete subsection reference.

<a id="canonical-3222112103312003-3032200001330003-1121121130311321-3122000131200003-1003001331002120-2012003022220213-3120230122001312-3310233233320033"></a>

<a id="canonical-0313112310013032-3031022212332230-2033231203221300-3020330333312202-1112010132003032-2020221133203303-1003332121033202-2131111213132012"></a>

## network_prefix property — dhcp_networks / 322201312031 / 6

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

<a id="canonical-1033320213023000-3023330012000312-0213123302000110-1120121033330233-1200121230222323-3233322001001032-1232132213122230-1013103033203322"></a>

<a id="canonical-2203131033111323-1012103031233100-1201002313210313-2010110120133231-0311012233021133-1030131313220213-0113110100312331-3130210333131221"></a>

## pool_settings property — dhcp_networks / 322201312031 / 7

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

- [pools](resources--network_interface--reference--group-001.md#canonical-0022313112313210-0123113032330113-1133331031313132-0123232333232223-0221003112011231-0100203021231322-1113212223121130-3203121100201013): complete subsection reference.

- [same_as_dgw](resources--network_interface--reference--group-001.md#canonical-1321333110000032-2221101020011322-3000213223311131-1111013300023003-2121013313122033-1013312133130021-2112121123101312-2113110010202333): complete subsection reference.

<a id="canonical-3100010230123022-0230320033300020-0333322223222003-2231110210122323-2313202113331100-3232110030211201-3032033311230321-2131103312232131"></a>

## Next pages — dhcp_networks / 322201312031 / 8

- [ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--network_interface--reference--group-001.md#canonical-2031022032003101-2020131013112220-3212133333103023-0102133203201100-3313332032213210-3301101122011321-0222102131031012-2100012213102022)
- [ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--network_interface--reference--group-001.md#canonical-1320103312023013-1333131110302200-3203333131321000-3001013020301113-2332133101132330-1021122332032122-0333213011033212-3320210321033130)
- [ethernet_interface.dhcp_server.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-0022313112313210-0123113032330113-1133331031313132-0123232333232223-0221003112011231-0100203021231322-1113212223121130-3203121100201013)
- [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--network_interface--reference--group-001.md#canonical-1321333110000032-2221101020011322-3000213223311131-1111013300023003-2121013313122033-1013312133130021-2112121123101312-2113110010202333)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2031022032003101-2020131013112220-3212133333103023-0102133203201100-3313332032213210-3301101122011321-0222102131031012-2100012213102022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103031321222203-1113031103132110-0332020113123322-0220330132032203-3130133111321202-0031210002020131-1331213310212111-2100201221010120"></a>

## ethernet_interface.dhcp_server.dhcp_networks.first_address — first_address / 321000101221 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-3311013102102330-1310033122331201-3120330333200032-0201112331031033-2002222123201110-3001030331302131-1110023013033020-2330233231103121"></a>

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

<a id="canonical-1001330301132031-2122101011032113-3202232030110312-0120333033320230-0313112103132121-2202313213101000-2322311030211002-0311233223001011"></a>

## Direct properties — first_address / 321000101221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201331310033033-1133031210301200-1131022210322330-1110030302330331-1203111303213323-2103300122203100-3313121123321313-2333011110011111"></a>

## Next pages — first_address / 321000101221 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1320103312023013-1333131110302200-3203333131321000-3001013020301113-2332133101132330-1021122332032122-0333213011033212-3320210321033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232121200101321-2202323302000021-3223300112233002-3322323230230311-1211322113100000-0320131203133022-2213310002111332-1003022120312320"></a>

## ethernet_interface.dhcp_server.dhcp_networks.last_address — last_address / 033221321320 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-2132331222111013-0311020021300330-2011222213033133-3320013300031302-1030031113201313-0231220322320011-0001100020020030-0122211101232323"></a>

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

<a id="canonical-3030333332002321-1202012033330203-2112011023230100-2112312002031303-3332210101132103-1133211213033312-1321323113320000-1231103110001323"></a>

## Direct properties — last_address / 033221321320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203133221113222-2233320321223222-0332023333123211-0130210131010323-0032010233223003-1302120022013301-3232322323020321-2311313302301211"></a>

## Next pages — last_address / 033221321320 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0022313112313210-0123113032330113-1133331031313132-0123232333232223-0221003112011231-0100203021231322-1113212223121130-3203121100201013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012202132321212-1030103022100201-3101111030010300-0113203322312212-3011233302032013-3103103312211021-0020312133322323-1001003231131000"></a>

## ethernet_interface.dhcp_server.dhcp_networks.pools — pools / 213312100033 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-2223202202331203-3330223231103122-1130331113212002-0022221122310333-1131013310321002-0100031302232313-1212133030030201-3103001210003011"></a>

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

<a id="canonical-1201010303332023-0312113113303031-0211001230131332-3002302100302221-3011101132233023-3311210332110230-1310211133011122-0111301012211221"></a>

## Direct properties — pools / 213312100033 / 3

<a id="canonical-0302321113203030-0210230031312113-3132111120012233-1103031021323110-2123310330022020-3101131202333203-0212121221202110-3332303303233010"></a>

<a id="canonical-0302302322110021-3002012000221112-2231000220303102-2132100223322031-2323200303101230-0033332303002112-1233202310223320-1110002112331330"></a>

## end_ip property — pools / 213312100033 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-2022332220232330-0020001123112101-2201033122011013-0212003321110202-1132130221112223-2223212312200120-1130120201310113-3231120211121120"></a>

<a id="canonical-0031020131121130-0322102320000000-3011130012131122-2133222311021111-2313110132021201-2103001222223003-0111130103310003-2112311031310230"></a>

## exclude property — pools / 213312100033 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1012230212031230-3303030113131033-3011331032003322-2113220313002011-3012121302220003-0311003030011200-0122232302121323-3210000112322123"></a>

<a id="canonical-0131212002331331-3102002321313103-0133212221023131-3030023102210321-3210303211113121-0032232020200330-0012310203202102-1200113000233303"></a>

## start_ip property — pools / 213312100033 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-3222123311230203-2023303001330020-0010003130011221-0130220103132002-3321200021012012-0113011122001023-3303123311333022-1010021113203321"></a>

## Next pages — pools / 213312100033 / 7

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1321333110000032-2221101020011322-3000213223311131-1111013300023003-2121013313122033-1013312133130021-2112121123101312-2113110010202333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221203001000012-1322232120030203-0013101001113200-1103020331332130-1303303103312012-1133010003202221-0233222103023033-3032112101301110"></a>

## ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 231310211301 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2222310032012312-3030103003133122-0103030231333021-3312032300333300-1103212221022331-0233121003211302-3303212322101203-1312102310213120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-3021002031300100-1322233132113311-3210101332201113-2303322333211101-0032321030003112-3220112232000030-0211212233323133-0302102130122012"></a>

## Direct properties — same_as_dgw / 231310211301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333200130123030-0101202113203111-1010321313203022-3320101023012332-2331123313332103-3122133223201101-2120220123132002-2220213222131211"></a>

## Next pages — same_as_dgw / 231310211301 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1032110213302010-1223013131211001-3222001111333003-0230130123012331-1230230111222213-1023111332010313-1301322111322232-3120011033031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100033132113312-1310312022102301-0100311102013331-3013233313110013-0330012120303103-0333311033133200-2023202310331121-3201203321332300"></a>

## ethernet_interface.dhcp_server.interface_ip_map — interface_ip_map / 332100210331 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-2222101101101123-1101002022233301-0132313311110220-3201100222031202-0121132031231001-3223213021120103-3313300300021132-1023110322122023"></a>

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

<a id="canonical-1111222220100321-1221020332310002-3120030013213202-0221232030202100-0223213322213122-0013211233132100-3321300223202312-2213011033013020"></a>

## Direct properties — interface_ip_map / 332100210331 / 3

<a id="canonical-3331002322130133-3113321120233111-2321122322322000-3323003000103130-3101011022330112-3120313133311310-3133111002012201-1123200131100113"></a>

<a id="canonical-1221111313321111-3321002201003231-3213022123003210-3300002233030003-0012000003223111-2133232210101221-1213002112213313-2320233230210300"></a>

## interface_ip_map property — interface_ip_map / 332100210331 / 4

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

<a id="canonical-2322030002221220-0230122320303222-0131100223313330-1231310213200320-2202120000111311-2111330001020212-2132000232133330-1112012230030231"></a>

## Next pages — interface_ip_map / 332100210331 / 5

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-1213321221100323-2203020113023012-1302133211310103-2211122222221022-2321023310300213-2331201231210033-2321200020302102-0332020332022122)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323121013023311-1221220010103303-0011211313002332-3203333320330133-0030003301130131-3013221203311222-2322202300020330-2012331013201210"></a>

## ethernet_interface.ipv6_auto_config — ipv6_auto_config / 230101300120 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.ipv6_auto_config

<a id="canonical-0003200332000001-2222121003003222-0011010201323213-1001233230310310-3231031012011123-3330110202333131-2312213323300000-0031203302233330"></a>

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

<a id="canonical-0201323103010320-3132211010302310-2233001310211322-0201022301313330-3011120301211121-1331310232233001-2302123013031122-1021211331033022"></a>

## Direct properties — ipv6_auto_config / 230101300120 / 3

- [host](resources--network_interface--reference--group-001.md#canonical-0210013010220110-2103131000100200-0020220111013130-0321133122222231-3001101223231232-0200220202022120-0000132001120013-2303013120031100): complete subsection reference.

- [router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222): complete subsection reference.

<a id="canonical-1313101332112110-3202213103121102-2302323300322233-1113222023220031-0123310103020321-1110112232103101-2103333133131002-1032021010021300"></a>

## Next pages — ipv6_auto_config / 230101300120 / 4

- [ethernet_interface.ipv6_auto_config.host](resources--network_interface--reference--group-001.md#canonical-0210013010220110-2103131000100200-0020220111013130-0321133122222231-3001101223231232-0200220202022120-0000132001120013-2303013120031100)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0210013010220110-2103131000100200-0020220111013130-0321133122222231-3001101223231232-0200220202022120-0000132001120013-2303013120031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103131022332222-0310321230303200-2113210212111102-3322223331232222-1031013032301320-3302030100132210-3103001123002023-2231101133220021"></a>

## ethernet_interface.ipv6_auto_config.host — host / 021313311211 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- ethernet_interface.ipv6_auto_config.host

<a id="canonical-0231120323200230-3012101212131132-2123120213132010-2003212212000023-1233013201301110-0011210322022132-2100203321323121-0320200332311212"></a>

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

<a id="canonical-0122232221001011-3222013222022031-3102323000122203-3022200202312032-2331300322202123-1322032212121133-3033303110330222-2012302131023312"></a>

## Direct properties — host / 021313311211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223013011021201-0023021232131020-0100033221130032-0331211032322210-1121023010120120-2113202121331023-1312330131202300-0210223133213203"></a>

## Next pages — host / 021313311211 / 4

- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030033011300312-3110323303133133-3133123320113233-1213012201330220-1202010323212302-2030122101001022-1032133213001311-0213000203212131"></a>

## ethernet_interface.ipv6_auto_config.router — router / 021030101011 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- ethernet_interface.ipv6_auto_config.router

<a id="canonical-1120213113131012-0303211103323100-1332321311203032-1122010002103233-1121120020231103-3030213020102331-2103020021202312-0101002210313030"></a>

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

<a id="canonical-1123213010233220-0120213121323012-2001023120232230-0201230113203210-2322310003102010-1213200212123031-2021111230101232-3020322221303322"></a>

## Direct properties — router / 021030101011 / 3

- [dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311): complete subsection reference.

<a id="canonical-0112210312131011-2230222012100303-2010233001311011-1321222110210333-2211133022021010-1031003301210320-3301303212333122-2023130323200100"></a>

<a id="canonical-2123201203311033-3123323101330112-1202032322223003-2210232112210100-2202333010133010-1133112113130101-3311031000231100-3031001211023100"></a>

## network_prefix property — router / 021030101011 / 4

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

- [stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102): complete subsection reference.

<a id="canonical-3200130322123301-0010333301121012-1230310320300032-1310113333211213-0300032002132302-1201231111032111-0112002033113121-1302100322013030"></a>

## Next pages — router / 021030101011 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301223132321202-1103021010000332-2323133110111231-1221323312103032-0220211133111001-0112133222030213-1202331101322211-1212110111032332"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config — dns_config / 222331120223 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-3130110331133221-0212302031333103-1100012112233101-3222121030133220-0311322231300213-2302021012120331-3231321022212311-1212011222221321"></a>

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

<a id="canonical-1032030302111320-3002033000132113-1000202101123021-2332301320213033-0130133220033220-3112013003300102-3223100232031301-1201321010313103"></a>

## Direct properties — dns_config / 222331120223 / 3

- [configured_list](resources--network_interface--reference--group-001.md#canonical-0231301112223203-3321303100222022-0010023322110221-1023200010001230-3031001110033330-0033213121333032-3333200331032100-2012132203112220): complete subsection reference.

- [local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033): complete subsection reference.

<a id="canonical-3320022003331020-0200130102330021-1210012311000221-3212001131102233-2011201220013100-0033111231101211-2031320021003123-0111013032212310"></a>

## Next pages — dns_config / 222331120223 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--network_interface--reference--group-001.md#canonical-0231301112223203-3321303100222022-0010023322110221-1023200010001230-3031001110033330-0033213121333032-3333200331032100-2012132203112220)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0231301112223203-3321303100222022-0010023322110221-1023200010001230-3031001110033330-0033213121333032-3333200331032100-2012132203112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032010223023321-0201232031223120-0113020113033210-3110123331332332-3111323100200000-0303332032312231-2320310213033131-2202122130300120"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 313010233000 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3012211213002333-3310231101202331-0211032123001113-1321213311202223-3310323320230230-2033100331232203-0012121312300023-0203223131301210"></a>

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

<a id="canonical-1302322231300213-3200211212210231-0200021213313213-1313332333332303-3312300322001021-2230033310030321-1012230110323232-1123222230220012"></a>

## Direct properties — configured_list / 313010233000 / 3

<a id="canonical-0033300023120332-3202030212120322-2301033220130112-2302013213201132-1213022010332302-1001120313203312-2212123220323203-0122132120231013"></a>

<a id="canonical-1322221322321121-1330132132221331-1112233010321032-0233232011012313-1131221132223313-2223300011220232-3023002313012123-0103322100003101"></a>

## dns_list property — configured_list / 313010233000 / 4

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

<a id="canonical-3212000332202221-1011323122231030-2321033132021010-0002020022312211-2300032112301100-0202123001030213-1103203120302332-3132022323321023"></a>

## Next pages — configured_list / 313010233000 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003333321020030-1332230120232201-1013203202212300-1211201103020130-0320213313131202-1232312303310200-0033123321201103-1231221132110310"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 012102023032 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2133212032322311-3201131202120223-0230011022001130-2013013011300032-1020011211210021-1323210233230233-3121032101113323-1301301030010312"></a>

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

<a id="canonical-0201302321321220-1303302022100230-3122033333002032-1001132102131110-2222112301211102-3130312113332122-1030022302031101-2220223310311333"></a>

## Direct properties — local_dns / 012102023032 / 3

<a id="canonical-0201032022210101-1211220031123321-2022000133323233-1211112113033211-0303210010213112-2220311122331222-2102211023320300-3020130323111323"></a>

<a id="canonical-1302230231021203-0002000011013022-0200313201232000-1203111301321300-3113000223001131-0213132011110103-1333121003133022-2131133133100311"></a>

## configured_address property — local_dns / 012102023032 / 4

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

- [first_address](resources--network_interface--reference--group-001.md#canonical-0203123000333232-3310121312202201-0301202233032233-3121212221013222-0132122130333211-3123221312101112-2122020123323133-3223131032203132): complete subsection reference.

- [last_address](resources--network_interface--reference--group-001.md#canonical-1021321302330133-3032032000020112-1312331032103010-3130102100321110-1303323320211310-1100122212103120-0332223130010212-1313331030303200): complete subsection reference.

<a id="canonical-0333323233032103-1310100120303022-3112302123303231-3111031000232332-2210121102030001-3303203011313022-0101232203330222-2302331313110002"></a>

## Next pages — local_dns / 012102023032 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--network_interface--reference--group-001.md#canonical-0203123000333232-3310121312202201-0301202233032233-3121212221013222-0132122130333211-3123221312101112-2122020123323133-3223131032203132)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--network_interface--reference--group-001.md#canonical-1021321302330133-3032032000020112-1312331032103010-3130102100321110-1303323320211310-1100122212103120-0332223130010212-1313331030303200)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0203123000333232-3310121312202201-0301202233032233-3121212221013222-0132122130333211-3123221312101112-2122020123323133-3223131032203132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031000120013313-2332330020233312-1120221123330331-1222213012023012-3321212112200210-2212220320210212-3233132132333220-0300312312130322"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 110000333113 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0002201233031130-2103033202331230-0023322000132230-0221131110101332-0312021002232212-3203312100011231-2300323002022221-3333010223030022"></a>

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

<a id="canonical-1010303003021123-2110332013102301-1333020202030000-1303012132203213-3330111331232033-1223323211121302-0122101130222222-2331303122222331"></a>

## Direct properties — first_address / 110000333113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113312230200223-0022301200123213-3101013230003130-0110221032321022-1302233303303230-0133300321001132-2123201321121123-0313223303213021"></a>

## Next pages — first_address / 110000333113 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1021321302330133-3032032000020112-1312331032103010-3130102100321110-1303323320211310-1100122212103120-0332223130010212-1313331030303200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100131202233003-0130130222231333-1311101122201110-2133333113103000-2121002000000313-0202303033312202-0133211203233030-0200223213103110"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 313210130311 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2110233021223032-2123023311322013-1312311300010320-3103233222131103-3221301312320030-3320310300323120-2333211212002312-1211003131323001"></a>

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

<a id="canonical-3132332103203121-1002312101021233-2112023313230002-0301313102032033-3022230023222211-2213210000231313-3010033001211322-0020303212330232"></a>

## Direct properties — last_address / 313210130311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020000303121110-0303023033222100-0133103101211321-3100010232100320-1132232100330021-1311302212021001-1031102010121012-1220201332000120"></a>

## Next pages — last_address / 313210130311 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-1201330132333331-0012102212122321-2023032103123202-3133332032130320-3221320302111300-1121000313201322-1013112001102311-1011133023223033)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330123110211230-0003011031012032-0213131131203032-1101333013010222-0101210001132110-3212011331003133-2302202333332123-1212123120113123"></a>

## ethernet_interface.ipv6_auto_config.router.stateful — stateful / 200310222220 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-0113300320232133-0200303131001233-0011303331001330-3221132011110100-0230122321003230-3211233103301320-2000003320333230-1123110310221122"></a>

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

<a id="canonical-1011130103013311-3320121001112330-0320001111102132-1213311021103121-1123103021121112-3221013300030010-0033302131310111-0321032020200333"></a>

## Direct properties — stateful / 200310222220 / 3

- [automatic_from_end](resources--network_interface--reference--group-001.md#canonical-2111313300312231-0133133012202220-0010021202000123-3133033320013323-1233221022331001-0011130311120233-1311131000322211-3111000031011030): complete subsection reference.

- [automatic_from_start](resources--network_interface--reference--group-001.md#canonical-1023022131232321-2132112210021013-1331132112010011-1113320330002122-0012303012123321-2331203013321210-1000102130111213-0032333203102021): complete subsection reference.

- [dhcp_networks](resources--network_interface--reference--group-001.md#canonical-0333310233021331-1112210231320333-0322200233232322-1331131300002310-3333013133232232-2233130010001302-2223313222310220-3021103030032012): complete subsection reference.

<a id="canonical-0031103130122233-1030012002020030-1312300231121033-2323031120110023-3130100320132312-3310002202330222-2000000102000221-1303003032130330"></a>

<a id="canonical-1302221022202303-3111201022032012-0012033333101221-0131320030002032-2122300332101033-2030131303002212-0023023112132132-2302311231222003"></a>

## fixed_ip_map property — stateful / 200310222220 / 4

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

- [interface_ip_map](resources--network_interface--reference--group-001.md#canonical-3002222031010110-3211012230203103-0310132220212232-3311011101001203-1212212130332013-1112130013200011-0211201033330322-2022132023000020): complete subsection reference.

<a id="canonical-3123232002231031-3020011130111312-3022331321131203-3130123123210033-1200122212012020-3100021122121122-1221323121003112-1222100230131002"></a>

## Next pages — stateful / 200310222220 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-2111313300312231-0133133012202220-0010021202000123-3133033320013323-1233221022331001-0011130311120233-1311131000322211-3111000031011030)
- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-1023022131232321-2132112210021013-1331132112010011-1113320330002122-0012303012123321-2331203013321210-1000102130111213-0032333203102021)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-0333310233021331-1112210231320333-0322200233232322-1331131300002310-3333013133232232-2233130010001302-2223313222310220-3021103030032012)
- [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-3002222031010110-3211012230203103-0310132220212232-3311011101001203-1212212130332013-1112130013200011-0211201033330322-2022132023000020)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2111313300312231-0133133012202220-0010021202000123-3133033320013323-1233221022331001-0011130311120233-1311131000322211-3111000031011030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031221012303013-2011131223002121-3111201013122311-0302123202103210-2113032232333303-2232021013121311-1203302030021202-3212001020323313"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 200020101122 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1230221233010222-3132312222120202-2100100011133110-0133303301200323-3231032302020033-1311330100023202-1320011012200122-0012022301111103"></a>

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

<a id="canonical-3032333331020100-2010002132230121-0202231223133321-1202003313111131-3202003312232310-2310120313213301-2130333012022211-2100021111011001"></a>

## Direct properties — automatic_from_end / 200020101122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300301000220130-3230313213210230-1223221322333032-1202300310200113-2213332123030221-3301220210211020-2003121320010322-3013130221100132"></a>

## Next pages — automatic_from_end / 200020101122 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1023022131232321-2132112210021013-1331132112010011-1113320330002122-0012303012123321-2331203013321210-1000102130111213-0032333203102021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321230001301322-2313323100302200-3011200111320232-2012131000000330-3122322011110130-2003232021123333-2003303121012301-0233112002120031"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 202322123022 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2023331031332113-2320320210333121-0211333111203111-1221310200222111-2111321111120021-2321032032200320-0130121333122322-1001112321321100"></a>

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

<a id="canonical-1213003310010111-0033120132132202-1030232121323320-1103101030312123-1221223302201211-1131212112233000-3311333202132132-1203313201012323"></a>

## Direct properties — automatic_from_start / 202322123022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121313220313133-3131313030333313-3312312331100123-2122312131030333-1320101000002012-0213020102312323-3111312033113032-3203212111032233"></a>

## Next pages — automatic_from_start / 202322123022 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-0333310233021331-1112210231320333-0322200233232322-1331131300002310-3333013133232232-2233130010001302-2223313222310220-3021103030032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013020012022222-0133001003220101-0222032212120300-3203201202222313-3233223221222221-2020022231021230-3203312031030010-1003300321323123"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 303033232200 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2023003303132312-0030112320201211-3132230331011122-3321233220131103-2222102013220313-1031132103031322-1022221311322031-2311311200133332"></a>

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

<a id="canonical-3210123112333303-1223111223000012-3202302130322023-0021233123313213-0311103032231212-3310311130311022-0223202101101031-0320303103233132"></a>

## Direct properties — dhcp_networks / 303033232200 / 3

<a id="canonical-0202011311333310-2121231020100001-3231020131023120-3020103123011030-3202011012031130-1012030112011023-2212120202110000-0332011130013121"></a>

<a id="canonical-2032010033200123-2321111131230022-2013202013101131-3000201133122203-1223220313130123-2310320002302013-0031212201103023-3333113311000032"></a>

## network_prefix property — dhcp_networks / 303033232200 / 4

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

<a id="canonical-0022131013323311-1010302310032300-1321110300010012-3120200012313322-1023012030212123-2132202310202033-0103333232300003-2323123333313021"></a>

<a id="canonical-2021020021320210-0022221030201130-3012220310222120-1111323312302210-2101312323330112-0123112110232312-1201322310032203-2301013210012012"></a>

## pool_settings property — dhcp_networks / 303033232200 / 5

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

- [pools](resources--network_interface--reference--group-001.md#canonical-2202223101031000-3030203121200100-1332022120203032-3120233110112303-2121212322000102-1110202011013113-2221102012311133-1232232123033322): complete subsection reference.

<a id="canonical-2121123231201222-1123231003120221-1212222033202320-3123000211332110-2331230321021311-2132232303013220-2112113220211210-1110313323132233"></a>

## Next pages — dhcp_networks / 303033232200 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-2202223101031000-3030203121200100-1332022120203032-3120233110112303-2121212322000102-1110202011013113-2221102012311133-1232232123033322)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2202223101031000-3030203121200100-1332022120203032-3120233110112303-2121212322000102-1110202011013113-2221102012311133-1232232123033322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102302121311331-1003313210233102-2233300031333221-2120200232310003-1312021111231220-0230010030003302-3310311100200012-0011101200220120"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 102312331203 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-0333310233021331-1112210231320333-0322200233232322-1331131300002310-3333013133232232-2233130010001302-2223313222310220-3021103030032012)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0013020213320311-0222020330030300-0320302003131030-0323222212022211-3223111001322200-0000303320003200-0220230003321200-0020211122132033"></a>

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

<a id="canonical-3102132331102033-2231012132022002-1302123330221000-3222312332102312-2132020231031032-2310103320200212-1232121230212011-0010320320000112"></a>

## Direct properties — pools / 102312331203 / 3

<a id="canonical-1233000101101012-3130321110102231-1230133132301211-1231003013123230-2211100132201113-0301133302013121-0331213010312312-3012310002010213"></a>

<a id="canonical-0320311021303230-2210211222033220-0233300310112303-3031110133120001-2030203313212020-1333013213012202-0113321211132222-1302203111013202"></a>

## end_ip property — pools / 102312331203 / 4

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

<a id="canonical-1301020033322311-1132022121001201-2121133010003102-2201120031003311-2012112000100121-1231303032013232-1023200202210231-3212210232103123"></a>

<a id="canonical-0022030210023202-3300313320003321-3101211020323121-1121333111322120-1213200222333311-1111301333211301-1002223102132021-3311301002133303"></a>

## start_ip property — pools / 102312331203 / 5

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

<a id="canonical-1033121131032202-2011131022122212-1002212323323200-2130023031112012-1100030331113101-3323131330101113-2321120021122313-1320001100132311"></a>

## Next pages — pools / 102312331203 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-0333310233021331-1112210231320333-0322200233232322-1331131300002310-3333013133232232-2233130010001302-2223313222310220-3021103030032012)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-3002222031010110-3211012230203103-0310132220212232-3311011101001203-1212212130332013-1112130013200011-0211201033330322-2022132023000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311012133001330-0102230213222313-3120133122130323-0213322231332322-3001023310132303-1310223120103012-1103012213022331-1221322001032001"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 213211200313 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-3111001330200033-1201101132233322-3022122131001003-3330012133233321-2313023311001213-3001032233030012-1021321321131100-2202111030013030)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-2132013220231010-2002322112213211-1023313231233111-3312321101222022-3132210120003020-1001312330212313-3203302123132300-0221313202102222)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2031011232113112-3302330330322032-2232133020302222-0200002110203012-3333222012202010-3312301103200000-3033001322201320-2030311313031001"></a>

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

<a id="canonical-0100313012213231-1303033011303200-2001223313110023-2021103123012023-2133310213012002-2212120301322302-3132030333000120-0212301221112212"></a>

## Direct properties — interface_ip_map / 213211200313 / 3

<a id="canonical-0332213212132233-1210301013132220-1331210031220212-3300212330020301-2200112121233202-1021232302010120-0003132211211211-3221223023230012"></a>

<a id="canonical-0111023122231120-3030300320222130-0113321131122110-2333012023230100-0233311233320213-1333133301200032-3130313111332222-1002130232123032"></a>

## interface_ip_map property — interface_ip_map / 213211200313 / 4

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

<a id="canonical-2032202003031123-3201131102033210-3321021021223123-3221303300300130-1122123322332330-0111102003002132-2000222033100202-3313002111212230"></a>

## Next pages — interface_ip_map / 213211200313 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-3223311332330012-2012301231200010-2230212022113101-2033112301221322-2310031032322203-1320220331120323-0323212012033033-3101302330110102)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-2112312133010331-0302023100100333-3322231323213110-3302123302312303-3320110310302222-0110233222321212-3002203300130323-0022002111303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132030333122302-2012211203133110-0232013111321330-0312212120321021-3221113331301122-2220032231223123-3133233212003112-0330100133011223"></a>

## ethernet_interface.is_primary — is_primary / 113331201110 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.is_primary

<a id="canonical-0222021130202303-0012130030002002-2312303031202011-3332232313113021-0213321223213131-1123101111121222-3011013120111331-2111223130232222"></a>

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
is_primary = {}
```

<a id="canonical-3212133212320130-2003221232231112-3010310000332010-2123213301132233-2121211320132011-2322310113220212-0310132212013331-3112313202110120"></a>

## Direct properties — is_primary / 113331201110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333333101111001-0102032010110322-1023312032300201-1110331022022210-0232130122222023-2220111120013222-2311212130101110-3003323322230000"></a>

## Next pages — is_primary / 113331201110 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1202200001221233-2202210301333102-2313310311302020-0113332023001130-0331031123021032-2313232302313220-1030023031111110-2103113103111213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132002121211232-0210131213212031-3332222310122133-1313223121101311-1300223001021323-2201113312320202-3311320220211332-1030310312010311"></a>

## ethernet_interface.monitor — monitor / 333010033130 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.monitor

<a id="canonical-0230023222211001-1001311133113012-0301131201202113-2310201320103332-3322121211130303-3012031321310000-2213221032112312-1012331221302133"></a>

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

<a id="canonical-1313332123001133-3303103032213020-0213221113033103-0212201022102331-0001012002230232-2123011023012322-2320103110130133-3110103331333212"></a>

## Direct properties — monitor / 333010033130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111001321301110-1100213221111121-1212201122232022-2002032300010220-2320122033233003-0321103012012330-3131002012300232-2130130032213233"></a>

## Next pages — monitor / 333010033130 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1012000310230201-3031310213013132-3001301202120013-0102310113031200-1300230302030222-3033021303212131-1012210313123122-3103102222102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323110111330321-0132020322230333-0132311121012000-2201320331000120-2010003133133301-1123123032302231-0001110210201322-3001323102130003"></a>

## ethernet_interface.monitor_disabled — monitor_disabled / 310232303032 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.monitor_disabled

<a id="canonical-0000230123130100-1020033121000221-0212002311321130-1211301011321031-1230313023010213-3321310120210202-3301021330121321-1202020133321323"></a>

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

<a id="canonical-0022222003202302-0010232221023030-0111123012011233-0222201101100130-3333010131121302-1211033231033311-2232013211021130-2123310321013021"></a>

## Direct properties — monitor_disabled / 310232303032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321313133220130-1322330320132003-1310323322012111-2323203011003221-2221222320233023-0211220123332303-3100320101212012-0222030321202022"></a>

## Next pages — monitor_disabled / 310232303032 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1333032111011330-0031133113021320-3311010213022003-0330032310232013-0203001200201102-0021320311032100-3220302232131331-2200200132331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321001213111223-2121032320103030-0020213101311311-0133302212110203-1111133201322201-0221320132212123-1230020030102302-3021001333012012"></a>

## ethernet_interface.no_ipv6_address — no_ipv6_address / 031001222133 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- ethernet_interface.no_ipv6_address

<a id="canonical-2001322023313200-3011100222320010-1220210030303113-2312302221132330-2133103012133230-1033000003002110-2323333202332011-1201003111200123"></a>

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

<a id="canonical-1221232301020331-1103111321020111-0222301110103302-3020302010021112-3211230233302313-0220301103012113-3222113330202202-2233322312023101"></a>

## Direct properties — no_ipv6_address / 031001222133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030030120030111-1312020201001100-2203033012022000-3003131030322301-1332312212023130-1202030312202031-2033222312303301-3233333221012203"></a>

## Next pages — no_ipv6_address / 031001222133 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-0212013301302220-1201200231123120-0000000232201111-2330233033122002-2203013033101202-2132013330003201-1213023112010321-1203321102031221)
- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)

<a id="canonical-1200321013310102-0032232013131303-0313231101032121-1032013113332222-3302320311120331-1320212223013321-0302331230222102-0102032230203100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
