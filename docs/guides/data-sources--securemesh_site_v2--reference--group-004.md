---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3033320221130332-2102030110033113-3021110323100200-1102113302220030-1113010313023020-3333031113102321-1123113100031321-3133223222002210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1102230220202223-2213322300303033-2312011002120112-0332000002310333-2031232223100302-0331033001033223-0332112211231133-3132100013000320)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1313212301311230-0211301001013212-1203211313301122-0331233300223213-2220010032001022-2130121213203212-1020331203000212-0332321133112210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1011332030310220-2331031120100200-3012131212101231-0310223301223010-0023200333130103-3111030233221323-2001220130310131-3023100323112300)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1233110230221213-3232211113312321-1020320331103303-0032330001012233-2112111320130210-1232113301231222-1233321322321003-1311122230131122"></a>

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

<a id="canonical-0113032121222100-0103313200132303-2213113122111201-0033030310112203-3222003023312202-0033032232321022-3113101121300110-2330112100300030"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-1121101320030113-1333122132203132-2310222211320111-2322332112211333-3202012021230003-3023110311321313-3111112133312100-1232013313021232"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1003112300021320-2203323231231313-1213020222332012-3020223231121021-1101100203111311-0002022313221000-2011300030112122-0111232201311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.monitor

<a id="canonical-0223123021133320-2120312331210231-2310200033322111-0123001330223010-0103100030003313-3301011122321222-3300312210001213-1311203303211333"></a>

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

<a id="canonical-1333020001033123-1020223030000320-2001200231333130-0021121223032302-3231020220033330-1010312101131202-3013202120000022-0333320200123123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0012200312233022-0220320133302300-0303310030012131-1233330300031012-0103203231120202-1312320102321133-3112330021302320-1023210110221322"></a>

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

<a id="canonical-1233311030230101-2220211320220003-2123131023021310-2130110232220202-0321031003011203-2031103311331312-1221123203000232-1323200321331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.network_option

<a id="canonical-1231000230311132-2201323233232210-0310310112313003-3122230001022003-1311010121102120-0120032322013012-2302321321120001-3203123022222222"></a>

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

<a id="canonical-3023300121212101-1131123231013010-0113223312121133-0002232013003111-2121022111101213-1120030212120001-3220332230011022-2013003231320330"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3300121001032321-0112121122311121-2331011211202221-1323312311220213-0203133123033222-3230302113130321-0211231120302310-3312021101310200): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0033021011330002-3212202012210003-0020311120301220-0000113303001310-0110331013300233-2112311312331222-1021011332322231-2300332331332011): complete subsection reference.

<a id="canonical-3300121001032321-0112121122311121-2331011211202221-1323312311220213-0203133123033222-3230302113130321-0211231120302310-3312021101310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1233311030230101-2220211320220003-2123131023021310-2130110232220202-0321031003011203-2031103311331312-1221123203000232-1323200321331030)
- aws.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0002301221323131-2202110222120023-1000321121013202-0013023113110332-1313031033333203-2321002311022311-2200200013301033-0212100322212132"></a>

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

<a id="canonical-0033021011330002-3212202012210003-0020311120301220-0000113303001310-0110331013300233-2112311312331222-1021011332322231-2300332331332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1233311030230101-2220211320220003-2123131023021310-2130110232220202-0321031003011203-2031103311331312-1221123203000232-1323200321331030)
- aws.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2031322213132122-0312112303011201-2222110011101103-3213213131301002-0312023221221123-0110132200102010-3322113122312120-2221200121301032"></a>

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

<a id="canonical-1201312031021133-3220222201020312-3212330322302332-0311331121330033-1320230010330302-1300320332213312-0311223222111013-0021102311133100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-2102032112201332-3231232330133201-2312133233020000-3322033223031202-0210101112111121-2131323220231321-0130230201323001-1011302303310331"></a>

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

<a id="canonical-3301103230310113-2120302300212230-2000231030110101-2200012331021010-3232113303100300-2302002120101333-0023220220200323-1222003323232101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3131201233320233-0331331311230132-0231203210103011-1103033321313022-3100333102102011-2312203122332130-0013013031122231-2320123103211302"></a>

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

<a id="canonical-2213100203312103-1300030003333202-2012121300231200-0330031112123210-3333132311222012-2201000003002031-0321330301032212-3101301033210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1131022233112121-3101023132112111-0023321020003021-0302332011120232-0330100301111011-3020303312330021-2030311221031130-0232231300021033"></a>

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

<a id="canonical-3230330233102120-0020132202333302-0013132300212132-3110003221201001-0312011332110323-1331331332220022-3312001002130230-0123113332130131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3202001223111200-0331311102121200-2102021031030002-1132220012331032-3320002013330222-3002102300021323-1222103000331302-2020032332022130"></a>

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

<a id="canonical-2202212012100012-3330130121131103-0211311012323021-0123132112231332-3321113312022121-2133201033103202-2212023020102222-0303232130130322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.static_ip

<a id="canonical-3300202323211101-3331102101122110-2200230212223302-2103312131211221-2300113013211101-1232202132311112-3302112113120011-3321011021231123"></a>

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

<a id="canonical-2032133102303030-3210203103100101-0310021120312203-1133300101201123-0131310211313221-2233210312002103-0233321233031023-3121032132102030"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1223222322031013-3232333230210232-0202200330121230-1020202132033331-0030301020120023-1112123320202032-1233130320302203-0103110311210333"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.default_gw` property

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

<a id="canonical-0322231300221111-1213230201310202-3021023032100202-0223113030103330-3032131222032210-2131331232230030-0213202302012201-3330301010001113"></a>

<a id="canonical-3232132001002102-3202131203122101-2001121131003022-0130111013103201-1112332123121312-3000130300020212-1213213032313020-0021233222330221"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3231322233111123-3103331303220122-0110201100020331-0213123333211033-2001203303323223-3311011121133210-2331221323302002-2001003211130022"></a>

