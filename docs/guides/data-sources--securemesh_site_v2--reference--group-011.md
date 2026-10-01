---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3113020001113322-2230023132331223-2101120202322111-3123031032231332-0133111313011122-3230210101302312-2100031031032131-0200312310132310"></a>

## Next pages — configured_list / 103123221033 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1123011102231212-3121032000301002-2220333102303322-3010301310230323-1300210233102131-3303120112113102-1012230232033102-2331132211111102)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103233302100232-0010222102121310-2103311031221220-3021213100233302-1123311013003102-2300303032002122-0122003110220210-0311031302322023"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 033000230032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1123011102231212-3121032000301002-2220333102303322-3010301310230323-1300210233102131-3303120112113102-1012230232033102-2331132211111102)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-0323131223100132-1233332201302031-0202130213321332-2100221001101202-3312322120231203-3322032000322321-2301012202210112-2101120021101231"></a>

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

<a id="canonical-2133123120231010-3002230021100123-0113101223132010-2221102032221021-1202211010202111-2201010311223011-3213012220131133-0111303320232222"></a>

## Direct properties — local_dns / 033000230032 / 3

<a id="canonical-0132132003031122-1100013312213203-3033201300321212-1200120120323010-3212123223301330-2103213101333110-3021303013310303-1122033113031323"></a>

<a id="canonical-2222010301212220-2103201303110302-3002220301200332-0132032010131102-1210031321100031-3122211120121211-2223310101130222-0031022203320200"></a>

## configured_address property — local_dns / 033000230032 / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1231102013211013-1022023332230011-2330201012032331-3223232112010320-1010032221023210-2121110323220333-2113311123032331-1133012033102321): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2012323122101122-2201301331213232-1310233002103330-0221102300311212-1110202123012020-1212130130112100-1312010132121210-3312030133230102): complete subsection reference.

<a id="canonical-0231213201032310-3101013330113132-0220103130033311-2022330010303321-0023110003131110-2113201103112132-2003301000211030-3013200000301321"></a>

## Next pages — local_dns / 033000230032 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1231102013211013-1022023332230011-2330201012032331-3223232112010320-1010032221023210-2121110323220333-2113311123032331-1133012033102321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2012323122101122-2201301331213232-1310233002103330-0221102300311212-1110202123012020-1212130130112100-1312010132121210-3312030133230102)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1123011102231212-3121032000301002-2220333102303322-3010301310230323-1300210233102131-3303120112113102-1012230232033102-2331132211111102)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1231102013211013-1022023332230011-2330201012032331-3223232112010320-1010032221023210-2121110323220333-2113311123032331-1133012033102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001011112012222-0010123312220131-0322021212002130-2110322301031312-3333311131331010-0031311311301331-2202303101310220-0001223233113302"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 330003311331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1123011102231212-3121032000301002-2220333102303322-3010301310230323-1300210233102131-3303120112113102-1012230232033102-2331132211111102)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1113032112100011-3203312200113100-3133112331133301-3210011323102213-2301231303200321-3213010323100331-2322213313123022-2131032131131031"></a>

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

<a id="canonical-3032302300131233-0013230111222301-3101003013213330-3230030220022003-1002212331331222-0031131301230232-3201113330020010-0101231202012302"></a>

## Direct properties — first_address / 330003311331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012201312021013-1021201122032233-1113220211012330-1330330231320322-2221200101220310-2333221022310220-1201010111322030-2232110322000333"></a>

## Next pages — first_address / 330003311331 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2012323122101122-2201301331213232-1310233002103330-0221102300311212-1110202123012020-1212130130112100-1312010132121210-3312030133230102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210302302330332-2000331133231122-2111121001133103-0312310013121133-0322122133103301-3233233100122011-3322012310003100-2123232233213110"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 223232111121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1123011102231212-3121032000301002-2220333102303322-3010301310230323-1300210233102131-3303120112113102-1012230232033102-2331132211111102)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2121022221231111-0210311230011213-0312032013312231-3031301123103102-3123122130300232-3211233313112123-3110012123301002-0213320010132220"></a>

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

<a id="canonical-0101231100122030-2221021133112203-2211331133002202-0030333131011322-1203333333302222-2332133201122021-3010021021132120-1032003303133133"></a>

## Direct properties — last_address / 223232111121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103330321121302-2103331103001012-3112230222032123-2313313223323111-0311023031130201-0002003032323322-0213100231300211-0013222001133233"></a>

## Next pages — last_address / 223232111121 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120222200332121-3310000323330001-0002210002121102-3003223013222231-1120312231103130-2313123102221320-0022021331033122-3101100233002003"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 222123033113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3322230122211300-2333231113030112-3112330330021003-0320103112302113-2021030213211112-1332201133223213-1121032211200101-2001100021222333"></a>

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

<a id="canonical-3212230031301230-2200132322103222-2202332232323131-2213023302021111-2221301200300311-0023320203031331-0020023230111232-0312222122033131"></a>

## Direct properties — stateful / 222123033113 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0202303331213103-2131232330301012-3311223123031333-1000010020331132-0122332231223023-2230112112123312-0021332021320201-1232111020010101): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0223232322233010-3100222321202000-3013122022032231-3011112122211311-2310012033032110-0203011011132100-1322203311113033-2303311301130300): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3302031331013120-1333323301022332-3102100102222230-2332330101323021-0322303223313330-3333122303022120-1201332013310313-2100121211030111): complete subsection reference.

<a id="canonical-0312120230132121-0310022003301003-3221122130133102-3020021000300230-2033011301033331-0001113021210100-2020301031102102-2332121022133033"></a>

<a id="canonical-0111232303321331-2000130002332333-2120010120033001-0123011210223231-1323122132010000-0201223200000312-2010301022322113-2102022000012313"></a>

## fixed_ip_map property — stateful / 222123033113 / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3030200111012320-2123020330102301-1330111332320101-2031321123203113-3210320102200222-0031300211020102-1113210311200320-3033223101133020): complete subsection reference.

<a id="canonical-0012232220111003-3312333333122010-2310132221311032-3210123033123032-1311321112032232-2133303203233133-0122112310111303-3332322323023112"></a>

## Next pages — stateful / 222123033113 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0202303331213103-2131232330301012-3311223123031333-1000010020331132-0122332231223023-2230112112123312-0021332021320201-1232111020010101)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0223232322233010-3100222321202000-3013122022032231-3011112122211311-2310012033032110-0203011011132100-1322203311113033-2303311301130300)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3302031331013120-1333323301022332-3102100102222230-2332330101323021-0322303223313330-3333122303022120-1201332013310313-2100121211030111)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3030200111012320-2123020330102301-1330111332320101-2031321123203113-3210320102200222-0031300211020102-1113210311200320-3033223101133020)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0202303331213103-2131232330301012-3311223123031333-1000010020331132-0122332231223023-2230112112123312-0021332021320201-1232111020010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201302120320231-0201002031010232-1112023230313130-3230312020101113-1202210211333113-2332320233110230-3120332322222301-2112332003130130"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 132221102322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1002121031112102-1031333303300320-3130131123320021-1232012113023230-3000201302133113-3211333332102302-3103023232333031-3031221333121322"></a>

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

<a id="canonical-0022013311232300-2121210121203010-3121011113030233-3003120231302202-1010300220200132-3011110220323233-3301320003113013-3220231133023302"></a>

## Direct properties — automatic_from_end / 132221102322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213120130331233-1103000033023203-2331032121220222-0021131323012333-2302202103023231-0220212300320133-2221022323120322-3211001102111312"></a>

## Next pages — automatic_from_end / 132221102322 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0223232322233010-3100222321202000-3013122022032231-3011112122211311-2310012033032110-0203011011132100-1322203311113033-2303311301130300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112112123002233-1111233031303233-0210122110200102-0012302210332323-1012033232001121-2213313132013022-1333211021100132-2013231210300332"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 322111212020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2110020211112122-2223123121213002-0213331131222223-2023123332223320-3100031021001303-3032003233302321-0031112203303112-1201312121110000"></a>

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

<a id="canonical-0122230313121012-2222323211212022-0213301122102031-0202232233002013-3201112030233100-2120231230003231-1103023030122333-1112000310302132"></a>

## Direct properties — automatic_from_start / 322111212020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111030300301011-2320320300330112-0221121200302311-0233222012020231-3201232310221032-3203103323033130-0302003323221331-1203000310030013"></a>

## Next pages — automatic_from_start / 322111212020 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3302031331013120-1333323301022332-3102100102222230-2332330101323021-0322303223313330-3333122303022120-1201332013310313-2100121211030111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323022131333021-1011321232330030-0110323033110320-0213331332312313-2332113221231110-1303111102111123-1210123133103310-2231302112203031"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 233331201323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-3321301232131232-1300213023131331-0301232221130301-0122312221202311-0022012010333033-0320332313012201-2113113313101221-2020010001230010"></a>

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

<a id="canonical-3033230021112121-0211113102201300-1122302013011201-2201012022320033-3233100121023131-0330211031220313-0333031222102120-2032113113233033"></a>

