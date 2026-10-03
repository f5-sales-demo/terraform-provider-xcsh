---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1220112221001030-1230303101000021-2130201200021133-3231300310010222-1231332133012010-1213321122121232-0011003311203211-2033301201311330"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 203202302230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1100220321312331-2002300003203111-2131323133312013-2333013221211333-0021213331022131-1033023132103230-2302322100130203-2213331230300132"></a>

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

<a id="canonical-0003030123113130-0333320303333313-3000002211013313-0211220002201123-3333113323201303-2021001021000231-3000223221211301-3020021202133310"></a>

## Direct properties — dhcp_server / 203202302230 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-006.md#canonical-1302131120011112-3033011310230132-2230313102233312-3320331302211201-2013102121210233-1110132010313002-0000010001323100-0032223012102323): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-006.md#canonical-2210212130203131-2231303112210110-3122112302320220-0010001001302000-2033012223333300-2132032111120000-0103020030231211-3311023302212033): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000): complete subsection reference.

<a id="canonical-1211312233221223-3012312110013331-0222132112123123-2121012301102101-2012331133303302-3001203231233103-1203301220111331-1331012003130301"></a>

<a id="canonical-3002320223033003-3210122220101123-1302302333223233-1011010112013232-0121303301211012-1120101313310231-3010312203032012-2212300022020020"></a>

## dhcp_option82_tag property — dhcp_server / 203202302230 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2310003223101111-3223130213330332-2333110302023311-3002033010111313-1323113322210003-2210212202103200-3001313201332330-0120000112023211"></a>

<a id="canonical-2210223323333221-0211023001013333-0332320020220013-3020020132211323-3232330133023312-2222033111132302-0321131001101003-0121020010310103"></a>

## fixed_ip_map property — dhcp_server / 203202302230 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-1012333133323023-1212231013003102-1213023313202203-1300210111210112-0131020330032312-3232222120113223-0133022012000223-3131033223322223): complete subsection reference.

<a id="canonical-2113331100233223-0222213000322332-0222003231013130-2032012000302232-2012321112200332-2200311001220031-3320220110001001-3332020232322313"></a>

## Next pages — dhcp_server / 203202302230 / 6

- [baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-006.md#canonical-1302131120011112-3033011310230132-2230313102233312-3320331302211201-2013102121210233-1110132010313002-0000010001323100-0032223012102323)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-006.md#canonical-2210212130203131-2231303112210110-3122112302320220-0010001001302000-2033012223333300-2132032111120000-0103020030231211-3311023302212033)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-1012333133323023-1212231013003102-1213023313202203-1300210111210112-0131020330032312-3232222120113223-0133022012000223-3131033223322223)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1302131120011112-3033011310230132-2230313102233312-3320331302211201-2013102121210233-1110132010313002-0000010001323100-0032223012102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222110021032322-0223300021212110-1132032012033323-0220322003010013-1302030033213010-2303300012221001-2332313021011002-1000212001203000"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 032001020021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3113033330331221-0231311033003211-2233023001202201-2301223213133330-2130020132322023-0122123121121120-1322323331300333-2113322210002233"></a>

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

<a id="canonical-1212201313230003-0201031311231111-0213130220132302-0301111232122120-2022133012003220-2232130102320100-3011210212120332-0002300112031210"></a>

## Direct properties — automatic_from_end / 032001020021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101030120201101-3310231003103223-2233222133101032-1130222132122222-1122230021200113-2100111311132303-0202222220111032-1021023311112131"></a>

## Next pages — automatic_from_end / 032001020021 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2210212130203131-2231303112210110-3122112302320220-0010001001302000-2033012223333300-2132032111120000-0103020030231211-3311023302212033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232223300301000-1311000233200332-1220121210300211-2220022132302103-3030303113302223-1011330021210302-3122201010321011-1211100222332101"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 201112232221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- baremetal.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3203233031222331-2231313202213030-1233201030222310-2203032310132032-0021003121032103-1030032213211122-1203123003000111-2231333031132013"></a>

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

<a id="canonical-2332022202030020-1232320332012001-1122101300003232-3203112220321312-0012032000110132-3013011001313020-3123333200001000-2201011132212310"></a>

## Direct properties — automatic_from_start / 201112232221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333221112310211-3330110133130333-1110003000310021-1221001032012313-3313210122003302-1121120220231213-2210202212230033-1230011331321302"></a>

## Next pages — automatic_from_start / 201112232221 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010130302123020-2221001201100203-1021202210322312-3022023311131000-1312001113003313-1101011203110222-3230200121133131-0232211102320320"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 330232301021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-0220002100011211-1231001323123022-3210331030023003-0202032001212233-3133002331113110-0230011222222330-2003132232011302-3312122113031333"></a>

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

<a id="canonical-3222223321010221-0101103112212112-3203120100013011-0032003223200103-0200201102203221-3111020032010202-3020302020212231-1200013220222210"></a>

## Direct properties — dhcp_networks / 330232301021 / 3

<a id="canonical-1023132330121123-0232331120113233-2201232011331120-3222331032332033-0213312023020331-1312331011212321-1102122331000121-1030010333333113"></a>

<a id="canonical-3303133313320331-1031000131323122-1032020210201301-2222000312231010-3123011202302231-1012320222120302-1033313310100222-3301333200103102"></a>

## dgw_address property — dhcp_networks / 330232301021 / 4

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

<a id="canonical-1323013103201210-2203322321121133-1312321030333120-2113101002102222-0133011121002010-2202031110233203-3123203121230212-3123013013131131"></a>

<a id="canonical-0023221010300122-3332323321301120-0103013131202112-2120331100201200-1220312211103210-0310233023030110-0320113303020023-1100121132320002"></a>

## dns_address property — dhcp_networks / 330232301021 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0202000032313210-0021030303001313-1312010131121112-1230211121111210-3132022330003032-3312233120033300-3111232100201331-3003323330212131): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0103220202101231-3120121223321020-3000021200230120-0131010021132103-0003201102301301-0232233132101222-0103100321202330-1013120101130311): complete subsection reference.

<a id="canonical-1322013031023232-1330210321033211-0223331102330112-2102211220002031-2233330001200222-2203333030212133-1132000332021221-3023223000303133"></a>

<a id="canonical-0222023231022002-0303312031210212-3220313103012221-3113012002022112-1001033000020033-2311332331332101-1023033301110011-3132312132330033"></a>

## network_prefix property — dhcp_networks / 330232301021 / 6

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

<a id="canonical-2300020231321113-2302121033212220-1102310213111301-3301303111311332-1032221313020232-0020203231231301-3312000030101110-1022010233332111"></a>

<a id="canonical-0011001230322032-0223123333323002-2031211022002003-0100323130221320-1122221120131111-0022120021121020-0331102112310033-0200212331022212"></a>

## pool_settings property — dhcp_networks / 330232301021 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-006.md#canonical-2130113033332211-2002023321332130-0120332010310010-3120331100000101-0330310231230220-2112313132201312-3310321113033021-0023132230230212): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-006.md#canonical-0012322100133031-2311300111201320-2000131322313021-3303310310001111-3200212301021122-3030303013231200-3012120212231323-0003110032120022): complete subsection reference.

<a id="canonical-0133111300121103-0320030201220233-3223122333212310-3012121331130013-1123321300220132-0202130011312223-1311120033220202-2010101331332200"></a>

## Next pages — dhcp_networks / 330232301021 / 8

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0202000032313210-0021030303001313-1312010131121112-1230211121111210-3132022330003032-3312233120033300-3111232100201331-3003323330212131)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0103220202101231-3120121223321020-3000021200230120-0131010021132103-0003201102301301-0232233132101222-0103100321202330-1013120101130311)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-006.md#canonical-2130113033332211-2002023321332130-0120332010310010-3120331100000101-0330310231230220-2112313132201312-3310321113033021-0023132230230212)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-006.md#canonical-0012322100133031-2311300111201320-2000131322313021-3303310310001111-3200212301021122-3030303013231200-3012120212231323-0003110032120022)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0202000032313210-0021030303001313-1312010131121112-1230211121111210-3132022330003032-3312233120033300-3111232100201331-3003323330212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102201310003222-2233220313021313-0002003011102312-1313130301122102-0222113123000222-2130313212121103-0112230010212030-0012223310300212"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 130112113120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2322321133200131-0013022311100100-3003213112230123-1223203302220210-2122122132303331-3123220333221201-0100020011230221-0332103332130231"></a>

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

