---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-3023122003201313-2110011202013023-2332203230303023-3322323131133003-0331201030013123-3100301131202322-0323132120201100-2012120201023012"></a>

## Next pages — automatic_from_start / 300012011002 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1302123303111020-3032120302131212-3312323033112210-1311221002221022-1332000103121310-0110130011303113-3002313002122333-1033231202210120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130331032102013-1233132313213233-0220033300032001-1310301033121110-0312130100313133-0030333323100011-1110110322232133-1002330132133232"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 203210033123 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-1120103031001332-2311300123233013-2212030230200220-3332331101320210-0320301302233212-3302221232122311-0023102131220322-1000003031203323)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-1001033021011103-2110300330002020-1210301200322110-2321120320011300-1331320231221231-0103323010021201-1010030010312310-3302310201012010)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1323312213010023-1113023330133013-3013300003201203-2233202013331323-3021333311231101-2113320020221331-0303323031330323-2102213210033010"></a>

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

<a id="canonical-2003300201233221-3323013322003200-1110311001012212-1301222300020000-0322233300213102-1010100001111232-2221302011020033-3201101331232333"></a>

## Direct properties — dhcp_networks / 203210033123 / 3

<a id="canonical-2311213123000010-1120322323132333-2100132130021101-2113332210020012-2312031212120233-0330321132330110-0120313121311233-1002032232222120"></a>

<a id="canonical-1110113122130102-1001330022230013-1212313010133233-3232302023231001-2132333233112302-3212131333003103-0333201322310030-1033123003301033"></a>

## network_prefix property — dhcp_networks / 203210033123 / 4

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

<a id="canonical-3133313223230011-1302212220200020-0202121212132123-0130223332131003-2200021223231120-0102103223313301-0311113122132023-2222001232332012"></a>

<a id="canonical-0223102200112033-0121333223313300-0330122230111213-3323313103110213-1133300213300222-3221010312111201-1122311333111110-2001102313222001"></a>

## pool_settings property — dhcp_networks / 203210033123 / 5

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

- [pools](data-sources--securemesh_site--reference--group-003.md#canonical-2311320332303021-0000200321212010-1313231212133003-0001213011102130-0130030013233122-2020102120322220-1320210222030002-1200101333020123): complete subsection reference.

<a id="canonical-2210331222110220-2012032331310123-1021230313301023-3301111101033232-2211333020222211-0021022101111110-3132323033231301-1010113133001231"></a>

## Next pages — dhcp_networks / 203210033123 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site--reference--group-003.md#canonical-2311320332303021-0000200321212010-1313231212133003-0001213011102130-0130030013233122-2020102120322220-1320210222030002-1200101333020123)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2311320332303021-0000200321212010-1313231212133003-0001213011102130-0130030013233122-2020102120322220-1320210222030002-1200101333020123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201311001103211-2130323220303131-2222120230102120-3032022101102303-1133100220232131-2110013220303313-0111103033300111-2301212133320233"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 312123111301 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-1120103031001332-2311300123233013-2212030230200220-3332331101320210-0320301302233212-3302221232122311-0023102131220322-1000003031203323)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-1001033021011103-2110300330002020-1210301200322110-2321120320011300-1331320231221231-0103323010021201-1010030010312310-3302310201012010)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-003.md#canonical-1302123303111020-3032120302131212-3312323033112210-1311221002221022-1332000103121310-0110130011303113-3002313002122333-1033231202210120)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1333000131323031-2101310123010020-1331010001313121-0001113223031030-2022112210112231-1101332120113102-2322030013121103-0120231321213300"></a>

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

<a id="canonical-0023033111303303-0332010303111112-0101300330302130-2220021332303100-0023131133120310-1222100132012210-0303310322231220-2301003300101020"></a>

## Direct properties — pools / 312123111301 / 3

<a id="canonical-2333331000032131-3202332123310111-3230000023322012-0212300233033211-3201222011012302-0332030201112123-3113200331111232-0012010033302222"></a>

<a id="canonical-3332310211300332-2203020133312120-2232213111322122-1310303202332030-2031322002001133-3000201300312223-0031130002101013-3100323031211012"></a>

## end_ip property — pools / 312123111301 / 4

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

<a id="canonical-2203230312311122-0203332033233330-1121213113311133-3223200121013203-1313201201033201-1213100210122310-1232111003122132-3112002001311231"></a>

<a id="canonical-2320111000020033-1010020122233223-3313211300022223-2112010131300313-3230011102030131-0103303002301303-0233113001301323-2121012332112300"></a>

## start_ip property — pools / 312123111301 / 5

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

<a id="canonical-0312103313010021-1232111122230332-3123120203212301-1000310322222001-1221310221213300-2033121303320110-2331313122023220-1232323022230013"></a>

## Next pages — pools / 312123111301 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-003.md#canonical-1302123303111020-3032120302131212-3312323033112210-1311221002221022-1332000103121310-0110130011303113-3002313002122333-1033231202210120)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1102111121011102-3010103203001032-0110123231330220-2113030000132231-0301122223132321-1211302113230202-0311221110022331-0330031030031200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113130200330000-2211201223222111-2311121322022221-1132110031210102-2011311032002132-1231102232303101-2002310212210031-1333131323112032"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 202300231101 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-1120103031001332-2311300123233013-2212030230200220-3332331101320210-0320301302233212-3302221232122311-0023102131220322-1000003031203323)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-1001033021011103-2110300330002020-1210301200322110-2321120320011300-1331320231221231-0103323010021201-1010030010312310-3302310201012010)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0302131131031330-2300333001112021-2230332300331111-0300003103222213-3023103111203122-1212211113220310-1213302101300223-3102231313033222"></a>

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

<a id="canonical-3233100123013032-1313123010122310-3032131213030102-0201013332031023-3320320011302213-0111310033113112-2330111322022010-1001032210330313"></a>

## Direct properties — interface_ip_map / 202300231101 / 3

<a id="canonical-1221030200112322-2101012230230010-3002012021310200-1132113233010223-0203223323310132-1211033211103232-0020103300232331-2003021302011113"></a>

<a id="canonical-3331030212013232-3010320211230030-2201321010222223-2131103133122233-3321311003130311-0230022312212231-3220103302133102-1132001021212121"></a>

## interface_ip_map property — interface_ip_map / 202300231101 / 4

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

<a id="canonical-3212133201221030-1101303332332333-0333221022300301-3301010112222310-3113231103122121-1101122231233133-3310100200130133-3303200320110132"></a>

## Next pages — interface_ip_map / 202300231101 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2000102311000103-3012010323231303-2112100311011332-1321233203303220-2013320221032202-1001011131212011-3033021211023100-0033023231122232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320321012000223-1311232110010213-3333302102213013-1331212030310020-3211100301020323-0103012121211330-3332011120001100-2020323222023103"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.is_primary — is_primary / 312010232320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.is_primary

<a id="canonical-3121220030212331-3122312203202111-0231313222110313-2300330132031022-1023200111313312-0301113203330132-3330110221203102-0312301313311111"></a>

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

<a id="canonical-2321222132132300-2013120023110333-3232321331132232-0233120120032021-2310122010323210-1031020132312321-1333132313111023-0213113323103103"></a>

## Direct properties — is_primary / 312010232320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311113210122210-1313222003011133-3013221023301231-0233310220322210-3331330220133212-0230112300120103-1112101130120112-3211321000012213"></a>

## Next pages — is_primary / 312010232320 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0100121333200100-1230110032122301-0212101332023212-2200302320332210-3211302220013100-0131301203121303-0021130323223021-3223010002012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113111123003310-3031122100011331-3303111032123210-1032031022030313-3223212301031223-3222001002310330-1230303203002312-2110130032012233"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor — monitor / 322312113320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor

<a id="canonical-2003211312301332-3132220001212230-2230300223312031-1032320001201113-0122231132133121-1111033013313001-1203301310211222-3021203313322203"></a>

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

<a id="canonical-2112132113233000-3332021230013020-2323032132200030-0321111011330133-2010113021300121-2012220013130321-3020331002332333-1223022321021333"></a>

## Direct properties — monitor / 322312113320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322121132130000-2030210203023021-1113213302101132-0200131112220100-2012331121233213-1333031221310130-1201113221110303-1001133101013022"></a>

## Next pages — monitor / 322312113320 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3013023002313210-3120000231323322-1023201031320102-2333112011310322-0000003133301200-2011000003111233-0121322022013113-0100111302033032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223132020133002-0233202101110012-1302310012011122-3203213023102233-0221003321202123-0231020320213321-2030331021130321-0021312020200302"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled — monitor_disabled / 021203302020 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled

<a id="canonical-1033213023213311-3121211313310311-3300312333312121-2013000301133222-2321333103133210-1300310221013100-2112123110203101-0012312021110310"></a>

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

<a id="canonical-0001102222231133-0322123000121221-0101300313003003-1230323030330002-2320111231222111-0312301300032122-3033200103310100-0133001020011120"></a>

## Direct properties — monitor_disabled / 021203302020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113231200332221-0120031312202330-1010313012012313-1320211000133113-1033212231222230-1022311201230102-1003200032022210-0001013102310332"></a>

## Next pages — monitor_disabled / 021203302020 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0021023001301300-2031020020121230-1011202320021230-0000101333030202-3222323000113220-2032203123011122-3011330030121020-3322300113000003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200033200201223-3203212210202131-0133001112320133-2230101020233021-3301112301103013-3301231201013233-2222103022110320-1123013321102132"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address — no_ipv6_address / 102321120020 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address

<a id="canonical-3203323101200331-2231200030221112-0210002220000232-1223021111032030-0321333203313311-3223001123230301-3113011332122002-3331232311122002"></a>

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

<a id="canonical-1020222202022012-0030131320200220-0132303023201331-0113020323213312-1232301002222201-3000120231003301-3303113131011000-2313130022210021"></a>

## Direct properties — no_ipv6_address / 102321120020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321020221200231-1231120000303330-2220300220231002-1033001120001123-3233120320011010-3121200000121331-0021232212223101-0310130323100301"></a>

## Next pages — no_ipv6_address / 102321120020 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2133113012200201-1023231232023010-0131332203103311-2002220033231122-2002311201203013-3211032112031212-1133321330202020-1111020030222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002223223202323-2030322122120020-3003013233223203-1131113323000201-0021110302110230-0120002321322213-1130330310320212-1221130131113133"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.not_primary — not_primary / 210100221131 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.not_primary

<a id="canonical-2232102323120102-0012302320122231-0020003021032322-2201213021201203-2233133332223310-2331312131130023-3031122302303110-1222220013000033"></a>

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

<a id="canonical-2121023230110132-2312110130222203-0301323110331310-1020032211200211-1211213122122102-1103310002301133-2301011330223303-2300221212211211"></a>

## Direct properties — not_primary / 210100221131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313311103213023-3320020313303300-2312033330032201-3121233202101223-0213212103110302-1230210023100310-2302032302322112-3013223013333103"></a>

## Next pages — not_primary / 210100221131 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0311230323132320-0102111232131003-1010222231213022-1221322230101210-3102112100013113-3202112210233333-0300112220120132-2321133030232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300332102313033-2011131012230012-3020322320100002-1213022230331223-1022000011313313-1223312131022232-2303111013001313-1223032310230003"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network — site_local_inside_network / 303032102132 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network

<a id="canonical-3223131023011013-2002133303330001-3003202020010120-1021101133002001-3111321223131122-3023300021001123-3010311123223311-3222030231333321"></a>

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

<a id="canonical-2010231133113020-2100002230303200-1033210230233122-1001222301221322-3230311020201033-3121001033211332-3101021230123233-1030232211211221"></a>

## Direct properties — site_local_inside_network / 303032102132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132111300332133-0301033023221300-1033330021323332-1113133120202232-3001131010332023-2022130232103221-2230222332220221-2311203011122220"></a>

## Next pages — site_local_inside_network / 303032102132 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3312031322030032-1122021022301030-2130313101300000-2220012112321030-0230023021300303-3121120123210231-2121023232121100-2231330030321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113000122021200-0022232012302301-0021002331110101-1211212102232301-2232311011213121-2303320201100022-3233233203223112-0011221030101302"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network — site_local_network / 103320012033 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network

<a id="canonical-1120103332113103-1230033332113021-0310122022301101-0202033021121013-0133112233211111-0220123320210002-0212033320201313-3201132112331222"></a>

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

<a id="canonical-0331333030333113-0232123110321222-2300301032032213-1323202220310011-3300221122220323-3021002321331313-0032233323320033-0132003212310233"></a>

## Direct properties — site_local_network / 103320012033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311013030003313-0023030020311020-3112011131010113-2002322012320113-0121120302132123-2122213121310121-0103111221223331-0102310013313001"></a>

## Next pages — site_local_network / 103320012033 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233332010020111-1110012221222200-3322201102113123-1323021103113001-3123213110330123-2131130031213221-1033203103233223-1111332313120221"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip — static_ip / 120201020110 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="canonical-0332212332202103-2032113102321301-0021302111313300-0021223211201310-3200300002202113-1212332033322020-2333302003102032-0213100131111312"></a>

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

<a id="canonical-3010331123222231-2211330231213311-0020221202233120-3132133300300310-2230102300030312-3113032221311100-0232301111231020-3120201300200021"></a>

## Direct properties — static_ip / 120201020110 / 3

- [cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-3101121311332132-1123233223220332-3030112112221101-3312110203202322-2200000020023302-0021122020123323-3131212302303110-3230100001132133): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0101201111001210-3211032312311310-3020200300111301-0220313023030003-3310210211132130-0222311112321331-3032303132113013-0201220301102221): complete subsection reference.

