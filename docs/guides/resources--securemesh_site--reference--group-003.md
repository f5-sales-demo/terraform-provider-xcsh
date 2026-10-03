---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-2321113123113110-1312003231303032-1332002203321312-1332221130113033-0320020333120231-3230311022012231-3313122010231201-1202130223010131"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful — stateful / 030303031331 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-0112113122003223-3232120223310300-2122331202030123-1120110020323212-0210202301003021-0203212201010232-1132012121131020-3103020200332320"></a>

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

<a id="canonical-0310000100112230-1223111131202120-3010021010220020-1303122220220112-2233320112103200-3001003200110031-2032202121013211-1313023122020203"></a>

## Direct properties — stateful / 030303031331 / 3

- [automatic_from_end](resources--securemesh_site--reference--group-003.md#canonical-3331021120111202-2002321133002112-1130222300233323-3222111102300322-1012001301202211-0322133111033202-3032330322031210-1023311030032112): complete subsection reference.

- [automatic_from_start](resources--securemesh_site--reference--group-003.md#canonical-2122201310121302-0231330013102333-0211122333302320-3232232300231011-3103212210030233-2001123202003300-0330023113023231-1310222232123210): complete subsection reference.

- [dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-1112211302213333-2121212323023303-0132333022332033-3031100231032310-2222023310031032-2220031322101033-0132102210013212-0011332331113300): complete subsection reference.

<a id="canonical-1230012333033112-2100220013102233-1113332020112013-0321120021202311-3212310112220233-0212332112122103-3211331122121102-1322103113030331"></a>

<a id="canonical-0313033320132202-3301101210321311-2211311202002312-0211132213130222-0023221013300100-0131323312301122-0130311101330132-0203213010110222"></a>

## fixed_ip_map property — stateful / 030303031331 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

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
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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

- [interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-1033211230113111-1322321313232033-1222313322200311-3113231013011220-1330313110121001-3212030302113103-2213321120233220-0232132301020012): complete subsection reference.

<a id="canonical-0023222202000332-1123212221233102-2032132022011111-0331333301203032-3330021313330103-2222002121031223-3303211300300132-2320311000200233"></a>

## Next pages — stateful / 030303031331 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site--reference--group-003.md#canonical-3331021120111202-2002321133002112-1130222300233323-3222111102300322-1012001301202211-0322133111033202-3032330322031210-1023311030032112)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site--reference--group-003.md#canonical-2122201310121302-0231330013102333-0211122333302320-3232232300231011-3103212210030233-2001123202003300-0330023113023231-1310222232123210)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-1112211302213333-2121212323023303-0132333022332033-3031100231032310-2222023310031032-2220031322101033-0132102210013212-0011332331113300)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-1033211230113111-1322321313232033-1222313322200311-3113231013011220-1330313110121001-3212030302113103-2213321120233220-0232132301020012)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3331021120111202-2002321133002112-1130222300233323-3222111102300322-1012001301202211-0322133111033202-3032330322031210-1023311030032112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120322310002021-2302222331313000-0313302233121010-0303133310220312-1330033303102212-0332321310121010-0321210100221033-2110213100312010"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 003013221310 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3200323330022333-0333323130310022-3111203112120220-3333032330231333-3323001330033113-3002332321203333-3302113301131223-3223211032320122"></a>

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

<a id="canonical-3212212303213302-2113033320203011-1110120220303011-0030022313031213-3301332200330231-2330333132110313-0132133231133111-0223020310331213"></a>

## Direct properties — automatic_from_end / 003013221310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020200032220213-3223011120320330-3310001022203001-3001011001333131-1132032332102231-3023032023033313-3302001103230032-1212123311003103"></a>

## Next pages — automatic_from_end / 003013221310 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2122201310121302-0231330013102333-0211122333302320-3232232300231011-3103212210030233-2001123202003300-0330023113023231-1310222232123210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021131001003301-0233011301023202-1130312000003121-0202013222222132-0102200312323022-3033312231012032-0231220210320310-1212031032101123"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 002103033332 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1330031220101333-2222133233022330-2200111110313313-3030123302230321-2130332310030320-1122312131231321-1031112212311003-3031120322123302"></a>

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

<a id="canonical-2303201133322200-3203133103100120-2101321123122032-2123000133030303-2300120230311303-2331113121101020-3220110113101223-3012201021020201"></a>

## Direct properties — automatic_from_start / 002103033332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032032111033232-1023031333001302-3021101121323121-2300031013123021-3232232003222213-3103312230321333-2333023332212301-2332211211330202"></a>

## Next pages — automatic_from_start / 002103033332 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1112211302213333-2121212323023303-0132333022332033-3031100231032310-2222023310031032-2220031322101033-0132102210013212-0011332331113300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000302220112312-1113023003320233-2031002133111312-2312011201332121-3212211112002132-3110300301000102-2310120311211132-1202123210033030"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 231223133331 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0032213200201323-2300121003320303-1121020321323331-1121103131132230-1233321030201322-3123021113102110-2332101202010331-0111001120010113"></a>

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

<a id="canonical-1223222012000001-3232120010003312-2001212101221112-0113102133113001-3023323333320322-0002333023201333-1031111213222303-1002013232230300"></a>

## Direct properties — dhcp_networks / 231223133331 / 3

<a id="canonical-3011232122222000-2002001310012220-0322322020133012-1010203223022132-0002333312021130-2201131001312332-1210020132310333-1113013120321312"></a>

<a id="canonical-3100110121302313-0330102303032221-3302121221000020-1323002021110233-3330323221031120-2133132131300131-2331202330210130-0322132130322120"></a>

## network_prefix property — dhcp_networks / 231223133331 / 4

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

<a id="canonical-1000000031023302-3023023210003232-2131322121112223-1201320213320311-1111000212311331-3133200003213103-0000211230332012-0121100102122031"></a>

<a id="canonical-2213002231322010-0022022102113222-1102121310010231-0020220001310301-0232302002231310-3331021201030103-1011233122233200-2032102311322233"></a>

## pool_settings property — dhcp_networks / 231223133331 / 5

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

- [pools](resources--securemesh_site--reference--group-003.md#canonical-2001333023100122-2321101121010123-3122321321010103-0232223302330322-1230232122232303-3120232313312123-0333311130130100-3203131210212101): complete subsection reference.

<a id="canonical-0331030221112300-3103033011202232-1030331023322131-0310330333231130-2022210121232211-0133011311322313-2200202323121222-2122230010333012"></a>

## Next pages — dhcp_networks / 231223133331 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site--reference--group-003.md#canonical-2001333023100122-2321101121010123-3122321321010103-0232223302330322-1230232122232303-3120232313312123-0333311130130100-3203131210212101)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2001333023100122-2321101121010123-3122321321010103-0232223302330322-1230232122232303-3120232313312123-0333311130130100-3203131210212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001223033112121-1310103023210333-3103212311311332-1330022120221312-3322233123022233-2111200013300102-1322133020303000-1301333320332322"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 013103232021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-1112211302213333-2121212323023303-0132333022332033-3031100231032310-2222023310031032-2220031322101033-0132102210013212-0011332331113300)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0303212002002121-0202203000033112-2002020310201020-3011232210002100-3021130032120313-3123112110302003-1312230010203300-1201322220102323"></a>

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

<a id="canonical-3103133013003231-1300302312111321-1213232310232311-1310311011212013-2100222213220131-0011320231110031-1232011032133202-1121100111310110"></a>

## Direct properties — pools / 013103232021 / 3

<a id="canonical-2012220332210200-2023011003021210-2212222202323222-3103101003032312-2131122013230010-3133310032022132-2322232032030022-2120300302210231"></a>

<a id="canonical-1231033330302002-3311321321010330-1103100301032101-1101112221221000-0200332111003320-1132021103230011-0100201301212133-1220131223003112"></a>

## end_ip property — pools / 013103232021 / 4

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

<a id="canonical-2313230000013231-3033220122310200-3302330211221123-1013131012331103-0212203312022222-2312333320320212-0012122030023210-3300030111203220"></a>

<a id="canonical-0011001333301113-0031321121101310-0321002231003303-1102303312201333-1110012101133000-2223111233030002-2331202122010122-0310120230312011"></a>

## start_ip property — pools / 013103232021 / 5

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

<a id="canonical-2321232221321003-0100322032210332-0200013313333112-0323200111133331-0311020022203302-0232332310120312-3023112323123302-0120032011013232"></a>

## Next pages — pools / 013103232021 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-1112211302213333-2121212323023303-0132333022332033-3031100231032310-2222023310031032-2220031322101033-0132102210013212-0011332331113300)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1033211230113111-1322321313232033-1222313322200311-3113231013011220-1330313110121001-3212030302113103-2213321120233220-0232132301020012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320321111302133-2313003010213030-0233302211003110-3330131333010211-0131033002020010-3012310311030313-2131100011212300-3201022131301220"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 100021310000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-3032102210100223-0110232120222231-3211121211101103-3112121133330023-1203000200111012-3132023012302021-0330011120130122-0002021202120133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0203120023130303-3312323333130300-2203102130201222-1122212300323200-3130210213212111-3121322302110313-1123131203031122-1332000120030310"></a>

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

<a id="canonical-1003103123311203-2333223222033113-2232213112221023-1113220303013100-2201223333222202-3110300221313310-1020030033321030-1001022121321000"></a>

## Direct properties — interface_ip_map / 100021310000 / 3

<a id="canonical-0023222100212021-0331121000300002-0031320002022232-2132301103001120-2201001100112130-2230210000211010-2232321021202211-0101211331301321"></a>

<a id="canonical-3123030220123023-3300003102102022-3302031032322211-2102230003131121-0222010123132122-1132300213122233-3203003010322211-1130202032302030"></a>

## interface_ip_map property — interface_ip_map / 100021310000 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

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
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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

<a id="canonical-0233210303202312-1300220022203033-3301331122201002-2110230123323323-2100022223103013-3313331102122323-3312330000131002-3130102013021210"></a>

## Next pages — interface_ip_map / 100021310000 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-2131303301112113-0223211232310202-3132301003332221-2233332033301021-1223010202312232-3312333022001030-1032202003333322-3332011033332211)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3011030111020311-3300213332323121-3231230031321023-0301121230030300-3100021313302312-2201021132221213-3123112332332103-0233110332301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201323321321313-0003330002133301-1110202100310213-2220022122131213-3222110101313222-3030121021112002-0023032111031313-1110231200100122"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.is_primary — is_primary / 022133232130 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.is_primary

<a id="canonical-2130312313210321-2003330020330030-2302021031033201-3030300131230031-2112321023212221-2021210132213003-1131320132000133-0112203113233313"></a>

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

<a id="canonical-1222222033331333-2310221210232111-3233311322330023-2213120030003111-3201323203201303-3331103022130123-3331220310103131-3000313220230033"></a>

## Direct properties — is_primary / 022133232130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130323301222303-1201103130111033-0102011322110301-0132330131112230-3301112122223323-0022202113333121-2002113133333330-1220303303113310"></a>

## Next pages — is_primary / 022133232130 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1020021101002011-1000113331021323-2023213103023313-0102213210333203-2100120231331211-2123023102011310-0013200131000330-1011301211211103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333213320230013-3122021020003301-3212222131002101-1231323202130021-3332032330130323-3202203113302002-3321000211121133-3102032211121013"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor — monitor / 002311111033 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor

<a id="canonical-1311011203132213-0000322233213311-2322120202302031-3303121320200132-3101012000231230-3002200102010003-3303102202201211-2232033233221303"></a>

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

<a id="canonical-1230301203032222-3323111123102031-3021033111303200-3111231003023331-0021210023203223-3200202003311233-3133212130033120-0323110000211333"></a>

## Direct properties — monitor / 002311111033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122100231223023-0202120331330231-2003023202030333-3033203032302302-2100222333232103-3220131000202022-2120323320312032-2212200323121322"></a>

## Next pages — monitor / 002311111033 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1303013201121301-0130130300113220-1020321001031001-0213012011113330-1103230113211333-1303231201331133-0221212120232313-2231322111002312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212012333221100-1203300110330313-3323333010003321-1203010031220301-0302323201112012-1202331333300222-2300000313101333-1001222030232200"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled — monitor_disabled / 303310020111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled

<a id="canonical-3323121121232201-1220323001031011-2213032033203313-0130003210101002-3022110002031211-1201021202131011-2221111202333221-1231110010230122"></a>

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

<a id="canonical-1123213000213001-3011302130203313-1212002330110302-3032312023202233-3312223031020212-3020033021300222-1102032003230311-0322312110021221"></a>

## Direct properties — monitor_disabled / 303310020111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320132322322210-1212330101232113-0233333223211031-0000302130222200-0133233332120102-0313233211011223-1232222010322121-1202231320130203"></a>

## Next pages — monitor_disabled / 303310020111 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2112233221302010-0212003311133123-3303222201001222-1011223313303222-0030211210333221-1022322202310310-1002110233021312-3330132111213123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012103130201320-3221221301300113-0110000011333203-3312300331230303-0031312000320131-2203211122020333-0131033322130223-3332331001001100"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address — no_ipv6_address / 303113212330 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address

<a id="canonical-1103211331021131-2302122210201333-1001111200322120-1333230332133002-0200321032130113-2211331012030300-2011232303122102-3013320103111333"></a>

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

<a id="canonical-0131312101332200-2123012303322022-0233333030223020-0203212001132112-3033123320100313-0313110011013331-0012212102332202-2323333001220230"></a>

## Direct properties — no_ipv6_address / 303113212330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111001311302210-3122132220021110-1112233333332311-2121101332003023-2112321121221032-0030310120111012-0232330131131330-2332200233102110"></a>

## Next pages — no_ipv6_address / 303113212330 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1312002312303010-0021100322330231-3103212331202213-3012013002111132-3203123230100320-1221132011312223-0310301021313303-2001301030221133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300231332101331-1100033022213103-1111132231233302-2003222013121312-2003312001223200-0133203001023331-3221031033310203-2231212120310122"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.not_primary — not_primary / 132231110011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.not_primary

<a id="canonical-2113110110022022-3022102102322020-3333021231121321-3302330231223133-2210312330231012-2012331122012001-0000231220212130-3122103312111323"></a>

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

<a id="canonical-3330330313312001-2301233103001022-1021013333313112-0013012303301110-1011313020133231-1223013301331323-2111330321323002-2111201330012022"></a>

## Direct properties — not_primary / 132231110011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110330202133130-3133101330023021-2032332322001133-3103331113221320-3113103023122111-3022130123310101-2330120021121320-2213132233131212"></a>

## Next pages — not_primary / 132231110011 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3221212220111121-0221231002312232-2211313013313322-0200322021111302-1003113223232130-0111332001032000-3112323122221001-3311321022032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322120322300022-2210021232223313-2302020211130230-3322213032000203-0212312320320120-2320220130310133-2111311122101331-1200010211003232"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network — site_local_inside_network / 100003310233 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network

<a id="canonical-3121201221101101-1102012011000201-0031201123233212-3222133320201110-0011021233121002-1033103023313113-2301312300122021-1131013300002010"></a>

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

<a id="canonical-1010133003313002-3212132232332033-2111013303020132-1222012320102223-1210113213213312-2132121010302011-2311303120301012-2332322033100023"></a>

## Direct properties — site_local_inside_network / 100003310233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320323112200132-3323212130030001-2111201200003130-3313011011300012-0030221132033032-0023213003003132-3202213303012010-1201201310113221"></a>

## Next pages — site_local_inside_network / 100003310233 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0302131022023122-2201211001033022-0033312030231003-0231210331200321-0311320000313032-3301231031020302-1132323102000021-3202320201212200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100203011330033-3103332332033221-1223003322130221-0033112003212233-2331322132331223-0131321311220201-3111030230122101-1210103201300201"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network — site_local_network / 320120000000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network

<a id="canonical-2020332031032130-2233331023323201-0003021131220330-2322331021130311-1223012120023311-1121301112121202-0322130200320333-2110212012112000"></a>

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

<a id="canonical-3003332000020302-2103320002221300-3203312031130330-0123323232010221-0321203223331303-3231000321100021-0031210113023200-2332222030311312"></a>

## Direct properties — site_local_network / 320120000000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323111320310110-1120230003030002-1132002200001113-3011012130003212-1312212313110231-2100201321121133-3113333202000233-3313230131012132"></a>

## Next pages — site_local_network / 320120000000 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321321111203011-2121323131222021-1301233003221010-1111321320220330-0022120103113000-2232310210002323-3132033230322021-1031220010033210"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip — static_ip / 132232031013 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="canonical-2013333031130003-3021120002010000-2002131103002122-3013021002201031-3102030111222301-2211310201203130-2321000311032332-3311320321332310"></a>

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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020200303233130-0313213221002202-3300031333100203-0302122112301022-0032303101023301-3303231221013323-1031332010112320-1221233101010130"></a>

## Direct properties — static_ip / 132232031013 / 3

- [cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-3032300233030002-3113011321302333-0320321211303111-3022212002201020-1332202031223020-0221120113311023-2220303003123223-2030200202320031): complete subsection reference.

- [node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-2122110133112122-3020303322010020-0010303003122320-3100030332210223-0211321021130103-0303333123311221-2021321320320132-2103313312200212): complete subsection reference.

<a id="canonical-1000121001031230-2010322302032101-0202311020133211-2321333023003102-3113032302132212-1202212032333112-0211213312132300-0111023122011100"></a>

## Next pages — static_ip / 132232031013 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-3032300233030002-3113011321302333-0320321211303111-3022212002201020-1332202031223020-0221120113311023-2220303003123223-2030200202320031)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-2122110133112122-3020303322010020-0010303003122320-3100030332210223-0211321021130103-0303333123311221-2021321320320132-2103313312200212)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3032300233030002-3113011321302333-0320321211303111-3022212002201020-1332202031223020-0221120113311023-2220303003123223-2030200202320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122323100133301-0202003310310211-2021303212032201-2033331302003201-1333203301221103-0123201010333122-1231120230000233-1102313211323321"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip — cluster_static_ip / 322323103110 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-0301302130000132-3013102232032033-0330031212211002-3303300112102220-2102301310021023-3132233120023212-2110310233000320-2202013333210100"></a>

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

<a id="canonical-3202220123121223-3232001020222011-0113011131203300-3320221203211120-2212011021032331-1300233301331100-3210110120300102-2132012002002112"></a>

## Direct properties — cluster_static_ip / 322323103110 / 3

<a id="canonical-2001110222202210-1230112311110022-1033033203123100-0032033100133322-1332233000002102-3113310310111112-0301301200321201-3032330330212213"></a>

<a id="canonical-1110012120103312-1121322102101120-0313331012222332-3232222112030310-1200030123021203-1303022300131112-1022013030121132-2100031213001321"></a>

## interface_ip_map property — cluster_static_ip / 322323103110 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
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
      "ves.io.schema.rules.map.max_pairs": "128"
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

<a id="canonical-3001231230030301-0211200031303212-2011002300102230-1122303222231133-0203203222210001-2202130113030313-0203323222020130-1323111001002222"></a>

## Next pages — cluster_static_ip / 322323103110 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2122110133112122-3020303322010020-0010303003122320-3100030332210223-0211321021130103-0303333123311221-2021321320320132-2103313312200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222023121111100-3213023320323010-3032033000103112-0102011202222111-0032310200311013-2302033133310122-1010201213201213-0121013030313100"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip — node_static_ip / 230203113321 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip

<a id="canonical-3222122212332010-0011102003003330-1031012320000321-1210012002123112-0112321300000033-1300000103100300-2200102013231001-0320020202031102"></a>

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

<a id="canonical-3310123100002221-0233331323113101-3201020220030023-3100111331321201-3332102113101221-1031310121331301-2331012333320300-2121031032011301"></a>

## Direct properties — node_static_ip / 230203113321 / 3

<a id="canonical-2211212103302200-3121210330020123-2330213232123312-0022132333232332-2201130221122333-1103320110232022-1013213102022211-0320112203330233"></a>

<a id="canonical-1303123223030011-0100213230200331-3112132221311112-2211002313121212-2002010220130112-0121330012233032-3112033230311332-0100010000330113"></a>

## default_gw property — node_static_ip / 230203113321 / 4

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

<a id="canonical-3103002230223231-1132233312000132-2202100020303121-3320030033301032-3221221201130212-0211101333131122-2332123131013233-0303220130322123"></a>

<a id="canonical-3100122031113011-2322230310210300-1322112301012233-0121200100121210-1111213322221013-0322101333310312-1012203232131303-0033311133221130"></a>

## dns_server property — node_static_ip / 230203113321 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3120000330103202-3031223031200233-0322222012200211-0300002232121220-2232031320311231-1133123121200202-1032001113322030-2223013001130201"></a>

<a id="canonical-0231203023121113-1211032102312333-1220321301111033-1131232320100132-0322211031313003-1221300210031233-0301013333303313-1013001130322123"></a>

## ip_address property — node_static_ip / 230203113321 / 6

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

<a id="canonical-2203111330332213-1020323011213010-0201013022001020-2101300200131220-0013212013013310-0110001103021103-1300223121333331-2302230001131302"></a>

## Next pages — node_static_ip / 230203113321 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-0312201332321112-1103221103313331-3101303123320303-2132101111111321-3300301332023210-1313031321303022-3112221000023123-0030002230201023)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001113031100212-1100301000113111-1302331222122312-0330110021323312-3213130023303233-2021112222332000-1212203200330012-1023112003130211"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address — static_ipv6_address / 202020231001 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

<a id="canonical-2131012233122132-1011222013320332-1013210031230223-2302101132021300-2210002122111020-3102221312111202-1133033212301321-2010032001311320"></a>

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

<a id="canonical-1301232130032111-0300002321203001-3332013003320213-1023221313323310-1232221221202330-2300220220203103-2020110122202221-2010102031133131"></a>

## Direct properties — static_ipv6_address / 202020231001 / 3

- [cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-1000111233132313-3131130113320001-2321010330322313-2302033130112012-3023313121021323-3310022011301000-3322332200132033-1301000123300103): complete subsection reference.

- [node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-3201220122111230-0221210030303011-2122132322201131-2322132122010331-0012022111101031-1022022330121110-2210203023202023-3020101203311021): complete subsection reference.

<a id="canonical-0201112220222322-2200322203331130-0300120330030130-0233222202232303-0002110031020120-0112012001330032-0112232020313231-0230002131100312"></a>

## Next pages — static_ipv6_address / 202020231001 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-1000111233132313-3131130113320001-2321010330322313-2302033130112012-3023313121021323-3310022011301000-3322332200132033-1301000123300103)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-3201220122111230-0221210030303011-2122132322201131-2322132122010331-0012022111101031-1022022330121110-2210203023202023-3020101203311021)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1000111233132313-3131130113320001-2321010330322313-2302033130112012-3023313121021323-3310022011301000-3322332200132033-1301000123300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111323231312200-3212122211002013-2212211211322010-3201213101110023-3122212223331200-3000223120332121-1201221230032212-1313001111330113"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip — cluster_static_ip / 220201013320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-1002331233313332-2032032030323222-1301313303300300-1223101131001030-0133030300232031-3300201122222112-3223320033212323-0102210201220010"></a>

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

<a id="canonical-2111003133203001-2222210313300111-3032310132333102-3031033233330223-0232233011220011-1131111001231331-2312033031111310-0123311301110233"></a>

## Direct properties — cluster_static_ip / 220201013320 / 3

<a id="canonical-3100231000120120-3033132333111003-1031201122223233-2232122322301212-0332320302323211-1223001022030331-0311121200013133-1300132131310121"></a>

<a id="canonical-1123023301103330-0032112300213300-1100131211103232-2203003220332220-3022311103333220-3301220111222033-3201321210001011-3011013113131033"></a>

## interface_ip_map property — cluster_static_ip / 220201013320 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
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
      "ves.io.schema.rules.map.max_pairs": "128"
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

<a id="canonical-1132231122311011-2000301031330100-1211131002130211-3122000123033231-3112001120023120-2120110331311113-1122301230311031-1312102033002003"></a>

## Next pages — cluster_static_ip / 220201013320 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3201220122111230-0221210030303011-2122132322201131-2322132122010331-0012022111101031-1022022330121110-2210203023202023-3020101203311021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311032133301033-2330203320021032-3002223223320301-3321223131001031-1213301230220120-0310100302003120-2333212221303001-0033322231011032"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip — node_static_ip / 201112000333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-2000323231121333-2133101122021032-1112021021221023-0011302130111122-3112323123231020-1233110120111011-0333001223023110-3311233000210030"></a>

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

<a id="canonical-0303333221313021-1021222311333330-3130222103302200-0223332031230320-1033132300321011-3000123201130300-1031120121033132-0313100132200322"></a>

## Direct properties — node_static_ip / 201112000333 / 3

<a id="canonical-0213333312320111-1323033120202323-2302031323312002-3021132221013222-0002102132311323-0030232320233010-0313100230023311-2220300321132332"></a>

<a id="canonical-0211113020131220-2220123123030323-0312023220102210-3112330100022001-0111302133001231-1200132212013133-3330300101330101-3233131012110330"></a>

## default_gw property — node_static_ip / 201112000333 / 4

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

<a id="canonical-2102231112022020-3003120032101303-3023331210301133-3213110320003211-0020130120112212-0001020131011010-0300311332001320-2323103131301202"></a>

<a id="canonical-2031221120122310-2302120111003120-2101100133010023-2231033000012103-2001013133100202-0332331311111120-3231102021233111-2022312310110121"></a>

## dns_server property — node_static_ip / 201112000333 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2320020213110321-0111112211131131-0313000320301102-2233020012232012-2011312101220312-2200223303331202-1332200303321012-2110211102221020"></a>

<a id="canonical-0221031021312301-3230322103133130-2031210122332013-1312000033033103-1201110133003112-1333230031230102-1321211112231131-3000320113122103"></a>

## ip_address property — node_static_ip / 201112000333 / 6

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

<a id="canonical-2022032022111300-3131003121310313-2003320103010023-3330022020200201-3111130010100202-1333103330130021-1032030312201133-0232211211321120"></a>

## Next pages — node_static_ip / 201112000333 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-1130030332311000-1001230200331300-1311030213002100-0131102120101212-2003200211110333-0010011233100110-0123022132221110-1232310223000210)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0302221212130113-3210032322320332-0131033110232323-2232202101222201-0120031231212101-1031312131332010-0122112232022201-1011320133013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311031122203223-2330001220220211-1200310022011133-0313102322002313-2313112313232003-0221200003130300-2232100020231132-2303323303003001"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.storage_network — storage_network / 121223100130 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.storage_network

<a id="canonical-2103230310230032-2233122030312032-1002103230222232-2111302120331210-0311222303311223-3230112232130231-1223212231103212-2221203013201230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for storage network.

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
storage_network = {}
```

<a id="canonical-0331221131001230-3113121030333202-3331201303213133-1122012330002202-3121322230101313-0102211223130210-3123220232010113-2202032331032023"></a>

## Direct properties — storage_network / 121223100130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112010023112000-3023300100211321-2011121011210333-3321033121313110-3101131222323320-3320121101201120-2010003213310230-3030101013030132"></a>

## Next pages — storage_network / 121223100130 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2030323210031111-2122220222013201-1221133130030112-2201300133321132-2131032320202110-0220030001210203-1202200303313102-0312132100233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303201302100333-0031302300221332-2233323200203201-0223023000020223-3203322212012303-3031202313013311-3223100122222133-3101003310020131"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.untagged — untagged / 331301330101 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-0210221233213102-1330203111311212-3210222123320201-0012103110001032-2110021221311000-3001212122330023-0213223233122222-0110312322332131)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-2233331111010003-1110010202110013-2311100311031210-0130311200003031-3331331300232031-0020110203132131-2030213131322111-2012032022331022)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- custom_network_config.interface_list.interfaces.ethernet_interface.untagged

<a id="canonical-0301331133230213-2011023212033110-3133230220223000-3012300230203001-2230212013333103-0330320023121002-0103023110300131-1322102302312113"></a>

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
untagged = {}
```

<a id="canonical-0220320012131322-3112033032031211-2223300302221220-2223233331220300-2221211121012121-0010021331222320-0300023312001301-2100030230010133"></a>

## Direct properties — untagged / 331301330101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221021103312332-0012121131112320-3102130322323221-0211213111333300-0211010020330221-3312233220132302-1311120012021213-0120012110112012"></a>

## Next pages — untagged / 331301330101 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-1111100000201032-1023333230021202-2220033101112022-0001002200333203-0031022101133301-3110031021010231-2031210321022333-0332232013020233)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2303112131223200-0330102000132110-3111303323230203-0033123201010303-1032300122232203-0301333310233012-0102331232312221-0331010221101021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112201312003330-3113321221313231-0001010312101230-2220201332020111-0221123032203012-0002013013010002-0010310232321203-0101210131030320"></a>

## custom_network_config.no_forward_proxy — no_forward_proxy / 213301332103 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.no_forward_proxy

<a id="canonical-3033333101130130-0311113313032123-1203213102322021-1301013300310013-3231001310203012-0300023323321123-3000300321231211-0012022302331111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-2302331102003122-2111032102010310-0131330233021333-3313302322233002-0023233310012113-3021003121012232-3331102122123233-3202212313222130"></a>

## Direct properties — no_forward_proxy / 213301332103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030200202232113-3300311030010000-0133331330132110-0313012001033020-2011212331011203-2231102331223110-1231011111110203-3232130333122232"></a>

## Next pages — no_forward_proxy / 213301332103 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1321020212311300-1111231333313020-3332223223121213-3103332312322312-1011030011231203-3212022023111201-3230122000201132-0311311333310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301101123103332-3010201221100311-3122013333122313-3203213320301322-2220332301133022-2130122012231311-1130123310212232-0000112030000111"></a>

## custom_network_config.no_global_network — no_global_network / 212210000132 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.no_global_network

<a id="canonical-3103312011210112-2230132221023122-2103133133233313-2022123012132200-2030110121212010-0323111332333001-0202122230121200-3131302330223210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-2111102322333013-2332213200120211-2030322000100003-2212012020230211-1013131113222313-0310322011010203-2220010103300223-3112000131102121"></a>

## Direct properties — no_global_network / 212210000132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212111010303100-2211121103320321-1103331030323310-0031232030312310-1232020332311233-1031201022020101-3000213020120101-1101211201321112"></a>

## Next pages — no_global_network / 212210000132 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1001313103113102-1000033221100312-2111311202310323-2312033332301000-0123101211223010-1023230323131023-2220113220032212-3113332130103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101310212222033-2212231321010232-2133331000203311-2003002033010203-3120222200001331-1013201230211201-3032230011132033-1333020211133011"></a>

## custom_network_config.no_network_policy — no_network_policy / 011100023033 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.no_network_policy

<a id="canonical-1322312130131120-1113130023002232-0023120023303013-1130313200313031-2012303012322101-2102222221031010-0312010333033230-1213130233322001"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-2231100131123013-3222122032011230-2112333111330333-1220311010122032-3130133121021312-1200020231011131-3302200231310120-1121120210021201"></a>

## Direct properties — no_network_policy / 011100023033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202302102010020-3213210303130132-0102320111131021-0311332031013211-0113033322023320-0131310100131130-3102202330323222-0232033101101333"></a>

## Next pages — no_network_policy / 011100023033 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230022223122212-3002323312302032-1113023012210112-1222203123221131-3021223123223112-3200210123003020-3010221133020203-3302200200011322"></a>

## custom_network_config.sli_config — sli_config / 222102331011 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.sli_config

<a id="canonical-1221312312123320-0103030110021301-3310022001112203-0213122002123300-3003333212130302-1323130020133011-0030011223030203-2133122033013013"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
sli_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203021333310300-1210012121102122-0131332223330330-2101302013013302-3013322113322221-3012002300003020-3033211002131321-0032110210313130"></a>

## Direct properties — sli_config / 222102331011 / 3

- [dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2010310131203210-2222303211202301-1112322113013202-3002000332201320-1303121113022031-3011011310021320-1300210302203301-2201331023333312): complete subsection reference.

<a id="canonical-1333212231110233-3313013302112231-3200223200130001-0311221332331131-1330311202013020-1230321123133301-2103230103331213-3203232212032213"></a>

<a id="canonical-0310022103133211-3201213103310312-3000120210033230-2202233111201323-0123323232210130-3202131111001311-3132033220211120-3123333200001023"></a>

## labels property — sli_config / 222102331011 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
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
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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

<a id="canonical-0122103111333233-1102333232122230-0202321003323120-0020300102132000-1013332021203321-3303210330311013-3312113213103302-3100033311022030"></a>

<a id="canonical-0011101030202301-1023011211013001-2112000031321311-1330103122323330-1321211012201322-0222131200333123-2313011223133133-1102222312122302"></a>

## nameserver property — sli_config / 222102331011 / 5

Type: `"string"`. Optional.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2321203110111323-1233031301311022-0001030130222221-2300031202113331-2201213331232321-2302220023130101-2302021103112000-3213200013201213): complete subsection reference.

- [no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120233231233001-0332310232113311-0102210212203021-0222021012130132-1231133131321110-3230103332132301-1001022030023101-3321221121010132): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-3131231112201021-0131203121303003-3001003131030301-3210020201311013-0221103033321320-3030222313012101-3322202211003223-2012203122133113): complete subsection reference.

- [static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013): complete subsection reference.

- [static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311): complete subsection reference.

<a id="canonical-2301202111113222-3001320213300220-3321320200022310-0110323321022033-1330223033101210-1212211122012222-3032222321223200-0332321103203313"></a>

<a id="canonical-2203110113103110-0102322110313012-3202123102223300-1003032223131230-3323111321301030-0111120223103322-0113132003302232-0013000030003023"></a>

## vip property — sli_config / 222102331011 / 6

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-0112332223111110-0333210222123223-0201001103232111-3030200120210110-2321131102120210-2222212003002000-2133311210130120-1011322120230332"></a>

## Next pages — sli_config / 222102331011 / 7

- [custom_network_config.sli_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2010310131203210-2222303211202301-1112322113013202-3002000332201320-1303121113022031-3011011310021320-1300210302203301-2201331023333312)
- [custom_network_config.sli_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2321203110111323-1233031301311022-0001030130222221-2300031202113331-2201213331232321-2302220023130101-2302021103112000-3213200013201213)
- [custom_network_config.sli_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120233231233001-0332310232113311-0102210212203021-0222021012130132-1231133131321110-3230103332132301-1001022030023101-3321221121010132)
- [custom_network_config.sli_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-3131231112201021-0131203121303003-3001003131030301-3210020201311013-0221103033321320-3030222313012101-3322202211003223-2012203122133113)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2010310131203210-2222303211202301-1112322113013202-3002000332201320-1303121113022031-3011011310021320-1300210302203301-2201331023333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101101323322232-2323222233120032-1310313133023210-3202100333313333-3231303232120131-3311200301102230-0001200023112311-3332123233031313"></a>

## custom_network_config.sli_config.dc_cluster_group — dc_cluster_group / 112031231313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.dc_cluster_group

<a id="canonical-1000020133210023-1302033220112300-0103002003221230-2033331113332102-0322232322010032-2203103210112221-2230122322102021-0320321010031232"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020221001131201-3110112322333302-3133311121203213-0301321302200303-1112003231022213-0330022101323120-3121033302133030-3332122010020012"></a>

## Direct properties — dc_cluster_group / 112031231313 / 3

<a id="canonical-0210302000131311-2202133120332002-2030122003133131-3010032222322023-0101012212200013-3310102132222110-3112300322311103-1301203310202011"></a>

<a id="canonical-2112111312331323-1230211023311022-1202001111321301-3003003303333302-3000001202213122-2301003021231021-2212321030131200-1301120021202330"></a>

## name property — dc_cluster_group / 112031231313 / 4

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

<a id="canonical-0330022212100021-1000012003222220-2213301301222003-1212132012101002-3300023100110332-3232212332211101-2131133332010132-3210111312331221"></a>

<a id="canonical-2313032001120132-0032020003101302-0010223023110311-0230210321010013-2311201322200112-1020012230112033-2210030113021332-1111100320221312"></a>

## namespace property — dc_cluster_group / 112031231313 / 5

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2312103002331133-0110122032323001-1132103311203222-0222232132220010-0332203202010111-2110301320202102-1121013232002031-3123221003130301"></a>

<a id="canonical-2301012101210233-1333032210003233-1322030210013133-2221323230120131-3023300112211233-2202022033003201-1220303102202332-3221010213102121"></a>

## tenant property — dc_cluster_group / 112031231313 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0131101310122321-0002322301332110-2123031020231133-3323030210311113-1010333022212132-1101001022323330-1333330201103223-3310002322010202"></a>

## Next pages — dc_cluster_group / 112031231313 / 7

- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2321203110111323-1233031301311022-0001030130222221-2300031202113331-2201213331232321-2302220023130101-2302021103112000-3213200013201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111210020123221-0121310030030121-1220203132323111-1231321131000333-3220110120333033-1012030200302100-0022303230000303-0013203220203200"></a>

## custom_network_config.sli_config.no_dc_cluster_group — no_dc_cluster_group / 013301110030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.no_dc_cluster_group

<a id="canonical-1112312322221032-3202020333111101-2001120331330202-2222100001010300-2223231113120321-3232031321120000-1001303121302012-1231130232320312"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-2310310132123102-0133111302102301-1202332112230031-0230003003013200-2103330013233100-0231000010221100-3332003331313203-1032201123222122"></a>

## Direct properties — no_dc_cluster_group / 013301110030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311033201330023-3332123112032012-2031332220302231-0313111212313020-2032303131203121-3120223102100102-2211001201100322-3323012020300220"></a>

## Next pages — no_dc_cluster_group / 013301110030 / 4

- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1120233231233001-0332310232113311-0102210212203021-0222021012130132-1231133131321110-3230103332132301-1001022030023101-3321221121010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221202002103032-1101003033011330-3033023011333122-0122030013123010-1201331112323200-2220112112303231-0022123020302132-2213211323033011"></a>

## custom_network_config.sli_config.no_static_routes — no_static_routes / 330311212031 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-0301321111133210-3223032221022100-0032211013121203-0221122203110202-2102211201202121-2122120301100303-3230002131213112-1322112320211113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-0312300313133011-2233322133310023-3303333210101103-2030100113332123-3033320000332022-3230222303321222-2121022032131100-0221130220222331"></a>

## Direct properties — no_static_routes / 330311212031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003212103100200-0103121131313021-0121212113132300-1231002011001320-0200003010030021-0111202330131322-2310230001102310-3332013123113320"></a>

## Next pages — no_static_routes / 330311212031 / 4

- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3131231112201021-0131203121303003-3001003131030301-3210020201311013-0221103033321320-3030222313012101-3322202211003223-2012203122133113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031131123320203-2331111030113300-0002120123311222-1230320033002021-2011131031213213-2002031302232322-2312131212330000-2112313220012223"></a>

## custom_network_config.sli_config.no_v6_static_routes — no_v6_static_routes / 230121211233 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-2033231302311222-1322210121332311-1233320111113323-1033220310112231-2320303020003310-3111131033111220-0332121223101011-1121332320002321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

<a id="canonical-2111232110021031-0202031232100313-0131003310130230-0130112303103311-2112112331221200-3121010131331011-3301230001103221-0210313231200320"></a>

## Direct properties — no_v6_static_routes / 230121211233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323211210133010-3020011013002110-1221101002301000-0011132033011130-0233231310203331-2223031312231020-2100232302112031-3123311202321030"></a>

## Next pages — no_v6_static_routes / 230121211233 / 4

- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322110112201333-0230131312321221-2020233323122221-1012312232010333-0210130002232032-2003232132031113-3023030022133132-0210101231210231"></a>

## custom_network_config.sli_config.static_routes — static_routes / 033210202022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.static_routes

<a id="canonical-0233103320310211-3131202121000330-0113032000300230-3201230232302003-3333213202023203-3323210310021120-2011211310200233-1003100220213111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323320300102221-3011122223231012-0331322321133103-1332322012230103-0101333000330103-0033313313021032-1021202011020331-2131032220232312"></a>

## Direct properties — static_routes / 033210202022 / 3

- [static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310): complete subsection reference.

<a id="canonical-2132110332331213-1303113022323300-0300013000330222-3333322203301020-1011121031033311-1303230321010100-0131030330211112-0200121003233233"></a>

## Next pages — static_routes / 033210202022 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222131133111110-0233320320031331-2101320222203002-3003000301131310-3001210211311233-1332323222321130-3210330333312001-1131110102222122"></a>

## custom_network_config.sli_config.static_routes.static_routes — static_routes / 002200201022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-1230123031130332-1000101310130020-2202103020330220-0212120312203113-0121032132101130-0322303122230312-2031201100120211-0110331202133230"></a>

Type: `"object"`. list nested block, Optional.

Static Routes. List of static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013200033023100-0301133332203112-0330303211210230-2202201320223103-2001132101331202-0021120030330322-0330300130222003-0112211012122121"></a>

## Direct properties — static_routes / 002200201022 / 3

<a id="canonical-0323001232320220-2221322002021321-1113333003203310-1101110112321202-2012220121202231-2132203321330121-3310022013103110-1222023323330031"></a>

<a id="canonical-3103232233112132-3113223203211033-1301130301311033-0221322133021312-2120002122200200-2302020233320322-3100230000021233-0300121131202210"></a>

## attrs property — static_routes / 002200201022 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site--reference--group-003.md#canonical-3333101021010211-0023301000223302-0222033100020212-0020310313030310-2302122331300000-1330033310313200-0101122130101332-3202133030331002): complete subsection reference.

<a id="canonical-3333022031022220-3232010121321333-0100000223132022-0231221213221100-2232310023331202-3033131233110113-2231311310313122-1232103123100221"></a>

<a id="canonical-3132301322121302-2201230321303220-1203000111102020-2221313210021133-3221322231122013-2120331100031213-1102111211003221-3122203013011300"></a>

## ip_address property — static_routes / 002200201022 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3200112122221113-0113101101321303-0122022313113102-1111120330111110-2023330131320223-2211301231010001-3100330022130232-2203303101000231"></a>

<a id="canonical-1003012030001123-1320111301133033-0303300021003212-0033120210301100-1010220203020330-1103203120103033-1302111021112032-0323100133212332"></a>

## ip_prefixes property — static_routes / 002200201022 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site--reference--group-003.md#canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102): complete subsection reference.

<a id="canonical-3130101021031100-0231211210210112-3113311233001020-3020012111133002-2103203222012102-0220331220220121-3312233121000023-0303010120130130"></a>

## Next pages — static_routes / 002200201022 / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-3333101021010211-0023301000223302-0222033100020212-0020310313030310-2302122331300000-1330033310313200-0101122130101332-3202133030331002)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3333101021010211-0023301000223302-0222033100020212-0020310313030310-2302122331300000-1330033310313200-0101122130101332-3202133030331002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331102000133002-1322111130021121-1222012032232212-2221132002221013-1030203301000320-3120232130102221-3223302023223220-3123110001303321"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — default_gateway / 322213121022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-0100321213212033-1020110131211103-1021102022121303-1333023232223302-0232020210221123-1233123133333123-3232111022323321-3321313100220213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-0332111310120232-0330323120211200-2301012133323333-0221023022011000-0202002120201022-0210210210303321-0010313110222130-0323303201102131"></a>

## Direct properties — default_gateway / 322213121022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223031302300031-0332010321123002-0231211302320233-1210120221130122-1302313230003122-3023222132130113-1102200111300111-3332013112330123"></a>

## Next pages — default_gateway / 322213121022 / 4

- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300112201012332-0011003201302331-2110203223020111-2011113101330203-1013021033112003-2223302110100312-2133321133031300-3322132220021202"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — node_interface / 133023102231 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-2131303320333211-0213221212013210-1111031302001201-2132303303222323-3011303231030021-3133102130231231-0030010132010012-3203223203023023"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132331210110333-1211031120103230-0231321303203331-3102013330121323-3111231021121023-1103002122220113-2331023200000102-2331001011231121"></a>

## Direct properties — node_interface / 133023102231 / 3

- [list](resources--securemesh_site--reference--group-003.md#canonical-3223230233020331-0110210300001113-3121120323113221-0111102302123120-1302001021230201-1012103022031000-3021113112002000-3331130032231210): complete subsection reference.

<a id="canonical-2212332003003331-3003012023120312-2110212010110310-3323000032133100-1311211201121230-1032312101102330-2012231202133323-0203211102021212"></a>

## Next pages — node_interface / 133023102231 / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-3223230233020331-0110210300001113-3121120323113221-0111102302123120-1302001021230201-1012103022031000-3021113112002000-3331130032231210)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3223230233020331-0110210300001113-3121120323113221-0111102302123120-1302001021230201-1012103022031000-3021113112002000-3331130032231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202132133121101-3213202133202300-3330223323222200-3220113313013021-1102130031332311-3030001121030323-0303011133311013-3320103130021201"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — list / 211202110021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-0331000312013131-1102133311322311-2200201223322211-3311202331213002-2100131132003302-3311123333103312-1312033022130122-0122321022201200"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033020131113130-0223122002312023-3011322330202110-1131300303123300-3220113232023033-1130132313211120-2002221202110131-1031130002330011"></a>

## Direct properties — list / 211202110021 / 3

- [interface](resources--securemesh_site--reference--group-003.md#canonical-3320113130112221-3000112103010323-1222031101001130-2103322212033320-1031300111122221-2111213022210102-2003303211132323-0222300130332232): complete subsection reference.

<a id="canonical-1323101321132201-3323211213212121-0321213101010131-0000223312301303-0202113313020201-1231321233220232-3023222320300113-3203010223103021"></a>

<a id="canonical-0020120022013032-2311220313312302-2111331000113120-1201031230303103-3331333312022133-0002312321003131-1331232102300103-0101111302310221"></a>

## node property — list / 211202110021 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-0101200123312220-3013122310210222-2100133303333320-0012102112200310-2131202231202222-1310211233230302-2220011120321300-0331201232022213"></a>

## Next pages — list / 211202110021 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-3320113130112221-3000112103010323-1222031101001130-2103322212033320-1031300111122221-2111213022210102-2003303211132323-0222300130332232)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3320113130112221-3000112103010323-1222031101001130-2103322212033320-1031300111122221-2111213022210102-2003303211132323-0222300130332232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221221122330013-1230311131001221-3132231020102311-2103302223210321-3113022102022031-1130300021301132-1103222330100122-0100222213133213"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 200230232332 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0112333301233311-1111313311012231-3333112122200131-0320112133311122-3102122323321132-3232202123100222-2232233333000012-2333132320022013)
- [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1331331032001310-0022200020112110-0320130111033133-2210222203010332-0103002321112122-3231211132202131-0210123312002010-1023232031003310)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-1320333210321000-0231103133213231-0231011101322300-3010230012223120-2321313230202102-0230201230223313-3320013000112210-2110131300120102)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-3223230233020331-0110210300001113-3121120323113221-0111102302123120-1302001021230201-1012103022031000-3021113112002000-3331130032231210)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-2030121333102222-3101030030113322-0111121230032201-3121123213313030-3202201020200213-3132222010011121-1213122300123133-3000012211323030"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133333333323033-2123001212021000-3112110301313010-0221311113032323-3011002020032110-3110031012301031-3033033031222311-0131001130211303"></a>

## Direct properties — interface / 200230232332 / 3

<a id="canonical-3220213303133332-0221033120110203-3302313011002210-0021331201301322-2231302230023231-2200003130133322-0123200113210333-3330331310222230"></a>

<a id="canonical-3301332323313211-1011100313202100-0322321330110103-0233010011211302-2012331310323000-2022203030123021-2220131023103112-2331122303113210"></a>

## kind property — interface / 200230232332 / 4

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

<a id="canonical-1221111002030132-0322311031033321-2023002333230023-1232313301221001-0330101320230312-2200101111130201-2021203303323001-0311031110310123"></a>

<a id="canonical-2110013230212320-0023211300222322-2203323133030210-3023110232331213-1132021123031030-2202012032223221-1333111101321122-1030313203220231"></a>

## name property — interface / 200230232332 / 5

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

<a id="canonical-3033102300022332-0000103232012022-2101201021223201-1021011111230200-1310112131020111-1212320313312122-1101021330021020-2300120131300001"></a>

<a id="canonical-2303223030323323-2232013232211322-1330131002331033-2313313131223030-1012022111001120-2112112323330331-1303322233333001-0333010100230131"></a>

## namespace property — interface / 200230232332 / 6

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

<a id="canonical-3122301033122312-1223122131330202-1100211030203300-2020230002120211-0200201232013110-1013322302303333-0200012313321012-0321112100111331"></a>

<a id="canonical-2113223120132331-0113032130012113-1330233323221102-3132230320232000-1320010131212222-3322110010201230-2330311113322203-2301332003213230"></a>

## tenant property — interface / 200230232332 / 7

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

<a id="canonical-2113013230021120-1002020021110233-3313212231030212-1103300123332310-1330223311311100-1303122000032022-2230120301222322-2330201223313103"></a>

<a id="canonical-2023102100213101-2122000301301122-1311023000003030-3111320221101323-0110312011011212-0203303001223003-3120101212101310-1212232232333033"></a>

## uid property — interface / 200230232332 / 8

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

<a id="canonical-3331022313121212-0121311021122322-2300320302200300-2001202312213110-1110101201131321-1300110220003110-3020013001213013-1113112023113202"></a>

## Next pages — interface / 200230232332 / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-3223230233020331-0110210300001113-3121120323113221-0111102302123120-1302001021230201-1012103022031000-3021113112002000-3331130032231210)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112033313022012-3101210222223022-0202102122202220-0233130012301213-0201210233230131-2302001220212321-0311113133011301-2100113210203231"></a>

## custom_network_config.sli_config.static_v6_routes — static_v6_routes / 001112122023 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-2111310301211013-2301313322200033-0201321230211221-3112133120233000-0231010013111030-3301212301000101-1332022312031021-0033322003221231"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201333213130031-0230003312110101-1122122211301232-1103301101130012-2123122323012030-1221231213220222-1031021230031221-3013230121201030"></a>

## Direct properties — static_v6_routes / 001112122023 / 3

- [static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332): complete subsection reference.

<a id="canonical-2101303200002113-0210203020200021-3230201003020232-0211202322303203-3302222113030121-3330123220330001-1112312313203110-0331011231113132"></a>

## Next pages — static_v6_routes / 001112122023 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133333202321012-3220021132103110-2003011301022210-2203332213213033-0100013300022213-1310032012020212-0113011333211221-1130201331203113"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — static_routes / 033132200221 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-0030201022032223-3201032102310001-0212110223320130-2120000211111200-1031312333310211-3301302112310113-3133100230131210-2222213232211301"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101320203221023-0102203003232301-1320321211031122-2110101001332001-3130113213200320-1201033001330102-3133112030233020-2021213130312110"></a>

## Direct properties — static_routes / 033132200221 / 3

<a id="canonical-1012100133230021-3001121021232310-2321320320021220-1212321020103012-0211132133000300-2301013333031013-3020023320123322-0332213002310230"></a>

<a id="canonical-3013322002013222-0313031122331002-0020310333111200-1311212111201220-2012203022003101-1011212101320130-3230201312001133-3320322323300031"></a>

## attrs property — static_routes / 033132200221 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site--reference--group-003.md#canonical-0021101012232322-1223312302000011-1231212032102001-3120211321321311-2132232221222013-2101220322221102-3233021111221221-1102201110032031): complete subsection reference.

<a id="canonical-2000213100231103-2331000012220230-0223020032233100-1203131131002320-2233031301030022-1233130320201020-0022323213031222-3010110130303100"></a>

<a id="canonical-1222020111322323-3220010133313023-1122001213001311-2023011311312211-0013100302223331-0302011312101122-1230023131132130-3121323312032232"></a>

## ip_address property — static_routes / 033132200221 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1120112320333000-1021311330200301-1221213011011231-2223111102100202-2321130200032003-0123210120300232-0032112011331233-0330110200312203"></a>

<a id="canonical-2330002230230222-3011203210111111-1100113031023323-3020210330213013-1130312213302020-2220320232113033-1320313230110100-0332122130022210"></a>

## ip_prefixes property — static_routes / 033132200221 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site--reference--group-003.md#canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120): complete subsection reference.

<a id="canonical-1320010101000100-0323332223233122-1033001232000030-3303102311100213-1002301203132313-0021133101123230-1023230123322333-1002131210301233"></a>

## Next pages — static_routes / 033132200221 / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-0021101012232322-1223312302000011-1231212032102001-3120211321321311-2132232221222013-2101220322221102-3233021111221221-1102201110032031)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-0021101012232322-1223312302000011-1231212032102001-3120211321321311-2132232221222013-2101220322221102-3233021111221221-1102201110032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232212321131031-0121022222013320-1201300332213313-1003311020323031-1112003233000203-3213020212303323-2123202310022120-3122100130002320"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 130320001202 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1333203222203013-3220233003321210-3121210232021000-1322013133222231-0233203323221122-3130213222232023-0223321131101330-0330332111133102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-2102021113233303-0220003012020021-2030203030022021-1313332101202201-3303301130232000-0300230203232200-1120010002313222-1203313213333131"></a>

## Direct properties — default_gateway / 130320001202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231220201033001-2210303321010010-2200321320320011-0302002111202212-3230020112232111-2222213223130312-1321132201223200-1333223301002311"></a>

## Next pages — default_gateway / 130320001202 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111223300302102-3020023013130201-3311120123121131-2000222020122322-3103211223221032-3312333221032113-0200231220332102-3223121311313211"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 123330131001 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-3200300132322113-0230030202321201-1333030103301123-3113200001102123-0133231110111033-3331312232332010-1131011102312310-1020032000222322"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311031331101023-3131300230102332-3011220123212312-3022300010122022-1332310123333012-3012220112212030-1331111033300203-2312300113303111"></a>

## Direct properties — node_interface / 123330131001 / 3

- [list](resources--securemesh_site--reference--group-003.md#canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230): complete subsection reference.

<a id="canonical-1012110332302101-1001213333301201-3000000132332133-1300020310320001-3303301322310113-2321132021221231-0121121123033212-1223031122103213"></a>

## Next pages — node_interface / 123330131001 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302033232103203-3113101121023101-1010323200113103-1203320201211232-3331323230330312-3130231012311003-0130312311312111-3000210121102012"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — list / 123012031301 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-1120230030012303-2001212101220221-0022123001100111-3132132311223211-3222110120120231-0013233013000012-0000033121232121-3211001200022322"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112131123000211-3013011220132013-3212122212303113-2313011212002301-2223310033023102-1001121300232233-0220012030111302-2102210010310002"></a>

## Direct properties — list / 123012031301 / 3

- [interface](resources--securemesh_site--reference--group-003.md#canonical-2023010321210223-1111301300101111-3121132103310001-1223200001000303-2222030103003022-0022123210023003-2331112331033013-2323012330230210): complete subsection reference.

<a id="canonical-3030330301102000-1210020333002031-3310030110033021-3302001033031020-3330102220101102-3312212323021300-2131133000312313-3313230302231211"></a>

<a id="canonical-3110130020011200-1000001011301010-3021233100003030-3130330002331112-3121332033222010-0131031121100313-3033102320201022-3132200022102302"></a>

## node property — list / 123012031301 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3303132101010133-0012321020132033-0200020312320110-1103222231132113-1113200201112321-3032322312312212-1113100322013101-3123301303202100"></a>

## Next pages — list / 123012031301 / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-2023010321210223-1111301300101111-3121132103310001-1223200001000303-2222030103003022-0022123210023003-2331112331033013-2323012330230210)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2023010321210223-1111301300101111-3121132103310001-1223200001000303-2222030103003022-0022123210023003-2331112331033013-2323012330230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012220010223302-3223030210322223-3013032023320231-0013123122321303-3223223230200002-2223121301103332-1221210303300220-1331312220212313"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 333311200130 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-3230101021120233-0312333123201221-3033003012302020-3311211000102123-3031201001232311-2322313202101300-1232311330003333-2320032030322202)
- [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-2120003233130110-1032303010231322-3113200321112031-1123011101120313-2301032102222231-2130032202000002-3333231200232232-1203112303021311)
- [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-1120021320130220-1102003303203133-1202100031000331-0110231223221200-0002223133111331-0102312303312222-0022233022010200-3030203320222332)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-3121232323112312-3002332323322023-3023301331001200-1231210302310202-3011132202201010-2322300222222213-2322310301010010-3001110030213120)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-1222000303032110-1230102100323001-3323220213233031-3332131003012001-0133303102030231-0011122313323000-0120100200101122-2233032311100031"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130303010203333-0132303121230232-3323021232302122-2112211123322323-0032201102030011-2211221322323132-3320010122220302-3123211133030011"></a>

## Direct properties — interface / 333311200130 / 3

<a id="canonical-0232111311232131-1300233211032011-0001323132102120-0330323312222002-1331312111200103-0110020320112223-3232203231123331-0213111110322300"></a>

<a id="canonical-2312201321231103-0010133000311313-3130212221213031-2133121320310200-3321102212220022-1010122001002030-1031032210313202-1301323011331011"></a>

## kind property — interface / 333311200130 / 4

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

<a id="canonical-1201123032300310-3213123010022301-2111102020323221-3201202101231003-1332022331222131-1123131323130202-0331211102232321-3010232313230332"></a>

<a id="canonical-0111103133220222-0230021132002320-3233123322120313-3222321013021133-0132312123323013-0232231122001002-1102110300021021-2120200311212320"></a>

## name property — interface / 333311200130 / 5

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

<a id="canonical-3323223202103332-0330003223130202-3010313023210221-0110231230330203-1012032311222002-2121123233310321-3331220200200232-1201120023011211"></a>

<a id="canonical-1020202120302103-1323133223110133-2312000200312022-2032030320332220-2131202200103110-0120110221130121-0220212111233103-1301302113302220"></a>

## namespace property — interface / 333311200130 / 6

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

<a id="canonical-2031312121221103-0333201310012122-2320320111132023-3023100332313332-0112213012221002-2030300021322130-2333001010233330-0231302213120332"></a>

<a id="canonical-0010010331220211-2110003110230120-1213300000221011-1103022222312122-0323321313033200-3122100321120003-3322213132220113-1021302020201321"></a>

## tenant property — interface / 333311200130 / 7

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

<a id="canonical-3303232002023010-2123023103131021-0311331022100113-3010212210032000-2011332313303322-2313200012203022-0211312323103012-3320011212100103"></a>

<a id="canonical-3120111223011210-2233101323212103-1213110001120120-1121313220310222-1322130202212332-0332133032000030-3032332012201133-0232103110103211"></a>

## uid property — interface / 333311200130 / 8

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

<a id="canonical-2102111210220003-0330303031000321-2012031120231311-1202033301032211-0301221301213211-3321211030121032-1230203010231232-0323313222321303"></a>

## Next pages — interface / 333311200130 / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202013320323322-2302220122122301-0113030330212103-2022123230200131-1121300010321212-2003120210121213-1301120222103303-3230111223313322"></a>

## custom_network_config.slo_config — slo_config / 320001103200 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- custom_network_config.slo_config

<a id="canonical-1332331002001001-1013113132220202-2132200030301021-3332011103110232-1023310301321202-0301033032212200-3121302222131300-2122311312010322"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121300002321032-1131030102130013-0032332122232101-1132010011331301-0310323111122222-1103103113311312-3301130011011203-0233311001003032"></a>

## Direct properties — slo_config / 320001103200 / 3

- [dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2100030221002000-2312112211033132-3211002030023102-3232011300232020-3120131002303132-0201312033300312-2203013033100223-1111012013232320): complete subsection reference.

<a id="canonical-2303212331122012-0202220223320012-1323101023222012-1312010322220332-2321113301321011-1110132312223120-3320310110211002-2222301333022022"></a>

<a id="canonical-3030021221100202-3300200213100220-1131003213223101-0123110003210322-1021112302233331-1132021221022330-3300331123322110-3033110302212200"></a>

## labels property — slo_config / 320001103200 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
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
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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

<a id="canonical-3110022030203000-1120223010033323-3203120022122312-0020320312132320-1220033310220000-0332023202120013-3000111123222121-1001101101100200"></a>

<a id="canonical-3213021311320202-1132233210021230-3203233222130133-2233101233122220-2323331010011322-2313331112120003-3011110331102113-2100231301332101"></a>

## nameserver property — slo_config / 320001103200 / 5

Type: `"string"`. Optional.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2211310112211203-3023202231313300-3230102112323221-0003322323202103-1103331232002220-2223023103313220-0113132033200003-3122023022231120): complete subsection reference.

- [no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-3300302132133100-3322212230012131-1112301131002213-3022013112211210-3120012311111221-1230011321220332-3302133111201132-3022033303030221): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-1300222022212012-0102110031001023-3013032102103021-2220113213320122-0321333110312122-2301132331013302-0213220232221131-3003020122133130): complete subsection reference.

