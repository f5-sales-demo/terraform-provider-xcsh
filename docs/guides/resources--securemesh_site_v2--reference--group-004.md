---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-003.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-003.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-003.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2123322302033321-0111010333321010-0210331221322023-0223232123002333-3003203013231303-3223232033310232-0303330023021222-0121213203032333"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010123103232320-0222202331020121-1000331231312203-2102023103123231-3112122130112010-1002020001230321-0110331233322122-3033010131111232"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-1322210211013111-1223203011222301-1301230200003023-2203320030120320-3232110213000132-3323122031330030-2222211200003313-0203301020230202"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-3020212000003323-1130110030033103-3222002031132132-0211300211003231-1010120033103133-2011323111133323-1303131131112000-3103320320022030"></a>

<a id="canonical-2013213111232330-0223230310013231-2322202010222322-1312300013331003-2331302312022212-1131131002111010-1033002022122222-0302323331023221"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-004.md#canonical-3000223213321230-1022320121030122-1123230031101333-2330302220022310-2220002023033222-3303233133220131-3223232122000010-2013120200333013): complete subsection reference.

<a id="canonical-3000223213321230-1022320121030122-1123230031101333-2330302220022310-2220002023033222-3303233133220131-3223232122000010-2013120200333013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-003.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-003.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-003.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2012320123203311-0302032203311200-3320100122122123-0202320332220111-0132301203311320-1133201022302233-3323203103030130-0023123031033230"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302230211302102-3001321203313001-0111121123312312-2113020023111310-1132021230221333-1333103031133211-2130120113133010-3202131310030201"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0022030110131312-2001233213102322-1030011222223322-2232000210222312-1220010111203002-2110313323133323-0101333223133221-3300001330001110"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-0230210220102313-1033122211021231-0032301003221112-2202332131220323-2212122002221131-1201311333111320-0300230321112333-1220021030031220"></a>

<a id="canonical-1221331022322111-1002203001031332-2301203102331300-0101120003312313-2330333320002312-3102312211000333-1323123020202233-0011000311233122"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-1330001013210123-1120220010020113-3202320223302023-2231201331210030-0101122001201110-3230010032333130-0003121112232331-0031022010210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-003.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-003.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-003.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1220231210331112-1320022003131313-0020113131220221-0222232233113031-3020120112030112-2231000332321203-2003230320012020-2302010023123330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210233113311132-3331113310032322-2213001323312321-3202001113101123-0123001102303321-0321303133122030-0030332101000112-1202113323231031"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-0111120002001103-3312232210112020-2112230032100211-0222032023023310-1023213111102231-0001333313232002-3222203002233033-0031323011201000"></a>

#### `aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

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

<a id="canonical-0202022030000330-3101201300011103-3023131130202300-1321001033331112-0130022201000002-0123211120301111-2013310230200001-0330223213013320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.monitor

<a id="canonical-2103102133300200-3101200031022323-3311311303020231-1021231333133222-0311223003133233-0233211212023120-2232332212331020-1010023022330123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203330022300012-3030233012313303-1203203000030122-0323130220203010-1002312111330030-2202230202112120-0330211133301321-2130311201310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1032233321320230-2011232121222233-2220220012303231-0213322111211332-3022322013322131-2211013122112333-2133111113301321-0200130222302210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.network_option

<a id="canonical-3030030013132130-1202033130300221-0100123011313223-2011300032112311-2121010021221100-2220232212133111-2003023322330303-2330033333210012"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333120001133033-0002311110230320-3111332213210111-3203313102221011-2130213030331021-0121121200303230-0221221011233020-3203030110220312"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2000120230120320-2021221112131120-1103013231233302-2020123001223230-3121032210320012-3220012201020031-3202202202320132-3011032112233213): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2020011030112311-1110111103312323-0002200102322300-0103131130332123-3030120131130222-1033303132301322-2300010210201223-2213013131003221): complete subsection reference.

<a id="canonical-2000120230120320-2021221112131120-1103013231233302-2020123001223230-3121032210320012-3220012201020031-3202202202320132-3011032112233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- aws.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3230010121111030-1131130311330130-1312101311003221-3022333322333203-2110311330231120-0012312033203033-0032330302320113-1203122010202000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020011030112311-1110111103312323-0002200102322300-0103131130332123-3030120131130222-1033303132301322-2300010210201223-2213013131003221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- aws.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1122220221012310-0031201003102222-0122020022012211-2001320300033201-3302321322301310-3020001301221013-1132213121213122-3133322212032131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002311221033103-0201212323100003-2322100030302123-2310132032331013-1132300330330102-3220122223323020-2001023033031213-2211312023003220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-2313232112023130-0012201133313112-3200210311320120-3221010013021030-3101103021100320-3232223232203333-1011322013231130-2013002332230023"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv4_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100230100302110-1330113121032003-3031310032013333-1200022322232311-0333211120100102-3031311232002231-0211032031122023-1011233232110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0120331102320211-0013213122313101-0320322111330130-3302231130320323-2030130111013032-2113210033202030-3331003222301111-0320311103001202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv6_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133113233101313-1032101200201102-0200310132002132-1201332321002212-0100330020311220-0221003203110220-1313130130120313-2230302313120113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3003233121310103-0022222323031111-2032032313201322-3022101223132132-3012101033010003-1000030320312010-3011323323330301-2220102303311323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_to_site_connectivity_interface_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313131032101202-3131121321022321-3302102310201112-1000003301331033-3313123101011223-2201030101012232-0001013321233122-2300312230322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1312012303113112-0033320032312112-1310311203230313-0212010200012111-0332303322222020-2023210123031203-3013201231211131-0100101202231200"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_to_site_connectivity_interface_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231330132121312-3133120010121331-3232322110010102-0123010002112203-0130321130031031-3232300002310113-2330021230332221-1113011321023231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.static_ip