<a id="canonical-3100020012012033-0032311121210232-2212031330322122-3200132322323230-3232220113011102-1113101103010300-2000333311301331-3323310122111033"></a>

## Next pages — static_ip / 120201020110 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-3101121311332132-1123233223220332-3030112112221101-3312110203202322-2200000020023302-0021122020123323-3131212302303110-3230100001132133)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0101201111001210-3211032312311310-3020200300111301-0220313023030003-3310210211132130-0222311112321331-3032303132113013-0201220301102221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3101121311332132-1123233223220332-3030112112221101-3312110203202322-2200000020023302-0021122020123323-3131212302303110-3230100001132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112000330312223-2130231230033122-2120223333310102-1000030003002333-1113310001232330-3010111220112303-0333312212210001-3233020220221120"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip — cluster_static_ip / 002313333002 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-2223200300303121-3001021110321112-3113131000201332-0201321013323122-0323030020220021-1301003322013003-1220303000102103-3220002223202003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1130313032130111-1231331210333201-1303013131231110-0112222213010231-0130020210232310-0322311103302230-3331001031113110-3103132332031221"></a>

## Direct properties — cluster_static_ip / 002313333002 / 3

<a id="canonical-0211113131311222-3030112323203213-0302322100230001-2100211323231132-0200321333212132-1201230332333321-1231003102102112-0033122302330133"></a>

<a id="canonical-2102120313202223-1201310201123001-3100110121210120-1221110202301300-3211212031020230-1331011221131003-1003213212032120-2301203013121202"></a>

## interface_ip_map property — cluster_static_ip / 002313333002 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-1133123032020220-0321223012012121-1302223202211232-1033323303132212-2132332100132232-3211211222202133-0101020213232003-1202100123332200"></a>

## Next pages — cluster_static_ip / 002313333002 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0101201111001210-3211032312311310-3020200300111301-0220313023030003-3310210211132130-0222311112321331-3032303132113013-0201220301102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010232321320320-2231002220130033-1221300303311323-1311020002120022-1220001130303020-2010210233031311-0321102313103302-3010233210102231"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip — node_static_ip / 121231011322 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip

<a id="canonical-3021132310001110-3212303202220312-0121123120302131-2130010210232113-0321112210310330-1231011002333222-0130031300123012-0113232312121301"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-0303123131013211-2020212221300103-3231320232310002-3131022121013101-1301020010223332-3302331133233300-2313130132213312-0032201032332010"></a>

## Direct properties — node_static_ip / 121231011322 / 3

<a id="canonical-0231101102111110-0230202220002320-1222031112100321-3131002330103301-3311011101010211-3321212112131011-0112231332012320-3220123020232000"></a>

<a id="canonical-3203013110112002-2120131201013211-3202011222301302-2012231222122021-1000300110312101-3011011003201103-0201123120022323-2000233132103233"></a>

## default_gw property — node_static_ip / 121231011322 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3120202322311110-0201131002210222-0103031311000003-1100003303320222-1123203101231100-2301132101132320-1301233311203133-2111011031210303"></a>

<a id="canonical-3032201030312130-0001110331013003-2020031012031111-2322301110122012-2031222000213222-1230202303012011-3210220032211120-3302200001103231"></a>

## dns_server property — node_static_ip / 121231011322 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2131313223223213-2101022222130122-1230132233200200-0111331210203112-3232210233030302-0002330032021121-0313220122031321-3202203011210202"></a>

<a id="canonical-2123102102132030-3010312231130000-1110112301301313-0323033123303023-1030120331023120-0021330201121001-2003312031221130-0020202313310132"></a>

## ip_address property — node_static_ip / 121231011322 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3112103320321213-1020313230110130-0133330213200133-0330120322112000-3010001020202112-0113320232001123-2121303010331130-2331322213112110"></a>

## Next pages — node_static_ip / 121231011322 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0120113002313030-1322302200111002-3301101312033012-1002300312011100-2030033212113011-1201131333202300-1312130301100302-2022121022000212)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3201201203322302-1032132021223012-3001303230210133-2201333222231310-0101013311022321-1233131121321330-2130232323203200-3131010312102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033110220102111-1012312102113113-1221122112312323-2133023212121212-2120231201131310-0312230120220312-0203001130012123-3322003223230322"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address — static_ipv6_address / 032310332323 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

