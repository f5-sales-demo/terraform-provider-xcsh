---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1131002133210310-0200203320202012-1132021320233322-1300231222031011-2230003210100232-3201300120031001-0103233303131111-2301203123323312"></a>

## fixed_ip_map property — stateful / 312323223223 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-0303203000211001-3200133013302030-1102211312123132-3000212230202313-0030300113110230-2310311300332202-1101231220112103-3332220011030312): complete subsection reference.

<a id="canonical-3203110121312311-1301130213112102-2020132130322321-1013013123103102-0130301231121013-0120322320323112-1330221223213301-2330231031303310"></a>

## Next pages — stateful / 312323223223 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-1310223221212020-2203212021200230-0221303023101330-1301310133310311-3032121213033212-1202203201003011-3220012231311123-1010330330003120)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-2211100321120010-2330031313123202-0132303113010022-1320222311120030-0033133130012002-2302131000322013-1300321113021231-1030002012203322)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-0303203000211001-3200133013302030-1102211312123132-3000212230202313-0030300113110230-2310311300332202-1101231220112103-3332220011030312)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1310223221212020-2203212021200230-0221303023101330-1301310133310311-3032121213033212-1202203201003011-3220012231311123-1010330330003120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232330110231113-0011232231003222-1331221221023221-0022212322033221-1320131333231011-1312022203310130-0330132131020212-1313002030223132"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 300132120312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3232120323022011-3010321110231310-2223012033222121-0221332223101201-3100223021232102-0132102212033132-1132320113131003-2002313233230332"></a>

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

<a id="canonical-2313220020133131-2103001131130300-2233311022321220-2210131120003301-0231012213010101-3322212131121102-1312320311300131-1120001233322102"></a>

## Direct properties — automatic_from_end / 300132120312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201131320030222-1211211233331233-2301123020333112-0201221102101300-2213200220032012-3100220122211300-0001132231223231-1131213000103301"></a>

## Next pages — automatic_from_end / 300132120312 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2211100321120010-2330031313123202-0132303113010022-1320222311120030-0033133130012002-2302131000322013-1300321113021231-1030002012203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300032030000120-1020023132013113-2211201102133212-1122303301031230-3132123232022102-1003121120300013-3212122100023303-0131123033322131"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 131122311202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1003031100121222-0320132003302210-1223020131221213-1021233311300131-0131303320231010-2230112103222030-3122122011033121-2123121300223032"></a>

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

<a id="canonical-0010022021030131-3303113200113000-3112002001203112-2012332223333110-1001331120320033-0100311022001222-2110022301213331-2021212220121000"></a>

## Direct properties — automatic_from_start / 131122311202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130322221220013-1102021302221212-3131231100003221-1022212023102110-0113111330313230-2110220101232321-3301131222111301-0301301222221103"></a>

## Next pages — automatic_from_start / 131122311202 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130331122210020-2111211001020103-3210323113313103-0101010122330102-2320023302013122-3032302310311220-0303010003323033-1333010121220301"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 211232010330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1131313212333001-0232330021110222-1223133200102121-3232110011333123-0312001122300330-0312220012310003-2002321013100321-0033210313330212"></a>

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

<a id="canonical-2120132011112213-3331010203320331-0213300331013121-1032002102001121-3200220120233123-1311211020133311-1200011102110231-3312101302221033"></a>

## Direct properties — dhcp_networks / 211232010330 / 3

<a id="canonical-0131210011013310-2320202330033113-0112200230003033-2212122331000021-3113300011000100-1011313301021213-1301003012111001-3113133212012300"></a>

<a id="canonical-0123233030222103-1033010121320113-3033222312302321-0130120100321311-2200030323303201-0133220331132200-1101011301200011-3010231100302013"></a>

## network_prefix property — dhcp_networks / 211232010330 / 4

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

<a id="canonical-2113000113120113-2013313001222222-1113303003112200-2003102011032203-0033031012131320-2302003220131033-1020230222230203-0003313222010110"></a>

<a id="canonical-0131132230220313-0300131132122332-3003013301021101-1101233113113130-0233020012001313-1133001131212202-1300121121001231-3311220013113222"></a>

## pool_settings property — dhcp_networks / 211232010330 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-008.md#canonical-0100312222123233-2223301211301333-3020313101300310-0211332132332333-3221033133201130-2210022031321313-0000033120103030-3330212000233203): complete subsection reference.

<a id="canonical-0302311022121313-1232130203113100-3222332120212110-0312311001203222-0310201202010001-1103330130030101-2232231221320212-2121332211313211"></a>

## Next pages — dhcp_networks / 211232010330 / 6

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-008.md#canonical-0100312222123233-2223301211301333-3020313101300310-0211332132332333-3221033133201130-2210022031321313-0000033120103030-3330212000233203)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0100312222123233-2223301211301333-3020313101300310-0211332132332333-3221033133201130-2210022031321313-0000033120103030-3330212000233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003032210332230-0103310113321133-2123001232212013-2031003113311030-2310202102231201-3011232203101123-3231133301020211-3121121113223033"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 200003330223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1213200232123102-3332121300322130-0313020033301122-0300023201113210-3221033331002212-1222010003303133-1032020111111332-2200103011323210"></a>

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

<a id="canonical-2010132130003130-1000323120312010-3001231222001110-2032020302331111-2001222201233130-1331213012331120-0122222012320101-0210131213332200"></a>

## Direct properties — pools / 200003330223 / 3

<a id="canonical-2011110103302321-0002001222110330-1031101013313201-0121201322133301-3301110231132113-3211233303302023-1230130022032130-1303202012112002"></a>

<a id="canonical-1012121132000133-0213010031103310-2212130120210132-3021100200110210-3000311001011033-0101221033111130-0021020032020003-2321033302020211"></a>

## end_ip property — pools / 200003330223 / 4

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

<a id="canonical-1302023322102110-1303122000220200-2132202120103313-3111022312010330-3001322330300310-2022223221121111-2002130130201010-2120002303130012"></a>

<a id="canonical-1312110010312222-2113202100332032-3100020100322121-0302032010233322-3122113101221313-0322301230021201-3000000323301132-3210230023102130"></a>

## start_ip property — pools / 200003330223 / 5

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

<a id="canonical-2100321130032233-1322222131223313-3333020310233130-0001000112200220-1003110100122100-2331303330211330-1300012013303233-2101221201221320"></a>

## Next pages — pools / 200003330223 / 6

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0303203000211001-3200133013302030-1102211312123132-3000212230202313-0030300113110230-2310311300332202-1101231220112103-3332220011030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130032221031102-2031300130010030-3321100330221221-3121122022330220-0331332100233103-2120102201212101-3020101003012220-3122223232231012"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 021033300000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1323031123212132-1203021332023311-1032103022222223-3232011112001232-2120013130200103-3031231223333021-2132223022300121-1232213221112210"></a>

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

<a id="canonical-0203332123111230-3120322203100131-0122321312101120-2130323133211130-0003332020313002-2030200212131102-1211122132301112-0310330031230111"></a>

## Direct properties — interface_ip_map / 021033300000 / 3

<a id="canonical-0032220013200301-0132231011030101-3111312231122002-1202202032130130-1032202123232031-0332132221330132-1003103220333331-2021323000200232"></a>

<a id="canonical-1203003113100011-0311210312221232-3033120223001300-0100002102030312-0010132211210103-2302011013011330-2122313321201122-1311000031302012"></a>

## interface_ip_map property — interface_ip_map / 021033300000 / 4

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

<a id="canonical-2320131132302221-2123022322323323-3120102201111022-1321322210030131-1300030210220002-0033011133222021-2003113000003311-2211123333113110"></a>

## Next pages — interface_ip_map / 021033300000 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0011102111120011-1022033123221032-1331321200332203-2301111122122233-3303333213003203-2021123332202300-1022321002112333-0023222221123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123331101012121-3212021220032333-3203021021203111-1222332211233222-1202123121322301-1332000011202021-1032023010311202-2033120231001113"></a>

## eks_k8s.not_managed.node_list.interface_list.monitor — monitor / 213100100222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.monitor