## Direct properties — dhcp_networks / 233331201323 / 3

<a id="canonical-2021133203201101-1303003102232023-0233333112130302-1113100101111323-3113201213302120-3021213223100123-2232222123333121-1233301120033211"></a>

<a id="canonical-2023333111211222-0311230121031221-2322030230211000-3232203231332022-2003131131221213-2111221230021302-2230002320123031-0301311300311130"></a>

## network_prefix property — dhcp_networks / 233331201323 / 4

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

<a id="canonical-0032212132133233-3121312000021220-1110201000312130-2202011212121333-0330120303003301-3312300121310022-2001320330020332-0002133120010111"></a>

<a id="canonical-0102021300210132-2201313133201312-0200332332203000-1120221323132132-3323210100131320-3323223013020000-2213101023330212-0202123011330032"></a>

## pool_settings property — dhcp_networks / 233331201323 / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3032031320322210-0322031100313310-3211130211232131-3301030321102023-1220112100023201-3112313120232333-2321112233022223-3122202310323212): complete subsection reference.

<a id="canonical-2332323223022111-1012010323003323-0120112020331213-2122033131223101-1010032102130302-0313021111130112-1322211222101323-0111032003221112"></a>

## Next pages — dhcp_networks / 233331201323 / 6

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3032031320322210-0322031100313310-3211130211232131-3301030321102023-1220112100023201-3112313120232333-2321112233022223-3122202310323212)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3032031320322210-0322031100313310-3211130211232131-3301030321102023-1220112100023201-3112313120232333-2321112233022223-3122202310323212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213300130323233-1000310331110101-1103113231203222-2033001021000010-0112130310132202-3111020000130302-0033233100230331-3221131213101230"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 130312330102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3302031331013120-1333323301022332-3102100102222230-2332330101323021-0322303223313330-3333122303022120-1201332013310313-2100121211030111)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0012322323230000-0123120003021312-1032101120332213-2011312223322223-3220201021303332-3032021223112213-1011203300233232-0230030002302122"></a>

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

<a id="canonical-3030311012021130-2300011210002200-2333231100230031-0022023231003110-0310023100320320-0023313121013112-3133312112022020-2011220232332013"></a>

## Direct properties — pools / 130312330102 / 3

<a id="canonical-2213212013000121-1310311031131330-3111333012233221-2213321113033210-3321033200231230-1033203111031201-0031331323010221-0111032313301022"></a>

<a id="canonical-1213222023121032-2330002122033101-1021102010120122-1010121301121332-3021231300031103-1012113321130303-0300000230123111-2133212210121201"></a>

## end_ip property — pools / 130312330102 / 4

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

<a id="canonical-0302232000110133-1220103122010232-1013312231031002-3100233021300011-3130002003312110-1102223222230010-3210210030333001-1121231032102110"></a>

<a id="canonical-3321232111303100-2103030033320112-0331121131133100-3120011133231221-0333121200130320-3210133120012331-3330312223033132-2012231200303032"></a>

## start_ip property — pools / 130312330102 / 5

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

<a id="canonical-3033003213202101-2312232023332022-2011311203232230-0310013211030122-3102210100001222-2101110333303112-2331000022103013-2010112000233011"></a>

## Next pages — pools / 130312330102 / 6

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3302031331013120-1333323301022332-3102100102222230-2332330101323021-0322303223313330-3333122303022120-1201332013310313-2100121211030111)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3030200111012320-2123020330102301-1330111332320101-2031321123203113-3210320102200222-0031300211020102-1113210311200320-3033223101133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232122330111031-1313312122211210-3103301323030101-2202220013002223-0130120321223221-0200021321231000-1232222303303223-0003122013220002"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 023330231121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3202131122022203-3112100320221100-0230110020003120-1220000311312203-3222303202011123-0220223310011103-0233103231103021-0102022212021123)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1203311112221030-2332020231232203-1221101002112312-0331030100302232-3011020210020020-3333113222321310-0212130032120330-0131330103120301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1022113003011212-0222021030120203-1021210010220033-0030020222232213-3220300130333100-0001211323123022-3101020123133200-1311323030310313"></a>

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

<a id="canonical-3322010131002321-2030003201323102-0310132123331211-2321301332232213-1020302211311011-1011230133232102-2310110032133103-1110000322030301"></a>

## Direct properties — interface_ip_map / 023330231121 / 3

<a id="canonical-3000101221330110-1020000230221011-0203011102333202-1131311001132101-2031211021032022-2220010130211231-1001130221100012-3110301113220312"></a>

<a id="canonical-2311312303100200-0312000030010321-1120331330113320-3233011202112220-3131221221202113-1330220321112100-1023333022212123-3123022011121203"></a>

## interface_ip_map property — interface_ip_map / 023330231121 / 4

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

<a id="canonical-1101100221200110-0002102300320022-0213310233201001-3002221123012300-3332212122230300-2201213103130012-0222310300302031-2132300001332201"></a>

## Next pages — interface_ip_map / 023330231121 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0333210002313113-1203111313030202-0323301203301323-2222312231030111-2133030020233013-0112021310212200-1032122001122122-3032111303200133)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2102111310322012-3021221112312110-1212311223121300-2201000020100210-0033302021033030-2322220013132133-1321112230320010-2111021130301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130002122013012-1132221003022013-2323000202102132-3312132303331022-0033020330020321-2321200231013111-0023210322000320-0120112113200000"></a>

## kvm.not_managed.node_list.interface_list.monitor — monitor / 231123122311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.monitor

<a id="canonical-1003322302123212-1313320300133030-3122002111020230-2313231031011102-3320310312312010-0200103003002231-3022113301033103-2022200233000023"></a>

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

<a id="canonical-0323300123231123-0323313300213133-2131203201123023-0331000110203113-2311232033213321-1013022301221202-2213113202102300-1303231232121013"></a>

## Direct properties — monitor / 231123122311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231131130002112-0123123301131330-3023301101002120-1000130120211201-1232311113321333-0133131121233302-1010301033023301-0110330310010331"></a>

## Next pages — monitor / 231123122311 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3222230312223023-0223223322113133-0131031131132200-3123333030100132-3121112233131212-0121020220332210-2313222012213022-3023331300200303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202212203323022-2002233130332332-3021310310330301-0320133303033213-0223200111302121-1320210303321210-1022030323101032-0302113102222311"></a>

## kvm.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 223313120221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1230221000320323-1111031201203131-0023301102102121-1023010232220002-1003033330333133-1322313333002033-3110132320323010-2111232021211222"></a>

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

<a id="canonical-2220202120022020-0011121103002300-3322023132221231-3223010212022301-3010120203020310-1113133222101323-2000112010311303-3010213323103022"></a>

## Direct properties — monitor_disabled / 223313120221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033333302302010-1011331332102020-1331330222312032-3122332012231031-0230330200331021-2113001113001203-2312321333313121-0302303230231010"></a>

## Next pages — monitor_disabled / 223313120221 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2202101011332233-0211310203120102-0310332311211322-2301312102000032-3201123322013331-0223330120212130-0202203301130330-3302112302000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210012330303121-1303113310021011-0130311122201332-2332230203302132-3011110211120133-2033101322100123-2012020303130031-3323132130333210"></a>

## kvm.not_managed.node_list.interface_list.network_option — network_option / 001012122132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.network_option

<a id="canonical-1001123111002022-3213133212320033-1121002302031023-2331301202333310-3110212121021210-0031033200320020-2003101331301120-2113313321131030"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2112001302033233-0023010030313233-2323021000123203-2203323220131112-3130023032032321-3101113211123113-0021322103113300-1132013302231231"></a>

## Direct properties — network_option / 001012122132 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3011021021103223-0222010222103300-3022302220133021-0200310230110232-3131102113212302-2313131203230100-0331212331332010-0032333021330213): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2221010313100223-3330322200200231-1232112012312200-0113300123102223-1231011103331103-2203233200311131-0001313330223203-1230220133101211): complete subsection reference.

<a id="canonical-1021221112010320-2112322113120002-3101021113202332-1001030321202022-1303233113030232-3013203323202100-3001200102112333-2332231321310302"></a>

## Next pages — network_option / 001012122132 / 4

- [kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3011021021103223-0222010222103300-3022302220133021-0200310230110232-3131102113212302-2313131203230100-0331212331332010-0032333021330213)
- [kvm.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2221010313100223-3330322200200231-1232112012312200-0113300123102223-1231011103331103-2203233200311131-0001313330223203-1230220133101211)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3011021021103223-0222010222103300-3022302220133021-0200310230110232-3131102113212302-2313131203230100-0331212331332010-0032333021330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112213010030112-3213210302112202-3110001300032130-3312021302021010-3200220332001030-0313323111311200-0021032232130321-2211013323202021"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 023320123313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2202101011332233-0211310203120102-0310332311211322-2301312102000032-3201123322013331-0223330120212130-0202203301130330-3302112302000202)
- kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2122220330311323-3322101132232320-2002311132332032-0013212331321130-3302230320031323-3020222333200302-2133300313022221-3031210013222123"></a>

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

