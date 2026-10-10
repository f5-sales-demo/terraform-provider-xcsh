---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2332030330301333-3201233300031002-1103321011310302-0031032130131231-0120121310220101-1110333220000132-2102030201313223-2303202312213300"></a>

## `eks_k8s.not_managed.node_list.interface_list.priority` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-008.md#canonical-0103332210102222-0101313012123113-3110223010330231-1233210031332122-1220201120130302-1030220311033000-1331310033002133-3213231231323310): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-008.md#canonical-2031202033300012-1103300021230131-2131322123021211-3111003121233331-0121331212100212-2222021220102331-2311321321112233-0313323220023112): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-3030211032102301-0013032213201203-2331222321333320-0011121300110003-0303323230202321-1230002033332221-3333131203032011-3301013013201112): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-0301032123232223-1100031103321323-0133220303233001-0223330231112002-3230302111330102-2033000033312333-2302233132112031-3122210031202101): complete subsection reference.

<a id="canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2222203202312133-0332331232200033-3301122010223112-0300330110312113-1203210321311032-2122133113022100-0013303010131221-0312032031001011"></a>

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

<a id="canonical-1013200021321120-1233320021110331-0102130001332222-1332100130203210-0332101011001023-0222033221203001-1200223330220332-1320132322222202"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-007.md#canonical-2123232033001200-0323200020121300-2303121110330103-1333122230220202-0101123010100133-2023003022010201-3101310213003032-3110302321023122): complete subsection reference.

<a id="canonical-0001201102330303-0100012320323213-2131211333201000-1331131230212300-0020001310123220-2212230001201313-2223032130022030-2222221110201120"></a>

<a id="canonical-0120102213311000-3330232322103231-1003111122022032-1100021223211233-2310012013312033-0202320203123303-3111100202313022-2213311012231120"></a>

#### `eks_k8s.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](resources--securemesh_site_v2--reference--group-007.md#canonical-2321301010010310-2102312313102310-1000313020113320-2121212002000110-0301221231033132-3310310130130203-2003301113001120-2133031010100130): complete subsection reference.

<a id="canonical-3130301331102120-0020311003032131-1012210121333333-1011032022221103-3013321030133221-1130233013121022-3302131202303111-1221322221102320"></a>

<a id="canonical-3201111213303110-3230231200022231-0201203023333101-0323323331023201-2100313232130301-3302013223020220-2001322130001222-3301320120033123"></a>

#### `eks_k8s.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

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

<a id="canonical-1201211213102112-3222120120230032-2220031112111003-0013333000023110-2322030220203101-3230030012213322-1303000332030002-1313013130201010"></a>

<a id="canonical-2120303333113303-1233212030103100-1031131013311220-2111203010312200-3213031232222022-3233102001320221-1001111221201010-2331311133221231"></a>

#### `eks_k8s.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-0232233100121233-1103123111232132-1311120333232332-1110111203013102-2101021030203122-0110323122203012-1001011123223110-1301333223132030"></a>

<a id="canonical-1331103223111013-0303120112311131-1203213332033333-3331213001121310-1321101312221130-3102332232103303-0112220330021123-2232223332101220"></a>

#### `eks_k8s.not_managed.node_list.interface_list.bond_interface.name` property

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

<a id="canonical-2123232033001200-0323200020121300-2303121110330103-1333122230220202-0101123010100133-2023003022010201-3101310213003032-3110302321023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-3220211020330112-2110011132230130-0010212301201313-0323202211103231-3013101030002330-0111203132012312-2330233110132231-0213231303211332"></a>

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

<a id="canonical-2321301010010310-2102312313102310-1000313020113320-2121212002000110-0301221231033132-3310310130130203-2003301113001120-2133031010100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1210323321310103-0023110103003103-1000100202103103-3113112021001021-0200223212230120-0120330223300132-1020103330313233-0231221232003202"></a>

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

