---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-0221000322321030-1113211202122102-0133131202033231-1331110121001123-0111220333300112-3122311132013122-0321033031322012-3310302200131001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface_list` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.interface_list

<a id="canonical-3010023100000333-0132232020110110-1313203002321022-1002210013020010-2301332303231233-2023100121232222-2332110202121120-3101221133121211"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2313020323222121-3002312110203312-3322021000303032-1010133303021121-1101333210232213-1321201333211333-1101112133133022-3321302210001031"></a>

### Direct properties for `peers.external.interface_list`

- [interfaces](resources--bgp--reference--group-002.md#canonical-0203222020113313-3333130022100002-3010001001331332-2333231321233111-2033211032302211-2031103132320120-2211013011202231-2312102111100000): complete subsection reference.

<a id="canonical-0203222020113313-3333130022100002-3010001001331332-2333231321233111-2033211032302211-2031103132320120-2211013011202231-2312102111100000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.interface_list](resources--bgp--reference--group-002.md#canonical-0221000322321030-1113211202122102-0133131202033231-1331110121001123-0111220333300112-3122311132013122-0321033031322012-3310302200131001)
- peers.external.interface_list.interfaces

<a id="canonical-1332320012301200-3311220113333022-3230210010301011-0113313313212133-1313012332203001-2022031331100230-1201112110330332-2032022133320331"></a>

Type: `"object"`. list nested block, Optional.

Interface List. List of network interfaces.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000111010131222-3302333000110112-3031100101110220-2033033100220031-0222310200232002-0111303010133021-3202212322311001-1201300103220302"></a>

### Direct properties for `peers.external.interface_list.interfaces`

<a id="canonical-2001231001102201-3302101131131132-1202112121023233-3332302003000213-1030320013031031-1100212033110111-3201322201301123-0313110121312331"></a>

#### `peers.external.interface_list.interfaces.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0131023001131032-1122133011320202-1020330113022131-0032223011311132-0001311302300220-3330001200232230-2332233300232103-3102000102323300"></a>

<a id="canonical-3111321222323313-1233031101021030-1021131212220113-2333013121002133-2012011222100020-2200311033011013-1221120101220021-2102310103130332"></a>

#### `peers.external.interface_list.interfaces.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2311221131010231-3320301333331102-3232000020100122-3213101221132010-1000200311031023-2102002321310200-0012330203021221-1010132323311112"></a>

<a id="canonical-3013032301222211-0230111100303322-3210212132330202-0322113031032111-0320320320103203-2201122210020333-2312320121220213-1012023323330310"></a>

#### `peers.external.interface_list.interfaces.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0330223011231313-0131302323100133-2232102110312200-0331033133303032-1301100120202101-2320101223111322-3100000322300223-0032212000310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.no_authentication` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.no_authentication

<a id="canonical-3331131031013000-0123322221321102-3001103211302301-1312030020112001-2130223211000220-2011002333003121-2100123022301231-1132312010002121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

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
no_authentication = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102332321320322-3302110300310021-2011013232000101-0212323123110010-0032202032311310-3223212030010113-2101332010233031-3022021202302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.metadata` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.metadata

<a id="canonical-3123021121030231-1321121201110330-2231213223012321-2231100202001102-1230331233033310-0120131322131110-2120302011121132-1313110021113300"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032011330032213-2032211123011302-0130330200322131-3212102031321210-0101102303102002-0012100232113121-2002233013032003-3120000010000100"></a>

### Direct properties for `peers.metadata`

<a id="canonical-2321232310222311-2033202210113321-3211230220001103-1122211013312211-3220200100023320-2013111001002320-3301312011201130-2302103300101330"></a>

#### `peers.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2201302213321233-1331331200311312-2212022110032301-2233330133133312-0121122030032230-0022112210120212-0100002020030330-2020020222033323"></a>

<a id="canonical-2012312031001203-0011233011003000-2033130002203333-2321001311021322-1332001021123322-0303000313211103-1332113020133201-1322001303122112"></a>