<a id="canonical-0002121023233323-3322323332132200-0123221311030233-2030101111121322-1132330322201231-2231220312010031-2003333212311303-3032000232000330"></a>

## Direct properties — site_local_inside_network / 023320123313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301320010220011-0201103022330233-1323222312301013-2222021023131202-2103222201102211-2323130103233330-0310013301202013-2310201013311130"></a>

## Next pages — site_local_inside_network / 023320123313 / 4

- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2202101011332233-0211310203120102-0310332311211322-2301312102000032-3201123322013331-0223330120212130-0202203301130330-3302112302000202)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2221010313100223-3330322200200231-1232112012312200-0113300123102223-1231011103331103-2203233200311131-0001313330223203-1230220133101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232003232223013-1011122213332103-2013221300220300-0021121013320232-0301111002333130-2112121032230201-1210311130231000-2220021322303030"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 231322232200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2202101011332233-0211310203120102-0310332311211322-2301312102000032-3201123322013331-0223330120212130-0202203301130330-3302112302000202)
- kvm.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0312232102211210-3131202131222023-2221001101232032-3011002131310102-3031002021032030-0111211132033332-1012000211312313-3020012231013210"></a>

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

<a id="canonical-0223123112203212-3002331003331010-2101231110221001-3301131133101322-2223312323331310-1210332122232233-0231100230211203-0010111121131331"></a>

## Direct properties — site_local_network / 231322232200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103131003003120-0112001012101101-1120113002103131-3012010321030112-1223230120202222-1310132132320311-2302030331323101-1223332022211031"></a>

## Next pages — site_local_network / 231322232200 / 4

- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2202101011332233-0211310203120102-0310332311211322-2301312102000032-3201123322013331-0223330120212130-0202203301130330-3302112302000202)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2020300023201321-0212200330301312-0001330013203303-1233112012310230-3131130323211230-2232132200003103-1100000101210322-0230110312031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231322210031323-0300321001331121-2210120233120002-2032131321230011-3033213112122103-2331302010022311-2122223133212111-0113120321331222"></a>

## kvm.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 020130321000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-1313333001201211-0002101311332330-2323210221122232-2111033303231202-0231131330313221-2320020310233210-1310301221310101-0113022210203330"></a>

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

<a id="canonical-0023120030012132-2012331103302021-3122100101320100-0001113130310000-3013222221122100-0310330113130321-3001130003102223-0233212020311322"></a>

## Direct properties — no_ipv4_address / 020130321000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221211003120011-1001021223222200-1002201123021023-2011012122100130-3211220013332122-3020330132000132-1102300302121001-2303313112022303"></a>

## Next pages — no_ipv4_address / 020130321000 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1033023130312123-2122223112031113-0233203222023211-3012331131320103-3313303303302001-0231220021031303-2303023102002021-1313130030113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112133320132332-2113321121103101-3133120102213131-3313223020213222-3303021122212201-3222320211322022-2102330112312333-0030012333221112"></a>

## kvm.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 100201020230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3211312321333321-1031032110220313-2220102113223223-1231032000001302-2232111122200130-3320113110202213-2120112200200023-2010222011303010"></a>

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

<a id="canonical-1310333102101310-1123330203313223-2011210221003321-0113033312000002-3113310301012312-0313232332202131-2113202003321301-3202331112302222"></a>

## Direct properties — no_ipv6_address / 100201020230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000202132333200-3213312120121312-2023231030130111-0320233201110101-1131312333133003-0302233111301000-3000122100122220-3001030232112002"></a>

## Next pages — no_ipv6_address / 100201020230 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1002030112212311-2102320320212210-1323311321303023-2223122003033113-2000332331110033-0100311023223022-2032212211121212-2002333232223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110112033202223-2321100031001213-3232000101031132-1232322100331312-0330301133131002-1021212130202222-1230202022012122-3221220130122231"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 033033001112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3212211100300120-1321231303330303-1232121032031230-3132002103131010-1223032212100221-3211010200010133-0032021111132130-1020131302013013"></a>

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

<a id="canonical-0013331211301211-0202321131022233-3010330333233112-0203331110222232-1020130333322331-3121111002113203-3233103210230012-1300132331111333"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 033033001112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200031323003132-1331103032213313-2021020233011220-0222132002230110-2303032133310222-1332202121022323-1332121223230313-2000001313122100"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 033033001112 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1002100132212213-2110331102311330-2300320012132212-2221202200122131-3302223001111032-0102312000321223-0230110321200132-0021232010122021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102121023312010-0202100222013313-0022100302132310-2222113103230311-0103123021202010-3100033232120322-0230220112301210-3010211330022213"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 223222322132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3222220133223111-3031323303203231-1221313133231123-3023322003001230-0033200123133010-2220100001202213-3112310211101303-0300222033313220"></a>

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

<a id="canonical-0013032300213131-3120010130321233-2232003002300110-2333120202231002-1220213221123123-0000232312111013-0013331333321030-3222133022022302"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 223222322132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312301303110303-0013111022120013-3302213012223101-1203223323112103-3323200111013000-0122022200100222-3033101213232310-3103023211021113"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 223222322132 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2222212320221320-1302020202113022-3200310330133232-1000121020033121-3113100221022300-1213103200211322-3111323002302222-2331220101002331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331011203320231-2212223021332211-1312010002031302-2222300221012201-0210022023032201-2332310331100031-3210121032221221-0112333002200320"></a>

## kvm.not_managed.node_list.interface_list.static_ip — static_ip / 321303031302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.static_ip

<a id="canonical-0011100020313103-3201331210300011-2120112201221113-1200311311300122-1212001202333222-1101203213103001-0331010221201101-1020012130320121"></a>

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

<a id="canonical-3320003031132023-0021212032120100-0123331333200310-1123133023132320-2102330310203120-2230203332232100-2333210132333231-0210111312203230"></a>

## Direct properties — static_ip / 321303031302 / 3

<a id="canonical-2202330113112030-0011301211233112-2112003131120011-3202301221212311-0121213200231102-3030213133222323-3331111013120331-2200211003212012"></a>

<a id="canonical-3223232013333022-1132313022101201-2301303233200320-1102221333302103-3232310022021013-3101103023201022-3210311103011123-2320100323311000"></a>

## default_gw property — static_ip / 321303031302 / 4

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

<a id="canonical-2111303000111303-1322200023011102-0010101131223033-1001310203101120-1203210212121032-3032310231232011-1110331123123200-3000130333221331"></a>

<a id="canonical-1201031320313212-0012012021331213-3030231130123100-3210002311133322-1320330013322100-3112003122112301-2033323010101333-2300110012122321"></a>

## dns_server property — static_ip / 321303031302 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1021223320020321-3310322213330211-3011320201003223-3133332302011303-1030012032030323-1121202122030131-3223213322132203-2221123303323333"></a>

<a id="canonical-2222003203221313-0033210011323111-1211120201013211-1113330032032020-3113200332001133-0012101322123232-3302010223010232-0330200132332213"></a>

## ip_address property — static_ip / 321303031302 / 6

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

<a id="canonical-2101301303202302-3002113023213333-2201033020300002-1111113202333212-1130220120100210-2102123122002110-0321231321330301-2120022033033003"></a>

## Next pages — static_ip / 321303031302 / 7

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0213120301011333-1130333011301133-1222230300201122-0022022002033023-3201203103002123-2300133200220003-2230201131003101-0122102103123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233232213012231-2202321122323312-2101102322121033-0021203000023100-3021122310311313-0111231021220223-2021202103001020-0022221200322113"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 022130132200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0221132021022011-1211333202200100-1211213121111300-2033011103211223-2110303102000200-0011301311223131-2121003132131313-2110013013200202"></a>

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

<a id="canonical-1030132220010230-1220233332210130-3100122222221120-2221222300300323-3000230103111012-0310003323231001-1031022300030303-3111020032233011"></a>

## Direct properties — static_ipv6_address / 022130132200 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0302232033111311-0332012113112302-3323000232101230-1223223233130212-3122012122221133-3323112211210210-0233013013212120-1012103131213320): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2031333110133323-3010100202123331-1312120000113022-2230220122032121-1301203321122301-1001311030011020-2123122223001323-0201233101032232): complete subsection reference.

<a id="canonical-1000320332130200-1023310112333003-0023230221012121-1211123313102013-1330120100220213-3032210002300022-0333030303122300-3023100130210221"></a>

## Next pages — static_ipv6_address / 022130132200 / 4

- [kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0302232033111311-0332012113112302-3323000232101230-1223223233130212-3122012122221133-3323112211210210-0233013013212120-1012103131213320)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2031333110133323-3010100202123331-1312120000113022-2230220122032121-1301203321122301-1001311030011020-2123122223001323-0201233101032232)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0302232033111311-0332012113112302-3323000232101230-1223223233130212-3122012122221133-3323112211210210-0233013013212120-1012103131213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333332003022201-0001301231202222-2110000331121130-3013211012203020-3130132020202132-1300022303030223-3230003122201232-1111331211011121"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 200100021010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0213120301011333-1130333011301133-1222230300201122-0022022002033023-3201203103002123-2300133200220003-2230201131003101-0122102103123022)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1231031122301120-1212103101032321-3320020210002020-1032331132130133-0123033201000323-0302120002011031-3121120322110221-0303211333232320"></a>

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

