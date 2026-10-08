---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1031313312122303-0000002311112331-0333031212322323-1030120011301132-1022033113332231-2010313303003012-0322210111131100-3320111032321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3230331011013231-0033311332220330-1113210312011120-3331333332332021-0122323223022230-2332033111310203-0213320023132333-3021312130012211)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1130121121013202-3031302020213203-1211300030021203-1002011130312303-0032330220102200-1212322002301013-3111023020110021-2130011213032232"></a>

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

<a id="canonical-1223220231000220-3320223121321112-2100102003122100-3300130312303313-2101220021023201-0311001220113220-3222303201011313-2032000310221132"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2313000000311032-2003301002321210-0101231331323222-0203013131211200-0232102120223303-0221310211123231-2000300101130120-3320201010331020"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3231012100101121-2133123332231013-0322112221222310-3132001231320222-2332002231111002-2033220210310331-0323010222001011-3121222311233303): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0321000032132330-1003311130200133-3331133232310202-3101111013020021-1131230202232133-0112303203320321-0232122322223330-1202030313300303): complete subsection reference.

<a id="canonical-3231012100101121-2133123332231013-0322112221222310-3132001231320222-2332002231111002-2033220210310331-0323010222001011-3121222311233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3230331011013231-0033311332220330-1113210312011120-3331333332332021-0122323223022230-2332033111310203-0213320023132333-3021312130012211)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1031313312122303-0000002311112331-0333031212322323-1030120011301132-1022033113332231-2010313303003012-0322210111131100-3320111032321300)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3223321110120020-3301003201031013-1130130320032313-3312120031321001-1031112210023232-2031030232002300-1122130001102021-3133031031113121"></a>

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

<a id="canonical-0321000032132330-1003311130200133-3331133232310202-3101111013020021-1131230202232133-0112303203320321-0232122322223330-1202030313300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3230331011013231-0033311332220330-1113210312011120-3331333332332021-0122323223022230-2332033111310203-0213320023132333-3021312130012211)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1031313312122303-0000002311112331-0333031212322323-1030120011301132-1022033113332231-2010313303003012-0322210111131100-3320111032321300)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2121321331133331-0002210232331211-3001301122001130-1123032221133000-1000210033130020-2233212322100011-0203311002033310-1022302012122012"></a>

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

<a id="canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0123210121102211-1211002221323312-1001021022203313-3101030110210222-3112102323010002-1021103321023201-0123233132232033-3210011203131023"></a>

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

<a id="canonical-1100122132312332-1212003101013301-3200213223312023-1303121131000112-3322302312300320-1113110000003001-2310022130032221-3011333230210210"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1312010100213310-3110110202310111-2331300301002323-1022013021201331-1013131320113233-0311032113232011-1313212320220323-2223022332302223): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2112310202123011-1003321133212302-3011101100112303-1312002031310003-0210313122300303-1200012311330021-2122323113103003-0201013023310130): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1232330033201122-0122313101321222-0321202311122220-2212331122310022-0000002001230032-3031302331301310-3200322321200202-3100202230032033): complete subsection reference.

<a id="canonical-1202233123221133-3102331000330233-1120112222121122-3123323101120233-1313333311231001-1212112001011033-2322020210301130-0303011202030222"></a>

<a id="canonical-0333123003201033-3123123131121122-2011120012011211-1123030112001330-2120112210202030-3010100213031330-3120301300210323-2130332132231000"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2010003020232130-0032132120132320-1121011223203332-1300210021131332-2130213021110012-3303220020030310-0103030213031321-2330110121313220): complete subsection reference.

<a id="canonical-1312010100213310-3110110202310111-2331300301002323-1022013021201331-1013131320113233-0311032113232011-1313212320220323-2223022332302223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2011213101332230-1113313031123212-3323023023111312-3120233010103001-2120032322333303-3100031102231111-1210003033220323-0110200330302031"></a>

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

<a id="canonical-2112310202123011-1003321133212302-3011101100112303-1312002031310003-0210313122300303-1200012311330021-2122323113103003-0201013023310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2012122213303212-2303302202213210-1131112200311301-0013123231302112-2303303012020211-2020122031202010-1002031300030030-0010123201203031"></a>

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

<a id="canonical-1232330033201122-0122313101321222-0321202311122220-2212331122310022-0000002001230032-3031302331301310-3200322321200202-3100202230032033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0332220123022010-0321321220001233-0110320303120321-3010302121323133-2020231310100131-3120122312310221-0010231121332330-1133122233000300"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2302002223203313-2222210331112103-3103333103213321-0301111321333010-2330022022003001-2002320121232220-3003313022133103-0202112000213110"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-3233001022201302-3003230203132322-0120312011230130-2332203112212220-2130221323123033-3013210323110010-3310022330122020-2131320100031310"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0110220031310323-3113012120011122-1122033132022221-3020333220113111-0203111221301301-0300031311032312-3033113201011120-2210021311233232"></a>

<a id="canonical-0101210030313211-3113223312320333-0120223322202303-3223133132312003-0332201222020110-2311100131113210-2101303210032323-1222301130311220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1212031012022113-1201030332311032-3313021322300031-1220320012031300-0113133010201113-2122311133333310-2113321112032301-3210103030012012): complete subsection reference.

