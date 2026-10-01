---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232212311231212-1111031300200321-0311232111131211-2311333323032230-1031310113113311-3311132130210103-0121223013003012-0020311130232220"></a>

## Property reference — Property reference / 000001200001 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- Property reference

<a id="canonical-0231231210330030-1032202230322011-0130031132332200-0220210300101303-0302331100200113-3131311330330000-3313131120131003-2122232110231001"></a>

## Direct properties — Property reference / 000001200001 / 3

<a id="canonical-3211030103000012-1331120300220301-3311102331310100-1232312121031132-3223303002313010-2221313222313032-3022012233323032-0300122300111201"></a>

<a id="canonical-0012223032233132-1111011320223030-1022001230012200-3232110330002132-2230200302232002-3033100310322030-3130123302010132-1213032102012002"></a>

## annotations property — Property reference / 000001200001 / 4

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

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321): complete subsection reference.

- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302): complete subsection reference.

<a id="canonical-3330131101312002-0000110202332320-2333113001120202-0203302333113031-2133110100023020-1030103021002123-0111330232023220-1112003110103333"></a>

<a id="canonical-0323102223301200-2110030310023223-2130332302100110-2212220122311013-2300320200012320-0010013211302000-2313011323012331-3223112001131103"></a>

## description property — Property reference / 000001200001 / 5

Type: `"string"`. Computed.

Description of the NetworkInterface.

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

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112): complete subsection reference.

<a id="canonical-1310210210213110-0000112020200133-1023333321001021-0200022320130003-2330123011023011-1331003200133311-0201021310122332-3110211213113232"></a>

<a id="canonical-2221132122121030-1113230311031200-0101230033111103-3220002312323200-1303233123012131-1022022133312301-2130032222321113-2100310130212001"></a>

## ID property — Property reference / 000001200001 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3300321303201002-1102201313323311-2131100200012330-2231330233132300-0323122313013300-0012110210023300-0023133323333332-0220202333003010"></a>

<a id="canonical-3311010202032220-0233121302202311-2111122310323032-3103213200211030-3313303103221221-3223121210010031-2111213101333132-3013220323000103"></a>

## labels property — Property reference / 000001200001 / 7

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

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030): complete subsection reference.

<a id="canonical-0110322303330303-3213131003322113-1133230103212113-3123331032212132-3032202223321313-3012203032223023-0311032131322202-0311010122021220"></a>

<a id="canonical-1030302112113230-1331121203110303-0033333033003301-2303000122121011-3003121203023331-2223012030203032-0221332212012333-3212201333310013"></a>

## name property — Property reference / 000001200001 / 8

Type: `"string"`. Required.

Name of the NetworkInterface.

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

<a id="canonical-0031200023311012-0003321231121010-2022201013012221-1123320023231132-0102221111322012-1011323133021122-3322011130133232-1023331333000122"></a>

<a id="canonical-2300021021321203-1120312302020012-2303330010123011-3131330203121223-2201301232000302-3121331120230023-0222030222103031-0121010323002022"></a>

## namespace property — Property reference / 000001200001 / 9

Type: `"string"`. Required.

Namespace where the NetworkInterface exists.

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

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313): complete subsection reference.

<a id="canonical-2230331120023132-1312311032222300-0312323021002123-2231313300231032-1111303210132233-1030030133010020-1012020102321223-1312200220013221"></a>

