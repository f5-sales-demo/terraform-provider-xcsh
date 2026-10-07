---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.outside_static_routes

<a id="canonical-3312331323320232-1123101132211022-0310221210303331-0212202212132310-2310113100003003-2223113312321203-0233211130330023-0202100123021131"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Additional upstream details:

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

<a id="canonical-0320102003212033-2122011200130332-1112132022112321-3022032222211322-1300312201330221-3033332022321122-2323013122011013-3323121221330100"></a>

### Direct properties for `vn_config.outside_static_routes`

- [static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022): complete subsection reference.

<a id="canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- vn_config.outside_static_routes.static_route_list

<a id="canonical-0113233320300210-2031112020313320-2102000211221331-0202033322100012-2002012213332231-2220020103030331-1130002210003113-3132221101321313"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0223013222020033-3103223121100231-0130033122120231-1132132123202223-3230110011122002-2113113102232321-1023332300232330-0313011331211032"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list`

- [custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233): complete subsection reference.

<a id="canonical-0302310312132022-1303002322032333-0013112310120111-2101132212221122-1023123220210213-1023230013002132-1301233203113031-0320233230331120"></a>

<a id="canonical-1101222120002221-2223021113020301-0232312322102201-1310120211131231-1023021203322030-0320012011220021-3232011113232213-1231220000122220"></a>

#### `vn_config.outside_static_routes.static_route_list.simple_static_route` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- vn_config.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2213212111110031-3020033110300133-3222312223101012-0331213222000212-0300223130123303-0201322001212013-3310103201202133-0131103112300101"></a>

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

<a id="canonical-0001310313311100-2022132012032331-1120000303212020-2112310131302212-1302010131110102-2033223213022313-3110101002300033-1012212030220110"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route`

<a id="canonical-0001303303211212-3001111213303202-3222021300102010-1103231103221211-1200001132301233-2113311333220121-3113202133223102-1201220321022233"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-1100111233320030-0031200000101021-2232320020200133-0001223333203233-3012003331023100-1022331221221103-0030323110101011-1003011312301130): complete subsection reference.

- [nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113): complete subsection reference.

- [subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201): complete subsection reference.

<a id="canonical-1100111233320030-0031200000101021-2232320020200133-0001223333203233-3012003331023100-1022331221221103-0030323110101011-1003011312301130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- vn_config.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1200203223130301-3233201020010130-3130331231202133-1313120013020000-0111010303222113-0002313002132032-0120332310110033-0031212223302232"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0303303111000031-3110221030332012-3211302230221131-2023221302013320-3330221133131302-3133122201021323-2221213033133133-1133332323203332"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022010331013230-0221333130101032-2320113033002032-0300222031010331-3023321033103123-3102333300000323-1022300303302033-2130300232101300"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop`