<a id="canonical-1212031012022113-1201030332311032-3313021322300031-1220320012031300-0113133010201113-2122311133333310-2113321112032301-3210103030012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1232330033201122-0122313101321222-0321202311122220-2212331122310022-0000002001230032-3031302331301310-3200322321200202-3100202230032033)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0023310211102130-2221100013301330-0131301332032000-2010231203013132-3313130002332002-2030013221220312-1311033202222313-0213332301321230"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1102023020200003-1111123331120030-3113112321032022-0201211210133300-1012020011131021-2300133323213222-2221121313002221-2223111301202221"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-3112013033313031-2321323133233132-3201011231230113-0002021323102113-0132113122010131-3331203330230101-2223011112132021-0333302301003002"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3212022000223030-3003211122033320-1121201133200220-1110203030001203-1303013110230231-3003102131101231-3012203300132233-0322001332001100"></a>

<a id="canonical-0032233222130212-1013202013120301-2223010310303323-1111111301230031-2113202211112110-1311112133013132-1122302320130230-0031012320022222"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2010003020232130-0032132120132320-1121011223203332-1300210021131332-2130213021110012-3303220020030310-0103030213031321-2330110121313220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0321012232003312-1012300003020031-1012020121233332-3030300203302302-3023220021111313-0113321222022020-1233331120111120-0233031212211332)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2312200323020001-3211313300320032-0100312023221311-1011010200020013-0023330313122113-1300133311021303-0220330303100330-0101013231211232)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2113201033202101-0302300121123111-0013013012023211-0130030301120102-1211300013311022-2231302303100032-3303131011210302-0333322031213003"></a>

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

<a id="canonical-1022121310312112-3312122102103230-3312003020020021-0000032010202320-2100332023102011-0031110023332032-3331233023233300-2022321332032012"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-3022132000303033-1112112233230202-1200232122201032-3020323332330301-1123322232033003-2121212030321011-2032123020200111-2032230022123120"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

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

<a id="canonical-0222232312333331-2302301333213120-0310101201131330-3321330213023320-1121131031303333-2303010100202133-1120311233100301-3332120303101310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.monitor

<a id="canonical-1221130000003111-1233300301123002-0313303320201023-0212223010132320-3212202301102132-2132311011200230-1121210101020322-3013301232030310"></a>

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

<a id="canonical-3230210033021102-1020113122103033-0132323330121033-2000102120023001-3112212100310210-2300222210200133-3201202230203033-3120330020100320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3322202230102011-0332101110121303-2012322112211012-2112102122232221-0012200010100322-2103233132010321-3100100323130301-1222133122330012"></a>

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

<a id="canonical-1013302121310111-2320131131212011-0000230233120322-2131110203001220-2322300221210212-1020212203213023-1101132233130122-2203322322321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.network_option

<a id="canonical-1321330032330230-0033130222123221-2010102121022221-1022321120121312-1232312023213022-0000333301212112-0133312233210022-3120130130230023"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

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

<a id="canonical-2103103320132331-0313222322030011-3301111311013222-1330232012002333-1221320101113212-1100001210000323-0303113031002103-1203132332021001"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2330133131012033-1322220322120011-0130320102020333-3020210133330203-3323103013030300-3200111000022331-3332110123011103-1211330230112210): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1221311102112103-2231120133221231-0203301201132312-3200022232031102-2311121331110231-3332302033233312-3113221231301032-3030230310301010): complete subsection reference.

<a id="canonical-2330133131012033-1322220322120011-0130320102020333-3020210133330203-3323103013030300-3200111000022331-3332110123011103-1211330230112210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1013302121310111-2320131131212011-0000230233120322-2131110203001220-2322300221210212-1020212203213023-1101132233130122-2203322322321233)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2233123230023202-3003120223103203-3230010032100123-2210330230133310-0021120121122001-0123303100310001-0022133202323303-2203133030203101"></a>

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

<a id="canonical-1221311102112103-2231120133221231-0203301201132312-3200022232031102-2311121331110231-3332302033233312-3113221231301032-3030230310301010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1013302121310111-2320131131212011-0000230233120322-2131110203001220-2322300221210212-1020212203213023-1101132233130122-2203322322321233)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0211031032223230-0321132003110122-2210023003333123-3211022302110110-0233212323221313-0311232121212210-0213210212002211-3021311012222030"></a>

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

<a id="canonical-2020220322120022-0022203022103101-0113331310113230-1020331032111000-2203213213020200-0023021302132021-3212310100313321-1301332011111301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3032203010302100-1002031211013121-2133333232310022-1221223100130101-1122310001222330-3312002201113013-0220332010033102-1121212310013023"></a>

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

<a id="canonical-1330220003032021-3122103222123230-3210011011131321-3331313010322312-0131211231121201-2330201230320103-0111130010211111-0013303203313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1232332203220133-1031203212123033-1111131200211111-2230121021011331-1220223133212120-0330010302121231-1330211100230000-0220023022013332"></a>

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

<a id="canonical-3331222003001103-3121013100322030-2021302333300010-0113323300102111-1210221331102202-1100202003321030-0200212231022330-1020320330320320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-0333331003011110-2233120011030030-1320130110030312-1213233001330322-2222103132032102-0302322113001212-1322332130011122-3112230022031203"></a>

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

<a id="canonical-2133233221331033-1021332103313132-2210221223320123-1211022011012222-3000321020331002-0130211322131302-2322203220101132-3202131311110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-2212110312121111-3331302322032323-0202113011021033-3022333202320332-1223112200230302-3012013033113233-1323030023020312-2103111002302132"></a>

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

<a id="canonical-3133123130132330-3021102311212120-3111023231321000-1021320210012323-1032211001033331-2003213213222002-0132200310233210-1311312210312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.static_ip

<a id="canonical-1000123032320300-1321301213021201-2322223330122301-0033312211010103-3332023113033100-3323121220103301-3203103102332230-3130031311003222"></a>

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

