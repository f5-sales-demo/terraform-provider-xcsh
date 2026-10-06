---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2332221212012321-0300321131323202-2132121330200312-2310202333131020-1221111303021213-2310113311020303-3110130101103103-0111210300303211"></a>

## Direct properties for `gcp.not_managed.node_list`

<a id="canonical-3102111110111312-3200211333101330-1221200121330111-0012313300221030-0222211211112132-3310333332111233-2221013212021330-1211313031200202"></a>

### `gcp.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331): complete subsection reference.

<a id="canonical-0302123330000010-0311331200033000-3322031101321200-1220311103333213-0110011322031032-3032233313212122-0223021120102333-3303000133330203"></a>

<a id="canonical-1033313332103130-1333222031210200-0223233033110113-1020222200020131-2122202221202310-3211001032301132-0302111322102111-2021011100333010"></a>

### `gcp.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1030021030033032-2102223202321131-0113333102333213-0201101330103102-1032120323010110-2223333202313000-0113032210332330-2313010203211322"></a>

<a id="canonical-1201021032003222-0201023122301132-1331232330300220-3232233320212323-0023330333223100-3110231333332221-3330303201111230-3001201303330330"></a>

### `gcp.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- gcp.not_managed.node_list.interface_list

<a id="canonical-1323313113202203-1202113333111123-0231133222302133-3310133323302022-0231211122200001-2310310110303000-2031331300213221-3100310032003030"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2310120033012330-1213220200132321-0223333232310100-2102221122333310-1330332221032112-0033322233132120-2303120010301231-2222100200120103"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002): complete subsection reference.

<a id="canonical-3312330221121222-3200302132110202-1231001301112003-1313112103113113-0001231312332312-0133012000203123-0131110302111113-1101123302131331"></a>

<a id="canonical-0302001212221303-1013311313311320-0130131103002021-2031032130331002-1001301313012130-2322113001300302-2131031111112310-3131112303000021"></a>

#### `gcp.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-0033122033012121-3302331120212023-3323021011132313-1123200320201001-0033122103112313-3030132200003013-2110330032333033-1232013121003210): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-2211332222002230-0122303122321113-3201122320313333-3310131113000011-1101321120213101-1203213233003121-2033310231011313-1012230033113111): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301): complete subsection reference.

<a id="canonical-2320031232313202-2010022000323211-1331220032100130-1312012203213033-2223233122000313-0232101000232323-3120120013222012-3303310031131033"></a>

<a id="canonical-0021212131201201-2000231203113321-1333003013332311-1210212022112101-2122233230230033-3013102301321231-1000130001200230-3030203231311103"></a>

#### `gcp.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3323101032132122-2331001133323211-2223032300102223-2201102211302121-2212010210212023-0320100231321320-2311221000121331-0203323201230200"></a>

<a id="canonical-3221110130033203-0230300020223111-1222111231011223-3232111110313013-1032200321003133-1230130303010021-1322010300212230-2330123202210003"></a>

#### `gcp.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-0133001220231023-3110021333021130-2223330321003100-0000300021222122-0203233321311321-2103232002123203-2321020232132221-1203132221120031"></a>

<a id="canonical-3130031123311113-3032121300001311-1303221011121011-1023013111021103-0300021013121113-1223023122300302-1213313333213311-0313322021012233"></a>

#### `gcp.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
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

- [monitor](resources--securemesh_site_v2--reference--group-009.md#canonical-1211223311200301-2113213233010230-1203012301302120-1022300331103323-0211120332111200-1030031000001303-2222221212332011-1213312211122110): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-2110013100330201-0203313102212200-0032133203312013-0101212301101020-2313122133233302-2220111322101030-0222330312032021-2111000113010023): complete subsection reference.

<a id="canonical-3311222223203222-2022021033323201-2221200213101330-0112000022011110-2102220002322300-2310130021120003-0010301021033323-0311021022202323"></a>

<a id="canonical-0333320003001103-1132311212312013-0213130103021130-2133311123100102-3320130321122320-0310132022130122-3013232221013210-1031221201100120"></a>

#### `gcp.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3113333132213021-0012022322311223-0233323303300101-1103211002321021-3232202220200301-2300031223020133-3312320212121010-0001113112312233"></a>

<a id="canonical-0311210320203101-0020223002011212-1122331002300101-0001132130332121-1111322130123110-1101021301000033-2133223020323200-0312021222102001"></a>

#### `gcp.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1203323220122012-2201322220300230-0230002310100030-3330031323321311-0300313103011131-2300131311231110-3013203313301020-2320101230133033): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-3000103023231312-2301010323330333-0333003302000002-1322000323123232-2002220330300031-1331112001201102-2321333303303203-2001023303022112): complete subsection reference.

<a id="canonical-0121230123100333-2233032212020112-3321110230100023-0123013202322310-0330000203301310-3122020110200211-3222201023003220-0230002011111231"></a>

<a id="canonical-3132233231201203-3220013322113230-2312213301222001-3333020232001303-0101032103100003-3333110031333311-3302222030111033-1233321130311232"></a>

#### `gcp.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-1133201233031230-3333012130033031-0221332312123131-3003211133222313-1013122333012331-0100231220131122-1212222211103002-0131010321033230): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-009.md#canonical-0220100020332132-3221320131203333-1233013032031323-3222113210012223-1121032113121201-1001013201031031-1210302012100032-0323200312201112): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-1030120100013303-3022221311322001-3212222021021102-1131222030210221-2220333320001121-0123333131111212-3121032230202313-3133313121120333): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-2322033023330322-3002023133132222-1111333320023023-1223323000130112-0000221001032331-3310323132320101-1112120210112123-1112101101200111): complete subsection reference.

<a id="canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2001222321030122-3333010021131112-2330230331112230-1130012211032103-2301220012233021-2131333320132300-2132203310033322-1330033120021221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2010200312032123-3221323033010132-1100001201310211-0030011200301222-2113013202230202-2003330033222030-1222003101211330-1220120310110203"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-2233223120230031-2320103303230112-0131121110300032-1333112312010101-2233103303000103-0332322300202330-3313021111031333-0330011212131203): complete subsection reference.

