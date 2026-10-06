---
page_title: "xcsh_alert_template reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template reference."
---

# xcsh_alert_template reference

<a id="canonical-1133313323321313-1021331100132113-3110202222120000-0021012302223200-2222322030331020-2031011132012122-0322111010203330-2101232233133233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113)
- Property reference

<a id="canonical-0300020023313112-3110031332022023-2001002111230000-1001001201323231-0100221023222330-3213331020103222-0201110203023321-2323003023311122"></a>

### Direct properties for `xcsh_alert_template`

<a id="canonical-2233030133032120-1012132310303020-0200302203022112-3022201332322011-3211230230113223-3331113021032020-0011203212111003-3023132111312003"></a>

#### `alert_message` property

Type: `"string"`. Required.

Alert Message. Alert Message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-0203020311200313-3222011112332101-1022232223311232-0022211211110321-0213201013200020-0201121332312220-0100300302330102-3213220232212131"></a>

<a id="canonical-1130231302030031-3232330332302233-2120331002312011-2123011222000302-3000021000121213-3321220321131321-0011113033321303-2022212031333013"></a>

#### `alert_message_details` property

Type: `"string"`. Required.

Alert Message Details. Detailed message of the alert.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1023031211113123-1222213002101310-2001111131233220-0311332210302200-3323131010022010-0110203211123300-3100210200010320-3003233002101310"></a>

<a id="canonical-3120001312002131-0123132210303233-0201203300302110-1002121112032213-3202020300012101-2130001111312023-2133231203133203-3120103331103102"></a>

#### `alert_name` property

Type: `"string"`. Required.

Alert Name. Alert Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-3333313210313330-2202010030103301-3231021030032200-1223210231132121-1120232112111330-2230000322210201-1203302222312213-0033311201223231"></a>

<a id="canonical-0323001230332030-0301123223220111-3212131202023032-2011303030112031-3120300100220300-1011333320131331-3202032032301300-3122301232022120"></a>

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

<a id="canonical-3001103320301103-2330203111313201-1212213110232310-1213030330023232-1033323102203131-1012320333130000-1313132012013011-1221213010102030"></a>

<a id="canonical-3321103313102132-1231313022132203-3033333200320022-0311100333230302-0310010020022032-1112322211302332-1222223322111123-3210000123211003"></a>

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

<a id="canonical-1000131110023121-0211212103103131-3301012131300303-3321202332201103-3223211311112222-1113130220112303-1012212202330130-2221331130100300"></a>

<a id="canonical-2213120103211111-3231103331100312-2121111321230132-1330323230022113-1030313113322031-1202121221330232-1012301311331203-0130132310102003"></a>

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

<a id="canonical-3230121231322121-3332321031133000-2200301033301232-1022211001100020-2033033312303030-1000021213223230-0200302301131003-2123201301222011"></a>

<a id="canonical-1133230111103221-2212132012222223-1031001110001331-2013220012333311-1231022010031230-1312331233213120-3120123231331212-2003211120000120"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3113111101211201-0022123102310203-2000313003130210-2310323130000301-2111133113020020-1032111033013001-0201213021122113-3322313221323303"></a>

<a id="canonical-2300311311002310-2301200002332331-0022031023300333-3112232313010330-1322100011113303-3223102332030133-1212003213222212-1130303031332302"></a>

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

<a id="canonical-3221300202131200-2131313012211231-0200322130323000-3331320301011002-3211211122021113-3132212222310123-2211022223012102-1203121301121222"></a>

<a id="canonical-1200321111222222-0300131030003132-3330033013022033-1210323220333310-2130301203211220-0210020213322000-3301310300210210-1311020312110202"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Alert Template. Must be unique within the namespace.

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

<a id="canonical-1323002031232230-2301021002301131-1032131222100231-1323121022122133-3321010002100022-2012133311132202-2103301001003133-0330223213222220"></a>

<a id="canonical-3021302033330112-0020232302321020-3102311211033030-0211321203233000-3002103132210202-2222211312203301-0212010110001201-2210012333303031"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Alert Template is created.

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

<a id="canonical-2310113033122132-0300300322030013-1023332332011123-1330312201331131-1113323222030302-2021022320203103-3131120312223012-1112021133021331"></a>

<a id="canonical-2203000212112123-0203200102231022-3101333123202131-0120300121100312-1123112310302031-1303102231321223-3013312310332321-0320230333213120"></a>

#### `severity` property

