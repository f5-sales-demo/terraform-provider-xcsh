---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-1022302033130123-0031311121222213-3023330000333111-2023320112322230-0120320011331133-3211100023112210-2033102001011302-1122230223220101"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 103001310311 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-3211232202201232-1031300221322332-3332100103100122-2313323303000230-3321130303320310-0121203121030022-1301200213222033-0120003101331032"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110101102013333-3120120203223011-1310313001002223-1023113231032323-2303222201111100-1120031222220223-1131332113221003-2002131133311030"></a>

## Direct properties — nexthop_address / 103001310311 / 3

- [dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323): complete subsection reference.

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-3012220220211011-0101111003200313-3103010211022000-1102023331221330-2310230232020330-3222202000331100-2020121310121311-3031332020021120): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-0102212312123232-2331013133111011-1130302133131120-2000330101011122-2020303120011212-3032111211221211-0002122112000302-2133201133212311): complete subsection reference.

<a id="canonical-2111133223022033-1131120111212111-3000310121233033-2023120002132200-0210010033012303-1121313320321202-2233322313303123-0312300123231002"></a>

## Next pages — nexthop_address / 103001310311 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-3012220220211011-0101111003200313-3103010211022000-1102023331221330-2310230232020330-3222202000331100-2020121310121311-3031332020021120)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-0102212312123232-2331013133111011-1130302133131120-2000330101011122-2020303120011212-3032111211221211-0002122112000302-2133201133212311)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201211302123323-0101133203330221-1102223222320013-3013031213323303-1113332010113202-3023032121021023-2000311103022231-0003320322131122"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 213230000211 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2010303032211013-1130002131031332-0110231002123323-1121003031220200-0120130301132312-1220212303220102-1133330300211303-2312201313101110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200113313011013-2220020211200032-1133330313303233-2210200301001103-2311323103132231-3230203321331021-0233023111203021-1301303001011212"></a>

## Direct properties — dual_stack / 213230000211 / 3

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-2301101203002112-2201022211203301-3113213310122323-3331220132331233-0010103232211230-2303222131033301-0002102123130031-3122112332230233): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-2200211323110000-3103123202121323-1102320321331321-0123013210110321-0111211003020121-3202001130111203-3321211030002101-1033330203213223): complete subsection reference.

<a id="canonical-2113121022132333-2000333323130313-1001012003303331-0021100311213303-0212002003130312-3221302330233013-0233233010100122-1013010212220302"></a>

## Next pages — dual_stack / 213230000211 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-2301101203002112-2201022211203301-3113213310122323-3331220132331233-0010103232211230-2303222131033301-0002102123130031-3122112332230233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-2200211323110000-3103123202121323-1102320321331321-0123013210110321-0111211003020121-3202001130111203-3321211030002101-1033330203213223)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2301101203002112-2201022211203301-3113213310122323-3331220132331233-0010103232211230-2303222131033301-0002102123130031-3122112332230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210020212221211-1223030213201010-1031202230032012-1102211012103222-3331020313332110-0021013011200130-3030021331233101-2333033332013022"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 322312323001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0122332203100202-1323121013131003-1120331030123202-1033001023032020-1210212323133000-1031313230332330-3221130333020102-2303003332132000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223102133323002-2022313103002012-0102230132021321-3333111132120202-3111133221013233-2300220121223130-1001302100021223-2133300022032320"></a>

## Direct properties — IPv4 / 322312323001 / 3

<a id="canonical-0132030300320113-0110132132331220-2022302221133112-3303312013130320-2332123312301202-0302212311213000-2303003230011013-3110233211133113"></a>

<a id="canonical-0022122112022030-0302111321100110-0022321331220233-2113102302330232-3032103102121031-3202303230212000-0220020232020130-2301003031132131"></a>

## addr property — IPv4 / 322312323001 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3310203320110313-3113213331202133-2222001303121031-1023221022111222-3311231121311111-2030033000311212-1031103231322012-0330033032312313"></a>