<a id="canonical-0110113023203232-3333130321121033-3201310001322312-0101330121121302-1011322002212233-3123031003112320-1233220000110321-0302001212100110"></a>

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

<a id="canonical-3203301210222220-3313231030021010-0330122312213023-2130110030303202-1310222003130233-3303030202302323-3110121201123031-0211012302020132"></a>

## Direct properties — monitor / 213100100222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003012023033020-2233001121220120-1010330013210223-3122033013111021-3201100013220332-0320303030313013-0103103323312001-0310323121011312"></a>

## Next pages — monitor / 213100100222 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3203001200120111-0313132010231330-3001221200201310-3122003203211002-0032332311222332-1021320120202100-3231020003010120-3230012212102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232102220302200-3003010030222322-2332131103023330-1120021221003311-2112131210221103-2332131332110032-1230012103003211-3010332331021211"></a>

## eks_k8s.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 203323032221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0233121130011132-0113223022011302-1310123130220221-1033213123112033-2131232301203213-0111222031002222-0333001113321010-3300332221133300"></a>

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

<a id="canonical-3001232130312113-0312332111212101-3221303212220012-0232123102111313-1101212003022122-0220130322123233-1022323133311332-0132022130313032"></a>

## Direct properties — monitor_disabled / 203323032221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121300213323112-1203220322311021-3003001203023033-3322301333021130-2122003200320012-2001101100120001-3030033102332302-2000211322301023"></a>

## Next pages — monitor_disabled / 203323032221 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032101202211023-0330111320331232-0032032013303012-0023121013331011-1303212303122313-2331220332201032-1332110212110212-2233113220013121"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option — network_option / 011010122112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.network_option

<a id="canonical-2210300003130312-3323012230220012-1122122012230322-3203213021003031-0212121332331101-3322002231332101-1132330000001122-2033232321123210"></a>

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

<a id="canonical-3131320311030312-3222120133230110-0130031112223031-1230113203012132-1120300133010003-3001303022132110-0121203013203001-3200232223020020"></a>

## Direct properties — network_option / 011010122112 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-008.md#canonical-3111010032020112-1021121231311022-2320112131300130-1300210122301010-3221013130300230-2332332012200102-0033023133122322-3122202333312023): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-008.md#canonical-2013022201331103-1232233332023020-3312111233210202-1000112330131023-2232330313331321-1303133322112130-2313100123301032-3112003223111131): complete subsection reference.

<a id="canonical-2102312133001122-1212031311012301-3130023303203110-1331020100130221-3202103132022332-3310003201331010-0103232022023233-1023031013002032"></a>

## Next pages — network_option / 011010122112 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-008.md#canonical-3111010032020112-1021121231311022-2320112131300130-1300210122301010-3221013130300230-2332332012200102-0033023133122322-3122202333312023)
- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-008.md#canonical-2013022201331103-1232233332023020-3312111233210202-1000112330131023-2232330313331321-1303133322112130-2313100123301032-3112003223111131)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3111010032020112-1021121231311022-2320112131300130-1300210122301010-3221013130300230-2332332012200102-0033023133122322-3122202333312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003002121021000-1103021033322020-2200233132310203-3103131000210101-3302131110113331-0321321211331012-0011330202013233-2311020300211130"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 301221003230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2122121202300213-3331303220113323-2113132230122323-3013312233321103-3001300132130030-2121201102113213-2110301320321232-1213120300101332"></a>

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

<a id="canonical-1303123210310110-1330032102211213-0022121223001003-1300212201131121-2313003331203121-1021230210122220-1222321203320113-3320301320111201"></a>

## Direct properties — site_local_inside_network / 301221003230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111030101113312-0030030223032130-1222112303010203-2332220110020211-0120113233012022-1201320221322012-2220202321022332-1023130332330013"></a>

## Next pages — site_local_inside_network / 301221003230 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2013022201331103-1232233332023020-3312111233210202-1000112330131023-2232330313331321-1303133322112130-2313100123301032-3112003223111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033013213313111-1032211211233302-3011031322212221-1201203032013230-2113101023032232-0200110112101212-0131023010313123-2112332320012231"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 100301230020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3020301123011331-1302323001032121-1331232330203202-1221112100323001-1113131312321112-2331231313203201-3131220210101200-0131332330221033"></a>

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

<a id="canonical-3121322313322321-3131223231121210-2021100132012230-2103321002321221-0200202102020031-1233312232120000-0232202021222202-1110033013333210"></a>

## Direct properties — site_local_network / 100301230020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010130211220302-2201232030213203-2100231330022332-1022200002111101-3010333302223031-1320302200202030-1113013010220030-3131022323030002"></a>

## Next pages — site_local_network / 100301230020 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2301031122023323-2312030233133232-0003013033030323-3223231102220222-3121003131102022-1010213001332131-0132020321212203-3121010111120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320301022030030-3301333303221022-1010101231033032-1001330112231330-2231322203302003-2201013023332322-3121121013331132-1333213032221000"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 033331303020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-2230220011102022-0032211113020030-0110002013322200-3011220122012300-3010000130333132-1303231131302213-0203010013031110-0201333011301320"></a>

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

<a id="canonical-2300130100313301-3333210200211221-3133233211201012-0100132032212301-2033202000020131-1201213131232230-1001013030011012-2002113032031003"></a>

## Direct properties — no_ipv4_address / 033331303020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312131210210000-0231020322103231-2022000333122201-2112110113011003-0133302333033311-2000132312213133-3011123211011213-2002003332331000"></a>

## Next pages — no_ipv4_address / 033331303020 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1130210113011010-0201323203331131-2111202121232233-0002232003002002-1113221321222220-3032221120013120-1301312312021220-3211333120012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002000113230123-1132031222130200-2220110322313200-1210300201101120-0303022223111130-3331330200213221-1211033301321332-3123310232301222"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 210302330000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0012222112231101-2121133312332213-3323221233312333-1032331333021000-2102320331002130-0132110203310032-1013132130120020-0323123211030031"></a>

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

<a id="canonical-2010323003221301-0331032000100002-3333010302201123-3020113121223213-3311032213331312-0302322001100112-2331122222103021-0013001020322122"></a>

## Direct properties — no_ipv6_address / 210302330000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022113230111301-0321002212031131-1231132113110112-0033231313233122-1321301331330322-3200001000023102-2201321103331223-1032023122323312"></a>

## Next pages — no_ipv6_address / 210302330000 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0103332210102222-0101313012123113-3110223010330231-1233210031332122-1220201120130302-1030220311033000-1331310033002133-3213231231323310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331113210120300-2103032231213101-3321200020111213-2133202220332031-0030233130213230-2032030302231032-0200200332312310-0230221330001022"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 223312033301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-2231002320103320-1200100301222021-2200230101001031-2132020032213320-2233232020332312-2000231001003211-0013213323330011-2030013130130302"></a>

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

<a id="canonical-3231123231120322-0323302120103112-1223002330010322-3331321013033302-2332103332030011-0031113230121312-1013003113322122-0031020213333222"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 223312033301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112002132330211-0313311221210232-1123121313331002-3111211011012221-2110133001002132-0310113033002322-1031213310131202-3110313331333130"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 223312033301 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2031202033300012-1103300021230131-2131322123021211-3111003121233331-0121331212100212-2222021220102331-2311321321112233-0313323220023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121112212001001-3232223322310110-0110132303112123-0221110020111023-2132020003032233-1222332213030131-3230201132133132-3111331102122202"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 100331223223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1333110222212121-3111023012322321-2320211021201200-2203203021202323-0203333122130213-3300332210301313-2021231013032122-1103331320012323"></a>

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

<a id="canonical-1120213312130323-2302310300233323-2312200123333301-1233032333212310-0002220310321311-1113333221202120-1112203333221323-0222213011203003"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 100331223223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212301133133323-1211223210122322-3123331312300131-2303113111112022-2003313021111023-0012301011330222-0122301100200132-1213101223021311"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 100331223223 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3030211032102301-0013032213201203-2331222321333320-0011121300110003-0303323230202321-1230002033332221-3333131203032011-3301013013201112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221232003000232-2122223332203323-2020210310322231-2311100100333212-3303221232221310-0310002113103103-2130002133200210-0210111123110312"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ip — static_ip / 022133123001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.static_ip

