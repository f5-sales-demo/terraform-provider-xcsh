---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-0320130330003210-3303331021101321-1000330020223030-0010033332003002-0122230011310020-3220302021132223-2030122012211023-0001031233101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

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

<a id="canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful` properties

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

<a id="canonical-2032301120021233-2202201021022302-0321310310302222-2321111131221020-0023221020320110-1111132120030303-2200211320322113-0303311301102123"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--network_interface--reference--group-002.md#canonical-0000012033230200-2123322310320231-3023332320131122-2113112303013230-3121022111002322-3101310220220213-3313302121303200-1220311003030223): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--reference--group-002.md#canonical-2010033133131333-2200020111311110-3233200132131213-1323112111122201-3123200131221030-3012302210110002-3233132121302011-0323120301311233): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--reference--group-002.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311): complete subsection reference.

<a id="canonical-0202210103210122-3111110112300232-0221230111301331-0120222112303133-0232113312333000-0031011330211022-1112032233321331-3233312210210320"></a>

<a id="canonical-0033102102221321-0231113230023320-2233331223311013-0201001121333131-0300222201223002-3201100312302122-1220120012111133-2103011211000031"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](data-sources--network_interface--reference--group-002.md#canonical-0313100122221133-0112300123320232-3110111202202033-2033121211010013-2332101313222002-0200022012131121-0232032031100332-1212210312222200): complete subsection reference.

<a id="canonical-0000012033230200-2123322310320231-3023332320131122-2113112303013230-3121022111002322-3101310220220213-3313302121303200-1220311003030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3212233301121232-2130123311011332-3100031100222122-2321300013113032-0121030012233301-3021101113011300-2322121331111003-1010220110022310"></a>

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

<a id="canonical-2010033133131333-2200020111311110-3233200132131213-1323112111122201-3123200131221030-3012302210110002-3233132121302011-0323120301311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3133223121002201-2331311321300100-2131011300122101-2232030221102101-3212132132301102-2112023122023002-3211310212131030-3313111233322311"></a>

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

<a id="canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1101111002210010-3110310301211130-0001121222011113-3220310002300303-3313203200032302-0212320001313320-2320233333221330-3322200110003002"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-2220320031213230-3230003202122201-3223322330203120-3023322233012200-2011003221000122-1330332233223032-2321232302220013-2103032221313222"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3323033131230022-2230223001113030-2122320232021033-1231011013122011-0012010201200222-3231031210313033-2222212311132222-2232211021003113"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--network_interface--reference--group-002.md#canonical-2122032230101010-0300100203020123-1230133021121321-3110131013003313-0022203123012130-0101232010230201-1222303132013023-0123303122130232): complete subsection reference.

<a id="canonical-2122032230101010-0300100203020123-1230133021121321-3110131013003313-0022203123012130-0101232010230201-1222303132013023-0123303122130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--reference--group-002.md#canonical-3231212111002233-0013103002113131-1200210212010331-0002131212313132-2023203102133010-1133011031123310-1200001010211332-0322313033312311)
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0003113100333120-2003022113113231-3012230001331323-0220110210031032-2221121311230012-1021213333333333-1203020100011003-3323010023133122"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-1120211201132033-0120203011322001-2321100002301310-3310331233211120-1002231002001220-2222123313222310-0322033231023102-0123133231200212"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0330002210032113-1032002112232310-3322001031131312-2311323000223330-3110002301311010-3311132300222112-0220202121131330-3222030202233130"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0313100122221133-0112300123320232-3110111202202033-2033121211010013-2332101313222002-0200022012131121-0232032031100332-1212210312222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--reference--group-001.md#canonical-1323010232231123-0020113031301313-0002210103011230-0020131232002220-2212322122113213-1222233132010112-3032131220301300-1300133010231001)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--reference--group-001.md#canonical-0021212102220102-3132201322323033-3312000213120233-3210202111213333-0230000030131231-3332223210331000-0220222023023030-3313011101202201)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--reference--group-002.md#canonical-2110300223322131-2103211300101333-3321030213032101-1102312213013133-1010130220220221-2030332303212213-3132301322132131-0201031123200020)
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

<a id="canonical-3102202303112100-3212202310200023-1211232210122033-2122331130312300-1013332111112233-3101333012312131-2000312310000032-1200110020323201"></a>

### Direct properties for `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2020322230121201-3210301200123321-0332220333133332-1102011212111332-3033020013103203-2110123310212220-0122133022023303-3001303030222001"></a>

#### `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-0131322022011012-2232101232210032-1121023033132330-2102122010330022-0303011212231321-0030120321323012-0002332223030223-2331132203033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.is_primary` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.is_primary

<a id="canonical-0213021100100001-0330331013011332-1301300210130322-2100302201203120-2333131111122332-0321010120202132-2110223133113210-3222133330102032"></a>

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

<a id="canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.monitor` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121330013003013-2100012130332222-1133031333201222-2021321321232111-2002301003121210-2322111332332230-0221011000030231-2023233203032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.monitor_disabled` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.monitor_disabled

<a id="canonical-1233112322031330-3011130002322221-0222223321331223-0210122323213220-2313020113230300-3112300233331000-1233113322320200-3102131010331130"></a>

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

<a id="canonical-1113032321321001-2310200133001032-3213200220203223-3000220113031202-1230221312200012-1202311001030122-0321131000121000-1111313323032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.no_ipv6_address

<a id="canonical-3310222213211122-3121201300312013-1031032213131200-1110202122301030-0013031232130220-1022323122000020-0201132201331330-3130312311231002"></a>

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

<a id="canonical-3312231302230300-0022030322113100-2201123131132030-3100011010120320-3310232100021031-1310303103110322-1031013222131303-3130102302313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.not_primary` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.not_primary

<a id="canonical-2231103231122332-0203212220131032-0332002321021100-0002002031012132-1310030202110020-0322100203202133-0001002210002211-2202103002010330"></a>

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

<a id="canonical-0212312303120132-1113000211222110-0213011322231122-3213230332311203-0133301012122321-3102222301230020-2232303202233323-3312131000300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.site_local_inside_network

<a id="canonical-3112330210230232-0131201333110111-3323130032020110-3103031120333133-0333111030221031-0100321022123313-2032200230112312-0103130122113323"></a>

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

<a id="canonical-3132210030013112-1102213013013232-0321001230003132-1320121232330023-0311112012212022-2113030330100121-0300331333001210-3033033130201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.site_local_network` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.site_local_network

<a id="canonical-1203322000313333-3321000133230210-3020130120303110-1021033120123202-3131132100110100-0002133230231110-2013002233012320-0223232323231110"></a>

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

<a id="canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.static_ip

<a id="canonical-3300133011111203-1330212301220211-2303202300012231-1111013302212331-2002000121213230-2332030021000223-2300131012003112-0110020231130323"></a>

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

<a id="canonical-2132310233011220-0210110313310133-0312021332100213-0312110002111010-3020002122022100-2213131303000100-1311003331033000-0201300220210301"></a>

### Direct properties for `ethernet_interface.static_ip`

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102): complete subsection reference.