<a id="canonical-0321130221322031-3302122222230121-2011303122323211-1033220222023230-0113002213322010-3301012202211312-1133231231111301-3320113202020231"></a>

## Direct properties — first_address / 130112113120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032010311133001-3111021121323013-3120030013110211-1311313120103202-3111122202223030-1220321002110101-2233200021000331-2111122130230201"></a>

## Next pages — first_address / 130112113120 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0103220202101231-3120121223321020-3000021200230120-0131010021132103-0003201102301301-0232233132101222-0103100321202330-1013120101130311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133013323021221-0110012321223332-3333031323323123-2111003310300012-0313331323133101-1030331310133323-1223013330011013-0020032202331200"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 131321033113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0000021303112220-2121222032011011-3300211030211203-3002330031003013-2332032221120031-1023032322200320-0100133203012010-3323032033202013"></a>

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

<a id="canonical-3021031113100231-3012123310333132-1120021012033230-2001010203003223-3113212021330223-3211223032133123-3032012101320210-3322302323220200"></a>

## Direct properties — last_address / 131321033113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222203023311302-3210012103112033-3231110301211000-1223323022000130-0213030202013123-1002113311003213-2301120032232232-3222333233323033"></a>

## Next pages — last_address / 131321033113 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2130113033332211-2002023321332130-0120332010310010-3120331100000101-0330310231230220-2112313132201312-3310321113033021-0023132230230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221003011102131-3020112332330002-2303210130121323-0203200033233221-2023100112131120-2112102202020310-2313331120030011-2313213112012310"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 220101012202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-0010222220113123-0003123032033110-2033200223030012-1120202020130201-2212133322131333-0121012110200002-0130103132220213-0012121311312302"></a>

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

<a id="canonical-3323221100213331-3201232313211230-3022022102003122-2212033321031211-1200201023221112-3103322202021012-3123203201303033-2331021131230002"></a>

## Direct properties — pools / 220101012202 / 3

<a id="canonical-1120100220111330-2323222022310101-0021022231210121-2330132001333032-3103323313023222-0120233303003213-2022033211210333-3123231210120201"></a>

<a id="canonical-1111303230200101-1311310320323232-3302301332332002-2001111310130121-1003332333211310-0202003331022111-1011103232303310-0033112002321002"></a>

## end_ip property — pools / 220101012202 / 4

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

<a id="canonical-0130013222003101-2012230230001321-0122310011230222-3110323132231013-2012310303211332-2233103201330232-2212300222013001-0300122230213103"></a>

<a id="canonical-3312102301223122-1032302110001103-1011022011232311-1213123220012020-2000323331133120-2322103111023031-3122022010312113-3303201230231101"></a>

## exclude property — pools / 220101012202 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3022131003100222-2322101103321231-1300120232230003-3030132132020101-2110311032001230-0011121102031300-1100131133203131-2011322001011110"></a>

<a id="canonical-3132122332310000-2133320012000003-0102221302323222-1030100103023103-3311221123320202-0321011131100012-3330120102320011-0012222113213030"></a>

## start_ip property — pools / 220101012202 / 6

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

<a id="canonical-2101322012120130-1320201312100223-1001200121312311-2000031313312213-1012132310203130-0032331303201302-3022102323220000-1032011230021320"></a>

## Next pages — pools / 220101012202 / 7

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0012322100133031-2311300111201320-2000131322313021-3303310310001111-3200212301021122-3030303013231200-3012120212231323-0003110032120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301000311101132-3311322110202010-0033332320300201-0123203311200131-0132320003131113-3312320102213233-0210300322200202-1212111220322311"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 311102123000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1030322013012123-3312013323202211-0210322022312302-3111222221130220-0101322010033003-0202211101333131-0220301011001212-0210001003223032"></a>

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

<a id="canonical-2320230210213331-0330033023201213-1210000200031321-2102220310020111-0313122032101020-0222002101312121-3300230312001010-2132311100332203"></a>

## Direct properties — same_as_dgw / 311102123000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012222320211132-1223022101210101-3332330200123011-0102332220131132-2033011321021021-0230220301332212-3233222230011211-2122013100103221"></a>

## Next pages — same_as_dgw / 311102123000 / 4

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-2322311322101313-0132030032212111-3012112133021332-1121211233130211-2233121311323022-3120123122332130-1303230113102021-0322031030020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1012333133323023-1212231013003102-1213023313202203-1300210111210112-0131020330032312-3232222120113223-0133022012000223-3131033223322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213221111303312-0220030300022330-1203303303023212-1303312223111023-3013323131213012-2031331313303020-1331131200331231-3102112201302331"></a>

## baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 321002330233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1321311331322103-2113312211023302-3310333001301200-0030110230323010-3121102101203332-1121210231201313-0112202021101101-1101033030232133"></a>

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

<a id="canonical-2131103000231131-0310101220322030-3101031310322111-1000313110021322-3321113110111300-2121110021030220-3020022133232110-3203303000201331"></a>

## Direct properties — interface_ip_map / 321002330233 / 3

<a id="canonical-0223102013231113-3311333312102323-0332332002003102-2112021002310130-3233210332210013-3223233133100023-3003111112331211-3000022021030213"></a>

<a id="canonical-3121333100001011-3013303133313333-2200110133121320-0002201013101302-2132210233312303-2332303230333030-0023301011213031-1332321012003313"></a>

## interface_ip_map property — interface_ip_map / 321002330233 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
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

<a id="canonical-0233232030013013-2312030312232120-0203333002320220-0331113121020032-0000002332330003-0132312230312220-3010220020210123-2313120232112131"></a>

## Next pages — interface_ip_map / 321002330233 / 5

