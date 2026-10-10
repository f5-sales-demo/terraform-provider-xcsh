---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3112313021221031-1311101222201313-3012222302033102-3313131003232023-3003033200000321-0021003123101102-3112130022202103-0002231030132303"></a>

## `openstack.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3123030311230312-2223203313103321-0310121332230121-2113331310013222-2030222322000320-0102100011323222-1132132130102321-1032311320030330"></a>

<a id="canonical-0100331232200311-2132020033201020-2021213103211022-3322033020333310-0121323232333130-1232033122323113-1210321313001111-2030313120203110"></a>

## `openstack.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1031033000212131-0003103011013322-3203202330303102-1322003020322232-0121121222121110-1131231010323111-0233310102212033-3010300032212233"></a>

<a id="canonical-1122120031333332-0311331300011211-3131012323123320-3022331223210112-1121110213103211-3021300132102133-3212231231110331-1310301003202002"></a>

## `openstack.not_managed.node_list.interface_list.bond_interface.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0313313111010312-1332003030032011-3332132101103332-0030123133323310-0100012132121011-0121233023100132-1102032120312012-1021002230013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- openstack.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1110120221223333-3301101230330102-0200202221010312-0210031211200101-1322001301102122-1111213103010221-0211011032222210-2231033332322332"></a>

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

<a id="canonical-3102223220023321-0111220332033310-1212222201122120-0302303033312332-0302021232311303-3210110002003212-0133322031022003-0032123012223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133)
- openstack.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1013111110012220-0132231333203211-2213011220100303-2331122310223222-0010011221020333-2333111121330022-3311213322301303-2033111223133110"></a>

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

<a id="canonical-3020012312020120-0233113300133112-1213111010210020-1120320113303103-2201303213113130-2203001033310020-1312101203201203-0101201321312302"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-1101300022010120-1030320023013120-0131222021330102-3322303012220131-0233231103133130-1121312332102030-2202000102200101-3210103300011033"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3230113312212313-1330233220212100-3203311232020130-3012012231200201-3230331132330213-1320323302203122-0310130112001121-3002210101233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3123020230132302-0312211133100030-3102303011030233-3330030212022012-1122302102022131-2003203012231020-2203300003120100-1132301122333331"></a>

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

<a id="canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2130113031122220-2112031113002113-2312332133031122-1212112200030200-2031330012313232-2211201133333012-1330011323312031-2112330132000030"></a>

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

<a id="canonical-0122130131121222-1303103012020302-3003330223303302-2032102221030330-2020333311030003-3033033003202203-3032032222223020-0000101011231311"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-2302311110230322-2331300112102023-1310333113223023-3101113100210310-2121020330213321-2023003011112013-0022332313112100-2102032032111321): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3321233201103332-2232003203311011-3323102233123000-3133102222202121-1313333012121122-3333000030031300-3020011013332300-0332303113201033): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210): complete subsection reference.

<a id="canonical-1220101221202033-3303012120332230-2222202203120110-3112001301231311-1021003203223333-1202001230302331-1031101023301203-2011133032111313"></a>

<a id="canonical-3201230011030022-2103101233202310-3113030332210011-2133131131203101-1230213312102210-0330113030023110-0102323100201231-1222030101220302"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0323031113123031-1112033003333333-0200301133123221-1211121010002222-2112323010231112-2223020312001312-3013200000002022-0320120033223111"></a>

<a id="canonical-3333133002332011-1231003312021333-3321033213102323-3031202013212332-2320002103011202-3210130233020322-2223031220303310-1130210121102011"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-2031202130120130-3313113302311303-3313032330222000-3222123212121221-1122131300323113-0111313233133023-2323232012022323-0133302123301321): complete subsection reference.

<a id="canonical-2302311110230322-2331300112102023-1310333113223023-3101113100210310-2121020330213321-2023003011112013-0022332313112100-2102032032111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0311003302301031-0320002032011110-0022133013201111-1310203100033332-0033312301211112-1223132323122200-2101123032113022-0210101310230310"></a>

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

<a id="canonical-3321233201103332-2232003203311011-3323102233123000-3133102222202121-1313333012121122-3333000030031300-3020011013332300-0332303113201033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3231221111001022-2123030000222011-2033020311303331-3300002222303230-2210313133132012-0132212103003301-1010012203001112-0013002321010031"></a>

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

<a id="canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1203223323032001-1011113333003031-0222031202310323-2132333332303230-3221222212223011-3233033101100210-2103200133233322-2100300002201302"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2333023113231323-0311030321011213-0231032302201321-2031031001130131-1200031011333223-1000323122202201-2200003133132102-1112000003221010"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3002130012200132-2003023102112232-3302333222133201-2021332021133311-3230303111012110-0031323213301112-2300312302133111-2202312330332122"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0311103220033302-1321303313321000-3333111320212321-1333111102013122-2130030312301132-1013331112100021-2321310233130000-1200032100002300"></a>

<a id="canonical-2332030020313102-3323320312223022-0320213221323222-2331000133230302-1011320231121332-3011222200002033-1102110223110012-2301003310103133"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2030321121302210-1132311132101102-1332201022213230-1000023223002301-1233320130001203-0101122001212122-2012003123113200-1221130033031333): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-1120122231310133-2021022122030222-2132203033012332-3333021233011031-3012221110122212-2022113012013020-3121302100021232-1211030212021330): complete subsection reference.

<a id="canonical-2120120110021120-0031003011011221-0320302130012333-2010200223010012-2213220200130121-3022312212132231-1011212220223131-0032323020310311"></a>

