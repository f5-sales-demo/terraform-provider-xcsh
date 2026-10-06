---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices.active_backup` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-001.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- [bond_device_list.bond_devices](resources--fleet--reference--group-001.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- bond_device_list.bond_devices.active_backup

<a id="canonical-2132113223210212-3220131323133320-0112021122002113-1331132211002101-2212032312213130-1230303323303103-3132122000332110-3231200233220012"></a>

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

<a id="canonical-1232131310312013-1102232002320220-3200203002001023-0003223112323331-1230230223333130-2100022101102310-1212322333003213-1131321102112223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bond_device_list.bond_devices.lacp` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-001.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- [bond_device_list.bond_devices](resources--fleet--reference--group-001.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- bond_device_list.bond_devices.lacp

<a id="canonical-3313323230011200-3112222003013333-1332013011020223-2101133212133213-0133030103102010-3231322320010231-0110131020223102-3201102113332030"></a>

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

<a id="canonical-2100331322111130-0331301312212221-3310121312013202-2303123123113033-0213110303130233-2122012220022013-1301032123221130-1100202301232102"></a>

### Direct properties for `bond_device_list.bond_devices.lacp`

<a id="canonical-0121011021031200-1203113213112302-2311030033311221-0001200330000302-2302020222001220-1313333320120131-2133220231103311-3311213310231231"></a>

#### `bond_device_list.bond_devices.lacp.rate` property

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

<a id="canonical-2220011112301110-2101322302212032-2331222020113323-3333113013301010-2200210103313131-3003210110031021-1011111332222330-3232220020111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- dc_cluster_group

<a id="canonical-3222132011302110-3223233110303022-0010111003310331-1232121213200230-0212010203312230-2220311223301131-1013331123331132-0011233232032100"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group, dc\_cluster\_group\_inside, no\_dc\_cluster\_group; Default:
no\_dc\_cluster\_group\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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

OneOf alternatives in this subsection:

- [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-3222132011302110-3223233110303022-0010111003310331-1232121213200230-0212010203312230-2220311223301131-1013331123331132-0011233232032100)
- [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-3113032020311002-1032323211212323-0330330230100021-2112321321021100-2313331222110031-1213131220011002-2013013013100021-3210322011302131)
- [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-0300111311003312-2320102111320110-2333221333222221-2321223000332000-0132220211220113-0221111301300211-3031311131013010-3001123212213133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133111233323012-3333032003030323-0311313333210011-3032030100220223-1003200200020322-3232320103321012-1313000333011022-2201222013201100"></a>

### Direct properties for `dc_cluster_group`

<a id="canonical-0101131011112223-1230130100133332-1220213300302321-2300002103000222-3121030302332303-2313202311120022-3112022200100121-2300011020133120"></a>

#### `dc_cluster_group.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2133010021202033-1133110231202332-3311230000223312-0103230123323001-2003333320312222-0331320310323013-0303101222201001-1323102123303313"></a>

<a id="canonical-1233131231231232-0001011222032003-0110201011100221-1003113330311103-2231112203120011-2333332212323332-1212113300031102-0320300231011302"></a>

#### `dc_cluster_group.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2023322113123002-3100320322200311-3001321002020130-3332211000200200-3223000030112330-3303213001022330-1221003223012102-0221221123233202"></a>

<a id="canonical-3123330313233211-3220333301132301-3233121132332300-2023113110103122-1320021112203031-2301022003103130-1120213203001133-2231122023130011"></a>

#### `dc_cluster_group.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1022322123131331-3133331231203323-3213011213101211-0310133103101302-3003011111020032-0112210323033012-3222000332033020-0103003311002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group_inside` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- dc_cluster_group_inside

<a id="canonical-3113032020311002-1032323211212323-0330330230100021-2112321321021100-2313331222110031-1213131220011002-2013013013100021-3210322011302131"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
dc_cluster_group_inside {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332322033101033-3221132031322203-0133330230130330-1323330021203101-1022011311023100-1103130011010100-1031100313031323-3121333210211211"></a>

### Direct properties for `dc_cluster_group_inside`

<a id="canonical-2011323303112120-1312121233303233-1311311230113121-0221023112223103-2320311333333331-0203021123121111-3322023001102331-1102313203031032"></a>

#### `dc_cluster_group_inside.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2100220030012020-0321102223000001-0102112221223102-1222121232320330-1020103111112312-1221001121120222-3010133001003222-0130001013213001"></a>

<a id="canonical-2301110023202212-2313232033310030-0111120022133203-0131011331000123-3321213133100210-2021312323213103-2102332100222021-0313211223120003"></a>

#### `dc_cluster_group_inside.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2012030333200310-0333322331231021-3020333102301233-1322113333332233-1121311121112010-2121323022000230-3023201203213011-2113201032312211"></a>

<a id="canonical-1313010022201301-2131322122122332-2030013022320310-1303012331200301-2012031222123202-3232011301302023-0300120222021203-0233213112000231"></a>

#### `dc_cluster_group_inside.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3310002022222312-0033100302223102-0112021232310302-1233201223101302-0221021103321211-0101113223023222-3000011131123213-1121030200103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_config` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_config

<a id="canonical-0222331031323120-0233230301313130-2222300312211303-2131330122312100-3030030103000312-3321300032012020-1221213113311200-3302320011013321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

OneOf alternatives in this subsection:

- [default_config](resources--fleet--reference--group-002.md#canonical-0222331031323120-0233230301313130-2222300312211303-2131330122312100-3030030103000312-3321300032012020-1221213113311200-3302320011013321)
- [device_list](resources--fleet--reference--group-002.md#canonical-3200322202323330-0323133003013331-1122130122320213-3221210232101013-1200323220311002-0232323111303111-2110131211011210-0020213222312313)
- [interface_list](resources--fleet--reference--group-002.md#canonical-1022102312212002-1321333300332120-1002202302210311-1003202220302323-2023331213223200-0220022300211302-2130200100001300-3021321322102123)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321320013020113-2010101112101310-3001121223201313-1212300320113031-2233220102201110-2003223123110010-3213112211322133-3220222221031301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sriov_interface` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_sriov_interface

<a id="canonical-3120231221133033-3133113300203332-3300212000321212-2230302302231033-3110323323131002-3333002303210003-1111030102211010-3220300203133231"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

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

OneOf alternatives in this subsection:

- [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-3120231221133033-3133113300203332-3300212000321212-2230302302231033-3110323323131002-3333002303210003-1111030102211010-3220300203133231)
- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0122221131232302-1301033322113301-0110223212011211-3233102320110033-0121020202101031-3330230010133021-2322320310031011-0300013133213110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sriov_interface = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311012123131011-1011230030130023-3102322012232302-1222231101331323-1232103111123230-0012022301103233-1003210001001231-0212032003000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_storage_class` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_storage_class

<a id="canonical-3232310021320301-3232001110331123-2122322312002112-3020302312022233-3221320130031123-1100300210011002-2302313031231330-2332332332120130"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_storage\_class, storage\_class\_list; Default: default\_storage\_class\]
Configuration parameter for default storage class.

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

OneOf alternatives in this subsection:

- [default_storage_class](resources--fleet--reference--group-002.md#canonical-3232310021320301-3232001110331123-2122322312002112-3020302312022233-3221320130031123-1100300210011002-2302313031231330-2332332332120130)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-2222300023023313-2122302313220111-2233300300012313-3303221100131111-0112033102321132-1321110322120221-0103222322331100-1100211021322120)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_storage_class = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012100320202111-1003330223212301-3032022132110231-1312231300111330-1332100132012122-1310200213231333-3313103222231312-0121120321130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all_usb` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- deny_all_usb

<a id="canonical-3031222031130102-0220323323303102-0001201301130320-0310333220032103-2020133202021100-3213120313202223-2013331220210232-1002130131230330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for deny all usb.

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
deny_all_usb = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- device_list

<a id="canonical-3200322202323330-0323133003013331-1122130122320213-3221210232101013-1200323220311002-0232323111303111-2110131211011210-0020213222312313"></a>

Type: `"object"`. single nested block, Optional.

Add device for all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
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
device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323323033213233-1123011032113102-0103212231301130-3022330121323232-0101223220230303-1321310030220033-3312013112331001-3131300220100131"></a>

### Direct properties for `device_list`

- [devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101): complete subsection reference.

<a id="canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- device_list.devices

<a id="canonical-0231031102013112-0023111223302101-2021310021322003-1200312221122032-1300223112001013-2331211030032020-3300200103100133-2033202112010213"></a>

Type: `"object"`. list nested block, Optional.

Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras,
scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet
and has an corresponding interface/device. The mapping from device configured in fleet with
interface/device in VER node depends on the type of device and is documented in device instance
specific sections.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023103201023012-3113120120203300-3031203133221101-3302211211012210-0201221321331030-2111002010131231-3021313232312211-2230310221032011"></a>

### Direct properties for `device_list.devices`

<a id="canonical-1100231202123210-2011110011112001-2113330223130011-2212213222220313-3201110322220330-3000120231022133-1111321100131301-0003123202132031"></a>

#### `device_list.devices.name` property

Type: `"string"`. Optional.

Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of
device in host-OS of node.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

- [network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310): complete subsection reference.

<a id="canonical-2211213033230302-1031233322012201-3022301211020123-0221130310223300-2010203301032032-3020303323311103-2233211130100012-1112111100021233"></a>

<a id="canonical-2333133021230121-3321233033101122-2303313231112331-0032202022100100-2212110131120122-2102003122321101-3103120101210121-3112220111310303"></a>

#### `device_list.devices.owner` property

Type: `"string"`. Optional.

\[Enum:
DEVICE\_OWNER\_INVALID|DEVICE\_OWNER\_VER|DEVICE\_OWNER\_VK8S\_WORK\_LOAD|DEVICE\_OWNER\_HOST\]
Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will
be network interface device or accelerator like crypto engine. Possible values are
\`DEVICE\_OWNER\_INVALID\`, \`DEVICE\_OWNER\_VER\`, \`DEVICE\_OWNER\_VK8S\_WORK\_LOAD\`,
\`DEVICE\_OWNER\_HOST\`. Defaults to \`DEVICE\_OWNER\_INVALID\`.

Additional upstream details:

Defines ownership for a device. Device is available to be owned by vK8s workload on the site, like
camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via some other
service. Like TPM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEVICE_OWNER_INVALID",
  "enum": [
    "DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices.network_device` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- device_list.devices.network_device

<a id="canonical-1233122212201230-3130021332033113-3121011121310113-3221100021110321-3012311120233122-0100013233333031-0122123300302103-3222001210333012"></a>

Type: `"object"`. single nested block, Optional.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interface")}
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
network_device {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111320102223212-3230313012322030-3133110023103100-3302322320100211-1100113031311022-0020320302320323-3001030103313230-3302130212321111"></a>

### Direct properties for `device_list.devices.network_device`

- [interface](resources--fleet--reference--group-002.md#canonical-2121333001221000-2320111310332123-3212003300122303-3001303200202330-3121002213211000-0230301230121010-1120321110201001-0223031210312232): complete subsection reference.

<a id="canonical-2231211212012001-1230322122322023-2101323312122103-0023012020120213-1103321033103103-2130100023103020-3130312223002122-3120003011331310"></a>

<a id="canonical-1023323032122222-3021300302001121-3312030222111101-1321210103230110-0332112232102223-2302110130210322-0022322033031122-1330122110300211"></a>

#### `device_list.devices.network_device.use` property

Type: `"string"`. Optional.

\[Enum:
NETWORK\_INTERFACE\_USE\_REGULAR|NETWORK\_INTERFACE\_USE\_OUTSIDE|NETWORK\_INTERFACE\_USE\_INSIDE\]
Defines how the device is used If networking device is owned by VER, it is available for users to
configure as required If networking device is owned by VER, it is included in bootstrap config and
member of outside network. If networking device is owned by VER, it is included in bootstrap
config.. Possible values are \`NETWORK\_INTERFACE\_USE\_REGULAR\`,
\`NETWORK\_INTERFACE\_USE\_OUTSIDE\`, \`NETWORK\_INTERFACE\_USE\_INSIDE\`. Defaults to
\`NETWORK\_INTERFACE\_USE\_REGULAR\`.

Additional upstream details:

Defines how the device is used

If networking device is owned by VER, it is available for users to configure as required If
networking device is owned by VER, it is included in bootstrap config and member of outside network.
If networking device is owned by VER, it is included in bootstrap config and member of inside
network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NETWORK_INTERFACE_USE_REGULAR",
  "enum": [
    "NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121333001221000-2320111310332123-3212003300122303-3001303200202330-3121002213211000-0230301230121010-1120321110201001-0223031210312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices.network_device.interface` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310)
- device_list.devices.network_device.interface

<a id="canonical-3100120003312303-0003110011302032-3320031212231121-2220210003310112-2103210012021210-3102323323031331-0133312302300322-1002303103203112"></a>

Type: `"object"`. list nested block, Optional.

Network Interface attributes for the device. User network interface configuration for this network
device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in
VER node If network interface refers to a virtual-network, the virtual-netowrk type must be
consistent with use attribute given below If use is NETWORK\_INTERFACE\_USE\_REGULAR, the
virtual-network must be of type VIRTUAL\_NETWORK\_SITE\_LOCAL or
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE if use is NETWORK\_INTERFACE\_USE\_OUTSIDE, the
virtual-network must of type VIRTUAL\_NETWORK\_SITE\_LOCAL if use is
NETWORK\_INTERFACE\_USE\_INSIDE, the virtual-network must of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-3312322310002120-0223120102011010-2212021333000122-1200002033002313-0113031321121110-3200332112032112-0330221301230210-3322200022210011"></a>

### Direct properties for `device_list.devices.network_device.interface`

<a id="canonical-1012010012103110-3030010120023321-1000231213302133-3332203032230213-0132012311012310-3320232210202323-3103000102220023-1013202021123323"></a>

#### `device_list.devices.network_device.interface.kind` property

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

<a id="canonical-3233333311000113-0020001313212012-1200231301321133-1222110310101332-1022220111201131-1012201330230012-1311132232333331-2033201310000320"></a>

<a id="canonical-1020231010033013-0120012312231232-2022220221120033-0322333310223333-1012123032301222-2122321323223223-3011111101330322-0100201132000202"></a>

#### `device_list.devices.network_device.interface.name` property

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

<a id="canonical-2213021333121212-1021213211210102-0122021231233013-3310220120110300-3122232133132230-2013200320003131-1021210103032312-0202312102310021"></a>

<a id="canonical-0123020112201132-2133302302302331-3100111322011232-2231200302223222-0303220100221333-2230312231031000-2322023010330001-2220120300132111"></a>

#### `device_list.devices.network_device.interface.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0230321303203311-1001303000233233-3313032111020212-3230212302231332-0223323100130231-3000333322323200-3001030112332013-0301123220030223"></a>

<a id="canonical-1121211202120122-1213301220322032-0022123321020032-2332233303011320-3101333103312203-3203121202201021-2002010102202022-3330011132211020"></a>

#### `device_list.devices.network_device.interface.tenant` property

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

<a id="canonical-0233221033313310-3010220003120213-2212031120312101-0220202203310302-0331302111111313-0012001102223032-3333320201013310-1123322110301202"></a>

<a id="canonical-2120211022032213-1301212331200223-3200201121223121-3233102330332230-0122132320303011-0003030333222031-0333310121012321-2312003020230003"></a>

#### `device_list.devices.network_device.interface.uid` property

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

<a id="canonical-0123032211233331-0131332231333212-3200202011200112-1001000300113120-0122200030312231-1201011230021221-0213033021011323-3032133220031003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_gpu` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_gpu

<a id="canonical-0213033021223211-2011031122232000-0031331130102311-0223100313001222-0032002333002023-3222030023220001-1220103333102311-3323302230201330"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable GPU.

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

OneOf alternatives in this subsection:

- [disable_gpu](resources--fleet--reference--group-002.md#canonical-0213033021223211-2011031122232000-0031331130102311-0223100313001222-0032002333002023-3222030023220001-1220103333102311-3323302230201330)
- [enable_gpu](resources--fleet--reference--group-002.md#canonical-2200303220000331-1001222020122121-1313202222301130-3330333133022123-3302320023303020-0201020112330320-2301302111030103-3220132132312202)
- [enable_vgpu](resources--fleet--reference--group-002.md#canonical-2022102100100133-3233222033211321-3200233210103333-1131101321130311-2120202203330101-1302300011010321-3003031331100131-3122303110120030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_gpu = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030313313212311-2001301220310203-0020001123321003-2121032212103232-1131303332211222-3020011200213223-1213020113230210-0100321121202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_log_anonymization` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_log_anonymization

<a id="canonical-2121001302333101-3211331033211213-3233213331231013-2301210300232123-0033110031101222-0312002133222130-0120100323223301-0331300320023001"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

OneOf alternatives in this subsection:

- [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-2121001302333101-3211331033211213-3233213331231013-2301210300232123-0033110031101222-0312002133222130-0120100323223301-0331300320023001)
- [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-3200320200030321-3033013222300332-2010230132310221-3032313333233322-0332313113200121-1201222331033303-3120202212101301-3020103131101003)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_log_anonymization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321001020131001-1211220132331302-1322103320333300-0031202311220102-1230303330102113-1311000031300313-1122113011301220-2122300130120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_vm` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_vm

<a id="canonical-3230331213222102-3010132211202323-0302300222121300-1322023031220311-3323322213002302-1122103131213313-1321302120133121-2211100120320133"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_vm, enable\_vm; Default: disable\_vm\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_vm](resources--fleet--reference--group-002.md#canonical-3230331213222102-3010132211202323-0302300222121300-1322023031220311-3323322213002302-1122103131213313-1321302120133121-2211100120320133)
- [enable_vm](resources--fleet--reference--group-002.md#canonical-0300310111033333-3323012300233121-0223320333212221-2033211212300122-0320122222321001-2130221000220223-1131213121203130-2103020131022232)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_vm = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002001130122322-1321022131203221-1122200222031023-1020213032321102-0222000230321311-3133232323322212-0110120030211330-3132231013210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_gpu` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_gpu

<a id="canonical-2200303220000331-1001222020122121-1313202222301130-3330333133022123-3302320023303020-0201020112330320-2301302111030103-3220132132312202"></a>

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
enable_gpu = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300130001320220-1113103202210133-1323332200222222-1011232231123110-0001121122301303-2210010201302013-1133102312333222-1032303221030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_log_anonymization` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_log_anonymization

<a id="canonical-3200320200030321-3033013222300332-2010230132310221-3032313333233322-0332313113200121-1201222331033303-3120202212101301-3020103131101003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable log anonymization.

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
enable_log_anonymization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023220101031001-1331310302022102-0213011031321233-3112103310211311-1212211111023103-3221232112323313-3332022200013311-1333313210022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_vgpu` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_vgpu

<a id="canonical-2022102100100133-3233222033211321-3200233210103333-1131101321130311-2120202203330101-1302300011010321-3003031331100131-3122303110120030"></a>

Type: `"object"`. single nested block, Optional.

Licensing configuration for NVIDIA vGPU.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_port")}
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
enable_vgpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103121101121101-0133302310120003-2303111201223123-3130032122301030-2120002102232112-3320110121200112-3213221100012332-0301333310123311"></a>

### Direct properties for `enable_vgpu`

<a id="canonical-2023210113311333-0320211122000010-1023321012002001-3210010223100003-1321020212111320-0122033120111313-0332033011031222-0102020333322310"></a>

#### `enable_vgpu.feature_type` property

Type: `"string"`. Optional.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNLICENSED",
  "enum": [
    "UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023131222202213-0331220323130333-0131111200110203-2033123101001111-3121102110021223-0101032120223231-0121001010101202-3330220300231212"></a>

<a id="canonical-1020112021030032-1331012201011012-1110201331000103-0210032313230113-1322132302330230-2022222231332111-0120221010102302-1132203011033133"></a>

#### `enable_vgpu.server_address` property

Type: `"string"`. Optional.

License Server Address. Set License Server Address.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="canonical-0110302323212200-1232132013132232-0011133101132231-1011111220312201-2020020031121032-0320011033100232-0210111310221232-0201022112313332"></a>

<a id="canonical-1103021010111000-2302203200103330-2331212233100333-2322100113302210-1203002022310110-3010211113002120-0112020020123101-1310311030333011"></a>

#### `enable_vgpu.server_port` property

Type: `"number"`. Optional.

License Server Port Number. Set License Server port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0322133113031201-3211102313221333-0210203311332212-0110112033311332-0120013321023303-1301023021020300-0230103122131203-2030002233213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_vm` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_vm

<a id="canonical-0300310111033333-3323012300233121-0223320333212221-2033211212300122-0320122222321001-2130221000220223-1131213121203130-2103020131022232"></a>

Type: `["object", {}]`. Optional.

VM Configuration. VMs support configuration.

Receipt-pinned upstream constraints:

```json
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
enable_vm = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333200312232013-0113002003132011-3330301312033033-2120022120220222-1123100201032211-1221013003232322-1100332121322133-3130123311000222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `inside_virtual_network` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- inside_virtual_network

<a id="canonical-2122322131012313-3130301003123021-1212130323302111-0203212231233110-0003213231120130-0211202223132033-2230322211113002-3013330103212130"></a>

Type: `"object"`. list nested block, Optional.

Default inside (site local) virtual network for the fleet.

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

Terraform syntax:

```terraform
inside_virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230331331311310-2231100122013223-3110221110101000-1223202131313133-2131331030222312-2112030102323302-1012001020120320-0021231120011011"></a>

### Direct properties for `inside_virtual_network`

<a id="canonical-2230030320120311-1301002313331101-1033031031133302-3130011021332011-1203003022220020-1011231331303022-0010212112311232-2230211222220020"></a>

#### `inside_virtual_network.kind` property

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

<a id="canonical-1011023102110022-2333212123222312-1202220112210103-3232201131211031-0133131312230321-2111002112331103-1120030101022102-3121000312223003"></a>

<a id="canonical-2323032112201003-1313202210302212-1101332211210201-2232031230101000-0321220330131001-2331100023113303-1321102212202203-0112233223100212"></a>

#### `inside_virtual_network.name` property

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

<a id="canonical-2111022200220232-1310232023101322-2300133301202011-1320312010121333-2030313313300120-0303122210120001-0022301332210111-3123113233212033"></a>

<a id="canonical-0132020303100313-1320003122113102-2120113320000311-3202033032031002-0021100322220122-1312101212011030-1103220002112202-1001330130230000"></a>

#### `inside_virtual_network.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2022230213020013-0333210331102312-2310211033222213-3330313003233032-3102003110321220-3011030031032203-1210003232121031-2130033312131303"></a>

<a id="canonical-0311132202011001-3120001213033022-0303000003200202-0221300111130323-0031030201213001-3013232303103113-0310011001203333-3220030221000111"></a>

#### `inside_virtual_network.tenant` property

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

<a id="canonical-3221222313112131-3131320132212203-3111322330230132-0023230111333013-3312113210322012-3021321332013130-1210013322202333-1213001110000322"></a>

<a id="canonical-1333233101030112-2321000320023300-2023102202121232-2203011123232102-0313231223110230-2020231112113210-0120202001322000-3333210311220301"></a>

#### `inside_virtual_network.uid` property

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

<a id="canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `interface_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- interface_list

<a id="canonical-1022102312212002-1321333300332120-1002202302210311-1003202220302323-2023331213223200-0220022300211302-2130200100001300-3021321322102123"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212220211013133-3032000223210033-1301021011303210-2133102200130221-0103201030033133-2010021312201223-2332223332002311-0313210133112012"></a>

### Direct properties for `interface_list`

- [interfaces](resources--fleet--reference--group-002.md#canonical-0212032112031011-3111023103022010-0033202112313203-0310232033202222-0112031123021210-0030212302331233-0133212012233013-0111201133002333): complete subsection reference.

<a id="canonical-0212032112031011-3111023103022010-0033202112313203-0310232033202222-0112031123021210-0030212302331233-0133212012233013-0111201133002333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [interface_list](resources--fleet--reference--group-002.md#canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302)
- interface_list.interfaces

<a id="canonical-2231333023223222-2300022113210131-0103201221101230-1302323202012100-0130121122330333-1001010010030123-0302102232132101-1103210332202212"></a>

Type: `"object"`. list nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303111332000200-3133222303230002-2111211101223132-3300221103011300-1233010233030300-1031331033201110-3030331331203130-0123222122320101"></a>

### Direct properties for `interface_list.interfaces`

<a id="canonical-0012002302301001-1302031001230033-0101120022101302-3220031301122001-3000330122031021-2112003213300001-0121301102032112-1110011131000000"></a>

#### `interface_list.interfaces.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1000221230233223-2330200323332330-2233300233233020-0223011131213102-1113102002130023-0000211223313220-2312231210132112-0330130220230133"></a>

<a id="canonical-3222131211033203-3313212021102112-3322131111003103-3003301222032130-1303223222301230-3012020100313110-2020301212103311-1010332103022100"></a>

#### `interface_list.interfaces.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0131003000100302-2302222210203211-1121332333131202-2003223101221013-1100033212023303-1333133232102301-0003312333312100-1131023003200221"></a>

<a id="canonical-2022010330212323-2210213300322201-2311311221131100-3121002003102020-2201233312023203-2313200011010323-1333302012100011-2211231131320213"></a>

#### `interface_list.interfaces.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- kubernetes_upgrade_drain

<a id="canonical-2021130120130131-0131201122032111-0100011213212223-3331321110111101-3311000033113323-1332022231022230-0220132320100133-0221011220232023"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131323200313301-2321000122202300-0030001233321212-1102201030113110-3210031003212033-1220121000103200-1112013222012022-1321311132222213"></a>

### Direct properties for `kubernetes_upgrade_drain`

- [disable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010): complete subsection reference.

- [enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122): complete subsection reference.

<a id="canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-1230301023033301-2231012323222023-1022233031331311-2121301231200022-2032231011102332-1103112103023220-3100110302212212-3223221220113201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-2200003321111012-2102312032003130-0331322023323100-1033223013033010-0132110322312003-2121320100203112-2220001120330032-0321201312231310"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331001200330323-2301032203020100-1210111203201310-2102313233332322-0200031032221003-3021302321111230-2002220112320312-1201001311300112"></a>

### Direct properties for `kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-1133220013133121-3331112130002231-2321003331133202-0201202120332211-3033030202301313-3102023213011322-1120023112300210-0012231231231000): complete subsection reference.

<a id="canonical-0312303200220300-2110120203113210-0133330021211203-3022113320111311-2211100210022330-1322023011232212-0220103030102123-2031210133020101"></a>

<a id="canonical-0100002013200220-3330331022002102-3220010230013023-3101303231200332-1001102101102032-1012010102302003-2020223331110322-0302110133110020"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1330022231112010-3133223023113013-1310120132011001-0331200003013230-2032223111103331-3000102103002220-2302031111131021-1330001033131320"></a>

<a id="canonical-0221211120220210-0220112013313323-3020310320211210-1331233020022133-3000333110033311-0031211023333211-2132013302333322-2220222230203313"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-3323030303221233-2100022310303231-0221302131302232-3011012213232112-0001210020220121-0302223122202021-3332000223033102-2023203132313122"></a>

<a id="canonical-3010233222102103-1031212211313310-0203231221210133-2332111201201021-1313032030033200-0312332120311100-1300103120221133-2323131303330301"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222): complete subsection reference.

<a id="canonical-1133220013133121-3331112130002231-2321003331133202-0201202120332211-3033030202301313-3102023213011322-1120023112300210-0012231231231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-3303103200232103-1120111120302102-2101210203313300-1121032310221130-1213002132012230-2233011102031033-0001200020220030-2212021010031011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-0001132232313221-1301212330111231-0201303232201201-0020323302213312-0210301323111031-3303033113100022-2131211311330102-2310322302233300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331330130330132-3110332130130011-0101221033032011-1321323001303111-1310212222202120-1000000110001011-0200233010121032-2221330200010033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- log_receiver

<a id="canonical-3031313201013102-3012032012303320-0312101313130103-3322003301113303-1220323132023022-0301030301323133-2222033021303010-2022122130110203"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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

OneOf alternatives in this subsection:

- [log_receiver](resources--fleet--reference--group-002.md#canonical-3031313201013102-3012032012303320-0312101313130103-3322003301113303-1220323132023022-0301030301323133-2222033021303010-2022122130110203)
- [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-2131023030210200-2311112100032332-0203003232113123-3101323130301101-1231112112221012-1033211311300001-2131011123301122-0232232301033323)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123211333133112-3102212002213103-3031303310111001-1212222121303200-0320202112212203-3132312113302201-1331001101022211-0020320212220313"></a>

### Direct properties for `log_receiver`

<a id="canonical-1130020113021020-3320232131202113-0302012111310003-0010301032313331-2123331221203013-2310211022130210-2303212231301010-3303012113202202"></a>

#### `log_receiver.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3203000220210021-2012022203121223-3231300120111030-2311331332031003-0123213133002111-1122201032101230-3010032103021231-3122320031310201"></a>

<a id="canonical-0120230133321123-3230201130021300-3200203121333013-3202113221332000-2011100000023331-0102220101331033-2210013213101121-2221033000310120"></a>

#### `log_receiver.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1112232211002303-3213210020311113-3230102202232032-3131313112232133-1330010131302300-3313312032212033-0302132231000212-0332111013211222"></a>

<a id="canonical-0112121030313303-1023002321332133-0133110321012312-1310313210010232-0312032022202121-3010113033332323-1021002203100003-1202200220013013"></a>

#### `log_receiver.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3212213123002023-0202231202121031-0011310312010030-0202200201110132-0022133203021121-1110002111002320-0011010213122003-0312100022301211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- logs_streaming_disabled

<a id="canonical-2131023030210200-2311112100032332-0203003232113123-3101323130301101-1231112112221012-1033211311300001-2131011123301122-0232232301033323"></a>

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
logs_streaming_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323111111030122-1303021120311223-1110003221103022-2020332013330231-3132231113011220-1121100020211332-0213211322111100-3303110210031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_connectors` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- network_connectors

<a id="canonical-2233111321103202-3320001311001223-1302200122123131-1120102313203322-1013031230100002-1200223111020300-0110311031222211-0030300120222110"></a>

Type: `"object"`. list nested block, Optional.

Network Connector defines connection between two virtual networks in a given site. Fleet defines one
or more such network connectors. The network connectors configuration is applied on all sites that
are member of the fleet.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
network_connectors {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302330323030330-3210110300133132-1321230300222320-3023231111122212-1331211200030201-2323121012322313-1311010331123231-3220123120210230"></a>

### Direct properties for `network_connectors`

<a id="canonical-2220233203213002-1010103303011202-0020302231231130-0332332120200001-3030001203330110-2201333303220222-2331012120200023-1333221330230123"></a>

#### `network_connectors.kind` property

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

<a id="canonical-1222331120102022-2213033310302010-1313011233223210-1123013312203001-1003010313200211-1110200101102231-3130101130101030-3300202123100332"></a>

<a id="canonical-1313023130122031-3312301030031312-3332222321130113-2200130313012003-1200031231122213-3102321303022200-1220103211233331-0201120320031031"></a>

#### `network_connectors.name` property

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

<a id="canonical-0033221133101031-1201020303111012-2030122312223333-2000210102121201-1111111123130132-2301313130220002-2321331110312333-3112011021021220"></a>

<a id="canonical-2013332233223303-1133322033122230-0233230332312121-0302211330131131-3102033211021031-3120100100230321-2112031102113133-1321013203231022"></a>

#### `network_connectors.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3232100332020232-0223211231000021-0132322330000322-0101123321202332-1302031103303010-0101032032132210-2323231122121121-3021000030311100"></a>

<a id="canonical-1021001333331311-0112000001012131-1022203303331303-3033000110221121-3233231023301021-2301130022211120-1023020302133223-2321321221001102"></a>

#### `network_connectors.tenant` property

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

<a id="canonical-3102022112233220-1000322113230013-0332312313232121-0103231013001200-1211023212000313-2323232300103020-2111202333330322-0220212330102320"></a>

<a id="canonical-1033230311323103-0012202210211221-0232230120111031-1232330313123132-2310000322133231-3311000111003213-1001031233221330-0002330322101321"></a>

#### `network_connectors.uid` property

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

<a id="canonical-2101130020120112-0133200133323012-0232110022111200-1113311010303213-1313213332310323-1200301201112022-1100302111320120-3222312321231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_firewall` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- network_firewall

<a id="canonical-2320210312033302-0003021301023012-0231131302212333-2022302003133123-1222032202030023-3301111031131101-1210122331032011-2320112313121321"></a>

Type: `"object"`. list nested block, Optional.

Network Firewall defines firewall to be applied for the virtual networks in the fleet. The network
firewall configuration is applied on all sites that are member of the fleet. Constraints The Network
Firewall is applied on Virtual Networks of type site local network and site local inside network.

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

Terraform syntax:

```terraform
network_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020301002032222-0023101121102003-2333021321010113-3123132211321132-1231221023100232-0331333201102111-3212111023011220-1233000003023322"></a>

### Direct properties for `network_firewall`

<a id="canonical-0223030201022003-3120030120002220-2201022022220103-0223220033120311-0121213212010021-3001100203102230-1030200003033131-2321320300330320"></a>

#### `network_firewall.kind` property

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

<a id="canonical-1021011111130311-2302113223123123-3120133233300333-1311320221333012-1312233201330110-3213202312310301-3323320320113112-3230332100202130"></a>

<a id="canonical-1132131030202222-3123303132300033-2003312031113031-2032023003313101-1121001203122202-1213122122002201-3323032030213003-1323122323110031"></a>

#### `network_firewall.name` property

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

<a id="canonical-3012301122232132-3310212310201133-0023123300211213-0200022233111102-1201113132100133-0122113022013111-3023032100001110-1133030001311331"></a>

<a id="canonical-3100030313012322-3300020201332230-3323232313101223-0203010100032210-2330031332220020-2303020310232310-3000202330102302-3022222131120031"></a>

#### `network_firewall.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2312010111200312-1013332103323313-1030102302222122-3323230000221131-1203030230001121-0100033232311223-0100320331210120-3111131221222000"></a>

<a id="canonical-3023230320020103-0311321122101310-1321113022013011-0101212223200212-3321213003203013-2123320020130301-0032133230022210-1201323000220202"></a>

#### `network_firewall.tenant` property

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

<a id="canonical-0000100210323231-3311001011020032-3322011200100021-3320011102311200-3213011023220103-0232113231202020-2230302313332210-1111320230113121"></a>

<a id="canonical-2202132233231302-1221310030110233-2332022203230232-3313131300202102-0022222202113322-1220211030030110-1120131301233201-3310300111322222"></a>

#### `network_firewall.uid` property

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

<a id="canonical-3203320110130321-1121100220230211-2020333200023102-0312320112101333-3102032302000203-1112313033231111-3102003021133322-0232123301133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_bond_devices` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_bond_devices

<a id="canonical-2233202132112121-1300211023330221-2212133210020112-2021322122322312-0321223001202132-3220130012222211-3210110102122303-1332222011023230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no bond devices.

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
no_bond_devices = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002230220210021-3122231112230020-3301210303100111-2000220001202332-3133200133012031-0222312213013001-0021203302213202-2210201320011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_dc_cluster_group` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_dc_cluster_group

<a id="canonical-0300111311003312-2320102111320110-2333221333222221-2321223000332000-0132220211220113-0221111301300211-3031311131013010-3001123212213133"></a>

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
no_dc_cluster_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323333231331210-1331121020203033-1001232023202221-0323221132032011-2122012311233313-0333321001230223-1230201132012120-3320022013311213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_device` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_device

<a id="canonical-0220230132023221-0012122010222212-2110320123231220-3000120023300020-3313010120130211-3322003111321031-1011003102230130-0132323201010321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_device, storage\_device\_list; Default: no\_storage\_device\] Configuration
parameter for no storage device.

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

OneOf alternatives in this subsection:

- [no_storage_device](resources--fleet--reference--group-002.md#canonical-0220230132023221-0012122010222212-2110320123231220-3000120023300020-3313010120130211-3322003111321031-1011003102230130-0132323201010321)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-1220331113303212-2010311203330203-1123010023313010-2333311011310323-2013213130203030-0232022330203123-2033201030211020-3221200223133131)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_device = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021101310113300-2330120320311311-0103112023302012-0111200030121332-2122032220110113-0122101022233212-3300032121002220-0102132112203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_interfaces

<a id="canonical-0132211203012133-3201031322322320-1332302222000203-1001303002202201-0323112222133113-3313003302132131-0010231220301220-2322322220001322"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_interfaces, storage\_interface\_list; Default: no\_storage\_interfaces\]
Configuration parameter for no storage interfaces.

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

OneOf alternatives in this subsection:

- [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-0132211203012133-3201031322322320-1332302222000203-1001303002202201-0323112222133113-3313003302132131-0010231220301220-2322322220001322)
- [storage_interface_list](resources--fleet--reference--group-003.md#canonical-2223313213113003-2010202220123320-3322302112013032-3001322211322310-2322330013330331-1010013302213030-2221310300112203-1322231330200313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_interfaces = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231011202030311-3303111022101200-2001103313313330-0212102002001302-2030113131220333-1311323330233023-2120212100232203-0031032331001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_static_routes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_static_routes

<a id="canonical-3333003330002323-2020312322133222-3003221302331313-3221312221202102-1100230230310222-2202101300132202-0113222211333133-1021333131121331"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_static\_routes, storage\_static\_routes; Default:
no\_storage\_static\_routes\] Configuration parameter for no storage static routes.

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

OneOf alternatives in this subsection:

- [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-3333003330002323-2020312322133222-3003221302331313-3221312221202102-1100230230310222-2202101300132202-0113222211333133-1021333131121331)
- [storage_static_routes](resources--fleet--reference--group-003.md#canonical-3300220221311210-2223123102101211-3220330103110130-2013011230213113-0302003112110032-0211100233331131-0000213033313103-2010020102112211)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223331122001220-3011311231202031-2020330101213300-3133211103110113-2113210013003020-2001322230121133-3212301331022221-2100121222312301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `outside_virtual_network` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- outside_virtual_network

<a id="canonical-3120120032101112-1320012311032120-3123110212013301-0222000011303300-0020203112032230-3201230233112233-0103200303010013-1233332133233032"></a>

Type: `"object"`. list nested block, Optional.

Default outside (site local) virtual network for the fleet.

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

Terraform syntax:

```terraform
outside_virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202030110010311-0312311300332203-2223030132103022-0133100222022123-0112300030012112-0100110122232011-1103200021112000-1300111021020223"></a>

### Direct properties for `outside_virtual_network`

<a id="canonical-0103321323202010-1112100230213231-2013222330100131-3100300130311313-2123110110201013-0013230323113200-1000111231213133-2131021032013111"></a>

#### `outside_virtual_network.kind` property

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

<a id="canonical-3111030303003033-1030111031000020-0202023121210221-0222210232033030-1330123013202110-2023112113130310-0321010213100132-1200310213311033"></a>

<a id="canonical-2131132201013110-1212223001323000-3010122222323113-0321323332332012-2031332330011220-0212301212121333-1012313111200213-3122231213310102"></a>

#### `outside_virtual_network.name` property

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

<a id="canonical-1122113103211203-2003130010301203-1212211233311010-3002233032130233-0030102201132022-2023022231111300-3001233132033331-3131211233030110"></a>

<a id="canonical-0000313322023332-1122102213300011-0222110111310222-3202233133221210-3120023233111310-1221012213333323-0131033223303110-1222233131000321"></a>

#### `outside_virtual_network.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0212012002130323-1011132302233203-2020131211303310-2022201103312101-0020133112000330-2300131020232101-0020022303021000-2212310200333010"></a>

<a id="canonical-2231213131201221-0001323100103000-1021212201321012-0313103003312233-2123300032330200-1101000033021021-3213320212211121-1121001003220331"></a>

#### `outside_virtual_network.tenant` property

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

<a id="canonical-1012003302101233-1223031302033032-3132122001101032-2312012212303001-1310122212210233-1303120200210121-3033311322021301-0330003121323333"></a>

<a id="canonical-0012103332010223-1031031033331203-2122320313102032-0010110101021002-3203312030110301-3022030331003322-0303202003322202-3332003023220211"></a>

#### `outside_virtual_network.uid` property

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

<a id="canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- performance_enhancement_mode

<a id="canonical-3030113113100330-0012133311113321-0021320010121312-3022300023333220-0130331330032012-2302300003220202-3021321212323322-3003213102113110"></a>

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

<a id="canonical-0211303102202002-1111131203121220-1200231221230303-3213201011123203-1013120212212212-3300201000011203-2121133312130330-2323112121233313"></a>

### Direct properties for `performance_enhancement_mode`

- [perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103): complete subsection reference.

- [perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003): complete subsection reference.

<a id="canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3203210222120212-1213232111221030-1330301330231323-3320223321203311-1210023122323120-1110002201030313-1323331031332233-0222312122231221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Additional upstream details:

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

<a id="canonical-3233131011312211-3213022000110301-3103000121302113-0001210032122303-0210123321122121-2302020311313211-1302003100233202-0232301222213301"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l3_enhanced`

- [jumbo](resources--fleet--reference--group-002.md#canonical-1001231011333031-2123112303131321-1213331333230000-2112333001033020-1012323323213332-3020133020200113-1020202000330121-0023131311333213): complete subsection reference.

- [no_jumbo](resources--fleet--reference--group-002.md#canonical-3230201301023222-1303110022122011-2300010221103110-1200010220323030-0320033011332130-1002301323003230-0230131213100330-2213101312101300): complete subsection reference.

<a id="canonical-1001231011333031-2123112303131321-1213331333230000-2112333001033020-1012323323213332-3020133020200113-1020202000330121-0023131311333213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2321122131222032-1311331312232023-1220132003213010-1313331233223101-2301010022103032-3321120112233322-2030223230311232-2111120030010120"></a>

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

<a id="canonical-3230201301023222-1303110022122011-2300010221103110-1200010220323030-0320033011332130-1002301323003230-0230131213100330-2213101312101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-0221131011213303-3333302211113030-0332103013132331-0313211313323002-1312111223120222-1330302333230123-1221022132130011-2131020022022031"></a>

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

<a id="canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0102133110100122-1213312131112131-3212313202300102-3010213002211013-1112333132023102-2303102120221133-1312032032313222-1303320333220101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Additional upstream details:

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

<a id="canonical-3301122030312022-2012012100121211-0030220031032200-3011110301210330-3030113330000213-3121202030323222-3021031023101211-2102012312131203"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l7_enhanced`

- [jumbo_disabled](resources--fleet--reference--group-002.md#canonical-2013030130220303-1200331321133120-2113203120312213-3222231223320221-2101303323100111-2330021213300330-1111101213210312-0213131131202213): complete subsection reference.

- [jumbo_enabled](resources--fleet--reference--group-002.md#canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220): complete subsection reference.

<a id="canonical-2013030130220303-1200331321133120-2113203120312213-3222231223320221-2101303323100111-2330021213300330-1111101213210312-0213131131202213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-2323033031003202-0123222000220312-1310020102232010-1321133223313132-3002221002330302-2223310003123311-2211110100030312-3233313331120110"></a>

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

<a id="canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-3100110103313003-2220220000200110-3121111023101023-1311130111113200-1120212313003100-0323210120321110-2000120303323000-2303102232221120"></a>

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

<a id="canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sriov_interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- sriov_interfaces

<a id="canonical-0122221131232302-1301033322113301-0110223212011211-3233102320110033-0121020202101031-3330230010133021-2322320310031011-0300013133213110"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

Receipt-pinned upstream constraints:

```json
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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100332211202032-1312200001220202-0013001013130002-2332323320112321-2132003130010303-2212110000132201-0201020130232213-3300233221210302"></a>

### Direct properties for `sriov_interfaces`

- [sriov_interface](resources--fleet--reference--group-002.md#canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320): complete subsection reference.

<a id="canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sriov_interfaces.sriov_interface` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021)
- sriov_interfaces.sriov_interface

<a id="canonical-0111313130111223-0101332233300300-2203102011300122-3303001100021221-1201313331221303-2320100023133330-2300300123213223-0013031000311111"></a>

Type: `"object"`. list nested block, Optional.

Use custom SR-IOV interfaces Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interface_name",
    "number_of_vfs")}
```

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
sriov_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123311220102321-3203312033123313-1000222011120111-0333113010211202-0220113322022220-3230010000121211-3133111210302230-0031130232120202"></a>

### Direct properties for `sriov_interfaces.sriov_interface`

<a id="canonical-2123331310213120-2210230032010112-2320220102201000-1013122020123333-2133301212021301-3302300131120120-2333320113133233-2011003233203113"></a>

#### `sriov_interfaces.sriov_interface.interface_name` property

Type: `"string"`. Optional.

Name of physical interface. Name of SR-IOV physical interface.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2023133122231310-0113222003213021-0313132300113010-3222133301301012-1123330223030202-1202220130201302-2123131332333001-3213231000110110"></a>

<a id="canonical-2231001013230310-3001221131011102-3333011132013100-3322103103232332-0003233312311301-0202131211300003-1212130012322303-1330232213122121"></a>

#### `sriov_interfaces.sriov_interface.number_of_vfio_vfs` property

Type: `"number"`. Optional.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233020321322311-2312010020020301-2120132130223210-2312002232001100-2331221210303221-3130003303230112-3330131212031312-1023031200111211"></a>

<a id="canonical-0132102321323323-3021211212012131-3323111300300203-1213111311222330-2203223320022001-3031303023200120-0133230121212000-0023131032110203"></a>

#### `sriov_interfaces.sriov_interface.number_of_vfs` property

Type: `"number"`. Optional.

Total number of virtual functions. Total number of virtual functions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_class_list

<a id="canonical-2222300023023313-2122302313220111-2233300300012313-3303221100131111-0112033102321132-1321110322120221-0103222322331100-1100211021322120"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022010331230331-2123100211020202-1211010210211212-2223331202110222-0222030203321133-0300322012230030-1321121003030012-2330003002203323"></a>

### Direct properties for `storage_class_list`

- [storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012): complete subsection reference.

<a id="canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- storage_class_list.storage_classes

<a id="canonical-1330213321222123-1121113101212033-0130331233330120-2323211313312210-3212020022203233-2102131212001131-0330220113011110-2023323211312113"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name",
    "storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103323201130130-0211022002333311-1310131321300112-2332323300232102-1221311131102211-2002310331322231-0133011122121022-2000031320201131"></a>

### Direct properties for `storage_class_list.storage_classes`

<a id="canonical-2221030123300233-0110230120011332-1303202011130000-2020112230013002-3221220130102100-0102221332011012-2113210020330311-1113012001022321"></a>

#### `storage_class_list.storage_classes.advanced_storage_parameters` property

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
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
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2202133111232311-1231331012012123-1121023221003231-2102332300032331-1132210010111111-3333320013013012-1330210010031012-0023121003112131"></a>

<a id="canonical-3213312020220000-0230313032031120-2330301102213112-1201320311313213-0333222222300221-0113002101330231-2120011120032312-3010201320103100"></a>

#### `storage_class_list.storage_classes.allow_volume_expansion` property

Type: `"bool"`. Optional.

Allow Volume Expansion. Allow volume expansion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [custom_storage](resources--fleet--reference--group-002.md#canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221): complete subsection reference.

<a id="canonical-0310131120333103-3131013023012023-3332321000031320-0230220103031110-1122331331303233-2320211131001321-3302202012303030-3111003212101112"></a>

<a id="canonical-0132112130003130-3121030010313012-1131221012321310-3001311111012203-0030332320121221-3110332331012110-0021112320310132-1313020031010330"></a>

#### `storage_class_list.storage_classes.default_storage_class` property

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1120201231001110-3312300231013311-1103123310101213-3023112130023312-1033000201010220-2131110100311320-3233213220021030-1320012310312203"></a>

<a id="canonical-3101132223223131-2103031001020311-2330200123133302-1222012213003313-0131313011302000-3221230220323233-3013000022033111-3220323130112312"></a>

#### `storage_class_list.storage_classes.description_spec` property

Type: `"string"`. Optional.

Storage Class Description. Description for this storage class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [hpe_storage](resources--fleet--reference--group-002.md#canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013): complete subsection reference.

- [netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102): complete subsection reference.

- [pure_service_orchestrator](resources--fleet--reference--group-002.md#canonical-1110101010012033-3333011200130233-2120103112121132-1312112202313113-2102222332321202-2231022203130123-1000222321310221-2201120133301120): complete subsection reference.

<a id="canonical-1221330002110232-1330210010213303-3122131210210101-2323203131323133-2210223112321131-3302203033113122-0303333011323223-1330111211100203"></a>

<a id="canonical-3202122012332233-0302012020023031-3033301011233102-0300211333021300-3100022201310230-3122233312302302-0100132223300000-2313112211233101"></a>

#### `storage_class_list.storage_classes.reclaim_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Reclaim Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-1122312030110112-2123300032021100-1302131230102013-1010202122131123-3030221101211202-2200120003120302-2321320302132103-1322300000230231"></a>

<a id="canonical-2320133203030313-0320202222333122-3133120130030310-1331313133113020-3333210210331121-1232230130212200-0132312222231203-2013221302200013"></a>

#### `storage_class_list.storage_classes.storage_class_name` property

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

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
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1011230231022323-1012133323300310-3301103320223322-0031203100122123-2023112220210130-0010032230233322-1221322301321213-0131000113332123"></a>

<a id="canonical-3001030232220130-0000002311123020-1010332111222300-3103321121133202-2332013233010313-2103011120233122-1222102113213021-2132033103302223"></a>

#### `storage_class_list.storage_classes.storage_device` property

Type: `"string"`. Optional.

Storage device that this class will use. The Device name defined at previous step.

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
    "byteLength": {
      "max": 64,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.custom_storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.custom_storage

<a id="canonical-3033311331231221-1311001010132332-0211110001021002-2302112231023110-2122322120033213-0223103231220001-0223003330203112-3310232133231010"></a>

Type: `"object"`. single nested block, Optional.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

Receipt-pinned upstream constraints:

```json
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
custom_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001120233323203-3032011232032213-3003303022121333-2132311023033001-2320132123213322-3100130321301322-0011000220231310-3300001330333002"></a>

### Direct properties for `storage_class_list.storage_classes.custom_storage`

<a id="canonical-0012223213023021-2022231023300113-2210222233333302-2300202213333101-1102133000202220-3303213120332332-3021112012210223-1220301023303211"></a>

#### `storage_class_list.storage_classes.custom_storage.yaml` property

Type: `"string"`. Optional.

Storage Class YAML. K8s YAML for StorageClass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.hpe_storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.hpe_storage

<a id="canonical-0030320200321221-1122032232023202-1010033100131301-0303232232220311-0120120111310233-1311333110233221-1032103003203013-3313223100132023"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for HPE Storage.

Receipt-pinned upstream constraints:

```json
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220131123002020-3021001220222310-2003021100000000-0023131201303131-2213233013331223-1020303231020002-2323213100301232-1030220211332220"></a>

### Direct properties for `storage_class_list.storage_classes.hpe_storage`

<a id="canonical-3021222303033311-0032023011212013-0131011032311111-1012211033111220-0121110212320220-3331320321311220-2111013133113221-2233032231230101"></a>

#### `storage_class_list.storage_classes.hpe_storage.allow_mutations` property

Type: `"string"`. Optional.

Mutation can override specified parameters.

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

<a id="canonical-1103003000332122-1130002230033202-0200130110002010-2133123131130111-3102010010002133-1233332023123312-3010132233003312-0323210003333100"></a>

<a id="canonical-0001300203223231-3320201331203212-3032100013123222-3231031321320321-0220323330313320-3013001312133111-0132303110100331-0201322132310210"></a>

#### `storage_class_list.storage_classes.hpe_storage.allow_overrides` property

Type: `"string"`. Optional.

AllowOverrides. PVC can override specified parameters.

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

<a id="canonical-0330213203012130-2131032223020120-0232111132100010-1013333302020001-0200300013211313-0200300133302132-1013002211223202-1032222331131232"></a>

<a id="canonical-3233332101021312-2322300320103102-1123302313100123-0221020300333313-0101232220121002-3222100022330303-3110331103101223-1023232113101003"></a>

#### `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` property

Type: `"bool"`. Optional.

Indicates that the volume should enable deduplication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212323122033303-1030320320313302-0131212130300330-3223123002300120-3123001121303220-0123010213302023-3013222132213321-3020203200131001"></a>

<a id="canonical-3302311333221120-0122001220301111-2032122130311230-3301011321120130-2321331202202102-1302123030312103-2122031222112320-0220201232013120"></a>

#### `storage_class_list.storage_classes.hpe_storage.description_spec` property

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

<a id="canonical-2032311201210020-2101301322313320-0330103020121311-1311021021023000-1011123330312332-0302022323023311-2033032323003223-1020310020313203"></a>

<a id="canonical-0010120233111333-2202113131012312-2222012200022321-3223102313332301-1231022212121330-1211100120321221-0231230102113211-2010122121131232"></a>

#### `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` property

Type: `"bool"`. Optional.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213233032032202-3222120021103130-1020000100120213-0203032012220101-0023033333301311-0110102123120122-1133210300102221-3132231231121312"></a>

<a id="canonical-2330030220323311-1022302030000312-2303302130010221-2100331103013103-1010311222222230-0231201131030302-1220010020303132-3010123223002112"></a>

#### `storage_class_list.storage_classes.hpe_storage.encrypted` property

Type: `"bool"`. Optional.

Indicates that the volume should be encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100223302023313-1230212121013111-3020013301330222-3000211003311213-3233111132103221-2310101332022212-1000112022203000-1012321220111313"></a>

<a id="canonical-0001213011102332-2123232133020333-1103221300301323-0330010023221213-3121001001003320-0122202000103302-2310310102232112-2231001123310213"></a>

#### `storage_class_list.storage_classes.hpe_storage.folder` property

Type: `"string"`. Optional.

The name of the folder in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0331312333303230-3213032021203001-3303331222003121-1211221103231122-0003210111232303-1203323130103202-2331131220233332-2020122331222211"></a>

<a id="canonical-1102210303220320-3010131221203323-0312123222333120-0330311312031101-0032330212213123-0320222221233331-1112313211013332-3101131202103101"></a>

#### `storage_class_list.storage_classes.hpe_storage.limit_iops` property

Type: `"string"`. Optional.

LimitIops. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-0000202201101113-2000113021101311-1312103120020030-1221112332321330-1020013122032001-2223202120311011-0101113002111033-3331022012310011"></a>

<a id="canonical-0110312203312222-1323100302233202-1131220230103121-0313113333231022-3332220222100203-1321330220210230-0103331202112003-1121311201022223"></a>

#### `storage_class_list.storage_classes.hpe_storage.limit_mbps` property

Type: `"string"`. Optional.

LimitMbps. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-2312032311313213-2103301311122211-3003002223320220-2113321111222123-0110023132211331-2123022212003132-0232101001233332-2113202313030022"></a>

<a id="canonical-2312230302023311-0200232311202331-2202120230031311-2103210333131200-2311220123010233-1202320101323123-3011302130303001-1012111212332130"></a>

#### `storage_class_list.storage_classes.hpe_storage.performance_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1123120313000032-3010213213201310-0201310123223002-2130203001013203-2113200321303231-1001020332010112-0232220012000023-3323000002030032"></a>

<a id="canonical-2030100003013320-2222100223010332-0211013010133313-2233020200221112-3330203232201100-2312210033022211-1312021132110312-0123023032111233"></a>

#### `storage_class_list.storage_classes.hpe_storage.pool` property

Type: `"string"`. Optional.

The name of the pool in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1130323311201032-1120010203311003-1323331001033313-2201120330021223-0312121332110113-0001010113003223-3102232331010231-3123121123013300"></a>

<a id="canonical-3202330030331300-3023202212230323-2203231112311330-2233313132323212-1113112223302211-0201002122112112-1320222202101212-0130330032023220"></a>

#### `storage_class_list.storage_classes.hpe_storage.protection_template` property

Type: `"string"`. Optional.

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0303123221003200-3100113303333222-3112311000012323-0222213103100310-1201100100133021-1001010210221121-1002012323101012-0310110101333333"></a>

<a id="canonical-2002312313111113-1021312200211210-1032030021230022-2313322030310030-2200312303102031-1110100013030032-0231322222200011-3030333311311232"></a>

#### `storage_class_list.storage_classes.hpe_storage.secret_name` property

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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

<a id="canonical-2021201030233111-0002201221112103-2101222223120030-1220313023023210-0200030331033230-0203023022310311-2112201211212102-0020331232210010"></a>

<a id="canonical-3111132023103211-2322310103031012-1333030211120013-2333221313201312-2022321312031231-3000330331331022-3132012102210031-3111010233202330"></a>

#### `storage_class_list.storage_classes.hpe_storage.secret_namespace` property

Type: `"string"`. Optional.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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

<a id="canonical-1221100221101202-0002132220221330-0232110121112122-3212132202120000-2321312103122110-3210013200122132-0323112100110030-2020031100302300"></a>

<a id="canonical-2000230323101323-0110201301312111-0112210300020001-0111333203300331-2102022031000111-1220022100101212-2012233113233332-3331132210121322"></a>

#### `storage_class_list.storage_classes.hpe_storage.sync_on_detach` property

Type: `"bool"`. Optional.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103130213311223-0331203131333031-2012120002033221-1020313133213322-1033323230020123-1001123233122311-0010201322001133-1111122310101123"></a>

<a id="canonical-2312303133330121-1021333000302003-3312111022331002-2300022313023300-3023021230102123-1122010130333311-2000001320311233-3010323031312213"></a>

#### `storage_class_list.storage_classes.hpe_storage.thick` property

Type: `"bool"`. Optional.

Indicates that the volume should be thick provisioned.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.netapp_trident` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.netapp_trident

<a id="canonical-3211323111122200-2022231133310222-2111133212330300-1010113231101330-0021303132000002-3100101322300112-2032232202110332-3231102032113310"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for NetApp Trident.

Receipt-pinned upstream constraints:

```json
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
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232023211320023-2032232311032212-3011200021232330-3111210133333103-0222131330330230-0022331102321131-3123012221113220-1330003112221130"></a>

### Direct properties for `storage_class_list.storage_classes.netapp_trident`

- [selector](resources--fleet--reference--group-002.md#canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003): complete subsection reference.

<a id="canonical-1323200222023300-3202123201121232-3200302013333131-1202122303320021-0313200002321110-1313123332203113-1103122333132232-1200231333320033"></a>

<a id="canonical-2323020330323323-3311302123331223-2030203011131121-3322223311012102-3303130300221231-2012232310130222-0213013300101020-0001323131212321"></a>

#### `storage_class_list.storage_classes.netapp_trident.storage_pools` property

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.netapp_trident.selector` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-1230001331232023-1121200110230223-1320210312010322-2222111303221202-1030032033322312-1213022300230231-0123203120230022-3133132202222111"></a>

Type: `"object"`. single nested block, Optional.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Receipt-pinned upstream constraints:

```json
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
selector {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110101010012033-3333011200130233-2120103112121132-1312112202313113-2102222332321202-2231022203130123-1000222321310221-2201120133301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.pure_service_orchestrator` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-2102032013030302-3313002330122221-3211033121302013-0310311210011323-0200201031123010-2233100223203002-1333301133030033-2001102003233021"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for Pure Service Orchestrator.

Receipt-pinned upstream constraints:

```json
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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222323033110022-1133012123122002-3012003321110222-3220333100201333-1202212010033202-2322230222011030-3200013221312010-0321101220101110"></a>

### Direct properties for `storage_class_list.storage_classes.pure_service_orchestrator`

<a id="canonical-0210010221120302-3002023220030311-0113300021331233-1333323112113322-2210123130321230-1122320202312310-3212230103230113-3032301202032312"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.backend` property

Type: `"string"`. Optional.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("block",
    "file"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-1230013231223200-0211121221313133-1122130202323311-0030031320222111-3200232032013311-2022311232030123-1000230010020201-0003222313232120"></a>

<a id="canonical-2310200120132201-1212201313020323-0320221100102111-3022133113210230-2322031202301210-3312123202000120-0303003303031233-0232022112223122"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` property

Type: `"string"`. Optional.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(12),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-0112023103302112-1303020021023310-2211221203122210-1211032313103023-0130211133203211-2302013020100201-3121003232131030-0311012310302131"></a>

<a id="canonical-2121003331203213-0031011212032003-0123321323212132-2010120132331302-0122323311123202-1110032323323222-2311202010220133-2211013000000112"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` property

Type: `"number"`. Optional.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 100, Maximum: 100000000},
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
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_device_list

<a id="canonical-1220331113303212-2010311203330203-1123010023313010-2333311011310323-2013213130203030-0232022330203123-2033201030211020-3221200223133131"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
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
storage_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221323220213322-2233203232120131-0131232202302312-3113311032101333-3021010311033212-2102320011021000-2030210121121022-2303120022323122"></a>

### Direct properties for `storage_device_list`

- [storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221): complete subsection reference.

<a id="canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- storage_device_list.storage_devices

<a id="canonical-2303022122102011-3003130203301312-2102021200031130-1013321101232010-3022110212312002-1300222232122012-3110113302112132-1103221033211211"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133021332010120-1300030312220330-1332101010220030-0101201131132102-2100210220102300-1332003312131330-3330233202302221-3120020021131203"></a>

### Direct properties for `storage_device_list.storage_devices`

<a id="canonical-3310031021310201-1130031213203122-2222103030112210-0221123102103022-1031001031020313-1321333202030101-2223223023022022-0123002131313001"></a>

#### `storage_device_list.storage_devices.advanced_advanced_parameters` property

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
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
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](resources--fleet--reference--group-002.md#canonical-1201010313101013-1022332222103330-0131311131312023-1221012021200303-3022303010333232-3021021003111313-2030220111213022-0221030232013033): complete subsection reference.

- [hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313): complete subsection reference.

- [netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023): complete subsection reference.

- [pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010): complete subsection reference.

<a id="canonical-1303002102023230-1212021212311133-3012312033032013-2102222101303330-1220300031023302-0213130200100133-0320110322222301-1232210033213310"></a>

<a id="canonical-2313212233001302-0032101212223121-2110131210102123-3013310002203110-1220330310322230-3023201022011113-0212222321133023-1103103001212002"></a>

#### `storage_device_list.storage_devices.storage_device` property

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1201010313101013-1022332222103330-0131311131312023-1221012021200303-3022303010333232-3021021003111313-2030220111213022-0221030232013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.custom_storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.custom_storage

<a id="canonical-0100120110322223-0112330323012102-0130122221222112-2000033202332103-2211220000120230-1202321200321112-1220322023101100-1233121231010130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom storage.

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
custom_storage = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.hpe_storage

<a id="canonical-1300123020201210-1210323131311333-3010122233320022-0003003322221212-2133201201131013-2033320331211322-0332300022020123-2323012132313022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for hpe storage.

Additional upstream details:

Device configuration for HPE Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server_port",
    "username")}
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320331322033121-2310013021100013-3310020223221131-0311131100301323-3012320221300222-2113133232130032-3320233011111322-3212132110103100"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage`

<a id="canonical-0100121002203112-3202321112103103-1303012230123123-1032323131130213-0012031022031020-0232110000113021-3232201211102332-3120202220033001"></a>

#### `storage_device_list.storage_devices.hpe_storage.api_server_port` property

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](resources--fleet--reference--group-002.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132): complete subsection reference.

<a id="canonical-1113023311332300-1111230300333131-2323221112111110-0103131033112032-2210122222300322-1023223200313311-3010102233111331-2020111222300303"></a>

<a id="canonical-3211203012311032-2232102002303033-1000311031322102-1310001110323032-3010023303000223-2330003120000021-3122221100033130-1010213331231022"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` property

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

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

- [password](resources--fleet--reference--group-002.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102): complete subsection reference.

<a id="canonical-3130131021220103-2310130023223033-1212220103022231-1112133102113030-2332213330212211-0313231133112020-2322303022023331-1002122021033120"></a>

<a id="canonical-3202003000103010-3222112321203320-3311333002232322-0133220331113212-1233132031320111-2010210021223320-0222131132010330-3303300330211002"></a>

#### `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` property

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

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

<a id="canonical-1032221020101332-1220231120222122-1123312121002001-1022131000123000-3021310100101310-1101203103020213-2020210021223303-3300013310231100"></a>

<a id="canonical-1301221011330011-0321103110303323-0100001032023311-1313202120212110-3321000011110000-3301012122111332-2130233331312203-3202223322203210"></a>

#### `storage_device_list.storage_devices.hpe_storage.storage_server_name` property

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2121200030301312-3111212101221202-2323001030322113-3033230130121230-3323321113002213-2200201021313103-3303013001302212-1220303033211111"></a>

<a id="canonical-3032123030300110-1130223212311001-2110311130300113-1332110100200220-1132103201323010-1222011223032101-2112111230220320-3330133102121133"></a>

#### `storage_device_list.storage_devices.hpe_storage.username` property

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-1333200000001010-1022131030213331-0211121200013033-1001011211103100-2300121021230232-2202133313323302-1103002233203000-1032302020222310"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233333212033103-3222332222313011-3112103202331232-0223311000012221-1032303201032221-1223220111130011-1303010020313002-1132303321021221"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password`

- [blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-2110222221121130-3323000022312202-1222221312330331-2220310022003331-2023131232330003-0302030033031223-2332203030132000-0330011232232203): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-002.md#canonical-1321102021232213-1301130012100310-0102103120131110-2203033023233011-0120110210130302-2103133301000030-1210101310333003-3120322303012013): complete subsection reference.

<a id="canonical-2110222221121130-3323000022312202-1222221312330331-2220310022003331-2023131232330003-0302030033031223-2332203030132000-0330011232232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-002.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-2011203201120113-1011203012002231-3321012033222312-3302200210000033-0101320221330033-1310220100320112-3313103113220210-3121102031033020"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130031120010031-1120332103211220-1022312222103131-2211322300311100-3021212312201222-1322122033102330-2102031321303332-1021002021301330"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info`

<a id="canonical-0310223111332103-3122221332110000-1101112020123021-1003212223101213-2222012210003310-1121120200303201-2031233002322332-3013330032133300"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0221310303202313-2121011001201312-0323032021231120-1231303310202303-0000212233133322-2133230331030213-3030000310233012-2130200100331231"></a>

<a id="canonical-2232230323300000-0332210220323123-1101122323333103-0301032032300103-2232100131323101-1200213100201203-2121012020031201-2131331211121100"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3311001322212321-2333002021021222-3322120012331121-0300200100032020-3102113012132020-2302233233033331-0112031130113223-2031302102120003"></a>

<a id="canonical-3122101330312231-3002213322212032-3302032010230320-2013303230032202-0122330121123000-2213233223130113-1311220100030331-2213132210020030"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1321102021232213-1301130012100310-0102103120131110-2203033023233011-0120110210130302-2103133301000030-1210101310333003-3120322303012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--fleet--reference--group-002.md#canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-2031303221102222-0030111031002030-2333101200112203-2000220011111223-1013231103320302-2220212302122131-0020002230132303-0210001210303021"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013322110310211-0102203310221100-2223331001012133-3320221110133213-2110311032012320-0110011301112032-3330213231310003-1131022113210233"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info`

<a id="canonical-1311031210011103-1312322211013031-2110122020221302-1303332022002303-2222122113131032-1123321300103331-2320031132310001-1302302232112203"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0023011130010020-0210202013100310-2302312101333331-3012210122330131-3330211000123200-3131033000332313-1331211013220023-2212023302123100"></a>

<a id="canonical-0232100203200101-0332311322313233-1133332001333323-0311020133230203-1003033222101211-0211110003320303-3120223002211022-1013122323300033"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-2131003020012222-1312020222122031-3033303322002020-2203010302123011-0230321123321121-3033331222100311-1330300022110313-0030033002310200"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021233212112130-1120302323030002-0213033232033223-3232220233200012-2033331000313302-3113100321000021-2303303213132011-3022131031210232"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password`

- [blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-0010013002333102-3332313022212223-1012222103020322-1300001011331221-2103322322102322-0130010331210030-3102131110222030-1313111002102211): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-002.md#canonical-1133000301131223-1230300011311221-0101232113102020-0200321001220323-0203213311200330-3102102112332231-2222313113210133-3102322100230123): complete subsection reference.

<a id="canonical-0010013002333102-3332313022212223-1012222103020322-1300001011331221-2103322322102322-0130010331210030-3102131110222030-1313111002102211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-002.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-1311211111133200-1013113321000123-1121102322020233-1202023132103002-0211322231023231-3112300311030220-3000321101121103-0111312132002321"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103210012013333-2331011200323022-1033312120201332-0202000030003302-3111020321221322-3331223310322133-1002110030010212-3112310302113013"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info`

<a id="canonical-1230200211213203-0023103220022320-2012121111222312-3202203202002331-3213332032113330-3112200111232331-0323033022111020-1122232100033103"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-2010013323212311-0312100003202302-2313221011112310-3221130010001122-3331321232001000-0013002322330120-3311002013321112-0332130332112010"></a>

<a id="canonical-3001312011103323-2121112003301220-1123332202310313-2332111003213013-3100012231103000-3130100033020332-2132300213130122-0201301321031302"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1331332100110312-0032030030111020-2021232333003123-0233203113223031-1333032232123212-2020111333300310-0020130112231033-1230102311132033"></a>

<a id="canonical-0102213212220322-2303302131031331-0323301122023232-3123002331302110-3100302213122232-1220223011120300-2312133102031333-0010332131223300"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1133000301131223-1230300011311221-0101232113102020-0200321001220323-0203213311200330-3102102112332231-2222313113210133-3102322100230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--reference--group-002.md#canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313)
- [storage_device_list.storage_devices.hpe_storage.password](resources--fleet--reference--group-002.md#canonical-1212101202211001-2200100120130110-1332213110330321-2213301011202321-1111330332131000-2222221012033230-1010112133212033-3012203021132102)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-3312223300230222-2221102330132003-3020131323013102-0132313003022123-1232103332312031-0003013222032311-1012223300320022-2031312120310132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210213303033223-1223123010120100-2322310103211103-3103131000221330-0322033213313331-1100102213310203-1113223300212001-1313321112020121"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info`

<a id="canonical-3122220201313000-0131330200201321-0302111212222202-0301312121203231-2002330300211103-2030011333102333-3101122221223110-1111212300012332"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1323030021203330-1133231111120013-3020313012220331-3330122100131201-1303332322133232-2220220002322022-2230322100011010-1000202301231102"></a>

<a id="canonical-0122320322320233-2310200130132103-0211213220211232-1120202232221302-1322323223122201-3120032131001211-0122103103123202-1211033211333301"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-3101021010223311-2100233103222002-2031222103120330-0032130300130100-1032231200121110-3313101222102111-2301311133330031-1001330210033033"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231333321313111-2333201130233302-2201230010300203-0320331320020212-2023320311231202-3122331302200020-2132233022020112-0112303202110213"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident`

- [netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131): complete subsection reference.

- [netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101): complete subsection reference.

<a id="canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-1300011332320212-0313000331011031-0302220301103130-2120310231022210-3113120211122333-2233132200230103-0133012320023111-1003312123221132"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300323203020332-0210001200332133-1220330312013203-1333312033323011-1010101133110200-2013222032211002-3331330133331300-0231211023313331"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas`

- [auto_export_cidrs](resources--fleet--reference--group-002.md#canonical-0011311003202231-1223000031133112-0130132012021312-2120200332320230-3122232220212010-0122313230210323-1103223303212302-0100333200330320): complete subsection reference.

<a id="canonical-1231313321302020-0000021011303123-2330330331320010-1223023221101211-0133213231123203-1131200031300320-3033310031323100-0112203230012230"></a>

<a id="canonical-3302211121101212-3132220222212110-0322231113131213-1033100121102132-1030323221121121-3002021322231203-2210200212221211-1120013310000123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` property

Type: `"bool"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Enable automatic export policy creation and updating.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1210031232131302-3211232322012102-0312312001221211-1303202120233112-0220230011320213-3122111223212303-3210013310230200-2023311101200233"></a>

<a id="canonical-0022030101301030-3021320100030033-1333210111302031-0302133123323233-1330330201002130-0103303321131333-1111300313231201-0201031310310102"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` property

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Additional upstream details:

Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3001312332233233-1101010101222020-2220131133121111-0220001230130021-1011023110132210-2331213312311031-0033300300221311-0000100202000220"></a>

<a id="canonical-3003132112211212-1202232320220331-0111023132202230-0030302230213222-3031320002132312-0311212333012002-0031222333002131-3021221132013303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` property

Type: `"string"`. Optional.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--fleet--reference--group-002.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221): complete subsection reference.

<a id="canonical-3210222313333120-1212302323213103-3133313032220232-3301321131112311-0131030311013101-2201212122100130-0301203031102202-0010102220002212"></a>

<a id="canonical-3130201230312012-2032330210221013-1022103131132031-3313221313201220-1211112213213100-2313031120330323-0323012303320312-2130311320212022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0011310011100021-1012023130231030-2300322111100111-2311310221223212-1232220332133103-2331331030102111-0301313200131123-3302320021212020"></a>

<a id="canonical-2301101103002233-3300321023203313-2110101010112031-2233031003020020-0322010000110320-1012032030131331-3311012030200032-2232122230132211"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-1131321231021330-0321000230001110-2123300312321330-1302102112312221-0131313231010120-0020130003000031-1010310322112033-0132022123002232"></a>

<a id="canonical-2332113321201213-2320220302200010-3321032003231000-3312111131222233-3223212301132132-3023003312002230-0332130021333120-0331032003111200"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
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
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1212012020302331-1100101013232022-0233323020333023-3322001031010031-2322102212323002-2030032231220110-3101022310322303-3320002213012332"></a>

<a id="canonical-0301001203111202-0223202113332010-1132030223223122-2120120032111212-0020003120122221-0313200133213023-3122301012033230-1323020210000130"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` property

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-2012300211032020-3100103233101013-1112312020230110-2232202023230000-3331130312300221-1031222231120030-1122333202112311-1311303202112303"></a>

<a id="canonical-1230002221002101-2111123321012113-0331333222202312-0123011221220231-1111302323000212-1132301103331010-1320323031000133-0223012223331023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` property

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-3301201200012332-1032312121230011-3010112210201220-1002102302101201-2122102002110231-1013031101220221-1011311032313100-3022100332012303"></a>

<a id="canonical-0213101123211033-2301123211020022-2333232312121212-0102311201310111-3323302321200311-0303020030312200-1010220013223310-0113200223131222"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0130113221022031-0231200122210132-2030110123103212-2221320021223333-3023300223020032-1213220212331122-3020101001201101-0110330111110102"></a>

<a id="canonical-2223232032212033-3312022203320311-2003202302223102-3010031003213232-1120111232110210-1202030330320123-2103331112002213-3123032200132301"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

<a id="canonical-2322331212203101-2032212003322321-2303333321222121-3213010111101213-3103221220001131-1310221300121132-3131331323311222-1000323120210310"></a>

<a id="canonical-0001021101113300-3223220211210301-0323221232310321-2120012121003031-1012000033111233-2130230021132233-3123310303200300-3332213222211200"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` property

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](resources--fleet--reference--group-002.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033): complete subsection reference.

<a id="canonical-3331213221131132-2023230121203212-2003000331121231-1131000100303122-0203332230231322-2211313103200331-1012110002220312-0000233210031123"></a>

<a id="canonical-2001201312031221-3111232021010323-1003123012323302-2010000010202322-2201200200131101-2111000332211003-3111313330300002-1222330130000130"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` property

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300): complete subsection reference.

<a id="canonical-3030322130013021-2111301300020323-2133223002133323-3201132312233323-3213023110301130-2011320110230132-2031332112331310-1311031120333330"></a>

<a id="canonical-2110333203331223-3323233030330301-2023101112010321-0000130010023220-0222311301123032-1003130330211230-3333121232121221-1030103022033212"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` property

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0232321133002011-1032223232113020-2012130130032112-3231232020120031-1021011101313130-1332203011210311-1333233012303100-3331303030013213"></a>

<a id="canonical-0120111131021121-2011003112323121-2121021223223203-0201032323133130-2122132003223003-3231310220230231-3312213101012230-0031003023030023"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` property

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-2100310330303223-1020310133303103-0313303300201020-1332121130232322-1111111122320223-3321120302111311-2111232030101112-2002202020311133"></a>

<a id="canonical-1130021230121202-3302332133310233-3310301201102130-1202301102210332-0102131022022232-2300110103212130-1103012011120131-2310101022333110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` property

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3320120122022020-3010323023113330-1330001220230111-3323222113112202-2330103233020310-3103030300323011-3012010200313232-2330213320122230"></a>

<a id="canonical-0202301100202132-0101300123113003-3201310121011211-0110331022101201-0322031113120010-1130312032203133-0100110132230000-2111111200311012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` property

Type: `"string"`. Optional.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-3012113013132330-1331210322023311-2031013220130030-1123311013230221-3210313012301333-1020213220003201-3030032222232121-3000022333031310"></a>

<a id="canonical-3333311221132103-3030332123103122-3020311201032311-0231212212203020-1131013223330201-0110320321202311-3023220131023231-3011010112023131"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` property

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331): complete subsection reference.

<a id="canonical-0011311003202231-1223000031133112-0130132012021312-2120200332320230-3122232220212010-0122313230210323-1103223303212302-0100333200330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-1021110303213110-3302031121103320-1332312323011213-2212202132032211-0321032102312102-2311210101332031-1000233032220111-3130222013121311"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
auto_export_cidrs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220103113233210-1111011010010113-0302222232313101-1113322232303312-0102122122303013-0011130133321130-3121033333200010-2223331211313023"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs`

<a id="canonical-0122131302011200-2233200123103300-2303003303103203-1112113321200021-3323333011212210-1030320002323023-3103333102210313-0231120222201103"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-1230113112310231-3223032113113331-1001113333123330-1010323232000312-1233301023032323-3102323330023311-2310212310031311-1033113201201300"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201321001300232-2331122110011232-1100312103313012-2110223321123002-3213320222231311-0232122013130132-3322231331110203-2111131223223131"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key`

- [blindfold_secret_info](resources--fleet--reference--group-002.md#canonical-0331023132232112-0331320023011112-0333332032213010-0300022123213023-1213202320223202-3303002001220113-0003321131232310-3000320032023032): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-002.md#canonical-1321101120300120-2212323121331132-3201213113030033-2101133022330202-0110030321310301-1333313310331212-2011120330113131-2030020001022003): complete subsection reference.

<a id="canonical-0331023132232112-0331320023011112-0333332032213010-0300022123213023-1213202320223202-3303002001220113-0003321131232310-3000320032023032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-002.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-2110132010000223-1001012310202032-1013231232032230-1202010011133322-1020200101303101-0222211103111022-3123010101021331-0200321000110323"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030331110321100-3113112131321011-2020300200030020-2332223120212120-2022203210111031-0333133213023211-0310211012002311-3310200331002130"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info`

<a id="canonical-0200303202212032-2320201022103332-2310031120131210-2113231302023112-1222302000123023-1101022303131223-3002011223210222-0101310212332331"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-1003221012303232-1222013221030303-2020113122130211-2322212002002133-1131111002032003-0302123331322131-3033212330231330-0322113020131300"></a>

<a id="canonical-3031120231320210-2330212111301323-3112130201132010-1321111130112102-3321013312021130-3233233012132200-3310301103332032-0310020211231302"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3110212323130123-3121320311010202-1201103020021211-2333121323312300-2102201223211113-2103331002020330-2003123210121313-1300101120112022"></a>

<a id="canonical-3303133132032331-1031232223122312-2331023200303202-3333103032333220-1003321111130333-2022303302120101-1213110310323203-1113022302100100"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1321101120300120-2212323121331132-3201213113030033-2101133022330202-0110030321310301-1333313310331212-2011120330113131-2030020001022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--fleet--reference--group-002.md#canonical-2021123321033031-2131220302333132-2013232223201100-1211122103213311-1123311311323332-1333300211323100-3332210130000322-1110303130220221)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-1301333012002230-1211121232220303-1001132330013210-1110312130033102-2111003110212300-2213310312211100-1210130322013100-1303032200332220"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213303123203202-3033122032013031-2201010021200023-0111233231311322-0033023013232130-1322020101121331-1102123130033030-2213011103313103"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info`

<a id="canonical-2212230131002122-2032211110233122-0111313302333023-1012210200303320-3113023123123120-0221220202333122-0132330021020022-2221110303232020"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3303021030021130-0302233013003223-2020021202013333-2222030310200212-3311333212013013-2032302133233233-0132212230300331-2310000023023131"></a>

<a id="canonical-3020300310102131-0112201230203132-1223100231203113-0012133223221130-2330133103313202-0313030102032102-2333222323002012-2312130233313323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-2310201301301300-3100121212331033-2232033100203311-1330302310003123-0232233021333003-2200330123100132-3300122313002212-2100311033121231"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212222021213321-0021333320132123-1033312323221231-1211331212213122-0321221213203113-2001231321221323-3220330320213210-0223213031321303"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0313031320220001-1022121101320112-3101002202203311-0230200323122201-0100310223000021-2112000302223133-3231231321013323-1120320102230330): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-0103103020330320-2211332020322103-0111232101323212-2320200313030001-1201233210331002-3331223222012101-2133203032230302-1232113003313111): complete subsection reference.