<a id="canonical-1211322112031130-2020031013232111-3111020320330211-0220131332302201-1131311302123111-2023133323122010-0020322331000331-0320003211131312"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232133200302321-2001330230033233-2331230323121023-1013022202302323-0203023113223302-2332313202212122-2123123331211213-1320331300000201"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1321301211011302-1003331311223320-3322012220320200-0202211303113132-2110323201203231-0213310110120012-3023203301233220-3323131330313232"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2022333123210130-2120022111321301-1022031100103300-0230303122202020-2221223011322212-2013132013020310-0310211203303112-0201103030331220"></a>

<a id="canonical-2210121012121323-2210010303310313-1133110201011332-3321221221233012-3302000101101232-1120133021022020-2231212203020100-2132312131322323"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2223132013021332-3201232303030220-2113002021232033-1321213233210121-0002300023312030-3231233323120301-1213300020233023-1102312332110330"></a>

<a id="canonical-0303210302213211-0230331203220021-3120011230031010-0230011010313200-0202131033111021-1113213003201323-2202101103121302-1230320131321022"></a>

#### `aws.not_managed.node_list.interface_list.static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3031322233110102-3112111213002112-2222001121112331-1332011330320131-1201233021011221-3210311303220002-1202213322300210-2011312233303103"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211221222012231-1031211323130022-3111113023220202-2101223223302200-2321322133120133-1301010220322123-3203302200011331-1100311012223013"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-3022220120333102-1313110210000232-1003030211221203-2213022030201202-0130221133131023-2131323010220102-2001223113312331-3013112231301223): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-2033311313310121-0130101223302032-0330200200112233-2200300112213220-1220320032323101-1130132100031330-3211210310210330-2132110220320031): complete subsection reference.

<a id="canonical-3022220120333102-1313110210000232-1003030211221203-2213022030201202-0130221133131023-2131323010220102-2001223113312331-3013112231301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-0301012030120230-2022110233322201-2021001001233203-0301202020030311-1121031023321022-0203023222001223-0222010112032213-3222323313213101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120201121003021-1302311312201202-3113333000233030-1130111331131233-1122100223003030-1331323203320122-2332102103221030-3002001212132311"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-0123203202112210-0130133020301112-2032313220102322-1022011130131230-1310322031230033-2330331233020232-2011313023222200-0322221303230020"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

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

<a id="canonical-2033311313310121-0130101223302032-0330200200112233-2200300112213220-1220320032323101-1130132100031330-3211210310210330-2132110220320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3133032011000031-2032212332331133-0011131100330113-0132111011101013-0222012021300113-1012331132312120-3020320033131200-0110300220211030"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310311003203020-1202010231110310-0300231022121133-1122232132210022-2121320313133130-1023000320031230-3010000201323032-0121230302230100"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3313303102333323-1000033232322130-3221300030222223-0002102230111232-3033312123002321-3232210220002221-2330230013221132-1301003211220213"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3211301103322020-3201312023012333-1012322302012122-1102000322033300-0020131222121022-0010132312113001-2321233223231011-2121133123010123"></a>

<a id="canonical-0223020222303322-3012210320111121-2032131001100212-1010120212112001-2120103012311132-2323130030013003-3322213020002013-1013211202122231"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3231101210212312-0013033123122113-0013023230011032-1332313330302202-1012301210112023-0231000322022301-3113123100111131-2002323011011100"></a>

<a id="canonical-0230011130331201-0300323221022000-2011230220120113-0031112300330123-3202100320301120-2003030231112212-1311310230210111-1100112002121100"></a>

#### `aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021322130033210-3013010303233300-1230322322220200-1030112032003212-0101211221120122-0231232330311100-1012100021331010-1131310012212000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2231121102200133-1231331133233320-1010103110021220-3312002310101212-3302330131200103-0010113033123231-1132233031300123-1211320331311110"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212223323331003-3331021333323233-3102133321130300-0302212233111223-3200123102232120-0333313003002232-3133003200021103-2231111122222202"></a>

### Direct properties for `aws.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-0011111023321321-3020311003023223-1101311113101022-0331303310332231-2013103111222230-1020232032322303-3033022110000030-1132023102230021"></a>

