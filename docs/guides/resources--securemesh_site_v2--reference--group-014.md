---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2311313123033232-2223103100323032-2003003030123300-3233212203220201-1102330330200231-3121330231210030-2321200201302301-1100103310133300"></a>

## `openshift_virtualization.not_managed.node_list.type` property

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

<a id="canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- openshift_virtualization.not_managed.node_list.interface_list

<a id="canonical-0133110031123132-0212113111332303-3320112132233213-3030032132313101-0213102111111313-2320120200032021-1032300103230212-3303133030232300"></a>

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

<a id="canonical-3032222202132200-0321202213100332-3000201313113211-2332300121213023-0003310130030021-0313133301330033-1201303103031322-3111130111110131"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-014.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322): complete subsection reference.

<a id="canonical-0312133301131232-3312302233001110-2113110100133021-2330202232213312-3321003230003103-1113023221011310-1121312332300133-0200331133120230"></a>

<a id="canonical-3111030200210103-2233033331103000-3010233111203223-3011213231230311-3111332201122111-3131233211031231-0311232030001111-2101331130232022"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-014.md#canonical-3120030010211121-3020323231113103-2102330210201023-0020121233001020-3123302220131333-0031310222210310-2222121123112312-3023231322012233): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-014.md#canonical-3211030230302131-3130002303002032-2001210201122113-0231133232100330-0113210301133023-2310223332332022-1330130120201221-2303012321020312): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200): complete subsection reference.

<a id="canonical-1333100131021021-3122011310000102-2331213120322001-3030111101231112-1033300010220110-0130203100002201-2233300000230011-0012202123201013"></a>

<a id="canonical-2003121231102131-1312312121333103-1102300303121102-1101002200310210-3313100132310111-3223111220120010-0213121111222101-3132202113220032"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2331223002322300-3213310331210003-2213100322300112-3320021020020211-3130020023313233-3332023203230131-1220121201003211-2023021310212333"></a>

<a id="canonical-0322102101202123-3110032221010130-0122322311031202-1021302023303212-3113221230031003-2232010032321011-0030320313213230-1113122223000322"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-2230323123311012-1012110222111132-0220010203122131-0203302202012020-1230212310203131-2223313231131202-0020222022323010-3103102122112320"></a>

<a id="canonical-3210131233032323-3202112122032312-0021032333330333-1121112101132201-1313230331200233-0312123311313313-0002302301303023-0112323222321003"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.labels` property

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

- [monitor](resources--securemesh_site_v2--reference--group-015.md#canonical-1213022213322102-3303133001122212-1230310310003310-3120313120012322-1033111331100220-3021211102013112-0233331131030233-1333203312131030): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-015.md#canonical-3323211302312000-3232023112321122-3023321012011011-1110301230322200-2331101002310313-0300011001120332-2221211310033210-0131321302233012): complete subsection reference.

<a id="canonical-3230303021101323-3201302220211110-2321132003100202-1133003223001330-2312231131111030-3032233303232111-3113302010121300-2201300131212322"></a>

<a id="canonical-1212332031202201-0302323301103101-1233120202030212-2000201120202220-1110022130320032-3301011102300321-3232110113112330-1132220303331213"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.mtu` property

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

<a id="canonical-0211033211121133-2312110313311223-3331232021333300-1212010223230322-1132311133123323-1103232303211003-1200010113331203-2002130111302133"></a>

<a id="canonical-0320202000233202-1200200131212232-0323212300303311-0201211010033130-2012302122011331-3001332310011302-1222201220220111-3321023330102122"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.name` property

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

- [network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-015.md#canonical-1011011213323210-3122010003102300-0330331301102032-3333301223303111-1201212112100210-3332111231323310-3202020132312321-3022102002123201): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3000222332330022-2133212331113332-3312321101221300-3212001103011211-3331313203222203-3013001232002331-3301220303332120-3023210001123102): complete subsection reference.

<a id="canonical-2103023311312232-1321312103233020-3133230202101012-3330130323233030-2312113133300311-3220331133231122-1200123000200200-0300001220332200"></a>

<a id="canonical-0000310011030003-1022000312120032-1103132121030030-1021110031121023-1020312313222023-2300010022112211-2201231303103112-3332233101013133"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.priority` property

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-015.md#canonical-1021203102231201-0231333013111010-1313030332200030-3022331203000230-3220220010313110-2231223230203220-1123111001231323-1033123310230230): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-015.md#canonical-0330203123013021-2220332110311313-1132132123211300-3302302130110331-2120002110212333-0101220202112101-0333103021012030-3022221021223113): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-2011222002031302-3200302122333032-3123332200133332-1213200011123210-2110002011322231-3321210212222320-3231013110330213-3131112113021112): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-0030030233110130-3213101333223331-3101211022101203-1022313322122110-2223021203331030-2030210001231300-2333232113003101-1103112122103212): complete subsection reference.