- [interface](data-sources--aws_tgw_site--reference--group-004.md#canonical-0313021002011322-2011123231000323-1300332012312213-0112120213110131-3023133002012130-0332102231120222-2131132231200233-1132223032311203): complete subsection reference.

- [nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100): complete subsection reference.

<a id="canonical-0332120320021321-0120002212111200-0321020213012333-3302102302002202-3003221012323210-2201031020223132-0220021113312022-0212100030313221"></a>

<a id="canonical-1022032011013320-0212112230330023-1102302303011123-3211321101033011-1202030310222220-3231230101103120-0210033000203200-3302333131210313"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` property

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

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

<a id="canonical-0313021002011322-2011123231000323-1300332012312213-0112120213110131-3023133002012130-0332102231120222-2131132231200233-1132223032311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-1202330203013031-1112313331003211-0111303132223300-2011103033201100-1321230321102111-3001021132320211-0310321003311133-1121012212230303"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3023302021223232-0223203101013312-1120003022031312-2121300230100023-2033032003103023-1310132231030213-2220232202322003-2202322131100221"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface`

<a id="canonical-1122300000320232-2111323223330210-1120011012210213-0020230302231310-3010030331012013-3122330020121132-2013312000223123-1301022232000131"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1003010012113301-2211123031012101-3231202212302110-3021121230122211-2123021002011130-3213302303201322-2330031230313100-2001120221101311"></a>

<a id="canonical-2321110102213203-0033331232020322-2211222003031100-2130120013000321-0111133222030213-3132130133302000-2331331323220002-1110300301111312"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1203111333202212-2232210323031203-1320313330032123-3323010300222323-0123233321110132-0311031022201001-0131121113301001-0332122032310220"></a>

<a id="canonical-3210233120032321-2222132322310333-1111200122000332-1202200200100222-1132303202011013-1221023031001201-2330001013333331-1113103323300131"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1311223023122001-3103311303102301-3302212201130312-0200013322330010-2302131321202033-2311011010110121-1123130330312110-3201032022010110"></a>

<a id="canonical-3030103310011103-3332103311133003-0321122032330331-1331131201031023-0303300000123213-1301200223203003-0002020313132112-2121031222310302"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0101101020123331-0130010332033310-1220021033101033-3332212310020202-0031020323200332-3200033331203023-1213130103030301-2131300100302221"></a>

<a id="canonical-3312032303202030-3310001131101312-2000003002201021-1020333023010122-2013330321211300-1022300132202300-0231332011002021-3330123231010230"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2233233232001103-0002203203030122-2123212203102233-3221131213133022-2112300231102300-3012111103001110-2232123230331020-0210031130200321"></a>

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

<a id="canonical-1022130000210300-3132011322013210-1033210003222100-0123312202322210-0213010113120122-0322000122132011-3031200201033303-3220112013201223"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address`

- [dual_stack](data-sources--aws_tgw_site--reference--group-004.md#canonical-0031213211233212-1200002322211301-0212223201031323-1310202332333100-1203222011300033-0310203001211013-2200223003000001-2002033121330022): complete subsection reference.

- [IPv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-2010332031222210-3223102310231020-2200123022012312-1120020113100313-3010112123102232-2232311011001303-1123312203030220-3202302221022313): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-1101202222212220-2113001211011111-2321022021102323-3102300332123213-0301102200200333-1223203210222210-0030121221102312-0212330101210030): complete subsection reference.

<a id="canonical-0031213211233212-1200002322211301-0212223201031323-1310202332333100-1203222011300033-0310203001211013-2200223003000001-2002033121330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0203002321101210-2231123310022232-2110012102133201-3323130013003103-3011010032210321-2021023330233101-2032001112302120-0000001202021130"></a>

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

<a id="canonical-0101011023131102-0213201033121101-1100111012011230-3310321011233122-2002222020322031-0102311333000231-1223313232000100-2011222221211103"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack`

- [IPv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-2311333300313121-0210321102022322-2321231303112312-0001233110112310-3231132213333301-0301131231100313-0233133320022333-1020332020200120): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-0321323123123012-0211000020102120-2121030000123200-1231122222113220-1331011023212233-0310312122121223-3201102011122312-2321312330020101): complete subsection reference.

<a id="canonical-2311333300313121-0210321102022322-2321231303112312-0001233110112310-3231132213333301-0301131231100313-0233133320022333-1020332020200120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-004.md#canonical-0031213211233212-1200002322211301-0212223201031323-1310202332333100-1203222011300033-0310203001211013-2200223003000001-2002033121330022)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0310230212201012-3103130231100011-2211110123223211-2102332201211301-3212232200112221-1001223001222330-2233032210030202-2113222230021202"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

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

