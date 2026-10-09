---
page_title: "xcsh_bgp_asn_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set reference."
---

# xcsh_bgp_asn_set reference

<a id="canonical-2221001120333021-0023132212300332-0300031102333013-0012022101331031-3232222323021231-0110023232233332-0120221100112302-0102233011221200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-3010010321231122-0133331023021212-0330320220102302-0102122320130321-1110011002112302-3010211202300010-0031103320123032-3230133030121103)
- Property reference

<a id="canonical-2210302211231021-0221321023003022-3120033113021132-2102013231100110-3231003300022013-1301133022210220-1033221322120211-3203221031212222"></a>

### Direct properties for `xcsh_bgp_asn_set`

<a id="canonical-0312212201202123-0332213311303311-2122111123033033-1220232210200122-3311032220010012-3102010332222220-1332310123112020-0331122202111323"></a>

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

<a id="canonical-2021332132301132-0112303310002203-1303130302033022-2113130121211311-2110003010002232-3012100103200102-0320013033310201-2011321013021100"></a>

<a id="canonical-0202232022021310-0333321020332011-1213111302201101-3323333112232023-1310303113303210-2211232130020103-0002203002331031-1100303212030131"></a>

#### `as_numbers` property

Type: `["list", "number"]`. Required.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1322232112321233-2311030333222333-0331003120233033-1322322230202123-0203230311232222-3022001222013321-2223300120012101-0301011311033123"></a>

<a id="canonical-1033232010020221-1032332110101110-3211122110012323-1213011321210101-1003100320111003-2231032000211133-2133221031222301-3333223120112202"></a>

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

<a id="canonical-0030310132003021-0002023332022201-0022213000103131-2313102002001032-1313001333032300-2322010013303330-2333021230000322-1121330010321222"></a>

<a id="canonical-2213023122232321-3202012223211312-2203323021303001-2333320313122203-3303121322020331-1321220311002313-0320113303220212-0120022013220301"></a>

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

<a id="canonical-1003302312030210-2211101023033201-1000102222010032-1123332011220331-0200120001301130-1103121132331120-1103100310010012-0021201313103111"></a>

<a id="canonical-3200032322331203-2222233132300133-3101303222030330-3333123130002012-1103101103202020-3312123211130313-2031131203022133-1313233031031333"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2201130032320113-0031333120002203-3123022321332023-3323232032123303-0000003001230300-0113120110203112-1013323320000232-1131310203220301"></a>

<a id="canonical-2321030313123203-2332323211230110-2010010111113103-0130313103123322-3003103300332313-3100100303030310-0022012332322133-1010121103030330"></a>

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

<a id="canonical-0211323223311301-3121311003022301-0121320012102231-0112033330300113-3330213003233123-1310112302101300-3030232001301123-3210012030020031"></a>

<a id="canonical-1330231203133301-0320131011130132-1320333123310321-1000122031310133-0013132133320130-0232301221320301-3111113323132320-2223122121021020"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BGP Asn Set. Must be unique within the namespace.

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

<a id="canonical-0031020021333100-0212003333200121-3322210110303033-2112310300131102-2132013011201100-3120313131200003-2302000201103113-3212001020031231"></a>

<a id="canonical-1300322112203132-1230133111101130-2030332333032221-1013222113103122-3000132322221023-0103033300311311-1200233232122233-0120320322001111"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BGP Asn Set is created.

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

