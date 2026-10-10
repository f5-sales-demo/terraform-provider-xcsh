---
page_title: "xcsh_trusted_ca_list reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list reference."
---

# xcsh_trusted_ca_list reference

<a id="canonical-2011100021221213-2030030021232013-1113200112210011-3323103123123311-0202132011000133-0020212102203021-2031323023203030-2231003231000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-0222000113113203-2212232323102203-2101230202220010-2030322030013021-2133220023123121-2113200100110031-3003002313103330-2200302323113112)
- Property reference

<a id="canonical-3233020002232232-0221312221333121-1321021212110213-3302231321122101-0301232313000320-0102310110121020-1302313032220111-3121212212020002"></a>

### Direct properties for `xcsh_trusted_ca_list`

<a id="canonical-2333102201310210-3323330030310333-1023133320112230-1112000031232302-2112013312201110-3101121233101021-1221000031030333-2111132303310213"></a>

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

<a id="canonical-0101321202103101-0321033020001322-3311232011213133-3120003322303123-0311202022303310-1131222022210300-3112022113320211-3322233032230202"></a>

<a id="canonical-3213021303223020-3211301000230003-3322133023001132-2312133200033012-1230000220211011-2012003010200211-0003121212130000-1011301310022332"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1133010130103020-0222120222120000-1201100211130112-2113010000230133-2011131301331012-2032110332122133-2333211323200201-2020321301321201"></a>

<a id="canonical-2011302111303220-1213102000303130-2103300311310121-2021003322031333-0130302202220003-2213111112023120-0322130312031312-1123300310233233"></a>

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

<a id="canonical-2123133200323121-0323011130012303-2123303001000111-0001030303032321-2001110313111233-3100232112331033-1222211010001023-0210002222010010"></a>

<a id="canonical-2210033321110011-3010331233100012-2122210132013332-0210121323211213-3110332010300211-2232003111010022-2230131200313133-3220211023203000"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0012033023221000-1130223121222101-0323313002333012-1031003111123331-3310112032331220-0030222111203030-2313113113023220-2331301000332013"></a>

<a id="canonical-1302120333332123-2131320011333021-1200300131112233-3033201313331332-2331113222001213-3120320232003120-3203110110232303-1022203210100221"></a>

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

<a id="canonical-2212221020311233-3131200323313111-0301332112111010-3002021020221131-2101220313201012-2003022010322021-3020323003131030-0130233003322200"></a>

<a id="canonical-1322103130131232-0020023110311301-0002110213010132-2032322312120033-0002202132201213-0212001212110220-1222233032231021-2211033201323030"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Trusted CA List. Must be unique within the namespace.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3031331233121222-3102031200003232-0232323320102003-1020330012013322-1213303231132133-1123003130220110-0231222002330113-2101330022332022"></a>

<a id="canonical-3031012101311021-0121131111123210-3021312032301111-1331100222111002-0020322113321332-1320312133020133-0120231322102100-2123223103002310"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Trusted CA List is created.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--trusted_ca_list--reference--group-001.md#canonical-3203210011233323-1022230202020011-3301302003301121-1113323201212111-2130030203222223-2222021013203211-3333100121100133-3133102001231313): complete subsection reference.

<a id="canonical-3332330300011230-0002013111020222-1110321222301031-0012200002203113-3030033112321210-0110101100222113-2001000032010123-1131022322221001"></a>

<a id="canonical-3201210020032331-0003010312020221-2200132002032220-2230302221322122-1023220333021212-1002213212230301-0210020211322033-3031303121022233"></a>

#### `trusted_ca_url` property

Type: `"string"`. Optional, Computed.