#### `aws.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2033222230111311-1200123010002111-1210013123100133-2011030330230133-3211130330133320-3233112120011110-0022220100030321-3123203021133133"></a>

<a id="canonical-1220130301322111-2232002332021323-3302123230021121-3302312132111101-2102322131101310-1002312320111100-3122321303103230-2021223000131331"></a>

#### `aws.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- Azure

<a id="canonical-2113313310322103-1200312233131310-1321033010323123-2021323211323313-1232210302330301-0320012003322032-2101302210230133-0231203211203233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
azure {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301322312102233-2233003103021331-1111032022012001-1201302320120002-0101313012210022-1001221333110023-0110201203131313-2331132332132302"></a>

### Direct properties for `azure`

- [not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121): complete subsection reference.

<a id="canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- Azure.not_managed

<a id="canonical-3311323323211312-2031030310212312-3303111233210120-2313313221333333-2322131221311100-2200310021203103-0021210031232003-2003320202133302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012301322210222-1320332100110110-3121113023111002-1133103311333103-2321231212112313-2121110012031331-1001300020010012-0100221013111013"></a>

### Direct properties for `azure.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300): complete subsection reference.

<a id="canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- Azure.not_managed.node_list

<a id="canonical-0000201231203121-1020322001201031-3331322333130221-0323030212101200-0220010320211223-0013233002033113-3122123000232003-1000331310301021"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120111333010231-1210202030032113-2330133012203332-3301201132233120-0301120003123320-3013122200223021-1302311332313001-0320003212110332"></a>

### Direct properties for `azure.not_managed.node_list`

<a id="canonical-1102330002203330-1303323223301011-0120230103213132-1210120321022223-1212231303322030-2301231303111023-1011023323232001-1112112012231303"></a>

#### `azure.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132): complete subsection reference.

<a id="canonical-1313330310032012-3222112332330102-0310230031103130-0222321011133333-0213102322033200-3330002033033313-3233233222233021-1312120301212031"></a>

<a id="canonical-3020300223230021-0100201112010123-1131331210231123-1130331230003122-2230212020310020-2113103121311101-1220331031203000-0313022132121200"></a>

#### `azure.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0302011221132222-1012002221310220-2333200010221222-0010103113203222-1122101133332003-3203223302123123-1023220333323220-3303210013332100"></a>

<a id="canonical-0103100232201202-0002231131110111-3330032320010233-2231331021302203-2320300332321011-0303332231331210-1021133212110122-2222222321223323"></a>

#### `azure.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- Azure.not_managed.node_list.interface_list

<a id="canonical-3001212221110111-0123102032113101-0331320323300033-0223200131001201-1213011121000031-3121020031202011-3103111030000110-1110002031101302"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

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

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022103300200033-1313023230002220-3001032012333102-2003312313131330-1220303123310113-2110123113000212-0203230110111101-3332002010232221"></a>

### Direct properties for `azure.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013): complete subsection reference.

<a id="canonical-3231020130113203-1320203303002332-0333212130111220-0001112021032131-0103210223203120-1130312010000031-1031103223330012-3101313013123302"></a>

<a id="canonical-1131312221200010-2233121110223113-3021333221012000-1003011232323102-0000332220322312-1202302100232303-3332301213121003-3001230010323100"></a>

#### `azure.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-004.md#canonical-0233210220123210-3110212232020123-2012020010311011-0203333030021302-2120031221013311-0021001010211031-0111131332203331-3112332000122100): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-0122222112012221-1303023111032023-0023302003202230-1000033332201233-2113112000202201-1120210213111000-3210131222311020-3130312203030233): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131): complete subsection reference.

<a id="canonical-2021123301033100-0332030210231320-2321111330230113-3322323310102032-1213132202121020-3112302231003112-2321221313031322-1002001312322220"></a>

<a id="canonical-0322321220323321-0022001221232101-3100110231101310-1022023013131233-3210301210111221-1200003212200212-1110233101223231-0000010323300130"></a>

#### `azure.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0331103302113330-1223121200033031-1123030301131331-0322332233120100-2110222221103202-2331030223322100-2333111111320222-0120000130201130"></a>

<a id="canonical-0301023330300310-1202013101302221-0030030102323312-1310012110120033-3310310101212222-2301313021230022-0222203213220011-3123102233230312"></a>

#### `azure.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-2301232312232313-1200122320002333-1322032213120200-0001133210003233-3102001123100221-2032122130121012-3213011102002013-3122221000032032"></a>

<a id="canonical-0033312010103130-0032332032331313-1212133030102313-3320020111332021-2102020000223031-2232313133000311-3020000201221232-2012102002121321"></a>

#### `azure.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

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

- [monitor](resources--securemesh_site_v2--reference--group-004.md#canonical-0012332332002200-3102111303112310-1332033032331322-2313132132223022-3223201320211202-0030102022000222-1313012110000230-0311203322223203): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-004.md#canonical-3212332310213213-3300033111113300-0101112023103310-2211012333331012-2302013210330030-0233133220330322-0103220102103300-3321310221201212): complete subsection reference.

<a id="canonical-3023222333001302-2203001322233322-1301100023000111-3021321012010221-3201211220330123-2330230023003303-0001110021033311-2123322332303203"></a>

<a id="canonical-3130020022200021-0230303132302033-1330232033031030-1213001120032012-3200223003133101-0222303301022023-1233221102222311-3112212002300121"></a>

#### `azure.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3002222123311230-2313113211100112-2332231330331233-2300321301231003-3121013122003011-2202213030312032-1202120132200323-2111112212003213"></a>

<a id="canonical-0323110000003200-3320333012112033-0003023123120213-3110330113232231-2122210020332311-1131231201122302-2310331221301321-2223000130031000"></a>

#### `azure.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-004.md#canonical-3301000232322233-3022320013231312-0113322101120131-1120100222101201-1312221213013300-0002303113231311-1330310113122223-0023303333221230): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0012211330030212-0013201021031130-0100233021002233-2232323210033022-0002100312201203-2112113030200010-1012313312221110-0032232310211201): complete subsection reference.

<a id="canonical-2223303000122103-0131012030331000-1211131033300313-1001110031210212-0011023131130210-1123231112322200-2101302310022303-2230021331103212"></a>

<a id="canonical-3013231221101221-0211021022302200-3112311201132233-3231232001003112-1223100312103302-3231130020132010-3313030331010033-0322323210020001"></a>

#### `azure.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1113102132323123-2300332322133033-1032231331233123-3200023012010021-0101001300032310-0232233203013222-3021310023000210-3113003022001010): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1311333023133210-0132231022023003-2101022021332322-1133323223020302-1303011201300233-1302230132320223-3232122321102132-2103323232033311): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-2110310311302232-0030010210011223-3113233222301032-3320311330331330-3033323202220023-2012311330121312-0210322021013113-3133120212032121): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-2211320003103001-3111012233312332-3202121123112303-3201112120030102-2311032010301133-1333023311102331-3301130003133313-2123002312121123): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-0331212221203100-2112013330321301-1013303220131023-0300102331331333-2120332302201122-3020013210210312-0111133211021002-3130221213131123): complete subsection reference.

<a id="canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3101212112023020-3121012313111230-2123130210332220-2020223311300220-3020202001012033-0012220101202113-2230220103322133-2330101100301032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223012111331220-1223220013212323-0100231102030120-2121021312003230-3211003231121001-2203001312311030-0231310131221113-3112203301321310"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-004.md#canonical-1003232300130320-1313033033131113-2333111300321201-2330222311021022-1323130233313000-2133112133002020-3011131200312232-0323010330000301): complete subsection reference.

<a id="canonical-0333200131003123-3121331311202201-1022230301231023-0313313201020122-0331112010131210-0212303323220201-1223213111131003-2113011223210003"></a>

<a id="canonical-3333033230012012-2303303112202032-1032233303031333-3330022302111133-3131203023330201-2220230012021302-1221201131000033-1122233131102033"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [lacp](resources--securemesh_site_v2--reference--group-004.md#canonical-1123013312002020-0313331111332212-0330332333312332-1311121123313131-0330220100212300-3321122031311320-1010302101330222-3203301200222322): complete subsection reference.

<a id="canonical-0003203320212103-1332002012020220-0032213131000310-1210132132233132-3103000133031020-0312201020322032-3231333311020320-3010201210032030"></a>

<a id="canonical-3302310030101021-0131122330232330-3130112211212110-2131321021201320-1311033210211302-2201202202112002-0230103321321311-2131123200223013"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2000232323311110-1010301123220221-2032222223120010-0020220221133330-3123113322332230-1212311101322223-1111113211212023-0102233303021000"></a>

<a id="canonical-2231323301122231-1311202332113223-0000021311120131-0312220022321322-1112211323011302-1210231202030300-3131331303322220-3312331033133303"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320230133003230-0113113322213310-1010233112302231-1022233020100133-2023131312321111-0031202001023030-2203232002223332-3021010110232222"></a>

<a id="canonical-0311103011323011-1230033320330001-3223001123000211-3021231201102033-3312013111333032-3021230012331112-1033132311112302-1221223200310321"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-1003232300130320-1313033033131113-2333111300321201-2330222311021022-1323130233313000-2133112133002020-3011131200312232-0323010330000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- Azure.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1313330020303331-3000103232120130-1200313103122213-2323323330013121-2020220021300120-2032332203012123-3033233312012032-1202131311121002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
active_backup = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123013312002020-0313331111332212-0330332333312332-1311121123313131-0330220100212300-3321122031311320-1010302101330222-3203301200222322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- Azure.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-3220123331000210-3021311020321321-0033222302112220-2010320103132213-0223230121000121-2000010222002010-0322232311002013-1301110323001123"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013111132230101-3331221331212032-3022013121012311-1201131313000023-1332302131022130-1031100220030022-2300331003100303-3132010022213020"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-1221131021313203-2301320332031133-3013223202032301-1232332322321323-2110230201311010-1003121031200132-2000211001230212-1021022130302012"></a>

#### `azure.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0233210220123210-3110212232020123-2012020010311011-0203333030021302-2120031221013311-0021001010211031-0111131332203331-3112332000122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3312123233123023-0320233302331013-1221000332222323-3003021213303100-1003023123131311-3000231120210000-3111330200003320-1302300022210133"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
dhcp_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1332211133002300-3300231303222202-2023032201131321-2013102223023233-2112301320031133-1113101120100002-3220111303302110-1031220210103013"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210302033023302-1010020303101130-3232133210222113-1223321312330123-2101320113201201-0202123221011223-2221232323111233-1012002122232332"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-2233313322033011-0211202300302003-2201100111320103-0111123031303310-2330300033120012-1002331320131300-3021022321331321-1331233100121111): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-2103312320321332-3123311030031210-1222033012111322-3331020310201100-2231222131113301-2032000010323230-1321301330000221-2312030220312131): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301): complete subsection reference.

