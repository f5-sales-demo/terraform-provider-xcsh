---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-1002201022001012-3122312230232322-1221131303323001-1130322310223220-2022221102320120-2301322300212300-3332131001022330-0321100200303303): complete subsection reference.

<a id="canonical-0233202102020033-1212330233223200-2020231002031320-3201230001200133-3130333300122331-2021002303222132-1301133322003002-0131321231011012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-2303302312013222-0232311030120233-2323120300103103-2331021200312332-2003130023231031-1211123131120230-1320102303201020-1032311023023103"></a>

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

<a id="canonical-0233312221032103-0330200223321020-3012102221120133-0203220111022302-1222123221112023-1023301210020321-2230020010130022-3203110210300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-0103201302301301-2202212331213013-1032330113210332-1110232131123322-0031123132312123-0332023303311202-2022121023233103-3023100331222302"></a>

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

<a id="canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2321102210201002-2322033221222111-0130213010330320-2302022332223133-1000321220222101-2220311033323132-3033101120312132-2302030000023131"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1123121303112222-0011310302100202-1212012020131203-1320022110201122-2303333001101133-0301113033132113-1332302332102023-3301013113003011"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-2231032301331111-0213112113120322-0200112123003230-3001210221223211-1322230130230010-2202222302122020-2130010123330202-1210022323313023"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0313022201302210-1131313303302321-0122322333130100-3022332010000302-3332323332202232-1323010110102331-1023330333011101-1221322220033110"></a>

<a id="canonical-3003200311030023-1113122002302222-1201011231021000-1333313233122030-0220201233103133-1133133302213023-1133323101333221-2013131100102300"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-0130010302110212-0131103231112310-2322101003132013-2330210103200131-1111332231021010-0133011323303031-0302021032321011-0223211100220103): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-0002331300232323-0211130123223120-2222032213223230-1230103002113023-3011321101201103-1003031330000023-0103331013232000-3320111011123123): complete subsection reference.

<a id="canonical-1230323023030111-1122231201323301-2132000133312113-1203333201122112-0310030330200323-2202332222100033-3021111221100031-0110203333130121"></a>

<a id="canonical-0210203300030030-3010332103111230-1230130012133012-1312211003121211-3030210102230011-3131130000101113-0031322103023310-0220210330311230"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2130101101331120-0023223013033121-2132302111222303-2221030231323122-0003330313121320-2002133300322223-2010220000031103-0111120120003121"></a>

<a id="canonical-2303011323010111-0313121002222031-3321021213332023-3123331203312022-3013332223320000-1110130031302301-0033111330323230-2130021302022221"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-2103231133112020-1221332201203232-2031023033311002-3112320023131232-0231002302021110-1000311231121110-2303003311332213-3202311003332212): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-010.md#canonical-0313020210031222-2021210110201020-0120313322030010-0132100303232222-3103020303200013-0201100300201132-3000120303131031-1133113320332301): complete subsection reference.

<a id="canonical-0130010302110212-0131103231112310-2322101003132013-2330210103200131-1111332231021010-0133011323303031-0302021032321011-0223211100220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0213321211020201-1322232200333212-1001011300003211-3021123333123210-1230100223311231-1220210023031003-3112322132223000-2332223130320001"></a>

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

<a id="canonical-0002331300232323-0211130123223120-2222032213223230-1230103002113023-3011321101201103-1003031330000023-0103331013232000-3320111011123123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1110103311112102-1330210111332001-3123032113000232-0022102331223203-1321322032121001-2033011121113023-2002313101201203-0321231100221211"></a>

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

<a id="canonical-2103231133112020-1221332201203232-2031023033311002-3112320023131232-0231002302021110-1000311231121110-2303003311332213-3202311003332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2021100010213120-3132202320132302-3022010202030313-2103102101001120-3002202322022013-0032200122103303-2012210032230231-2030302311100022"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0003200032011222-2322121033222322-1322002321112201-1000130110101302-2030130020013330-3221020320233011-1002203121013222-3220132022302032"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-2313202111223230-3123130110011121-1223132011231013-0032130003300333-1022213323000103-2022012133013123-2322133211133032-1320010331030210"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1110211311123000-2002022033332020-0201131021020322-3103030011031132-0202121200020103-0310333313101232-3223200020012333-3310310021300033"></a>

<a id="canonical-1120211223232030-0211123332210011-0100101112032233-0000003030332003-1200210201322230-3032133323311132-1323211311223032-0020202312111023"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3223112332011102-3111313122232233-2033111021322212-3020133332320030-2210020112201121-1103212013102302-3313332123132130-0013112320010112"></a>

<a id="canonical-1301231032000300-1213000120302302-1302102323130100-1132013311321332-0131101000221111-2200022231223111-0002212013112313-1320113223122002"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0313020210031222-2021210110201020-0120313322030010-0132100303232222-3103020303200013-0201100300201132-3000120303131031-1133113320332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-3120300311211122-2001110033210003-1121122012012003-2321311333120302-2130030330221211-3103121120102130-0313210103010111-2102202203331303)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2122221300213021-2312113220103002-3032012312023230-3113022013000023-2122023112113221-1011131130012110-2030022303212130-2312212012232311"></a>

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

<a id="canonical-1002201022001012-3122312230232322-1221131303323001-1130322310223220-2022221102320120-2301322300212300-3332131001022330-0321100200303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-2320301323000311-2112211203102131-2000111011221233-3232301120203301-3223202000113120-0320021001233231-0201202332110231-1221031203102333)
- kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1220002312103301-0332330222111031-2021100211122020-1313323213123030-2322032112012220-1311310220231201-2233111321111023-0213103233203223"></a>

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

