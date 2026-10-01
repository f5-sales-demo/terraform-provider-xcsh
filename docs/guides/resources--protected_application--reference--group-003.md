---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0130220212323001-3203032213333001-1231023302322010-2003120130330022-2233233030112331-0130100300221300-2032003001321333-2131321222000331"></a>

## name property — metadata / 223202103302 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1112212031011221-1022331030210001-2320212230301310-2231200223220230-3121313102133230-0230023121300113-3302001330232221-0333000211012321"></a>

## Next pages — metadata / 223202103302 / 6

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3313102332000123-2222313311313211-0013120212332103-3201031211032333-2332132333320123-1122123222002123-2323011023213133-1102033113123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313331133001112-1011022012212202-2100033131322233-1103002200001203-3033223333002011-3322221012300132-2313322233100032-3321320203202102"></a>

## cloudfront.manual_js_insert — manual_js_insert / 310322302201 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.manual_js_insert

<a id="canonical-3111311000322030-0022100000021020-0232123211332032-3221310303123120-2102331112132202-3131330031332023-2332203032012203-2102001302122230"></a>

Type: `"object"`. single nested block, Optional.

Insert JavaScript Manually. Insert JavaScript manually.

Upstream description:

Insert JavaScript manually.

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

<a id="canonical-2310210320021113-1233321201101331-1201310013323031-3022032331221202-3033330220302221-1221330102113112-1333313002203122-2100100312202022"></a>

## Direct properties — manual_js_insert / 310322302201 / 3

<a id="canonical-0030001220222311-1221333222033303-3133111222220222-1111301122022302-3023123312000231-2301133132020010-0211211111333111-0123003022312231"></a>

<a id="canonical-1231301211302200-2001003111031011-2221102113023130-0230221210322000-3130030032032330-1231322130222320-1231102100200021-2303201112311013"></a>

## javascript_mode property — manual_js_insert / 310322302201 / 4

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3111111310200211-1032320032223200-0002301013233313-3333311133113320-0210202032201320-2302311202121022-1313110203221202-2312130211011313"></a>

## js_download_path property — manual_js_insert / 310322302201 / 5

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1123130013212003-2303323130310220-2002201113312000-0210030332313010-1321320322011002-1013003232300311-1033323132122033-3310200110231201"></a>

## Next pages — manual_js_insert / 310322302201 / 6

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201320313002102-3002321332311302-3303110231001030-3310031210312123-2200302330230223-1033231203131320-1002303113302300-1102211323301021"></a>

## cloudfront.mobile_sdk_config — mobile_sdk_config / 132023122120 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.mobile_sdk_config

<a id="canonical-0003102230233322-1123130302302312-1310011103113111-1100130222313321-3032223220113130-1200231202311210-1023311301023022-2330013100022023"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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

<a id="canonical-1010213112111111-2323000332311321-2000111013113003-0332333121020202-0213003202213131-1103030303013020-0121233230111110-2233100112122231"></a>

## Direct properties — mobile_sdk_config / 132023122120 / 3

- [mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123): complete subsection reference.

<a id="canonical-1300132213232211-1223101032013112-3210203111012013-0000211300222302-3222120211111302-3230122103221010-3132111332111011-1010110211320302"></a>

## Next pages — mobile_sdk_config / 132023122120 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112013331013030-2110001332312132-3120111130011300-0213111321022232-1020133121331223-1021302202212000-3222022213102000-0001220012323011"></a>

## cloudfront.mobile_sdk_config.mobile_identifier — mobile_identifier / 010031112001 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- cloudfront.mobile_sdk_config.mobile_identifier

<a id="canonical-1023231021012323-3021220201203210-3212312302100032-1202130310111132-0203300131111000-3300021023130102-0313033101021011-0211110232200321"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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

<a id="canonical-3311031002323203-0022312010200233-2000322013332212-0313210320011020-0130130002322002-3232313221202100-0312110212213033-0031001110232230"></a>

## Direct properties — mobile_identifier / 010031112001 / 3