<a id="canonical-0303322213130221-3233020331221230-0203021113113303-2220232023331023-1333320231230001-1000101220212020-3133232121010311-2113332130120123"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.ip_address` property

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

<a id="canonical-3320032303130010-1131032120200303-0121123200123210-2220223003202120-2301113312232121-2103300002230201-0302303203310101-0012100300033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0131122100130222-1303001000103120-2213030212310203-0133022032001012-0131302122331310-3033200322323003-3200301233301322-2010120213310312"></a>

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

<a id="canonical-0011101210333323-3322203110132102-3310013120123313-3122210120302002-2000003223302200-1123133032102102-3121122030333231-2000032321133203"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0133331232322022-2033232120122301-0132122033312332-1010321123111331-1131222220121121-1103310002202122-1331230313202232-0031101101132023): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1130013022323120-3123322001221203-1131022220120120-3221333213113212-0003122321203001-3131130300333110-2020310310001322-0311223230221331): complete subsection reference.

<a id="canonical-0133331232322022-2033232120122301-0132122033312332-1010321123111331-1131222220121121-1103310002202122-1331230313202232-0031101101132023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3320032303130010-1131032120200303-0121123200123210-2220223003202120-2301113312232121-2103300002230201-0302303203310101-0012100300033333)
- aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1023322110221131-2310203203202020-3223301133103010-0323222003300210-0220033331003000-0331210303233121-2030113211312233-2020110132310301"></a>

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

<a id="canonical-0212223220310030-0100032220123121-2123121303312230-1030010302330013-2022230222023313-3012301231030223-1120131203132011-2301120333012301"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-1302101022133112-0013122032102230-2303303023131201-0330122313020322-3320310222311321-1000332010112313-2030032221000202-1333323330223011"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-1130013022323120-3123322001221203-1131022220120120-3221333213113212-0003122321203001-3131130300333110-2020310310001322-0311223230221331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3320032303130010-1131032120200303-0121123200123210-2220223003202120-2301113312232121-2103300002230201-0302303203310101-0012100300033333)
- aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3321303312103332-2021110332032110-0013100201312133-3203002310103322-1113200302013012-0131121130130213-1002232112023021-3220320303221310"></a>

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

<a id="canonical-3201010302223031-1222130100121231-3221230213312031-3022112101233203-2200123010230311-3310211201312103-3112331313321302-2203102031123320"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-1221112323320300-2131113112021013-0121113110331313-0012231010001210-0103202311303132-3113332331122331-3300103102200012-3011223120020130"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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

<a id="canonical-1322101320202212-1202030031223021-0120133022020231-1322313122203013-3111233300212031-1312230133101202-3300313002310313-3203210330122033"></a>

<a id="canonical-3102012222232022-0021203021121112-3213020131121332-2102013303211211-3230303233210331-1321213030221011-0301323320000233-2302112223133330"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2321112001311311-1030203221111111-0213303321000201-2132333202213222-0212303031023203-3100311211010201-0133220211323311-2131021221320200"></a>

<a id="canonical-0111133321202312-3323001210301132-2013312121223132-2223001010003013-2303000211203312-1021203320222032-1212122313330320-2200300220220311"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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

<a id="canonical-2232101303130011-0300122130123332-3320203030101201-2201233120021220-1320011231222131-1012223220231022-3330131012310312-0021203201322022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2132201313231011-0001110132101133-2000030121011102-2000310230301332-1010032022222021-1323320021031332-0301312000322330-0121133110301310)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0331311101202333-2332012020331021-2210003331033112-0100213313211001-2120123322223310-1033001322031000-0021312321121030-2033311302120313)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-0121201223000303-1123230003302031-0112010000011001-3132201020102011-2030231220132021-0012111330212032-3023130112232003-0123122103112033)
- aws.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3313010103103103-0301213311313201-1203132232032002-0300003312131020-0200330330211102-0232021112010011-0000232001011231-1000303100032002"></a>

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

<a id="canonical-0000012303302330-3300220322320012-2122021302211230-0311003220000211-3230102120013010-2101000131131100-3013023311030013-1311330310201111"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1112322130312221-1121310333020302-3210311130210022-1002123313231122-2010311013203131-3201133002003313-1012132101303300-1121122133231332"></a>

#### `aws.not_managed.node_list.interface_list.vlan_interface.device` property

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

<a id="canonical-1303132010122213-3230300321333223-0231332022031231-2311023203131133-0110213311002023-3010300000310121-1112031232003331-1230011300300221"></a>

<a id="canonical-1132230203223010-1031130101212113-3010011233101002-1223322220233101-1033131300312101-0331330220320201-2303321213011002-3301210103300011"></a>

#### `aws.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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

<a id="canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- Azure

<a id="canonical-3222211031011021-0310022230210203-0123011232002132-0330332123331122-3221213013101000-3232100111321112-2320133201322131-2121223001332101"></a>

Type: `"single"`. Computed.

Azure Provider Type. Azure Provider Type.

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

<a id="canonical-1003111031222102-2101030113130031-0313132331000033-2111013302300331-0003123213303121-2120131303032230-2223320200012222-3311301201203302"></a>

### Direct properties for `azure`

- [not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033): complete subsection reference.

<a id="canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- Azure.not_managed

<a id="canonical-1030332333102300-1121323001301033-0033130032221111-2221312231303013-1213000003000321-3021113110333232-3310330313003300-2322322001222002"></a>

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

<a id="canonical-0333212112213121-0230201213312322-0021100013033121-2123331300131100-2111001310000223-2203321032312122-2003003200201021-2020330120113320"></a>

### Direct properties for `azure.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323): complete subsection reference.

<a id="canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- Azure.not_managed.node_list

<a id="canonical-2130013122132311-1022200330323231-3210103023312201-1233210023200113-3000313202112312-0223111012302112-2231233221200213-0000031011111201"></a>

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

<a id="canonical-3133031333130023-0031013001131003-0102330102330312-1312202030013010-2302133301303212-2200203010100211-3212101232302233-3000211212230020"></a>

### Direct properties for `azure.not_managed.node_list`

<a id="canonical-0123200131232130-0231331103031201-1233333103032103-0032121221021122-0112310321122210-2233132102011020-1333223132232201-0213312211033100"></a>

#### `azure.not_managed.node_list.hostname` property

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110): complete subsection reference.

<a id="canonical-0212310001331100-3211033203012231-1103221333320023-1011012210313213-1021012031203111-3222003012312210-3310112310031122-2302032302202011"></a>

