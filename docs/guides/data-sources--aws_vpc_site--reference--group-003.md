---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110030233123010-2200300313012003-1122300130212123-1113022111323222-3133010001123321-3203131321113212-1201232122211123-3230132133010320"></a>

## ingress_egress_gw.inside_static_routes — inside_static_routes / 031331223332 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.inside_static_routes

<a id="canonical-0303201210020022-3112000332220201-3110101202301103-3120212323223123-0320233213033023-1211321233331100-1233231023012310-1323110022223330"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

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

<a id="canonical-3030312310322331-3101120020111022-3222102201001030-1233312111021301-2001002011320310-0212133013223101-2203213220202213-1300103231210313"></a>

## Direct properties — inside_static_routes / 031331223332 / 3

- [static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122): complete subsection reference.

<a id="canonical-2301222333202102-2303230011032203-1033311102300203-3320013032231331-0122320321322110-1032020112010101-0100031213033302-1122200010200022"></a>

## Next pages — inside_static_routes / 031331223332 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111312111101221-0201100201011223-2321213100032120-1230322121302322-3121320223131130-1312131012002203-0003131220031302-1312132101102212"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — static_route_list / 203323222320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-3223303033002210-0022100213110110-2031211220002131-1321301310202212-2123013112230213-0111001120033312-2301023032010232-0300323010022020"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

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

<a id="canonical-1110220003122203-2302213110003122-1330302021313211-2112313102332200-0313303130330130-1233321120033323-1023201002210002-1123312022302232"></a>

## Direct properties — static_route_list / 203323222320 / 3

- [custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302): complete subsection reference.

<a id="canonical-3331302211212012-2213121333131332-2302222120203011-3212213232112212-2130332121001202-3003322311101322-1211012020210013-0111211232100123"></a>

<a id="canonical-2212132011322021-0012230230202123-0100101313103031-1011101002210332-1302023022130110-2102333110330222-2122200211112102-1312111032003021"></a>

## simple_static_route property — static_route_list / 203323222320 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3201130123233123-1123312033200021-1300103220101313-3110231313320312-2033003113030031-1200131111220133-1122003011113332-3213022312132022"></a>

## Next pages — static_route_list / 203323222320 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011110300033113-0033301132200233-3130213212101111-1220010121123000-1010112110221311-2020011021300333-0021032102221121-3222312113001310"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 301230131113 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-1202023013102313-2001210030220202-1012121233031113-1311123012322123-2032010023111033-3021200120231320-1200010330202200-3030122012113130"></a>

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

<a id="canonical-2022031022212000-3303021232032000-0103112301031000-3220031321001131-3120123130030033-3102102032110112-0313312321332221-1031332330330002"></a>

## Direct properties — custom_static_route / 301230131113 / 3

<a id="canonical-3202332310023032-0122320302331321-3223232200001332-0311203212222323-2223020313222132-1013200120233112-3000002122211031-2233333033021131"></a>

<a id="canonical-2111001013003220-0011001023232313-1202302220113000-2330233111331020-1013213001102122-1303121032302001-0213130131120001-3330001310203013"></a>

## attrs property — custom_static_route / 301230131113 / 4

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

- [labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-1000122221022200-2103030023000002-1001230031032113-2132303210223320-3231122233331311-2330111200321001-3012032102130012-0223212331021100): complete subsection reference.

- [nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013): complete subsection reference.

- [subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323): complete subsection reference.

<a id="canonical-3120122123110202-2011313323332212-2002130202322120-1332021002133022-2120332301120220-2213231230000330-1221312323013101-1033231023210020"></a>

## Next pages — custom_static_route / 301230131113 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-1000122221022200-2103030023000002-1001230031032113-2132303210223320-3231122233331311-2330111200321001-3012032102130012-0223212331021100)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1000122221022200-2103030023000002-1001230031032113-2132303210223320-3231122233331311-2330111200321001-3012032102130012-0223212331021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220031103223033-1103210121022103-2322120232212102-0331302111020113-2220301311001000-0021333313230230-2302200120202301-0313131322212213"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — labels / 023302321221 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1011332011000223-3232021330300110-2031311312101212-2231121210211113-3001003322322032-0323112013103101-0123220120032021-2111232101133010"></a>

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

<a id="canonical-1121322320012203-0320121201020211-2323020111130021-0013221201011013-2002133312111132-2230110313301223-0312023232031000-0013230112003332"></a>

## Direct properties — labels / 023302321221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112330323002211-3002320002313222-1300121222132131-3302131312203132-3133312021332323-3100033022002123-3312331111321223-3101011300210031"></a>

## Next pages — labels / 023302321221 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013030333102033-2333011333023203-2213021212130313-2311003103111121-3133121331202313-3023210211201010-2201130112033112-2233013031200313"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 101100002132 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-2212001302131213-3301322201333302-2221031301132011-0100213023312332-2033311130322022-3023220010002110-2030133300313203-0012101333223010"></a>

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

<a id="canonical-3123321333033013-0210223123121302-2321102022311223-2002200213201310-2100302202231020-1300330223021310-0332122112311132-2221303232001212"></a>

## Direct properties — nexthop / 101100002132 / 3

- [interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-1112110222201220-2211332020102031-1120123100103013-1113030110020213-1000330333012213-2123032010113030-1032030100332300-1010200220220002): complete subsection reference.

- [nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120): complete subsection reference.

<a id="canonical-3211303330121003-3032323113202133-3022031021021131-0322021323013332-3210123330212012-1030020302120320-1312030023101023-2132320323321131"></a>

<a id="canonical-0312333122121123-1230321221222313-1023332033122310-0113210312120221-2332223302233020-1232313000023120-0023313123132112-3002012322310123"></a>

## type property — nexthop / 101100002132 / 4

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

<a id="canonical-0012103201002132-0011333031202202-2112333030120133-0233213123322132-0333202112310022-2330230130201213-1000013211301321-2313032212110232"></a>

## Next pages — nexthop / 101100002132 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-1112110222201220-2211332020102031-1120123100103013-1113030110020213-1000330333012213-2123032010113030-1032030100332300-1010200220220002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1112110222201220-2211332020102031-1120123100103013-1113030110020213-1000330333012213-2123032010113030-1032030100332300-1010200220220002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210202010110111-0110022113312003-2320303322032123-3013003021120332-1300233320121113-3033101020110203-0311100231112003-1230312210002211"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 133233112223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3320113303031103-3331022333010302-2131310032011133-1233221123002021-3022123031103313-2133321300010020-0213221032203101-1101020003133233"></a>

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

<a id="canonical-1200310013332112-2120221122101102-3203302331213203-2223112333020101-3222311131233233-1122203303130313-0120110201122001-0313211223203102"></a>

## Direct properties — interface / 133233112223 / 3

<a id="canonical-2011310213121320-0030132300023123-1233023113303212-0022003012100121-0332211113200111-0303300030300321-3333020220230100-3012103320210232"></a>

<a id="canonical-3223323221320122-3221102113033000-3123032201301023-2000312030213023-3120233031021220-2211132202221021-3301231211021312-1213200120023313"></a>

## kind property — interface / 133233112223 / 4

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

<a id="canonical-0223010130003011-3230132101032012-1130332121232303-1112020012302320-1223332221120303-2223333322323013-1022123312232213-3310121232000102"></a>

<a id="canonical-2003222121312000-2131011231202122-3102330132332102-0320103220023000-3210322302002031-2303331331323013-0220310300312102-2331013322330201"></a>

## name property — interface / 133233112223 / 5

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

<a id="canonical-3230210030232131-0100313133121230-2203231023013020-0223001020121101-0013002033020303-3303130103130212-2110220022110231-0001022110202003"></a>

<a id="canonical-1013110320202330-2210331330120300-0213213203031303-3233023132010102-0330302231231311-0312030333010111-2112220021111301-3330330222112200"></a>

## namespace property — interface / 133233112223 / 6

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

<a id="canonical-0322023013130313-0321122233302232-3323313311330101-1000303302021201-2031213231203023-0220303110331012-0322231221203323-3331021100021311"></a>

<a id="canonical-3333301222301123-1130011301031023-2102312202332313-1333033000231102-3313203131121233-0210233003103012-2132331331230301-3011121213230330"></a>

## tenant property — interface / 133233112223 / 7

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

<a id="canonical-0303012320021131-1223211200203121-3102200322323200-1323303101023303-0232333130123200-2001020121313100-3211033211211130-2130003221120103"></a>

<a id="canonical-0000120303202212-3101310121001023-2002120011002101-2102312303212221-3301033113321330-0210220230233123-3102001122220330-2332323211223132"></a>