<a id="canonical-1312030021101203-2221322310022020-3030120321323131-1030211333000123-1223333123301331-2310211322012113-3312202100313303-3211332001200322"></a>

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

<a id="canonical-1221301031023302-1132122111030120-3322103331110000-2312020021200233-0013120331022332-2300033021103310-2222320202202002-3132312123011003"></a>

## Direct properties — static_ipv6_address / 032310332323 / 3

- [cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-1000033200100200-1300201223210120-1332121201020313-3330000333131103-3221102312000201-1311011113123302-0212312030203003-2323321023212010): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2233220230200203-0321102022132220-3200221020112011-1321333232202213-3013002112023101-0312300020123133-0131102030032112-3133233300230120): complete subsection reference.

<a id="canonical-3332112230200320-2133322031221222-0222310030033111-2122002113003221-1122022003221110-0221030013001031-0133333321323233-3233330313000211"></a>

## Next pages — static_ipv6_address / 032310332323 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-1000033200100200-1300201223210120-1332121201020313-3330000333131103-3221102312000201-1311011113123302-0212312030203003-2323321023212010)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2233220230200203-0321102022132220-3200221020112011-1321333232202213-3013002112023101-0312300020123133-0131102030032112-3133233300230120)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1000033200100200-1300201223210120-1332121201020313-3330000333131103-3221102312000201-1311011113123302-0212312030203003-2323321023212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301233033223132-3231310221100233-1122131210301332-1130323222130132-2220233010210011-3322131103312300-1121130121300021-1331132333012301"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip — cluster_static_ip / 023203302111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-3201201203322302-1032132021223012-3001303230210133-2201333222231310-0101013311022321-1233131121321330-2130232323203200-3131010312102031)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-1100031010302103-1303010010313231-3303111301221230-1312100122311221-0113021132102202-0230310212111123-0313233220103223-3030130210010223"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1020300221021131-0232233112313133-2133331002333002-1003123013110011-3002022212113220-0102331200003121-3222032202000223-3111021221300023"></a>

## Direct properties — cluster_static_ip / 023203302111 / 3

<a id="canonical-1221000031121220-0101301331202111-3032320100322211-3311231103121300-3303232000020302-2011332232030100-1000002330003130-0303033132200331"></a>

<a id="canonical-2320313103113113-3121333020112233-1313331003222102-3200003130010213-1012010333321102-3332202011323323-3211320033220212-3212021200120213"></a>

## interface_ip_map property — cluster_static_ip / 023203302111 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-1031131212020130-3202320111121102-0313302321230130-2010030112023322-0310330322321031-1120223013103003-3322211010123213-3301330222310113"></a>

## Next pages — cluster_static_ip / 023203302111 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-3201201203322302-1032132021223012-3001303230210133-2201333222231310-0101013311022321-1233131121321330-2130232323203200-3131010312102031)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2233220230200203-0321102022132220-3200221020112011-1321333232202213-3013002112023101-0312300020123133-0131102030032112-3133233300230120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332310302300212-0223112303312232-1113301112330230-2102233101121112-1031100201001302-2230201113332312-0111202311323301-0233303130022031"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip — node_static_ip / 030223221312 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-3201201203322302-1032132021223012-3001303230210133-2201333222231310-0101013311022321-1233131121321330-2130232323203200-3131010312102031)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-2201132020311010-1233210100322101-1320121112332223-2001103100000132-2123213231330033-2210311133112311-2330303023231000-3203330200210102"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-3030322232321102-1030112011013321-2213133010120122-3201201232001310-1030212030233202-1211233100110201-3103130202313100-2323223313000130"></a>

## Direct properties — node_static_ip / 030223221312 / 3

<a id="canonical-1020031120133120-2202000302311230-0102012120012212-3111000321212212-2011032231221221-0022023103203110-2203333032003211-0313023333122233"></a>

<a id="canonical-1210320033112123-3002212230101120-2032300220233221-0110203310121100-2020100101023133-0020321321213002-0101211231033131-3121120021221302"></a>

## default_gw property — node_static_ip / 030223221312 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1220022201311002-1323130132121130-3120122333211322-0003010213001220-2211322203200030-3133211133302131-0131313222213100-1303113321233122"></a>

<a id="canonical-0133000201231310-3321013131023302-1323013322120120-2311023220211030-2222113123322110-2120002231230223-0303111010030022-2330220123303321"></a>

## dns_server property — node_static_ip / 030223221312 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1120121003000331-1202331000213133-2023022332221223-0130031031002003-3103012323221321-3220203031303322-3231101332100100-0202111213000202"></a>

<a id="canonical-2231100001230002-0021023210312131-0313002213013333-3031123012220121-3133320012232333-0032223123031311-3221001200002333-0020313330202311"></a>

## ip_address property — node_static_ip / 030223221312 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2001110122222300-1002202222100201-2201010201013330-0202133030233020-2223333000221320-2302130202330321-3220210300300022-2100100323313320"></a>

## Next pages — node_static_ip / 030223221312 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-3201201203322302-1032132021223012-3001303230210133-2201333222231310-0101013311022321-1233131121321330-2130232323203200-3131010312102031)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3131103110322120-0030222211230020-0310213322002033-0123021000100201-3303000131103112-2130321232100211-0011223000231310-2202322113132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132102022030020-3232222102332030-0103222322313213-1102103032032302-0210232330103220-1103022330021231-2022313213330221-2201210222002000"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.storage_network — storage_network / 311123310222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.storage_network

<a id="canonical-1022100303122012-0113121323103320-0232110132001322-2220301230203210-2122302022031232-1011123223330310-3313301000310231-1210320202023002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3301000333113020-2032303201033222-3011333313333333-0123103213321301-3233312133310203-0212012231113200-2013110221100033-0022000010222032"></a>

## Direct properties — storage_network / 311123310222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222021311220221-0223130001020011-2223111233231313-1202033120121001-2113130030010222-3112310230121212-3211221213212230-1000020310322222"></a>

## Next pages — storage_network / 311123310222 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2300212021202032-3100313111201113-0022300212212223-3011122321203303-1312121201033322-2313031031223222-0020313320020132-3331213303220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131003232122313-0120202302032212-1011120012031010-2111003101130022-2010000001032311-3001000213222021-1031133112002100-0220102033111320"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.untagged — untagged / 111122330213 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-1122303030301200-1022322213231100-2121110100132321-3203211333133310-2033030110010033-2202200233331012-0030323312310323-0323130121130221)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- custom_network_config.interface_list.interfaces.ethernet_interface.untagged

<a id="canonical-2130000133101322-1021033312120211-2010122313020132-1203201320032021-0001011231012110-1321100212211120-2021031211000322-0130303331321000"></a>

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

<a id="canonical-0031332300223020-0131200323101312-0213003100301303-3333012322312302-1300102322133020-1313203310211321-2231003331020131-0311223032020322"></a>

## Direct properties — untagged / 111122330213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010232122231223-2132323222103232-1022212003300010-2201033332333011-2302212130330303-1113130233232000-0011333200321311-2120121123212210"></a>

## Next pages — untagged / 111122330213 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3323103022122213-2312213103212102-3223200302230303-2310320111131323-0203333312003121-2232100200333030-2023001212023213-2121332313021321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2211132100230232-0332032112001113-2331331231012012-3320322002220323-0120102222300321-3113101202102323-2322232121101001-2230333201203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121202120023100-3011202111203301-2221203200302121-2121031230120001-0123303322103223-3332033111133333-1122111221330102-0023013102212003"></a>

## custom_network_config.no_forward_proxy — no_forward_proxy / 202012130021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.no_forward_proxy

<a id="canonical-3002103003330320-2131230000033032-2322102012023310-0313013311312032-1220000310202121-3301331331132001-0220313213323023-3231012202103213"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2302011130301200-2231002130112020-0121001022113032-1311332111030022-0132331312200021-1000333003131230-3031012213000332-1112312102000002"></a>

## Direct properties — no_forward_proxy / 202012130021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211221130322332-3012100112131032-2313231002101320-0101203231313123-1130200200123200-0331210000310331-1200102120013031-1231110201132210"></a>

## Next pages — no_forward_proxy / 202012130021 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3203231233310210-3312303100312211-0212030103112311-2211012220123012-0310031332022133-2222302312121122-2022301033230003-3123301021221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312123023010222-2130332003331323-2013120132311101-0332120010122323-2101013122010331-3111201330132033-1220132332311302-0333201233021232"></a>

## custom_network_config.no_global_network — no_global_network / 130131203130 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.no_global_network

<a id="canonical-1111330110231302-2333020021330122-1133202012223120-2202231311301133-3331131300001310-3312202103033323-1222130233120112-3112230230100201"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3200003332301303-1223203111220221-1012022123133222-2203322021232132-3131223320110100-0033310002023323-3300103202231300-1031020030101212"></a>