<a id="canonical-3213000311021312-3100330012330110-3020312022333020-3012003203300032-1013330002303333-2012132001003322-1031303031031323-3310100302203212"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1323300002213122-3121203001330300-0301111103030100-0122020023030332-2022021001033310-0022022030221211-0323111223333220-2022212221012222"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1301200210023210-3010030110330033-1013322203203102-1100001013223323-1200221310223010-0111301233132322-0123132323012121-2000033001231303"></a>

<a id="canonical-1300132102122332-3232102212021321-3333311113303012-3300310031121213-2210322310101211-3121010300313000-0112030213102133-1320333213300212"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3230020031030132-1332132101100032-1200301310321303-1033001201102022-3223121223111030-3030321111001100-1322013122121223-2020132221213313"></a>

<a id="canonical-2213030120110301-3332233132220203-1232130021013201-2123331121011330-3200300231222333-0110130113313212-3013213231110230-0113303232133211"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1031013013203003-0331110020310301-1023330132300331-2131022132331001-3230022032230030-2330302202212012-0130230013030022-1300201020111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0032100120213010-1320313123023032-2231103013103203-1203311203113100-2133331332033013-3203003201021203-0222121211303001-3311001211222302"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

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

<a id="canonical-2230231111200210-0223221000003212-1213110101201110-0233023331133121-2212031133313302-1210002031310210-3211311221131312-1231301122002120"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1112003103133201-1321232203021203-1312323202110221-2011103332021222-1232201200131222-0131110320110331-3102103213313202-3131333112222131): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1232112321030032-3032210320030203-1303112222013231-0201310232120202-0233032211303101-3231123231300002-2021312223030130-3101000310103231): complete subsection reference.

<a id="canonical-1112003103133201-1321232203021203-1312323202110221-2011103332021222-1232201200131222-0131110320110331-3102103213313202-3131333112222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1031013013203003-0331110020310301-1023330132300331-2131022132331001-3230022032230030-2330302202212012-0130230013030022-1300201020111002)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1221231232211330-3300303012322033-3012000013322223-2213020201111202-0023123103230201-2111230233330302-2111123120320200-0231122003122330"></a>

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

<a id="canonical-2121032003131321-0222101232133023-2023210113012111-1203033101212012-1001313321321300-0110201320313312-2233122200332100-2110022212211233"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-2331032101132333-0033303102302002-1103212132010100-2303120333230203-0033203012331311-2221223010330200-3310333220212210-2223022302020002"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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

<a id="canonical-1232112321030032-3032210320030203-1303112222013231-0201310232120202-0233032211303101-3231123231300002-2021312223030130-3101000310103231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1031013013203003-0331110020310301-1023330132300331-2131022132331001-3230022032230030-2330302202212012-0130230013030022-1300201020111002)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3322022100330131-3111001332301202-1303003322332203-3223222320023101-3010310213300310-1023003323220011-0110012031113202-3221333210330132"></a>

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

<a id="canonical-2130021002133333-1231121300310023-3111330221010112-3133131320112201-3323013302013320-0112313002122212-0333022130132130-3211210031002013"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3303122200001301-0203010310201102-2332031001021212-0303132223000000-1102300013210103-1033312221122310-3330013321313021-2311231311232012"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3201321030213031-0101022223323221-3213200331210331-1220330121133022-3103031022100002-1001201023232012-2331002132211103-2032001030101220"></a>

<a id="canonical-2003331301310121-3330213022122323-1202222122132000-2023311320233231-0223011032310022-0303023031000213-2123103202000111-3222323010000303"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2111023230323211-0210210021021130-1131222023133103-0133310200113120-3111122312301222-1331300131200011-0001312123012030-3231302311210011"></a>

<a id="canonical-3333323202002213-3021302222132130-0101212120010133-0221133132212012-1200013032233030-1001133201211300-2012231020213332-0103120022022100"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1200120313200112-0333103032211123-0311231323211303-0001233210331211-1322200320020031-2120113332302220-0232222032210322-3022303111121232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2002020103210111-2231310121323311-3313200320012223-2113130021103131-0100222203301301-1101222120010132-3111112011223321-0002310030101210)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3302322331002102-2002320330032132-1111311313102033-3033021101103230-2122130203101221-1321100111333011-1202321011211332-0021112021323013)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1332120013232103-1222211031233103-1230010331002222-3230033113213210-0100222230112112-3101300222023201-2101302220122113-1013320223211212)
- openshift_virtualization.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0031010120221301-2233010302120332-2102133111201112-1101023113213133-3023133012302031-0310200201010302-3203331021012000-3122032020302110"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

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

<a id="canonical-2201121232330003-1221322132023022-0013030110010132-1012313011321211-1333123131312220-3032230213031011-0113203012033202-0321010232333323"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-2121113232313013-0010220112111211-2133232233220113-2032211231023222-0310311000112013-3111213031131303-0121233322101313-3110032121003010"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0030133020231303-1301011033022300-3322231302212013-3312013101121010-1323221122302230-1233220033101133-3201210001132212-0313212102030123"></a>

<a id="canonical-3231103012022332-2233003303323121-1101210332321120-0200001213210101-0101003233030110-3001231312300111-2323032213020132-3102303012221221"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- openstack

<a id="canonical-0000320003110311-1320021332020301-0130102333312030-1213003310131021-2232010101101321-3332100030200031-1122120011223313-0031332301020013"></a>

Type: `"single"`. Computed.

Openstack Provider Type. Openstack Provider Type.

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

<a id="canonical-3203013022110010-3203233332013320-3033233030013220-2303233201231203-3113223321110113-0231020220123001-2030211300033101-1103123111113301"></a>

### Direct properties for `openstack`

- [not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033): complete subsection reference.

<a id="canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- openstack.not_managed