<a id="canonical-1110320223112212-1002132232021210-2023033211203013-0202100022031201-2201231322012213-3321300110121330-0131332213110223-1112132230213032"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-1331133030120112-3012003001023331-2222323313002223-2302202232211221-0122322032323021-2122102033313323-1222111320303033-1312301313111003"></a>

#### `kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-3230303333012202-1021300033312101-1113102213332132-0123230310002231-3301112312020310-0003210212213203-3321200333100110-1110210321322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-1220011232132322-0230322021232322-1011213201332202-1300113203123110-2101131131310322-2322012133102233-1032121302210103-1110332111131033"></a>

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

<a id="canonical-0311201323132002-0210311032132013-2321132210311102-0202323303222331-0333101110023010-2021322002310120-0333130302002030-3312120310101133"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-2222320322103023-2323202110322332-3333210201032102-2322332020100311-3101301222221101-0011002330302200-0330133223022333-0012120130233222"></a>

#### `kvm.not_managed.node_list.interface_list.ethernet_interface.device` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1230122033102213-1233121001020203-1013300300333322-3203212021002110-1231030211301311-0001303220033031-3030330103112312-1102201110330200"></a>

<a id="canonical-3102122313320120-0221322012113322-0221033012101221-0202312302320032-0333003003023222-3133230102131312-3001000223302121-2311123030222221"></a>

#### `kvm.not_managed.node_list.interface_list.ethernet_interface.mac` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1312013210011313-0331231003011200-3223322013122220-0230320111301211-2311313112312110-2023113212020210-0013222220323133-1021201332301111"></a>

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

<a id="canonical-1111130233322213-2200301013213200-2012210002022102-0220020210112213-1312301320223212-3121123031323031-2211032231033302-2022220310113132"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-010.md#canonical-0221210313131330-3032301220130211-2222103213131102-1112103123123000-0212113332302203-2231322101323331-1220312102233200-2120222030230132): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301): complete subsection reference.

<a id="canonical-0221210313131330-3032301220130211-2222103213131102-1112103123123000-0212113332302203-2231322101323331-1220312102233200-2120222030230132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-3113322011213002-0221100120301122-2301220133013201-0203211310013300-3010320132113031-2130201122002321-2023313030021321-2111031220222231"></a>

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

<a id="canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1021220221211201-1011112233303331-2212130003133221-1220011120320320-0003300333120000-1112032020300011-2010221311012021-1113122022032011"></a>

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

<a id="canonical-3301121032323310-3323202022323332-1320012302331301-1223323220103330-1203113102220012-1001133331302330-2111030300230011-3210223200312031"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020): complete subsection reference.

<a id="canonical-0113231123233323-1330120001121030-3203312020301000-1132110321231230-2131122100301333-2013222103310313-1222301011332130-3132032311302312"></a>

<a id="canonical-1101323033103230-2213123201230210-1010133112302111-0022302303223031-1003313202122102-2130202133121103-3121030032321113-2321211230000303"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313): complete subsection reference.

<a id="canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3201233223220331-2133233031213320-0322013012021201-1203301230211212-1332233012131313-3223113003122121-3003223233020302-0110201303320113"></a>

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

<a id="canonical-3033202232032220-3212120100122002-1220123000213030-2323001021233120-0210311331122211-3322120020102021-2133001123332120-3013323122310201"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-010.md#canonical-1221100330021131-2022001323213120-1103133223000320-0221112012023033-0033331300220113-1332332302112121-1113011110201332-2311323331332002): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-1121011130323011-0220303133201201-1032230122011103-0232302122122302-2021132010323000-0202212200330031-3020203030020030-0320030111011332): complete subsection reference.

<a id="canonical-1221100330021131-2022001323213120-1103133223000320-0221112012023033-0033331300220113-1332332302112121-1113011110201332-2311323331332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0003320000322022-0200130212131210-2130200021330203-0321200113111303-0310103312303023-2121333303223032-0232012112110000-3311020000311332"></a>

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

<a id="canonical-0231112020313130-2200220323013302-0231011110021323-0202003132311110-0223022322231122-3121212323022323-0112330022121021-3011131121012310"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-3313223013302303-3321232213030332-2233122200202322-0113213211211301-1313010023120002-3223313100112002-1301221021223021-2111011103013123"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1121011130323011-0220303133201201-1032230122011103-0232302122122302-2021132010323000-0202212200330031-3020203030020030-0320030111011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2333212301021012-2223122122101132-0332310330313320-1233222001003322-1202113002210013-0111112300033130-1010113230303212-3131232313313132"></a>

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

<a id="canonical-2113002001201330-1103021133133311-1200312122211232-3211103120112000-2203300000030022-1133230103100332-0123220232222310-0301323132322100"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-0312100310311223-3231331303332000-1222312221103021-1000100021101311-0312220313003131-3203222330232032-1000232213113210-0010332010333321"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-0123320230322210-3302131303100330-0113000220201023-1221123222023323-1100332312013222-3030212122131022-0210302123002021-1313313301211311): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1131030101330000-0020002312100200-0030330021030303-0233211113230110-2220322030010312-2222032220021211-1223211320012203-0110203132303320): complete subsection reference.

<a id="canonical-0123320230322210-3302131303100330-0113000220201023-1221123222023323-1100332312013222-3030212122131022-0210302123002021-1313313301211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-1121011130323011-0220303133201201-1032230122011103-0232302122122302-2021132010323000-0202212200330031-3020203030020030-0320030111011332)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2303122010100123-2213230223112311-2000333121031202-2021200011021333-3321303113120323-1133103010101223-2200220103200301-1123133110110032"></a>

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