<a id="canonical-0011102203133331-2031311021102332-2112320311123320-0210331303030301-2201232223211223-2112301302122210-1203300230112133-1100111333331102"></a>

## Direct properties — cluster_static_ip / 200100021010 / 3

<a id="canonical-3310233333122120-2223110033100321-3010213231330333-1131301330113310-0312312133021122-2102010313000132-2032310132320230-0222222200102002"></a>

<a id="canonical-1111331113033222-1202332013012132-2220330022010110-1010003200111311-1000002201013032-0212311110101131-0332131210310102-2303223130021023"></a>

## interface_ip_map property — cluster_static_ip / 200100021010 / 4

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

<a id="canonical-1203001311320202-1330002303310032-3112100131223111-1301213131120002-2011031121302113-2300222310210312-0130230123220130-1331212111112223"></a>

## Next pages — cluster_static_ip / 200100021010 / 5

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0213120301011333-1130333011301133-1222230300201122-0022022002033023-3201203103002123-2300133200220003-2230201131003101-0122102103123022)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2031333110133323-3010100202123331-1312120000113022-2230220122032121-1301203321122301-1001311030011020-2123122223001323-0201233101032232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200322030133311-0202301021130023-0320033033131030-2102311200310300-1321102021103200-2222220201213211-0331210330310211-3122320203203021"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 322310302023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0213120301011333-1130333011301133-1222230300201122-0022022002033023-3201203103002123-2300133200220003-2230201131003101-0122102103123022)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-0123322232213020-1120220333130023-3132002320002300-1322202011023023-2032232012103102-1020112133120101-0222111113320122-3210223103001120"></a>

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

<a id="canonical-1030122321023031-2111220000002213-2001311033313310-3322333232322000-0021312311313130-0332320230321030-3000030133120313-1133301122001302"></a>

## Direct properties — node_static_ip / 322310302023 / 3

<a id="canonical-2303203202212332-1113222132132302-2102102111023112-2213100111001300-3020320011220103-0002201201302020-2101302000023013-3311221033123023"></a>

<a id="canonical-1210032221103033-0211302133213010-1323013132123110-0332001323231322-3320321030200213-3221221221021320-2023012330313323-0213220111132323"></a>

## default_gw property — node_static_ip / 322310302023 / 4

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

<a id="canonical-0223010303311031-3002002100213213-2322301201321321-2210233331210232-1312232212133122-0332020012130213-2021001111110021-2222022023010002"></a>

<a id="canonical-0131303123222210-1330123202131123-3030033131102330-1000310310000102-1121331100002101-2010310301102221-3012303333301102-2113332120112311"></a>

## dns_server property — node_static_ip / 322310302023 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3320133102223312-0021221122002003-3130112123223032-1022101302322300-3012330223331232-1300230213111132-1121303013000332-2032321131101321"></a>

<a id="canonical-0313330303220103-1200132223302330-2321311033333010-1031332123300010-2031010313032223-2001310330113012-2111302331311211-0202320121121013"></a>

## ip_address property — node_static_ip / 322310302023 / 6

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

<a id="canonical-0122221223000002-0311302221323121-3301232010201130-0311222120013322-0320122222001313-3301100030200332-0321120022321233-2201311133303123"></a>

## Next pages — node_static_ip / 322310302023 / 7

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0213120301011333-1130333011301133-1222230300201122-0022022002033023-3201203103002123-2300133200220003-2230201131003101-0122102103123022)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1332012121210311-0212112111221311-1030013112321030-1210210102010012-3301231301113222-2223002330222030-2210233312033023-2023021212202233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200112222032112-1133022301213300-2220310100200312-0333303012201101-2332312313223120-0223123210222133-3123110322201020-3030220320201031"></a>

## kvm.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 203120212000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0010331222031302-3300202313011311-3220101011233333-3002312303322302-1232020032103301-1000332003321211-3020000233323330-0332200000202101)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-3331123301013303-0131230003020311-2213123010100120-3200023330331302-0311301112301112-3310130010331032-1213312321330212-1033221012113000)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- kvm.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2112122010120103-3303132331302301-1133113132013122-2110021313102311-1103201321231103-0102121001222312-2231032130233002-0310000231210133"></a>

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

<a id="canonical-1121313102203202-3311101330212331-1010210033012302-1001110301112031-3313011230033310-2232232103313101-0003202010222021-0122312310231100"></a>

## Direct properties — vlan_interface / 203120212000 / 3

<a id="canonical-3021212133332011-2111112111222321-3321021132031232-1202233113303211-2013111331200020-1312321312320021-0232003323123111-1011230211011220"></a>

<a id="canonical-0231031032301030-3232030131301320-0221221031232222-1211230231003332-1312032100303113-3122310021322312-3011202110330021-0010221031230112"></a>

## device property — vlan_interface / 203120212000 / 4

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

<a id="canonical-2100010303031012-0030323130202233-1231203310333101-2001113112222033-1300133120100213-0003211013321303-3131202013200331-3013223200221302"></a>

<a id="canonical-1022012101133301-2002312302232103-3012213113212020-3111210103123110-3312003310330200-2322311213130232-1231103021032111-2022313333220330"></a>

## vlan_id property — vlan_interface / 203120212000 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3122122333001331-0100103013033011-0102303102012231-0211223100201010-3331231110020332-1230230020312133-1320203123303323-0130101112012331"></a>

## Next pages — vlan_interface / 203120212000 / 6

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-0110310013223321-1332303032203321-0211010100200211-2211132011022132-3303022220030100-2122203230313031-2202021000110320-0001313131103321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0121223111122001-3321121211131212-0130203002210033-1220320201001121-1121022313021321-2102030003320310-3112312203330120-0133131122212110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033122202133213-0021322222202300-2310332223233320-1231203333102213-3000012103331103-3212203110313003-1032112331232133-1323330001103310"></a>

## load_balancing — load_balancing / 031102003130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- load_balancing

<a id="canonical-1033011302031120-0322000222111231-2313022130110113-3100231321003333-3330013211123230-1133131220021233-3323010322322310-3122030132220311"></a>

Type: `"single"`. Computed.

Section contains settings on the site that relate to Load Balancing functionality.

Upstream description:

This section contains settings on the site that relate to Load Balancing functionality.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021300033321122-3133232212323210-1120002032113310-3132302032210113-2012101303130013-1023300200000310-1002211220312111-3210021030320222"></a>

## Direct properties — load_balancing / 031102003130 / 3

<a id="canonical-0330333000120123-3222332301212003-3233203121330201-1320033222231220-3110133130030021-2121200301101313-2201203133301331-0003303103020300"></a>

<a id="canonical-2021322121003022-2222123102320120-1321033023112332-2113103103112201-0023210331100212-1010232310220303-0101213230031131-1101312303233321"></a>

## vip_vrrp_mode property — load_balancing / 031102003130 / 4

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200200013200231-3321231202320133-2333212111122322-1000210013313021-1003120211102331-2032321131202221-1103303022323112-1000003100011313"></a>

## Next pages — load_balancing / 031102003130 / 5

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233101203000003-2013102111122002-2302213200302211-3330010332000232-1112000323122203-1201102222332122-2020331303013331-1310223103230131"></a>

## local_vrf — local_vrf / 303001030001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- local_vrf

<a id="canonical-2010033300111012-3120330330233133-0331332123123210-3102332013122323-2213303212110022-3013033031001203-0032011313300312-2122000300022123"></a>

Type: `"single"`. Computed.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF.

Upstream description:

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

<a id="canonical-0100000203211113-3231222103220002-2021221022032030-0001023022000233-2020211020031113-2033011202332310-1311312103201122-1001223313132112"></a>

## Direct properties — local_vrf / 303001030001 / 3

- [default_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1312212213100201-3212031020210331-2321023200012020-1120311200012030-0021112200313121-2030021100323103-0113111110223030-2221020100022210): complete subsection reference.

- [default_sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0030030111322113-3202133313133012-1222032133311310-0003010333112300-3201312210233203-1230121300033331-2020311100232223-2331120231113233): complete subsection reference.

- [sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122): complete subsection reference.

- [slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300): complete subsection reference.

<a id="canonical-3331302120313232-1101302100112211-1222022033032221-0231312003310130-0233112032110333-0111003101020220-1200200033120323-1033302213113222"></a>

## Next pages — local_vrf / 303001030001 / 4

- [local_vrf.default_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1312212213100201-3212031020210331-2321023200012020-1120311200012030-0021112200313121-2030021100323103-0113111110223030-2221020100022210)
- [local_vrf.default_sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0030030111322113-3202133313133012-1222032133311310-0003010333112300-3201312210233203-1230121300033331-2020311100232223-2331120231113233)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1312212213100201-3212031020210331-2321023200012020-1120311200012030-0021112200313121-2030021100323103-0113111110223030-2221020100022210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221010211300233-0030001312302133-0301201000100011-3321121031202110-1123002023300111-2202223223122203-2023121100221220-0022122310113100"></a>

