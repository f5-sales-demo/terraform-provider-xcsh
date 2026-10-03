---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-2312231023030112-0031001202231000-0301113111232212-1122321000011322-0031013132323002-2101012101313233-0113030131310323-0233033323302130"></a>

## Next pages — static_route_list / 302033300111 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103112031122202-0123000202123213-1103130031211030-3222213003323331-0030022113200010-1310123331113221-0230102302303302-0200021332301301"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 222111013212 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2203033201320200-0231123101310210-1311102011013311-2003233100303111-2003323301101012-3121212210330200-3131213130013331-1232001220010220"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033301220020302-0123210323200223-1021213102232133-2030030031131312-1101300200323122-2022303113100003-0312300301012221-2012123110223022"></a>

## Direct properties — custom_static_route / 222111013212 / 3

<a id="canonical-0201333032122032-0033233011200111-3010132302321132-2210002302321321-0321113133122033-0102213201132222-1310223021301213-2102300121211113"></a>

<a id="canonical-3231320031200102-2123132221133011-2221331130203121-0022122011100111-3123001232001231-2330320212220110-1121302203211223-3231132020310123"></a>

## attrs property — custom_static_route / 222111013212 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--azure_vnet_site--reference--group-005.md#canonical-1020311221122111-3223023010200232-1332131002112033-3312013012101231-0312302231112113-2212003012111132-0101211102310201-1100313213023310): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000): complete subsection reference.

<a id="canonical-1202001323102200-3013220212112333-1013323213200033-1223200321311303-0022130000333122-0200023213012331-1001330211003133-3230320200231201"></a>

## Next pages — custom_static_route / 222111013212 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-005.md#canonical-1020311221122111-3223023010200232-1332131002112033-3312013012101231-0312302231112113-2212003012111132-0101211102310201-1100313213023310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1020311221122111-3223023010200232-1332131002112033-3312013012101231-0312302231112113-2212003012111132-0101211102310201-1100313213023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311231123312102-2030330130030202-2303001033032200-3202320000013232-1033021331113221-3100220311332133-2120232013230210-1210312123110033"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — labels / 013011321131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1201221111310110-0000113300130013-1323333030233203-3321220132101213-1122112203302131-3120333223112000-0223110331200303-3032013120303111"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020030021332220-1021233011231003-3233210120211230-3001000012023303-3233101200030211-1032001212221110-0002133312331222-3331322312222100"></a>

## Direct properties — labels / 013011321131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101103203313033-1220313331013213-1310303230110232-1123110013312301-1132310112332003-3323312023020210-0132130222323032-0010013300213012"></a>

## Next pages — labels / 013011321131 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133011332221102-0101102201110330-1313012202230222-3221130200123100-3301033103212101-1213032320001232-0302133221010201-3103233322132331"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 212022120311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1101132330113131-1102120313201230-2212012101030313-2203112320020023-0122313201303102-3001101300222020-0011130113011333-2331212130010310"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211313311122101-0220301332210010-0123022232122332-0203000223020210-1323302302221102-2312101203022223-3211033232222110-0121021221133133"></a>

## Direct properties — nexthop / 212022120311 / 3

- [interface](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110120220303112-1200013100231222-3311321112231322-0120000323310103-1200013320210220-0022232033113200-1022330102231122-2322332133202302): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123): complete subsection reference.

<a id="canonical-3012210103103213-1111000010032312-3212221233033032-3133303033323112-1021110201013102-2300131021322001-3302133333111031-1102013132203202"></a>

<a id="canonical-1332233232323122-0330323333120212-1000231300032030-0211333221030230-2222222120230202-1303131113012121-2012313311231103-2300023011233322"></a>

## type property — nexthop / 212022120311 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120001032132201-3132311100213111-1110133031031002-3321120211223111-2023302202301231-2023011133311010-2200103031010212-0201023220220201"></a>

## Next pages — nexthop / 212022120311 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110120220303112-1200013100231222-3311321112231322-0120000323310103-1200013320210220-0022232033113200-1022330102231122-2322332133202302)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1110120220303112-1200013100231222-3311321112231322-0120000323310103-1200013320210220-0022232033113200-1022330102231122-2322332133202302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021213323200002-1202133012122301-1313110331011312-1113123330323332-1223231200331210-1322003030122332-1001201322111010-3111001321221331"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 123313130202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-0000030112312212-0313122111222202-1020132213022301-2133001230213210-3000210203030233-1212231011103003-3031321013210222-1331333001332313"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-1120010222222323-3232100231111230-1203302332231332-2113323111230010-0301002312031310-3333122110232312-2222011132122102-2223320010230203"></a>

## Direct properties — interface / 123313130202 / 3

<a id="canonical-0012330002013321-1000131322111033-3210301303210013-1321211331103120-2311300103132331-0101032330101101-0233102203231232-2323110310000201"></a>

<a id="canonical-2131333223103312-1032303113203330-0211332002023033-1213301012010022-2213013222133202-3023331203030120-3210333233301112-2011022011232302"></a>

## kind property — interface / 123313130202 / 4

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

<a id="canonical-3310032303113100-3302130020203113-3020113121023003-2202202322013123-1113321000030133-0220220300333322-1030302313001303-2123120132330030"></a>

<a id="canonical-0032312102330323-2112111203033321-3011003333212100-1001012122333133-3033321321311310-0132031203111300-2211112133232013-2313200021030213"></a>

## name property — interface / 123313130202 / 5

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

<a id="canonical-3011301323310002-1213123201110121-1013222022331323-3100302101302323-3010032212322322-1201222103132220-2223222010233230-0011132300210300"></a>

<a id="canonical-3223310033121323-3313230122312301-2330313130121333-0223301111223230-3311111312331002-2100010311112032-2310201210202131-2233333321332103"></a>

## namespace property — interface / 123313130202 / 6

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

<a id="canonical-0300003010222023-2220210301001021-2130101323020013-0100133120020202-3221122200333320-3213213110230032-3232332020011300-1003212230331211"></a>

<a id="canonical-0132131323322013-1102201030301330-3322330230221323-0012223022330113-1322221333133232-2133012311201101-2003020102203023-1003202210212002"></a>

## tenant property — interface / 123313130202 / 7

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

<a id="canonical-3122133033101031-2212311212021312-1120103211101013-2032020023222112-2011313302231230-0202000312310202-0222023213003223-0211011203110002"></a>

<a id="canonical-3230003033233120-2023120322032302-0233130322303202-2122323200330330-1111122212131122-1300003132232301-2022311033202002-0203321230020333"></a>

## uid property — interface / 123313130202 / 8

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

<a id="canonical-1121310010013223-1003031111100032-3323301233333222-0210223013211031-0222022323223102-1330212233203031-0322132122331023-1031032100103130"></a>

## Next pages — interface / 123313130202 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213133312232110-2301003033321230-2112113031200313-2220030132102302-1333001213200010-3320230323212203-2122103112123312-3110220332001132"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 230323130213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2220000000333223-2113113030232323-0011231302021301-2223111301010113-2123303000310210-2312310010230213-2102302000002003-0233212220111221"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-0120133310303123-0303232332200221-3120031020121203-3300232000210033-2121031002213231-3103103233231033-1300113301331330-3233210123301203"></a>

## Direct properties — nexthop_address / 230323130213 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130): complete subsection reference.

- [IPv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-2133100220312332-0302232023311232-0303221233111211-2020310032302120-3033120021131001-0033203021231302-2003111313023220-1002322020030230): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-0313311112102301-3301222313221003-0021023220222122-2002222120222202-1030322122031011-0110201023132003-3011331032330102-3201201112320012): complete subsection reference.

<a id="canonical-0330112001000101-0201213303313220-0231003303203010-2201102021033231-0103330201023110-2110120012223322-2003213310212321-2113210132120232"></a>

## Next pages — nexthop_address / 230323130213 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-2133100220312332-0302232023311232-0303221233111211-2020310032302120-3033120021131001-0033203021231302-2003111313023220-1002322020030230)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-0313311112102301-3301222313221003-0021023220222122-2002222120222202-1030322122031011-0110201023132003-3011331032330102-3201201112320012)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010233120213001-2122111213222021-0201013230320000-2131021320211030-1103222220110032-3203110221012103-1303110322032230-1222002101001011"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 002223102201 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-1003032030320210-0210311203112303-0010223302200010-2232113110303103-3312320132021100-1010310310021300-0302110310211112-3221120312201103"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3203230131223102-3232130011213121-3311021002000031-1333333030003303-0201312103113303-1110322332033312-1302012300000312-2103102332023001"></a>

## Direct properties — dual_stack / 002223102201 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-0003322010003000-0223303000331102-2211331032321120-3100130033031300-3210230000222223-1012223211101111-3320230222212023-0012100131330012): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-3031321101012212-2220103010200031-2220022101020201-2203230320301133-2113110211102332-3112223022220202-1102111112102123-0101330303200132): complete subsection reference.

<a id="canonical-2312213330100203-2230002330211000-2000323213221232-1033302102013302-1133321132323123-0121010333323130-2131131233113001-1122003303021122"></a>

## Next pages — dual_stack / 002223102201 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-0003322010003000-0223303000331102-2211331032321120-3100130033031300-3210230000222223-1012223211101111-3320230222212023-0012100131330012)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-3031321101012212-2220103010200031-2220022101020201-2203230320301133-2113110211102332-3112223022220202-1102111112102123-0101330303200132)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0003322010003000-0223303000331102-2211331032321120-3100130033031300-3210230000222223-1012223211101111-3320230222212023-0012100131330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302031330302102-3311321113230111-2030103331220013-0110030130011322-3221211120333020-1221330123222302-1322323211132102-3213012131022323"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 210001222203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0232212310211213-2021030102213310-0122303301223110-1013223012300123-1323212210121220-0110012333112030-2303010002030302-1211133210320302"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1033301230202202-0333001133101020-2110310033212230-2231002001113332-0010113230112310-2230331222233233-2331232331031222-3333320132311312"></a>