## Next pages — IPv4 / 322312323001 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2200211323110000-3103123202121323-1102320321331321-0123013210110321-0111211003020121-3202001130111203-3321211030002101-1033330203213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230322123103001-1113033010021020-1012103313031310-2321032231220010-1113032103300221-1011032213311112-3133133020033021-2120133133130030"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 021300233300 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3103021333233110-1011210131322000-2123321202132102-1033132312123001-1332332020200203-3113112232100312-2202200210302311-0130210312320211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212311233100201-0330010300200203-3003201200302231-3311022130211023-2202230023303220-1311331113331230-0010131331203220-1222132323112320"></a>

## Direct properties — IPv6 / 021300233300 / 3

<a id="canonical-3000222210112033-0313011023210122-2101122210233303-1102323223021113-3010010101022002-2302113201110111-3130020112301230-2020301312032030"></a>

<a id="canonical-1022201202103301-1332222100221222-3011311000121133-2031201022302212-2002010103031320-3302200103030210-0230212213122213-0311333301020002"></a>

## addr property — IPv6 / 021300233300 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3203310332010302-0132330111320201-3301031132103203-0212030122320002-0331101132112113-0211333103231131-1010023300012333-0320313212330001"></a>

## Next pages — IPv6 / 021300233300 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3012220220211011-0101111003200313-3103010211022000-1102023331221330-2310230232020330-3222202000331100-2020121310121311-3031332020021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300030231221200-3032331021210320-0132131133031301-0022320320301001-2303021013210203-0103110302122102-0010200322233322-1222223121013200"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 100312103333 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0022102320330031-1301233021001322-0133021130031201-1313122222013012-1133113200001310-2320302220302023-3312133032223203-0233122221012212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321201031231032-2330012220102132-0010303232301301-1331320111133321-0031320201232100-3101121132101310-3000322111332131-2223331303013110"></a>

## Direct properties — IPv4 / 100312103333 / 3

<a id="canonical-3012032120033230-1002320131101230-0302310110031322-0222123111033310-1130011333121033-2331211001320323-2021223223010332-1200133300013111"></a>

<a id="canonical-3333003200113103-2022322222202333-0000231003311203-0123200030110003-3110211222221220-1131220310331231-2223013222022000-2120230312202123"></a>

## addr property — IPv4 / 100312103333 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0310230301101302-0032310211123133-2003002111210320-1213030211320222-2012103202212311-2001112102102302-0110301012012023-3312231023001303"></a>

## Next pages — IPv4 / 100312103333 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0102212312123232-2331013133111011-1130302133131120-2000330101011122-2020303120011212-3032111211221211-0002122112000302-2133201133212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220120121100330-2232323100222330-2120220032330220-3110302203013332-3230230221312123-0031123220122012-0131222130133223-2303231222333033"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 101021321333 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3110230121213301-2031211111333120-1223001300021321-3210231131001300-2222113113003332-2221101013021333-3132203201131120-2001231132120122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313113333320301-3303223211212023-0223212321313120-2332012031011022-0333133322332232-0010212121213122-2023130002011121-3212311132020112"></a>

## Direct properties — IPv6 / 101021321333 / 3

<a id="canonical-3122323132320201-3302232002130000-0110111200113013-1022010232320320-2003321102313023-2131003022312220-0033023222122033-1221100021023201"></a>

<a id="canonical-0310220033010323-3323323000121320-0321002031002222-3232333221212300-0210210031032113-0220202331031221-3023203133102313-2112310203222101"></a>

## addr property — IPv6 / 101021321333 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3233323021020021-1132030222300003-2322332333213200-1302220200223023-3202200113100301-3210302300322130-3223221113013221-1311311021123323"></a>

## Next pages — IPv6 / 101021321333 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020233300302330-1020312123122231-2120103213131200-3003030222233311-0113021321113333-1331123030310013-0031312213003012-0110333312203111"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 322103330333 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3033333132110230-0101321032301313-2200333332220220-3330322331100021-3221102322011032-2113022010001113-2220330301031031-3232030221323301"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202331100102320-1210013221022323-0210133313131022-2211201012121230-2233313010111013-3130102022013021-0201210301200013-2112112100003221"></a>

## Direct properties — subnets / 322103330333 / 3

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-1100132202210331-2030031320112010-2231332212322120-3223102330101330-0301101201211030-0312223003210130-0033321303211333-0223103020031213): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-1022033101311020-1133132121221020-0101021213303113-2332100001221011-1311202021023223-0303103210313230-3102301121021330-3202323332110100): complete subsection reference.