## Direct properties — no_global_network / 130131203130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010301313220012-2122122000310031-0331110130121111-3002132203010012-0312323222001223-2003130022010212-0012003211020102-3022130131320132"></a>

## Next pages — no_global_network / 130131203130 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0211120021231100-0200211311001101-0221332130221012-3033122221023220-0003033302132322-2132311332201133-2020001113203320-1003003232002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320000303110220-3203302201331012-1333213132232012-3012333303032012-1331323002213011-0000132101233101-2220112111223130-0101023310312233"></a>

## custom_network_config.no_network_policy — no_network_policy / 300202033223 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.no_network_policy

<a id="canonical-3301231103303223-3023033332232001-1223010023210301-3133310220311032-3121222222323000-1020211010203131-1333110033110213-0320132302211112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2230321111230022-3322002233312202-2000322022121200-1102332002002101-1033100131001010-0012223133201230-1133103302100333-1203000110032212"></a>

## Direct properties — no_network_policy / 300202033223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113101320212023-2221122210300231-2022211023322033-3102133223010010-3230210113023122-2323230233132321-3231120102233003-3021231302210032"></a>

## Next pages — no_network_policy / 300202033223 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003311003330302-2113023122002130-3021000231132200-0330332111222100-2131221202232021-0112001132223333-3031232332001121-2320131200120220"></a>

## custom_network_config.sli_config — sli_config / 321223321111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.sli_config

<a id="canonical-2231020131223023-2310221002020202-2033232322303321-1230123330232230-0220223213033322-3313020230211112-2201312013110103-0001232100102322"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

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

<a id="canonical-1112032302311233-3230321123222131-1312230313223233-2113230010200013-2311231322211203-3230132023101031-3322100010222212-2002120230031303"></a>

## Direct properties — sli_config / 321223321111 / 3

- [dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3212331332331113-2033030212033023-3311032332120221-2210322303201233-2302111132023030-3111303213031223-3103230120111312-1311012321313001): complete subsection reference.

<a id="canonical-0020232022012222-1331113231131031-0310203331102000-3031000211201102-3300210010133101-1202333131010002-1031121023313221-1320032111232303"></a>

<a id="canonical-3021210131232301-2100110213023311-2232300320132202-1231212330323113-2222030003302313-0210223321201222-3311331323112203-1110000022303220"></a>

## labels property — sli_config / 321223321111 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-1331123212201131-2102011002323312-3123323200010332-0012232233103333-3210332003033233-1202210022101303-3131102232103322-0323220310210221"></a>

<a id="canonical-3233112101300000-3200001110310311-2133200211310201-0012213111113212-3331331010221200-0222221013033211-3312202101020122-2323300112013201"></a>

## nameserver property — sli_config / 321223321111 / 5

Type: `"string"`. Computed.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3110012202003332-3011113110031011-0211332003201332-0133002321131023-1100002123013331-1230131221100302-1123300022002013-1311120213212231): complete subsection reference.

- [no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3301200312111120-1233323211331000-3300310013113320-2012302310030310-3012002331303012-2033300210110331-0200000113212031-1203023331202211): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3001232121312021-3202121002021100-1210101311303103-2231022010030331-3120200121330103-3003113111223303-3123213112112321-0121130313112000): complete subsection reference.

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221): complete subsection reference.

<a id="canonical-0100111223313012-0313210321233023-1132301322012123-1023022130231201-0022223031211113-0312110322320222-3320000000133330-2321211033013032"></a>

<a id="canonical-1211012301130032-0131100002300332-0223231320321110-3031030100132223-3331211003223322-3022301330321032-2011230000002011-0222322312220223"></a>

## vip property — sli_config / 321223321111 / 6

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-2223123201201201-3232020001232300-2111102332010112-1200200003022010-0033133120222131-3310311132223011-3201012013221033-0321030102311000"></a>

## Next pages — sli_config / 321223321111 / 7

- [custom_network_config.sli_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3212331332331113-2033030212033023-3311032332120221-2210322303201233-2302111132023030-3111303213031223-3103230120111312-1311012321313001)
- [custom_network_config.sli_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3110012202003332-3011113110031011-0211332003201332-0133002321131023-1100002123013331-1230131221100302-1123300022002013-1311120213212231)
- [custom_network_config.sli_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3301200312111120-1233323211331000-3300310013113320-2012302310030310-3012002331303012-2033300210110331-0200000113212031-1203023331202211)
- [custom_network_config.sli_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3001232121312021-3202121002021100-1210101311303103-2231022010030331-3120200121330103-3003113111223303-3123213112112321-0121130313112000)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3212331332331113-2033030212033023-3311032332120221-2210322303201233-2302111132023030-3111303213031223-3103230120111312-1311012321313001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112321010123000-2010200003321010-0213333211331232-0202100301231120-1111210030231313-1133113121003032-2100213321300333-1112210110132312"></a>

## custom_network_config.sli_config.dc_cluster_group — dc_cluster_group / 132313022100 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.dc_cluster_group

<a id="canonical-3100202303332203-2000302303301220-2311022013132030-3203003200033220-3203202231112111-0101312312121123-1232212021021131-1111301323231012"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1013333203103202-2010113132132310-1021230032202113-1020201223131033-1123012132230020-0003010113230313-0131030322303312-3313101333011131"></a>

## Direct properties — dc_cluster_group / 132313022100 / 3

<a id="canonical-1303233011311132-3103232201120233-2131203311322121-3321323131010302-0032223301223232-3321221123312022-3313031000113100-2313231001111122"></a>

<a id="canonical-1012011323220001-2220302302203131-1013230033101001-1321201130230011-3103000000300010-2010202033203203-3022013220003103-2133012001212032"></a>

## name property — dc_cluster_group / 132313022100 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3033311110013010-1012302223232100-2210212301223313-1103022222333103-0223210013310102-3102120023033033-2132022103211211-1110321333310110"></a>

<a id="canonical-0131022112103321-1030331221020303-0111232122203300-2120330023331320-3200120123231120-2002113210333323-1311031311133222-2303101020102101"></a>

## namespace property — dc_cluster_group / 132313022100 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2212003300210333-1030201133223202-0300021121122123-3333322031002322-3100213020112001-2032330220311023-1313321223222322-3231031000211311"></a>

<a id="canonical-1332110012312332-1003310023011102-3112310120213301-2101203021121132-3023313101110331-3330122230311013-2312231033031203-0023311332213100"></a>

## tenant property — dc_cluster_group / 132313022100 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1123012032330313-2013020300231331-1312222112121311-3020213331101122-3023132113110322-1021132320321023-1233010011210320-0010220302113331"></a>

## Next pages — dc_cluster_group / 132313022100 / 7

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3110012202003332-3011113110031011-0211332003201332-0133002321131023-1100002123013331-1230131221100302-1123300022002013-1311120213212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022011230321122-3233201110013303-0203023320231323-1301322213321022-2021311001000302-1122222021033220-2121330112133131-2002203300230202"></a>

## custom_network_config.sli_config.no_dc_cluster_group — no_dc_cluster_group / 032321331202 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.no_dc_cluster_group

<a id="canonical-1332302013221103-0200201310332201-1121031320221112-2223100202020231-0022201321021300-2312103211120031-2100323333222332-2330303113012023"></a>

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

<a id="canonical-2102103002233222-3031232332000233-2001333220213002-3211111332011023-0223133302001033-2203101002030030-2310233302133112-2200010133031011"></a>

## Direct properties — no_dc_cluster_group / 032321331202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231112221200100-3333133111100023-2323200112002202-0320220002332231-2333322321333223-0223331212211030-0123330132222001-1110311221130203"></a>

## Next pages — no_dc_cluster_group / 032321331202 / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3301200312111120-1233323211331000-3300310013113320-2012302310030310-3012002331303012-2033300210110331-0200000113212031-1203023331202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102210322221213-1332103223230333-3121222022333233-1102213020330221-1000000303202032-0002320022222330-2101211330201211-3221231013011113"></a>

## custom_network_config.sli_config.no_static_routes — no_static_routes / 311212302022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-3230310013132213-2002201233330311-1030020213233111-2301012111331323-2222132322030220-0023312232113303-0213101012002023-3131212131210202"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1223203211013322-3032232310331310-1030211321022232-3003312033233032-1110321210322103-3023110103232020-0300132111021103-2321032030210212"></a>

## Direct properties — no_static_routes / 311212302022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321301231211020-2011000323301001-1012322022210110-1310120011210323-2023320211321322-3021232311030232-0202332201003012-3111321123130002"></a>

## Next pages — no_static_routes / 311212302022 / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3001232121312021-3202121002021100-1210101311303103-2231022010030331-3120200121330103-3003113111223303-3123213112112321-0121130313112000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003001211131222-2211302223220132-0333011113311021-1313002231023313-3001131030331013-3021020311120031-3332202120221231-2130010020322113"></a>

