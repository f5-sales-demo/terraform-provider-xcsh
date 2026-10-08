---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0223132131303030-1021321002133031-1322031020210023-3021012231212331-3323212220200021-3333310133301130-1033301220110133-0232301203331323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- CloudFront.js_insertion_rules.exclude_list.domain

<a id="canonical-0110023101111330-2100003101201210-3220210301313112-3122012123320131-2203001331203003-0000122002021020-1212112020000021-0032011100222031"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321103001302133-2233102201310302-2213300223210120-1301102103321002-0012331213013301-0112102300201103-0021220311212012-0223212231110113"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.domain`

<a id="canonical-3333212031332210-1233001202012210-1000012210011312-0211203132313301-1030032121202301-0202332231300320-0210222211000213-2130312303020320"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1212312013033113-3312110132332300-0231320133000112-0012322322210203-3020120112230002-3112222233213232-2233232100203321-0331231103333122"></a>

<a id="canonical-1201220123121330-1301021201313320-2011220030331130-3330112312231323-3111101110310031-3303031331211031-0112121022303001-0131210200130201"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3010312020320321-3021213012322230-3031023033233131-2321330012130212-0011120322101223-1232122311121032-1312213211013311-1033232103200021"></a>

<a id="canonical-0221301003322033-2310011123133031-3303131101032111-3303223323200102-1213000131123113-3001123003100131-2130111311302123-1313310001322233"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2010012310003231-3020100210200133-2210333023120033-1323211001220022-1201221122322313-0012223022121231-2311120202101323-3322032232313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- CloudFront.js_insertion_rules.exclude_list.metadata

<a id="canonical-0331131003120202-3321122131300201-2120330321122301-3103100301232002-0000223312033102-3121300220220003-1112331321203033-1100312310231133"></a>

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

<a id="canonical-1110201111021230-0001310021221121-2130120233221113-1320010131031113-3111320213231321-3103013322020133-3100321131023312-1200221003110213"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.metadata`

<a id="canonical-2100000131003312-2022111030200000-3111123023030023-2000323110003021-3132300120130122-3032321112010011-0230111020001131-1333320233222210"></a>

#### `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2231111332221102-1022111101313032-1023111230022112-0233102111121013-3330311302003012-0212321203211303-3132222012130122-2203321011313212"></a>

<a id="canonical-3032203310301330-1230120323213211-3220001132100330-2011013000330111-3302323010321131-1303131321010302-0011321200100223-3333003313032032"></a>

#### `cloudfront.js_insertion_rules.exclude_list.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2021213021001211-1331122120210300-3310333130003330-2200033112003033-1101020130331102-3033121122222031-1213332313232000-1103013002303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- CloudFront.js_insertion_rules.exclude_list.path

<a id="canonical-1212333221130231-2003301110201313-2232333231213322-1023001213302123-1301011201211301-2220213111211321-3222130333003200-3221103010032323"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130020213101230-0321202100002330-0002113203110010-2213220213132132-0333333033122230-3230013213001200-1212023233100201-3311013111220322"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.path`

<a id="canonical-1332230000021302-3310310103030120-1023000022321332-3310201023011011-2102301103111110-0033112122000113-3113220300033031-0322111222120020"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0312021110333212-0133023012130221-3313130013330000-0223010212101331-3111002101320321-2230203003231321-3322011120020212-0203202302032220"></a>

<a id="canonical-0203200230102301-3203230313200320-3033033001031333-1021000002332123-3320300302221333-1010310023311101-3233012013233013-0111210122020313"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1330210002002220-3200033311123313-1313330130330212-2020313310103020-1011330330212113-3120023031110102-1102112211311022-0233010132000313"></a>

<a id="canonical-1311313301133111-2012023312230102-3123232321031333-1332231023301323-3303120102333110-2021222102033223-0200330222210301-1021113230012213"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- CloudFront.js_insertion_rules.rules