## All schema paths — Property reference / 000001200001 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_interface--reference--group-001.md#canonical-3211030103000012-1331120300220301-3311102331310100-1232312121031132-3223303002313010-2221313222313032-3022012233323032-0300122300111201) |
| `dedicated_interface` | [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2313131122020221-0233301323112002-0302000311302211-3112301213001122-3301001203003313-3122131012332030-3221323232202001-3330300001121330) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-0120233122333101-0022201212300031-0001202002111303-2003020000003100-1021011212112301-3020301023301103-2002312031220103-1303120003111133) |
| `dedicated_interface.device` | [dedicated_interface.device](data-sources--network_interface--reference--group-001.md#canonical-0111130032120211-0202231123231000-3310100012301202-2101312100031211-1231222021030232-2232222102321031-1102110132322013-0131103131333330) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-0202201123102032-2012020231303321-0232310331000222-1122303003120220-3132000330130303-2122002221323133-0121133212102020-3232021322312023) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-1001302302332331-1003132212132213-3113212310300300-3323123332202212-0222132032110120-2020113330002011-1131202002123332-3122013010322020) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-0202320301012200-1233101231000012-1210113321331312-2233132332031222-3201103030121131-1033130033022333-1011102102123330-0313230322231131) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-2321003010103230-0022132023113301-0223003130000003-2233303200102023-1010212023031021-3231131021320012-2100021120302300-3101203320210110) |
| `dedicated_interface.node` | [dedicated_interface.node](data-sources--network_interface--reference--group-001.md#canonical-1321300022322331-1002221100200311-1332233022033321-3100222210333223-3103300021020023-0233213030230311-1211033013332333-3303111110220320) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-1110302013001312-0211111010001230-3230020030011312-0231320331102230-3000131022230222-1223223313302233-0123012123323301-0323232003010000) |
| `dedicated_interface.priority` | [dedicated_interface.priority](data-sources--network_interface--reference--group-001.md#canonical-2010323000320110-3201130210013112-1202002130233213-1010333100013112-3221102013103120-3232320023320023-0301030121222233-2211032113133120) |
| `dedicated_management_interface` | [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-3101133230301021-1301002312033022-0023111132213123-2132231010101102-0322332231113123-1031303333011300-0310111103301120-0123201333330121) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-0303230121133202-0320322110103122-2022230011323123-1032313312330201-3022301320300221-3130132011303302-3132320100113013-3002312112211232) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](data-sources--network_interface--reference--group-001.md#canonical-2011232323321203-3132201102330001-3002012100000200-0011230101010020-3103320220012321-0013032201011031-2111201333320130-1200210031110110) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-1200320013012221-0101333320011000-0302003122133300-3132031221211033-0233021012313330-0223331213130102-0331023300022333-1302002110330211) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](data-sources--network_interface--reference--group-001.md#canonical-0130033312333112-1003233203333232-1110021132111001-0111100321201201-1203021023212020-0211031322232010-3321213020030231-2121203113101112) |
| `description` | [description](data-sources--network_interface--reference--group-001.md#canonical-3330131101312002-0000110202332320-2333113001120202-0203302333113031-2133110100023020-1030103021002123-0111330232023220-1112003110103333) |
| `ethernet_interface` | [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-3231211313330313-0022132220201221-0013000032323022-3300222221132203-2012302133001301-0323102322100320-0311033002022030-2102121202121010) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-0200013332221021-3003113022031200-0231212300131133-3223223322231220-0320331200311330-3331102301200100-0221323011210322-3112211032320202) |
| `ethernet_interface.device` | [ethernet_interface.device](data-sources--network_interface--reference--group-001.md#canonical-1101112103132201-3113333302321232-3013202220332111-3020210321133021-3333130033211033-0320200120120112-2013011300300111-3233330202321002) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-2322032301300010-1302320200133131-3302013213312130-2113012210320311-2031030103203202-1313313023032331-3031032002330223-1120232033113002) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-0201002310300232-3201102100230313-1013113121133313-1002120131233210-2200101003320132-0313122302213003-2230001313232210-1322012022302131) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-3320130022001202-0322311233012202-3211222321303230-1321022313333023-2033100022332001-2331301220332221-3132230323023021-3001311203022111) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-3313323210320100-1033100311222113-1113131013100022-3313221121213102-3223221112210202-0201233103302131-0213011233301203-3311222013022232) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-2302031210123303-0301220333223021-1031020220013221-1001001112123012-2232002122312322-2122102200232022-3331223310030112-3320321312121021) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--network_interface--reference--group-001.md#canonical-0330033301012221-3221003222032110-0020103002103322-2232130022033313-1231302113020220-3102012022122233-3302213121012002-1223122113110101) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--network_interface--reference--group-001.md#canonical-1021210132330030-0320113001012320-1220120231301102-0102203312021212-3022011132000003-2332303103220331-3301313132321121-1112331000202021) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--network_interface--reference--group-001.md#canonical-0200132021100203-1310210311200020-0330021111212120-1213112102030202-0330022000303333-3113013312111032-3233032200212201-0300223203023101) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--network_interface--reference--group-001.md#canonical-0230121010121220-2030230331123032-3201021101202202-3111212011310120-0003132201313232-3233032010022211-3212121312130030-2311100130023000) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-2001321001030233-0122003322031001-0200120312223200-0120132110330013-3301233122012222-1010302033231120-2100211203300000-2323002330011010) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--network_interface--reference--group-001.md#canonical-3000332012101113-0310003303101112-0023203232201001-3101232303332300-3201112010220023-0020103120211331-2221232133210113-3302102333021230) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-2011003312200012-1201013310023133-1120132313210101-2110232003220011-3210223021202011-0203001003220012-0120133112303312-1302300200211001) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--network_interface--reference--group-001.md#canonical-2100101213030320-0311120222231300-2011310102333322-0132221210330101-0333200313212321-3203201201322331-3100112312301233-2010200100301001) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--network_interface--reference--group-001.md#canonical-0211023133321221-2033230201031013-2210023321212233-2133122323011000-0232111133201212-1133231122122110-2020023320302031-1121103013212132) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--network_interface--reference--group-001.md#canonical-2330223010221132-2030012313002103-3103032110330030-1200123012221222-0110303112002222-2321330220313300-1031032313021331-3131021300103103) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-0000120110111221-0022213312102202-1020130213213032-1131202021220332-3001212133020331-1022233031110300-0201033130233202-0103320000233101) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--network_interface--reference--group-001.md#canonical-2200220103202222-2203100112023232-2023010021123332-1122312322011103-1303101232311023-2133321230301212-0201020323010023-2322120023123220) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0003032121012210-3322130123001021-3100233202223310-3302021011020031-1312213010302300-2101200112221202-3121113002322012-3331312112033012) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-3312202101330131-0022223121220211-3011023330021010-2031231103002023-2100032301213011-3023101231221211-0200002300123003-0110231020332103) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0020031131211031-3021331303300100-0100113022312110-2201310010213302-1132310111111311-1211231310310231-2322102300202103-2302023323122310) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-0302020210100300-3221301212010300-0120300101011001-0003332121311303-0330322202212003-0311013102130210-3321211000031230-3323302313202013) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--reference--group-001.md#canonical-3331221110103103-3201212303332230-0333010223210202-0103302132212210-0000000110322111-2323100313100101-2110312103123131-3111330333323302) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-3201033233203023-3322121320002122-2213313113131111-1313212101013203-2301310100033123-3012210212001002-1102003312322302-0011213021011100) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-3300212221303020-1220310012310230-0232200220220213-1231033323321011-2123110220202003-1001110311310221-0203232112111020-3230003100312130) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--network_interface--reference--group-001.md#canonical-2211000233103000-1203032011311202-3012300201203210-0313111233300131-1313210003312010-1121330033031011-2233323031110321-2001330001210201) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--network_interface--reference--group-001.md#canonical-1110010201200130-0321312110032100-1312031113001313-1202233002223310-3110131200122103-3032212311311221-2233202000212032-3233100000220020) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-2203102230100313-0230121002311320-1123033200312321-0331200121121203-0123032333010000-3030210100000130-0032232233300222-3221212212320301) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--network_interface--reference--group-001.md#canonical-3311331101200121-1322231201333223-3232110101002322-2200222021222322-2223020010110213-3201212010302113-0303130013232203-3320032302133220) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--network_interface--reference--group-001.md#canonical-3010001321030301-0122231103201122-0230120221213002-0231122010112032-2223203333330132-1132102000331222-1023230112203230-3122201301123303) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--reference--group-001.md#canonical-0133213203213220-1230012020100121-2100232030010203-2120121220012201-2133313120113031-1300223032230310-3131100033332201-1222213302121200) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-0022001133123231-3230133323121222-3113303122032123-3200122003113110-2231033203123223-0020322110002330-2300311110103203-0121303012231312) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-0211210133333101-1212112313030220-0013000103333121-0330213320123022-0012330113323220-3221032022321212-3122213131011121-2113011020023331) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-3212233301121232-2130123311011332-3100031100222122-2321300013113032-0121030012233301-3021101113011300-2322121331111003-1010220110022310) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-3133223121002201-2331311321300100-2131011300122101-2232030221102101-3212132132301102-2112023122023002-3211310212131030-3313111233322311) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0211113313111232-3031222211130022-1202100220220211-2302022031021232-2330231132133233-0032020311301210-0231310001331232-1021213223213000) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-2220320031213230-3230003202122201-3223322330203120-3023322233012200-2011003221000122-1330332233223032-2321232302220013-2103032221313222) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--network_interface--reference--group-001.md#canonical-1121132332210102-1110301102220113-0221223330332330-1203113331133113-0100331110203113-1213202100020221-0002111013200103-1023023202333001) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-0033220313221110-3021213303323113-1020220020332113-3111021021033033-2110132022031133-2202113202220222-1000123111301023-0022223120013302) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--network_interface--reference--group-001.md#canonical-1120211201132033-0120203011322001-2321100002301310-3310331233211120-1002231002001220-2222123313222310-0322033231023102-0123133231200212) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--network_interface--reference--group-001.md#canonical-3211031130330112-1230233331033100-2032230211223123-3312223311013211-2323001020222223-0023223113003010-2230013302012032-1312112113031010) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0202210103210122-3111110112300232-0221230111301331-0120222112303133-0232113312333000-0031011330211022-1112032233321331-3233312210210320) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0110112110023301-1331211011310003-2200302313102001-1302100331033111-1002011202200203-3321301220322303-0021310222130033-2100102311020010) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-2020322230121201-3210301200123321-0332220333133332-1102011212111332-3033020013103203-2110123310212220-0122133022023303-3001303030222001) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-0213021100100001-0330331013011332-1301300210130322-2100302201203120-2333131111122332-0321010120202132-2110223133113210-3222133330102032) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-1033313201201102-1202031310100130-2200001133331030-1000023213110002-0312220033120210-2301122322130010-3112000033013131-3110200002213011) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-1233112322031330-3011130002322221-0222223321331223-0210122323213220-2313020113230300-3112300233331000-1233113322320200-3102131010331130) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-3022021211012011-1230021012001233-0131113132130013-1002000211011001-2102210212012323-1020213131031033-1000210022301133-1032322223131203) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-3310222213211122-3121201300312013-1031032213131200-1110202122301030-0013031232130220-1022323122000020-0201132201331330-3130312311231002) |
| `ethernet_interface.node` | [ethernet_interface.node](data-sources--network_interface--reference--group-001.md#canonical-0302021222022010-2311303133323210-1203123030111010-0300331331103201-1110001110302331-1033122133310022-3021310310213122-3211110010232333) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-2231103231122332-0203212220131032-0332002321021100-0002002031012132-1310030202110020-0322100203202133-0001002210002211-2202103002010330) |
| `ethernet_interface.priority` | [ethernet_interface.priority](data-sources--network_interface--reference--group-001.md#canonical-1300000122323121-1301300303231132-0301001111230011-1012102100113100-1331200011220013-3331203322303331-1311210210203132-2202132223003310) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-3112330210230232-0131201333110111-3323130032020110-3103031120333133-0333111030221031-0100321022123313-2032200230112312-0103130122113323) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](data-sources--network_interface--reference--group-001.md#canonical-1203322000313333-3321000133230210-3020130120303110-1021033120123202-3131132100110100-0002133230231110-2013002233012320-0223232323231110) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3300133011111203-1330212301220211-2303202300012231-1111013302212331-2002000121213230-2332030021000223-2300131012003112-0110020231130323) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1031130201223312-1002010011231021-1023131120012231-3312313020123003-1013130302310212-2111103023022103-3302133032320000-1212322300221320) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-3010023201022000-0002103200211213-1131101320013211-0230300231230300-1111111333312033-0002202002313212-1232103100112123-0320213033110323) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3220110030331230-1032002202032023-2333010201201130-3203233103311303-1132321012031032-0130311200010112-3311001102201303-0202012112021300) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--reference--group-002.md#canonical-3110022012211100-1000330221221332-2010203022002301-1112121212213000-1132203221022123-3303102022113022-0330020230001310-3230312302113303) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--reference--group-002.md#canonical-1130111221303111-0002310333300023-1331131011020111-3233233323100022-3310131131300331-2030130313023332-3320222132110323-3110030012310013) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--reference--group-002.md#canonical-2103113210201222-0011033303212021-2102120302002230-2131021130330311-2302320233322121-2203212323322031-1203033120323103-2300131000300201) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2012012100233103-0020302021020220-1021110301020301-0211320203011110-3033201122200101-2033021031201011-3311223230020021-2323030221003131) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2131010032210103-0301310100010222-3331312032212000-3120323223031120-0001221020222220-0122032103310201-3113330222323322-2203011002101131) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-2211030020312330-1133122213031302-0120020203213133-3223132310330312-0223122211320132-3130022030310231-3330103113121130-0203223322321220) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1032002110320220-0320210300113000-2123213232011231-3102030233010020-2100200122333303-0212032303210331-1212023010232210-0111211101203020) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--network_interface--reference--group-002.md#canonical-2000211102002202-2121133233233201-2131201331212102-1120013231111300-2332320211210120-2320130111311010-2102313123202221-3202002211323302) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--network_interface--reference--group-002.md#canonical-2003223203303030-0220312301313001-1310212102112201-0302012020010100-3100231221010200-2312133032331302-3133130011321220-3213220212301133) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--network_interface--reference--group-002.md#canonical-0203013002223113-0201100122031301-0120300221012000-0202031322203110-0302200300032231-2003330001110313-3113332222230321-2321221020013201) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](data-sources--network_interface--reference--group-002.md#canonical-2222221302222022-1230033212101101-1222323331022000-0130003210020222-0102133221030111-3331300022113330-0010133020123323-3310123231323310) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-3332202232033323-3023222020030011-3230200311123331-3021012122032332-0131010220030001-1332031123233013-3120333302302030-1101110130010102) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](data-sources--network_interface--reference--group-001.md#canonical-3013023100320231-1122010101010013-2212330311302303-3032203312221011-2010010201010210-0330000331113012-3112223113301311-3303302212323101) |
| `id` | [id](data-sources--network_interface--reference--group-001.md#canonical-1310210210213110-0000112020200133-1023333321001021-0200022320130003-2330123011023011-1331003200133311-0201021310122332-3110211213113232) |
| `labels` | [labels](data-sources--network_interface--reference--group-001.md#canonical-3300321303201002-1102201313323311-2131100200012330-2231330233132300-0323122313013300-0012110210023300-0023133323333332-0220202333003010) |
| `layer2_interface` | [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-3202113010230203-3202320132300102-1102221210032110-3313003010222230-0232130003102200-0100010011021322-1003230323121233-0033100200303223) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-3132200021010003-1031232001333112-3313022332200113-1303021311010002-2023210112130201-2020211011321002-2333213110330102-2203023111202012) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](data-sources--network_interface--reference--group-002.md#canonical-1112033311201310-1131020021213012-0100012132231000-0130010112313332-0031102113231231-3321011302331130-3330230122113322-3320322303010203) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-2011230012310221-1233300113013111-1133230331120100-2233131231222101-0121332213122112-3213032120301030-1102122303112302-0002001231033222) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-1330333313200301-1210020213320221-2332100321130123-3101002311233003-2121222133120011-2221121223301102-1323010113313302-0213333111232100) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-0232122010331133-3302132312333320-1102001130210213-1032031132220333-3011020220212311-3331102000223133-1021122021012133-1332033031100210) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](data-sources--network_interface--reference--group-002.md#canonical-0113023300301001-0130030320211111-0003311210221212-3211123101010302-2210122010030313-3203111331033322-3002200201331222-1230123323211000) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-0022022123232202-2000212030132302-0200123333331320-3212123321320020-0230211320333102-2321210321033112-1332212100230022-3101123223010331) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-2213122321301030-2303330210333230-3330100000120121-2302332212012100-0303020220331233-2032310200101322-3000320320133331-0303110002021132) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](data-sources--network_interface--reference--group-002.md#canonical-2113113000333321-2320012032112303-3133231003003030-0331223223212110-1022012222203300-2220123202300201-0222320222200022-0332302011311111) |
| `name` | [name](data-sources--network_interface--reference--group-001.md#canonical-0110322303330303-3213131003322113-1133230103212113-3123331032212132-3032202223321313-3012203032223023-0311032131322202-0311010122021220) |
| `namespace` | [namespace](data-sources--network_interface--reference--group-001.md#canonical-0031200023311012-0003321231121010-2022201013012221-1123320023231132-0102221111322012-1011323133021122-3322011130133232-1023331333000122) |
| `tunnel_interface` | [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-0202122131102011-3232213112222323-2223112333132222-1101010011000212-3201033332312313-1112311332113213-3310013212203003-0200303300303233) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](data-sources--network_interface--reference--group-002.md#canonical-2301122233103120-3112112022331222-0231123032211032-3110312111020213-2203203103133233-0200001103313123-2101303013211020-2000122113310232) |
| `tunnel_interface.node` | [tunnel_interface.node](data-sources--network_interface--reference--group-002.md#canonical-2233332222021020-0212203121222212-1201333202121302-3322330033213101-2223000322030020-2113221132000223-2321300020302312-1312133313133201) |
| `tunnel_interface.priority` | [tunnel_interface.priority](data-sources--network_interface--reference--group-002.md#canonical-2122101203312001-2323121303032301-1322312322302221-3322311331232333-0230320212003212-0201120000322321-2120111221233331-0313003222333233) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-3131333032013223-1021201131223103-3121100210012223-3003312331101032-3123131112323102-3222200300101301-2233222321233032-0223102001110123) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](data-sources--network_interface--reference--group-002.md#canonical-1000312200201102-3330320200321211-2212132111202100-3030330320110133-1111232132221101-0133103232233200-0112330201110013-1031001312110121) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3113112132103132-3203032102300021-0103202132212110-3322112313010211-0122300231201223-3310202132113311-2301231320023232-2332012333313130) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0321110100210030-1113001033321023-0210230320101121-1321230122332123-3011201332031200-1012330232211212-0022111022033113-3021201120332210) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-3101130211022311-1023301213212323-2130213331001301-0033121110020201-3100330113310312-2123012113023310-0011222131002331-1120131102032133) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1311332322200012-2123002113110032-1012212100031111-2220101010130003-1001110031223110-1022310030023002-3103222230003121-2202101023300133) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--reference--group-002.md#canonical-0030132022001213-1330101023110212-3111001212032032-0312030032331330-3232323222212132-3122012331303112-3131301211130202-0021330312203021) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--reference--group-002.md#canonical-2303121121210330-1303230133303032-1230031333121233-3023220211332103-0310001121230032-3303030210323320-0303001222030121-2031313201031323) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--reference--group-002.md#canonical-1223020222311233-0203332000321022-1122103323030102-2020211220110101-3213000310321132-3103311321100021-0111220100031330-3112033020100120) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](data-sources--network_interface--reference--group-002.md#canonical-0313310202220321-3100221300010201-0322011131122303-1332000011200001-0312032200123323-0111102223233001-3011312020223322-0022010333122233) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](data-sources--network_interface--reference--group-002.md#canonical-2231123023210121-3020002101303113-1032320312300323-0312120110111231-3213200002003223-2122202203203130-2033222030312013-2010033111001111) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](data-sources--network_interface--reference--group-002.md#canonical-2130220312101112-0223322103302131-1212020212132213-0031212331230103-1300202103021033-3113231311210031-2120332213101003-0033303310203330) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](data-sources--network_interface--reference--group-002.md#canonical-3011221000031031-2331131120023320-0002311032110300-3023121111001333-3120310103122023-3102131011323302-0032032203210313-2121012013321331) |

<a id="canonical-1033122233221321-1301133320301113-1122121203301013-3300313132203122-2103131201003221-1311112133203333-1200301300330221-2032210023103203"></a>

## Next pages — Property reference / 000001200001 / 11

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332333123330222-0030022111120323-0302313032100021-1101301003110213-0222321022220023-1221222222110103-1220132312033310-1030220003203203"></a>

## dedicated_interface — dedicated_interface / 033122303113 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- dedicated_interface

<a id="canonical-2313131122020221-0233301323112002-0302000311302211-3112301213001122-3301001203003313-3122131012332030-3221323232202001-3330300001121330"></a>

Type: `"single"`. Computed.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

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

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2313131122020221-0233301323112002-0302000311302211-3112301213001122-3301001203003313-3122131012332030-3221323232202001-3330300001121330)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-3101133230301021-1301002312033022-0023111132213123-2132231010101102-0322332231113123-1031303333011300-0310111103301120-0123201333330121)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-3231211313330313-0022132220201221-0013000032323022-3300222221132203-2012302133001301-0323102322100320-0311033002022030-2102121202121010)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-3202113010230203-3202320132300102-1102221210032110-3313003010222230-0232130003102200-0100010011021322-1003230323121233-0033100200303223)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-0202122131102011-3232213112222323-2223112333132222-1101010011000212-3201033332312313-1112311332113213-3310013212203003-0200303300303233)

Select alternatives according to the provider validators above.

<a id="canonical-0033000320122031-0213111310000311-0022232100233310-0312123101322012-2110212121132202-2011323303110010-2132002222103002-0300111310131213"></a>

## Direct properties — dedicated_interface / 033122303113 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1211132330113332-3212031033012102-0222101221133332-0311302120203213-2032132300102321-3222013032130332-1210023202332201-1100023032023022): complete subsection reference.

<a id="canonical-0111130032120211-0202231123231000-3310100012301202-2101312100031211-1231222021030232-2232222102321031-1102110132322013-0131103131333330"></a>

<a id="canonical-3120200103212133-1203202101023332-3112222222330201-3001223022000133-2331113132133213-2030202032232230-0322212220030003-1201212130320112"></a>

## device property — dedicated_interface / 033122303113 / 4

Type: `"string"`. Computed.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

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

- [is_primary](data-sources--network_interface--reference--group-001.md#canonical-0022131012331121-3101130212200322-2211202021211213-3113002032312120-1022122112303230-1223110011121212-3323100130212100-2330110111233331): complete subsection reference.

- [monitor](data-sources--network_interface--reference--group-001.md#canonical-1301300033020023-2330210123101223-3133232003002121-1100233103130100-1100331332330110-0203112300000220-2321033203313011-2310121123120321): complete subsection reference.

- [monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-0030313302121202-1233123111231332-3332201001311232-1022220000121011-2130021022031232-3321001222222110-1130010012122033-2033000003310300): complete subsection reference.

<a id="canonical-2321003010103230-0022132023113301-0223003130000003-2233303200102023-1010212023031021-3231131021320012-2100021120302300-3101203320210110"></a>

<a id="canonical-1321020223212320-0033000002122331-2331330320021312-3210121332330203-2113220020012100-3123021221333112-3103213031000032-3312212033303330"></a>

## mtu property — dedicated_interface / 033122303113 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-1321300022322331-1002221100200311-1332233022033321-3100222210333223-3103300021020023-0233213030230311-1211033013332333-3303111110220320"></a>

<a id="canonical-2002103021200003-2011232213123101-0113133310033302-0303300112032013-1222001200222101-1131323132202312-1313311313212020-3203101330120230"></a>

## node property — dedicated_interface / 033122303113 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [not_primary](data-sources--network_interface--reference--group-001.md#canonical-2101232323003203-3203010310132201-2231031113103111-2121012100323012-1302210102202103-2001313101033321-2133320212012010-3330031331120011): complete subsection reference.

<a id="canonical-2010323000320110-3201130210013112-1202002130233213-1010333100013112-3221102013103120-3232320023320023-0301030121222233-2211032113133120"></a>

<a id="canonical-2011300330032310-1320312113101011-0010021033021020-2031310132330002-2003110332012013-1202212322113021-3202111001321330-3112102301210100"></a>

## priority property — dedicated_interface / 033122303113 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331012113321120-3003230231111122-2003221220031332-1330101001332122-1222211132323111-2000212321333323-3022103333122003-2230230000023100"></a>

## Next pages — dedicated_interface / 033122303113 / 8

- [dedicated_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-1211132330113332-3212031033012102-0222101221133332-0311302120203213-2032132300102321-3222013032130332-1210023202332201-1100023032023022)
- [dedicated_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-0022131012331121-3101130212200322-2211202021211213-3113002032312120-1022122112303230-1223110011121212-3323100130212100-2330110111233331)
- [dedicated_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-1301300033020023-2330210123101223-3133232003002121-1100233103130100-1100331332330110-0203112300000220-2321033203313011-2310121123120321)
- [dedicated_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-0030313302121202-1233123111231332-3332201001311232-1022220000121011-2130021022031232-3321001222222110-1130010012122033-2033000003310300)
- [dedicated_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-2101232323003203-3203010310132201-2231031113103111-2121012100323012-1302210102202103-2001313101033321-2133320212012010-3330031331120011)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1211132330113332-3212031033012102-0222101221133332-0311302120203213-2032132300102321-3222013032130332-1210023202332201-1100023032023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032130003023230-3002233311021210-1111211003101331-1011122330231230-0330000011233332-0303031121301022-3220221133313123-3101121301031132"></a>

## dedicated_interface.cluster — cluster / 300302303021 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.cluster

<a id="canonical-0120233122333101-0022201212300031-0001202002111303-2003020000003100-1021011212112301-3020301023301103-2002312031220103-1303120003111133"></a>

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

<a id="canonical-2000211113132202-3303302323231001-2112231322222031-1000203313202012-2013310330112223-2302300232302023-1133303212230010-0030221121133012"></a>

## Direct properties — cluster / 300302303021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120222330221321-3232220121103030-1311102103312231-2010320201110120-1330202101113331-2232210130020212-1131201312132313-2331112310102101"></a>

## Next pages — cluster / 300302303021 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0022131012331121-3101130212200322-2211202021211213-3113002032312120-1022122112303230-1223110011121212-3323100130212100-2330110111233331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133302311123322-3302300222003021-0321013012222121-2020113200120213-3330122312131103-3300213012000303-1021310102023111-1032132021030032"></a>

## dedicated_interface.is_primary — is_primary / 022031212113 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.is_primary

<a id="canonical-0202201123102032-2012020231303321-0232310331000222-1122303003120220-3132000330130303-2122002221323133-0121133212102020-3232021322312023"></a>

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

<a id="canonical-3303332223220123-1020120332231103-1110130010023010-1210032120011210-1211001132100332-3300023201223102-2212123321333011-1231133001222123"></a>

## Direct properties — is_primary / 022031212113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222211301012133-2311003123203320-1101213203130003-3000000021200300-1200220102120113-2200030301333211-1103110123221312-0111101200012210"></a>

## Next pages — is_primary / 022031212113 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1301300033020023-2330210123101223-3133232003002121-1100233103130100-1100331332330110-0203112300000220-2321033203313011-2310121123120321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131310122201323-3030300120110133-2303200102320031-3212232012003331-3031130230023021-3030133110321121-1332012132233032-1230103022131210"></a>

## dedicated_interface.monitor — monitor / 123003103023 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.monitor

<a id="canonical-1001302302332331-1003132212132213-3113212310300300-3323123332202212-0222132032110120-2020113330002011-1131202002123332-3122013010322020"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1102110133210323-2002331332003321-1032323013023032-0302101113103010-3121001122003230-0303201002322212-3021211002331110-3030121313333002"></a>

## Direct properties — monitor / 123003103023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232021320023013-1210313200133220-1203300312101200-1003303210102313-1120333203231111-1200330210220212-2111301200323312-2112002103131132"></a>

## Next pages — monitor / 123003103023 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0030313302121202-1233123111231332-3332201001311232-1022220000121011-2130021022031232-3321001222222110-1130010012122033-2033000003310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202201312210212-2323032112123001-2201101113010321-3221213300312122-2211222322133130-0103310103003011-2331301322311312-1231113122211311"></a>

## dedicated_interface.monitor_disabled — monitor_disabled / 202333200021 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.monitor_disabled

<a id="canonical-0202320301012200-1233101231000012-1210113321331312-2233132332031222-3201103030121131-1033130033022333-1011102102123330-0313230322231131"></a>

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

<a id="canonical-2102300031020210-3333230303002021-3231332230230131-0010220213331331-3300013312021300-3302010211021200-2022000113111111-1000221033221021"></a>

## Direct properties — monitor_disabled / 202333200021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200231102303112-2231020223333021-1022201222230313-3010131001133232-1000110012111023-1220211200210220-3112032001313231-2123001123000102"></a>

## Next pages — monitor_disabled / 202333200021 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2101232323003203-3203010310132201-2231031113103111-2121012100323012-1302210102202103-2001313101033321-2133320212012010-3330031331120011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013123131221112-2111030023331323-2100331000110223-3201033101022031-2121100332210011-2312111301131311-3000121110200233-2330302200333231"></a>

## dedicated_interface.not_primary — not_primary / 000013101301 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.not_primary

<a id="canonical-1110302013001312-0211111010001230-3230020030011312-0231320331102230-3000131022230222-1223223313302233-0123012123323301-0323232003010000"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1311121111310223-3031000103030213-1312031113222021-1020002233113022-1022230313200232-1321210000310023-0332121211130232-2122303130010231"></a>

## Direct properties — not_primary / 000013101301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220112211331121-2202111003322103-3030003110310220-2000211323311001-2321102003230220-0000101002120311-0233133101300021-0101000111201032"></a>

## Next pages — not_primary / 000013101301 / 4

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002231100101331-0033120002333111-3021333320011012-1003110012121012-2011002120123203-0230210200323210-2112130130012101-3310033211211300"></a>

## dedicated_management_interface — dedicated_management_interface / 103232102200 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- dedicated_management_interface

<a id="canonical-3101133230301021-1301002312033022-0023111132213123-2132231010101102-0322332231113123-1031303333011300-0310111103301120-0123201333330121"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

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

<a id="canonical-2103211023210221-3013022123320201-2233112300121332-0220000210101223-0322003222311321-3320000003323030-2200113100311200-2002133101131302"></a>

## Direct properties — dedicated_management_interface / 103232102200 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1121030012320221-0012303030303003-3203112120202321-3003221000313133-3022112200032123-3312331333222030-3212201010232230-3221200033002033): complete subsection reference.

<a id="canonical-2011232323321203-3132201102330001-3002012100000200-0011230101010020-3103320220012321-0013032201011031-2111201333320130-1200210031110110"></a>

<a id="canonical-0320330023320032-2312033210300232-3000320300110222-1322303221323132-2222110330131321-3200311303103031-1113110233202021-3201020302002300"></a>

## device property — dedicated_management_interface / 103232102200 / 4

Type: `"string"`. Computed.

Name of the device for which interface is configured.

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

<a id="canonical-1200320013012221-0101333320011000-0302003122133300-3132031221211033-0233021012313330-0223331213130102-0331023300022333-1302002110330211"></a>

<a id="canonical-2213123233131233-0130233331313131-2333010302202322-2100130120330303-0002333032210101-0133012330331031-0321320320130010-3023211032001102"></a>

## mtu property — dedicated_management_interface / 103232102200 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-0130033312333112-1003233203333232-1110021132111001-0111100321201201-1203021023212020-0211031322232010-3321213020030231-2121203113101112"></a>

<a id="canonical-0000100022331322-0130021311112120-2023221333322023-3010101133221003-0021031123321100-3002331312310031-0030033212111032-3203332322231021"></a>

## node property — dedicated_management_interface / 103232102200 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0020113330110023-3011001322303223-0033320303121033-3013332232302332-0033003221130300-1131333323123231-2022223031110102-2223333020132230"></a>

## Next pages — dedicated_management_interface / 103232102200 / 7

- [dedicated_management_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-1121030012320221-0012303030303003-3203112120202321-3003221000313133-3022112200032123-3312331333222030-3212201010232230-3221200033002033)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1121030012320221-0012303030303003-3203112120202321-3003221000313133-3022112200032123-3312331333222030-3212201010232230-3221200033002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131333313321232-0133313200033003-1002021322223130-2002320102000123-3310202311101033-0320002031313221-0331002211132312-3131100001202333"></a>

## dedicated_management_interface.cluster — cluster / 020110201020 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302)
- dedicated_management_interface.cluster

<a id="canonical-0303230121133202-0320322110103122-2022230011323123-1032313312330201-3022301320300221-3130132011303302-3132320100113013-3002312112211232"></a>

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

<a id="canonical-3131122333112112-0033023331000233-3213003333311113-0120123313012312-1123230113000131-2122222332302111-3311202333031013-1322133232100112"></a>

## Direct properties — cluster / 020110201020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223010100201231-0223231110113113-0102321000321122-0023332332233022-1001202312003000-3320000201323122-0200232211230123-2331302133331333"></a>

## Next pages — cluster / 020110201020 / 4

- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221232332012323-3233303223203312-0201030020333211-0322300202330001-1321332232332110-2002103011223133-3021330133011133-3321211031120313"></a>

## ethernet_interface — ethernet_interface / 221132323302 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- ethernet_interface

<a id="canonical-3231211313330313-0022132220201221-0013000032323022-3300222221132203-2012302133001301-0323102322100320-0311033002022030-2102121202121010"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

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

<a id="canonical-2211312313032133-3131003111022130-1200033311200233-2020223303113232-3023300320310200-1300231123212302-1120310021210112-3002311223311222"></a>

## Direct properties — ethernet_interface / 221132323302 / 3

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1000310232301022-0020330010030121-0022102133110121-2323122310112011-0022100321232112-1021210323112113-2211012232330000-3121002333213032): complete subsection reference.

<a id="canonical-1101112103132201-3113333302321232-3013202220332111-3020210321133021-3333130033211033-0320200120120112-2013011300300111-3233330202321002"></a>

<a id="canonical-0000313022032102-2022121033120120-2101222130021322-0110321011010303-3211221323220023-1033132333033210-2223231210213113-0301213210133123"></a>

## device property — ethernet_interface / 221132323302 / 4

Type: `"string"`. Computed.

Interface configuration for the ethernet device.

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

- [dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-2111203310333301-0203133201310122-3302311102203221-0033033020213331-3321331221303031-1022320231103221-0030211110001111-3102232032320223): complete subsection reference.

- [dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033): complete subsection reference.

- [ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001): complete subsection reference.

- [is_primary](data-sources--network_interface--reference--group-001.md#canonical-0131322022011012-2232101232210032-1121023033132330-2102122010330022-0303011212231321-0030120321323012-0002332223030223-2331132203033201): complete subsection reference.

- [monitor](data-sources--network_interface--reference--group-001.md#canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201): complete subsection reference.

- [monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-3121330013003013-2100012130332222-1133031333201222-2021321321232111-2002301003121210-2322111332332230-0221011000030231-2023233203032113): complete subsection reference.

<a id="canonical-3022021211012011-1230021012001233-0131113132130013-1002000211011001-2102210212012323-1020213131031033-1000210022301133-1032322223131203"></a>

<a id="canonical-2201332110311011-3133030021302221-2101233212331212-0301300113201122-1230232212310333-2203312203202122-0200101201301001-2212023113332210"></a>

## mtu property — ethernet_interface / 221132323302 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-1113032321321001-2310200133001032-3213200220203223-3000220113031202-1230221312200012-1202311001030122-0321131000121000-1111313323032111): complete subsection reference.

<a id="canonical-0302021222022010-2311303133323210-1203123030111010-0300331331103201-1110001110302331-1033122133310022-3021310310213122-3211110010232333"></a>

<a id="canonical-0223203133212013-0021312030322102-1220032213132331-2230022011221300-2330323002310001-0320301120301000-3110213232012130-1223003230203111"></a>

## node property — ethernet_interface / 221132323302 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [not_primary](data-sources--network_interface--reference--group-001.md#canonical-3312231302230300-0022030322113100-2201123131132030-3100011010120320-3310232100021031-1310303103110322-1031013222131303-3130102302313330): complete subsection reference.

<a id="canonical-1300000122323121-1301300303231132-0301001111230011-1012102100113100-1331200011220013-3331203322303331-1311210210203132-2202132223003310"></a>

<a id="canonical-1231022221233020-3203331000312103-3111233020333022-2112133312102230-3202212103331330-1013130020220222-2231301211203202-1032021233221203"></a>

## priority property — ethernet_interface / 221132323302 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-0212312303120132-1113000211222110-0213011322231122-3213230332311203-0133301012122321-3102222301230020-2232303202233323-3312131000300010): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-001.md#canonical-3132210030013112-1102213013013232-0321001230003132-1320121232330023-0311112012212022-2113030330100121-0300331333001210-3033033130201021): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231): complete subsection reference.

- [static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310): complete subsection reference.

- [storage_network](data-sources--network_interface--reference--group-002.md#canonical-0333033210010321-2132211112223013-1230003210301302-3330133302221210-0113332030013030-1331320233211220-3001013200001121-2203312010010122): complete subsection reference.

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-2303100002223300-2211310120020011-3133100223301321-2323322300103112-1230330031231133-3230103021330133-1023322211311012-0001221012231312): complete subsection reference.

<a id="canonical-3013023100320231-1122010101010013-2212330311302303-3032203312221011-2010010201010210-0330000331113012-3112223113301311-3303302212323101"></a>

<a id="canonical-3023200203323321-1312023021112121-1332002310131312-0231332233302112-3212211100200030-3322311120331320-1223331103212223-3011021221323201"></a>

## vlan_id property — ethernet_interface / 221132323302 / 8

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311212032012320-0020120013321310-1130002000213122-0010220301211203-0210000020312103-3011000022310332-0330023230020202-0130102211113013"></a>

## Next pages — ethernet_interface / 221132323302 / 9

- [ethernet_interface.cluster](data-sources--network_interface--reference--group-001.md#canonical-1000310232301022-0020330010030121-0022102133110121-2323122310112011-0022100321232112-1021210323112113-2211012232330000-3121002333213032)
- [ethernet_interface.dhcp_client](data-sources--network_interface--reference--group-001.md#canonical-2111203310333301-0203133201310122-3302311102203221-0033033020213331-3321331221303031-1022320231103221-0030211110001111-3102232032320223)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.is_primary](data-sources--network_interface--reference--group-001.md#canonical-0131322022011012-2232101232210032-1121023033132330-2102122010330022-0303011212231321-0030120321323012-0002332223030223-2331132203033201)
- [ethernet_interface.monitor](data-sources--network_interface--reference--group-001.md#canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201)
- [ethernet_interface.monitor_disabled](data-sources--network_interface--reference--group-001.md#canonical-3121330013003013-2100012130332222-1133031333201222-2021321321232111-2002301003121210-2322111332332230-0221011000030231-2023233203032113)
- [ethernet_interface.no_ipv6_address](data-sources--network_interface--reference--group-001.md#canonical-1113032321321001-2310200133001032-3213200220203223-3000220113031202-1230221312200012-1202311001030122-0321131000121000-1111313323032111)
- [ethernet_interface.not_primary](data-sources--network_interface--reference--group-001.md#canonical-3312231302230300-0022030322113100-2201123131132030-3100011010120320-3310232100021031-1310303103110322-1031013222131303-3130102302313330)
- [ethernet_interface.site_local_inside_network](data-sources--network_interface--reference--group-001.md#canonical-0212312303120132-1113000211222110-0213011322231122-3213230332311203-0133301012122321-3102222301230020-2232303202233323-3312131000300010)
- [ethernet_interface.site_local_network](data-sources--network_interface--reference--group-001.md#canonical-3132210030013112-1102213013013232-0321001230003132-1320121232330023-0311112012212022-2113030330100121-0300331333001210-3033033130201021)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- [ethernet_interface.storage_network](data-sources--network_interface--reference--group-002.md#canonical-0333033210010321-2132211112223013-1230003210301302-3330133302221210-0113332030013030-1331320233211220-3001013200001121-2203312010010122)
- [ethernet_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-2303100002223300-2211310120020011-3133100223301321-2323322300103112-1230330031231133-3230103021330133-1023322211311012-0001221012231312)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1000310232301022-0020330010030121-0022102133110121-2323122310112011-0022100321232112-1021210323112113-2211012232330000-3121002333213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102330123012103-0110300313123032-3222101030303002-0333323102231103-0100323001023230-3230010022311002-3120312332330303-3002021033213202"></a>

## ethernet_interface.cluster — cluster / 202002011000 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.cluster

<a id="canonical-0200013332221021-3003113022031200-0231212300131133-3223223322231220-0320331200311330-3331102301200100-0221323011210322-3112211032320202"></a>

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

<a id="canonical-2102011330212221-0000331201000313-3131200333132211-0113100013231133-3332232201301103-1322222310310011-3211302300033331-0203220031213201"></a>

## Direct properties — cluster / 202002011000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130130103330012-0100031231111330-0312323210222303-2323100203311112-3313333230000231-2211222131221002-3321212131213003-3130013001011220"></a>

## Next pages — cluster / 202002011000 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2111203310333301-0203133201310122-3302311102203221-0033033020213331-3321331221303031-1022320231103221-0030211110001111-3102232032320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113002321313300-1201312232103222-2230011213131231-3310001021011213-0023021302322323-0201111221332122-1200331213312002-0331130301300322"></a>

## ethernet_interface.dhcp_client — dhcp_client / 011032320103 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.dhcp_client

<a id="canonical-2322032301300010-1302320200133131-3302013213312130-2113012210320311-2031030103203202-1313313023032331-3031032002330223-1120232033113002"></a>

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

<a id="canonical-0002023211303323-3010223313332031-1311131310202102-1330020131312212-2230002213310203-0320121012202103-2003203301113303-2100123120002301"></a>

## Direct properties — dhcp_client / 011032320103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213021310012211-1123012001311213-0001103020232133-2131233201003302-0100332233212020-3123322200331203-2203131020301322-0213130221333320"></a>

## Next pages — dhcp_client / 011032320103 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023200022220030-2332020311121121-1121132222030100-0310131210002020-2000010213310323-3212302110130003-2000212212330001-1132210122212231"></a>

## ethernet_interface.dhcp_server — dhcp_server / 310200012012 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.dhcp_server

<a id="canonical-0201002310300232-3201102100230313-1013113121133313-1002120131233210-2200101003320132-0313122302213003-2230001313232210-1322012022302131"></a>

Type: `"single"`. Computed.

Configuration parameter for dhcp server.

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

<a id="canonical-0212030113120322-1122213202111112-0212232220031213-3202222123313303-0312302311303122-3222022211102212-2012232111203221-0220023300001122"></a>

## Direct properties — dhcp_server / 310200012012 / 3

- [automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0122220112022100-1110200101331310-3120100032020331-3000303033202113-3333320102233320-0230121232130021-3200320032100023-1132030210220311): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-1213132133310111-3012101232022020-1230330321131133-0221221213002212-2110020031120002-3003123011123212-1013322212033222-2110032231323132): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121): complete subsection reference.

<a id="canonical-2200220103202222-2203100112023232-2023010021123332-1122312322011103-1303101232311023-2133321230301212-0201020323010023-2322120023123220"></a>

<a id="canonical-2311200231011020-2100022211001233-1120033203021221-1120233113301230-2222300202033312-3230010123313213-0001210330032303-2113311211111322"></a>

## dhcp_option82_tag property — dhcp_server / 310200012012 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-0003032121012210-3322130123001021-3100233202223310-3302021011020031-1312213010302300-2101200112221202-3121113002322012-3331312112033012"></a>

<a id="canonical-0113212231230313-1330013013331221-1131203031110033-0010221000333231-1203101302333000-3002132210310333-2030330030203122-3121220131202203"></a>

## fixed_ip_map property — dhcp_server / 310200012012 / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0102031303123133-3210033231031111-3332001310200032-0302121200120133-0231111002210130-0021232021310111-0301232112000213-3003102103212232): complete subsection reference.

<a id="canonical-2020022100200132-2033121212120020-3123002311212303-1200131222211210-0230220211001312-1031112003030330-1101323311113132-1002303303310013"></a>

## Next pages — dhcp_server / 310200012012 / 6

- [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0122220112022100-1110200101331310-3120100032020331-3000303033202113-3333320102233320-0230121232130021-3200320032100023-1132030210220311)
- [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-1213132133310111-3012101232022020-1230330321131133-0221221213002212-2110020031120002-3003123011123212-1013322212033222-2110032231323132)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0102031303123133-3210033231031111-3332001310200032-0302121200120133-0231111002210130-0021232021310111-0301232112000213-3003102103212232)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0122220112022100-1110200101331310-3120100032020331-3000303033202113-3333320102233320-0230121232130021-3200320032100023-1132030210220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110222131323102-0011101221222131-3130003322010320-3132031110310220-1013032222332310-1301012031302032-0011002212300032-0030312013102011"></a>

## ethernet_interface.dhcp_server.automatic_from_end — automatic_from_end / 122002223003 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-3320130022001202-0322311233012202-3211222321303230-1321022313333023-2033100022332001-2331301220332221-3132230323023021-3001311203022111"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1313330320022131-2310031012102122-0231122022023223-3113021012111031-3023212210022310-1313110013013202-2333101230332223-3100230103312132"></a>

## Direct properties — automatic_from_end / 122002223003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130320023100012-0230102110231023-3331010303120331-3200021010022000-3331110313020320-0212300120201001-2303100020031310-1322303100201110"></a>

## Next pages — automatic_from_end / 122002223003 / 4

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1213132133310111-3012101232022020-1230330321131133-0221221213002212-2110020031120002-3003123011123212-1013322212033222-2110032231323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203320203302020-3120220121222013-0020211313032202-1030120201011130-0112310130030210-2000123332220221-0100132032211012-0121010211331331"></a>

## ethernet_interface.dhcp_server.automatic_from_start — automatic_from_start / 122333112021 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-3313323210320100-1033100311222113-1113131013100022-3313221121213102-3223221112210202-0201233103302131-0213011233301203-3311222013022232"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3320130101221323-0203323331112033-3113100021122133-3300113101101310-2320110030001133-2110312122213320-3111031002202121-1022332032222113"></a>

## Direct properties — automatic_from_start / 122333112021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033223232022203-0321312302333333-1121110131032031-3322310223030122-2022323331323021-2220100312332223-0130011023310210-1011320112233303"></a>

## Next pages — automatic_from_start / 122333112021 / 4

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022102101330320-2033131111232132-1311121301311131-0020131231223302-2202010231112321-2221010031023211-3220320323221013-0230111200000133"></a>

## ethernet_interface.dhcp_server.dhcp_networks — dhcp_networks / 203000121031 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-2302031210123303-0301220333223021-1031020220013221-1001001112123012-2232002122312322-2122102200232022-3331223310030112-3320321312121021"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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

<a id="canonical-3103113100310321-1023031202220333-2120303332122020-1033230311123133-0100021001133021-3000313001121221-2303321031310121-1132301023123311"></a>

## Direct properties — dhcp_networks / 203000121031 / 3

<a id="canonical-0330033301012221-3221003222032110-0020103002103322-2232130022033313-1231302113020220-3102012022122233-3302213121012002-1223122113110101"></a>

<a id="canonical-1203030132321203-3313113233200123-3201130023013310-1030321132102030-3023323021002031-0103330202200231-2110231031113332-2220003131003023"></a>

## dgw_address property — dhcp_networks / 203000121031 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1021210132330030-0320113001012320-1220120231301102-0102203312021212-3022011132000003-2332303103220331-3301313132321121-1112331000202021"></a>

<a id="canonical-3111322032102100-2122312012231001-2101300221111300-1100201013102220-1010133203000233-3012203010321110-1221230033103102-2031222331002232"></a>

## dns_address property — dhcp_networks / 203000121031 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--network_interface--reference--group-001.md#canonical-3123310101130231-2320101120233203-2220003303302101-1213302031220223-3123011331301303-0131230131202121-3320220012120331-0231031312013221): complete subsection reference.

- [last_address](data-sources--network_interface--reference--group-001.md#canonical-2003030231203230-1000320233212023-0200100232232103-0001120101003303-3030023223201200-1023230300221132-1311121130112322-3220003330310022): complete subsection reference.

<a id="canonical-2001321001030233-0122003322031001-0200120312223200-0120132110330013-3301233122012222-1010302033231120-2100211203300000-2323002330011010"></a>

<a id="canonical-0133001121311011-1221333002210120-2013113311232031-0332232102031013-1022103300322120-1110210131101022-2233212112111001-2320010213311221"></a>

## network_prefix property — dhcp_networks / 203000121031 / 6

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3000332012101113-0310003303101112-0023203232201001-3101232303332300-3201112010220023-0020103120211331-2221232133210113-3302102333021230"></a>

<a id="canonical-3022231320311030-3133303133013213-0012032313113301-1211030123020102-0010311313221012-3213021030231303-3100221031120310-3221130003111303"></a>

## pool_settings property — dhcp_networks / 203000121031 / 7

Type: `"string"`. Computed.

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

- [pools](data-sources--network_interface--reference--group-001.md#canonical-1302301201311002-1320301322113132-2000100103202213-2210303302320331-2323233210101001-0001332021013311-0323120033310013-0300311330201332): complete subsection reference.

- [same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-0312033220223213-2031132300300320-2110022121023323-1113220300222231-2110323210012211-1201303002102033-2312201121300001-3032031301201131): complete subsection reference.

<a id="canonical-0013313223333330-1030101010313012-3113233311332303-0130331303331313-0221102030331010-3232003101102301-3131301330312302-3001033111023231"></a>

## Next pages — dhcp_networks / 203000121031 / 8

- [ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--network_interface--reference--group-001.md#canonical-3123310101130231-2320101120233203-2220003303302101-1213302031220223-3123011331301303-0131230131202121-3320220012120331-0231031312013221)
- [ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--network_interface--reference--group-001.md#canonical-2003030231203230-1000320233212023-0200100232232103-0001120101003303-3030023223201200-1023230300221132-1311121130112322-3220003330310022)
- [ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-1302301201311002-1320301322113132-2000100103202213-2210303302320331-2323233210101001-0001332021013311-0323120033310013-0300311330201332)
- [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--network_interface--reference--group-001.md#canonical-0312033220223213-2031132300300320-2110022121023323-1113220300222231-2110323210012211-1201303002102033-2312201121300001-3032031301201131)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3123310101130231-2320101120233203-2220003303302101-1213302031220223-3123011331301303-0131230131202121-3320220012120331-0231031312013221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233312323222111-1101002120320202-0321030221010020-1321021112233211-2020010211203220-3030110300003221-0332010100330230-2200131120232023"></a>

## ethernet_interface.dhcp_server.dhcp_networks.first_address — first_address / 200123102233 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-0200132021100203-1310210311200020-0330021111212120-1213112102030202-0330022000303333-3113013312111032-3233032200212201-0300223203023101"></a>

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

<a id="canonical-2311130100130221-1333230232311233-0132021321033323-0030133002312121-2201022110101130-2023023130111331-0303012200330303-0002202222102331"></a>

## Direct properties — first_address / 200123102233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210200331123022-3113033311220032-1011003022000210-2011111333231330-0010110121211333-3222123122201313-2032122312130133-1020122203310232"></a>

## Next pages — first_address / 200123102233 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2003030231203230-1000320233212023-0200100232232103-0001120101003303-3030023223201200-1023230300221132-1311121130112322-3220003330310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100312002132323-3121121103220333-2013020203322311-0213321111221121-2021300120211120-3001102133300122-0301103332310033-3213100313232130"></a>

## ethernet_interface.dhcp_server.dhcp_networks.last_address — last_address / 133212000013 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-0230121010121220-2030230331123032-3201021101202202-3111212011310120-0003132201313232-3233032010022211-3212121312130030-2311100130023000"></a>

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

<a id="canonical-0131111302231300-3003111000211000-2323303131131031-0021331133333302-1200202310330003-3132033223030103-0110102132200032-0333203203312310"></a>

## Direct properties — last_address / 133212000013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121332123123322-1032201200310203-1232313033332110-3200220210221112-0100223233320121-0322323303023332-2233022312012010-2013311103313301"></a>

## Next pages — last_address / 133212000013 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1302301201311002-1320301322113132-2000100103202213-2210303302320331-2323233210101001-0001332021013311-0323120033310013-0300311330201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200001121102310-2013311313231201-0210213033032200-0203220213202132-3113101222311311-1220013002202230-2100002310031320-3133030033212102"></a>

## ethernet_interface.dhcp_server.dhcp_networks.pools — pools / 033123023222 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-2011003312200012-1201013310023133-1120132313210101-2110232003220011-3210223021202011-0203001003220012-0120133112303312-1302300200211001"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2301133020111100-1021223202332102-1023001213303322-3232313102230202-2032131110332121-2231310331100030-2101113211030110-0131331322330011"></a>

## Direct properties — pools / 033123023222 / 3

<a id="canonical-2100101213030320-0311120222231300-2011310102333322-0132221210330101-0333200313212321-3203201201322331-3100112312301233-2010200100301001"></a>

<a id="canonical-1221313022003230-2101022331211320-2222102101123330-3303232321133220-0320333331222322-3032112030310000-2211303003302202-3002003332312220"></a>

## end_ip property — pools / 033123023222 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0211023133321221-2033230201031013-2210023321212233-2133122323011000-0232111133201212-1133231122122110-2020023320302031-1121103013212132"></a>

<a id="canonical-0332031211113302-0322133011300030-2011311212302210-0222231303322122-1333110103232131-0221200302012203-3121300310320332-3021113322330131"></a>

## exclude property — pools / 033123023222 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-2330223010221132-2030012313002103-3103032110330030-1200123012221222-0110303112002222-2321330220313300-1031032313021331-3131021300103103"></a>

<a id="canonical-1210333312132022-0311023320022301-2011313301033310-1321202312130023-1201033023111200-1321001011120002-3133122312320003-2202232023223300"></a>

## start_ip property — pools / 033123023222 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0110020233231031-2220333133303021-0030232102010022-2313210010301302-2133300022213300-2303213202233222-2331020120020003-2111312022013103"></a>

## Next pages — pools / 033123023222 / 7

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0312033220223213-2031132300300320-2110022121023323-1113220300222231-2110323210012211-1201303002102033-2312201121300001-3032031301201131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221323100330231-3223001330121101-3333012130311220-0123000233321100-0331112010133102-3303200332332110-1212101320210131-2332032100220312"></a>

## ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 110233112101 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0000120110111221-0022213312102202-1020130213213032-1131202021220332-3001212133020331-1022233031110300-0201033130233202-0103320000233101"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1103020320023012-2001013022002233-0312101323011112-1112303132231100-3312213312023302-1103220002102310-0130232301023020-3000113013213022"></a>

## Direct properties — same_as_dgw / 110233112101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123131330203220-1033103033121312-3213011022010133-0030030200022313-0200132130333233-1023232321322331-2120223203330022-2120013011213133"></a>

## Next pages — same_as_dgw / 110233112101 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0102031303123133-3210033231031111-3332001310200032-0302121200120133-0231111002210130-0021232021310111-0301232112000213-3003102103212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332002011332200-2303323032233001-3112302122233131-3200101032033133-3010121012232210-0222113213121112-1113022312032213-3011023221211223"></a>

## ethernet_interface.dhcp_server.interface_ip_map — interface_ip_map / 322013020030 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-3312202101330131-0022223121220211-3011023330021010-2031231103002023-2100032301213011-3023101231221211-0200002300123003-0110231020332103"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3303310312003211-0232220210100212-1233103031301333-1230000010322110-0333103310111303-3211010122331002-2222322323011311-2002003313300311"></a>

## Direct properties — interface_ip_map / 322013020030 / 3

<a id="canonical-0020031131211031-3021331303300100-0100113022312110-2201310010213302-1132310111111311-1211231310310231-2322102300202103-2302023323122310"></a>

<a id="canonical-2121100110303021-0230130230200330-3231123231101313-3132221211210001-1010213012323232-2310230210201121-2210021113130330-2201330020310111"></a>

## interface_ip_map property — interface_ip_map / 322013020030 / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

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

<a id="canonical-0230031322202013-1103333320022103-1133313101202031-1032110133130322-3003031332011332-2113321022031332-0111130320132111-0010032310011310"></a>

## Next pages — interface_ip_map / 322013020030 / 5

- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312312321222001-3002302212322332-2002302110031112-3221010101221312-0222200123320332-2331001123111111-2331002030230211-0330000012010231"></a>

## ethernet_interface.ipv6_auto_config — ipv6_auto_config / 020220122012 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.ipv6_auto_config

<a id="canonical-0302020210100300-3221301212010300-0120300101011001-0003332121311303-0330322202212003-0311013102130210-3321211000031230-3323302313202013"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

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

<a id="canonical-3020230303122303-2110010100332333-1332220233120322-2112212033200210-3233233100130020-3000223303010233-1321213310331320-1023223100002023"></a>

## Direct properties — ipv6_auto_config / 020220122012 / 3

- [host](data-sources--network_interface--reference--group-001.md#canonical-3131222310022111-2212231101112200-3221122330332001-1000003310221013-3003113102023211-0030331000001311-3112101313002313-3201302310133130): complete subsection reference.

- [router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201): complete subsection reference.

<a id="canonical-1210023022110321-2302312131121132-3002303110122330-1310223203220320-0302013031222012-1011333212322230-3111322230033210-0230300331202331"></a>

## Next pages — ipv6_auto_config / 020220122012 / 4

- [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--reference--group-001.md#canonical-3131222310022111-2212231101112200-3221122330332001-1000003310221013-3003113102023211-0030331000001311-3112101313002313-3201302310133130)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3131222310022111-2212231101112200-3221122330332001-1000003310221013-3003113102023211-0030331000001311-3112101313002313-3201302310133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201111101033232-2211203020032012-0123031230322022-1321332001230320-3320123223311233-3121023322031132-1120012000333201-2000011032223332"></a>

## ethernet_interface.ipv6_auto_config.host — host / 002110313220 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- ethernet_interface.ipv6_auto_config.host

<a id="canonical-3331221110103103-3201212303332230-0333010223210202-0103302132212210-0000000110322111-2323100313100101-2110312103123131-3111330333323302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0032333223122111-3222001322221033-3312021230033330-3002322210013201-1031132312311211-1030030020022112-1031020023200311-1031200002222111"></a>

## Direct properties — host / 002110313220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321112123132123-3032323221030113-3012311020103110-2102023111102010-0120220102000031-0330322031021300-0010321201320322-3011323123210103"></a>

## Next pages — host / 002110313220 / 4

- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202000010331130-3232103103131022-1321300210220023-1210233012321210-3022221201311222-1211300013330222-0223200131102000-1311333213022221"></a>

## ethernet_interface.ipv6_auto_config.router — router / 213303331222 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- ethernet_interface.ipv6_auto_config.router

<a id="canonical-3201033233203023-3322121320002122-2213313113131111-1313212101013203-2301310100033123-3012210212001002-1102003312322302-0011213021011100"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

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

<a id="canonical-0101300233303202-1302220330301330-2232233310210202-2011323300123200-1123330312022201-0113001302021031-1100020101231211-3031201013313330"></a>

## Direct properties — router / 213303331222 / 3

- [dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033): complete subsection reference.

<a id="canonical-0022001133123231-3230133323121222-3113303122032123-3200122003113110-2231033203123223-0020322110002330-2300311110103203-0121303012231312"></a>

<a id="canonical-3102131033101333-0333212201021220-1321313111032030-2232332312312010-1331013300032002-3202330333210112-2000202313320022-2312332111331113"></a>

## network_prefix property — router / 213303331222 / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020): complete subsection reference.

<a id="canonical-1112010032323210-1233001020032320-3103023210233310-2223102111121230-3000133311101131-3231230132131112-0121233030002303-2100011303121111"></a>

## Next pages — router / 213303331222 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000123131022100-2310131123333011-3223023231320031-1201030030201330-0111111132121010-3010210210123200-2210223212023330-2010103331033320"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config — dns_config / 333233311203 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-3300212221303020-1220310012310230-0232200220220213-1231033323321011-2123110220202003-1001110311310221-0203232112111020-3230003100312130"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

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

<a id="canonical-1212021232200232-1323200011023323-0101123322203113-0313302200213301-2332233010201323-3121210110210111-2123212200230213-1213222122322103"></a>

## Direct properties — dns_config / 333233311203 / 3

- [configured_list](data-sources--network_interface--reference--group-001.md#canonical-2020032120123230-1213332300010231-2012301301013001-2111202322121303-0332000203211102-1303223203320202-0130103330312030-0132123002103133): complete subsection reference.

- [local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332): complete subsection reference.

<a id="canonical-3100313313100001-3200120121321000-1203313122331322-1220223330331120-0112031003302022-0112122002120002-1002010121111133-2021202333133030"></a>

## Next pages — dns_config / 333233311203 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--network_interface--reference--group-001.md#canonical-2020032120123230-1213332300010231-2012301301013001-2111202322121303-0332000203211102-1303223203320202-0130103330312030-0132123002103133)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2020032120123230-1213332300010231-2012301301013001-2111202322121303-0332000203211102-1303223203320202-0130103330312030-0132123002103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110222121202111-2233111100020010-2131102230203303-1311211000301013-3331120321101302-0100300320223020-1012320223001332-1332012301022102"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 320110012031 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2211000233103000-1203032011311202-3012300201203210-0313111233300131-1313210003312010-1121330033031011-2233323031110321-2001330001210201"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-1313311110030113-0320202020303330-1131223232331320-3102122330330031-0332120312231302-1100130223330132-1133012133130221-1213303023030330"></a>

## Direct properties — configured_list / 320110012031 / 3

<a id="canonical-1110010201200130-0321312110032100-1312031113001313-1202233002223310-3110131200122103-3032212311311221-2233202000212032-3233100000220020"></a>

<a id="canonical-1331302310013233-0200202133322032-3312000102133320-3023331100331220-1223033313010310-2032210003000003-2323100020121020-2230011330102300"></a>

## dns_list property — configured_list / 320110012031 / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

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

<a id="canonical-1313311210310201-2022102003223332-3200333303211310-2321310211301110-0321310123321213-1222121221022102-1100222310212033-1103021033103030"></a>

## Next pages — configured_list / 320110012031 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121021310300231-1013322231201131-1120202303201101-2103101102012103-2010303111113221-0313222231103332-1333211202101313-3030003000030321"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 212003201321 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2203102230100313-0230121002311320-1123033200312321-0331200121121203-0123032333010000-3030210100000130-0032232233300222-3221212212320301"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

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

<a id="canonical-0010232130232223-2013013213023232-2110202000200232-3313032023201012-1211011311312130-1023310213122311-0133120333221012-0121202320022032"></a>

## Direct properties — local_dns / 212003201321 / 3

<a id="canonical-3311331101200121-1322231201333223-3232110101002322-2200222021222322-2223020010110213-3201212010302113-0303130013232203-3320032302133220"></a>

<a id="canonical-0231200300103222-2031033023322322-1200120302002011-2333130203231321-3122011301332100-0203132303032121-2210110011223300-3011302220023033"></a>

## configured_address property — local_dns / 212003201321 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](data-sources--network_interface--reference--group-001.md#canonical-2300300312202001-1330222030310001-0032123200032033-2010230023312032-2130331230030203-3030010233033310-0210103200211000-1113022111013112): complete subsection reference.

- [last_address](data-sources--network_interface--reference--group-001.md#canonical-0320130330003210-3303331021101321-1000330020223030-0010033332003002-0122230011310020-3220302021132223-2030122012211023-0001031233101332): complete subsection reference.

<a id="canonical-2032023211101331-1332013122210003-0332313322030130-0301003133122200-3230111112211311-3302002230323302-3010303312233002-3321302321121001"></a>

## Next pages — local_dns / 212003201321 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--network_interface--reference--group-001.md#canonical-2300300312202001-1330222030310001-0032123200032033-2010230023312032-2130331230030203-3030010233033310-0210103200211000-1113022111013112)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--reference--group-001.md#canonical-0320130330003210-3303331021101321-1000330020223030-0010033332003002-0122230011310020-3220302021132223-2030122012211023-0001031233101332)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2300300312202001-1330222030310001-0032123200032033-2010230023312032-2130331230030203-3030010233033310-0210103200211000-1113022111013112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030231311230130-2233021021002233-1201030000222023-2122130131121113-3200011202223231-3312012321303120-1002120322310313-3130112110102132"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 310302301100 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3010001321030301-0122231103201122-0230120221213002-0231122010112032-2223203333330132-1132102000331222-1023230112203230-3122201301123303"></a>

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

<a id="canonical-2201011100001131-3210011002021113-2021003022330011-3120123212221313-0121010123031300-0231012120233103-3301222111103002-3111203330031101"></a>

## Direct properties — first_address / 310302301100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322303300333302-3311202011211332-1101303123003133-3130010312113130-2023311202113322-0111101321310220-1332123230313001-3321221120120300"></a>

## Next pages — first_address / 310302301100 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0320130330003210-3303331021101321-1000330020223030-0010033332003002-0122230011310020-3220302021132223-2030122012211023-0001031233101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102233331003233-2211001120310103-3233231202021031-2022333133232003-0121112101002122-1122302122301010-2313200201220300-2200032132201223"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 310203220130 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0133213203213220-1230012020100121-2100232030010203-2120121220012201-2133313120113031-1300223032230310-3131100033332201-1222213302121200"></a>

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

<a id="canonical-1000232301331330-1032121322022301-3302231010020330-2220112113012203-1333300222020010-0131202321213003-3022313022201203-0103301303210323"></a>

## Direct properties — last_address / 310203220130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321233201101013-1312303031202322-3122211213001322-3022233113211232-2102010200311301-0322212200232232-2112210023212011-0130300202121120"></a>

## Next pages — last_address / 310203220130 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032301120021233-2202201021022302-0321310310302222-2321111131221020-0023221020320110-1111132120030303-2200211320322113-0303311301102123"></a>

## ethernet_interface.ipv6_auto_config.router.stateful — stateful / 032220010022 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-0211210133333101-1212112313030220-0013000103333121-0330213320123022-0012330113323220-3221032022321212-3122213131011121-2113011020023331"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

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

<a id="canonical-0033102102221321-0231113230023320-2233331223311013-0201001121333131-0300222201223002-3201100312302122-1220120012111133-2103011211000031"></a>

## Direct properties — stateful / 032220010022 / 3

- [automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0000012033230200-2123322310320231-3023332320131122-2113112303013230-3121022111002322-3101310220220213-3313302121303200-1220311003030223): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-2010033133131333-2200020111311110-3233200132131213-1323112111122201-3123200131221030-3012302210110002-3233132121302011-0323120301311233): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311): complete subsection reference.

<a id="canonical-0202210103210122-3111110112300232-0221230111301331-0120222112303133-0232113312333000-0031011330211022-1112032233321331-3233312210210320"></a>

<a id="canonical-2010102223212033-1131031231211311-3332223301213222-2032012101200220-2031223121202033-0303100302023001-1212301203133321-1122010300330212"></a>

## fixed_ip_map property — stateful / 032220010022 / 4

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0313100122221133-0112300123320232-3110111202202033-2033121211010013-2332101313222002-0200022012131121-0232032031100332-1212210312222200): complete subsection reference.

<a id="canonical-1011322313003313-1102320230000221-1221011323312030-0213202130321020-0103001133300221-2113200110313303-1122130033103012-2131031032311122"></a>

## Next pages — stateful / 032220010022 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0000012033230200-2123322310320231-3023332320131122-2113112303013230-3121022111002322-3101310220220213-3313302121303200-1220311003030223)
- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-2010033133131333-2200020111311110-3233200132131213-1323112111122201-3123200131221030-3012302210110002-3233132121302011-0323120301311233)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311)
- [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--reference--group-001.md#canonical-0313100122221133-0112300123320232-3110111202202033-2033121211010013-2332101313222002-0200022012131121-0232032031100332-1212210312222200)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0000012033230200-2123322310320231-3023332320131122-2113112303013230-3121022111002322-3101310220220213-3313302121303200-1220311003030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211122030010011-2020303300022302-2010123032322200-3331032311302213-1233030311133303-0011011222311300-2302302111300313-1233111030321023"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 313210000220 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3212233301121232-2130123311011332-3100031100222122-2321300013113032-0121030012233301-3021101113011300-2322121331111003-1010220110022310"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3321330322210111-2313111200103211-2013133322232233-1100033310231022-0111130232303230-2120301223133210-3330310211110230-2302222003023023"></a>

## Direct properties — automatic_from_end / 313210000220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110022312000220-3320212231102103-1201010213023320-2222000212030030-0320003121021111-0333333223220021-0110310233201301-1013213302030333"></a>

## Next pages — automatic_from_end / 313210000220 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2010033133131333-2200020111311110-3233200132131213-1323112111122201-3123200131221030-3012302210110002-3233132121302011-0323120301311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311230333112200-1223300101021112-3102333212332301-3221232303220212-3113112002300010-3333131310231021-2311023110012100-3023001201020200"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 211321033320 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3133223121002201-2331311321300100-2131011300122101-2232030221102101-3212132132301102-2112023122023002-3211310212131030-3313111233322311"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2101212311103022-3321231231311001-3033213323112231-2212011123120313-3111022313012003-2302013212203030-3230332010312211-0033120033321030"></a>

## Direct properties — automatic_from_start / 211321033320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102002111303020-3213121123123021-1101111110002221-1003033201010031-0222001230321132-0320333001223231-2223320303122031-1100012210032301"></a>

## Next pages — automatic_from_start / 211321033320 / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101111002210010-3110310301211130-0001121222011113-3220310002300303-3313203200032302-0212320001313320-2320233333221330-3322200110003002"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 121031220000 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0211113313111232-3031222211130022-1202100220220211-2302022031021232-2330231132133233-0032020311301210-0231310001331232-1021213223213000"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3323033131230022-2230223001113030-2122320232021033-1231011013122011-0012010201200222-3231031210313033-2222212311132222-2232211021003113"></a>

## Direct properties — dhcp_networks / 121031220000 / 3

<a id="canonical-2220320031213230-3230003202122201-3223322330203120-3023322233012200-2011003221000122-1330332233223032-2321232302220013-2103032221313222"></a>

<a id="canonical-3211220200102211-2132133221010113-1313000231300223-3203201210112120-3032001231010032-1110132000120102-1220130220332032-1333310333310312"></a>

## network_prefix property — dhcp_networks / 121031220000 / 4

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-1121132332210102-1110301102220113-0221223330332330-1203113331133113-0100331110203113-1213202100020221-0002111013200103-1023023202333001"></a>

<a id="canonical-2233313233300330-1002312320020001-3320330110312300-3232210232323022-1120032320022223-1321032102031031-0121103120122010-2331302100332233"></a>

## pool_settings property — dhcp_networks / 121031220000 / 5

Type: `"string"`. Computed.

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

- [pools](data-sources--network_interface--reference--group-001.md#canonical-2122032230101010-0300100203020123-1230133021121321-3110131013003313-0022203123012130-0101232010230201-1222303132013023-0123303122130232): complete subsection reference.

<a id="canonical-0333011112110331-2130310020203232-0131133123033200-3302331231310012-1310310310021220-1032122121312221-3000202110322013-2122113012322300"></a>

## Next pages — dhcp_networks / 121031220000 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--reference--group-001.md#canonical-2122032230101010-0300100203020123-1230133021121321-3110131013003313-0022203123012130-0101232010230201-1222303132013023-0123303122130232)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-2122032230101010-0300100203020123-1230133021121321-3110131013003313-0022203123012130-0101232010230201-1222303132013023-0123303122130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003113100333120-2003022113113231-3012230001331323-0220110210031032-2221121311230012-1021213333333333-1203020100011003-3323010023133122"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 223031130311 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0033220313221110-3021213303323113-1020220020332113-3111021021033033-2110132022031133-2202113202220222-1000123111301023-0022223120013302"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0330002210032113-1032002112232310-3322001031131312-2311323000223330-3110002301311010-3311132300222112-0220202121131330-3222030202233130"></a>

## Direct properties — pools / 223031130311 / 3

<a id="canonical-1120211201132033-0120203011322001-2321100002301310-3310331233211120-1002231002001220-2222123313222310-0322033231023102-0123133231200212"></a>

<a id="canonical-1312032331233213-2323132120201333-1333030203020123-0113022311213110-0013323012332301-2231310012213223-3132230130232133-3310110331132330"></a>

## end_ip property — pools / 223031130311 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3211031130330112-1230233331033100-2032230211223123-3312223311013211-2323001020222223-0023223113003010-2230013302012032-1312112113031010"></a>

<a id="canonical-1213203232213233-2333002203133131-2231033010310213-2012221223022122-2303000200130001-2231331112122233-0203111220002001-2133230121200130"></a>

## start_ip property — pools / 223031130311 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2110311033223213-2131003201130332-1233001020210232-3301100313010311-3330312011012231-1223313202323123-0122213332121311-2112120311113323"></a>

## Next pages — pools / 223031130311 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0313100122221133-0112300123320232-3110111202202033-2033121211010013-2332101313222002-0200022012131121-0232032031100332-1212210312222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102202303112100-3212202310200023-1211232210122033-2122331130312300-1013332111112233-3101333012312131-2000312310000032-1200110020323201"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 033202110222 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0110112110023301-1331211011310003-2200302313102001-1302100331033111-1002011202200203-3321301220322303-0021310222130033-2100102311020010"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2121333003033101-0222221200122233-2202112330101312-2231101233331203-3102211331232122-1232100210100223-2131211303003131-0111220032001020"></a>

## Direct properties — interface_ip_map / 033202110222 / 3

<a id="canonical-2020322230121201-3210301200123321-0332220333133332-1102011212111332-3033020013103203-2110123310212220-0122133022023303-3001303030222001"></a>

<a id="canonical-2223022233310023-3123321330203330-0002122000223110-2302013032001023-1111122012330020-0321210333122323-3131103322210322-1001031103230302"></a>

## interface_ip_map property — interface_ip_map / 033202110222 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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

<a id="canonical-3301303120211020-0013002022102332-0332201130300200-2233321113200221-2030201230312300-1122021310223102-3022231130032010-0210101201100212"></a>

## Next pages — interface_ip_map / 033202110222 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-001.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0131322022011012-2232101232210032-1121023033132330-2102122010330022-0303011212231321-0030120321323012-0002332223030223-2331132203033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212221332020031-2021333033311132-0233112310021111-0201231023031200-0111130322031213-1102003120013123-2322001311102312-2323211303010311"></a>

## ethernet_interface.is_primary — is_primary / 301120300210 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.is_primary

<a id="canonical-0213021100100001-0330331013011332-1301300210130322-2100302201203120-2333131111122332-0321010120202132-2110223133113210-3222133330102032"></a>

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

<a id="canonical-3200200231213320-2110023021310210-0332020112010201-1113002010313032-0013200231203103-1010201202001130-0310000102211202-0000023102303210"></a>

## Direct properties — is_primary / 301120300210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102232013030103-2010132130212233-1330111321223202-1320121033102320-2123002000213001-0331210033112312-2011103322030323-1032211202323300"></a>

## Next pages — is_primary / 301120300210 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030223202022031-2131012121132132-2012331032013123-1133132001132120-2203332011220121-1002100203333130-2230220232013220-1231223033202302"></a>

## ethernet_interface.monitor — monitor / 231211323321 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.monitor

<a id="canonical-1033313201201102-1202031310100130-2200001133331030-1000023213110002-0312220033120210-2301122322130010-3112000033013131-3110200002213011"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3233110322213001-1130033320122102-2023032323223232-0012031223220103-1203013100011112-1213211122012123-1200311221233030-0211231302131302"></a>

## Direct properties — monitor / 231211323321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311313212012322-3001320120233213-0230212213332301-1111222100220030-2211032331010112-2013221230130313-0133210211101320-1322221331010301"></a>

## Next pages — monitor / 231211323321 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3121330013003013-2100012130332222-1133031333201222-2021321321232111-2002301003121210-2322111332332230-0221011000030231-2023233203032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321323301020312-1112200133330220-3132200221122012-1312212003232010-3300212111221222-1230311001002103-1132011300310311-0303332331122312"></a>

## ethernet_interface.monitor_disabled — monitor_disabled / 203313210222 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.monitor_disabled

<a id="canonical-1233112322031330-3011130002322221-0222223321331223-0210122323213220-2313020113230300-3112300233331000-1233113322320200-3102131010331130"></a>

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

<a id="canonical-0212200110133012-0033011000132201-0022301001130122-0112131300131121-3220123300223332-3022003313323203-3101010121130223-1031000133200102"></a>

## Direct properties — monitor_disabled / 203313210222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131100110332222-0322332112322100-0132032130131021-2102100111000300-1320010012011212-3320201023221330-1023010102102222-2321301303002020"></a>

## Next pages — monitor_disabled / 203313210222 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-1113032321321001-2310200133001032-3213200220203223-3000220113031202-1230221312200012-1202311001030122-0321131000121000-1111313323032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332003220212211-1023033313011132-0312031013033302-1321313210333321-0303132111221123-3111001323002132-3311021323033002-3322213320011103"></a>

## ethernet_interface.no_ipv6_address — no_ipv6_address / 000231313030 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.no_ipv6_address

<a id="canonical-3310222213211122-3121201300312013-1031032213131200-1110202122301030-0013031232130220-1022323122000020-0201132201331330-3130312311231002"></a>

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

<a id="canonical-3031031100032230-1321310220310013-3020022111013223-3023132030012332-1132032021203100-0213133232123130-2201232220100000-2120222230302203"></a>

## Direct properties — no_ipv6_address / 000231313030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022331000231012-0232300112313022-0321100210300022-1322231231101211-2312323323232020-3303132000310102-0131221133011130-1231130331203123"></a>

## Next pages — no_ipv6_address / 000231313030 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3312231302230300-0022030322113100-2201123131132030-3100011010120320-3310232100021031-1310303103110322-1031013222131303-3130102302313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131022311023113-1120100023131323-1113110133332103-3130101132013200-0300012112203103-3002302100010102-2300103012021221-1003322003032000"></a>

## ethernet_interface.not_primary — not_primary / 330230032100 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.not_primary

<a id="canonical-2231103231122332-0203212220131032-0332002321021100-0002002031012132-1310030202110020-0322100203202133-0001002210002211-2202103002010330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3302000110333203-3212101301210133-2210122111230231-0232311101233111-1101223202322010-3111102003221110-1032032010212120-0320300211312001"></a>

## Direct properties — not_primary / 330230032100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013202221023223-0101221012113211-2113113202303023-0203210012302101-1303133132032000-0122132231001002-0111032022003211-1333032112200332"></a>

## Next pages — not_primary / 330230032100 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-0212312303120132-1113000211222110-0213011322231122-3213230332311203-0133301012122321-3102222301230020-2232303202233323-3312131000300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132231231322332-1311203020323321-2023222301021023-3002330030232120-3122130203323031-1130032020000311-0113213131110331-3313030300120213"></a>

## ethernet_interface.site_local_inside_network — site_local_inside_network / 120231212210 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.site_local_inside_network

<a id="canonical-3112330210230232-0131201333110111-3323130032020110-3103031120333133-0333111030221031-0100321022123313-2032200230112312-0103130122113323"></a>

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

<a id="canonical-3211212110022322-0022030101130000-3113212132033200-1121300211120302-0001301211111122-0300033011233013-2200011211031201-1101330203130001"></a>

## Direct properties — site_local_inside_network / 120231212210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110100222212211-2301011222223211-0030300100323130-2320001101203132-2130200123302203-0012320003213223-2130231230101201-0003023020002221"></a>

## Next pages — site_local_inside_network / 120231212210 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3132210030013112-1102213013013232-0321001230003132-1320121232330023-0311112012212022-2113030330100121-0300331333001210-3033033130201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102212201220223-0102111201230211-2211212201313032-3212110231101133-2313132133022332-1033323032001013-0330101233122000-0111222011201003"></a>

## ethernet_interface.site_local_network — site_local_network / 023220333113 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.site_local_network

<a id="canonical-1203322000313333-3321000133230210-3020130120303110-1021033120123202-3131132100110100-0002133230231110-2013002233012320-0223232323231110"></a>

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

<a id="canonical-3122011130110032-3311013131312320-3332032321311032-3320123001013220-1201213221203002-0330301211020020-2233103231302322-2100103302233311"></a>

## Direct properties — site_local_network / 023220333113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030122020223221-1233132013321020-2130113133313221-0302003331031000-2213110230223013-2303110202120120-3011011230133031-2231122113301002"></a>

## Next pages — site_local_network / 023220333113 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)

<a id="canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132310233011220-0210110313310133-0312021332100213-0312110002111010-3020002122022100-2213131303000100-1311003331033000-0201300220210301"></a>

## ethernet_interface.static_ip — static_ip / 001303133230 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.static_ip

<a id="canonical-3300133011111203-1330212301220211-2303202300012231-1111013302212331-2002000121213230-2332030021000223-2300131012003112-0110020231130323"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

<a id="canonical-1021212310123203-1103103233320023-1232033231212231-1323032030311131-2312131013102210-3031303101112101-3303300230201333-0133202213101333"></a>

## Direct properties — static_ip / 001303133230 / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102): complete subsection reference.
