---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-2113220200130130-3123332030221021-0110121023312202-1332013312100001-3200022131030233-2333122201222112-2223313223023020-0000222332021001"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 022232131101 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3220110210311121-3022012202112021-3122120111220131-2023200001102101-2311022032322213-3001131023213100-3023213301223013-3101333330023101"></a>

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

<a id="canonical-3222313323132130-0333223031000313-2102011121132231-3113333223333202-2122010001332321-0001120220100123-2101200020031223-2121232022202302"></a>

## Direct properties — IPv6 / 022232131101 / 3

<a id="canonical-3201112131210203-2131102211032122-2102133331002220-0003133032133033-0123303112210130-0003002232333332-3101012033332322-3012212320021323"></a>

<a id="canonical-1203332330013133-0311111112303101-0222010310201213-3110012201111002-3303113113330123-3211000222303020-2320311231233122-1213100232322033"></a>

## addr property — IPv6 / 022232131101 / 4

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

<a id="canonical-0023310302223130-1121130212320113-2331321103132000-2020113000201211-2232313311032230-3002321002023302-0332301323321123-1212021011132013"></a>

## Next pages — IPv6 / 022232131101 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101110112331221-2032030322033032-1302110122320213-2200001310000103-3131023000110033-2030321120131310-1133120303331110-3023320320301022"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 223310020112 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-0211132213222123-3313121302300003-1033302311201033-2101302103331323-1230011112201332-0322112133300301-0331211333132110-2020230101322322"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1321210011222122-1330122312201002-0133010233013001-3222112130311020-0132101213131121-3000123101330132-2220322233111313-2231111333122213"></a>

## Direct properties — subnets / 223310020112 / 3

- [ipv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-1333133010122201-1100233003331212-1200232313020213-0222010221120201-2113130301132022-0233303132311111-3122000223210202-2300011102111003): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-2312131120031202-0100232332000210-1033132122210330-3213300011330210-0132302211101221-3213032023311311-3231001002130233-1102322100133222): complete subsection reference.

<a id="canonical-2110303030310302-2013003321033312-2000212100310003-1111310010101202-0313220331123111-1003032111100021-1301300203120330-0130303233111123"></a>

## Next pages — subnets / 223310020112 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-1333133010122201-1100233003331212-1200232313020213-0222010221120201-2113130301132022-0233303132311111-3122000223210202-2300011102111003)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-2312131120031202-0100232332000210-1033132122210330-3213300011330210-0132302211101221-3213032023311311-3231001002130233-1102322100133222)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1333133010122201-1100233003331212-1200232313020213-0222010221120201-2113130301132022-0233303132311111-3122000223210202-2300011102111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230202311111111-3033032130032102-1320110322102201-0302022313312300-1023231212132332-3123320301100030-3011113302302232-0023233133222011"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 121031122330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-2110132003201010-0123021110032021-0131103030113213-1211030313121031-1210311021323310-0100020101300023-2030200301222012-3021212230223221"></a>

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

<a id="canonical-1002300332320201-3203303132320330-0000320100211201-3013220121031313-0220310021002210-2232133200222202-2233220203312033-2232221200121133"></a>

## Direct properties — IPv4 / 121031122330 / 3

<a id="canonical-3101030202221301-0310210022010030-2212003233000112-2322313321230121-1030132212033302-2200331333111021-1230333232231230-1110032021103213"></a>

<a id="canonical-1023021130000220-3001233221132210-2301022220222321-1222100020011232-1221111010132231-0110320223322203-1113332323010001-1101210113211123"></a>

## plen property — IPv4 / 121031122330 / 4

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1002010002023332-2203202332021021-3220032113232112-2020113002223010-1321131201011022-2133021230122232-1311203333301203-0233332313323213"></a>

<a id="canonical-3211220130012020-1000101022121121-2113210311322002-0233321120033311-1102301320010120-2210222332301232-3123121023111002-2001332311011211"></a>

## prefix property — IPv4 / 121031122330 / 5

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

<a id="canonical-3333330013223110-0030312300223031-3023210323312021-3102301213101301-3100021322220132-2130133233023220-0210000231103230-2320020312033212"></a>

## Next pages — IPv4 / 121031122330 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2312131120031202-0100232332000210-1033132122210330-3213300011330210-0132302211101221-3213032023311311-3231001002130233-1102322100133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100320203310103-1221113011313202-1022133032120322-3101011022020213-2220033100212132-2020013331133103-3132012323033323-3103002030030113"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 020103001010 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1013100231330102-3203031332213110-0103010113231000-2003233002031230-3333110300223113-1210330013211000-0131222103013000-2323113031132130"></a>

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

<a id="canonical-0120030320113232-2331022221311201-3101202132012012-0132120030113322-3132133012121101-2223323013221311-1223221321101333-2012333111001303"></a>

## Direct properties — IPv6 / 020103001010 / 3

<a id="canonical-1332131310302023-2101212321221302-2311322023101202-1023020000013000-2310303010222232-2320330200002012-1303333201210112-2031022312310303"></a>

