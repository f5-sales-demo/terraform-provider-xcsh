---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-0213231110010001-2331101103232102-1111022032201121-0311030222301202-3232210332210120-0321012222032102-0331311110302011-2301000230103331"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 111131002002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1131331320100030-1223211103020112-2322302212013031-0200122302023112-1330103220213230-1202032031133220-1201333323101000-3021130202020101"></a>

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

<a id="canonical-3122333033231033-3112111330200300-0012230330330022-3310033200113202-2300222001300202-2231113331132112-3013320211301101-1233301020212202"></a>

## Direct properties — IPv4 / 111131002002 / 3

<a id="canonical-0210033310100010-0303030222011013-1021001300103322-3012131222113331-1203212223102032-1320103233101323-1303332312303101-2210323030200130"></a>

<a id="canonical-1302310012301000-0030332203020303-0300201303230313-0303100213230301-0120001012332112-0201231131123031-1111200120102213-3322012000000300"></a>

## plen property — IPv4 / 111131002002 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1122111330311200-1300301033003103-3100121121133323-2322201210002001-2231102332230001-0013322331332201-2323131101322311-3332213312222223"></a>

<a id="canonical-1013223323011223-0033131112022320-3323131132302321-3233010332220232-1312223212110112-2233230011322312-3203223113321121-0200202012200110"></a>

## prefix property — IPv4 / 111131002002 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-1310102031223230-3012101300211000-1030331220222302-2201111301001001-3230232230010310-0132000300222212-1012021302220120-3232320021022332"></a>

## Next pages — IPv4 / 111131002002 / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2021131022311213-1123101110132000-1122210101013221-1313333311323200-1001322003031000-2000000333030131-2023321003113112-1220030211210010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133210032112233-1331131203013102-0230230232133113-3112302330110103-2111310020121021-1102003230302003-3303212123123132-3233222131031211"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 220211220230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2301220030110210-1321112001300320-0333200310332010-1301212132322033-3220330211323302-1020202122231033-0130023223121011-0331112320011013"></a>

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

<a id="canonical-0123333321223022-2013321012330122-0030310131111000-1332302213101310-1223201010013310-0220333223123110-2302120130030113-3323023130200033"></a>

## Direct properties — IPv6 / 220211220230 / 3

<a id="canonical-2130123302202011-0010311001211002-0200100003203032-1033031133120302-0221122221310123-2320312312301303-1003132303210211-0311131002222302"></a>

<a id="canonical-1202013200130211-3121130320311220-1302200110300101-3330320231120031-3010101223121312-2233133103310132-1230132003133320-2300023311331303"></a>

## plen property — IPv6 / 220211220230 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2212133111010133-0313303333000123-1231232310021310-0332102033002133-1000021022003013-1203230102320023-1130111210320001-1111113332032231"></a>

<a id="canonical-3233322300120210-2300230321233023-3310222312232311-3121203123030123-0201233210130211-0212120300321332-3010332321113333-0213232000223332"></a>

## prefix property — IPv6 / 220211220230 / 5

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

<a id="canonical-1301221311303032-2002131303031120-3211220122321331-2320300020303202-3131121032011212-1013001221012231-3010220221303202-1131232101212220"></a>

## Next pages — IPv6 / 220211220230 / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2001202302330310-0100103301103103-0200111330302202-3211031001200231-2122300231030322-1103313310222221-2313011322130121-2231220032230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232012302022211-0331330032310331-1300100002333031-2222012232102003-0010231130001202-0332231122302211-0121312230101303-0200232330211330"></a>

## ingress_egress_gw_ar.no_dc_cluster_group — no_dc_cluster_group / 112303233230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_dc_cluster_group

<a id="canonical-0320221222213233-1320211310330023-3012210232031332-3300332311102003-3021001221223101-1222111020222203-1210321210220120-1003120333013313"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-2233022211203232-2031301002121300-1010030122232102-2021033300032203-2331200121030030-1121223010003300-0331323131313221-1113110113132320"></a>

## Direct properties — no_dc_cluster_group / 112303233230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231231333210310-1122311322313311-3020032310031103-3103121200312310-1212022022231020-3113001030010133-1223121012110102-2312100200323111"></a>

## Next pages — no_dc_cluster_group / 112303233230 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2123231301113210-0122001001313303-2112121031000301-2123012332320130-1332020030012221-2020000031000210-2112333312133330-0113023022123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031233212221103-3123332100310101-3201320021332310-1031321132300200-1221331222122322-3231022021030212-1223213233213222-0200313113122010"></a>

## ingress_egress_gw_ar.no_forward_proxy — no_forward_proxy / 203310032123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_forward_proxy

<a id="canonical-2301013333103213-0233301220123320-1221333230030210-0222010303113113-2231330203223032-1101320313222102-2323103121123200-1001310120210233"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_forward_proxy = {}
```

<a id="canonical-0313021103021000-0231001203111223-1131133113203312-1232131200123212-0010000331202330-0201032202123333-3303232300032211-2303110123221032"></a>

## Direct properties — no_forward_proxy / 203310032123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023100002002330-3211031231232201-2103033302213001-1101212223120023-2113213211331032-0020332130310013-1002123223320323-3233102132312331"></a>

## Next pages — no_forward_proxy / 203310032123 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3311012300023210-2010002130332111-0320210012231201-3123320103300322-2103110032133331-1221301023021132-0102030002121113-3322231013020111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231121011313132-0223122030012211-3323123220123113-0133020000133211-3301123332112101-3122013113111232-3221210013021120-3003031033010133"></a>

## ingress_egress_gw_ar.no_global_network — no_global_network / 313233102330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_global_network

<a id="canonical-0033233103330102-0333110102122232-1002221221102121-0201030022103311-0203300003120300-1330031010323200-3322301221023333-0012212031103110"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_global_network = {}
```

<a id="canonical-2213330303010131-3200322301123020-0113212233003200-2203301232222021-0011220001000020-0033323103130311-3300111110313320-0032332222011101"></a>

## Direct properties — no_global_network / 313233102330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023320332001200-0022031231020030-3123001302033033-1011111210131010-3121202221010322-3311221100200021-2033113322301103-2220003313210221"></a>

## Next pages — no_global_network / 313233102330 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0002310301230132-1203201302113031-2021313310120113-2011122130300330-2302020233202012-1012332100332233-0100313311012011-3202013230312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231120130023121-3312231112033200-0112130001201233-2001230121120033-3321321012211121-0120320201202210-3203121201303022-0310221223232030"></a>

## ingress_egress_gw_ar.no_inside_static_routes — no_inside_static_routes / 311031110013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_inside_static_routes

<a id="canonical-1021031032122120-2310330110113002-0200330100321120-3322133310302232-0101213213000201-3110331022231130-1021301012021011-0030331010201030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-0320100102010213-2103002220202003-2001200230331211-2012113112222222-3321331010310302-0001313020101200-2221323332311203-1320101210201031"></a>

## Direct properties — no_inside_static_routes / 311031110013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211021332323211-3131331322322012-0330132302322133-3202221230210333-2332133012011122-0001110133113211-3202003323031122-2232023223031030"></a>

## Next pages — no_inside_static_routes / 311031110013 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2012222100111221-0022030020231112-1213122020130101-2230123113020302-2220213303201223-0230232021120010-2332212210121320-3331311123103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031310201200301-2033002020003113-0133222011330300-3231100322301211-2200033031232322-1211001031103313-0211302032321223-1131010030313001"></a>

## ingress_egress_gw_ar.no_network_policy — no_network_policy / 133200122232 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_network_policy

