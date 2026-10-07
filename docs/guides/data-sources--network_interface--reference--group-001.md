---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- Property reference

<a id="canonical-2232212311231212-1111031300200321-0311232111131211-2311333323032230-1031310113113311-3311132130210103-0121223013003012-0020311130232220"></a>

### Direct properties for `xcsh_network_interface`

<a id="canonical-3211030103000012-1331120300220301-3311102331310100-1232312121031132-3223303002313010-2221313222313032-3022012233323032-0300122300111201"></a>

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

- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321): complete subsection reference.

- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302): complete subsection reference.

<a id="canonical-3330131101312002-0000110202332320-2333113001120202-0203302333113031-2133110100023020-1030103021002123-0111330232023220-1112003110103333"></a>

<a id="canonical-0231231210330030-1032202230322011-0130031132332200-0220210300101303-0302331100200113-3131311330330000-3313131120131003-2122232110231001"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NetworkInterface.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0012223032233132-1111011320223030-1022001230012200-3232110330002132-2230200302232002-3033100310322030-3130123302010132-1213032102012002"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3300321303201002-1102201313323311-2131100200012330-2231330233132300-0323122313013300-0012110210023300-0023133323333332-0220202333003010"></a>

<a id="canonical-0323102223301200-2110030310023223-2130332302100110-2212220122311013-2300320200012320-0010013211302000-2313011323012331-3223112001131103"></a>

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

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030): complete subsection reference.

<a id="canonical-0110322303330303-3213131003322113-1133230103212113-3123331032212132-3032202223321313-3012203032223023-0311032131322202-0311010122021220"></a>

<a id="canonical-2221132122121030-1113230311031200-0101230033111103-3220002312323200-1303233123012131-1022022133312301-2130032222321113-2100310130212001"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NetworkInterface.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311010202032220-0233121302202311-2111122310323032-3103213200211030-3313303103221221-3223121210010031-2111213101333132-3013220323000103"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the NetworkInterface exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1030302112113230-1331121203110303-0033333033003301-2303000122121011-3003121203023331-2223012030203032-0221332212012333-3212201333310013"></a>