<a id="canonical-1321203203030003-2030221103222223-0221020223322300-1111013213003332-3213133331312000-1232223230330233-1233111122012013-1231003102112010"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-3203030330231030-1111231012023301-1120131332100002-3312301200001333-0333122022121313-0230223123101000-0200130023121013-1023330320311210"></a>

#### `eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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

<a id="canonical-3012211120012130-2120220310021002-0011001021013213-1201330020123013-2321130021032122-1223300321221131-2310211131333311-0310233021221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-1330233210022311-1132132111111112-3010210332131033-1003222332320003-0311210330321002-3103222320230120-1111021002321200-0110311132103120"></a>

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

<a id="canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3310202212001323-3302310311001201-3322102003122011-1222102233202300-0101100122023332-3112201203120013-0311132233013233-1020211310132112"></a>

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

<a id="canonical-2012333102200032-2231302213322113-1121122103211202-3010100313022131-0311103011333033-2213301110002030-3101111123022102-0101210000232011"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-007.md#canonical-0220100333223003-2111303331103023-0111313021333210-0212302311320110-3200102011323320-3313221010202202-0213311230130320-3313200121322030): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-007.md#canonical-2132110321321323-2332002331132213-3300210233310101-0023012003303313-1032311202200131-0032131110030211-1233013233233303-0212120112232020): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002): complete subsection reference.

<a id="canonical-1130211101231002-0301330333223133-0310300230320231-1333012222203323-0322130220223101-2032120310223232-3102212130133032-3210023101311002"></a>

<a id="canonical-2321011012201021-3311030213003011-1302203320031231-0331321000033210-0210201231223002-1110020101010323-0211211132123103-2333102203031111"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-1122030221132223-2031112012130213-1211032220222313-2013112320102121-3101113010201002-2320102332213010-0121303233202013-1211102131002131"></a>

<a id="canonical-2232101110102032-0030102333020210-1021320001130222-3023112300213210-1100013220303032-0003102233020013-1211123002230103-3002212001110103"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-007.md#canonical-3303302301230200-0011330321312322-1022200030121323-1132113323220100-3121232202222222-2233333120113121-3000322210011300-0330300230101301): complete subsection reference.

<a id="canonical-0220100333223003-2111303331103023-0111313021333210-0212302311320110-3200102011323320-3313221010202202-0213311230130320-3313200121322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1020101330310103-1100302302201002-3302012010201322-0323003231302303-0101321032310102-1213123113303003-1020300311101031-2310220323120023"></a>

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

<a id="canonical-2132110321321323-2332002331132213-3300210233310101-0023012003303313-1032311202200131-0032131110030211-1233013233233303-0212120112232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-1300102333323221-2031332312232212-0300202021113132-0201221022210311-3020220220013301-2112023011230311-2110010011101213-3320213333301201"></a>

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

<a id="canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2321101212212010-0212301333302333-1332133231203300-0331232022023210-1003230101003223-1332021311102110-0331121200103201-3312311330232020"></a>

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

<a id="canonical-1333031323312311-2223202322201203-1231321323213100-1213122112110022-0132003232113032-3233010012123111-2121112133322012-3302013131133112"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3300033311201211-2013110312120131-3211331212103222-3103231233120202-2200001113332203-1130132331101033-3113123221101121-1201033030321333"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

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

<a id="canonical-3312102012211201-1120022202132132-3320102112202210-1023230233301300-3302031110203003-1133221231000200-1322112321233031-1311302032102301"></a>

<a id="canonical-3320332112112213-2232313111122133-1230323222030120-2002122103322000-1210210102320133-2310102131023112-2300230333011010-3303200030320301"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

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

- [first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-2030103303132012-2013131220223312-3303133321333100-1203121131211200-2030130000113211-0320010212133000-3201032210323020-2012102032021001): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0021303331303301-3031221212331311-1100310330232131-1303312212002221-1020110333313012-3102031331122112-3111113310000103-3133313132221002): complete subsection reference.