<a id="canonical-2213001301231220-1223123323133233-3332221230022110-0031102320332223-0010200020333223-3020300330120011-3303223221230122-2131323120321031"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-2223103022020330-1302110210312023-2323321020023021-2311232000002013-2013011220111333-0301223212330323-3100201201130001-1113212330121333"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0321323123123012-0211000020102120-2121030000123200-1231122222113220-1331011023212233-0310312122121223-3201102011122312-2321312330020101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-004.md#canonical-0031213211233212-1200002322211301-0212223201031323-1310202332333100-1203222011300033-0310203001211013-2200223003000001-2002033121330022)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-1221100310122323-1210112302002031-1123333322000202-1112220132100203-2131202031302033-3122310112023222-1120222031021100-1013333002231122"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1110233331033102-3202323121323313-2203213320312002-0110112332133133-1323320132213022-0212332332322012-1120013121121001-2011330210103211"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-2320321323203023-2312201211021003-2323013110122011-1312321210320121-0002033212011202-2220223201130122-3322033131332023-2021001000133222"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2010332031222210-3223102310231020-2200123022012312-1120020113100313-3010112123102232-2232311011001303-1123312203030220-3202302221022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0010211210103132-1323322113103230-0233202212330001-0223003022112012-3011003032000221-2101201101011112-2323013201133033-0301312200322320"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

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

<a id="canonical-3130002212100133-3111332101300030-0212313233230132-0020012111103100-3333030011122201-2020121002220021-2003031023133233-2102033332103300"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4`

<a id="canonical-1201323120011333-1020302002302301-3100031202203313-1120233012302232-2100010331000323-1102201313010010-2201313130002211-1010001011011123"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1101202222212220-2113001211011111-2321022021102323-3102300332123213-0301102200200333-1223203210222210-0030121221102312-0212330101210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-004.md#canonical-2213332013013302-0333030212111111-2012011103132322-1010223002210333-3023112033332023-1132002021213023-3103223301002133-2223220013323113)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-004.md#canonical-1031011102122230-3032213212302232-1113011013012232-1002300101132223-2011230123231010-3332211100130322-0130223301320331-3200303020311100)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3220110210311121-3022012202112021-3122120111220131-2023200001102101-2311022032322213-3001131023213100-3023213301223013-3101333330023101"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113220200130130-3123332030221021-0110121023312202-1332013312100001-3200022131030233-2333122201222112-2223313223023020-0000222332021001"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6`

<a id="canonical-3201112131210203-2131102211032122-2102133331002220-0003133032133033-0123303112210130-0003002232333332-3101012033332322-3012212320021323"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2130223120300123-2311203113123123-0320021303100201-0312003013031220-3330333012110322-0231002232210011-0230032203201323-3311023030002201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-0211132213222123-3313121302300003-1033302311201033-2101302103331323-1230011112201332-0322112133300301-0331211333132110-2020230101322322"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0101110112331221-2032030322033032-1302110122320213-2200001310000103-3131023000110033-2030321120131310-1133120303331110-3023320320301022"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets`

- [IPv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-1333133010122201-1100233003331212-1200232313020213-0222010221120201-2113130301132022-0233303132311111-3122000223210202-2300011102111003): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-2312131120031202-0100232332000210-1033132122210330-3213300011330210-0132302211101221-3213032023311311-3231001002130233-1102322100133222): complete subsection reference.

<a id="canonical-1333133010122201-1100233003331212-1200232313020213-0222010221120201-2113130301132022-0233303132311111-3122000223210202-2300011102111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
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

<a id="canonical-2230202311111111-3033032130032102-1320110322102201-0302022313312300-1023231212132332-3123320301100030-3011113302302232-0023233133222011"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4`

<a id="canonical-3101030202221301-0310210022010030-2212003233000112-2322313321230121-1030132212033302-2200331333111021-1230333232231230-1110032021103213"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1002300332320201-3203303132320330-0000320100211201-3013220121031313-0220310021002210-2232133200222202-2233220203312033-2232221200121133"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2312131120031202-0100232332000210-1033132122210330-3213300011330210-0132302211101221-3213032023311311-3231001002130233-1102322100133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-3200213222000312-1132321020310201-0133121023310311-2321303121102310-3132320200033201-2233300203321003-1330303232232321-2320302332203022)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-004.md#canonical-3333003103022313-2330001300111033-3331231112112121-2032130030000003-1131022120000222-3132333201231130-1000131130210231-1220133203322233)
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