## uid property — interface / 133233112223 / 8

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

<a id="canonical-3120222111121000-3320202103001212-2222201212113212-3212030010111010-2232313123123222-0132010320312333-0122131221303312-2220022122203011"></a>

## Next pages — interface / 133233112223 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110033330103312-2101230111332230-1302302030020120-0003001031222303-3123100211110023-2302311021203002-2130332113120323-1022033332212211"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 001200101320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1003313213312333-3203022002112212-0312203220321323-2331211202033020-0121330033213312-0302101023110001-2302233003310313-1011313102111223"></a>

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

<a id="canonical-1022013211122012-1223012310313212-1222013033110210-2101001231002130-1220221200123313-0003003303222311-3203100012211321-1220310102013130"></a>

## Direct properties — nexthop_address / 001200101320 / 3

- [dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201): complete subsection reference.

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-1230020232111213-0311010222312213-3113112000331230-0010220333133023-1011331113011033-2321031312220011-3022330222113220-0113220000033003): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-1100011231131333-2123202230130112-0233110031301133-1301003102112201-1221031213212100-0303002120132121-0012020223133002-2201312301310001): complete subsection reference.

<a id="canonical-0310221112121132-3222000021113031-0113212023231110-2100211110200332-2303222101113210-0033311011100123-1310021001332120-0331331120000013"></a>

## Next pages — nexthop_address / 001200101320 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-1230020232111213-0311010222312213-3113112000331230-0010220333133023-1011331113011033-2321031312220011-3022330222113220-0113220000033003)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-1100011231131333-2123202230130112-0233110031301133-1301003102112201-1221031213212100-0303002120132121-0012020223133002-2201312301310001)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212321230103320-0031000222301332-1123222130110000-3320023011131011-1101301231211331-3100121100230112-0102113102110323-3020210023232020"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 110303200112 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2232302230031203-2001121332210120-0100111330030302-2130212023231120-0311023333220302-2122203120111222-0130120210323013-1211303000023231"></a>

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

<a id="canonical-2132100131002311-1313012212233211-2112033130033310-2232003032323333-1010030120302123-0300233031003113-0102023333121230-0021212122311332"></a>

## Direct properties — dual_stack / 110303200112 / 3

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-3001113102220002-0133020122000121-0212213001221112-1200122100223111-1031202221211100-2022212113232000-0230203202013022-2011130123333022): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-0103312320133220-1022212221233320-1132020033131220-1100110301330222-1133231201333333-0231222120132030-2003202030030132-3300002001321330): complete subsection reference.

<a id="canonical-1312120130201211-3310322210033330-2011123321112031-1310012320133200-0220023133000000-3313001213210231-3323103202321322-2230201310121232"></a>

## Next pages — dual_stack / 110303200112 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-3001113102220002-0133020122000121-0212213001221112-1200122100223111-1031202221211100-2022212113232000-0230203202013022-2011130123333022)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-0103312320133220-1022212221233320-1132020033131220-1100110301330222-1133231201333333-0231222120132030-2003202030030132-3300002001321330)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3001113102220002-0133020122000121-0212213001221112-1200122100223111-1031202221211100-2022212113232000-0230203202013022-2011130123333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320111220310201-3313131333313230-0012002130133102-0030220122323020-2311330233310013-1132112131011011-1133020312232102-3222133301201020"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 022111233212 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-2122331132310200-0311102121133210-2123310020100323-0121322103223322-3301013321301112-1310020010303200-1321321110223202-1012023002223021"></a>

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

<a id="canonical-3133131121132130-1100010203303013-0221212032222222-1223333231102332-3313111010131333-1230221101231230-0123021331321300-1012003121022313"></a>

## Direct properties — IPv4 / 022111233212 / 3

<a id="canonical-0233303332223011-2321122113222132-0111023200202220-2022301030111133-0223120202113232-1233221203022313-2322032330322210-0231102133001103"></a>

<a id="canonical-3232033022130303-0320031122321112-0230033201320011-2300232011103220-1302320022110302-0111330132322332-1011313000200131-3102313003103011"></a>

## addr property — IPv4 / 022111233212 / 4

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

<a id="canonical-3332301331232233-0230301002020312-3021110130202303-3223330213011220-0123033112033312-3223320122011111-2223210133222220-3033221030311222"></a>

## Next pages — IPv4 / 022111233212 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0103312320133220-1022212221233320-1132020033131220-1100110301330222-1133231201333333-0231222120132030-2003202030030132-3300002001321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212332011133110-2020022011002332-1333112311200012-2330320301213302-3323132330312033-0320233200200330-1203310301333003-3213022323220323"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 232233222203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-2003332033112320-1133013122310032-3031331101011101-2112220313110100-2121033231022321-2022122312322031-2132222103231202-0332322332300312"></a>

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

<a id="canonical-2013323020232232-3211213033211231-2311021310300103-2132031222000010-3110212102000322-3113003113311211-0200303111201112-2311202003101232"></a>

## Direct properties — IPv6 / 232233222203 / 3

<a id="canonical-0111031220122120-0303132311113010-2313320101300023-0103101023112322-0102103022002100-1032033221133320-0123113310102123-3211312200103223"></a>

<a id="canonical-0311013130033113-2300012232031032-3332223210020123-0023023020001132-1132230232130201-2322202102322311-1332321103000021-0133130231233211"></a>

## addr property — IPv6 / 232233222203 / 4

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

<a id="canonical-1102312311312112-2130121331201120-2103332002222222-2132021322313230-1201122330023232-3100000111022101-1130232203133113-1102020222232203"></a>

## Next pages — IPv6 / 232233222203 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2301332321120012-0022320231222032-2123202323012120-1111232310313311-0100121100023103-1213121210310030-0230331320121210-0100121131121201)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1230020232111213-0311010222312213-3113112000331230-0010220333133023-1011331113011033-2321031312220011-3022330222113220-0113220000033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323220212222223-2221310120321331-1030012003332023-0010202102303101-0023330231111023-0120101302001100-1102102211222123-0120022321111320"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 002212113101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-3312311231300221-0231301103122221-0113100321031003-1033201203313102-2312121102023213-0000331222200312-1312303202103221-2221101311222320"></a>

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

<a id="canonical-0011213202102120-0300311220331312-2302102111311312-1330021302300312-1123310121122130-2231103013131201-1031123222221310-0232200332310012"></a>

## Direct properties — IPv4 / 002212113101 / 3

<a id="canonical-0110312233313020-3113203331323302-0021211130012011-2013032133303312-0030011030122233-3020122222131302-3031331123031230-3321030002023302"></a>

<a id="canonical-0012112133311101-3021101110332303-2130131012123012-3220231120121231-1313113102310201-3222102023301123-1000120011111031-3000123212123331"></a>

## addr property — IPv4 / 002212113101 / 4

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

<a id="canonical-2133310233220233-0231322211121302-0212212303132332-3103232322000030-3232001033013021-2022102120200130-2113300312231213-2223031110111322"></a>

## Next pages — IPv4 / 002212113101 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1100011231131333-2123202230130112-0233110031301133-1301003102112201-1221031213212100-0303002120132121-0012020223133002-2201312301310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323323233103231-3100212022332300-3330231212220120-2021100021323102-1131222010202132-1221131001012020-3021322230130232-0223120233333233"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 300023131031 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-1310222012311322-1202000121210100-0123122330033223-2113123320113021-3120001222333131-1002323323231031-3012211312023233-1011332313113013)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2313301203310203-2200011331303120-3230030212300302-0233130300311232-2030312000313032-1331330303110002-2201332332021311-2013201210312320"></a>

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

<a id="canonical-0310231130112033-1302120110010111-3022023123011313-0232130103022013-3323201113021230-2311230300322233-3133331002330211-3332310302010013"></a>

## Direct properties — IPv6 / 300023131031 / 3

<a id="canonical-2130203101322122-1300021311203121-1313323201221012-3100023002222012-3020200300113220-2110232023021103-2200202103301022-2221120313232211"></a>

<a id="canonical-1103331013211233-1331121200133100-2313101022322103-2231323321010101-3311131300112202-2023003020032000-3330031222122221-0130112122112000"></a>

## addr property — IPv6 / 300023131031 / 4

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

<a id="canonical-1330231012120030-0012232233321312-3030231301301010-2001211003103232-1201022313231311-3221221132113111-0301103313021202-3012233230022231"></a>

## Next pages — IPv6 / 300023131031 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1201131233100003-0332321102333220-0323303121133231-3132201011002132-3232230032331313-3322223013232311-2302322121102313-2101032330332120)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200013001132202-3111123120103321-2022013300311010-1233231301300323-1311331023201110-1301011031230332-2300313302011330-0301021020002021"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 002033331133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-2122200210312103-0223331310120031-2121202031213023-0122203230210302-3233023233303211-3222222200103111-1220013223032000-3331301220113301"></a>

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