<a id="canonical-3010010223021030-0331302130021131-0131110023120323-0120320323203222-3323312310033102-2030300032233002-0300021222230200-3332201103213020"></a>

<a id="canonical-2023030223310110-1200231222301023-3131220000302302-0330133030012323-1331000010130313-3330223021230332-1030013330001220-2213210312002330"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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

<a id="canonical-1120100023302103-0031021233232222-3332300333130210-3223011002000103-2030002121120023-1123133100032120-3103330113130110-2023303222222322"></a>

<a id="canonical-0013122311323322-3201331130023203-3210202131112322-3133201020202122-2202000311031031-2002032101221002-0003222300103020-1121310013122131"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-007.md#canonical-2223322130221332-3301201323222230-3221311211100132-2013033103333321-1103010331130331-1310211301002312-3333232223112233-3023020132031121): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-007.md#canonical-0132200210332303-1211133112023112-2300232030020112-1231130312023013-2300100032223203-2020203223000223-0000232123222003-0331210123232222): complete subsection reference.

<a id="canonical-2030103303132012-2013131220223312-3303133321333100-1203121131211200-2030130000113211-0320010212133000-3201032210323020-2012102032021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1001030323120133-0133011102000233-2231203030112211-0212223202111301-3031333131020200-1103313022220121-1102202033322202-0113211301032330"></a>

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

<a id="canonical-0021303331303301-3031221212331311-1100310330232131-1303312212002221-1020110333313012-3102031331122112-3111113310000103-3133313132221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1310000001101032-3130023311320230-0132301000023230-0301030203221202-1002000113021130-3122331322320132-3010211111113020-2212222200113320"></a>

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

<a id="canonical-2223322130221332-3301201323222230-3221311211100132-2013033103333321-1103010331130331-1310211301002312-3333232223112233-3023020132031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2203223133310332-0313321233310021-1121021212330000-3233232200020233-0101010212332132-2131321321110222-2021000221300313-2013011213003231"></a>

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

<a id="canonical-3201011200002320-3332232312230023-3312000302301131-0221121020003013-1230221213123113-1323031121121132-1223132120302133-0011120130331203"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-0021200131101133-3022013231032010-3212011320233000-1233103002311310-2321303130212311-1231212321021221-1132030102133102-3033322322223232"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

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

<a id="canonical-1002101232032003-1332021002012102-3020132133213100-0121023202022301-3133200131322200-2110210112223113-3023001101300201-3223030001110022"></a>

<a id="canonical-0330303110130011-3113310131320330-2120121212120023-3302212132221110-3230020221331022-3312023101112100-2001012131312100-0020313210200112"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0131001300212220-0223330212322301-1322221131020012-2122100111000230-2000303302120021-1021112301120031-3303333100323232-0332022012022213"></a>

<a id="canonical-1230322203013101-2022121203203332-1211230031332300-3112003100331120-2302230223032003-2210123133313230-1031201013133011-3020123312311201"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

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

<a id="canonical-0132200210332303-1211133112023112-2300232030020112-1231130312023013-2300100032223203-2020203223000223-0000232123222003-0331210123232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0111023323231233-1031312030320130-0001311013032110-3120102232132313-0103011223333220-1102121021323103-0233121120331132-1112002113010011"></a>

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

<a id="canonical-3303302301230200-0011330321312322-1022200030121323-1132113323220100-3121232202222222-2233333120113121-3000322210011300-0330300230101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1132130013121312-3313320310022310-3302313111121213-0111321303011100-1333111212112230-2301220022202022-1001003233013333-3330022333201213"></a>

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

<a id="canonical-3112302313003110-1113221133322110-3020122023232133-3003130203100323-1110113201230103-1330310132223033-1100321221312322-3021001212213013"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-3032232333223021-2003002000120002-0131102003331211-1122211230121323-3032312133221300-3202031033133001-3222330233100320-3233011321333230"></a>

