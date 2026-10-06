---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3312321011133301-1021113131203010-2300030210021321-3112002121211301-2012230221200110-0011233300021110-0201122330032213-2022112021201223"></a>

## Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-0331331123030333-3000212003300011-3123132231010102-0223032020120330-3200013020031313-3110211200301121-3230223310301113-3210202021110323"></a>

### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
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

<a id="canonical-0123210020021113-3011132220021232-3122301132030200-1023301132121011-1211213232212020-2013232020232211-1200101123122012-2220100212032111"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2022102321120220-0232022013001320-1032003112003211-2303332200300013-0311220313312200-1032121222323101-2211312023121313-0223213300121113"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033): complete subsection reference.

<a id="canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
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
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-014.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
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
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-1203003203230303-2012002030222300-1312221333200331-2120030231303133-3111221222110221-2221130010301112-3000113120303222-1033131220013323"></a>

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

<a id="canonical-3312320133210211-2030122232212023-1001030010320111-3022302301331001-2323032100331132-2010110210122020-0201113132211302-3312133310130333"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013): complete subsection reference.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231): complete subsection reference.

<a id="canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
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

<a id="canonical-1321123213122322-2300323310330011-2311022323013232-1211200021210331-1300230102122213-3232311012121132-3200230033102302-0101121202011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3331130110200321-3122031010321023-1130010010121112-3002221022002302-2311212120122022-3200123003223333-2300303201332202-1020011200032132"></a>

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

<a id="canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0112333322212313-0233332103200012-3031232321321003-1101233222201202-0300021333213132-2110221011202113-2103022330312112-0221112320102233"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1200201230213223-2301113100033321-0323101232312302-2230211013222100-2203210102011201-3200000000322023-1331200312212011-1013021221033323"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0102023133312102-1122030103331131-0211300131013113-3022021133303112-2202023332320223-3102032321301233-2221211221101032-2110302110230213"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2010202133113330-1120221312220323-1031033031320020-3321311220212101-1130101131203333-0122131021221302-1312322302210123-3300033121212202"></a>

<a id="canonical-0230232002101100-2331130021012112-1310102232102121-3012233331230312-0033222110202302-2123022313121133-2310000232312031-3302300123221130"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-015.md#canonical-2313331132321311-1132001221220132-0331222201103203-3013021030002312-2230033321133221-2103103201110223-3222131103110330-3130220113222121): complete subsection reference.

<a id="canonical-2313331132321311-1132001221220132-0331222201103203-3013021030002312-2230033321133221-2103103201110223-3222131103110330-3130220113222121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0222030101031120-2203232113010333-0031323121201323-1033211230202232-3031003323212223-1311220103112001-1123311100321031-2310030112121100"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1002333332231120-3313301222303100-2012000320200201-1123303133301200-1001212012021203-1331201022330113-0121310302231013-1103000023231110"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0212300233002111-3002130333313233-2101231333210011-2032022121030332-1013113020212012-0330213323211213-1312211232321000-0233322110232000"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1203321322310333-3102232203211111-3323103123121132-2330202002030202-1202021022232301-1002321103331012-3100133100110011-2130213222110232"></a>

<a id="canonical-2031223112303132-2302303013130311-0202033300211012-1220332203033330-2111133021301313-2200331232103231-2032300132121021-2331001212112032"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1320100020131231-3031030122003123-3122112320312000-1122020000233020-3021323023000330-2330213310102330-0300022023230203-1122100132010333"></a>

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

<a id="canonical-3310121002013202-2231031232122230-0320332301302200-0232202220330031-0121020122011212-1033333123012302-3131003033211200-1311233033101331"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2022021113023112-1220102003110003-0003220133202330-3133312101100031-3310220023021302-3201120003212212-0302120212330010-1302320333202220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1213022213322102-3303133001122212-1230310310003310-3120313120012322-1033111331100220-3021211102013112-0233331131030233-1333203312131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.monitor

<a id="canonical-3031111122310211-3302310023020210-2311313021023323-0202111322222101-1003333103103222-1223000301323122-0200200313311032-2220201121222130"></a>

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

<a id="canonical-3323211302312000-3232023112321122-3023321012011011-1110301230322200-2331101002310313-0300011001120332-2221211310033210-0131321302233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3333113010210320-0033330300013112-1032111100222230-3200321110031102-0003333323203131-0121100130202311-3333113333121210-3000303300232210"></a>

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