<a id="canonical-2331103333030313-3113301001233133-0303001031321311-2002023231321000-1211113311313012-2021031002312101-2003211330032231-1311201311201032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_network_policy = {}
```

<a id="canonical-0100221003222212-1222023321201030-0201210213102130-2202122231021233-3220101211022122-1001310212021313-2232222331003331-3233211310033223"></a>

## Direct properties — no_network_policy / 133200122232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333300210202231-3333033201101313-0223323220203211-0001203002032013-0231303020003323-1100000300301022-3223321220030132-2011313101202123"></a>

## Next pages — no_network_policy / 133200122232 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3231233332203021-1001200031020211-2213332032231112-2133203233300102-0222320203313231-0103011311201332-0231313213301300-1203311030110310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322031332233312-3110212023112013-0330000131321310-3122102311220223-1222121112112333-3122102100122221-2002201310133110-3300110223032103"></a>

## ingress_egress_gw_ar.no_outside_static_routes — no_outside_static_routes / 010001022300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.no_outside_static_routes

<a id="canonical-3033222113333313-2233013313312223-0210212331321310-2120332103233121-0203012120221203-0232030310033133-3212000100120101-1300222023030033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-1203202112222322-0232102103312320-1231332120002300-1232221300001131-2013232030122313-3330022011211111-2310123131210131-2332122231233311"></a>

## Direct properties — no_outside_static_routes / 010001022300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323023102323101-2310310123202313-1020203330301130-3131101330200332-3331230111002230-2123031331021021-3312112100202200-1113233302211103"></a>

## Next pages — no_outside_static_routes / 010001022300 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331303210030230-0002210100232103-0332220022301131-3333030211123003-3330001233101120-3001021110200110-1312030011230012-3321231330232000"></a>

## ingress_egress_gw_ar.node — node / 000021202022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.node

<a id="canonical-0232311010031230-1200030133101102-3030100210103233-2331201223303330-3230032132033121-0220012133323313-2132130123133231-2222332133333111"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating two interface Node in one AZ.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fault_domain",
    "node_number",
    "update_domain")}
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
node {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130111231002323-1003032121023211-1210130230133231-1333312233201300-2102012322001102-2122111102213312-1001023233002322-1322131212000022"></a>

## Direct properties — node / 000021202022 / 3

<a id="canonical-0123021330131030-3033020233212320-3110133211321111-1013130322103322-0232230231132103-0133302230121311-3103333231133000-1231322200000022"></a>

<a id="canonical-1221331301000232-0222020103213121-2110023103332222-2311221301001013-3131202012211313-2202103311023133-1131131322213233-3233213313213322"></a>

## fault_domain property — node / 000021202022 / 4

Type: `"number"`. Optional.

Namuber of fault domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121): complete subsection reference.

<a id="canonical-3020031213103213-3133311320112122-1220113222232230-1302000321003203-0222103033333021-1320102212230321-3213320211323201-0113313100132020"></a>

<a id="canonical-0010002203031102-2210222100223010-2111133002312133-0323312102012233-2022023312033103-2122322020130003-0321122121130323-2113232012130132"></a>

## node_number property — node / 000021202022 / 5

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032): complete subsection reference.

<a id="canonical-1022313023131331-2022230023100312-1223110010013122-0221212011033013-2131120103013023-0012122322210011-0022332220212210-0120303030100111"></a>

<a id="canonical-2022312131330222-0111032122100133-3100330201000321-3130323002223320-1221300231210131-3033211133320020-1001010303231123-2000200230011010"></a>

## update_domain property — node / 000021202022 / 6

Type: `"number"`. Optional.

Namuber of update domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-1203020101200210-0113122010021123-1023032321332220-2310230001032322-2322113010011012-3221333133123312-2132132032013222-1333333321311100"></a>

## Next pages — node / 000021202022 / 7

- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213130011120021-2310111031331103-3031033233331130-3310301310232131-0123112130003333-3200111202013101-1100212133003103-1101133321313122"></a>

## ingress_egress_gw_ar.node.inside_subnet — inside_subnet / 213001100310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- ingress_egress_gw_ar.node.inside_subnet

<a id="canonical-0001000000322012-0121221203232110-0301333130220130-2023300233032223-3302300213221223-0031302003233013-3300220301233020-3030233023231303"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100003233130030-1212111030023232-2331013220101313-2110112101012012-1311123332012012-2110010232233311-0332133220301120-2010132023111020"></a>

## Direct properties — inside_subnet / 213001100310 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2100112303311011-0330210033212312-1032233302013233-0020320030321102-1222201121332013-0122002203312333-1013133333333310-0111001332123331): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0033130311203002-0021232331303131-1122332111320121-2313003210133301-2013333333322201-3023222132110133-2033213010101330-1033002103332233): complete subsection reference.

<a id="canonical-1020232220323202-1302320113030303-1003000120021001-1000300033203011-3133120220023010-1210332212120111-2303222121120213-2113023203130123"></a>

## Next pages — inside_subnet / 213001100310 / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2100112303311011-0330210033212312-1032233302013233-0020320030321102-1222201121332013-0122002203312333-1013133333333310-0111001332123331)
- [ingress_egress_gw_ar.node.inside_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0033130311203002-0021232331303131-1122332111320121-2313003210133301-2013333333322201-3023222132110133-2033213010101330-1033002103332233)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2100112303311011-0330210033212312-1032233302013233-0020320030321102-1222201121332013-0122002203312333-1013133333333310-0111001332123331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103331220300133-1103213331311113-0112331033313003-3223231231320321-1332032220211013-1131003223331302-1300303330223002-0003013103130311"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet — subnet / 131002023101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- ingress_egress_gw_ar.node.inside_subnet.subnet

<a id="canonical-0030333122200333-0212023302022130-2000003323020333-3210311032323132-3030010222013222-3212320032003312-3223331311113220-1301110222222111"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032223110011302-2232330231323133-2022010110013002-1010031222322033-1202220130113301-0102310112112123-3313221113032003-3313131312100220"></a>

## Direct properties — subnet / 131002023101 / 3

<a id="canonical-2223101330103333-2301202303103011-2233003033331101-2220330210330331-2300120211122121-0213210111200333-0300223123033100-3222323002213013"></a>

<a id="canonical-0100011120331300-2312333211120111-1000111322220102-3001100330321231-1103130221003300-2032011313020311-3102310210300102-0332133012231312"></a>

## subnet_name property — subnet / 131002023101 / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3310332322310011-0232331030111120-2232330200131211-1333033113111233-0032232003003111-3013000230022112-3030013000300221-2012230103310230"></a>

<a id="canonical-0221010202311023-0212332101212300-0022231201200323-0132211113003112-3232210331101210-0222330111300032-1030110022001132-2223012302332323"></a>

## subnet_resource_grp property — subnet / 131002023101 / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-2131013201233220-0132321211202211-2023012310120022-3023013323103220-3032133100300100-3102021112321233-3332030122332222-2201303331102132): complete subsection reference.

<a id="canonical-1323022021302102-3003003231302321-3010033123322310-2101003210333122-0313222303212120-0211300123231133-2102122323012013-1311323123211121"></a>

## Next pages — subnet / 131002023101 / 6

- [ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-2131013201233220-0132321211202211-2023012310120022-3023013323103220-3032133100300100-3102021112321233-3332030122332222-2201303331102132)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2131013201233220-0132321211202211-2023012310120022-3023013323103220-3032133100300100-3102021112321233-3332030122332222-2201303331102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231200210033030-3331211111320222-1013320210333322-0001000121321321-0120232000031130-0001031031313022-0030222330330230-1201112022233000"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group — vnet_resource_group / 101000021133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2100112303311011-0330210033212312-1032233302013233-0020320030321102-1222201121332013-0122002203312333-1013133333333310-0111001332123331)
- ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group

<a id="canonical-2210031100312002-0002123111112231-3122110011021111-0103033110113201-1023003320302011-1330323223002301-2100031023002333-2203023233022320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-3012301333233030-1121102321020320-1103320123012130-2320022333102030-0133333003300120-2002022003023021-1330320021320123-1103202003131032"></a>

## Direct properties — vnet_resource_group / 101000021133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130122123203123-3212320121101121-2231331220312112-3033010300000020-2123123302212022-2220020022133231-2203221030312330-3221120332133213"></a>

## Next pages — vnet_resource_group / 101000021133 / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2100112303311011-0330210033212312-1032233302013233-0020320030321102-1222201121332013-0122002203312333-1013133333333310-0111001332123331)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0033130311203002-0021232331303131-1122332111320121-2313003210133301-2013333333322201-3023222132110133-2033213010101330-1033002103332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002203030231111-1001001320221230-2332113023023300-3022213122103210-0020011302131320-3311033110221102-1310330203322133-1203333311030031"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet_param — subnet_param / 123213133222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- ingress_egress_gw_ar.node.inside_subnet.subnet_param

<a id="canonical-3200101120203122-2333133033230132-0320103231220231-2303202300020012-1303002233232231-0113123111123230-2100101001033033-3113101311132120"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333030331310100-0302312102333323-0030111020222013-2131110331013302-0233210320203103-0310202310130103-2022000220112211-2231113012321003"></a>

## Direct properties — subnet_param / 123213133222 / 3

<a id="canonical-3231313330222321-2200021233303133-3111031311011121-1003313132201200-2112232300010213-0220003111120220-0101031231322200-2201111212123322"></a>

<a id="canonical-0132101130130121-0031201221130231-3030033201210001-0321312030001312-0310222032323202-2230003301303211-3301130323131211-0121030331300131"></a>

## IPv4 property — subnet_param / 123213133222 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3302131213300331-0210321031110013-0303233311301131-3210302001311223-0211333022320213-0002331033312233-2003120332002320-2202211120221103"></a>

## Next pages — subnet_param / 123213133222 / 5

- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3123033001331333-0110023212132131-1211100322223012-1223013001233103-0310122321200010-1320000331033012-1033320103220030-1112123121221121)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231320313020313-1133023111101121-0211201003031030-2233322110322300-3303300003210002-1131131300030132-2020033121330213-0122020233030100"></a>

## ingress_egress_gw_ar.node.outside_subnet — outside_subnet / 221133021302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- ingress_egress_gw_ar.node.outside_subnet

<a id="canonical-1330033212111032-2033110131132013-3223123222213212-0231210200220223-1223131011310301-3120030202232322-3121200111133031-1213300132121030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100223301021131-2112031222012023-1112032131133023-2301132020220013-0330123300223233-2222132101002203-0100333113122331-1112022111232121"></a>

## Direct properties — outside_subnet / 221133021302 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3223321000223231-0212220023133021-1310330331331101-0003012212202030-2002102313130000-0333310232211303-0021023332122331-1123232321013310): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0003330123323110-2322100002213313-1030032333303112-0102301003121010-3000102211123201-0311321321132112-1333000310200132-1203012000333223): complete subsection reference.

<a id="canonical-0230310201232113-3230310120100113-1011331000012303-0012200021211101-1000120332000012-1323302311023012-2330102222133231-2312231131321110"></a>

## Next pages — outside_subnet / 221133021302 / 4

- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3223321000223231-0212220023133021-1310330331331101-0003012212202030-2002102313130000-0333310232211303-0021023332122331-1123232321013310)
- [ingress_egress_gw_ar.node.outside_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0003330123323110-2322100002213313-1030032333303112-0102301003121010-3000102211123201-0311321321132112-1333000310200132-1203012000333223)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3223321000223231-0212220023133021-1310330331331101-0003012212202030-2002102313130000-0333310232211303-0021023332122331-1123232321013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321213300233310-0222033033012112-1313013113320320-0301111032222133-2120313210211100-0331302302221012-1203322210012201-0021100010333100"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet — subnet / 000131103222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- ingress_egress_gw_ar.node.outside_subnet.subnet

<a id="canonical-2323033033200011-0213322330013233-0131230302333333-3323112220312022-3003213223233022-0221112332023120-2000132112303211-0123133122101103"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212302213121102-0111232032032333-0222321121233131-0023223121311202-1223200223311003-1010231011012220-0231032113003020-1302230331211311"></a>

## Direct properties — subnet / 000131103222 / 3

<a id="canonical-2011100122023032-2312122013220320-3110312313130322-3030033030322211-0323133321322201-1220320302310023-1100210311032321-2130310031003131"></a>

<a id="canonical-1101303210232101-3322201323201232-2023311321110121-0102031020130000-3012033301003130-1211203030221203-3031033013310200-1200212311233232"></a>

## subnet_name property — subnet / 000131103222 / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3223330331330012-2302320133112201-3003112122310312-1021131301212210-0122020130202300-1203022123313100-0211312330033001-2311012320333000"></a>

<a id="canonical-1132311231022331-3210223331322310-0332133132300011-1332030022110331-2001012031310331-3211220103200330-2200131322012232-2133213030202012"></a>

## subnet_resource_grp property — subnet / 000131103222 / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-0311013012011322-2012112210131010-3010032301223131-0231312012100302-3100020102312332-0313013101032312-0030322302220321-3311133233130111): complete subsection reference.

<a id="canonical-1233020133322113-3311332131001332-3233320300333002-2111321332321330-0222131022133200-0300221313010011-3132112102120033-1122302132301231"></a>

## Next pages — subnet / 000131103222 / 6

- [ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-0311013012011322-2012112210131010-3010032301223131-0231312012100302-3100020102312332-0313013101032312-0030322302220321-3311133233130111)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0311013012011322-2012112210131010-3010032301223131-0231312012100302-3100020102312332-0313013101032312-0030322302220321-3311133233130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122002132130021-2033132321212022-3010123201002223-3323220301323210-3220321003111233-3100313130332013-2232321312000311-3102231132201112"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group — vnet_resource_group / 013211323032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3223321000223231-0212220023133021-1310330331331101-0003012212202030-2002102313130000-0333310232211303-0021023332122331-1123232321013310)
- ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group

<a id="canonical-2211203023130130-3311330222122132-3103130020012220-1001322222130120-2313032110220201-1023210303120233-0120121010021113-0112322322111100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-1120321323112100-1012230203230130-1033101203010231-1120213122331323-1332303322200312-1311031000233121-2032013320003130-1103231121320221"></a>

## Direct properties — vnet_resource_group / 013211323032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330331133013311-2311211102331101-1201002012211200-3222122300130111-3333023112020322-1312000120013313-1131213032030122-1313200103132031"></a>

## Next pages — vnet_resource_group / 013211323032 / 4

- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3223321000223231-0212220023133021-1310330331331101-0003012212202030-2002102313130000-0333310232211303-0021023332122331-1123232321013310)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0003330123323110-2322100002213313-1030032333303112-0102301003121010-3000102211123201-0311321321132112-1333000310200132-1203012000333223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021233100010332-1010012222330101-0211202320001130-0013330231120003-1112003230131213-0123033033303132-1001001232321311-2203110233111211"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet_param — subnet_param / 310301032302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-0310131323310002-1113100332312212-0131001031010203-0202323321300311-3211003202030030-2113230231030203-0202121113112301-1121202003033332)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- ingress_egress_gw_ar.node.outside_subnet.subnet_param

<a id="canonical-2122001131103031-3310003301301232-2233020210320310-2012103333002333-1031320031003212-1111020300202133-3332110210103310-0113232312120223"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031023101123022-2311212201112002-3221220300211110-2322203110002012-2230023010201032-3113200002313123-2023120132301331-3310212001121021"></a>

## Direct properties — subnet_param / 310301032302 / 3

<a id="canonical-3213210112331201-0223033101112323-1032202311310230-2020321001211102-3113003330303020-2313301333223221-2000203122302122-1103232121211320"></a>

<a id="canonical-2213300230101221-0313202110331133-3131231100201011-1301103233223131-1113122133003133-3123311321033023-0222032331133111-2030132210332112"></a>

## IPv4 property — subnet_param / 310301032302 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-1110032220302233-3222213232232010-1012021313033133-1322010130002020-1102222331013220-1023203322211031-2112031123300122-0010303232220010"></a>

## Next pages — subnet_param / 310301032302 / 5

- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2033013201111023-2101333312321223-0120203132330323-2121123212123222-2012221233100212-0322023313322312-2232011101000110-1000302032131032)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0332132312230000-3021223331310100-3131113231132221-2110000203012020-2303122033223133-0322321103113032-2233122110201001-1101220011203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023030132302110-1122001113003300-2320120133110102-3101223301223203-2133302023303330-3312002231310132-1300001001013233-2232212030023300"></a>

## ingress_egress_gw_ar.not_hub — not_hub / 032200133220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.not_hub

<a id="canonical-1101010000201021-2012301112021222-2300023032233120-2103302013022122-3020331123121332-1133320101003120-3332123310130202-2011100002133331"></a>

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
not_hub = {}
```

