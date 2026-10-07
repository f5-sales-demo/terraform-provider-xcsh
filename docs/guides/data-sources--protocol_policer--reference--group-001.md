---
page_title: "xcsh_protocol_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer reference."
---

# xcsh_protocol_policer reference

<a id="canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- Property reference

<a id="canonical-3202321213203232-2031213213332122-0302021320312113-0220030101220032-0210021022320213-3100201310000310-2112022310322013-1232332230103120"></a>

### Direct properties for `xcsh_protocol_policer`

<a id="canonical-2211233020212233-1100033121313100-3311120110111032-2303130302001031-0011023100211201-3230031322302121-0010013322010003-2202011221220122"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-1022323322230201-3333122200303120-3022321001330133-3132202300330321-3231120022232303-0131200330111311-2132023223020120-3221111223123202"></a>

<a id="canonical-3222010332210333-1201220322332131-1320102303132111-3303220031230210-0032111020100011-3102113102112022-2312203230132010-1032031023212203"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ProtocolPolicer.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030121011221212-0330223113213220-2030202323102003-3302120131333011-2131200220133211-0322103021223111-3031312101210330-0300231110113123"></a>

<a id="canonical-3103103101200022-2321122301120210-2200213011313222-3012010210321203-3221213323223313-1100122300111003-3311023213011000-1232230012002220"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2110133330220103-3212101220201320-1223131203333303-0102132230223333-0131201020113030-2112110212221332-3213213101230111-0021333121123313"></a>

<a id="canonical-1013211312320101-0022101003230002-0122310330331232-2321031030223303-3300110033221103-2132323131212100-2120101220302223-1100012221301331"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-2121002223202223-0303010332323320-2301123301120030-3303231310231322-1222223012312330-0302110012313321-3101303113211010-3123313313303122"></a>

<a id="canonical-3123130011102110-0003013211002003-2312222132311103-3012133132201211-1001321232230233-1213010202022021-2332312123013101-2221130033011102"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ProtocolPolicer.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3300101002303320-1133023012112132-3301331012302231-2122021200312113-1320001023123132-2132101320012100-2011011302113220-0333012321023102"></a>

<a id="canonical-2213120110003000-1022302111022113-1323301010023133-1212132300002120-0112201130110321-1303130011213022-0131120022322033-3030312232113031"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the ProtocolPolicer exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031): complete subsection reference.

<a id="canonical-1131302030100002-3103213310022033-2322113002323132-3331220002231100-3301112323131220-0012330112022312-0102031100231320-0031111300031320"></a>