<a id="canonical-1311121102002022-2300331320202203-2331323113022312-3013001333123130-3333120112232330-1032023311030032-1000112312323332-1011222111020203"></a>

## Direct properties — subnets / 002033331133 / 3

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-0013332313230013-3103203003211020-2130033303213013-3322211201233012-1201211111303201-1000323303320313-0013120011130302-1302011023210311): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3000112302010123-3011301203302023-3101232111220131-1101321023112230-3331221131010031-0311303311000332-1223000311102300-3222330300210130): complete subsection reference.

<a id="canonical-2212013300200203-3110330021033213-1003130022111110-0303131130103313-2002331112220121-0310233202010220-3223231311003323-2301212312211032"></a>

## Next pages — subnets / 002033331133 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-0013332313230013-3103203003211020-2130033303213013-3322211201233012-1201211111303201-1000323303320313-0013120011130302-1302011023210311)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3000112302010123-3011301203302023-3101232111220131-1101321023112230-3331221131010031-0311303311000332-1223000311102300-3222330300210130)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0013332313230013-3103203003211020-2130033303213013-3322211201233012-1201211111303201-1000323303320313-0013120011130302-1302011023210311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231210211203302-3001111102322221-2111322230220132-0123100312303132-3220030230220323-0313331333220311-1033320222212300-0122020201023001"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 001303212122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0312111313032003-2031011103010023-0013020023300003-2010210332131021-0313231301332123-1310001110110022-2230023221233102-2303231002301030"></a>

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

<a id="canonical-2311230120330132-0321310110323103-3111311210220100-2000220121101330-3220102222322011-1101033020022121-1211231211302320-3123120012130331"></a>

## Direct properties — IPv4 / 001303212122 / 3

<a id="canonical-1010122101031211-1021032123220222-3320302330101131-0121301130001122-2100023102233311-0021022222003000-1200100100213111-3230331132122010"></a>

<a id="canonical-0123300200303002-2302322300322020-2301210130031100-0110331002001001-1322030220030011-3210010202011232-1013233300113012-3201220220320303"></a>

## plen property — IPv4 / 001303212122 / 4

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

<a id="canonical-3001223213230113-2032323133301320-1012012222332022-3123330121313000-1201330210133003-3213132210030301-1231301221330210-3210222333000000"></a>

<a id="canonical-3300020013332222-2232021201102133-3000020230010332-0022310103231220-3203230231313213-3030103301211223-1120010202301232-2333302222003111"></a>

## prefix property — IPv4 / 001303212122 / 5

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

<a id="canonical-2203130002203131-2022122021222002-0100200111022023-3121022132331020-2022113213321223-0303222320100320-2010000002303011-2102202121212201"></a>

## Next pages — IPv4 / 001303212122 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3000112302010123-3011301203302023-3101232111220131-1101321023112230-3331221131010031-0311303311000332-1223000311102300-3222330300210130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221231100133233-1222233003310022-1120111123011203-1320312302330200-0030133130110121-3310221232201121-2120001231101102-0023212020123322"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 013010331222 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211130102210130-2213223313000000-0033311000200311-3320020103213201-1120322131003333-0220302102223232-3110102222022110-0122223020030122)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3031023311032222-3220103210132201-1300033023303201-3201200300200001-0222300332110010-1033202001000223-0011101030313232-1202122321322302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-0232203033002011-3220010100013203-3322313131320323-3303121131121333-3010123203220202-3003010032022312-2321202130201320-1102200220233003"></a>

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

<a id="canonical-3133022303101321-2130032113030303-2011213323000112-1221200002120103-3311012110203203-0203203212200211-1020313203102023-3222213003131330"></a>

## Direct properties — IPv6 / 013010331222 / 3

<a id="canonical-3022100200030312-1333132012223332-2320023113131123-2222223303013200-1333303130032102-2221110021210332-0123033220002122-0012322213301312"></a>

<a id="canonical-1021100020201013-0133313133332320-2003321102001001-1033321213222211-2000321010322310-2133312313012022-3331330122100030-0301001003330303"></a>

## plen property — IPv6 / 013010331222 / 4

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

<a id="canonical-3111231331213302-0013323322130133-0022200321032213-0013320320300101-3003002221323032-3100311122212211-1020323131113221-2020022020121213"></a>

<a id="canonical-2202020203130203-1132231011000031-2333230201221001-1213202313033101-1101213101202122-0201032312031222-1233301333123213-2232010320312312"></a>

## prefix property — IPv6 / 013010331222 / 5

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

<a id="canonical-0301023202113302-0033031130010030-3011100011023103-3300201133222310-0131131313033232-2222131003011222-2000300122311321-1123210033013012"></a>

## Next pages — IPv6 / 013010331222 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-0022100322333210-0003222202022213-2311102310023013-2331323133221203-2320301011312130-2212321131321333-2231211132333102-2331022011023323)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2113321200233231-3011122321321213-2322321230202210-0323003112201000-0212212222101310-3103103013032210-3111111010332102-2202300302033122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232033111132020-0120300232113033-3232031131300030-1211331021211110-1132300221320301-1120020223303312-2303011212121213-3103210120130331"></a>

## ingress_egress_gw.no_dc_cluster_group — no_dc_cluster_group / 212102022310 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-2211210312030103-0200011301120030-0222121203033211-1131233312000102-2200330312200311-0022031003202322-3023022202220211-0100313201022231"></a>

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

<a id="canonical-0231223010102332-1123020122303002-0010200013013313-1132313002203203-3103320001031230-0102313332300101-1013131210120001-3021031013011312"></a>

## Direct properties — no_dc_cluster_group / 212102022310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322200030011011-1200031110020223-0320122033002021-2232111320102233-2230123301232211-3333233221312333-0332102330110320-1123133320001130"></a>

## Next pages — no_dc_cluster_group / 212102022310 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2312012003101212-1030213320311213-0123230122232202-1310002303110300-2311102011011110-0312001221202332-0122021121031220-1230121022010103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331020010012203-1211103333313311-2310112320020002-1003333331230222-3302323010233233-0010321123211013-1231221133213210-3013333233330210"></a>

## ingress_egress_gw.no_forward_proxy — no_forward_proxy / 221323201320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-0300230323332103-2321221333323132-1303221023121013-1320233201132331-2103111302103131-2231203203310302-1010212303132101-1232001233332032"></a>

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

<a id="canonical-0303121110222000-1320010313211100-1103133320232301-3221132121202201-2112032131310323-3230213011121132-0110222311203000-1201010013112130"></a>

## Direct properties — no_forward_proxy / 221323201320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121231300032033-3032032002101331-0222211030112310-0302330003332333-0211130323203101-0123221011003210-1200300112221133-1322313310203201"></a>

## Next pages — no_forward_proxy / 221323201320 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3012322203311023-1122021233203330-0131320222010020-1011203310102133-1133311130333020-0131031232013230-2110210300100131-2131013320322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000313231121101-2311311201311210-0011020321202022-0032233332112301-3230221320300203-1032212313032231-3232120121002322-3323222002211302"></a>

## ingress_egress_gw.no_global_network — no_global_network / 121201012301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_global_network

<a id="canonical-1311103311110320-1323110302101001-2113332322101123-0300311232310233-1020313323030212-0001311230111210-1210322232323031-0022013131332023"></a>

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

<a id="canonical-0020313123103302-0301202323321032-2133112112003103-2331122321011230-1032223100100000-1223020030333300-1002112301322010-3211203133131023"></a>

## Direct properties — no_global_network / 121201012301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023203023230133-3200312310033110-2300010111011033-0133223123303330-3132333220220022-0101323210022232-3022303323023312-2212120103103120"></a>

## Next pages — no_global_network / 121201012301 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1122332310030122-2301203331321331-3312021331313032-0110313021220211-3000212223310100-3103123301113003-2201001303021001-3310113103222230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212012222030200-3332033030221011-0332223123321232-1223111011133100-1321103220232100-1021011302221110-3210201102030131-2101311330100232"></a>

## ingress_egress_gw.no_inside_static_routes — no_inside_static_routes / 211300210112 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-0232311121302013-1331020113311123-1013112310011320-0032100223000202-0132110323300013-2220101302321112-0223101031012321-3131101102121311"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2101012330002211-3332212210213030-2132302132110213-0121212303013203-0031222103320333-2102320032200123-2221331330133100-0000023300331221"></a>