<a id="canonical-0303321320121331-0032222113332301-2222031200120120-2230311222012011-3330021312001310-0110002100213310-1113033332331001-0330213110312033"></a>

<a id="canonical-0322013223102301-1302113121213121-3120222033020011-0132332313100202-1223311300012212-2010230101032301-3110301002011330-3211112330330302"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2213113310320323-2122221323313022-1133311103302212-1130122323332201-0110001022331122-1202013233003010-1122121013101130-1020012012301111"></a>

<a id="canonical-0310031100222002-3222333210203231-2212123030020310-1203331101330122-0133200111000112-3022323223032130-2013311303220203-2331221001132111"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-004.md#canonical-1220020203222101-3221122312100013-2111102322111013-0132220221333302-0010000311221302-0303001000030231-3332010302330313-2100010101223221): complete subsection reference.

<a id="canonical-2233313322033011-0211202300302003-2201100111320103-0111123031303310-2330300033120012-1002331320131300-3021022321331321-1331233100121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3001133331122323-2002221223201201-0032222121323010-1233333002303123-1020201113032312-2112110120012213-3230210131122012-2221013010311122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103312320321332-3123311030031210-1222033012111322-3331020310201100-2231222131113301-2032000010323230-1321301330000221-2312030220312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3010331130300002-0220130213203222-0011212120110132-3012313123211213-0310012202001000-3112211011131300-1223020333200133-3203201221000010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3010332202033300-2213203233231012-2331321000202000-0012033031213120-1311221210200003-3202300212023113-3001123010120032-3133132332120333"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200320233112232-0123320100202130-1103112221322221-2200002313220202-3211030102232102-3230110121220210-1021301200300232-3120131333022113"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2003201232100110-0120102221312112-1322031031233213-3213022021023230-0320100222111202-2033312210330211-1202100221123120-3200312313120003"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-1201130332003213-2232032202303111-1023013123321103-0133020033321012-1020330000130213-2022330131010123-0233233300333301-0211133103121022"></a>