<a id="canonical-2322200131123200-0002012303112120-2102323223020220-1202313110122130-2321120311202122-2223111231311120-0202222100010001-3313212122101113"></a>

## Direct properties — not_hub / 032200133220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300130330030202-0322301320213330-2131113301212103-1000313010113310-1111331303102211-0221012030031100-0311203331111013-2220020030012121"></a>

## Next pages — not_hub / 032200133220 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133000301220212-3132212133031310-3020303011102123-0300012311221113-3021113101220003-1310012303231112-2131110023111213-2222200033220333"></a>

## ingress_egress_gw_ar.outside_static_routes — outside_static_routes / 101121223330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.outside_static_routes

<a id="canonical-0013021033312121-2000221111020333-2301213002200200-1300032000313020-3010030000212321-0312332133133321-0131103310002302-2310333303023211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231120213303210-2301100320001022-3213021331222313-3111002113312021-1313111231321330-3311131333010233-2030113130330301-0231112311301221"></a>

## Direct properties — outside_static_routes / 101121223330 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120): complete subsection reference.

<a id="canonical-0112031313023020-1100012201022323-1123130320311100-3313200303231111-2120330112332101-2003131111123103-1120223221021303-1302033130331300"></a>

## Next pages — outside_static_routes / 101121223330 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013103312131320-1230212113213203-2010132011222011-3133221223220211-0212020130030300-0132123220032000-2312202003300133-0010322021231111"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list — static_route_list / 300022012031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- ingress_egress_gw_ar.outside_static_routes.static_route_list

<a id="canonical-1010112002220312-0313133032130033-2122012202012130-0101331120322003-2221320201300302-3231331120020333-0123001313022203-2102102113132130"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112013010222123-0223011003201001-0011210201321313-1101031313212131-1102221120300032-1230231020303330-1311331300132332-0332331203221202"></a>

## Direct properties — static_route_list / 300022012031 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332): complete subsection reference.

<a id="canonical-3100200213222012-1320231322312021-0100302032223113-3001200212113110-1210010202333331-2230211321100220-1202103000220131-0302312123301333"></a>

<a id="canonical-2130313020311121-3233222332322023-0311120021100201-3313022303220202-3112012130233001-2012131013323201-3300300122231103-2332333323020300"></a>

## simple_static_route property — static_route_list / 300022012031 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3302122021230020-3230331111110132-1001103221113202-1322033122322000-0310321010123122-2022212011030211-3030101233133011-0001100020020303"></a>

## Next pages — static_route_list / 300022012031 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131221120021333-2321213030300032-1120001013300232-2321220102031220-0112003322103100-2330122301330013-2112212330102001-0020133332201222"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 013200202103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-1110201221202212-2200233133220122-3231211200311313-2103100233320000-2301210001102231-0320020131000333-2112212211111200-1322310010113332"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111302203212133-0113200110111231-0212022112100131-3010022203102233-1221032123032300-3000111202322202-1321311320332012-1220010232123020"></a>

## Direct properties — custom_static_route / 013200202103 / 3

<a id="canonical-0300011331122031-0212312022123322-1312332132322033-0331023022321110-3133211003301311-0201333312313213-1313110120020332-1231303102130023"></a>

<a id="canonical-2300100313023000-1320223011003223-3310101210211000-2132112200121121-0200313102030110-3123122113313310-0202322321111020-1031203102333203"></a>