## Direct properties — no_inside_static_routes / 211300210112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100021033333320-0112300212100203-2012003000112123-2012211022320221-1121321100311201-2203300132110231-0030332130221023-3303133301212333"></a>

## Next pages — no_inside_static_routes / 211300210112 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2322012331311110-3013210130000330-0300331233210121-3001320133210330-3030230023030101-3300221022323202-0233223101003130-3122210101130021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023200200320030-2122232031133220-2003313232112231-0121221302201003-0303020110303132-0111210130021103-0100012031331032-2011233021233003"></a>

## ingress_egress_gw.no_network_policy — no_network_policy / 123310022203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_network_policy

<a id="canonical-2333111013203202-1013230111323233-3230112231011201-2130110312220200-1321202212323212-0303233022220332-1220000102131033-0022233132003111"></a>

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

<a id="canonical-1101320322023002-0223220311032332-3310213321121013-0133122331130320-1022001212100121-2120213100313033-3133013031020202-3333220001213232"></a>

## Direct properties — no_network_policy / 123310022203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321000200113212-3113232001230033-2330113301201330-2211123013233323-3121300010000011-3203100003221300-0303033123131213-3131111320333332"></a>

## Next pages — no_network_policy / 123310022203 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0122030330020231-1020100201333223-2023110003022323-2222101300102211-3213133230113230-1320232212211222-1130122001323210-0110110010011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322213310312030-0100313113310122-1331202001011113-2303020103131103-0002333011220122-2320320003213333-0232002311011200-0011123332010213"></a>

## ingress_egress_gw.no_outside_static_routes — no_outside_static_routes / 100220110130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-0110003023310303-3132020021102112-1132323232032132-3122322210220310-1220000333013130-3133121031203321-3200020012230131-2130231130331002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3300002212120030-2122020132111011-1221221012313221-2320030030302323-2132030232001032-0321211122313113-1020013211220021-2203303323113303"></a>

## Direct properties — no_outside_static_routes / 100220110130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220202011210013-0232001213222023-3123321320032330-1022120231021132-1230201113211301-3021001012131111-1313331023021103-1030123312222010"></a>

## Next pages — no_outside_static_routes / 100220110130 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303330120121023-3223201101211301-1131100020103113-0011002210203233-3130020101110100-3033120103302223-0000310033121030-0301021013212231"></a>

## ingress_egress_gw.outside_static_routes — outside_static_routes / 022002132003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.outside_static_routes

<a id="canonical-1333002123321123-1223022100322122-1203123000033321-0103313310110323-1023012211133212-3200322100022111-1331022121132122-1323232113102203"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

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

<a id="canonical-3132331122110322-1301031033301210-3213211133113103-3202311212013120-1132321331302013-3313120202313331-2012322010000223-2213302131011321"></a>

## Direct properties — outside_static_routes / 022002132003 / 3

- [static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333): complete subsection reference.

<a id="canonical-0120321013231220-1110210231313333-2331101101300122-0223020211013323-2010122113313312-1330101121103020-2200133123220332-2322112012302330"></a>

## Next pages — outside_static_routes / 022002132003 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013011011023312-2311002122033232-1303312011203231-3031220121213101-2322031331300122-0101023101122202-3131330001333022-3323110120321021"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — static_route_list / 002031023120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-3231031330121321-1220322300111003-2231002230130302-2100110021321111-2002020132220130-1301303322000303-1002031300131022-2220101221120220"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

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

<a id="canonical-2211322301202111-1333203330112103-0103201003313332-2331131000223231-3211023210030013-3213121030330201-0231113322203203-2032130302033230"></a>

## Direct properties — static_route_list / 002031023120 / 3

- [custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020): complete subsection reference.

<a id="canonical-2100332222003021-1313101132213132-2000221233332030-3123311333211132-2111330223032011-2320102332030011-3100332112013031-0331013002220111"></a>

<a id="canonical-3130203100301113-3201101330120202-3032311120020332-2300111131112322-2210223233130223-1121022332011110-1232011031312012-1330002100102331"></a>

## simple_static_route property — static_route_list / 002031023120 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1101332312032100-3231121331331222-2122331022130302-2212310233333131-3111112123203023-2300012310031101-1302222320311113-3230333220100333"></a>

## Next pages — static_route_list / 002031023120 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112003003111221-1121321020033201-3212313020203222-3333121100313101-3212211322123220-1203111120122020-0330213013032201-1111003202012323"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 102102230301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-0031213112111102-0133130003131133-3313200223113332-0102321022023232-2113101301110133-1030002220123230-1031303021021032-2113221321202213"></a>

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

<a id="canonical-2320331003131030-2220221122331303-3003011130003100-1020231111030101-1013001333022333-0113030321301203-1100032122300122-3102110021333201"></a>

## Direct properties — custom_static_route / 102102230301 / 3

<a id="canonical-3120222030013233-0003031200133121-1200122320303300-0320013103333211-2322321033203010-0103100000021213-2231201212330121-1302012333332030"></a>

<a id="canonical-0002321313111013-0112121203322120-0023023231333213-3333232232213012-2303010111112113-1302233121300021-3101312133110112-0030221232030111"></a>

## attrs property — custom_static_route / 102102230301 / 4

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

- [labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-1322111133201102-1023202301111311-0021331133121031-0121111020120011-0013230331230002-3003200132223313-3002320320023233-0212220220020333): complete subsection reference.

- [nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310): complete subsection reference.

- [subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033): complete subsection reference.

<a id="canonical-2200213332311232-3101202121311230-0101333231302312-0222113321120221-2213112120002333-2301301001131000-3303020121321031-1102121230101311"></a>

## Next pages — custom_static_route / 102102230301 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-1322111133201102-1023202301111311-0021331133121031-0121111020120011-0013230331230002-3003200132223313-3002320320023233-0212220220020333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1322111133201102-1023202301111311-0021331133121031-0121111020120011-0013230331230002-3003200132223313-3002320320023233-0212220220020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212303320303302-0300203211233323-0311003210231031-3210330232022123-1020330001213310-0233012200220122-2222222112230211-1013020220101210"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — labels / 212121003130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-2030221023203032-2111320232330210-1021231100132023-1013013113302012-1131221122310111-2133222010122120-2310012313232022-0121132033132101"></a>

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

<a id="canonical-0311122322232230-2011201123230201-2122011132020121-2123210232311333-2221020202213100-0112301323230230-0133232332331033-2111112313201230"></a>

## Direct properties — labels / 212121003130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022210200000012-3310201030133110-3213133020002012-3001111323132113-0202020021100021-3021220321103103-1011112201213232-3233233011002103"></a>

## Next pages — labels / 212121003130 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302023333022130-0103331000112322-3200310310103320-1111120032030123-3200131033021302-3011030113321323-3333030231112003-3203320203012310"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 110211330022 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0131001232003312-2101011301303310-0011333231213122-0212322303302112-3102330300031223-0101003002011033-1320010123000331-2123030232003130"></a>

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

<a id="canonical-3333203300321110-3200023121121101-1122323313000020-2303211222332130-1310133232213322-3001110110023002-2220132001322002-0030102111302032"></a>

## Direct properties — nexthop / 110211330022 / 3

- [interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-0131331003200331-2121001320203013-0101033032020230-0200023111033021-2122313213210231-1100303122030130-0100320231221310-2210333122110300): complete subsection reference.

- [nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331): complete subsection reference.

<a id="canonical-2132211123102210-2323013223330102-3023212100223212-1000323111111111-1310332121131111-1311132312330321-3030112322121213-1020202313013303"></a>

<a id="canonical-1011310011023023-2302201123011202-3121112212220022-2021020210111111-0322103300300031-3132111031110211-1201203303111222-0001300202011330"></a>

## type property — nexthop / 110211330022 / 4

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

<a id="canonical-3121230031101001-3021013221102102-2332122210302301-3213220103132002-1120002133132303-3210111110213133-2200011111233011-2122022322022220"></a>

## Next pages — nexthop / 110211330022 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-0131331003200331-2121001320203013-0101033032020230-0200023111033021-2122313213210231-1100303122030130-0100320231221310-2210333122110300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0131331003200331-2121001320203013-0101033032020230-0200023111033021-2122313213210231-1100303122030130-0100320231221310-2210333122110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000301301033331-2101132132213033-2023023132231032-1223213130230120-3303221033332232-3222100301301201-0200303213022333-1101212331110333"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 102300001001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3001021122202022-2300212333100312-2332022133111121-0101031131302012-2210213313311312-3011232320220300-3213213031320022-3301231110032303"></a>

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

<a id="canonical-3121332020112000-1333233003033231-0323312123102201-3013231321311213-2031011201321130-2113313002211221-0030032333232013-0000321121310013"></a>