<a id="canonical-1300222202023103-0320023231000323-1032333122112223-0003113210301232-2331303301212223-3101213133010021-3010213012123202-1200001202230331"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0013330021001122-3130130213123100-0100210112200213-1123310301110002-0202331111322230-3213313200301320-2230102202312102-0122200000112303): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0212132301200321-1023021333033300-2220131210220133-0023312212200021-3302122202001011-2331021321310302-2300222233002131-0322201023100223): complete subsection reference.

<a id="canonical-0222321023001312-0300203200332323-2111011120033301-0321221132330220-3302311112200311-0030200011030211-2122133313300002-2122322001131331"></a>

<a id="canonical-0132103001103221-1200222200123002-3001232303322221-0021330131212310-2011132022131021-2121330200321133-2110320202322003-2222112331100231"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-2200010010230103-1330021001030003-2130232302303221-2020032123121023-3210102112003321-1000322210123321-0332010002121023-0203011011300120"></a>

<a id="canonical-1323232333300010-0333213013012300-0301302001022103-1003121023332120-0011011023201330-3221230000313112-0110032202300202-2330203313210312"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-004.md#canonical-0210122333231303-0102210303200330-3221211103031100-2221122100301320-0320231133200130-0032221223133120-1321322321033330-3302012313202031): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-004.md#canonical-3100133112120033-1033112212330322-3212302013103101-1011210331123200-2210313030112322-2112002203311231-2221212222031322-2333131221033331): complete subsection reference.