<a id="canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2122012313313212-0103233320222101-0020011223010311-3011111132002330-2021033111301113-2203210210013133-1133232213131103-1103123103310120"></a>

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

<a id="canonical-0321212100313302-0212102300120133-1312300231130302-1131210312000333-3020011202332230-1112311311312311-1031300212212010-2303220223222323"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-014.md#canonical-2033310201130303-2332210213321010-1000222221121002-1200130210032320-2312232023311131-0221303022331110-3310230003202133-0303301213112131): complete subsection reference.

<a id="canonical-0320203333322013-2231120301031222-1303020002011330-1223223000221213-0322030331103320-3331130221120033-2233310201311332-2202312001100020"></a>

<a id="canonical-3303330130203233-2200333332122300-1311330030003002-3103131222113220-1311221221103231-0020102220111222-3230302003230311-0031333301321312"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](resources--securemesh_site_v2--reference--group-014.md#canonical-3210132312220113-0202310201202011-2301210123103023-0133323331330202-2210201022023300-3103313033312102-1000021131102300-3122313302223021): complete subsection reference.

<a id="canonical-2130202023211212-1102220121221112-2013031130311012-0333132230012203-3011320330110321-0310203303032103-1022312110013330-0131310323320032"></a>

<a id="canonical-2120103021333213-2330212110001201-2221030033112310-0300120231130203-3213011333103201-2323203033311121-3031320221120230-1110102123312020"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-3302210122233321-2103110200103202-1333222030310103-1302123210301123-1302332133022301-2300130331032302-1033202030132323-2113221011233300"></a>

<a id="canonical-0123002033102333-3100303102133312-2031111212013232-3232001202100131-3132201011131121-0031202302000220-3120121233320003-3223013131301021"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-1112312013203202-1230010320220330-2020013023333023-0222130232113033-1220031002320211-3310220312231200-2013222103321120-0200322322033021"></a>

<a id="canonical-1303120120111133-1012331003230333-2301302302002113-1323232112300001-1232030030123230-2133102210200012-0032223131200033-2122031131222220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-2033310201130303-2332210213321010-1000222221121002-1200130210032320-2312232023311131-0221303022331110-3310230003202133-0303301213112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-014.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1022130202310100-0231113132220011-0210221020210313-1300311001011112-1233233333320122-0202030012000301-2121112231222323-3011130133003222"></a>

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

<a id="canonical-3210132312220113-0202310201202011-2301210123103023-0133323331330202-2210201022023300-3103313033312102-1000021131102300-3122313302223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-014.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-2213032231012001-0121000130301120-2130322021100130-1323013212332211-3030020032032001-1220302111030000-2211221332033000-0100332310232230"></a>

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

<a id="canonical-2023302223123021-3132333122003010-3102111110133310-2122302220233122-0123322311030113-2002013301100202-0120133301233001-0222022211331221"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-0120012312022332-1020211131202232-3000220323213322-1011120232331130-3030033000101300-1033032212133220-1213110211220331-1203212201120013"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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

<a id="canonical-3120030010211121-3020323231113103-2102330210201023-0020121233001020-3123302220131333-0031310222210310-2222121123112312-3023231322012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0222212231303101-0121300112310310-3203033003230011-3300310123230113-0100332223103301-3132031300203300-2002220133233133-2202311303202232"></a>

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

<a id="canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2000220210211231-3120302311332311-3000213233332031-2132303320130123-3132321212032202-0203231230020123-2111220220002321-0210100112331210"></a>

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

<a id="canonical-1300103202213202-3010311001212310-0311121203213201-1203231001233220-2301331320031203-3223222131202113-2011120300302010-2331302032033112"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-014.md#canonical-0323210112011302-3333133111330013-1311303320022201-1011322122200100-0021221110131111-3112111200203023-1103202220220033-0220310012233213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-014.md#canonical-2021221123033000-3232132230203103-1201011301313320-1003020232200313-0113210202113232-2110010001023313-1103110310122113-2320112301010032): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-014.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133): complete subsection reference.

<a id="canonical-1213301230131000-0321110132201123-1003313131230231-3132213303311313-2131321010122231-1200023031312003-3021202111211323-0330210220210001"></a>

<a id="canonical-1302221130112130-3133232113100032-2320223110122220-2212031103300003-3321020320113112-0331310111132113-1231221201002203-0322330020202122"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0333111313122303-3223203023230322-1220023021011230-0000233102321110-0033133013122103-1332300222003332-1012110232322123-3121020002012102"></a>

<a id="canonical-2132032211303001-2132231220111102-0212323232120021-1033101312123223-1021101011311232-3022230320323303-0033022231233011-0103033213312030"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-014.md#canonical-2223230122102210-2333032312101132-2211310112202313-1203223211311122-0021030230323003-3332101001331233-2020320223212330-3103111020323011): complete subsection reference.

<a id="canonical-0323210112011302-3333133111330013-1311303320022201-1011322122200100-0021221110131111-3112111200203023-1103202220220033-0220310012233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1202220313230323-3003202011222320-3312003230321210-1133021210122200-3031303313131302-1201121310301210-3122010032312302-2132203011132232"></a>

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

<a id="canonical-2021221123033000-3232132230203103-1201011301313320-1003020232200313-0113210202113232-2110010001023313-1103110310122113-2320112301010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-0223203111202001-3222001102313010-0101122210200112-0331112313023302-3202323101133020-2202303020221002-2133303222030033-1212021221211023"></a>

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

<a id="canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-0010300122212131-3131313330032112-0032320003113020-1210130010101000-3120321022322223-3210110222323320-3211201033113112-0302011301320321"></a>

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

<a id="canonical-1233001222313332-2123122222203213-1333331213131331-0002120310311213-0130102101131330-3031200012330022-0322210310132302-0131103300312002"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-0033112233103100-2012020313112003-3033030300113333-1100230033330203-1031320110300202-2002112003311123-1312013313202301-1021120100321202"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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

<a id="canonical-3233021301211131-3223001010123322-2232031203010002-2330133001311331-2123232000223003-1231021303203210-2122323210021331-0003302120133122"></a>

<a id="canonical-1130122320301220-2102022232203301-0101010120233222-1313323332030130-3101010201032200-0333131200200010-0303323100300332-1001120222003300"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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

- [first_address](resources--securemesh_site_v2--reference--group-014.md#canonical-0331200211020023-3330122302320202-1010302311122223-1023222021331211-3301021123312331-3332303013332233-2331002231121220-1300022211013313): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-014.md#canonical-3222003331300113-2320331123001013-2300130320023200-1323030310331220-0321320103201312-3210213013222322-3012131302013332-2101103310320011): complete subsection reference.

<a id="canonical-1320332332030220-2332311302230330-2223110130100001-1323221031202121-0123201230123021-3230321300312330-3302102123032221-0100031032303303"></a>

<a id="canonical-2002103133002200-2130010230232031-1333011030102031-3302331231120223-0200221022330201-1011212101320313-0331331102312111-3213011222210103"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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

<a id="canonical-1013212223120102-0123303033222033-0100110033322221-2032101210303320-1203112013013202-1022200230223003-0011021202231222-1013331321123203"></a>

<a id="canonical-3111031003223202-0011233322010112-0223012210212201-1113212032302120-2321032033101020-2101222202313101-1132131330113322-0312001101321311"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-014.md#canonical-3221032312233123-1211121022123310-3303002100132210-2002122012332001-1222033213110133-3203002033301002-0203310030111120-3111232010021213): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-014.md#canonical-1221122231010011-2031131320331303-3111132232220311-0133313033111330-1221130003132303-1031033301302101-2113311013101031-3201113330331203): complete subsection reference.

<a id="canonical-0331200211020023-3330122302320202-1010302311122223-1023222021331211-3301021123312331-3332303013332233-2331002231121220-1300022211013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-014.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2212333212211020-0021001300030320-3131300100303132-1031023031123220-0210213023131001-2023212310332320-1112022202103021-0202203210320231"></a>

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

<a id="canonical-3222003331300113-2320331123001013-2300130320023200-1323030310331220-0321320103201312-3210213013222322-3012131302013332-2101103310320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-014.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0223103003300202-2130302011211002-3210101323200220-0011001033013002-2321333110302202-2012212222000032-2133300031322100-0300012310031111"></a>

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

<a id="canonical-3221032312233123-1211121022123310-3303002100132210-2002122012332001-1222033213110133-3203002033301002-0203310030111120-3111232010021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-014.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2222321001132202-1332001231030003-3231231201200112-0320232312323230-1130333320110201-2303023122000300-3202023201300330-0130330201030110"></a>

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

<a id="canonical-2032233321100031-1322230202222033-0300020233002002-1201032103313213-2121320112331320-3312032030232023-1232010112130231-2000202022032330"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3331213220303012-1301220312111020-3102212330332323-0221213122211011-2311203231033023-1021123013302033-1331111313313133-3021003202322200"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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

<a id="canonical-1220210101333202-2210232330130320-0313121121232001-1112312110231110-2110002002233132-3012223102012312-2121023312230233-1211102231232213"></a>

<a id="canonical-1312112323331230-1312221131311201-1010210123231331-0232010122002131-3103303101000103-2221220001030220-2033023201210323-0101033201012313"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-2021232322103313-2310003002202331-2013032212130311-0101301113301122-0120011330201231-3211333313022101-3203133220230133-1312002330203112"></a>

<a id="canonical-1030102002200103-2033131111233300-1033221232020101-0111320031121133-3031312133130003-1333112123010130-1111310200320003-1311333311232002"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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

<a id="canonical-1221122231010011-2031131320331303-3111132232220311-0133313033111330-1221130003132303-1031033301302101-2113311013101031-3201113330331203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-014.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3330003002303000-1132231113300121-2330101220002031-0232023302021002-3102211331131331-3200021010331301-3300201310313112-2312011200021230"></a>

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

<a id="canonical-2223230122102210-2333032312101132-2211310112202313-1203223211311122-0021030230323003-3332101001331233-2020320223212330-3103111020323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-014.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-3232322300123021-2332323211221330-1300103203302101-2003123022311311-1001131032311200-3220120010123320-3112233322021113-2023130332313322"></a>

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

<a id="canonical-0000033111332300-3321200300200300-1032122132312020-0030001112121302-1321022030230112-2200102031113011-3122222030322001-1311020030333303"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-0002310222132022-2302332012221200-2230213033232213-1030222102323111-1220310210203330-3113332233300231-1300021311100231-3302310313030220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-3211030230302131-3130002303002032-2001210201122113-0231133232100330-0113210301133023-2310223332332022-1330130120201221-2303012321020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-1110202022021132-2013321232002121-3233011101321102-3111101223332012-1220200112133220-2323203022223211-1000330010132010-0230232010233100"></a>

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

<a id="canonical-0223113103303233-2322202332002333-2202002123021113-3221303202213113-0001311220021302-2311123300302331-0301233101103122-2211121331022200"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1122331230230312-1210131201010333-1311121221103332-1310121211332303-3133313332121312-2201212300223010-3003332212101112-3323012230322203"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface.device` property

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

<a id="canonical-0213302221111121-2020120213201023-3221113300102021-2220121132110110-0122220201202203-3130323320202111-0211323202313220-2012102003111012"></a>

<a id="canonical-0022113122110123-1121022210312210-2311131100210222-3233111002020133-2023112000013000-3001212221201012-3110002321222032-1100332211131100"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface.mac` property

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

<a id="canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1032311020103023-3133121330130220-0303323201212030-3231022323013321-1032201012231030-2112111201133101-2012122303123123-2102330000321231"></a>

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

<a id="canonical-3323212101330031-1331123003023331-1003211330121333-2310333212331012-0310330110233302-0110132032230023-0111122033112121-3332310001302333"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-014.md#canonical-0100011021021321-2033021100023100-1232021332200320-2032012310211123-1213023333123121-2001032022231122-1212312230002132-0202210201203020): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010): complete subsection reference.

<a id="canonical-0100011021021321-2033021100023100-1232021332200320-2032012310211123-1213023333123121-2001032022231122-1212312230002132-0202210201203020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2301102331321213-1333121200332212-1303313110202033-1330310200013033-2001101012331221-1030131002322332-2110202110030032-0020323212022303"></a>

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

<a id="canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1221313221000302-3200120023110103-3012123231322032-2122011112322101-3030213232013230-1323302132323321-3023233101333102-2101221001331232"></a>

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

<a id="canonical-0313112233030130-3222210011223320-3000111100233301-3301321302032320-2102312130300323-1313311021213122-3030013033233303-2310302003030233"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120): complete subsection reference.

<a id="canonical-0202103223021130-0123112200033121-0002330133132030-0333330003233303-3030133132223113-0120113221213231-2030332321310113-1102133121302133"></a>

<a id="canonical-0302203331021023-2223120221122101-1032112013132010-0100010131012220-1330220200012001-2021031323233101-0133201000013020-3332021133322231"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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

- [stateful](resources--securemesh_site_v2--reference--group-014.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200): complete subsection reference.

<a id="canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1310310200221112-0223013311023230-2122230312302002-1330211202132233-3123213210022022-0100311011310013-0033013330232123-0300213222120012"></a>

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

<a id="canonical-3032321121332303-2103120133131133-2030011010020203-2031120111202132-1213332230123200-3232201131023233-0023122033131011-3313231031222012"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0003210330003222-3221133323332110-2203230221223302-0111203111321321-2131120202330121-0033022130232003-2022201213300000-0003302022113333): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-014.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230): complete subsection reference.

<a id="canonical-0003210330003222-3221133323332110-2203230221223302-0111203111321321-2131120202330121-0033022130232003-2022201213300000-0003302022113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2001303312301232-1102023132311122-3131102311010002-0233010111001300-1233321333021030-1003112321231031-0202212233011111-1323101213331132"></a>

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

<a id="canonical-3312321011133301-1021113131203010-2300030210021321-3112002121211301-2012230221200110-0011233300021110-0201122330032213-2022112021201223"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-0331331123030333-3000212003300011-3123132231010102-0223032020120330-3200013020031313-3110211200301121-3230223310301113-3210202021110323"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2023213000031310-0011120011001232-2103311002212210-0330203212032130-0301003120131031-3020133320221030-1022012131200332-2120223030320033"></a>

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

<a id="canonical-0123210020021113-3011132220021232-3122301132030200-1023301132121011-1211213232212020-2013232020232211-1200101123122012-2220100212032111"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2022102321120220-0232022013001320-1032003112003211-2303332200300013-0311220313312200-1032121222323101-2211312023121313-0223213300121113"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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

- [first_address](resources--securemesh_site_v2--reference--group-014.md#canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-014.md#canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033): complete subsection reference.

<a id="canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-014.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0131000302230131-2000223012203011-3003000130113231-1133232132100232-1203303122022201-0221203031220022-0202133122121321-2013132022131112"></a>

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

<a id="canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-014.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2111132102221232-1311321332030233-3211122301311030-3222211313032111-1012212221023212-0202132320121020-1230212231030233-0021022323000031"></a>

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

<a id="canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-1203003203230303-2012002030222300-1312221333200331-2120030231303133-3111221222110221-2221130010301112-3000113120303222-1033131220013323"></a>

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

<a id="canonical-3312320133210211-2030122232212023-1001030010320111-3022302301331001-2323032100331132-2010110210122020-0201113132211302-3312133310130333"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-014.md#canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-1321123213122322-2300323310330011-2311022323013232-1211200021210331-1300230102122213-3232311012121132-3200230033102302-0101121202011222): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303): complete subsection reference.

<a id="canonical-1123033031232300-3020103321120122-1013221100132302-0100132333113303-2121000222132000-3300012003012210-3030310123112332-3232211122221212"></a>

<a id="canonical-3210312010023002-2021203221310202-0211312320331133-2112102220011120-3003000333103001-1121100023320030-1001100021200201-3320333200131220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231): complete subsection reference.

<a id="canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-013.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-014.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0101330300332111-3200311310020232-0211320203220012-2200310311210202-0303013301001201-0210133302110011-3210022021100103-3222301131102101"></a>

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