## Direct properties — interface / 102300001001 / 3

<a id="canonical-3210303323020332-3122203122120213-3020011200330032-2223120330002003-0211310301223003-1001111103230303-0123001302112203-0300321033213001"></a>

<a id="canonical-0200331302000322-1203331201112032-2110310231133311-3010232120112231-3212233203130000-0010120332300001-1121332001110130-2111002231111310"></a>

## kind property — interface / 102300001001 / 4

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

<a id="canonical-1001010020001221-2011102311223010-0233231202203021-0311232302033331-3300011120231023-1331111211202130-1022213013213000-1130032300030312"></a>

<a id="canonical-0331301200112213-0213130122311010-3001332303320212-0000320111303223-3130011110011221-0313233211222123-1220031213302332-2221201012201113"></a>

## name property — interface / 102300001001 / 5

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

<a id="canonical-3220323030013321-0322010323133010-3110333112333112-1010211022011102-1130000020120001-0110123021333013-1112332311220023-2203001132100330"></a>

<a id="canonical-2200300112131020-1303213233232011-2232110112111120-2211210022232202-3331230230223101-1120221032302023-3303330121010213-2000000302121122"></a>

## namespace property — interface / 102300001001 / 6

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

<a id="canonical-2330303113200222-3201231033000013-1120220111123200-2212010101010013-3202102012012102-1333331133312010-3111011012313212-1001311321131131"></a>

<a id="canonical-3301031211331010-3203310112302231-0012130101332011-3210300102033033-0212133103310132-0111002020101033-1200300031110020-0201222000001220"></a>

## tenant property — interface / 102300001001 / 7

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

<a id="canonical-2032320120333302-0110012123331010-2121002010233320-0211220311233012-2233023211221322-1310111110121313-0011321232001210-0002222200201321"></a>

<a id="canonical-1112023331231032-0220011010210133-2123002030222230-3233203221100111-1231000000030110-0202331021010102-3003233113222223-3301031331103013"></a>

## uid property — interface / 102300001001 / 8

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

<a id="canonical-1232310102200322-2001311110122200-3012110211020021-2010112202312133-0313000133312112-2320333323302200-0230332002232000-1201102012313103"></a>

## Next pages — interface / 102300001001 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102233023202220-2111132021123320-0320203303322032-2012003233310002-1330001130000310-3131213100331130-0210233232111210-1331303302112332"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 010121211203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-3113020201132233-0001101021311220-1202311302222213-3223311123220223-0021010321232231-0203231121012230-0030203230002020-0112122233032223"></a>

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

<a id="canonical-0323122000231322-2222012101030323-3231110222122203-3033001323122120-2120030130030311-0213231221100202-2011301122120330-2222211133132203"></a>

## Direct properties — nexthop_address / 010121211203 / 3

- [dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020): complete subsection reference.

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-1110030320303122-2200133000020133-2210330220113023-3103102321003221-3110101000233230-0113333203202220-1103212200300320-2221131003122033): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120233230113331-2002102033130230-3303101311131222-1011223002001013-3101311212032232-0233110013101113-1331221032303100-2003021130200223): complete subsection reference.

<a id="canonical-2232332231101233-2223102232123011-0113113122232021-0303111211210012-2012321211320113-1302213202213120-1003320333102231-0231201023110223"></a>

## Next pages — nexthop_address / 010121211203 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-1110030320303122-2200133000020133-2210330220113023-3103102321003221-3110101000233230-0113333203202220-1103212200300320-2221131003122033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120233230113331-2002102033130230-3303101311131222-1011223002001013-3101311212032232-0233110013101113-1331221032303100-2003021130200223)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120021210231331-3213213001333033-0212132131330330-3232111032021331-1203232002021132-3311123203301331-2222033001312222-2331300122111002"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 010100211321 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0313132331203131-0121232001311310-1313232301210231-2011013322200003-0112112023020201-0102033210331221-1031031201323021-0303102030013332"></a>

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

<a id="canonical-1022222311121012-0220312103101200-3013113313330202-3323110213310320-1202210122010303-3102222020202202-3012033103311103-0013321123100110"></a>

## Direct properties — dual_stack / 010100211321 / 3

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-2032101231102123-2312103102333231-1300202322001302-3303033300231020-2102012312132011-2231203301332013-3110012133321103-1013102131223132): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2011033320231123-3303120313022322-0032211202002233-0320020322232220-2222230011113030-2220301322031202-0231132033232000-0220021123231111): complete subsection reference.

<a id="canonical-3200032032010111-2302210332220033-1221230202113301-2222132121010110-0013220312332332-2000100002333023-0121000001023220-2103322221231011"></a>

## Next pages — dual_stack / 010100211321 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-2032101231102123-2312103102333231-1300202322001302-3303033300231020-2102012312132011-2231203301332013-3110012133321103-1013102131223132)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2011033320231123-3303120313022322-0032211202002233-0320020322232220-2222230011113030-2220301322031202-0231132033232000-0220021123231111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2032101231102123-2312103102333231-1300202322001302-3303033300231020-2102012312132011-2231203301332013-3110012133321103-1013102131223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030003222221330-2103101300333210-2301130112021120-3110103000031001-1100331013301313-3000202202211312-1112002000001111-3101223001223220"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 312032000122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-2323313020123101-0121123122110321-2333330212221210-2023332011111202-1101211122330123-1200223312013111-1213012203011321-0121231031202230"></a>

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

<a id="canonical-2211121210321021-2220031020213311-0201333230330302-1233131012333221-1321121332220203-2330220210113233-2333022332220103-3310330011022100"></a>

## Direct properties — IPv4 / 312032000122 / 3

<a id="canonical-2222222211302300-1320032002231002-3101130023211302-2300033023323203-2220003011330213-0013023231111202-2002101103113102-2310200222213310"></a>

<a id="canonical-1211320101330010-1321320031331001-2210120031001100-0000100323033133-0111132131100201-2301200220020102-3301010101200011-2300001230031221"></a>

## addr property — IPv4 / 312032000122 / 4

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

<a id="canonical-3213323033020110-2022303110132212-0123320000021131-0002323302330101-3100301003020333-0313221000320200-3201313222322212-2212310230222322"></a>

## Next pages — IPv4 / 312032000122 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2011033320231123-3303120313022322-0032211202002233-0320020322232220-2222230011113030-2220301322031202-0231132033232000-0220021123231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103002302223021-1103300311001003-2033301303002231-0330111130002013-2102021201003112-0223002101121303-1231031011010020-0313202032100310"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 332221101021 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3103310110011221-0220001331023321-2210032201000332-1100301010201121-1130302133331111-2312212213331230-3102121131323322-2102212202121122"></a>

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

<a id="canonical-1212111120221021-0323302200131331-2102331323031023-0113331312011313-1003220230101320-1321201111232330-1000232201100220-2030201312112301"></a>

## Direct properties — IPv6 / 332221101021 / 3

<a id="canonical-3123311020212103-2132312301211221-2233201032130322-0332033001330023-1130113111233233-2101112132031331-1322212202003022-1200303313111021"></a>

<a id="canonical-1300103213202102-0132202302001300-2310031011221311-1322230002212202-1202320303200311-3321110031021130-0331303231001202-3210111322232021"></a>

## addr property — IPv6 / 332221101021 / 4

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

<a id="canonical-3310210223100201-3131223122122310-3220033321231022-2200212010103103-2112201322003103-1310111121131010-1031103130100202-3311123320100302"></a>

## Next pages — IPv6 / 332221101021 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0320102012022320-3002310001203132-0130230110033032-2000010123010103-2020313133202122-3320232112130333-1202112210101113-2101021311101020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1110030320303122-2200133000020133-2210330220113023-3103102321003221-3110101000233230-0113333203202220-1103212200300320-2221131003122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013201011210012-2032321200001003-1312020100103122-3122001332022330-2033212130120012-0220303320233332-0113230320013021-1221213210032230"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 132231312003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1132021100000313-1320101113220233-0103103220033201-1211302001123201-1323212021001100-3330022001213203-2332213010333320-2330202221330300"></a>

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

<a id="canonical-3112031310333130-0301023022013123-2113332031032002-2213200202012311-1310032110222113-2310333101323202-2001000031210121-0200101322011001"></a>

## Direct properties — IPv4 / 132231312003 / 3

<a id="canonical-0020323323312111-1311112211320332-0321103333023100-1233010320010120-1101021302200131-0003322101121130-3103101331333132-0203321220021013"></a>

