---
page_title: "xcsh_alert_gen_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy reference."
---

# xcsh_alert_gen_policy reference

<a id="canonical-0113323010121111-3310031311010033-0120221122230330-1032002302130100-3011003320313103-0332032231313220-3233111110000200-2012111221210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001)
- Property reference

<a id="canonical-0300020013023103-2301303322221202-0313323321333031-2322320221201110-1031113002031332-0021300211103130-1022010103123130-1110321102223022"></a>

### Direct properties for `xcsh_alert_gen_policy`

<a id="canonical-0002030222312112-2300332011102002-1121000313131021-2012132030313211-2222121210033011-2313323320300132-2132322200323312-1001032010033220"></a>

#### `alert_status` property

Type: `"string"`. Optional, Computed.

\[Enum: ALERT\_ACTIVE|ALERT\_INACTIVE\] Alert Status. List of alert statuses Active Inactive.
Possible values are \`ALERT\_ACTIVE\`, \`ALERT\_INACTIVE\`. Defaults to \`ALERT\_ACTIVE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALERT_ACTIVE","ALERT_INACTIVE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALERT_ACTIVE",
    "ALERT_INACTIVE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ALERT_ACTIVE",
  "enum": [
    "ALERT_ACTIVE",
    "ALERT_INACTIVE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210332111020021-3011333313133333-3330231303132023-0030031101232031-0202000021212320-1002211022310002-3111230023211012-1002321121333332"></a>

<a id="canonical-0230131032222113-0120000222200231-0210132133310103-3312020213230133-0033221300100121-1113023103020023-2331232231133330-1312231223310032"></a>

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

<a id="canonical-0113132311213223-1331320130120131-0221031003021102-0233201003132023-1001032200312220-3032210132033303-0301132231013112-1203120301323000"></a>

<a id="canonical-2323221311112213-0230222121230000-2023100133313231-1010300212323132-2211312213322203-3233331330202032-0110130231000100-1130331100110110"></a>

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

- [details](resources--alert_gen_policy--reference--group-001.md#canonical-0320213020013121-1010322233231212-2300332320130011-3121330303302010-0001233132322230-3323233211022310-3113330122121321-2101130202032320): complete subsection reference.

<a id="canonical-3130132201021003-1312223131211101-0333201120003112-2300321200321112-0332202302102110-3110033012230011-1032132333033222-3330103200013323"></a>

<a id="canonical-1122302120300323-3223211231103113-2222021323331122-3230022111133312-2310302201012021-3000000032320111-0100123122131200-2202221301223001"></a>

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

<a id="canonical-3020031101003301-3131331020122201-2121223022130212-3321313032101010-1111223311230312-3213112212130301-0111301031333113-3103121300123222"></a>

<a id="canonical-0120202111012030-1122102023013232-2223133333201130-2111213110313313-1113112311000031-2020212113110002-2311333200132102-0200220110233100"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0010101320112131-2203022023030002-0100201301222033-0221331030000122-3301123300330122-0222113201032331-0311310000110123-1213222332322303"></a>

<a id="canonical-0130122012022211-1300221212302231-0100101330321012-2211313322132302-2010120200223202-0312321231002213-0223321013133112-1220212202223202"></a>

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

<a id="canonical-2222013333112130-3122222032213102-3223213010131223-2201320233030220-0310120320022021-0102100130233122-2033232132101331-0120213230322032"></a>

<a id="canonical-1121112023010332-1033202003300011-3300311320220002-0332323000210220-2013033122312332-1201013023033322-2112023131311202-2210213300200303"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Alert Gen Policy. Must be unique within the namespace.

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

<a id="canonical-2023310303003302-2001302301003313-3121000303021213-0130311013322113-3223221331102320-3132021303013012-0001103133030102-2231131231203121"></a>

<a id="canonical-0233313120313231-2013000323130111-3001233222332103-2233101223113213-1132030013212223-3303123032320103-0232032203322212-3231211110000323"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Alert Gen Policy is created.

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

- [timeouts](resources--alert_gen_policy--reference--group-001.md#canonical-3222010321001100-3030000323213301-1100232023312222-3000012332332203-0302210320302130-2311330301223302-0000133310200101-1021231310310200): complete subsection reference.

<a id="canonical-3220022130210223-1200222032210210-1133303212133333-2003220330020313-2033031300122201-0200221133201200-0332032230202021-1131110322203111"></a>

### All schema paths for `xcsh_alert_gen_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_status` | [alert_status](resources--alert_gen_policy--reference--group-001.md#canonical-0002030222312112-2300332011102002-1121000313131021-2012132030313211-2222121210033011-2313323320300132-2132322200323312-1001032010033220) |
| `annotations` | [annotations](resources--alert_gen_policy--reference--group-001.md#canonical-3210332111020021-3011333313133333-3330231303132023-0030031101232031-0202000021212320-1002211022310002-3111230023211012-1002321121333332) |
| `description` | [description](resources--alert_gen_policy--reference--group-001.md#canonical-0113132311213223-1331320130120131-0221031003021102-0233201003132023-1001032200312220-3032210132033303-0301132231013112-1203120301323000) |
| `details` | [details](resources--alert_gen_policy--reference--group-001.md#canonical-2012233223312123-3103312003132310-0331113320110100-1232111000301212-0123332320030230-1002322210320130-3102302300133010-2111112300230230) |
| `details.alert_message` | [details.alert_message](resources--alert_gen_policy--reference--group-001.md#canonical-0100303213311213-1101110122223130-1100302130333000-0332100011033102-2002032112021010-1033310123210000-1312201031333031-1103003312303201) |
| `details.alert_message_details` | [details.alert_message_details](resources--alert_gen_policy--reference--group-001.md#canonical-1123303013222223-3111313332103022-1102032210230303-0333313203030101-3132221022312312-3330311101233222-3330131320130232-3122301131311130) |
| `details.alert_name` | [details.alert_name](resources--alert_gen_policy--reference--group-001.md#canonical-1113201120131230-3111102223031123-0233320000310122-2213201320111200-1212300223220003-1213311003111313-1201213212013023-1232003303103000) |
| `details.severity` | [details.severity](resources--alert_gen_policy--reference--group-001.md#canonical-0112032032323021-3021303102322103-1131320022320122-3320031323130033-1221100331131113-0322102333332121-1230202221000310-2213302322010031) |
| `disable` | [disable](resources--alert_gen_policy--reference--group-001.md#canonical-3130132201021003-1312223131211101-0333201120003112-2300321200321112-0332202302102110-3110033012230011-1032132333033222-3330103200013323) |
| `id` | [ID](resources--alert_gen_policy--reference--group-001.md#canonical-3020031101003301-3131331020122201-2121223022130212-3321313032101010-1111223311230312-3213112212130301-0111301031333113-3103121300123222) |
| `labels` | [labels](resources--alert_gen_policy--reference--group-001.md#canonical-0010101320112131-2203022023030002-0100201301222033-0221331030000122-3301123300330122-0222113201032331-0311310000110123-1213222332322303) |
| `name` | [name](resources--alert_gen_policy--reference--group-001.md#canonical-2222013333112130-3122222032213102-3223213010131223-2201320233030220-0310120320022021-0102100130233122-2033232132101331-0120213230322032) |
| `namespace` | [namespace](resources--alert_gen_policy--reference--group-001.md#canonical-2023310303003302-2001302301003313-3121000303021213-0130311013322113-3223221331102320-3132021303013012-0001103133030102-2231131231203121) |
| `timeouts` | [timeouts](resources--alert_gen_policy--reference--group-001.md#canonical-0000203111012130-2003121101032013-3323212000100132-2023330203110233-1333203011311111-3003002323023121-0012101120032210-2102230013311301) |
| `timeouts.create` | [timeouts.create](resources--alert_gen_policy--reference--group-001.md#canonical-2011231332110200-0103203030002030-2312022303033211-1230130221300011-0131232222113022-3230222012113323-1023323021220233-0003213132310213) |
| `timeouts.delete` | [timeouts.delete](resources--alert_gen_policy--reference--group-001.md#canonical-0122323002002132-1110012130130110-3020023302213011-0111100122102020-3123213301112003-2223312103321301-3103120011221111-3012131203130232) |
| `timeouts.read` | [timeouts.read](resources--alert_gen_policy--reference--group-001.md#canonical-0001213320023020-3320320122210331-0101202013231321-2110120111022332-3300032310132011-0102132110010131-3123031000020230-2113333000322213) |
| `timeouts.update` | [timeouts.update](resources--alert_gen_policy--reference--group-001.md#canonical-0031230202220122-1323300221131213-1130102033030130-0102233311032110-3131322332220220-3230313120311133-1130312333201133-3300002123203210) |

<a id="canonical-0320213020013121-1010322233231212-2300332320130011-3121330303302010-0001233132322230-3323233211022310-3113330122121321-2101130202032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `details` properties

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001)
- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-0113323010121111-3310031311010033-0120221122230330-1032002302130100-3011003320313103-0332032231313220-3233111110000200-2012111221210303)
- details

<a id="canonical-2012233223312123-3103312003132310-0331113320110100-1232111000301212-0123332320030230-1002322210320130-3102302300133010-2111112300230230"></a>

Type: `"object"`. single nested block, Optional.

Notification Details. Notification Details.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("alert_message",
    "alert_message_details",
    "alert_name")}
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
details {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030300200303121-2233031103313113-3203013120121003-1102010323332031-2212010130120020-2202310203331130-1012221132120002-0323331313320111"></a>

### Direct properties for `details`

<a id="canonical-0100303213311213-1101110122223130-1100302130333000-0332100011033102-2002032112021010-1033310123210000-1312201031333031-1103003312303201"></a>

#### `details.alert_message` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123303013222223-3111313332103022-1102032210230303-0333313203030101-3132221022312312-3330311101233222-3330131320130232-3122301131311130"></a>

<a id="canonical-2321030011311302-0022021203133031-0311121202111022-0112022332023103-0212113203323302-3011132021003201-0331221321301222-3033020022130123"></a>

#### `details.alert_message_details` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1113201120131230-3111102223031123-0233320000310122-2213201320111200-1212300223220003-1213311003111313-1201213212013023-1232003303103000"></a>

<a id="canonical-2012122300331323-2320122312022130-1113000232020221-3222303013023111-2331031330321020-3232303220213111-0033330312201221-0112311333011110"></a>

#### `details.alert_name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0112032032323021-3021303102322103-1131320022320122-3320031323130033-1221100331131113-0322102333332121-1230202221000310-2213302322010031"></a>