### All schema paths for `xcsh_protocol_policer`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--protocol_policer--reference--group-001.md#canonical-2211233020212233-1100033121313100-3311120110111032-2303130302001031-0011023100211201-3230031322302121-0010013322010003-2202011221220122) |
| `description` | [description](data-sources--protocol_policer--reference--group-001.md#canonical-1022323322230201-3333122200303120-3022321001330133-3132202300330321-3231120022232303-0131200330111311-2132023223020120-3221111223123202) |
| `id` | [ID](data-sources--protocol_policer--reference--group-001.md#canonical-0030121011221212-0330223113213220-2030202323102003-3302120131333011-2131200220133211-0322103021223111-3031312101210330-0300231110113123) |
| `labels` | [labels](data-sources--protocol_policer--reference--group-001.md#canonical-2110133330220103-3212101220201320-1223131203333303-0102132230223333-0131201020113030-2112110212221332-3213213101230111-0021333121123313) |
| `name` | [name](data-sources--protocol_policer--reference--group-001.md#canonical-2121002223202223-0303010332323320-2301123301120030-3303231310231322-1222223012312330-0302110012313321-3101303113211010-3123313313303122) |
| `namespace` | [namespace](data-sources--protocol_policer--reference--group-001.md#canonical-3300101002303320-1133023012112132-3301331012302231-2122021200312113-1320001023123132-2132101320012100-2011011302113220-0333012321023102) |
| `protocol_policer` | [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1232213220202310-3110310311202330-0023122120011230-0231303232211113-3001310112030333-3120201200322000-2132031130003100-0022121110112113) |
| `protocol_policer.policer` | [protocol_policer.policer](data-sources--protocol_policer--reference--group-001.md#canonical-2322203012020332-2112012112133113-0333113220110210-3301213311002131-1203120021221112-1211331310233303-2312022100113203-0013301331200131) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](data-sources--protocol_policer--reference--group-001.md#canonical-0331121333122013-0013013132102230-0011330320200303-2332321333002313-0332202100111132-2130102013132303-3211001312221330-3222103201212300) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](data-sources--protocol_policer--reference--group-001.md#canonical-2033323233122221-1033211230320100-1031030132130322-0013201002302221-2120202210001231-0200303103132210-0020322112101202-2203202203323020) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](data-sources--protocol_policer--reference--group-001.md#canonical-3123110311220221-0100033230231001-3230030102221120-1120011321111312-0133203213120323-1301023002322110-2331030331231321-0222230331120031) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](data-sources--protocol_policer--reference--group-001.md#canonical-1023200023212300-1321213333133122-3021330211100332-0300023101002012-1011311222101031-1001102010331121-0313012112103020-2233003001322331) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](data-sources--protocol_policer--reference--group-001.md#canonical-0210031100013002-3300130200113000-1320313020001001-1131300333002030-1131130003111032-0330301222301333-1112120201113021-1123101330211022) |
| `protocol_policer.protocol` | [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-3022201232302131-2113023312032121-1023320322221022-1312220112211020-2131211100221110-2221020031001010-3033003000223033-2120103011321033) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](data-sources--protocol_policer--reference--group-001.md#canonical-0023132023223221-1121032023333210-2333133131113020-3110220301132330-2002321011222320-2223221323232311-1101323103200330-1010022113130200) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](data-sources--protocol_policer--reference--group-001.md#canonical-0012301010101010-0330021031303200-0131003222102103-3023001332103233-2211332200102220-1301300033113021-0101031111320031-1220003111231333) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](data-sources--protocol_policer--reference--group-001.md#canonical-1312030023110323-2230002102232020-1032122311230121-2310233002300313-1233303310021320-0132323021223332-3111223123021012-0102130311010022) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](data-sources--protocol_policer--reference--group-001.md#canonical-1112122211311010-1130303131223203-3132330230021320-1230030212030020-1131011231021110-3322203123233331-3332301021233330-1230033112022130) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](data-sources--protocol_policer--reference--group-001.md#canonical-1023101102331111-2232110202013000-2112323231233120-2000212212023000-3013221232120310-3213023002230110-1002113101310221-3020313223012210) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](data-sources--protocol_policer--reference--group-001.md#canonical-3313020323132333-0101230023222330-1013321023223111-2000023310113102-2101210230132123-2122312232330120-2300121321111203-3233002032300210) |

<a id="canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- protocol_policer

<a id="canonical-1232213220202310-3110310311202330-0023122120011230-0231303232211113-3001310112030333-3120201200322000-2132031130003100-0022121110112113"></a>

Type: `"list"`. Computed.

List of L4 protocol match condition and associated traffic rate limits.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033113023130301-1100232223113011-1010032233132213-1131300320231132-2030013332021322-0202033100021201-2310022111030303-2221313112123031"></a>

### Direct properties for `protocol_policer`

- [policer](data-sources--protocol_policer--reference--group-001.md#canonical-2132230333102203-0322122020020230-2120203031302132-1302300322113020-3232313111200032-0222220003012123-0120231201220120-3330310100303110): complete subsection reference.

- [protocol](data-sources--protocol_policer--reference--group-001.md#canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131): complete subsection reference.

<a id="canonical-2132230333102203-0322122020020230-2120203031302132-1302300322113020-3232313111200032-0222220003012123-0120231201220120-3330310100303110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.policer` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- protocol_policer.policer

<a id="canonical-2322203012020332-2112012112133113-0333113220110210-3301213311002131-1203120021221112-1211331310233303-2312022100113203-0013301331200131"></a>

Type: `"list"`. Computed.

Reference to policer object to apply traffic rate limits.

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

<a id="canonical-3201122321230311-1101122333000102-2322313212020121-2310132311033013-2110221120210132-1231012211123012-1033133100013220-2203130201113001"></a>

### Direct properties for `protocol_policer.policer`

<a id="canonical-0331121333122013-0013013132102230-0011330320200303-2332321333002313-0332202100111132-2130102013132303-3211001312221330-3222103201212300"></a>

#### `protocol_policer.policer.kind` property

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

<a id="canonical-2033323233122221-1033211230320100-1031030132130322-0013201002302221-2120202210001231-0200303103132210-0020322112101202-2203202203323020"></a>

<a id="canonical-2333223203132230-3301100102022021-3212300332203131-1120303031122213-1100200011211030-2112233110010201-2202202121232203-0230232123212301"></a>

#### `protocol_policer.policer.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3123110311220221-0100033230231001-3230030102221120-1120011321111312-0133203213120323-1301023002322110-2331030331231321-0222230331120031"></a>

<a id="canonical-2320321211033023-2302001103022000-0131333003111201-2101133210022232-0003020022023131-3103310220211022-1131123313320210-3212103232321023"></a>

#### `protocol_policer.policer.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1023200023212300-1321213333133122-3021330211100332-0300023101002012-1011311222101031-1001102010331121-0313012112103020-2233003001322331"></a>

<a id="canonical-2302200213103200-3021102321312010-2102121130033203-1321313310110220-1313013000330313-3233202000103010-3010221021201112-3310030213031323"></a>

#### `protocol_policer.policer.tenant` property

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

<a id="canonical-0210031100013002-3300130200113000-1320313020001001-1131300333002030-1131130003111032-0330301222301333-1112120201113021-1123101330211022"></a>

<a id="canonical-1323012130320121-3131222121030123-1320302212222213-0210331130020232-3211111013202002-2030111332333120-1112213101131111-0002202211132230"></a>

#### `protocol_policer.policer.uid` property

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

<a id="canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- protocol_policer.protocol

<a id="canonical-3022201232302131-2113023312032121-1023320322221022-1312220112211020-2131211100221110-2221020031001010-3033003000223033-2120103011321033"></a>

Type: `"single"`. Computed.

Protocol and protocol specific flags to be matched in packet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"dns\",\"icmp\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-0122023112102110-0131213020331020-1232001100322100-0223132001303203-0313100323123331-1321000111230302-0120333301223002-2031222203322011"></a>

### Direct properties for `protocol_policer.protocol`

- [DNS](data-sources--protocol_policer--reference--group-001.md#canonical-3102301222102230-1322023030013331-2203100131023231-1212203321232311-0100320103220032-0320213012012012-3222120130122321-3232030302213211): complete subsection reference.

- [icmp](data-sources--protocol_policer--reference--group-001.md#canonical-3010200132302232-3303103021120313-0321121120213302-2202132312133203-2223333213312221-0232313320330001-1032002102302323-3101123033311201): complete subsection reference.

- [tcp](data-sources--protocol_policer--reference--group-001.md#canonical-3001021011222003-2321001123332012-3102133112033010-3212213222211030-2121313332320032-0213311321231211-3312212131132220-1123101033122330): complete subsection reference.

- [udp](data-sources--protocol_policer--reference--group-001.md#canonical-1330111033333100-0003103222010321-2101033020101020-0301111332211012-3231000011212031-2332310100002032-1013211310221333-2310321231100130): complete subsection reference.

<a id="canonical-3102301222102230-1322023030013331-2203100131023231-1212203321232311-0100320103220032-0320213012012012-3222120130122321-3232030302213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.dns` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131)
- protocol_policer.protocol.DNS

<a id="canonical-0023132023223221-1121032023333210-2333133131113020-3110220301132330-2002321011222320-2223221323232311-1101323103200330-1010022113130200"></a>

Type: `["object", {}]`. Computed.

Match all DNS packets including UDP and TCP.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010200132302232-3303103021120313-0321121120213302-2202132312133203-2223333213312221-0232313320330001-1032002102302323-3101123033311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.icmp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131)
- protocol_policer.protocol.icmp

<a id="canonical-0012301010101010-0330021031303200-0131003222102103-3023001332103233-2211332200102220-1301300033113021-0101031111320031-1220003111231333"></a>

Type: `"single"`. Computed.

ICMP Packet Type. ICMP message type to match in packet.

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

<a id="canonical-1302221202122330-2000312210133011-3202320003220132-0131211210311002-0310203212000122-1302032210122100-0203131313021300-1220211001301122"></a>

### Direct properties for `protocol_policer.protocol.icmp`

<a id="canonical-1312030023110323-2230002102232020-1032122311230121-2310233002300313-1233303310021320-0132323021223332-3111223123021012-0102130311010022"></a>

#### `protocol_policer.protocol.icmp.type` property

Type: `["list", "string"]`. Computed.

\[Enum: ECHO\_REPLY|ECHO\_REQUEST|ALL\_ICMP\_MSG\] ICMP message type to be matched in packet.
Possible values are \`ECHO\_REPLY\`, \`ECHO\_REQUEST\`, \`ALL\_ICMP\_MSG\`. Defaults to
\`ECHO\_REPLY\`.

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

<a id="canonical-3001021011222003-2321001123332012-3102133112033010-3212213222211030-2121313332320032-0213311321231211-3312212131132220-1123101033122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.tcp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131)
- protocol_policer.protocol.tcp

<a id="canonical-1112122211311010-1130303131223203-3132330230021320-1230030212030020-1131011231021110-3322203123233331-3332301021233330-1230033112022130"></a>

Type: `"single"`. Computed.

Specification of TCP flag to be matched in a TCP packet.

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

<a id="canonical-2322332211333322-3110202132232331-3122120032233030-3200312203220211-0233013000110110-2112032330200003-3121333311221313-1201020131210103"></a>

### Direct properties for `protocol_policer.protocol.tcp`

<a id="canonical-1023101102331111-2232110202013000-2112323231233120-2000212212023000-3013221232120310-3213023002230110-1002113101310221-3020313223012210"></a>

#### `protocol_policer.protocol.tcp.flags` property

Type: `["list", "string"]`. Computed.

\[Enum: FIN|SYN|RST|PSH|ACK|URG|ALL\_TCP\_FLAGS|KEEPALIVE\] TCP flags. TCP flag to be matched in a
TCP packet. Possible values are \`FIN\`, \`SYN\`, \`RST\`, \`PSH\`, \`ACK\`, \`URG\`,
\`ALL\_TCP\_FLAGS\`, \`KEEPALIVE\`. Defaults to \`FIN\`.

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

<a id="canonical-1330111033333100-0003103222010321-2101033020101020-0301111332211012-3231000011212031-2332310100002032-1013211310221333-2310321231100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protocol_policer.protocol.udp` properties

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-1101013031120213-1133133221120212-3213211320100310-3121322130301110-0301300000312303-0220121310301021-1102321132231311-1031011233101031)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131)
- protocol_policer.protocol.udp

<a id="canonical-3313020323132333-0101230023222330-1013321023223111-2000023310113102-2101210230132123-2122312232330120-2300121321111203-3233002032300210"></a>

Type: `["object", {}]`. Computed.

UDP Packets. Match all UDP packets.

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

This is an empty object or choice marker. It has no direct properties.