- [static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022): complete subsection reference.

- [static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333): complete subsection reference.

<a id="canonical-1121110122012012-0222302320000323-1032302212033230-1011321213010311-1133323232222211-0023312311222312-2111103022301331-3221000013310211"></a>

<a id="canonical-2303200101001220-0101110211012333-2202211133213010-3201231202302010-3213101133000302-0001211130001132-3112132112002211-2012220000102112"></a>

## vip property — slo_config / 320001103200 / 6

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-2203331310313113-0101220120333132-2221111123220102-3100031022202310-1113101313000003-3213022310020322-3011103102011301-3220001300212002"></a>

## Next pages — slo_config / 320001103200 / 7

- [custom_network_config.slo_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2100030221002000-2312112211033132-3211002030023102-3232011300232020-3120131002303132-0201312033300312-2203013033100223-1111012013232320)
- [custom_network_config.slo_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-2211310112211203-3023202231313300-3230102112323221-0003322323202103-1103331232002220-2223023103313220-0113132033200003-3122023022231120)
- [custom_network_config.slo_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-3300302132133100-3322212230012131-1112301131002213-3022013112211210-3120012311111221-1230011321220332-3302133111201132-3022033303030221)
- [custom_network_config.slo_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-1300222022212012-0102110031001023-3013032102103021-2220113213320122-0321333110312122-2301132331013302-0213220232221131-3003020122133130)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-0232330332331102-1133303311300232-2010001000330022-1121131113110121-0013201011132020-1330301333221030-3223310303222020-3212230102133333)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2100030221002000-2312112211033132-3211002030023102-3232011300232020-3120131002303132-0201312033300312-2203013033100223-1111012013232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230200302120013-3012030012120321-3323220323300321-3021333001210122-0010202011220102-2223023303021210-3312023201200232-0010212111012010"></a>