<a id="canonical-0113102122011032-2211332320221312-2330312232313110-2332031303231230-2023033211212100-3122132011132223-3211013020321311-3201131010333133"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1103301303031133-1323232331023211-2001010332302133-2221223112331110-3123111030312020-3131022033303121-1213021021103310-1131112323332310"></a>

<a id="canonical-0302210300032300-3201220310011033-0023003010303331-3301102003112033-1201210302322330-2231311223031022-3311322102200220-1330030000203210"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-016.md#canonical-2330021230203122-3212121232232300-1003313223311032-2010132213210202-1032231111212023-1111213002222030-2200023232201132-3112020313023310): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-016.md#canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023): complete subsection reference.

<a id="canonical-2030321121302210-1132311132101102-1332201022213230-1000023223002301-1233320130001203-0101122001212122-2012003123113200-1221130033031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0201220113110120-2010020113311330-3033322333203210-0002122122020203-0113032002103230-0213131213210030-0223312033213232-2122102032202331"></a>

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

<a id="canonical-1120122231310133-2021022122030222-2132203033012332-3333021233011031-3012221110122212-2022113012013020-3121302100021232-1211030212021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2303020121330123-1001110321201201-0131200322221002-3021230222033000-1002130031300313-3303100302020113-2313223213032232-0122310131301033"></a>

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

<a id="canonical-2330021230203122-3212121232232300-1003313223311032-2010132213210202-1032231111212023-1111213002222030-2200023232201132-3112020313023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2312000301221222-2312221230331123-1022331233110221-1310222323102303-2133131201130122-0321122113313220-2323330212003120-0330123012122010"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3030012111132100-1000202032302321-3221033203112003-2032230202202231-0023302011301111-0210330010213021-0211130022313220-3331103001330112"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1023021121220321-1232100213001200-3230232202330130-1331200330333000-0301220322311231-0213312012211312-3113023212220321-3120033313131230"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3200132320001310-3022303223303322-2132012002201003-2023320330021311-0023100130330013-3332312203213001-0122222203202302-2220201221212111"></a>

<a id="canonical-1100120002122331-2211101300000001-0011032021301022-3123200013222330-2330232203111110-3121312002223132-2232310311122030-2023210332330123"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1202001113232012-3112302030211222-3122230032203013-1020030021220131-2332010330132002-1031330320011033-0200302133210310-2332301130102123"></a>

<a id="canonical-0310320323022300-1211201031321110-0313200301000201-0012100103121130-2031010012200222-3232320311323311-1333233003221023-2213223022323022"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2233032222111100-2020101101123322-1233331313323223-2101330211233213-3102100023130203-2221301312112311-2002220103231000-3101020302332023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-3133133322202103-1203031302033021-0330011121223213-0000033001010221-1321011031131113-3222333101002312-3321002223233302-3121030131321210)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2331020133330021-0331230100331001-3003331130133133-3213222013011132-1120003311312011-3013332201031313-3130022333331322-1022312002322202"></a>

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

<a id="canonical-2031202130120130-3313113302311303-3313032330222000-3222123212121221-1122131300323113-0111313233133023-2323232012022323-0133302123301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230)
- openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2121302300323232-3033312130231012-0302301130123033-3002313201003320-0100311003321311-0220313113113133-3010231022031203-0123000310113102"></a>

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

<a id="canonical-1013303100300203-2332133300120112-0330110220002133-2102112013312230-2021323031232131-2232230131130232-1112230113230110-1321313212033012"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-0303030333301303-2233003002011332-1121123311122212-0330302012033210-1231210301221123-1122123032322210-0121301101320122-0123303113311012"></a>

#### `openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1330203230212300-2103311130233331-1103133013123230-3202300010230033-3332211103233312-2123113330001101-0223131232213221-1022112130113323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3033113113133001-0230321130230020-2202231302303211-0110130213311233-0303223313030002-2220230203330100-1110313332023103-1230231320221112"></a>

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

<a id="canonical-3312221203131120-0210200202121222-0030103022032112-3031201321010002-1102303310023220-2102003330210021-1221103002212301-1322020210301332"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1303032220333330-0302212301311110-3202110100113311-2120011233202130-0301303333320311-0133011232033033-0000202110023030-3012000102001001"></a>

#### `openstack.not_managed.node_list.interface_list.ethernet_interface.device` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2000022321113303-3202203302113123-3013320321100231-1032303101301131-3201303210223302-0001103132013110-1331133133002122-1021000223202310"></a>

<a id="canonical-0230322211102303-2110133302112233-1302211100001323-1030221201321220-1331021300230121-2022122110033233-2033312130310301-2032333013232021"></a>

#### `openstack.not_managed.node_list.interface_list.ethernet_interface.mac` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3003013202210011-1321002233223101-2332231012022101-2203013101213032-2213033010223113-2202102312001033-0030002023031130-1000330332113201"></a>

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

<a id="canonical-1022302021212312-1321201031100331-0033211031000112-1302202100220010-0113323203212110-0310013130130210-2313323013030002-2221000211221133"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-016.md#canonical-3101302123110213-2302210111232033-2033131103122200-2313321121312001-1300002132312021-2213200322300011-2202312130002102-0311220131010110): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001): complete subsection reference.

<a id="canonical-3101302123110213-2302210111232033-2033131103122200-2313321121312001-1300002132312021-2213200322300011-2202312130002102-0311220131010110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-3210030021330003-0313101311311301-0103202331332101-3030032112213221-2332300102232213-0012322002013223-0020200112003213-3303301013111120"></a>

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