## custom_network_config.sli_config.no_v6_static_routes — no_v6_static_routes / 230103203030 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-1200200111011100-0101312303120211-1213313102212320-1011023132212002-1120133211103013-3133220202113210-3320213112333221-3133023030120122"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2130131013032201-1013210000332102-0122033211012132-1132211320011212-1030013022003200-0202323220112133-1032121222202221-0203313021123323"></a>

## Direct properties — no_v6_static_routes / 230103203030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133001000023230-2001301010131010-3110233201202103-1220132213320023-1331030030130220-1222210311322013-0332021101001121-1003130332131031"></a>

## Next pages — no_v6_static_routes / 230103203030 / 4

- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231001113303023-0120301131333231-0211030321021021-2101223330301211-1212111131232002-1023322233320312-2000021210121012-2331031111311030"></a>

## custom_network_config.sli_config.static_routes — static_routes / 302233312333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.static_routes

<a id="canonical-2231231200033032-3000233021002020-2032220320010111-3102232203323312-3212131022033203-2021012200003232-0303313223003331-1300232221330203"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-1023232220210231-1100133010302120-3211331300301233-0213122022313030-3112123313012012-2032122231201303-2120121110011331-3013132231211320"></a>

## Direct properties — static_routes / 302233312333 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121): complete subsection reference.

<a id="canonical-0113122130223222-1232002200301001-3033111233201000-3300033203003301-0202231131322301-1313120210231221-2020300231010030-3333221031013321"></a>

## Next pages — static_routes / 302233312333 / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233223000321222-1130232320110110-3210022233123112-0300113020002132-1103230230021213-3033131013233210-1311100310212333-0313211021121203"></a>

## custom_network_config.sli_config.static_routes.static_routes — static_routes / 100222130200 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-3221103033130222-1020231203122332-1101230320200200-2002120131102031-0111111130330330-3112020201330231-2313313002230230-3330012031302003"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-1231110132212011-0311132001222311-1313223122022202-2312202130131222-3202013330220231-1313302112313221-0232213322301113-2311320302313100"></a>

## Direct properties — static_routes / 100222130200 / 3

<a id="canonical-0103131211212203-3312301113212012-0230102131322002-1223231233332233-0011310301322221-0221133000102020-1211300303031011-3202103312001132"></a>

<a id="canonical-3323023030223232-3300230012231003-0013002322122322-3003203330301103-0233323231303301-1113321002121112-0001312212211231-3033030023102100"></a>

## attrs property — static_routes / 100222130200 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-1101330030220310-1210031131310300-1023031213322200-1112312230012213-1312323100331311-1221212013301023-3333230301001121-2030122333022022): complete subsection reference.

<a id="canonical-0112230100301123-2322032012203323-0021211010302311-0201232030013203-0123310220312001-3103311212100222-2022312111303001-1120202321212130"></a>

<a id="canonical-3303022100201003-3032120230233332-3120002023010020-0101113112223223-1231232220030103-1132022222313000-3201311122303121-3013033131020113"></a>

## ip_address property — static_routes / 100222130200 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3313011323231212-3322330333323231-2033212013121001-0000032312303021-3133221022301223-1000303210220231-0333310133101230-1033312101230012"></a>

<a id="canonical-0312201210233101-1230231132033121-0301131013211101-2223222300132002-3003011132210001-2331133022123133-2033200023332122-2112211023232203"></a>

## ip_prefixes property — static_routes / 100222130200 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222): complete subsection reference.

<a id="canonical-2232032300210310-3310203132011113-0312313210031032-1031221203302223-1233000121133030-2031301013310031-3332330232013221-1223011201132112"></a>

## Next pages — static_routes / 100222130200 / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-1101330030220310-1210031131310300-1023031213322200-1112312230012213-1312323100331311-1221212013301023-3333230301001121-2030122333022022)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1101330030220310-1210031131310300-1023031213322200-1112312230012213-1312323100331311-1221212013301023-3333230301001121-2030122333022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201332010021321-2301232103333102-1332120320320012-0120000201213233-0131111021202100-0011201033321312-3321023303020210-2000030322002323"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — default_gateway / 010311203200 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-1233013322333232-2112311321223331-3031012220311123-0112310130322112-3321122302322102-0232013120332230-1000031221233120-0302020212202103"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0323233223112222-2313201010001312-2103100010210003-0222133123213221-1123323331311223-0322332232103002-0312231230200103-0203122331122102"></a>

## Direct properties — default_gateway / 010311203200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013220113101112-0120330133123102-0101311211131331-1121311202232111-2010232033332202-1113201102003131-3132302320332333-0200032313112100"></a>

## Next pages — default_gateway / 010311203200 / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301212020033313-3111221303311012-3102313023023033-1321033113320101-3200211321201233-3210032023133312-1010311003312020-2102210103231130"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — node_interface / 023220030322 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-1031010033021031-0231302212001012-1122003321233222-1333232322120000-1013301233003222-0000123122122131-1223030001122023-1011201012333301"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2112310101001002-1131300220211033-0300310303030232-3000013113330232-3120121210031231-0220231222213211-3212033321233323-1213121200232321"></a>

## Direct properties — node_interface / 023220030322 / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-0313111031203300-1002123202131132-1030022011330220-3221102311212012-2322322320223303-3103033130011201-2331130323131010-2132322013130111): complete subsection reference.

<a id="canonical-1120231330001211-0313223130232331-0201111130021003-1100203312101203-1031201210330113-1323030223100300-0332100230022000-2122321032300212"></a>

## Next pages — node_interface / 023220030322 / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-0313111031203300-1002123202131132-1030022011330220-3221102311212012-2322322320223303-3103033130011201-2331130323131010-2132322013130111)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0313111031203300-1002123202131132-1030022011330220-3221102311212012-2322322320223303-3103033130011201-2331130323131010-2132322013130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013203021133012-3322320000031033-0102100131131230-2303131321122322-3031211121231021-0002202300231302-1123102221011321-1233103000022132"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — list / 023013103021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-2330013130122333-2123333021221301-2223002333201013-3223301111030231-2332212321202320-0030021333300000-1230102310102300-1201233023321110"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-2232002202201300-2310312213330001-3312023201023310-1231330333230123-0311312312003120-2012323013103110-1300212222030223-2302201202323111"></a>

## Direct properties — list / 023013103021 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-3231310210313212-3112023001002213-1020003330120302-1101003122021231-1233031233210220-1010001111212300-1322130003312030-0030222032103011): complete subsection reference.

<a id="canonical-2301203302200231-3003021103021121-2002330010112213-3302333220200021-0303102310301333-0100200301131231-2210121322202122-2303032313121300"></a>

<a id="canonical-2031301002033122-3323131210111222-3010312102333000-0003023131203022-2123133000301033-2320221321003102-3113103311323202-0131030103330010"></a>

## node property — list / 023013103021 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232321121233131-1221231121333033-0212233300021313-0321303332303132-3211323210100020-1123100322210231-0302112220131122-0302113012200120"></a>

## Next pages — list / 023013103021 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-3231310210313212-3112023001002213-1020003330120302-1101003122021231-1233031233210220-1010001111212300-1322130003312030-0030222032103011)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3231310210313212-3112023001002213-1020003330120302-1101003122021231-1233031233210220-1010001111212300-1322130003312030-0030222032103011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112312013121121-0322213113011101-2201112212231111-1132323311011220-1201230323303031-2020133012323333-0213011203332330-3102301333211131"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 312211020111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0302302102013123-0012103201103131-2212032303310113-3323232303311100-0302001301112101-0031310111103302-3300310300201110-2331011022200100)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2203120311210111-1200212231323300-2200220310133012-2312031131021320-1220131023001032-0233112120322121-0232310211213200-3312102103002121)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3133220021121000-3202233113233330-1033132300200101-0031100121013022-0311001213030200-1123013123232222-2130321322310323-2010211303123222)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-0313111031203300-1002123202131132-1030022011330220-3221102311212012-2322322320223303-3103033130011201-2331130323131010-2132322013130111)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-2331033110302220-3300032333220302-1331031000312320-1331112231023230-2230321030230200-0311201011221231-3133001332321212-2322102211132302"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1301101023103022-2120012002131222-1133002001100132-1002232002103322-1112133003201230-2232101033322312-2200101001312301-0312203320303102"></a>

## Direct properties — interface / 312211020111 / 3

<a id="canonical-0100313212123300-0010121120311223-0222013322130102-3111313223201201-3033302110002102-0133113213223021-1010110123132133-1213102021113220"></a>

<a id="canonical-1033101231103023-2021103011101012-0230033033122011-1312121031003132-2302201230000223-1132212310233313-3030202322030023-1223112021112101"></a>

## kind property — interface / 312211020111 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3011002232001130-3013110020123321-2232302200201213-0230321302303202-0012030111000303-2003211021031002-0233030211313323-0021313201321113"></a>

<a id="canonical-2200000333121301-3323122003031001-0133112032123120-0231322221000303-3031201212130232-1033101111302200-3213232033023202-1121222133031033"></a>