#### `peers.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3001222122101221-1132332210123021-3233102220031330-0130111121100131-0022302320313231-1211122232300220-3033321021113002-1312113030132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.passive_mode_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.passive_mode_disabled

<a id="canonical-2212020233023212-2222232202312122-2020000113303321-2021001111300031-3313301122332212-1020122230002331-2032310231112310-0212231103210122"></a>

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
passive_mode_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232120302301113-3232111320111310-2033121311313201-3030333321022232-1011013210322100-2321202101133221-3333110022231200-3330323000303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.passive_mode_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.passive_mode_enabled

<a id="canonical-1211231311112210-2302200201102202-3222212121032320-0033302331200303-2310002333132111-2333330210220302-1220233203300200-0331131300012001"></a>

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
passive_mode_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.routing_policies

<a id="canonical-1222011310120302-1020320101232012-1012212323223221-3030013303323030-1100221312103133-3013110032213120-1222103032012133-2301000320111113"></a>

Type: `"object"`. single nested block, Optional.

List of rules which can be applied on all or particular nodes.

Receipt-pinned upstream constraints:

```json
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
routing_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322202033130203-3000011033223001-1231223013132201-1301102013202210-2022011313202113-2300223113100113-1002022100103020-2003200222311312"></a>

### Direct properties for `peers.routing_policies`

- [route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312): complete subsection reference.

<a id="canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- peers.routing_policies.route_policy

<a id="canonical-3022131123031203-1002112023220131-3200102201102031-1232113311200030-3211302303313123-2210130111233323-3303102011112021-0031313200022303"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Additional upstream details:

Route policy to be applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("object_refs"),
  validators.ConflictingListObjectAttributes("all_nodes",
    "node_name"),
  validators.ConflictingListObjectAttributes("inbound",
    "outbound")}
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

Terraform syntax:

```terraform
route_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003330123213301-2000330211033130-1033312032023231-1223313301232301-1222020023221122-0133032200323311-2230023323321220-1302103020102331"></a>

### Direct properties for `peers.routing_policies.route_policy`