<a id="canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.network_option

<a id="canonical-1301032021211102-1213331211122003-1303320213313313-3020030301220200-3201300323310332-2323232330000122-0130212321222200-3303213032020312"></a>

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

<a id="canonical-1110310101033322-1031322221302112-0102011032330323-3300003003332301-0022313021210213-3000232300301222-0301032001103331-0230103232003031"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-015.md#canonical-1203311300332013-2200303332030210-1223311321323021-2130302111211220-0213013012330212-2120112212100221-1230231131201220-0320011001121103): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-015.md#canonical-0203123203103313-1122110333123123-1023231222022003-2010033112013230-0301020022023132-2310201213001300-3301023000020313-3201133103311311): complete subsection reference.

<a id="canonical-1203311300332013-2200303332030210-1223311321323021-2130302111211220-0213013012330212-2120112212100221-1230231131201220-0320011001121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0231102121103130-2122332220023100-0211210100300033-1020122030312203-0211223103100013-0223310122020130-1112030233331033-2033322233022220"></a>

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

<a id="canonical-0203123203103313-1122110333123123-1023231222022003-2010033112013230-0301020022023132-2310201213001300-3301023000020313-3201133103311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0200132003320231-0113033330212101-1211221133200003-3130033312000003-3102010202100031-3201032112002230-0010033023232213-2120012333023011"></a>

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

<a id="canonical-1011011213323210-3122010003102300-0330331301102032-3333301223303111-1201212112100210-3332111231323310-3202020132312321-3022102002123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0031331331223330-3310203030223333-3133022200211231-1010210101122001-1132111013012322-0323020002001231-3200012230232202-2333102210303012"></a>

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

<a id="canonical-3000222332330022-2133212331113332-3312321101221300-3212001103011211-3331313203222203-3013001232002331-3301220303332120-3023210001123102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1322212300030203-3201131303130233-3322103321033223-3102211100323202-1033313210033002-1301332321320203-3323332100103103-2100001100030320"></a>

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

<a id="canonical-1021203102231201-0231333013111010-1313030332200030-3022331203000230-3220220010313110-2231223230203220-1123111001231323-1033123310230230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-2333122203200021-3222111312203312-3210222002312120-2001313102112223-0221131112311110-3231130220011220-2103331022131100-2002101111201031"></a>

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

<a id="canonical-0330203123013021-2220332110311313-1132132123211300-3302302130110331-2120002110212333-0101220202112101-0333103021012030-3022221021223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3000200211111120-2310231031112303-3001213033030001-0221121321121233-1223021300001203-3001222302202212-0131232300212223-0020022021000310"></a>

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

<a id="canonical-2011222002031302-3200302122333032-3123332200133332-1213200011123210-2110002011322231-3321210212222320-3231013110330213-3131112113021112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.static_ip

<a id="canonical-3223010202101321-1123111002213223-1302102203130230-3231332122332021-0012320103302122-1211201312002133-3000222133232010-2102110330003202"></a>

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

<a id="canonical-2321032311223333-2331020302313313-1333202222310001-2301033100012233-3103321033011132-2230101230323312-0021001011302100-0033221321111021"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1313112221303302-2222323012120101-3010210220101033-1021130221331030-0303132133113310-1101230012230123-2333302021322122-1300310103110022"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1030131023322311-2101102321000012-3023320101000311-1232012131231013-2102130333110000-0320102310033011-3223011220022311-0120202220223303"></a>

<a id="canonical-3212223200203221-2212312311311203-2132321111131100-3030112331111333-1212201222302133-3033332132022100-1123010122232131-2100333313331130"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0030031200313312-3012233221003021-3222122030211312-0121321231230033-1032202321313013-0101200303331323-1003131301023011-1012001130111101"></a>

<a id="canonical-1032323203110030-1001312332121110-2101002332323111-2121200110222132-0221232010101030-1013023332021303-0312232211220003-2122131030102111"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3000011220012321-0120100022313311-2323300310312031-3211003301011202-2322232032101211-2130001233232313-3331113031132100-3023102311330213"></a>

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