<a id="canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-3103003233133202-0230112212301131-1200331120211210-1030300121010100-1313212020301132-1001132320211220-3200303301220311-3131100213332101"></a>

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

<a id="canonical-2323101120210121-3310003213101110-3333102001321302-2120331201212223-1010020010103320-0001230312233020-3032333311112121-2020112020032110"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232): complete subsection reference.

<a id="canonical-2100321022021333-2232221230133232-1303120113100102-0310123302320332-1101301002023300-2300103011300211-1210221310311310-0213112202331103"></a>

<a id="canonical-1002010111032013-2022312232232013-3112121211230102-1112012302211031-0213230200020230-3123111113222033-3320223133102311-3033122112303030"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201): complete subsection reference.

<a id="canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1332311200303030-1121112303200321-0100222120103230-3011333032000122-0332033123113031-0020232301031112-0013230213102323-0310010200121210"></a>

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

<a id="canonical-2111032013021103-3230102322013003-1003302032013133-0203101223310110-2330210123210303-3222320012101201-1111002103111112-0333222121321013"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-016.md#canonical-0333223122131111-1321002232230113-2030111233101131-1013020301312013-2012011332223310-0030232101132133-1330310000103200-3312112110323330): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132): complete subsection reference.

<a id="canonical-0333223122131111-1321002232230113-2030111233101131-1013020301312013-2012011332223310-0030232101132133-1330310000103200-3312112110323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0023120322021322-2301300120113302-3212123002322202-1111230010113011-3333211331200212-0120223102003213-3232011232213322-1130332023222010"></a>

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

<a id="canonical-3321322123221003-0312223223331323-0212013321010023-2110033300031202-0031331122231021-0211201312031320-1202021301131213-3332230013030201"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-2003332233010223-1303333113232111-3131203133302320-1000322102100232-0032313200301313-0122030231133300-3223133101330332-3202113212132203"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-3213132210132301-3030013330110303-3323300021131232-2320300321001201-3000332002000301-2001112200100313-3303100012011131-2233321311330011"></a>

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

<a id="canonical-3111211122312311-2003311013123232-3032210110332311-2111300232201102-1023131133000313-0003223010300221-2213131120023102-0311002231120321"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-1303011120221230-2112022112312102-1330302113320322-3131103110100203-1211001300130331-0021223003300012-2103220233131331-3310332000223231"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3120101230221223-0101223001131310-2100002221320213-1031101332220310-2313332131011122-3332111233323300-0021100331123221-0122203023222001): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3110023203302313-2030221003021200-2202333000122302-3312000300223033-1122032220211022-1103220201110333-0223021213320223-0222011331320020): complete subsection reference.

<a id="canonical-3120101230221223-0101223001131310-2100002221320213-1031101332220310-2313332131011122-3332111233323300-0021100331123221-0122203023222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2000330102310302-1320023302312302-0002212332003300-0002332132301212-3021130223023002-0321100310323231-3222232012132210-1033023103033330"></a>

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

<a id="canonical-3110023203302313-2030221003021200-2202333000122302-3312000300223033-1122032220211022-1103220201110333-0223021213320223-0222011331320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-1012130020200311-2131100303013122-1233120222123322-3203001313322022-1233311111201112-1102101002031100-0102211233011131-3233003031223232)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2331332001033010-0212300302211000-2100120110230030-2220033123310130-1013011323001003-1111201310021100-3323201332011122-1032002101323200"></a>

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

<a id="canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0020332200200101-1010113111101103-2103113201300123-3310002133003023-2012321302333303-1002112223032031-1202112023210013-1002123011213011"></a>

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

<a id="canonical-1113010132120131-0320133101222200-3020103112212131-2120213012121032-0002033223221000-1131111312313333-0312113020312132-3030133113213330"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-0101230031010110-0202202331010300-1113112001210101-1123203121202331-3201211011322323-2110320010212110-3121311012213011-2311130132001030): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-3223222203121003-3022102222321010-2113202330312222-0022230323211122-0003003133231000-1222121020020323-3312000002322220-2102121211130103): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203): complete subsection reference.

<a id="canonical-1220331032200230-3130020022103033-0123122322320031-0332000311210212-0310332331021113-0101003210013123-0133330211221332-3001313010032222"></a>

<a id="canonical-0001113211022302-3230313132332002-0323022233203211-2003123011021300-1111313330011331-1022013013110313-3012123300010112-1202103102232221"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-1233203333000302-3323203123130330-1302220232001130-0023020001111022-0120031233220223-0313230211120030-3332301111301311-1122212112311013): complete subsection reference.

<a id="canonical-0101230031010110-0202202331010300-1113112001210101-1123203121202331-3201211011322323-2110320010212110-3121311012213011-2311130132001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3010230202023302-3103322332031021-1232010102110303-2312230230323303-1220033221030000-1103133012221223-2303110132201122-2001222033232321"></a>

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

<a id="canonical-3223222203121003-3022102222321010-2113202330312222-0022230323211122-0003003133231000-1222121020020323-3312000002322220-2102121211130103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3130103331300132-3001111001221033-1200101220221100-1030232201320213-3333030030011110-1201331312203022-1313220031321123-0012010300330022"></a>

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

<a id="canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0320013031003230-1101202100321100-0332102303223003-2232102202200302-1223130223110332-1210212311132210-3210233210202232-3303013203211001"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0220010300001130-2220300121200233-0300213033231031-0011130111011311-1012233011020022-1133320232313211-1013311100302103-2132110012000301"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0130122133331210-0130320303213123-2131301033032010-0300113313033003-3103212222130103-0310022211010002-0223302303131000-1201030220232213"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2222302133330120-0021112223212012-0011123132202022-1323200300132231-3202100112323321-3312022030133100-1320022203222021-0231221201222321"></a>