<a id="canonical-0200102303012313-1111323123000102-1000233120232210-3320213110100223-3223213102033202-3023022220300202-1012120332323310-2023130123313301"></a>

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

<a id="canonical-0032311323220020-2000322322231020-1322001112130031-2100221300033302-3303123132123013-3303320012310013-1331132000202203-2033123133221221"></a>

## Direct properties — static_ip / 022133123001 / 3

<a id="canonical-3100120311212113-0002102322233221-3323212030223121-1220032313132132-3211021310120230-1202322012300002-3001020023323013-3311230001032131"></a>

<a id="canonical-1131010231300233-0230002310100022-3321103012022203-1120222220323013-3331133122233333-0022100211302330-1121131331110301-0001030023022213"></a>

## default_gw property — static_ip / 022133123001 / 4

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

<a id="canonical-3322323223301220-3322332312211013-3311230033212002-1310312031132201-0100011330203111-3021332003001012-3331210022333010-1121323202322010"></a>

<a id="canonical-0320022101013110-2213311022033131-1213120310110020-1120322210102331-3312310123003301-2023230133112230-3131121301213231-1121321122022320"></a>

## dns_server property — static_ip / 022133123001 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3320123323232312-2300232331102303-2033132212211010-3220131331301322-2100311010002323-0331123120200210-1202210100203203-3212120021300330"></a>

<a id="canonical-1203331132201231-0020031322011013-2011033102010302-2220220000002003-3020202123131110-0030311123220103-0020000202130123-3200123202220323"></a>

## ip_address property — static_ip / 022133123001 / 6

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

<a id="canonical-2032220313332110-2100123333300120-3032230132233120-0110303130031301-0110132023331110-1310221213200133-0333131321220113-2223111312003021"></a>

## Next pages — static_ip / 022133123001 / 7

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003212313303320-0112322333323320-1022303111002033-0312223112031301-3030033220323300-1220101313130103-2201323100202211-3202100203323012"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 102223233133 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0222131223332001-2100310222011330-1313002320332022-2012000132313021-1023033321303332-3301300213323302-3302231301132302-0023303133120321"></a>

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

<a id="canonical-2000011303121223-2111131110210003-1333012312231001-1303320112112010-1113020001302123-1110011220013213-1201131212303111-3012000020201033"></a>

## Direct properties — static_ipv6_address / 102223233133 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-0322300120000311-1000000000003213-1112321011321233-0112203200300030-0301303311210033-2112011313212110-1003112231103123-0233203000020323): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-1232320223313313-0023112012020333-1123231211120300-1221213021231133-1200233031111120-2302020112301223-1003300133133011-3020130121200311): complete subsection reference.

<a id="canonical-1120303321022122-2203210131303213-3222032200302201-2023110213233133-3031222122121131-2212010132032212-0230202302203201-0211031311310033"></a>

## Next pages — static_ipv6_address / 102223233133 / 4

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-0322300120000311-1000000000003213-1112321011321233-0112203200300030-0301303311210033-2112011313212110-1003112231103123-0233203000020323)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-1232320223313313-0023112012020333-1123231211120300-1221213021231133-1200233031111120-2302020112301223-1003300133133011-3020130121200311)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0322300120000311-1000000000003213-1112321011321233-0112203200300030-0301303311210033-2112011313212110-1003112231103123-0233203000020323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203121020131103-1112001322332121-2113102321210210-3100133222131011-0302321032011321-3201122002211322-2131103331221002-0322033003110313"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 130332102231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2133211030102032-0120322232220321-3100022012131102-2321132320101203-3003133001033013-3220330121302101-2011332111130211-3202231232133000"></a>

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

<a id="canonical-1200232131222302-3100310223132303-0101031131120322-0331021121313332-3310222311003322-2200132001311222-0012202202220310-2232232122022110"></a>

## Direct properties — cluster_static_ip / 130332102231 / 3

<a id="canonical-3123023013121121-2002000333211310-1231133002022113-2231323223221200-1303323310111021-2313131121330133-1221123310211202-3322333131023330"></a>

<a id="canonical-3021231302013230-1002331323231213-3030332121331122-2001000111203002-3312212330022102-0333120312212130-1233322232012331-0232313123003012"></a>

## interface_ip_map property — cluster_static_ip / 130332102231 / 4

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

<a id="canonical-1132203202010022-0012023113222132-2110233023223101-0310300303103323-3322003321311300-0303011323332230-1221213132202110-0133030110310121"></a>

## Next pages — cluster_static_ip / 130332102231 / 5

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1232320223313313-0023112012020333-1123231211120300-1221213021231133-1200233031111120-2302020112301223-1003300133133011-3020130121200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101210222102120-1002332331030100-0311302001230110-0000122321102012-3202233232120330-3032010121101013-3232213202210213-0020301113113231"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 332022201132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2021223322200000-1310231130201012-0010011331332002-2232031320001230-1230000302022122-0001330132320131-2003000021312200-3113031000022132"></a>

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

<a id="canonical-3223333323111310-1101302021222310-2221100220313003-1203120322100232-1301120133201311-3310001311303103-0303202132022203-0010201231202331"></a>

## Direct properties — node_static_ip / 332022201132 / 3

<a id="canonical-1002123202000203-0313113033133312-0332233210111200-1002030203101121-3121320021023032-1232313323230203-0122131321001101-0230303323222201"></a>

<a id="canonical-0131303300133220-3312131003302310-3033213322022203-1001123313023333-3200302200112212-2012212232021320-2011013120102021-0033231320102010"></a>

## default_gw property — node_static_ip / 332022201132 / 4

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

<a id="canonical-3133311102001120-3321302311221120-1300213101301123-0202203221331012-0321223121022131-1331130210010200-3202023232102031-0311133000113131"></a>

<a id="canonical-2302220201210133-0220330032312203-3320132111132231-1233133213211113-1020123031233203-3031203301322211-1222200011002313-2102012233231302"></a>

## dns_server property — node_static_ip / 332022201132 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2022122220021022-3232113032121133-0233001130321031-1222233021110031-3221010030333223-0200211112022111-0102012301200130-2213120010113031"></a>

<a id="canonical-0003322101003111-0221023130300211-2033100033032120-2313121212303000-2320230122130320-0221232213003331-1130220321303011-0013223010112032"></a>

## ip_address property — node_static_ip / 332022201132 / 6

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

<a id="canonical-1301321000033020-1313120211001002-3232010221011210-3312113122322130-2123103021013120-0331233002222323-0201003112323122-0213302231322213"></a>

## Next pages — node_static_ip / 332022201132 / 7

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0301032123232223-1100031103321323-0133220303233001-0223330231112002-3230302111330102-2033000033312333-2302233132112031-3122210031202101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123333311102021-1333111302221011-3202233203101223-2202122100211133-3322323310131301-0031003100222312-0013233231111321-1003332110212221"></a>

## eks_k8s.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 103000020323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0202322111313001-1002230001133123-2121333331200003-2321020312112012-1131030020220000-0033020230320312-2210032100230231-2321131122030031"></a>

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

<a id="canonical-3213023033331231-1313021210131030-0111000311100303-2130113103200123-3023121310331332-0103322223120003-3021102323012302-3032100230302331"></a>

## Direct properties — vlan_interface / 103000020323 / 3

<a id="canonical-1232103033112132-2002301102301330-3232133202212133-0201330110021113-3222202002233103-1132311323112213-0003132312123231-0033322213110101"></a>

<a id="canonical-0112221100011311-2201002321203130-3313200112013323-3222012022212232-0312210211032330-3233110102030033-1210312202103111-1211022010210120"></a>

## device property — vlan_interface / 103000020323 / 4

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

<a id="canonical-3111213322120131-2332030321132212-3013103010332022-0202312111120030-0211322103202110-2111012000310111-1313033032112322-1320132120332002"></a>

<a id="canonical-1210033300122121-1222221012130301-0221201301131230-2110013120211213-0013300011113210-1232333103203032-3201033100032001-2323022321000221"></a>

## vlan_id property — vlan_interface / 103000020323 / 5

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