<a id="canonical-2011203330022311-3313031213321322-1121132113001312-1111203113320031-3112223233201123-3131011023022002-3132112313303320-2232213303121021"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-3220011103312033-0132120323300332-2223101102302120-0013033000022120-2011211112032100-0112320312300001-3210122022112312-3002122132123013): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-2220312333300111-0223021033223300-0222200111230011-3030331223022220-0103202313122203-1203332030322331-1103011122001100-3233320130330010): complete subsection reference.

<a id="canonical-3220011103312033-0132120323300332-2223101102302120-0013033000022120-2011211112032100-0112320312300001-3210122022112312-3002122132123013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2101331232212213-2232310202222133-0101011120321210-2312011221113332-1031320012323202-0011031131233211-2300121333300320-1100233211321213"></a>

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

<a id="canonical-2301122210021211-2200311001123330-0321323220302231-0301022033311212-2010221313301302-0313303002212122-0310221121323122-1123020300122020"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-0023131211312100-2113022123202033-3203231331032301-2230201203212211-0203210312301210-1103213222201332-0112101222323121-2322332222210032"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-2220312333300111-0223021033223300-0222200111230011-3030331223022220-0103202313122203-1203332030322331-1103011122001100-3233320130330010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2012013121133210-2030202002320222-2022130130311223-3111221011211301-2133322131122302-0001203123213211-3133110121203030-2300202102130003"></a>

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

<a id="canonical-1230021303123030-3310012011330013-3322031012130110-3012123312312032-1133000310120033-3202321130102011-3222323211133203-0323132103133021"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3013012112012031-1001203233103031-2330231032201223-2331213210312022-2020213031303033-3130012213131123-2110211120131002-2300301122211220"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3202101122112320-0323002200033011-0110330003130000-0100130102131010-0230013310030133-3302203310022000-1222021021210221-0130233210023200"></a>

<a id="canonical-2112031320021113-0003331102110122-2333211220123303-1320122032330311-3322323032131100-0122203102012132-0211220100122300-1110331121010130"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-1320132013303111-2301300221003020-1330330032011122-1332101113101331-2232003200213212-2031233021332101-0313030112313012-3302023110310333"></a>

<a id="canonical-2333100210032201-1320130133120121-1200331003230121-0011220122301020-0202310001331130-2110211002113230-0001312201230220-0303220211113303"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0030030233110130-3213101333223331-3101211022101203-1022313322122110-2223021203331030-2030210001231300-2333232113003101-1103112122103212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-1301020203021233-0020112222331322-1213201023333012-2123323220011322-1021210121321220-0321021230013301-1233303330011133-0002022130003131"></a>

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

<a id="canonical-0331031210101311-3202133321231202-2032332010330132-3113103331300120-1123303002201300-3210101332330010-2130022003023311-3223010001221022"></a>

### Direct properties for `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-0031332221123023-0023112210210131-2002022032123210-3332022123133322-2013030113332331-1020111222321310-3210323311310122-1100111021021120"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3100123121010332-0131320130302233-2132120330011323-1110031331320023-3110213011130001-3033001232101000-2230333233201210-3231231302000311"></a>

<a id="canonical-1133121030301232-2022230113230302-1200032332133221-0010331003210123-1132332331022122-2020030131113132-2210020220100111-1232222020320021"></a>

#### `openshift_virtualization.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- openstack

<a id="canonical-2201321203200030-2221013332323000-2322311113130010-2100020321100012-2121310021131033-1123013102021223-1310301202122112-3313312103100001"></a>

Type: `"object"`. single nested block, Optional.