<a id="canonical-3002201223003111-0011203100210201-0213333332111101-1331210331200321-1112230002122032-2011112013320310-3223211300022200-0312231013330132"></a>

<a id="canonical-3200121301010101-1300321323012030-0231103001222323-3003203013013210-1230200013303112-2032002031120110-1320203120321210-0321003332133310"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330): complete subsection reference.

<a id="canonical-3332101220301211-3201332233031312-2322332232110131-3222321130131302-3131102121131313-1223211123110333-0211311010200310-2122122003013321"></a>

<a id="canonical-0313210321131001-1322223220321222-0103312030200331-0112203030223210-2120020231033232-1320313220112311-1120311332033313-2010000230123010"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1021322000122023-0221033322000222-0322330123012200-1132022121331212-0120303032103123-3130232132300301-1010100110010000-3000100011100012"></a>

<a id="canonical-0120312302100131-3310011111122033-2020000012001013-1122201331301013-1003021020321000-1323212023312312-0330112121030013-3301330100011330"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0020011012022332-1130333220033023-2213322122032200-1102300033223222-1010201223301222-3210003101313032-0021000000233222-1010132130013303"></a>

<a id="canonical-3132323202200332-2123030003101303-1002310320000222-1221302321323210-3122220112310001-1103320112101013-3132301131301023-3100311022233013"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2233223120230031-2320103303230112-0131121110300032-1333112312010101-2233103303000103-0332322300202330-3313021111031333-0330011212131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- gcp.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-0103310101200112-3103110103223223-0232211211023323-3111333131210333-2113010031111133-3333202101022300-2211133112311230-2111033212030021"></a>

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

<a id="canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1031110200131023-2031130001331322-2113313211122031-3213222332211230-3210100211310222-1313033111331131-1000333031323131-3322122320202200"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1221113331023113-1103202013023013-2121200002013132-3311013032123113-1320022213010232-0233102120330031-3213100203002232-3102201201300210"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-3120000010132321-3322331103231003-3311300130100031-3113311202023321-2223132300203110-0221002021230313-1020200102210303-2333003123121003"></a>

#### `gcp.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0033122033012121-3302331120212023-3323021011132313-1123200320201001-0033122103112313-3030132200003013-2110330032333033-1232013121003210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3213131300211013-0031022002201033-1020323021211022-1002333123032201-0103102000312110-1123200113300020-1301111130120322-3001030031121230"></a>

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

<a id="canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1202231222103132-2213223331320022-1111320230100300-0333302030203330-3230013310203011-1213333113103222-3220011132000132-1211200313212132"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0221203221012201-2210212122110011-2303012002321201-0030120230310030-1232121332311231-3013131201013102-3203230013220310-1232320001203113"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-3230201023201000-0031203031300322-0130003111203121-0221031120033200-2320100230310111-3203202120223003-2200122002201120-1130102300303323): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-3021112221211102-1120023232212220-0102020133000323-3312003300330210-2120230223212123-3221201200132033-1322200123220301-3212121021112120): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132): complete subsection reference.

<a id="canonical-0021201321301330-1022203301131303-2003200120030122-3003313111030333-2031130202021313-2032003323023021-1121200030201031-0023002131102003"></a>

<a id="canonical-3332333210330110-3100021002123201-3122022020131133-3220333322203331-0323003313312223-1022002103133231-2011323023222123-2223112230203021"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0021030113311002-1331300302000301-3010330030023303-0332110313233021-2222321313023121-2101201103030113-3002230231311030-3203120301123000"></a>

<a id="canonical-3110320003120102-0302310101123330-1313321021312302-3021032030322012-1301103211230123-0131131021100100-3022000133100231-2022111322203313"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-2111220011212201-2210211320212123-0002230223223032-0103310130311120-3030102221113221-3203102202021320-2333322310332131-2200113023033310): complete subsection reference.

<a id="canonical-3230201023201000-0031203031300322-0130003111203121-0221031120033200-2320100230310111-3203202120223003-2200122002201120-1130102300303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0320320020011200-0310321333330223-2113000331220213-1033130131230110-3322330200013331-0001330333320133-0120211320223013-1100300002021323"></a>

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

<a id="canonical-3021112221211102-1120023232212220-0102020133000323-3312003300330210-2120230223212123-3221201200132033-1322200123220301-3212121021112120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-2222122000020320-2213023112023330-0332021300111131-0102100202221302-2133112023132003-2220201111133312-0122001310212332-0200122300031321"></a>

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

<a id="canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2103120333232032-0222310333022131-1232130203013230-3021323300303222-2012010001203230-2300010233120202-0222123312120020-2222303200012102"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1211212103321322-3330023320011021-1101020101222120-3023111320311122-1022311213120331-0112111021123231-2233303012302232-2000212122030332"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2323103330202101-3313320003331123-1320131023002211-0302033213311332-3133201122013021-2021113110310033-0110020230310233-1300322203201130"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-0320232001133133-0211000120101202-3230113200010213-3031332203032303-3000222222030022-2033001100130012-2213120100020100-1233033313002112"></a>