Trusted CA certificates for validating certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(512000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512000,
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
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3212323201222230-2023232200331130-3202302201032121-3230010102111133-3020003313113032-1130301122210011-0313113222112322-0301200030222003"></a>

### All schema paths for `xcsh_trusted_ca_list`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--trusted_ca_list--reference--group-001.md#canonical-2333102201310210-3323330030310333-1023133320112230-1112000031232302-2112013312201110-3101121233101021-1221000031030333-2111132303310213) |
| `description` | [description](resources--trusted_ca_list--reference--group-001.md#canonical-0101321202103101-0321033020001322-3311232011213133-3120003322303123-0311202022303310-1131222022210300-3112022113320211-3322233032230202) |
| `disable` | [disable](resources--trusted_ca_list--reference--group-001.md#canonical-1133010130103020-0222120222120000-1201100211130112-2113010000230133-2011131301331012-2032110332122133-2333211323200201-2020321301321201) |
| `id` | [ID](resources--trusted_ca_list--reference--group-001.md#canonical-2123133200323121-0323011130012303-2123303001000111-0001030303032321-2001110313111233-3100232112331033-1222211010001023-0210002222010010) |
| `labels` | [labels](resources--trusted_ca_list--reference--group-001.md#canonical-0012033023221000-1130223121222101-0323313002333012-1031003111123331-3310112032331220-0030222111203030-2313113113023220-2331301000332013) |
| `name` | [name](resources--trusted_ca_list--reference--group-001.md#canonical-2212221020311233-3131200323313111-0301332112111010-3002021020221131-2101220313201012-2003022010322021-3020323003131030-0130233003322200) |
| `namespace` | [namespace](resources--trusted_ca_list--reference--group-001.md#canonical-3031331233121222-3102031200003232-0232323320102003-1020330012013322-1213303231132133-1123003130220110-0231222002330113-2101330022332022) |
| `timeouts` | [timeouts](resources--trusted_ca_list--reference--group-001.md#canonical-0032001131310223-2123110302002223-0010320002111230-1301000320121110-2102013333131332-2132332201030222-1022311221233231-3312310012201013) |
| `timeouts.create` | [timeouts.create](resources--trusted_ca_list--reference--group-001.md#canonical-0003111022202303-2300121011003232-1023001203111121-3002213223231110-0213220012021022-2102211130231023-1232120312020313-1332020033300203) |
| `timeouts.delete` | [timeouts.delete](resources--trusted_ca_list--reference--group-001.md#canonical-0130033221201320-1031220221331221-3303032111013300-1233012103101122-3330032133320001-0120302331131012-1302130001323320-3212312332231301) |
| `timeouts.read` | [timeouts.read](resources--trusted_ca_list--reference--group-001.md#canonical-3010211111230300-0310211210201301-0201022222030213-2320012130231013-2122323303330012-3112323122100133-2212333111133030-2311300133012002) |
| `timeouts.update` | [timeouts.update](resources--trusted_ca_list--reference--group-001.md#canonical-3302231002322321-1101130231203110-2321000110013133-0033211130103221-0032030030133230-1312031011232233-0200102303223012-3201330032031023) |
| `trusted_ca_url` | [trusted_ca_url](resources--trusted_ca_list--reference--group-001.md#canonical-3332330300011230-0002013111020222-1110321222301031-0012200002203113-3030033112321210-0110101100222113-2001000032010123-1131022322221001) |

<a id="canonical-3203210011233323-1022230202020011-3301302003301121-1113323201212111-2130030203222223-2222021013203211-3333100121100133-3133102001231313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-0222000113113203-2212232323102203-2101230202220010-2030322030013021-2133220023123121-2113200100110031-3003002313103330-2200302323113112)
- [Property reference](resources--trusted_ca_list--reference--group-001.md#canonical-2011100021221213-2030030021232013-1113200112210011-3323103123123311-0202132011000133-0020212102203021-2031323023203030-2231003231000112)
- timeouts

<a id="canonical-0032001131310223-2123110302002223-0010320002111230-1301000320121110-2102013333131332-2132332201030222-1022311221233231-3312310012201013"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211022132023221-1233233203113102-0013212011232213-3002231300212210-2200101111302321-0023221133333101-3331030003021222-2201030121022331"></a>

### Direct properties for `timeouts`

<a id="canonical-0003111022202303-2300121011003232-1023001203111121-3002213223231110-0213220012021022-2102211130231023-1232120312020313-1332020033300203"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0130033221201320-1031220221331221-3303032111013300-1233012103101122-3330032133320001-0120302331131012-1302130001323320-3212312332231301"></a>

<a id="canonical-2301020331323020-2122011233133132-3233133323330333-1130011112210032-2322120222331021-0101011101021201-1131301310322222-1313203220000333"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3010211111230300-0310211210201301-0201022222030213-2320012130231013-2122323303330012-3112323122100133-2212333111133030-2311300133012002"></a>

<a id="canonical-1202222021010012-1231202102223203-0020232220031300-0001201110210332-3330301012311023-1220123132322032-0223302233230233-0033331302132203"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3302231002322321-1101130231203110-2321000110013133-0033211130103221-0032030030133230-1312031011232233-0200102303223012-3201330032031023"></a>

<a id="canonical-1222312312222021-2101011200333303-0122110322110202-1201213130030301-2103203032000003-0122120233301131-2113013031230121-0113332031032313"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
