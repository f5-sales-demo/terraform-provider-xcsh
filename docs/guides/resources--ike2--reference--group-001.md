---
page_title: "xcsh_ike2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 reference."
---

# xcsh_ike2 reference

<a id="canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- Property reference

<a id="canonical-1212221012333223-2020203031333301-1100222100120101-3213331213130321-2300123123313202-3303220331330123-1221032023111301-3312201013233010"></a>

### Direct properties for `xcsh_ike2`

<a id="canonical-3000100111320012-3121232131201023-3212012120101310-2102120013233200-1130103220200202-0011022022103203-1103031212331031-3110032223101122"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2122332320220322-1230013111232123-1131232013100113-1133202123212203-1102011120311221-3200212210232103-1200110310020331-3322232301122032"></a>

<a id="canonical-1022310010001220-0210002331320213-3020122203203111-0130210221030303-1232330301232210-1303102033202031-0322222112200303-0213012312332213"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [dh_group_set](resources--ike2--reference--group-001.md#canonical-1132030321233131-0112333220310223-0120130001103312-1033130032311312-1222102313212001-3232201300232232-0030223131101212-3011200112322331): complete subsection reference.

<a id="canonical-2111123112312303-2331302122310122-3001111113200332-0010233202010230-2030301112022213-1130322012201221-0300110231301033-2321001101322031"></a>

<a id="canonical-2221120222202113-1021320022112331-0230123333013202-2033022001313101-0123222310031023-0312121222231220-2000031330212100-1120010332232012"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_pfs](resources--ike2--reference--group-001.md#canonical-3123101102221122-2012103322303011-1203112322120320-3320111010111123-2232303211033321-3302202122311112-0322000001010023-2203320122033312): complete subsection reference.

<a id="canonical-0222220132010301-3310222023103001-2131121012332233-0313330013233331-0113321200201210-2130322312100300-2312303332113313-3121030101210010"></a>

<a id="canonical-3323120212011303-3002022011233001-3123330323323030-0103323003030212-0032310232011000-0211130113222231-0000023321301120-0131010121300031"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-1233220103012212-3122103302311332-2213202110011333-0222131130103113-0311231022333011-3331312030001233-3022310301221023-0130120002011032): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-3100020103332033-2200320203233111-0013310012201310-1000300323210330-3312223323313112-2110101022322212-2102301010112130-3110302112233330): complete subsection reference.

<a id="canonical-0211320022032131-3302030113022200-3022021311111200-0221332101323003-2010130320232032-3102321211221212-3020123301213131-3021102320022120"></a>

<a id="canonical-0330021111301313-1230222213003223-2223231112330020-0213232021310211-3233330332222003-2320201321020321-2323313130003103-3301212002210023"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0023220222132021-0201023002100302-0211103012202132-3300020232013313-3311200111310310-2020112231312310-2133001302232330-3203122232130010"></a>

<a id="canonical-2313211210320023-1322313312322020-2322302302032232-2211331312122333-2231323203101330-0100303200221013-3022102320321300-1103123110003323"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Ike2. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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

<a id="canonical-0202200101111000-3113123333313221-3023033333202230-1103131130131103-1333322003300001-2331232010100302-1110231200330110-1222020132230211"></a>