Type: `"string"`. Optional, Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CRITICAL","MAJOR","MINOR"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("MINOR",
    "MAJOR",
    "CRITICAL"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "MINOR",
  "enum": [
    "MINOR",
    "MAJOR",
    "CRITICAL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--alert_template--reference--group-001.md#canonical-2320323033312231-0231121320223030-0122312332202102-2230311233010203-2101322220112001-2022220233000003-3022200323123230-2221323010212232): complete subsection reference.

<a id="canonical-2113323303330013-3310103000023323-3001112310030312-1230030031132132-0333210320222213-3301122032031100-1120112331222103-3132230231013030"></a>

### All schema paths for `xcsh_alert_template`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_message` | [alert_message](resources--alert_template--reference--group-001.md#canonical-2233030133032120-1012132310303020-0200302203022112-3022201332322011-3211230230113223-3331113021032020-0011203212111003-3023132111312003) |
| `alert_message_details` | [alert_message_details](resources--alert_template--reference--group-001.md#canonical-0203020311200313-3222011112332101-1022232223311232-0022211211110321-0213201013200020-0201121332312220-0100300302330102-3213220232212131) |
| `alert_name` | [alert_name](resources--alert_template--reference--group-001.md#canonical-1023031211113123-1222213002101310-2001111131233220-0311332210302200-3323131010022010-0110203211123300-3100210200010320-3003233002101310) |
| `annotations` | [annotations](resources--alert_template--reference--group-001.md#canonical-3333313210313330-2202010030103301-3231021030032200-1223210231132121-1120232112111330-2230000322210201-1203302222312213-0033311201223231) |
| `description` | [description](resources--alert_template--reference--group-001.md#canonical-3001103320301103-2330203111313201-1212213110232310-1213030330023232-1033323102203131-1012320333130000-1313132012013011-1221213010102030) |
| `disable` | [disable](resources--alert_template--reference--group-001.md#canonical-1000131110023121-0211212103103131-3301012131300303-3321202332201103-3223211311112222-1113130220112303-1012212202330130-2221331130100300) |
| `id` | [ID](resources--alert_template--reference--group-001.md#canonical-3230121231322121-3332321031133000-2200301033301232-1022211001100020-2033033312303030-1000021213223230-0200302301131003-2123201301222011) |
| `labels` | [labels](resources--alert_template--reference--group-001.md#canonical-3113111101211201-0022123102310203-2000313003130210-2310323130000301-2111133113020020-1032111033013001-0201213021122113-3322313221323303) |
| `name` | [name](resources--alert_template--reference--group-001.md#canonical-3221300202131200-2131313012211231-0200322130323000-3331320301011002-3211211122021113-3132212222310123-2211022223012102-1203121301121222) |
| `namespace` | [namespace](resources--alert_template--reference--group-001.md#canonical-1323002031232230-2301021002301131-1032131222100231-1323121022122133-3321010002100022-2012133311132202-2103301001003133-0330223213222220) |
| `severity` | [severity](resources--alert_template--reference--group-001.md#canonical-2310113033122132-0300300322030013-1023332332011123-1330312201331131-1113323222030302-2021022320203103-3131120312223012-1112021133021331) |
| `timeouts` | [timeouts](resources--alert_template--reference--group-001.md#canonical-2001100130320100-3200233123000023-0030330313221333-0102321201121023-2212010222300111-0032233230321202-2233102113321300-2213100201231021) |
| `timeouts.create` | [timeouts.create](resources--alert_template--reference--group-001.md#canonical-0202130131302100-3100021330303011-3212202020001022-2110201122333032-0012231203021212-1013201300002110-0223130102333232-2222020102030300) |
| `timeouts.delete` | [timeouts.delete](resources--alert_template--reference--group-001.md#canonical-0311011212300212-2002222322121012-3203311101322323-2121233123130132-1312200001230330-1023020130301211-3133033223223230-0220113323010303) |
| `timeouts.read` | [timeouts.read](resources--alert_template--reference--group-001.md#canonical-1132000010333030-0112102032303230-1330310100200323-3333331011023331-0133022032211113-0300021303213132-2322322213010223-1300333210200220) |
| `timeouts.update` | [timeouts.update](resources--alert_template--reference--group-001.md#canonical-0223333313103100-2321032022003022-0011331200332030-3313112223201333-3303121311302323-1310211323221101-2302132132130100-0311211220302213) |

<a id="canonical-2320323033312231-0231121320223030-0122312332202102-2230311233010203-2101322220112001-2022220233000003-3022200323123230-2221323010212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113)
- [Property reference](resources--alert_template--reference--group-001.md#canonical-1133313323321313-1021331100132113-3110202222120000-0021012302223200-2222322030331020-2031011132012122-0322111010203330-2101232233133233)
- timeouts

<a id="canonical-2001100130320100-3200233123000023-0030330313221333-0102321201121023-2212010222300111-0032233230321202-2233102113321300-2213100201231021"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313030122210023-1320301232133222-1233230210203000-1030210231320221-2131121003303003-2022201003010300-0313010011102111-1202132203201103"></a>

### Direct properties for `timeouts`

<a id="canonical-0202130131302100-3100021330303011-3212202020001022-2110201122333032-0012231203021212-1013201300002110-0223130102333232-2222020102030300"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0311011212300212-2002222322121012-3203311101322323-2121233123130132-1312200001230330-1023020130301211-3133033223223230-0220113323010303"></a>

<a id="canonical-2001213232311000-0000233302222303-2230220213003132-3031021322330211-1012201311012322-3232032103003301-1323220031022102-3222002103233120"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1132000010333030-0112102032303230-1330310100200323-3333331011023331-0133022032211113-0300021303213132-2322322213010223-1300333210200220"></a>

<a id="canonical-2231222302210201-0211313121010123-2323233111233010-3010113203133113-3233320322020033-1123001020010210-1300213010011303-1001112021222330"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0223333313103100-2321032022003022-0011331200332030-3313112223201333-3303121311302323-1310211323221101-2302132132130100-0311211220302213"></a>

<a id="canonical-0020003320311033-0300110031130020-2102313032001011-1312313313211232-2011211223233231-2331002100112001-3323232011310002-3131023231012112"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