<a id="canonical-3211111302301010-0220303320013100-2201322030113002-0231323021200332-2101301213213333-3032323130132031-1211220301130320-0331103200233111"></a>

## Next pages — subnets / 322103330333 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-1100132202210331-2030031320112010-2231332212322120-3223102330101330-0301101201211030-0312223003210130-0033321303211333-0223103020031213)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-1022033101311020-1133132121221020-0101021213303113-2332100001221011-1311202021023223-0303103210313230-3102301121021330-3202323332110100)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1100132202210331-2030031320112010-2231332212322120-3223102330101330-0301101201211030-0312223003210130-0033321303211333-0223103020031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210023012222010-3103101010301311-1212031211330212-3133002122000100-2313213102301133-2023230002021120-0012030000002100-3011002200013332"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 331220210111 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0323231030102312-2031232301232010-0220222200330232-0231223122131211-3213023310322223-0313313013310030-2301132112303002-1120303112133022"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033213033330233-0321223210100031-0320103211110000-3010123213210303-3013013200301122-2131202333003013-2130020201333211-0200222321130021"></a>

## Direct properties — IPv4 / 331220210111 / 3

<a id="canonical-1232010211121032-0232021230333032-1022300220202133-3323230101203020-3011111101320201-0201313213303000-1333310333033223-0310100330212221"></a>

<a id="canonical-3202211031312220-2100012322321132-2232211220110023-0202200003003230-2002310311203031-2131230112331020-2103121232201121-0022101212322123"></a>

## plen property — IPv4 / 331220210111 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

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

<a id="canonical-3103022313302202-0002120301011210-3010302030010020-1012013123113010-0123300332121311-2300003102102120-3320212010212200-0013310312101123"></a>

<a id="canonical-1001330103321102-0331232003212111-1011313100212323-0131311222010031-3311202122232132-2031203102331221-1311232320001200-0323333112301112"></a>

## prefix property — IPv4 / 331220210111 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2021023031001213-3112012312031223-2121011331103120-3301301210130231-2312222122111313-0112121012112312-2111323302003031-3232010013133210"></a>

## Next pages — IPv4 / 331220210111 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1022033101311020-1133132121221020-0101021213303113-2332100001221011-1311202021023223-0303103210313230-3102301121021330-3202323332110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310130030322333-1221112332220221-2230122210100312-3113203032300232-3032001311321133-2222313311130110-3111112033331222-1200123030202011"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 310310021200 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1023232322123303-2120112113232100-0103223303310332-1313030210220223-0032302220033330-2330203030211331-3333100021200202-3112121010123032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020301221123221-0321110011212331-0223011121232000-2110301000032322-3133200033131123-2222012031001011-0222013202111211-3011000210003313"></a>

## Direct properties — IPv6 / 310310021200 / 3

<a id="canonical-0132122202103201-0211030311300202-1132011203322003-2220101232103030-0110020121313133-2210002100223302-1203103120023210-0300123021322323"></a>

<a id="canonical-0303213321000333-1013110232023203-2100101230213303-3120221320001310-1120301103311210-2132332220322003-0000102101320121-3332210210213032"></a>

## plen property — IPv6 / 310310021200 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

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

<a id="canonical-3111311323101023-3311301333110110-2330200200323033-2120232102212013-3333301203210010-0223001013321213-2111032000323001-0320313323013221"></a>

<a id="canonical-1123012301330113-3022231000201030-3033123130030211-3201100303001122-2102000320113011-1102331123332112-3012001101122200-2031220012112322"></a>

## prefix property — IPv6 / 310310021200 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3011031203000103-2222322200311001-0311003221032110-0202032323322310-2032203313131301-1103033320203010-2000100121131312-1302000211302210"></a>

## Next pages — IPv6 / 310310021200 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3303020033221031-0121001232113232-0331101030300303-3303101011322302-1303033333020021-2321013112020223-0312203301333130-3001033131210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210001201333212-2313032020211321-2000321120203320-3300322332311221-2203023121013020-0300220323123031-1131213033023201-0002200122100232"></a>

## vn_config.sm_connection_public_ip — sm_connection_public_ip / 121130311203 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.sm_connection_public_ip