<a id="canonical-2111123133223303-1201012220333301-3311300212302301-0303010021102012-0212230111120120-1010330302331013-1131202223221133-3232130213111101"></a>

#### `azure.not_managed.node_list.public_ip` property

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

<a id="canonical-2233231210332002-2323230322102102-3113330121231232-1132132133131221-3203331110302132-2112322131331032-0223120321320102-1221113132223101"></a>

<a id="canonical-2231202102001221-0100231230213132-0110333112202313-0130122212222213-0323010223003223-1230221133100210-3331232231232300-3011212222201000"></a>

#### `azure.not_managed.node_list.type` property

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

<a id="canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- Azure.not_managed.node_list.interface_list

<a id="canonical-1313221130130103-0002202310211202-1203002032110230-2133210203201302-3311212210221330-0202313320210102-1322102220221232-0222030001032300"></a>

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

<a id="canonical-3102231001232210-1233300322012133-1011000110213032-1210301103311033-0100203000223311-1221210312120102-3023023223000033-1203020331321103"></a>

### Direct properties for `azure.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1312022313112103-2102113232202332-1310102121123312-2020120232102111-3030201032010021-1131322121110311-3130310013033101-3130200313010200): complete subsection reference.

<a id="canonical-1001333002100010-2211312333123212-3020222113211130-3232332103301312-3221130303130033-0231100012003013-3311112333320111-1120133000003323"></a>

<a id="canonical-2010211211122201-2232023111233013-0131233033210202-0022100010330102-3212223101013320-3311233322301120-0120120223112212-0000301321232122"></a>

#### `azure.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3230023301000013-2012212302121011-1003100230132113-2332233300000222-2330230001001333-1120221331323100-0003011101331010-2003031302302113): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0021003321201122-1212311211220333-1323301013310202-3000120223023303-2012132220202100-0030103213002332-1000201120001023-2131201220133201): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300): complete subsection reference.

<a id="canonical-1201100112222120-0330332303022130-1122320233123321-2311302022020110-3300330100021213-0130100122220331-2011112112002323-0303323310020030"></a>

<a id="canonical-3022313102231202-1032120210120132-1233222312020202-0223313310321223-1032121103132102-0120230311313021-3001113312100122-1223221312321213"></a>

#### `azure.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-1312131310131032-1331101031031301-0103012312113222-3012233321002022-0112112032222001-0310002020212102-3331333312323000-0131032230200003"></a>

<a id="canonical-0313012011303020-1200203231300333-2013332311223330-2322033121331103-3323113111310232-1210012213021120-0300211310312023-2220310133233130"></a>

#### `azure.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3203321003130100-2303211012113220-1030231111302011-1121120131303320-0000101223333310-1200322320212110-1101120011020013-0011022231201102"></a>

<a id="canonical-2131023000113330-2302102112302131-3132302023020003-3221223011132232-1003102012212113-3111302112201332-0333211013230332-2022221133131223"></a>

#### `azure.not_managed.node_list.interface_list.labels` property

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

- [monitor](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1131313313000333-3211233010002003-1230232300100203-1012021222020133-2113000210302220-2302123323101332-2201302021102331-0033222003010331): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1300100302021332-3303213332003322-2313211100300103-3120320132323010-0013302303111322-2030223032101001-3320032330122321-0033300222133320): complete subsection reference.

<a id="canonical-2122212203211211-1121323102212012-2011133023000230-0222032322022120-2213202112123202-3131231222023102-0021031223023112-1002320323010321"></a>

<a id="canonical-2231232210302133-3003002013331120-3133301130121231-0030332111030030-0301030232132012-0032122330112032-0320021220000233-2001123010103222"></a>

#### `azure.not_managed.node_list.interface_list.mtu` property

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

<a id="canonical-0202211002322221-2100022211302331-3113120213122021-1012113132231112-2230311132100031-2023302333213322-1323202232120001-3012231100123031"></a>

<a id="canonical-0010322023013010-1220203111022302-3223113311133310-0133001302130220-0001233023122122-0020313310001333-2211032323213033-3002021302203231"></a>

#### `azure.not_managed.node_list.interface_list.name` property

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

- [network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1133322231300022-2223132221112121-3213201012123030-1013321032023312-3113311100202210-0111222231130203-2201332102232123-0133010110112013): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1303330031222030-2322131123213211-0211022223123330-2200101233221013-0301123113211110-1033312333022032-3022113122303131-0100311221221112): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1230222333021311-0203030103301232-3312231020322303-2231322120222112-1012313323303130-0111002013213333-0110101201133332-2221202233202021): complete subsection reference.

<a id="canonical-1212022222110123-2001212101310303-3301102011012203-2123113322131011-3321111221220130-0123110213221203-1133212201231131-3220030333202000"></a>

<a id="canonical-3312201303303123-3220231303032231-3030032010322021-1112222023133130-3311323220102130-2230212103123110-2111021200232200-0011110100110120"></a>

#### `azure.not_managed.node_list.interface_list.priority` property

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3312300103000112-1003121230323300-1010211000031032-3212003011301121-0201321310320333-2113330021330013-2113210102132233-3310222230002203): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1330131100210003-2302112112121101-0033313031321102-1101002100230302-1031110230002221-3101332131213300-1030000102331322-3032012100030323): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2112223023131102-1132132230021333-3133002200110113-3330231020030223-1203102203223221-0132103100330111-0300223123020201-3203321302130101): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1211120233231203-0110323123233100-2010300030212231-2030001200102110-0212003021130111-1213230011333211-2011220310200020-3211101133303300): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302322001010313-1200112121331130-2022203212131331-3303211013223201-3212310012133301-0311212203213311-2202110222001300-0120102320112003): complete subsection reference.

<a id="canonical-1312022313112103-2102113232202332-1310102121123312-2020120232102111-3030201032010021-1131322121110311-3130310013033101-3130200313010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.bond_interface

<a id="canonical-0222011013032112-1320213112331200-2031313022120323-0132320232112330-3130223032010323-3332032323310030-3032203132302233-0303332222000200"></a>

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

<a id="canonical-0120312333013130-2103233032330211-0303120333323202-2022030311311132-3311003220230211-2002110222122113-0003323131322302-0202100121322103"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0213001122323011-3213310100032000-3322213202001223-0223122313110122-2301320220330232-0333331011002300-1102010212120103-3123113112200322): complete subsection reference.