<a id="canonical-3032231322001331-2203002131210130-2302331313012012-3033120021233112-3301002201003010-3330031000000303-1002201130103211-0030122013030023"></a>

## plen property — IPv6 / 020103001010 / 4

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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1000212111231002-3320232222202000-1100300120121132-0332301013231133-0233322002232023-2100002313300123-0202003130023223-1312212030313033"></a>

<a id="canonical-2033300221132332-0220321033321311-3301321120021100-0201213311320300-2312133322023233-2123300000131100-3331312122211130-1222030013311110"></a>

## prefix property — IPv6 / 020103001010 / 5

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

<a id="canonical-2101101323032131-3121011122220211-1332300011223310-2133110203312200-1222312031233312-0120330302101312-3321112230300133-0300131110020123"></a>

## Next pages — IPv6 / 020103001010 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3310333131011322-0330100020123222-2030322120221110-2111221333212323-0130210132002022-1102312332302113-2031133120302121-0210233110231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110203100013131-0032232121122300-2121130102322221-2221120202233000-2033122103221230-0312310330130033-1011202022201003-0302231310323122"></a>

## vn_config.sm_connection_public_ip — sm_connection_public_ip / 233303322000 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.sm_connection_public_ip

<a id="canonical-1100310330332101-0230332232221303-2122130112022222-1310322111320331-2211022002330202-0322231113130100-0022011100120321-1020200030213102"></a>

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

<a id="canonical-1232301022101213-1331222111113033-1131020232321020-1001313321213200-2012101323000001-1010010301103020-1110312102010032-0200101220020310"></a>

## Direct properties — sm_connection_public_ip / 233303322000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232323222011311-0012000023033320-2322331331232231-0110001212102222-0321102030201012-3100112001213000-0010001011323133-3123101123032111"></a>

## Next pages — sm_connection_public_ip / 233303322000 / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0330213031123210-3033120010312021-0023130021033220-2313022211120033-1111033320230122-0232222032233123-2302100330110023-3000000033133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001031323312130-0012302110011001-0111001223200001-1310031002132201-2320031330231032-2213232010230102-2202120312012323-1321110321320021"></a>

## vn_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 112212000133 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.sm_connection_pvt_ip

<a id="canonical-0003013021032322-2210300002230132-1231002011213030-3122001130312031-3131201223003311-1302201313021331-1223322112022211-1320320301021323"></a>

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

<a id="canonical-1301103020102230-3212113233130131-1233011302311331-0031012102200332-3331100020233333-3102013121331321-0132221210110022-3103123122230023"></a>

## Direct properties — sm_connection_pvt_ip / 112212000133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332220102102102-0130011013230001-2200032031013133-3001303332212022-1203311123111012-3010310220133011-1110213130233112-2301221221120321"></a>

## Next pages — sm_connection_pvt_ip / 112212000133 / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302022011112110-3123110111333110-3012132222233300-1012330220203101-1300001301230103-1010233301120310-0020233031310133-2233101023131110"></a>

## vpc_attachments — vpc_attachments / 130232323232 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- vpc_attachments

<a id="canonical-2032121101122220-3001123200110112-1100220002312213-0333230202320131-0130022322132321-3303232323223223-1023210032011101-2130332210311003"></a>

Type: `"single"`. Computed.

Spoke VPCs to be attached to the AWS TGW Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301200010133313-0310322113210102-3030000231332222-3312332121102032-0010222222203021-2031210131110233-0301122303330212-3033033021101033"></a>

## Direct properties — vpc_attachments / 130232323232 / 3

- [vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020): complete subsection reference.

<a id="canonical-2210030030120010-0103201221310222-2221301320102023-0222010231112312-1211133233320311-0111131213033333-3012101330321332-0002222201231013"></a>

## Next pages — vpc_attachments / 130232323232 / 4

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333131211320120-2212200312113002-3333210203330330-3200023231110011-2131331330020320-0001211113231033-1103130123023232-1032000211233023"></a>

## vpc_attachments.vpc_list — vpc_list / 301330232133 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303)
- vpc_attachments.vpc_list

<a id="canonical-2200213132321220-0322000303311301-1201222033223222-1311001311321011-2310321200330030-0203210200102312-0202112230320030-3112323011313012"></a>

Type: `"list"`. Computed.

List of VPC attachments to transit gateway.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-2303013102113311-0122123301031022-2223131222101133-2031330233331231-0022212301313332-1220021332320331-3110110311013011-2033022312311012"></a>

## Direct properties — vpc_list / 301330232133 / 3

- [labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-1330000012313312-1331303300303332-2012220031130103-2103102121312023-3031002023201112-3333231211231021-1221010000211212-1123320111301031): complete subsection reference.

<a id="canonical-0003020000210021-2302311213221331-1222202130232200-2313023213123201-2233301033102213-2210001200131021-2033003301033323-1213122200323301"></a>

<a id="canonical-2300020020321122-1022120310113203-0232201332331221-3122010100230131-1033203222010123-1302222201032213-1221203211102020-3123121202131331"></a>

