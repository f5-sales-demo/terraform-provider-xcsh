---
page_title: "xcsh_certificate_chain reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain reference."
---

# xcsh_certificate_chain reference

<a id="canonical-3021131230120120-2001003003312102-0331033112303122-3121321212202103-3132210200101210-2100113131211103-3003002012030023-2121301110001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322)
- Property reference

<a id="canonical-2332331210300310-2210022220013001-0130301132311123-1210111001201301-2113323022210312-0211130032023101-0310332333213112-3331103132230311"></a>

### Direct properties for `xcsh_certificate_chain`

<a id="canonical-0030112233131233-0303021313123303-0000021132200223-1230132121303020-1112213201132211-0111322213300131-1131332312101031-3020212010202213"></a>

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

<a id="canonical-3013131231001110-0232203310110002-3220200331010003-0001300011123013-3300332302102012-0011211222022110-1021211113232201-2320023131021010"></a>

<a id="canonical-0031232023313322-3032001221323001-2122230203301003-1220011033323031-0312212300102302-1202213220311301-1011012001220032-2321033103130120"></a>

#### `certificate_url` property

Type: `"string"`. Required.

Certificate chain is the list of intermediate certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3321233020002113-1202120130300013-0330312233002103-2313233013022032-3103112320133232-1002111313120000-3303233001213203-3231100320321211"></a>

<a id="canonical-0321111122200212-3322030313101003-2031331211310230-3122121303101230-0222133101310231-0023333031032332-2031101230301011-3212000003312232"></a>

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

<a id="canonical-1023213100230223-2132111030111323-0013002310110010-2210021312311020-0230031031002303-3113231222023233-2321103212203120-0013231011121021"></a>

<a id="canonical-0023322000301122-2223021023101322-2032333030003112-2320013322023133-1313001032301333-1303131310132110-0031212131310022-0123111113111012"></a>

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

<a id="canonical-3200322200222020-3101323120011100-3000111001113000-1122300121102211-2211023221121130-3001312230333003-3301203113203123-3332203312033102"></a>

<a id="canonical-1320311010133312-3303322022330110-0013312001011120-0222113312023011-1330020203201310-2022333022202132-2321211211221002-1121321232233233"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1302222002202212-0021002202023232-3322223003321213-1010322103120230-0311320110121330-0103031323233121-0210200322222130-0012230011021233"></a>

<a id="canonical-3222222002221133-0200013323033222-3232321303120002-1200303003113301-2130132321133033-0310321103103220-1313030011011303-3102011333033200"></a>

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

<a id="canonical-3333123231221031-2010233202002201-0213201021123202-3203333020211122-2122102012203122-3112122221113123-1003022100300001-1101131100003112"></a>

<a id="canonical-2202102132113103-3233130221023100-0211213312223200-1132312000320013-0033300322023110-3202201231003102-0210333002001322-1210023301022321"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Certificate Chain. Must be unique within the namespace.

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

<a id="canonical-1121301020113030-2232112100133030-2321122020312010-3131311323212113-0132001320012131-2011021133112221-0021212011213313-1033212130013122"></a>

<a id="canonical-1222123200301120-3012210330021332-0321001333332102-3320112003030322-1310220203132231-2230130321321133-0131300032212300-1121021312302331"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Certificate Chain is created.

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