<a id="canonical-2012311123132131-1301121030102312-1033332301232133-0230120130330301-1320322120100212-1202133111213032-0300213333011100-2231010101301133"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-0312003331312110-3111001100022001-2222300312023132-2122133211103120-0020032122130203-0112333013130301-2230113023320120-1221112300121112): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1311222001020112-2112133132212013-0031233202312313-2013003011202013-0320120020020030-2030111033010310-1202132011310200-1222312310232123): complete subsection reference.

<a id="canonical-1210102102322200-2111122012322032-0122020032122000-0210102023122211-2210012132213303-0201312311010302-3001233123132321-3022200321331011"></a>

<a id="canonical-0132220212320222-3003132020210122-0213100321220200-2303320003111202-0232320300311013-2232033011300223-2133302030202103-1310303102312331"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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

<a id="canonical-0201302303301202-0013021300131023-3022300123320113-2101112122011000-1221300003212110-1313231020202032-3003200002223203-2230310100233100"></a>

<a id="canonical-2302011302120220-1231323320012133-2031013123100300-0213030223311212-2223021302303322-3133011223132132-0300001011001021-3312200013022323"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
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

- [pools](resources--securemesh_site_v2--reference--group-009.md#canonical-1230113202102212-1123120311220322-1031123313022131-0112023321111331-3301330310322223-0033132130330122-0331013200321001-1202011312012121): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-009.md#canonical-3131011201201112-3113332132333013-2033231312003100-1121111323123320-3120202232101232-1003132033323032-0332020332132112-3010301022003223): complete subsection reference.

<a id="canonical-0312003331312110-3111001100022001-2222300312023132-2122133211103120-0020032122130203-0112333013130301-2230113023320120-1221112300121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2311222320322023-0121203331023323-3130232331301011-1003332131221212-1120223203003301-1211330133011200-2000112110313311-1322002310313212"></a>

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

<a id="canonical-1311222001020112-2112133132212013-0031233202312313-2013003011202013-0320120020020030-2030111033010310-1202132011310200-1222312310232123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2311120210113303-0321303213222131-3311111222032301-0303023101333300-2333213102130223-0133312031023023-3112312033321113-3303110021130312"></a>

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

<a id="canonical-1230113202102212-1123120311220322-1031123313022131-0112023321111331-3301330310322223-0033132130330122-0331013200321001-1202011312012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-1032321201211003-2201222210101202-2122211013112232-1101030021233021-3100331001003113-2001023220332302-3022102230103011-2013122231333031"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2302111013221322-1101000000111230-2032221033321233-3111012000313011-2300221323322122-0311223313323110-3012323212101331-2100000322022302"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1330220301202033-3200322233010123-0313303112000321-3312233212222230-2012130130122203-1300322123302311-3123023013311011-2000300230100310"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-3013303132022201-0231023021232122-3103002000323123-3012110132201310-1123313310132101-0121001120111223-3130130231210320-2313023132222311"></a>

<a id="canonical-1203302333211211-2322221101102311-3321221220220203-1133212102312023-1030211123023333-3320102311330323-1202223031202300-3230003303302001"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1010220111022002-1110303130310012-0330322012310010-2001203110002031-0110101022202132-2210200333220100-3030020123320102-1213032030212321"></a>

<a id="canonical-1320110333013321-0221013000032211-0111303332131012-3300011002213032-1332322211133033-3223201003301130-3121220103231302-3200200112111023"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-3131011201201112-3113332132333013-2033231312003100-1121111323123320-3120202232101232-1003132033323032-0332020332132112-3010301022003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2311200333231211-2221202112320303-0313300221013032-2303332302201122-0322130113232031-3010232121130332-1130032033032233-3321113312312122"></a>

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

<a id="canonical-2111220011212201-2210211320212123-0002230223223032-0103310130311120-3030102221113221-3203102202021320-2333322310332131-2200113023033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1313302310211211-0113000002100212-1230302020211203-1330131333101033-1202032031212013-1220122010102001-1322323201222003-3212101223201010"></a>

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

<a id="canonical-3231313122013121-2311023230023202-3030203231102013-2232010210223031-2122202021201100-2101203112321322-3213221000101002-2012223322330032"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-0030300301030203-1120100213002013-3032021333122330-1311013303013220-1201322231023021-3011301320320303-2301333013302011-1131113201103001"></a>

#### `gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2211332222002230-0122303122321113-3201122320313333-3310131113000011-1101321120213101-1203213233003121-2033310231011313-1012230033113111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0023133312311013-1020102120132013-2031120023102022-1211021223300330-1023121303100112-2223203122111221-2102322312220112-2111001012331101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0203021110000210-3013032221122322-1302013030001012-2000001223120232-3312021321131301-1102021033101102-3212030312313021-3120113013313132"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-2202131113031322-2202201233333312-2100212212322300-0120200110203310-0210133313210132-2120111033323012-1232132323300323-3332212003132231"></a>

#### `gcp.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2033232212010332-3213011001213120-3320310112330011-1022111033111320-1112031131320012-1300001230032031-3100221130230213-3122323200300212"></a>

<a id="canonical-3103003321201033-0110312332323131-1011222001213001-1233001133222303-2023111323321231-0211020010311031-0130301210133332-3223213311203203"></a>

#### `gcp.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3121303103203100-1101323212230321-3230012332132301-0232101222130000-3122012312020000-2133013130121123-3013331110201202-2233232321003333"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1310330212312213-1020322321000203-2221131002021013-0203111133111211-2333010020211010-2323200230311221-0023310120021313-2221132321123133"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-009.md#canonical-2112021120030020-3231333020020100-1131200101210103-0002113133023111-0231212221103103-2232011230301021-2320301133223222-0030010301313122): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020): complete subsection reference.

<a id="canonical-2112021120030020-3231333020020100-1131200101210103-0002113133023111-0231212221103103-2232011230301021-2320301133223222-0030010301313122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0302330002233201-0301231330211202-3230300313203330-3131120002000011-3130030003000102-3300230131112320-1023333003302222-0312220133101033"></a>

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

<a id="canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-3210112001333003-0330021231033233-3213023112222300-1031231133322111-2310133000031301-1112133102213312-1022130100321013-2131012312110121"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0130123211210333-1013001111102000-1032200332203301-2020133211011322-1030231310122023-3211232101022022-2203030322113113-2010010221122332"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011): complete subsection reference.

<a id="canonical-2301212332222121-1302223311302010-0301321012231103-3302110312321200-1120323333210321-0221013211111323-2333130012210221-0012220300110130"></a>

<a id="canonical-1113011213030021-3201113011230222-3123113232231210-0002212332003301-0110223132021111-1233130133101210-1212023333201030-0113111130320030"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130): complete subsection reference.