<a id="canonical-3222022210100132-0231310210312030-3123332101211000-3333120331002210-3132013211110303-1302321222203022-0113101312003233-0111010211103203"></a>

#### `details.severity` property

Type: `"string"`. Optional.

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

<a id="canonical-3222010321001100-3030000323213301-1100232023312222-3000012332332203-0302210320302130-2311330301223302-0000133310200101-1021231310310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001)
- [Property reference](resources--alert_gen_policy--reference--group-001.md#canonical-0113323010121111-3310031311010033-0120221122230330-1032002302130100-3011003320313103-0332032231313220-3233111110000200-2012111221210303)
- timeouts

<a id="canonical-0000203111012130-2003121101032013-3323212000100132-2023330203110233-1333203011311111-3003002323023121-0012101120032210-2102230013311301"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211122031311312-0122301021223132-1022113213320313-2122200322331310-3003003310013323-3101222110313000-1113031333313330-1102213110322211"></a>

### Direct properties for `timeouts`

<a id="canonical-2011231332110200-0103203030002030-2312022303033211-1230130221300011-0131232222113022-3230222012113323-1023323021220233-0003213132310213"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0122323002002132-1110012130130110-3020023302213011-0111100122102020-3123213301112003-2223312103321301-3103120011221111-3012131203130232"></a>

<a id="canonical-3203311102233323-3302001232020010-3132312002231013-3323121122203203-0030101023220100-0022113331123101-1121220032210021-1021112112333203"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0001213320023020-3320320122210331-0101202013231321-2110120111022332-3300032310132011-0102132110010131-3123031000020230-2113333000322213"></a>

<a id="canonical-1313010031110011-0213201211300321-3222001201331313-2121231120032113-2130202112111032-0220202322323202-0333123233000111-1131212112223130"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0031230202220122-1323300221131213-1130102033030130-0102233311032110-3131322332220220-3230313120311133-1130312333201133-3300002123203210"></a>

<a id="canonical-1301320020001010-1200122321112212-3311222013013113-1123210102330132-1320322331110130-1331302131222302-1230312110323301-0222113130331302"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