- [all_nodes](resources--bgp--reference--group-002.md#canonical-3200212313320220-0011130101333033-2330233320300232-0123330021123323-3210221001303022-0013331303101320-3010223121032312-3303033101001030): complete subsection reference.

- [inbound](resources--bgp--reference--group-002.md#canonical-0331112133300211-0220113302003321-0003030201203110-3331322233022123-2313003003223123-3331312211120012-2213002203022013-2002322121000333): complete subsection reference.

- [node_name](resources--bgp--reference--group-002.md#canonical-0223023301201233-0013302203332313-2100312202110011-3322232232303000-1022321200012330-1323321022220112-0122032220100333-0313033203333220): complete subsection reference.

- [object_refs](resources--bgp--reference--group-002.md#canonical-3030211023130222-3112031333003232-2333303021322312-2100100200012032-3002310002022331-1200003323031121-2312021001220331-1100300220233123): complete subsection reference.

- [outbound](resources--bgp--reference--group-002.md#canonical-3221333323021330-0323000002011212-1023032220130110-0122031210303123-3212000211131011-0330300333213322-0130121001222101-1102121322001321): complete subsection reference.

<a id="canonical-3200212313320220-0011130101333033-2330233320300232-0123330021123323-3210221001303022-0013331303101320-3010223121032312-3303033101001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.all_nodes` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312)
- peers.routing_policies.route_policy.all_nodes

<a id="canonical-2102120230022023-0322332220121120-0303223021202313-3320210222313113-1233032133012203-1123133131002201-1032131113301020-1203010003212100"></a>

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
all_nodes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331112133300211-0220113302003321-0003030201203110-3331322233022123-2313003003223123-3331312211120012-2213002203022013-2002322121000333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.inbound` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312)
- peers.routing_policies.route_policy.inbound

<a id="canonical-3232300212311103-0231023320112211-0201112100310020-2333233111102223-3012310210213101-1310221031332321-3100320211213013-1321223320020023"></a>

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
inbound = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223023301201233-0013302203332313-2100312202110011-3322232232303000-1022321200012330-1323321022220112-0122032220100333-0313033203333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.node_name` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312)
- peers.routing_policies.route_policy.node_name

<a id="canonical-2100303231002121-2031220231011032-3202001321030203-1232023203231321-3332100302333032-3101231023133131-1030132221110023-0322132230213121"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

Receipt-pinned upstream constraints:

```json
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
node_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233112120222010-0213233132031301-2000312313130330-1133231300010001-3002231123020223-2101321212233122-0022023311100030-0102303013112121"></a>

### Direct properties for `peers.routing_policies.route_policy.node_name`

<a id="canonical-2012331002333101-1133322030310302-3110002120212201-1233332023121132-2001200233030321-3221301011212101-0211030013313000-0312130222232111"></a>

#### `peers.routing_policies.route_policy.node_name.node` property

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3030211023130222-3112031333003232-2333303021322312-2100100200012032-3002310002022331-1200003323031121-2312021001220331-1100300220233123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.object_refs` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312)
- peers.routing_policies.route_policy.object_refs

<a id="canonical-0231022312211000-1133122020231123-3012212311023001-3020023201200312-1301103120120133-2202222323223231-3311330120312211-1211101030131030"></a>

Type: `"object"`. list nested block, Optional.

BGP routing policy. Select route policy to apply.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
object_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203320133031332-3301231331320203-1303210110113322-3010101101033212-3031312013011320-3331110022223303-1232113011122033-2120113020202133"></a>

### Direct properties for `peers.routing_policies.route_policy.object_refs`

<a id="canonical-3303100000322320-0303320303313123-0213023320110113-1201332110022123-2023302122021133-1101301223022233-2112233031020201-3323333023222233"></a>

#### `peers.routing_policies.route_policy.object_refs.kind` property

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

<a id="canonical-3213230102102032-2202232011323010-0031220011332033-0220030231033212-3233322133022122-0002301302303210-1220023330323330-2321230232021301"></a>

<a id="canonical-0112021333133202-1101123202033320-2110032003311300-3033131110211212-0110220212212100-0312032100301103-3333331130300212-1002321013020013"></a>

#### `peers.routing_policies.route_policy.object_refs.name` property

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

<a id="canonical-0202323012200310-1101123320331130-0011133130002310-3220120300202002-0313101032033320-1312122133113210-0021030030312023-0013100320223020"></a>

<a id="canonical-3201333111133321-2202120230211220-1130230200000010-1231100210103110-3221230030233231-0021313313130212-2013303201100013-1333010332110032"></a>

#### `peers.routing_policies.route_policy.object_refs.namespace` property

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

<a id="canonical-2200300120323213-0223001322322320-1011322031112311-2002021312000021-0023230311232222-2302032132111303-3033312120022331-2011222331013200"></a>

<a id="canonical-3303012101132321-2031123212032121-1313331201321201-2000033201232103-3321210022032001-3111221112011033-1222012100112302-2020311022011130"></a>

#### `peers.routing_policies.route_policy.object_refs.tenant` property

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

<a id="canonical-3112023030132101-0120002322112123-0133222113213330-1231120313200122-3200202020300312-1331112302230032-2313130001032101-2320011311301022"></a>

<a id="canonical-3200311230002023-0003023332212322-0320302330121122-2102313113332002-0122132321031232-3302200220111010-0130032123111331-1213012032023013"></a>

#### `peers.routing_policies.route_policy.object_refs.uid` property

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

<a id="canonical-3221333323021330-0323000002011212-1023032220130110-0122031210303123-3212000211131011-0330300333213322-0130121001222101-1102121322001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.outbound` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-0122030323303113-0012010203020021-3311132111121031-2023000321132100-3023211100310310-0121330020331232-2320000302322232-2200030131311312)
- peers.routing_policies.route_policy.outbound

<a id="canonical-1132311230203222-1203220010222010-2021121231222300-0202223200102012-3202330200123302-0233023022112021-2102110211111202-0001212201102310"></a>

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
outbound = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300212323303233-3123033100000031-0113223222020021-3122201130331211-0030003011113133-0300212300011301-1323133010230310-2232202103011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- timeouts

<a id="canonical-3021132332201311-2300111330103202-2031210111010222-2323023300232102-1121013002001221-2302111301321311-1300210212212221-1133120311021213"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302331203213203-2300111101123321-0321311332120331-1101331313132331-2303332220121123-2132201123001330-3121200230023323-2330221013011312"></a>

### Direct properties for `timeouts`

<a id="canonical-0001123213110033-1101132023002101-0301032133310021-3110122311333033-2002121333210203-2031310220331120-1002212000023023-0113020122302013"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1303300133212212-3113003210211101-0213130023022033-2130101330100030-1131132212033020-3330223222213231-3031021332213321-0123130130333021"></a>

<a id="canonical-2310012010003202-1231312223300002-3302032320113231-2123010032332100-1231322122221032-0332202010303033-1101103031202223-1012112120102000"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0213202220010231-3132220332213101-0032133111301032-3023113333022022-1311132202320330-2300023113101013-3322233133200220-0210211330323113"></a>

<a id="canonical-3110220013111320-1213232120121312-3122303330130000-0102001110121023-0100320111221023-0231211122033233-1221330221133323-1102103302101020"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1211222311012113-0322013113212200-3032010131111123-1130310123120112-2332301123030032-2003033213012030-1031032222213220-3031300023320130"></a>

<a id="canonical-3110220220202322-0112230223131313-0010320323021233-1022131323203002-0010113112313030-0020112131301013-2031201003131210-2121213302030302"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- where

<a id="canonical-3220103110002133-1333320020233033-3313320203230220-1020021110111113-1203113112011300-1131212123120133-0030303222123001-0021021002032023"></a>

Type: `"object"`. single nested block, Optional.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323312130103010-0222332023301120-1213201203023303-1120331012320032-0121210313011312-3210212110311223-3021103112300332-1232323332220122"></a>

### Direct properties for `where`

- [site](resources--bgp--reference--group-002.md#canonical-3003103202123011-2133102220310110-2313130323201112-3020222031201202-3230211211322333-1320210100002013-3011300020032303-3300020332210021): complete subsection reference.

- [virtual_site](resources--bgp--reference--group-002.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200): complete subsection reference.

<a id="canonical-3003103202123011-2133102220310110-2313130323201112-3020222031201202-3230211211322333-1320210100002013-3011300020032303-3300020332210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- where.site

<a id="canonical-3230002111332233-2023012332221333-0300112011323011-3021321320110121-3301033130223130-2320000113002011-0312111033010033-0230123331200302"></a>

Type: `"object"`. single nested block, Optional.

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312213330330200-0003220023301131-0321222112202003-2232132220133300-0213031103030323-2131210022303300-1213333300303010-2011100201000112"></a>

### Direct properties for `where.site`

- [disable_internet_vip](resources--bgp--reference--group-002.md#canonical-2010120022210101-0231321322203131-0323101320233313-3031000003210202-2032021131102102-3220021302000223-3020301013303332-3300300032103102): complete subsection reference.

- [enable_internet_vip](resources--bgp--reference--group-002.md#canonical-3113303301231310-2031210322012313-3211201222233211-2323132121320120-3232332121000323-3012033300100303-1131211310322012-1300002332121003): complete subsection reference.

<a id="canonical-1131203123322230-2323011230100021-1101000301011201-3202201332003031-1002312203020221-0330322033223131-3013133331210020-2322112211102113"></a>

<a id="canonical-3222131330200230-2023231320312021-2003330220103132-1101133133003022-1203212301200033-0003122121002311-0212231331001022-0333300101100132"></a>

#### `where.site.network_type` property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--bgp--reference--group-002.md#canonical-0122130011100021-0123001230112303-1211321023202201-0013002033110302-1110231213130112-1310100303031313-2223021320310003-1312201001032221): complete subsection reference.

<a id="canonical-2010120022210101-0231321322203131-0323101320233313-3031000003210202-2032021131102102-3220021302000223-3020301013303332-3300300032103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.site](resources--bgp--reference--group-002.md#canonical-3003103202123011-2133102220310110-2313130323201112-3020222031201202-3230211211322333-1320210100002013-3011300020032303-3300020332210021)
- where.site.disable_internet_vip

<a id="canonical-1223003311301221-3303101102323112-0130122122213132-3213033313020013-0332110231202011-0311010303133210-1233303323101111-2213302212020112"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113303301231310-2031210322012313-3211201222233211-2323132121320120-3232332121000323-3012033300100303-1131211310322012-1300002332121003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.site](resources--bgp--reference--group-002.md#canonical-3003103202123011-2133102220310110-2313130323201112-3020222031201202-3230211211322333-1320210100002013-3011300020032303-3300020332210021)
- where.site.enable_internet_vip

<a id="canonical-2130001020322233-1313231130113100-2213110233332300-0132130310331332-2030132012303100-3111212022301103-1210023031320011-1221101302100320"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122130011100021-0123001230112303-1211321023202201-0013002033110302-1110231213130112-1310100303031313-2223021320310003-1312201001032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.site](resources--bgp--reference--group-002.md#canonical-3003103202123011-2133102220310110-2313130323201112-3020222031201202-3230211211322333-1320210100002013-3011300020032303-3300020332210021)
- where.site.ref

<a id="canonical-1022131212131330-1012011213112133-1133032212312132-1321322130010022-2133012301012301-1003202133302321-1230212333013323-1213113201102213"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123211322300012-2210032032313301-0011102213121113-1311022123331331-2311301120123233-2312033222210013-0223023013001333-1021303120220311"></a>

### Direct properties for `where.site.ref`

<a id="canonical-1221313221202101-3133032112000212-0303210001112330-2023023231010330-1120002203333000-1322013211120023-2320223311030212-1230110231102333"></a>

#### `where.site.ref.kind` property

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

<a id="canonical-3000033133131221-1033311203321011-2231120033010230-1002123011200331-2002130211031233-3200033123022322-2102032200020030-3312023010130101"></a>

<a id="canonical-1212230101013232-0300321333013020-2123101232320033-1211203212202203-0303320000330220-1021020312222002-3032012123213310-2130013320011321"></a>

#### `where.site.ref.name` property

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

<a id="canonical-2131232113200123-1302022101201111-2210303202010221-2131101120132310-3220223320020301-2302301112132333-1111131232201011-2032312203111202"></a>

<a id="canonical-1331213213012120-1030330012112333-0201211230320321-3203133201033022-0100020031232233-1312000203233011-1030322311123232-1032322320133111"></a>

#### `where.site.ref.namespace` property

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

<a id="canonical-0121001320211210-3211223111302030-2210313003210231-3102231201020320-3013323321333302-0233013000300232-1311110103321330-1221301131112300"></a>

<a id="canonical-0020122332221321-3332321221101222-2232220313212010-1231211133023102-2321311333000310-2021010232120033-2010213021022223-2331203133033222"></a>

#### `where.site.ref.tenant` property

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

<a id="canonical-3313011120031311-3202203133010111-2231011312210220-0021201021230130-3323302331233231-2203003223322033-0202233310321100-0310130100132322"></a>

<a id="canonical-2132231030222212-2233130200201203-0112122322332222-3322332232323020-3303233230031312-2212001003012231-0012033213202233-2201111201013120"></a>

#### `where.site.ref.uid` property

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

<a id="canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- where.virtual_site

<a id="canonical-1133111233031221-3102311311110012-2201202300020232-0322320321311213-3110230103020321-0230221330332120-2111011013030311-3121133232001101"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033312123103100-1003313120012012-2002000011001221-3210202233200311-3132221022220301-2120330323220202-3130113300310303-0131113311303331"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](resources--bgp--reference--group-002.md#canonical-2001312332332021-1211103132101230-1211213112201110-0333023221322020-3331321010220030-2211022013211130-0220112223203322-0320003031312021): complete subsection reference.

- [enable_internet_vip](resources--bgp--reference--group-002.md#canonical-3112200331130322-3302030033020110-3102222103210303-2023131011102130-3110310002210231-1300212211231023-0330331201321103-2110110102133123): complete subsection reference.

<a id="canonical-0112202232222320-0110213112121033-3210233131322023-1110011200012203-3011200100021022-0010021013122001-0211110200113232-1313312213312210"></a>

<a id="canonical-0221200000032221-0213230200011110-0013001212003313-3302110100302303-3223212301100303-3113201330031222-0000033111322011-1202011110100112"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--bgp--reference--group-002.md#canonical-2110210021323002-1032333330101203-0210010203103202-0311133000022303-0300003012201032-0202230323011002-1322313223002123-2101112131103311): complete subsection reference.

<a id="canonical-2001312332332021-1211103132101230-1211213112201110-0333023221322020-3331321010220030-2211022013211130-0220112223203322-0320003031312021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-002.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.disable_internet_vip

<a id="canonical-0300131333311032-0222313321100222-3221230201010202-1003022200310020-0022010312020023-0331303101133331-3123213210333210-3201130230330330"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112200331130322-3302030033020110-3102222103210303-2023131011102130-3110310002210231-1300212211231023-0330331201321103-2110110102133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-002.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.enable_internet_vip

<a id="canonical-3003323211221211-0232122113322013-2220301222332123-1121311313220012-1033303201311333-1330130101023011-3122132312022100-2220231213322111"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110210021323002-1032333330101203-0210010203103202-0311133000022303-0300003012201032-0202230323011002-1322313223002123-2101112131103311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231)
- [where.virtual_site](resources--bgp--reference--group-002.md#canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200)
- where.virtual_site.ref

<a id="canonical-0102011002200311-0011100111101313-2313132101001233-0331122211313201-2132120201311322-3102033101132231-2132102022023013-1223102030322302"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331232101223313-2011201220011201-2012030000300101-0010031233002331-1232303212300313-1032111012330111-0022223301301021-3102113003011313"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-0032033333233030-2131302231320222-2003130333112330-1203210132201301-2011133212311313-1023111001111032-1020012120220001-2022103030000230"></a>

#### `where.virtual_site.ref.kind` property

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

<a id="canonical-2110132321123303-1002131313200122-0200200333011121-1330001022021312-2110221012111332-3001101223012032-3210030012111203-0323031033120030"></a>

<a id="canonical-0220213313130130-0032011323200000-3123232012220033-1303322221133033-2130222122202001-3031202111030201-3002112020311233-0231230021201012"></a>

#### `where.virtual_site.ref.name` property

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

<a id="canonical-3032301222230330-1203133332231311-2332313130213303-3210102002210120-1313013221312100-0021001203323312-3322221230021222-2111102101313010"></a>

<a id="canonical-2210331111121222-0111122210023232-0122110102202001-2310210200130322-1303231030322012-3303331310113110-3312322231131332-2201121033201212"></a>

#### `where.virtual_site.ref.namespace` property

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

<a id="canonical-0133120002032030-2133313313320012-3312233100321231-0102201113022123-2222223322221030-2203100021113101-2300222123323111-2031003200200200"></a>

<a id="canonical-2031112323301321-1021222223133322-3023233232211312-0301300003212112-3030313320121301-0011313202232003-0101300011123313-2203223220311023"></a>

#### `where.virtual_site.ref.tenant` property

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

<a id="canonical-3101100121323210-3021302301200202-1122032321131321-3222132031031123-1131102123030131-2033013230032131-1021110103302222-3103322230113331"></a>

<a id="canonical-0112130002112122-3313221212213001-2132310303010312-2000231023232011-2010231103011232-1230033230111003-0300101222232331-0223232012132031"></a>

#### `where.virtual_site.ref.uid` property

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