<a id="canonical-0100320203310103-1221113011313202-1022133032120322-3101011022020213-2220033100212132-2020013331133103-3132012323033323-3103002030030113"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6`

<a id="canonical-1332131310302023-2101212321221302-2311322023101202-1023020000013000-2310303010222232-2320330200002012-1303333201210112-2031022312310303"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120030320113232-2331022221311201-3101202132012012-0132120030113322-3132133012121101-2223323013221311-1223221321101333-2012333111001303"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` property

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3310333131011322-0330100020123222-2030322120221110-2111221333212323-0130210132002022-1102312332302113-2031133120302121-0210233110231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.sm_connection_public_ip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.sm_connection_public_ip

<a id="canonical-1100310330332101-0230332232221303-2122130112022222-1310322111320331-2211022002330202-0322231113130100-0022011100120321-1020200030213102"></a>

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

<a id="canonical-0330213031123210-3033120010312021-0023130021033220-2313022211120033-1111033320230122-0232222032233123-2302100330110023-3000000033133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.sm_connection_pvt_ip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.sm_connection_pvt_ip

<a id="canonical-0003013021032322-2210300002230132-1231002011213030-3122001130312031-3131201223003311-1302201313021331-1223322112022211-1320320301021323"></a>

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

<a id="canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments` properties

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

<a id="canonical-3302022011112110-3123110111333110-3012132222233300-1012330220203101-1300001301230103-1010233301120310-0020233031310133-2233101023131110"></a>

### Direct properties for `vpc_attachments`

- [vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020): complete subsection reference.

<a id="canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments.vpc_list` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0333131211320120-2212200312113002-3333210203330330-3200023231110011-2131331330020320-0001211113231033-1103130123023232-1032000211233023"></a>

### Direct properties for `vpc_attachments.vpc_list`

- [labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-1330000012313312-1331303300303332-2012220031130103-2103102121312023-3031002023201112-3333231211231021-1221010000211212-1123320111301031): complete subsection reference.

<a id="canonical-0003020000210021-2302311213221331-1222202130232200-2313023213123201-2233301033102213-2210001200131021-2033003301033323-1213122200323301"></a>

<a id="canonical-2303013102113311-0122123301031022-2223131222101133-2031330233331231-0022212301313332-1220021332320331-3110110311013011-2033022312311012"></a>

#### `vpc_attachments.vpc_list.vpc_id` property

Type: `"string"`. Computed.

VPC ID. Information about existing VPC.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1330000012313312-1331303300303332-2012220031130103-2103102121312023-3031002023201112-3333231211231021-1221010000211212-1123320111301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments.vpc_list.labels` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- waf_signatures

<a id="canonical-2002121323312202-1100332113202321-0200320322132330-1320113233311320-1103132121103200-3321312301302332-3101302103212110-2021233200122010"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0320110232032220-0221000230010231-3032102022203210-0320123201022231-0111203213000331-3020310203332300-0201210032001113-1300330200222113"></a>

### Direct properties for `waf_signatures`

- [automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-0313312313022101-3320033033012311-1003310022120103-1131230031302300-2030321212200323-1222331233013130-2110130123101032-1022022201032110): complete subsection reference.

- [manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-3233121232010331-3030013302023002-0130221132123031-1311021223032020-2330032230332123-0303313111231231-0322313222311202-3030313211322022): complete subsection reference.

<a id="canonical-0313312313022101-3320033033012311-1003310022120103-1131230031302300-2030321212200323-1222331233013130-2110130123101032-1022022201032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures.automatic` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- waf_signatures.automatic

<a id="canonical-1303233131211010-2113312203131111-1201110023200102-2101320311203333-3311033011212212-2001012312212311-0322322101200002-0303323033101201"></a>

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

<a id="canonical-3233121232010331-3030013302023002-0130221132123031-1311021223032020-2330032230332123-0303313111231231-0322313222311202-3030313211322022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures.manual` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- waf_signatures.manual

<a id="canonical-1202003002301220-3032303103333101-1212030300222032-2330000212022123-0130012232212100-3331210133113330-0202132311220223-0103123211201313"></a>

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