<a id="canonical-3313232331221303-1103230132113301-0313012300323003-0011101031312112-2122103211112303-3203030112113322-3332112010020003-3112100222311122"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-016.md#canonical-1120202130022001-1221101210333031-2113300303310313-0220131201023200-1333212323233323-2103331103132212-3112303330300303-0233011131222120): complete subsection reference.

<a id="canonical-1120202130022001-1221101210333031-2113300303310313-0220131201023200-1333212323233323-2103331103132212-3112303330300303-0233011131222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-2310133332320111-1332331321133203-3121313212302033-1112011133002223-0230123210010001-3323032113330211-0320232231313102-2100301211202203)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1313021222022102-3100330011020110-1223023200133020-2202021321131332-2330312322303332-0120100001122221-2032231133020012-2322030011100220"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2230121312220121-2031211200100333-2122010011021333-3122301201200203-0200333232013021-1213010003203232-1022132233231111-0320312120002003"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-1322331321113333-2100030022230031-3212013330230023-0011213201121310-0132111313131213-1201203103223332-0022232132133122-3013020332022022"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0303021133000203-3031111023133023-3230110100133212-0110110133303031-2000232301133100-0301013332113101-3210011220123131-1121022230311313"></a>

<a id="canonical-0231030121032303-0313330031321123-1230120020301333-0010112110013202-0322102001301022-0023021201331322-2002112003131302-0113101112013100"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1233203333000302-3323203123130330-1302220232001130-0023020001111022-0120031233220223-0313230211120030-3332301111301311-1122212112311013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-3132231233332132-2102122213221230-3313013002311202-1203333210031313-0021320121131030-3023310132121323-2213202102303030-3331333212012001)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-3000301321230010-3111332131201213-2110232013333201-2132212120133331-2302031003333120-3321000300302311-0013231133211303-3031112320112201)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3133120020122330-0222011130310000-3212301032022101-2121112021230000-0301230311331301-3133111101210033-0103310100301100-1022210320122330"></a>

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

<a id="canonical-0112313210013032-1010131013003202-0113101122300213-2201011203211321-3220000321131213-1333223100022033-3102213220011132-2313013010122320"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-3212133303011123-0222332100010122-2002210102102313-3220231101303111-0220030012103312-3203121120112033-3310333030123121-0000331222120100"></a>

#### `openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-3101113233000123-3211233033203031-3230112121131320-3021032013030320-2211023113300001-0002133320033300-3100233023020320-2021301303200202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.monitor

<a id="canonical-3002020310100322-0033123131303030-2321312302001100-0203032321210222-2300100110000330-3202120310033003-1012333101323122-1133322010100220"></a>

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

<a id="canonical-3222123322300001-0110133122012110-0111313122300121-0010031303320211-3130112222100330-1130133233002013-2230032311102211-0211232000012200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0102301002233113-3313121101120203-3002200232123313-1111001300200222-1210001200100320-1131220101312010-1313200230210002-3031011323211133"></a>

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

<a id="canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.network_option

<a id="canonical-1221022311300321-1302203013101332-2003330120001330-2210102123031300-2033332021210133-1312230133210033-1313223301031113-2023220310201111"></a>

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

<a id="canonical-3311312311300223-0132330103031000-0013213011131023-3121131102003113-3321002200132201-1220231230311321-3220113200021132-1212230322331033"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-016.md#canonical-0032303011130121-0221223212220232-3032001112003222-3210312230001020-0123130031231220-0301113022022211-0111203021013101-2121013332011302): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-016.md#canonical-1121221200221011-0102213201120211-3113331233332011-2121300111333010-1310331032212121-3100032230120321-2101221333230302-3200112330103012): complete subsection reference.

<a id="canonical-0032303011130121-0221223212220232-3032001112003222-3210312230001020-0123130031231220-0301113022022211-0111203021013101-2121013332011302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0311002312212122-3111320231110010-3232231201222130-1103301232212110-2013212003032121-3320230210333001-2113301023322113-3203113021133112"></a>

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

<a id="canonical-1121221200221011-0102213201120211-3113331233332011-2121300111333010-1310331032212121-3100032230120321-2101221333230302-3200112330103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010)
- openstack.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2022212123101203-1121021231131322-0122210030311202-1302000201333301-3330020300012222-0203210321012103-3203212000323323-3201023032333012"></a>

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

<a id="canonical-2332311111121233-1023200032232001-3122133030130130-1123313102323230-0200221122211002-3230133312201011-0122132200032030-0013213232231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3301210213120202-3113030232301101-0211310101031210-2101022233211000-2312113233013010-2111223012210331-2032032013032200-0003210012200330"></a>

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

<a id="canonical-2110313201023123-2310200020000232-1311232001030303-2222100103031323-0211313131023030-2300033103012320-1233022303303331-0011200211020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3222332033122201-3211302030330232-3333123003212103-1021003121110323-3302101320112313-1020213222020001-0021013103130200-1210330011210311"></a>

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

<a id="canonical-0210012301212222-0111212112333310-1312311313231131-0323100032010230-0030031102013120-1311123222203330-0103003211100302-2202220211100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1221110132120232-3011130131213201-1200322023322102-2122000120133233-3322100003323301-1232222011133203-0330221211011300-3032213220230331"></a>

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