- [headers](resources--protected_application--reference--group-003.md#canonical-3222330311330230-3311302303032020-0033203020300230-0131033100232230-0222232233322013-0200101033201122-2230110201020130-3012230303003032): complete subsection reference.

<a id="canonical-3323210100320031-3110201102303211-3321133313032231-0011033313331100-1300332222321002-2130121213021210-0210123221201202-3003002313110001"></a>

## Next pages — mobile_identifier / 010031112001 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-003.md#canonical-3222330311330230-3311302303032020-0033203020300230-0131033100232230-0222232233322013-0200101033201122-2230110201020130-3012230303003032)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3222330311330230-3311302303032020-0033203020300230-0131033100232230-0222232233322013-0200101033201122-2230110201020130-3012230303003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000233231112220-2123013130302220-1232130022210311-0020120131323232-2331010302313232-1221002312232303-1320032133022210-0112220300222320"></a>

## cloudfront.mobile_sdk_config.mobile_identifier.headers — headers / 321100321021 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123)
- cloudfront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-0010301331001010-3302032031132133-1213010000220323-2312312212233301-2330302330313100-0110133110010001-0223222323333030-3211032012112333"></a>

Type: `"object"`. list nested block, Optional.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2112032333021310-0022023310033210-2213230210100122-0231122011203002-3120112011032113-1100122233300131-2202313322230233-3222320113102201"></a>

## Direct properties — headers / 321100321021 / 3

<a id="canonical-1110312221023320-1233012033220103-2021320222013312-3131222112300002-3122013313301022-2120321223222332-3132223023200200-0100113132031102"></a>

<a id="canonical-3111102023120013-1022200301310201-1003323003232202-0122313331101301-0001333110332200-2031020030103310-1232033123033032-2102023310110001"></a>

## exact property — headers / 321100321021 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0013121111112311-2231031023132132-1031303331103323-3131333120221332-2301133010222032-0130331233001302-0303322313213310-3022312113132220"></a>

## name property — headers / 321100321021 / 5

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0011122010320221-2321012023231231-3233220023331313-0101123322100303-3032012313101133-0311021303032311-0123220110131102-1231012300302320"></a>

## regex property — headers / 321100321021 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3123020223330031-3130023332012031-2301012132123323-2213010112121003-3201331123312312-2000300121200210-0031201332121200-3001313020231000"></a>

## Next pages — headers / 321100321021 / 7

- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-003.md#canonical-0213223320221112-2122101302230222-1313301132220012-2212300132233123-3031030123331203-3003133332121301-1031000123123233-1033210032003123)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000202330122301-2303130233201303-3030321110333022-3221032101201313-0310121112311022-2210131300302221-3333112310121333-1300220320112310"></a>

## cloudfront.protected_endpoints — protected_endpoints / 302032320020 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.protected_endpoints

<a id="canonical-0130322112222030-3010310233213230-0222002100231013-0112310133121111-0103222031032213-0333230030103221-3213332030100332-0030303213021213"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2213000111103110-3232220022022202-1300213200121213-0201110103030321-1112331322130330-1211233023333122-3000120331302032-1002112300122211"></a>

## Direct properties — protected_endpoints / 302032320020 / 3

- [any_domain](resources--protected_application--reference--group-003.md#canonical-1231101122233203-0201223233310200-1021322101112322-1300122131023113-1213022023011002-1322233103132220-3212101123210011-2013103111301322): complete subsection reference.

- [domain](resources--protected_application--reference--group-003.md#canonical-2021302022322203-1201213130101213-0121203302122111-1133202213130300-1121111330212131-0212232321003031-3320131323133301-0323031310100222): complete subsection reference.

- [flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321): complete subsection reference.

<a id="canonical-3312011312103121-1203001110201303-1013212002130311-0102102303021330-1213003211000202-1230230002103232-2122213311313311-0130022312320232"></a>

<a id="canonical-1130010131321202-1012012322332301-1103123311211231-1020033213120212-2233320103030133-2002113113313130-0203031230102011-3112000020323031"></a>

## http_methods property — protected_endpoints / 302032320020 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0302112030030232-1201222330111022-2112230333312311-2122003000332321-3033311222232311-1213321133012202-1233102302113001-3230211000231013"></a>

## path property — protected_endpoints / 302032320020 / 5

Type: `"string"`. Optional.

Accepts wildcards \* to match multiple characters or ? To match a single character.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210220233303103-1022102002032320-1213131122123101-0300022121200120-3322022002120003-1120023111301033-3310323213313310-1101113013131330"></a>

## query property — protected_endpoints / 302032320020 / 6

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121203232200312-3320033121203330-3133032123111131-1203102123210332-1312031121211131-3322121033100113-0200000211320011-1320232203322321"></a>

## Next pages — protected_endpoints / 302032320020 / 7

- [cloudfront.protected_endpoints.any_domain](resources--protected_application--reference--group-003.md#canonical-1231101122233203-0201223233310200-1021322101112322-1300122131023113-1213022023011002-1322233103132220-3212101123210011-2013103111301322)
- [cloudfront.protected_endpoints.domain](resources--protected_application--reference--group-003.md#canonical-2021302022322203-1201213130101213-0121203302122111-1133202213130300-1121111330212131-0212232321003031-3320131323133301-0323031310100222)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.metadata](resources--protected_application--reference--group-004.md#canonical-0212111302332332-2200231122211333-2121203211303103-0331301100223111-3032301232231102-3220220133320233-3132313100200232-0031102133110112)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-004.md#canonical-0213232003133201-2213121300112120-1230023220211230-1331323331121131-3322331313113311-3231221012003322-0013220102232300-1333301123112210)
- [cloudfront.protected_endpoints.undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-0131023103211321-1320023232011020-2100333323130231-0001322233320320-3331120133222213-2223200322102221-1312101233013012-2013320230011222)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-1020130212202303-3003102010033232-1022131031112222-0333010110210231-0301020212010331-3300211020023001-3213011302113301-2310103111001033)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-2301202303022003-3310033230133332-2300012130112012-1211013023313222-3323320221311301-3221310030232212-0133121013213301-3111320012220112)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1231101122233203-0201223233310200-1021322101112322-1300122131023113-1213022023011002-1322233103132220-3212101123210011-2013103111301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210111112132313-2031003322201023-3123131212131131-0010131032003111-1122321301300013-0311011210311021-1133210011000312-3302231230311011"></a>

## cloudfront.protected_endpoints.any_domain — any_domain / 221220221013 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- cloudfront.protected_endpoints.any_domain

<a id="canonical-0102331320201233-3232020032100022-0321033121230120-0323031320233133-3122103330110110-2233111033112222-2113302030111310-3010101103210113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-0121223313013301-3003232001012312-3330223303033220-3212112103322101-1110131320333023-1233333021300002-0203113023210102-0211321020023302"></a>

## Direct properties — any_domain / 221220221013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023103232300011-2203113002203313-3000120013300202-0102130000303003-3020012132200333-1321213130010231-1230301012213222-0222233312303221"></a>

## Next pages — any_domain / 221220221013 / 4

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2021302022322203-1201213130101213-0121203302122111-1133202213130300-1121111330212131-0212232321003031-3320131323133301-0323031310100222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100102110020311-2301222333012023-2201220300103300-3111010230332320-3101303300310310-3323221032013203-1330112320130201-0010321100102102"></a>

## cloudfront.protected_endpoints.domain — domain / 203202313322 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- cloudfront.protected_endpoints.domain

<a id="canonical-1313110312021023-0123030231003101-1313310231010102-0032233321311031-1320300301233022-2311030020313023-1333303301302310-3010010332011201"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2223102222331110-0101332032223013-1301300231133132-3233203302222013-0321200000033210-2310210110121032-0320100031030012-3120233301232213"></a>

## Direct properties — domain / 203202313322 / 3

<a id="canonical-3120122233021112-1131311312011113-3201013020212100-2000310000110132-3132033010000031-0032302122122331-0103333323020323-3001212013012103"></a>

<a id="canonical-1322320032233012-1111322312331122-1112133000032132-2221120021210110-3303311231001320-0022333321021201-0200213131111331-2322131300211000"></a>

## exact_value property — domain / 203202313322 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3133022212111322-3212101112230103-3223333311212110-1112211303003122-0233010020213123-3131123321012113-3120213130121102-2131210133000011"></a>

## regex_value property — domain / 203202313322 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3133331222311332-3013100011213013-2123131013131132-2223111211233310-2231221022000221-1211212120023030-1333132120103212-0201013213101221"></a>

## suffix_value property — domain / 203202313322 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1311300303002010-1221232303103002-2022230121233321-3100032021210022-3010331002333330-3021222203121232-0312303000311211-3002133313212023"></a>

## Next pages — domain / 203202313322 / 7

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130001333100020-1030022032010312-0100130021312302-2001203320202103-0131320331213223-2301020013101310-2110203010321122-1310220220233100"></a>

## cloudfront.protected_endpoints.flow_label — flow_label / 110013000203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- cloudfront.protected_endpoints.flow_label

<a id="canonical-2003220302020001-2223030300311220-0103320100333031-0001113022333230-1000133323103202-2303321023221102-2323311330302202-2333120220201021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0331230311111312-2321030020101310-0020232103311200-3013310021312231-1122212211213232-1310002332313132-1300111313302121-1231131311030320"></a>

## Direct properties — flow_label / 110013000203 / 3

- [account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102): complete subsection reference.

- [authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110): complete subsection reference.

- [financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012): complete subsection reference.

- [flight](resources--protected_application--reference--group-003.md#canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020): complete subsection reference.

- [profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013): complete subsection reference.

- [search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033): complete subsection reference.

- [shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223): complete subsection reference.

<a id="canonical-2031013212330201-2101231301012221-0131133032333022-2013021222013002-0300030002312010-2112032321001203-2130032321110101-3101012311003223"></a>

## Next pages — flow_label / 110013000203 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012)
- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220102231332112-3223211223211302-3320033121303320-1110102012101010-3323101313320013-3132102132110213-3202131233131121-3212120112010202"></a>

## cloudfront.protected_endpoints.flow_label.account_management — account_management / 231230313201 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.account_management

<a id="canonical-2222223332333010-2101003220111022-3331301103100110-0100323032022002-1211131202302023-1011322320032110-0300232020112203-3201231331233113"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0333331320212300-2230121302130322-0333203201011021-0311333132232030-0302103213033301-2221202323213223-1121132231203321-2032203010102013"></a>

## Direct properties — account_management / 231230313201 / 3

- [create](resources--protected_application--reference--group-003.md#canonical-0010011231211302-0003112002121330-0022012020011220-1103322001233112-3321301302223320-2210012101311322-1101033320202013-1221012322210231): complete subsection reference.

- [password_reset](resources--protected_application--reference--group-003.md#canonical-3311323233200230-1233101111300330-3200133202001002-3011100013000000-1113130113133202-3231023301112032-3112030331002223-0211002222323002): complete subsection reference.

<a id="canonical-2331020020023010-0112301112012033-1222230112032012-3210220201312323-2301111111312301-1033001221021232-1011312131032003-3302210311132121"></a>

## Next pages — account_management / 231230313201 / 4

- [cloudfront.protected_endpoints.flow_label.account_management.create](resources--protected_application--reference--group-003.md#canonical-0010011231211302-0003112002121330-0022012020011220-1103322001233112-3321301302223320-2210012101311322-1101033320202013-1221012322210231)
- [cloudfront.protected_endpoints.flow_label.account_management.password_reset](resources--protected_application--reference--group-003.md#canonical-3311323233200230-1233101111300330-3200133202001002-3011100013000000-1113130113133202-3231023301112032-3112030331002223-0211002222323002)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0010011231211302-0003112002121330-0022012020011220-1103322001233112-3321301302223320-2210012101311322-1101033320202013-1221012322210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001113003232302-1203220231320012-3301112321023130-2202002310210121-2022123312302213-1221320300322123-1322100022002322-1200033300012210"></a>

## cloudfront.protected_endpoints.flow_label.account_management.create — create / 031110031030 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- cloudfront.protected_endpoints.flow_label.account_management.create

<a id="canonical-1123023232301222-1213312321112321-0321330131330212-2312032020211233-0212123021011113-0021312031032021-1213010300202102-0021033312221312"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1312223232230330-3211313210113321-0110103110030111-1123203320132021-3302023030311233-1202103323211222-3032332103312313-1121210311133101"></a>

## Direct properties — create / 031110031030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302010331022313-3301320233301103-1221031030133301-2023113212022011-1311113102230311-1223323310101220-0000031211331331-1010213222033323"></a>

## Next pages — create / 031110031030 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3311323233200230-1233101111300330-3200133202001002-3011100013000000-1113130113133202-3231023301112032-3112030331002223-0211002222323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130021322123000-3223101021301303-2312231221020321-2212120323121030-0023210003123210-3012030200000310-3220101011231112-0323223032132213"></a>

## cloudfront.protected_endpoints.flow_label.account_management.password_reset — password_reset / 333033131231 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- cloudfront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-3113321021003202-0010000333020131-1213112333220030-0131213131112103-2012323220231000-1010231010220220-3223202131113111-0200300100201122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

Upstream description:

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

<a id="canonical-1231201301112131-3200310321002112-3013100210333032-2012210023221223-1222010012302032-1230331200102312-1103223101332100-0323213121131032"></a>

## Direct properties — password_reset / 333033131231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122321111333012-1312223231230122-3311321203000020-3211200030323310-1033313032122200-2333003233232031-2323313131023210-1212113022023110"></a>

## Next pages — password_reset / 333033131231 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-3211101131023120-1123312220031210-2022320321213311-3111233311133331-3331303121322001-3100323103210023-2230022212122213-2233213001313102)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322030233303130-2003231320001011-2322112320030100-3203001300230130-3310301113131020-2101333311202101-0211001101022012-3023322120211112"></a>

## cloudfront.protected_endpoints.flow_label.authentication — authentication / 233331133020 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="canonical-2120011210000210-1312030131323033-1323203212003010-0111030211221013-0200303123202213-2011010032323302-0010223122220331-3020031320323210"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2011032023023221-1301230111010101-2120200323233212-3223011202330112-2300223122303033-2032322301312031-1222101311000011-0302030113233032"></a>

## Direct properties — authentication / 233331133020 / 3

- [login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123): complete subsection reference.

- [login_mfa](resources--protected_application--reference--group-003.md#canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033): complete subsection reference.

- [login_partner](resources--protected_application--reference--group-003.md#canonical-3021331301322233-0202121213222120-1202301321321011-0001020133303301-3231212320132102-1233320130021010-0102332203222333-1331232032201022): complete subsection reference.

- [logout](resources--protected_application--reference--group-003.md#canonical-2023230121133030-2201121101121230-0220332300301121-2302020013003231-2020332313030301-3013310311001322-2300002322121000-0203111111331023): complete subsection reference.

- [token_refresh](resources--protected_application--reference--group-003.md#canonical-0221220001332301-1120022323032010-3223222333000222-3122210013003033-1211103333223233-2023023223100310-1203113132202212-1301101303333202): complete subsection reference.

<a id="canonical-3323010220210211-3010231122233123-3323313111000111-3310201230210231-3231230203131013-3110023113021030-0103323130312312-2230020110202010"></a>

## Next pages — authentication / 233331133020 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](resources--protected_application--reference--group-003.md#canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](resources--protected_application--reference--group-003.md#canonical-3021331301322233-0202121213222120-1202301321321011-0001020133303301-3231212320132102-1233320130021010-0102332203222333-1331232032201022)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](resources--protected_application--reference--group-003.md#canonical-2023230121133030-2201121101121230-0220332300301121-2302020013003231-2020332313030301-3013310311001322-2300002322121000-0203111111331023)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](resources--protected_application--reference--group-003.md#canonical-0221220001332301-1120022323032010-3223222333000222-3122210013003033-1211103333223233-2023023223100310-1203113132202212-1301101303333202)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302202321022212-1103300202201232-1313202020312200-0001222031102302-0031002000000322-2021122221132022-1233323032000231-0310321033330033"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login — login / 312202032130 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="canonical-2100220030133031-2222020120132102-2303203130021202-0211202303130230-0023131020330202-2331013021222200-2313113322231331-1321200303120202"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0311021220220220-1200010132312213-3123003030213223-3032122230203000-0032013303112122-2303111201233013-0332300210222220-2103322103222132"></a>

## Direct properties — login / 312202032130 / 3

- [disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-1011330022313201-2010222013023213-2121231222200200-2103302111012200-3303301221330112-3300222132111000-1113310001123222-2321032232133312): complete subsection reference.

- [transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112): complete subsection reference.

<a id="canonical-0210003031121133-1312310103032203-2033022320331120-1023220313320211-0210102221210210-3132303121121012-3103230220323322-1032023002230212"></a>

## Next pages — login / 312202032130 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-1011330022313201-2010222013023213-2121231222200200-2103302111012200-3303301221330112-3300222132111000-1113310001123222-2321032232133312)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1011330022313201-2010222013023213-2121231222200200-2103302111012200-3303301221330112-3300222132111000-1113310001123222-2321032232133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301202031003210-3022330112302030-1102003001111123-1000201202010111-2230020200013031-2103012133233131-0321332222121101-2200132311111100"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result — disable_transaction_result / 301232123120 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0023303301013131-1120311233013031-2201231101300033-0033123210213203-2103022323110122-3232002223103223-2003331010003112-1233303012210200"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-0103120210333133-0132331003302301-0200302221102020-1211023312223231-2033313213033111-0321032023130120-3201102312031302-3003311120100122"></a>

## Direct properties — disable_transaction_result / 301232123120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312310133331030-0333100313003300-0321101312232323-2010033013003000-2230000200202301-2330300030202311-2132222201212323-0202131011212001"></a>

## Next pages — disable_transaction_result / 301232123120 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010012020131212-2030213101301012-0212330310022233-2021312102023220-0110320002233313-1331000020100003-3300112313023222-0111112010002113"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result — transaction_result / 223101002131 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-1312321011003321-0331102232310132-0213001002210213-3222001021211133-0101312320033223-2113212211103033-3202012121333121-0331211222013331"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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

<a id="canonical-2112011212130013-0303002220011212-2021123012223232-0022130222120301-3303313220103131-3233313000122303-3003221011021121-3220210221023313"></a>

## Direct properties — transaction_result / 223101002131 / 3

- [failure_conditions](resources--protected_application--reference--group-003.md#canonical-3020000011113220-1111003231013330-2033011320012111-2001121321120203-2200221023320332-1101303020220113-3000000222300121-0111230130032120): complete subsection reference.

- [success_conditions](resources--protected_application--reference--group-003.md#canonical-1232102212131233-1330303022023022-0020022022100122-0323203102120032-2013212001003312-0311003300313120-3022300233013111-3310333001321002): complete subsection reference.

<a id="canonical-3330003323110221-2010200333223121-3111033011232230-2021311332013323-0300003201020010-0020100002300101-2230323103031302-1332222311320330"></a>

## Next pages — transaction_result / 223101002131 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--protected_application--reference--group-003.md#canonical-3020000011113220-1111003231013330-2033011320012111-2001121321120203-2200221023320332-1101303020220113-3000000222300121-0111230130032120)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--protected_application--reference--group-003.md#canonical-1232102212131233-1330303022023022-0020022022100122-0323203102120032-2013212001003312-0311003300313120-3022300233013111-3310333001321002)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3020000011113220-1111003231013330-2033011320012111-2001121321120203-2200221023320332-1101303020220113-3000000222300121-0111230130032120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003133033220102-0100131132111013-3211030102132323-1331031333330223-1030333222113311-0112210331133313-2213331110300002-3030311100223222"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — failure_conditions / 123233012331 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-2322020233033113-1232313311222022-2301130220121332-1320023013012212-1321010130320103-2232322123211213-3332021202003010-0321130300032302"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0313303101101030-0200032101103003-2121130033132203-1332213223132302-3021232111111310-1201303232100333-2221200201221302-2103032212231332"></a>

## Direct properties — failure_conditions / 123233012331 / 3

<a id="canonical-2003111202101211-3323130303000020-0321001231332101-3032103303220312-3110310213011012-1120131233021313-0100021323100300-1223222000200033"></a>

<a id="canonical-2120223213322001-3332210220301011-2021333201320111-3000300220003003-0111010112030011-3233212312231222-2231111231303100-2122333322101132"></a>

## name property — failure_conditions / 123233012331 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311222002100332-1223102322322132-0012021003202210-3033010001021120-0232130230320300-2222011200123202-1111000303312021-1031302130313230"></a>

## regex_values property — failure_conditions / 123233012331 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0030021000130133-3213010003002330-2030323030030013-1302133233112100-3132113033210330-3231121002130000-1133213130311031-0232203021022031"></a>

## status property — failure_conditions / 123233012331 / 6

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

Upstream description:

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

<a id="canonical-2310021022203023-3231201321223202-1033323123003021-3130202013333212-0020003232233300-2230013132111102-0301322320310131-0120212232223311"></a>

## Next pages — failure_conditions / 123233012331 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1232102212131233-1330303022023022-0020022022100122-0323203102120032-2013212001003312-0311003300313120-3022300233013111-3310333001321002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213101123021130-3023332023131302-3320013320210301-1312232121230113-3032102021021213-3330101213220013-0221213311323220-3232020003232213"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions — success_conditions / 210321213311 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2110110120000110-1211111300013013-1012332313111333-0002322001030310-0100030321111311-2130220020103010-1013210131330023-0332131102302123)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-2303233312312033-3332223123122300-0321130001300301-3222023120031103-0130030330110120-3030332333100003-3101023013330233-0313012231320010"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320211312011002-1310121112020313-1220000221330023-0123332221213320-3030221203113130-0321011211222210-3232211101322220-1100203010002223"></a>

## Direct properties — success_conditions / 210321213311 / 3

<a id="canonical-3122223223030113-1301132113013330-3032301010310201-3310020213232120-0203231232223121-1003131130002011-2102000313112302-1212320321122332"></a>

<a id="canonical-0112322132023131-1221301012333230-2010323010301303-0332311031222212-0230123223001110-2020013322221002-1210210203220310-1030123201121330"></a>

## name property — success_conditions / 210321213311 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2111200302223322-3321131213123000-0031302303203330-3103121331232223-2100220031300211-2022333311001331-3033312301033022-2231100302323322"></a>

## regex_values property — success_conditions / 210321213311 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303031100002322-3201303310111212-3111310200312223-0011213110230231-0222201001201022-3330022322031021-3122221033102101-3022113303031333"></a>

## status property — success_conditions / 210321213311 / 6

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

Upstream description:

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

<a id="canonical-3331333121031313-1020123103031130-1130101111232213-2201123101302013-0330320301233212-3003220300332333-2222201123130113-0110233123333012"></a>

## Next pages — success_conditions / 210321213311 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-3022212120330211-0100001121312213-0122300301330011-3023122010321221-0120233110103322-2120030301020021-0232321230310013-0101011022111112)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012132230102110-0232301321133322-2020310101120111-0311222300210020-1222220211231001-3313312313032233-0132123123103301-0031111011223110"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_mfa — login_mfa / 031023320323 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- cloudfront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3003300202212103-1231230001310010-1300203213002131-1023312020311022-3101031300203030-1301333320232322-1333210212302311-3011231321020000"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1210121203001022-0020203110103020-2003101002212002-2232110301111330-2103000012203101-2013332103031033-0030102222223201-0000323233212301"></a>

## Direct properties — login_mfa / 031023320323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013320211200100-1220120333133201-2002320330220212-2033032101220323-1103320100320211-1230122012131113-2113203130013310-3131033132323201"></a>

## Next pages — login_mfa / 031023320323 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3021331301322233-0202121213222120-1202301321321011-0001020133303301-3231212320132102-1233320130021010-0102332203222333-1331232032201022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231122323221121-3310023210211310-2202201000210202-0020231202301201-2102133223321220-1201333001321002-0002012020213232-0331222030322032"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_partner — login_partner / 032112132331 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- cloudfront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-1132303132132333-3012301232210213-3232320112121110-1202320312223012-3230300032233001-2021211011230333-1130123211131332-1322122122333033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

Upstream description:

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

<a id="canonical-3331022232020311-1023200000220011-3121010112022110-2202222130103022-2213231122132021-3030000303320023-2110132001122023-3113130112030102"></a>

## Direct properties — login_partner / 032112132331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021123210221232-3322111003131130-1033031222200311-3323102110022330-2330130303313313-2010213211000011-3313213213111201-2231230232201311"></a>

## Next pages — login_partner / 032112132331 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2023230121133030-2201121101121230-0220332300301121-2302020013003231-2020332313030301-3013310311001322-2300002322121000-0203111111331023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021222113232313-2303102201121231-2230233320220103-2313220132331133-2301130232220022-1310302332100231-3231123123010212-2313221312220230"></a>

## cloudfront.protected_endpoints.flow_label.authentication.logout — logout / 300002212231 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- cloudfront.protected_endpoints.flow_label.authentication.logout

<a id="canonical-3210322333030011-0220230330311302-0311003201123012-2231031002312301-1022313021002022-0301133021100133-2032030300032320-1212200213323031"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
logout = {}
```

<a id="canonical-0112230010300232-1221033300132310-2232131012010333-1130210033033013-2221100330020300-3202221032321311-1112320001122023-0232000120001302"></a>

## Direct properties — logout / 300002212231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333323233223012-2132313300113222-1210213300210020-3110330231030323-1032103223332320-2220001313111033-0230230322003131-0011333320213213"></a>

## Next pages — logout / 300002212231 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0221220001332301-1120022323032010-3223222333000222-3122210013003033-1211103333223233-2023023223100310-1203113132202212-1301101303333202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303231203312320-2023302131202012-3111211022202312-1113332322201201-1022011313330032-3112202230300220-0230013322031133-0131221001233203"></a>

## cloudfront.protected_endpoints.flow_label.authentication.token_refresh — token_refresh / 220310103102 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- cloudfront.protected_endpoints.flow_label.authentication.token_refresh

<a id="canonical-0133020132230213-1032230031202221-0233133023130312-2112323011122320-1221133031211230-0121113031103022-2320110121300230-3332212320130031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

Upstream description:

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
token_refresh = {}
```

<a id="canonical-0211311330323032-0033003022113321-1230230210330202-3202300330213333-2323211112213331-0230331123233200-1130320233101332-3112323112012322"></a>

## Direct properties — token_refresh / 220310103102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321323120232311-1202301302102101-2212203322020302-1202131233001100-0312303201003112-1203332211010110-0311100211323123-2122300230330121"></a>

## Next pages — token_refresh / 220310103102 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-0323111212011331-2233302330122121-1312022311201112-3020213013212312-2311332012311021-1220300222331032-0121102320033210-2310022210020110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102202121223032-3202013033330002-3100222110300031-2333202032123302-1213103032033201-3300231030330001-1222200333122131-2032000111223131"></a>

## cloudfront.protected_endpoints.flow_label.financial_services — financial_services / 200102222023 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.financial_services

<a id="canonical-3311023122210020-0012103130131013-2112303010103103-0223101301311113-0230213211331112-3322023232310120-0131230302031121-2033011023231133"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
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
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330033332323222-3302221223031101-2203033210323113-3303230311132023-3301301230333112-2002300202132222-2211023011113301-0212231320130330"></a>

## Direct properties — financial_services / 200102222023 / 3

- [apply](resources--protected_application--reference--group-003.md#canonical-3233130101131012-2230231323120303-1112132001020310-0331221320303132-0133030121333222-2103013232210132-2000323131320330-0130212132311231): complete subsection reference.

- [money_transfer](resources--protected_application--reference--group-003.md#canonical-2123133020032313-0001102203232031-2301333103103301-2233222202131001-0301301130233211-1232100322300112-3330100302320022-3311203301301333): complete subsection reference.

<a id="canonical-1211303311321211-2223313123132313-0220023122230310-2131010213032332-1113023032310101-0031100022210000-2123201031233312-1302021132032110"></a>

## Next pages — financial_services / 200102222023 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services.apply](resources--protected_application--reference--group-003.md#canonical-3233130101131012-2230231323120303-1112132001020310-0331221320303132-0133030121333222-2103013232210132-2000323131320330-0130212132311231)
- [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](resources--protected_application--reference--group-003.md#canonical-2123133020032313-0001102203232031-2301333103103301-2233222202131001-0301301130233211-1232100322300112-3330100302320022-3311203301301333)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3233130101131012-2230231323120303-1112132001020310-0331221320303132-0133030121333222-2103013232210132-2000323131320330-0130212132311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121212111001001-1010301112233300-0123031113133332-3302232003301312-3103013320000023-1000112332100212-3111032030022221-3200231310103233"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.apply — apply / 021330121203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012)
- cloudfront.protected_endpoints.flow_label.financial_services.apply

<a id="canonical-1110203300103111-3200322301300013-3101120200310203-3103111322303212-3310333133301011-0333302103303311-2032120232031001-3030211000303332"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
apply = {}
```

<a id="canonical-0300100313212031-2303221320203331-2311320200232032-3210321120101031-1331033130030232-2003030333023221-1121321313000122-3321111113231200"></a>

## Direct properties — apply / 021330121203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221033201132233-0021111303211330-1011301103021103-2131212033200002-1300133130012123-0301132202111220-2133220312222033-1110131003233122"></a>

## Next pages — apply / 021330121203 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2123133020032313-0001102203232031-2301333103103301-2233222202131001-0301301130233211-1232100322300112-3330100302320022-3311203301301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111100133030111-1323232220122113-0301112201202020-3302002200233213-1223022220311121-1133130211031131-3130330110020311-3130201303021212"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.money_transfer — money_transfer / 230220330101 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012)
- cloudfront.protected_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0012330332312330-2122100113301031-2222200313101032-2200323331210200-2032303022212221-1211033301033222-1233212210301313-1102223030121200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

Upstream description:

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
money_transfer = {}
```

<a id="canonical-2001022013321230-2130230133112222-2130303002300231-2101031001110213-0300220121112102-1302312310112021-1121003130112101-1010121132032003"></a>

## Direct properties — money_transfer / 230220330101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323103111033321-0101131003333203-2133230311230022-1023132010333221-2331332321130003-3122213230020213-2022111012102230-3101332231201331"></a>

## Next pages — money_transfer / 230220330101 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-1221023222000133-2202131211101230-3002001310200213-2332103310233321-1101212233310220-2110032311022233-2010231232203303-3311033032212012)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203021211231333-0300122022022020-2322213031122222-0123020223120231-3120021333033312-0121233303233112-1112231212120330-3330100331212303"></a>

## cloudfront.protected_endpoints.flow_label.flight — flight / 220100320313 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.flight

<a id="canonical-3301321131023313-0302330113013030-3310232110030222-3030313232312020-1320032232323023-1332203032111130-3030023102213320-1202033201333023"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010023112021323-3031022210101002-3133122311201223-0100222213302222-0131032102033201-0122323310210031-1022212030120203-0003011302331132"></a>

## Direct properties — flight / 220100320313 / 3

- [checkin](resources--protected_application--reference--group-003.md#canonical-0113120030020110-0011033131131333-2010121201103222-3212110233112131-0130010203121022-0020303320232311-2123221132313032-1030102031331213): complete subsection reference.

<a id="canonical-1020133211303233-3130101003013311-0320120330020200-0331213300121222-2202011312311011-2011133130303001-0231102003210012-0003321221333101"></a>

## Next pages — flight / 220100320313 / 4

- [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--reference--group-003.md#canonical-0113120030020110-0011033131131333-2010121201103222-3212110233112131-0130010203121022-0020303320232311-2123221132313032-1030102031331213)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0113120030020110-0011033131131333-2010121201103222-3212110233112131-0130010203121022-0020303320232311-2123221132313032-1030102031331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130130211221321-0231133033012232-3101001133202302-0030031121231332-0133120302221030-1312023021131121-2102331100032201-2022032021333313"></a>

## cloudfront.protected_endpoints.flow_label.flight.checkin — checkin / 320300113103 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020)
- cloudfront.protected_endpoints.flow_label.flight.checkin

<a id="canonical-2121032003132110-1132231101023120-1112231211031222-1111032101213012-1113331010201312-2123103221033222-2313130313220023-0303132332000302"></a>

Type: `"object"`. single nested block, Optional.

Enable this option

Upstream description:

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
checkin {}
```

<a id="canonical-1233112301200003-1101100133022001-0132102031231203-3320012222111213-3213313133323302-2220031203330222-1123112320003123-2121122010320000"></a>

## Direct properties — checkin / 320300113103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022322010301011-2313000122211313-0210211203321301-3212122122220102-0111203323232132-3213100020332330-2203312120211013-2323031033003121"></a>

## Next pages — checkin / 320300113103 / 4

- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-1222012302233310-3002202302031210-3333311020003022-3111110222100121-0022220203120113-3212033313021202-1110103012131232-2221322022111020)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223120102111203-1121230123221112-1103000332303312-3013310333222223-2330010033303220-0002211010221103-2003211113310211-0112313113213112"></a>

## cloudfront.protected_endpoints.flow_label.profile_management — profile_management / 022323032201 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="canonical-3201201231303312-0232023323010330-2030123010331000-3111113012323111-0230302312122221-1112332120211321-2032232201311203-3011220030110030"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301320211303310-3111213112213222-3301211101201132-3313023211311101-1023012102321212-0230100133223011-2202231233033323-2301021232110222"></a>

## Direct properties — profile_management / 022323032201 / 3

- [create](resources--protected_application--reference--group-003.md#canonical-2100323120303101-1231303211310202-3012003221011133-3323101101331301-3021110232320012-1023000312330310-1320302133312101-3212032133103230): complete subsection reference.

- [update](resources--protected_application--reference--group-003.md#canonical-0201123110301210-0210120100110231-1033012112300313-3333023100222310-1113200023020321-0102020023323231-2300003033200121-0032232222102023): complete subsection reference.

- [view](resources--protected_application--reference--group-003.md#canonical-1233222100133011-0002022330323113-3101330101022203-0102003111221232-3312132032130331-0011030012202200-2331030301002313-1332002132123022): complete subsection reference.

<a id="canonical-2000310202303320-2312023101221311-0232303031011030-0301013220102123-2203301303031333-3101332012110223-0211011022103132-1010002121202113"></a>

## Next pages — profile_management / 022323032201 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--reference--group-003.md#canonical-2100323120303101-1231303211310202-3012003221011133-3323101101331301-3021110232320012-1023000312330310-1320302133312101-3212032133103230)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--reference--group-003.md#canonical-0201123110301210-0210120100110231-1033012112300313-3333023100222310-1113200023020321-0102020023323231-2300003033200121-0032232222102023)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--reference--group-003.md#canonical-1233222100133011-0002022330323113-3101330101022203-0102003111221232-3312132032130331-0011030012202200-2331030301002313-1332002132123022)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2100323120303101-1231303211310202-3012003221011133-3323101101331301-3021110232320012-1023000312330310-1320302133312101-3212032133103230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300323121331203-3103233300103310-1203031301332300-1112131202232330-3100100132311330-3112321031201322-2032000300230222-1131011132223113"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.create — create / 113122102130 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- cloudfront.protected_endpoints.flow_label.profile_management.create

<a id="canonical-0002003313032132-1031130003310220-1121310123223313-3120202003033212-3121300202100133-2201200333311232-1012210312200303-3131212133121220"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-3102301031102331-1002201300020312-1010032131030312-2123001323301232-2210200200002230-0030121221023201-2101231001302201-0121222130322130"></a>

## Direct properties — create / 113122102130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330300012110220-3202222100313223-1013030220112333-2000203132233130-0032311301103333-2023121223211301-2122131230110101-3203123233210010"></a>

## Next pages — create / 113122102130 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0201123110301210-0210120100110231-1033012112300313-3333023100222310-1113200023020321-0102020023323231-2300003033200121-0032232222102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032123333333333-3230221120330202-0030300030111110-0230013323032320-1011203201122211-1211332032320222-2121121313110223-0101310200320323"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.update — update / 320321213232 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- cloudfront.protected_endpoints.flow_label.profile_management.update

<a id="canonical-1222122222100100-0122031012302131-1110020310222101-2232022232023022-3330233030011322-1033220003003331-0132310101222101-0020211102210130"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
update = {}
```

<a id="canonical-2131031313211011-0313130332312113-1333212101322113-1132323313023030-1330220230231101-0111200131200231-2313030200123031-1201002223132302"></a>

## Direct properties — update / 320321213232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021221202201202-0222000011331222-0131003003210132-3010132111230123-0012230103301100-2000303323003310-0310303113201321-1001011303201311"></a>

## Next pages — update / 320321213232 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1233222100133011-0002022330323113-3101330101022203-0102003111221232-3312132032130331-0011030012202200-2331030301002313-1332002132123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323333121211100-0232322111111201-0322332210123031-1133231122220013-1233333322021023-0000022023021022-2202201123102131-2311223022132111"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.view — view / 203200130111 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- cloudfront.protected_endpoints.flow_label.profile_management.view

<a id="canonical-1112203213212331-2131201122320122-0212301200031202-0101321021221031-2330010313321102-0302110220131030-1031320220332223-1311121322213133"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
view = {}
```

<a id="canonical-3112101102010022-2122210222201202-2312321020012010-2310310211201121-1023312032320021-3322110101203113-2133313123330021-0101102013033031"></a>

## Direct properties — view / 203200130111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220023302230222-1112122103131122-1032120303020030-2312300310101122-3010123010131020-0320322133233200-3321211032120031-0030210322303331"></a>

## Next pages — view / 203200130111 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001312223130333-2321110220131223-1332302001011233-3323211220212120-3320302020210111-1032303213322313-0021130131330131-1213012231120123"></a>

## cloudfront.protected_endpoints.flow_label.search — search / 103313230102 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.search

<a id="canonical-1033122000013003-1200113000130310-0310300330202030-3203011102111231-3201320111230302-2010023012131133-1100320212203323-2122322022311023"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("flight_search",
    "product_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "room_search"),
  validators.ConflictingObjectAttributes("product_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("product_search",
    "room_search"),
  validators.ConflictingObjectAttributes("reservation_search",
    "room_search")}
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
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301311132330203-1011031000020221-1132131330120213-0133123220301331-1121210033103031-3212320331302130-0101001103202101-0112201132010010"></a>

## Direct properties — search / 103313230102 / 3

- [flight_search](resources--protected_application--reference--group-003.md#canonical-3223030031332010-2210120101102300-1213021232212133-3020320231330311-2013302201101102-3211303223113233-0231210011130331-2332230230212211): complete subsection reference.

- [product_search](resources--protected_application--reference--group-003.md#canonical-0201223233323331-2332201113330022-2111002002123021-3332302222333212-1100300032130013-0213220111223331-2231233122312012-1021232002011020): complete subsection reference.

- [reservation_search](resources--protected_application--reference--group-003.md#canonical-3113211010123310-1120033320002210-1130303233131111-3122300330031100-0200023320320233-3302012021130212-1000332233211013-2222123233012202): complete subsection reference.

- [room_search](resources--protected_application--reference--group-003.md#canonical-2133130311222322-2122123012212012-2333030011312022-3333111013103030-1021202123313110-3113101001121103-0001310322011023-2303030103021202): complete subsection reference.

<a id="canonical-1213033331010223-1111013011120332-3302132110321021-1031312001200213-2030101012320130-3100221202310221-3222200230310333-1003333111012210"></a>

## Next pages — search / 103313230102 / 4

- [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--reference--group-003.md#canonical-3223030031332010-2210120101102300-1213021232212133-3020320231330311-2013302201101102-3211303223113233-0231210011130331-2332230230212211)
- [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--reference--group-003.md#canonical-0201223233323331-2332201113330022-2111002002123021-3332302222333212-1100300032130013-0213220111223331-2231233122312012-1021232002011020)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--reference--group-003.md#canonical-3113211010123310-1120033320002210-1130303233131111-3122300330031100-0200023320320233-3302012021130212-1000332233211013-2222123233012202)
- [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--reference--group-003.md#canonical-2133130311222322-2122123012212012-2333030011312022-3333111013103030-1021202123313110-3113101001121103-0001310322011023-2303030103021202)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3223030031332010-2210120101102300-1213021232212133-3020320231330311-2013302201101102-3211303223113233-0231210011130331-2332230230212211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122201231002333-1301033030111013-2030100212313310-0033321001233223-0323022003211221-3221201001213100-0010333331233231-2331320220030013"></a>

## cloudfront.protected_endpoints.flow_label.search.flight_search — flight_search / 013131103131 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- cloudfront.protected_endpoints.flow_label.search.flight_search

<a id="canonical-2032120321100120-0223320231030200-0103022021031132-0322211323312123-3110010111030123-2033133101223323-1323213101222033-2021222203330102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for flight search.

Upstream description:

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
flight_search = {}
```

<a id="canonical-1110131303312033-2331020110303013-2022230011230032-2011310202231121-3131123001232333-2133311300210023-1130020320113323-2320021330322110"></a>

## Direct properties — flight_search / 013131103131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232230312222303-1101302233121323-2203331113200112-1222022101331032-2121013310223003-3331133230212103-0133220023120011-3221120121323323"></a>

## Next pages — flight_search / 013131103131 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0201223233323331-2332201113330022-2111002002123021-3332302222333212-1100300032130013-0213220111223331-2231233122312012-1021232002011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321031201311033-0211101211210203-0123033131113032-2322213211021230-3132013201302213-3212323001012001-3002333332132130-2030130232301223"></a>

## cloudfront.protected_endpoints.flow_label.search.product_search — product_search / 203230320120 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- cloudfront.protected_endpoints.flow_label.search.product_search

<a id="canonical-3121202110033003-2010031021211002-1303322113120130-2120233010133213-2021020301301201-1013022002212021-3111213030133102-1233120032300121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for product search.

Upstream description:

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
product_search = {}
```

<a id="canonical-3003003120121323-2222003230322222-3212231120111233-0302122322103330-2311110120122133-2313120012210213-1100020223133223-1312333232121003"></a>

## Direct properties — product_search / 203230320120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210313220001331-3031103013222330-1011232133012110-3310102110302022-2313211020100302-2132323321130100-0133231222003020-3020001123310230"></a>

## Next pages — product_search / 203230320120 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3113211010123310-1120033320002210-1130303233131111-3122300330031100-0200023320320233-3302012021130212-1000332233211013-2222123233012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231332310212011-1123020121030210-3121233301013212-3211223110122002-0033331302201303-0333201220131203-3131131331201223-1131220313013033"></a>

## cloudfront.protected_endpoints.flow_label.search.reservation_search — reservation_search / 123330312131 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- cloudfront.protected_endpoints.flow_label.search.reservation_search

<a id="canonical-1320033130302113-0311213122120002-1300101111033202-1103102012013333-0123201300233201-3313112021331231-1200120001210012-2202020113212020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reservation search.

Upstream description:

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
reservation_search = {}
```

<a id="canonical-1212130012301103-1102110222322201-0330032302030210-2132013122000121-2303132333001330-3001312001310132-3323230023233013-1323000331013023"></a>

## Direct properties — reservation_search / 123330312131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232311110333312-2112013010033030-0221101112313103-1001233210322133-0011213320032320-1333333312022111-2320101203202223-3120213203210002"></a>

## Next pages — reservation_search / 123330312131 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2133130311222322-2122123012212012-2333030011312022-3333111013103030-1021202123313110-3113101001121103-0001310322011023-2303030103021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233001012031020-3021311000131312-0233322331032003-0123122312212310-3003221131133320-1330110201110212-0232012323322200-3111202311032323"></a>

## cloudfront.protected_endpoints.flow_label.search.room_search — room_search / 001102300312 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- cloudfront.protected_endpoints.flow_label.search.room_search

<a id="canonical-2322003113001213-3333313133323000-0300212122032113-1330232313311313-1331030131231222-1021200111102121-1310221130333013-0113222013330020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for room search.

Upstream description:

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
room_search = {}
```

<a id="canonical-2331131012233323-1103133023203112-2002021022331122-2210123203202012-1010103030212112-3332132033121121-1213202331303311-1213313210133212"></a>

## Direct properties — room_search / 001102300312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233110023213131-1332010313300303-3122231230321332-0121130101113221-3011302122301021-3033233002300110-2230000100200101-2111213123123300"></a>

## Next pages — room_search / 001102300312 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1221120003231031-3310301213120220-0032220323011233-2110122200010201-3313222021201123-3020102102131103-3123330110003123-2121330223202033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110231120103122-2221221020221231-2013130010210021-0303132113132121-0330231220321001-1100310111100231-3320310323233102-2002313020132122"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards — shopping_gift_cards / 211102103202 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards

<a id="canonical-1332302023322230-2010133223021213-1222003233112000-0203312001221213-3121011231323300-3231331201211230-0331321121300010-3232120003221202"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "gift_card_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_purchase_gift_card",
    "shop_update_quantity")}
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
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021210023330220-1303110131130312-2111321311311032-2112210002010203-0130312020203231-3200211100033203-2210332322331022-0310023010303010"></a>

## Direct properties — shopping_gift_cards / 211102103202 / 3

- [gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-3120300330320222-2200121031020023-2031111213330001-0032220013011210-1112300310330300-3220131101020312-0211231022201210-3010212102012332): complete subsection reference.

- [gift_card_validation](resources--protected_application--reference--group-003.md#canonical-2133331233201132-2300303213300021-3222222320210030-0313012211302123-0302330312201101-1233211113011123-2133223300130311-1122132232330330): complete subsection reference.

- [shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-0120013113211021-1203331311220032-0202222202011130-0022203100103021-1223001222301011-1322132001233033-2331023001022110-3020132201323311): complete subsection reference.

- [shop_checkout](resources--protected_application--reference--group-003.md#canonical-1123100023311133-1133333032010301-3322221222020333-0333103200011030-0032022332123333-1202310321312313-0222300021120110-0311023231132111): complete subsection reference.

- [shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-1212333323023112-1101101213223222-3003022211323120-0322033321102010-3012003012320023-1331310301333210-3132320002011333-0100333011123113): complete subsection reference.

- [shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-1122201233020223-0123123102301002-2032003312331213-2000032030330122-1322100021333303-3313030000100110-3020220333031310-2323203301123323): complete subsection reference.

- [shop_make_payment](resources--protected_application--reference--group-003.md#canonical-1012102320222231-0230000030130332-0220311102022112-2203013110323333-2000320030132232-3233013121232111-2120131001302000-2131323333311300): complete subsection reference.

- [shop_order](resources--protected_application--reference--group-003.md#canonical-0011213111221013-2100013223312023-3020012303323230-2220033102001002-1331103022210111-1221001020032200-1233110301023232-3113330322021100): complete subsection reference.

- [shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-1110300023112033-1013101033311200-1311322021330322-2012200323001223-2300000312031120-0310022301323221-3311223311300022-1012311111122002): complete subsection reference.

- [shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-1130100110130001-1312332013300100-0211030011002200-2202102212232001-0131102320020033-2010311230213022-3102113133113010-1023212100303111): complete subsection reference.

- [shop_purchase_gift_card](resources--protected_application--reference--group-003.md#canonical-0023110333210212-1012020102233112-1311311023032233-2020213001333003-2030012213031301-3203002203030133-1001311010222330-0302110102302201): complete subsection reference.

- [shop_update_quantity](resources--protected_application--reference--group-004.md#canonical-2312112312301300-0313211201121023-2222332212221230-2122020003310213-2122033101132113-2020320021120032-2201312120300031-3011303213113303): complete subsection reference.

<a id="canonical-0112101210132301-3301120322133231-0201010112203330-2230333021031330-0133100002310022-0331312323312122-0103032323332000-3132110110103302"></a>

## Next pages — shopping_gift_cards / 211102103202 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-3120300330320222-2200121031020023-2031111213330001-0032220013011210-1112300310330300-3220131101020312-0211231022201210-3010212102012332)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--protected_application--reference--group-003.md#canonical-2133331233201132-2300303213300021-3222222320210030-0313012211302123-0302330312201101-1233211113011123-2133223300130311-1122132232330330)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-0120013113211021-1203331311220032-0202222202011130-0022203100103021-1223001222301011-1322132001233033-2331023001022110-3020132201323311)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--protected_application--reference--group-003.md#canonical-1123100023311133-1133333032010301-3322221222020333-0333103200011030-0032022332123333-1202310321312313-0222300021120110-0311023231132111)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-1212333323023112-1101101213223222-3003022211323120-0322033321102010-3012003012320023-1331310301333210-3132320002011333-0100333011123113)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-1122201233020223-0123123102301002-2032003312331213-2000032030330122-1322100021333303-3313030000100110-3020220333031310-2323203301123323)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--protected_application--reference--group-003.md#canonical-1012102320222231-0230000030130332-0220311102022112-2203013110323333-2000320030132232-3233013121232111-2120131001302000-2131323333311300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](resources--protected_application--reference--group-003.md#canonical-0011213111221013-2100013223312023-3020012303323230-2220033102001002-1331103022210111-1221001020032200-1233110301023232-3113330322021100)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-1110300023112033-1013101033311200-1311322021330322-2012200323001223-2300000312031120-0310022301323221-3311223311300022-1012311111122002)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-1130100110130001-1312332013300100-0211030011002200-2202102212232001-0131102320020033-2010311230213022-3102113133113010-1023212100303111)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--protected_application--reference--group-003.md#canonical-0023110333210212-1012020102233112-1311311023032233-2020213001333003-2030012213031301-3203002203030133-1001311010222330-0302110102302201)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--protected_application--reference--group-004.md#canonical-2312112312301300-0313211201121023-2222332212221230-2122020003310213-2122033101132113-2020320021120032-2201312120300031-3011303213113303)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3120300330320222-2200121031020023-2031111213330001-0032220013011210-1112300310330300-3220131101020312-0211231022201210-3010212102012332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123231102230313-2213020322121320-0102033022022303-1220203231031131-2230300200200332-2233032312302212-0113122010332321-3211331001030312"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — gift_card_make_purchase_with_gift_card / 013320212023 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-2221032332113013-2121320230101312-2332002130322213-0310233211233210-0210301233012033-0032202120112201-1010123223222231-0112301020221233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card make purchase with gift card.

Upstream description:

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
gift_card_make_purchase_with_gift_card = {}
```

<a id="canonical-3232101032103133-0111323110332211-1211210121033132-1300303113110322-3001200213013032-3010320233233301-1102122312110021-2110313023120111"></a>

## Direct properties — gift_card_make_purchase_with_gift_card / 013320212023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331331023331103-2021213020012030-2201102210201101-2111122101333312-2203230032022121-1302130003220011-2010123131211113-2233003210012012"></a>

## Next pages — gift_card_make_purchase_with_gift_card / 013320212023 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2133331233201132-2300303213300021-3222222320210030-0313012211302123-0302330312201101-1233211113011123-2133223300130311-1122132232330330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110131331020003-2203202132310123-1321301023313132-2232301330331100-1022112032132331-0133310220101221-3033213131313233-2310032323223112"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation — gift_card_validation / 321231130310 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-0220132213210103-3023321222211031-0121322331311112-3111132203010113-2313021122132323-3310131020121203-1131212301323331-3003221333313023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card validation.

Upstream description:

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
gift_card_validation = {}
```

<a id="canonical-1030111211223211-0001230301110231-2013123110321122-3303113221032132-0102323301313311-3033030201133003-1003333120231320-2222302112302113"></a>

## Direct properties — gift_card_validation / 321231130310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012133333221002-1221210221102312-3103110300103113-1122211130001233-3030011120331213-0121021321303132-1102011210022202-2221303233132203"></a>

## Next pages — gift_card_validation / 321231130310 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0120013113211021-1203331311220032-0202222202011130-0022203100103021-1223001222301011-1322132001233033-2331023001022110-3020132201323311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121301011021301-0312300132023223-1301230002120110-3103111332122133-0300131011233321-0121222013032212-3223330023023020-1311201001203201"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — shop_add_to_cart / 321020233230 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1012031230320210-2230212003000030-0113100220102223-0130121203310203-3100112320033323-2303023221112223-1021332020213103-1212111030022331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop add to cart.

Upstream description:

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
shop_add_to_cart = {}
```

<a id="canonical-2133302100213130-0021022200331100-0320123022201303-0322113011202223-1331220021122033-1013032120333110-0323030220201100-1011031101030120"></a>

## Direct properties — shop_add_to_cart / 321020233230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001113000103011-2323000100102300-3031030310323310-2310113122223300-1023211221233332-2213301200231310-2301203220101231-3201311310121101"></a>

## Next pages — shop_add_to_cart / 321020233230 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1123100023311133-1133333032010301-3322221222020333-0333103200011030-0032022332123333-1202310321312313-0222300021120110-0311023231132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131203202233032-2032100223023021-3300301313112221-0302032302301012-1113133133302323-1221303000113211-0103233103121022-2032011300030333"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout — shop_checkout / 211301013321 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-2202132013210321-0102003121332010-2311211103012313-2313121101223222-2322223131031333-0031120201300233-2310312201201032-3031203113320011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop checkout.

Upstream description:

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
shop_checkout = {}
```

<a id="canonical-0223133133210101-1130012302111233-2022133223313332-2210202332202303-2300013110230210-2003021201331312-2113002210211200-2022330112303311"></a>

## Direct properties — shop_checkout / 211301013321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233233010203021-3021203001111331-3112203321330302-0003111022330212-1202113123212302-1111223311220110-2211020000101231-0212013233310020"></a>

## Next pages — shop_checkout / 211301013321 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1212333323023112-1101101213223222-3003022211323120-0322033321102010-3012003012320023-1331310301333210-3132320002011333-0100333011123113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102013020310220-1232023121223330-0233031132020000-2213330011132301-1101310323313231-3333330121323331-1322200302112112-0321100233322131"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — shop_choose_seat / 331013030230 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-3103311323333301-0102113113121023-3200330012333230-1000200113110233-1123320001133123-0323113323333012-3220001320113212-1310103311332033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop choose seat.

Upstream description:

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
shop_choose_seat = {}
```

<a id="canonical-3300333212311011-2110223302113101-1300023220312011-2021332201111112-2231100221130311-0121330203110123-3303122330023331-3232100311303010"></a>

## Direct properties — shop_choose_seat / 331013030230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123020321331233-3120012203312231-0230230232300303-1003201123012012-2223120233103313-3010200122130220-2233102121321010-0333021023030023"></a>

## Next pages — shop_choose_seat / 331013030230 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1122201233020223-0123123102301002-2032003312331213-2000032030330122-1322100021333303-3313030000100110-3020220333031310-2323203301123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312310002111332-3120130003232212-1220233000113230-3120122330322122-1230111330002023-2023023030211231-2102331122232222-0213113301012303"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — shop_enter_drawing_submission / 231212210300 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-0101100103001210-3023212213331231-2201322332002003-2110030333321010-1112303001102320-0220200330211302-1201033113131320-3300231032030112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop enter drawing submission.

Upstream description:

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
shop_enter_drawing_submission = {}
```

<a id="canonical-0002323222223332-2010300002220330-2221311221222222-1200030333300131-0332332323312201-3002222030030101-3300113203302331-0031331003000121"></a>

## Direct properties — shop_enter_drawing_submission / 231212210300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320300223300131-0012230220233211-3123132100300123-0120012031213101-2112222303203311-0001302133022121-2033201312322302-2022003231130301"></a>

## Next pages — shop_enter_drawing_submission / 231212210300 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1012102320222231-0230000030130332-0220311102022112-2203013110323333-2000320030132232-3233013121232111-2120131001302000-2131323333311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300033030012212-1002122200333133-2110203011320113-2032213220001210-3202333300020112-0121232310021011-1310023310132220-3222113231220122"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment — shop_make_payment / 332023330031 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-0023010221111200-1330212031130113-1223223000213223-2301302033222120-3030103103230103-1221022133110233-1333221102130002-2011020102221123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop make payment.

Upstream description:

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
shop_make_payment = {}
```

<a id="canonical-2233223012222121-2310322312313033-3311301003123010-2010233000001102-2210202303032212-1331233202101021-3100322200223133-3320333111333213"></a>

## Direct properties — shop_make_payment / 332023330031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031202020222130-0311201323111023-3000220101031321-0111001033213211-3300102303302331-2112332002333332-3220011332123131-1001030221202030"></a>

## Next pages — shop_make_payment / 332023330031 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0011213111221013-2100013223312023-3020012303323230-2220033102001002-1331103022210111-1221001020032200-1233110301023232-3113330322021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323012200023030-2020020300231112-2123333102120103-1003210112021001-0200012100331132-2321303102102222-0231002203211011-1300111032032231"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order — shop_order / 111033123003 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-0202312022130002-3322132233123022-3010013302333211-2131202101222120-3123333311311130-1122021230303322-2001200122113202-3021011031212120"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
shop_order = {}
```

<a id="canonical-0232302102332300-3212211031222302-3220312001123323-0021320222303212-2220110313101020-0000202312233112-3102023130221033-1221230001021030"></a>

## Direct properties — shop_order / 111033123003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100212112131210-3210202322123121-2032001000211112-2031110121013133-2223200233002130-1221301300332230-1100302210231323-3311130031033112"></a>

## Next pages — shop_order / 111033123003 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1110300023112033-1013101033311200-1311322021330322-2012200323001223-2300000312031120-0310022301323221-3311223311300022-1012311111122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323320120310031-2230313221320131-3001111332120100-2030222201331313-2101220131032111-1132011131100103-3223330323102310-2132210211211212"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — shop_price_inquiry / 031130101022 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-1033333122321133-0102111123210322-0220002122232303-1321301303223032-1111132110012323-1313303120031032-0312212001133112-1102301301011000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop price inquiry.

Upstream description:

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
shop_price_inquiry = {}
```

<a id="canonical-2103101223213212-3211002131102211-2033113033211132-0220132311321220-1231230101033231-3220301002022022-0220102112221200-3002100333221123"></a>

## Direct properties — shop_price_inquiry / 031130101022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210311312120302-2233120320220110-3130301030020011-0120331111020020-1110322201111213-1103330323122000-0111232220212202-0323130232001123"></a>

## Next pages — shop_price_inquiry / 031130101022 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1130100110130001-1312332013300100-0211030011002200-2202102212232001-0131102320020033-2010311230213022-3102113133113010-1023212100303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202130213022133-0332020133231221-1031120320320200-1120131132230201-0333322322113302-3200010231112132-0300202332232230-0013323231032133"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — shop_promo_code_validation / 021121120113 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003301112323021-1003101203221303-3200130031112010-1113223021202211-0112323031302303-0233032001012210-3211303300322001-1101322113022321)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-3003231110112122-0001210302023011-3233230030321310-1310123230121101-0310230231133223-0212123021300103-3022301313213123-2111210200311202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop promo code validation.

Upstream description:

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
shop_promo_code_validation = {}
```

<a id="canonical-1231101310221233-3111330022000232-2231213132203010-3331021101020030-3332000222301323-2111212113231310-1133333321200301-2310313221100211"></a>

## Direct properties — shop_promo_code_validation / 021121120113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110321313323331-3230223231120323-3203210133333133-0331130312222103-0223331103110032-2323211121021213-0012033201001310-3301120231001012"></a>

## Next pages — shop_promo_code_validation / 021121120113 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-0303211230231300-2003111202022220-3313000233300133-0120313112201012-1221202223231121-0123123120133032-1312011231322232-1111310130000223)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0023110333210212-1012020102233112-1311311023032233-2020213001333003-2030012213031301-3203002203030133-1001311010222330-0302110102302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