## vpc_id property — vpc_list / 301330232133 / 4

Type: `"string"`. Computed.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-0231122031300111-3200230212102330-2131323010133111-3101203020021001-2133112032000210-1010001023231233-1003002111031012-2023320112302002"></a>

## Next pages — vpc_list / 301330232133 / 5

- [vpc_attachments.vpc_list.labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-1330000012313312-1331303300303332-2012220031130103-2103102121312023-3031002023201112-3333231211231021-1221010000211212-1123320111301031)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1330000012313312-1331303300303332-2012220031130103-2103102121312023-3031002023201112-3333231211231021-1221010000211212-1123320111301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123131332301213-2002233100223021-0232110332102112-1332310301112013-2113322303333023-0233313302321330-0320232300112020-2010100013030101"></a>

## vpc_attachments.vpc_list.labels — labels / 101121330201 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303)
- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020)
- vpc_attachments.vpc_list.labels

<a id="canonical-3111123302102312-1011330022033313-0123230031120033-3200002230203112-1020332211232000-2021203001003301-1321003320311301-2333022111203013"></a>

Type: `"single"`. Computed.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303131223110001-0113130310033003-1310302131101030-1313222312322322-3212112001101211-3333320102302203-0202232233100222-0201230310120131"></a>

## Direct properties — labels / 101121330201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030321230220013-3011220110102113-3311220231211301-1031322112021030-2030330121202100-0203322332003333-1002011311013001-2111101120321230"></a>

## Next pages — labels / 101121330201 / 4

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320110232032220-0221000230010231-3032102022203210-0320123201022231-0111203213000331-3020310203332300-0201210032001113-1300330200222113"></a>

## waf_signatures — waf_signatures / 022213110030 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- waf_signatures

<a id="canonical-2002121323312202-1100332113202321-0200320322132330-1320113233311320-1103132121103200-3321312301302332-3101302103212110-2021233200122010"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-2322232102333121-2303022011000311-2102323313211211-1310102213310201-3210001223123021-1302333023033222-2102222102011100-2133312233010300"></a>

## Direct properties — waf_signatures / 022213110030 / 3

- [automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-0313312313022101-3320033033012311-1003310022120103-1131230031302300-2030321212200323-1222331233013130-2110130123101032-1022022201032110): complete subsection reference.

- [manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-3233121232010331-3030013302023002-0130221132123031-1311021223032020-2330032230332123-0303313111231231-0322313222311202-3030313211322022): complete subsection reference.

<a id="canonical-0323131021033010-1021210103220200-2310322202233203-1321331331001203-3130211030112000-3120220112023333-2023302220100202-1311102023200200"></a>

## Next pages — waf_signatures / 022213110030 / 4

- [waf_signatures.automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-0313312313022101-3320033033012311-1003310022120103-1131230031302300-2030321212200323-1222331233013130-2110130123101032-1022022201032110)
- [waf_signatures.manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-3233121232010331-3030013302023002-0130221132123031-1311021223032020-2330032230332123-0303313111231231-0322313222311202-3030313211322022)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0313312313022101-3320033033012311-1003310022120103-1131230031302300-2030321212200323-1222331233013130-2110130123101032-1022022201032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031112313001211-1313322130002233-1210213301020030-0200112032210332-2313003222110203-1132301120133330-0220110110011121-0133101300123323"></a>

## waf_signatures.automatic — automatic / 233010132011 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- waf_signatures.automatic

<a id="canonical-1303233131211010-2113312203131111-1201110023200102-2101320311203333-3311033011212212-2001012312212311-0322322101200002-0303323033101201"></a>

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

<a id="canonical-1303302111113333-1210032230012132-0233100211212231-2102313310323202-0020322023221321-0212010020033332-3000100210332121-0213120231231300"></a>

## Direct properties — automatic / 233010132011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113211133030221-2303303131133232-2021202132030100-2301300113303323-1121232003302313-0120200001031030-3102030020120300-2212033331333100"></a>

## Next pages — automatic / 233010132011 / 4

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3233121232010331-3030013302023002-0130221132123031-1311021223032020-2330032230332123-0303313111231231-0322313222311202-3030313211322022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020103332313202-2122233233232323-1232301121333120-2130230011013323-0322003331112312-3013103032130020-1203322002203020-2223330210301303"></a>

## waf_signatures.manual — manual / 212312031130 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- waf_signatures.manual

<a id="canonical-1202003002301220-3032303103333101-1212030300222032-2330000212022123-0130012232212100-3331210133113330-0202132311220223-0103123211201313"></a>

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

<a id="canonical-1020110101013333-1321201020331111-0121033012030312-2331131000212221-1332100023221130-2011320012030023-3300013020322112-0133122003222022"></a>

## Direct properties — manual / 212312031130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113032012332311-2113332030010202-1003312332111232-0200100030333333-2321220301300323-0002131331202223-1210102103020103-3122233231221003"></a>

## Next pages — manual / 212312031130 / 4

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