## local_vrf.default_config — default_config / 131002131130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- local_vrf.default_config

<a id="canonical-2330012233133202-3001031222020000-1312222000331022-2310133203030121-1023331233021011-3120122132011100-2001011121113311-3230323203002303"></a>

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

<a id="canonical-3321321331111312-2102232223213132-3312323303223123-3103110011202223-1021233301033033-3313311300312233-1101032300031301-3023132331000030"></a>

## Direct properties — default_config / 131002131130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133023110310030-3200020000210231-3112331120200101-2030130320303123-0122201201220001-3333012300021232-2223320313232130-3133003103031122"></a>

## Next pages — default_config / 131002131130 / 4

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0030030111322113-3202133313133012-1222032133311310-0003010333112300-3201312210233203-1230121300033331-2020311100232223-2331120231113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211300121130232-3030001232031031-2302031103313030-3000010230112031-2021123021122031-3310132211121032-0113010230122331-1232011123200112"></a>

## local_vrf.default_sli_config — default_sli_config / 032222020133 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- local_vrf.default_sli_config

<a id="canonical-2031013311230021-0323122112321012-3003323313210102-2110232101323302-2232323201102220-1023233113121232-1011121032322210-3323332001313213"></a>

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

<a id="canonical-2000201212330102-3112131321123202-2020323330002002-0022232113013022-0131303131022113-1122302332311303-0321111330123111-0200232012131321"></a>

## Direct properties — default_sli_config / 032222020133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222222100132221-0202311300311222-1301312233001123-1103113322321132-1020123020113021-1332211332233303-2323120110031203-1231023301310012"></a>

## Next pages — default_sli_config / 032222020133 / 4

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230023132202302-1220221231131122-0131030302223222-1213302212033021-3323131200013031-2312102201322231-2011211113210031-0102223130102223"></a>

## local_vrf.sli_config — sli_config / 200300012020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- local_vrf.sli_config

<a id="canonical-3220113013233302-1112103103330110-2122232331131213-0023123022222123-0210203002310103-3121213111123322-0030210202022031-2102001302032321"></a>

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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-2213301010110020-1323130300122002-2310120103301003-2031022201113200-3331223330211022-2200010320010111-2301133212000310-3203021210003132"></a>

## Direct properties — sli_config / 200300012020 / 3

<a id="canonical-0220133300003302-1311333330301212-2131231122201123-0301211010003320-3101301320121031-2132132022003223-3310212300103221-0210033001201303"></a>

<a id="canonical-1200311101233031-3200223311012010-1302132131112123-3011330132322323-0023323002220323-1303123230101111-2021321220101022-2232011012313120"></a>

## labels property — sli_config / 200300012020 / 4

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

<a id="canonical-3232301100013213-1200323132331220-1011202210000110-1101323323013002-2323022120302030-2321201002210212-3031303212012113-1223130323311103"></a>

<a id="canonical-1320010201231311-2323233330212102-2203100200320230-1002013212032300-2132020100031002-1032020122320121-2320112322200312-2320200231323221"></a>

## nameserver property — sli_config / 200300012020 / 5

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1013031010323200-3313120002130121-1303333121001101-1310212202220010-0133222210211012-1103322133332201-1102333303021121-1311320002322020): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2002031323032002-3003033333323223-2023330220230331-2200100020312222-0302030011020000-2031101220210112-1223313131121222-3012303101200313): complete subsection reference.

<a id="canonical-2232310120102002-0103211100031231-0011112330200323-0231110330023010-1322231120112321-0013311313020013-3303120111210323-3132020301112111"></a>

<a id="canonical-1321301110103132-2002010232202320-0003101232332101-0223021031203332-1301131101020013-1321002332130030-1033130202012312-3101212001010231"></a>

## secondary_nameserver property — sli_config / 200300012020 / 6

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012): complete subsection reference.

<a id="canonical-1220333032010222-2321203013233121-0021023232012311-1021213112231010-3012120003221002-0110223332123310-2031113120321103-1021203013311012"></a>

<a id="canonical-1313221110022112-2103310230301022-3120122320222120-0230122212122213-1022100311002210-3212322121011331-2202100032303210-1132130232223012"></a>

## vip property — sli_config / 200300012020 / 7

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

<a id="canonical-2010003221002022-3222300211202330-2000002232302031-0231202012200022-2303120321123302-3013211202233031-1201200311101322-0031312220010000"></a>

## Next pages — sli_config / 200300012020 / 8

- [local_vrf.sli_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1013031010323200-3313120002130121-1303333121001101-1310212202220010-0133222210211012-1103322133332201-1102333303021121-1311320002322020)
- [local_vrf.sli_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2002031323032002-3003033333323223-2023330220230331-2200100020312222-0302030011020000-2031101220210112-1223313131121222-3012303101200313)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1013031010323200-3313120002130121-1303333121001101-1310212202220010-0133222210211012-1103322133332201-1102333303021121-1311320002322020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323331111011101-3031333201012011-3211223220012033-0233211322221323-0133133132011011-3312333232120301-3230011131120223-2121223021033203"></a>

## local_vrf.sli_config.no_static_routes — no_static_routes / 131210032322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- local_vrf.sli_config.no_static_routes

<a id="canonical-0033021333032220-2001123320122201-3210222202201131-1202323003303313-0120123322231001-0313213312221131-0010002013001212-3230211100020112"></a>

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

<a id="canonical-1331122311221002-0003113331300103-2311323203300301-2122110310113310-1313103020310202-2210010232302103-0013001131132223-1320302312110121"></a>

## Direct properties — no_static_routes / 131210032322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100021122303003-2200331001331322-1021220302131310-3230320100102113-0233321023311313-0232321200223203-3032012032313033-3312202012331120"></a>

## Next pages — no_static_routes / 131210032322 / 4

- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2002031323032002-3003033333323223-2023330220230331-2200100020312222-0302030011020000-2031101220210112-1223313131121222-3012303101200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121232211210222-3203002213313331-3200120300300121-1212020130002230-3203302322030100-1110210222231210-2120113031120120-3110230033201312"></a>

## local_vrf.sli_config.no_v6_static_routes — no_v6_static_routes / 332302310110 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- local_vrf.sli_config.no_v6_static_routes

<a id="canonical-1300303320000331-3130022332120132-0020312112210010-3320333132313001-2332202120310133-0032000110312012-0211222311302132-2120220131130212"></a>

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

<a id="canonical-1310310030100012-0102031332003133-0313223303313101-2020032220300102-3201330221301333-2112001222313033-1311303200302223-3201323223301113"></a>

## Direct properties — no_v6_static_routes / 332302310110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333032103003011-2111200303301312-3130120002010213-2312201222222101-3330300032120031-0211002000212122-3203212320331200-2003032231130332"></a>

## Next pages — no_v6_static_routes / 332302310110 / 4

- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100003230112230-3221121110021110-3203101021032300-2133322120011221-0233232022023132-0301303210133223-0201033012130211-3312232303013030"></a>

## local_vrf.sli_config.static_routes — static_routes / 022303201030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- local_vrf.sli_config.static_routes

<a id="canonical-0203130333111013-3221121130113211-3230210113000122-3131110331012111-2120110201322320-0313221200220103-2201002123001021-3310221210313122"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322213022332022-3231023332021121-2330212312022313-3120011320020211-2101301220210032-0230033300330200-1231110120311033-1102023123001112"></a>

## Direct properties — static_routes / 022303201030 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010): complete subsection reference.

<a id="canonical-1133232002130012-0320023012110302-1112000320022003-3301330101322001-2230222231033002-2331123233232330-1210203303110032-2113130301123220"></a>

## Next pages — static_routes / 022303201030 / 4

- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112023323132110-3233321303012312-0312133220302002-2302112111120101-2333013323313103-0012310331202310-0323311000323000-0322012003021013"></a>

## local_vrf.sli_config.static_routes.static_routes — static_routes / 330000121312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- local_vrf.sli_config.static_routes.static_routes

<a id="canonical-2211231332222002-2002112022113223-1202332200212011-2012213233031310-1320202231201113-1302010322312202-2332023022222021-1313333113232203"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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

<a id="canonical-1000121030020013-1000002102200120-2133212201132232-3233200001111103-1312321310120203-2230322202000312-2101222023102022-2023322230031213"></a>

## Direct properties — static_routes / 330000121312 / 3

<a id="canonical-3201232031102222-1003031001011030-0232023303022230-3030230303121301-0320320230100013-1231300302022211-0312101010201330-2312112133330112"></a>

<a id="canonical-1000310013000300-1102022330003211-1011022230323210-3023103201123030-2130011003113012-2313330011100322-0231313011001000-0200132020203220"></a>