<a id="canonical-1131030101330000-0020002312100200-0030330021030303-0233211113230110-2220322030010312-2222032220021211-1223211320012203-0110203132303320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0322013022033303-2012132220211023-2302333203313321-1002013320022110-1201110030002112-3332312100032323-1230111320220032-1301122003313020)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-1121011130323011-0220303133201201-1032230122011103-0232302122122302-2021132010323000-0202212200330031-3020203030020030-0320030111011332)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1113021003132302-3013222221033111-3300031200230123-1110111120012302-0201310002330231-3230231101203012-0030121203201031-0123132231321231"></a>

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

<a id="canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0023122012131023-2320303032223033-0003223012331213-1010020121103012-3313323331033030-1300121133322101-1023021033000020-3312131322130123"></a>

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

<a id="canonical-2223212102013222-3321312300320103-2113203110130102-0133321202001211-3331100102331332-0000012312231310-0022023321021001-1302033322312113"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-1312101132100331-1223123303033230-2323023132130210-1111130112303331-2213021330023301-2023023020211110-2030023032132313-0123332122321233): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-0023010133333303-3302310130122022-0200112103032132-1031123031301302-0232232001101103-2113110103303201-1211112010131022-1130030121132202): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-1231300310020122-0311021130200210-3211312333033111-2033302233021203-2222301000013312-0122100212302111-1302000100223000-1023101113011022): complete subsection reference.

<a id="canonical-0023312230120031-2312122003231032-1230201201310321-3220301013032101-3310332300310003-2123103200220100-2222202010113110-2032000030012323"></a>

<a id="canonical-3121233030231313-1111303013023030-0220312023330122-2010221312323012-1311201211123121-1133300303031210-0101233031031123-0322111310130130"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-3221310321212231-3303030331102303-3201322132131332-1133301323102030-3222322012030223-2102233322221231-2113130103130302-3031000231333021): complete subsection reference.

<a id="canonical-1312101132100331-1223123303033230-2323023132130210-1111130112303331-2213021330023301-2023023020211110-2030023032132313-0123332122321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2231020213023132-1230310233311130-3211113223200131-0133330132010312-3313213001121233-1211033322320132-3031000233302311-1101311201311033"></a>

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

<a id="canonical-0023010133333303-3302310130122022-0200112103032132-1031123031301302-0232232001101103-2113110103303201-1211112010131022-1130030121132202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0233212131202211-1022133331313332-0233331130313001-0112023020101300-3001323003330022-3133112100012002-3212111201321230-1021103110030321"></a>

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

<a id="canonical-1231300310020122-0311021130200210-3211312333033111-2033302233021203-2222301000013312-0122100212302111-1302000100223000-1023101113011022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0301211011221022-2323110222202023-3311202013112013-2032322302112320-3320002201001101-1123301301220003-3220102130300312-1313321103001312"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3111220030312113-2100110031031002-1301221323000100-1000222022120000-0100221021013311-2211321311313120-0223113033330021-2113123222301322"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-3333110022222000-3101213103123312-2232303133201121-1312203102200010-3101322310220000-3130122303120210-0121221331000333-3030111231230002"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3103012100000302-3201100011310120-3112212213011110-1030232311132200-0310232031130211-2022320202300223-2203030111131202-0131232203133223"></a>

<a id="canonical-0232030013201322-1302110302103120-1200010103211331-0223320333000310-1222113013133130-2102133220130231-0202232202101202-3312203302113221"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-3320220330210013-1123230020012323-1200231320033220-3020021323000032-3311313221002222-2100111123130322-0301030102313123-3310110220001102): complete subsection reference.

<a id="canonical-3320220330210013-1123230020012323-1200231320033220-3020021323000032-3311313221002222-2100111123130322-0301030102313123-3310110220001102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-1231300310020122-0311021130200210-3211312333033111-2033302233021203-2222301000013312-0122100212302111-1302000100223000-1023101113011022)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-3211322130021003-1033313030232000-0313233211011013-2301001323103011-2222301133101223-0110103311310200-0220211101321032-2311311332220322"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2323220121111210-2301130230003322-2213313230130000-1113133010221123-3303211213122302-2000011120120331-1230010201020130-2201013013100201"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0032133223321022-1111230133322012-0012322201300230-0030223001132132-3332001132222310-2030313222232112-1301130023210223-0312110322322220"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2321331232101000-0013210332100301-0332232002223103-2210333102013132-3211330212222111-0212131233223021-3023310011030111-3320032122001221"></a>

<a id="canonical-2311111013330102-1033221100020202-1231032332230301-0021131011131130-3133010002002202-3321100120022110-1101231310223202-2300122233300131"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3221310321212231-3303030331102303-3201322132131332-1133301323102030-3222322012030223-2102233322221231-2113130103130302-3031000231333021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0310013202310301-0312200201113322-3233203132331000-0010233300012231-2031210303032201-0300021222001302-2001010333101303-0130101022122311)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-010.md#canonical-0132213012122020-1001332030211202-3031311010020332-2133230330100233-3232220021210223-1111321231321122-0213011331103132-1212323202022313)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3101012031211302-1111011021301021-0302122110122332-3222012312033002-0332010332122003-3221020011300013-3322111202131130-3311323223220312"></a>

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

<a id="canonical-2021311321323022-2221123012133212-3131332310002213-2000003322332202-3232032221310213-2322013230300013-1212303223031103-2120120213010011"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-3231001210012111-2331003231333231-0020221000003311-2022012130313203-3020021032033110-1100101232331223-1320222113100120-0222320100232222"></a>