<a id="canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2213011013201101-2003230021030321-3323100213110131-3123221023310300-1001112002020010-3022032332301110-2221103300221213-3312003012121320"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2102213311000311-3000002320301321-2312312132302000-3023202032222003-1200331222331111-2310223201023111-3220133200133031-2013030302313012"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-009.md#canonical-0102032310030203-1222102212030202-3102110031010021-1011113132210110-1231001002302130-1110122112302011-3013300020220032-2102312101013033): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010): complete subsection reference.

<a id="canonical-0102032310030203-1222102212030202-3102110031010021-1011113132210110-1231001002302130-1110122112302011-3013300020220032-2102312101013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3123210121001102-3023303213112331-0331313333133230-3113031202111230-1312021301120101-0233022232000230-1322322110222010-1130023212320222"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1013023121113213-1020012023010223-1112212213002303-3023201203123231-0333013222223301-0312230233222202-2031013100221133-0020103303102201"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1112231303131122-2032000212231133-2121033310320230-3310312332230020-2333321310312222-2113122210130123-3212333231103033-1030020132210223"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1333202111011101-2322313311320030-3322221213033032-1100320130213130-0022201020232123-0313311032011212-3230323103223112-0303212322210322"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0013002322021312-1031102021011121-2200020120031221-0300230133123001-3211032321000121-1033123003312313-2301331223200222-3323213303312221"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-0230000303331021-3330000101003213-1132221100232322-3032102013110103-1221303232312113-3120233223023003-3211303212132202-0030020013000122"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-1232122223032023-2013112221302031-0133111212023311-1012220213120000-1200313130123323-2313031311232022-3012133131113323-1220233303023232): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2210013013203202-3031130011301312-0023012332231020-2231032133323111-0220213130000200-1232011331021120-2111021033331100-0321031223002303): complete subsection reference.

<a id="canonical-1232122223032023-2013112221302031-0133111212023311-1012220213120000-1200313130123323-2313031311232022-3012133131113323-1220233303023232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2323120301123201-0031101110322311-3221210102333331-2312131030022202-2001333002123330-3113110300112102-1121313110201123-2101111230022110"></a>

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

<a id="canonical-2210013013203202-3031130011301312-0023012332231020-2231032133323111-0220213130000200-1232011331021120-2111021033331100-0321031223002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1001003103003212-2131100332301323-1032322120303112-2033322002213021-1021202022132010-3011212131333123-1203310231333011-0000331332203011)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-3112133003230022-1321230122331200-2213000130222333-0033010322123331-0321111000312000-2121232333003210-2322312012021103-2100230322300010)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2221011102121103-0302332000001202-3101322231233000-1120031101131011-3333031313320232-3310000233231301-3321223023323322-1222202130210102"></a>

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

<a id="canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3001330023321311-3203102220022223-0311033020031020-2231211301120031-3110222311202121-2111123011323230-1021332132220201-3320331210022322"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3113001123101303-1112321112012010-2033300300012103-1112133022222212-1031120321303113-1132121000330202-0222103001000121-1223010012130331"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-3212320310102033-0301112300311120-1023311303202320-3233102100033320-2301032121021000-2230222231131023-0302320031301132-3110201000111213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-2333333223310132-3110110313231231-1312033123112312-2332001202002112-2332300300203331-2300212013213232-3331122003212013-2213303000300330): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133): complete subsection reference.

<a id="canonical-0021332011031012-3110302121130110-0123223222003313-0331200121321200-0011020120202203-0012002201231322-1321333323002300-1312231302122220"></a>

<a id="canonical-0330302113001023-1321212001100313-1011021211312101-0221213331133113-1232100233001331-3232321303030302-2020222131232110-1221130301002300"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-2123230302111331-0223331032011030-3331322202310202-2202033132112023-0103133102323022-3233330203303323-1200202112333200-2102203333031113): complete subsection reference.

<a id="canonical-3212320310102033-0301112300311120-1023311303202320-3233102100033320-2301032121021000-2230222231131023-0302320031301132-3110201000111213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0212130303131131-1102021233313300-2023333013021232-1211011113311032-2120123130300011-3210320202010121-0121323331332112-3011302001300021"></a>

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

<a id="canonical-2333333223310132-3110110313231231-1312033123112312-2332001202002112-2332300300203331-2300212013213232-3331122003212013-2213303000300330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2303201111302100-3301232333111023-3301301331002202-3013011322310302-3021013121132031-3323223133323001-3030303313322303-2001102320332130"></a>

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