<a id="canonical-0133120302310230-3112011022020110-2212301300100010-3200021301000220-0202020000132300-0023233133013112-2103031120130301-2233202332102013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ip.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-1031130201223312-1002010011231021-1023131120012231-3312313020123003-1013130302310212-2111103023022103-3302133032320000-1212322300221320"></a>

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

<a id="canonical-3023220223012300-0223330010333123-2213221001321310-2321031301302033-0330122032223122-1112333130103130-3121011013112201-3111001220211000"></a>

### Direct properties for `ethernet_interface.static_ip.cluster_static_ip`

<a id="canonical-3010023201022000-0002103200211213-1131101320013211-0230300231230300-1111111333312033-0002202002313212-1232103100112123-0320213033110323"></a>

#### `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3230031122333310-1130013023032301-3303011331212201-2231332320203231-3323013210221212-3302000020222320-0210111121010031-2312311202031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ip.node_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231)
- ethernet_interface.static_ip.node_static_ip

<a id="canonical-3220110030331230-1032002202032023-2333010201201130-3203233103311303-1132321012031032-0130311200010112-3311001102201303-0202012112021300"></a>

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

<a id="canonical-3203130001133021-0002120131333200-2111101013113010-1013212323210032-2321201221102213-2130133203020311-3212312003231332-1320100133310133"></a>

### Direct properties for `ethernet_interface.static_ip.node_static_ip`

<a id="canonical-3110022012211100-1000330221221332-2010203022002301-1112121212213000-1132203221022123-3303102022113022-0330020230001310-3230312302113303"></a>