## custom_network_config.slo_config.dc_cluster_group — dc_cluster_group / 203201020222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-2221310222200112-1031302213311311-2012033211320022-3010233122031103-0303332213312223-2322032330032120-2312322202310133-2301120211132230"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021013322010011-3102331210101002-1231021012232301-2331111330110301-3033002322121211-1100120010110333-1032303120221223-1201113030121320"></a>

## Direct properties — dc_cluster_group / 203201020222 / 3

<a id="canonical-1121122111300133-1033201121103011-1222003331222322-0020011131010223-1013331312202110-3331222011112033-3331133320203220-0023010321211300"></a>

<a id="canonical-3302123203230203-1303300123031221-2330203121023002-2022323221133112-1302100203112231-1200103220010120-3132033120000131-1300313232231032"></a>

## name property — dc_cluster_group / 203201020222 / 4

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

<a id="canonical-0001303230121110-2131013212222232-2003300030020111-1311011023003023-0330012021101303-2302200033201332-1231120012133111-1002312032333210"></a>

<a id="canonical-0110201323123202-2202201211321123-0033011313301213-2233003301213022-0223210133010332-1222101003012100-1022201133122233-3021320313213102"></a>

## namespace property — dc_cluster_group / 203201020222 / 5

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1212303002202220-1003133231301110-3300222112111121-2003230121023320-0122112031112003-2212221313112013-1022222232231002-0320213102202220"></a>