<a id="canonical-2031330202202022-3331000031001100-1200313120202122-3313120233311233-3322331012221303-0023120222112321-1130321210023233-3011313203321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0130231101201023-0003213102020302-2220300023020211-0013323201310120-1313232332022233-3310130130001111-1102230113330032-3031221321100313"></a>

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

<a id="canonical-2112231302320012-3123103320233130-1022202103132302-0121321121313211-3111011202010200-2222311200010133-3021032133223331-0323330130002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.static_ip

<a id="canonical-1033022200100312-1131212220101123-1311023230320222-2111311033110121-0111103202000321-1130210300001223-3110111203333001-2332211200212120"></a>

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

<a id="canonical-3133131112002020-2130102231103030-3001321121231032-2033311303012022-1220010002310213-3212333001322022-1030110133110001-0212012100333221"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1023320000320110-1100023312013200-1000323132202103-1211332302030031-0010233223120331-3313313332100333-2113222230110211-2120121220011111"></a>

#### `openstack.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2320202312331301-3012123122200230-2031201020030220-1020102132310320-0332130132322201-2301030030223102-0031332030132030-1201301131221131"></a>

<a id="canonical-0030002320320222-3300001111123212-1100311102212102-2303222120220233-0110302120011003-3303010122220313-2121301120133020-3010330200212122"></a>

#### `openstack.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3213303211031123-1033212100220232-1112222321023010-3320233303212013-3021002323320020-1332032012021210-3223121013233210-0213013110101131"></a>

<a id="canonical-2010023003231021-2211103112112102-0111200203303131-3003001002100203-3000000032012101-0302122232233230-2221213131211322-2013101131011003"></a>

#### `openstack.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3332220323123022-1031200022030110-1231303222230020-0010210223021323-0032333132312202-2311223210223011-0303121031303013-0211122122203300"></a>

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

<a id="canonical-3303011033102311-2121002323300102-3221130032330101-3100201021231301-3331032133333213-0232303330322310-2330222333000302-1120321311123203"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-1322213312132232-1203013222023130-2031122231033203-3120003000321000-1213123021002112-1133013012331102-2010000101223331-2200312331302000): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231): complete subsection reference.

<a id="canonical-1322213312132232-1203013222023130-2031122231033203-3120003000321000-1213123021002112-1133013012331102-2010000101223331-2200312331302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-3101312120002222-0110222110021310-0310330231222212-3130110231020101-3120200121113201-0333102202030213-2002111232322230-3102303021030131"></a>

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

<a id="canonical-2131332223120123-1331310230223021-0120223131022300-2202213220223201-0231211222203322-3102001103202322-1311302023213133-2303012332013102"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-3303301300031012-3011012330020213-1232231133001321-1220133023320312-2033033133013200-1311021000022211-1130303001110003-0332221211303313"></a>

#### `openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3023020130101312-2030330013132000-0020030020300102-3131211312001203-2222322302130311-2222233121220012-3003330123213000-0031033102330231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1030122030003320-3322100330010000-2103322223111332-1322232210221331-1103002133132333-2130120302000313-0233120220103232-2213131300130033"></a>

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

<a id="canonical-3031310221302023-0030230221001033-0111132200312100-0002000022010002-1022123320210320-0121002013200000-1020300220013020-2012311221313220"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-2313303231101001-3223111110310123-1033310223001301-0002021211101300-1132120110031000-2200232203300032-1123103002123320-3233322023132122"></a>

#### `openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1012110020311330-0121213003011323-3210103220212231-2033323111222120-2333221000013233-2332322332012120-2132012312230233-0121030210233000"></a>

<a id="canonical-2112000031222300-0231203221021132-1313100201002130-1322221033303330-0200331110012212-2212231031133002-0223120223302000-3200331121010011"></a>

#### `openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0312231111001013-1001022211313010-2310102201011112-2320211203200200-1201023210031322-0023113321330330-3231103131013302-2103122031133303"></a>

<a id="canonical-3033030211333203-3201200311132013-2012133113130021-0133210022212211-0120223023022002-0301101213112200-1102113310312231-1122301132100232"></a>

#### `openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0020330021133011-0122223222303323-2332322122302031-0103032300320231-1322120221332020-0002202212303101-1102001132023322-2011031201011130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2100130333300131-3300313233000132-0221033331031333-1000032322313100-0131132022003223-0202302311013122-3102233123200312-2300313131022101"></a>

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

<a id="canonical-2302002311231313-3122330213201221-3003202233311033-3121223201111010-1121201201031222-1100112121330102-0031113123201320-1031210231001310"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-3132202332331122-2000202131213103-2300310301232103-0221300332210330-1122232311230131-2131030220120321-0002130002013233-2212121102232312"></a>

#### `openstack.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0333011012111302-1023311230112221-3200132133332113-0311312112211303-0122311202212130-1023333303312000-3102321211323202-3101012010202000"></a>

<a id="canonical-3102220103122121-0330123123031120-2311201100211022-3121031332211111-0213003020322012-2200213012002301-2320133002333320-0302200220030211"></a>

#### `openstack.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- performance_enhancement_mode

<a id="canonical-3222313202323110-0220121122313311-2112010132221321-2303210302002101-3211220111320132-1310211311100123-2122213331100323-0021000321222211"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2202023023133233-0022113033303322-1010210101112131-1130120130002003-0312232302211312-2313230221310312-0212311112212223-0330332023321013"></a>

### Direct properties for `performance_enhancement_mode`

- [perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320): complete subsection reference.

- [perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312): complete subsection reference.