<a id="canonical-0013330021001122-3130130213123100-0100210112200213-1123310301110002-0202331111322230-3213313200301320-2230102202312102-0122200000112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0301112133213120-0102100113013022-0030112020100010-0120101221010021-0113201103223131-2031022311200131-2323203311310123-3211322011101200"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212132301200321-1023021333033300-2220131210220133-0023312212200021-3302122202001011-2331021321310302-2300222233002131-0322201023100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2321112110212003-3211203102000233-2201203023310311-0222310033320123-2021122103310200-3000222201322212-0023132102313210-1001103132013131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210122333231303-0102210303200330-3221211103031100-2221122100301320-0320231133200130-0032221223133120-1321322321033330-3302012313202031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3030010332112021-2031130232112110-0002112211320131-2111013320301220-1013123111200101-3103233220302230-2200303113100110-0132223211103000"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131112112200330-2311230113131333-1202332122211002-0112322310333331-2332200030113310-0112222010223110-1132321210201032-2332020030231322"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3033223033323210-1101332223322031-1313121101213220-0302322222020321-0133230022002130-0003300022001212-3230020332310301-1310123222312203"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-2331011201020313-2221322010010000-1003033033031020-3323133130010320-3230032023133220-3331110223303000-0033312213013302-3023011112003100"></a>

<a id="canonical-3003130132210321-3231220021313223-0001232233120322-3103101131100131-2110230310001001-1101001332001213-3023310011201002-2110000131002210"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1001200222133201-2301211313231113-2202333202213233-0023213212023310-0332210300100313-0321223133032223-1331213122300231-1202221202000233"></a>

<a id="canonical-0030022120232002-2332101112223312-2303322323200233-3231232130100320-1103221310333313-3322013031232310-0333102211301322-3332210231303030"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-3100133112120033-1033112212330322-3212302013103101-1011210331123200-2210313030112322-2112002203311231-2221212222031322-2333131221033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2120230312323113-0322100220323330-0222230321023133-2301100001212111-3213110113002111-1113131330212233-2013220031222323-3220213222001032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220020203222101-3221122312100013-2111102322111013-0132220221333302-0010000311221302-0303001000030231-3332010302330313-2100010101223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1211032332101030-1123321221111332-1032203320322232-3321312211200031-1013211101313033-2231312210221201-0130001132131021-1210012020312023"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233103311310213-0003103330033323-1212332331132312-0332321001020312-0332100023322031-2101110101110221-1213011111310323-2300213121121021"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-3023030130223300-2331022330222301-3311133030212233-1112231330203121-1202332022030212-2331021233131130-1212121301111010-0320010320012021"></a>

#### `azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

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

<a id="canonical-0122222112012221-1303023111032023-0023302003202230-1000033332201233-2113112000202201-1120210213111000-3210131222311020-3130312203030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0221022222030212-2223123110330020-1120300211301001-1320001013213123-3212111230210222-1300231300101232-0120320111231330-1333300031022203"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320233020003223-3232302223101013-1000301233122221-2031201002312310-1111232201201302-3122320200201310-1012022111113300-2120003213130232"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-0313010111000330-2032112123112222-3122232031213100-0130310100030023-1000303032323233-3010230310112123-2103210103330000-3121203033033301"></a>

#### `azure.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0210000212312303-1111312100001103-3112020302331021-3223222110122203-3201010111333303-0102102112301303-2010212221322013-3101032332313210"></a>

<a id="canonical-2013103220123310-3103300311212033-0331033212223232-2022033033000031-2131122320012010-0131203311212312-0322221121031202-3200211311310211"></a>

#### `azure.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-2200123231313031-2032203100032232-1110301220212330-1200120223031013-1120001102201301-3010102323002220-2332103323100000-2322030210222230"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103233102023200-2002323233010120-3330102020113102-3022130103120222-2031111301323023-2223011212323101-2030010101133301-0212033000322302"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-004.md#canonical-2202033320320023-2210222221233221-2132101132020312-2122020112231320-1001103321010213-1002011323221230-3002210301321031-0232102121120302): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021): complete subsection reference.

<a id="canonical-2202033320320023-2210222221233221-2132101132020312-2122020112231320-1001103321010213-1002011323221230-3002210301321031-0232102121120302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1020022031220113-0320302333212130-2313021031033110-3313202003312101-1012232112002100-1313330332222031-1023222121101220-1133333001331032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
host = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-3032311301312032-3232021020112211-0233320213203300-0011231133230203-2320330200330021-1102123011210010-3321203033102122-3211322002322200"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013031001303302-3313312322320001-2112133320330211-3310021111021313-3201103312122211-3232113122231303-2323100231113101-3331210201003122"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131): complete subsection reference.

<a id="canonical-1210211100112221-3322212011101112-2330300022122031-1123101031311232-1001222322101310-3113132333012303-1222232202023232-1101332212303313"></a>

<a id="canonical-1010120321103300-0331020320223202-0103303101113313-0231010203033111-0001330210212012-3201010212111131-1100323213121221-2012112223000120"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100): complete subsection reference.

<a id="canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-0110223032102113-2323023213113133-2201002320022223-0103131210103020-0200113002130122-0023120003232003-1121333323020313-2012320232230131"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023130231312002-2323022300022222-3030233212322120-2221131031223012-3332300130002320-0133101311232212-0011100323121022-0331221122211100"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2022331003023121-2110013031123330-2200210022302202-2303033132223003-1323322100230303-0112320321322321-1033120003023011-0110211011311033): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-3333220210333330-2210031200322032-1333010121101222-3322120023020133-0233133030232231-2213331213011213-2202311222331222-3132222121023001): complete subsection reference.