<a id="canonical-2121121212223302-2122211323133001-2101111023210012-3120310110210201-0030203121121222-0211320021200300-1113021120330032-3202021000312032"></a>

## tenant property — dc_cluster_group / 203201020222 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2331320212002220-2231231310111330-2111003210212113-3021033111212030-1101102102302203-0202213123030030-1030222211213110-1223020022200200"></a>

## Next pages — dc_cluster_group / 203201020222 / 7

- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-2211310112211203-3023202231313300-3230102112323221-0003322323202103-1103331232002220-2223023103313220-0113132033200003-3122023022231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333203023130320-3020120200330232-0022002212320102-2331012111310213-2221200130110222-3202103303111021-1103112123011020-3211301022221131"></a>

## custom_network_config.slo_config.no_dc_cluster_group — no_dc_cluster_group / 321012130313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-2300300123211131-2002221213030310-1011233130323300-3033300112312011-2132030022331301-1131203003221111-1003301230202213-2300333313320303"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-1001332022020321-2102100321131123-0122223020000120-1131120300200001-2200212221301103-2232301230121212-2222323122232210-2123113031223323"></a>

## Direct properties — no_dc_cluster_group / 321012130313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120303310220323-0300201002103122-2023323331111230-3301320300310113-3033330123220232-2101000122302313-0132210300012131-3030012320330203"></a>