<a id="canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2120233022030200-2010331323303213-1223333001220002-3033033032233211-3301210133311332-1212032100300102-3230011020220120-2003313013110201"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Additional upstream details:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2031200131001301-1303132130032200-2333303132002012-3101011102000001-3030122011330011-0303011311210233-0133310010212031-1211300130032122"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l3_enhanced`

- [jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-1112021030022202-0232100133302322-3102000223021212-2321131233012100-3220003032033132-3202333331002300-3120120222132213-2103200001332202): complete subsection reference.

- [no_jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-2133331031330223-1233231002111321-0301013031120330-0031202100130012-0301322330210123-0132031223223212-2211212012302112-3111223202313311): complete subsection reference.

<a id="canonical-1112021030022202-0232100133302322-3102000223021212-2321131233012100-3220003032033132-3202333331002300-3120120222132213-2103200001332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2210130332023122-0011311033212320-3121232113323113-2233230021130203-3110232022321233-0323323133102201-3002101110330111-1023103323012200"></a>

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
jumbo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133331031330223-1233231002111321-0301013031120330-0031202100130012-0301322330210123-0132031223223212-2211212012302112-3111223202313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3312312232232002-0001300013223113-0011101121302323-0123310030003312-0200203123300011-1132222102120330-3230231102310132-0021110010222320)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-1033102131303321-2000000100212310-2303201003313222-2212301021213111-0012133301000312-2112103203212333-0112003132300301-1103212232320212"></a>

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
no_jumbo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-1012202230032201-2331310300300132-3321111301013300-1111111131301002-3132321303200003-1032122310212220-2221233330000211-2210332333221031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Additional upstream details:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1323013220003321-1220013311332132-3211302021332111-2020130212132121-3101321013021313-1303313023031320-2100211210313101-0130102333112210"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l7_enhanced`

- [jumbo_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-2032323233303102-0031210333001023-1212220111301003-3003133001012120-1302303322121131-3030311211121320-2313000212102310-1213102112001122): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0033222113030132-2031300202112222-0023320013221202-1000232133103101-1210312200013311-0131122203331020-3300022233010021-0201021103003133): complete subsection reference.

<a id="canonical-2032323233303102-0031210333001023-1212220111301003-3003133001012120-1302303322121131-3030311211121320-2313000212102310-1213102112001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3111210213011001-0122233221230302-1223203231122302-2130011203231320-3112312120131110-0020133012132311-2120032223333231-2111000200200032"></a>

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
jumbo_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033222113030132-2031300202112222-0023320013221202-1000232133103101-1210312200013311-0131122203331020-3300022233010021-0201021103003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-0303120011111203-0120313102113202-1232323203231220-2303211023320102-1130102203310213-3200020102102001-1121031130131032-0202223232203312)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1002220121023331-3110130323330202-3111030031130111-1020100132011120-0333333003321031-2120111223223223-0302112132200331-2230011313013203"></a>

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
jumbo_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132333020000033-3121313020113130-1113011323131231-1303130201313202-3331202201110330-3232030020101012-0011123211010231-1301321320123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_adn` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- private_adn

<a id="canonical-3001100011332303-2231203223000211-1230200300233313-2012022223033331-1002001201232231-1312011231123311-2020311101200103-0010032010320220"></a>

Type: `"object"`. single nested block, Optional.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("private_adn")}
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
private_adn {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123000231332122-2100303210002320-1130130223203132-0213000212002223-1123013013222030-0130110032021010-1220232212130210-2332230300231131"></a>

### Direct properties for `private_adn`

<a id="canonical-3332013313231212-1022231232201203-2121210130100131-0233132032220133-1333021033100333-2012332313132132-3332331001010123-3333301103230101"></a>

#### `private_adn.private_adn` property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `re_select` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- re_select

<a id="canonical-3101110123032132-0220020232313132-1303330213001120-1201313031301302-1121213131113310-1002022030113211-2022033012223103-3202010003233221"></a>

Type: `"object"`. single nested block, Optional.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("geo_proximity",
    "specific_geography"),
  validators.ConflictingObjectAttributes("geo_proximity",
    "specific_re"),
  validators.ConflictingObjectAttributes("specific_geography",
    "specific_re")}
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
  "x-ves-oneof-field-re_selection_choice": "[\"geo_proximity\", \"specific_geography\", \"specific_re\"]"
}
```

Terraform syntax:

```terraform
re_select {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111010032222333-1030323011220013-2020220312310001-2200110131200201-2321113233221201-3321010012310221-1231220211213102-1202220002322022"></a>

### Direct properties for `re_select`

- [geo_proximity](resources--securemesh_site_v2--reference--group-016.md#canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100): complete subsection reference.

<a id="canonical-1110021000312201-3310112332202210-2203130220120113-2020321232223022-2320113331312121-0310322130302231-0122333131013223-0312223032313330"></a>

<a id="canonical-1213132331320010-2302311302122022-2030211200100213-2220030023233211-2212202031111002-2111110123102030-1122022211111122-1233111123230020"></a>

#### `re_select.specific_geography` property

Type: `"string"`. Optional.

Geographic selection for the site's Regional Edge connections.

- [specific_re](resources--securemesh_site_v2--reference--group-016.md#canonical-1211312003122330-1331202120220330-1302000112320001-0232023111130032-0221001111330011-0220021320230210-3033030101123302-3320121132333331): complete subsection reference.

<a id="canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `re_select.geo_proximity` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- re_select.geo_proximity

<a id="canonical-3101000013030220-2131313232331012-3200020122110230-0013110120023003-3121301102020230-0130012322300332-1332321222231030-0110311332100102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for geo proximity.

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
geo_proximity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211312003122330-1331202120220330-1302000112320001-0232023111130032-0221001111330011-0220021320230210-3033030101123302-3320121132333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `re_select.specific_re` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132)
- re_select.specific_re

<a id="canonical-3223112100302232-0223330113130200-3020201331033102-3301022011320213-2233222113210102-3221320313123030-3220102201232020-0032211213013103"></a>

Type: `"object"`. single nested block, Optional.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

Receipt-pinned upstream constraints:

```json
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
specific_re {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023302120203120-3202200203110131-1210131100101213-2223312032230223-0220331322330103-1301100010330002-2230223213323131-3332301120122310"></a>

### Direct properties for `re_select.specific_re`

<a id="canonical-0331201030103001-1330202110331302-0222123113011303-1233223331112031-2001113321122301-1032321001221110-1110100022033113-1033123333031022"></a>

#### `re_select.specific_re.backup_re` property

Type: `"string"`. Optional.

Select backup RE for this site, cannot be the same as Primary RE.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1303002323130123-3012331323333310-2213203302023132-2233000212003230-3001202012132122-0011311001200101-0210111121300021-3103302213102331"></a>

<a id="canonical-0132003010303102-1030011211332133-0332100321321001-2000031210302013-0123211110132223-0112023102130111-3022012121301001-0101120030031201"></a>

#### `re_select.specific_re.primary_re` property

Type: `"string"`. Optional.

Primary RE Geography. Select primary RE for this site.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- segment_vrf

<a id="canonical-3230021233323013-2212321323223310-3232303201223222-0011210110010230-0233321201312301-3332031300130202-0022120223300213-0221022133112223"></a>

Type: `"object"`. list nested block, Optional.

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
segment_vrf {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221020121132032-0111122133233320-2201110313233000-1322000000033023-0011301013100320-2302123230010320-3020020311102002-0200232031312012"></a>

### Direct properties for `segment_vrf`

- [segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021): complete subsection reference.

- [segment_network](resources--securemesh_site_v2--reference--group-017.md#canonical-3110213102100123-2213032002131131-2001112022032210-1133200211232232-0332033311333113-3210132330223010-0302230001102032-0313121312022203): complete subsection reference.

<a id="canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- segment_vrf.segment_config

<a id="canonical-1121321000213231-2132012020102021-0011333001013023-0121310021021100-3311123331331231-1030310311010003-1030133002003332-3022011233330320"></a>

Type: `"object"`. single nested block, Optional.

Segment Network Configuration. Segment Network Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
segment_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122010003022320-1212220210112003-1122311121101100-0310331023033202-3223221213312012-0210103223301020-2011211312131132-1022223032202033"></a>

### Direct properties for `segment_vrf.segment_config`

<a id="canonical-1000220003133312-2202300201020122-3032233113120302-1012300132131221-1233310103101230-2310323222200100-1202120020211331-0301012202231313"></a>

#### `segment_vrf.segment_config.nameserver` property

Type: `"string"`. Optional.

Optional IPv4 DNS server to be used for name resolution.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [no_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-1330032102221030-2100130132333313-2133103331222101-0010033002132332-3100131030120203-2012301200112221-3122103032223220-1000131121111122): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-1122111130130121-1310322230113120-3320103131022213-2330113120211023-2213120101110331-0110301202123320-3210123023220122-3121012032012303): complete subsection reference.

<a id="canonical-0020121300313303-2110013112123120-1010221222031113-1213220021110310-2321013213320103-2203213010003212-3032121133132210-2331201231200130"></a>

<a id="canonical-0213002331202222-2013030322033103-3301321123000001-1213231001212002-3230113032300020-3030111333100111-2211302312213303-1033322121303123"></a>

#### `segment_vrf.segment_config.secondary_nameserver` property

Type: `"string"`. Optional.

Optional Secondary IPv4 DNS server to be used for name resolution.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022): complete subsection reference.

<a id="canonical-1330032102221030-2100130132333313-2133103331222101-0010033002132332-3100131030120203-2012301200112221-3122103032223220-1000131121111122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.no_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.no_static_routes

<a id="canonical-2312001031300222-1013111310221000-2011023223032021-0310023131113221-3322301220033301-2022210113332103-1000202023233132-1311312333121110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

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
no_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122111130130121-1310322230113120-3320103131022213-2330113120211023-2213120101110331-0110301202123320-3210123023220122-3121012032012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.no_v6_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.no_v6_static_routes

<a id="canonical-3120210021311230-1112123322101302-1233203222233211-1120002123310212-3203300023033333-2222332110232020-2011100233200233-2321123202212013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

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
no_v6_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.static_routes

<a id="canonical-3102132021022121-2112330102023031-3303212003111203-1022313103120323-2202213030000330-3201223130230133-3130322120020322-1103202103000332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020320011231330-2301231132223210-2312210330300020-1201333322110222-2123223123122130-3313221231022012-3223111130313023-3111320130232112"></a>

### Direct properties for `segment_vrf.segment_config.static_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311): complete subsection reference.

<a id="canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- segment_vrf.segment_config.static_routes.static_routes

<a id="canonical-2300223320301032-1212311322133003-1301033111133331-0011012211032223-3103022231332030-0201331201010131-1031332110230300-0031211323331113"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013213112201122-3020023013012221-3113323331001111-2001120122300100-1023220130312003-2232302010200121-3123301223120121-1303213120131013"></a>

### Direct properties for `segment_vrf.segment_config.static_routes.static_routes`

<a id="canonical-3013111120011010-2333333033002031-0022023221131002-3202311023120021-2331230321211100-3123300232303112-2321212122022300-1313130112321301"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [default_gateway](resources--securemesh_site_v2--reference--group-016.md#canonical-3212121032131230-3022010321123000-0312130021130002-2133123102020332-3311233020100301-3132031121022320-2102010111311230-0101000301331130): complete subsection reference.

<a id="canonical-2130301332322320-2003222322033122-0210301232223301-0121231231102202-0132320331123203-3221310132101011-3130230130322102-2233010031103030"></a>

<a id="canonical-3010111313213230-0103023000022013-1203222023010120-1333210302331113-2113030120013332-1120223022213310-1020320112100122-2333100212322213"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.ip_address` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213131301232031-2133021313210212-0310001311313323-3201203013331232-1302320002012133-0003122110131301-3201212203320101-0321323222021030"></a>

<a id="canonical-0120232023231111-0220132303013331-1030303313112030-1301020010120323-3013221111031332-1120210303321300-2131212002303311-0202120000022331"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023): complete subsection reference.

<a id="canonical-3212121032131230-3022010321123000-0312130021130002-2133123102020332-3311233020100301-3132031121022320-2102010111311230-0101000301331130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="canonical-0213122212210021-2311112221232101-3012331113102302-2212032011331320-1012330121123133-1031120212211102-1322033311023323-0231233111030202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="canonical-1320203221023300-0021031131002102-1021311003003020-3323323011211201-1021121210020133-1131111201132123-0032330123022032-1222301301331310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201022310112112-2321211033201000-1010301321023122-2213010132213133-2332301001222311-2020332123103310-0101000231001210-1333321212313201"></a>

### Direct properties for `segment_vrf.segment_config.static_routes.static_routes.node_interface`

- [list](resources--securemesh_site_v2--reference--group-016.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102): complete subsection reference.

<a id="canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="canonical-2103303232131212-3003030331112110-3300020230202223-0111330033021210-3101131302200201-1023102032001201-3022232021200213-3212333031002210"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220101330102133-0112312201232321-2301131001221211-1120321201010312-1311033202031102-2022123012321330-0323331222323333-0231013320033120"></a>

### Direct properties for `segment_vrf.segment_config.static_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-016.md#canonical-1123232233033330-1210330312203023-0122110200133222-0312023012313302-3333123331132221-2031133013030111-3100211011203313-3132310132103032): complete subsection reference.

<a id="canonical-3101000120200330-3201002301032330-1303323012133011-0110322021333030-0003131031000122-0203313311030201-0022121232230331-2130023222310212"></a>

<a id="canonical-3210121332110330-3211220232220012-1023030331010120-0330333132121111-0033120131202133-1310103323302330-2321330003300003-2103312310232333"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Optional.

Node. Node name on this site.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123232233033330-1210330312203023-0122110200133222-0312023012313302-3333123331132221-2031133013030111-3100211011203313-3132310132103032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2330103322203031-0103103110002132-3203111103232123-0302131133211202-0331201333113222-2030101130133030-1303203110102113-1231230111202311)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-2103023112031113-2310110120120230-2222310321031013-1100213230000333-2200212222210032-1212001131211320-3320033012321222-3313133011113023)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-016.md#canonical-2112323303131120-1220033330203231-3303200130131131-0210313120310332-1123303300112221-3332103330131331-3321132330102211-0001320321313102)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3022211220121201-2022013101133301-1233201333313120-0121211230012123-0120220333030001-2332232210030033-3010212323113030-3302220320233020"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2121032302321021-0020001221223020-3000203023132300-1200010231022321-3310011210000311-0302300332311121-3201332131103212-0313121213011002"></a>