<a id="canonical-3002303212022300-3231313111011303-1022202201031223-0333300032213003-0020210310020233-2120330312000020-0232123001010313-2021220123230203"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3331233322111133-2032221321231033-2211311223333330-3202002210323101-1130013110211120-1032000313001312-3101003302123032-2233012133011322"></a>

### Direct properties for `openstack.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003): complete subsection reference.

<a id="canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- openstack.not_managed.node_list

<a id="canonical-0131322321220211-3010203232123200-2211132000120320-3203203312203303-0331133320101002-1131021210023310-1232023220330110-1313031321101102"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2010030020033021-1330230222030122-0311103132311223-2121322123232100-0130320132232312-2001212133100300-0022200322323000-0212102201100221"></a>

### Direct properties for `openstack.not_managed.node_list`

<a id="canonical-0230032230131320-2131201012333323-3121010011221111-3213132312223300-2313303313302313-1200202100121213-2221103113301233-1301010221100301"></a>

#### `openstack.not_managed.node_list.hostname` property

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102): complete subsection reference.

<a id="canonical-1313220013012213-2111010331122133-1122320213210012-2330010012121203-2133232130333213-1223121002120233-1232321100021133-1112003130113023"></a>

<a id="canonical-2300033213013000-2033202333100101-1131313023301001-3022210332220330-2120300310322100-1120333220322301-3000333121220320-3110200320311211"></a>

#### `openstack.not_managed.node_list.public_ip` property

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1111201223001130-3100022000100033-1310330030020312-2022021032311302-0131032200313330-0333013100003111-0123012231130322-3113102101023331"></a>

<a id="canonical-1113200231211030-3001130011123113-3021130211012031-3001300302322103-3231131310312022-3200300221011213-3113220321200211-1120010003021210"></a>

#### `openstack.not_managed.node_list.type` property

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- openstack.not_managed.node_list.interface_list

<a id="canonical-0332100122201200-3200212003120210-3112012312003210-2212100021203323-2112022221020001-3233212032301131-3213033321131200-2222210111200213"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3111230222203030-0103032232200303-2301320211030103-3021300110303322-1323102032130132-0321131213132310-2010101133030300-3230231202002122"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3303323332321002-2313023131102203-0030202133133101-0131313021311203-3311311223113211-1221213023012331-2100033133221010-0213121101210312): complete subsection reference.

<a id="canonical-0223100231010030-1133232210303201-3220230003323300-0321213213022213-1102313121313000-0311133110311302-0331110331131312-3012302212023100"></a>

<a id="canonical-1313232021123302-1320313221101000-0311220310222222-0220023311132101-1210031212211223-0311302221302313-3022302223301333-1223330201032003"></a>

#### `openstack.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0033012012013111-1223332331302012-3132032311123303-1322200212211112-2112111010202020-2230003221111033-0320033013103232-3113120330303213): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2033333113321230-0122003033301003-1302133221203002-3112230201223110-0223201223311010-3320113322202022-2310231230003111-0312220200021023): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301): complete subsection reference.

<a id="canonical-1223222132311223-0333102223323030-3033202222132203-0301001233000301-0122211012021011-1112013032333103-2102320011131210-3333032120103333"></a>

<a id="canonical-2332111211320301-2230003110102002-2213313301001031-0303332032100032-0210303023311322-0202321101200200-3303131012310123-1103200333230213"></a>

#### `openstack.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0120312100113111-0002022112300331-0013311230311102-1221021202021310-3212303222102111-3220301122221100-3303302302331030-1111333323212021"></a>

<a id="canonical-0012303111302322-2213133113111233-0332013232020131-1201311201330123-2231011022133111-0211311232030330-0100130103203223-3011213202313022"></a>

#### `openstack.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3120010033120030-3000303213213122-1101023120111330-0300010023121012-0223223220021103-0222220320102113-0100202121322020-0020022123223302"></a>

<a id="canonical-2003212122212110-1321322331101220-3212202111320220-1301311333221023-1220121133303300-2332212231301323-2002210001231033-2332312103031331"></a>

#### `openstack.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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

- [monitor](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0320123220131011-0213321013300220-0002002112223113-1300010102212000-2313202231000223-0202113132301323-0030312202220200-3103222000313122): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0111101012331023-3321222033002313-1201222031230213-1111002200302333-3021112320122202-1310113230030012-2203131302103031-0001322113201103): complete subsection reference.

<a id="canonical-3233132322220112-3111230300011021-3020202132011320-1311131330302321-0011010210221113-2313113221001332-0222230011203133-0132231122112100"></a>

<a id="canonical-0331033302022001-2230212133033122-0223223220321333-2131231320110312-0232332103102222-1232111120132213-0233322100221112-1213033131212332"></a>

#### `openstack.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3211300102311012-2203012111020030-1000321020230310-0121001302230211-0010200231311102-1020221113232013-3220300130001120-3221111221320032"></a>

<a id="canonical-1323130002321333-1300213302030101-1023220230302000-3100001021030201-2133232020103000-2103012213211123-0201022331133100-0030031213130030"></a>

#### `openstack.not_managed.node_list.interface_list.name` property

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1010322322302200-0233212100000030-1100303000000031-0321322221023131-3233322021311021-3302201133120003-0003233320103012-2001021303220221): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0020121230311220-3310233121122123-0023131100022131-2320032311030002-2220323203020033-3103111020121003-0021011110230103-2222131303303102): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0023212030220320-1233000023302112-1320113121333103-1230033002232130-1002202231300230-1123112030021020-0032311010311101-0331230330310003): complete subsection reference.

<a id="canonical-0322010213311132-2312002321331221-1020230000323211-1311000012113121-2222131321333112-2020302200131233-3220332002203323-0330313320011311"></a>