## Next pages — no_dc_cluster_group / 321012130313 / 4

- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3300302132133100-3322212230012131-1112301131002213-3022013112211210-3120012311111221-1230011321220332-3302133111201132-3022033303030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110022133301223-3223233233110332-2313013032101302-2123322132031203-0133103320333200-2103223020322011-3033212233331012-2301323232231331"></a>

## custom_network_config.slo_config.no_static_routes — no_static_routes / 012003023121 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-0312222212133232-2123321103323020-3330121001122021-0001102032203021-0123321103120321-0011301312220203-0203201202110102-3303323202223010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

<a id="canonical-0121112022031100-3202213210013223-2030210132103211-0000232113003110-1320002312032100-2210012231101300-3032203030331210-0221303332032220"></a>

## Direct properties — no_static_routes / 012003023121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200320201312103-3333231320202021-3133022211221321-2013010301211333-0300021123121010-2113302032121230-0132131210321030-3231022020302212"></a>

## Next pages — no_static_routes / 012003023121 / 4

- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-1300222022212012-0102110031001023-3013032102103021-2220113213320122-0321333110312122-2301132331013302-0213220232221131-3003020122133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201301112033003-2333222211220230-3223131333031320-3230232221331023-0200001031020222-3000233122012023-3203320033301001-3313123031121010"></a>