## name property — interface / 312211020111 / 5

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3001332020231303-0003100130323132-2202131103023213-1101310333312023-3323010021020011-2113303311203333-2020022122132002-1110033120013022"></a>

<a id="canonical-1123031002330003-1113220113122100-2120301120331302-0302302003102212-1101213223103232-0312213212301233-1323132121012213-1130011001011131"></a>

## namespace property — interface / 312211020111 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1212320111013001-2102303330301003-0020232321220133-0313323032132102-0201033101202121-2012211133232331-2013022230201203-3123332210113120"></a>

<a id="canonical-3123323010133131-1123032132311300-0301220112203101-0213322120213220-1112313001220231-3333022330301302-2201101230301301-1002133112330023"></a>

## tenant property — interface / 312211020111 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3033111021020020-1300302223301002-0110303021021001-2230322201303111-2112233210300031-0021302211220332-2302331110110321-3011302031021321"></a>

<a id="canonical-0201002222331320-2120021210302323-2333023220031102-1112320230231033-0110001023233203-2023003322321202-2210302001122210-1101222301320312"></a>

## uid property — interface / 312211020111 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2301003331030113-1103323132230311-3311033031312103-2032321211321313-0012310221332021-2000101320233323-2021220312322332-3311123021211200"></a>

## Next pages — interface / 312211020111 / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-0313111031203300-1002123202131132-1030022011330220-3221102311212012-2322322320223303-3103033130011201-2331130323131010-2132322013130111)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311011003131312-3321312222130232-3113131021220121-1100032021323102-1323113133301131-3033322112223201-0003202330112011-3220221332313301"></a>

## custom_network_config.sli_config.static_v6_routes — static_v6_routes / 032230211311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-1011312002112110-0130130332021012-3113221003200012-2131203023101232-2210132233201133-1032031320002021-3300313303231333-2232221331100311"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-0020010013200211-0122112003313021-2322231111301033-1222011230130012-1130300010310311-0022123221002133-3002000300331023-3121323113220323"></a>

## Direct properties — static_v6_routes / 032230211311 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000): complete subsection reference.

<a id="canonical-0320201101110202-0013100332133333-3222302103030120-0333221012031310-3331101301033210-1031021313321101-3120232110001321-1203102013222310"></a>

## Next pages — static_v6_routes / 032230211311 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102300313321010-2020211231010011-1102101323323021-3132021002112210-3122310233210201-1320010030122213-1010203030233332-2102013213023023"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — static_routes / 310011202012 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-1133333030101312-1133202100031013-2003012022130030-2220202020332211-2213111020011011-3231323001100011-0102333201122221-0102032120131301"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-2312001133303021-0032122233210223-2033333311030110-2030323120201223-0123310321331023-1110120331133313-0011300000313013-2001201303032200"></a>

## Direct properties — static_routes / 310011202012 / 3

<a id="canonical-2210100103121300-1233011331301210-0220132330232332-0102312002233213-1000030201300120-1200120112231020-1011310102002103-2123233012210332"></a>

<a id="canonical-3030302131212000-3300020002321012-3201133131331003-0111220313331232-1031113332312222-3030220011003111-1221333320312320-2302203222000313"></a>

## attrs property — static_routes / 310011202012 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-0221110032131033-2330122332011202-2132200223211213-3112213113112202-0003101011301303-3132301222231300-0321321132231012-3331121333233222): complete subsection reference.

<a id="canonical-2101320010231020-3201210202133222-2120212003311000-1023102113300320-2232010003313010-2103021111213013-3233103220132202-2331003121323213"></a>

<a id="canonical-0231003220133203-1232023333102120-2210000030030213-2112103320022332-2212002110310231-1210020232313303-3300013221311031-0223331200131132"></a>

## ip_address property — static_routes / 310011202012 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3213010123211112-1103332010010020-2132013312100103-2301322130213222-3021101212012321-0012130210110101-3121020112120002-1333001002012202"></a>

<a id="canonical-1023120310223112-3003112212113220-2032002222003021-3020112023232323-0333113113030332-3211120232022221-0011300113212112-0111221003122102"></a>

## ip_prefixes property — static_routes / 310011202012 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012): complete subsection reference.

<a id="canonical-0323323113131010-1111232100110212-3011003123022020-0032031012013101-3333331313232313-1330313222121211-0320213130101102-3013332030022203"></a>

## Next pages — static_routes / 310011202012 / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-0221110032131033-2330122332011202-2132200223211213-3112213113112202-0003101011301303-3132301222231300-0321321132231012-3331121333233222)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0221110032131033-2330122332011202-2132200223211213-3112213113112202-0003101011301303-3132301222231300-0321321132231012-3331121333233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213013201112320-1300003032102111-2123120312020012-1323302221202010-0330231030220311-1123321213322233-1213102332331120-0331100201111202"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 010113100031 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-2133331301002003-2301321203010013-2123131002332302-0320001202323133-0030120203020302-2000330232232103-0122103200012303-0231202123303120"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0121321301201220-0003302222330003-0211312231130101-0111211231310030-2222110230330103-0012220312322232-0223113123022021-1231223013211202"></a>

## Direct properties — default_gateway / 010113100031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111232111021311-0312023030022323-2313311200322012-1022320310320002-2022230103001122-3102001301220112-1010102321020110-0231303022130030"></a>

## Next pages — default_gateway / 010113100031 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023000110123001-3232101120022103-1111112001202320-1121211213113202-0033031103030110-1210232221121210-3011111111220202-0122132132332231"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 210113011120 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-2103002313232112-1313131031213213-2201302300202320-1112001303020233-1310021233230113-2110332232323231-0013210133333232-2312210120200020"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1020303222210203-0332211133003223-1000030112110310-3033002010210211-0012230223312001-0122121100321233-1313221323332301-0113213200312322"></a>

## Direct properties — node_interface / 210113011120 / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-2123311213031103-3300132313101131-2230102102022111-3312022212232200-3322020220031313-0030221001220323-1123013000310300-3202033333133223): complete subsection reference.

<a id="canonical-3201202010130030-2213331303312112-3103322221023113-0122223202022000-1032231123330311-3012230320001122-0310301033001022-0022111100312232"></a>

## Next pages — node_interface / 210113011120 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-2123311213031103-3300132313101131-2230102102022111-3312022212232200-3322020220031313-0030221001220323-1123013000310300-3202033333133223)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2123311213031103-3300132313101131-2230102102022111-3312022212232200-3322020220031313-0030221001220323-1123013000310300-3202033333133223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232112201220011-3210230033233131-2311232220122033-2111110123002310-1331311200101033-3200202103210113-1033312101201221-0121230113213023"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — list / 322121103321 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-3232211231021112-3200031110201232-1031222232330332-3231201023100110-2133110013002333-3221210102102031-3201010203130221-1023001312122112"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-1320020002121231-2200222232323031-2132110320201323-3030133133120031-1033120321200113-3010211113112012-2033102301032322-1221300012023023"></a>

## Direct properties — list / 322121103321 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-1122330030112322-0122112222330111-0331100132311020-1023211121302312-2222033010123310-3221211300112332-2102110332200211-1021121102311220): complete subsection reference.

<a id="canonical-3220202121010322-3322310030322221-0022301301232310-2112220202231323-2123210201100001-1201133233132303-0020102311322030-0222223122201103"></a>

<a id="canonical-0333110233323113-3333210120102033-3033021333111333-0123103231131021-2313203220212223-2111313222333103-1311230220100101-3302220203202231"></a>

## node property — list / 322121103321 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1230020200123212-3222030012020310-0203323130120320-1002103231201122-2000121110112233-3010011102312320-0132231130321010-0020000201221312"></a>

## Next pages — list / 322121103321 / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-1122330030112322-0122112222330111-0331100132311020-1023211121302312-2222033010123310-3221211300112332-2102110332200211-1021121102311220)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1122330030112322-0122112222330111-0331100132311020-1023211121302312-2222033010123310-3221211300112332-2102110332200211-1021121102311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233233012312030-1122232120222032-1112301031022232-2011302131011221-2013022033203330-2130200121121303-1321313320033331-1120331132002211"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 222000023310 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3202131112032311-3302311331031111-1222020130102101-0133123232012133-2111120301100023-2200001323001020-2010312103020302-3320313120011221)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0210310121121310-0210002200020232-3310032100003133-2203103030210022-2130021302323313-0330231232133011-0010301321333022-3000112033212000)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-3030111232132023-1131002233233232-0120210120122331-1210111302313030-3011303113233002-3200323200330132-0110033113322111-1303300130010012)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-2123311213031103-3300132313101131-2230102102022111-3312022212232200-3322020220031313-0030221001220323-1123013000310300-3202033333133223)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-1203110011313231-0203121232121101-1010021310130330-1323130032132310-2110202131201231-2322013322133012-0233203020331101-1023231001113102"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2111231032212211-3111312312322202-2030121331222320-1131020133211300-3110111320213302-0112030011032010-2212331312102302-3112020223321011"></a>

## Direct properties — interface / 222000023310 / 3