## attrs property — static_routes / 330000121312 / 4

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0300321210320232-2202101312312300-1201111020132030-1021111110201102-3213010210132233-3033223013320220-0323313013100033-1233200131203121): complete subsection reference.

<a id="canonical-0120003220202131-1121332101123233-0232302132333330-2330232001031311-2023230003030303-1233111011231312-2120003131233212-0310210320221332"></a>

<a id="canonical-0120010232121321-3011101033321022-2320201031320021-3112211002110122-1312321002321232-1322110202221130-1323122002321113-0301101102131133"></a>

## ip_address property — static_routes / 330000121312 / 5

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

<a id="canonical-1030212130101221-0132231313211312-1203311223300330-2011323110133222-3002321030100313-3321102123230332-2203312222001023-1323120330021321"></a>

<a id="canonical-3212103030001013-1000313111101221-2200200212330331-3212022223123303-3310103303213313-1121001031033301-2133032112321331-0321323313312000"></a>

## ip_prefixes property — static_routes / 330000121312 / 6

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101): complete subsection reference.

<a id="canonical-2123232100002231-2310010002123211-2133313230211202-1323033323333201-0113113111012223-2302130311203232-0311300013232332-0233301123013221"></a>

## Next pages — static_routes / 330000121312 / 7

- [local_vrf.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0300321210320232-2202101312312300-1201111020132030-1021111110201102-3213010210132233-3033223013320220-0323313013100033-1233200131203121)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0300321210320232-2202101312312300-1201111020132030-1021111110201102-3213010210132233-3033223013320220-0323313013100033-1233200131203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001300332122133-3212331112112231-0332313300213311-1321101322201032-2113200031202101-1133101322122121-0331223021213003-3100330311302302"></a>

## local_vrf.sli_config.static_routes.static_routes.default_gateway — default_gateway / 201132132332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- local_vrf.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-1012031121013012-0002031012103322-3100231201110112-2333133332103301-1321110121212101-3120312122132313-3011220300110022-2021212301002220"></a>

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

<a id="canonical-1113233020302011-1111212300211023-3020223231020130-3211111200212312-0330312230222210-2102202122312033-1033111132202033-3313010013110300"></a>

## Direct properties — default_gateway / 201132132332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102222010103011-1130120031203031-2001022113323003-3313330120030332-2321132333132310-1013221033302033-1002223313203021-2203332000133030"></a>

## Next pages — default_gateway / 201132132332 / 4

- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200333301000230-3020133213130132-2321111301303312-3001202301211320-1322030031032100-1032003003201020-0203221031313112-0201231031112012"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface — node_interface / 120020313202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="canonical-3233333111320201-0322332032302301-3110301001013302-2200031301230011-1312001302032010-1102312121233213-0312032132011033-3230220033131123"></a>

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

<a id="canonical-0122012311203010-3003312220220120-3210030302223201-3323303231113121-2323103330210120-1112231110322121-2223232301203013-3200310012020000"></a>

## Direct properties — node_interface / 120020313202 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0221010011332321-3330212232121203-2333133130211103-2311121203220321-0220120221111113-2011211002023201-3003322003313232-2103003322003232): complete subsection reference.

<a id="canonical-2200303033102023-2200103131003233-3032120021322332-0122212313103313-2000202213310121-2133033100221300-3312032210233222-0101023012221111"></a>

## Next pages — node_interface / 120020313202 / 4

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0221010011332321-3330212232121203-2333133130211103-2311121203220321-0220120221111113-2011211002023201-3003322003313232-2103003322003232)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0221010011332321-3330212232121203-2333133130211103-2311121203220321-0220120221111113-2011211002023201-3003322003313232-2103003322003232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221021002100011-0101030220320321-2133002230000012-0230011000301010-2112232311121300-1032012213013023-3003113111123300-3202211112011203"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list — list / 031310203323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-0123310003210021-1233111302132223-1311210020311311-1311212301101123-0122100313310023-0222122011031031-2222001021220322-2202313102103330"></a>

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

<a id="canonical-1033122300301200-1031303102111121-0033321122312021-1313321130030211-0010213201333123-2232122310213031-3300221210030211-0003201203003201"></a>

## Direct properties — list / 031310203323 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3033102032200002-3130212212121232-0112012213023320-0330330132303130-0120211023103331-3133323232310103-0123211032310031-1310030120120103): complete subsection reference.

<a id="canonical-3031010112320210-3013322113130312-0220132102010120-0222122322302103-3011232200012203-2002230013213123-0110222122011222-2001011122232033"></a>

<a id="canonical-3333112233033033-0012332103211111-0312133010311321-2032120302023111-0201233220332032-0003122110312323-1012331203311331-1112010013300121"></a>

## node property — list / 031310203323 / 4

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

<a id="canonical-0332330023333031-2110102022301000-2322131110223223-2202231212121303-2032030002013211-2033330222131122-1001103010011210-0030030201220323"></a>

## Next pages — list / 031310203323 / 5

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3033102032200002-3130212212121232-0112012213023320-0330330132303130-0120211023103331-3133323232310103-0123211032310031-1310030120120103)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3033102032200002-3130212212121232-0112012213023320-0330330132303130-0120211023103331-3133323232310103-0123211032310031-1310030120120103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003133301112130-2332021312102030-0001110102021230-2331322210203012-1330132303210201-0110331312221113-3310132021022001-2131121312321100"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 123231001323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2103322220010200-3110133223321003-1321113301233013-1203013211022013-3212103301311101-2232313030100130-2211100113023223-3120220102132033)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1133101233100321-0301302300200211-1313232320323023-2233020201302001-1212002120013111-2233100000320001-0210222020033131-2030320021213010)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1033103223101320-3303122123233322-2221100102000133-2223312022121311-2221001312310032-2120003113222020-0223203010302223-0121111102112101)
- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0221010011332321-3330212232121203-2333133130211103-2311121203220321-0220120221111113-2011211002023201-3003322003313232-2103003322003232)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3323003212012223-0312130022203102-3210313332201330-1033321110021222-1033303111033103-3131322123103322-0231330321330020-2211012113320113"></a>

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

<a id="canonical-0132011212022032-0132223330122321-0020222011132331-3313321210000223-3123222132331000-2103332203010011-0300200302310021-2201131333022320"></a>

## Direct properties — interface / 123231001323 / 3

<a id="canonical-3232322131213001-1310310103321023-1120223100333323-1222103001320000-0133201213333331-0321100120213302-3201323222021120-2101200302330111"></a>

<a id="canonical-0101101021320320-3321211212113233-2003322131311202-0013200212323032-3301311103322121-2333112010231302-3132012010210232-2023201130321030"></a>

## kind property — interface / 123231001323 / 4

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

<a id="canonical-2303233003313032-3301003101311031-3103123313002130-0322102101112132-0031221330300020-3330312120100122-2203310320231301-3233020231100111"></a>

<a id="canonical-2321110012232100-1233231110322303-3033313322233102-0030031121113013-2300013220232111-0221333131003013-1300033221331300-1203002103301320"></a>

## name property — interface / 123231001323 / 5

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

<a id="canonical-1232202110320313-0333330201031212-2011311013100310-2111212202102312-3200222022331110-2221002110300212-3031320330032100-3023300100033322"></a>

<a id="canonical-1310201200232312-2231133113012321-1201331233010321-0111212021001211-3312032031103021-0032121320333113-0220333310231310-3013203313111111"></a>

## namespace property — interface / 123231001323 / 6

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

<a id="canonical-2102213102102100-2323112322010212-1103212203031210-2311320322213202-0232002002011100-1212121011100032-2313232000322323-0103202113111322"></a>

<a id="canonical-2230133031112232-0102131102301000-2122330323100032-1011333012300011-2033202313033032-2223102302331002-2130131023103311-3030033321230332"></a>

## tenant property — interface / 123231001323 / 7

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

<a id="canonical-3310231131010132-2230102113300202-0102023010231120-1230332132321202-2123023332212222-2311201300333123-0031110313212331-1033103001322021"></a>

<a id="canonical-3222312302032202-0122202123031333-2022001320011002-3322323030331321-0232313310103311-3001200112111013-3022320303223002-3313201030010010"></a>

## uid property — interface / 123231001323 / 8

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

<a id="canonical-1232321120121213-0123101223320010-1323312111220002-3101302021020011-2232220231301122-3220121200222321-0003232233301121-3303223211210023"></a>

## Next pages — interface / 123231001323 / 9

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0221010011332321-3330212232121203-2333133130211103-2311121203220321-0220120221111113-2011211002023201-3003322003313232-2103003322003232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231020211210212-3102200320030200-1130222100301211-3013302122302133-1203011210200330-1303130033333121-2103010321030313-0201231022130030"></a>

## local_vrf.sli_config.static_v6_routes — static_v6_routes / 231000320113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- local_vrf.sli_config.static_v6_routes

<a id="canonical-3323101030110303-2213330221123220-1331303332313122-0131231203000220-1210322031230000-0232233221211213-0223002000103100-3122001012301223"></a>

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