## custom_network_config.slo_config.no_v6_static_routes — no_v6_static_routes / 011201002311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.no_v6_static_routes

<a id="canonical-2001020220310202-0332033311232313-1111003110231220-0021132100301201-0022310320121002-1011101213332111-0033023031220323-1022313031010001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

<a id="canonical-0101131022131112-3330303313132123-3223023330122030-2320131013131010-1312212302011221-3330111300332330-3223112000000133-1121021011002122"></a>

## Direct properties — no_v6_static_routes / 011201002311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310201210203101-0222012220132023-3122032122100033-0230332322320230-1321023122333302-0001231320203331-0212121113013033-2131201302112111"></a>

## Next pages — no_v6_static_routes / 011201002311 / 4

- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)

<a id="canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233310013011112-0310120212131110-0033000301011313-0002331320131133-0300030131010322-3023102210313331-0210200302130032-0013332321223300"></a>

## custom_network_config.slo_config.static_routes — static_routes / 033122002300 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-0221210100300032-1323301000110203-0003200133200010-1121101320113033-1203000110031002-2332032100223330-3220112032322030-1011202210112330)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-0333201331313102-2002003312232331-2123110122101233-3032121221303003-3323310312332103-2230011000003222-1111013311021313-0232110102113230)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-1032112231012201-0211322232030132-0302222020321232-2322231023000213-3102302203112121-2233121130211202-3310233301010120-0130020313303312)
- custom_network_config.slo_config.static_routes

<a id="canonical-1033003123123002-1102020233203110-1313031031331002-0322323003333131-1300332303130320-3012023130030333-3010021032232231-0213030130031233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```