## attrs property — custom_static_route / 013200202103 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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

- [labels](resources--azure_vnet_site--reference--group-007.md#canonical-2321220201001331-0311200312312111-1200130312232231-1112032210221121-3021133133221111-1212201312130102-3022132223221013-3020301220023211): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023): complete subsection reference.

<a id="canonical-1012122020220011-1320033303113220-0231102112130211-0221200301203120-0221213001002313-0211121101121221-3122020123010132-0003330130220231"></a>

## Next pages — custom_static_route / 013200202103 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-007.md#canonical-2321220201001331-0311200312312111-1200130312232231-1112032210221121-3021133133221111-1212201312130102-3022132223221013-3020301220023211)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2321220201001331-0311200312312111-1200130312232231-1112032210221121-3021133133221111-1212201312130102-3022132223221013-3020301220023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103313223220333-0103311311121220-1302120003110202-2223121221333331-1030323003122000-1320001313312000-0113322212333113-1202020031233303"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels — labels / 231023303102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1013300212301130-3102333113110302-0300202303300300-0100300102221020-3000210320320003-0310033032331132-2120131010301120-2200132120222123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-1102310210110031-2230003332323000-1021123212110131-1212021023012202-3000303231322133-0223221323321031-0031323302021000-2311323100001023"></a>

## Direct properties — labels / 231023303102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311023103301321-1033031133203300-2221212202321023-1012230011230203-1330323320101312-3133202031233313-0122230332121323-3110333231313301"></a>

## Next pages — labels / 231023303102 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021300202322222-1102320110111220-0132131331112133-3323302333030031-1022231132103100-1122333220203221-3232223300230233-1323031122313301"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 213222213213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3101202332010013-1221113003200320-3123333212100223-0231300203123102-1002120320000231-1002231301000021-0023312222303213-0203203123011112"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321020220022310-3333230232120000-2020121201213102-0122212201022213-0220332313320222-3230103311332222-1002120013013021-2013013111222112"></a>

## Direct properties — nexthop / 213222213213 / 3

- [interface](resources--azure_vnet_site--reference--group-007.md#canonical-3133212331033212-1223331131112102-0201320211320010-0101130300310030-1132321003001132-1010022121022311-3112310233230312-1001133103022211): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232): complete subsection reference.

<a id="canonical-3232331233101121-1230112011302211-2130323113230023-1202010332113112-0003111013100320-2021303312330132-0203011320031312-3232232203100030"></a>

<a id="canonical-0120322300131213-2303031323201013-0202312223032123-1232220212012022-3103102321130220-1303211210003213-2212002222200022-2332302320212021"></a>

## type property — nexthop / 213222213213 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

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

<a id="canonical-2032333203030011-0302000331100202-0003002020213103-2301221102100133-2003022323003030-1311103012230300-1330203123312012-1031301033301333"></a>

## Next pages — nexthop / 213222213213 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-007.md#canonical-3133212331033212-1223331131112102-0201320211320010-0101130300310030-1132321003001132-1010022121022311-3112310233230312-1001133103022211)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3133212331033212-1223331131112102-0201320211320010-0101130300310030-1132321003001132-1010022121022311-3112310233230312-1001133103022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003122023020312-2121202012303012-3313100131112130-1330030012000101-2102210213023033-0211021110100310-1230312301021122-1031232313011301"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 023313031312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-2301303010000100-1031101313313221-1323122030323232-0021233200131233-3123333321231121-2132312312032110-2011303030220031-3112101002231333"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233320130333332-0102331111012233-2032111222212210-1020310131012332-1320210112033313-0002222023211021-2223202321030322-1230213223231030"></a>

## Direct properties — interface / 023313031312 / 3

<a id="canonical-0332131311020032-0300202120012312-1020010233213013-1311020310333322-1302121300312023-2201200110301020-1300232111301223-2103320000120121"></a>

<a id="canonical-2012331111123201-0030320331020302-3211112310230031-0302010310032303-1222313022122321-2132322001011300-1013033211201321-3122200202132312"></a>

## kind property — interface / 023313031312 / 4

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

<a id="canonical-0320211212210032-3213020331230212-2000303331203310-2020123231212220-3103222112301223-1000312201221220-3200213102201230-0001003310133003"></a>

<a id="canonical-3100002232200123-1220233031333032-0101120333201332-0011333013103123-3323112321320011-2321323202100003-2332232031030103-0122330021033031"></a>

## name property — interface / 023313031312 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0232203033113303-3000102201130313-0331201000300232-3320110001022123-2130033200003101-2212130232200010-2132010033033010-1001202032003301"></a>

<a id="canonical-3233210200200202-0233002202033230-1232203311003221-3303030322333321-0020323302030033-0031223231210331-2101000100320130-3121120222110303"></a>

## namespace property — interface / 023313031312 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-1033011132023312-0030202321300020-1021202302120230-0201013302213322-3202133331202131-3031310112123110-0030332323313111-2033122203322230"></a>

<a id="canonical-0212003212033201-2123021203222000-0221212330311212-3102031000312200-0123212223220220-0311011300012220-3021031031210311-2011210232033123"></a>

## tenant property — interface / 023313031312 / 7

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

<a id="canonical-1221001023203010-0012121212311332-1312332100110102-2133331301200121-3231233122330022-2103023302120131-3321011022130112-0201120331321232"></a>

<a id="canonical-2032331330311223-3112311331013200-2113223113320031-1211010312002023-2232102303121101-2112103211331330-0130232133010123-3012222210012100"></a>

## uid property — interface / 023313031312 / 8

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

<a id="canonical-1311202202313200-1123313313013320-0222222001223310-3021003330321021-2222203312003012-0333321200231120-2210111231111112-0332020001231023"></a>

## Next pages — interface / 023313031312 / 9

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311310311233330-2322223011303011-1002200203321321-0033220220231323-0102032001212320-3113232011101000-2321333222331111-0330201003330110"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 110013230230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-3201323220203300-2211201332113212-0121303311203113-3130121112022111-1131031002323022-3031331312220321-0321320203321121-2132010310121100"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2202220120203300-0200213231102201-0000221100320133-0013003202030103-2112021111320230-0103223110230023-1112031102200330-3233103111332033"></a>

## Direct properties — nexthop_address / 110013230230 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211): complete subsection reference.

- [IPv4](resources--azure_vnet_site--reference--group-007.md#canonical-1032322231023133-1003220023003131-3101323202130212-3122211122222210-3012121203010002-1311221022102131-0320212223122211-2302302311111113): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-007.md#canonical-3013221202222130-2123210310322211-3310020100210120-0300112031332300-3232122233222203-1300312013321223-2231333212230310-2321002301133203): complete subsection reference.

<a id="canonical-0111201333230010-0103222112223021-0111331212221000-2223002221032222-3032013022322333-0032331012012200-3012212100130213-0002012110111033"></a>

## Next pages — nexthop_address / 110013230230 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-1032322231023133-1003220023003131-3101323202130212-3122211122222210-3012121203010002-1311221022102131-0320212223122211-2302302311111113)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-3013221202222130-2123210310322211-3310020100210120-0300112031332300-3232122233222203-1300312013321223-2231333212230310-2321002301133203)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032133201312202-3020032003233310-0303010003330320-3233101030222112-2023100322300021-3133210133311203-2301313332131033-0233003030213202"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 020012132213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-1030200130103313-3111033300020302-2213002320331122-2223300033203320-0321103332302021-2123331222302102-0322002220132230-2031301013133001"></a>

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

<a id="canonical-3023230122211230-3311120113313301-2321313132123320-3132121200121302-0212203231210103-2331011330102001-3103313331303032-1202031102332132"></a>

## Direct properties — dual_stack / 020012132213 / 3

- [IPv4](resources--azure_vnet_site--reference--group-007.md#canonical-0213013011131330-1001303231013130-3231001330310212-3202100311202212-3123100321100201-3222230003301203-1113203101011013-2301103302111100): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-007.md#canonical-3220303133011202-2313033133310222-3023232300302020-0131333320221211-1332311123111302-1132012103212311-1323233013031003-0103221021220022): complete subsection reference.

<a id="canonical-3121121102211201-3313213301123232-0232211013123003-2032100000001023-0322201202233113-1301311133020333-3033200112130310-0101310221133010"></a>

## Next pages — dual_stack / 020012132213 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-0213013011131330-1001303231013130-3231001330310212-3202100311202212-3123100321100201-3222230003301203-1113203101011013-2301103302111100)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-3220303133011202-2313033133310222-3023232300302020-0131333320221211-1332311123111302-1132012103212311-1323233013031003-0103221021220022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0213013011131330-1001303231013130-3231001330310212-3202100311202212-3123100321100201-3222230003301203-1113203101011013-2301103302111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311013000132103-1312210323232331-1130311330221030-1220313200133202-2130033212331213-2302200333130130-1021122132223202-3122210202330333"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 120221002032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-2030102213220021-3010112202032021-1213200123013330-1232132212032300-1000033320331223-3233231321011112-0021232033302221-3301011102230301"></a>

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

<a id="canonical-3132200112333302-2233231120020122-2122031120033113-3220130302323223-1220321103120221-0113021020002313-2103101232032323-3021001223132301"></a>

## Direct properties — IPv4 / 120221002032 / 3

<a id="canonical-0312323130021332-0331000202310123-3012331320000120-2302302003312032-1211301001032232-2302211220001321-1211301200302221-1013330010322233"></a>

<a id="canonical-1121133131220313-3212131203321011-3011331100112213-2000021202331220-3330003320031020-3312003032121033-2331022202303333-0220221303320202"></a>

## addr property — IPv4 / 120221002032 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-1213210021011012-2003301313112112-3020231231221020-1232102122301013-2312022312321230-1133232211111322-3213221310330033-1333112021200130"></a>

## Next pages — IPv4 / 120221002032 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3220303133011202-2313033133310222-3023232300302020-0131333320221211-1332311123111302-1132012103212311-1323233013031003-0103221021220022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001303231111020-2111112131012110-0123130120300030-3011322123323102-2312100212231323-3100302100030132-0333113313020213-0130213022311311"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 300123032311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3032132333223102-1021222333002331-1102311320020321-1232221013023331-3300220200130210-3003302002022001-1211013220322232-2032321101100221"></a>

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

<a id="canonical-3222022003123112-2100303010221121-0331312301201112-0210103220003300-3002310100012323-0020131012303212-2100302101210121-3000300123200000"></a>

## Direct properties — IPv6 / 300123032311 / 3

<a id="canonical-3020231220010111-3202013213113230-1233003020202003-1213020232200013-1000300233201211-0320002203300111-2002213333012022-1023131232103120"></a>

<a id="canonical-0120001111121221-1011333222312222-2032222101231012-0212122203100102-0313321231230303-1213202003310131-2310002233233101-2313120123000120"></a>

## addr property — IPv6 / 300123032311 / 4

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

<a id="canonical-1102223310103100-3323200331033032-0203013020211311-0330210231110033-2221221221110200-2122203103131221-2310213311023100-0313130221222122"></a>

## Next pages — IPv6 / 300123032311 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0012202030333102-2013212132121122-1211111010013220-0330303201230022-2202200121130112-3111120331202031-3322130202230011-1210103301223211)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1032322231023133-1003220023003131-3101323202130212-3122211122222210-3012121203010002-1311221022102131-0320212223122211-2302302311111113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001130300013212-0303122012331133-2022232210010123-1003303012002132-3003123013033223-2321002222133331-1033031203303100-3303321211130022"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 011003022131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1323132331231201-3113231333301031-0321012010231001-2033111111003233-3113321203323201-1102300123223212-2333213330002103-2122230021011211"></a>

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

<a id="canonical-2221220001310230-0123210120000113-1232001102223213-3122022120332033-2310300112110002-0202322303000110-3231221313011001-1233202300300210"></a>

## Direct properties — IPv4 / 011003022131 / 3

<a id="canonical-1231032001232301-3322331303022321-0021021003132310-3331012113102130-2210201333311223-0021330201210213-1132111232311331-3110122232020211"></a>

<a id="canonical-1313012201122201-3213330213210333-2120002313211120-3100133232221302-1113001202122113-0133113322223203-1112233202310033-3103331001033112"></a>

## addr property — IPv4 / 011003022131 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3231302211223123-2211110122211233-0302133103011120-2303103232310102-2302320222311201-1002320300201303-1122223101321033-3010012313112102"></a>

## Next pages — IPv4 / 011003022131 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3013221202222130-2123210310322211-3310020100210120-0300112031332300-3232122233222203-1300312013321223-2231333212230310-2321002301133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301333202101021-0123303213203220-2213123131310102-2113022010100322-2130301123203121-0102100332201320-3200313120332213-2130010312122023"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 220321103313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-2311203003301300-2113320323312310-1230311232200303-0302232220122303-2200220312012220-2220121121230131-1132132100123330-0002031122211303)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3112233133120030-3032000322302101-2230123133130001-3111001311332223-2332003301012111-0021112310132012-3132102120030303-1211231013111323"></a>

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

<a id="canonical-3023022332230221-0323201020322101-2132112022011311-0122301112020120-3300220331011011-3310030020010123-2031201022220101-2211221331001003"></a>

## Direct properties — IPv6 / 220321103313 / 3

<a id="canonical-1212111123322201-1111302313100332-3302120030012101-2033021103231000-1110330302031310-2111323321221312-3031023123332020-1332001113303023"></a>

<a id="canonical-2201210332111121-0211312323230321-0020000123123232-0130031313202010-0311203220333121-0003301103021223-3301302332011020-0321030202300030"></a>

## addr property — IPv6 / 220321103313 / 4

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

<a id="canonical-0231112120023001-0230112222001121-0032332310010032-0002032013222102-0311003130130212-2233200200202310-3032011312332102-0003030301101010"></a>

## Next pages — IPv6 / 220321103313 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-0010033000031212-1013103032330131-3013132212113101-2231113303012211-1311332121032121-1130010200100111-2123320321230012-2200302213331232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032133110203231-2231123223022003-3002002003102113-1301020303300100-2322131011220220-3103100213032120-0331323313130210-1001012201032103"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 331123033233 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-1131331111230012-2311220212200213-1301031021111322-3020031233110220-0310301113130022-0103323333310301-3020321220110213-3111221120123013"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
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

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301012233131222-1211322130102321-3222231303323032-3323202233003111-1233210103100022-2000301012010332-2323121323132010-2100022302300011"></a>

## Direct properties — subnets / 331123033233 / 3

- [IPv4](resources--azure_vnet_site--reference--group-007.md#canonical-0220032111330030-3131033312322303-1310022011201331-1313233102320022-0102231210113203-0120311010000303-2230030320113110-0330302123101233): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-007.md#canonical-1010312212233130-3101132331301121-3303102201231020-3030112323131123-2231031102101201-1300312303301010-1311202212020213-3001301213310300): complete subsection reference.

<a id="canonical-0131303230132331-3003111113102233-1210303233103330-0333312133301013-1130211330211231-2300012011100033-3201000323013022-2210201212102211"></a>

## Next pages — subnets / 331123033233 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-0220032111330030-3131033312322303-1310022011201331-1313233102320022-0102231210113203-0120311010000303-2230030320113110-0330302123101233)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-1010312212233130-3101132331301121-3303102201231020-3030112323131123-2231031102101201-1300312303301010-1311202212020213-3001301213310300)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0220032111330030-3131033312322303-1310022011201331-1313233102320022-0102231210113203-0120311010000303-2230030320113110-0330302123101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332012000112132-2213301100130333-1200332300212122-1333111023013313-0000132310023311-2022323021212002-1232011112320010-1013223023100222"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 201112232001 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0030011322030213-1120202122111023-3100211210223201-2221200130232320-2201121200230131-2233022220223013-1312131221132213-1220032022221200"></a>

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

<a id="canonical-1132033331231012-3023103130012011-0313332001022112-2103220011113012-3311202111103021-1322030313030312-3303211200221122-3210331201002002"></a>

## Direct properties — IPv4 / 201112232001 / 3

<a id="canonical-3032220132010231-3001121022112300-3323232103031223-1133021130022203-2233113210232203-2202121200103132-3310302210113123-3223133220133231"></a>

<a id="canonical-0302023103233220-2110120030103000-0222231331013201-2121003231200321-0201232231233301-3120310031110321-3133332021222022-1310320221112111"></a>

## plen property — IPv4 / 201112232001 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1102303200110202-3201312311333020-1213033232112120-2112331000232133-0223123331111103-1301102003233002-0031100200113123-3000022003100130"></a>

<a id="canonical-2001132323121031-2231101001201302-3313222332203223-2123232030331132-1021213321321201-3210232100311230-3111112113210223-2313001000133303"></a>

## prefix property — IPv4 / 201112232001 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-0012011122112002-3013012102310331-0210331010330111-0220232220030022-1332023230021032-3020113201200213-3001202200211231-3101030231121130"></a>

## Next pages — IPv4 / 201112232001 / 6

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1010312212233130-3101132331301121-3303102201231020-3030112323131123-2231031102101201-1300312303301010-1311202212020213-3001301213310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122313101202210-2032100322321222-0102122032231131-3032310121101311-3101113310033003-2331020201223113-3220022023312021-3312003201103231"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 023313032011 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-3110121201232121-3333131110120223-1331320010322113-3031012210210213-3213331210301222-0231011311210132-2210110200203130-0231323232012022)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-1102221113313032-2123220330303131-1203032002001320-3301110311302113-3211033322320231-1113003131031031-1301323301002323-3130320000301120)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-0200122022312030-3301022211203110-0121311112130013-0333331123223221-0300021023010221-0300002001030122-2321202130302020-1220012302223332)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1023223121012201-2232100123002123-2310203111021033-2322223213000102-1232313023231022-0021222023130332-3031011132102301-1322312301231021"></a>

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

<a id="canonical-0231101012312023-3022013211322321-3313323301121330-2310231100031111-0121131322301320-2023223022221211-3013200123033110-3121332200312103"></a>

## Direct properties — IPv6 / 023313032011 / 3

<a id="canonical-3231322211013030-1113321333233123-2201021201033010-2011200013223110-0101113203011111-2113111331022223-3213110032000122-0220313213213122"></a>

<a id="canonical-0111322003211312-3003203121010112-3230311011202122-1101122033003133-2233133101110021-2020223300331100-3322101133203331-2022312101311123"></a>

## plen property — IPv6 / 023313032011 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0303030113023211-2301101132200021-2002010333333311-2312320030122120-3311312130003133-1202232012301333-1322222022323212-2300322331330122"></a>

<a id="canonical-2133301310021220-2002232011021010-2201033133132222-0232321130112100-2203132212113323-2003002031212203-2322123220121133-1331023213003122"></a>

## prefix property — IPv6 / 023313032011 / 5

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

<a id="canonical-3303123010001301-2120010223023022-2002330232101303-3211330231333131-0333323102230023-1131211001211221-3132321121211120-2202213111202120"></a>

## Next pages — IPv6 / 023313032011 / 6

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-2323232230223031-2212323122231102-1332020202010002-3311131133310203-0203021312231200-3220032001210212-3113301101210203-0030331233203023)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112331002111102-0011111131203123-0012210323332222-1101102020231230-2123223232121121-2013310220020311-0110202213001120-3332223032021102"></a>

## ingress_egress_gw_ar.performance_enhancement_mode — performance_enhancement_mode / 232231310322 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.performance_enhancement_mode

<a id="canonical-0201321002310102-2120122331120333-2012322333200001-3122320310233123-3110323230211023-3210322223222233-2311121002220132-3010000323011233"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121001231332201-0131010202212001-1201132310322312-2313302010302311-0121110321312333-2232322000210321-3023031303131120-3221113133302033"></a>

## Direct properties — performance_enhancement_mode / 232231310322 / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312): complete subsection reference.

<a id="canonical-3201302122011103-1332123103322221-3230122310212103-0212023011103303-2022013202213222-2032231030331032-0102102002203200-1020231320032202"></a>

## Next pages — performance_enhancement_mode / 232231310322 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321123323231132-3123022221312332-3131303310122301-1310112022001231-3032003012003122-0100111031113310-0321303310233033-2222132221031320"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 112331230210 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3013122333302113-3112201112101332-2210301022021313-1312310211101001-2321221210221030-3300030211113320-2311303102310213-3112232322311221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321210110321322-3313112023121310-3113130111230120-2303123021121012-2301133001213001-3133112213031300-3210001331203203-1312332232201030"></a>

## Direct properties — perf_mode_l3_enhanced / 112331230210 / 3

- [jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-1330130203300001-3112031221213023-0222323133201203-3022112321130323-2030022012210212-3132021021023012-1012321011312323-0003233001102312): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-1320231102323231-3132323200330011-2311110110101300-2113121330020313-2212330100120123-2312220233211023-0102010102221003-3202323031211320): complete subsection reference.

<a id="canonical-0202013023320032-2032100112221101-3303111203302120-0001000000311312-0311132133323020-2322013333322202-2313223333312330-0010222313021012"></a>

## Next pages — perf_mode_l3_enhanced / 112331230210 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-1330130203300001-3112031221213023-0222323133201203-3022112321130323-2030022012210212-3132021021023012-1012321011312323-0003233001102312)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-1320231102323231-3132323200330011-2311110110101300-2113121330020313-2212330100120123-2312220233211023-0102010102221003-3202323031211320)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1330130203300001-3112031221213023-0222323133201203-3022112321130323-2030022012210212-3132021021023012-1012321011312323-0003233001102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012230122310111-2113031320331002-2300013323320032-0110210132130331-2212311003020303-0121132101010301-3112120033231233-1332100011320000"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 133222000203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-3122123332330021-2001201102201030-0222200212121231-1222120131112000-2011321031102102-2130031111011300-1233231110211001-0130000202102300"></a>

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
jumbo = {}
```

<a id="canonical-1302103323032220-0011021311310311-3110010222330210-3211203202101332-1120111001011022-1112232120332032-3132223322300111-3030221111211313"></a>

## Direct properties — jumbo / 133222000203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230021211220120-3101030033302123-2010121220023212-1132130231031131-0121222111300301-1233230310112331-0001301212212131-1201033031203201"></a>

## Next pages — jumbo / 133222000203 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1320231102323231-3132323200330011-2311110110101300-2113121330020313-2212330100120123-2312220233211023-0102010102221003-3202323031211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320321033103311-2101323012333020-0210103233011001-3010321110211300-3201313221302032-1033322301113111-1122232023220230-3003213300113022"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 300122103101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-3130323313302331-3213223331230301-0001113223133121-2332000300020311-0100110312001300-0111112301302131-3312221210312233-0130332223110320"></a>

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
no_jumbo = {}
```

<a id="canonical-1031202230123001-1331231301103313-3210303310111213-1303120302311103-0313010313332012-2021022031320321-1100311331031131-3223103210101131"></a>

## Direct properties — no_jumbo / 300122103101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313201320312002-0231011032001120-3220222020332102-1112010222003211-3031122011213323-0233211333321002-2031030100230313-3022210203212122"></a>

## Next pages — no_jumbo / 300122103101 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-2131003231322333-3200021102222030-1211330023203301-2233211313200201-3203232220102233-2210110232002122-0001002011330231-2301210012202311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003200001030201-3130021123210102-0112332000020330-1021111303333200-0322112121331233-2302031130121011-3030233230101202-0310001223203012"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 020312312110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0002100130331120-2133302032201333-0021232010011122-1231131330213013-0020310130121303-3320130312021113-0211012210320013-1213132021133003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000333213200012-2303013212030331-2123302011232210-2121022120331022-2220323322033001-1020202200213230-1222220222010030-2233020232210220"></a>

## Direct properties — perf_mode_l7_enhanced / 020312312110 / 3

- [jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-3211211303022023-1210223123010002-2210131113021132-2302133001202211-3300311302202221-2121103102003031-0110233110020011-1333031033121033): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-3221132022113323-3023200012030102-3223002021020322-3033130123203132-1220222103110310-2300223303101103-3210331110332001-0302332131202220): complete subsection reference.

<a id="canonical-1033323212030013-3131033212030333-0000102301012001-2112101003300232-2032212220332201-3031332020331003-0033012223111022-2101122313320312"></a>

## Next pages — perf_mode_l7_enhanced / 020312312110 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-3211211303022023-1210223123010002-2210131113021132-2302133001202211-3300311302202221-2121103102003031-0110233110020011-1333031033121033)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-3221132022113323-3023200012030102-3223002021020322-3033130123203132-1220222103110310-2300223303101103-3210331110332001-0302332131202220)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3211211303022023-1210223123010002-2210131113021132-2302133001202211-3300311302202221-2121103102003031-0110233110020011-1333031033121033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333300230100003-0202000310332012-2230132103211311-1113003323202303-3332233232033313-2023331023102102-2033122232121303-0010003031221021"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 103310111310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-0110102113032121-0201110132033020-3222031121312330-0201110121023032-1200330221003012-2203012111222222-0303331010020012-1303301310113312"></a>

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
jumbo_disabled = {}
```

<a id="canonical-0102333202031101-2310201131010022-2113013221332232-1002200121022002-2332133010112222-1220231301023223-2310323322333201-2301122133011132"></a>

## Direct properties — jumbo_disabled / 103310111310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200121021133031-1000031021200200-3001212001310322-2021212203303230-3101012113020210-3131110211021213-2221022102023211-0212221102230102"></a>

## Next pages — jumbo_disabled / 103310111310 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3221132022113323-3023200012030102-3223002021020322-3033130123203132-1220222103110310-2300223303101103-3210331110332001-0302332131202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212032220303303-0313221020000101-1003033222322113-0020011020000303-0013112221221321-1103310101311300-0302201330311112-1120113101232002"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 113020200202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-2030013332020020-2302012001213021-3202220113322202-0321122013212332-3101232002120130-1311002313113310-1022212221110332-3313000003013023)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-2211233122332012-2111013202233302-3202330313330221-2100220312131113-1302212013201001-3120223133320202-2313312201030012-2233102102310323"></a>

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
jumbo_enabled = {}
```

<a id="canonical-1132002202311102-1101131011131101-1330311132200313-0210030300223301-3000110120333302-1323011221113323-1130120233013022-1322113030311230"></a>

## Direct properties — jumbo_enabled / 113020200202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022311300223032-1233233112303321-3222033122323111-3313321001021022-0021022322333300-2022030001022101-3012201111211132-0220302332020300"></a>

## Next pages — jumbo_enabled / 113020200202 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3203013020200321-1203022331202302-3103010323230121-3102001221302313-0201133221222221-1130030103000321-3320321223133130-1202311223003312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1002220320321130-1001012212223211-0300100323210121-2210210103200203-1100221203122113-2010203213231021-3310010103223202-2200012223303032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111322220030212-3203222101020213-3010231022122012-0023230010023320-0121120121112013-1230023200213021-1022013032133011-0013311231030210"></a>

## ingress_egress_gw_ar.sm_connection_public_ip — sm_connection_public_ip / 102002001120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.sm_connection_public_ip

<a id="canonical-2331322313032001-0100103322033033-0100200230201303-3231132322130031-1110103302322110-3002223212313032-3131313313223231-2202120330221312"></a>

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

<a id="canonical-1201210011302121-0101132131003311-2221201110233301-0120003010022133-3121301130213202-2321212121013322-1132231102101220-0112303322032133"></a>

## Direct properties — sm_connection_public_ip / 102002001120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231113033213301-2103222102131313-3121013230131012-0003201321212210-3103203123120213-3122003300201302-3100113011133032-3323021111112330"></a>

## Next pages — sm_connection_public_ip / 102002001120 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0303232313013312-0103020012001131-3212301122000033-0122212122201120-3002021131121330-0113313002023101-2133102132121111-3223200302322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302302132201120-1002331213311221-3223112223223111-1001230312001032-3213013121012132-1223033110112300-0303111313103330-3203013322001321"></a>

## ingress_egress_gw_ar.sm_connection_pvt_ip — sm_connection_pvt_ip / 331303133202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.sm_connection_pvt_ip

<a id="canonical-3112322323010320-3302231121330013-3123110333220230-0323130203223032-3010123011201210-0113313310331332-1023120132022232-3001233023101120"></a>

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

<a id="canonical-2132213233023030-2123223111232232-2110332000321220-1310222011223102-0010003130023131-3021232313221111-1222023211021131-3302103332321010"></a>

## Direct properties — sm_connection_pvt_ip / 331303133202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212330312132330-0002000010221321-2332002312023100-3212032123001013-1033232030003301-3312331202132201-1312001003103303-1332121123302102"></a>

## Next pages — sm_connection_pvt_ip / 331303133202 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333333102013230-2022131023102221-3002020210130301-3123110103002320-2201111323112313-0133001112232003-0023123110000312-1202322103022330"></a>

## ingress_gw — ingress_gw / 320213123221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- ingress_gw

<a id="canonical-0233032011112223-2023011231031022-2312132332312301-0203210101332212-3103130302231103-3301333101222321-3013212220021211-0003031010021130"></a>

Type: `"object"`. single nested block, Optional.

Single interface Azure ingress site on on Recommended Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("az_nodes",
    "azure_certified_hw")}
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
ingress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321333131132023-0233103201112103-3031201032033000-0121312210112233-1201332030313113-1103321002113333-2300232200301331-3100032000101030"></a>

## Direct properties — ingress_gw / 320213123221 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320): complete subsection reference.

- [az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021): complete subsection reference.

<a id="canonical-3002320122213103-1310110013012333-0331131021131101-0303321331123110-1310311320222322-0220200203333232-2133011122013102-2323012232022211"></a>

<a id="canonical-0221122110032022-3031200333113113-2213233302100023-2112333020012222-0201113132103233-0223322210001012-2232302122232033-2221132130332002"></a>

## azure_certified_hw property — ingress_gw / 320213123221 / 4

Type: `"string"`. Optional.

\[Enum: Azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213): complete subsection reference.

<a id="canonical-1221233133322011-0300111203331113-2021332222013030-1321203130220232-2101100331121131-1011311233110302-2122023021323123-1312111003211101"></a>

## Next pages — ingress_gw / 320213123221 / 5

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022132013003220-2001222012112210-1232300030213213-0020232131001120-0013001121223003-2210230013001213-1120221322220133-0322332010112323"></a>

## ingress_gw.accelerated_networking — accelerated_networking / 311332030022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- ingress_gw.accelerated_networking

<a id="canonical-1211331330113000-1203030233233011-0303333331023321-3322311333130120-0201020210223233-3320010231113301-0120111021232130-3322220112201131"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.
Server applies default when omitted.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122102121110230-0130131211200313-3221221332001022-3213001022103110-1320311230012330-3032022323121201-0233303033210200-3200000333131331"></a>

## Direct properties — accelerated_networking / 311332030022 / 3

- [disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-1222323233203133-0321320031202011-0113223323223102-3221122212332212-2233101030110002-3121103230231133-3031211331113121-2120311202113332): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-007.md#canonical-2312222001123302-0233023100201110-3132200010201230-2312321213003222-0231112232311102-1033130321333022-0311300200203012-0313210200323313): complete subsection reference.

<a id="canonical-1013122023130220-0010202201300301-2302111310231130-0021301012000231-3102202131222320-3310201022332311-2300202302003212-0232101313202310"></a>

## Next pages — accelerated_networking / 311332030022 / 4

- [ingress_gw.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-1222323233203133-0321320031202011-0113223323223102-3221122212332212-2233101030110002-3121103230231133-3031211331113121-2120311202113332)
- [ingress_gw.accelerated_networking.enable](resources--azure_vnet_site--reference--group-007.md#canonical-2312222001123302-0233023100201110-3132200010201230-2312321213003222-0231112232311102-1033130321333022-0311300200203012-0313210200323313)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1222323233203133-0321320031202011-0113223323223102-3221122212332212-2233101030110002-3121103230231133-3031211331113121-2120311202113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121021202312200-0313023023021021-1021332033130103-0011323102232012-3203203111212231-3031130133313013-2103222020300200-0001300112233223"></a>

## ingress_gw.accelerated_networking.disable_spec — disable_spec / 111210022102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320)
- ingress_gw.accelerated_networking.disable_spec

<a id="canonical-0330012002013011-3310133103222111-0230111230212201-2021010312303220-0120333221111110-2032222033002222-3130213321231111-0113320133312310"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-2300313012021201-3310223301130002-3330300011023201-3221023330301302-3122012221023310-3212030002102032-3222333023010231-3131333231312122"></a>

## Direct properties — disable_spec / 111210022102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332030111311211-3132121213322000-1020122121303202-1330302211322232-2223100303111103-1303122101221102-1130002133223122-2101123233113302"></a>

## Next pages — disable_spec / 111210022102 / 4

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2312222001123302-0233023100201110-3132200010201230-2312321213003222-0231112232311102-1033130321333022-0311300200203012-0313210200323313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002031221222001-3313133002101012-0332213033231221-1110202202310312-1020321131122203-2300222011022223-1320231012003133-1103012201011212"></a>

## ingress_gw.accelerated_networking.enable — enable / 103000211123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320)
- ingress_gw.accelerated_networking.enable

<a id="canonical-0131201021122220-2313322330123320-1103103223010330-1100120311121323-0011002321212003-2012101003032303-2321000330231013-2220331003201232"></a>

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
enable = {}
```

<a id="canonical-3022222111223310-0101300232133110-1102201033023203-3322032123120100-1100101110320022-0013112021002301-0032033000323321-2223101201221323"></a>

## Direct properties — enable / 103000211123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131301300011110-1320220110103103-0312012133022100-3222121301231020-3012121213331220-2121012113030021-3330032130112230-3203330311030320"></a>

## Next pages — enable / 103000211123 / 4

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-0131311002000031-1200100123330122-2330213012000120-3122110310221331-3232121203223210-2021122211223001-1001213201330210-1211013212203320)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313322223000233-1012203111111233-1113121121131003-0002222302323223-0130213211230200-1221003300033300-3001100333323023-2102333211220312"></a>

## ingress_gw.az_nodes — az_nodes / 022332011122 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- ingress_gw.az_nodes

<a id="canonical-2223002202333110-2021231330021102-1000133113223333-3320010002131210-0131310311220313-1330112020331103-0033103130323201-2100231123223332"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("azure_az")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323230322301321-1131021113232031-0322000123223210-0033222133223110-2110111023002220-2223212031332223-0133130033133033-2100001322330321"></a>

## Direct properties — az_nodes / 022332011122 / 3

<a id="canonical-3230010100312021-0221023211003231-2003031031200020-2021102202310203-2132020100202311-3133200012132012-0023111333010221-0211111122331032"></a>

<a id="canonical-2222023012230330-1203023101123102-0312100203002200-1023213310200210-3103032310122120-3332022033230122-3002330020133231-3332213102112301"></a>

## azure_az property — az_nodes / 022332011122 / 4

Type: `"string"`. Optional.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("1",
    "2",
    "3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "1",
    "2",
    "3"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  }
}
```

- [local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011): complete subsection reference.

<a id="canonical-3312210103120122-3322330323122111-1331001313010200-2222000312201231-1100322121113123-0322000202020320-2220201311111133-3100102021000330"></a>

## Next pages — az_nodes / 022332011122 / 5

- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000213210200321-3002021013002120-0031021213310210-0121210003111121-3100201203011202-1300313100011323-1121023232122210-0231332023332333"></a>

## ingress_gw.az_nodes.local_subnet — local_subnet / 321212233221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- ingress_gw.az_nodes.local_subnet

<a id="canonical-2023123033103122-1213303223121200-3133223221333332-3220203203201111-1222022331123100-1132222113010022-3201310220013321-3003103222003202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333120202120110-2123100310320012-0033022213103230-2021023331000310-1322300233021011-3322323201332032-3122001211120012-3213332130030232"></a>

## Direct properties — local_subnet / 321212233221 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2221003113313011-1300111111323221-2010310231033032-2102100111203111-2323132212323103-3003113123021222-1133111010332023-2131330330100312): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-1223303031213132-3312001003032233-0301221010032331-1121321111130331-3323120123020123-3310001101313022-0112103322012213-1131002021102222): complete subsection reference.

<a id="canonical-0023023300130202-2301210311331121-3313001301013302-1031211222331210-3133221301201323-3032013213121230-2223110332222133-2031133131100312"></a>

## Next pages — local_subnet / 321212233221 / 4

- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2221003113313011-1300111111323221-2010310231033032-2102100111203111-2323132212323103-3003113123021222-1133111010332023-2131330330100312)
- [ingress_gw.az_nodes.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-1223303031213132-3312001003032233-0301221010032331-1121321111130331-3323120123020123-3310001101313022-0112103322012213-1131002021102222)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2221003113313011-1300111111323221-2010310231033032-2102100111203111-2323132212323103-3003113123021222-1133111010332023-2131330330100312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202001220023113-1010201102321131-0232011131033133-0313302211211310-1131213023120222-2003310322130222-0133021101320331-1002010322112021"></a>

## ingress_gw.az_nodes.local_subnet.subnet — subnet / 100222332020 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- ingress_gw.az_nodes.local_subnet.subnet

<a id="canonical-2230300002210210-0121223211232323-2033013331020320-2222323220112223-0033200133021012-1330110211213211-3102213331113230-2323011202110333"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220231111222212-3303022110133002-3303200112300230-2132110113001032-1130010033130011-0213013002100103-0201021121033021-1032122310133332"></a>

## Direct properties — subnet / 100222332020 / 3

<a id="canonical-0312123102303201-3212112002311330-0102032010122102-3003213301001321-1101323003031321-3131200222111010-3022321210113311-2031103200200323"></a>

<a id="canonical-1112013321210130-1122023231033303-0003333022203113-2131313030102312-3312111333302001-0331132331231222-0320102221301123-2102022131211301"></a>

## subnet_name property — subnet / 100222332020 / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3331203112133123-0023032133011211-0022210210101022-1313322330001102-1330233100120133-1121302120213201-3011323133000311-0202303212131232"></a>

<a id="canonical-3310301120131122-0230013133321011-2103010320131230-1321232003012313-2211112010001031-3213223122021223-1121320112130312-0313122133322000"></a>

## subnet_resource_grp property — subnet / 100222332020 / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-0233210323321312-3013303321112112-2232133002123013-0230300122230132-0130133130311320-3111333123112203-1010332321033131-2203120312330321): complete subsection reference.

<a id="canonical-1111102203033210-3002332220031132-1210111123111200-3312211021002013-1311222202101301-2323123233102112-2230103232023101-0033033000100330"></a>

## Next pages — subnet / 100222332020 / 6

- [ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-0233210323321312-3013303321112112-2232133002123013-0230300122230132-0130133130311320-3111333123112203-1010332321033131-2203120312330321)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0233210323321312-3013303321112112-2232133002123013-0230300122230132-0130133130311320-3111333123112203-1010332321033131-2203120312330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202033312002002-2302113211202310-2312023020010321-1213200300103110-3002000002203121-3203311122222320-2332133120031331-0101223023010313"></a>

## ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group — vnet_resource_group / 002130333130 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2221003113313011-1300111111323221-2010310231033032-2102100111203111-2323132212323103-3003113123021222-1133111010332023-2131330330100312)
- ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group

<a id="canonical-0333000232330203-2031300201012130-1120312102331021-1233113122031220-3311033301103200-1333332123000230-0232200133131332-2211231210121133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-3001321311003222-2211121300310230-0212012022022301-0122110021120103-0211302100312313-1122232200121123-1103102020111031-0211303011223022"></a>

## Direct properties — vnet_resource_group / 002130333130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032313211210003-3201102302331232-0033200211103112-3001122231123001-0300331100032000-0000132311211333-3231231000122011-2321103303122203"></a>

## Next pages — vnet_resource_group / 002130333130 / 4

- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-2221003113313011-1300111111323221-2010310231033032-2102100111203111-2323132212323103-3003113123021222-1133111010332023-2131330330100312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1223303031213132-3312001003032233-0301221010032331-1121321111130331-3323120123020123-3310001101313022-0112103322012213-1131002021102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330001221232032-1030203113111333-3101023222101210-3103032030232031-0032303330022320-3022112303232321-1203210320321022-0332333033123211"></a>

## ingress_gw.az_nodes.local_subnet.subnet_param — subnet_param / 101130033313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-2330110122000232-0222010200013221-3203113021331110-1203201211222002-3102120021202102-0320003231200111-0312113323331332-1323001302221021)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- ingress_gw.az_nodes.local_subnet.subnet_param

<a id="canonical-3323312300230110-1031133302233111-2111332112113300-3230223100213200-3121101123111033-1023330021211323-1210012133330331-2321020011030002"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203103321220323-1233113102212012-3202331120103103-3021322100313131-1200313320021130-1111330002022221-3113221301000223-1022033100211000"></a>

## Direct properties — subnet_param / 101130033313 / 3

<a id="canonical-3120312332212213-1102001211233001-2323300023022102-0222020132212222-3321100122233200-2303023023320301-1000021000222231-2120322232003231"></a>

<a id="canonical-0220101121211120-0332232222332333-2001122302321033-1131100230033211-2221012203302002-1303003112333233-2223113321000001-2332332301230232"></a>

## IPv4 property — subnet_param / 101130033313 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3313111132201211-1110011132211131-3301111101320122-1132310211020011-1111320312120112-0022233020002030-0212003113002203-2301103300311013"></a>

## Next pages — subnet_param / 101130033313 / 5

- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-3011020203220300-0201303121130101-0312131121113100-3121131130031130-3020210311013320-3233231120302103-3011312332302032-3122200312023011)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102322133102130-2113021212111032-2121131030303102-1110230230333121-2202203233102332-0112131302231110-2120301001120201-2210012001033122"></a>