<a id="canonical-0230313200113231-0321113221102101-0330001300302113-2210130202120002-3123331203312011-2021222212203222-0303310220211211-2302220113222232"></a>

## Direct properties — static_v6_routes / 231000320113 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230): complete subsection reference.

<a id="canonical-0112133103332301-2132001132203013-1121101322220302-0023231102033230-2312011210333100-3132130202300123-1311020131010130-1111001100110133"></a>

## Next pages — static_v6_routes / 231000320113 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020002222331131-3312221221021122-1213121021002333-1300100112102123-1330302201232120-3102131123113121-1121103300320010-0113101100223023"></a>

## local_vrf.sli_config.static_v6_routes.static_routes — static_routes / 302333130312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="canonical-0222103223321002-1032023200201302-2031113322200112-0031111310001003-1200033210103222-3300220200011230-2010232302212323-1111023220133132"></a>

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

<a id="canonical-2102000212230130-0203220310300221-0131303210013303-0011001301200131-3022310032122322-0032221032010103-2330222221112221-1301131233110302"></a>

## Direct properties — static_routes / 302333130312 / 3

<a id="canonical-2303331030021203-0011333221013202-0201032000223203-1131220311121023-2122202232313013-0120222033321302-0221112230220121-1123303010232222"></a>

<a id="canonical-1111110331123222-2121303323303223-0221200121212213-3133200230223200-2302231322311210-3201312113033023-2102312320312221-3203310332111303"></a>

## attrs property — static_routes / 302333130312 / 4

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2133110011230222-3321011122100030-0220001112102130-1201121323312320-3321332002321312-0131310321203012-3102231023022112-1131222020030320): complete subsection reference.

<a id="canonical-3313230231001002-3311032221112101-1033213312102310-1010211313031323-1200002312111111-2023032200323020-1100310101021210-0330112233102130"></a>

<a id="canonical-2120032300213013-3220123313323000-1131100033021213-0031111012212330-3223302201310131-1201133101322223-3002120210131032-3301132012300332"></a>

## ip_address property — static_routes / 302333130312 / 5

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

<a id="canonical-0320133211130331-3202201013111010-3132232112312000-3231122333123301-0030301003212003-3113111031101331-3230211300221102-1131001231331130"></a>

<a id="canonical-3131311122021120-1111112320132312-3013101232323102-3122311230102113-1223013013300010-1133301223101200-1110130013003231-0122031212010302"></a>

## ip_prefixes property — static_routes / 302333130312 / 6

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132): complete subsection reference.

<a id="canonical-1332122123300313-2001232323322303-3123131121122030-2010223210302310-3133223133102013-1003320211333100-1003222332322333-1301013203103320"></a>

## Next pages — static_routes / 302333130312 / 7

- [local_vrf.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2133110011230222-3321011122100030-0220001112102130-1201121323312320-3321332002321312-0131310321203012-3102231023022112-1131222020030320)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2133110011230222-3321011122100030-0220001112102130-1201121323312320-3321332002321312-0131310321203012-3102231023022112-1131222020030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213333122100220-0202002333010333-2311311203332313-3112100212310123-2000303300130322-1102102102331010-3130000022100312-0202231013032020"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 010010103012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- local_vrf.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-2210231001031322-1021223230131320-0332021010223023-1311230010022303-0333233311001332-0233313322231303-2220233330303312-1011123202100223"></a>

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

<a id="canonical-0301130211231122-3213003320030001-0332130013310013-2023312122133221-1131102230230101-2110113120313030-1333130003323111-2132311222110210"></a>

## Direct properties — default_gateway / 010010103012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130210330201203-1300022203110313-0131222300321313-2313131210210020-0223300200002022-1322300031332130-2221323121133000-3222311133030032"></a>

## Next pages — default_gateway / 010010103012 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210111220320032-2130111010321232-1203303102311210-1012333220101302-0212132203030111-0330033332313203-1131210103131012-0321212220300113"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 333031221012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-2223223101201033-3202003031222021-3102022220332023-1101103122003121-2131230320231213-0013122322231010-1033313233200021-0312330102212111"></a>

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

<a id="canonical-1031330110320112-3031210333012231-1233030122123311-2033333113202120-1133323023303302-1101232010333010-3232300002132021-1313211110203310"></a>

## Direct properties — node_interface / 333031221012 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111): complete subsection reference.

<a id="canonical-2122201120110202-2123012301333330-2320222201323023-2022230322021131-1232021223312320-0013310123312321-1011230031033303-1231110231101311"></a>

## Next pages — node_interface / 333031221012 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201130233120220-3200313302313223-1232322332203121-0013221321202000-0100003220123222-2231103323221103-1031201311221212-1223202212122323"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list — list / 223103132203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2203133003210120-2231112330111110-2030222311211302-1102201321301033-0331013123002130-3332120021121130-3212320331331020-2101001020112320"></a>

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

<a id="canonical-2333320230002122-2213110320030311-2330233300022330-3022232113123323-3202102212123331-3011103330223020-3011211012001233-1320022303230010"></a>

## Direct properties — list / 223103132203 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3320203210031330-3201020130111221-1210003000320233-3220230121132000-1230112313023102-2223033221002103-3011331113213111-0113010300113013): complete subsection reference.

<a id="canonical-1222010011300012-2022232310022023-0121300110011310-0310000032312333-3112023301023311-3310211020312021-1021330313232200-2330233112233103"></a>

<a id="canonical-1022310133033121-1303022303311200-3233003233032220-1131112302003012-0230032320111232-3300033020122132-0211131032212113-3223312103312013"></a>

## node property — list / 223103132203 / 4

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

<a id="canonical-2102033110103311-2101011023321033-2110212013011310-1001231210303121-2013203033313010-2300021102203321-0113301222300313-0220311102220011"></a>

## Next pages — list / 223103132203 / 5

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3320203210031330-3201020130111221-1210003000320233-3220230121132000-1230112313023102-2223033221002103-3011331113213111-0113010300113013)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3320203210031330-3201020130111221-1210003000320233-3220230121132000-1230112313023102-2223033221002103-3011331113213111-0113010300113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202221021200000-2013031230030121-3020310321332223-3003112132100001-1120220102203022-0312320333103001-2103330311030013-0220032213300130"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 010303213322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1002322003022320-1110313312310031-0112223130021222-1322213210312211-1322210031203333-0032330321100121-1100210313020101-2103113223000122)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0313001032212020-0211332230221220-3200003221232133-3013021033010012-0220021022322113-2120001232221122-3213331212121002-0112313322230012)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0203212003013020-2201330032220330-3312321101021230-3211010303023120-2313210031312330-0231012013333133-3311103100321321-3023302000213230)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2213120302211010-3213221121001311-0133133230123103-1212022203001332-2100120320210010-1132102323211000-3001221020000330-1321132030203132)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2003320121033102-2031122111320223-0022232112102003-3013133031311101-3220220022203002-1023213120011103-1110321122300201-0132322321302200"></a>

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

<a id="canonical-3331231113103320-1220003122110321-3331331023222312-1023112100311333-1103101022333122-2301220200313222-2100022330123201-1111233200300231"></a>

## Direct properties — interface / 010303213322 / 3

<a id="canonical-1112310021231002-3330011330202323-3310323012021303-3122313213120033-2320303213322021-1133221112331313-0332100013033301-0023332023131122"></a>

<a id="canonical-0003030313021010-0003310030321313-3000231130031221-2033222101001320-2030113220200322-1113113320022231-1130220020233231-2130120300313101"></a>

## kind property — interface / 010303213322 / 4

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

<a id="canonical-0222023130032112-2031033320123123-0323213130133110-1321333332033012-2133012033203313-2000131331330202-0203102130123122-2032001132112121"></a>

<a id="canonical-3223202302331213-3113322112333102-3021102120022301-2200010113232210-2222323120122123-0220211203311302-2031311002300232-1132012333101200"></a>

## name property — interface / 010303213322 / 5

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

<a id="canonical-0231223020031033-3223032220031031-2312120030032121-1022131210030120-1220000202010202-1301220230020312-3032132223113003-2133321033311023"></a>

<a id="canonical-2123102310002030-2113333300203202-3020120013130233-1232122121131130-0022220322130110-3100022211100330-0322323222311113-2001222000120012"></a>

## namespace property — interface / 010303213322 / 6

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

<a id="canonical-1312013100113120-0032112331103330-3213222100202313-2201022002112222-0102112313133110-2030021132222132-0330321312331121-3120333221130021"></a>

<a id="canonical-2210103120123013-0000021022011212-1201333021310012-3002013101302021-2100030102233233-1221332300320231-3222013120313100-0020221330330002"></a>

## tenant property — interface / 010303213322 / 7

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

<a id="canonical-1102312102213333-1002310122130312-0223211321110233-2210302313202222-2213132101030122-1330121123012000-3032321003121021-1032312132230220"></a>

<a id="canonical-3211321223001203-0112330311330121-2203333222302010-1203020233103313-0002221110313303-0331032131233131-1112000213312130-1133201202001131"></a>

## uid property — interface / 010303213322 / 8

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