<a id="canonical-2202131002033110-2202132310200203-3110103113331213-2012222031030023-0312202202002303-2001301330030332-0232231213012022-3200010201100331"></a>

#### `openstack.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0211231110302111-1002001101222012-2211003313303323-1323003121111220-3311233231303113-0033131130212230-2002221202303013-1133212201113212): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-3100213310330310-1000030113112223-3113220320331312-3132032333013100-3212112021322220-3300002103200312-2313302121213221-3023121203221202): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0110012232231101-0220321323222030-3020001320121330-2013223003003122-2032123111300121-1100210330022120-1301331022323330-0300133130213111): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1133333032223312-3320220331330233-1222020303202123-2201303332330121-2012211031332331-1230220320332103-3101223001210203-3210321032322031): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2223333332301211-0322331121023331-1123111101312313-0002201031122231-3011021200020110-2021012313230000-1202202030100201-3032331301102310): complete subsection reference.

<a id="canonical-3303323332321002-2313023131102203-0030202133133101-0131313021311203-3311311223113211-1221213023012331-2100033133221010-0213121101210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- openstack.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3130012211020031-0002120313322321-2130230333013331-1212023133122230-0330322031203131-0002333033110223-2013031313301113-2020012000320312"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

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

<a id="canonical-0232323220300233-3023231332120223-3320120011120020-0233301211100323-3111110211131220-3030331102120002-0323103320211311-0132233102222330"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3203122230223200-1120230202122012-1212020130023221-0033210111112300-3120301122203012-3032001303311123-1110000310223020-2031000003012230): complete subsection reference.

<a id="canonical-1231102211321100-3101010101221131-0103212322211233-1020332100003323-3000213132330333-1232002320000220-3120003332303100-0332123010331023"></a>

<a id="canonical-1113103010301330-2203110030123022-3220120203113101-0233031213113320-0222032220200312-2202122323310032-1312031200032112-3113303312310300"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [lacp](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0232120202321320-1333102000301122-1101201122132332-2001323303023323-2121213002001331-1300231033223001-3023312202210022-1102313301312103): complete subsection reference.

<a id="canonical-2320303221301033-0030201232123302-2311213322312213-0332000332212020-0113000222000111-3022031231222012-3033233122320200-0022102010112113"></a>

<a id="canonical-2320003021321321-2301122120330323-0003203120323121-2233111031012330-0231211322233112-0122011203333302-2212000000112000-1301232330121031"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1120021213222000-2021232303202310-3020002212333222-2022322122031321-2031331000011233-1310223310222332-0323130233020103-0312001113121312"></a>

<a id="canonical-3101201303032331-2030121333021213-2302232212332031-0001322113123100-0200310313032213-1302002322001000-2111102012022120-1111123323013302"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1303211001221102-3032230211100322-3133313021300112-3230302313120003-2012033203001122-2132110222011031-0230311233130113-1313301123333311"></a>

<a id="canonical-0331212030022201-0113110010310321-1021013032223121-0002122301012103-2022131203032110-3133113210000231-1121212210111222-2312221000201322"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3203122230223200-1120230202122012-1212020130023221-0033210111112300-3120301122203012-3032001303311123-1110000310223020-2031000003012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3303323332321002-2313023131102203-0030202133133101-0131313021311203-3311311223113211-1221213023012331-2100033133221010-0213121101210312)
- openstack.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-0321331102031100-0330303110302310-2023220200322331-0103013220021331-0122330111120333-2012030303211230-3002301031000033-3231012202220333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

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

<a id="canonical-0232120202321320-1333102000301122-1101201122132332-2001323303023323-2121213002001331-1300231033223001-3023312202210022-1102313301312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3303323332321002-2313023131102203-0030202133133101-0131313021311203-3311311223113211-1221213023012331-2100033133221010-0213121101210312)
- openstack.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-2132323230120021-3202132121330320-3131122120221030-1000201212112020-2122211131132013-1002323311312023-3211300202000112-3021303220031030"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

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

<a id="canonical-0301031020331011-3100103033103323-2131010203332100-2030120311030232-1010203012210000-3333022320200111-2201023220023210-2022021132030031"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-3330200320302230-1323231030110333-1003230321020222-2032023232122222-0230302231203023-2222303033321000-2231130122302020-2211213010010120"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0033012012013111-1223332331302012-3132032311123303-1322200212211112-2112111010202020-2230003221111033-0320033013103232-3113120330303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- openstack.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3103221031003211-2101312123000310-3022121020231022-3033030221210311-3133312021203303-2313201212332013-3303001230122203-3013023123323203"></a>

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

<a id="canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2110330111013202-0323231220121021-3201012121121132-2002232302120201-3300320000332301-2123122023320121-0222301223303230-2222302332201202"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-3301132133223003-1312301020330012-0302010212112321-1131122101320200-0130222132332231-1331011310002002-2032230320123233-3102203223201321"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302213011221101-2301011021031131-2001301203223312-3002332223210013-0032023301233333-0303100130032220-1122330222130100-2020221111003121): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3121333210031012-3000010121232113-2302221031020111-2131122310323021-3113313302023321-3102101313310313-3111131123022231-0330213002221320): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323): complete subsection reference.

<a id="canonical-3110010212033133-0022030031110302-3023212111230022-3202312201113232-0023033203203002-1110223100300200-0100030003311120-0020300000221012"></a>

<a id="canonical-3022213120220213-0210003321322032-1301203323102102-1302210000002121-3231120133000323-1020033023100313-0312310323102200-3012131133300110"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-2213113021230103-3210310310331002-3332331101013230-0020201013200100-3122302311111233-0302212010011032-1231212021030320-1032302212212023"></a>