<a id="canonical-2022300102322210-3220322021221233-2301200023112113-0330322233202122-1311130020302322-3030202102203030-1133122131123033-1121211233010003"></a>

## addr property — IPv4 / 132231312003 / 4

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

<a id="canonical-0003220222130201-0022301212000013-1132100203331200-3130210001101322-2133323221321033-0223002100003221-0203202312102301-3210203111133111"></a>

## Next pages — IPv4 / 132231312003 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3120233230113331-2002102033130230-3303101311131222-1011223002001013-3101311212032232-0233110013101113-1331221032303100-2003021130200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333123330110100-2121000210020311-2111030220020203-2323211302210301-3133203233020231-3310012001220202-1203332111000202-1032131103030202"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 323222211310 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331033311123230-3310100110021000-0313022311222202-0100322111112003-1213203031131111-3332222001220121-3303231223332033-1312232312100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2111003000031110-1231302102110033-2102113033201330-0203330003032023-0211220100313103-1023203333330031-3320330130012213-1333312330112033"></a>

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

<a id="canonical-0230123111010230-0232313222323010-1313100030120122-0132311110030333-1231020001002113-2032020201300002-2121211303013013-2111113230200012"></a>

## Direct properties — IPv6 / 323222211310 / 3

<a id="canonical-2111310121221311-1022313311313231-1210031001121312-1200312231112032-0133032232110213-2330032131320332-2002121221230122-0111120032322110"></a>

<a id="canonical-2213202323113223-0212331020103123-1011200320330233-0232102200320010-1001002013020120-1110003313320000-0011231213021032-0131200211112010"></a>

## addr property — IPv6 / 323222211310 / 4

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

<a id="canonical-1331222322213331-1212201232020311-1231212013000003-3323222212112003-3222201212012330-1201210013332330-2303220220110331-1133010322330031"></a>

## Next pages — IPv6 / 323222211310 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200202312330033-1300112323301121-1303211230122012-0212310332022310-1300103120112301-0200101120312333-3033033313012301-1112333130231331)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021123322300113-3001032220213310-2200233102023332-1011332310330021-1213122333221302-2131203210212322-3223311100310010-2132223031003333"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 011313322311 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-1113332033200311-0121101231101103-2202131323103201-0330120331203232-0003013202312111-3223112011312102-1212332113032000-3230112321203103"></a>

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

<a id="canonical-1120002201332131-2203123221213013-1112202120022131-1003313010013032-3220321122022012-0333111122123111-1103201202303022-3110303000101312"></a>

## Direct properties — subnets / 011313322311 / 3

- [IPv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-0012102012003023-0101300133113120-0232013022123320-3331001131020131-2122031100320233-2030012011033300-0022120313211120-3202011101313321): complete subsection reference.

- [IPv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2031213010321011-2013103332322000-1333021230331033-3210021332021312-0303003203031012-0203113002021303-2311111111301033-1303131330101300): complete subsection reference.

<a id="canonical-2003300322121333-0111003322302133-2302023020113023-2112200213023311-3020222333122321-2002311012211130-0102300323231023-1212210230030213"></a>

## Next pages — subnets / 011313322311 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-0012102012003023-0101300133113120-0232013022123320-3331001131020131-2122031100320233-2030012011033300-0022120313211120-3202011101313321)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2031213010321011-2013103332322000-1333021230331033-3210021332021312-0303003203031012-0203113002021303-2311111111301033-1303131330101300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0012102012003023-0101300133113120-0232013022123320-3331001131020131-2122031100320233-2030012011033300-0022120313211120-3202011101313321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222311021322122-1221111300111201-3133220312231001-2233023210230031-3031330211033313-1033012113133322-1332323313331201-1322122101202222"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 330211033111 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-2230121203020333-1301032321113120-3031232023233110-0011000132131210-0022323020023202-3101111310313230-2320331001322313-3132311213030233"></a>

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

<a id="canonical-3031122320201210-2033233003312203-3201330102220300-3033221123023322-0230331121010213-2200103031322131-1013023021312311-2330202210232232"></a>

## Direct properties — IPv4 / 330211033111 / 3

<a id="canonical-0321012001211203-0020323333302000-2101221022202030-0313232200322301-0011032213310130-1113313002113022-0000320130230320-3313121332303303"></a>

<a id="canonical-0200321012122023-2002231032130010-3222221132203013-1011332220321012-0300101000301133-2302331322203020-1323222102023111-2220213202331210"></a>

## plen property — IPv4 / 330211033111 / 4

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

<a id="canonical-2000122210202301-2032201111012212-1130123012013210-2210201103312022-2333333020130103-2322323231332013-0102020212310002-0121011132113032"></a>

<a id="canonical-3302022123220333-0031232031033003-1221002221130013-2210103120222201-0022211302301312-1322211101220010-0001223310300102-2221223000202211"></a>

## prefix property — IPv4 / 330211033111 / 5

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

<a id="canonical-3222322303102030-1000311032031333-3322200310230121-0222110121313010-2210111132300020-1001130232110131-3113131221110000-0021211000031012"></a>

## Next pages — IPv4 / 330211033111 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2031213010321011-2013103332322000-1333021230331033-3210021332021312-0303003203031012-0203113002021303-2311111111301033-1303131330101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320110031323123-2312003301321122-3010300332322301-0132220202310000-2321210033222212-2232132222213101-3302132110100213-2331323113001120"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 200100012200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3200210331323022-2330213330121300-0030032020102311-3211010103220111-3002012223113332-1332211103333320-1100031021310133-3232232002100333)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3300022112221023-1220203130020311-0023233201220000-3223213333212123-1110023111100000-1331311130222100-1300102122332102-3000033102302020)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-3223112313021333-2132302123132202-3101200330331313-3312022102202003-1210030031202003-1230210012121320-1111301002133233-3000013132003131"></a>

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

<a id="canonical-1223122010330012-2330321012211331-2222322222101030-3300000110020230-1201211030233322-2331200301311021-2023212301123012-1202302221033031"></a>

## Direct properties — IPv6 / 200100012200 / 3

<a id="canonical-1330012003021032-1332332003201100-1301023211012011-1020200112030220-1013320022131223-2320212001311320-0321210123213113-3232301213013331"></a>

<a id="canonical-1331030120001222-0203102123203231-2132301131213032-1302331300031221-0020112210000101-2202001222223233-0220221002012113-1100033313203333"></a>

## plen property — IPv6 / 200100012200 / 4

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

<a id="canonical-2110333211030022-2120022223113233-1333000322211201-0130011202231102-2033122311212223-2103002013321231-1032112020111310-3313002301313113"></a>

<a id="canonical-1000210033232000-0313020000320132-3200332232012132-3011231302003302-0020220221311000-2313210101220223-0003033112022103-3310012333023231"></a>

## prefix property — IPv6 / 200100012200 / 5

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

<a id="canonical-1120112100020130-1313221233211111-0001310230232233-2101203012203031-3322000003301021-3023013132213201-2013212121121112-1212200101131331"></a>

## Next pages — IPv6 / 200100012200 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120232201101102-3133132312030212-3010003321303332-1211001322320002-1221310311203201-1100223110122333-3213310333311132-2112213113103033)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011323111031213-2330223322322102-1300312103022033-0023133212002221-2131002232232212-0202320003103203-1213110011133323-3130201300212112"></a>

## ingress_egress_gw.performance_enhancement_mode — performance_enhancement_mode / 033321333330 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-1000022222111013-1101030000222200-1330102312303203-2330113330323212-1331221321000333-1101033122221110-2130221133222011-2031332022023202"></a>

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

<a id="canonical-0310222222011222-1012323130231002-3230011003021320-3332300311313230-2012201221223112-3231300323111303-1221132032031013-1003023321220000"></a>

## Direct properties — performance_enhancement_mode / 033321333330 / 3

- [perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200): complete subsection reference.

<a id="canonical-0110033223112101-2320320201011101-2303010103110003-1003301320203022-3321100023312112-3302212011013023-0102323302010010-3112232021031301"></a>

## Next pages — performance_enhancement_mode / 033321333330 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321012000231313-1300311001031003-3013211220333133-2032103033132320-3332001123210202-1321221112220131-0312302333131031-2133132233121231"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 320301012022 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2020230303210112-1103012312113022-2112010030201210-3000011132023232-1322121233001120-1233333331210122-3000000111010110-0232200321231213"></a>

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

<a id="canonical-1032023311303233-3212030303002022-3130203101212121-1303220223333221-0230110330320012-1121332133202222-1033230333022131-0001323212302132"></a>

## Direct properties — perf_mode_l3_enhanced / 320301012022 / 3