#### `kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-2012311020303313-1030231232300200-2221030123233323-1122220201031103-2332220103012331-0300232332201133-2300231200212201-2020333320021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.monitor

<a id="canonical-2231330201102021-1011121333302030-3112110101111110-0231103023312002-1202303213232012-2021113202211101-1300321232213312-2223010132300320"></a>

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

<a id="canonical-1100202122033020-2002013131322332-3123310302210202-3133301323002113-0313101011311302-3121103222332202-0200231131113331-2121223300012122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1202031321332020-2312333112230021-3121013331300323-1031213320001220-1232313323033000-3112101130203320-0322301320013211-3012010110013220"></a>

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

<a id="canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.network_option

<a id="canonical-3113223022120220-2230012231121230-3132023313320100-2333310012211010-2222313102313013-3203111122020110-3022300311331322-0133212313110212"></a>

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

<a id="canonical-1302033332300330-2230111232031102-2030122222302031-2302201230002032-3001111220101010-2223331203000121-0301112330003113-0101303313230330"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-010.md#canonical-2102132122313221-1212333033203320-1202112330232221-3311020032302033-0220313030220220-1330211303300200-1013303002102200-0221303033232131): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-010.md#canonical-3302030030303322-1232200100231101-0223313303101312-0110202111201331-0100031200130130-1313123232032330-3302011310111031-3103120312111231): complete subsection reference.

<a id="canonical-2102132122313221-1212333033203320-1202112330232221-3311020032302033-0220313030220220-1330211303300200-1013303002102200-0221303033232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113)
- kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1022123201111102-1220320300112102-2300232000330012-0113131122111211-2110323330201331-2012202313022331-2121001202021130-2331020232011211"></a>

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

<a id="canonical-3302030030303322-1232200100231101-0223313303101312-0110202111201331-0100031200130130-1313123232032330-3302011310111031-3103120312111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-2111012111300220-1101022333120303-1210013102220311-3130230133301023-2230000210021021-3003013233210230-3223322330002111-2123130321000113)
- kvm.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0313202122200211-1322130301021032-3322301301310321-2321203101110133-2003002022032300-1123130213103002-3002011331010300-3220231002230233"></a>

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

<a id="canonical-3123131131203320-0311213232102301-1111221021223032-0003223303123123-3221233231103221-0223301101112102-2323310000121122-0302333222323201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3222210110003221-0221220321100332-2210133331300030-2110113213101002-3321011011202113-1301210212100131-2311121031330032-0111003030002231"></a>

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

<a id="canonical-0311010103223322-1303203103333303-3322330303011033-3111120302320230-1100013130210201-0011001111033233-1112101120100212-2202332313302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0120021222113311-0330132010101000-0312212130311000-0320333312202003-3202012200003021-1030010102100102-2121202110301232-0301321213331210"></a>

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

<a id="canonical-3332312113013203-3301201000000220-2221200231232212-2203012012230202-3321023213020032-0233201012101203-0132330002212211-3013111113220233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-0021303033021112-1021201233101021-0201021301011213-1333120211330213-1301310121012233-2010302220333130-3331323202210122-1223303332032020"></a>

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

<a id="canonical-3013032133102011-2021301032001100-1323130231102113-1023120032112202-1012023001233332-3220030033113122-3220021230120310-3031131022131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3223031133301000-0203010212023313-0313332322032001-2010001021203200-1231020233210021-0330100223030103-0301203101312032-3031101201310020"></a>

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

<a id="canonical-0131210010111310-2321312120022211-1112020000212131-3230212132003223-0200222021213212-0100033002013220-1222130322202101-1321121211213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.static_ip

<a id="canonical-1301012203030212-2213211013222310-2220121330333200-1000131330322122-0011210132022223-1113130320110130-0133132221123311-0032022201211333"></a>

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

<a id="canonical-0330002030103003-3333331013232133-2211201022100101-2202311301232103-2020001000121120-1100113212031203-2210230130002231-1012113321211313"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.static_ip`

<a id="canonical-2300310033302221-2032033010233123-0013000203103021-3211103333011303-1300012213023222-1013221202311100-3230111301013223-0303312302301233"></a>

#### `kvm.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3113020321210011-0002101132123223-3010200320013220-1312303012022021-3202321020302102-3311203313003211-2003133220130021-1101301011121230"></a>

<a id="canonical-1120332103020313-1021120203332223-0010211303322232-1311232010030202-2333003001233210-2330322120222000-2233133113220020-0323312131330003"></a>

#### `kvm.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3230312301122221-1023213020001222-2103222112013032-3130123232020333-3013202113212213-3211301332223302-2103300231322001-0303131013032030"></a>

<a id="canonical-1332013210231202-3111303013120132-0203311123310032-2303113032212012-1031011111023101-2113200312033233-1112223000212031-3131020223212010"></a>

#### `kvm.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2103010032211231-0201210111212020-1302323311323123-0012332332030202-2130103032023212-0113300210303212-1233002221002031-1132223220202312"></a>

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

<a id="canonical-0322313123112132-1010110311231312-3112222033301111-3133030023030311-0120003231332313-3202023200313310-1010023312210123-1223301233333010"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3123032023110220-0022210303013101-1112222232003232-0221201232133311-3131222201113321-3012310302200133-3003203121223122-1203030200033132): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-3232333213213110-1002312103100310-3210213021012112-1023100133130021-2233333321201132-3303011231323313-2213333021010132-1013312133120313): complete subsection reference.

<a id="canonical-3123032023110220-0022210303013101-1112222232003232-0221201232133311-3131222201113321-3012310302200133-3003203121223122-1203030200033132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2323000231331021-3123111020331012-1030231323023101-2133111003223012-0312322020222320-2130010303313030-0023110102301102-2120112012322300"></a>

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