<a id="canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-3032131303303123-3031021131012011-1303332333121322-2021100231011123-2132132301300023-1322020110231131-0200112331213101-0213322023030112"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1013323000100211-3322031013120032-1220311321031201-0010100231010231-2333112313111331-1300000233133231-3210123220221333-3021210010202121"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-2302110133322310-1302122313132313-1131232013211321-3001303121201112-3010302230131303-0312310133311013-1333203102302111-2023012300200222"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-0222320203033033-0033111031133331-0220110232320102-3232101133010200-3302130213212232-3033112211103130-1300220100020332-0100103222232202"></a>

<a id="canonical-0012110220121203-2002031103311112-3101010330302022-1030233312231000-2232101023110302-3001030111201021-3033102012333212-0213301101213312"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
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

- [pools](resources--securemesh_site_v2--reference--group-009.md#canonical-1230002211003033-1030231113023001-3001220103002231-1203301211030111-0212030230320320-2203213023322302-0333001233001230-1323332133221101): complete subsection reference.

<a id="canonical-1230002211003033-1030231113023001-3001220103002231-1203301211030111-0212030230320320-2203213023322302-0333001233001230-1323332133221101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2121002223100232-1123021102103010-0213321022230301-1122233023301012-1302333021001122-3131001022321212-0101033123202232-1221032200311133)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1301331111221302-2231132131002011-3202012230122031-3022201120301320-0303030221002112-3111233031010002-0300233320210100-0212203002230111"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132010212120013-3032303322110031-1131232102323020-2032112211220322-2120030123000102-3002112000001220-0033112010031211-2310033200100213"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-3031012131130022-2331211101022312-2301100110131311-3332120113130111-2121030013012213-1033120212301120-3120333333323131-3233000311323123"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-0300021023312213-1313021321023300-2200233100033202-0233202100310330-2203200020121020-1202131310100023-1120011211123323-3030131323232310"></a>

<a id="canonical-2230220112302332-2033213011101010-0131213313301333-1203110012320003-2100202112012021-3220112302200021-3301112103322302-0112022131021032"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-2123230302111331-0223331032011030-3331322202310202-2202033132112023-0103133102323022-3233330203303323-1200202112333200-2102203333031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1120330000013123-3133121330020123-1002030233012021-2330113333213200-3120202210332130-0013033120020121-1101203323010122-1310223332011020)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3223212030322111-3002302101120010-3002213133012033-0220320211200331-2330102200332301-3321213020302333-1202330303302310-3002112102332130)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0013010012301321-1123311311233212-3001101303321311-1010131111012030-2303032020003212-0202203120213213-2322331112301301-2000023020000300"></a>

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

<a id="canonical-3112120003102123-3320003022101133-0012312330120133-0303300110230230-1010113011301001-1310033303331013-2221310220333000-3200331111203011"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-1332001203121113-2131303203010000-1230133333000223-0213111001100201-3322223303331023-3002002301110021-3033323021010333-3312120001030001"></a>

#### `gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1211223311200301-2113213233010230-1203012301302120-1022300331103323-0211120332111200-1030031000001303-2222221212332011-1213312211122110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.monitor

<a id="canonical-3011133112101111-1322133303013213-1220021130330333-2013232133102032-2211133323113012-1303332000003312-1233312330202120-2001001230310112"></a>

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

<a id="canonical-2110013100330201-0203313102212200-0032133203312013-0101212301101020-2313122133233302-2220111322101030-0222330312032021-2111000113010023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3122133323111102-0022203232223232-0101211033330010-3123332022111121-0213332222223023-2031020102303122-1013330110203130-0101312331313132"></a>

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

<a id="canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.network_option

<a id="canonical-1330310330110230-0303011333132301-2230311210222301-1331313130303233-3030211121111233-0031031131020001-2232101330303303-3113130033321111"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3210001110232132-2023233123002131-1023120103302102-3000012331100321-2032213022213202-3213301122320300-1001230032212003-2311132303000120"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-009.md#canonical-3333320202012202-0002100001111310-0103330113000112-3112313330312312-2001110311222301-3203002120212201-0201300202312003-0120323202330110): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-009.md#canonical-1002102000211210-2202011122321031-3100013213211323-2200032023330011-0013000130200301-0012233320202130-2320203003312011-3021003110001201): complete subsection reference.

<a id="canonical-3333320202012202-0002100001111310-0103330113000112-3112313330312312-2001110311222301-3203002120212201-0201300202312003-0120323202330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3232031223301220-1213100132311120-1112330021032320-3112320303030321-1022311322333030-1023231022220202-1321322333013331-2003312132323331"></a>

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

<a id="canonical-1002102000211210-2202011122321031-3100013213211323-2200032023330011-0013000130200301-0012233320202130-2320203003312011-3021003110001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- gcp.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1101130122001330-3200313132311102-0012333132101320-1012003123203130-1021002313023211-3101101032010201-0211013312312312-1132332022032300"></a>

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

<a id="canonical-1203323220122012-2201322220300230-0230002310100030-3330031323321311-0300313103011131-2300131311231110-3013203313301020-2320101230133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-1221003113232322-3112200022301210-1213101331310110-2002021133031101-0001100213131323-2103311331231201-0303222121302212-2321322032021313"></a>

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

<a id="canonical-3000103023231312-2301010323330333-0333003302000002-1322000323123232-2002220330300031-1331112001201102-2321333303303203-2001023303022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2232231013323133-3223302223010013-2303030223033313-0000223120321231-3320113130310310-3021130112031220-1333020210212230-3233012201232202"></a>

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

<a id="canonical-1133201233031230-3333012130033031-0221332312123131-3003211133222313-1013122333012331-0100231220131122-1212222211103002-0131010321033230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3201230020221132-3322132103310220-0113330033100031-2030100300232321-0303012020132333-0212222101002330-2100330100100002-1110313022130022"></a>

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