- [jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-3122211231001302-0003010112130023-0302131310213030-2220031231120011-2000110300203130-0032333301302303-1212022333012120-2301133303031211): complete subsection reference.

- [no_jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-1002223013113300-3321202012130020-2203312030231323-0233032312022220-0131331322330013-2012323003211230-2310233000311322-1003200313012101): complete subsection reference.

<a id="canonical-1310022233001102-0311132111002331-2003021212100132-2112222113230110-1113233132101213-2200022302020233-3313200233110013-3322031120230231"></a>

## Next pages — perf_mode_l3_enhanced / 320301012022 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-3122211231001302-0003010112130023-0302131310213030-2220031231120011-2000110300203130-0032333301302303-1212022333012120-2301133303031211)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-1002223013113300-3321202012130020-2203312030231323-0233032312022220-0131331322330013-2012323003211230-2310233000311322-1003200313012101)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3122211231001302-0003010112130023-0302131310213030-2220031231120011-2000110300203130-0032333301302303-1212022333012120-2301133303031211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333320320100311-2310230003210311-3132313103011303-1311101022211131-3033112023011213-1031111313022231-3022313021331102-3311130130201012"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 210300222100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2131233032001322-1220313110210102-1112033102100132-1303112310111311-3202011232203201-0330323211023302-0323200023103203-0311012310130133"></a>

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

<a id="canonical-2120310320030121-2133000321032020-0112022313113100-1203323033121133-2120001331323103-1020323122331233-1120123100303330-1100201330012121"></a>

## Direct properties — jumbo / 210300222100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310203202333010-1032332320303130-3113010211200222-2312332131202000-2011211123011112-2233331021132010-3322103230312030-0123032123031131"></a>

## Next pages — jumbo / 210300222100 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1002223013113300-3321202012130020-2203312030231323-0233032312022220-0131331322330013-2012323003211230-2310233000311322-1003200313012101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332322303111310-0232121321022023-0222011132003011-3220011232121001-1310212210022333-3320021011031232-3130313022232103-3201130330123112"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 310212132112 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2013222301311230-2012221000330123-2131033113333012-2032331310321230-3133233120023033-1130233100021222-3030003011213332-2333012200030020"></a>

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

<a id="canonical-3323032213223202-2320230023312201-3211222003323302-0311233323321100-1210220101130321-2302110231301113-2111012300223021-0023113022103221"></a>

## Direct properties — no_jumbo / 310212132112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210003332312212-1201222110210322-3131230232001100-1302002000111213-0101322301023323-2311102123221320-1302101031031230-0232022200111112"></a>

## Next pages — no_jumbo / 310212132112 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322221313323123-0332121113000010-1202303002330003-2230310012312022-0231210301203122-0032022211323112-1030020121100311-0122011322113210)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010213021313233-0200111030123332-2110310032020113-0103100203133311-2103310102311102-0312003231311221-1022320313100310-0330000202202021"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 200222310013 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-2002133213032303-1003301102230332-1313300003012022-0323333101111123-2000220313200201-3302300320123122-2310113211032203-1100301111332313"></a>

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

<a id="canonical-2302133123210230-2302311003312302-2320200003203320-2210003310210230-1230102303131003-2011122020103033-2303323303230033-1233232323101321"></a>

## Direct properties — perf_mode_l7_enhanced / 200222310013 / 3

- [jumbo_disabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-0111210021331211-3333123032110321-0013232313211102-3231112120302302-2330223313213021-1202133121310011-3012133120031031-1310132132022201): complete subsection reference.

- [jumbo_enabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-1133113331031033-1223233232110003-0123102120112031-1223132102022122-2123213011311330-3131210232202002-0101212211322011-2120121123112300): complete subsection reference.

<a id="canonical-0111123333222200-0101203313033331-2223032303203132-3303230033211210-3131333032021311-0113033300021012-1221202311302033-1100322130311331"></a>

## Next pages — perf_mode_l7_enhanced / 200222310013 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-0111210021331211-3333123032110321-0013232313211102-3231112120302302-2330223313213021-1202133121310011-3012133120031031-1310132132022201)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-1133113331031033-1223233232110003-0123102120112031-1223132102022122-2123213011311330-3131210232202002-0101212211322011-2120121123112300)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0111210021331211-3333123032110321-0013232313211102-3231112120302302-2330223313213021-1202133121310011-3012133120031031-1310132132022201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032123133310332-3321332311120111-2220130113322133-3333132331303212-2002321130130310-3212002330011303-2313311023110302-3332201030011133"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 212103011123 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3002200332030030-0013201012203103-2302221102231130-1331021233201132-1200302211302113-3320222022202332-3331303111030301-0320001122102002"></a>

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

<a id="canonical-3203121222301120-2211301022230001-1033011003220231-1130232121113212-1301202013203320-2330103211232133-1003000321213120-0312301131131003"></a>

## Direct properties — jumbo_disabled / 212103011123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033230100211231-2003010001021011-0132031200322121-3332133331221233-1011022113331303-0322223031310300-0222011000200211-3221203001300310"></a>

## Next pages — jumbo_disabled / 212103011123 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1133113331031033-1223233232110003-0123102120112031-1223132102022122-2123213011311330-3131210232202002-0101212211322011-2120121123112300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223001321110020-1230031131303320-0230031020201031-2330323022131330-2123201322332032-1212002122301221-2331200100021333-2222003003230222"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 133112133301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-3112311003230003-0320321132131021-0110001230202100-1102003123010013-1330222111111023-0033233303012203-2210230302123013-3010023021113121"></a>

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

<a id="canonical-2203233033213231-3212321320012322-1332101012311133-3311000100231111-3212112221200101-0132310210020121-1020220232212132-1102003112311213"></a>

## Direct properties — jumbo_enabled / 133112133301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033033220201223-1003232030313303-1202220112313233-0031031123302211-1022000301020033-1212321112121212-3132011313333130-2221201011230310"></a>

## Next pages — jumbo_enabled / 133112133301 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-3321221121010111-3201033012232222-1331121313113022-1333221102300002-1322322021310202-2023213002001033-1010301131330013-1030210123130200)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1032302033302121-2121112031020130-2120010231320100-1213021300031120-0122002321131212-2311120210210313-3103011330111033-2012123323331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212231201331321-3100110023113013-1232303330101033-2020113311121320-0311233221101311-3221020003003023-0213210032231230-3301013000023313"></a>

## ingress_egress_gw.sm_connection_public_ip — sm_connection_public_ip / 101103132303 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-3320023331003320-2112032331232113-1013011122122022-0133023120012002-0212311210233022-2303030011112023-2330011121112333-1201113311110311"></a>

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

<a id="canonical-2102111303212122-2223331023200110-0233213201030333-1233212022112202-1102100020102312-3110133213032121-0320100303310032-3220100011122002"></a>

## Direct properties — sm_connection_public_ip / 101103132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332331300202103-2331333132011323-1222302301303113-1201200132203113-0302121011233220-0211200023230313-3130333100000020-3012100330021230"></a>

## Next pages — sm_connection_public_ip / 101103132303 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1031100033200031-2003100101032231-1010002231031223-1103230320030133-0110003322321200-3003223200303311-1023021301231301-1232111011220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230321210123223-1220230103032012-1013220333023013-3032233003321201-1222001201232201-3023032022101012-2302002022203023-2201103201120000"></a>

## ingress_egress_gw.sm_connection_pvt_ip — sm_connection_pvt_ip / 113100013000 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-0120333321313220-1323333211030301-3200022223300310-0301211201110312-1220133101332230-0321012221333101-1230121331031031-0103121331211333"></a>

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

<a id="canonical-0112310001213031-0233331300322231-2233001303332001-0102002202323113-3200113002301222-0102111032033300-2232111010201303-1211300313110022"></a>

## Direct properties — sm_connection_pvt_ip / 113100013000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131330131003211-1322033210313232-0313211020301232-2032110100202202-0000001322233013-1001203312002221-2330312132233322-0130230001110101"></a>

## Next pages — sm_connection_pvt_ip / 113100013000 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002300120021302-2221000200320230-1202323113010333-3301110313301102-0030233303023133-2101132013202311-3203002133231012-2212202021102202"></a>

## ingress_gw — ingress_gw / 310011322213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- ingress_gw

<a id="canonical-0211201331003211-2132210332011032-1003100231033223-3330301033201113-3213013312201031-1131230330130022-1313121113320212-3103122022210202"></a>

Type: `"single"`. Computed.

AWS Ingress Gateway. Single interface AWS ingress site.

Upstream description:

Single interface AWS ingress site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203202311331131-2130002213313023-1320303303033333-3133012323331123-1111312210301010-3022330322320210-1311332322010200-1203211202213132"></a>