#### `ethernet_interface.static_ip.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1130111221303111-0002310333300023-1331131011020111-3233233323100022-3310131131300331-2030130313023332-3320222132110323-3110030012310013"></a>

<a id="canonical-0112323311130033-0012031313203133-1210333131130211-2332111320202021-3111210330032300-3222333011221322-2310021210301221-0200120221321023"></a>

#### `ethernet_interface.static_ip.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2103113210201222-0011033303212021-2102120302002230-2131021130330311-2302320233322121-2203212323322031-1203033120323103-2300131000300201"></a>

<a id="canonical-3302220010003331-1011220101320022-0012031101300200-2331302321201032-3210320023103010-1231010332022331-2333131121200111-0011200033010111"></a>

#### `ethernet_interface.static_ip.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.static_ipv6_address

<a id="canonical-2012012100233103-0020302021020220-1021110301020301-0211320203011110-3033201122200101-2033021031201011-3311223230020021-2323030221003131"></a>

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

<a id="canonical-1001220331203023-3101220111311011-2021101121331213-3102113100302030-2221030300112032-2323133001220012-2302131113021312-0201021221311332"></a>

### Direct properties for `ethernet_interface.static_ipv6_address`

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2000333032022001-3322313322223203-1013133100320320-3033310202132103-2203113332121223-2322020003032112-1131120000310220-2132120112013312): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-2302202323030311-3302100320112003-2033232230321102-1101111301001110-1220330230220000-3231210201001121-0323103312331311-0111020120231012): complete subsection reference.

<a id="canonical-2000333032022001-3322313322223203-1013133100320320-3033310202132103-2203113332121223-2322020003032112-1131120000310220-2132120112013312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-2131010032210103-0301310100010222-3331312032212000-3120323223031120-0001221020222220-0122032103310201-3113330222323322-2203011002101131"></a>

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

<a id="canonical-1031202203211103-1222122112320030-1100102210233210-0330313233020022-3011202230213002-2103200301302310-3100323032121000-1232332130102102"></a>

### Direct properties for `ethernet_interface.static_ipv6_address.cluster_static_ip`

<a id="canonical-2211030020312330-1133122213031302-0120020203213133-3223132310330312-0223122211320132-3130022030310231-3330103113121130-0203223322321220"></a>

#### `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-2302202323030311-3302100320112003-2033232230321102-1101111301001110-1220330230220000-3231210201001121-0323103312331311-0111020120231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310)
- ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-1032002110320220-0320210300113000-2123213232011231-3102030233010020-2100200122333303-0212032303210331-1212023010232210-0111211101203020"></a>

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

<a id="canonical-0322101110121222-0221023200221211-0013010210121332-3320200030311311-1333310300001233-0333000123213032-3312001203310012-0122211133303010"></a>

### Direct properties for `ethernet_interface.static_ipv6_address.node_static_ip`

<a id="canonical-2000211102002202-2121133233233201-2131201331212102-1120013231111300-2332320211210120-2320130111311010-2102313123202221-3202002211323302"></a>

#### `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2003223203303030-0220312301313001-1310212102112201-0302012020010100-3100231221010200-2312133032331302-3133130011321220-3213220212301133"></a>

<a id="canonical-2201201331121231-2103012101301010-0223132333020223-1211213320123110-1123131123333210-0213022203102022-1133232233111011-1232223323112131"></a>

#### `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-0203013002223113-0201100122031301-0120300221012000-0202031322203110-0302200300032231-2003330001110313-3113332222230321-2321221020013201"></a>

<a id="canonical-1012012133100321-1312001233000021-1103011003222021-3030110322003023-0320330102332013-0121313100031112-2210333213333332-2220310130020122"></a>

#### `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333033210010321-2132211112223013-1230003210301302-3330133302221210-0113332030013030-1331320233211220-3001013200001121-2203312010010122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.storage_network` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.storage_network

<a id="canonical-2222221302222022-1230033212101101-1222323331022000-0130003210020222-0102133221030111-3331300022113330-0010133020123323-3310123231323310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-2303100002223300-2211310120020011-3133100223301321-2323322300103112-1230330031231133-3230103021330133-1023322211311012-0001221012231312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ethernet_interface.untagged` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-1231221132033202-1203021230002130-0223011322110002-1101210333012230-1031331120120320-3210100301333310-3003231230333101-0332232310023112)
- ethernet_interface.untagged