<a id="canonical-1202210310030101-1012103013201211-1002001200202022-3330223100321212-1113102330110213-1102313323310013-3212020313210310-2211230323011322"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-3223030232033122-2331320121010020-3112330231302121-0233323002323032-0330220203220001-2331302200010221-0011212022133303-2100111321300112"></a>

## Direct properties — sm_connection_public_ip / 121130311203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000311313133232-3231202022001232-3133030310303221-2122201322020313-2321033200003130-2331003302302220-2220022132221213-3033332322202310"></a>

## Next pages — sm_connection_public_ip / 121130311203 / 4

- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0021012122100010-3002323221023231-2002011020202100-1001130223230002-1100232013102230-1031022221333310-1130233020000232-2303021330003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301011112131221-2121321330313111-0000302110101223-3212333203302302-1330030133311001-3131230313301002-1210002301302110-0131223300203000"></a>

## vn_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 203023203212 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.sm_connection_pvt_ip

<a id="canonical-3001222103030002-2113002333332213-3101301320231220-2101231001300030-3330022330113332-2130013113113300-2122332310131203-3002112231320021"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-1033132303331111-0103113111021132-2102321032210320-2202313022111021-0333100002012231-2323330110311102-1302210130122023-2103322033101130"></a>

## Direct properties — sm_connection_pvt_ip / 203023203212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033220212132033-0333230323203201-2203002012123221-2332301011310100-0203220202100313-2311232300232122-2123111130030111-2200213302012230"></a>

## Next pages — sm_connection_pvt_ip / 203023203212 / 4

- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322031103121310-3300011133312033-1021021201032323-0323320232220313-3123301130001223-1211021030313112-1120110021101111-0022012023130102"></a>

## vpc_attachments — vpc_attachments / 003311020221 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- vpc_attachments

<a id="canonical-2330320031031130-0223131233013102-0031030323103312-1323110303231202-0032131311202231-2300322232210321-2211310022213121-2021200310311132"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vpc_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102010203032211-1111311030200010-2003132123122201-3312023212123110-1031130213000332-0233103002013332-2032102202131303-2100113301220311"></a>

## Direct properties — vpc_attachments / 003311020221 / 3

- [vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302): complete subsection reference.

<a id="canonical-3003130021001321-1320003101301302-1110231201300130-1020000323232221-1223221102301011-0302230123033201-2001123213313203-0001233120333023"></a>

## Next pages — vpc_attachments / 003311020221 / 4

- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021213000033012-1301012201022010-3302032312233212-1320121131032233-3023313030222310-0130121130230123-1213332121333202-0013110012112301"></a>

## vpc_attachments.vpc_list — vpc_list / 222120123031 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303)
- vpc_attachments.vpc_list

<a id="canonical-1210321101230210-2200013012202212-0211221230322001-2121112201021310-3211233213301020-2002023033112131-3133321132020220-0222221111321220"></a>

Type: `"object"`. list nested block, Optional.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012312123331212-1231231111302003-3010113110312201-1022033202030222-3231123232311323-1213000222210231-2120212011020122-0102230231301310"></a>

## Direct properties — vpc_list / 222120123031 / 3