<a id="canonical-2013331310000032-0301303231210210-3212110121103322-1032300322110110-2031103311120021-1202330031120312-3002131313103213-1013320310133313"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-3130022132032222-2301223130211220-1330100011220313-0201001110033312-2122212203012003-3102320230010213-2312221222030120-1033023202012021"></a>

#### `kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3232333213213110-1002312103100310-3210213021012112-1023100133130021-2233333321201132-3303011231323313-2213333021010132-1013312133120313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2333320220330222-3223331321123102-1230203212000212-2200200003213010-2321301123233220-2002223331120023-3032023201233013-2313023210133321)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1101102232232130-3122222111232221-0222201010323322-3211122111021131-0302312112032301-1101303330312033-2131123220021222-3111020220023130"></a>

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

<a id="canonical-2233021222230001-0132302020111132-2211110323332211-0300101002332020-0223321222122130-1010310110213221-1020022320221323-3200132010203101"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-2332133131022332-3213230331312102-0101132020103230-2101101013200130-2232000132113133-1302103132301010-1130213201000310-1230212121202232"></a>

#### `kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3323131302200321-1101231213213203-0332022330103312-3330022030111302-3032133331212102-2032220132022323-1022030230010301-3301210121121113"></a>

<a id="canonical-1302123202003312-1203202301300320-2033012332002022-1321203202012302-0231321313133121-3102021200221013-3330202123121031-3001101033330003"></a>

#### `kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2020331122231023-3030131210312313-3321131303313023-3302330121322021-3330023303210331-0320301012331300-1002211211120310-2101020323010223"></a>

<a id="canonical-3033303101131003-3211002212123301-2011301123023321-0300220010331021-0311200112221031-2313213230203221-2103032030312001-2023310300201313"></a>

#### `kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1203031212100311-0311310330301021-3222021132312202-2201122021003312-0323010310321021-1223212321312130-0322333300122112-1313302312011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kvm.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-0312133013000012-2303232001313133-0100112232102302-2320201313200222-0221033330222231-3203021210222101-3200000131331020-2000023302101230)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2122230313320002-3111003313322021-1233023203103120-0311011310331203-1302330331220021-1000101220131132-1001010112002213-0133033230112112)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-3313322110112303-2301203213130310-2121003323131120-2012031010232031-2012320021033113-2303311011113333-1223130321101012-1123201332020203)
- kvm.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-1332111221010100-2212323313001321-0310311323122132-1013310112322200-1210001221031010-1032301032312313-1033010010201220-1301133330121310"></a>

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

<a id="canonical-1231230220132111-3231012233311112-1121233222332031-3032011313232211-3100320001303300-0002011210333300-0210303221310101-1030312103130101"></a>

### Direct properties for `kvm.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-1230102201012000-0011310023130003-1133303123131311-0010211020121010-1331101221111301-0022121132231022-1113200232200230-1003020102301103"></a>

#### `kvm.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1301223103203223-0133032313110121-1001211202210221-1230213220031003-2011022102130311-0210310101231011-0332112120113133-3010221301230232"></a>

<a id="canonical-1122231231321230-2201230123012120-2122133233120100-1202102120211320-3033110002332230-0031332223213232-2221122203322221-1223020020200331"></a>

#### `kvm.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2013023210221220-2012002230302233-0332210211312003-2210303202232112-2101201000333121-3101333310310123-2112213032220120-1121200031201313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `load_balancing` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- load_balancing

<a id="canonical-1100113000012320-2310200021323201-2031033330000000-3231321103211333-0333210101010220-2302213021313223-1020213032223211-0121210013032132"></a>

Type: `"object"`. single nested block, Optional.

This section contains settings on the site that relate to Load Balancing functionality.

Receipt-pinned upstream constraints:

```json
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
load_balancing {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111303020021322-1200123131233230-2030332012331032-1012013310213213-1023310112132032-1233003120100122-0031031100220031-0320000132010002"></a>

### Direct properties for `load_balancing`

<a id="canonical-3031003122010301-2332020101021331-0201001131001122-3000121222201211-0003320121111230-2120103011323231-1201131230012031-3111023201101132"></a>

#### `load_balancing.vip_vrrp_mode` property

Type: `"string"`. Optional.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIP_VRRP_DISABLE","VIP_VRRP_ENABLE","VIP_VRRP_INVALID"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- local_vrf

<a id="canonical-2012010333200333-2100321022022131-0222200301012002-2330003222303212-2110310223232011-2101001212330200-2310121203230122-2030012222332012"></a>

Type: `"object"`. single nested block, Optional.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_config",
    "slo_config"),
  validators.ConflictingObjectAttributes("default_sli_config",
    "sli_config")}
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
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

Terraform syntax:

```terraform
local_vrf {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101210220111231-1313001103012320-1300020321001223-1213221113323113-3021022321312211-3033200011212313-0110312220031031-2102021320113112"></a>

### Direct properties for `local_vrf`

- [default_config](resources--securemesh_site_v2--reference--group-010.md#canonical-2021333233233210-2031212213120013-3221220332022221-2112133230132320-3031001302032301-1032032113213200-2013003213333332-1110330323033023): complete subsection reference.

- [default_sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3200133233230023-0133322030132023-0100110312332222-0330232102000210-2032223313231302-2300232112000000-3311200013030013-0223120322203010): complete subsection reference.

- [sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000): complete subsection reference.

- [slo_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331): complete subsection reference.

<a id="canonical-2021333233233210-2031212213120013-3221220332022221-2112133230132320-3031001302032301-1032032113213200-2013003213333332-1110330323033023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.default_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- local_vrf.default_config

<a id="canonical-1300233211122122-2300300030010333-3313310323212212-3332200330332132-3012223211032021-2013010302010121-0012121223230230-2323233032132122"></a>

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
default_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200133233230023-0133322030132023-0100110312332222-0330232102000210-2032223313231302-2300232112000000-3311200013030013-0223120322203010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.default_sli_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- local_vrf.default_sli_config

<a id="canonical-3333333000102221-1103003110112132-3300222021313312-2303323001021312-0023312220333010-0002012310012301-1222332133131020-0021201110031201"></a>

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
default_sli_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- local_vrf.sli_config

<a id="canonical-2010011201330320-0201311201021311-1230102201312132-0131001310301221-0231100011020101-0002030102311231-0232230231221103-2130310010021013"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

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
sli_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003003032301222-1202120001301320-2211033023331301-3130220032023302-3231032200011221-0101211301230211-2202123113011111-3020211203112331"></a>

### Direct properties for `local_vrf.sli_config`

<a id="canonical-1333020320110211-2303031003211103-3030200221003210-0130111010232312-2133213210231003-2110321302131031-2212213023221322-0020303303012201"></a>

#### `local_vrf.sli_config.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-3233120303011011-2223213330110000-3011003012133330-3313101232030102-0003100123223203-2033131031030300-0102211121321233-0000032001221200"></a>

<a id="canonical-0002120321021003-1212102131333230-0332100032231213-2022203202223202-2313203012323103-0021221111111320-2101031023012203-3230233111220302"></a>

#### `local_vrf.sli_config.nameserver` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [no_static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3000023101031031-2323300021012011-1302012331031101-2322121210013032-3313120222032013-0221123200120103-2221110113030333-1222320110310010): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0302012111002013-0323030023332122-2311002312031013-0232033121100301-0132022023202323-3330310122130020-3221002322023011-1200203230222211): complete subsection reference.

<a id="canonical-1322111133212132-1321332321121321-1212032102320313-2002121123302032-2221031220233323-1130220223012320-3010010231321011-1312111331321300"></a>

<a id="canonical-0133311233001203-0203122103323122-0010333002111202-1013233112012111-0300123312213011-1131220100301000-3311001131312100-3213013331033021"></a>

#### `local_vrf.sli_config.secondary_nameserver` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000): complete subsection reference.

<a id="canonical-3111020220201133-2303022022322112-0331032011001101-3113012310213101-0321321333312012-3332331100313322-2232031300333303-0310012232030201"></a>

<a id="canonical-0230212033320231-0322133030102210-0100122203212033-3002232011232010-1133131000111100-3002000133100231-3210120231122202-1230303013203011"></a>

#### `local_vrf.sli_config.vip` property

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3000023101031031-2323300021012011-1302012331031101-2322121210013032-3313120222032013-0221123200120103-2221110113030333-1222320110310010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.no_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- local_vrf.sli_config.no_static_routes

<a id="canonical-3223301020231332-0112302011301122-2100303303112302-2002233133203312-1210023231312001-2322232300130220-2321330201302021-3023310110011321"></a>

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

<a id="canonical-0302012111002013-0323030023332122-2311002312031013-0232033121100301-0132022023202323-3330310122130020-3221002322023011-1200203230222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.no_v6_static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- local_vrf.sli_config.no_v6_static_routes

<a id="canonical-1133203311223103-1102311321033000-3323232203301101-1313023003131011-0021322233010133-1323100200111103-3332200112322300-1020201023233220"></a>

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

<a id="canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- local_vrf.sli_config.static_routes

<a id="canonical-3302301113100211-1312011221332300-0332233300232102-2013223211232010-1300122330203112-0110122330211032-0323120130110013-0303030202033213"></a>

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

<a id="canonical-1323030132031223-3331231102302313-1130222102223123-3310323113023212-3130111022331322-0110113103223321-1330120013322032-1220033213112113"></a>

### Direct properties for `local_vrf.sli_config.static_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033): complete subsection reference.

<a id="canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- local_vrf.sli_config.static_routes.static_routes

<a id="canonical-3231310330232213-2320233132030030-3323331013223103-0102002031032010-1113321301311033-1223210301310111-1333101111133212-1330320303310201"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0010331010102203-0222122220211001-2120002321220023-2202322230333321-0131321221001110-2310321121000332-2123323012113022-0323011021312100"></a>

### Direct properties for `local_vrf.sli_config.static_routes.static_routes`

<a id="canonical-0300233110330131-0321230130313120-1120002302213021-2130121020221211-3032231132330000-1331302201322032-1310111310231120-2201323213212133"></a>

#### `local_vrf.sli_config.static_routes.static_routes.attrs` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_gateway](resources--securemesh_site_v2--reference--group-010.md#canonical-1222221301321120-1200310123133111-0213323301322130-2003112113202102-0121012030201313-2310012320203130-1031203233313201-0131332200013123): complete subsection reference.

<a id="canonical-0103002332302122-0213013322233333-3301203212203330-2201221101021130-0310302101003202-0310232020033310-1110100302223320-3332112313110121"></a>

<a id="canonical-0130031123332312-0202231302321312-0320231321301011-1021121310031101-1230211310131121-3311131210130311-1132032202221200-1122321200010211"></a>

#### `local_vrf.sli_config.static_routes.static_routes.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0131311032003220-2122201031322012-1322120301133133-1031001322103231-0131003021130231-2201102310120110-1123032223011220-3021020303130301"></a>

<a id="canonical-3013210102202032-0100133310321012-0103001310213110-2330033333312312-1030322002112110-3121302013132020-3131232110033331-3123321120213211"></a>

#### `local_vrf.sli_config.static_routes.static_routes.ip_prefixes` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011): complete subsection reference.

<a id="canonical-1222221301321120-1200310123133111-0213323301322130-2003112113202102-0121012030201313-2310012320203130-1031203233313201-0131332200013123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- local_vrf.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-0030001023012333-2123303232313020-2301301122300030-3133301313031201-3301302310032331-1302003322131000-3323112230211322-3030200201030121"></a>

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

<a id="canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="canonical-2232033031322202-3212121121011002-1120021113113032-3030322223131103-2131113312202110-3331100022300302-2303122303213312-1210203222322300"></a>

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

<a id="canonical-0021220101021031-3031311120310231-3031321210211101-2032033133310212-1203200011332232-0332132303113032-3133101102331202-1121311133120030"></a>

### Direct properties for `local_vrf.sli_config.static_routes.static_routes.node_interface`

- [list](resources--securemesh_site_v2--reference--group-010.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123): complete subsection reference.

<a id="canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-3033331131102212-3010312310202210-0101100012010111-0321331213002012-0123002213302021-0212112323000123-3020032220212330-0301012310322212"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2331012112313101-1000223323201310-2130030300121231-3121121102112021-0013311313330121-1110202233002120-3320002110310022-0213122333002013"></a>

### Direct properties for `local_vrf.sli_config.static_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-010.md#canonical-1130110231120131-0100113210122102-1003331032013330-2101001331221111-1121330010202101-3300032213031131-2332111130100122-0132201331033103): complete subsection reference.

<a id="canonical-0311031232103310-2201103130322230-0313022333023320-3212130330331222-0301002032101101-3130123233030330-0203113320321031-1222201100300311"></a>

<a id="canonical-0301133212003003-3130233031200213-1002131310131213-0012131320113003-2210232123133122-1021013030121001-3011110022221313-3133000303113121"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.node` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1130110231120131-0100113210122102-1003331032013330-2101001331221111-1121330010202101-3300032213031131-2332111130100122-0132201331033103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-010.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3222221200101033-1211120322310023-0322303333120310-0102010033113013-0020023021023102-1030010301030223-1231003322330222-1132102300300232"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2112300122321113-2033012330120310-3232130312120212-2221003203310233-2132232030122321-0021033113221233-0030103000231223-1102121122003310"></a>

### Direct properties for `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface`

<a id="canonical-2313312210021202-2033001221200233-2130231231113030-3033310112223122-3121121220013102-0133120121112103-3012311031103030-1221012003000020"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2102232000121202-3002112320011330-3313010033200112-1102111333130212-0132210310123213-2232323201001131-2011103213003031-0323013233031123"></a>

<a id="canonical-1202122221232320-1320013110130020-2211120113212102-2000120012330032-1132232330230210-2320223120323021-0021113000023121-3033011230210112"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0020102230031230-3130231001232321-0323023201111230-3113211133011021-1333011302122123-0021003320312203-2210331102011131-3303210102333333"></a>

<a id="canonical-2303120111132031-3310322311100310-0012221101313102-0023210111020320-0201312013331333-0321222133311023-0301331133311013-3230313100202212"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1302031311010221-1301201132031123-3212120321013312-1133232212113211-2323313011302321-2231131103133313-2130311222320122-2030011313011123"></a>

<a id="canonical-0220101230302011-1200222210302123-1102000231221222-0110333003233001-0323010122303322-0101303102202233-3213001222302120-2231330203201110"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2223312331232320-3312100010231130-1100223222033022-2110000212211230-3212232120001000-0323213210113331-2200023110121013-2100003313221132"></a>

<a id="canonical-2300312032023133-2213300102201111-3201231203111032-2200231202330101-0333112311333233-0301220232033231-0311300131010200-1231130000330332"></a>

#### `local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- local_vrf.sli_config.static_v6_routes

<a id="canonical-0303001002330011-2330332120220020-3223120101123222-1112201302002220-0310101211033123-3231130332101112-1213221111012000-2211230313202133"></a>

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

<a id="canonical-2332002202121300-1111213331310223-2300210003001031-2311222133002130-1030201113300231-0013011022303001-3203021031111030-2113000231333101"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes`

- [static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212): complete subsection reference.

<a id="canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="canonical-0321111330330212-3202110332033300-1310333010022021-2232031012230012-1211121312002223-0031031303133213-3230322120203133-0200312101303322"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0222003201311131-3112221133231311-3321312200312131-3121201331133313-0311012010210132-0201010030111213-1022321120001001-2133101203023011"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes`

<a id="canonical-1023032310320231-1203230302022233-2033113020032213-2331133122231301-1120301110321120-0211311323030013-2230331300322121-3102212333211232"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.attrs` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_gateway](resources--securemesh_site_v2--reference--group-010.md#canonical-1220231010001020-2320320123020211-2133332032032202-1331213101223030-1102302301130123-2112112111310002-0020030003100311-1213302303300120): complete subsection reference.

<a id="canonical-3012323111223321-0030103313021223-1120100012302322-3023310233110031-0133103023030021-0332020300303330-1000001013213213-3132331111013021"></a>

<a id="canonical-0301200121210112-0300200332100021-1103122012000033-0301300000312222-3011121312021331-0301212330311123-0123312322222033-3123111213120102"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.ip_address` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0033233210133320-1310133030202001-0220031110323000-3112030121223310-0213010001132121-0020223022001012-3010112333202101-0323130020023122"></a>

<a id="canonical-1233313333223123-2302010112110130-0033020201120312-3121020233303022-1100321013013032-2323032130231333-3301220333330212-2210033210113101"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.ip_prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023): complete subsection reference.

<a id="canonical-1220231010001020-2320320123020211-2133332032032202-1331213101223030-1102302301130123-2112112111310002-0020030003100311-1213302303300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- local_vrf.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-3111012110000222-3103332312111120-0203220332331203-3001111321103310-2302233332131012-1221111221033021-1032322010113033-2013203133233203"></a>

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

<a id="canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-1120331223212220-1002022123222022-3113111030110021-2232203003001011-3210331103103200-1031221301132003-1031110020213202-0122313322202012"></a>

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

<a id="canonical-2232320330211110-1301012200332222-2023210213121323-3110020232132221-0113211101102203-1313033000002331-3011121220212310-3300333203320233"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes.node_interface`

- [list](resources--securemesh_site_v2--reference--group-010.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301): complete subsection reference.

<a id="canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-3031111100323322-1003211101031210-3202121122321110-1332213310232301-0203111032220010-2320231330332120-2311322330101311-2132332230110120"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3002003323213102-1123220223312130-2223003303010101-0010000122123123-0303020330123200-1103301000210311-2033310003130330-1002322001301031"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-010.md#canonical-1012310210121023-2123121011130310-1033321320002210-3023333323013011-0231323021112311-3013021312101210-1100203112003330-0211201221110132): complete subsection reference.

<a id="canonical-3101123220130302-0012223211321202-2212322331013230-1213320012003313-1113101103121113-1203303103331310-3020011202003133-2202303321000321"></a>

<a id="canonical-0211210123331020-1133001032031300-2132223032123123-1201301310102202-3022313310012213-1000013021211100-1313011112200130-2101230020013301"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.node` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1012310210121023-2123121011130310-1033321320002210-3023333323013011-0231323021112311-3013021312101210-1100203112003330-0211201221110132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-010.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-010.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-010.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2110203031100002-1301103321000003-2013113011122303-3101103322123012-2321222221201100-1223330212203312-0333103313123032-1033211032310031"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0223003210133330-1120200311301030-3232333033312200-1330223131131133-1013001223332001-2330121131331122-1311323123301112-2130111301132111"></a>

### Direct properties for `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-3023212112331000-3332322122021213-0213120231133022-3313030312323223-1331021012120310-2012201200312103-2223030000310302-2221032000300021"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3113212012130100-1112203121030021-3111231232303132-3013120321120333-1123101130002001-2303022311323320-3103111110212303-3200210122103310"></a>

<a id="canonical-3211112002110303-1113111031333331-0031001312011120-0201033333203212-0203223233101233-1121002301311320-3121320322212212-0231111222230030"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3130222220013332-0200211233020103-1131303320202001-0320303331032211-0001000123303221-2332331323301131-3311333301332310-0233120121032310"></a>

<a id="canonical-0310302331020310-0331301203333332-1102001002321123-2211231212102320-2332100213033121-0122313020323113-3010011201130021-0011131001032202"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3001313212231110-0020231100332213-1221300202320112-3323131031231202-0300211103112221-0222211103100113-2133223323330020-3223311100033210"></a>

<a id="canonical-0221010111232222-1313233300121133-1133201101100003-2000313033321031-0033303031100332-1323000210023330-2121100232010023-1030301010102002"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223130033130110-3023021012330030-3212310112222033-3101211223323223-1223120330212201-3221210203222110-3331320110130333-0032121112131130"></a>

<a id="canonical-2111303312033210-3332201023210021-2221113202031323-2110112332211212-0332301322120131-3031332201022102-1011032033113220-3111101321300120"></a>

#### `local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_vrf.slo_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- local_vrf.slo_config

<a id="canonical-0023201332223302-0123003312310202-1011110332203000-2033320102323201-1012322100101103-1222331131223002-3000230221222010-1111322013103330"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

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
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131201102301310-1032231002121213-0013011332132200-0331000301100223-3222332030200310-1203221323033331-3012323123103221-3130030300033302"></a>

### Direct properties for `local_vrf.slo_config`

<a id="canonical-1233112121331011-3223100010222210-2011210232333203-3330013230322302-2303111220011102-0210211233222012-1102101332002132-3311102003312021"></a>

#### `local_vrf.slo_config.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-0120310020200323-0122333221010312-2221213121211001-0233013331302033-2301210121011113-2301110230123102-2011201222320022-0100300233113311"></a>

<a id="canonical-1130023121211230-3312120320330300-2210301221310032-0200110322311103-3013102332032120-3111113302131202-3313011122303023-3212001002102030"></a>

#### `local_vrf.slo_config.nameserver` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [no_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3010131300322320-1003302333332300-1020211101201223-1330230321111000-3302100331030033-2131211010320101-0113010111321222-3110323122220301): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0012230222212130-0202203112130311-0010030110131100-0230201032312322-0332203222000101-2001310101203203-2120032012002122-3001033010123112): complete subsection reference.

<a id="canonical-2312230010322012-0111232111233023-3321003321230110-2102313323320322-2012313033020020-0311310322221222-0101230110331313-0311100122112202"></a>

<a id="canonical-2130113121011012-0332301320031221-3320212232300302-1112032333311331-3231030321202030-1132121103100102-0201300331232013-3023221030023232"></a>

#### `local_vrf.slo_config.secondary_nameserver` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020): complete subsection reference.

<a id="canonical-3203210313231201-1032100131022131-0312202012012021-1010233122333200-3103101131210112-0023302230020231-3221101220202003-2220210130003312"></a>