<a id="canonical-1011300020011320-3300021211202231-3331032300331023-1011023003031220-2330001130303130-1233110310313011-3113123002312110-1022131311310313"></a>

<a id="canonical-1332032002013011-1023011031332020-1303233202311310-3120312213320232-0031301020220223-1101121101102201-0301311110331001-0001320131122200"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2010213103201331-1213000030222012-3030020013211123-3323102333221231-2031112012123221-0101212230122012-2300010111203113-2202030233003113): complete subsection reference.

<a id="canonical-0001102322001020-3331023003003333-3310322010010020-2100013131223201-2202102130300131-3323032130013002-3201312121021021-1331021010330130"></a>

<a id="canonical-3333032322233112-0111121330203203-3320333322203230-3330232323132233-3012030311112213-2331200213131300-0133303032213030-2301110013120300"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-3022131330211330-0121030231002003-1322111220012322-1323211301101110-1201113123032221-0202310032232001-3220013132220022-1011001010301332"></a>

<a id="canonical-0332110310020311-1200011231021230-1113131022203302-2003201012012312-1022111303221003-2033011030033302-2300323200230200-2223001300230303"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-0303320002222110-1100313131021213-0332200112333132-2123011002232223-1021000312133223-0323110000011112-2101320301332232-1011101021330230"></a>

<a id="canonical-3003331311313111-2301030322003313-3223001011322222-2331033110202023-1223310310021211-3323231311221011-2021113012302332-3000021001321122"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-0213001122323011-3213310100032000-3322213202001223-0223122313110122-2301320220330232-0333331011002300-1102010212120103-3123113112200322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1312022313112103-2102113232202332-1310102121123312-2020120232102111-3030201032010021-1131322121110311-3130310013033101-3130200313010200)
- Azure.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-0121013110203211-1033111311020031-3320130002322312-3010223000230033-0213301123211113-1121223012033212-2313121312322232-0020310013321001"></a>

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

<a id="canonical-2010213103201331-1213000030222012-3030020013211123-3323102333221231-2031112012123221-0101212230122012-2300010111203113-2202030233003113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1312022313112103-2102113232202332-1310102121123312-2020120232102111-3030201032010021-1131322121110311-3130310013033101-3130200313010200)
- Azure.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1102013312120003-1131013332200312-0023021220100201-1030323303022312-2020322023321213-2333322120313121-3230221121200001-2210030330001010"></a>

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

<a id="canonical-0002012222233232-3212202011203320-1131131000123212-3130332212332121-0013131323120321-0301002202320112-0032120300120032-1113010133033321"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-0200101222310311-0330312232003333-3313220311203210-2321001103120322-1020210302212100-0202301213320013-2220221102120102-0023000132311223"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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

<a id="canonical-3230023301000013-2012212302121011-1003100230132113-2332233300000222-2330230001001333-1120221331323100-0003011101331010-2003031302302113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2011132330311300-0020011010000022-2120302032033131-2112331323203001-0330021122213100-3302333310130231-2121020132221303-2300120133310020"></a>

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

<a id="canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-0112023030012310-0322323201210210-3010031030100121-3200203101103120-3120231010002300-3220333331003202-3203121000023031-0112031213331332"></a>

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

<a id="canonical-2331232031022113-0001220100132101-1132132332210131-3030230103330110-1320233212030232-3121331223323121-1002332131311223-2022321132313310"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3231312023133022-3320230301100110-0022310031002113-3321112223203210-0123233311311301-2002320000011223-0231120103122210-3311010301230213): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2102132100130031-3010010100013030-0003312122232323-2322310312012113-0031330033232222-1310012022130330-2211221101131230-1110030200110332): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103): complete subsection reference.

<a id="canonical-0022003113230011-3222133221100102-1103010111230001-0332110233310102-1222000203103111-2130303220311002-3033131202031221-1322103020211232"></a>

<a id="canonical-3332231120122023-0103201232003123-0212103013021010-1123232003203221-1103331300022330-3323002200002122-3303331000033023-2312332322322011"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-3133013302301132-3310201103111003-0032101311103221-1010133033332120-2003320322322102-0202322322110311-2201030020012210-2322011100300110"></a>

<a id="canonical-1103120122201023-2002332212212031-0022021333022123-3013002123011222-1220230323213111-1311001122120230-1311202231311212-3322230003322230"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0213203231231223-3131302031221111-3033110312202221-2133320310100112-2302223201000300-3101032230200130-3130112012222012-2222000322021111): complete subsection reference.

<a id="canonical-3231312023133022-3320230301100110-0022310031002113-3321112223203210-0123233311311301-2002320000011223-0231120103122210-3311010301230213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0230032012311232-1223311030323031-3012003302001103-2022212221220022-0002033200030023-0221103230231122-0220213220323213-2033033110303211"></a>

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

<a id="canonical-2102132100130031-3010010100013030-0003312122232323-2322310312012113-0031330033232222-1310012022130330-2211221101131230-1110030200110332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3131002203311310-3132110332220010-1213002310011331-0103100330330310-3111210332031323-2103230211033020-2213120031232013-1313101100213012"></a>

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

<a id="canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1030100000033213-0231321313102033-0011100332000102-2201012112330111-2123102100220321-0131133111101011-2220002122220002-0330013201311302"></a>

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

<a id="canonical-2322220000110303-1103112131321023-2223321220331103-2313100130322121-0010113313220002-1323033130131110-1131131230220122-0103320011000133"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2111032031303032-2022002202211332-3113221011100330-2122313121220210-3120132110210012-0130230101111100-0123023010110123-0211023030031103"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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

<a id="canonical-1113301003210133-3100010313303100-2110103020232112-3121020302200323-2111002013221130-3112200230012203-0200101233331223-2101132312232313"></a>

<a id="canonical-1031303232221010-1300202102121113-2231200122221112-3332321203323033-3213312111232303-3221312020222131-1001333011333331-1003312122012210"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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

- [first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2333032120232213-0131121333122103-0030311133020103-1032203231213333-1110321223111023-2130110123333203-0131120123332201-3330322101210032): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2203032322231200-0210200333232331-2123220130321311-0001201333001030-1333201012212003-0131110123123032-2020123131332231-2022332211123003): complete subsection reference.