- [labels](resources--aws_tgw_site--reference--group-004.md#canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002): complete subsection reference.

<a id="canonical-3200001003301011-3223122010122023-1300221302312301-2031311121211021-1112132231112213-0300212231022323-3302203123022121-1010223002332223"></a>

<a id="canonical-3333111230130100-0010022122012102-2203013233112100-1333302213110330-2333233310223010-0203313203330131-1223021301202100-1021212132201310"></a>

## vpc_id property — vpc_list / 222120123031 / 4

Type: `"string"`. Optional.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3131003232030313-0223313012002130-3230302202322031-2022302301231013-2223011000210023-0120221011321201-1302020023121121-1232112003013001"></a>

## Next pages — vpc_list / 222120123031 / 5

- [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--reference--group-004.md#canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303131220131111-0312130231330013-1121310221032201-3030313033010301-3133030033011120-2122131103030020-0333312202232230-0030010233230221"></a>

## vpc_attachments.vpc_list.labels — labels / 312020231322 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303)
- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302)
- vpc_attachments.vpc_list.labels

<a id="canonical-3210011200110111-1123303222023031-3121103133003212-3301201031101331-2302321202100113-1102100130033303-2110013002113232-1233213302200202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-3320230101233231-1232103320121130-2003330231120002-1302020200222203-3310000310103131-2020120110333233-0111030113332213-1131001330233203"></a>

## Direct properties — labels / 312020231322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223123203200322-2231130110112230-0311103131133123-1111103311303333-1022111331110001-1000321333202020-1321231233331033-3221113310230210"></a>

## Next pages — labels / 312020231322 / 4

- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120232220013203-2121020233033000-2201313220222131-2011013230110022-0013033103101222-1201032130311223-1102133013321233-3110010010112123"></a>

## waf_signatures — waf_signatures / 120010222312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- waf_signatures

<a id="canonical-3222213203010330-0133312022232213-1321000200323331-1122003320120230-1011033320302333-3322120000321332-0231130203100112-2321011322233231"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321020120031023-1201010333323210-2033003213313312-0103020302302021-0102300312320120-0322322111021113-2330131212130013-2331330301103201"></a>

## Direct properties — waf_signatures / 120010222312 / 3

- [automatic](resources--aws_tgw_site--reference--group-004.md#canonical-3232131113232120-1030302320302203-0212112203333322-1200310110032030-1232102131102032-1333003102012002-2202203111032311-0020021100203011): complete subsection reference.

- [manual](resources--aws_tgw_site--reference--group-004.md#canonical-3131311212203203-0323201303302230-3101003023223233-1001210013131211-2110130201001003-3302131020220333-2201200202321013-1020212313120130): complete subsection reference.

<a id="canonical-2202211302131323-2313122312302133-2333330332323002-1301022100223101-0130303003230101-0323313301110310-2133120023133002-3032101132232001"></a>

## Next pages — waf_signatures / 120010222312 / 4

- [waf_signatures.automatic](resources--aws_tgw_site--reference--group-004.md#canonical-3232131113232120-1030302320302203-0212112203333322-1200310110032030-1232102131102032-1333003102012002-2202203111032311-0020021100203011)
- [waf_signatures.manual](resources--aws_tgw_site--reference--group-004.md#canonical-3131311212203203-0323201303302230-3101003023223233-1001210013131211-2110130201001003-3302131020220333-2201200202321013-1020212313120130)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3232131113232120-1030302320302203-0212112203333322-1200310110032030-1232102131102032-1333003102012002-2202203111032311-0020021100203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011323311323011-3222302012111003-1301000202200330-0110310030023013-2230113130032013-3231123023132223-0333220302330330-2133002313330221"></a>

## waf_signatures.automatic — automatic / 110021303020 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- waf_signatures.automatic

<a id="canonical-0313330010000201-2032202213233101-2201101031123001-1023232213222320-2013230313222013-2101323121320132-3031123323200121-2030232013223213"></a>

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
automatic = {}
```

<a id="canonical-0003322000130321-1130000221323302-2021101301231021-1003031011113311-0021121221331233-1012122312031213-3313120233221031-2331202211200010"></a>

## Direct properties — automatic / 110021303020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322022330202220-3103320133201300-3233103302220331-1212111002323213-3321022030130221-1113133110210133-0000333130032121-0021230300221030"></a>

## Next pages — automatic / 110021303020 / 4

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3131311212203203-0323201303302230-3101003023223233-1001210013131211-2110130201001003-3302131020220333-2201200202321013-1020212313120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301202110323012-3333020232200232-2103203303110203-0023212103202323-1202030010011013-2123023131020132-1133302110102221-0123303302203310"></a>

## waf_signatures.manual — manual / 333021322103 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- waf_signatures.manual

<a id="canonical-0012101302001113-2310030110220330-1112010323013232-0102111022123101-1032033001110320-1222021331023312-3302323013130322-0211021220212203"></a>

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
manual = {}
```

<a id="canonical-2103100012332131-0211033003232132-1201302221130231-3132301003013013-2033111120032100-3332123033101033-3112211323223201-1321021103211033"></a>

## Direct properties — manual / 333021322103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310322000112023-0210012211211031-0222231023023302-0032012213311120-0212012302333331-2330323322103030-0332211123233101-1213201220212230"></a>

## Next pages — manual / 333021322103 / 4

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