<a id="canonical-0220100020332132-3221320131203333-1233013032031323-3222113210012223-1121032113121201-1001013201031031-1210302012100032-0323200312201112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0333100100202200-1032122232320002-2320331101223032-3010301203011122-1333021033220321-3302321030121220-0133000323223100-2301223332112220"></a>

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

<a id="canonical-1030120100013303-3022221311322001-3212222021021102-1131222030210221-2220333320001121-0123333131111212-3121032230202313-3133313121120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.static_ip

<a id="canonical-3322232032012023-3111223102312021-2220310300031211-1123112210221212-3131221102230120-2022320323303220-1331013100101100-3011021002121002"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2220212000312130-1030123323222113-2313033010330311-1301011322232123-0000220021023113-3103013310312111-1011200300022133-3102033303302203"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.static_ip`

<a id="canonical-0012032321231030-2231232231130333-3012020210002013-1202201203330030-3102101311221332-3100032101010022-0021032303111100-3011002033201330"></a>

#### `gcp.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3020120111232210-2102110103323202-0010231223121032-1203213213222112-0012120310230320-3331021020120323-1112112321231121-1230333110220110"></a>

<a id="canonical-3302210312103131-1202101301031001-1320322332211223-2021213211301313-0011222001201132-2322111120021011-3232221232133110-1122030112323332"></a>

#### `gcp.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3113320003112323-2023200210222112-0231103320303012-3003333000330332-2120201113112323-1123130120312120-1111033303223320-0310223122221311"></a>

<a id="canonical-0201310103313313-1012223000122123-0021320130332231-3333312020030201-2323311113330331-1032222302010331-1001030113313332-3212300230022203"></a>

#### `gcp.not_managed.node_list.interface_list.static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3203201112332013-2123213103300133-3111100312201222-0122232323111013-1131322330103333-3311010222120313-3031010030320213-1010212220012232"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1010303301000221-3220301322311313-1312021001323010-2220322233230222-3203013131020201-2112003322003012-3001110323323332-2110322210032311"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-3020020320210103-3302022223103100-3202322112032131-1311112321131031-0313211102210332-2300213000303300-1200033322223220-1232222222212002): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-3212312131231313-1103210201121200-2232230110113111-3022100023003203-2030230112233301-3002203022132033-1021303122001320-2212221313021301): complete subsection reference.

<a id="canonical-3020020320210103-3302022223103100-3202322112032131-1311112321131031-0313211102210332-2300213000303300-1200033322223220-1232222222212002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2332222231303130-3110133130001320-3122102203023232-0212000332330203-0311310133321311-2320010022302101-1320320312012210-1221101321321103"></a>

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

<a id="canonical-3200023132232023-2232021233201212-3310310000003010-2032123031011220-2013303200311303-3001310231000022-2130201131133021-0031101022320102"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-0011332330010000-3333112302132303-3333132311113310-1331333002020330-2110303131112111-1323113132101032-2033002001332203-3323030223130321"></a>

#### `gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3212312131231313-1103210201121200-2232230110113111-3022100023003203-2030230112233301-3002203022132033-1021303122001320-2212221313021301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-0322132333223300-0101202020231122-1322010031300301-1133112212123112-2021031023222133-1013210110123212-0202021020323201-1332033310120130"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2022201003100222-0331131310001122-2311220223013333-0302003333231101-0201231011322323-0010122020302100-3031311301300131-2011000303322202"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-2133113323021033-2021231212310330-2020002122300310-3031321323220102-0023033201300221-2132231123123030-1313031123001330-0313303110233000"></a>

#### `gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2203010120021321-0012122132132023-0311001223032011-1023302201001000-2323220312032130-3330023121101312-3010121130000332-0203020210231200"></a>

<a id="canonical-1113210131011231-0311122321120010-3231221223102320-0321223131000122-1121123331130203-1020211223000130-1303221030010331-3313020201000012"></a>

#### `gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3111001121103231-3221330003020002-1120021320332011-0020302211002113-3121123113212232-3233210321312112-3302012203100121-1110200113332023"></a>

<a id="canonical-3322103321303002-0230012221321010-2121010202110122-0000130131112312-0301331013210213-1331231110000023-0232100310331220-2203021330131232"></a>

#### `gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2322033023330322-3002023133132222-1111333320023023-1223323000130112-0000221001032331-3310323132320101-1112120210112123-1112101101200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-008.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-1332210301322023-2303020332221120-1002032331033200-2001110113233121-2223002310212110-1320221130123100-2300113003123123-2211213003330310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3101331012201120-0111322330113331-0132300221232210-3012102001011010-2120020011230001-3223012022301003-2203301022000222-1310232000010112"></a>

### Direct properties for `gcp.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1302030203310223-2020122133303310-0000133203122212-1033112012332322-1333222100231230-1211200120031033-2111301203131110-1222220120323133"></a>

#### `gcp.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1232101330033202-3023013122202030-2211121022020220-1300213320323300-2322021013031013-3223313021233310-3020330010320010-2131211103331023"></a>

<a id="canonical-3320123210123331-2211113231003023-1022220111320332-0223202102030121-1222230112120231-3023021230022122-2032130133232330-0221113232323313"></a>

#### `gcp.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- kvm

<a id="canonical-3200011220002023-3002321213101032-3320012101030001-0002211010233312-1133321102002011-0013220032121103-1330120002111201-1130031202103111"></a>

Type: `"object"`. single nested block, Optional.