<a id="canonical-3023032001323332-0232320013320033-1003200231131000-2212233120110231-0313103032023312-2101101000032303-3201212003011300-0233220302002012"></a>

## Next pages — vlan_interface / 103000020323 / 6

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120312013323111-2001313211213303-3013322321101221-2302001312022101-1002300000201112-0101023103222121-0113230102321331-1332233101200023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303032233111213-0301233123331310-1302323123212210-2121023101123000-1230313003200031-1220022022102230-2321301123113311-2210002130232021"></a>

## enable_advanced_delivery — enable_advanced_delivery / 312232203201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- enable_advanced_delivery

<a id="canonical-1220131302000301-1122033110230313-0322113111321333-3220130110202322-2023020022301130-2201202233132330-3211202131020111-2200131101320320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable advanced delivery.

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
enable_advanced_delivery = {}
```

<a id="canonical-2023320001122120-2031122013203200-3121033020111030-2212112213203320-0132002222303213-0301033221223032-3111332221020002-0331230330231100"></a>

## Direct properties — enable_advanced_delivery / 312232203201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023331232310201-2012002220313303-2313221232131003-1203001231110321-3301103100301200-1020233232312021-1230302010323321-2112121023210333"></a>

## Next pages — enable_advanced_delivery / 312232203201 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1220133303231202-3120323110021310-0221103231002213-1302123301122110-0113201221110233-2033320311323331-0000121000323311-3123322022023132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233022100231331-2011113130313113-2222203121321233-0022213203111203-3000102031032020-2100313333030311-2131232321100213-0211013310321202"></a>

## enable_ha — enable_ha / 103230330111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- enable_ha

<a id="canonical-2330001010130310-3032020220122111-2022223300332103-0310000223001331-1202013231320033-2100001200103323-1020221333133201-2231302313002023"></a>

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
enable_ha = {}
```

<a id="canonical-2322332133221101-3311203030012230-0301302130131131-0013030001220313-3322313220133201-0102213333203220-0322202003003313-1022010333201130"></a>

## Direct properties — enable_ha / 103230330111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022233320003121-1321300103303200-2212123011310123-3113031120320333-1302113221101321-1312220030102001-0323223330303331-1021113231111231"></a>

## Next pages — enable_ha / 103230330111 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3313333110020022-1323331321323133-1131231332001000-0330200311301223-2112002130210103-2212202210323103-0001322331303020-1112333301131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303131311310320-3222000201101103-0023202001131123-1333331233132330-0021302121231311-0203223310132230-3112010110310111-1212100231330221"></a>

## enable_log_anonymization — enable_log_anonymization / 313020320223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- enable_log_anonymization

<a id="canonical-1132023020201001-3210133312102302-2003130003203131-1011133331132330-0002300031323000-3311011021302323-3122332102101131-1011013211131323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable log anonymization.

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
enable_log_anonymization = {}
```

<a id="canonical-0321231021112211-3122100003011211-2332223021032012-3012213130010302-1120332022202212-3123321032302330-1012232331322301-1232113323133123"></a>

## Direct properties — enable_log_anonymization / 313020320223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231320033322120-1122112030112101-1302100102121012-3132220301030000-3001113303330013-3300021031333123-2032011320133102-3302121220003233"></a>

## Next pages — enable_log_anonymization / 313020320223 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3101121100030333-3231330312100021-1120003212311331-1332211112201100-3020232221303003-0132200233113110-1011133120133331-0113033310032120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233231121010033-1002231203303003-1203322300000213-2123320010312122-1021032223021220-2201132011011310-2211333011011022-2203111123012012"></a>

## enable_management_network — enable_management_network / 112222203201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- enable_management_network

<a id="canonical-0312212202231003-0331322130101202-3220012301310333-1330003212202002-3111120113022031-3333231233121120-0000311222301201-3130230321003022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable management network.

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
enable_management_network = {}
```

<a id="canonical-0130310311300213-3212102013233230-3012233120031303-1101010300011100-3103113023131100-0130221113011112-0210012231323100-1133320223320331"></a>

## Direct properties — enable_management_network / 112222203201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003013010033313-0313233132103320-1211003120123031-1011010210220121-3010313131230013-3100033310300323-0020002103310202-0313322321203011"></a>

## Next pages — enable_management_network / 112222203201 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2110022213030001-1113110301320221-2101200333030220-3333200320222031-3001312302021121-3133213020013103-1112002120223030-0213300210232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102131103012222-1301122013102211-3232121001033022-1120311322333203-0111003003223201-2033300000331313-0312333010112200-0011001110121222"></a>

## enable_url_categorization — enable_url_categorization / 133103320210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- enable_url_categorization