- [baremetal.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-005.md#canonical-3223303200132003-2132013013022112-0013021010122202-0122111221303030-1000333003103020-3112202230103233-1121011130103323-1313333330100220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2311111313301331-0003100322113221-3121001301113123-1112322020331333-0232203130212122-2332020132212131-3020132101002321-2233212102001123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311002002232031-0201232221123113-2121311223002232-1300113102032331-0120201321301002-3032121010210133-0200002001011013-2303332030222202"></a>

## baremetal.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 333032031302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3332311003111011-0320020333221003-3020322322030202-0121210330011213-0330332111310323-0031021313001021-2313030313021330-1132331123232020"></a>

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

<a id="canonical-2301012322320220-0211233132233220-0332311230223223-2322210011002003-1333313320032312-2233321003013332-2220003322113132-1111133020021203"></a>

## Direct properties — ethernet_interface / 333032031302 / 3

<a id="canonical-0122220132133101-0321102332013303-3312333300120231-2012122002333323-3302021333032111-1313221221122132-1211123100232110-3202122031201033"></a>

<a id="canonical-1130001232333313-3023211310133230-0210102100131321-0032203303310030-0030310023333202-0330233203233033-1202031203023003-3212123323333011"></a>

## device property — ethernet_interface / 333032031302 / 4

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

<a id="canonical-0102222330230130-0230203223023033-2120231202102133-1222303203313030-3222220122301231-2130111110113311-2010312231320331-2331310122221202"></a>

<a id="canonical-0120031003313002-1301331211013201-1330330310133332-0202000022201112-0320333103232000-3311031231012201-2233212333011110-0310112010211331"></a>

## mac property — ethernet_interface / 333032031302 / 5

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

<a id="canonical-1010020023012020-3113212030122210-0223323322311230-0213231102232033-3020320223002230-0131133202220133-0031131200131101-3222323222303232"></a>

## Next pages — ethernet_interface / 333032031302 / 6

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032300331131331-3100013021222031-1202103111213200-2233321001322201-1222123322022023-2330313201233222-2031233100332111-0322101011313111"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 020032133320 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0220021021221233-2302123031223212-1003321101131313-0200200120333231-1220120013123220-1312230131223222-0002020133210110-3000120210202131"></a>

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

<a id="canonical-1333211001110311-2112010003110002-3111112312030233-3110301103003102-2300001131211112-0212200110011303-1132031131133112-0202200222312221"></a>

## Direct properties — ipv6_auto_config / 020032133320 / 3

- [host](resources--securemesh_site_v2--reference--group-006.md#canonical-1032201112110110-3233101303131203-2311133223020310-0021003231020331-1023223103130030-3210001131102202-1303300303303303-1111332013203011): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031): complete subsection reference.

<a id="canonical-3220223102000110-3223321232210330-0331320221132302-3200311022223322-2323330011010133-0131221011011301-3000323331020210-0003123312002321"></a>

## Next pages — ipv6_auto_config / 020032133320 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-006.md#canonical-1032201112110110-3233101303131203-2311133223020310-0021003231020331-1023223103130030-3210001131102202-1303300303303303-1111332013203011)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1032201112110110-3233101303131203-2311133223020310-0021003231020331-1023223103130030-3210001131102202-1303300303303303-1111332013203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332112031100020-3232322123301303-1033333103013113-2331221221321231-1312123130103221-1230121023011120-1132123122300323-2213310030200320"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 113100301021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1032002201301331-0323013123122011-2302333000320002-2221002203333122-2102231121233222-3211201230212302-3330032201013332-1301133122233020"></a>

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

<a id="canonical-3122303003230001-0333333210110331-2103310010322333-3313100210131112-2202210220013302-1323331320121132-3230033010130231-3212023201101013"></a>

## Direct properties — host / 113100301021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120220320121011-1221122213123110-3013022101203333-0301123312023310-2312112312100310-3310320131223111-3200301023213300-0300202330313112"></a>

## Next pages — host / 113100301021 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203312331103333-3221200210022121-0122023301203212-1311331103321011-1313113200331212-2023333103330012-2302321313212203-3200210003202303"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 313231001030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2221123231100202-3011003121313121-1030321100321332-0210222311001221-2210231333013210-1210110313010201-3303331311110101-0232222321222022"></a>

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

<a id="canonical-3122233122331133-2210011331302101-3321230223321023-0131332023222312-3320320223202310-2331302220213021-2231130221321100-1200032210003211"></a>

## Direct properties — router / 313231001030 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033): complete subsection reference.

<a id="canonical-3331013121210311-2231133320100021-2022231132210122-1311013333331000-0000031231200230-0330231102021111-3003221320031321-3121032113232311"></a>

<a id="canonical-1313231200223000-2021133222001102-0202120333221010-0230020132013231-1003311312021030-3301103001011202-1221123012002303-3233031131011121"></a>

## network_prefix property — router / 313231001030 / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013): complete subsection reference.

<a id="canonical-2010300233200133-1130213120320102-3103212321122031-2122023111321111-0223333030001122-0223211033020113-0002010121212301-0303111002132032"></a>

## Next pages — router / 313231001030 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120333332020323-3021120311111213-0103210123030203-2323231130003310-3101121112220301-1003001010200132-0101330002223203-2033130230003220"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 233220120201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2010331212120132-0230022222001130-3112332300320221-2302121113132313-2003213313101012-3130330333313320-2210231332201223-3023131300003121"></a>

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

<a id="canonical-0033022321023202-2023332012230123-2302310201002210-1230113133313323-1211331103011103-2302131002101102-3022202132102311-3310032021300232"></a>

## Direct properties — dns_config / 233220120201 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2222032213100032-2120020013100300-0203021210131121-0201231303032322-1300031110220232-1021331002100011-3331121030201130-1012020210302022): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331): complete subsection reference.

<a id="canonical-2312031300233010-1303100323130202-1013321102223210-2323123302301113-1133221201003122-1220232100233012-0133120111222202-3211003230022230"></a>

## Next pages — dns_config / 233220120201 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2222032213100032-2120020013100300-0203021210131121-0201231303032322-1300031110220232-1021331002100011-3331121030201130-1012020210302022)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2222032213100032-2120020013100300-0203021210131121-0201231303032322-1300031110220232-1021331002100011-3331121030201130-1012020210302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101022231230300-0221103010111001-2223220120321021-3203212130002130-1132201010321212-0000030301013033-1020231222102110-0122331012011311"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 233021121120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0331202112102322-1101000302013333-0103033201222223-1302221021112323-3302203332113030-0230222133011031-0303211133012120-1231202312001232"></a>

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

<a id="canonical-1301322321232221-3030000320020313-2201232231130221-2123133211231130-3231003333230103-3100132113032010-0212120312310012-3323013023133221"></a>

## Direct properties — configured_list / 233021121120 / 3

<a id="canonical-3211013010330321-1313213023301223-0301123102023020-0003110321021130-1130112323001331-2101213122331210-0132332313023001-0200013130320221"></a>

<a id="canonical-3201100323112131-1102110311221330-1132300303032012-0302331233211223-1322111011310122-3233113021011132-2233321020232313-0020031121102133"></a>

## dns_list property — configured_list / 233021121120 / 4

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

<a id="canonical-1232123033030130-2310222213203212-2210113100033321-1200331130202021-2332213223123310-2012333012003312-1122221132022111-1101111103110123"></a>

## Next pages — configured_list / 233021121120 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123232002111032-0223003103032330-1001200030003322-0201111001000101-3302000022110301-0331003213320330-2110312122133003-1301121312110322"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 223300230202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-0221011032200230-0300323020220320-0023231213302130-0302023213222312-1230020231330020-3231213300200022-2122320300330313-1233010201330002"></a>

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

<a id="canonical-2022012310220330-2202032232033003-3113003001321113-2321023013233010-3022233311031100-2202103123030220-2100212201100300-2013021322111212"></a>

## Direct properties — local_dns / 223300230202 / 3

<a id="canonical-3121302232032331-1030310330221100-3103110110001113-0120011010011221-3031200030001100-1103001010300230-0103211132321032-1133000122033121"></a>

<a id="canonical-2330313233122023-3103032021300333-3002100113303010-3301332111211220-0323213130102310-1303023233120123-3013332103220123-2330203010130011"></a>

## configured_address property — local_dns / 223300230202 / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-006.md#canonical-1323102212103022-2033122100222320-0213102322330312-0030231000210023-3133101203133202-2022030330330122-1113333330323201-1102213201320030): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0322200010320022-1211002002100133-2210003223112332-1302100333333231-0132020222132322-0301203122212221-2102221011000023-0301121122233000): complete subsection reference.

<a id="canonical-1121203213001223-2330212101010202-3222332301231213-3120012010103123-3023000103321012-2102012213031131-2102030220013023-0131302033233203"></a>

## Next pages — local_dns / 223300230202 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-006.md#canonical-1323102212103022-2033122100222320-0213102322330312-0030231000210023-3133101203133202-2022030330330122-1113333330323201-1102213201320030)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-006.md#canonical-0322200010320022-1211002002100133-2210003223112332-1302100333333231-0132020222132322-0301203122212221-2102221011000023-0301121122233000)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1323102212103022-2033122100222320-0213102322330312-0030231000210023-3133101203133202-2022030330330122-1113333330323201-1102213201320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023102201021111-3122130131322102-0121102321303100-3310102322130022-0300233030323321-0211132321022330-3222231112123112-0132203300012311"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 103320020123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1022312323001302-1013123122033232-3121100112301322-0320122211321310-3223202023111103-2031200101321223-2223232010131021-2123301331133311"></a>

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

<a id="canonical-0301120121012230-1000020333233012-0301123323302210-0220323332102201-1211321223322022-3012010022210310-2330222332020002-2230023031331310"></a>

## Direct properties — first_address / 103320020123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202010113220033-3312223323122113-3321120033133001-2003022201003133-3301212200030311-0323333303213200-1012332023312022-1232011113212010"></a>

## Next pages — first_address / 103320020123 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0322200010320022-1211002002100133-2210003223112332-1302100333333231-0132020222132322-0301203122212221-2102221011000023-0301121122233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133213010333101-3321020011031323-1231302211232303-1002221132121330-0203132102121311-1011111231132221-0001331021210313-0102102300101203"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 212202020002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3022020113031011-2020011010202232-1332320200003221-1110011131231233-2201213302300212-1202231201001223-0121300213103330-3032002103333033)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1113312210132130-0130222120002121-0031213212333230-2323213010300111-0221332211102122-1300323222302133-2123110023011323-3100113103210302"></a>

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

<a id="canonical-0023021313230210-3133302213313001-1133330213211113-2232331123200220-0300301223020110-0323023300001021-0313103032110103-2310001232101330"></a>

## Direct properties — last_address / 212202020002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023012333300103-1201311232233120-0322000331300031-2102232023330321-3322111111022233-3011323122323323-1132030101330031-0101013033223032"></a>

## Next pages — last_address / 212202020002 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-2133300300133032-1002223232120012-3133021010321031-3302013111123111-2020133200211033-1310210011102002-1221003110023012-1133011223102331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000330321212221-3302002111220121-2300222301231213-1333122001311123-2000211222012200-0213301003222022-0221031212000131-0003112310222231"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 030331321002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0102230330012013-2300233200223211-1200230320002133-2122101210032321-3312131201031022-0233121203233232-1123333001220002-1301012220233013"></a>

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

<a id="canonical-2332232132312311-3123332311231132-1210030031011231-2221112131013212-1030331112012031-0102121330032221-2001002212013322-3302330103032233"></a>

## Direct properties — stateful / 030331321002 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-006.md#canonical-0032310333323021-3332322201303023-0231120322201123-2322232231001311-0131333223130013-0001021233203001-0011101120203102-3033103301121221): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-006.md#canonical-2122233013102120-1022222301001100-0313110122230320-2003300102310303-3110303222303111-2110032010301303-3313330030210211-0011033113001120): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-1200020222223321-2121201133022222-2300020021300012-1021032202101010-0133313121122112-3012312011113220-2022313331320201-2223332011221211): complete subsection reference.

<a id="canonical-1330021200021312-1212112321130110-1001133232030302-3213120011302201-3211320232213031-0123201231100031-0113102220133223-1000111101013210"></a>

<a id="canonical-1110221222010330-0222203120100103-3301313020300010-2231303002133220-0031132011132200-0312201201332310-2211130130000011-2333330011032031"></a>

## fixed_ip_map property — stateful / 030331321002 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-2300212221022321-0300213330033033-2011032013323222-3303212003332320-2030331210320222-3102330000203311-0003020333320012-3313030300210331): complete subsection reference.

<a id="canonical-2222211031200011-0030121212030111-2102102012131001-0211120032101102-2320320331022000-2233330220022001-3333330302122131-2031112233231212"></a>

## Next pages — stateful / 030331321002 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-006.md#canonical-0032310333323021-3332322201303023-0231120322201123-2322232231001311-0131333223130013-0001021233203001-0011101120203102-3033103301121221)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-006.md#canonical-2122233013102120-1022222301001100-0313110122230320-2003300102310303-3110303222303111-2110032010301303-3313330030210211-0011033113001120)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-1200020222223321-2121201133022222-2300020021300012-1021032202101010-0133313121122112-3012312011113220-2022313331320201-2223332011221211)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-006.md#canonical-2300212221022321-0300213330033033-2011032013323222-3303212003332320-2030331210320222-3102330000203311-0003020333320012-3313030300210331)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0032310333323021-3332322201303023-0231120322201123-2322232231001311-0131333223130013-0001021233203001-0011101120203102-3033103301121221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023211100330000-3030113322102300-1123233003121003-2222023111112311-1202200202213312-0310201103220020-0323222102201023-0230030232322101"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 223033330033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0010032203213100-0201222322123303-3310022102031100-2002331001131033-0200301303310222-3030031031322032-3333212302021132-1322311212222322"></a>

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

<a id="canonical-2303100303223200-0322203312303021-1231213000211323-1303101303313212-0211001120300012-1033100312023232-1120330320132210-0131300021223232"></a>

## Direct properties — automatic_from_end / 223033330033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111030331213112-2113123010232231-0031013302121232-2310123110133332-1231333023333201-3212110313023313-2120100321210113-2311321011200120"></a>

## Next pages — automatic_from_end / 223033330033 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122233013102120-1022222301001100-0313110122230320-2003300102310303-3110303222303111-2110032010301303-3313330030210211-0011033113001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031113213221310-3113231113101323-1210122333110021-2111231110201331-1022101023323202-3013233030220030-2013011202210230-1131211332110103"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 131312101203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1033300013013220-0231003321223111-1123120302133132-3123123021201303-2332021302231230-0031032333303221-1200232221011302-2333103132203333"></a>

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

<a id="canonical-0130121220312211-2233112223002130-0233111331011320-1321310210110212-3130310010100313-0023120313210011-1233013331202010-2113203212200010"></a>

## Direct properties — automatic_from_start / 131312101203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331011003011102-0031101313123030-0223102212132123-3323033120031130-0103131202111320-0231110031031021-3020300200131133-3101220011101020"></a>

## Next pages — automatic_from_start / 131312101203 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1200020222223321-2121201133022222-2300020021300012-1021032202101010-0133313121122112-3012312011113220-2022313331320201-2223332011221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102231111120210-3310122020200322-3302123203321202-2032302011130220-3021213301000310-3122010223032120-0113002112113321-0023330322221320"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 122322320023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2102131100313112-1012111231102321-0122102010030223-0312122230201011-0200201302020132-1203330333313031-2210331133101030-2310113023200320"></a>

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

<a id="canonical-2300131030121220-2223022123310322-1031330331323200-1310031222120221-3023302102110321-2132310310113311-2332021203311113-0202222111103210"></a>

## Direct properties — dhcp_networks / 122322320023 / 3

<a id="canonical-0032211300230130-1213301310201210-3112201133330001-0121022132310002-0001213313133201-2122130110233202-2333132021232000-2033001312111032"></a>

<a id="canonical-3332210321220231-0310302222113020-2131313313110323-3200212233203200-1212102011323002-0223201313003003-0323323201022022-2222330121130112"></a>

## network_prefix property — dhcp_networks / 122322320023 / 4

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

<a id="canonical-1103033033313212-1200300203231122-2030102103300013-2312002311212333-1102310310123233-3021200210323333-3222130011312331-1310023233330121"></a>

<a id="canonical-1331031302132330-1102223002001120-1322222032112201-2022232021132332-2303310203213330-0231032121010111-0213001112033011-3323311303112103"></a>

## pool_settings property — dhcp_networks / 122322320023 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-006.md#canonical-2123120320123110-2202322022023202-1122230321203001-2213023112120200-1012200130221230-3321211213010223-1002133130303030-3021010300112003): complete subsection reference.

<a id="canonical-2233301123113300-2111313310031230-1133112211000033-1132321223311322-3112301210300123-0211123313302101-3212001301122103-0111313000000212"></a>

## Next pages — dhcp_networks / 122322320023 / 6

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-006.md#canonical-2123120320123110-2202322022023202-1122230321203001-2213023112120200-1012200130221230-3321211213010223-1002133130303030-3021010300112003)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2123120320123110-2202322022023202-1122230321203001-2213023112120200-1012200130221230-3321211213010223-1002133130303030-3021010300112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203011312203112-0102121321110113-0033100030011311-1103312102333113-1110133032202030-2101301033221121-1332103010101210-1113333010203210"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 213301330231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-1200020222223321-2121201133022222-2300020021300012-1021032202101010-0133313121122112-3012312011113220-2022313331320201-2223332011221211)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-3210300110023000-2021303301100233-3032201103103223-0232012111300100-0331031110030210-1320230000013203-1111110133120300-0333102320232301"></a>

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

<a id="canonical-0201320131213013-0103112333333311-0231202102211311-2203231300310012-0321233323023220-1010102310202311-1030122021220200-3020011303201031"></a>

## Direct properties — pools / 213301330231 / 3

<a id="canonical-0311022102000312-0012210012312120-0002222311231021-1233332113322221-1321310133311221-2331023120103123-0230203231131221-3020231232231233"></a>

<a id="canonical-2333311213320302-0023012310310111-2320112110011302-1101230003001120-1201322031012110-3231133303330031-1332232221223113-0300113111220132"></a>

## end_ip property — pools / 213301330231 / 4

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

<a id="canonical-0313110310123103-2012201311032130-1302230013202101-3131103001230322-1112222003313033-1103001030301112-0121003022002303-3213200021233103"></a>

<a id="canonical-2202330331012120-0023332103302020-3103033322211330-2023313032101100-2023311232320022-2023030003222101-2231121313001230-2220321303101300"></a>

## start_ip property — pools / 213301330231 / 5

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

<a id="canonical-2130321213112130-2023332200302333-1103131122033103-2133022121233131-1103201113331110-2301131112030301-2310000100033111-0332032011300313"></a>

## Next pages — pools / 213301330231 / 6

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-006.md#canonical-1200020222223321-2121201133022222-2300020021300012-1021032202101010-0133313121122112-3012312011113220-2022313331320201-2223332011221211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2300212221022321-0300213330033033-2011032013323222-3303212003332320-2030331210320222-3102330000203311-0003020333320012-3313030300210331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030032220231101-3230131330230230-3330210301110221-3101233331101013-3231001021200032-1010310132300101-2203200011212103-0312212223112122"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 122322121301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-006.md#canonical-3031332110213110-0103112101000313-2211131231230211-2201030003013123-2102201211320103-3120032030320312-0011110123320202-0200210103203301)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-006.md#canonical-3130023300033323-3301131230223103-3132122101131302-0323220000013322-3302120231012022-3321030110032320-3302332110121300-3322220112230031)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1032003331013013-0230121232223112-0303222112022333-0323133121123232-0022021122103331-3213132322031100-3332222012310201-2312031330232120"></a>

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

<a id="canonical-3103033003001112-1233010130220011-2332030221220320-2133331201221301-3020002133033233-0311331322231313-3130001113300321-2012221103320032"></a>

## Direct properties — interface_ip_map / 122322121301 / 3

<a id="canonical-3001310032232003-0320302300213032-2121212300210012-2101113103301122-3122201011023323-2320003013232132-3023312122100312-3332100233320011"></a>

<a id="canonical-2220102323220113-3203122101021011-1031300312021131-3120330302311212-2223102103023003-1310011112223002-0202301102210020-3012131122111003"></a>

## interface_ip_map property — interface_ip_map / 122322121301 / 4

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

<a id="canonical-2123131120112210-3000222130121111-0100312133121100-3031102003103233-2123132211232221-1103213002100122-1220323113201312-0301203323100321"></a>

## Next pages — interface_ip_map / 122322121301 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-006.md#canonical-0222001133030221-2212031303321201-1223102331131211-1302230233310303-1022320010001113-0201331023132303-1220003220301020-3312101112203013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1001301210321323-3322002230003132-1101231323313123-3331010111310110-1332021203103302-1101102211202132-2012212310002132-1101213030301023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202320312111210-2201001311323111-3020330002031330-3003100333310332-1313113130103021-2020011001212200-1102010221121202-3333211133303201"></a>

## baremetal.not_managed.node_list.interface_list.monitor — monitor / 002300333131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.monitor

<a id="canonical-1110223220220103-2020013132000221-0100320122232203-0130100130122113-1322021021323122-0000110003030021-3303022212222212-3211010213131022"></a>

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

<a id="canonical-0232113230211313-1203310213113232-3231033020230201-1121322033230210-3130022332000200-0221113033132122-3230112233033211-0332331103220320"></a>

## Direct properties — monitor / 002300333131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201103203233321-2103101231203222-0231332013310100-3100131333222232-3203311302011322-1021223030030031-2231323222012113-0113322133221011"></a>

## Next pages — monitor / 002300333131 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1211030121310013-1130120210121103-2120032103102101-3123220033113121-2110023302003100-1302003333230221-2232213313123321-3233121302310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130212233012031-0133202100320112-3101122312332201-1111310311012113-2320123100023033-1111122002311310-3311011133332101-0220303122200202"></a>

## baremetal.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 311321122303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3013112210303213-0221123211122332-3000103030022313-1102220220323020-2001111212333312-1101232012220332-1001211231023021-2301232023012132"></a>

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

<a id="canonical-2131202003100312-3221202301201233-0212203331123221-0010232100212313-2020111310112303-2000113232320010-1122203330110213-3110211331023233"></a>

## Direct properties — monitor_disabled / 311321122303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102000100330133-2130000101030013-3320331033333011-0022301003313330-2120230302133320-3031001100121020-2232321012022013-2320301230300012"></a>

## Next pages — monitor_disabled / 311321122303 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2303100010301113-0233113232223213-3302303332333131-2030023313201112-0203211231121320-0000002032310303-2220233012033022-1203010112300032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011212233313311-2211032301212000-3312311103101231-3322301101121101-0121011031320202-3331013010130310-3123012303122011-1331131133130103"></a>

## baremetal.not_managed.node_list.interface_list.network_option — network_option / 112013302303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.network_option

<a id="canonical-1300303330123002-0311023121030103-2133123303130131-3023203320230011-1002313302303123-3122330220012131-1030030032232131-2310123333033012"></a>

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

<a id="canonical-1032321132003020-1310213132132032-2012231330032320-2233121231310323-3203101322223030-2120012321233001-0202031011322120-3212031003201331"></a>

## Direct properties — network_option / 112013302303 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-006.md#canonical-3210311333000123-1222323022202011-1011312033233323-1133123002123012-1130001212332110-0231203101021003-1102213030221331-3121030012200130): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-006.md#canonical-2010323102130102-1312033101122300-0100113011000103-0132023031333022-0002122303211312-1003000322203333-2013030301212120-0232200113030221): complete subsection reference.

<a id="canonical-3301211330301333-2202323012122012-2022333230210032-0301032100303312-0201220220132222-3331113210231023-1200200131011111-3000030022100022"></a>

## Next pages — network_option / 112013302303 / 4

- [baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-006.md#canonical-3210311333000123-1222323022202011-1011312033233323-1133123002123012-1130001212332110-0231203101021003-1102213030221331-3121030012200130)
- [baremetal.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-006.md#canonical-2010323102130102-1312033101122300-0100113011000103-0132023031333022-0002122303211312-1003000322203333-2013030301212120-0232200113030221)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3210311333000123-1222323022202011-1011312033233323-1133123002123012-1130001212332110-0231203101021003-1102213030221331-3121030012200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032112011212333-3032123310110210-3313230303010133-3200301111301120-0301333202022312-2332033002102332-0000122133022321-3011022211002031"></a>

## baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 201230213211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-2303100010301113-0233113232223213-3302303332333131-2030023313201112-0203211231121320-0000002032310303-2220233012033022-1203010112300032)
- baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1220202021211231-0221103032303311-3210213112010320-1103200303303311-2022213212013332-2231001301303333-1331012121323320-1023231300010000"></a>

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

<a id="canonical-2201101203220202-0210022332202211-1232303013332033-3331032112323230-2332003123102110-1212001321102120-3033323313203220-3121111310202123"></a>

## Direct properties — site_local_inside_network / 201230213211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013000330212331-1033122202030032-1020232120301033-1000100320003013-1322013301020021-3001330032001210-1112113233213223-1101201321310312"></a>

## Next pages — site_local_inside_network / 201230213211 / 4

- [baremetal.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-2303100010301113-0233113232223213-3302303332333131-2030023313201112-0203211231121320-0000002032310303-2220233012033022-1203010112300032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2010323102130102-1312033101122300-0100113011000103-0132023031333022-0002122303211312-1003000322203333-2013030301212120-0232200113030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302213200303201-2212012012103221-0132132213212321-0120001100302200-2101132133310020-1113331013131231-2013103111220023-0002003021011301"></a>

## baremetal.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 302233221031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-2303100010301113-0233113232223213-3302303332333131-2030023313201112-0203211231121320-0000002032310303-2220233012033022-1203010112300032)
- baremetal.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2211203332201210-0330102211101213-2230321323231213-3320120122213022-3033300021021103-0023321033003102-3003200221322010-3301120213031300"></a>

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

<a id="canonical-2223101222332211-0112301211103230-0213001132322012-3213332002303033-2301111012233123-0120002330132231-0331300131231311-0313121122010201"></a>

## Direct properties — site_local_network / 302233221031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002310101313131-1110030103211110-1213312201321031-3000313323001102-0202221031033223-0222023230123032-2100310331211000-3110102321012233"></a>

## Next pages — site_local_network / 302233221031 / 4

- [baremetal.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-006.md#canonical-2303100010301113-0233113232223213-3302303332333131-2030023313201112-0203211231121320-0000002032310303-2220233012033022-1203010112300032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2201331300131201-0012302323220212-0211302032300311-1212103101112300-3311013013222000-3301002111121013-2123013213200033-1212131023302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110133013111200-0122301121030020-1332213333220221-3133221100223211-2020133000111111-3033003100223220-2032220211311230-0300101301200230"></a>

## baremetal.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 132002321313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-1321022031021311-3122110200133101-2330012000211113-2020102231302020-2320103221310233-2311131312121132-3211111102012112-2201322300011122"></a>

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

<a id="canonical-3301301111303213-3122112133222333-3111032221203123-0020030302102213-0002313113121202-3233310010030332-2122232332031220-2211333133021313"></a>

## Direct properties — no_ipv4_address / 132002321313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232121123020000-2331131232303231-0120312121013203-3120223300323131-3031030210203201-1130322222322203-3312112220030120-1322311322102201"></a>

## Next pages — no_ipv4_address / 132002321313 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0013113320310130-0312013000020112-0001321310333233-1203222320200000-2021003311121330-1012312120223201-1230231001333303-2212321101001233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300000330112113-1221330102011131-1320021330332213-1100200103033110-2202132213330230-1332211323011122-3200122210131133-0223330310113212"></a>

## baremetal.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 223313211313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1122303120133021-2332320221120113-0313103122321032-2003223023203322-3202103013223110-0122212300111203-0102223200320110-3020133011011121"></a>

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

<a id="canonical-1223303120330301-3120202223010132-3033023112232332-3002303022110002-3320001330013102-3131020020323113-2112313232200332-0022201101331220"></a>

## Direct properties — no_ipv6_address / 223313211313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322112013123231-3001120023023233-0221033030112031-1231033023030202-3110301301323120-3123021300100302-3232030331103001-3113202201101333"></a>

## Next pages — no_ipv6_address / 223313211313 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2030112120231320-3233101122212103-0003203330130213-0301311111201020-2120110102112231-3233312120020001-0111331303000332-3223230100323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333222211332031-0230220301231012-0012213203110111-1102231122212021-0312321121301101-1230101213132011-2132132112233211-1130201123311213"></a>

## baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 032131101222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-0120110203312120-2312013212013203-2011103223332030-0322013210020203-0302101331210202-0200320010222113-1210233022021130-2113100020112032"></a>

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

<a id="canonical-1123121021011221-1212120220210022-2333103312132201-1120202312320021-1000111313013030-0233022113120302-0132012121101001-1301210213330221"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 032131101222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331130330123013-0222320131303013-1121110120122121-0123302323331222-1112202220313232-0033001020222013-1012331100131310-0200313131021001"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 032131101222 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1232201111031220-3223013313221311-3003301312332332-1321130012213223-3210331030110332-2303003220303313-2032130212000301-0013113013323131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020103222200223-3103130003033120-2013013021111230-3310021201003033-1310131122120021-0032303122033122-0330111220130133-2103020003220312"></a>

## baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 220120333330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0330232201001332-3003112110202011-1321301330020011-1010222302132220-3032332320131021-0211023331312121-3311230321002203-0332232232103023"></a>

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

<a id="canonical-1011110133113301-0211331130233223-2120203133233313-2120103331223101-0332131231232201-0200203331023020-0121133222022210-2233221322323311"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 220120333330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331301112132300-1200003321212110-1222111310301110-1101200301030000-0023320211123313-3123033213120320-2322222220123330-1032133103111303"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 220120333330 / 4

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2132111032233212-2202122112312213-1120101222132313-1003303200032320-1022222200300010-1213011202310302-0230022332310230-2233301110320231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123210303133033-0213203020303032-2320332212332320-2033120130201330-1032213110320113-3321131320022111-1212322103223302-0113110330123011"></a>

## baremetal.not_managed.node_list.interface_list.static_ip — static_ip / 202212113112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.static_ip

<a id="canonical-0321302230123132-1222010200032002-3303100031201231-2223232100131123-2012312232201333-1103130303033223-1233121200321013-2321220113102112"></a>

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

<a id="canonical-0113332322231223-2012333221111120-0203122011011303-1213221122120223-1213023200102223-0120012303121033-3213320211103022-3211222323200213"></a>

## Direct properties — static_ip / 202212113112 / 3

<a id="canonical-1313231313120100-2132000003033303-3020133230112320-1321213131010230-3131302302100231-1122200011101130-2033330203013331-0200010201301131"></a>

<a id="canonical-2131112020332220-1320231102310010-2310122020123133-1300002331020111-2000233020233032-0002200113012223-3221133302323302-0311200231311300"></a>

## default_gw property — static_ip / 202212113112 / 4

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

<a id="canonical-3030203200032110-0010123110233222-1310323022332320-0201331022031201-0330121302001022-3012312223121301-0122020013320201-1121302302002332"></a>

<a id="canonical-3113321000311113-2321323210313332-3032330101113201-0232333322120022-0212020103021322-0002231120300111-3303311201000000-2130123210023211"></a>

## dns_server property — static_ip / 202212113112 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2301122100030012-0000332221113223-1131211100003221-3130010111113133-2103311000231123-1212000232033213-2230012333113003-1330311022310113"></a>

<a id="canonical-0113200012002330-2200311130032220-1033220203030313-3011213321323233-0331302101003332-3131111130022221-2303233013112113-1100212212131122"></a>

## ip_address property — static_ip / 202212113112 / 6

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

<a id="canonical-0230321133020303-1213223230300222-3221030131210010-0121100213112103-2221202320223201-0022010310203201-1313223113002232-2331013231220120"></a>

## Next pages — static_ip / 202212113112 / 7

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2013322300003332-1300323301233323-3200013330313220-1131211002312220-0021020331001333-2303010102321020-2110131312211002-0301010223222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033232331211000-1212020031220333-2001301321223002-1300322101012013-1311112321132031-3003032010312032-0322301032200001-3122003012213200"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 012130233313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3030231110201323-3032032131221220-2033330232310313-2032223002113300-1001101301120320-3011022200212302-3010033013100212-0303320010123021"></a>

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

<a id="canonical-3011220211031330-0103330030220022-3032301123122002-3302310213000011-1333123020122102-3220322000231012-3202312023030332-1223233123232020"></a>

## Direct properties — static_ipv6_address / 012130233313 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-2000110130121013-2211103002132130-0210020211201210-0011330122011032-3122020213200333-3010120300223333-3203003311030232-0133022231322333): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-2313120331023333-1313220030200210-0221321213103031-3122121233312201-3222313110113210-0003002323201001-3223013312230231-3233123132233231): complete subsection reference.

<a id="canonical-3332222120101223-1230213300130220-1022112122313010-1330320013312113-2200332020211213-1221213210312121-2300201220101322-3332120103203221"></a>

## Next pages — static_ipv6_address / 012130233313 / 4

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-2000110130121013-2211103002132130-0210020211201210-0011330122011032-3122020213200333-3010120300223333-3203003311030232-0133022231322333)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-006.md#canonical-2313120331023333-1313220030200210-0221321213103031-3122121233312201-3222313110113210-0003002323201001-3223013312230231-3233123132233231)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2000110130121013-2211103002132130-0210020211201210-0011330122011032-3122020213200333-3010120300223333-3203003311030232-0133022231322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131000223330023-1031211133300102-0310133313110300-0202103133020331-1332101023212322-2120032303012232-3331113213321303-3003131020321001"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 130323332012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-2013322300003332-1300323301233323-3200013330313220-1131211002312220-0021020331001333-2303010102321020-2110131312211002-0301010223222013)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-3022022130330112-2030332003011213-0112010311311200-2200022311102320-3330212210211123-0310210133232201-3200320321321003-2110121332323111"></a>

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

<a id="canonical-2230131223313011-1003010032020223-1012320110030220-3021232220032312-1000310220303321-3322110321122013-1301232030110223-2302112220323001"></a>

## Direct properties — cluster_static_ip / 130323332012 / 3

<a id="canonical-3112223302112122-0220331301130023-0203300113211311-2332120230110132-0202032001330131-1230010332133331-3301112020121331-3113002211230000"></a>

<a id="canonical-3101032030333021-1202003220123033-0310121202321021-0303303030013220-3032111020322230-3030333221213320-1233022022203332-1112200130103102"></a>

## interface_ip_map property — cluster_static_ip / 130323332012 / 4

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

<a id="canonical-0230333310120201-3001130201213121-3100013110203030-2321203131313302-2012032023112312-3213202223302023-2010322332213330-1123333300003232"></a>

## Next pages — cluster_static_ip / 130323332012 / 5

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-2013322300003332-1300323301233323-3200013330313220-1131211002312220-0021020331001333-2303010102321020-2110131312211002-0301010223222013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2313120331023333-1313220030200210-0221321213103031-3122121233312201-3222313110113210-0003002323201001-3223013312230231-3233123132233231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012013022121012-3201221101232113-2000330000313130-3010220332012032-3333101201233023-0111311203303021-1111000131311223-1232301011330102"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 321333232122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-2013322300003332-1300323301233323-3200013330313220-1131211002312220-0021020331001333-2303010102321020-2110131312211002-0301010223222013)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3032301300012321-3133333312132233-0210301033002111-3333330221303022-3113223123213213-0023002211100030-2210003022320003-3312133003203220"></a>

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

<a id="canonical-2223000313021021-1001301223312330-3010302100110021-3113310213010122-1022310021012133-3233110112130112-2321000202230012-1111333001102212"></a>

## Direct properties — node_static_ip / 321333232122 / 3

<a id="canonical-0100020110221321-1021112300122300-3333300323201030-0312310010032020-0300211103000101-2102221100303121-3330302312310131-0130120323221113"></a>

<a id="canonical-3021321231110321-0201030023133123-0131113213232303-3323023230023202-2032202210030301-1322322223321002-0101232011212232-2102312222212131"></a>

## default_gw property — node_static_ip / 321333232122 / 4

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

<a id="canonical-3112123131233030-1201302131110202-0113300001020003-3200002223003030-2323200303121321-3202001110010321-1301223020211230-3013212132211323"></a>

<a id="canonical-3002320120132212-0331212332110321-0021133000330121-3313311212301313-2300120133332312-0320112321210022-0031311310021133-0012222220002310"></a>

## dns_server property — node_static_ip / 321333232122 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0102210303212333-3100111323120001-0331023110212201-0231123011130303-1122333322333113-2311323323020123-2002020102121201-3202100313123221"></a>

<a id="canonical-0332002310331310-2120102300301131-3011102010001022-0220022012122002-3330133320132021-0211102030312303-0203230312223200-2330013311220202"></a>

## ip_address property — node_static_ip / 321333232122 / 6

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

<a id="canonical-1120310120331310-0120322230103221-1020223031213200-3322331331021132-0120320203230233-1001322220233202-1311033123323220-1003300300320132"></a>

## Next pages — node_static_ip / 321333232122 / 7

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-006.md#canonical-2013322300003332-1300323301233323-3200013330313220-1131211002312220-0021020331001333-2303010102321020-2110131312211002-0301010223222013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132231201130320-1200202033231312-0303021022313100-2303232000202303-2100031023110203-1201120320200211-3300312131130022-0312233032322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300112110333332-3313013002000312-2103323310301212-3003010323312020-2312200231302321-1022330021220222-2333320333322300-3103233032330212"></a>

## baremetal.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 031023000023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303)
- [baremetal.not_managed](resources--securemesh_site_v2--reference--group-005.md#canonical-3330133313120320-1110031122210111-3100332013003112-2121102133112122-1000110010212301-0333220323131210-0113311122000302-1123230003311032)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--reference--group-005.md#canonical-3020310101222311-3213000120333023-0300113133233301-3003003302313011-1102011203021302-2232110113332333-2320021101333123-0120322231031011)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- baremetal.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3132220132331330-2220212312333120-0212331032103310-0312102022200011-3002100311322211-1102033312111333-0121201113021312-1202113102132132"></a>

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

<a id="canonical-3232131203233103-0212030023001020-1022212233102103-3110300133313032-3210202310302302-1101313313033232-0003332133331011-2131201232220221"></a>

## Direct properties — vlan_interface / 031023000023 / 3

<a id="canonical-2323010020121221-3300331133132001-2011113312330211-1223231322113031-2331132300332030-3112203212121030-2211303013201033-3312120122011130"></a>

<a id="canonical-2130203331122202-3020211120231130-2021020023333123-3022003000231122-3202231001023020-2331103230132233-2212332123011221-2220332320030233"></a>

## device property — vlan_interface / 031023000023 / 4

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

<a id="canonical-1001132010220000-1010300013023221-1112110113323133-3220301220320212-0303211132100111-2100311131331033-0010113003112233-3321131223010013"></a>

<a id="canonical-0030110323010110-3011011002312010-1103221020132003-0321313212110032-2323013031230311-1312230302321121-0321201022332303-0030230032111233"></a>

## vlan_id property — vlan_interface / 031023000023 / 5

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

<a id="canonical-0310123231001100-3010021033133102-3122122223033131-0011331100012223-3001202230102303-0303331211200023-0030031133212230-3120200320211211"></a>

## Next pages — vlan_interface / 031023000023 / 6

- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-005.md#canonical-1223122220311222-1133301100222133-1122010023011123-1011023011233011-0211201203002123-1221231102332022-2332301211302302-0003100133203001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0213320123112011-2202110211322101-0013303221220031-1312221223201333-2013012110330312-2222232103123133-2323003000320332-2232323130313122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233032302013222-1332222313011031-3232022000321323-1212302031011223-0320031321133022-3212213300310310-1303320310333313-1013212103333010"></a>

## block_all_services — block_all_services / 302120111300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- block_all_services

<a id="canonical-1200230033201000-0130212221222131-2010122123110320-1230032221232220-1221122201130022-3132113001133302-1101001131023303-2110200201212032"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1200230033201000-0130212221222131-2010122123110320-1230032221232220-1221122201130022-3132113001133302-1101001131023303-2110200201212032)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-3213310320320000-2012310120300130-1103100322231110-0230322301220121-2021132132312022-2123023223130232-3232123003320121-1003332012302230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-2022223021120132-0100022113121101-0213232233002303-0202101130331121-3110023202222033-2332322012210203-3330212110303211-2112130113212113"></a>

## Direct properties — block_all_services / 302120111300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020212332331032-2322332103220101-1223122011132120-1100130311011133-2020322112120033-2201020321310030-3322002003202011-2020221321012020"></a>

## Next pages — block_all_services / 302120111300 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032311201102200-1200331132022122-0202003121132001-0133201230112200-3320330103332230-0021302323101220-2313031300321201-1312102312021001"></a>

## blocked_services — blocked_services / 000133233023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- blocked_services

<a id="canonical-3213310320320000-2012310120300130-1103100322231110-0230322301220121-2021132132312022-2123023223130232-3232123003320121-1003332012302230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023033313031032-0003212110110323-3123120331033200-0122003232210003-0201202132111312-0001030120201002-1323301221023320-0333001230002121"></a>

## Direct properties — blocked_services / 000133233023 / 3

- [blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032): complete subsection reference.

<a id="canonical-0230311012202031-1312310102223201-3210033101013123-1001002333201002-2011323221203222-3322120222002231-1031121320102103-0020300311102133"></a>

## Next pages — blocked_services / 000133233023 / 4

- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011111030220132-3333121011233031-2230122212022003-3300000003222100-1121322221113103-3312320103131022-0200301001120023-3022102332200211"></a>

## blocked_services.blocked_service — blocked_service / 103111231002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- blocked_services.blocked_service

<a id="canonical-2211003122121232-0323100031001011-3322303001313312-1213231311310210-2300320012213211-1220200030233001-3030230131211023-2133313200320311"></a>

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

<a id="canonical-0202130003301202-0320023231321301-0301223332230001-0131131302123121-0002232101322320-1032011023232022-2112030310302211-0331210112202121"></a>

## Direct properties — blocked_service / 103111231002 / 3

- [DNS](resources--securemesh_site_v2--reference--group-006.md#canonical-3131030030231103-3100013311313213-2312122111023223-1332210200121031-0330203011101002-2022022223113302-3031131031332103-1112001030211132): complete subsection reference.

<a id="canonical-0010131031320012-1032123113200302-3123213012232200-2001232110000033-3310220331321000-1111100233203123-3022321011310230-1321002310003121"></a>

<a id="canonical-1133201301012113-1310300333001123-1131202323233211-2011110233003123-0330201030321112-1002220312100010-0100010122010331-1003303322112132"></a>

## network_type property — blocked_service / 103111231002 / 4

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

- [SSH](resources--securemesh_site_v2--reference--group-006.md#canonical-2013313321133020-3000322300332212-2023330213032300-1101011303221033-0003302332133321-2133300013200133-1031111333331013-0230020002023301): complete subsection reference.

- [web_user_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-0311113333110132-1332331102333301-1033301322332030-2201030230223030-1032031310230003-2320012320231200-1301031130212110-1031030131310223): complete subsection reference.

<a id="canonical-1002113210301010-0320320133130332-1033101020211101-1332002122231301-0302301100230110-3300002231002021-1023311113333301-0200112110112310"></a>

## Next pages — blocked_service / 103111231002 / 5

- [blocked_services.blocked_service.dns](resources--securemesh_site_v2--reference--group-006.md#canonical-3131030030231103-3100013311313213-2312122111023223-1332210200121031-0330203011101002-2022022223113302-3031131031332103-1112001030211132)
- [blocked_services.blocked_service.ssh](resources--securemesh_site_v2--reference--group-006.md#canonical-2013313321133020-3000322300332212-2023330213032300-1101011303221033-0003302332133321-2133300013200133-1031111333331013-0230020002023301)
- [blocked_services.blocked_service.web_user_interface](resources--securemesh_site_v2--reference--group-006.md#canonical-0311113333110132-1332331102333301-1033301322332030-2201030230223030-1032031310230003-2320012320231200-1301031130212110-1031030131310223)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3131030030231103-3100013311313213-2312122111023223-1332210200121031-0330203011101002-2022022223113302-3031131031332103-1112001030211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303233021101222-3303200321213112-3010203031112112-2202212202130331-0131232210033302-2223301021313011-3202233120323010-3312333023333220"></a>

## blocked_services.blocked_service.DNS — DNS / 201012310330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.DNS

<a id="canonical-1313133131030320-1310101202031021-2323201013102223-1101220122201111-0031010332323001-2311001211030200-3232332222111122-0213033220132212"></a>

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

<a id="canonical-1023011022210331-2303321123103320-0211323130331030-2323001120122011-1131100321321003-2130122001231221-1301332010102100-0330310030021210"></a>

## Direct properties — DNS / 201012310330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310020330221303-2313003202213313-0221112003012302-1000221330132032-3132330111032103-2110130000030301-2120320133133132-1333320200020023"></a>

## Next pages — DNS / 201012310330 / 4

- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2013313321133020-3000322300332212-2023330213032300-1101011303221033-0003302332133321-2133300013200133-1031111333331013-0230020002023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322201023201201-3020031120210003-2320000331001121-3102231201133220-0200332132102113-2133203010012103-2221002023111122-3220211013000213"></a>

## blocked_services.blocked_service.SSH — SSH / 313312132303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.SSH

<a id="canonical-0333302002332220-2132030323021223-1301110310321313-2110121232033303-2002113030200230-2113313030002011-2313231000122210-2121111030201112"></a>

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

<a id="canonical-3033212320112211-2131130021131132-0313200232330312-1133303001323212-2130222212211033-3000133103313331-0310233112333123-0332031000122222"></a>

## Direct properties — SSH / 313312132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002311322320222-2233223033201111-2102111233103231-1132331211302320-2303310223002312-1320112033311003-2111332202121103-3303223132213313"></a>

## Next pages — SSH / 313312132303 / 4

- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0311113333110132-1332331102333301-1033301322332030-2201030230223030-1032031310230003-2320012320231200-1301031130212110-1031030131310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013332301113001-0213322330221220-1013202312021332-0313020231323321-0110032100313322-0332030003223301-3001000322312120-0223231210103100"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 202032222012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-1331200200130100-0002102132110323-2011012011321222-2333332031310103-1000022000122032-2033303011202321-3231220110301310-2202321210013111"></a>

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

<a id="canonical-0302013003310132-1102301213203302-1021011332313301-0232231001223232-0023012211322203-3132121132212201-0311320120320310-0322332312223031"></a>

## Direct properties — web_user_interface / 202032222012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330222322213003-2213330110003010-0113212131122123-3202110322220302-3000123110033302-0212223331013311-2330231103312200-3020223220100033"></a>

## Next pages — web_user_interface / 202032222012 / 4

- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-006.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212133223201312-1032302310112133-2111023110131021-2031310312232213-1203212010032103-3333032001233011-3333220231000032-3022312111303100"></a>

## custom_proxy — custom_proxy / 313033120102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- custom_proxy

<a id="canonical-0133132100102112-1321232211122100-3231121222211020-1323020332320220-3100221323323313-2002012322321002-1030330013212022-1310232101232330"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_proxy, f5\_proxy, private\_adn\] Configuration parameter for custom proxy.

Upstream description:

Custom Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("proxy_ip_address",
    "proxy_port"),
  validators.ConflictingObjectAttributes("disable_re_tunnel",
    "enable_re_tunnel")}
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
  "x-ves-oneof-field-use_for_re_tunnel_choice": "[\"disable_re_tunnel\",\"enable_re_tunnel\"]"
}
```

OneOf alternatives in this subsection:

- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-0133132100102112-1321232211122100-3231121222211020-1323020332320220-3100221323323313-2002012322321002-1030330013212022-1310232101232330)
- [f5_proxy](resources--securemesh_site_v2--reference--group-009.md#canonical-0323202033202102-2203130323103033-2120333100022103-1031333331023101-0001233230123023-0302313331331233-0112221333211032-1203223313013302)
- [private_adn](resources--securemesh_site_v2--reference--group-017.md#canonical-3001100011332303-2231203223000211-1230200300233313-2012022223033331-1002001201232231-1312011231123311-2020311101200103-0010032010320220)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203231201022130-1300022033213003-3331012033330023-2233100203120300-3223011130221201-0003121220013233-1232112230313101-2133322330101320"></a>

## Direct properties — custom_proxy / 313033120102 / 3

- [disable_re_tunnel](resources--securemesh_site_v2--reference--group-007.md#canonical-2211130001212222-0222130033220113-2331102010101102-1233031020203303-0233103111010211-2321212021332000-1232110123223311-2232011002301213): complete subsection reference.

- [enable_re_tunnel](resources--securemesh_site_v2--reference--group-007.md#canonical-3111102123022002-1031033320333131-2212221101213001-2311200122113001-1132310110311032-0212310303232030-2022223013213220-2022110121020030): complete subsection reference.

- [password](resources--securemesh_site_v2--reference--group-007.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010): complete subsection reference.

<a id="canonical-0020010033100320-0210101320112332-1100201312321323-2212132123303213-2101112333232013-1232201123033220-3021211221013003-2033110213333020"></a>

<a id="canonical-2123011022030230-1002032333122321-2332323132233222-2102123310012123-0300132101023111-1020212000310122-3102130303311212-2202013203112131"></a>

## proxy_ip_address property — custom_proxy / 313033120102 / 4

Type: `"string"`. Optional.

Specify the IPv4 Address of the internal Enterprise Proxy.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1003102021201221-2122201000221001-1230301333000122-2100000003000032-1130233323330221-0330202123023113-2230032100020132-0021202303321113"></a>

<a id="canonical-0131103121032220-2032102302211022-0033311030133203-0010311010010023-1333210030111120-1231032112001222-1121120032210011-1033003200133333"></a>

## proxy_port property — custom_proxy / 313033120102 / 5

Type: `"number"`. Optional.

Specify the Port of the internal Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0130213202312113-2000012122332021-1303013322111301-2001202112302232-3012213023300212-3123230323231110-2033012300102123-3003101213332232"></a>

<a id="canonical-2100001113213202-3302233213203311-1122120212303221-3131133100313301-2121331021313023-2223102311003100-0300001001333102-3213013031123123"></a>

## username property — custom_proxy / 313033120102 / 6

Type: `"string"`. Optional.

If the internal Enterprise Proxy is using basic authentication, specify the username. This is an
optional field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