## Direct properties — IPv4 / 210001222203 / 3

<a id="canonical-2313000021123312-0330321012201100-2112322122203223-2222220120133012-1010201000122310-2121103013312210-1013312210103223-1223230332001033"></a>

<a id="canonical-3013130232133331-2202331111302321-0301032202101202-0010210301301233-1320332020131330-3103321121111301-2100010103331102-0333330221320031"></a>

## addr property — IPv4 / 210001222203 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3023022002110103-1003213203000000-2112220201301332-3333330203103203-2033322331302220-3112133000220030-2211002120331013-3022212023110223"></a>

## Next pages — IPv4 / 210001222203 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3031321101012212-2220103010200031-2220022101020201-2203230320301133-2113110211102332-3112223022220202-1102111112102123-0101330303200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220013122112112-3331010023110311-0220200112130321-0003113232221301-1102301200123112-3110133311313111-3032223023310212-3133302010102001"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 032230101202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3202212032230103-3301210333032210-3302032203122122-3023321032310131-1301311120313223-3102112101131023-3312201313100111-2230001311213332"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3030010121320201-0022023323321211-3110110213201033-1020313220313332-0102333302333320-0323222230001231-0002303220310313-1112213301103303"></a>

## Direct properties — IPv6 / 032230101202 / 3

<a id="canonical-3321322232211031-1022132131201313-3032101330021110-2131210300003322-1202203103233111-0033003132323123-1101211021033312-1002002320320303"></a>

<a id="canonical-2003111100333112-3230231110100230-2331232321210033-2120102322023001-2322001110110232-3201001213203223-3332102322000033-3300301103121330"></a>

## addr property — IPv6 / 032230101202 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-0031133332021312-0131312132003331-1323011212101000-0300110312223201-3100213120023102-3021303220133212-3222021103003212-0202201003031023"></a>

## Next pages — IPv6 / 032230101202 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312322103232133-3101031103002303-3112322310012133-2223113222102013-0112230032122223-2221023002200210-2130103012103023-0220112023123130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2133100220312332-0302232023311232-0303221233111211-2020310032302120-3033120021131001-0033203021231302-2003111313023220-1002322020030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233031330003310-3202102333102323-2213120212222110-1121330133032023-0111221302021320-3222100302123021-3120112012033331-3220311313102001"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 031321320022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-2133100222232133-3123211202221032-1323231023122313-3323002313213010-3001101113301101-3322210303323221-3012030121302122-0210001033211023"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2232203023232000-2311013130130333-0222222133010113-2022120310021000-0112210123300331-3302012131100132-1030132010321112-3310000302000312"></a>

## Direct properties — IPv4 / 031321320022 / 3

<a id="canonical-0220231233203011-0011311101223211-0130211033301101-0210313030100202-2203202011112232-0210002331303220-3311220002233313-0120103302012201"></a>

<a id="canonical-1312222213111033-2013221112300303-2010210031110220-3131231222322002-0121011331122231-3312302003221300-0310111313002110-3333111103200322"></a>

## addr property — IPv4 / 031321320022 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3320013000013133-1100312311300302-0311100033120122-2312300033200321-0002330220312013-2221012230023333-0230332230300311-1311320231323032"></a>

## Next pages — IPv4 / 031321320022 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0313311112102301-3301222313221003-0021023220222122-2002222120222202-1030322122031011-0110201023132003-3011331032330102-3201201112320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313133103110102-1321332230223102-3020023000113320-0300321001130310-1120120111332003-2312301311121310-1320332131003311-1122200202121210"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 202320301222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-005.md#canonical-3021021003330111-2030112002200133-1132021103231133-0122301111330132-1310020232120332-0213302223000020-3101220131120322-0321322123203021)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3330123000313321-1211212002222121-3233301332223233-1223333331132013-3323032002133132-2232111101302001-2033101103200022-0120231002312132"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0231320001322231-1223332311003311-0202230032322213-1132311202220200-3021213103000233-0322312000302001-3023233132010012-0133010230122223"></a>

## Direct properties — IPv6 / 202320301222 / 3

<a id="canonical-2031230300013023-0322212113222213-2233130133132322-3012330123313221-1102230301332120-2031012221203001-0200302013010032-3202211321111033"></a>

<a id="canonical-2123103131111201-2132001113300113-1311231303222033-3212221131213320-2122330311123033-1312300200302300-3303312312220212-1030001100000100"></a>

## addr property — IPv6 / 202320301222 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-2001102001330202-1311200331102203-1033030010330001-3101031233332333-3011230131301301-3233302023301202-1320023230121110-1232022001311201"></a>

## Next pages — IPv6 / 202320301222 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320333001222010-2332303022232312-1331023302231011-3132013020121102-3031022021233102-2230233123320100-0333312320033012-3123131023132123)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300230133123132-1220012313311322-3330303121121133-1110332000131303-1211223231131232-2123210011331133-1303011221333323-1200003002230122"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 223301010200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-0312213221312130-3131331103121321-1032220021313021-0020103323133002-1221030121232302-0020003001223021-2223312010130010-1301312202332212"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0102130302233101-0221101023102120-2231023133022320-1331300231301332-0020311012332221-1221333320320211-2323111122323312-3302120111211023"></a>

## Direct properties — subnets / 223301010200 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-2103322303220112-1232031003121330-0000231233000220-1132120100200001-2332302023331312-1313003231100123-1232112102011210-3113013030222233): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-3103000113320301-0330102131032103-2323221001301331-3123211220300210-2202103001222020-2033003233202030-0231123300223333-1322011223310032): complete subsection reference.

<a id="canonical-1221320221112232-2113132330120030-2212121311231330-0300002122033300-1222113011302301-0000322210111223-3011130132332200-2330133012202311"></a>

## Next pages — subnets / 223301010200 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-2103322303220112-1232031003121330-0000231233000220-1132120100200001-2332302023331312-1313003231100123-1232112102011210-3113013030222233)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-3103000113320301-0330102131032103-2323221001301331-3123211220300210-2202103001222020-2033003233202030-0231123300223333-1322011223310032)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2103322303220112-1232031003121330-0000231233000220-1132120100200001-2332302023331312-1313003231100123-1232112102011210-3113013030222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013011103102230-0210330332323131-3220320220323232-1211031101213211-3300101312033111-2231013120211132-2102132110210313-2302120231303222"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 210232120330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1102010101120122-0303210320020021-2330112023011202-0300030032032031-2132222300021332-3311322021220303-1121323113033031-2200010011103013"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213011201100302-1010310032121110-3101223102312120-0030201012032303-1133222030232132-2322312300001103-1103123010102232-2122210212303312"></a>

## Direct properties — IPv4 / 210232120330 / 3

<a id="canonical-2112100111012302-2202322001012002-2211300111213231-3130013112102330-1230300002211113-3323011113231121-0010313020212122-3101301130020200"></a>

<a id="canonical-3310011300122100-1310202313332120-0033332133221223-2130320322300303-3233113302021023-0003221120023022-2102121131020301-1120230301000331"></a>

## plen property — IPv4 / 210232120330 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1012121213031212-2001030013112310-1223002010321102-3113232201320120-1200222202323113-2313001310200303-3000203202022120-2130032100102120"></a>

<a id="canonical-3222101220031111-2131113032101303-1220010031111031-0312231000131103-0232200030022212-0130223221301210-1220011011223132-0132332331123210"></a>

## prefix property — IPv4 / 210232120330 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-0132231030232122-0033201002210010-3130201113211320-3200011033103003-1230212220302131-2321221022120001-3113133113203233-0131323102231303"></a>

## Next pages — IPv4 / 210232120330 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3103000113320301-0330102131032103-2323221001301331-3123211220300210-2202103001222020-2033003233202030-0231123300223333-1322011223310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020122121323231-1310110313013113-3210131213232221-2233121122031222-1033021301030111-0311012001002302-1211031000203300-2013113310313033"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 333020130312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-3331300101102232-2233302033200230-2221021301230030-1213303003022032-2033032032203331-1312101202001000-1120322010301313-3003322000030230)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-1111102202331220-0213001220320111-3321111231031230-2222001132232302-1331203220130311-2322211131323331-1132231000012231-0302310000213311)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-005.md#canonical-1332000221022112-1001103020332333-3333022220033222-2010010132133130-3310320102213222-1123022310210021-1302031030212002-0133013022010300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2002020310320220-2131212330230010-3102020010323321-1113320113322222-0001201232321230-2112301003033003-1311113012210321-1210023003330202"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313110001333110-2303103103010102-1002121032302232-1300310302313302-3210022000130120-0012001013031111-1123022310212110-3112032021113113"></a>

## Direct properties — IPv6 / 333020130312 / 3

<a id="canonical-0133230331110111-2123102021201202-1131001222022213-2320110221111102-2233131301100012-0130220210010212-0123020233311231-0331232230012302"></a>

<a id="canonical-0130122331322002-0201012103021012-1033203233002331-1012321023022121-3103130101013311-1130120001121130-1031331321001203-1111330010101011"></a>

## plen property — IPv6 / 333020130312 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-3213001001003203-2000003202202313-0031103333323211-1100233332031312-0031321321001222-2201110133233112-1132221303100331-0303303011213222"></a>

<a id="canonical-2333131011021010-0113101231301003-0323133323322303-1201300332331230-0133213230321332-2010102320011032-3312310123323310-2001110231013312"></a>

## prefix property — IPv6 / 333020130312 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-3033132003213130-1100310201023002-3030203013313023-2312232111131203-2130212021000310-1011203221302022-2030102101011112-1000320332132100"></a>