<a id="canonical-0221231011000130-3210230021321033-3001001223020301-2323202210013303-3031101202120022-1131312302003200-0231203331330203-3113010312313030"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1020210110201002-2232212100303010-2231211112030210-3033022011032200-3021302120010313-3210300231002213-1333132210312333-2123123122233133): complete subsection reference.

<a id="canonical-1302213011221101-2301011021031131-2001301203223312-3002332223210013-0032023301233333-0303100130032220-1122330222130100-2020221111003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1231112132103321-3111120301330301-2002213113323333-0022233011002113-3112123212310221-3100231123030132-3032033000010221-2010300232130032"></a>

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

<a id="canonical-3121333210031012-3000010121232113-2302221031020111-2131122310323021-3113313302023321-3102101313310313-3111131123022231-0330213002221320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3100100011313023-2320001230312320-1002231011213133-3312301213323331-0233232123201103-3112103333032221-0220330311322021-1301102021301011"></a>

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

<a id="canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1331012200212121-1112023121303323-2010322032031210-1013310112323100-0021002133021132-0001232211213033-0322030010003130-3033113222100320"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1223111030011231-1022020111230322-2032112231110133-2201003222230202-1112220323121123-0321321332211333-0221320300212331-1022131320202201"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2012313020033313-1002213311302222-3033320233030313-0121101302130121-0103112330330211-2102103310030203-2220001003302220-1313201103101203"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2331311330023120-0131323102012131-2230010301302332-3333303112332321-0232320212001101-0323120010213121-1232023212012201-3131030113311121"></a>

<a id="canonical-1211002111123002-0100132331113201-2211123313033211-2303300332302302-1331121203102023-0203300003010232-3230001323122031-0200020023010231"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3202012200311313-1203120321002320-0012321012333010-3131130113201023-3133023102213013-0032230131000011-1230112003110220-1011003030231233): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2211121032002213-1102310102110012-1113322012222113-2021330030022200-0011210101231200-1201232000233001-0303011320313200-1010122201333213): complete subsection reference.

<a id="canonical-3332201331222302-2332201211331222-3233133110230230-1100311023100000-3330300021021332-2003213333022111-3233120313022333-1102230023000012"></a>

<a id="canonical-2001001013113121-1000310012003010-1120321022313310-2112331021231130-0331123113100202-0222123212231200-1131030322200232-1220022120313033"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1303201012000200-3222000002130030-0222121201020100-3232332210132030-1320033303111023-1010102311113201-1020031022210003-3010122012121120"></a>

<a id="canonical-3031303320100233-1001301222001300-0213210210220112-3023113203332322-3120132313232132-0103133020233331-1201330122101300-3320111233132122"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1032313101120000-2201310131111223-2223120120322113-1001330331222133-1002310232300003-0321221131000110-0311301323121102-2210302300012011): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1021033232103131-2220230010133203-0331023221102132-3301321023133232-3113301110301332-1013012010112013-2313022101303101-2222211001121110): complete subsection reference.

<a id="canonical-3202012200311313-1203120321002320-0012321012333010-3131130113201023-3133023102213013-0032230131000011-1230112003110220-1011003030231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1212300302033302-3113302022002031-0211002100030111-2201332320212031-1222320032303021-3313300003331332-2111100210300311-2211233123002232"></a>

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

<a id="canonical-2211121032002213-1102310102110012-1113322012222113-2021330030022200-0011210101231200-1201232000233001-0303011320313200-1010122201333213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1300303131210002-3322202113122300-3223130311020111-3232100202201013-3301320322130213-3201202201013300-0100210102103332-2233203313030131"></a>

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

<a id="canonical-1032313101120000-2201310131111223-2223120120322113-1001330331222133-1002310232300003-0321221131000110-0311301323121102-2210302300012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3112322200102032-3230221221200233-0301231030001212-2001200110210213-1102001012033102-2201331223000202-3001200021210111-2013211010010301"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2210220000210020-0310233031222032-1203322202000313-1120011131201311-2213313231110221-0223002223210113-1312012102312322-0131230130020203"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1200133013203011-2002002003222113-0002110003301100-3220213211301232-3212023301021100-0201210223230022-2322010312100121-2032111023311213"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0101122200002201-0130202012133003-1230300310222310-2200012022201333-3221303012303020-1332330213002301-2232112331112032-1132333130000110"></a>

<a id="canonical-2210300113130302-3311322313121220-3202333031301012-2032103120032322-3013000201322310-1333102330101322-0122222323323210-2300230302100002"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-3003302233222031-3203123220030121-0010113213010102-0311111231021223-2031103330311320-2201122311000032-0000111101211112-1110002302333313"></a>

<a id="canonical-3211200223201020-0310023210123321-1100210322321102-0330332233311001-2223332000111331-2322113202212113-2123131231231130-3232300221201320"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1021033232103131-2220230010133203-0331023221102132-3301321023133232-3113301110301332-1013012010112013-2313022101303101-2222211001121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3201010203231023-1131230302120302-3212113132310032-3312131103002100-2123132313303002-1102211200311301-2313233132112222-0000223031333323)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0112320031222311-2230013113001023-0223221333230003-2132202033301123-3312003201301032-1303132022002313-1313030320111310-2112211302313013"></a>

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

<a id="canonical-1020210110201002-2232212100303010-2231211112030210-3033022011032200-3021302120010313-3210300231002213-1333132210312333-2123123122233133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312)
- openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2202301122130011-3320322203032023-1112321323323221-0213122322223103-2103202001021220-3321033302112111-0311233331221300-3111231232300220"></a>

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