<a id="canonical-3220330230031000-3111131023200310-1321210233012010-3102333313111212-1132301320023313-1033302213100022-3332201121312223-2033313100003331"></a>

<a id="canonical-1012030102210110-3232221032110132-2033131331332123-1300321312222223-3210231021112031-2333133321313203-1113101123232232-3330121322322203"></a>

## kind property — interface / 222000023310 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1120323321313012-1311103000211032-0303322201133000-0213030113322300-1203022001202203-2230201211323221-0022031213100102-2130322120000213"></a>

<a id="canonical-0321303211322002-0102112110001323-2123112101311121-3312223130112121-0023332332011032-3102223100210021-2313032231032033-1103201312111203"></a>

## name property — interface / 222000023310 / 5

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311303010100123-0123003333313112-0210212330212202-2103001131132031-2130013322023223-1232323222123203-0133313031103021-0222002302131201"></a>

<a id="canonical-2212212023020212-1300112332322012-0133122103302221-2113232333311302-2321223010110322-3002031213101311-0133331222123201-1122302120211231"></a>

## namespace property — interface / 222000023310 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0130131302303003-1302320211101211-2213233201031321-0202100102123033-0213010233232130-2130021003311113-3021302202201030-2021033202333322"></a>

<a id="canonical-3213013113331030-2132230310301012-2230211312012332-2103303010310333-3210132023203112-3203022020223000-1020320323123203-0322033301011200"></a>

## tenant property — interface / 222000023310 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0221302202122021-3130231020301023-0203222120030021-3102013131332113-0133002203232232-3020103022221330-3302331303111113-1123330201332312"></a>

<a id="canonical-0212100131323322-0300223131113221-1302110100333030-0020322012131113-2202332011021120-0312133120012233-3222101322100212-3320311032202211"></a>

## uid property — interface / 222000023310 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1021002301122323-3233102211203133-0102101330333112-1212132223111102-1020312302101111-1132210132020031-2010023002012012-2212333223003311"></a>

## Next pages — interface / 222000023310 / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-2123311213031103-3300132313101131-2230102102022111-3312022212232200-3322020220031313-0030221001220323-1123013000310300-3202033333133223)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223300111021130-2102121223003113-0232011113223123-2010001032011011-2322323100310332-1031010012013010-3203233103102213-2122211300331331"></a>

## custom_network_config.slo_config — slo_config / 210101101210 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.slo_config

<a id="canonical-3203230011213303-1200213003033022-0312002200201320-1301113021001202-1202030213233131-2103313030131111-3200310231211010-3120233132212000"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

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

<a id="canonical-0213133113021203-3012100301010323-0110122120121131-1023023323201210-2100223222121002-1122312013133120-1023012320130020-1312021313332100"></a>

## Direct properties — slo_config / 210101101210 / 3

- [dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-0021013231212111-2323011010303030-0220030222310311-1233300002101310-2010301121112103-3331030220323313-1011320013212013-0322020320300130): complete subsection reference.

<a id="canonical-2210302203212202-1222110322010001-0110013321002222-3132021133030301-0211123001113121-1321111202031032-2330302301032022-0021131100102133"></a>

<a id="canonical-3033001023103311-1203220332332211-2101233010033202-3112131322332103-1200223301312210-0012021332312212-1233012301123122-0211123330010222"></a>

## labels property — slo_config / 210101101210 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-2221003221331011-3320132222300020-0100021300131332-1111300003000103-3311011230111033-0301030300132221-1233103211203303-3313211330302321"></a>

<a id="canonical-3220033311013123-0320000130231122-3223211003202130-3211132130002321-3133300133102213-3030111000210011-2332231320302203-2110102223312010"></a>

## nameserver property — slo_config / 210101101210 / 5

Type: `"string"`. Computed.

Optional DNS V4 server IP to be used for name resolution.

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

- [no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-0331220313031123-1130231111010331-3222321211223331-3020221030123223-2122103130312222-0001021030322313-2121310322131203-2331222222130222): complete subsection reference.

- [no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1311123201303212-0213023032011103-3323032320012231-2332020123113332-2323113320013330-3322022113122133-1313022103211202-0011223130232322): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0332233001113101-1101130031112011-0130120231011033-0011000023120220-0232121131121110-3010321210232103-1331010032220200-0323210030211120): complete subsection reference.

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331): complete subsection reference.

<a id="canonical-3210021312300201-2310101333123103-2323233222221030-3301211022311111-0001133202023331-2320013031003330-1031231000323130-0221301111212332"></a>

<a id="canonical-0301213013322220-1123022011323021-0330221303311221-3222320011133212-2311331312201001-1202331031332030-1030100002122112-2312322033012333"></a>

## vip property — slo_config / 210101101210 / 6

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-0210133011220010-2102211010221320-2232333312002102-0112100003023220-2323102001102211-0131233013010010-3101202310223323-0133111321211311"></a>

## Next pages — slo_config / 210101101210 / 7

- [custom_network_config.slo_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-0021013231212111-2323011010303030-0220030222310311-1233300002101310-2010301121112103-3331030220323313-1011320013212013-0322020320300130)
- [custom_network_config.slo_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-0331220313031123-1130231111010331-3222321211223331-3020221030123223-2122103130312222-0001021030322313-2121310322131203-2331222222130222)
- [custom_network_config.slo_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1311123201303212-0213023032011103-3323032320012231-2332020123113332-2323113320013330-3322022113122133-1313022103211202-0011223130232322)
- [custom_network_config.slo_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0332233001113101-1101130031112011-0130120231011033-0011000023120220-0232121131121110-3010321210232103-1331010032220200-0323210030211120)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0021013231212111-2323011010303030-0220030222310311-1233300002101310-2010301121112103-3331030220323313-1011320013212013-0322020320300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330023320003213-3331311101232003-2213032102131112-1003320002222233-1310232211331013-0302332211120100-2310133322033211-0133332320032013"></a>

## custom_network_config.slo_config.dc_cluster_group — dc_cluster_group / 021113231333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-2112133333311322-0020211102020101-1003010110102313-0231323232131011-3322131312301331-3013331021200232-0100231230013120-1111221211313001"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2300232231122001-1120032003322101-1213331000113302-3010310232201221-2122110221212022-3121221333002223-2023132032020301-0122023201221033"></a>

## Direct properties — dc_cluster_group / 021113231333 / 3

<a id="canonical-2022112201232222-3131100111132130-1300211012000011-0232232112022201-0020210130131121-0203100012210323-2331012103332003-3002002112110303"></a>

<a id="canonical-3331212023003202-3011202312211233-0333230112030312-0100302331333101-1101323300111312-1323033130313132-3130323103001331-2223232002013331"></a>

## name property — dc_cluster_group / 021113231333 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2032302231220201-0301331131221111-2333010221313003-1202331113132103-2333120313211012-3020313003310230-2210233022113330-0321031232211112"></a>

<a id="canonical-0011312323312010-2230303200131112-0011111020120003-1033322320213003-2002213231121021-0213013321331013-0321100012132303-1303112310012331"></a>

## namespace property — dc_cluster_group / 021113231333 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2001011230300211-1222200003013122-0303100003322310-1012200000112102-2011032113301320-2300200031010031-0121011311300221-2110111202032301"></a>

<a id="canonical-0233012200303000-1110130103100121-1233301231103002-3131101202032310-1210202323011231-0323133103000013-3130013103121303-0331121033113331"></a>

## tenant property — dc_cluster_group / 021113231333 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2132311131031111-3300023212110312-3211311322002322-1132223020103102-2222332223300123-1330012020232200-3303331232321231-3001110112333212"></a>

## Next pages — dc_cluster_group / 021113231333 / 7

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0331220313031123-1130231111010331-3222321211223331-3020221030123223-2122103130312222-0001021030322313-2121310322131203-2331222222130222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122001302333312-1203123233121223-3131213111030313-1100332221210222-2133130233011010-3223321312002323-3302333112111221-0030303102201322"></a>

## custom_network_config.slo_config.no_dc_cluster_group — no_dc_cluster_group / 320233122101 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-0132312211030001-1331133323030321-1112220112320311-1210112133321021-0330201101110210-1332130300103212-3320233200221320-0201230310022132"></a>

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

<a id="canonical-2223322100313322-0330210000103021-2111211020033220-2312223032032330-3020122102011221-3131103233002033-2203213223302100-0302333112122212"></a>

## Direct properties — no_dc_cluster_group / 320233122101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312110013013331-3200123121101330-0233001300102030-1020211301001132-1310301233203111-2210122103312032-1310100212310212-1300110021110210"></a>

## Next pages — no_dc_cluster_group / 320233122101 / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1311123201303212-0213023032011103-3323032320012231-2332020123113332-2323113320013330-3322022113122133-1313022103211202-0011223130232322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210311320032021-3212300213232230-0122301001202311-3013122111231222-2102223013310122-0120021101211013-3011103223130032-2220121010002001"></a>