## Direct properties — ingress_gw / 310011322213 / 3

- [allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230): complete subsection reference.

<a id="canonical-0332130213330220-3213103201230030-2122213301102202-2212110331110030-0312330332300103-3222021232133112-0001332210031123-3122133230320001"></a>

<a id="canonical-3332101333110313-1132003020333113-2203321323123220-3013310230031312-1232223001220003-0302023233003221-0001001020310111-0112313013100210"></a>

## aws_certified_hw property — ingress_gw / 310011322213 / 4

Type: `"string"`. Computed.

\[Enum: aws-byol-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-2030121122110111-0110300000301102-1313222122231003-1323213211300101-0203321132303001-0331031331110311-1122120030120133-0210011013223123): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-3000012212001333-1122300021203100-2303200102332013-1000033233322203-1233130010232011-2130322020102303-0232201320213132-1202313200021300): complete subsection reference.

<a id="canonical-1302022213022203-1021233123021221-0222201121000012-2330101123003220-2000220303203211-3231023202322333-1033003323032230-1123113221310212"></a>

## Next pages — ingress_gw / 310011322213 / 5

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- [ingress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-2030121122110111-0110300000301102-1313222122231003-1323213211300101-0203321132303001-0331031331110311-1122120030120133-0210011013223123)
- [ingress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-3000012212001333-1122300021203100-2303200102332013-1000033233322203-1233130010232011-2130322020102303-0232201320213132-1202313200021300)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302221321003102-3233200112032210-1221223333333330-1130230210030203-3203202011030221-0231330002311223-3021310231312223-1312332021133001"></a>

## ingress_gw.allowed_vip_port — allowed_vip_port / 012213013232 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- ingress_gw.allowed_vip_port

<a id="canonical-2202201322020111-3232021323231100-0321213223112133-3031301232111231-1220012211000223-1320012021333310-0012322212302222-0202332322232231"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-2030233133312133-3203020030123330-0013321301131200-1212300112123101-2002012310131031-2103210211303202-2003020121113311-1330300311132013"></a>

## Direct properties — allowed_vip_port / 012213013232 / 3

- [custom_ports](data-sources--aws_vpc_site--reference--group-003.md#canonical-1111201203230103-0011121102131020-0312230010201111-0203023232102311-1013301003300201-2023213010031330-0023232210003120-2103310011122123): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-2332223212001012-2231101032001130-3300302310031102-1200120301020203-0133113222032313-0020312002011020-3010020100302122-0213110232020213): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1303303202210103-1301310013032130-3021210312031301-1202003131332112-2320110121001233-0002111223303000-1122010131122001-2122301123330221): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1211013112322210-1030133330222333-0223020133003031-1222030210032211-1023012220012200-0221212202122132-0230011201030330-1321032000122213): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-2310031033231133-1203001123131103-2003211201033100-1031120301232112-2113211222023330-3220213312101023-1113310303211130-0221132202013212): complete subsection reference.

<a id="canonical-2003320223222021-1022210321021110-3301220311133100-3110312333010013-1222012002032032-1310311113132000-3303103311321230-1001321303323022"></a>

## Next pages — allowed_vip_port / 012213013232 / 4

- [ingress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-003.md#canonical-1111201203230103-0011121102131020-0312230010201111-0203023232102311-1013301003300201-2023213010031330-0023232210003120-2103310011122123)
- [ingress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-2332223212001012-2231101032001130-3300302310031102-1200120301020203-0133113222032313-0020312002011020-3010020100302122-0213110232020213)
- [ingress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1303303202210103-1301310013032130-3021210312031301-1202003131332112-2320110121001233-0002111223303000-1122010131122001-2122301123330221)
- [ingress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1211013112322210-1030133330222333-0223020133003031-1222030210032211-1023012220012200-0221212202122132-0230011201030330-1321032000122213)
- [ingress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-2310031033231133-1203001123131103-2003211201033100-1031120301232112-2113211222023330-3220213312101023-1113310303211130-0221132202013212)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1111201203230103-0011121102131020-0312230010201111-0203023232102311-1013301003300201-2023213010031330-0023232210003120-2103310011122123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213000121303122-3203121111133320-0232303133221031-2323020120011321-1303230001330210-2301332020331101-3000300102022311-3030322003231312"></a>

## ingress_gw.allowed_vip_port.custom_ports — custom_ports / 220113213122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- ingress_gw.allowed_vip_port.custom_ports

<a id="canonical-1202101003321332-1020201310011232-2203311111002123-1231031103103203-3331203102222123-1213123311001001-1223303130133312-1012210121011213"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331112012033313-1121010120010223-1013302202223132-3320323322331031-2122230132211132-2133303121323022-0232202303032123-3123020033233331"></a>

## Direct properties — custom_ports / 220113213122 / 3

<a id="canonical-1311121133210302-3330122010123120-0301332210311002-2231303120202203-0211030012332130-3100221332120010-3300320220133210-1212203112230303"></a>

<a id="canonical-3021332230323213-0120333332013232-1223303130313022-2130301003102233-1000030031301312-1000202221323012-2311011013220332-2132301132220011"></a>

## port_ranges property — custom_ports / 220113213122 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-2102310010213312-0030121101333112-1113203300112020-3111011003030210-1111310030332333-2132200013132121-3201001120233331-3200133303200100"></a>

## Next pages — custom_ports / 220113213122 / 5

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2332223212001012-2231101032001130-3300302310031102-1200120301020203-0133113222032313-0020312002011020-3010020100302122-0213110232020213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130031033202312-0322002322220331-1130103031113131-0230133022030020-2201110311321111-3012301102232320-0021321022111231-0030003031131302"></a>

## ingress_gw.allowed_vip_port.disable_allowed_vip_port — disable_allowed_vip_port / 102302313023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- ingress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-0210133201012121-2101321101330200-2023212223333123-2113101130011310-2233132121111231-2003021011000211-3102121323230102-1020330130301132"></a>

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

<a id="canonical-2010320130212102-0231220311220232-1320110212131132-1321310223320113-2221301200013022-0232133122023101-2121023130202112-0100332311211203"></a>

## Direct properties — disable_allowed_vip_port / 102302313023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210032223320210-3311031131102002-2230301310000113-3313211102020021-3010303321101022-1302300120130033-3032312012033010-0131211022103301"></a>

## Next pages — disable_allowed_vip_port / 102302313023 / 4

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1303303202210103-1301310013032130-3021210312031301-1202003131332112-2320110121001233-0002111223303000-1122010131122001-2122301123330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320131020201132-2200222002211303-3010311020032322-2000332222013012-0022121132233001-0302221330211030-0202223001033320-2020212021232132"></a>

## ingress_gw.allowed_vip_port.use_http_https_port — use_http_https_port / 300202222122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- ingress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-1103011202102001-1301021132122103-2020031132231303-1233200201033330-3100330303123111-1110222202111101-0031211211132013-1202112133210012"></a>

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

<a id="canonical-3123300133311313-2121110131103110-0311301213313220-3300111233113320-3202302103020112-1332212131323031-2000121333310013-0310311211001321"></a>

## Direct properties — use_http_https_port / 300202222122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330132211131001-3103123000201121-3230123200303011-1133213330122132-1201220322002123-0031333100102313-1111323231021132-0211132033113302"></a>

## Next pages — use_http_https_port / 300202222122 / 4

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1211013112322210-1030133330222333-0223020133003031-1222030210032211-1023012220012200-0221212202122132-0230011201030330-1321032000122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201102302000023-3311133123222013-2123223332230020-1223132202000133-0101302002322200-1102220132023100-3122011131332112-0301220233001133"></a>

## ingress_gw.allowed_vip_port.use_http_port — use_http_port / 032122223231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- ingress_gw.allowed_vip_port.use_http_port

<a id="canonical-1323112202301300-0330111132223322-2003132322301300-0111312100312032-3033032033210221-1132303112021101-1330333320312023-0102201003201120"></a>

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

<a id="canonical-0132103131103303-3112132123131121-2100210003222220-1301300230100322-2003300311103001-2222331233011330-2303231020200300-0301021313011300"></a>

## Direct properties — use_http_port / 032122223231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100023203231212-1211310101232102-1112131303332301-3012102000330112-3320232313000303-3232022103330200-3333313211302211-2210300001303203"></a>

## Next pages — use_http_port / 032122223231 / 4

- [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0121301122311222-2132200220113310-3132022103301103-0011121120323013-3233110013232220-3210112123102030-0231103322300301-2003332201002230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2310031033231133-1203001123131103-2003211201033100-1031120301232112-2113211222023330-3220213312101023-1113310303211130-0221132202013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