<a id="canonical-0030320302203312-2011013131132320-2210112303332222-3030010212321032-1111123111133233-0100310103011202-1000311000121230-2330213103100022"></a>

## Next pages — interface / 010303213322 / 9

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0321123102010132-3313210020023013-0213223230011022-3131223301100103-2231030231222221-1132003132110032-2003300010113300-3333303022033111)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010230311121000-1322011102010213-2110100303211223-0311103122212132-0102120322033331-0103103023130313-2102221123103133-1300221122230210"></a>

## local_vrf.slo_config — slo_config / 103312220121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- local_vrf.slo_config

<a id="canonical-1013302003210320-1010321312202012-1113002230031012-0313322103110330-0130330110021332-0323032233130022-3103102233330120-2102022322220331"></a>

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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-0202200323303030-2321233022110302-3330012300010230-1302110013222110-0011303112323211-3012121002231211-1231123033301122-1100123323010023"></a>

## Direct properties — slo_config / 103312220121 / 3

<a id="canonical-0003123121022021-1013130021332101-3031320132131101-3310013302032120-3000013031212022-2100203033213000-1222322300100011-0022322221221010"></a>

<a id="canonical-2331111223311323-2002031331302302-1102200313112300-2301001322221021-0200122320231231-2321220130213130-2111133203230001-3221133123013110"></a>

## labels property — slo_config / 103312220121 / 4

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

<a id="canonical-2100000031210301-3230203023002212-2231110221100212-0212131323310321-3003023031311232-1312133000102332-1110111333303201-2001121103321111"></a>

<a id="canonical-2111001023000032-3131330102232223-0323332202011300-2030110002031222-2203123132100101-2210023230032023-3323210220111012-3132233203123111"></a>

## nameserver property — slo_config / 103312220121 / 5

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2132113300200203-0000123232033102-3232001122302031-2003222000012132-3003101002012012-0302212131100313-3321333020133012-1300133231211202): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1121322113210100-2321101323310121-2303211003010113-2020132223121213-3103101202011320-1012033032022001-3233330203020010-0220002333022120): complete subsection reference.

<a id="canonical-3121312133022332-0101220202232212-0302013101310122-0201222131010013-2133000231332121-3233123223111223-0312213222231333-1323123302111311"></a>

<a id="canonical-2113201312103022-3001130311003333-0101230123211222-0013213031122022-2300232332302202-1020030222013113-2211110320332012-2221120202312003"></a>

## secondary_nameserver property — slo_config / 103312220121 / 6

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222): complete subsection reference.

<a id="canonical-3011120203301231-3323122103212010-0330131333212013-0232013330121133-0020200200013021-1213021202230313-3222101000010120-3033321121120310"></a>

<a id="canonical-0131233033123233-1213232100200132-3223230300021210-3210131211003203-1123303132200001-1130210201103232-0111113032221321-0302002011212132"></a>

## vip property — slo_config / 103312220121 / 7

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

<a id="canonical-2311223310320013-1003030030310320-1120301212230203-0202221131110132-2211211232122210-3322130331211311-0022320133333300-3103320010021332"></a>

## Next pages — slo_config / 103312220121 / 8

- [local_vrf.slo_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2132113300200203-0000123232033102-3232001122302031-2003222000012132-3003101002012012-0302212131100313-3321333020133012-1300133231211202)
- [local_vrf.slo_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1121322113210100-2321101323310121-2303211003010113-2020132223121213-3103101202011320-1012033032022001-3233330203020010-0220002333022120)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2132113300200203-0000123232033102-3232001122302031-2003222000012132-3003101002012012-0302212131100313-3321333020133012-1300133231211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332301001203210-3311032201001022-0011230312333030-1322132131333011-2131220013303221-3322212123010302-2121131210020212-3001103111100320"></a>

## local_vrf.slo_config.no_static_routes — no_static_routes / 032310010221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.no_static_routes

<a id="canonical-0002212233003002-0322113100013000-3302133202131001-2302111033033130-2000013323302313-3031130110223001-1302333023201033-3003010011013200"></a>

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

<a id="canonical-0331311221221013-3203102100022321-0132323312003101-2311302100132233-0232133223022113-0002202212121310-3121323200101212-3000201010100332"></a>

## Direct properties — no_static_routes / 032310010221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111200032113330-1223021031312120-3032200023132323-3311223303320201-1011222220002202-3100213310112233-2011020102133131-0230033032120301"></a>

## Next pages — no_static_routes / 032310010221 / 4

- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1121322113210100-2321101323310121-2303211003010113-2020132223121213-3103101202011320-1012033032022001-3233330203020010-0220002333022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321110012323102-2022230212132110-1003020023321133-0203223103221110-0302013031300023-0321220012230231-2013121231122332-1000120300103231"></a>

## local_vrf.slo_config.no_v6_static_routes — no_v6_static_routes / 312110120322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-3113133221211201-1030232331033303-0323111112303330-0102331120232112-2212003230131113-1102200300123002-1003312120131222-1001311300132103"></a>

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

<a id="canonical-2231023103220131-1202332311010200-3213332203023222-0312103023110301-2330311121211332-3121022200213303-0220110221113021-3320113323300312"></a>

## Direct properties — no_v6_static_routes / 312110120322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221013323132310-1100103333231230-3201231310121012-0333203222211010-2103000131111101-0000031101021032-3202232120323200-3302300121023200"></a>

## Next pages — no_v6_static_routes / 312110120322 / 4

- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131331132012320-1030033130113010-2301001311011303-2113003332223133-3333311202202212-2131110301102202-0002322021110312-2022313022100023"></a>

## local_vrf.slo_config.static_routes — static_routes / 333020111213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.static_routes

<a id="canonical-1211031221132030-3133211012021313-1213200303110122-0102031232111012-3210020221002201-2231011023221302-0030331020230032-1330113131322112"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312200130020320-2302030301330032-1302030131020132-1231110330022121-1323223012333102-2102322330321300-0113122200023232-2301012113232000"></a>

## Direct properties — static_routes / 333020111213 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013): complete subsection reference.

<a id="canonical-3323101232113331-0211110201111011-0111002023231311-1311213131103121-1311300231031301-3030120222033122-1113020310221001-1322321213133103"></a>

## Next pages — static_routes / 333020111213 / 4

- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231120003000222-1322321311101001-1100032030202033-3121323121123113-2031011223230331-2211021212112211-3202200023122222-3120100312233303"></a>

## local_vrf.slo_config.static_routes.static_routes — static_routes / 311201202010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-2322111233332011-2221112233033231-0000010132122210-2222010201011333-0111331220032211-3010030020221120-1213011321211302-2223200202032112"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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

<a id="canonical-0023001333312323-0132030101321301-1312012200003031-0332102231302132-3113201033033121-1030010330033010-2103031100200133-3230211112320201"></a>

## Direct properties — static_routes / 311201202010 / 3

<a id="canonical-2033020313220012-0120010320332210-3321201232131230-1210113301300232-3322321000231312-0102312322303033-2112003313101100-1300330312200213"></a>

<a id="canonical-2122300302331332-3132121322103132-3333320100211331-3212202033131021-2000121123021120-1032223120303013-2021311301021303-0323330000030221"></a>

## attrs property — static_routes / 311201202010 / 4

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121): complete subsection reference.

<a id="canonical-1212213113121112-0132323130032223-2321232113330030-3212300300212220-2103000132200311-0233321231331123-3102330211111232-0312323231322122"></a>

<a id="canonical-3310303222202202-2033210122101322-1012321122231220-2012303213213130-0313022333313220-2231001303220023-0220020013132013-3112230200102031"></a>

## ip_address property — static_routes / 311201202010 / 5

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

<a id="canonical-1131131131121101-2310103200012303-1212132133031311-2223211113100021-2203011201302000-0310310133032133-3010023023233311-0003130023312030"></a>

<a id="canonical-1322013200123302-1103111331220033-3221231330301222-1223010011102312-2302011311033101-0332001210323112-1232103312221101-3113331011023132"></a>

## ip_prefixes property — static_routes / 311201202010 / 6

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001): complete subsection reference.

<a id="canonical-0020020013002312-1103032201310313-2102312100323130-1200231121202303-2331120001110313-2331330313021302-1130311213002121-2130313210013032"></a>

## Next pages — static_routes / 311201202010 / 7

- [local_vrf.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122131313130131-0300000332003002-1323323333032032-0103012110300311-0201101121332212-0311303122202033-3203210323121302-3112221302022013"></a>

## local_vrf.slo_config.static_routes.static_routes.default_gateway — default_gateway / 023110131013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-3220330001301102-0211021232103133-1220312322120331-1132120200002130-3310210310221021-1221131131120012-3032103110011130-3023012230320200"></a>

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

<a id="canonical-1131312312132012-3120332232202313-3301221220030201-0033303321033202-3332120120130130-3201311131332333-3211012320332232-3331022200003010"></a>

## Direct properties — default_gateway / 023110131013 / 3

This is an empty object or choice marker. It has no direct properties.