### Direct properties for `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface`

<a id="canonical-3030031033011122-0331231131221032-1220311000131130-2213201220020030-2131330003110212-3310232331201032-0122330332222322-1031010003123233"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3211232013022131-1130010212233230-2111032110202120-3112320221221020-3021202111300301-1303202231113000-2320330332230232-3100000213302212"></a>

<a id="canonical-1131031231200333-3332013223211102-0110000023023010-1201013113112123-2001000332110302-2002003100020233-2231201112233310-2330231333332210"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1202233113223212-1211203112100200-1110202232130030-0323112112231330-3322200100100103-0020020102121120-0013331120313013-2001213211331023"></a>

<a id="canonical-2030303323111110-2000003231133010-0320132003233232-1310332032113130-1010100023030110-3030302133130033-1010212100120122-2320113302023330"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300321312233111-0230022131003202-2302131001302212-2011021310111001-3010321333102323-1303103322333213-0322230322321033-1110002300033010"></a>

<a id="canonical-0103330333123203-1331300033332002-3110302223301311-3230321022101010-0113113201013301-1322310201020332-2321011011213023-2021002232032233"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1011231133313212-2112201102222001-3321011131123333-0222112011221013-3313221333010202-3213221303302331-2103210211000220-0332110101121101"></a>

<a id="canonical-1013303321302200-2000022101030022-2112010230133300-3311200002200123-0301133032301132-2122313201113220-2321002301331021-0012321021231032"></a>

#### `segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_v6_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- segment_vrf.segment_config.static_v6_routes

<a id="canonical-0132220013132210-3203210302102232-3223110310330213-2010032111133012-1003320321111023-3032032123131122-0012231032103221-2302013111211211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Additional upstream details:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011212201110311-2012320003022021-2113220033101312-2302011120213021-1232201102301301-3332100030223013-2330201003320003-0213032302313333"></a>

### Direct properties for `segment_vrf.segment_config.static_v6_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312): complete subsection reference.