### All schema paths for `xcsh_network_interface`

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
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--reference--group-002.md#canonical-0133213203213220-1230012020100121-2100232030010203-2120121220012201-2133313120113031-1300223032230310-3131100033332201-1222213302121200) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--network_interface--reference--group-001.md#canonical-0022001133123231-3230133323121222-3113303122032123-3200122003113110-2231033203123223-0020322110002330-2300311110103203-0121303012231312) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-0211210133333101-1212112313030220-0013000103333121-0330213320123022-0012330113323220-3221032022321212-3122213131011121-2113011020023331) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--reference--group-002.md#canonical-3212233301121232-2130123311011332-3100031100222122-2321300013113032-0121030012233301-3021101113011300-2322121331111003-1010220110022310) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--reference--group-002.md#canonical-3133223121002201-2331311321300100-2131011300122101-2232030221102101-3212132132301102-2112023122023002-3211310212131030-3313111233322311) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-002.md#canonical-0211113313111232-3031222211130022-1202100220220211-2302022031021232-2330231132133233-0032020311301210-0231310001331232-1021213223213000) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--network_interface--reference--group-002.md#canonical-2220320031213230-3230003202122201-3223322330203120-3023322233012200-2011003221000122-1330332233223032-2321232302220013-2103032221313222) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--network_interface--reference--group-002.md#canonical-1121132332210102-1110301102220113-0221223330332330-1203113331133113-0100331110203113-1213202100020221-0002111013200103-1023023202333001) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--reference--group-002.md#canonical-0033220313221110-3021213303323113-1020220020332113-3111021021033033-2110132022031133-2202113202220222-1000123111301023-0022223120013302) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--network_interface--reference--group-002.md#canonical-1120211201132033-0120203011322001-2321100002301310-3310331233211120-1002231002001220-2222123313222310-0322033231023102-0123133231200212) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--network_interface--reference--group-002.md#canonical-3211031130330112-1230233331033100-2032230211223123-3312223311013211-2323001020222223-0023223113003010-2230013302012032-1312112113031010) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--network_interface--reference--group-002.md#canonical-0202210103210122-3111110112300232-0221230111301331-0120222112303133-0232113312333000-0031011330211022-1112032233321331-3233312210210320) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-0110112110023301-1331211011310003-2200302313102001-1302100331033111-1002011202200203-3321301220322303-0021310222130033-2100102311020010) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-2020322230121201-3210301200123321-0332220333133332-1102011212111332-3033020013103203-2110123310212220-0122133022023303-3001303030222001) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](data-sources--network_interface--reference--group-002.md#canonical-0213021100100001-0330331013011332-1301300210130322-2100302201203120-2333131111122332-0321010120202132-2110223133113210-3222133330102032) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](data-sources--network_interface--reference--group-002.md#canonical-1033313201201102-1202031310100130-2200001133331030-1000023213110002-0312220033120210-2301122322130010-3112000033013131-3110200002213011) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](data-sources--network_interface--reference--group-002.md#canonical-1233112322031330-3011130002322221-0222223321331223-0210122323213220-2313020113230300-3112300233331000-1233113322320200-3102131010331130) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](data-sources--network_interface--reference--group-001.md#canonical-3022021211012011-1230021012001233-0131113132130013-1002000211011001-2102210212012323-1020213131031033-1000210022301133-1032322223131203) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-3310222213211122-3121201300312013-1031032213131200-1110202122301030-0013031232130220-1022323122000020-0201132201331330-3130312311231002) |
| `ethernet_interface.node` | [ethernet_interface.node](data-sources--network_interface--reference--group-001.md#canonical-0302021222022010-2311303133323210-1203123030111010-0300331331103201-1110001110302331-1033122133310022-3021310310213122-3211110010232333) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](data-sources--network_interface--reference--group-002.md#canonical-2231103231122332-0203212220131032-0332002321021100-0002002031012132-1310030202110020-0322100203202133-0001002210002211-2202103002010330) |
| `ethernet_interface.priority` | [ethernet_interface.priority](data-sources--network_interface--reference--group-001.md#canonical-1300000122323121-1301300303231132-0301001111230011-1012102100113100-1331200011220013-3331203322303331-1311210210203132-2202132223003310) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-3112330210230232-0131201333110111-3323130032020110-3103031120333133-0333111030221031-0100321022123313-2032200230112312-0103130122113323) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](data-sources--network_interface--reference--group-002.md#canonical-1203322000313333-3321000133230210-3020130120303110-1021033120123202-3131132100110100-0002133230231110-2013002233012320-0223232323231110) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3300133011111203-1330212301220211-2303202300012231-1111013302212331-2002000121213230-2332030021000223-2300131012003112-0110020231130323) |
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
| `id` | [ID](data-sources--network_interface--reference--group-001.md#canonical-1310210210213110-0000112020200133-1023333321001021-0200022320130003-2330123011023011-1331003200133311-0201021310122332-3110211213113232) |
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

<a id="canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- dedicated_interface

<a id="canonical-2313131122020221-0233301323112002-0302000311302211-3112301213001122-3301001203003313-3122131012332030-3221323232202001-3330300001121330"></a>

Type: `"single"`. Computed.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Additional upstream details:

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

<a id="canonical-3332333123330222-0030022111120323-0302313032100021-1101301003110213-0222321022220023-1221222222110103-1220132312033310-1030220003203203"></a>

### Direct properties for `dedicated_interface`

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1211132330113332-3212031033012102-0222101221133332-0311302120203213-2032132300102321-3222013032130332-1210023202332201-1100023032023022): complete subsection reference.

<a id="canonical-0111130032120211-0202231123231000-3310100012301202-2101312100031211-1231222021030232-2232222102321031-1102110132322013-0131103131333330"></a>

<a id="canonical-0033000320122031-0213111310000311-0022232100233310-0312123101322012-2110212121132202-2011323303110010-2132002222103002-0300111310131213"></a>

#### `dedicated_interface.device` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3120200103212133-1203202101023332-3112222222330201-3001223022000133-2331113132133213-2030202032232230-0322212220030003-1201212130320112"></a>

#### `dedicated_interface.mtu` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1321020223212320-0033000002122331-2331330320021312-3210121332330203-2113220020012100-3123021221333112-3103213031000032-3312212033303330"></a>

#### `dedicated_interface.node` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2002103021200003-2011232213123101-0113133310033302-0303300112032013-1222001200222101-1131323132202312-1313311313212020-3203101330120230"></a>

#### `dedicated_interface.priority` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1211132330113332-3212031033012102-0222101221133332-0311302120203213-2032132300102321-3222013032130332-1210023202332201-1100023032023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface.cluster` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.cluster

<a id="canonical-0120233122333101-0022201212300031-0001202002111303-2003020000003100-1021011212112301-3020301023301103-2002312031220103-1303120003111133"></a>

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

<a id="canonical-0022131012331121-3101130212200322-2211202021211213-3113002032312120-1022122112303230-1223110011121212-3323100130212100-2330110111233331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface.is_primary` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.is_primary

<a id="canonical-0202201123102032-2012020231303321-0232310331000222-1122303003120220-3132000330130303-2122002221323133-0121133212102020-3232021322312023"></a>

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

<a id="canonical-1301300033020023-2330210123101223-3133232003002121-1100233103130100-1100331332330110-0203112300000220-2321033203313011-2310121123120321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface.monitor` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030313302121202-1233123111231332-3332201001311232-1022220000121011-2130021022031232-3321001222222110-1130010012122033-2033000003310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface.monitor_disabled` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.monitor_disabled

<a id="canonical-0202320301012200-1233101231000012-1210113321331312-2233132332031222-3201103030121131-1033130033022333-1011102102123330-0313230322231131"></a>

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

<a id="canonical-2101232323003203-3203010310132201-2231031113103111-2121012100323012-1302210102202103-2001313101033321-2133320212012010-3330031331120011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_interface.not_primary` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_interface](data-sources--network_interface--reference--group-001.md#canonical-2111100003310000-3312321232302132-3230210113110323-1133111201200010-1211100011233201-2312033201321122-3130121003002133-2332211313132321)
- dedicated_interface.not_primary

<a id="canonical-1110302013001312-0211111010001230-3230020030011312-0231320331102230-3000131022230222-1223223313302233-0123012123323301-0323232003010000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_management_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- dedicated_management_interface

<a id="canonical-3101133230301021-1301002312033022-0023111132213123-2132231010101102-0322332231113123-1031303333011300-0310111103301120-0123201333330121"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated management interface.

Additional upstream details:

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

<a id="canonical-1002231100101331-0033120002333111-3021333320011012-1003110012121012-2011002120123203-0230210200323210-2112130130012101-3310033211211300"></a>

### Direct properties for `dedicated_management_interface`

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1121030012320221-0012303030303003-3203112120202321-3003221000313133-3022112200032123-3312331333222030-3212201010232230-3221200033002033): complete subsection reference.

<a id="canonical-2011232323321203-3132201102330001-3002012100000200-0011230101010020-3103320220012321-0013032201011031-2111201333320130-1200210031110110"></a>

<a id="canonical-2103211023210221-3013022123320201-2233112300121332-0220000210101223-0322003222311321-3320000003323030-2200113100311200-2002133101131302"></a>

#### `dedicated_management_interface.device` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320330023320032-2312033210300232-3000320300110222-1322303221323132-2222110330131321-3200311303103031-1113110233202021-3201020302002300"></a>

#### `dedicated_management_interface.mtu` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2213123233131233-0130233331313131-2333010302202322-2100130120330303-0002333032210101-0133012330331031-0321320320130010-3023211032001102"></a>

#### `dedicated_management_interface.node` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1121030012320221-0012303030303003-3203112120202321-3003221000313133-3022112200032123-3312331333222030-3212201010232230-3221200033002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dedicated_management_interface.cluster` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [dedicated_management_interface](data-sources--network_interface--reference--group-001.md#canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302)
- dedicated_management_interface.cluster

<a id="canonical-0303230121133202-0320322110103122-2022230011323123-1032313312330201-3022301320300221-3130132011303302-3132320100113013-3002312112211232"></a>

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

<a id="canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- ethernet_interface

<a id="canonical-3231211313330313-0022132220201221-0013000032323022-3300222221132203-2012302133001301-0323102322100320-0311033002022030-2102121202121010"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Additional upstream details:

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

<a id="canonical-1221232332012323-3233303223203312-0201030020333211-0322300202330001-1321332232332110-2002103011223133-3021330133011133-3321211031120313"></a>

### Direct properties for `ethernet_interface`

- [cluster](data-sources--network_interface--reference--group-001.md#canonical-1000310232301022-0020330010030121-0022102133110121-2323122310112011-0022100321232112-1021210323112113-2211012232330000-3121002333213032): complete subsection reference.

<a id="canonical-1101112103132201-3113333302321232-3013202220332111-3020210321133021-3333130033211033-0320200120120112-2013011300300111-3233330202321002"></a>

<a id="canonical-2211312313032133-3131003111022130-1200033311200233-2020223303113232-3023300320310200-1300231123212302-1120310021210112-3002311223311222"></a>

#### `ethernet_interface.device` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [is_primary](data-sources--network_interface--reference--group-002.md#canonical-0131322022011012-2232101232210032-1121023033132330-2102122010330022-0303011212231321-0030120321323012-0002332223030223-2331132203033201): complete subsection reference.

- [monitor](data-sources--network_interface--reference--group-002.md#canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201): complete subsection reference.

- [monitor_disabled](data-sources--network_interface--reference--group-002.md#canonical-3121330013003013-2100012130332222-1133031333201222-2021321321232111-2002301003121210-2322111332332230-0221011000030231-2023233203032113): complete subsection reference.

<a id="canonical-3022021211012011-1230021012001233-0131113132130013-1002000211011001-2102210212012323-1020213131031033-1000210022301133-1032322223131203"></a>

<a id="canonical-0000313022032102-2022121033120120-2101222130021322-0110321011010303-3211221323220023-1033132333033210-2223231210213113-0301213210133123"></a>

#### `ethernet_interface.mtu` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [no_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-1113032321321001-2310200133001032-3213200220203223-3000220113031202-1230221312200012-1202311001030122-0321131000121000-1111313323032111): complete subsection reference.

<a id="canonical-0302021222022010-2311303133323210-1203123030111010-0300331331103201-1110001110302331-1033122133310022-3021310310213122-3211110010232333"></a>

<a id="canonical-2201332110311011-3133030021302221-2101233212331212-0301300113201122-1230232212310333-2203312203202122-0200101201301001-2212023113332210"></a>

#### `ethernet_interface.node` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [not_primary](data-sources--network_interface--reference--group-002.md#canonical-3312231302230300-0022030322113100-2201123131132030-3100011010120320-3310232100021031-1310303103110322-1031013222131303-3130102302313330): complete subsection reference.

<a id="canonical-1300000122323121-1301300303231132-0301001111230011-1012102100113100-1331200011220013-3331203322303331-1311210210203132-2202132223003310"></a>

<a id="canonical-0223203133212013-0021312030322102-1220032213132331-2230022011221300-2330323002310001-0320301120301000-3110213232012130-1223003230203111"></a>

#### `ethernet_interface.priority` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-0212312303120132-1113000211222110-0213011322231122-3213230332311203-0133301012122321-3102222301230020-2232303202233323-3312131000300010): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-002.md#canonical-3132210030013112-1102213013013232-0321001230003132-1320121232330023-0311112012212022-2113030330100121-0300331333001210-3033033130201021): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-002.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231): complete subsection reference.

- [static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310): complete subsection reference.

- [storage_network](data-sources--network_interface--reference--group-002.md#canonical-0333033210010321-2132211112223013-1230003210301302-3330133302221210-0113332030013030-1331320233211220-3001013200001121-2203312010010122): complete subsection reference.

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-2303100002223300-2211310120020011-3133100223301321-2323322300103112-1230330031231133-3230103021330133-1023322211311012-0001221012231312): complete subsection reference.

<a id="canonical-3013023100320231-1122010101010013-2212330311302303-3032203312221011-2010010201010210-0330000331113012-3112223113301311-3303302212323101"></a>

<a id="canonical-1231022221233020-3203331000312103-3111233020333022-2112133312102230-3202212103331330-1013130020220222-2231301211203202-1032021233221203"></a>

#### `ethernet_interface.vlan_id` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1000310232301022-0020330010030121-0022102133110121-2323122310112011-0022100321232112-1021210323112113-2211012232330000-3121002333213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.cluster` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.cluster

<a id="canonical-0200013332221021-3003113022031200-0231212300131133-3223223322231220-0320331200311330-3331102301200100-0221323011210322-3112211032320202"></a>

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

<a id="canonical-2111203310333301-0203133201310122-3302311102203221-0033033020213331-3321331221303031-1022320231103221-0030211110001111-3102232032320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_client` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.dhcp_client

<a id="canonical-2322032301300010-1302320200133131-3302013213312130-2113012210320311-2031030103203202-1313313023032331-3031032002330223-1120232033113002"></a>

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

<a id="canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server` properties

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

<a id="canonical-3023200022220030-2332020311121121-1121132222030100-0310131210002020-2000010213310323-3212302110130003-2000212212330001-1132210122212231"></a>

### Direct properties for `ethernet_interface.dhcp_server`

- [automatic_from_end](data-sources--network_interface--reference--group-001.md#canonical-0122220112022100-1110200101331310-3120100032020331-3000303033202113-3333320102233320-0230121232130021-3200320032100023-1132030210220311): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-001.md#canonical-1213132133310111-3012101232022020-1230330321131133-0221221213002212-2110020031120002-3003123011123212-1013322212033222-2110032231323132): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-001.md#canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121): complete subsection reference.

<a id="canonical-2200220103202222-2203100112023232-2023010021123332-1122312322011103-1303101232311023-2133321230301212-0201020323010023-2322120023123220"></a>

<a id="canonical-0212030113120322-1122213202111112-0212232220031213-3202222123313303-0312302311303122-3222022211102212-2012232111203221-0220023300001122"></a>

#### `ethernet_interface.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-0003032121012210-3322130123001021-3100233202223310-3302021011020031-1312213010302300-2101200112221202-3121113002322012-3331312112033012"></a>

<a id="canonical-2311200231011020-2100022211001233-1120033203021221-1120233113301230-2222300202033312-3230010123313213-0001210330032303-2113311211111322"></a>

#### `ethernet_interface.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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

<a id="canonical-0122220112022100-1110200101331310-3120100032020331-3000303033202113-3333320102233320-0230121232130021-3200320032100023-1132030210220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-3320130022001202-0322311233012202-3211222321303230-1321022313333023-2033100022332001-2331301220332221-3132230323023021-3001311203022111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-1213132133310111-3012101232022020-1230330321131133-0221221213002212-2110020031120002-3003123011123212-1013322212033222-2110032231323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-3313323210320100-1033100311222113-1113131013100022-3313221121213102-3223221112210202-0201233103302131-0213011233301203-3311222013022232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-0002220131312232-1111230002220223-2331300112311102-3121320021111033-2100023311210123-1333033000000011-1200010213322320-0121002323322121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.dhcp_networks` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3022102101330320-2033131111232132-1311121301311131-0020131231223302-2202010231112321-2221010031023211-3220320323221013-0230111200000133"></a>

### Direct properties for `ethernet_interface.dhcp_server.dhcp_networks`

<a id="canonical-0330033301012221-3221003222032110-0020103002103322-2232130022033313-1231302113020220-3102012022122233-3302213121012002-1223122113110101"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3103113100310321-1023031202220333-2120303332122020-1033230311123133-0100021001133021-3000313001121221-2303321031310121-1132301023123311"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1203030132321203-3313113233200123-3201130023013310-1030321132102030-3023323021002031-0103330202200231-2110231031113332-2220003131003023"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3111322032102100-2122312012231001-2101300221111300-1100201013102220-1010133203000233-3012203010321110-1221230033103102-2031222331002232"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

<a id="canonical-3123310101130231-2320101120233203-2220003303302101-1213302031220223-3123011331301303-0131230131202121-3320220012120331-0231031312013221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.dhcp_networks.first_address` properties

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

<a id="canonical-2003030231203230-1000320233212023-0200100232232103-0001120101003303-3030023223201200-1023230300221132-1311121130112322-3220003330310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.dhcp_networks.last_address` properties

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

<a id="canonical-1302301201311002-1320301322113132-2000100103202213-2210303302320331-2323233210101001-0001332021013311-0323120033310013-0300311330201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.dhcp_networks.pools` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2200001121102310-2013311313231201-0210213033032200-0203220213202132-3113101222311311-1220013002202230-2100002310031320-3133030033212102"></a>

### Direct properties for `ethernet_interface.dhcp_server.dhcp_networks.pools`

<a id="canonical-2100101213030320-0311120222231300-2011310102333322-0132221210330101-0333200313212321-3203201201322331-3100112312301233-2010200100301001"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2301133020111100-1021223202332102-1023001213303322-3232313102230202-2032131110332121-2231310331100030-2101113211030110-0131331322330011"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-2330223010221132-2030012313002103-3103032110330030-1200123012221222-0110303112002222-2321330220313300-1031032313021331-3131021300103103"></a>

<a id="canonical-1221313022003230-2101022331211320-2222102101123330-3303232321133220-0320333331222322-3032112030310000-2211303003302202-3002003332312220"></a>

#### `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0312033220223213-2031132300300320-2110022121023323-1113220300222231-2110323210012211-1201303002102033-2312201121300001-3032031301201131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` properties

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

<a id="canonical-0102031303123133-3210033231031111-3332001310200032-0302121200120133-0231111002210130-0021232021310111-0301232112000213-3003102103212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.dhcp_server](data-sources--network_interface--reference--group-001.md#canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033)
- ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-3312202101330131-0022223121220211-3011023330021010-2031231103002023-2100032301213011-3023101231221211-0200002300123003-0110231020332103"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

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

<a id="canonical-3332002011332200-2303323032233001-3112302122233131-3200101032033133-3010121012232210-0222113213121112-1113022312032213-3011023221211223"></a>

### Direct properties for `ethernet_interface.dhcp_server.interface_ip_map`

<a id="canonical-0020031131211031-3021331303300100-0100113022312110-2201310010213302-1132310111111311-1211231310310231-2322102300202103-2302023323122310"></a>

#### `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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

<a id="canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config` properties

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

<a id="canonical-0312312321222001-3002302212322332-2002302110031112-3221010101221312-0222200123320332-2331001123111111-2331002030230211-0330000012010231"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config`

- [host](data-sources--network_interface--reference--group-001.md#canonical-3131222310022111-2212231101112200-3221122330332001-1000003310221013-3003113102023211-0030331000001311-3112101313002313-3201302310133130): complete subsection reference.

- [router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201): complete subsection reference.

<a id="canonical-3131222310022111-2212231101112200-3221122330332001-1000003310221013-3003113102023211-0030331000001311-3112101313002313-3201302310133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- ethernet_interface.ipv6_auto_config.host

<a id="canonical-3331221110103103-3201212303332230-0333010223210202-0103302132212210-0000000110322111-2323100313100101-2110312103123131-3111330333323302"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

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

<a id="canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router` properties

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

<a id="canonical-2202000010331130-3232103103131022-1321300210220023-1210233012321210-3022221201311222-1211300013330222-0223200131102000-1311333213022221"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router`

- [dns_config](data-sources--network_interface--reference--group-001.md#canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033): complete subsection reference.

<a id="canonical-0022001133123231-3230133323121222-3113303122032123-3200122003113110-2231033203123223-0020322110002330-2300311110103203-0121303012231312"></a>

<a id="canonical-0101300233303202-1302220330301330-2232233310210202-2011323300123200-1123330312022201-0113001302021031-1100020101231211-3031201013313330"></a>

#### `ethernet_interface.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020): complete subsection reference.

<a id="canonical-1123300211221313-0221201311100130-2111213002121232-2001132003331220-2130103202001321-1121012031132202-0120203022123031-0033130032303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.dns_config` properties

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

<a id="canonical-2000123131022100-2310131123333011-3223023231320031-1201030030201330-0111111132121010-3010210210123200-2210223212023330-2010103331033320"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--network_interface--reference--group-001.md#canonical-2020032120123230-1213332300010231-2012301301013001-2111202322121303-0332000203211102-1303223203320202-0130103330312030-0132123002103133): complete subsection reference.

- [local_dns](data-sources--network_interface--reference--group-001.md#canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332): complete subsection reference.

<a id="canonical-2020032120123230-1213332300010231-2012301301013001-2111202322121303-0332000203211102-1303223203320202-0130103330312030-0132123002103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` properties

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

<a id="canonical-0110222121202111-2233111100020010-2131102230203303-1311211000301013-3331120321101302-0100300320223020-1012320223001332-1332012301022102"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1110010201200130-0321312110032100-1312031113001313-1202233002223310-3110131200122103-3032212311311221-2233202000212032-3233100000220020"></a>

#### `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3011320002132222-1233021013113230-3012233031331233-2120030233312312-2000222130011230-1031022121233020-1230112020202132-1002312132302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` properties

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

<a id="canonical-2121021310300231-1013322231201131-1120202303201101-2103101102012103-2010303111113221-0313222231103332-1333211202101313-3030003000030321"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-3311331101200121-1322231201333223-3232110101002322-2200222021222322-2223020010110213-3201212010302113-0303130013232203-3320032302133220"></a>

#### `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [last_address](data-sources--network_interface--reference--group-002.md#canonical-0320130330003210-3303331021101321-1000330020223030-0010033332003002-0122230011310020-3220302021132223-2030122012211023-0001031233101332): complete subsection reference.

<a id="canonical-2300300312202001-1330222030310001-0032123200032033-2010230023312032-2130331230030203-3030010233033310-0210103200211000-1113022111013112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

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