## custom_network_config.slo_config.no_static_routes — no_static_routes / 222130300310 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-3332113130332130-3032332301301200-1102000113322310-3022202300211330-1111221101331320-0232111132110232-3330020201033233-0002321232303302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3103133022302030-2203223313030313-1200023210020131-1101312303100130-3120323120102333-0333033223220223-0321323033310230-0313310310233121"></a>

## Direct properties — no_static_routes / 222130300310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222131033303332-2231313220223021-2231333233112003-2231003013213100-0330202010101002-0230232132332012-0201203211112211-1010022110301223"></a>

## Next pages — no_static_routes / 222130300310 / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0332233001113101-1101130031112011-0130120231011033-0011000023120220-0232121131121110-3010321210232103-1331010032220200-0323210030211120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233030231020001-0110132202321201-0303132300200230-2001112221021202-2220001302203100-3211203121123101-2300232202102120-1310323313202130"></a>

## custom_network_config.slo_config.no_v6_static_routes — no_v6_static_routes / 011113122130 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.no_v6_static_routes

<a id="canonical-1221000212123021-2130332100220221-2303233201303231-2203030023213111-1211002320231020-0022111112033123-0313000301030020-0103202112112003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2030311231303222-0010221332031212-1301111303213110-2223113111233323-3200103312223103-0131211110120203-0221223320123000-3230223303013221"></a>

## Direct properties — no_v6_static_routes / 011113122130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002112100301202-0021011132122202-0122322032130313-0000010223012003-3232233030131320-1211312203323321-0332101223033221-2101203001220201"></a>

## Next pages — no_v6_static_routes / 011113122130 / 4

- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330011030023212-2332023210301200-0322113300020203-3330103320223130-1033331020311322-1013031313222310-3232110333301033-1011312322103122"></a>

## custom_network_config.slo_config.static_routes — static_routes / 130312232300 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.static_routes

<a id="canonical-2210011211232111-1333210323302323-3100022003031013-0322323220313130-1231011202202010-2022330320223121-0131213000030300-3013220003103020"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-1100200010212021-3012311303112112-0103330222330230-3002021303333310-2203102202102332-3122302132301113-1203220210113333-1222321221123301"></a>

## Direct properties — static_routes / 130312232300 / 3

- [static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220): complete subsection reference.

<a id="canonical-2301313222001321-0320102031203021-1120103020323123-3213322013213320-0130333021020120-1200332331221103-3133022220213310-2010110131110202"></a>

## Next pages — static_routes / 130312232300 / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230330311312333-3211022032123213-0123202121033200-0132121133113133-0321113111013322-0002110301030213-1230012110112123-3012330022030203"></a>

## custom_network_config.slo_config.static_routes.static_routes — static_routes / 131100232110 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-3321332102212130-1332122021312020-3120012232332212-0312311130201202-0000102122010002-0331100102013201-3121021012302232-2333213003121133"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-0301001222101302-3202201012310312-1322101323123132-2132103212310102-2013213323312121-3213331131302111-1000020212133023-2100113111030002"></a>

## Direct properties — static_routes / 131100232110 / 3

<a id="canonical-1132133002032133-2211333100000321-0031231232133213-0103311232322132-2303003131010130-2112010331123313-0332131222130021-2231323121300011"></a>

<a id="canonical-1222322221301303-0300310010321102-1203012333322221-1020103032311020-3212112220213330-1111201302212033-0111233013302302-0200003302302333"></a>

## attrs property — static_routes / 131100232110 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-3313121110210001-1013311223202132-3103333332102200-0020111330013032-1222120020223300-2030033012230122-3101323113210201-2323221203023100): complete subsection reference.

<a id="canonical-3133302001003230-2221003220131131-1101031122011212-3311022320023113-1322103203013130-0112020133301221-3211331332023200-1313200100010013"></a>

<a id="canonical-0330113100310231-0300000013310033-3330222331231213-1321200000132013-2322210312222113-1121113313233032-3330320111332333-3312111030123331"></a>

## ip_address property — static_routes / 131100232110 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0001112002313223-2213203203102300-1301333233333302-3123102003111202-3001130230013133-1130003303011032-2313221230323112-2230313013301302"></a>

<a id="canonical-0230220301210002-0030112300011101-2023101001001131-1032321011331332-2303022030232122-3122210002031002-3310130132332111-1331333210321023"></a>

## ip_prefixes property — static_routes / 131100232110 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323): complete subsection reference.

<a id="canonical-0310000321023000-1133112223211110-3032330020330002-3320130031301022-0120111322201121-2202311220013131-0210220032103301-2203303223113212"></a>

## Next pages — static_routes / 131100232110 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-3313121110210001-1013311223202132-3103333332102200-0020111330013032-1222120020223300-2030033012230122-3101323113210201-2323221203023100)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3313121110210001-1013311223202132-3103333332102200-0020111330013032-1222120020223300-2030033012230122-3101323113210201-2323221203023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213130222033322-2103113013133101-0131101121103331-0010330231002122-0323030102223030-3102210000023113-1212332211311231-3020100132232302"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — default_gateway / 303203213133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-3321101231101020-3010010021322020-1121322101230021-0300313213110020-1322233331131210-1023223311111013-1013103112302210-1013022131113111"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0231211133232332-0030313011131311-2210132222310001-1212312032023323-3031311023320031-2332233301303231-2133321033302122-0222233123303231"></a>

## Direct properties — default_gateway / 303203213133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311213013021300-3031330230302123-0010103103203232-3232100131032021-1331333311021033-1031011122131102-2130011112332221-1010211303232310"></a>

## Next pages — default_gateway / 303203213133 / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103321022102222-1313202222302213-1313211213012312-3320311001012021-1132133222231103-1201032022101101-1012201000222123-0001313321312202"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — node_interface / 110132102033 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-0311223002220100-1223212022301101-3110321321330320-1011010200122323-2200321031332332-2231133221031212-3123133010301321-2010000322332001"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1311233322032230-3122033021101022-0031211301210022-3011311122101210-2003232213123333-3202021202222222-2003331133322233-3003022211202110"></a>

## Direct properties — node_interface / 110132102033 / 3

- [list](data-sources--securemesh_site--reference--group-003.md#canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133): complete subsection reference.

<a id="canonical-0021120011021133-2301211112232231-1320003211132223-3000023221022000-3223001200211300-3032132102030130-0032200230111331-1020120102011230"></a>

## Next pages — node_interface / 110132102033 / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003001320203333-0200130021331022-0003312031003332-2221133302312021-3121110202330233-2312110320011113-0010021331310231-0311000210311010"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — list / 032001300000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-2220023102321031-2121321211111311-0130001332231202-0121333013311031-0011120300221320-2032013021221232-2013213332033312-2331300112000112"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-1313322030110110-0130023010101323-0023013003233210-0313103323301231-1132032101002313-1030010210303112-2213313322303011-1130322221101212"></a>

## Direct properties — list / 032001300000 / 3

- [interface](data-sources--securemesh_site--reference--group-003.md#canonical-0103030001310321-1332113332321113-1032123020022013-2313202333113112-0112303002231022-1232113330310331-3033123313200020-1002002132002220): complete subsection reference.

<a id="canonical-0323313301310110-1223112101312131-1231212202312231-2003220321233223-0222033231003020-3012003321210113-0100130203212200-1103132102222210"></a>

<a id="canonical-2312202020111300-0001222213010331-1022010113310212-0301003121032333-2023212300232032-3103322200011103-0033132203001001-0222111213023010"></a>

## node property — list / 032001300000 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3312012311331211-0232301011000030-1003231212110203-0311101202012320-0331323032000001-2210320123331111-0220300200130001-2220021111333332"></a>

## Next pages — list / 032001300000 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-0103030001310321-1332113332321113-1032123020022013-2313202333113112-0112303002231022-1232113330310331-3033123313200020-1002002132002220)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0103030001310321-1332113332321113-1032123020022013-2313202333113112-0112303002231022-1232113330310331-3033123313200020-1002002132002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132120103123332-2222031311311203-2221321301301120-3220101312220131-1112012311022003-3220301300303031-3001111220113333-1211101023102000"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 203003033320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-1211212133113102-0023112010301203-1021010111002333-2313202322132321-3331302133120101-2122022323032132-3010210120101020-0323230102331123"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1121100013303132-0311330031230003-3002223102132222-2112102201103012-2231123022102100-1000333323101222-0311113103233020-2322020022000312"></a>

## Direct properties — interface / 203003033320 / 3

<a id="canonical-2300102000201103-3310311302122211-1120301132300312-0022110231210002-0000131203000321-0222311312010110-0323011110120200-2111003022210001"></a>

<a id="canonical-3123323203031322-3212302212021201-3002203030003003-3100101020023213-1003200231130212-0110202222030200-1123333023311000-2301133023011212"></a>

## kind property — interface / 203003033320 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0330201010332002-2121131101233310-3323301103310112-3002320033313210-0312230211000201-2003231012103231-1133032123102033-2201000102031201"></a>