#### `eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

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

<a id="canonical-1332210212311330-2022022112011131-1010231320331201-2022030010222231-3332011213032033-0302230033003033-0112221032213112-1223000100220010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0130110330021022-2223312311132201-3111323210330032-1220002213013301-0133123032122201-3133310013312222-1322032132103022-3000022232022323"></a>

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

<a id="canonical-0301131100221230-1002022331230031-2123300222320303-3100213211322222-3312212120300012-3230003223203132-1300123223330121-1301212313203313"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-1232012212122011-1203003131301320-2331323131331003-3030220011113320-0113303321201203-2033200200030300-1222012312322032-0213311230111103"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ethernet_interface.device` property

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

<a id="canonical-1131020101222301-2233230231331332-2120033012203033-0002023312333101-3131333213123313-2223013013332121-2012221200213333-3220133012222102"></a>

<a id="canonical-0211123230202231-2122013032132101-3200212100031310-3221220201001312-3011103220213033-2320323233020131-1232012021020210-2320302210321222"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ethernet_interface.mac` property

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

<a id="canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0323323312231303-3112111130123133-1123221310031112-2132222032030101-3301303110330133-2002120312303033-3022303202210000-1320130212323300"></a>

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

<a id="canonical-3233022132121112-0311322322322101-2311313121011311-2132033300231130-0221012011232032-1023203130122310-3213010302321112-2032132033201311"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-007.md#canonical-3210110102113333-1210220300312303-2102123210011322-3211201311123203-3013323200023210-2000130102311123-0213000301220301-1230221231101130): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213): complete subsection reference.

<a id="canonical-3210110102113333-1210220300312303-2102123210011322-3211201311123203-3013323200023210-2000130102311123-0213000301220301-1230221231101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2001223103233312-0332131201322031-0320323233000031-2211100033132232-2213022010231102-3022113311130200-2003222011330322-0123313120220032"></a>

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

<a id="canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2201312203012123-1331333003133111-2120220033200323-2330312102210002-1322113331200230-2231331020000202-3231022001121202-3132220130003222"></a>

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

<a id="canonical-0200311230212100-0303202323322033-1131230230230022-3120020220233221-3012031020322313-2101323111102011-0231231203033110-1013011233323303"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132): complete subsection reference.

<a id="canonical-0023112320201230-2233013333033101-3232132032002110-1232022320030220-3333111123121212-2200010031032003-0111132321121333-2012033130202013"></a>

<a id="canonical-1131323132320303-3203312130311002-3120120113011021-0010020333021032-2330201320122330-0011220031011002-2220332131020100-2010230132310200"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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

- [stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333): complete subsection reference.

<a id="canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3000222303312220-2021333323121002-0013223010210013-0000203012031301-1032030322301000-1023223023022112-2212010331231302-3311212013032103"></a>

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

<a id="canonical-2020323223213000-2133001211301303-3110300200002001-2112020211002233-1231102101200333-0001323203301311-3210000132213320-2201100130220330"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-007.md#canonical-0301133201300303-0002301023011110-3300133332313333-0021101333021112-0311201222203131-0032301032123223-3302112000333120-1001322302013211): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113): complete subsection reference.

<a id="canonical-0301133201300303-0002301023011110-3300133332313333-0021101333021112-0311201222203131-0032301032123223-3302112000333120-1001322302013211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1230030202311210-1013122031230322-2301033232033330-2220230011132202-1221130301233032-1102113121013232-2312110031331023-1023312222123123"></a>

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

<a id="canonical-0113201000121113-0001110112221232-3013030003020230-3000013001223000-0131013012033321-0112033101322323-0121101032113333-2100102211323100"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1313201313331020-3102022210220322-2101303002220311-3001120302112323-0011100113002112-1111123302222223-2212203010213133-3233011031311100"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2303031310223303-2021102133020203-0112303202033210-1321012122101112-3322113010332003-1302202031112310-1323200132220221-2230000102113210"></a>

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

<a id="canonical-1213323301112123-2223012312220311-2321330013302111-1001311001301300-2321203120112120-3203201112103120-2201102231032133-1113112310320201"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2201312001211202-0132322001223130-3012333031123213-1133320330002203-3132202213100233-2013031030001022-0121202212312300-1012131312013031"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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

- [first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0131022310030131-3333020223333011-1230003233111300-3031222033302211-1333133132132301-2133201232130223-3232221221021313-1123010122311332): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-3121131031322320-0033222122012233-1233012103023202-0001131112000212-2212222210103130-3010022222203320-1002303321232003-1301012211110221): complete subsection reference.

<a id="canonical-0131022310030131-3333020223333011-1230003233111300-3031222033302211-1333133132132301-2133201232130223-3232221221021313-1123010122311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1231001132102110-1003310130010213-1313202010132021-0100123222202202-3233021210021320-3013121130313033-2233300010320331-3331112013122030"></a>

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

<a id="canonical-3121131031322320-0033222122012233-1233012103023202-0001131112000212-2212222210103130-3010022222203320-1002303321232003-1301012211110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1212110222100113-2023233133211302-1221000311202233-3001211101110232-0333211200223212-1222111122331202-0232020101322323-3003121301221031"></a>

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

<a id="canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3312101300231202-2100001031211330-0123103002333311-3011312101332111-3132111232030033-2110112230330320-0323331321310130-0133120112121210"></a>

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

<a id="canonical-2202100102310313-1223333300102200-0230002032133221-3230330013333122-3201032103220310-1300021330233301-3103323120123320-0111013210300202"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-007.md#canonical-1310223221212020-2203212021200230-0221303023101330-1301310133310311-3032121213033212-1202203201003011-3220012231311123-1010330330003120): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-007.md#canonical-2211100321120010-2330031313123202-0132303113010022-1320222311120030-0033133130012002-2302131000322013-1300321113021231-1030002012203322): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330): complete subsection reference.

<a id="canonical-3301100301232221-3220000311010211-0010120231203111-0023210203223221-0033102012101210-2202123302010122-0031013001330312-1320233120002021"></a>

<a id="canonical-0122022323320222-1311303330303111-1030102100103210-2231310233333110-2313231213322031-2132302011322313-2231030210203303-0103213121212213"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-007.md#canonical-0303203000211001-3200133013302030-1102211312123132-3000212230202313-0030300113110230-2310311300332202-1101231220112103-3332220011030312): complete subsection reference.

<a id="canonical-1310223221212020-2203212021200230-0221303023101330-1301310133310311-3032121213033212-1202203201003011-3220012231311123-1010330330003120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3232120323022011-3010321110231310-2223012033222121-0221332223101201-3100223021232102-0132102212033132-1132320113131003-2002313233230332"></a>

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

<a id="canonical-2211100321120010-2330031313123202-0132303113010022-1320222311120030-0033133130012002-2302131000322013-1300321113021231-1030002012203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1003031100121222-0320132003302210-1223020131221213-1021233311300131-0131303320231010-2230112103222030-3122122011033121-2123121300223032"></a>

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

<a id="canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1131313212333001-0232330021110222-1223133200102121-3232110011333123-0312001122300330-0312220012310003-2002321013100321-0033210313330212"></a>

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

<a id="canonical-1130331122210020-2111211001020103-3210323113313103-0101010122330102-2320023302013122-3032302310311220-0303010003323033-1333010121220301"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0131210011013310-2320202330033113-0112200230003033-2212122331000021-3113300011000100-1011313301021213-1301003012111001-3113133212012300"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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

<a id="canonical-2113000113120113-2013313001222222-1113303003112200-2003102011032203-0033031012131320-2302003220131033-1020230222230203-0003313222010110"></a>

<a id="canonical-2120132011112213-3331010203320331-0213300331013121-1032002102001121-3200220120233123-1311211020133311-1200011102110231-3312101302221033"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](resources--securemesh_site_v2--reference--group-007.md#canonical-0100312222123233-2223301211301333-3020313101300310-0211332132332333-3221033133201130-2210022031321313-0000033120103030-3330212000233203): complete subsection reference.

<a id="canonical-0100312222123233-2223301211301333-3020313101300310-0211332132332333-3221033133201130-2210022031321313-0000033120103030-3330212000233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-1213200232123102-3332121300322130-0313020033301122-0300023201113210-3221033331002212-1222010003303133-1032020111111332-2200103011323210"></a>

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

<a id="canonical-3003032210332230-0103310113321133-2123001232212013-2031003113311030-2310202102231201-3011232203101123-3231133301020211-3121121113223033"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-2011110103302321-0002001222110330-1031101013313201-0121201322133301-3301110231132113-3211233303302023-1230130022032130-1303202012112002"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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

<a id="canonical-1302023322102110-1303122000220200-2132202120103313-3111022312010330-3001322330300310-2022223221121111-2002130130201010-2120002303130012"></a>

<a id="canonical-2010132130003130-1000323120312010-3001231222001110-2032020302331111-2001222201233130-1331213012331120-0122222012320101-0210131213332200"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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

<a id="canonical-0303203000211001-3200133013302030-1102211312123132-3000212230202313-0030300113110230-2310311300332202-1101231220112103-3332220011030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1323031123212132-1203021332023311-1032103022222223-3232011112001232-2120013130200103-3031231223333021-2132223022300121-1232213221112210"></a>

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

<a id="canonical-1130032221031102-2031300130010030-3321100330221221-3121122022330220-0331332100233103-2120102201212101-3020101003012220-3122223232231012"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-0032220013200301-0132231011030101-3111312231122002-1202202032130130-1032202123232031-0332132221330132-1003103220333331-2021323000200232"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-0011102111120011-1022033123221032-1331321200332203-2301111122122233-3303333213003203-2021123332202300-1022321002112333-0023222221123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.monitor

<a id="canonical-0110113023203232-3333130321121033-3201310001322312-0101330121121302-1011322002212233-3123031003112320-1233220000110321-0302001212100110"></a>

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

<a id="canonical-3203001200120111-0313132010231330-3001221200201310-3122003203211002-0032332311222332-1021320120202100-3231020003010120-3230012212102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0233121130011132-0113223022011302-1310123130220221-1033213123112033-2131232301203213-0111222031002222-0333001113321010-3300332221133300"></a>

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

<a id="canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.network_option

<a id="canonical-2210300003130312-3323012230220012-1122122012230322-3203213021003031-0212121332331101-3322002231332101-1132330000001122-2033232321123210"></a>

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

<a id="canonical-3032101202211023-0330111320331232-0032032013303012-0023121013331011-1303212303122313-2331220332201032-1332110212110212-2233113220013121"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-007.md#canonical-3111010032020112-1021121231311022-2320112131300130-1300210122301010-3221013130300230-2332332012200102-0033023133122322-3122202333312023): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-007.md#canonical-2013022201331103-1232233332023020-3312111233210202-1000112330131023-2232330313331321-1303133322112130-2313100123301032-3112003223111131): complete subsection reference.

<a id="canonical-3111010032020112-1021121231311022-2320112131300130-1300210122301010-3221013130300230-2332332012200102-0033023133122322-3122202333312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-007.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2122121202300213-3331303220113323-2113132230122323-3013312233321103-3001300132130030-2121201102113213-2110301320321232-1213120300101332"></a>

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

<a id="canonical-2013022201331103-1232233332023020-3312111233210202-1000112330131023-2232330313331321-1303133322112130-2313100123301032-3112003223111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-007.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3020301123011331-1302323001032121-1331232330203202-1221112100323001-1113131312321112-2331231313203201-3131220210101200-0131332330221033"></a>

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