<a id="canonical-0322300022333021-3031321220220003-3121001123003331-3210300311220231-2113221133000010-1111302322003320-3311113331332300-0113223000130223"></a>

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
enable_url_categorization = {}
```

<a id="canonical-1003001201322003-2121200020312130-1003300233212110-3200223303323323-0210331012133303-3021320212301002-1201303001012032-1322023101003102"></a>

## Direct properties — enable_url_categorization / 133103320210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023132301211000-1020322312123102-2320330001221001-3112123230001122-2021322311230312-2010333120021310-3300011203100010-1111002203131133"></a>

## Next pages — enable_url_categorization / 133103320210 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310112121100102-2033112231122100-3121331133003103-2131003220320223-3322301330011120-0220332333331200-2101112332203132-1101321020003310"></a>

## equinix — equinix / 011321202331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- equinix

<a id="canonical-1130000031010100-0011332132110332-2231000200201323-3113131110112033-1110321221310020-3310203010022321-1322323013012111-1203221300312231"></a>

Type: `"object"`. single nested block, Optional.

Equinix Provider Type. Equinix Provider Type.

Upstream description:

Equinix Provider Type.

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
equinix {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121030103022001-3002202202223020-1121230011332203-0133322213002331-3113123233202032-2222002122210220-0320121013021022-1003310133002200"></a>

## Direct properties — equinix / 011321202331 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100): complete subsection reference.

<a id="canonical-1112012100223202-1210000112123002-3320221030201331-3131023333200111-2310032131101021-1011211321021301-0122231313220110-3102310220200302"></a>

## Next pages — equinix / 011321202331 / 4

- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300302121020011-3223200001003001-0220230220102302-3021032033100122-2103133230120021-0310231203332221-2201121300003030-0021002020330131"></a>

## equinix.not_managed — not_managed / 130331300323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- equinix.not_managed

<a id="canonical-0331333131010203-2212030330002031-0221102320232202-2133332202010122-1303102032210102-1203113202013112-0212031011230332-1000021121220021"></a>

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

<a id="canonical-0000230301233320-0132002013330321-0022303103210022-2320321103030033-2012100312200211-1113222202322301-2213300102230331-3210333213013211"></a>

## Direct properties — not_managed / 130331300323 / 3

- [node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333): complete subsection reference.

<a id="canonical-0102312221110322-2321212102333003-3031112020212330-1003322200110123-2022031222032100-1321211103330212-0203200303133023-2121112233212201"></a>

## Next pages — not_managed / 130331300323 / 4

- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211211222312300-0021113023023113-1333332323000130-2302011002120230-1320310311123312-3312333111113002-3303122331313133-0313003132311112"></a>

## equinix.not_managed.node_list — node_list / 110011330323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- equinix.not_managed.node_list

<a id="canonical-3212000012311003-3321011123201023-3200231102031020-0131322033321303-1202022102202331-3000222031120310-1332203320211202-2220233213031321"></a>

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

<a id="canonical-0121003200221213-1302212132101313-1031123321123231-3211023300001112-1231132333002332-1120321120320010-3130300210111132-3332213001002010"></a>

## Direct properties — node_list / 110011330323 / 3

<a id="canonical-3223330031333203-2131321030233022-1113323331321102-3311231000130300-2102112201232020-3000312322031000-1222320010023003-3213113322311232"></a>

<a id="canonical-2200223022333333-3121211231103230-1013200132100312-2010001233313122-2232330213301120-1330020120320203-1232222033133222-0132212322133220"></a>

## hostname property — node_list / 110011330323 / 4

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

- [interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130): complete subsection reference.

<a id="canonical-2320033203311000-3010103121102122-2111200323211023-0231310103120311-2123021332120033-0031111210300320-2222011132022211-3123331323010130"></a>

<a id="canonical-3133000113133313-1120030321323132-1010030102203132-1212131302230311-2323312012011212-0230323100123011-0032132201113302-3302102300213312"></a>

## public_ip property — node_list / 110011330323 / 5

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

<a id="canonical-3203330120101123-1233222032000123-2112221201102003-0131221331031130-0023130031132102-1232223123223310-2131320121213221-3011302131102321"></a>

<a id="canonical-2010003002303301-2011002300130101-1220220310202002-2212303020120313-0201122331220230-2232201000222121-3223231021033110-0213232001201221"></a>

## type property — node_list / 110011330323 / 6

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

<a id="canonical-2231033013031232-3031120313233011-0003033100132133-2211121103321231-3131021022232103-3132200302032031-2100201131112321-1101323111302113"></a>

## Next pages — node_list / 110011330323 / 7

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122103322032022-1031231013110303-2222011013003300-2211201301111311-2200113130002002-0110220203112330-0023323333023230-3022333022010130"></a>

## equinix.not_managed.node_list.interface_list — interface_list / 301023213102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- equinix.not_managed.node_list.interface_list

<a id="canonical-3101212033100332-2001110302100231-2332020320302300-0011200212023322-3100203323312320-3222123132202123-0212023130132031-2110320112112312"></a>

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

<a id="canonical-2120313022111122-0110321111202232-0312121103003312-0020121231111311-1011000333111200-3302222133021022-2001100010322130-1203030211033112"></a>

## Direct properties — interface_list / 301023213102 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220): complete subsection reference.

<a id="canonical-1333010111102102-2223023123312100-2103103323203213-2102333220003311-2003123000021001-1313012323201112-0130331311023201-3112310031230210"></a>

<a id="canonical-2002303101232300-0022301032333202-0223302100100101-0013111303303100-0311310203330332-2233233001001100-0200130012033330-1202123311101132"></a>

## description_spec property — interface_list / 301023213102 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-008.md#canonical-0001022313310230-3203003012133033-2120111101311103-2303113011101123-3200332010101103-2133212122110332-2122023200122330-1331111131211222): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-3322031010311131-2122203121111130-0001310132003200-3021103022000310-1201230022322121-3033022033102132-0120012203211320-3321123232013002): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012): complete subsection reference.

<a id="canonical-2111023232301220-0221222223310312-3310222130020002-0033002101311001-1301002101331330-3201233000333010-3211022111110202-0012333333223012"></a>

<a id="canonical-2332322200332322-0131222223212230-3120012311323223-1000000111122001-3322302212030330-3332322321120312-2302100221112021-2331333201123220"></a>

## is_management property — interface_list / 301023213102 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0032202021130202-3322010201033231-2202220232311013-2131113132303223-3101303333120113-0111221112302300-3330322103232132-0310000112133213"></a>

<a id="canonical-2203323121131200-0331213000233313-2321213220211102-3333301101331110-0213203120200303-0110130233033212-0212303232133211-3222302032331203"></a>

## is_primary property — interface_list / 301023213102 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3010230203111000-3001033330300033-0323002333222001-0201001322132100-0210000333101233-1133003032113330-0112123321101033-3003230031031120"></a>

<a id="canonical-0312133021112113-1112322111320311-3221321003222321-2230031023110312-1101123110030132-1002230210023023-2333011300232032-2313113012121101"></a>

## labels property — interface_list / 301023213102 / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-009.md#canonical-3323331120001023-2103203213002302-1100311012120223-3131032211023223-3322031201322212-2101212123020100-2001123011023213-0222232203101210): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-1120332122331202-0100310320123121-0010230211321301-3033030022201102-2013202121322201-0330121012102232-3221023301222122-2031133113330102): complete subsection reference.

<a id="canonical-0220213331023112-0313303023230000-0230100022121322-1002103210013312-1001002302100001-1111001002022020-1311102130003313-3201003100221333"></a>

<a id="canonical-3013310131332321-0232220213302013-1000023103201131-0022000202231001-1113020110020011-1211223323232100-3100102113333013-2312202123123223"></a>

## mtu property — interface_list / 301023213102 / 8

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

<a id="canonical-1321232213031230-3201032011232130-3311100101003230-1023233222210330-2323100203313312-3302212131112331-2021231120322131-3222100330201130"></a>

<a id="canonical-0123132210203222-2203200130220131-3301320331332022-0223221012022311-1102202332022033-0312321303010322-2003300010113110-0101110220101202"></a>

## name property — interface_list / 301023213102 / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1113123221220321-3001202320101212-1013301013311033-1011012033233210-0322231000233322-3000000333010321-0300220022322333-3200122212120133): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1211220320012111-1133221130312110-3030113322023323-3002020201110233-0000233011131032-0321210201001020-3002031131313100-0220310210201110): complete subsection reference.

<a id="canonical-2302311131103112-3030320200210203-2121203001301321-3032131213310003-3010110323132022-0303120312100030-2020313122023201-3201231021010221"></a>

<a id="canonical-3010032313232130-3133321302203313-0022310231223120-3021133020210333-0003232211212012-3031033310130000-0103211111323231-2002331101221231"></a>

## priority property — interface_list / 301023213102 / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-1000302322130123-3211101132112110-1100312012320333-2033331000011310-3203123313213200-2120322131323103-1133023202200222-1130131331211101): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-009.md#canonical-2313031021311203-0012332220211021-2322000220201103-1311103011023020-3033310110123113-1223223231112210-0130200212132323-1300112200130222): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-2021020013311333-2023023000123300-2103322212112131-3300003132202332-1001030123222102-3322101101033120-2333000113130302-3232232321120023): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1320201032311302-2010322011120020-2300202030302120-2310032210331223-3102333123002122-0011212110103133-2031023020012131-3013110312323121): complete subsection reference.

<a id="canonical-0301011121202203-3022132011310323-3300030121303333-1221301030210310-1212301323133303-3311121033321023-3233011121111313-2100232210313100"></a>

## Next pages — interface_list / 301023213102 / 11

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220)
- [equinix.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-008.md#canonical-0001022313310230-3203003012133033-2120111101311103-2303113011101123-3200332010101103-2133212122110332-2122023200122330-1331111131211222)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [equinix.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-3322031010311131-2122203121111130-0001310132003200-3021103022000310-1201230022322121-3033022033102132-0120012203211320-3321123232013002)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-009.md#canonical-3323331120001023-2103203213002302-1100311012120223-3131032211023223-3322031201322212-2101212123020100-2001123011023213-0222232203101210)
- [equinix.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-1120332122331202-0100310320123121-0010230211321301-3033030022201102-2013202121322201-0330121012102232-3221023301222122-2031133113330102)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120)
- [equinix.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1113123221220321-3001202320101212-1013301013311033-1011012033233210-0322231000233322-3000000333010321-0300220022322333-3200122212120133)
- [equinix.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1211220320012111-1133221130312110-3030113322023323-3002020201110233-0000233011131032-0321210201001020-3002031131313100-0220310210201110)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-1000302322130123-3211101132112110-1100312012320333-2033331000011310-3203123313213200-2120322131323103-1133023202200222-1130131331211101)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-009.md#canonical-2313031021311203-0012332220211021-2322000220201103-1311103011023020-3033310110123113-1223223231112210-0130200212132323-1300112200130222)
- [equinix.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-2021020013311333-2023023000123300-2103322212112131-3300003132202332-1001030123222102-3322101101033120-2333000113130302-3232232321120023)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330)
- [equinix.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1320201032311302-2010322011120020-2300202030302120-2310032210331223-3102333123002122-0011212110103133-2031023020012131-3013110312323121)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331013332202212-2201120102033101-0011232002221303-2021223303003123-0323033300001132-0322210102331022-3233232323012023-3322201332132300"></a>

## equinix.not_managed.node_list.interface_list.bond_interface — bond_interface / 130013100233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2012000121311132-3130211012100113-1212002130211301-0310103030310010-1212101130131212-3003020003103130-3121033132030320-0301132101111230"></a>

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

<a id="canonical-1203211203210333-1012202012132221-2002012022003000-3033321201000210-2322233133331103-3020013103230321-1331100013303300-0110223223222002"></a>

## Direct properties — bond_interface / 130013100233 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-008.md#canonical-0120122111031101-1110223101111232-0312203021212013-2000100000213033-3001233233200300-2120000330132233-3323012033231000-2031010313133000): complete subsection reference.

<a id="canonical-1302020131102133-2213131313300110-3011221232323100-3002303112120020-3112123332312133-3023032303001322-0322013110330021-0313311222213012"></a>

<a id="canonical-1232001013201112-0213233302133000-1303102120220231-2311130102110320-2002010020001013-0001300110002330-0312330212113311-3110110220023023"></a>

## devices property — bond_interface / 130013100233 / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-008.md#canonical-1201311013303321-0202122210222212-2233322300111012-2332220202331113-1233010221130220-2013132330330211-0301001101023000-3132010211010233): complete subsection reference.

<a id="canonical-3322333230203321-0110002021332330-2222122303202222-1110222131302013-0312331223321113-1131033312123031-3311131100101003-1232002323333203"></a>

<a id="canonical-2300021322201330-1113220223013311-2032120031011313-3131132023123002-1231002103211120-1221210232231011-2011300213210201-2100230013322020"></a>

## link_polling_interval property — bond_interface / 130013100233 / 5

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

<a id="canonical-1103011210013323-2132303202123322-1110221310210332-1301112210311123-1212331022203301-1333022110100001-3003030220301120-3211222231332020"></a>

<a id="canonical-0110211002221200-2332111210103031-2332323112201200-1203111132331231-2112200213333012-1312303210223331-1001112131230033-0301121110103303"></a>

## link_up_delay property — bond_interface / 130013100233 / 6

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

<a id="canonical-3231131212003121-1213003123133133-3231212213010201-0032023222321202-0001301223212110-2033101330200211-2010030123122032-3132200233311210"></a>

<a id="canonical-1100103023002122-0121130202310330-1031010011220130-3230331021211110-0013032203011023-2313303032210123-0000111122103132-3300112201322011"></a>

## name property — bond_interface / 130013100233 / 7

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

<a id="canonical-3203303022103212-0130001110201030-1033003110320010-1021000212113201-0123101331020120-1310003033121011-1310312333303121-1303113302113310"></a>

## Next pages — bond_interface / 130013100233 / 8

- [equinix.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-008.md#canonical-0120122111031101-1110223101111232-0312203021212013-2000100000213033-3001233233200300-2120000330132233-3323012033231000-2031010313133000)
- [equinix.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-008.md#canonical-1201311013303321-0202122210222212-2233322300111012-2332220202331113-1233010221130220-2013132330330211-0301001101023000-3132010211010233)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0120122111031101-1110223101111232-0312203021212013-2000100000213033-3001233233200300-2120000330132233-3323012033231000-2031010313133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222333102130110-3032212223220103-2231113300311003-1303232313120111-1312313213113201-0322031311022123-2330112031222022-2023221303321020"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 010011113133 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220)
- equinix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1232003100111220-1102100300310133-0103233122023100-0333120013121233-1102001320333201-2102030203122311-1231121130103311-2032331233221320"></a>

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

<a id="canonical-2230022300030111-3213312031023333-1212211321320100-1313332201223232-0001200013203320-2233023302332031-1020303303012130-2110223010203002"></a>

## Direct properties — active_backup / 010011113133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101301020302312-1011133323003020-1110323002223320-2110221222211220-3133323313102233-1220301013221313-2121220213203111-3110202130220303"></a>

## Next pages — active_backup / 010011113133 / 4

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1201311013303321-0202122210222212-2233322300111012-2332220202331113-1233010221130220-2013132330330211-0301001101023000-3132010211010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321100023221001-2201013100313321-3023300330122301-2033203123231003-0031330113333002-3311221130302203-1312000131130002-1003220200211021"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 213332123112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220)
- equinix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-2022011020201111-3021230101020333-1130203222203211-2223102230133212-1222112020033213-0210232020230002-2201312033131132-2133012110313002"></a>

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

<a id="canonical-3300203223200312-2201302330100122-2002200232323300-2320032202113123-2202130122322332-3112230021021222-0122033110030230-1013322201302113"></a>

## Direct properties — lacp / 213332123112 / 3

<a id="canonical-1121313230323200-0012312332012303-3032320132321120-2212222000202300-1231122223111130-3331300202030123-0132231221231203-3100322321232331"></a>

<a id="canonical-0023203320202112-2123030231312330-1323130010032331-0322331030110130-1123213102100001-3220010313113112-0212222323233301-2312021003311200"></a>

## rate property — lacp / 213332123112 / 4

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

<a id="canonical-1130302003312212-3000010232133322-0201323102003322-1122320131000202-1110032012202031-1332301030300133-1001210003100122-2002012100202110"></a>

## Next pages — lacp / 213332123112 / 5

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-2023021031211122-1333203010202230-1003231310211011-3233213303300012-3033212010122033-3211022023110322-3032313202213122-1301332111211220)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0001022313310230-3203003012133033-2120111101311103-2303113011101123-3200332010101103-2133212122110332-2122023200122330-1331111131211222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302233011132131-1103223330221332-1130110231123123-0000222033022223-0123121230220220-0203203031321001-3331222010123031-2201301121330103"></a>

## equinix.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 220300030113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2203221312121132-2002132003303022-0320321010010332-3321113212130033-3103122130003322-2222121111300312-3222222311002002-3220313101213021"></a>

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

<a id="canonical-2112123333113001-1321332021123132-3230331222010301-3132331003223110-2021211302033001-1011321213321330-2132021100222311-0022313103200332"></a>

## Direct properties — dhcp_client / 220300030113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113320330032003-3312313033312131-2120102012101131-2113321002132133-0122131102231102-0010003313110232-3102333220211320-0133120330310101"></a>

## Next pages — dhcp_client / 220300030113 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200122121011020-1203131112003233-1031113112301033-0013121323130130-3103311133303233-3022313123123232-3330101133021011-2310112011012032"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 131010010030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2031030020221321-0130001013200103-2333233222101233-0223110313303200-3310220322022023-1202232131123231-1232122303312121-3300233200232103"></a>

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

<a id="canonical-2231110000301312-3322010011033320-0002213212211202-3313100020123000-0020113311021223-0113312003223313-2132321321222130-3011013121103032"></a>

## Direct properties — dhcp_server / 131010010030 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-1031123303113210-3210211222221023-3333200222101020-2300000112231223-1301133021032101-0100313333010021-3102233320120320-0021032100101213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-0220000000223332-0212110202303231-2001310300213212-0320022113330131-1121010320102030-2300330012211320-3313332101122233-2121200210030131): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210): complete subsection reference.

<a id="canonical-0002310012201110-1232031012023100-1123010121303132-2000232321033121-2231113230303303-0330032312123322-3203232103031103-0010122122131230"></a>

<a id="canonical-1010003301333033-3213132311203201-0001230231020200-2002330022113012-3213010113032232-0011120133333130-2121120323031022-0322000030231122"></a>

## dhcp_option82_tag property — dhcp_server / 131010010030 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3212111012022131-3210230123113201-0303220203130021-0233122112000100-2030133312000123-3112203203120200-0000323012323121-3020233120310031"></a>

<a id="canonical-0032213111311121-1321100123300200-3331030031023033-0031300212300322-1111113121331303-2200120300322021-3132221313211110-2331313202210213"></a>

## fixed_ip_map property — dhcp_server / 131010010030 / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-3312322311131021-0312111313333130-3213012310000102-1321132321003223-3132030001223133-0232311232101312-1010203230303310-2210320121203123): complete subsection reference.

<a id="canonical-3223102021230112-2030003110112202-1330123122230030-0223132012120110-1013222333200213-0321022303131231-3231021133011123-0101232201032003"></a>

## Next pages — dhcp_server / 131010010030 / 6

- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-1031123303113210-3210211222221023-3333200222101020-2300000112231223-1301133021032101-0100313333010021-3102233320120320-0021032100101213)
- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-0220000000223332-0212110202303231-2001310300213212-0320022113330131-1121010320102030-2300330012211320-3313332101122233-2121200210030131)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- [equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-3312322311131021-0312111313333130-3213012310000102-1321132321003223-3132030001223133-0232311232101312-1010203230303310-2210320121203123)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1031123303113210-3210211222221023-3333200222101020-2300000112231223-1301133021032101-0100313333010021-3102233320120320-0021032100101213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122212230230122-2303130021100332-1013230032200313-1020111210302133-3112130333220322-2122113112330313-3232033333233302-2110131121300301"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 211113213201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1232112313100222-3322200302023300-3210112333031311-3000210002313110-3002332301001033-0310211233131220-0010331333221212-3223222220020121"></a>

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

<a id="canonical-0001012230110203-0122013301200023-2300121030111120-3320300112212311-1201320021220303-0123130131220300-3321133302310203-1000003313311111"></a>

## Direct properties — automatic_from_end / 211113213201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113102301031302-2111032332311120-2000331010301213-3030330233023011-2220301030301120-0030222232122212-3021303013123330-1021303101212023"></a>

## Next pages — automatic_from_end / 211113213201 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220000000223332-0212110202303231-2001310300213212-0320022113330131-1121010320102030-2300330012211320-3313332101122233-2121200210030131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232230212130120-2132103013113211-1221210311211232-2130213010101302-0100210012131210-0201121122133130-0110002031033100-1032303020331101"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 310001131112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3021220102300123-2232030003323121-3123310012231223-1212031223010302-2321002113303223-0033332133020102-1331123211223002-1331212032332021"></a>

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

<a id="canonical-1010010120101201-3000112313320213-2032020023333232-2121022320203001-2110322322230313-2032031303210323-2231302303311332-2312030000030322"></a>

## Direct properties — automatic_from_start / 310001131112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201021031330012-1312000223303210-2323011201022112-0121110013113130-2312000123122120-1310331200020120-2130303201203323-0003203130110033"></a>

## Next pages — automatic_from_start / 310001131112 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212032320010310-3021110200330313-1203110313313200-3232202133003221-2031021333101210-0020121221311100-0032300010231203-2310210300130320"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 230121222303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2112022302222033-3312003122332320-3333132130323122-0212331201100130-1210001210232121-2222132323112032-0221022311110201-0011222333213230"></a>

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

<a id="canonical-3310302311021033-0101201033021112-2320303231330012-1330312033101213-0301232322101232-1211212133212223-2301020033312000-0232130210121113"></a>

## Direct properties — dhcp_networks / 230121222303 / 3

<a id="canonical-3020120120130102-0022302002112023-2320331300121201-1013123301101023-2111030122013220-0102131301320223-3111101113212101-0023120332303101"></a>

<a id="canonical-0323021103332300-0311202002023012-0013310133230020-1102321111213021-3110010302210103-3121113302102323-1110100022100101-1222131110122032"></a>

## dgw_address property — dhcp_networks / 230121222303 / 4

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

<a id="canonical-0232202120223111-1201331232131330-2312330021302220-3032311312311021-2232132030220103-0032010031123230-0013203223311102-2102203302310133"></a>

<a id="canonical-0211100321131033-3020031113021302-3322203100301011-1332310213003120-2130231030113023-3320123223003033-1132001311303223-1010202131011233"></a>

## dns_address property — dhcp_networks / 230121222303 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2312322332222330-1123132210233210-0112230320130221-0312300120102203-0133030112222131-0122031202220303-0100100323230223-1030223133221213): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2132220331022200-3120010202103213-3012200200023332-1123102120113212-0130231131332101-1221211211222200-3001011013132202-0001023111232301): complete subsection reference.

<a id="canonical-2031011002010332-2021200230233012-1131310221201311-3223102330032121-0021322332220211-3230103322003302-3010320130311111-2020323023323022"></a>

<a id="canonical-2321301101022103-2001031301020003-0313123002112231-0031100111333021-3321233202233323-1222101202322332-0002320222231222-2310312323320032"></a>

## network_prefix property — dhcp_networks / 230121222303 / 6

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

<a id="canonical-1321101200201210-2201121101033211-3033221211003302-3223122003201030-2023311200123032-2003202031003030-1303312310303011-2313230322102133"></a>

<a id="canonical-1231312111031323-0213132311210303-0131102013100221-0223310311023322-1232230022230030-0031301321011110-3301221233100302-2011023100003111"></a>

## pool_settings property — dhcp_networks / 230121222303 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-008.md#canonical-3132311222212330-1202203212020111-3030122222103200-0312132323310101-1122003112132321-0331321133131202-1221232131020002-0023211030030332): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-008.md#canonical-0222113122023023-1113213110202200-0233003022302233-3111130121012103-3112100202122331-2120022331323022-2312031002122123-2123211023113013): complete subsection reference.

<a id="canonical-1102230231201313-3233011320213311-0201032103103310-2331322333311212-0131120222231232-1121010000200232-2110223331013201-2222120031331130"></a>

## Next pages — dhcp_networks / 230121222303 / 8

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2312322332222330-1123132210233210-0112230320130221-0312300120102203-0133030112222131-0122031202220303-0100100323230223-1030223133221213)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2132220331022200-3120010202103213-3012200200023332-1123102120113212-0130231131332101-1221211211222200-3001011013132202-0001023111232301)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-008.md#canonical-3132311222212330-1202203212020111-3030122222103200-0312132323310101-1122003112132321-0331321133131202-1221232131020002-0023211030030332)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-008.md#canonical-0222113122023023-1113213110202200-0233003022302233-3111130121012103-3112100202122331-2120022331323022-2312031002122123-2123211023113013)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2312322332222330-1123132210233210-0112230320130221-0312300120102203-0133030112222131-0122031202220303-0100100323230223-1030223133221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202221231230200-0110000103310133-3001302001010030-2203132012012011-3111022313011330-1300023111332301-1212202012010101-1310133120121301"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 233013221303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-3002122201113201-2302213002133120-1013231021100213-3033130131333112-0311332111110203-1001030000322121-3312000013333313-1213301320103110"></a>

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

<a id="canonical-1220213122210021-3130210312101230-0320033310013200-0303211013022332-1233131001322222-2321212111122000-3321222311012102-1322003110031332"></a>

## Direct properties — first_address / 233013221303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111210221203033-0101011310002130-1031132122121031-1101010121222011-2230111300321300-2302013132203323-3331201202221322-0220100110211112"></a>

## Next pages — first_address / 233013221303 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2132220331022200-3120010202103213-3012200200023332-1123102120113212-0130231131332101-1221211211222200-3001011013132202-0001023111232301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200111010112010-3231010123110031-3323020201230323-0100033220010311-2020110212300120-1012300000233103-0203121010311231-2023122223001303"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 313220302230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0232231133001231-0110021003303002-3300133201011323-2330213323233332-0310331103101221-1110330032201131-2032032213102012-2333132310100331"></a>

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

<a id="canonical-2033023210233131-0012311321203013-3201310333103301-0203233131200021-0321222020200120-0210022130310102-3103003202003101-1121011000212112"></a>

## Direct properties — last_address / 313220302230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010113322201302-3220310023302232-0222130202120302-3022031122223010-2301133110021032-0321211112110301-3131132311202320-1331032123012031"></a>

## Next pages — last_address / 313220302230 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3132311222212330-1202203212020111-3030122222103200-0312132323310101-1122003112132321-0331321133131202-1221232131020002-0023211030030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331031103202322-1300221201130201-1312220221132122-1331130332031321-0033122313223301-3011201011232321-3031031312023300-3323320122113230"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 123320211203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-1012211212202100-3003300013333232-2133300303011021-3210103210011132-2203301132311321-2222103300000111-1101112110012103-3122323130133302"></a>

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

<a id="canonical-0021222213100033-2020112133113332-1110012310012220-1101232333303333-1131132001300212-3022010103030033-2333232202020331-3022232222130221"></a>

## Direct properties — pools / 123320211203 / 3

<a id="canonical-1211201001222112-3010300321033023-0032111131113102-3130122002331022-3301103200203113-1201002033020331-0312221232221002-3313021330101032"></a>

<a id="canonical-1221130121222111-1030220203301112-3201120332223021-2101231232310210-0320131121211100-0110133211302103-0123011031331030-2020221022103000"></a>

## end_ip property — pools / 123320211203 / 4

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

<a id="canonical-2213113112032121-1033010332333322-3321223003031112-2021101300000000-0303013013311110-3211223330032313-2011022310130022-3301120201320122"></a>

<a id="canonical-3213333323110312-2320233200000011-2133013003303211-3133321332312112-2111102000333203-0203003130103312-2223202231301103-1020231011223102"></a>

## exclude property — pools / 123320211203 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0333010013121223-3123300112120222-2120110001320333-2013003332131010-3333021113310021-2003003230211222-0011120110211233-1001000011110212"></a>

<a id="canonical-1311323120110322-1220200211022111-3303233013202110-0321032102032030-2221101121331021-3003132223000030-0103203122001012-3021033220020012"></a>

## start_ip property — pools / 123320211203 / 6

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

<a id="canonical-3233303020131301-2222101121001332-3121012030210111-0110132203030121-3330122010030323-3101212031202003-2002123300133103-2002122230122221"></a>

## Next pages — pools / 123320211203 / 7

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0222113122023023-1113213110202200-0233003022302233-3111130121012103-3112100202122331-2120022331323022-2312031002122123-2123211023113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121023232102011-0320021110131312-2332011221022032-1130233030021203-3102122132203330-0331011301013102-1311012210031220-2202322301131200"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 123102300100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1130101220012011-0121130302233023-0100102321011003-2102012303032333-0003021312231222-0113222302011322-0233300103011022-0330310213111223"></a>

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

<a id="canonical-1010020311101011-3003201322111312-2303220023121113-3323033332100020-1123212320221230-1303032311010120-1023222302021111-1312133132100133"></a>

## Direct properties — same_as_dgw / 123102300100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331111103110010-1201011010030131-2221030033102302-3032020222002100-0223123113320221-0120010201211220-2122201023301233-1013123202213101"></a>

## Next pages — same_as_dgw / 123102300100 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0213020131020030-2203320331002230-3101033033221031-3011201331112131-2220222110123222-3133101201013030-1222033231101113-0101133302112210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3312322311131021-0312111313333130-3213012310000102-1321132321003223-3132030001223133-0232311232101312-1010203230303310-2210320121203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333010213002103-2123001133121021-1210322001222230-2003332230120012-0301210323113031-0111302032121322-2302123321233301-1031230123310211"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 112322213232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2132111133320222-2202011023230220-1112021312302212-0113020020313232-1023300020033110-0100301002031213-0322231122011033-3202311231200311"></a>

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

<a id="canonical-0232011132211023-3102020322212101-2310132333323130-1230333101321132-3010120333103210-0103011031233223-0313233023103030-2100210221032300"></a>

## Direct properties — interface_ip_map / 112322213232 / 3

<a id="canonical-1101212003213001-1110331133312001-3312111031210203-2221310322013213-2002220013230031-2131303311332331-2012111001000031-2210133201100113"></a>

<a id="canonical-0331103312323320-2302201010012010-2101202220310301-1032100010132003-0022030112320312-3332310001231100-0200313010132001-2332001130031003"></a>

## interface_ip_map property — interface_ip_map / 112322213232 / 4

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

<a id="canonical-0203123230120220-1133332130103131-0312011110003210-2121303211103331-1112310310323030-2233332300310113-0221302002213232-2111210201201003"></a>

## Next pages — interface_ip_map / 112322213232 / 5

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-3130022222102221-2031233310010323-0031031023132101-1122032302321233-0200123331223113-0030130332033000-2300122200132122-2131101310033113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3322031010311131-2122203121111130-0001310132003200-3021103022000310-1201230022322121-3033022033102132-0120012203211320-3321123232013002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001033111213323-2102313212320020-3131111111011000-1123222122302030-3310023201300030-0211000100121003-2030200302320020-3013301013310123"></a>

## equinix.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 022021303232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0101023321101321-2021312313030032-0230013132202301-1011202300111200-0102333232033112-0130333120012130-2230132201300313-0300331010101031"></a>

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

<a id="canonical-1232131021321132-3120032202030030-2001203111101101-2312303033003203-2223233320000100-3103002213001302-1032103132030313-2332312222210211"></a>

## Direct properties — ethernet_interface / 022021303232 / 3

<a id="canonical-1233303313233130-2111122203301002-1203211302333020-0123120310000200-2221111011113123-2101323001011003-2120022322133112-3222012303110330"></a>

<a id="canonical-2131231113113102-1131302001122332-2212300010111111-0311112123212100-1131301311230020-0323101330223312-3122121303330032-0210021223110021"></a>

## device property — ethernet_interface / 022021303232 / 4

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

<a id="canonical-3033020232100110-3012033002211033-0121010201333332-0330312002132303-0321022100200120-1110100333020213-1323110211322223-0023002232202011"></a>

<a id="canonical-2101223002133302-0010232010212311-0132101330113122-2002230122100031-2222202012130022-0211301110001220-3230301103022132-0201321201312202"></a>

## mac property — ethernet_interface / 022021303232 / 5

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

<a id="canonical-1223201021301030-3012012001232033-3122223321312103-0100213000130020-2203212130102212-2233221133313030-0113130111022001-1233033311312233"></a>

## Next pages — ethernet_interface / 022021303232 / 6

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012022213030222-0133113131331232-0321230102211322-3012311203003021-0233311312210013-0330022112033230-2233302201023013-2201020330122303"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 021222132332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3301330003033013-0130203002212230-0323332112111110-2312230222113123-2001310200201103-2021222321121133-3012332220210230-2031002210020020"></a>

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

<a id="canonical-0300122031013232-1130300031011030-3220210311323033-2311022033210031-2120100330122330-2200032301003023-1223230302320123-1331133133321311"></a>

## Direct properties — ipv6_auto_config / 021222132332 / 3

- [host](resources--securemesh_site_v2--reference--group-008.md#canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-008.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211): complete subsection reference.

<a id="canonical-1212202311001200-1321131231120112-1202210130211312-3313111230230112-3131112300100100-1333220312021023-0021222101301301-2012020230201233"></a>

## Next pages — ipv6_auto_config / 021222132332 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-008.md#canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112121203201021-0203000303113010-0301323122132222-3313300312331321-1020313201131212-3222303103032303-1332310113131231-2022121213231022"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 231332101000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2201033211321231-3113311032033220-3021111121020211-1131032001130212-1110120122022212-1020113112222333-2203120320113102-2120031003211213"></a>

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

<a id="canonical-2200232203100321-0211112113033200-1100133322332232-0111132132121222-1221331010302133-0311031132101120-1301111321012031-2020023311130032"></a>

## Direct properties — host / 231332101000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310312301313321-0311132231230100-0003301020100013-3230203122110112-1120311132202303-3123212321333111-0000001320021231-3030213322020321"></a>

## Next pages — host / 231332101000 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303102220111121-3201131023220131-3013010102030020-3030223201111132-2102321223312112-3103221002103212-1022013120301131-0103022323313131"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 022301330331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0201232012002311-2130231221000113-2120010213312010-0103220303022022-3002232132202113-0132122111122023-3013122132121313-0002131111121220"></a>

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

<a id="canonical-1230121330021332-2322203222130230-1320030010122022-2203301203122031-1001023212131313-2222220320202132-2313223010122121-3311302000102330"></a>

## Direct properties — router / 022301330331 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222): complete subsection reference.

<a id="canonical-3112310302200012-1312013030103133-0032102200320121-2233220312333233-3133220331032333-1231332200321100-3021113212111320-2222311111233002"></a>

<a id="canonical-1300200131110011-3111010302232012-0112213110031211-1032002200103010-3210131133231013-2323323103031312-0200023330120102-3302101321122211"></a>

## network_prefix property — router / 022301330331 / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001): complete subsection reference.

<a id="canonical-3112322010112032-0310111313112122-3000312201230002-1201203231203030-1211102113131330-3321212112111130-0002320201303223-3111022013113030"></a>

## Next pages — router / 022301330331 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