## ingress_gw.performance_enhancement_mode — performance_enhancement_mode / 032023230300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- ingress_gw.performance_enhancement_mode

<a id="canonical-0223210100123322-1333222020021011-1110131001020323-0221012101001323-2230201131230232-1011002320100122-0203030111331133-2302212333213103"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. Server applies
default when omitted.

Upstream description:

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023111000233100-0312032233011033-2030322222333333-0311231233032101-3211131310103122-1013021013123023-3332022211020100-2331303332233112"></a>

## Direct properties — performance_enhancement_mode / 032023230300 / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3100323223120223-2203120201312121-1310220303232121-3221023323212020-2220030000012221-2031121122023020-1212200303222020-3221311111002103): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-2132233103112012-2222110332232113-1300103101321333-1033100122023201-2230011222321010-0320300111021312-2122103010233230-0321322011302132): complete subsection reference.

<a id="canonical-3012101303030312-2210131122303023-0213202230221031-0301233221111003-0221022322322301-0023010332200003-1201211310300230-0130003131210011"></a>

## Next pages — performance_enhancement_mode / 032023230300 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3100323223120223-2203120201312121-1310220303232121-3221023323212020-2220030000012221-2031121122023020-1212200303222020-3221311111002103)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-008.md#canonical-2132233103112012-2222110332232113-1300103101321333-1033100122023201-2230011222321010-0320300111021312-2122103010233230-0321322011302132)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3100323223120223-2203120201312121-1310220303232121-3221023323212020-2220030000012221-2031121122023020-1212200303222020-3221311111002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123321013223013-0001123323033010-1122321103322331-3102311122103123-3202010011103203-2210002121332120-0311021201230330-0330321301120133"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 032220111101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2111231233031133-0321022012300110-2030320321232211-0231323323130033-3111033320000102-0301102311200233-1332201100220003-1030300130130301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110132132021030-3030111023101121-3112110200121102-3211132211200312-3321100302103033-2332202323213000-3323222303100230-3011021102312203"></a>