- [timeouts](resources--certificate_chain--reference--group-001.md#canonical-1301121200323322-1002233302113101-0302201121221101-2311303131132232-3020133003311033-0110322002333132-2031103203320320-3022201130110221): complete subsection reference.

<a id="canonical-2320100203230033-1232310332020331-0323222021233113-2123212000331133-3120130001023101-2120123011122323-0200213302123001-3000332003301202"></a>

### All schema paths for `xcsh_certificate_chain`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--certificate_chain--reference--group-001.md#canonical-0030112233131233-0303021313123303-0000021132200223-1230132121303020-1112213201132211-0111322213300131-1131332312101031-3020212010202213) |
| `certificate_url` | [certificate_url](resources--certificate_chain--reference--group-001.md#canonical-3013131231001110-0232203310110002-3220200331010003-0001300011123013-3300332302102012-0011211222022110-1021211113232201-2320023131021010) |
| `description` | [description](resources--certificate_chain--reference--group-001.md#canonical-3321233020002113-1202120130300013-0330312233002103-2313233013022032-3103112320133232-1002111313120000-3303233001213203-3231100320321211) |
| `disable` | [disable](resources--certificate_chain--reference--group-001.md#canonical-1023213100230223-2132111030111323-0013002310110010-2210021312311020-0230031031002303-3113231222023233-2321103212203120-0013231011121021) |
| `id` | [ID](resources--certificate_chain--reference--group-001.md#canonical-3200322200222020-3101323120011100-3000111001113000-1122300121102211-2211023221121130-3001312230333003-3301203113203123-3332203312033102) |
| `labels` | [labels](resources--certificate_chain--reference--group-001.md#canonical-1302222002202212-0021002202023232-3322223003321213-1010322103120230-0311320110121330-0103031323233121-0210200322222130-0012230011021233) |
| `name` | [name](resources--certificate_chain--reference--group-001.md#canonical-3333123231221031-2010233202002201-0213201021123202-3203333020211122-2122102012203122-3112122221113123-1003022100300001-1101131100003112) |
| `namespace` | [namespace](resources--certificate_chain--reference--group-001.md#canonical-1121301020113030-2232112100133030-2321122020312010-3131311323212113-0132001320012131-2011021133112221-0021212011213313-1033212130013122) |
| `timeouts` | [timeouts](resources--certificate_chain--reference--group-001.md#canonical-3003311030011131-2012302212132121-2122022213212121-1233320200200102-3210121110302022-1330030002032033-2032031330013310-0202023101301323) |
| `timeouts.create` | [timeouts.create](resources--certificate_chain--reference--group-001.md#canonical-3223121312202310-2301110010310222-3221002033311230-0011121112221033-1210330330210222-1113030233120123-1301302023013333-1320332132031202) |
| `timeouts.delete` | [timeouts.delete](resources--certificate_chain--reference--group-001.md#canonical-2233213322003031-2302222110001200-0000122233011100-0122211113313012-3033001011231200-1121110112021001-3202122231332321-0101331021111310) |
| `timeouts.read` | [timeouts.read](resources--certificate_chain--reference--group-001.md#canonical-2332202332110212-2213313001122030-1300301021221122-0230212130311010-0100303021233003-3222222003221300-3201131032013323-3202110232021020) |
| `timeouts.update` | [timeouts.update](resources--certificate_chain--reference--group-001.md#canonical-0300131033131021-1030030232033210-1010212303301223-2322120221023211-1000312011213223-3221322122200303-2010200300333121-2103002100121212) |

<a id="canonical-1301121200323322-1002233302113101-0302201121221101-2311303131132232-3020133003311033-0110322002333132-2031103203320320-3022201130110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322)
- [Property reference](resources--certificate_chain--reference--group-001.md#canonical-3021131230120120-2001003003312102-0331033112303122-3121321212202103-3132210200101210-2100113131211103-3003002012030023-2121301110001321)
- timeouts

<a id="canonical-3003311030011131-2012302212132121-2122022213212121-1233320200200102-3210121110302022-1330030002032033-2032031330013310-0202023101301323"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032022111310202-0132121010303131-2311233230313330-2110202333301121-3203130133221103-0012130023023033-1231212101222320-0201312131331000"></a>

### Direct properties for `timeouts`

<a id="canonical-3223121312202310-2301110010310222-3221002033311230-0011121112221033-1210330330210222-1113030233120123-1301302023013333-1320332132031202"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2233213322003031-2302222110001200-0000122233011100-0122211113313012-3033001011231200-1121110112021001-3202122231332321-0101331021111310"></a>

<a id="canonical-1100021313002112-2021112003010110-0320302000012001-0031123230233321-1201313133233023-2200332031131311-0110222320202121-0312300003110311"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2332202332110212-2213313001122030-1300301021221122-0230212130311010-0100303021233003-3222222003221300-3201131032013323-3202110232021020"></a>

<a id="canonical-3000222220210013-3220330010110203-2331311033300130-3223023102232330-2033322033113110-0123202130300030-1010233121323331-2311131221003330"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0300131033131021-1030030232033210-1010212303301223-2322120221023211-1000312011213223-3221322122200303-2010200300333121-2103002100121212"></a>

<a id="canonical-0020033222230222-2022113131213032-0131123300122120-3201130120020213-2232130303331131-2020003331022302-2312311033311112-3302031232301123"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