<a id="canonical-3230121321101213-1132110003103013-0332100001000121-3302202001010230-3113033113302122-2333022203231332-3110321303100003-2032302303213000"></a>

<a id="canonical-3001112133110122-3030130210200010-3230132020031303-3210102301312321-1322212010330012-1320321131220302-1002011303222333-2223010230331310"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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

<a id="canonical-1023011312221203-1312012221212132-3203030321212003-1110001032222031-3123100332023030-2110103111113202-1100203011002112-2132003001323101"></a>

<a id="canonical-2330323221212223-0232011223313131-3203313112322110-1101201213322000-1030103003012032-1313121002013301-3032210031032023-3321032032301201"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2312110223221221-1020112121110002-2333010021012332-1323230113103321-1220113301330122-3012231112320100-1132300001122332-1033200111111210): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3201221001000102-2222223001120213-1013230011110002-0302031102330021-3212111022111001-3221303122101033-0203022232112211-2013030012320011): complete subsection reference.

<a id="canonical-2333032120232213-0131121333122103-0030311133020103-1032203231213333-1110321223111023-2130110123333203-0131120123332201-3330322101210032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0103230222122301-1002201300202333-1002313321013222-0110120230321030-2300333230332111-0330111221032112-1330313102322311-0113223131032333"></a>

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

<a id="canonical-2203032322231200-0210200333232331-2123220130321311-0001201333001030-1333201012212003-0131110123123032-2020123131332231-2022332211123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2202312033323232-2003010011010202-2012303020221133-0301233312031333-2012110210003112-0310113232310131-0210132020211103-0331030013302233"></a>

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

<a id="canonical-2312110223221221-1020112121110002-2333010021012332-1323230113103321-1220113301330122-3012231112320100-1132300001122332-1033200111111210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3220030132302111-2032331222032210-3100111130200102-2312003102201210-0111330011332220-3213203030230221-3031223011112320-0220201110033122"></a>

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

<a id="canonical-1021002122133302-1321013223211102-0212223122311113-0133031230021301-0210002210300021-1100200120130003-2321130110320220-2210032020202122"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1213332322030221-2311133030310230-3012031111221103-0200333321130302-0011322113233022-3301101112201231-0013120101030322-3220012312301213"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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

<a id="canonical-2223100320210001-3132131311221021-0310123013100222-3223323113023003-3302111002123111-0130002020032123-1332232223000100-0103011221212130"></a>

<a id="canonical-1121322323111021-1321131023112223-2203100300221001-3001112123110122-1110131300330322-3031302303030130-2103231331132132-1213203213023030"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-3112130220220032-0112120012213223-1311112300311121-2221300221102301-2120110010130312-1130330321012023-2321031221220113-0000213033100231"></a>

<a id="canonical-2322121012130230-3032310102102313-3211202312032111-2321221021332123-2213103030100310-3111013311031313-3132002122233000-1212203330201031"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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

<a id="canonical-3201221001000102-2222223001120213-1013230011110002-0302031102330021-3212111022111001-3221303122101033-0203022232112211-2013030012320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2130310120001112-3131122031013220-3003023123203301-1032012300132202-0210113203321120-1201330310201221-2010102331033112-3132031000033103)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1032021331313012-2313231002231033-0210013222100000-2212231231102232-3031000132322300-3313330012121210-2121030231303221-2200312332323223"></a>

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

<a id="canonical-0213203231231223-3131302031221111-3033110312202221-2133320310100112-2302223201000300-3101032230200130-3130112012222012-2222000322021111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0232210202031132-1322211132030002-0100022103201131-0222123111132120-2013233002221310-1103102212100330-2003320122221002-0231220120102311)
- Azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0131301211122001-1201121310310001-3213012221332301-3003220033002232-0130201302231221-1020130313001122-1121311031233011-2302022302211220"></a>

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

<a id="canonical-3101002213300002-2031231320211311-2332210130200032-1330110011223221-3320231103003123-3312000102102131-3022130132330202-3231111131110312"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-2221013233300332-2123200100320222-0012202220121331-0223133002103021-2201200032231122-0233203010300311-0202121123333133-1202331003310101"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-0021003321201122-1212311211220333-1323301013310202-3000120223023303-2012132220202100-0030103213002332-1000201120001023-2131201220133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0113133301233020-3032202310003101-2310130003023222-0121132210133120-2330332001021323-2332121131133131-1232202301201002-1302200311233310"></a>

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

<a id="canonical-3012002110021201-2222220333303220-3210202203102132-2230300200012331-1133113012110230-3222323333133331-0323200022133001-1213222302222322"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-0221222302312222-3302233300333320-1100301003331221-0232122321303222-3232331321023203-3213121001333231-0020321123033021-3323103230032203"></a>

#### `azure.not_managed.node_list.interface_list.ethernet_interface.device` property

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

<a id="canonical-0012320211132132-2112003003322211-2102011323130203-2202121020120132-1022330323301213-3111122033131033-1301210222231133-0323212330221113"></a>

<a id="canonical-0232021333323203-0202212110200203-0010200030210022-1302101231130112-1100300132312203-1030101023330333-2311321023012332-3301313130233221"></a>

#### `azure.not_managed.node_list.interface_list.ethernet_interface.mac` property

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

<a id="canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3131201210110002-0033001110312301-2022322300013203-1033221013130323-3201323202112130-2033213222200010-3112321013132010-0310131003201003"></a>

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

<a id="canonical-3011113021300210-0112313102110323-2230102211033003-1300100302313213-2202033113230120-2221331001113121-3031021301301332-3112220131223300"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3011023232310333-0110013021001020-3130033221011022-0223201000313231-0313101103303131-1331203131221101-3311330111113132-2002231011000233): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303): complete subsection reference.

<a id="canonical-3011023232310333-0110013021001020-3130033221011022-0223201000313231-0313101103303131-1331203131221101-3311330111113132-2002231011000233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2132102312202300-0333330010130123-0323122021212200-1123313322112321-2303331321302200-2212021330102111-1003220222200120-3112232020332031"></a>

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