<a id="canonical-2232301211223321-3221213223300310-2233200033002102-0010210130210031-1102112312112332-1322001133222310-3302100211330323-2001211333132331"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("exact_path",
    "glob"),
  validators.ConflictingListObjectAttributes("exact_path",
    "prefix"),
  validators.ConflictingListObjectAttributes("glob",
    "prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332220013312213-2233231223211231-1030020213300122-2013201311023122-0120223130323113-3122220120220300-0023003202222220-3033012122201302"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules`

- [any_domain](resources--protected_application--reference--group-003.md#canonical-3113132303113221-1313322211002210-1201130223003310-2203023013231322-2212023120011313-1030032202011113-3201110011121201-2030233302323303): complete subsection reference.

- [domain](resources--protected_application--reference--group-003.md#canonical-0023212320322203-0213121232200112-2220021122213223-3223312330031332-2202130332023231-1020232012200112-3022013203103133-1303001233321310): complete subsection reference.

<a id="canonical-0323311013030222-3112110323330323-0233023332333221-2331300003222313-3013231300321133-2211110231220200-3103011023320312-1000212103101020"></a>

<a id="canonical-1031000332000132-0011030021021113-1011230112021022-2310210112301313-1332013103331220-1001223031031020-0000031002002303-2032113002113200"></a>

#### `cloudfront.js_insertion_rules.rules.exact_path` property

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2023020032320232-1121333112011120-0031322303211220-1001303331022202-1133300331031122-2330202221130231-2123101233211112-1012213100011300"></a>

<a id="canonical-3330131120302223-0221013101212300-1012120133023031-3311103111123221-2211310132030123-1310120232213120-1303333320113013-2200102120101322"></a>

#### `cloudfront.js_insertion_rules.rules.glob` property

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](resources--protected_application--reference--group-003.md#canonical-1223022223311200-0231213311321001-3003302023112203-3232220111011133-3011123103101123-0333000231303333-3023031110022220-2000201223130122): complete subsection reference.

<a id="canonical-3223031203233032-2003012201202212-2232133100013233-3202123130232202-2330313013202000-1013211002131002-2302122321032232-0311121222322110"></a>

<a id="canonical-1001010131301222-1132200112120101-1113233011322333-3111122120130211-1012223020312233-0132301113011202-0133211301001023-1323333111330330"></a>

#### `cloudfront.js_insertion_rules.rules.prefix` property

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3113132303113221-1313322211002210-1201130223003310-2203023013231322-2212023120011313-1030032202011113-3201110011121201-2030233302323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-003.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- CloudFront.js_insertion_rules.rules.any_domain

<a id="canonical-2132311030110110-1123210001112230-1120313320130332-3013011220000113-3030012003101131-1121000303311021-2113331331203333-3322333030110332"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023212320322203-0213121232200112-2220021122213223-3223312330031332-2202130332023231-1020232012200112-3022013203103133-1303001233321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-003.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- CloudFront.js_insertion_rules.rules.domain

<a id="canonical-0013100022002003-0302333331023221-2001211133322011-3033012111230113-0122131231331000-2302133312331300-1210000301330103-0112000111012303"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232120221130201-2200211330131212-1303120203212332-1000022000021101-3010031303220310-1002201330212012-3032033201331102-1332302203132031"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules.domain`

<a id="canonical-0313103222213221-2320320302100212-1203013322231112-1203012021312131-1020203233213233-3002120033100003-2130303013103211-1310320113202121"></a>

#### `cloudfront.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1200233221202121-0322022332012233-1123223020103322-3223110021030232-1113113123120113-2112030201033010-3310211110332131-3133120321230221"></a>

<a id="canonical-3200120210133203-0020033003232021-3101032001333032-1333122332110000-1111011201012222-1020110200210320-2223103210222223-3021013003203202"></a>

#### `cloudfront.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3210320010130001-0220111310231323-3000133310023222-0133103323012132-1122323030322200-1030012132302010-1020122231303030-2323210032031102"></a>

<a id="canonical-3033012322020302-0020210230003300-2310202001122020-2000322313122322-1230013022123102-0102001302201332-2111301300033231-1103120101122110"></a>

#### `cloudfront.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1223022223311200-0231213311321001-3003302023112203-3232220111011133-3011123103101123-0333000231303333-3023031110022220-2000201223130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-003.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- CloudFront.js_insertion_rules.rules.metadata

<a id="canonical-1031010132131112-1022320021232131-3131213322032030-1122331330333012-2122113332003112-2220333131220202-0022122112001111-3303133012331302"></a>

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

<a id="canonical-3301101000223210-0020112331320221-0130300032302213-0313333133101000-0332311210112011-3003323031010020-2311313212101113-0111203200223100"></a>

### Direct properties for `cloudfront.js_insertion_rules.rules.metadata`

<a id="canonical-1012322111320232-0210033131222102-3113200102233330-3221112010312322-1123233302313332-1200001101020331-2231022303312001-2101130001202201"></a>

#### `cloudfront.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1101222332210323-2303302130210230-0113003103232131-0330223020030202-3200221223103311-3303020310000332-3212333020211031-0221311202310323"></a>

<a id="canonical-0330030032123311-0113001112321121-2101302133120203-3201200132101021-2100000010030023-3223232020002321-0133130320013221-1003202132231020"></a>

#### `cloudfront.js_insertion_rules.rules.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3313102332000123-2222313311313211-0013120212332103-3201031211032333-2332132333320123-1122123222002123-2323011023213133-1102033113123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.manual_js_insert` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.manual_js_insert

<a id="canonical-3111311000322030-0022100000021020-0232123211332032-3221310303123120-2102331112132202-3131330031332023-2332203032012203-2102001302122230"></a>

Type: `"object"`. single nested block, Optional.

Insert JavaScript Manually. Insert JavaScript manually.

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
manual_js_insert {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313331133001112-1011022012212202-2100033131322233-1103002200001203-3033223333002011-3322221012300132-2313322233100032-3321320203202102"></a>

### Direct properties for `cloudfront.manual_js_insert`

<a id="canonical-0030001220222311-1221333222033303-3133111222220222-1111301122022302-3023123312000231-2301133132020010-0211211111333111-0123003022312231"></a>

#### `cloudfront.manual_js_insert.javascript_mode` property

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASYNC_JS_CACHING","ASYNC_JS_NO_CACHING","SYNC_JS_CACHING","SYNC_JS_NO_CACHING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011003301222102-1310233120312210-0121013002313230-0023220030200033-1022123020103032-0023121000313121-3021222010133130-3302112032133333"></a>

<a id="canonical-2310210320021113-1233321201101331-1201310013323031-3022032331221202-3033330220302221-1221330102113112-1333313002203122-2100100312202022"></a>

#### `cloudfront.manual_js_insert.js_download_path` property

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

<a id="canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.mobile_sdk_config

<a id="canonical-0003102230233322-1123130302302312-1310011103113111-1100130222313321-3032223220113130-1200231202311210-1023311301023022-2330013100022023"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201320313002102-3002321332311302-3303110231001030-3310031210312123-2200302330230223-1033231203131320-1002303113302300-1102211323301021"></a>

### Direct properties for `cloudfront.mobile_sdk_config`

- [mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123): complete subsection reference.

<a id="canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- CloudFront.mobile_sdk_config.mobile_identifier

<a id="canonical-1023231021012323-3021220201203210-3212312302100032-1202130310111132-0203300131111000-3300021023130102-0313033101021011-0211110232200321"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112013331013030-2110001332312132-3120111130011300-0213111321022232-1020133121331223-1021302202212000-3222022213102000-0001220012323011"></a>

### Direct properties for `cloudfront.mobile_sdk_config.mobile_identifier`

- [headers](resources--protected_application--reference--group-003.md#canonical-3222330311330230-3311302303032020-0033203020300230-0131033100232230-0222232233322013-0200101033201122-2230110201020130-3012230303003032): complete subsection reference.

<a id="canonical-3222330311330230-3311302303032020-0033203020300230-0131033100232230-0222232233322013-0200101033201122-2230110201020130-3012230303003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123)
- CloudFront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-0010301331001010-3302032031132133-1213010000220323-2312312212233301-2330302330313100-0110133110010001-0223222323333030-3211032012112333"></a>

Type: `"object"`. list nested block, Optional.

A list of headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000233231112220-2123013130302220-1232130022210311-0020120131323232-2331010302313232-1221002312232303-1320032133022210-0112220300222320"></a>

### Direct properties for `cloudfront.mobile_sdk_config.mobile_identifier.headers`

<a id="canonical-1110312221023320-1233012033220103-2021320222013312-3131222112300002-3122013313301022-2120321223222332-3132223023200200-0100113132031102"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1133001011303302-1320233333013220-3130223023133000-2333030031131101-3032233310123311-1232122311001202-0300030223323311-2013021221111312"></a>

<a id="canonical-2112032333021310-0022023310033210-2213230210100122-0231122011203002-3120112011032113-1100122233300131-2202313322230233-3222320113102201"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0202020213311113-0123123032120011-0131032202020323-3310113301300312-1311102312222020-2022031101231223-2230122112101212-3201012030230230"></a>

<a id="canonical-3111102023120013-1022200301310201-1003323003232202-0122313331101301-0001333110332200-2031020030103310-1232033123033032-2102023310110001"></a>

#### `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.protected_endpoints

<a id="canonical-0130322112222030-3010310233213230-0222002100231013-0112310133121111-0103222031032213-0333230030103221-3213332030100332-0030303213021213"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods",
    "path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_client"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_mobile_client"),
  validators.ConflictingListObjectAttributes("web_client",
    "web_mobile_client")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000202330122301-2303130233201303-3030321110333022-3221032101201313-0310121112311022-2210131300302221-3333112310121333-1300220320112310"></a>

### Direct properties for `cloudfront.protected_endpoints`

- [any_domain](resources--protected_application--reference--group-003.md#canonical-1231101122233203-0201223233310200-1021322101112322-1300122131023113-1213022023011002-1322233103132220-3212101123210011-2013103111301322): complete subsection reference.

- [domain](resources--protected_application--reference--group-003.md#canonical-2021302022322203-1201213130101213-0121203302122111-1133202213130300-1121111330212131-0212232321003031-3320131323133301-0323031310100222): complete subsection reference.

- [flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321): complete subsection reference.

<a id="canonical-3312011312103121-1203001110201303-1013212002130311-0102102303021330-1213003211000202-1230230002103232-2122213311313311-0130022312320232"></a>

<a id="canonical-2213000111103110-3232220022022202-1300213200121213-0201110103030321-1112331322130330-1211233023333122-3000120331302032-1002112300122211"></a>

#### `cloudfront.protected_endpoints.http_methods` property

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--protected_application--reference--group-004.md#canonical-0212111302332332-2200231122211333-2121203211303103-0331301100223111-3032301232231102-3220220133320233-3132313100200232-0031102133110112): complete subsection reference.

- [mobile_client](resources--protected_application--reference--group-004.md#canonical-0213232003133201-2213121300112120-1230023220211230-1331323331121131-3322331313113311-3231221012003322-0013220102232300-1333301123112210): complete subsection reference.

<a id="canonical-0212211313212221-2231220220232323-2301013100022303-1232220021231120-3120213233130301-1310120231311131-1320311121012003-3203211131233131"></a>

<a id="canonical-1130010131321202-1012012322332301-1103123311211231-1020033213120212-2233320103030133-2002113113313130-0203031230102011-3112000020323031"></a>

#### `cloudfront.protected_endpoints.path` property

Type: `"string"`. Optional.

Accepts wildcards \* to match multiple characters or ? To match a single character.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="canonical-1131322110103100-2310231121012331-3202011010121210-1020311122113120-1232031322202123-0120000323301212-0121032121111003-2031001022330102"></a>

<a id="canonical-0302112030030232-1201222330111022-2112230333312311-2122003000332321-3033311222232311-1213321133012202-1233102302113001-3230211000231013"></a>

#### `cloudfront.protected_endpoints.query` property

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-0131023103211321-1320023232011020-2100333323130231-0001322233320320-3331120133222213-2223200322102221-1312101233013012-2013320230011222): complete subsection reference.

- [web_client](resources--protected_application--reference--group-004.md#canonical-1020130212202303-3003102010033232-1022131031112222-0333010110210231-0301020212010331-3300211020023001-3213011302113301-2310103111001033): complete subsection reference.

- [web_mobile_client](resources--protected_application--reference--group-004.md#canonical-2301202303022003-3310033230133332-2300012130112012-1211013023313222-3323320221311301-3221310030232212-0133121013213301-3111320012220112): complete subsection reference.

<a id="canonical-1231101122233203-0201223233310200-1021322101112322-1300122131023113-1213022023011002-1322233103132220-3212101123210011-2013103111301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- CloudFront.protected_endpoints.any_domain

<a id="canonical-0102331320201233-3232020032100022-0321033121230120-0323031320233133-3122103330110110-2233111033112222-2113302030111310-3010101103210113"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021302022322203-1201213130101213-0121203302122111-1133202213130300-1121111330212131-0212232321003031-3320131323133301-0323031310100222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- CloudFront.protected_endpoints.domain

<a id="canonical-1313110312021023-0123030231003101-1313310231010102-0032233321311031-1320300301233022-2311030020313023-1333303301302310-3010010332011201"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100102110020311-2301222333012023-2201220300103300-3111010230332320-3101303300310310-3323221032013203-1330112320130201-0010321100102102"></a>

### Direct properties for `cloudfront.protected_endpoints.domain`

<a id="canonical-3120122233021112-1131311312011113-3201013020212100-2000310000110132-3132033010000031-0032302122122331-0103333323020323-3001212013012103"></a>

#### `cloudfront.protected_endpoints.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1203301210030200-0102012033303010-2112213203030121-3110302223131003-2202203011110300-3223031322302200-3331231333202330-0022302132223023"></a>

<a id="canonical-2223102222331110-0101332032223013-1301300231133132-3233203302222013-0321200000033210-2310210110121032-0320100031030012-3120233301232213"></a>

#### `cloudfront.protected_endpoints.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3203233103200221-0323013331312231-0130021311300333-3301203002322332-3312022001231332-3012000232330302-2211103120313301-1220121300323122"></a>

<a id="canonical-1322320032233012-1111322312331122-1112133000032132-2221120021210110-3303311231001320-0022333321021201-0200213131111331-2322131300211000"></a>

#### `cloudfront.protected_endpoints.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- CloudFront.protected_endpoints.flow_label

<a id="canonical-2003220302020001-2223030300311220-0103320100333031-0001113022333230-1000133323103202-2303321023221102-2323311330302202-2333120220201021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130001333100020-1030022032010312-0100130021312302-2001203320202103-0131320331213223-2301020013101310-2110203010321122-1310220220233100"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label`

- [account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102): complete subsection reference.

- [authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110): complete subsection reference.

- [financial_services](resources--protected_application--reference--group-004.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012): complete subsection reference.

- [flight](resources--protected_application--reference--group-004.md#canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020): complete subsection reference.

- [profile_management](resources--protected_application--reference--group-004.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013): complete subsection reference.

- [search](resources--protected_application--reference--group-004.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033): complete subsection reference.

- [shopping_gift_cards](resources--protected_application--reference--group-004.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223): complete subsection reference.

<a id="canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- CloudFront.protected_endpoints.flow_label.account_management

<a id="canonical-2222223332333010-2101003220111022-3331301103100110-0100323032022002-1211131202302023-1011322320032110-0300232020112203-3201231331233113"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220102231332112-3223211223211302-3320033121303320-1110102012101010-3323101313320013-3132102132110213-3202131233131121-3212120112010202"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.account_management`

- [create](resources--protected_application--reference--group-003.md#canonical-0010011231211302-0003112002121330-0022012020011220-1103322001233112-3321301302223320-2210012101311322-1101033320202013-1221012322210231): complete subsection reference.

- [password_reset](resources--protected_application--reference--group-003.md#canonical-3311323233200230-1233101111300330-3200133202001002-3011100013000000-1113130113133202-3231023301112032-3112030331002223-0211002222323002): complete subsection reference.

<a id="canonical-0010011231211302-0003112002121330-0022012020011220-1103322001233112-3321301302223320-2210012101311322-1101033320202013-1221012322210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- CloudFront.protected_endpoints.flow_label.account_management.create

<a id="canonical-1123023232301222-1213312321112321-0321330131330212-2312032020211233-0212123021011113-0021312031032021-1213010300202102-0021033312221312"></a>

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311323233200230-1233101111300330-3200133202001002-3011100013000000-1113130113133202-3231023301112032-3112030331002223-0211002222323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- CloudFront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-3113321021003202-0010000333020131-1213112333220030-0131213131112103-2012323220231000-1010231010220220-3223202131113111-0200300100201122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- CloudFront.protected_endpoints.flow_label.authentication

<a id="canonical-2120011210000210-1312030131323033-1323203212003010-0111030211221013-0200303123202213-2011010032323302-0010223122220331-3020031320323210"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322030233303130-2003231320001011-2322112320030100-3203001300230130-3310301113131020-2101333311202101-0211001101022012-3023322120211112"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication`

- [login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123): complete subsection reference.

- [login_mfa](resources--protected_application--reference--group-003.md#canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033): complete subsection reference.

- [login_partner](resources--protected_application--reference--group-003.md#canonical-3021331301322233-0202121213222120-1202301321321011-0001020133303301-3231212320132102-1233320130021010-0102332203222333-1331232032201022): complete subsection reference.

- [logout](resources--protected_application--reference--group-004.md#canonical-2023230121133030-2201121101121230-0220332300301121-2302020013003231-2020332313030301-3013310311001322-2300002322121000-0203111111331023): complete subsection reference.

- [token_refresh](resources--protected_application--reference--group-004.md#canonical-0221220001332301-1120022323032010-3223222333000222-3122210013003033-1211103333223233-2023023223100310-1203113132202212-1301101303333202): complete subsection reference.

<a id="canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- CloudFront.protected_endpoints.flow_label.authentication.login

<a id="canonical-2100220030133031-2222020120132102-2303203130021202-0211202303130230-0023131020330202-2331013021222200-2313113322231331-1321200303120202"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302202321022212-1103300202201232-1313202020312200-0001222031102302-0031002000000322-2021122221132022-1233323032000231-0310321033330033"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login`

- [disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-1011330022313201-2010222013023213-2121231222200200-2103302111012200-3303301221330112-3300222132111000-1113310001123222-2321032232133312): complete subsection reference.

- [transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112): complete subsection reference.

<a id="canonical-1011330022313201-2010222013023213-2121231222200200-2103302111012200-3303301221330112-3300222132111000-1113310001123222-2321032232133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- CloudFront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0023303301013131-1120311233013031-2201231101300033-0033123210213203-2103022323110122-3232002223103223-2003331010003112-1233303012210200"></a>

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
disable_transaction_result = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-1312321011003321-0331102232310132-0213001002210213-3222001021211133-0101312320033223-2113212211103033-3202012121333121-0331211222013331"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

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
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010012020131212-2030213101301012-0212330310022233-2021312102023220-0110320002233313-1331000020100003-3300112313023222-0111112010002113"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](resources--protected_application--reference--group-003.md#canonical-3020000011113220-1111003231013330-2033011320012111-2001121321120203-2200221023320332-1101303020220113-3000000222300121-0111230130032120): complete subsection reference.

- [success_conditions](resources--protected_application--reference--group-003.md#canonical-1232102212131233-1330303022023022-0020022022100122-0323203102120032-2013212001003312-0311003300313120-3022300233013111-3310333001321002): complete subsection reference.

<a id="canonical-3020000011113220-1111003231013330-2033011320012111-2001121321120203-2200221023320332-1101303020220113-3000000222300121-0111230130032120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-2322020233033113-1232313311222022-2301130220121332-1320023013012212-1321010130320103-2232322123211213-3332021202003010-0321130300032302"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003133033220102-0100131132111013-3211030102132323-1331031333330223-1030333222113311-0112210331133313-2213331110300002-3030311100223222"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-2003111202101211-3323130303000020-0321001231332101-3032103303220312-3110310213011012-1120131233021313-0100021323100300-1223222000200033"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

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
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1011020021211122-2023323010222303-1112320303200130-2103202313010222-3033133331200032-0100130013110331-2132020020122130-2233100320012200"></a>

<a id="canonical-0313303101101030-0200032101103003-2121130033132203-1332213223132302-3021232111111310-1201303232100333-2221200201221302-2103032212231332"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1133313212032100-3113020223310312-1200221030112031-3213303310310221-2322000000330033-1131223013203131-0102313001231300-3101122121011202"></a>

<a id="canonical-2120223213322001-3332210220301011-2021333201320111-3000300220003003-0111010112030011-3233212312231222-2231111231303100-2122333322101132"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1232102212131233-1330303022023022-0020022022100122-0323203102120032-2013212001003312-0311003300313120-3022300233013111-3310333001321002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- CloudFront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-2303233312312033-3332223123122300-0321130001300301-3222023120031103-0130030330110120-3030332333100003-3101023013330233-0313012231320010"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213101123021130-3023332023131302-3320013320210301-1312232121230113-3032102021021213-3330101213220013-0221213311323220-3232020003232213"></a>

### Direct properties for `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-3122223223030113-1301132113013330-3032301010310201-3310020213232120-0203231232223121-1003131130002011-2102000313112302-1212320321122332"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

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
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2222222000010333-1012102201231123-1210022332033022-0130113033213011-3321100120132333-2313120231310032-0102320332330131-3111323222121133"></a>

<a id="canonical-3320211312011002-1310121112020313-1220000221330023-0123332221213320-3030221203113130-0321011211222210-3232211101322220-1100203010002223"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0323030010303113-2122111223300123-3301213302121013-3002123213120323-3102302033332022-3312212001013332-2003201022102113-0200221310323031"></a>

<a id="canonical-0112322132023131-1221301012333230-2010323010301303-0332311031222212-0230123223001110-2020013322221002-1210210203220310-1030123201121330"></a>

#### `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- CloudFront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3003300202212103-1231230001310010-1300203213002131-1023312020311022-3101031300203030-1301333320232322-1333210212302311-3011231321020000"></a>

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
login_mfa = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021331301322233-0202121213222120-1202301321321011-0001020133303301-3231212320132102-1233320130021010-0102332203222333-1331232032201022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.protected_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- CloudFront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-1132303132132333-3012301232210213-3232320112121110-1202320312223012-3230300032233001-2021211011230333-1130123211131332-1322122122333033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

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
login_partner = {}
```

This is an empty object or choice marker. It has no direct properties.