Openstack Provider Type. Openstack Provider Type.

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
openstack {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130313211322111-0013130200311233-0302222120320100-0013320113200120-0323313202121301-2213202222102013-0331203210123122-1213331233221310"></a>

### Direct properties for `openstack`

- [not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211): complete subsection reference.

<a id="canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- openstack.not_managed

<a id="canonical-2203331222332312-2200330311121332-2133230111222310-2311312101130212-0100132121100230-2011200322030111-1112130131322233-2100022232331213"></a>

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

<a id="canonical-2233131320131013-0231311301310000-3320131311330203-1320032303221131-0323123133102103-3021303001312101-0322222121310312-0201011300112230"></a>

### Direct properties for `openstack.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332): complete subsection reference.

<a id="canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- openstack.not_managed.node_list

<a id="canonical-2110022221221312-0101320313103300-1233113013000200-3001111123002100-3011100220210120-3203003322222323-0230131313323312-1212313120010310"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3110200012013330-1011233331120110-0332132113333131-2132210103030002-2022000010033312-3302230323313203-1100322233323033-0122020102103311"></a>

### Direct properties for `openstack.not_managed.node_list`

<a id="canonical-1322331111302103-0221112022330010-0300023333311030-3232311021203333-0210130220302233-1320123220211103-3303021332211313-3133030331312120"></a>

#### `openstack.not_managed.node_list.hostname` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132): complete subsection reference.

<a id="canonical-1110031220332222-1012123312232302-2302103110320033-3102130002221110-2011102213302023-2332130101022201-3123010123213101-3323202213303022"></a>

<a id="canonical-2120202311233133-3131320202310313-1211221123033031-1133212231333102-1320000300022132-3132000202112333-3001030323223021-0201023230120132"></a>

#### `openstack.not_managed.node_list.public_ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2123221013120333-0021223130221310-0331203303231111-2300102013303223-3233032100012011-3111102032220300-1120201320001233-0200023221212332"></a>

<a id="canonical-0132322301201112-3010012113222333-1301332103021123-1012221310230003-1232013023122112-2230110121212210-1022201332302030-0233311131111103"></a>

#### `openstack.not_managed.node_list.type` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- openstack.not_managed.node_list.interface_list

<a id="canonical-0310323121211222-2103231002130122-2220220203200233-0211001023320111-0322000121121010-3303202103333120-3023232310211223-2012120220221013"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0301311331032330-2233000312013010-0023222002312223-0231111020032032-1201230332202201-3321320230033021-2302223203230002-0201103212312101"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133): complete subsection reference.

<a id="canonical-2113111103022211-2133211022301112-1000302202312010-1301220202021100-1012010301333131-2212101321122233-2100223302300323-1103112320321123"></a>

<a id="canonical-0021011111310312-2001311133332133-2123321103301210-2132200100101303-0110300133100022-0332302211111300-1112112011133233-3222101331302003"></a>

#### `openstack.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-016.md#canonical-3230113312212313-1330233220212100-3203311232020130-3012012231200201-3230331132330213-1320323302203122-0310130112001121-3002210101233322): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-016.md#canonical-1231132310211101-3332301131231222-0313132121011130-2020001223032222-2332000001331301-0210220312003331-2120001302232122-0231331300020230): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-1330203230212300-2103311130233331-1103133013123230-3202300010230033-3332211103233312-2123113330001101-0223131232213221-1022112130113323): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-0230211231030113-0120120332331000-2101100010222322-1130111123330322-0333212310313033-2113233200233201-2113313201210232-2221311223132113): complete subsection reference.

<a id="canonical-3210023300301021-0303013233213222-3321002033232231-1320332111220011-0230311303030121-0021002232222303-2211310033312202-2300301332313003"></a>

<a id="canonical-1202212101100021-1321122223132201-2322200233333320-3031101112123211-0212333211133113-2011122312301223-3321133032101203-0230000103010312"></a>

#### `openstack.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2102003002112121-3121021002202213-3332001322031311-0323022321120112-0223212230032300-1023311011312201-0122231121011020-2321001300013202"></a>

<a id="canonical-2322013323312023-1033322111001323-0222102310213122-2212132320113012-0032020202312310-0101223211221002-0301013111201230-1030211203320232"></a>

#### `openstack.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1112101020300012-0103023203230210-3112311112232131-2223331212312313-1330132031021302-3211212332101030-0001233213100120-1101322110012231"></a>

<a id="canonical-0030023031231312-1220123232311110-1212003323023210-1303303322022320-3323332013202121-1331102032012032-3031322100333231-3200022223113123"></a>

#### `openstack.not_managed.node_list.interface_list.labels` property

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

- [monitor](resources--securemesh_site_v2--reference--group-016.md#canonical-3101113233000123-3211233033203031-3230112121131320-3021032013030320-2211023113300001-0002133320033300-3100233023020320-2021301303200202): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-3222123322300001-0110133122012110-0111313122300121-0010031303320211-3130112222100330-1130133233002013-2230032311102211-0211232000012200): complete subsection reference.

<a id="canonical-2010220332112102-2003332213113202-3002133231210110-0110302131103002-0022012212230321-3323031022330033-0003310102133023-3123100031313201"></a>

<a id="canonical-0101111030131102-2333031113112022-2101021120010111-1012023222313210-1322321322223133-3003203103213201-2320202011230123-1132031312301013"></a>

#### `openstack.not_managed.node_list.interface_list.mtu` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333102331110023-3131123113022320-0320110323310103-0333000032132102-0200111311310223-0203322230302031-1323133131200000-0322203133212202"></a>

<a id="canonical-3121222133130230-3103033123100110-2022131113121031-1113300120202123-0101011120132022-2002103200122021-2013100310302133-1023323320110031"></a>

#### `openstack.not_managed.node_list.interface_list.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-3023211112201130-1201110122221122-2300322103102332-2310233121202112-0120030230132112-1023123001333201-0132113301223201-1123330213022010): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2332311111121233-1023200032232001-3122133030130130-1123313102323230-0200221122211002-3230133312201011-0122132200032030-0013213232231232): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2110313201023123-2310200020000232-1311232001030303-2222100103031323-0211313131023030-2300033103012320-1233022303303331-0011200211020330): complete subsection reference.

<a id="canonical-3033101202223200-1110333202033113-1000013321203121-2003312312330012-0010000233111230-0110122123012102-0323220322310300-1110101323321013"></a>

<a id="canonical-2033123333313233-2202003231003002-3121330020012320-3110112220010202-0132101013112303-1310230333113123-3032333223123103-3213001003000111"></a>

#### `openstack.not_managed.node_list.interface_list.priority` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0210012301212222-0111212112333310-1312311313231131-0323100032010230-0030031102013120-1311123222203330-0103003211100302-2202220211100133): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-2031330202202022-3331000031001100-1200313120202122-3313120233311233-3322331012221303-0023120222112321-1130321210023233-3011313203321011): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-2112231302320012-3123103320233130-1022202103132302-0121321121313211-3111011202010200-2222311200010133-3021032133223331-0323330130002321): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-2210213202121222-0311220023211303-2120210023332220-2213022210322011-1330301132011020-0300322201021313-0130100222303313-1132220120132003): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-0020330021133011-0122223222303323-2332322122302031-0103032300320231-1322120221332020-0002202212303101-1102001132023322-2011031201011130): complete subsection reference.

<a id="canonical-3222110123222101-1102322120120302-0131000011103102-0212101302323031-0012020121021222-0302112012301031-2200003033230301-3303221213101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `openstack.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1110010322131320-2221002311300011-0022132301131322-0323121331112011-0212213332032002-2031000230123001-3201021000300133-0203213021231332)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-1131102132020103-0331213303112023-3300232213212020-0202202222312002-2212022002021011-1030210222131310-3312211020203212-1300033030310132)
- openstack.not_managed.node_list.interface_list.bond_interface

<a id="canonical-1101112320333211-1213301001313000-0100202023313011-1121202113133320-3203333021212003-1232311120032021-2020133222303211-1132300133202211"></a>

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

<a id="canonical-2213323313201110-0201331230331232-0231221330110100-0002021323032211-3122232012211213-3122132332021212-2001102030320300-0202132302022312"></a>

### Direct properties for `openstack.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-016.md#canonical-0313313111010312-1332003030032011-3332132101103332-0030123133323310-0100012132121011-0121233023100132-1102032120312012-1021002230013121): complete subsection reference.

<a id="canonical-1200121322232002-0012131332032231-2203132031023323-0221323331010322-3112222100302331-1331100031013303-0222021122112201-2101001201010013"></a>

<a id="canonical-0030022130122210-0121200103001330-2302312313102323-0021123102212002-2301103332331122-1011012220001120-2320303001203201-2302332132013100"></a>

#### `openstack.not_managed.node_list.interface_list.bond_interface.devices` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [lacp](resources--securemesh_site_v2--reference--group-016.md#canonical-3102223220023321-0111220332033310-1212222201122120-0302303033312332-0302021232311303-3210110002003212-0133322031022003-0032123012223003): complete subsection reference.

<a id="canonical-2113102110120322-0011023310220310-2233112122301130-1110033003020033-3022331322322220-1013321202030311-0101330021232010-2310220111320023"></a>