KVM Provider Type. KVM Provider Type.

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
kvm {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222132123022222-3131320000113211-1310102203111333-0231011113211110-3133211011101210-0300223232103112-3201232130202002-3201021112330010"></a>

### Direct properties for `kvm`

- [not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230): complete subsection reference.

<a id="canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- kvm.not_managed

<a id="canonical-1220031302201113-0112001003002033-0111323113131021-0122330230130101-1121310000313202-1333130033213203-2111131211010322-2231022321321213"></a>

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

<a id="canonical-2222103031110130-0021020320102031-3030223002321221-0330002232120321-2313121110021110-1120303223102031-3102020113032303-2003331020312202"></a>

### Direct properties for `kvm.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112): complete subsection reference.

<a id="canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- kvm.not_managed.node_list

<a id="canonical-2001222030202010-3013102220100331-2030012022032322-1032103311000332-1311200020021031-1130222311112122-1102022002301110-0021030320001133"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3030311312133123-3333301122213333-1231321022123333-2323030221211212-2313001212112001-3102203022233311-0100111330202120-3201333001223202"></a>

### Direct properties for `kvm.not_managed.node_list`

<a id="canonical-0320232333031230-1002302023133212-3001131212020110-2013112121323030-3003222131012323-1332100233320233-1101322103332221-2111233132211321"></a>

#### `kvm.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203): complete subsection reference.

<a id="canonical-0123223130332203-0302210031001133-2232010300331030-0301303300103330-0300222201312331-3002023002322103-1001011310322123-1100200331001033"></a>

<a id="canonical-2311313321313133-2022302203222312-0213003310302333-3011103120030003-1103101112020212-3213013332120012-1222230320310323-1130010033233302"></a>

#### `kvm.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2231132103032013-1203201302301021-3122012021310302-2010213110131002-3200331000121030-0011122021231303-0021100021032201-3123331011022323"></a>

<a id="canonical-2023300132300022-0102303302001331-2031310320131320-1221032220012002-2323213233000201-3221213122322003-0021123320031323-0110332231233110"></a>

#### `kvm.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- kvm.not_managed.node_list.interface_list

<a id="canonical-3231220113131302-1202000222223103-0320221211232311-3100212200010030-0112232120220031-3013210133001212-3311302130123011-3121121030030033"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0031331333013013-0311133120003132-3233200122030233-0302312021010203-0322313321033011-0213320322110113-0332303132211000-1330003010130200"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302): complete subsection reference.

<a id="canonical-1013113010203211-3222013230322010-2023320003222111-0300112111321010-1302101321033131-2310220001102031-0302103030313333-3112330330310032"></a>

<a id="canonical-3023120213033103-0033313013133032-1121103031203310-0311123221031331-3010013002012222-2120101122002101-3220102223011133-1321010111030311"></a>

#### `kvm.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-2120300022110211-0201330330303132-3312130313332333-3232322022112121-1121032202220011-0230013130201002-1102333120120132-0330321020121012): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3230303333012202-1021300033312101-1113102213332132-0123230310002231-3301112312020310-0003210212213203-3321200333100110-1110210321322322): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311): complete subsection reference.

<a id="canonical-2223310020032302-0101112133221002-1332003221200302-0021231221212023-1032002323123223-3000303110233031-1123222220212102-0131130020312211"></a>

<a id="canonical-2212103123132033-3313311011112011-2223222100023003-1332020333323122-1102020300332200-2032213213021322-0012332023130131-2132013121132122"></a>

#### `kvm.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2031133333322020-2201212133023222-0203113212121312-2302010223012101-2230321130203103-1102233102003303-3120330322301333-1123022001132201"></a>

<a id="canonical-2303231303233123-3310013322131230-2231210200333221-3011331223103111-0321321111003110-0112122133211003-3303302213230313-3002121012221223"></a>

#### `kvm.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3331120301013222-3113001233123322-0331002112320330-0020112311030321-0022120311131032-3323221022313123-2221202322031033-0211200331231011"></a>

<a id="canonical-3120113122020230-2032230030221313-0013001320000313-3123133222131110-3121200320002301-2221331331102113-1200212020222312-3333312131020103"></a>

#### `kvm.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
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

- [monitor](resources--securemesh_site_v2--reference--group-010.md#canonical-2012311020303313-1030231232300200-2221030123233323-1122220201031103-2332220103012331-0300232332201133-2300231200212201-2020333320021213): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-1100202122033020-2002013131322332-3123310302210202-3133301323002113-0313101011311302-3121103222332202-0200231131113331-2121223300012122): complete subsection reference.

<a id="canonical-1003103030031012-0121022032000311-3111010121030003-2003332320213011-1111231022122000-2332121031303211-3311301220130230-0300222330113212"></a>

<a id="canonical-1323230010111133-3111210202012320-3230302022303223-2100001221131332-3211210022221021-1211311321213311-2222332333012323-3133300122201102"></a>

#### `kvm.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-0133120023022023-3011322131222123-0001301110132312-0131013311223320-3301010200031302-2032311031222131-3003200133130211-1200012220031202"></a>

<a id="canonical-0213101210203231-3323320121311202-1332120000230333-3210230330000203-2301223232033303-2002131320012203-2012020231203030-3120020223320000"></a>

#### `kvm.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-010.md#canonical-3123131131203320-0311213232102301-1111221021223032-0003223303123123-3221233231103221-0223301101112102-2323310000121122-0302333222323201): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-0311010103223322-1303203103333303-3322330303011033-3111120302320230-1100013130210201-0011001111033233-1112101120100212-2202332313302101): complete subsection reference.

<a id="canonical-2233230112022301-2331123302321111-0021333020130020-1301220010211322-3033220220310122-3211320122320001-0003001232230232-0321002111000221"></a>

<a id="canonical-0230002111300331-2203133301233001-2313222223032212-2221201221330313-2330213021000333-2201302113012211-0333101200111332-2322223203211200"></a>

#### `kvm.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-3332312113013203-3301201000000220-2221200231232212-2203012012230202-3321023213020032-0233201012101203-0132330002212211-3013111113220233): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-010.md#canonical-3013032133102011-2021301032001100-1323130231102113-1023120032112202-1012023001233332-3220030033113122-3220021230120310-3031131022131111): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-0131210010111310-2321312120022211-1112020000212131-3230212132003223-0200222021213212-0100033002013220-1222130322202101-1321121211213033): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-1203031212100311-0311310330301021-3222021132312202-2201122021003312-0323010310321021-1223212321312130-0322333300122112-1313302312011123): complete subsection reference.

<a id="canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3201312210120212-2101330120000123-0331032013302120-3323221223132131-2011202120100023-1013103122032233-1311331013100011-2310301133220333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1331300312013312-0201102330121221-1133030132032102-1033211101022320-1213200121022112-3232112001001133-0122002112210133-2100020102313210"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-0332101012330232-1031102131202332-3133121301210002-3231023203210331-0200020121313032-0012213210320030-1003220131133320-1111330221001100): complete subsection reference.

<a id="canonical-3223120311203201-3001303200231121-0332001002213232-3102321333122322-2012000231202033-3103100121331201-3031011210020001-0212323232031201"></a>

<a id="canonical-3103033200110113-2113111202300030-3221101131130111-2322022220002100-1131113111013313-3131223132320202-0121231231220233-2001322311133012"></a>

#### `kvm.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-2030322113302030-2222013010030213-2223312220000212-1022113023332023-2313233111032021-0323121312021332-1233221013003220-2013210131130100): complete subsection reference.

<a id="canonical-2011003312203302-1212303331310303-2221131002032331-3231023131333333-2313030020123130-0321313223131301-3010300000100221-3132020330131003"></a>

<a id="canonical-2012123012023023-0220123223301122-1013020222121130-3012011320221001-0112222312222113-0032200011312203-2013000130233102-2121032112202020"></a>

#### `kvm.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1321220213123023-3302321010311101-2122033301233031-3321202203000300-0321120130013002-1211300031303131-2101022330033032-1011320323121032"></a>

<a id="canonical-1033322333122323-1232023102303010-0100120232332002-1210231331122200-0121311002022223-3121111200123302-1133312212312120-1101111020222121"></a>

#### `kvm.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3122313223222000-3130210002122233-3003213033022213-0203132112311001-3211133032012131-2112103003020321-2123003330022023-3111200230201002"></a>

<a id="canonical-3203310303132101-3103033133320131-0313100311332222-1131111200121133-2033331112030222-0020110121212102-2210020230132100-1221310203333211"></a>

#### `kvm.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0332101012330232-1031102131202332-3133121301210002-3231023203210331-0200020121313032-0012213210320030-1003220131133320-1111330221001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- kvm.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1030301102223111-2131000102132012-3112320011021101-1021202330332223-3220332123113302-2113131311303021-1113221323031301-0231221133131130"></a>

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

<a id="canonical-2030322113302030-2222013010030213-2223312220000212-1022113023332023-2313233111032021-0323121312021332-1233221013003220-2013210131130100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-0113003200331131-2200121310001132-0320130023311211-2102122300203103-0130303103023033-2232023303210312-0110310112312313-0232301132220302)
- kvm.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1123033111000220-3333212131033010-2012321323032222-2002221300022020-0213112313200133-1111013230231210-2013130330132223-2310021233203012"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3200202122113321-1232001310203123-2330223311111233-2003110121220011-2032133210332111-1112331123032222-0330220312320313-3311223311003030"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-2131011301221230-1133202131203301-1103100301211011-3232100313031021-3023331233013302-1212103103002102-2011002203332212-1213031003121000"></a>

#### `kvm.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2120300022110211-0201330330303132-3312130313332333-3232322022112121-1121032202220011-0230013130201002-1102333120120132-0330321020121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0323330113201303-1033023113023003-2333031322000130-1211112221033022-1022331021322101-1103313113203111-0013311012223032-3303003303333133"></a>

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

<a id="canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2100033213100331-0301233002121230-2301223311323231-2300023300212303-2312233212013003-1303320020201203-1331003110302111-3022311203210131"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3321321003022223-3203211020002331-2110200130101312-0302220313000330-2333303323222313-0012310232012302-3102302200202310-0002102233233131"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-0233202102020033-1212330233223200-2020231002031320-3201230001200133-3130333300122331-2021002303222132-1301133322003002-0131321231011012): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-0233312221032103-0330200223321020-3012102221120133-0203220111022302-1222123221112023-1023301210020321-2230020010130022-3203110210300232): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303): complete subsection reference.

<a id="canonical-0321021311031303-0232303120012121-1331322102011013-2020230310011322-2102120200233123-3033101122132021-1323012203222123-1101231003033130"></a>

<a id="canonical-2123011202311233-0013021230303001-0010231213330031-2001312302232212-3112202333000103-2211223012332311-3000102021123233-1110202232001112"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3223202122121111-3313331220001012-3110013110023102-0031131102010032-2201033133133110-3202012221000303-2300200020013000-0010022221320320"></a>

<a id="canonical-0322021313202222-1323331233021101-3112310022022303-3332213013030112-0323210013220033-1010133223030132-0033032102220103-3312000020200221"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-1002201022001012-3122312230232322-1221131303323001-1130322310223220-2022221102320120-2301322300212300-3332131001022330-0321100200303303): complete subsection reference.