<a id="canonical-2022331003023121-2110013031123330-2200210022302202-2303033132223003-1323322100230303-0112320321322321-1033120003023011-0110211011311033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2232233331213320-2310200300322102-1020030212113123-2230300310321330-0032111103221000-1211122322022022-1330021330300011-1111233123120003"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132332011101130-2223120012032322-2201023202113030-1012010232102301-2311032203023331-2003212231222010-1011212312133103-3330103010100210"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1232122122013003-1112102202211130-3313121213030213-0031120231230321-0203110320302201-1101103330121213-2310312132112311-3032302322013210"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3333220210333330-2210031200322032-1333010121101222-3322120023020133-0233133030232231-2213331213011213-2202311222331222-3132222121023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1131010210232213-3330110310320310-1133100321333001-0113001322232313-3012231221113313-2002333212312230-3230131013223333-3220321221223203"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132211011331120-3022301101222311-2103000021322221-3312330233121012-0022021222031102-3012223221323012-3312203100212021-1332103122120201"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-0200103320232321-0023202223130030-3300102012120330-1000112223123013-2232200112032023-1021123323232002-2013123122133322-0101313300303133"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-2302202211311222-3203311130221333-2333002113111313-1123300111030001-3102333102130213-2212011301300321-0230320030112203-3221113321032203): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-004.md#canonical-2030012323210020-0000201200230223-2230101120123100-0013301300200102-1011231032230333-1110103130113330-0010230310021002-3302211221111000): complete subsection reference.

<a id="canonical-2302202211311222-3203311130221333-2333002113111313-1123300111030001-3102333102130213-2212011301300321-0230320030112203-3221113321032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-3333220210333330-2210031200322032-1333010121101222-3322120023020133-0233133030232231-2213331213011213-2202311222331222-3132222121023001)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2211330010211232-3201022101310003-2210321200220123-0221331002312130-3113001132031300-1003100230023013-0130211312031000-0002312233233201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030012323210020-0000201200230223-2230101120123100-0013301300200102-1011231032230333-1110103130113330-0010230310021002-3302211221111000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1123120021210000-3213110123231000-2120022100112113-0001132312000123-0302303113303101-1031003202023001-2100003023202133-0210100202212131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-3333220210333330-2210031200322032-1333010121101222-3322120023020133-0233133030232231-2213331213011213-2202311222331222-3132222121023001)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-3321031230012212-2212131132230133-3200021231032332-3233213301001220-2320200202323100-0133013031013020-0203103013023223-1332011312210213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2013032312330203-3322010331313323-1111130312012001-3231033122211130-1203201302122311-3223301311310130-2222322321330022-2101101203312012"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121301322013330-0001031222212312-1111003023323233-3102212012301020-3313203201300321-2222310022121313-1010313133022233-0102221300011032"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-0030232013310123-2121302010322131-1303002100323330-2232310230320120-2013102231222201-0133232331222012-3131121021023010-2111232031302213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-0101111211030213-0232210123112322-3231301230331310-3023112001033232-3030133232302213-2002001303122010-3231333203330122-3132010300120322): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132002321111010-1313232123232332-1320000132300222-3131303120330201-3332331201222223-0302121220123321-2001130302330111-3333223113100233): complete subsection reference.

<a id="canonical-3101203203323113-3102000221020001-0200003300233113-3312312112212132-1132233221230233-0202033133020003-0301120331022000-2013123022020103"></a>

<a id="canonical-3101003113301322-3002000301033113-3321203132030311-3123312032121333-0223221120300231-3120021000033113-1130213103202133-3320032203033330"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-004.md#canonical-0332003113133122-2211000003103311-2010100023022023-1231223200230020-0203001132123022-1101323211131120-0122212311300200-2001002012030322): complete subsection reference.

<a id="canonical-0030232013310123-2121302010322131-1303002100323330-2232310230320120-2013102231222201-0133232331222012-3131121021023010-2111232031302213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2302312112201230-0301101002222022-3000011123033102-2113023221310001-3133213331112101-1111231300100110-3232101112220131-2200010301230331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101111211030213-0232210123112322-3231301230331310-3023112001033232-3030133232302213-2002001303122010-3231333203330122-3132010300120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0311113001311113-3231331222323313-2220310011111222-0231110323201013-2222031213000120-2211121300223111-0122213033030033-2132223213203221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132002321111010-1313232123232332-1320000132300222-3131303120330201-3332331201222223-0302121220123321-2001130302330111-3333223113100233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1132223133100021-1223211123233301-0330333230232222-1303211102310313-0231023110211011-3020332111013220-0121020302323012-1212310310323331"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330301221221112-2002320102030332-2212320303103111-2033033001111033-2033013002020312-1010033310310001-0333312321011221-0222202103331212"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0302000001132023-2030301210032101-1232012113230231-3223201213301102-0001333023021111-3013120032302112-2023132312313122-0021312013200000"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-1231211233312333-1302300003021320-1010101022231103-3322010330110310-2100030020022311-0321111202321102-2203222332020320-0031310030010333"></a>