<a id="canonical-3332202232033323-3023222020030011-3230200311123331-3021012122032332-0131010220030001-1332031123233013-3120333302302030-1101110130010102"></a>

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

<a id="canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `layer2_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- layer2_interface

<a id="canonical-3202113010230203-3202320132300102-1102221210032110-3313003010222230-0232130003102200-0100010011021322-1003230323121233-0033100200303223"></a>

Type: `"single"`. Computed.

Configuration parameter for layer2 interface.

Additional upstream details:

Layer2 Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

<a id="canonical-3123323132031212-0313033111031330-2112010121121202-0023313012211111-2110321122102210-0112133131103213-3322121111132021-1003223121323132"></a>

### Direct properties for `layer2_interface`

- [l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021): complete subsection reference.

- [l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232): complete subsection reference.

- [l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-1310020032101313-3031201121133011-1230010330213320-2222133123313220-0222120030331301-1221220002202101-1300220222313213-3101332120332022): complete subsection reference.

<a id="canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `layer2_interface.l2sriov_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2sriov_interface

<a id="canonical-3132200021010003-1031232001333112-3313022332200113-1303021311010002-2023210112130201-2020211011321002-2333213110330102-2203023111202012"></a>

Type: `"single"`. Computed.

Configuration parameter for l2sriov interface.

Additional upstream details:

Layer2 SR-IOV Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-1132122011203323-0012120102021113-1222133010230021-2232110322321232-0003023013020323-3203122301202303-0231121202231321-3003320322122133"></a>

### Direct properties for `layer2_interface.l2sriov_interface`

<a id="canonical-1112033311201310-1131020021213012-0100012132231000-0130010112313332-0031102113231231-3321011302331130-3330230122113322-3320322303010203"></a>

#### `layer2_interface.l2sriov_interface.device` property

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-2103020311113321-3031302111201022-1221222221011223-2111222311100331-0012102101211202-1212301213000001-3301323103311010-0210321110232213): complete subsection reference.

<a id="canonical-1330333313200301-1210020213320221-2332100321130123-3101002311233003-2121222133120011-2221121223301102-1323010113313302-0213333111232100"></a>

<a id="canonical-1210123232200330-1032010301312332-3032203032201032-2010110230313121-3300010020013031-3000323202103100-2001331311312230-0002313031111102"></a>

#### `layer2_interface.l2sriov_interface.vlan_id` property

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged interface.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2103020311113321-3031302111201022-1221222221011223-2111222311100331-0012102101211202-1212301213000001-3301323103311010-0210321110232213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `layer2_interface.l2sriov_interface.untagged` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-2321233030331021-2213203223212012-2300223322311023-3130133333303222-0211302321321120-2232321231200220-2302011030310300-3221333030301021)
- layer2_interface.l2sriov_interface.untagged

<a id="canonical-2011230012310221-1233300113013111-1133230331120100-2233131231222101-0121332213122112-3213032120301030-1102122303112302-0002001231033222"></a>

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

<a id="canonical-0330330120020103-1222300012103030-1232232331201001-2303012223312000-1223302103300000-1221113313122123-1013010302300132-0102113300031232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `layer2_interface.l2vlan_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2vlan_interface

<a id="canonical-0232122010331133-3302132312333320-1102001130210213-1032031132220333-3011020220212311-3331102000223133-1021122021012133-1332033031100210"></a>

Type: `"single"`. Computed.

Configuration parameter for l2vlan interface.

Additional upstream details:

Layer2 VLAN Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102333122330213-2333233332200000-3233232000231103-3001211330120203-3300123001111220-3122202121022312-0112132320232320-3300001103110203"></a>

### Direct properties for `layer2_interface.l2vlan_interface`

<a id="canonical-0113023300301001-0130030320211111-0003311210221212-3211123101010302-2210122010030313-3203111331033322-3002200201331222-1230123323211000"></a>

#### `layer2_interface.l2vlan_interface.device` property

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0022022123232202-2000212030132302-0200123333331320-3212123321320020-0230211320333102-2321210321033112-1332212100230022-3101123223010331"></a>

<a id="canonical-0121200122133323-3111301101011021-3223113301301211-0022201332023220-3332232200113311-0323012301102100-0103223020303102-1013102003030021"></a>

#### `layer2_interface.l2vlan_interface.vlan_id` property

Type: `"number"`. Computed.