<a id="canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0132232103002010-0032213202112223-2303012001122111-3321322322333103-3003013233331120-2000010102031013-0210013201102002-2230301003322331"></a>

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

<a id="canonical-1103122202311133-2210200113333231-3231231113310001-3323120311313212-1201122020320022-1221111013022002-3301231130002201-3312012232331132"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200): complete subsection reference.

<a id="canonical-2031132223100330-0222302001023201-3311110202022221-1103313322302021-1031100023233303-0003321221202132-3311230201111303-1121130231003011"></a>

<a id="canonical-3101101111103110-2003103003302102-2321002311021031-2302012013103013-0000302033020013-3020313111131000-0321011330031110-3203132021300303"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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

- [stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321): complete subsection reference.

<a id="canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2310011012302030-0031220331121211-1120002130132123-3311232023310021-2330022200100203-1123202331310131-0121121123020200-1203100001220130"></a>

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

<a id="canonical-1003122111102233-2303103030311313-3112033123122332-2010302300100101-2033320010113332-1232130001221222-2110301132032323-1110211130223222"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0321201312321312-0300301100310022-3202200013322121-1032230000320313-2223202200310231-0130311132320030-1103323233102031-3133201110002303): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2020010122100101-0202102213100132-0200200332101011-2230013010301202-3000030131330201-2111110303122300-2321331210310210-0222003203000301): complete subsection reference.

<a id="canonical-0321201312321312-0300301100310022-3202200013322121-1032230000320313-2223202200310231-0130311132320030-1103323233102031-3133201110002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2212022011131030-0000300331112031-3230133223010003-2011201221122131-0232223111323013-3330021231333313-1023000300020020-1000011323301032"></a>

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

<a id="canonical-2111020121132212-0012332311230202-0221210022230201-3021233222202303-2230131301002232-1320222202132230-2200110330103120-0220313203010102"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1332103232230032-3133233220213331-3202000011022033-1203021113322000-3312103033030203-1300102221032211-2222112221213112-0301100300102333"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-2020010122100101-0202102213100132-0200200332101011-2230013010301202-3000030131330201-2111110303122300-2321331210310210-0222003203000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-3311013300300100-3000110211113013-3033101331032030-3113002323212111-2310132331332120-2111311300022210-2111222232230321-1311112133100323"></a>

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

<a id="canonical-2303022322310332-2311033321303010-3123301202000122-0223022130312002-1222300213331222-3311100010000133-1333000112132101-0322231213022300"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2231231103031211-2033223111120000-0102013013222211-3230303121010003-3313221131201103-1010231111022122-0131001020303010-3032020013130332"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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

- [first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3220021011231221-2312322231213133-3210103201210310-2321111123311332-0132132201112311-0022110011120330-0131002323202122-0122000030223221): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1320002312122331-3223131223133320-0113001111333222-2113331210222012-3232021100123232-0302131002123310-1220122001320121-1000010033111010): complete subsection reference.

<a id="canonical-3220021011231221-2312322231213133-3210103201210310-2321111123311332-0132132201112311-0022110011120330-0131002323202122-0122000030223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2020010122100101-0202102213100132-0200200332101011-2230013010301202-3000030131330201-2111110303122300-2321331210310210-0222003203000301)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1311202003321311-0133333312010120-2122213310123201-0220102312302332-1102131120220200-1031320301213100-3021301133120201-1211122210112233"></a>

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

<a id="canonical-1320002312122331-3223131223133320-0113001111333222-2113331210222012-3232021100123232-0302131002123310-1220122001320121-1000010033111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1311123323320122-3030122330032333-3001132301321133-2203203220003011-3011000112313033-0123030021211001-1120030312132212-3302100123133200)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2020010122100101-0202102213100132-0200200332101011-2230013010301202-3000030131330201-2111110303122300-2321331210310210-0222003203000301)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1033213102102103-0321131313103322-2010100131212323-3312230111300000-1301322103113011-1103230333030113-0123311201313101-1323303131322301"></a>

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

<a id="canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3223023330001233-3123213232222310-1101030121131213-2110031023113332-2121303300123212-1000331221131331-1233120331223033-1313200102300121"></a>

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

<a id="canonical-2002120130333303-2020202102032033-3210322100023323-3322111133200322-1132312210321212-0132302033103333-3233211111310130-1101213013210130"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3312030101311210-2323122312313321-0122330000210103-2222332211120301-3313130200030020-1120020032302133-0200033330332003-1110012130322212): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0011011120313133-1333111230311031-0323103322012122-1303032132100213-2020232223031230-2303101001231001-0101022100001213-1311120200213032): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023022103103320-3322311130031010-3030010313233022-0321130211102330-2032102011111101-2111121033131333-3011201231022100-3030022331032132): complete subsection reference.

<a id="canonical-1011333020302333-2203202030131201-3220230021312301-2130110123203322-3310103003011132-1211332323022020-0312330303331202-2123132332211223"></a>

<a id="canonical-2011200002320031-1300031133320012-3003033110121331-3322002123220003-3200023202032233-3130223033212311-0130223123013111-1013313132000130"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2331012011003033-3322211023320203-0120010023010231-1330221022012330-1300233211300211-0123130121303322-1333231123002010-0312003000013102): complete subsection reference.

<a id="canonical-3312030101311210-2323122312313321-0122330000210103-2222332211120301-3313130200030020-1120020032302133-0200033330332003-1110012130322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3321301330033020-2220002032332203-2110131002112130-0211121213020302-1103002300301311-3003300332020212-1330030301130102-3001130211231322"></a>

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

<a id="canonical-0011011120313133-1333111230311031-0323103322012122-1303032132100213-2020232223031230-2303101001231001-0101022100001213-1311120200213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0210220202011333-2002320211003223-1113002321322131-1310110103333202-2022032322301032-0101313122200011-0230223311032221-0122112001201101"></a>

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

<a id="canonical-1023022103103320-3322311130031010-3030010313233022-0321130211102330-2032102011111101-2111121033131333-3011201231022100-3030022331032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-3212101031230000-1331223031233311-1123123122330231-3233203333330131-0200303110121123-3020222321331313-2012103220322031-1211333330112122"></a>

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