## Next pages — IPv6 / 333020130312 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-1203331130003312-1310000002301320-1211321230130001-1221023023330121-0213200331011033-2121211213213233-0032222321323120-0213212131320000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031131322122233-2300012220112230-1023202323332220-2200330302302233-2320100133322132-3230323322001323-2131222030022013-2132231130110333"></a>

## ingress_egress_gw.performance_enhancement_mode — performance_enhancement_mode / 212032230130 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-1032103103212211-0330323132031232-2130130000021202-3022212211232023-1110031030022011-1312303003331221-0313311112033111-2133212302300003"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-1223310022313112-3332212113132112-2211202030202022-3031021231333221-1302323030322232-3130120013210121-3131022201123231-2233011102032000"></a>

## Direct properties — performance_enhancement_mode / 212032230130 / 3

- [perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322): complete subsection reference.

<a id="canonical-1013203111002123-0110120221233323-2322303320200021-2123120030103003-3113321322311331-3222030013101022-3003210222201010-3213313001131222"></a>

## Next pages — performance_enhancement_mode / 212032230130 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212111122002001-0303311302133220-1313303210133010-3301212331200100-2110123100003022-0123123330100021-2321122030333132-0210200332121310"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 203200331321 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2020033322200311-1121312133021121-0221213313110011-3210200222112021-0111101223332010-2002112330110111-3103221021133000-2231031003030320"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-0313103302133023-1331301023013213-2133001031312220-0112221222031301-1032123200201010-0110110203010002-2210122330323023-1212203111030132"></a>

## Direct properties — perf_mode_l3_enhanced / 203200331321 / 3

- [jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-3231123323302033-0321303132100103-0022012021213202-1312321132313300-0112001002220123-0321023130000001-2122201001002320-1301210032220013): complete subsection reference.

- [no_jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-3312021210011212-2101333003101013-1100211113131111-3030230313013311-3301021111220202-2113113331231013-1331131132300001-2133321133203302): complete subsection reference.

<a id="canonical-3000120022223001-3311332233031030-2032300230233332-1101222301212230-2010123310011021-2310322102322032-1131211203211130-2311031221231013"></a>

## Next pages — perf_mode_l3_enhanced / 203200331321 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-3231123323302033-0321303132100103-0022012021213202-1312321132313300-0112001002220123-0321023130000001-2122201001002320-1301210032220013)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-3312021210011212-2101333003101013-1100211113131111-3030230313013311-3301021111220202-2113113331231013-1331131132300001-2133321133203302)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3231123323302033-0321303132100103-0022012021213202-1312321132313300-0112001002220123-0321023130000001-2122201001002320-1301210032220013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330333223333030-3313012023320320-3030133111203122-3313031021313120-3303120110020123-3120230213103211-0130302320111111-0220031211132012"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 221330112303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-3332112331210102-0021313211112230-0323231130203220-3110010131003302-2022301213331203-0021121031212322-3302110132212010-3030030001020123"></a>

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

<a id="canonical-1000332122332232-0033312030311221-2101322112032111-2332010312231001-1123330300322133-3233113212031330-1311220321122130-2220203332123103"></a>

## Direct properties — jumbo / 221330112303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111233012321233-1120223101033310-3303332213111333-1111020202102321-3202002030232020-0030020201022230-0111012013000113-1011302013313223"></a>

## Next pages — jumbo / 221330112303 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3312021210011212-2101333003101013-1100211113131111-3030230313013311-3301021111220202-2113113331231013-1331131132300001-2133321133203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310302031301022-1332010210302011-2233221020311212-1322213222010012-1100022123211322-2231033213022213-2222123302123303-0103101031133322"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 331212333301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-0133313021111323-2203031001003012-0111330211111302-0223002121120300-2001103021100310-2211022332110021-0303320313322111-3232231022020332"></a>

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

<a id="canonical-1221330010020123-2022131213032121-0233013120130130-2330220033221033-1212212310003121-0303011102213330-2312012132211203-3110023002323103"></a>

## Direct properties — no_jumbo / 331212333301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123100022020310-1131020022022112-1300000020200223-3201031311233333-1131222320022032-1121013301032103-1023012103330212-2210323211110031"></a>

## Next pages — no_jumbo / 331212333301 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-2303022230012323-2002333203123321-2000200122121000-1200310030311303-1203111200003310-3222311210311122-3233100013221003-0021102200010232)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233023233210302-1022313132210010-1210031003031331-0131210323223322-2211300330033221-3120330030311313-2322133230210033-2110322301333300"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 012100302022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-1232221131203333-1121300302112022-1112103222002323-1133332220030230-1000311212210133-3300232110010210-2121230330321302-2012303133111332"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-2211111213113030-2002032122103133-0311132100020322-1113022320301223-2212102030320102-1210022121101210-2302031301003333-1312001330320031"></a>

## Direct properties — perf_mode_l7_enhanced / 012100302022 / 3

- [jumbo_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-0303003030113120-1210213322213311-2001233121130303-1313103202210103-0223222110221222-3011302013123133-0112222123302210-3131312302331322): complete subsection reference.

- [jumbo_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312211112211222-0230030301110022-1222103123000122-1210021220131321-0111010312331322-2033332323212310-1330121212331031-3131333312101113): complete subsection reference.

<a id="canonical-2330231210303323-2331030300303310-3003123332333132-3202213302312202-2331013132210003-1120002110102332-3232300010202003-3330022200331102"></a>

## Next pages — perf_mode_l7_enhanced / 012100302022 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-0303003030113120-1210213322213311-2001233121130303-1313103202210103-0223222110221222-3011302013123133-0112222123302210-3131312302331322)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-1312211112211222-0230030301110022-1222103123000122-1210021220131321-0111010312331322-2033332323212310-1330121212331031-3131333312101113)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0303003030113120-1210213322213311-2001233121130303-1313103202210103-0223222110221222-3011302013123133-0112222123302210-3131312302331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030212333322230-3120213112312230-0310332220300213-2130102012323111-0110133202332102-2220223000120120-1113203213321203-0200212003211021"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 031120021021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-2021101100013213-2200022123230221-0112303300113132-0203331031310223-1202303201123232-2133110310103222-3302330032202310-3032031133312303"></a>

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

<a id="canonical-0031023232010220-0300211010123200-1012123321333230-0130123301121110-0313331330312213-2230001300312110-1110102103331332-3203101130331300"></a>

## Direct properties — jumbo_disabled / 031120021021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233102031303231-3102210330303202-1203231302133032-0103202313202103-0212002101313211-0130330300331020-1333230111101031-1330010300012032"></a>

## Next pages — jumbo_disabled / 031120021021 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1312211112211222-0230030301110022-1222103123000122-1210021220131321-0111010312331322-2033332323212310-1330121212331031-3131333312101113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112330023101231-3001003123332233-2331132232030220-0232211311300111-2001102223301302-0032311133013021-3203332022131131-0223330002301330"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 133300111001 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-3320123322323001-3302323320333033-0003330030321331-2012103231212303-0121222032123320-3331003213000221-2203210310111310-0200023221221321)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-3130200312330030-0232232322010120-0013130033202022-1202011000100331-1032100231122001-2222320333001122-1230103002231213-2302202301120003"></a>

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

<a id="canonical-0111231010203332-3000221111111032-2121200300031212-3132002211120300-1213312110102002-1211201202330020-1203002111230311-1100211233111020"></a>

## Direct properties — jumbo_enabled / 133300111001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011223301013112-1113212323131231-2200002310113123-2032222031103302-3133031220211322-2011212231200200-3303120222030302-2132220331023222"></a>

## Next pages — jumbo_enabled / 133300111001 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-3322002111011330-0302000022133021-0303320301212101-1020313203323300-0112101002231333-1301000212310021-0132320222231033-1220113212313322)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3201123132331302-0333103222323001-0032022201210111-1113330003231121-1131221312033333-1200020130101331-0020131022023013-2333121222102311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013200211203222-1121323323322220-1100000201123030-3032132122102200-2002020021221210-3320320213020100-0002222301121330-0121201301020313"></a>

## ingress_egress_gw.sm_connection_public_ip — sm_connection_public_ip / 223121000102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-3102001321330210-3331332122231022-0332212331221303-2313333323000133-2231123220020011-1110301032310231-0021331202300310-0012021313033022"></a>

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

<a id="canonical-2123122322131230-3000002103002202-2122011012223113-1231200000130202-2131302232203311-1222211221113333-0213130302031020-1302212333321202"></a>

## Direct properties — sm_connection_public_ip / 223121000102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102310312001302-2311121103001222-1112220300301003-0201213110011002-3031133301000132-1300321311201220-3023231131121201-1302032202031320"></a>

## Next pages — sm_connection_public_ip / 223121000102 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0113111200033103-1311211113112302-0232213232320333-0030201113220331-0200130302231021-1021231221210013-2323112023102313-2101221130133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320221003332230-2103001221023323-0300023103232123-0030021113120110-1102331202110223-2122022130200120-1110211022030300-3001330210333233"></a>

## ingress_egress_gw.sm_connection_pvt_ip — sm_connection_pvt_ip / 001002301320 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-0022110322013223-1131211303222033-1030300022310121-3022120120313303-0312120333111133-1101302131013002-3212203001122131-1000132000313212"></a>

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

<a id="canonical-0032120300311223-0212220123112322-0330101311213203-3331332210202133-0030101213200333-3303103010023221-0320021233201131-2213321022022002"></a>

## Direct properties — sm_connection_pvt_ip / 001002301320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221231210030103-3230020030313110-0213230133103212-1111311121113301-0301202233200001-1010132300233030-1230120101312022-0003121312111133"></a>