<a id="canonical-2121010202200032-2221033112000203-0010013300003302-0131033001102322-3033211201313121-3011311320021321-1010233000313202-0121121103000003"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Ike2 is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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
  }
}
```

- [timeouts](resources--ike2--reference--group-001.md#canonical-0213233033132133-0200210123301111-3220032211230311-1021220110033120-1113001221313131-1311002133000201-3311231331022313-1012202102322311): complete subsection reference.

- [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-0203033021002212-3132000022022110-2300020230323322-3233000320123033-3332011010230320-2323323330132300-2212011002210320-0110220022320303): complete subsection reference.

<a id="canonical-3103033030110002-3001221221311000-0033110300230033-1321203313221213-1333211101332322-3303032110101033-3233130131023322-1022232100003213"></a>

### All schema paths for `xcsh_ike2`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike2--reference--group-001.md#canonical-3000100111320012-3121232131201023-3212012120101310-2102120013233200-1130103220200202-0011022022103203-1103031212331031-3110032223101122) |
| `description` | [description](resources--ike2--reference--group-001.md#canonical-2122332320220322-1230013111232123-1131232013100113-1133202123212203-1102011120311221-3200212210232103-1200110310020331-3322232301122032) |
| `dh_group_set` | [dh_group_set](resources--ike2--reference--group-001.md#canonical-1211323211200023-3202101300031233-1031122023213330-1110131323123010-1223302231211101-2112330221211223-0101000202230310-3233303210130002) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](resources--ike2--reference--group-001.md#canonical-1033201112133220-0300032202032103-3303333102112121-3310330012311313-2001003211010020-0132201000033000-2232113321213211-3133000332213331) |
| `disable` | [disable](resources--ike2--reference--group-001.md#canonical-2111123112312303-2331302122310122-3001111113200332-0010233202010230-2030301112022213-1130322012201221-0300110231301033-2321001101322031) |
| `disable_pfs` | [disable_pfs](resources--ike2--reference--group-001.md#canonical-2300022300133200-0100210031303130-0132023012032030-2031210122121031-0222200101311233-3223301320303010-0203333332321320-3121223331022002) |
| `id` | [ID](resources--ike2--reference--group-001.md#canonical-0222220132010301-3310222023103001-2131121012332233-0313330013233331-0113321200201210-2130322312100300-2312303332113313-3121030101210010) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-2133032203120332-1030012223331322-3113031330220313-1330323111003121-3203032123320133-2300213312321222-3231312313210113-0300030222220001) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike2--reference--group-001.md#canonical-1213332023203132-0121103312231030-3303332213303133-3230230232331201-3032300322131013-2000002113331011-0220100012230010-1212132232011331) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-3021313001222133-0202313302331213-0223201311010112-2012321220103331-1021132120120030-3111222311330222-0112131132123310-0113222021112332) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike2--reference--group-001.md#canonical-1130332110031130-3232330213202213-2310101102220111-0331023120111202-3300312020231300-0321330203323310-1132313320321310-2031130320002033) |
| `labels` | [labels](resources--ike2--reference--group-001.md#canonical-0211320022032131-3302030113022200-3022021311111200-0221332101323003-2010130320232032-3102321211221212-3020123301213131-3021102320022120) |
| `name` | [name](resources--ike2--reference--group-001.md#canonical-0023220222132021-0201023002100302-0211103012202132-3300020232013313-3311200111310310-2020112231312310-2133001302232330-3203122232130010) |
| `namespace` | [namespace](resources--ike2--reference--group-001.md#canonical-0202200101111000-3113123333313221-3023033333202230-1103131130131103-1333322003300001-2331232010100302-1110231200330110-1222020132230211) |
| `timeouts` | [timeouts](resources--ike2--reference--group-001.md#canonical-1320020213010133-2312230201210230-1320122000212130-3211033301031322-3230223221202322-1030312001332022-1212333322303000-0030312102310310) |
| `timeouts.create` | [timeouts.create](resources--ike2--reference--group-001.md#canonical-3200333211322233-3212333322221230-0332003221311122-0300330201123313-1120030101320222-0113013131312333-3221120111223110-1131002021302131) |
| `timeouts.delete` | [timeouts.delete](resources--ike2--reference--group-001.md#canonical-3322203330121021-2131121201211030-2013023132230003-1133313132301021-2300333213102112-3233313101003203-1121112022322200-0122212212332022) |
| `timeouts.read` | [timeouts.read](resources--ike2--reference--group-001.md#canonical-1001033103223311-3020230301020120-1212301203310313-1311121120333122-3320203112010232-3133100033311223-3300223232022330-3020110310113032) |
| `timeouts.update` | [timeouts.update](resources--ike2--reference--group-001.md#canonical-0201103333023300-1022133213030320-1132303110201020-3200202131032031-1131210222001201-3231201331310323-3030120102322003-0010022010321313) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-2202301222032201-3121303001020011-2132030002032110-0301211221303231-2332321122200222-1331022232311013-0310002132120133-1030332032130213) |

<a id="canonical-1132030321233131-0112333220310223-0120130001103312-1033130032311312-1222102313212001-3232201300232232-0030223131101212-3011200112322331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dh_group_set` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- dh_group_set

<a id="canonical-1211323211200023-3202101300031233-1031122023213330-1110131323123010-1223302231211101-2112330221211223-0101000202230310-3233303210130002"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Receipt-pinned upstream constraints:

```json
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

- [dh_group_set](resources--ike2--reference--group-001.md#canonical-1211323211200023-3202101300031233-1031122023213330-1110131323123010-1223302231211101-2112330221211223-0101000202230310-3233303210130002)
- [disable_pfs](resources--ike2--reference--group-001.md#canonical-2300022300133200-0100210031303130-0132023012032030-2031210122121031-0222200101311233-3223301320303010-0203333332321320-3121223331022002)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dh_group_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311330131302203-3013231311211101-2301231103133032-3031203332320101-2103122012003111-2003022311312002-1100132120332122-0331113322022001"></a>

### Direct properties for `dh_group_set`

<a id="canonical-1033201112133220-0300032202032103-3303333102112121-3310330012311313-2001003211010020-0132201000033000-2232113321213211-3133000332213331"></a>

#### `dh_group_set.dh_groups` property

Type: `["list", "string"]`. Optional.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Diffie Hellman Groups. Group or collection configuration. Possible values are
\`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`, \`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`,
\`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`, \`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`.
Defaults to \`DH\_GROUP\_DEFAULT\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3123101102221122-2012103322303011-1203112322120320-3320111010111123-2232303211033321-3302202122311112-0322000001010023-2203320122033312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_pfs` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- disable_pfs