<a id="canonical-3312110102212020-3133323312001113-2332133032330001-1201312130321311-0113033001013031-3302233102300331-3133221000230102-1312231333010013"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-1332132200003210-3000221200211122-0311222303210112-3112333001330310-0100012033313023-3320322302123031-2010021122132230-2101010300323032"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-2033333113321230-0122003033301003-1302133221203002-3112230201223110-0223201223311010-3320113322202022-2310231230003111-0312220200021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-1122122013331330-1201223033201010-0031001111131011-0133123302213110-0331200331122303-0000321331300201-3223102301131212-2303323200333131"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-3101033330012111-0330321222330022-1111311131311233-1221011321100000-1100130033120100-3001333113010210-1332232102121333-1323312233322303"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1230220220030223-3032020002103300-0122313010230003-2211130010110200-2023313313001000-1333020222213021-1200312013033202-2303230122113223"></a>

#### `openstack.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1203203030013100-1120123023200333-0011120001320302-3322120132211113-0103000131111101-1120203223133103-2112001030102012-2023123022003000"></a>

<a id="canonical-0033003322002022-2202220313203112-0321220220021202-1230032212022211-1102230220102110-1230020121331130-0232102323300001-0310013011203333"></a>

#### `openstack.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-2010230211002310-3201222332322213-3332002231200221-0101131213001321-0202313112013001-1332303221102333-3001132133330221-3003013232302310"></a>

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

<a id="canonical-3011231131101303-1201033300223333-3313311102221002-3011131203031220-2121021303203022-1230333230232121-2221310020321311-0321021103113130"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2330030100330310-2311030111211203-3310131231202220-0313002230320120-1001323210010202-0021231121132212-2130300203010032-2202023033123121): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121): complete subsection reference.

<a id="canonical-2330030100330310-2311030111211203-3310131231202220-0313002230320120-1001323210010202-0021231121132212-2130300203010032-2202023033123121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0200132101110012-0100133030123111-1232203202112113-3232021010212101-3101002222032302-2101230332033220-1121021032002102-0021221130310120"></a>

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

<a id="canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1030000032003101-3232110232221103-3120221311121330-3010011022323021-0122002032000032-3210101323322221-2102032230120223-2003303103023212"></a>

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

<a id="canonical-1302001332211103-1322021231000310-0112003101110221-3321133003012000-2133011233033222-1231131032212210-3123133031323332-2232122210310132"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121): complete subsection reference.

<a id="canonical-1003212210012031-0023332103302231-0023033120300021-0130003021313003-0311010010320033-2212213113110231-3131030130011330-0032212111101223"></a>

<a id="canonical-2012031231333303-0101101023100131-1021313001033000-2030212321302101-0323101122332202-3033202200032201-2210303010222111-0030033013101121"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120): complete subsection reference.

<a id="canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2312122002231210-1121021131011201-2001222232111322-2322133211002222-2313221100230312-1022302300201103-0220330123320221-0002222120321001"></a>

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

<a id="canonical-3221013102231230-3113010211223012-3223121020010020-2020033003200303-3220113211332202-3203003101302331-1220331010003303-3311313323121102"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0311202102302100-1203033000323202-1320311102113030-3331220303232132-0200003003312133-2301200233113021-3000213131121010-3312020021002312): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0020321320311321-2120233200113332-0333032222200313-1133311320211303-3111233222120230-1313233110102322-2123332130212221-3221130203320200): complete subsection reference.

<a id="canonical-0311202102302100-1203033000323202-1320311102113030-3331220303232132-0200003003312133-2301200233113021-3000213131121010-3312020021002312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1002233131110122-0230113031311101-3312322121212333-2332311033302221-1010103130300212-1031121033220231-2013300013030033-1110213223232112"></a>

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

<a id="canonical-1123111220001212-1010333130322311-3232221311300003-1301202123322303-2101211310132013-3233333130221012-0032101013223210-1221210302323120"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-2320233323022130-3223123111333010-0203333233203322-1300000211323201-3212103010001110-0321303000202203-3321103231110030-2330203021221220"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0020321320311321-2120233200113332-0333032222200313-1133311320211303-3111233222120230-1313233110102322-2123332130212221-3221130203320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2232100313022132-0330112321223013-3121323132100312-1113300131103223-0310020331032300-0133121331120132-0332230031013102-1002200231232331"></a>

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

<a id="canonical-3100320320213310-3301331101230003-3131323220020310-3211033103321300-0101330231123120-3020321220220322-3200232331100303-1120021023032010"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-3200203230333133-0303233312013201-0310213230031013-1213022232101232-3332122122300001-1332230012122211-3333113230023102-0221020013303101"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1133001333010213-2210203302033102-2232331101033031-2130033323320301-2011201222321112-2303011232302011-0303223030120211-2111023001020033): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0213310023122123-1320002031103022-3100122103322020-3012121121100200-2211112132100023-1322100022202333-1223213221311010-1230321223330333): complete subsection reference.

<a id="canonical-1133001333010213-2210203302033102-2232331101033031-2130033323320301-2011201222321112-2303011232302011-0303223030120211-2111023001020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0020321320311321-2120233200113332-0333032222200313-1133311320211303-3111233222120230-1313233110102322-2123332130212221-3221130203320200)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1122223121201321-2332332133011102-3300301002203201-2120223203311012-3002212001011223-2111023203233122-1223000230230102-2023123232220022"></a>

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

<a id="canonical-0213310023122123-1320002031103022-3100122103322020-3012121121100200-2211112132100023-1322100022202333-1223213221311010-1230321223330333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3230013300132322-3231312221122122-1221022312310030-0033200133210320-3031211030311012-3203003321300300-0202200010000332-2020111233230121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0020321320311321-2120233200113332-0333032222200313-1133311320211303-3111233222120230-1313233110102322-2123332130212221-3221130203320200)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1230323012321321-2323001331301002-3312132203021203-1200003210210311-0203333023210232-3021021332223301-0201032232100022-0121300030323112"></a>

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