VLAN ID. VLAN ID

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-1310020032101313-3031201121133011-1230010330213320-2222133123313220-0222120030331301-1221220002202101-1300220222313213-3101332120332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `layer2_interface.l2vlan_slo_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030)
- layer2_interface.l2vlan_slo_interface

<a id="canonical-2213122321301030-2303330210333230-3330100000120121-2302332212012100-0303020220331233-2032310200101322-3000320320133331-0303110002021132"></a>

Type: `"single"`. Computed.

Layer2 Site Local Outside VLAN Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0222231203300001-1212103332023002-2103200210231310-1101330332323300-1312300113030121-2021122122113303-2011333100103111-1201323133110031"></a>

### Direct properties for `layer2_interface.l2vlan_slo_interface`

<a id="canonical-2113113000333321-2320012032112303-3133231003003030-0331223223212110-1022012222203300-2220123202300201-0222320222200022-0332302011311111"></a>

#### `layer2_interface.l2vlan_slo_interface.vlan_id` property

Type: `"number"`. Computed.

VLAN ID. VLAN ID

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- tunnel_interface

<a id="canonical-0202122131102011-3232213112222323-2223112333132222-1101010011000212-3201033332312313-1112311332113213-3310013212203003-0200303300303233"></a>

Type: `"single"`. Computed.

Configuration parameter for tunnel interface.

Additional upstream details:

Tunnel Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

<a id="canonical-3301313023312122-3111110023101331-0210212021021002-3311231200132130-0121030320000033-0331001303133002-2003303101210330-0333111223213213"></a>

### Direct properties for `tunnel_interface`

<a id="canonical-2301122233103120-3112112022331222-0231123032211032-3110312111020213-2203203103133233-0200001103313123-2101303013211020-2000122113310232"></a>

#### `tunnel_interface.mtu` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2233332222021020-0212203121222212-1201333202121302-3322330033213101-2223000322030020-2113221132000223-2321300020302312-1312133313133201"></a>

<a id="canonical-2002203202222023-1120211022001032-2113303311121231-2303030213030031-3120213110122223-0233211332103210-3120102033112331-2123202123310131"></a>

#### `tunnel_interface.node` property

Type: `"string"`. Computed.

Exclusive with \[\] Configuration will apply to a given device on the given node.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2122101203312001-2323121303032301-1322312322302221-3322311331232333-0230320212003212-0201120000322321-2120111221233331-0313003222333233"></a>

<a id="canonical-2223312000122111-1003311331023011-0302221312223011-3321310320103012-0223213031031022-1201011303010131-2123111110233311-1033012211033113"></a>

#### `tunnel_interface.priority` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-3200332131130022-3120200023013333-2002211203321130-1322310032233303-1003303131302120-0101303201220213-2222001110031020-2221303132033100): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-002.md#canonical-2221222222220313-0103032102101022-3232123232013030-1030202202321210-1300333113001022-0000302010201210-1310302133012213-1002100222303130): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032): complete subsection reference.

- [tunnel](data-sources--network_interface--reference--group-002.md#canonical-0102301312001330-2110301120232212-0101133301010002-3123313010311003-0221110020121013-1330123002102212-1022122002001323-1202110330310231): complete subsection reference.

<a id="canonical-3200332131130022-3120200023013333-2002211203321130-1322310032233303-1003303131302120-0101303201220213-2222001110031020-2221303132033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.site_local_inside_network

<a id="canonical-3131333032013223-1021201131223103-3121100210012223-3003312331101032-3123131112323102-3222200300101301-2233222321233032-0223102001110123"></a>

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

<a id="canonical-2221222222220313-0103032102101022-3232123232013030-1030202202321210-1300333113001022-0000302010201210-1310302133012213-1002100222303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.site_local_network` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.site_local_network

<a id="canonical-1000312200201102-3330320200321211-2212132111202100-3030330320110133-1111232132221101-0133103232233200-0112330201110013-1031001312110121"></a>

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

<a id="canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.static_ip

<a id="canonical-3113112132103132-3203032102300021-0103202132212110-3322112313010211-0122300231201223-3310202132113311-2301231320023232-2332012333313130"></a>

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

<a id="canonical-0322122112122131-0332122312231210-3221013322031003-3213001000132020-2102002310003110-0021313213122010-2320212010020032-2210222122123310"></a>