- [timeouts](resources--bgp_asn_set--reference--group-001.md#canonical-1002101111103310-1033330011222221-2210011131311112-2110231213210210-3333333231210333-0223020101301332-3103232220312003-1133200111301011): complete subsection reference.

<a id="canonical-2221302010330132-2331133102102101-3231132012121101-2103321213023103-3030010210110223-1323210320032213-1330003232103320-3011021001231333"></a>

### All schema paths for `xcsh_bgp_asn_set`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp_asn_set--reference--group-001.md#canonical-0312212201202123-0332213311303311-2122111123033033-1220232210200122-3311032220010012-3102010332222220-1332310123112020-0331122202111323) |
| `as_numbers` | [as_numbers](resources--bgp_asn_set--reference--group-001.md#canonical-2021332132301132-0112303310002203-1303130302033022-2113130121211311-2110003010002232-3012100103200102-0320013033310201-2011321013021100) |
| `description` | [description](resources--bgp_asn_set--reference--group-001.md#canonical-1322232112321233-2311030333222333-0331003120233033-1322322230202123-0203230311232222-3022001222013321-2223300120012101-0301011311033123) |
| `disable` | [disable](resources--bgp_asn_set--reference--group-001.md#canonical-0030310132003021-0002023332022201-0022213000103131-2313102002001032-1313001333032300-2322010013303330-2333021230000322-1121330010321222) |
| `id` | [ID](resources--bgp_asn_set--reference--group-001.md#canonical-1003302312030210-2211101023033201-1000102222010032-1123332011220331-0200120001301130-1103121132331120-1103100310010012-0021201313103111) |
| `labels` | [labels](resources--bgp_asn_set--reference--group-001.md#canonical-2201130032320113-0031333120002203-3123022321332023-3323232032123303-0000003001230300-0113120110203112-1013323320000232-1131310203220301) |
| `name` | [name](resources--bgp_asn_set--reference--group-001.md#canonical-0211323223311301-3121311003022301-0121320012102231-0112033330300113-3330213003233123-1310112302101300-3030232001301123-3210012030020031) |
| `namespace` | [namespace](resources--bgp_asn_set--reference--group-001.md#canonical-0031020021333100-0212003333200121-3322210110303033-2112310300131102-2132013011201100-3120313131200003-2302000201103113-3212001020031231) |
| `timeouts` | [timeouts](resources--bgp_asn_set--reference--group-001.md#canonical-1030322211210110-0200132210303012-3023310322312323-2020221233113132-0201131312121030-1213001330211011-0222210301130323-0311330210211030) |
| `timeouts.create` | [timeouts.create](resources--bgp_asn_set--reference--group-001.md#canonical-3323223122001203-0313333333122133-1121133322313013-2012332222230013-1110010300233011-2120103322033323-3122323223211112-0030121320001210) |
| `timeouts.delete` | [timeouts.delete](resources--bgp_asn_set--reference--group-001.md#canonical-1033102303203102-1300002320120211-0211230333110230-3030001101212320-2023002300131001-0111232103231310-3132002330222330-1222131112232322) |
| `timeouts.read` | [timeouts.read](resources--bgp_asn_set--reference--group-001.md#canonical-1212112312210210-1332001001011022-3131132323001010-3103113133321230-3323031310213133-1202122101133220-2201333311203120-2103201322233231) |
| `timeouts.update` | [timeouts.update](resources--bgp_asn_set--reference--group-001.md#canonical-2001120302230102-1030121120222221-2123332121131200-3032303321023021-3030031010231313-0311230323011222-2223233311113312-1331112111112221) |

<a id="canonical-1002101111103310-1033330011222221-2210011131311112-2110231213210210-3333333231210333-0223020101301332-3103232220312003-1133200111301011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-3010010321231122-0133331023021212-0330320220102302-0102122320130321-1110011002112302-3010211202300010-0031103320123032-3230133030121103)
- [Property reference](resources--bgp_asn_set--reference--group-001.md#canonical-2221001120333021-0023132212300332-0300031102333013-0012022101331031-3232222323021231-0110023232233332-0120221100112302-0102233011221200)
- timeouts

<a id="canonical-1030322211210110-0200132210303012-3023310322312323-2020221233113132-0201131312121030-1213001330211011-0222210301130323-0311330210211030"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222220031103030-1121111003300133-0000132032113232-1123102111031321-0021312233300210-1021123031030303-1223321322120321-2202013330103123"></a>

### Direct properties for `timeouts`

<a id="canonical-3323223122001203-0313333333122133-1121133322313013-2012332222230013-1110010300233011-2120103322033323-3122323223211112-0030121320001210"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1033102303203102-1300002320120211-0211230333110230-3030001101212320-2023002300131001-0111232103231310-3132002330222330-1222131112232322"></a>

<a id="canonical-3222110022311330-0332233031320332-2011103201013231-1022133123010230-2311110301031313-3210023132120322-1130102031221131-3032033103230201"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1212112312210210-1332001001011022-3131132323001010-3103113133321230-3323031310213133-1202122101133220-2201333311203120-2103201322233231"></a>

<a id="canonical-0013320201001103-3231023010231033-0223002212312001-1103322020103210-2321223233301213-0201111233233231-1303113211310013-3032032032031320"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2001120302230102-1030121120222221-2123332121131200-3032303321023021-3030031010231313-0311230323011222-2223233311113312-1331112111112221"></a>

<a id="canonical-1302002103322312-3102213301302021-0110001112203201-1310103210111011-2313332230302032-2111331130131301-1331010110012221-0133102100302210"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