<a id="canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2311002202200221-3323010201332110-2211113113231012-1023302131320032-0103200322230233-2131112120022210-3213310110230121-1033103223121201"></a>

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

<a id="canonical-0122301202223123-1113212202311112-2001030323320310-0231313120322223-0312211130131303-3030233230032331-1200313210202113-1100221312023000"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0210230130033023-2100132303120003-1011113210213313-1100132311031321-1031001013203130-1332210003332123-0221210321022100-1303223003020321): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0113221022021220-3123312010301112-1221202003123311-0300020103211310-3100233323202311-3023122200330313-0023312101332112-2233000230330323): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1101021020002321-0210303232330113-3021100120011010-2233103102133001-2200332120102122-3113020020330021-0231133231033303-3012110221010223): complete subsection reference.

<a id="canonical-2311202310100211-0101132133233101-2331110022112313-0210002212112312-3203302211302022-1321322331300133-0033312301300230-2132320001211023"></a>

<a id="canonical-2203232320301331-0023331133231023-3003032302113110-3301020111121333-2232113201001001-0013113132213030-0313100022330033-2031332333011020"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1013132323313101-1210000202312202-0111012023220233-0212121303111202-2111123221030131-3220121110332133-1020333001121233-2330002300121313): complete subsection reference.

<a id="canonical-0210230130033023-2100132303120003-1011113210213313-1100132311031321-1031001013203130-1332210003332123-0221210321022100-1303223003020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1232020320321201-2312130323210022-3233023002213122-0102312133211100-0100333110212331-3232310012033120-0100121233101021-1122012002210012"></a>

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

<a id="canonical-0113221022021220-3123312010301112-1221202003123311-0300020103211310-3100233323202311-3023122200330313-0023312101332112-2233000230330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2310330222312231-0010302210303312-1101101221201330-3001121221212223-3311031301002033-1133103300001311-3011020322312302-3011222232211121"></a>

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

<a id="canonical-1101021020002321-0210303232330113-3021100120011010-2233103102133001-2200332120102122-3113020020330021-0231133231033303-3012110221010223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-3102131212013301-2220202301320322-3321313313012012-3132003331221031-3312201210303213-2113033331010322-1200021320220122-3323300000032021"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1011320322311033-0221013103233121-0133233322132110-2102111120022133-0233130230013022-1123130320001223-3221232013321312-2303101130202030"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0010001113213201-0002031100112331-0231031110213012-2312032030120302-0011332031103333-2003011101313313-0010130320222101-1121330002122211"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1322031230020213-2111112311212333-3130011201202200-1001320111332213-2211303303231021-3202323322322202-1213121002012033-0002300333220230"></a>

<a id="canonical-3123101322332223-3021221002021122-3132013122032312-3001122120102031-3222300123330022-1033301322213302-1131133110201320-3022323332102223"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0113010323001301-2101201210232100-1112020313112123-2022211310102302-0333200022313013-0333221133313330-1212001011322300-0121311322221213): complete subsection reference.

<a id="canonical-0113010323001301-2101201210232100-1112020313112123-2022211310102302-0333200022313013-0333221133313330-1212001011322300-0121311322221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1101021020002321-0210303232330113-3021100120011010-2233103102133001-2200332120102122-3113020020330021-0231133231033303-3012110221010223)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-3002323322111033-3210200003020331-1121312000302100-0012100300001112-2220102330103332-1001322231210113-2102333333013302-2312300132023003"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3322320202323110-2032222311121200-1120012132132200-2101213102323300-1012021023101231-3013333211133321-2130332223030323-3032221322030132"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-1231033033010022-2002232313003102-1313122032003312-3113330031112003-0331123102302230-1201331030132100-1203022110031130-0221022233013033"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1002102322210303-1221002112012231-0200211320232230-3130002030130122-3222220023111122-2012230122121031-2303013333002102-0312310012330303"></a>

<a id="canonical-3300032013123110-3202122203332020-2213013001232023-3123132330033123-0223100012022123-2023013001130312-3221003112131013-3101010230210322"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1013132323313101-1210000202312202-0111012023220233-0212121303111202-2111123221030131-3220121110332133-1020333001121233-2330002300121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0121012021310311-3023130022320203-2010210230130202-3233220123322102-3012313121001233-1103333001201020-2230123322011133-3233022212100033)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0132331133112013-0312301003323220-0202111011213220-2322200301303301-1311322333120100-3303222321023022-1130202221210311-1300230301131003)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0323113011320310-3332313100022212-2233033302223031-1111031301112313-2112030312330222-3303302322301031-0233032032123332-0320021320010102)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1320011103211231-0220113101112201-1031222100012220-3112202323332112-0010110132030003-3331021102110233-2311013233232000-0022002012231301)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1302001330032220-2021112110202212-3232202101322221-2212022033012303-2300111122321323-0000312231030030-2322132210211023-3301313103033121)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3013100313020132-1131123321032033-0032320231303211-0132323210333012-2131202232200121-2223122321331000-1203330013200232-0112003131311120)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2012023002333013-0221122010001220-3131003331001011-2102210121131110-2302023123312323-3310111213130200-3330212121012321-1213311030011320"></a>

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

<a id="canonical-0331023200322232-2201333032121011-2230102203322011-1213130031312101-0230002011302302-1300033212123001-1011302021123311-2333223013310210"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-0003200103123303-2333303310111122-1301233321002212-1313000300102203-2023131322201013-1320022320111330-0330333013032310-3213232222331101"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

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