### Direct properties for `tunnel_interface.static_ip`

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-0000300202223022-3302100312022303-0202020302222110-3320022230230331-3210013203103021-1231321223111312-1003233220110010-3332031202310001): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-1113231101321112-3303000123310131-2000230003120133-1313010210333012-2131222233000220-3000230022123310-1032020033112020-0223203121322031): complete subsection reference.

<a id="canonical-0000300202223022-3302100312022303-0202020302222110-3320022230230331-3210013203103021-1231321223111312-1003233220110010-3332031202310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.static_ip.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-0321110100210030-1113001033321023-0210230320101121-1321230122332123-3011201332031200-1012330232211212-0022111022033113-3021201120332210"></a>

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

<a id="canonical-2023200011202230-1101330111122313-0311330221312311-0012312131301311-0130230201023033-3023123110213311-3211312231002010-1331203021310132"></a>

### Direct properties for `tunnel_interface.static_ip.cluster_static_ip`

<a id="canonical-3101130211022311-1023301213212323-2130213331001301-0033121110020201-3100330113310312-2123012113023310-0011222131002331-1120131102032133"></a>

#### `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-1113231101321112-3303000123310131-2000230003120133-1313010210333012-2131222233000220-3000230022123310-1032020033112020-0223203121322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.static_ip.node_static_ip` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032)
- tunnel_interface.static_ip.node_static_ip

<a id="canonical-1311332322200012-2123002113110032-1012212100031111-2220101010130003-1001110031223110-1022310030023002-3103222230003121-2202101023300133"></a>

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

<a id="canonical-1332323133230320-1223122223213203-1130012021020322-2230111010000220-3013223013221132-3002321302000022-3120210303131131-1213222122021200"></a>

### Direct properties for `tunnel_interface.static_ip.node_static_ip`

<a id="canonical-0030132022001213-1330101023110212-3111001212032032-0312030032331330-3232323222212132-3122012331303112-3131301211130202-0021330312203021"></a>

#### `tunnel_interface.static_ip.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303121121210330-1303230133303032-1230031333121233-3023220211332103-0310001121230032-3303030210323320-0303001222030121-2031313201031323"></a>

<a id="canonical-1202212223233111-0303303102303302-0203220222021312-1312130210312013-0331032102200331-0120202221121210-1132131011002233-1331233222110303"></a>

#### `tunnel_interface.static_ip.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1223020222311233-0203332000321022-1122103323030102-2020211220110101-3213000310321132-3103311321100021-0111220100031330-3112033020100120"></a>

<a id="canonical-2320201101012210-3010001030023321-1110323213002211-1033100321103331-3322313113003231-0210130011230203-2101000100131032-0133103232303333"></a>

#### `tunnel_interface.static_ip.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0102301312001330-2110301120232212-0101133301010002-3123313010311003-0221110020121013-1330123002102212-1022122002001323-1202110330310231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tunnel_interface.tunnel` properties

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-1030102220310031-2101220021010303-1132023020221330-1013302003330133-2331301123030223-3233302202221123-3320020030202302-1320303030232313)
- tunnel_interface.tunnel

<a id="canonical-0313310202220321-3100221300010201-0322011131122303-1332000011200001-0312032200123323-0111102223233001-3011312020223322-0022010333122233"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3122331031330020-3301110221220300-0313020100000131-0213330120110321-3132013312203132-1100223301001211-1122220010012120-0010002310212113"></a>

### Direct properties for `tunnel_interface.tunnel`

<a id="canonical-2231123023210121-3020002101303113-1032320312300323-0312120110111231-3213200002003223-2122202203203130-2033222030312013-2010033111001111"></a>

#### `tunnel_interface.tunnel.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2130220312101112-0223322103302131-1212020212132213-0031212331230103-1300202103021033-3113231311210031-2120332213101003-0033303310203330"></a>

<a id="canonical-2110103201001012-1213201220232220-2201333331100010-3331320313010103-3220233330132302-2021010101321200-2302220011312030-1123210001303030"></a>

#### `tunnel_interface.tunnel.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3011221000031031-2331131120023320-0002311032110300-3023121111001333-3120310103122023-3102131011323302-0032032203210313-2121012013321331"></a>

<a id="canonical-3300022100322312-1032102332131011-3312212013233002-2200212131321310-3213300113121102-2131130133302023-0020031300230013-0031121002003233"></a>

#### `tunnel_interface.tunnel.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