## Direct properties — perf_mode_l3_enhanced / 032220111101 / 3

- [jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-3302023313113310-0203323102302333-1102211030301102-2100201213230032-0300301312310230-1332100222133202-1002202100032303-2231321003110221): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-3213130230020313-0331020033022222-1122213121320310-1322233020322303-1000330030322001-1100112001002103-2120010232010000-0023101221230212): complete subsection reference.

<a id="canonical-2220203231023230-2210010001212133-0112332313033200-2232021030003101-0100322221121331-1130220010232003-1232323301223103-2233111020200300"></a>

## Next pages — perf_mode_l3_enhanced / 032220111101 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-3302023313113310-0203323102302333-1102211030301102-2100201213230032-0300301312310230-1332100222133202-1002202100032303-2231321003110221)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-008.md#canonical-3213130230020313-0331020033022222-1122213121320310-1322233020322303-1000330030322001-1100112001002103-2120010232010000-0023101221230212)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3302023313113310-0203323102302333-1102211030301102-2100201213230032-0300301312310230-1332100222133202-1002202100032303-2231321003110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303021121332221-2021032031003332-0312112300132231-0211031303011101-3031121032223211-0231130332103011-3221133113200233-2010221300022032"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 211002101222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-1333211010132022-0031302333131201-1300202012331203-2322121022221122-1122103331112201-1320012020113002-2311131132332112-1220121001020213)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-3100323223120223-2203120201312121-1310220303232121-3221023323212020-2220030000012221-2031121122023020-1212200303222020-3221311111002103)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-3301221022023131-1232303200122020-0213311201222312-3011123113120310-0300120013213022-1230220023311110-1322111210100223-2112123200231021"></a>

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
jumbo = {}
```

<a id="canonical-2313100033213021-2121010330012300-1032003303000032-3003033031133111-2012233231301112-1311103331311001-2223300033120103-2203020000030333"></a>

## Direct properties — jumbo / 211002101222 / 3

This is an empty object or choice marker. It has no direct properties.