<a id="canonical-2233211020100103-1131211031023223-2200030232131022-0201331203023212-1023203020323213-2220301003331323-0001013133322213-1012132331023303"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-3202230032301012-3223023310221000-0032223320100232-2132321033230003-1312301110230010-3011023123001103-2011221231122203-3333300311312203"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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

<a id="canonical-0133033103100102-2320200232032100-2202220332301003-2010132112310312-2220012003302111-0100220332312032-0002131022031103-3022132201222101"></a>

<a id="canonical-2030101030331203-3122120012122231-1312200130231121-3113322313121021-2032233001321331-1322301012022230-3123333301012123-1300333323212013"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3022100202330331-2220002212133032-0233200313120023-2200003323100111-0003120102302210-1321110331133013-1323003020122311-2000003220013033): complete subsection reference.

<a id="canonical-3022100202330331-2220002212133032-0233200313120023-2200003323100111-0003120102302210-1321110331133013-1323003020122311-2000003220013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023022103103320-3322311130031010-3030010313233022-0321130211102330-2032102011111101-2111121033131333-3011201231022100-3030022331032132)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0323103233111223-2223132120302311-2112031220110022-2021330213132022-0001122022202333-3103021001033123-0102102011113133-1312301310203112"></a>

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

<a id="canonical-0323230101230331-3222102231313003-1202212002322130-2031132132121302-2123030303000102-3202002020310333-3213223001103131-1221300202000310"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-1100320031331113-2233230021203231-0202311013120031-3112120232320100-0000000332221102-1323133331120201-3032302312221310-0132331123210303"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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

<a id="canonical-1012332021301120-0101120212320002-1130223011211332-3103022233310122-2031303332033311-3003001021121121-1220320222333003-0231313202012303"></a>

<a id="canonical-1022333113233332-1320303121013231-3230313011131003-2010022103203221-1200200130220323-2321232320030011-1133000221311100-1111120330201012"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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

<a id="canonical-2331012011003033-3322211023320203-0120010023010231-1330221022012330-1300233211300211-0123130121303322-1333231123002010-0312003000013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3302202020013110-0210303221321123-0202110311223303-0311231303332112-1100213303200212-1232210231202030-0030122301122222-1131113321133300)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1023113130310002-2033002032223300-3133210301211332-2323302122321122-1230221001230000-2023113313133031-1020201321312332-2221110021102303)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1331003103202301-0133030233110213-2031003230123113-1321001200230202-2203301311120223-0313122202002013-1322333020112303-1012031003022321)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2033112122211221-0111020032311330-0200300331230132-2112311311223321-2201102021102101-0230333220022100-1112321212222301-3230113202203221"></a>

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

<a id="canonical-0022010330321112-2102213201213010-3130031100013031-3330221130011121-2323323020030231-2230322130302121-0223231222321032-2313232202203222"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-1333322311323223-3201222330333331-2123022301102131-1003120210233031-2323103133000303-1221100230112031-0210322232332121-3102302201310131"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1131313313000333-3211233010002003-1230232300100203-1012021222020133-2113000210302220-2302123323101332-2201302021102331-0033222003010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.monitor

<a id="canonical-0020322021120102-0033311310231203-2323211201330301-2311012123230100-3021313130201313-3031310130232223-0221122100331312-0033001132021122"></a>

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

<a id="canonical-1300100302021332-3303213332003322-2313211100300103-3120320132323010-0013302303111322-2030223032101001-3320032330122321-0033300222133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-2200011330131120-0301112330113133-2313012030310122-3320131101130102-1023301130330132-1130002212310201-3230011130200222-3210012320313323"></a>

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

<a id="canonical-1133322231300022-2223132221112121-3213201012123030-1013321032023312-3113311100202210-0111222231130203-2201332102232123-0133010110112013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.network_option

<a id="canonical-1002223132023002-0200122201333022-0302211320333003-1111010003132011-1002222323102203-3122010022101232-2131320320322300-2233200331120032"></a>

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

<a id="canonical-3302231032321320-3202121002231122-0112133232312020-3300213300033311-1233220223313102-0110131201201131-1010002211112312-2100013331010232"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3323313030221331-1030331221021333-3020121132303201-0203101222200123-2111122131201003-1311112121300220-2123231021301210-2030210000311012): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1032030011332310-1231130132131130-2010332302010220-0311203220132011-1300222123020030-1212313012013332-1031103313120322-2311313021023121): complete subsection reference.

<a id="canonical-3323313030221331-1030331221021333-3020121132303201-0203101222200123-2111122131201003-1311112121300220-2123231021301210-2030210000311012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1133322231300022-2223132221112121-3213201012123030-1013321032023312-3113311100202210-0111222231130203-2201332102232123-0133010110112013)
- Azure.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2321121101011032-3033033133130331-3231100201123322-0020011030132201-0202201121201120-2030030201130100-3113023113000203-1022223302021331"></a>

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

<a id="canonical-1032030011332310-1231130132131130-2010332302010220-0311203220132011-1300222123020030-1212313012013332-1031103313120322-2311313021023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1133322231300022-2223132221112121-3213201012123030-1013321032023312-3113311100202210-0111222231130203-2201332102232123-0133010110112013)
- Azure.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2103100010203210-1313303222132101-1320022002101001-1233211100320101-2323030103112232-3310232010011202-3201231320310111-3321102310103233"></a>

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

<a id="canonical-1303330031222030-2322131123213211-0211022223123330-2200101233221013-0301123113211110-1033312333022032-3022113122303131-0100311221221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0232200323130131-3212233103020232-3123122003033020-0330333323122220-0332103121231011-2000231101200332-1231221003330230-1032222233313100"></a>

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

<a id="canonical-1230222333021311-0203030103301232-3312231020322303-2231322120222112-1012313323303130-0111002013213333-0110101201133332-2221202233202021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2302231333200133-3133002121011130-0220001100212210-3101001102200031-2101232220200301-1030232210220121-2120210313120011-0323302221113030"></a>

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

<a id="canonical-3312300103000112-1003121230323300-1010211000031032-3212003011301121-0201321310320333-2113330021330013-2113210102132233-3310222230002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-0032313303021010-3203300312000313-1201030122121230-0203311103122102-2221022001012010-1221101001020201-2301203312120112-0201102202213310"></a>

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

