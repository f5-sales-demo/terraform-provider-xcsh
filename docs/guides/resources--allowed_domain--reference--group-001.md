---
page_title: "xcsh_allowed_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain reference."
---

# xcsh_allowed_domain reference

<a id="canonical-2310331131020212-1322032001003131-0111302203012211-3201023230322111-0103212210213030-2323300331323331-1322110232101022-3321121103131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300)
- Property reference

<a id="canonical-3331101000030231-0330000332011121-2030230221232203-2332310320031132-1131132033330211-0303033312333022-3033301212100100-3013200100032301"></a>

### Direct properties for `xcsh_allowed_domain`

<a id="canonical-2321023201002031-0230133333120022-3310300033202323-2120012330130100-3311111113223030-0332121311123103-2121001301033200-3231013031021022"></a>

#### `allowed_domain` property

Type: `"string"`. Required.

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.
Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2301012231211323-2200120030003020-2001023033131310-1011101230202130-1012333333033100-1131210311000021-3233313322200231-0102102020222030"></a>

<a id="canonical-1310303310033212-2021333122001311-2021120021113313-3202132120123112-0311100011213313-0023202001200331-3322332111330122-0002113310131203"></a>

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

<a id="canonical-1102203003210213-0033223311203302-3111030302322200-3102122112011133-3113112023202022-0101210303021031-1130310133232012-2123002100223130"></a>

<a id="canonical-2310031022321010-3321233301122113-0230000310223302-0313103121303203-0121011310002030-3312303102330031-3013203001020101-3111112332200120"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0321300113211201-3112222111210333-2333200333122321-0211032232233101-3010223300123011-2330200312120132-2303011223113232-1232323032301120"></a>

<a id="canonical-3003330002200122-1111222202202321-3113223311203201-3301220000031301-0323301231100010-1331133321221123-0130313011313112-3022010010301023"></a>

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

<a id="canonical-1210120231121233-2121333321113132-1200301331023320-3303022220223002-1300301331302132-1033312301013012-0102211101333011-1010323231333120"></a>

<a id="canonical-3100320130021130-3323313123323110-3310021222230231-3210102113211200-3221212312133212-2130121130221232-1001120202233003-2330013320131013"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1133320200032213-0300231022011033-2101300211220033-0201233200203203-3310202011333232-0310303201113100-0011102231323230-1313331203322230"></a>

<a id="canonical-0231230120003020-2002003201303302-2011203132130120-1330021212001111-0103033120002033-1011230101100133-3213221211310110-0021003200001032"></a>

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

<a id="canonical-2313322223320332-2223303313023310-2312112212331230-1031010110131001-3323110110222121-0202311332021130-2210230322311133-1303131022022222"></a>

<a id="canonical-0113121312322230-0120011110322330-1112121312000031-3100220301301030-0222022220113333-2020130112110203-1122322303103222-0113312021323310"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Allowed Domain. Must be unique within the namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1103113331231122-2011130220311201-3010020122023212-1113101312321023-3030131032232112-1121112333310110-2132003121031110-0320233132323331"></a>