<a id="canonical-2223001121222001-0321120201103213-3232112302123232-2332202303132223-3113122231013222-1132233220131302-2012100013103010-1110310211032130"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-004.md#canonical-1112322320221212-0311113332131102-3012012111020211-1122110323100120-0222103031100002-0320001123212320-1312012012213130-1212012000012302): complete subsection reference.

<a id="canonical-1112322320221212-0311113332131102-3012012111020211-1122110323100120-0222103031100002-0320001123212320-1312012012213130-1212012000012302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132002321111010-1313232123232332-1320000132300222-3131303120330201-3332331201222223-0302121220123321-2001130302330111-3333223113100233)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1111310131323230-1222113022313230-2301131302033333-3320200320010203-0223023210000022-1210221032012123-3111201233001023-0211001332111003"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221032021120222-1102312132321310-2311022301110020-1010033322310012-3312110323001032-3100200232200320-2002202322131232-3133230323102220"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0100103122212230-2213112220011222-3231213231031002-1301013133120011-3223130130002133-1330130013112320-0301223300033123-0223313102120333"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-2133100302130021-1333333223122313-3010121201211121-2330010133103310-1201230300211130-1001031010213202-3002322213301123-0110022133320303"></a>

<a id="canonical-2122320123031011-0032121122301120-3312233011110303-0102113010013211-1001302010222222-1200302200320030-2112021001113202-0021233113023310"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-0332003113133122-2211000003103311-2010100023022023-1231223200230020-0203001132123022-1101323211131120-0122212311300200-2001002012030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-0033213033000302-0103232303103213-0032220101203332-1003011301031203-1201023301333301-2302331310331321-2302000233302132-3213212220011100)
- Azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1223102301223212-2011301030000322-1233021231113101-1213113021213233-3011222130200133-2101201330133110-0021033021102003-3110102313132003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032222331132323-0120221120223200-2201122301020111-0000201321012010-3012322203112022-0331322301333030-0131121301213312-3232010313210003"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-3330130111000121-1333010220212313-2210231120031321-2123231000230013-3100313130310032-0322211331020101-2002311313322223-3320220212231000"></a>

#### `azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

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

<a id="canonical-0012332332002200-3102111303112310-1332033032331322-2313132132223022-3223201320211202-0030102022000222-1313012110000230-0311203322223203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.monitor

<a id="canonical-3210223021231130-1330302033010101-3311110123302200-2330231303012222-2323331012020000-2230032213213301-3332110132333303-3213323212311113"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212332310213213-3300033111113300-0101112023103310-2211012333331012-2302013210330030-0233133220330322-0103220102103300-3321310221201212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-2231123011120120-3300112232030010-3103331312113132-3133320001333202-0232321133032112-2122323320013210-3101300220011323-3001200222022131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.network_option

<a id="canonical-2221121020013103-0100031230012210-2233132101230333-3221021102213021-1211103203230323-0213101133201210-2223120322121010-1303221101110300"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323313301231233-3002113102000000-1230312011010311-3210200103220220-0331222122022210-0032300201123322-2020103201010011-0203332123100121"></a>

### Direct properties for `azure.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2013002113031011-2233000033301131-0011010132010031-3023011310121010-1333120230102330-1222302201210102-0110311331130303-1222030010031122): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-004.md#canonical-0130313122102203-0110120310130103-2133211112313300-0031002232013131-0031303112102122-0103201100200330-2221132213223310-2031211303233003): complete subsection reference.

<a id="canonical-2013002113031011-2233000033301131-0011010132010031-3023011310121010-1333120230102330-1222302201210102-0110311331130303-1222030010031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000)
- Azure.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3221033313131323-3221011032311113-1302311001032232-3113032001102300-3321002103123100-3212301220102003-2203332022000200-2120030013201022"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130313122102203-0110120310130103-2133211112313300-0031002232013131-0031303112102122-0103201100200330-2221132213223310-2031211303233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000)
- Azure.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3023033312022122-0323120303120131-0011200100121101-1121302313020012-2020330300131320-1333102211303022-0120330323313021-2103031103313131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301000232322233-3022320013231312-0113322101120131-1120100222101201-1312221213013300-0002303113231311-1330310113122223-0023303333221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0323231131011110-0230010220132203-2333302321302031-2122110303301033-1103320210121321-2302133313301003-2223133333113221-3000113113332333"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv4_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012211330030212-0013201021031130-0100233021002233-2232323210033022-0002100312201203-2112113030200010-1012313312221110-0032232310211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1122233013233201-3111203320211313-2203322230012301-0032321010202032-3212233332121112-0012112112002300-2221221301101302-2231322332001311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv6_address = {}
```

This is an empty object or choice marker. It has no direct properties.