## Next pages — sm_connection_pvt_ip / 001002301320 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302313102003112-3312230223333233-1321211003120211-1000022332212323-3002203202101130-1203202112102130-3122330131133133-0121132020233302"></a>

## ingress_egress_gw_ar — ingress_egress_gw_ar / 213102132323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- ingress_egress_gw_ar

<a id="canonical-2232000000230233-1020013233322310-0013320032012323-0013102210010000-2000310221103000-1223100322220230-3111210301320332-2220020201303132"></a>

Type: `"single"`. Computed.

Two interface Azure ingress/egress site on Alternate Region with no support for zones.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-hub_choice": "[\"hub\",\"not_hub\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-2303223100223320-0213311133203022-3130230310030313-1320022322002103-1333110000101200-0320131131101032-2122121030133301-2001330103031203"></a>

## Direct properties — ingress_egress_gw_ar / 213102132323 / 3

- [accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030211311333313-3211311232111311-0310303011112103-2010300202311313-0100133211120231-3111313221022121-3321203311033301-1101101200101300): complete subsection reference.

- [active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310): complete subsection reference.

- [active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2332222211221303-3022232320120201-1131331120200322-1002200101232002-2030221222122013-0102231321010032-2313011133320013-1122203203221210): complete subsection reference.

<a id="canonical-1220031212221300-3022121130032211-2321003122230132-2131200002200012-0132020300012332-1221010212020202-0131033233112000-1122101223333031"></a>

<a id="canonical-0311032202203133-2121320003300310-0122220213320013-2313213113233303-3300323133100210-2120220000321100-3013320232212321-2203213220133013"></a>

## azure_certified_hw property — ingress_egress_gw_ar / 213102132323 / 4

Type: `"string"`. Computed.

\[Enum: Azure-byol-multi-nic-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-multi-nic-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group_inside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-1222233130210323-3312211322001012-2102103131032231-3302223230112331-3322313221013230-1223021320331030-3230302300133000-3111113113132001): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-2320012003131112-2332113233303202-0201001102333331-1323023010122323-0321210213031332-3003130022133113-1203332120223212-3230301211303123): complete subsection reference.

- [forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030123200030233-1233100121323010-3102113203031210-3320110303132003-3310332222232311-2032033311032031-3022031231020123-1310112032033130): complete subsection reference.

- [global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003): complete subsection reference.

- [hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012): complete subsection reference.

- [inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213): complete subsection reference.

- [no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-0122313011113033-1032211210323022-2001323122102232-2100030323111021-1320002130003131-0232100113332232-3211201112331302-3113210102210312): complete subsection reference.