<a id="canonical-3000111323202110-3333222320231112-2020133321303103-2320232031220232-1303201020323130-2002301011121212-1120321301220001-3010212011023201"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Allowed Domain is created.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [timeouts](resources--allowed_domain--reference--group-001.md#canonical-0133103013313212-3101013322333211-2212003013031120-2013033122010330-0313300303001020-0210312331302011-2311313202120323-3010031322023112): complete subsection reference.

<a id="canonical-0011323112123233-0032311220030131-2321032101202332-2321223020121102-3213211003133002-0223311031233310-3012030110123013-0210020120010013"></a>

### All schema paths for `xcsh_allowed_domain`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_domain` | [allowed_domain](resources--allowed_domain--reference--group-001.md#canonical-2321023201002031-0230133333120022-3310300033202323-2120012330130100-3311111113223030-0332121311123103-2121001301033200-3231013031021022) |
| `annotations` | [annotations](resources--allowed_domain--reference--group-001.md#canonical-2301012231211323-2200120030003020-2001023033131310-1011101230202130-1012333333033100-1131210311000021-3233313322200231-0102102020222030) |
| `description` | [description](resources--allowed_domain--reference--group-001.md#canonical-1102203003210213-0033223311203302-3111030302322200-3102122112011133-3113112023202022-0101210303021031-1130310133232012-2123002100223130) |
| `disable` | [disable](resources--allowed_domain--reference--group-001.md#canonical-0321300113211201-3112222111210333-2333200333122321-0211032232233101-3010223300123011-2330200312120132-2303011223113232-1232323032301120) |
| `id` | [ID](resources--allowed_domain--reference--group-001.md#canonical-1210120231121233-2121333321113132-1200301331023320-3303022220223002-1300301331302132-1033312301013012-0102211101333011-1010323231333120) |
| `labels` | [labels](resources--allowed_domain--reference--group-001.md#canonical-1133320200032213-0300231022011033-2101300211220033-0201233200203203-3310202011333232-0310303201113100-0011102231323230-1313331203322230) |
| `name` | [name](resources--allowed_domain--reference--group-001.md#canonical-2313322223320332-2223303313023310-2312112212331230-1031010110131001-3323110110222121-0202311332021130-2210230322311133-1303131022022222) |
| `namespace` | [namespace](resources--allowed_domain--reference--group-001.md#canonical-1103113331231122-2011130220311201-3010020122023212-1113101312321023-3030131032232112-1121112333310110-2132003121031110-0320233132323331) |
| `timeouts` | [timeouts](resources--allowed_domain--reference--group-001.md#canonical-3133010321303012-2303312120202332-2200011300102021-1030012132332023-0021330022102030-3011330310213222-2211313320101300-0203020303231230) |
| `timeouts.create` | [timeouts.create](resources--allowed_domain--reference--group-001.md#canonical-1312232312213223-1320031210031332-3133331221230021-2303200111131002-0122120020231032-0131233021032020-0111030021121013-1121030112310021) |
| `timeouts.delete` | [timeouts.delete](resources--allowed_domain--reference--group-001.md#canonical-2133322120310110-1133021223000213-3122023112211011-2322223322223301-3131213102222313-1123201112210002-3103120001131020-1021322100303210) |
| `timeouts.read` | [timeouts.read](resources--allowed_domain--reference--group-001.md#canonical-2333102330130210-3313203000100023-0120132212331202-0230110132312311-1110103333001112-1011103301303023-2020332030232303-3303131311201030) |
| `timeouts.update` | [timeouts.update](resources--allowed_domain--reference--group-001.md#canonical-2123023122223333-0101102130210022-2210102310022310-0232030233020021-0220021132113133-0030023202131322-1321021231311312-3333013331301132) |

<a id="canonical-0133103013313212-3101013322333211-2212003013031120-2013033122010330-0313300303001020-0210312331302011-2311313202120323-3010031322023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300)
- [Property reference](resources--allowed_domain--reference--group-001.md#canonical-2310331131020212-1322032001003131-0111302203012211-3201023230322111-0103212210213030-2323300331323331-1322110232101022-3321121103131313)
- timeouts

<a id="canonical-3133010321303012-2303312120202332-2200011300102021-1030012132332023-0021330022102030-3011330310213222-2211313320101300-0203020303231230"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201122320132002-3303230023233130-2301320121202101-0333112312310300-3223322212331010-3200231203222332-2111012032022112-0003002220213202"></a>

### Direct properties for `timeouts`

<a id="canonical-1312232312213223-1320031210031332-3133331221230021-2303200111131002-0122120020231032-0131233021032020-0111030021121013-1121030112310021"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2133322120310110-1133021223000213-3122023112211011-2322223322223301-3131213102222313-1123201112210002-3103120001131020-1021322100303210"></a>

<a id="canonical-0321130033133030-1321213213301013-2130323311213313-2333333011002231-1031033001003011-3201021131300102-1020000131120313-1003122121321222"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2333102330130210-3313203000100023-0120132212331202-0230110132312311-1110103333001112-1011103301303023-2020332030232303-3303131311201030"></a>

<a id="canonical-1012301003221333-3132001110000133-3202112122322110-1210022302302022-0002121223003130-0232210200311030-1213331110231231-3223100120110201"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2123023122223333-0101102130210022-2210102310022310-0232030233020021-0220021132113133-0030023202131322-1321021231311312-3333013331301132"></a>

<a id="canonical-1311223013213323-1202331100031210-2232002322123101-0011333011112101-0031123030133001-1312022011323000-3003000303300113-2211222030133311"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