<a id="canonical-2300022300133200-0100210031303130-0132023012032030-2031210122121031-0222200101311233-3223301320303010-0203333332321320-3121223331022002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

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
disable_pfs = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233220103012212-3122103302311332-2213202110011333-0222131130103113-0311231022333011-3331312030001233-3022310301221023-0130120002011032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_hours` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- ike_keylifetime_hours

<a id="canonical-2133032203120332-1030012223331322-3113031330220313-1330323111003121-3203032123320133-2300213312321222-3231312313210113-0300030222220001"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Additional upstream details:

Input Hours.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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

- [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-2133032203120332-1030012223331322-3113031330220313-1330323111003121-3203032123320133-2300213312321222-3231312313210113-0300030222220001)
- [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-3021313001222133-0202313302331213-0223201311010112-2012321220103331-1021132120120030-3111222311330222-0112131132123310-0113222021112332)
- [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-2202301222032201-3121303001020011-2132030002032110-0301211221303231-2332321122200222-1331022232311013-0310002132120133-1030332032130213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032303110020223-0330111001102132-3233211203132131-3300102002220322-2300112111332303-1213311323013003-2223303112100131-1332322013232220"></a>

### Direct properties for `ike_keylifetime_hours`

<a id="canonical-1213332023203132-0121103312231030-3303332213303133-3230230232331201-3032300322131013-2000002113331011-0220100012230010-1212132232011331"></a>

#### `ike_keylifetime_hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3100020103332033-2200320203233111-0013310012201310-1000300323210330-3312223323313112-2110101022322212-2102301010112130-3110302112233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_minutes` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- ike_keylifetime_minutes

<a id="canonical-3021313001222133-0202313302331213-0223201311010112-2012321220103331-1021132120120030-3111222311330222-0112131132123310-0113222021112332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002303223303111-3000110100111212-3012220233220300-3112102002200303-1030131200301010-3312202301333131-3013130332100312-3222223000000111"></a>

### Direct properties for `ike_keylifetime_minutes`

<a id="canonical-1130332110031130-3232330213202213-2310101102220111-0331023120111202-3300312020231300-0321330203323310-1132313320321310-2031130320002033"></a>

#### `ike_keylifetime_minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-0213233033132133-0200210123301111-3220032211230311-1021220110033120-1113001221313131-1311002133000201-3311231331022313-1012202102322311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- timeouts

<a id="canonical-1320020213010133-2312230201210230-1320122000212130-3211033301031322-3230223221202322-1030312001332022-1212333322303000-0030312102310310"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222121001100101-2322030203120332-1213130000012332-3020111123322203-2100212212100013-1032031131202023-0121320003310211-1113131122231321"></a>

### Direct properties for `timeouts`

<a id="canonical-3200333211322233-3212333322221230-0332003221311122-0300330201123313-1120030101320222-0113013131312333-3221120111223110-1131002021302131"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3322203330121021-2131121201211030-2013023132230003-1133313132301021-2300333213102112-3233313101003203-1121112022322200-0122212212332022"></a>

<a id="canonical-1103113012202111-1130000123031330-1000220021000331-2321132311033131-2212312230203312-3031211302031110-0220012122132111-2022132133210321"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1001033103223311-3020230301020120-1212301203310313-1311121120333122-3320203112010232-3133100033311223-3300223232022330-3020110310113032"></a>

<a id="canonical-2121031312022020-0032100210202133-2302000303311031-2003232110213233-1120322012200030-3022003010033300-0301013302133110-0021322203302223"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0201103333023300-1022133213030320-1132303110201020-3200202131032031-1131210222001201-3231201331310323-3030120102322003-0010022010321313"></a>

<a id="canonical-2211321030013132-3312233001223012-0011013121030101-3221030012101020-2231221131122333-0322031132022312-2223112331310113-2320323200110021"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0203033021002212-3132000022022110-2300020230323322-3233000320123033-3332011010230320-2323323330132300-2212011002210320-0110220022320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_keylifetime` properties

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232)
- [Property reference](resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- use_default_keylifetime

<a id="canonical-2202301222032201-3121303001020011-2132030002032110-0301211221303231-2332321122200222-1331022232311013-0310002132120133-1030332032130213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use default keylifetime.

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
use_default_keylifetime = {}
```

This is an empty object or choice marker. It has no direct properties.