- [no_forward_proxy](data-sources--azure_vnet_site--reference--group-006.md#canonical-2101001023122203-0033102212111023-3321233031301033-0112330111120123-3303230032203232-1311222010033233-0330013132102111-2230120032221332): complete subsection reference.

- [no_global_network](data-sources--azure_vnet_site--reference--group-006.md#canonical-2232112110113101-2323121302332100-1211003200203012-2020123331102113-2011101102020333-0332203211013120-0210311311132122-3213310233001310): complete subsection reference.

- [no_inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2301333331223003-0020022320230012-2200002031110230-1000130230302230-0222210021000323-0103131213302220-2033010212231030-2212010320212120): complete subsection reference.

- [no_network_policy](data-sources--azure_vnet_site--reference--group-006.md#canonical-3331312201030030-2102232110000103-2312000101122210-2323232200100100-3212200013122310-2210133012313013-0001233030022123-2323012323202231): complete subsection reference.

- [no_outside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-1222031321110021-3122213001323211-2212013221020221-3323033000300023-1202211110032303-2200030231102133-3311202032023120-0233210203332303): complete subsection reference.

- [node](data-sources--azure_vnet_site--reference--group-006.md#canonical-0301332221030113-2231320220003100-0121223000301012-0100011110302031-0120322021303223-1103033132113331-0130013200110001-1200020003302130): complete subsection reference.

- [not_hub](data-sources--azure_vnet_site--reference--group-007.md#canonical-1010201301211330-0021302200213021-3110300131303330-0232202210131321-0013312121313011-1023231031113200-1213022332011033-1131121320002302): complete subsection reference.

- [outside_static_routes](data-sources--azure_vnet_site--reference--group-007.md#canonical-3303311101210221-2121200010011123-2221213331221231-1123020303323210-3320300300201030-3003333322123210-1013332111120320-3310002222032200): complete subsection reference.

- [performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-007.md#canonical-0303230233121030-3103201033220111-0233123223101131-0213111321231012-2023310312122003-3323333232113332-1132313300132231-3211130323012233): complete subsection reference.

- [sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-3331210200012303-1313102220132131-1030323122220302-0002301310231131-0130033133132201-0033211323032020-2233020220102133-1112222213002011): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-0202223033000312-0220233201100301-1102001122103313-1020221312131302-2120332313200200-0030012203123033-2021332102103220-3320131231220211): complete subsection reference.

<a id="canonical-0002112203333230-2033301333322331-3313331220102121-0013230331311122-1332330123003313-3110331003000202-2021132133100100-2313320221311011"></a>

## Next pages — ingress_egress_gw_ar / 213102132323 / 5

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030211311333313-3211311232111311-0310303011112103-2010300202311313-0100133211120231-3111313221022121-3321203311033301-1101101200101300)
- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310)
- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2332222211221303-3022232320120201-1131331120200322-1002200101232002-2030221222122013-0102231321010032-2313011133320013-1122203203221210)
- [ingress_egress_gw_ar.dc_cluster_group_inside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-1222233130210323-3312211322001012-2102103131032231-3302223230112331-3322313221013230-1223021320331030-3230302300133000-3111113113132001)
- [ingress_egress_gw_ar.dc_cluster_group_outside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-2320012003131112-2332113233303202-0201001102333331-1323023010122323-0321210213031332-3003130022133113-1203332120223212-3230301211303123)
- [ingress_egress_gw_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030123200030233-1233100121323010-3102113203031210-3320110303132003-3310332222232311-2032033311032031-3022031231020123-1310112032033130)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-0122313011113033-1032211210323022-2001323122102232-2100030323111021-1320002130003131-0232100113332232-3211201112331302-3113210102210312)
- [ingress_egress_gw_ar.no_forward_proxy](data-sources--azure_vnet_site--reference--group-006.md#canonical-2101001023122203-0033102212111023-3321233031301033-0112330111120123-3303230032203232-1311222010033233-0330013132102111-2230120032221332)
- [ingress_egress_gw_ar.no_global_network](data-sources--azure_vnet_site--reference--group-006.md#canonical-2232112110113101-2323121302332100-1211003200203012-2020123331102113-2011101102020333-0332203211013120-0210311311132122-3213310233001310)
- [ingress_egress_gw_ar.no_inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2301333331223003-0020022320230012-2200002031110230-1000130230302230-0222210021000323-0103131213302220-2033010212231030-2212010320212120)
- [ingress_egress_gw_ar.no_network_policy](data-sources--azure_vnet_site--reference--group-006.md#canonical-3331312201030030-2102232110000103-2312000101122210-2323232200100100-3212200013122310-2210133012313013-0001233030022123-2323012323202231)
- [ingress_egress_gw_ar.no_outside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-1222031321110021-3122213001323211-2212013221020221-3323033000300023-1202211110032303-2200030231102133-3311202032023120-0233210203332303)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-0301332221030113-2231320220003100-0121223000301012-0100011110302031-0120322021303223-1103033132113331-0130013200110001-1200020003302130)
- [ingress_egress_gw_ar.not_hub](data-sources--azure_vnet_site--reference--group-007.md#canonical-1010201301211330-0021302200213021-3110300131303330-0232202210131321-0013312121313011-1023231031113200-1213022332011033-1131121320002302)
- [ingress_egress_gw_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-007.md#canonical-3303311101210221-2121200010011123-2221213331221231-1123020303323210-3320300300201030-3003333322123210-1013332111120320-3310002222032200)
- [ingress_egress_gw_ar.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-007.md#canonical-0303230233121030-3103201033220111-0233123223101131-0213111321231012-2023310312122003-3323333232113332-1132313300132231-3211130323012233)
- [ingress_egress_gw_ar.sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-3331210200012303-1313102220132131-1030323122220302-0002301310231131-0130033133132201-0033211323032020-2233020220102133-1112222213002011)
- [ingress_egress_gw_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-0202223033000312-0220233201100301-1102001122103313-1020221312131302-2120332313200200-0030012203123033-2021332102103220-3320131231220211)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312222131331212-0113031132312111-0011021130012100-3021310332122202-1122012032011112-2023002123100011-2032301220113131-0323102130101102"></a>

## ingress_egress_gw_ar.accelerated_networking — accelerated_networking / 220122000013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.accelerated_networking

<a id="canonical-3202022121030133-2202002322301122-1323321123113133-0233200313133300-1111231302101301-3232230121231303-3031300232323011-2212011003210112"></a>

Type: `"single"`. Computed.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-0233332202321030-2112331223133033-3130331301003123-0220112012021220-1012301033012132-2003322003021210-1001302202022013-3121233211133230"></a>

## Direct properties — accelerated_networking / 220122000013 / 3

- [disable_spec](data-sources--azure_vnet_site--reference--group-005.md#canonical-2312230133023213-1200101011031233-3132302202031132-3303033332221102-1300103022223003-0223212321013210-3230110111023213-3122230321022013): complete subsection reference.

- [enable](data-sources--azure_vnet_site--reference--group-005.md#canonical-2333213133010121-1121303013230220-0010231331213202-3003203133113220-3101112112223313-2103333122101332-2303021100200031-2033000002323123): complete subsection reference.

<a id="canonical-0123221331000203-1020010131230033-0330311002233223-2210332113121030-1320133212103032-0333213211110233-3021301100010222-0222031202002032"></a>

## Next pages — accelerated_networking / 220122000013 / 4

- [ingress_egress_gw_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--reference--group-005.md#canonical-2312230133023213-1200101011031233-3132302202031132-3303033332221102-1300103022223003-0223212321013210-3230110111023213-3122230321022013)
- [ingress_egress_gw_ar.accelerated_networking.enable](data-sources--azure_vnet_site--reference--group-005.md#canonical-2333213133010121-1121303013230220-0010231331213202-3003203133113220-3101112112223313-2103333122101332-2303021100200031-2033000002323123)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2312230133023213-1200101011031233-3132302202031132-3303033332221102-1300103022223003-0223212321013210-3230110111023213-3122230321022013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213013211332201-2023130312312100-2020200020122311-2333323102201213-1103021201212313-3331131300123130-0022130022211000-1021330213002033"></a>

## ingress_egress_gw_ar.accelerated_networking.disable_spec — disable_spec / 110101331220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130)
- ingress_egress_gw_ar.accelerated_networking.disable_spec

<a id="canonical-0313230321103000-0122033210003021-0323003012132231-2303322233032003-0332023022300212-2112233121010321-1232021123112222-2132113323301213"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0102021333023233-1003300332001002-2332333021133121-2121011021302003-1333223211321133-1203332313232003-2200013231031000-3031131022031232"></a>

## Direct properties — disable_spec / 110101331220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021112323010113-0230133030310133-0202001101120301-1220022123001032-0213013000211332-1221300111233021-2111331121233310-3300233332013323"></a>

## Next pages — disable_spec / 110101331220 / 4

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2333213133010121-1121303013230220-0010231331213202-3003203133113220-3101112112223313-2103333122101332-2303021100200031-2033000002323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300213201031120-3131112222030021-3023112130211032-2030320120200323-3033123311010132-2232030000201133-3121001021111101-2310303021001002"></a>

## ingress_egress_gw_ar.accelerated_networking.enable — enable / 111000322002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130)
- ingress_egress_gw_ar.accelerated_networking.enable

<a id="canonical-0200003032302132-3200131122130101-3311002000011132-2121003022131110-3322333000302212-1122221222233331-2112022232001103-1233131212031213"></a>

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

<a id="canonical-2200113000201031-2333113212023122-1023213032302312-0131312222002231-0231123100003120-0113022112332220-0100013101101320-1001311103022022"></a>

## Direct properties — enable / 111000322002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300311111121121-0211123001310130-3233221231123303-1313033220230303-0020030123231302-0222001201201201-0021032100031100-0202233012221021"></a>

## Next pages — enable / 111000322002 / 4

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-3230123202320031-0330100100313132-0213310202010211-0302120331330030-0033030130223203-2210103220111122-2032112330131121-1210310103312130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2030211311333313-3211311232111311-0310303011112103-2010300202311313-0100133211120231-3111313221022121-3321203311033301-1101101200101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030001011113303-1112110202102310-3330032221100032-2211020132122112-2231030313130002-3111131320020202-2100120331201112-0130331123012031"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 313332023322 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.active_enhanced_firewall_policies

<a id="canonical-3011233003203021-1133123011222303-3132133203102101-1033101212210202-1310111031231303-1332220132203101-2130113022033213-2033200302210212"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2311232231202130-2331120210121311-2102233001130030-1123320203112123-0111213012010223-1233233313230000-1012321233021331-3101020100233012"></a>

## Direct properties — active_enhanced_firewall_policies / 313332023322 / 3

- [enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-0000301302321122-1202330231110320-3233103333330212-1332311112021010-1232110112102330-1022322220111320-1332020231302331-2021223303110311): complete subsection reference.

<a id="canonical-2101221231320110-2130203320132231-0300111112202202-1021313130233200-0133010200110113-2330131013321111-1112322303123211-2323111303012233"></a>

## Next pages — active_enhanced_firewall_policies / 313332023322 / 4

- [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-0000301302321122-1202330231110320-3233103333330212-1332311112021010-1232110112102330-1022322220111320-1332020231302331-2021223303110311)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0000301302321122-1202330231110320-3233103333330212-1332311112021010-1232110112102330-1022322220111320-1332020231302331-2021223303110311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020311312332212-2332310200231013-3113302100112300-1103230001300021-3123201100331030-2231013332312310-2011323013102033-0132022330201213"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 122200030320 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030211311333313-3211311232111311-0310303011112103-2010300202311313-0100133211120231-3111313221022121-3321203311033301-1101101200101300)
- ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1203120302302303-1332101330013231-1212303000020321-2201212030230033-2202211102103001-0331312122203203-3101000113110212-0020132333102002"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2002332010203122-0331001310312131-3000131301122013-1301132200223133-3121133303103202-0302333302221213-3230010013022113-2123211212303113"></a>

## Direct properties — enhanced_firewall_policies / 122200030320 / 3

<a id="canonical-0223001112001302-2111110330010123-3103122121333013-3030133333210013-3110123001232231-2300032312011103-3001103100303323-2231303113313300"></a>

<a id="canonical-1113110013332020-2020113023012110-3003312202200200-0222311200012100-3323332302323101-2303333112110313-0013333132132202-3112222221220301"></a>

## name property — enhanced_firewall_policies / 122200030320 / 4

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

<a id="canonical-2212121231013122-0030311112201012-2332111320130001-2111201003023312-0032022320322000-3000212231121123-2030212011011321-3120000013313121"></a>

<a id="canonical-3101002130321331-1010203300103113-1002221032120332-3211013010031100-3201131013212123-3332002130103031-3003221232031130-3222321101330002"></a>

## namespace property — enhanced_firewall_policies / 122200030320 / 5

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

<a id="canonical-0113133211100203-2312311331223010-0023201011120122-1002300312220133-2311212030003300-2312331302322232-0133311230320201-2023130011123021"></a>

<a id="canonical-1223020221033333-3212020100222331-0211032321200330-2301013131103210-2111232033322332-3030332322232103-2233031121123321-2223213011122131"></a>

## tenant property — enhanced_firewall_policies / 122200030320 / 6

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

<a id="canonical-2121031102323112-0121310101300200-2021122321000223-3100221310323131-3131133303103330-2102222021202131-2110130103023211-2213320230121223"></a>

## Next pages — enhanced_firewall_policies / 122200030320 / 7

- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2030211311333313-3211311232111311-0310303011112103-2010300202311313-0100133211120231-3111313221022121-3321203311033301-1101101200101300)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111202000330111-1031201323022331-3333101322333323-0112013302233330-3132010131110310-1212120203313332-0311303320011222-0101230001010233"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies — active_forward_proxy_policies / 222130002320 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.active_forward_proxy_policies

<a id="canonical-3301313230221322-2131030233020130-3312113032232010-2333111003033310-1003233101320323-1213111112121331-0100002021230323-0213330102121321"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121001233201023-3211230212322110-2322003133332302-3000123120030123-0111122020130331-3310120010130033-1211011220301221-0220103033200230"></a>

## Direct properties — active_forward_proxy_policies / 222130002320 / 3

- [forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-0202210021000113-2001223101133121-3130300213312123-0022033012200122-3210011213201331-3301213313133200-2330322232100022-3200020010331103): complete subsection reference.

<a id="canonical-1233302210032100-3300010031313113-2011121331331221-2313213113331300-1201133110012100-3332110010311011-0003211321212112-2012011000002203"></a>

## Next pages — active_forward_proxy_policies / 222130002320 / 4

- [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-0202210021000113-2001223101133121-3130300213312123-0022033012200122-3210011213201331-3301213313133200-2330322232100022-3200020010331103)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0202210021000113-2001223101133121-3130300213312123-0022033012200122-3210011213201331-3301213313133200-2330322232100022-3200020010331103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130130121200113-1330310221201121-1110231100202101-3131030033103322-3313313330120220-0113323330102001-0311010211300300-3300033031021303"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 023312201012 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310)
- ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0130311233311032-2223010331012113-0302012321003331-1201303102112311-0002323022102230-3331220113101023-0332110013100201-3101103031112002"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1333323322002311-3020311220213131-1011022031311001-2032122132132020-0030122333121133-2100112210032322-1300011000233013-1031312322320222"></a>

## Direct properties — forward_proxy_policies / 023312201012 / 3

<a id="canonical-0123010012103010-0021031031212322-0111111022003102-1113013330232211-2100013132313222-1210100021023333-1310221311321121-2021102020233012"></a>

<a id="canonical-1232100303113003-3003001330102102-3010333323212221-3230333221232003-0211303232322232-0112002320211300-1103133220331022-0030030310301111"></a>

## name property — forward_proxy_policies / 023312201012 / 4

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

<a id="canonical-3111301011300031-2211202003213311-1030022301033002-3233001230212023-0332101021020113-3212201333331223-0211011230333100-2030011312333020"></a>

<a id="canonical-0012223002002121-2201331223133333-0021210333013211-3303232131032310-0223332010012003-2122220211020113-1132320130122102-3332320230233113"></a>

## namespace property — forward_proxy_policies / 023312201012 / 5

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

<a id="canonical-1131031133232213-1330000030221210-2133311130232112-1113133033200202-0103013333323230-0031323130110232-0200002132130331-1323001122333313"></a>

<a id="canonical-0010210201121213-3110130110311321-0303311222033220-2313030101320303-2111220230303000-1031002023130321-2230101132301012-2122110122202132"></a>

## tenant property — forward_proxy_policies / 023312201012 / 6

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

<a id="canonical-1010233321000121-2211213010212131-3323203031110012-1132121310101210-1110102220312312-0332100120322132-3130211122020223-2200033210232220"></a>

## Next pages — forward_proxy_policies / 023312201012 / 7

- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2332222211221303-3022232320120201-1131331120200322-1002200101232002-2030221222122013-0102231321010032-2313011133320013-1122203203221210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000133013102223-1131322300121312-0131000331300023-1202032202322210-3233330103021321-0231003003220312-0100001211211131-0121101311001333"></a>

## ingress_egress_gw_ar.active_network_policies — active_network_policies / 231303003323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.active_network_policies

<a id="canonical-2000133020113122-1220321102011112-0211201223321333-1303100101322332-0333222312323320-0301331200330220-2131103213032123-1300212232302333"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021011300031031-2310313100332323-0121012212201310-0230132031033212-0013121231302322-2303030102103220-1100002031321301-0022203220113301"></a>

## Direct properties — active_network_policies / 231303003323 / 3

- [network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2010301011210300-3122330012021330-0302000222130213-0323120111030133-1122312022230120-2101111113223200-1110211230112200-3300010120023131): complete subsection reference.

<a id="canonical-2233203223222101-1130310213032121-2321102301203133-1222131330222132-0003132130132120-1022311232031032-0031112001103201-0220221120201232"></a>

## Next pages — active_network_policies / 231303003323 / 4

- [ingress_egress_gw_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2010301011210300-3122330012021330-0302000222130213-0323120111030133-1122312022230120-2101111113223200-1110211230112200-3300010120023131)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2010301011210300-3122330012021330-0302000222130213-0323120111030133-1122312022230120-2101111113223200-1110211230112200-3300010120023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010101033312031-2133231332011203-3203320201100022-3313203302223121-3113110202212002-2122223301302021-1100333013200333-3033220033213010"></a>

## ingress_egress_gw_ar.active_network_policies.network_policies — network_policies / 123002113122 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2332222211221303-3022232320120201-1131331120200322-1002200101232002-2030221222122013-0102231321010032-2313011133320013-1122203203221210)
- ingress_egress_gw_ar.active_network_policies.network_policies

<a id="canonical-2022131111222323-1223303002231130-2332121031130132-3012310322031102-3031222123130033-0312220311220022-0222102021221302-2231312133110220"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0030131322010021-1223213101133311-0011023203132221-0211112131222010-2101120130230313-1121023033001120-1211123303123310-0331200213130001"></a>

## Direct properties — network_policies / 123002113122 / 3

<a id="canonical-2020122212212021-0232102223220130-2300202311002001-3131211223200303-0131003210330023-2302231300130223-2031020200030021-0132032320133010"></a>

<a id="canonical-3302211333301213-1322200103312310-3120232333033022-3233212020033131-0213311202022333-2230113101202220-0312011001001222-0103103121322123"></a>

## name property — network_policies / 123002113122 / 4

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

<a id="canonical-0020121102230303-1301302123133121-2121032213113110-0303313200312231-2303021101002322-2311013003110020-3311102231023212-3133202101301012"></a>

<a id="canonical-3233232232023132-2313033213321010-1001111131202112-0100103002100130-1223121210022131-2011122313211312-2123333112002030-2112031120103323"></a>

## namespace property — network_policies / 123002113122 / 5

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

<a id="canonical-3011232130101233-2000230030130230-2110312130221110-3133031322311012-3321303033002302-2301223213103033-2133223122000212-0233213131330211"></a>

<a id="canonical-3303202103000003-0210331231110211-0101003201322013-2001333023100010-0112002321221300-2202023020022311-3022303002100111-3132320001000223"></a>

## tenant property — network_policies / 123002113122 / 6

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

<a id="canonical-3203013021033311-2110323210100223-2231113103333231-2031010101313221-1000200011023313-2123130022000230-3023122101200003-2031110330031333"></a>

## Next pages — network_policies / 123002113122 / 7

- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2332222211221303-3022232320120201-1131331120200322-1002200101232002-2030221222122013-0102231321010032-2313011133320013-1122203203221210)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1222233130210323-3312211322001012-2102103131032231-3302223230112331-3322313221013230-1223021320331030-3230302300133000-3111113113132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301302203330301-3201102120101001-1223021311333002-3330200131010102-0123332312232213-0331013221101023-3001310031221202-1122021303201000"></a>

## ingress_egress_gw_ar.dc_cluster_group_inside_vn — dc_cluster_group_inside_vn / 200000330031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.dc_cluster_group_inside_vn

<a id="canonical-2021212133022210-3112220122110010-2202210303130032-3123000032220011-3131110333120102-3200103311122221-2122233102031201-2002230021122300"></a>

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

<a id="canonical-0302323220301302-3223320200101323-2013120021133320-2101203133021333-1333023232331020-1231230013023131-0311101302111332-2031103110011313"></a>

## Direct properties — dc_cluster_group_inside_vn / 200000330031 / 3

<a id="canonical-0030103122220033-2230120222103221-1320222113222311-1202301211330112-0322013313131010-0011133110322213-1021231002231202-2010233022030300"></a>

<a id="canonical-2300313123221223-0213121200032113-0123112232303201-3010023310230211-3221121303303210-3123202002231020-2222111013320132-0110331000221132"></a>

## name property — dc_cluster_group_inside_vn / 200000330031 / 4

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

<a id="canonical-2331111310230032-0232221022322110-2020020200022133-3030023030203203-1211023012001000-2201303310112012-1212120212122200-3213022020233211"></a>

<a id="canonical-2101113321112221-0232121231130022-1230233312330311-2303303310110133-0023311210332001-3032013030223312-2320100121310100-1311120300010202"></a>

## namespace property — dc_cluster_group_inside_vn / 200000330031 / 5

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

<a id="canonical-3120212020011032-2202002133033320-3032233021302111-0330113312102001-2011032013103310-1100003003133200-0011213212003323-3301001200210123"></a>

<a id="canonical-3033210211001111-2030310230313011-2122203101333113-3032313102232012-0133231131201323-3333200210220113-3200321220331301-3313113001220303"></a>

## tenant property — dc_cluster_group_inside_vn / 200000330031 / 6

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

<a id="canonical-2221011002023312-2033220301313032-2332302002202130-2121321000121333-3222220331321202-0103031230202132-3112212101003113-1001213213131332"></a>

## Next pages — dc_cluster_group_inside_vn / 200000330031 / 7

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2320012003131112-2332113233303202-0201001102333331-1323023010122323-0321210213031332-3003130022133113-1203332120223212-3230301211303123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300102133323103-1111000133232203-2032003200003200-0332122010103212-0120011233330300-1012000231201132-1023030003103101-0230333212032211"></a>

## ingress_egress_gw_ar.dc_cluster_group_outside_vn — dc_cluster_group_outside_vn / 002031200102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.dc_cluster_group_outside_vn

<a id="canonical-2323310011111131-1330133110210103-1210323010110201-1200130030030102-1213010312031103-1023023133102230-3120231210330232-0030313001330030"></a>

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

<a id="canonical-2320312020333102-2220223321232100-3113022023120010-3233232121013230-0310300101313222-2020222013210132-0023103103100232-1300231011111022"></a>

## Direct properties — dc_cluster_group_outside_vn / 002031200102 / 3

<a id="canonical-1212331100300123-2013232330130102-3003233022130231-2122213303001111-2321101002000222-0210322120031100-3111010211222301-2032233231220023"></a>

<a id="canonical-1202001121112032-0022330123321022-1313321331300001-2323012311233230-1201002233203010-0222131222010130-3132110301332121-1012321003223233"></a>

## name property — dc_cluster_group_outside_vn / 002031200102 / 4

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

<a id="canonical-3001220033102332-2220232212103102-2013213232203113-2013022023101212-3101331222322112-1122132310231332-1131130223102032-1233123100131103"></a>

<a id="canonical-3312101322332333-1301111120232131-2112033232131231-0303032230203310-2232213321301332-3103321303230210-1123130030023001-0320133132231203"></a>

## namespace property — dc_cluster_group_outside_vn / 002031200102 / 5

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

<a id="canonical-0203312120212021-0323011303113113-1131003202220310-1201000022222003-2133001100312120-2002133320122203-3323330130311202-3010300002320010"></a>

<a id="canonical-0300322020333203-3223003313010322-3110130022022200-0200230313331333-3203002000000200-2233022332210100-2131030202002102-1210321311103111"></a>

## tenant property — dc_cluster_group_outside_vn / 002031200102 / 6

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

<a id="canonical-2101123232013303-2031110110021003-0010112220121122-2321233323122001-2222032012322131-3031200100212023-2110223010010130-1223331111101030"></a>

## Next pages — dc_cluster_group_outside_vn / 002031200102 / 7

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2030123200030233-1233100121323010-3102113203031210-3320110303132003-3310332222232311-2032033311032031-3022031231020123-1310112032033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101101102020133-1132020123022230-3110301021131231-2201322222320323-1333322013332301-2213133000031202-1312103120103101-0223111130002133"></a>

## ingress_egress_gw_ar.forward_proxy_allow_all — forward_proxy_allow_all / 102300203002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.forward_proxy_allow_all

<a id="canonical-2330122103303132-0003120323003210-2230121320000120-2303030232321211-1123013122330120-3112002300133123-0220301001001310-0030233313121120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for forward proxy allow all.

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

<a id="canonical-1231302311033000-0013201220001212-2331323023131221-2112131001213222-1323320322110011-2203230233312020-2220320011023121-2213133122021333"></a>

## Direct properties — forward_proxy_allow_all / 102300203002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022113010301101-1301332110233203-3332100121102320-2013111221320230-2312013112301220-2301320100221131-3033222113102021-3032323322013120"></a>

## Next pages — forward_proxy_allow_all / 102300203002 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100112100031301-1313230030122032-0210121302032023-2302123120200310-3201122200102003-3321223301223111-2130111333111013-2211332321331300"></a>

## ingress_egress_gw_ar.global_network_list — global_network_list / 210311023021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.global_network_list

<a id="canonical-0301032203132220-2103022302120010-2022011312222312-3211220203133131-0203112000011130-3320233023311221-1112012333113213-3111123313130021"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031002123102031-3133012100322332-2202321121133110-3223210113232313-2302230221202123-2112132323330013-2223123332002321-0002100002122000"></a>

## Direct properties — global_network_list / 210311023021 / 3

- [global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203): complete subsection reference.

<a id="canonical-3130213220233333-3311122123113232-3202210222003031-2032012230012100-1321123301223111-0022201033230332-3033001232032311-0303020120032321"></a>

## Next pages — global_network_list / 210311023021 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031222231032123-3201203131231030-3321313203011301-0301132301210020-0110232011223002-3311231023003030-0300233033222020-0313201323220121"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections — global_network_connections / 232201110301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- ingress_egress_gw_ar.global_network_list.global_network_connections

<a id="canonical-2302322310020331-1012113200021220-1312202111002313-1220033103131003-0331202311032033-3300303032030223-3300111312310012-3013323223013123"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0103231120033102-2110001122020313-0210010023330032-0121301031003332-2232023232231231-2031113310102203-0120121330110021-3302120002020231"></a>

## Direct properties — global_network_connections / 232201110301 / 3

- [sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-1211200102020212-3011021211103111-2020313202332131-0032110123300200-2101103231220032-1033001131303331-0013122301221233-1102002211102110): complete subsection reference.

- [slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-2032132232111000-3120203213031331-3030222332231330-2201200203102113-2002020323221121-3113300321212210-2011312010330301-2131012303011030): complete subsection reference.

<a id="canonical-2230031033010200-0230221121103002-0221202332102331-3321210223020223-2211131121111210-1030311023211200-3113113111121302-0313231332002131"></a>

## Next pages — global_network_connections / 232201110301 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-1211200102020212-3011021211103111-2020313202332131-0032110123300200-2101103231220032-1033001131303331-0013122301221233-1102002211102110)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-2032132232111000-3120203213031331-3030222332231330-2201200203102113-2002020323221121-3113300321212210-2011312010330301-2131012303011030)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1211200102020212-3011021211103111-2020313202332131-0032110123300200-2101103231220032-1033001131303331-0013122301221233-1102002211102110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303113012203332-2113001132300213-3032101030323113-1013322303032333-1003110311323031-1231023222112213-3310331100010220-3010331320212222"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 023333203330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-2331201132310302-3221013022332200-3102301003221332-0032331123233123-2021200033122212-0002232031212201-0000200112000321-1311113013030331"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310200233303333-0010302230022323-3220132213112301-0300000333000112-3311221112110010-2301213113201002-0303201302200332-2103122033001103"></a>

## Direct properties — sli_to_global_dr / 023333203330 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3312023131010301-1013312302222012-1210012301102321-3331331211310203-1232022033212013-1011230313211113-3222302032121103-2323000131102011): complete subsection reference.

<a id="canonical-2332333111022003-0221232032310203-2103033332020021-0223100320011030-2222203333112123-3031012203203111-0321022003101131-0101100322331300"></a>

## Next pages — sli_to_global_dr / 023333203330 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3312023131010301-1013312302222012-1210012301102321-3331331211310203-1232022033212013-1011230313211113-3222302032121103-2323000131102011)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3312023131010301-1013312302222012-1210012301102321-3331331211310203-1232022033212013-1011230313211113-3222302032121103-2323000131102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022123201010111-2300003101031020-2322300200013301-1010301030333121-3301220323103110-1313230002203322-3323023220320311-1200112301230012"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 031010000303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-1211200102020212-3011021211103111-2020313202332131-0032110123300200-2101103231220032-1033001131303331-0013122301221233-1102002211102110)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-3231131101022210-3221120133102020-2323231031023222-2113301002230123-0330022302103200-3330033131031031-3100012121032100-0232023301212222"></a>

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

<a id="canonical-1310132310013032-1031022102221210-3103230311323132-0301221332001112-1231130110113330-2232312112311221-1131321300032122-0021031013010300"></a>

## Direct properties — global_vn / 031010000303 / 3

<a id="canonical-0222231231021232-1211002121101222-3221203303021222-0002020311113100-2231101321321203-0332021223232000-0222022101313012-3223211010011102"></a>

<a id="canonical-0010001211022303-0103221020032113-2231330220330101-2010002203322212-2122003023211111-1103113122302210-0130123232310330-1322031213312031"></a>

## name property — global_vn / 031010000303 / 4

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

<a id="canonical-0221110111202211-0303333131113202-3211133332002331-3331203210200123-3131123021121220-1001323122301203-0011113033130022-3133020110302200"></a>

<a id="canonical-0000202101011333-0203133332112332-2232120302320021-0130200103012021-2022013220232320-1132102000001003-2210032230122202-1311231213132113"></a>

## namespace property — global_vn / 031010000303 / 5

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

<a id="canonical-3012212110322023-0322121200110113-0221330103010301-1301210003323132-0312133101233101-2122132200033210-0331220312131010-3223210122103003"></a>

<a id="canonical-2231331003113330-3303131022233122-1300310121022032-2312332222002110-2211103203130330-2213312323112212-1001312023230323-0211101033031301"></a>

## tenant property — global_vn / 031010000303 / 6

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

<a id="canonical-2331202112312301-3122122131332300-1203131023223231-3010201310212310-3232003221010331-0200301132133232-2202233220203020-3323300231131122"></a>

## Next pages — global_vn / 031010000303 / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-1211200102020212-3011021211103111-2020313202332131-0032110123300200-2101103231220032-1033001131303331-0013122301221233-1102002211102110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2032132232111000-3120203213031331-3030222332231330-2201200203102113-2002020323221121-3113300321212210-2011312010330301-2131012303011030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123333333103322-1212201223232330-1212303012013120-2131103322230213-1212310230220131-1003322032220110-3032101121233200-3303100303231031"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 011321001303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-0123331123231333-0212311013023001-1011102023303102-3031323213030033-2223310100333001-1222223102113220-1102332010000011-2312013000332113"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0320122311333003-2231013113023021-3122212213010322-3020030011223030-3212123230233201-2330320010201321-1001131023233312-1201211222100310"></a>

## Direct properties — slo_to_global_dr / 011321001303 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3100020203022120-1031231020310232-0311011033032222-3211301102020212-2133132000210333-1321020323221010-2212300233200102-0232233130232331): complete subsection reference.

<a id="canonical-3010003121322233-3222330110230133-0332102321311300-0303201333020130-0323112232321233-3312122103213110-1113332012102311-0111021001303323"></a>

## Next pages — slo_to_global_dr / 011321001303 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3100020203022120-1031231020310232-0311011033032222-3211301102020212-2133132000210333-1321020323221010-2212300233200102-0232233130232331)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3100020203022120-1031231020310232-0311011033032222-3211301102020212-2133132000210333-1321020323221010-2212300233200102-0232233130232331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300002333022231-2322102000223313-3330131010300003-0010221310232331-1332222202321010-1023312310221300-2030110000311020-0003021202012120"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 132330300012 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-1110301233310220-3112203032220000-3212012220021302-0120020002010213-2101133031330312-0112023203300123-2001201102102222-0011101211012003)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-3311021230233203-0000320011223133-1110100202020031-3121010222313321-0021323012012023-0012211123022133-3033203022221323-2012030031133203)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-2032132232111000-3120203213031331-3030222332231330-2201200203102113-2002020323221121-3113300321212210-2011312010330301-2131012303011030)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-2022001330320303-2222111132001100-1000213021311130-3301303200333030-1303221220222230-0130220230120033-3210223031120033-0310001233300121"></a>

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

<a id="canonical-0121212023123320-0023223312231003-3010030023001031-2232321202033100-0201112012113230-1032002331321132-0032211132033322-0203333003011123"></a>

## Direct properties — global_vn / 132330300012 / 3

<a id="canonical-1330320210310312-0200212012123301-1312203123223212-2301132122202330-3302223230303100-1011002213012033-3020211233031012-1011103222222313"></a>

<a id="canonical-3030101102230021-0313111000322200-1310323223102321-3012122303120212-3320300023101103-3130310113021111-1211213201330023-2230232220022123"></a>

## name property — global_vn / 132330300012 / 4

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

<a id="canonical-0100021110220220-3110232031222113-0313111113223012-3222100323110132-2330302222221113-3332110132010212-3022132111020033-2301030312012232"></a>

<a id="canonical-3213033133033013-0233031332321022-3230012020212200-1210312231220231-2023222111021012-1322200231320323-0220111003021022-2033201333103200"></a>

## namespace property — global_vn / 132330300012 / 5

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

<a id="canonical-0331312322111232-0021331310012333-3211021331211030-2223313020113120-3203330210031111-2220311031322313-0111011023210230-2031310021313201"></a>

<a id="canonical-0220211303301100-3322222321302130-2101300020330100-0301012331011032-0301230321020103-0201213233331213-3200102001103031-3120201011200002"></a>

## tenant property — global_vn / 132330300012 / 6

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

<a id="canonical-0003333301331132-1020013233123222-3222103221101010-1310021000011111-3302303313122311-0202200203130000-1001132012103201-0120010031132223"></a>

## Next pages — global_vn / 132330300012 / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-2032132232111000-3120203213031331-3030222332231330-2201200203102113-2002020323221121-3113300321212210-2011312010330301-2131012303011030)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103033301233302-1312332313031101-2020023312131231-3332202110223330-1330123203210122-2112133113023000-0122231110012320-1030020032132230"></a>

## ingress_egress_gw_ar.hub — hub / 113020001302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.hub

<a id="canonical-3301002022132012-2112123111031102-2303121011331230-2120101131213001-1122121112310223-1332321032200213-1012123122201001-1030322030311201"></a>

Type: `"single"`. Computed.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

<a id="canonical-2120111331221023-1332303303113311-3032021323002311-3233022211011121-1033121300000220-1311310312213030-2103302312103030-2201222320133201"></a>

## Direct properties — hub / 113020001302 / 3

- [express_route_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-0011200213301220-2212032213011303-2201110121021011-2030102012022310-1001210120330301-1222302312230330-2331301321013333-2030312110220020): complete subsection reference.

- [express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311): complete subsection reference.

- [spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011): complete subsection reference.

<a id="canonical-2303132212000203-1320231021321122-3012123210111201-0231302023033033-2212031020110120-0120320213202012-3331311013102131-1111112212211113"></a>

## Next pages — hub / 113020001302 / 4

- [ingress_egress_gw_ar.hub.express_route_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-0011200213301220-2212032213011303-2201110121021011-2030102012022310-1001210120330301-1222302312230330-2331301321013333-2030312110220020)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0011200213301220-2212032213011303-2201110121021011-2030102012022310-1001210120330301-1222302312230330-2331301321013333-2030312110220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102002321132200-1101323033013100-1232123311011033-1123002203023020-3122303322322203-1001002120101332-3210120220103102-0031330201032000"></a>

## ingress_egress_gw_ar.hub.express_route_disabled — express_route_disabled / 013001322310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- ingress_egress_gw_ar.hub.express_route_disabled

<a id="canonical-2303130112220220-0303102132100122-3001102221033322-0111102213013300-2302221102200131-1330201221001013-3010302301320311-2321021311030030"></a>

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

<a id="canonical-3313122321213313-1023012010011011-2110202032213220-2331003002223222-2302321113100112-3100023210331331-1310233133223033-1133122323222312"></a>

## Direct properties — express_route_disabled / 013001322310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232320011202213-1213102010220310-3233033003110122-2322213022121101-2132123131320103-1110331220003211-2120000302221203-0011311210022110"></a>

## Next pages — express_route_disabled / 013001322310 / 4

- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332011312020332-0112313310012212-0222312321101133-3122120220013022-1310231001220113-1310230013123303-0000233313310311-2121122313320332"></a>

## ingress_egress_gw_ar.hub.express_route_enabled — express_route_enabled / 020211322311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- ingress_egress_gw_ar.hub.express_route_enabled

<a id="canonical-3313300333120023-3021133110230122-1030033232102030-0100111112133200-2321113232221111-1212211133010121-0301213323202303-1202000031231002"></a>

Type: `"single"`. Computed.

Express Route Configuration. Express Route Configuration.

Upstream description:

Express Route Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_express_route\",\"site_registration_over_internet\"]",
  "x-ves-oneof-field-sku_choice": "[\"sku_ergw1az\",\"sku_ergw2az\",\"sku_high_perf\",\"sku_standard\"]",
  "x-ves-oneof-field-spoke_vnet_routes": "[\"advertise_to_route_server\",\"do_not_advertise_to_route_server\"]"
}
```

<a id="canonical-3200313333200211-2022200201103022-1000022032203303-1031022012301111-0333302121232110-0001313222012230-1112010003311312-3123223122001000"></a>

## Direct properties — express_route_enabled / 020211322311 / 3

- [advertise_to_route_server](data-sources--azure_vnet_site--reference--group-005.md#canonical-1120011011221000-0102300131100323-2230112132032331-2011212321113023-1003211113301102-3131311220320021-1132002123313303-1131123031301311): complete subsection reference.

- [auto_asn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3232203113321331-3303001331022311-0130012321201122-0030210211003233-1031000023322013-2310333002331221-2003320212301031-3301111022313320): complete subsection reference.

- [connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323): complete subsection reference.

<a id="canonical-2210322000111102-0120323122010010-2331331013233333-0312021321331332-0330011111333132-1131321013210102-2311221110223130-2120011103010203"></a>

<a id="canonical-3331031331233011-3102110330311012-3322213110223312-1021032322222310-1132012130213121-0113232203100031-2132210122310312-2021013330021003"></a>

## custom_asn property — express_route_enabled / 020211322311 / 4

Type: `"number"`. Computed.

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Upstream description:

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  }
}
```

- [do_not_advertise_to_route_server](data-sources--azure_vnet_site--reference--group-006.md#canonical-2023033331210222-2230333023012012-2001203323310220-2100322210231333-2212333100032331-1022332113301211-2222220012113010-3123232322303111): complete subsection reference.

- [gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010): complete subsection reference.

- [route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202): complete subsection reference.

- [site_registration_over_express_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-0221210213203222-0012113111230331-2303312003212102-0222102310133123-0310332030312123-1311033112100312-0133112203120001-0200302031023221): complete subsection reference.

- [site_registration_over_internet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3322112221103022-1322032123023103-1233300213112123-2101013332232100-3203120023210100-1131102111120331-3321310121112303-3020232121203332): complete subsection reference.

- [sku_ergw1az](data-sources--azure_vnet_site--reference--group-006.md#canonical-3011100300322203-0212310322003212-0333213120000101-3112233012222022-2210121212200313-0131103002013221-2001133020311021-1013131000230232): complete subsection reference.

- [sku_ergw2az](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222311000330111-3000222023332301-2311333110200220-0123102130323230-1031310102321111-1031000201221003-0130002130113011-1300232321133322): complete subsection reference.

- [sku_high_perf](data-sources--azure_vnet_site--reference--group-006.md#canonical-0211212312303130-0012322203100312-2011200011033121-3011100032023122-0113332120113011-2113302022010020-2201111300101233-1301011121211120): complete subsection reference.

- [sku_standard](data-sources--azure_vnet_site--reference--group-006.md#canonical-1210133331031012-3233002103231131-3133110300112012-2211120200022111-0321110333200100-0203120003322011-3312001001323113-2310132223020220): complete subsection reference.

<a id="canonical-1021021213320101-0320311312013232-3312023203033322-1303330220002013-3321303023023321-1111300223331110-3232121323330021-3013220111002012"></a>

## Next pages — express_route_enabled / 020211322311 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](data-sources--azure_vnet_site--reference--group-005.md#canonical-1120011011221000-0102300131100323-2230112132032331-2011212321113023-1003211113301102-3131311220320021-1132002123313303-1131123031301311)
- [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](data-sources--azure_vnet_site--reference--group-005.md#canonical-3232203113321331-3303001331022311-0130012321201122-0030210211003233-1031000023322013-2310333002331221-2003320212301031-3301111022313320)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](data-sources--azure_vnet_site--reference--group-006.md#canonical-2023033331210222-2230333023012012-2001203323310220-2100322210231333-2212333100032331-1022332113301211-2222220012113010-3123232322303111)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-0221210213203222-0012113111230331-2303312003212102-0222102310133123-0310332030312123-1311033112100312-0133112203120001-0200302031023221)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3322112221103022-1322032123023103-1233300213112123-2101013332232100-3203120023210100-1131102111120331-3321310121112303-3020232121203332)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](data-sources--azure_vnet_site--reference--group-006.md#canonical-3011100300322203-0212310322003212-0333213120000101-3112233012222022-2210121212200313-0131103002013221-2001133020311021-1013131000230232)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222311000330111-3000222023332301-2311333110200220-0123102130323230-1031310102321111-1031000201221003-0130002130113011-1300232321133322)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](data-sources--azure_vnet_site--reference--group-006.md#canonical-0211212312303130-0012322203100312-2011200011033121-3011100032023122-0113332120113011-2113302022010020-2201111300101233-1301011121211120)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](data-sources--azure_vnet_site--reference--group-006.md#canonical-1210133331031012-3233002103231131-3133110300112012-2211120200022111-0321110333200100-0203120003322011-3312001001323113-2310132223020220)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1120011011221000-0102300131100323-2230112132032331-2011212321113023-1003211113301102-3131311220320021-1132002123313303-1131123031301311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300332323022132-2333230231022030-0313112233131131-2232131231032120-0332103013210120-2033203032122323-1202200000113021-0002312312131011"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server — advertise_to_route_server / 230031112033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server

<a id="canonical-2223312131223120-1021030102113312-3203031210322311-2332013120010102-0120111103021212-0010212110331120-2030131301020000-3332222203033200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise to route server.

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

<a id="canonical-1101221002103033-1101221133021231-2221100330123310-1020001032102211-1211232221112313-2020130313300010-1231203323033230-1132333113223031"></a>

## Direct properties — advertise_to_route_server / 230031112033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013303102123021-2303202002322113-3002213022012113-1223301021320110-1212222020121020-2101113123022001-1303313323020313-3100202132110310"></a>

## Next pages — advertise_to_route_server / 230031112033 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3232203113321331-3303001331022311-0130012321201122-0030210211003233-1031000023322013-2310333002331221-2003320212301031-3301111022313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222322300300221-1210320100302033-0120012332223122-0200231301221030-0032323203121132-2111020211113003-2033133301323031-2223002102112202"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.auto_asn — auto_asn / 212100002111 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.auto_asn

<a id="canonical-2301321103030300-0300231031023111-2103121000033112-3332302303101031-3000212033011312-0321330213321312-2110121313200130-2312133231021230"></a>

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

<a id="canonical-3322130032133233-1111222210021112-0122220203100312-1330110111222332-1032130230202020-3203300032103302-3120031002102031-0223133030300113"></a>

## Direct properties — auto_asn / 212100002111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130010022320330-1330222213031130-3223001332112223-1332303012302321-2111000001330302-3033212120212331-3003313323201023-1023301310112301"></a>

## Next pages — auto_asn / 212100002111 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
