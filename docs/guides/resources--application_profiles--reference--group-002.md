---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-3012033202323021-1022322102201000-2212332322100103-0201333313111121-1122003311231001-3132012311120222-3323322322032102-3200101300320330"></a>

## Direct properties for `virtual_server.connection_rate_limit_mode`

- [per_destination_address](resources--application_profiles--reference--group-002.md#canonical-0311220212331211-3011202012101113-0232011333222301-2203212100122311-0131333201111200-3210010320312213-1012131321030100-2210101133000220): complete subsection reference.

- [per_source_address](resources--application_profiles--reference--group-002.md#canonical-3112311033110011-0231210000133330-2311033311023222-3312112003323111-0211331222300131-2100100002310130-3022112310120230-1310100121122011): complete subsection reference.

- [per_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0112133010331231-3221202200001131-0320310312102001-2223300122300312-2222302002211231-3033121312002302-0330102121121023-2331210003103123): complete subsection reference.

- [per_virtual_server](resources--application_profiles--reference--group-002.md#canonical-0002110030200031-3021230220301233-2301020300210212-2012003220332212-2101023222030132-2332012301011100-1321330200112001-1110103230220112): complete subsection reference.

- [per_virtual_server_destination_address](resources--application_profiles--reference--group-002.md#canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220): complete subsection reference.

- [per_virtual_server_source_address](resources--application_profiles--reference--group-002.md#canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201): complete subsection reference.

- [per_virtual_server_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0113212332311100-1310301131000332-1202003101033103-2322132100123101-0101000111312210-2220101121121130-3032010102130123-2011222003213122): complete subsection reference.

<a id="canonical-0311220212331211-3011202012101113-0232011333222301-2203212100122311-0131333201111200-3210010320312213-1012131321030100-2210101133000220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_destination_address

<a id="canonical-3313232302320002-1322131102311312-3202132010023031-2321130021120013-2003133200130013-2122320003212321-3020011300103011-3302222023132003"></a>

Type: `"object"`. single nested block, Optional.

Destination Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110021330133333-0211331320320212-1310113030101011-0322200330011313-0132013210230301-3122213130322132-1110101201023312-0302210123301322"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_destination_address`

<a id="canonical-1011231323321110-0220010012111331-0001312032130211-2230012002022120-2203231322201101-2003032111331110-0222310201000131-1002031220203333"></a>

#### `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` property

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3112311033110011-0231210000133330-2311033311023222-3312112003323111-0211331222300131-2100100002310130-3022112310120230-1310100121122011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_source_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_source_address

<a id="canonical-1032313202203331-2213320301221002-2211210101321220-3133123230322220-1301210201332131-0321031301121220-2332221203212311-1323130120131323"></a>

Type: `"object"`. single nested block, Optional.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_source_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202301133000223-0200131200332310-1300000322230233-3221333023211203-2032000312312312-0033133203330330-2213210130300003-1110012302032012"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_source_address`

<a id="canonical-3030132331223231-2311121121311021-0023330023232220-2130321003300202-0321323130233121-2333301200322333-2132323112102003-1021023322120103"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` property

Type: `"number"`. Optional.

Configuration parameter for source mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0112133010331231-3221202200001131-0320310312102001-2223300122300312-2222302002211231-3033121312002302-0330102121121023-2331210003103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_source_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_source_destination_address

<a id="canonical-1121321010302331-3210012021133201-1311000332111023-0102212333311320-1302133003221333-3332301131221312-1322321301203313-2200323113202003"></a>

Type: `"object"`. single nested block, Optional.

Destination and Source Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_source_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023212201103031-3222302020022033-3121203130011010-0111232300300022-1300120203220130-2311222221320333-1113301100012121-1220321301110133"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_source_destination_address`

<a id="canonical-3130120312101122-1301101010333203-1321211202203331-3233022013120312-1312002132212303-3103202112111212-2310220211133201-1222332310111100"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` property

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3131112320301313-1101232211113011-3001121131001302-0101312032323111-1000000131131032-2202312000332200-1131023233110021-3031121111323311"></a>

<a id="canonical-3112333113010202-1300121321230031-0102001110303210-2121321301122010-2002133213320220-1131033003121113-0332231222230203-0121110022121303"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` property

Type: `"number"`. Optional.

Configuration parameter for source mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0002110030200031-3021230220301233-2301020300210212-2012003220332212-2101023222030132-2332012301011100-1321330200112001-1110103230220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server

<a id="canonical-3130113030210232-3312223213230030-2322321202130103-1302122101233331-0120310132320003-1010003012102203-1003023002131030-1022333022323012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for per virtual server.

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
per_virtual_server = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

<a id="canonical-2012022113300113-2301301113011120-2020123111303122-1010131133200020-1212100222003022-2110131011230123-2303210200233110-1300223013230020"></a>

Type: `"object"`. single nested block, Optional.

Destination Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_virtual_server_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310313320001012-1202300300320330-1110111012130333-1113011002010120-1112001012300300-1012221313102223-2312233331022313-1103312100101300"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address`

<a id="canonical-1000011102212103-3113131323031300-1033321133312130-3011111121312303-0032332330103213-2130301222003102-0232130310131322-0012233033333233"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` property

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="canonical-2101303130210111-1220102123110101-0231331023010311-1203131013010310-3023132221000003-3032200132201300-3222303021122200-0133330120331001"></a>

Type: `"object"`. single nested block, Optional.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_virtual_server_source_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321320330123021-3321103010301133-3123131132031122-2313310232233320-3122021011110203-1021112233312311-1111302333231312-2301010200003312"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address`

<a id="canonical-0333231212030220-1021031331111010-0321211202222011-3221220033112032-2233122210303320-0132212011202312-3223331303230003-0201112110123333"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` property

Type: `"number"`. Optional.

Configuration parameter for source mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0113212332311100-1310301131000332-1202003101033103-2322132100123101-0101000111312210-2220101121121130-3032010102130123-2011222003213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-001.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address

<a id="canonical-0331310010022310-2302213032223221-1223232003201100-3021010332102311-2211130113110022-0223001011223003-3212310223321212-0023232130313231"></a>

Type: `"object"`. single nested block, Optional.

Destination and Source Address Mask.

Receipt-pinned upstream constraints:

```json
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
per_virtual_server_source_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212002201223103-3210210032001132-2131000202202333-0200013101122212-2233123221211101-3020222133030133-1132132101110201-2010030331303013"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address`

<a id="canonical-1131102111332120-3131131320131203-3203120021321221-0233121222103203-2213311330001211-2100023100332312-3133011120012220-3321112230102313"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` property

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1231033213102212-0020130033122133-0330222102122121-2211323131230012-0200201101111020-1020212310313223-3322201022120002-1000312130312113"></a>

<a id="canonical-1000223030213030-3332032121012112-2001110301222322-1130211321111001-2333303311323203-2032232311221332-3121123221233200-1223133332121300"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` property

Type: `"number"`. Optional.

Configuration parameter for source mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1130200120303132-1310331020320222-2223332303331002-1001311303211331-1130223022103331-3332031011100321-0310011023011303-0330123013203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.default_persistence_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.default_persistence_profile

<a id="canonical-3003330212123022-1202023302201020-2222131130302220-3103112023103232-0120013021311210-1301120010220333-0031231203030233-2001000121103120"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for default persistence profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_persistence_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100133310032103-3002120311011123-0002020122020211-2033312300133121-0203030212122322-1131131211330303-2331021203221003-2221220332010030"></a>

### Direct properties for `virtual_server.default_persistence_profile`

<a id="canonical-1032030113301233-3012310320011110-2321311321123130-1010033223211110-3002330031100210-1033101231011032-0021331301222213-1312000100200102"></a>

#### `virtual_server.default_persistence_profile.kind` property

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

<a id="canonical-1013302000100300-0300223331100030-2301030003001313-3211212333200301-0220220033023021-2132120331210312-3201021221330201-0231103211111301"></a>

<a id="canonical-2212330122330123-1110120203033103-2310012300231333-1121100203222233-1320122100022002-2132023020202111-3032013123000003-0121123130011323"></a>

#### `virtual_server.default_persistence_profile.name` property

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

<a id="canonical-0123110133310023-2330220232011220-2010113312321210-1213112030223011-2031102300213301-3331013321030323-0212110123021301-2022130110313013"></a>

<a id="canonical-2312113011003231-3223202220300332-0133130321302233-1033113322320313-3300331232032020-1113121103020011-3332311023012221-2133130010302132"></a>

#### `virtual_server.default_persistence_profile.namespace` property

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

<a id="canonical-2132010111023003-1310103300310111-0023222132330330-3203010001133001-1112031101200301-3312011210302002-1201102312131100-2131122132111300"></a>

<a id="canonical-1203110303100311-3130000133033233-2230301101113103-0330031133123013-2120122203023223-0002330023101300-1210312123203211-2333000020022111"></a>

#### `virtual_server.default_persistence_profile.tenant` property

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

<a id="canonical-2302120021022132-1013123122013002-2213233113310232-3033120213200102-1032003123003033-2103113033103000-3223303121211332-1320121123222331"></a>

<a id="canonical-0230232123210320-1130223002021103-0312310002210220-0202303012200233-2032302112300000-0011302000313320-2100330003003321-2213020021213121"></a>

#### `virtual_server.default_persistence_profile.uid` property

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

<a id="canonical-0220212320310233-0020030230010211-1101200002230023-3233033023031122-2320003100323101-1300230000023000-1322202320211120-0101202210330311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.default_pool` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.default_pool

<a id="canonical-3323300103233100-2200233303311301-1122232232202202-1300201201233111-3320200010030210-3001213020011320-2330222130011223-0013330220133002"></a>

Type: `"object"`. list nested block, Optional.

Specifies the pool name that you want the virtual server to use as the default pool. A load
balancing virtual server sends traffic to this pool automatically, unless an iRule directs the
server to send the traffic to another pool instead.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021010310231133-2133233201030321-2200003210030333-3010013202331222-2122133200332213-1203001220203020-3131100330331123-0213333330332322"></a>

### Direct properties for `virtual_server.default_pool`

<a id="canonical-3233011302111323-3010331222123322-2113221113201212-0300101020323000-3113313031321130-0131030001233023-0213100200220203-2230003310200201"></a>

#### `virtual_server.default_pool.kind` property

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

<a id="canonical-1002132220121110-1130032020003033-2002033121132232-2120023033003231-0022110202310230-3332022011103131-0213321001113130-0200221001031032"></a>

<a id="canonical-3020210033012120-0021121310033331-1213003210012233-2202312313001320-0000003222023113-0000330133110100-0312031322323332-3111322112020302"></a>

#### `virtual_server.default_pool.name` property

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

<a id="canonical-0110023303023330-3310132013030213-2011022132022302-3230220213233313-1301133323333101-2033203332112300-3333200010023021-0213013001211022"></a>

<a id="canonical-0320021230212222-1031301031100212-1121223203110120-0231223320002231-3010230300011222-1113031011322000-2022312100212321-1320021200321203"></a>

#### `virtual_server.default_pool.namespace` property

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

<a id="canonical-3302123123312220-1320322313001213-0101323301020221-0202312333130231-1110112320010232-1200110201300211-1012213220310203-0033010123332220"></a>

<a id="canonical-2211213010201311-3023033113110101-2033230220001321-0020333031332222-1200310303121000-0313132330213022-2021311203013322-0220313121332013"></a>

#### `virtual_server.default_pool.tenant` property

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

<a id="canonical-0012023333200002-2030020322213323-3012222013333221-1221232000013020-3222212301313023-1313110223331002-1302313223301110-3332300221000103"></a>

<a id="canonical-3133231232203023-2122200030213011-1200120032203203-1323300313311330-1233222133010300-2333303201001312-0230331322110302-0123102300022113"></a>

#### `virtual_server.default_pool.uid` property

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

<a id="canonical-3103321022122021-1320313333200212-3310120231330011-3010133213120203-2222103302302221-3011000212110223-0013202223013102-1230032013233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.fallback_persistence_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.fallback_persistence_profile

<a id="canonical-3021300023002001-2000330111223011-0331013022010230-3330313113132111-3323032232021003-3122301013122023-2212102223011130-3300130111221331"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for fallback persistence profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
fallback_persistence_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230230211102201-2311211002302123-1103033213133123-3011203332130023-0101100233323230-0101131203213000-0211232033211132-3022020220301131"></a>

### Direct properties for `virtual_server.fallback_persistence_profile`

<a id="canonical-1330112230020101-0021310312221331-0103121112122020-0210321112111223-3103001323110320-2032210302321233-0011320310231023-1030033230022203"></a>

#### `virtual_server.fallback_persistence_profile.kind` property

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

<a id="canonical-2030221001130301-3122031011310203-0303101300031310-0000132113321210-1223121302130113-2012233200300030-2301032112033000-0112101012230311"></a>

<a id="canonical-3102310020023133-3311130303121332-2200230103200020-3010001003031131-1221222030312230-2222000201310210-1321202011223002-2203331002331010"></a>

#### `virtual_server.fallback_persistence_profile.name` property

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

<a id="canonical-1123323201121001-2013133301202012-0113102110322300-3331131322120030-3231222132013022-2323000332133031-1301002101110230-3121330230102010"></a>

<a id="canonical-2303133321320230-0103132221130112-0211030321321132-2200101021032000-2312302021333002-2332103013211013-0230203202233210-3032230010020311"></a>

#### `virtual_server.fallback_persistence_profile.namespace` property

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

<a id="canonical-1333323303123233-2212133033033032-2120232102123002-1220002122200002-1123310211023213-1031113002012221-0120113120223010-1111003031133230"></a>

<a id="canonical-0220300320031111-2202111030013013-2302312300330010-3211331321023210-2332213020301001-1001012302212323-1120002323033022-3103332012222322"></a>

#### `virtual_server.fallback_persistence_profile.tenant` property

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

<a id="canonical-3301210113012320-2212320330000330-1011212000223311-1302222001220112-2032220300132021-0121012331200331-0131023223302302-3332022203003302"></a>

<a id="canonical-0103200103231132-0113000211113103-2202131310313301-2212132131301203-3322022320303131-0301002223301100-1231323133021313-2003222312201130"></a>

#### `virtual_server.fallback_persistence_profile.uid` property

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

<a id="canonical-0300032301211222-3312312313011013-2022223230130323-3110133300220133-1121320101021121-2303320221132211-1112303221011231-2020012203010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.fix_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.fix_profile

<a id="canonical-2220330223223032-2003200010320322-3132202123100121-2211120320000322-1013120130311031-3101111201103333-3311130121031320-2122213302333231"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for fix profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
fix_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113302322233233-1133231112021322-0022100030323123-2001103102111331-3213001102312102-0102110112100300-2033110023221123-3130132113201332"></a>

### Direct properties for `virtual_server.fix_profile`

<a id="canonical-2332022332123013-0013231020120012-1231010201023101-2132002021110223-0001312112002302-1330333131031221-1210320030303330-2122211033200233"></a>

#### `virtual_server.fix_profile.kind` property

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

<a id="canonical-2233022010323031-2220201322202013-1002120001010322-0010201330310101-0200120100333021-1333030301222310-3033101302323022-0321333123322100"></a>

<a id="canonical-2313202233303020-3313320232111201-1211330213211110-0102123312133202-2001213232202300-0120021000222233-0113332021322300-1010011313332032"></a>

#### `virtual_server.fix_profile.name` property

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

<a id="canonical-0203032130000311-2321212301033301-2331313011220000-1003032130330230-3330310231333222-3121233002203331-2231213300212122-3213330112031323"></a>

<a id="canonical-0203111303120132-0003000130130300-3210212323222311-2031212023202321-0303302103213320-3222331121312202-2112302112020003-1001013101001032"></a>

#### `virtual_server.fix_profile.namespace` property

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

<a id="canonical-2121121222100033-1120220212111320-2310301202130120-1222201030031022-0130123120131231-2101223003221201-0301320022220202-3323110232102322"></a>

<a id="canonical-1303333202122330-1322002002013100-3030220331003323-0100211202313202-0012220222311013-3030112132312312-1310332001113110-0313321222230100"></a>

#### `virtual_server.fix_profile.tenant` property

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

<a id="canonical-1202223323102200-0012100002112132-0012200300132221-1313220330133232-2130023120233002-2230212102100031-3023333312023321-0022200132010232"></a>

<a id="canonical-2100212233212003-2232020313202031-0210132033130332-2322223321103130-2320333320012112-1232010112102300-1002312112213201-3210022123213223"></a>

#### `virtual_server.fix_profile.uid` property

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

<a id="canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.http

<a id="canonical-2332010000332020-0110331212332300-3330311200330200-3222000010231313-1221311303330220-0313132223022231-3301130303131001-0033103233213202"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
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
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313220031113031-1110133233301202-3101102331231310-2032233212100122-3013033231132030-0102131013223123-0303201000331212-1313030130103201"></a>

### Direct properties for `virtual_server.http`

- [client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3130120121301111-0232122113022331-0133313310022200-2330232310032203-1332022200120302-1321130233030333-0131333010323133-2230331203221003): complete subsection reference.

- [http2_client_profile](resources--application_profiles--reference--group-002.md#canonical-0332330302121333-0200202323301323-2011232011103020-2101231222032323-2011023122121013-2201030101010301-0301300001200132-0033123212100321): complete subsection reference.

- [http2_server_profile](resources--application_profiles--reference--group-002.md#canonical-0022331112021010-3110203310013002-0212322031113123-3313230120011121-0103101002003131-0033200323322031-0033010030130201-3323002202213230): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-002.md#canonical-2122303303313200-2013130123113121-1031102010022023-1131131323001322-2322123310020201-0021023332013321-2103212212003311-1320211331131022): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-002.md#canonical-2322322120023310-2331230320202223-2321331132010123-1120303200320012-2132131011333223-0311000303222001-0111212033133333-2113032201233303): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-002.md#canonical-0212200100013300-2010330003202013-2113300023302120-3201222310221023-0313220131110103-1312321001211122-3021120111312003-3200020321033213): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3232011302112010-0200312011113331-3101111203223133-3020113313300300-0023113130010221-0133232012023323-1002302300221201-1331300202001202): complete subsection reference.

- [stream_profile](resources--application_profiles--reference--group-002.md#canonical-2003122102331011-0102223311211013-1113200013103313-3101102121131201-2101310003131021-3232021132033212-0313132302121023-2222132212203322): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-002.md#canonical-2223311101321210-2223313131302002-2310010033333222-2122013110101211-1021221123210103-1210221022002331-2101300030310020-2331210030200231): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-3313330121201111-0233011211100030-0312301221013133-0221322322010322-0221231321001010-2010112300003100-1220230110101232-1323330130100112): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--reference--group-002.md#canonical-2103003320213123-2333013120322010-2312232213313300-1031221031100011-3332121220121033-0110233202012310-0011130201222110-1123332212132030): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--reference--group-002.md#canonical-2133310101022000-1330312111332132-3210102331201100-1113232033330012-0200001232322231-1010212032131322-0301022332301131-3030020121130200): complete subsection reference.

<a id="canonical-3130120121301111-0232122113022331-0133313310022200-2330232310032203-1332022200120302-1321130233030333-0131333010323133-2230331203221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.client_ssl_profile

<a id="canonical-2003221132223331-1212333232201023-3211311333310031-2221202111301221-1323210013230332-3312220301221300-0303321020102203-2122120110233333"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023213311221322-3021210021013213-0221233130220332-3100220223031030-0213230133033330-2302230030310223-0021132010030321-1221200203023231"></a>

### Direct properties for `virtual_server.http.client_ssl_profile`

<a id="canonical-3300332013002312-3210111301031111-3012202101030120-3211323013131232-1312222001320310-2211103133103111-0100030301231002-0323210022012010"></a>

#### `virtual_server.http.client_ssl_profile.kind` property

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

<a id="canonical-2031111110320233-0033303223312000-0013213013110100-2001102310113302-2100311030223303-0200333220310101-3222220300022122-2131230211322113"></a>

<a id="canonical-2102100313200100-3033121113003321-2032022232201300-2213110203130002-2201031023312103-3012222001311110-2021322300333311-2031102002230101"></a>

#### `virtual_server.http.client_ssl_profile.name` property

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

<a id="canonical-3333032332002022-0023210022110332-3202002121330201-3123123001320133-3000221013122223-1222221202331031-3211103312010311-1101121200102033"></a>

<a id="canonical-2123311032221023-3130220302220101-3123322001322301-0120113202010032-0120322311121333-1123111222230113-1232030001311300-0001222202333231"></a>

#### `virtual_server.http.client_ssl_profile.namespace` property

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

<a id="canonical-0201312331013222-3021233111130210-0310113000300213-2310122230233121-1211210333210031-0320203032213200-1312031203132210-2200022313012022"></a>

<a id="canonical-1332223121312201-0121321100003302-0213201013200110-3201010330330113-1311123313302112-0102002102301113-0212230031022123-0031302113100021"></a>

#### `virtual_server.http.client_ssl_profile.tenant` property

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

<a id="canonical-2032321300031201-1102321000211001-2000213130132310-3231330313312130-3012223333302123-3102320013222011-3002233220302122-1212100222200310"></a>

<a id="canonical-3331201111323030-1203321010200100-0322010303212021-2322023212223333-0021210201331122-1001110302012200-0032201003023123-3021110201101121"></a>

#### `virtual_server.http.client_ssl_profile.uid` property

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

<a id="canonical-0332330302121333-0200202323301323-2011232011103020-2101231222032323-2011023122121013-2201030101010301-0301300001200132-0033123212100321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http2_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http2_client_profile

<a id="canonical-1022312112332200-0222223130011023-1211321323203210-1222020210311023-1120011101100232-2023112312302133-0012200131203330-1232112223113013"></a>

Type: `"object"`. list nested block, Optional.

HTTP/2 Profile Client. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020130001030320-3011010122313102-3311120002302100-3032301132323033-0202213101213211-2113300223001003-3031303302030000-1203332332200130"></a>

### Direct properties for `virtual_server.http.http2_client_profile`

<a id="canonical-3101002331323112-2223133331130222-3111323123123223-3111021303031302-1213213010321303-3300003323313303-3101322130311202-0103130121323332"></a>

#### `virtual_server.http.http2_client_profile.kind` property

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

<a id="canonical-3322131032121130-0313332202013030-1032331120313213-1011200231110213-1103222011113212-1013311131311203-1032300021000333-2033300322023010"></a>

<a id="canonical-0110011212021201-2030003022200302-3120003033231031-2210223110200120-0010213122211323-2330111001302103-0103011310111323-0332322310213130"></a>

#### `virtual_server.http.http2_client_profile.name` property

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

<a id="canonical-0333330322231131-1123322131032031-2311313301022110-1110211102003001-3211120023113332-2233233300331313-1302302210003320-1301121102112330"></a>

<a id="canonical-2122130123303200-2020130130200330-1310312331000131-1232122001123321-1002300202022311-2212101202033223-1221033033233102-1032120113321301"></a>

#### `virtual_server.http.http2_client_profile.namespace` property

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

<a id="canonical-2103320101311310-1013130302023321-1102003220102122-0132001321212111-3022030210201221-3211331321013022-1011331003111231-0032320230233002"></a>

<a id="canonical-3102313032220121-0302023123003230-2301102023202100-0302223320120121-3133112033200001-0232131320201232-1313022320000131-3310323233230310"></a>

#### `virtual_server.http.http2_client_profile.tenant` property

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

<a id="canonical-1121111232030031-2313113321020320-3211322311103002-0321310331030301-3201231100210230-3002203221331333-2001030022122330-2321022033330322"></a>

<a id="canonical-3203233132011333-0133330102332233-1322003031103320-2322222020323030-0000132023222211-3211322012022110-2213232332130320-2001311013101003"></a>

#### `virtual_server.http.http2_client_profile.uid` property

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

<a id="canonical-0022331112021010-3110203310013002-0212322031113123-3313230120011121-0103101002003131-0033200323322031-0033010030130201-3323002202213230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http2_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http2_server_profile

<a id="canonical-1021113223331302-0220202131331210-2123132120213301-3310212101201220-2111233331121033-0022122331121111-3201000211233123-3120021232320232"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http2 server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230331033013131-0310021313101022-1012211323102010-0003123301032203-0221123313223013-1023201110022222-0022032313011312-0130033223030302"></a>

### Direct properties for `virtual_server.http.http2_server_profile`

<a id="canonical-2221023020223311-0113331122232003-3123201221211132-0000330112100001-1313311120012110-2000001110113330-2112022021323332-3310022020330032"></a>

#### `virtual_server.http.http2_server_profile.kind` property

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

<a id="canonical-3032011303010020-2103020203002013-0100103303220001-2203233132100002-2312302030322320-0132213311010201-2211233302013030-2302330322212133"></a>

<a id="canonical-1302333133030032-1030003322113212-1032121221032331-3233231313202233-3100121001321003-0032120231201121-1212300232132212-1010303201020320"></a>

#### `virtual_server.http.http2_server_profile.name` property

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

<a id="canonical-3001103022033230-1333001103131111-1022222132003100-0113211001231031-2121321120001103-3220003103310033-2101212200111231-1010221213303121"></a>

<a id="canonical-2211220033000101-2011010202233020-0232300302311021-3101233330011013-3001321122002020-1013012331330033-3013312232033230-2200021101102002"></a>

#### `virtual_server.http.http2_server_profile.namespace` property

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

<a id="canonical-1031220312220123-1011233100132310-1230331123101311-1233120220201033-3223312101201130-0031002211102313-1312010113110301-0111010221020213"></a>

<a id="canonical-1322203112131320-0220013310212223-1312331213222032-0313132003202230-1202113031122313-2031131310312230-0230301231121100-2200311123210022"></a>

#### `virtual_server.http.http2_server_profile.tenant` property

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

<a id="canonical-0330130332102210-1110033022220012-0031132320113020-3303220320210011-2132210110312230-2003232211203320-3103023211322002-1010220333130301"></a>

<a id="canonical-0131102101112201-0133110000332023-2133131221120030-1201331103211131-3210000122123122-1111300330111222-1311133120131022-0210211031300010"></a>

#### `virtual_server.http.http2_server_profile.uid` property

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

<a id="canonical-2122303303313200-2013130123113121-1031102010022023-1131131323001322-2322123310020201-0021023332013321-2103212212003311-1320211331131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http_client_profile

<a id="canonical-1113100222021131-2011310210010022-0120023101322100-0330202311321000-3132200232113230-2323110010233002-3321311201232300-1222113233120233"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301302200300301-3202322203313131-0330030210211131-0102331223030321-3333222221310333-1231130203330212-2023031220230110-2023313110303202"></a>

### Direct properties for `virtual_server.http.http_client_profile`

<a id="canonical-0101021333322322-3303132201311230-2100031032312123-1333110312111133-2220121011121020-1301131002333022-1313130302222233-2223123131313002"></a>

#### `virtual_server.http.http_client_profile.kind` property

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

<a id="canonical-3220303301223023-2123211100031220-0200233111011000-1111223322003323-0100301203001321-3133110102230110-2101212210012010-0230313301323011"></a>

<a id="canonical-2321131020111311-3030232223012320-2200020001203320-1032322022112220-2130331233303103-0311211222232312-2031113101323102-2200301301111230"></a>

#### `virtual_server.http.http_client_profile.name` property

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

<a id="canonical-1030111303102331-1302321220333001-1201010133021102-3213303223210011-3231003312330102-1332331322310330-2201211132010113-1312111020233022"></a>

<a id="canonical-0001223020013013-0230200233133232-2212221122102321-3322332223010130-0033020223333112-3210013223103310-3102020023000003-2333111120023131"></a>

#### `virtual_server.http.http_client_profile.namespace` property

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

<a id="canonical-3300100233013102-1012010320002031-3302131002012222-1210310113300010-0220032232010232-1103131003232011-0321201022302000-2112030031233100"></a>

<a id="canonical-2010001100031221-3201332111331100-3231000023022201-0133310110030002-1020232021022211-0300221312312121-1101010020203310-0112123132300221"></a>

#### `virtual_server.http.http_client_profile.tenant` property

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

<a id="canonical-0311132203322311-0233130001201300-1233022103211333-2133013320020233-0111320303133032-2212132112122010-0233220013131303-2002022300220012"></a>

<a id="canonical-2031311301013220-1121102111132031-1103103213201221-0102301013220101-2210210331301000-0103120130023133-0031023222100321-2301011302331112"></a>

#### `virtual_server.http.http_client_profile.uid` property

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

<a id="canonical-2322322120023310-2331230320202223-2321331132010123-1120303200320012-2132131011333223-0311000303222001-0111212033133333-2113032201233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http_server_profile

<a id="canonical-1213110300000223-2122321330223221-3123002011000213-0021132300332002-0032022023223023-3133330032100221-1300020310321322-0223322210021032"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201022131130031-1231022302023111-3001122302001221-2033111030230110-2100130012321023-0132013331203310-2220202133303232-2012203230231110"></a>

### Direct properties for `virtual_server.http.http_server_profile`

<a id="canonical-0123112123120200-0102021220033201-1200130211101330-1032020001301303-0032013331130301-2130212010113221-1112202111232321-1123232123133233"></a>

#### `virtual_server.http.http_server_profile.kind` property

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

<a id="canonical-1300111231210032-1020010010131011-3012310202020113-1312211021223123-3133001111312332-0201300000232011-0133320013331312-2302133101002003"></a>

<a id="canonical-3312210220220323-0000011302201031-3301313013001032-2020001210203233-1211222220231221-1002001202310302-1022122313201120-3332003110011312"></a>

#### `virtual_server.http.http_server_profile.name` property

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

<a id="canonical-3121232322132020-3332023201332123-0233203032212023-0220012213023322-2313011333231313-2313001112232003-3300222222031020-3100101322120233"></a>

<a id="canonical-2221122232203013-1331021011111011-0102020330203003-2212011310101111-0033123210132312-3101132201230113-1300300013312231-1012123220011030"></a>

#### `virtual_server.http.http_server_profile.namespace` property

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

<a id="canonical-2323312323003313-3033312331322211-3203313301020221-2132111232212003-1333322322230012-3332113030213222-0231223211023212-0121210300123022"></a>

<a id="canonical-2312101212200121-0233303132100131-3333032033112131-2203310233031000-0033330301331232-1003223032202133-0323013003022023-0121012333321003"></a>

#### `virtual_server.http.http_server_profile.tenant` property

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

<a id="canonical-0122133312001331-1022310203010302-0233012103332213-1311201022311230-0112112322122231-0231010330131111-0203100211213130-2232202123322202"></a>

<a id="canonical-1230231200230103-2123023222211002-0200203312023122-0300311010200003-3030333003223110-2203213221020220-0100302013123320-0132313120100133"></a>

#### `virtual_server.http.http_server_profile.uid` property

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

<a id="canonical-0212200100013300-2010330003202013-2113300023302120-3201222310221023-0313220131110103-1312321001211122-3021120111312003-3200020321033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.ocsp_profile

<a id="canonical-3100301002030221-1333133211330332-3023213213021130-1030100001100123-2231320331330011-2103212330310100-0332130113313302-2323130220000223"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000011023130033-0131003130032033-0021233231002212-1233013223232203-2133023322111020-0132313122302023-0313103002103212-2202021312101302"></a>

### Direct properties for `virtual_server.http.ocsp_profile`

<a id="canonical-0213213000023030-2031213303111300-2301011013312220-3233030330333331-2303320233112223-2302213020122111-2022131003003223-0212121233202012"></a>

#### `virtual_server.http.ocsp_profile.kind` property

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

<a id="canonical-0130120301011100-3200223212121012-2202001302003132-0303010233320232-0011223121021233-0301013012310000-2003313030001033-1213100102333200"></a>

<a id="canonical-0313010323120223-1223013231133001-2301212023202020-1232002002213201-1101021201232200-0003011113231032-0201012021003031-0230102112231222"></a>

#### `virtual_server.http.ocsp_profile.name` property

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

<a id="canonical-0300202233010331-2211211132210022-3103231202213033-2333120002112232-2102132133211300-3331102322031120-0030033330013213-3222201230021123"></a>

<a id="canonical-1332100130111223-0232120101132220-1210102221330130-2133032232020012-2310201201333130-1123030130212023-0322110123030222-2002313222132001"></a>

#### `virtual_server.http.ocsp_profile.namespace` property

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

<a id="canonical-0311301120303300-3322213301301012-2130303131032301-0032202323001220-1212123331330100-0101022012112321-3020003212023300-0300201210012211"></a>

<a id="canonical-3233133320303322-0102313220220022-3013111113103123-2311113113023132-3331123001102120-1223313303103001-0010020112212003-3003233302301123"></a>

#### `virtual_server.http.ocsp_profile.tenant` property

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

<a id="canonical-3001331120213330-1222231232223200-1110021010300110-0130323311033013-0013120233222121-3123011202013232-3111003112311002-2100103111312221"></a>

<a id="canonical-2300320023212112-2321023223131101-3211332010211331-3310312332203000-1210210022211032-3322311202022202-1122013321033202-2032110332121132"></a>

#### `virtual_server.http.ocsp_profile.uid` property

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

<a id="canonical-3232011302112010-0200312011113331-3101111203223133-3020113313300300-0023113130010221-0133232012023323-1002302300221201-1331300202001202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.server_ssl_profile

<a id="canonical-3010133303300011-3222112300303320-1201223020122330-2133200201010012-2123103122120202-1121303321332202-3320312331123202-2000333302030120"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123100110320230-2111010203122031-2203110222202310-1031231312213211-1300132322132303-3202300032012102-2321202301103321-1010031002103331"></a>

### Direct properties for `virtual_server.http.server_ssl_profile`

<a id="canonical-1012310300203232-2210311031133003-1322030232220323-2001000301031231-3120022313323033-2022301133213233-3330000101120030-0303300012102210"></a>

#### `virtual_server.http.server_ssl_profile.kind` property

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

<a id="canonical-0113130033330003-2001102213130102-1213331212222101-2132022032322212-2310003201003202-3321223130121131-1321110310331233-3103312332332022"></a>

<a id="canonical-2312333122121122-2231122003111010-2220130020031003-2331020011120121-1003122222123231-1111010033010300-1223131202131223-0010212232231031"></a>

#### `virtual_server.http.server_ssl_profile.name` property

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

<a id="canonical-3103311222322332-2313101313122333-3303311220230231-0012211131102130-2213100321122230-3332020232033121-1031011203311033-0122011212102310"></a>

<a id="canonical-3222302020330130-0222222230010203-0232011131322201-2232223013022200-3333003100100333-2022233312333113-2000300120123232-1213322321033131"></a>

#### `virtual_server.http.server_ssl_profile.namespace` property

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

<a id="canonical-0311331321301221-2032133113211303-2112213231301110-0201103110220200-3003031010213011-1133332202222133-2110211303220333-3312310130001133"></a>

<a id="canonical-1103212121131300-0110033331223002-3333201012203101-0110201213012202-2100000103301201-2102213131103300-2013031021101300-0030001110322322"></a>

#### `virtual_server.http.server_ssl_profile.tenant` property

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

<a id="canonical-3000002222332213-0130233303102212-3121023212100231-2021021123110012-1303203313201212-1130232013301210-2011121003131103-0312332303203130"></a>

<a id="canonical-3001333212033021-3323331012202022-3133003322333322-0000111312210132-3112130232201133-3102032302002033-3330311001122112-2310130133121111"></a>

#### `virtual_server.http.server_ssl_profile.uid` property

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

<a id="canonical-2003122102331011-0102223311211013-1113200013103313-3101102121131201-2101310003131021-3232021132033212-0313132302121023-2222132212203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.stream_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.stream_profile

<a id="canonical-3102013030110002-1100221012323030-3312013033112103-3012023300122022-0132211310102010-3301002032211022-2110002101003312-1033331011012212"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for stream profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
stream_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210212331321112-1330313212003231-0320310223131111-0212110302133321-0331013130013100-3213113303302121-3202220213031032-3222020130313203"></a>

### Direct properties for `virtual_server.http.stream_profile`

<a id="canonical-0033323323112113-3121331332122111-1311120203312000-3203113121021202-1322103133020021-3203230201010012-0001223001120103-2232223001321020"></a>

#### `virtual_server.http.stream_profile.kind` property

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

<a id="canonical-2333233112332102-0020003323010312-1121221332221311-0012213223102203-2311322033020333-2130013013203211-3313231221300103-0132200032003133"></a>

<a id="canonical-3311231002233221-2010311310313122-2222222322000320-1112023202330010-0213233023021023-3211310112231033-1121030121312221-1203033310222310"></a>

#### `virtual_server.http.stream_profile.name` property

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

<a id="canonical-2020130232330213-2101101323212120-0010000313203013-2002103202003130-3233023012013332-2122102130203323-2321210222122320-3011210113103112"></a>

<a id="canonical-0032312101131320-0120111002002230-2023003100221031-1300330211113111-0133033132020322-3003202221011211-2202312231331312-1003013332202232"></a>

#### `virtual_server.http.stream_profile.namespace` property

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

<a id="canonical-3133301103203232-3203121101200023-3133323320130221-1121000120110131-0131212201012111-0220131222022302-2233121332221033-2300300131010313"></a>

<a id="canonical-1203202022323131-0220131021212103-2010322131321020-1001333031031030-0321231320332300-0023021313002012-1332230130121220-3102031032203302"></a>

#### `virtual_server.http.stream_profile.tenant` property

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

<a id="canonical-0102131022332311-1301310130321013-0101322320030322-0130333101203301-2331030033010331-2310202032210231-2301021110032231-0131310220211212"></a>

<a id="canonical-1010122120012323-2310303122021031-1311330300311202-1202330230331211-0211322100333222-0033112123331100-1321322102011032-0102220203031331"></a>

#### `virtual_server.http.stream_profile.uid` property

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

<a id="canonical-2223311101321210-2223313131302002-2310010033333222-2122013110101211-1021221123210103-1210221022002331-2101300030310020-2331210030200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.tcp_client_profile

<a id="canonical-2222103002213222-1321203331302022-1212333301101011-1232112221113131-2021123322210003-3330013232203221-1323211110021210-3203312102321133"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323130130302222-0031000121001101-3311120323113312-0212022212210001-0001011311320010-2230133221030203-1202031000233211-0100213331100222"></a>

### Direct properties for `virtual_server.http.tcp_client_profile`

<a id="canonical-3100210000103122-1130131103112020-3303032232032300-1233220333223213-0023323313231002-1203010022010012-1001123231120003-3331313322111012"></a>

#### `virtual_server.http.tcp_client_profile.kind` property

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

<a id="canonical-2101110333101233-1300123303133222-1202110230223302-1321102021300203-3002310000031003-2322223030230020-0202121331130120-1023311102201100"></a>

<a id="canonical-0033011301121000-1212320330302321-0231333232032211-1233300131032010-1331100312101321-1223110010022303-2121013330222302-0110101131322102"></a>

#### `virtual_server.http.tcp_client_profile.name` property

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

<a id="canonical-0013311201323312-2211312312312022-1022302111030002-1331110120012212-0023300031202200-2032313321123311-1030033232201332-0020120103230000"></a>

<a id="canonical-1333312231201102-2123111133300013-0122012002231213-2022212313223313-3020222122220332-0333113312110301-2033213100132110-3003320100321013"></a>

#### `virtual_server.http.tcp_client_profile.namespace` property

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

<a id="canonical-0321200223302101-0100131032311111-0222202033010030-3033233333010002-1332322212302022-2311011010122100-3020331310201322-2132231001330320"></a>

<a id="canonical-1030203012220130-0312130202112120-3010321113333013-1113031121233201-3333113222312313-0000232023201133-2133230012001321-2122131030231233"></a>

#### `virtual_server.http.tcp_client_profile.tenant` property

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

<a id="canonical-2003010303310013-0031210011230133-0321133231132022-3110212212101030-2200021033221232-3213310211310021-0301001302231330-3112322321131231"></a>

<a id="canonical-3232311110222012-1311211333130213-0123223033313001-3103002021203210-0302332220233011-3333020010101113-1203200321220223-0020130323000220"></a>

#### `virtual_server.http.tcp_client_profile.uid` property

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

<a id="canonical-3313330121201111-0233011211100030-0312301221013133-0221322322010322-0221231321001010-2010112300003100-1220230110101232-1323330130100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.tcp_server_profile

<a id="canonical-0032322120000231-2113212203313122-3322322010120231-3301111113111021-2031230313322100-2220221313011021-3303121310113132-0022300100310032"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203101123131222-2301002321233231-0201233131201011-0223022003200002-1221221311012300-3221100000013211-2232133130200021-0303133131200010"></a>

### Direct properties for `virtual_server.http.tcp_server_profile`

<a id="canonical-3330231201203003-1031222202310302-0120220020202323-2322323012032011-3311300120301120-0222301302122033-2301321023001223-0011000202001000"></a>

#### `virtual_server.http.tcp_server_profile.kind` property

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

<a id="canonical-3101211303223032-1231020320220223-3022020310322030-3133221330020101-1201001203121132-2101202221332231-0103020131322130-3012002001120232"></a>

<a id="canonical-3002111100133300-3323220321200231-1303021300103000-0230322232010203-0203002002233200-1132302223120111-0203333103002001-1122103122322300"></a>

#### `virtual_server.http.tcp_server_profile.name` property

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

<a id="canonical-3021032011003322-1332330220020233-1023122033131112-1312120123200023-0213320013313013-3132003313111030-2213012133020212-0220220010130321"></a>

<a id="canonical-3300222100231220-0121311321102311-2010213303122321-2113303213123130-0100210210230000-0001012313122023-2333212123120120-3223232103110211"></a>

#### `virtual_server.http.tcp_server_profile.namespace` property

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

<a id="canonical-0223011121300023-2212232213313220-2202230031232023-1332223320023032-2211330212023002-1322301003313011-1013332330122023-1313100031320223"></a>

<a id="canonical-0012221302203023-1113110331321323-1000020001221301-3022333332200013-0312113003010323-0123322101013323-3131210233130130-3120331021302122"></a>

#### `virtual_server.http.tcp_server_profile.tenant` property

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

<a id="canonical-3012110322231301-2021220300300201-0311100000133220-3032032033330021-2132321233031013-3232202101330013-2012220323221322-2023123201012322"></a>

<a id="canonical-0321000312101000-1231111221300120-1230221321200230-3322200102200023-2303323232300330-0011130001120200-3012100110000103-2222033110303002"></a>

#### `virtual_server.http.tcp_server_profile.uid` property

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

<a id="canonical-2103003320213123-2333013120322010-2312232213313300-1031221031100011-3332121220121033-0110233202012310-0011130201222110-1123332212132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.websocket_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.websocket_client_profile

<a id="canonical-2223110220120301-1213223332122020-1320112133333333-0201211010202031-1022132031331300-0332032202323012-1303202122321332-2231221120220032"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Client. Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011211310013002-0020200033212310-1001212331202101-2013111311033003-0212103300131331-0130321012221222-3223132323212103-3103223001201302"></a>

### Direct properties for `virtual_server.http.websocket_client_profile`

<a id="canonical-3030032010131311-1001103223310323-0011013221123013-0013102111032322-2023230033303313-3033131033103131-0013000113201101-1313021313110222"></a>

#### `virtual_server.http.websocket_client_profile.kind` property

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

<a id="canonical-1122302322322001-2102320321021121-3110330220332222-0011032330322310-0311312112011030-0003212023003011-3103113013001330-3120132110033101"></a>

<a id="canonical-0023333003330133-0300233000322203-0323311130210103-0132132331331133-1223030133030320-2132022301000103-2032131331010301-3101233033121133"></a>

#### `virtual_server.http.websocket_client_profile.name` property

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

<a id="canonical-3101102123202032-2113102303113130-3020312322030012-3101312332022031-1100313233303333-1131310021301322-2233032011011230-0001033110103012"></a>

<a id="canonical-2202221113212222-1120222002020102-0222130330113022-3113120233310300-0121023022021122-2302302233211223-0202330200021300-3222022220113021"></a>

#### `virtual_server.http.websocket_client_profile.namespace` property

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

<a id="canonical-3200203130131213-2211031120033003-1103221001102300-0202201030013133-0023201102112131-1110003113111032-3301000331321032-3010303322220222"></a>

<a id="canonical-1212032132010230-3032013233330210-3000121320222332-1122001312321301-2200312330012221-1131230010001003-3030201321111031-2000120130332332"></a>

#### `virtual_server.http.websocket_client_profile.tenant` property

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

<a id="canonical-1331111301122210-1023033323332121-3112312330002232-1223210212221120-2332030012232213-1102232323103320-2032322201002202-3221031013320323"></a>

<a id="canonical-0203022222201130-3011131101023012-3302023200001000-0310301231320202-1002330001002233-1232122212302032-2110331110001003-3303032020132201"></a>

#### `virtual_server.http.websocket_client_profile.uid` property

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

<a id="canonical-2133310101022000-1330312111332132-3210102331201100-1113232033330012-0200001232322231-1010212032131322-0301022332301131-3030020121130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.websocket_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.websocket_server_profile

<a id="canonical-1202232302031303-0013323212010003-2020001202002320-1103221331022212-0113213022111122-2102330111012133-0022131122303201-0103012213310211"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Server. Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120331323220112-3000020020233111-0221132003311331-0323123102201100-1111000322023313-1313310230100210-2113032230333210-0100002103130303"></a>

### Direct properties for `virtual_server.http.websocket_server_profile`

<a id="canonical-2131021222322302-1310112022132011-1112213013312312-0223320031030002-2101020321001130-1121113321202122-1212200122031320-2311103013020331"></a>

#### `virtual_server.http.websocket_server_profile.kind` property

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

<a id="canonical-1231113113020320-1002213012123130-3102331221323121-2101130331232211-2001000333102302-0032102100002210-1210213211303030-3310211303022230"></a>

<a id="canonical-0312032123121303-1110333011320002-0130020303231032-2121000122023202-0132032322201311-0130132233323221-2012010132101011-0133022122310033"></a>

#### `virtual_server.http.websocket_server_profile.name` property

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

<a id="canonical-3333101202000331-0331320333223120-0310321333023120-1330033023011123-2300203111011200-2121130100230100-2133221300332113-2231023220210223"></a>

<a id="canonical-0021210202321311-0301011031002331-0012022000220310-1320222233322321-0122213111103203-0121000301203021-3300023030121032-3010023321211003"></a>

#### `virtual_server.http.websocket_server_profile.namespace` property

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

<a id="canonical-1111213312131202-0002000100232122-2301033120001112-1233023323203101-0230010112233302-0321021031112231-1123033221223311-0312211321133331"></a>

<a id="canonical-1112312111013103-0010001011011301-0330213032312322-1210000002232110-0013230213102010-0112021013112213-1231312110101330-3200201022023030"></a>

#### `virtual_server.http.websocket_server_profile.tenant` property

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

<a id="canonical-0031201230102221-3321111012120001-3223303333010333-2303231200023111-3333102132002302-0000301232013021-0231233121031122-3021213231131011"></a>

<a id="canonical-2322203130000032-2332003301101331-1332132331232021-0102322111222011-1002221300022131-1331201003030313-1033202133202231-0102113023302232"></a>

#### `virtual_server.http.websocket_server_profile.uid` property

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