<a id="canonical-1330131100210003-2302112112121101-0033313031321102-1101002100230302-1031110230002221-3101332131213300-1030000102331322-3032012100030323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0312330033210132-3331031302020032-3202231031321213-1201232332002101-1333331312120002-1222323011330230-3022012221222212-0123200213120111"></a>

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

<a id="canonical-2112223023131102-1132132230021333-3133002200110113-3330231020030223-1203102203223221-0132103100330111-0300223123020201-3203321302130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.static_ip

<a id="canonical-3330210011033021-2200132300112203-3231120211130223-2322100232211132-1323111312103111-3222203001132320-2221130001102010-1132223131223130"></a>

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

<a id="canonical-2303322123103320-2203321313233303-3333003313210320-1233013003302212-2120122020301010-2002202300012302-0131233313201130-1100012031203233"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1011223110201031-3110201103030032-3112323000303302-2123003332023320-2000202330110012-2310032320231303-0001220123021331-3103231223323301"></a>

#### `azure.not_managed.node_list.interface_list.static_ip.default_gw` property

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

<a id="canonical-2310002330101021-3133113002310211-1120010120111023-3021302320132030-2013310120133112-0323003110301012-1321202213330023-3310330211112323"></a>

<a id="canonical-2223212033223110-3110302212301023-0321132211300132-0011202030230332-3010233100103231-2013112210202010-3111023010230203-1001210220130133"></a>

#### `azure.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2122030011322101-0011323321203131-3122100232133011-0101020322122002-0112030101023302-0330003111301212-2322320100220021-0203023320220302"></a>

<a id="canonical-3122322222003302-2312333110033110-2133210330211330-3333300021230331-2101131311022011-1003130310001310-2003113100010212-2001003312232331"></a>

#### `azure.not_managed.node_list.interface_list.static_ip.ip_address` property

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

<a id="canonical-1211120233231203-0110323123233100-2010300030212231-2030001200102110-0212003021130111-1213230011333211-2011220310200020-3211101133303300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0323130033213101-1102003203230102-2103332102300223-1221311102023021-1310313230110332-1211311113023333-0202233100003110-2300123030213222"></a>

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

<a id="canonical-1331301031222132-0021012231300332-2010312211310301-0321221303130302-0110333112102021-1233133122100332-3332233012313232-3301000312211231"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2000222223212203-1111220022023023-0200020320331300-0211133321320300-0332010013111001-3322100131312302-2333320202011130-3102203223032023): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0100101201020203-0100303110013123-3333121303322331-1221322301322110-3230003220033333-0300100332022323-0223113122210022-0210002333331200): complete subsection reference.

<a id="canonical-2000222223212203-1111220022023023-0200020320331300-0211133321320300-0332010013111001-3322100131312302-2333320202011130-3102203223032023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1211120233231203-0110323123233100-2010300030212231-2030001200102110-0212003021130111-1213230011333211-2011220310200020-3211101133303300)
- Azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-3130320000103130-3222313012022123-2300320310031022-0300003303133300-2311333122133223-1021132310202102-1130102212003223-0012203220330311"></a>

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

<a id="canonical-1102332322020130-2010310303220120-2200020113301331-2213302230321333-3121020201212301-3123120210033131-3231302333323202-3211111223312223"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-0131323031200000-2120303022113323-2130212101301033-0133332302000101-2211011102301301-2333111320332131-0302213111132012-3131210011200302"></a>

#### `azure.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-0100101201020203-0100303110013123-3333121303322331-1221322301322110-3230003220033333-0300100332022323-0223113122210022-0210002333331200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1211120233231203-0110323123233100-2010300030212231-2030001200102110-0212003021130111-1213230011333211-2011220310200020-3211101133303300)
- Azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3010301312010311-0332322100112320-2102211321200023-3311211111331303-0310021123103013-2002311131002220-1201310333103323-3131233310312010"></a>

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

<a id="canonical-3101210312020213-0120112101021231-1133300203322021-3131020313020213-2322113322003320-3201122330002222-3221221010223230-0222311130111031"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-0121312301021230-3022102212222122-1311212131122010-2211031033331022-3112233323233020-1312022123302112-1210122012232112-2132121022033100"></a>

#### `azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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

<a id="canonical-0221323231111210-1311230300220131-0212000212300210-0223130030212023-1200130111202131-0001123201211223-3302321100201113-2333330132200023"></a>

<a id="canonical-0130201222013111-3212112322331330-3230003311323123-0100011123133130-0201103331003320-2233313013203011-0303210332210030-1122013023000010"></a>

#### `azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2100302123020232-3003211103032120-3211310000113023-3331001301101233-1010303322300223-3020330233011213-0330200031123133-1321112210123120"></a>

<a id="canonical-0000133231022100-3330203320300003-0313112132202012-0031010322220101-3030113131021310-3210002110312030-0332022222033220-2122330221332200"></a>

#### `azure.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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

<a id="canonical-3302322001010313-1200112121331130-2022203212131331-3303211013223201-3212310012133301-0311212203213311-2202110222001300-0120102320112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2030111322112020-3332023121223220-1110001210322333-0023013121333330-0220202000203020-2233311312123232-1010221200332332-2022322000003033)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2100030312311201-3121222012313110-1013011110011020-2210322003200020-3000320130023013-0111113313233202-1213221213111312-1231332213100323)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-3110310023102310-3303131013221013-2030030013013300-2023033110233113-0210333201301023-0311103121013123-3031132132310012-1102212323013110)
- Azure.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2003120320301011-2120332301333033-1022220011330003-3320010121211130-0111120313021212-3210212010221132-2223122212222112-0221202122213113"></a>

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

<a id="canonical-3101233230031301-1102223213111030-1230013211003032-1023103022130002-3213213231211032-1210231032210332-0022000131312332-3230233222322132"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1230332030212320-1220301313030233-1201333211000022-0302211023112302-1233030203000333-1223303030322221-0131212200012333-1303010023102011"></a>

#### `azure.not_managed.node_list.interface_list.vlan_interface.device` property

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

<a id="canonical-0133113020121302-3330323030031003-0123333032032012-0303323032003120-2213122320222100-1022231331131212-1131023320322110-2112121202200333"></a>
