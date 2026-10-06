---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.web_client

<a id="canonical-3121310223103111-0301010320012013-3120110030132220-3202231230330102-1023232201201213-0320310200133333-1122201101321010-0333232030022222"></a>

Type: `"object"`. single nested block, Optional.

Web Client. Web client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("continue",
    "redirect")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
web_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013110231203301-3300113310332011-2020230022120113-1331301121301300-1021210032003332-0030220331323113-3112010100120323-3133023133203232"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client`

- [block](resources--protected_application--reference--group-002.md#canonical-3112012233110303-1213332313103120-2021013130012232-1213300123323111-0033233013221231-2213110231312112-1122201022120013-3213300223331123): complete subsection reference.

- [continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220): complete subsection reference.

- [redirect](resources--protected_application--reference--group-002.md#canonical-0331132031332122-3211021101303101-1130201233222223-2210221222322200-0212103121001101-0022333331320010-1333322112111011-1111113123101332): complete subsection reference.

<a id="canonical-3112012233110303-1213332313103120-2021013130012232-1213300123323111-0033233013221231-2213110231312112-1122201022120013-3213300223331123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.block` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.block

<a id="canonical-3113230031312132-1013333030110310-1103201101313101-2010020112300212-3221033123032102-0033012120323122-2132012332211031-1132012021211023"></a>

Type: `"object"`. single nested block, Optional.

Block Response. Block Response.

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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230212022013021-3131013122120303-0011231033130212-2011103012131333-2330313033210011-1020033221302312-3011030331230022-1232323210113013"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client.block`

<a id="canonical-1000303100301120-1313211001111013-3203111203313013-2010223223103301-0300221221100200-3013101311313000-1030032320321303-3200212200122001"></a>

#### `cloudflare.protected_endpoints.web_client.block.body` property

Type: `"string"`. Optional.

Body. Custom body message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-1033011230312201-2203003033100320-3223111330102330-1201133210210331-1033202101120213-3213323121201311-2100221113300310-1201122212320333"></a>

<a id="canonical-3110011011320011-2112011013210202-2310100122013301-3311103132300211-2003212131023212-3303101010321221-2330102021020003-2312011303101201"></a>

#### `cloudflare.protected_endpoints.web_client.block.content_type` property

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2211320110302003-3110131210202212-3111200323130010-1201030132030103-0220012122300213-3022212210331323-1303223030333102-3323030012231101"></a>

<a id="canonical-3201113301003221-3301222233011000-2010013003101333-3303123120303233-2000130221113220-0221011103100033-1303033331321300-1232223301021022"></a>

#### `cloudflare.protected_endpoints.web_client.block.status` property

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

<a id="canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.continue

<a id="canonical-0103301000301232-2102202133001231-3101010102310102-0132321123013312-2233113022313221-2322021020021312-2320120323301222-3010203201332102"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311113010003322-1120033313231102-2212210113112331-3333032220321313-0023121331001302-3221112013333123-1220032123022331-1203110132232311"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client.continue`

- [add_header](resources--protected_application--reference--group-002.md#canonical-1212021322320303-3230022010211001-0330201213203001-0122013012130331-2012233102113331-2212023101132322-2021101030321122-2223321213202312): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-1301121201301132-1110113100101213-1133303312330101-2030211022222133-0102033221132121-1001133101222112-1131311333302220-3230120013030303): complete subsection reference.

<a id="canonical-1212021322320303-3230022010211001-0330201213203001-0122013012130331-2012233102113331-2212023101132322-2021101030321122-2223321213202312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- Cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-2321303303011121-3102333113320202-0102303311323302-1220232333322011-2222131000233023-3323030232210031-2230332230102130-0000021200033300"></a>

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
add_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301121201301132-1110113100101213-1133303312330101-2030211022222133-0102033221132121-1001133101222112-1131311333302220-3230120013030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- Cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-0311232032311101-3201331022132222-0311230333032322-0300332202203301-2201021130112012-3312200211101303-3123032023133000-1231202233203030"></a>

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
no_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331132031332122-3211021101303101-1130201233222223-2210221222322200-0212103121001101-0022333331320010-1333322112111011-1111113123101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.redirect` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-2330331123021232-2310211101130032-3220122333323223-1212121020101113-3120103311003003-2322300032133210-2011013003012000-2331103030003211"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203213120022310-2201302132300023-2133323313321033-2121202323323102-0322031232232120-1202102031202003-2112031103023313-2111032223303310"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client.redirect`

<a id="canonical-3303322023003220-0220321223113110-2132311020120323-3102233031222011-2022222132023003-1130120000030212-2200011212301301-1302323303310203"></a>

#### `cloudflare.protected_endpoints.web_client.redirect.location` property

Type: `"string"`. Optional.

Location. URI location for redirect response.

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
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-3200300102212301-2232031211301231-2210311321000113-1133101333331113-3223322220001323-0332001323333222-3313001313000320-2232123133132102"></a>

<a id="canonical-0220323103022332-0333012311331010-1223121221213301-3130020132331302-1002213212312230-3323103120011131-3012301310012203-2231212102012122"></a>

#### `cloudflare.protected_endpoints.web_client.redirect.status` property

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

<a id="canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.web_mobile_client

<a id="canonical-2123131020201201-3112310030020222-1000120231311011-2121330012333300-0322020211303031-3333221011223102-3030032332030222-0120232031011120"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_mobile",
    "continue_mobile"),
  validators.ConflictingObjectAttributes("block_web",
    "continue_web"),
  validators.ConflictingObjectAttributes("block_web",
    "redirect_web"),
  validators.ConflictingObjectAttributes("continue_web",
    "redirect_web")}
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
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

Terraform syntax:

```terraform
web_mobile_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110110010002220-1323122012003120-2331232111111103-0232312122032202-0331010101010230-3322130322331222-3312003302330233-3330323112021211"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client`

- [block_mobile](resources--protected_application--reference--group-002.md#canonical-3011122022300310-2132312002222213-0001223030312021-1101320311023132-2133103130122310-3311010113322001-3312323102002103-3332020000013313): complete subsection reference.

- [block_web](resources--protected_application--reference--group-002.md#canonical-0300020103002300-3330212321220103-3000022131130233-3202203013212133-1031103010321203-0023213133313221-2131121330233012-2232323300132001): complete subsection reference.

- [continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131): complete subsection reference.

- [continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312): complete subsection reference.

- [redirect_web](resources--protected_application--reference--group-002.md#canonical-3223112032330233-3211013131233032-3321031030031233-1211001113012013-3010223031232211-2123312332113311-3100023303213011-0010021220103331): complete subsection reference.

<a id="canonical-3011122022300310-2132312002222213-0001223030312021-1101320311023132-2133103130122310-3311010113322001-3312323102002103-3332020000013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.block_mobile` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-0133011313331013-1123020310131212-1220023213112000-2210111021201323-1011301000302311-0112002121211201-1331110331131223-3232113110202010"></a>

Type: `"object"`. single nested block, Optional.

Block Response for Mobile. Block Response.

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
block_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022011200221121-1130130023013003-0100133032132320-1331031123200022-1213332121122300-1223010021120333-1002011200100201-3331033311321032"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.block_mobile`

<a id="canonical-0303122013133012-1311200103102331-1033010130233000-1033133110301332-3002110223131000-3001310232330030-0203323300323320-1330213020332333"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` property

Type: `"string"`. Optional.

Body. Custom body message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0003011101233020-1102323323011113-0211200332231122-3333330221323123-1101333130213100-2103323322113101-0103021003113131-2031213203313112"></a>

<a id="canonical-0113133100330011-2103310231322102-2123022131211330-2213330102200122-2213030303022002-3021132103230003-2330020123002113-2230212300201121"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` property

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1111102302303232-0302333100310210-0102323001312231-3031110122222102-0312230133002222-0103321311201202-2022122321031121-3321220222120131"></a>

<a id="canonical-2133123133222310-1031010031333333-0132033033013233-2303003222232211-1322023132033132-3133333332023003-2033100203212130-2133113110231301"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` property

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

<a id="canonical-0300020103002300-3330212321220103-3000022131130233-3202203013212133-1031103010321203-0023213133313221-2131121330233012-2232323300132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.block_web` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.block_web

<a id="canonical-2121130001221211-0013111003301112-2133210303201331-1133310022221301-1000321321310303-2311123121010010-2210200222132000-3030230201230212"></a>

Type: `"object"`. single nested block, Optional.

Block Response. Block Response.

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
block_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202000310330003-3113313003211323-2111220033230000-1010231333120301-0303020333021201-3200022322223310-2213321331023333-3303010202323133"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.block_web`

<a id="canonical-2302110221123123-3332211302310103-0121133000032001-2021002300103130-2002032303031301-1011021112301100-1111303211110012-2303302031001000"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.body` property

Type: `"string"`. Optional.

Body. Custom body message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0230230222313103-3122022003233301-0100003013301013-1013112130311330-2123020110012301-1332123333202032-2133230201212221-2013231302223121"></a>

<a id="canonical-1123023222111213-3033230003012011-2013002133013032-0003020222033123-3130231120003320-1202331033313032-1331230222232312-2132032332132012"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` property

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3123331322020200-0030320132330200-3200132200233223-0130333321201231-0323233332223322-3013131003120010-3001232120301002-2303220332333203"></a>

<a id="canonical-1021113332311023-3231002310113331-2300121223302113-0220010030202112-3302302010131003-2020120030011201-1023013122203223-3211321122013133"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.status` property

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

<a id="canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-0020312020023110-0332102220321200-3230300020103030-1132122301212201-0233213121010131-3203022302000111-0123111032103312-2102120101012013"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333133222320233-1102322133322200-0021310303331001-3123202232201230-3020101333030221-0023131031132003-0211100013323123-3311111102212233"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.continue_mobile`

- [add_header](resources--protected_application--reference--group-002.md#canonical-2102330112200223-3232310120303122-0221020231032121-0203100321232220-3201010221220023-2231023201203310-2220102323301002-1302333012122101): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-3211131321032221-0023223022233123-1200201002323332-0330121120210022-1313323322220013-0303302010313103-0123032330210210-2011221122101301): complete subsection reference.

<a id="canonical-2102330112200223-3232310120303122-0221020231032121-0203100321232220-3201010221220023-2231023201203310-2220102323301002-1302333012122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-1202022210201011-0123213303012133-0021013221331133-1130233300313303-3012203120122012-1022313201231031-3133230211232202-2211212333113303"></a>

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
add_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211131321032221-0023223022233123-1200201002323332-0330121120210022-1313323322220013-0303302010313103-0123032330210210-2011221122101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-3303302132232022-1331033131311302-2121311121201222-2030211302222123-3302123132132333-0131003203330201-3020330222201202-0302132133330303"></a>

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
no_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-1231113030022102-2000200322232122-2321312123020001-0113033202330212-0330111012330220-0303213222231103-1010020132330221-2102320221012112"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011033011101122-1233113031310022-2013233110010310-3120002200011331-2321233303311030-2103311300122113-1100033010120012-2031122322231130"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.continue_web`

- [add_header](resources--protected_application--reference--group-002.md#canonical-3232112211331123-0232322102013112-2011011120101230-1332323131200220-3300113010212303-3022010100230100-0212222101122000-3002112312002200): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-2123113300302133-0133103220000233-1323120313122230-2303121030233303-1111203210320201-0133130331321320-3331221103003212-3323132312222133): complete subsection reference.

<a id="canonical-3232112211331123-0232322102013112-2011011120101230-1332323131200220-3300113010212303-3022010100230100-0212222101122000-3002112312002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-3112123101221322-0223303102103030-3230202330103013-1233213330203102-1010233232312332-0001122023102003-3021210221121023-3001311310310102"></a>

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
add_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123113300302133-0133103220000233-1323120313122230-2303121030233303-1111203210320201-0133130331321320-3331221103003212-3323132312222133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-0211301332003020-1133323021030330-0010121000200202-2310012110103130-0000210211103002-0331032303232133-3220200013231022-0002333302010130"></a>

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
no_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223112032330233-3211013131233032-3321031030031233-1211001113012013-3010223031232211-2123312332113311-3100023303213011-0010021220103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.redirect_web` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-3321121013300203-2221300202100130-2001121130003103-2223110132233113-0030113330210212-3033221230011011-3021233113031113-3001101201201123"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
redirect_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011301202031032-3023020213012030-3303131131123313-1021003323111320-1120301331301100-0130133323003131-2022220031001321-0021010131230012"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.redirect_web`

<a id="canonical-1033102313130121-1212111200320312-2103010022303313-0320211001010201-2231122222210312-1130002201302220-1022010303201320-1022331001203312"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` property

Type: `"string"`. Optional.

Location. URI location for redirect response.

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
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-1231212322331113-2002101310332100-3002103011131103-0233313001232011-0113123001313021-2301311123033110-3200212033223130-2311023300212001"></a>

<a id="canonical-3333223202020201-3130101233212033-3332103023130303-1103301302022122-0010032330330023-1330113310110322-3133032300110211-0133323303330222"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` property

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

<a id="canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.trusted_clients

<a id="canonical-1131333132113202-2301303320003333-2310101333111312-1111122101112132-0233321030333333-2010121213002123-3233022020022211-3220210000021010"></a>

Type: `"object"`. list nested block, Optional.

Define your allowlists to skip Bot Defense inference processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323232003320130-2203231113131100-2312232131321130-0203303302130101-3321001000023201-2322023220323103-2213213032030103-3320331201301013"></a>

### Direct properties for `cloudflare.trusted_clients`

- [http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303): complete subsection reference.

<a id="canonical-2201222032231023-1100111113322000-0310001022010001-0102110101131113-1102200003213313-0102301312330023-1222313303121203-2221110231322320"></a>

<a id="canonical-2111201130311002-2311112230110123-0221230232013111-0111000223231301-2111223103301122-0321220110102310-0311113210302220-2132303311230102"></a>

#### `cloudflare.trusted_clients.ip_prefix` property

Type: `"string"`. Optional.

Exclusive with \[http\_header\] IP prefix string.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](resources--protected_application--reference--group-002.md#canonical-3110220033333323-0312031131030203-3323100222312233-0022112032113311-1113103231330102-2133330301210221-0020222332120210-2130011333123321): complete subsection reference.

<a id="canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- Cloudflare.trusted_clients.http_header

<a id="canonical-2213020301213230-2223211032033033-0202320213002123-3210032202222300-3202103221203122-3213131331002000-1012201022100103-3200023101103113"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013002211120101-3333302013331221-2001100312120012-1120011302020022-3032322331021031-3000011130112202-2112033301123001-2020032320310220"></a>

### Direct properties for `cloudflare.trusted_clients.http_header`

- [headers](resources--protected_application--reference--group-002.md#canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031): complete subsection reference.

<a id="canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303)
- Cloudflare.trusted_clients.http_header.headers

<a id="canonical-3330231022313213-3223033010203031-0313200320102002-0022023110303003-0302212332312310-2113001220001210-3233023111233123-0002131211310310"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-2111330011203032-2011102201030001-3222233011311202-1002202203102121-2321230020132111-3022101123002030-2300333220223332-0332210300213321"></a>

### Direct properties for `cloudflare.trusted_clients.http_header.headers`

<a id="canonical-1333300212333232-3301012321211102-1103332103020112-3213011313013220-1223202100120100-1331223101203002-1302302210112010-0022002331123111"></a>

#### `cloudflare.trusted_clients.http_header.headers.exact` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2321222303030310-2210030221113123-2032101101310213-1003332013310120-0020312120133031-2303222032133101-3331133000332130-0223201100032231"></a>

<a id="canonical-3032232121310111-0030112031300320-0003230012032032-0110202311302103-2021031032202132-1233121030103013-3111030211313030-3130312221213313"></a>

#### `cloudflare.trusted_clients.http_header.headers.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2230121101211233-0032331121321033-0300113110313301-0121212211202000-1322132012300033-3320013310323303-0022113313233133-3111203311012013"></a>

<a id="canonical-1212312313302301-2101132330120301-2302010033120323-3020220110233013-3202022020120210-2131020020031032-3000123003302000-0023301202020121"></a>

#### `cloudflare.trusted_clients.http_header.headers.regex` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3110220033333323-0312031131030203-3323100222312233-0022112032113311-1113103231330102-2133330301210221-0020222332120210-2130011333123321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- Cloudflare.trusted_clients.metadata

<a id="canonical-2201322213223112-0221310303320031-0310103213131102-3310301311330211-3030033213020102-3313111313313131-1033300012313331-2023323332223203"></a>

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

<a id="canonical-3111203322210231-1121231212331210-3132333131231130-1113030312231011-1013001312321311-1033012323113313-3020312221333221-0313332301111021"></a>

### Direct properties for `cloudflare.trusted_clients.metadata`

<a id="canonical-0111211300022002-3211121013120131-2301311331223031-2012223102013002-1301321232223013-0030102321103232-3100220330330301-2123320202310321"></a>

#### `cloudflare.trusted_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1323133113332330-1010301320001032-3201230023321303-3132322322330310-3332130233322011-1203321023322012-2203311333222201-0010112022030203"></a>

<a id="canonical-0103332033220320-0333210220223300-2213110032100332-1121233100103330-1101112321113102-2032300010211030-3303030310313232-1003120231110033"></a>

#### `cloudflare.trusted_clients.metadata.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- CloudFront

<a id="canonical-0323011203311321-0002123031021111-2112223320220213-1231321321221222-3112000200110112-2331021213200101-2130200001232020-2310023233330330"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for AWS CloudFront.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_endpoints"),
  validators.ConflictingObjectAttributes("aws_configuration_id_selector",
    "aws_configuration_tag_selector"),
  validators.ConflictingObjectAttributes("aws_configuration_id_selector",
    "disable_aws_configuration"),
  validators.ConflictingObjectAttributes("aws_configuration_tag_selector",
    "disable_aws_configuration"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "manual_js_insert"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insertion_rules",
    "manual_js_insert")}
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
  "x-ves-oneof-field-aws_configuration_type_choice": "[\"aws_configuration_id_selector\",\"aws_configuration_tag_selector\",\"disable_aws_configuration\"]",
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
cloudfront {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301323032122310-1023012122002120-0312303102202223-3322211012022132-0321110213011231-1210320302313031-0221221103013013-0111311121310030"></a>

### Direct properties for `cloudfront`

- [aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113): complete subsection reference.

- [aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322): complete subsection reference.

<a id="canonical-3023102100102300-2020110030232123-2220003032101000-0131310021331123-3232113012330302-1231031333101223-0311223313303212-2202020002221333"></a>

<a id="canonical-0012302202210201-1320222031110213-3131020310123133-2100222221300311-1131110321303311-0101103312102232-3232210201303230-3312210012000312"></a>

#### `cloudfront.continue_mitigation_action_hdr` property

Type: `"string"`. Optional.

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2122303003123230-3100332312320120-0100202001332333-3031321110321312-3123012101232213-0112120220231323-2003032103203231-0300310223013010"></a>

<a id="canonical-3303232113322331-3330210000232103-2110333030032212-3021030103220131-3030101221223223-2313233022222220-0110212030322231-3003130311111030"></a>

#### `cloudfront.data_sample` property

Type: `"number"`. Optional.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 1048576),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1048576,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1048576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1048576"
  }
}
```

- [disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100): complete subsection reference.

- [disable_js_insert](resources--protected_application--reference--group-002.md#canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-0021320220121330-1123001003212312-0103101010131121-2011033121210101-0313032101220000-3302332203012213-3111202202113020-2302220232120222): complete subsection reference.

- [js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210): complete subsection reference.

<a id="canonical-0200213222021013-2012212303233100-0013222033230112-2123103232303031-3332032132212033-2213310230003002-0123211310101320-2332102230131121"></a>

<a id="canonical-0322203331031123-1330123133021210-2313232020233000-3211002231200112-3030021101102131-1011120210003102-3233212312100230-1012120132000230"></a>

#### `cloudfront.loglevel` property

Type: `"string"`. Optional.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOG_DEBUG","LOG_ERROR","LOG_INFO","LOG_UNDEFINED","LOG_WARNING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOG_UNDEFINED",
  "enum": [
    "LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [manual_js_insert](resources--protected_application--reference--group-003.md#canonical-3313102332000123-2222313311313211-0013120212332103-3201031211032333-2332132333320123-1122123222002123-2323011023213133-1102033113123031): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021): complete subsection reference.

- [protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232): complete subsection reference.

<a id="canonical-1021002221210131-3333001132021221-3031121310230132-3310032330123023-2111231223310301-0111013032033001-3111132201000221-2010022331122030"></a>

<a id="canonical-1313133102222013-3003010202031131-1320302310002323-0303031300123000-2232323130031033-0323031002330223-0030132333311212-0121000003000030"></a>

#### `cloudfront.timeout` property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [trusted_clients](resources--protected_application--reference--group-004.md#canonical-1331313303213213-1122122223112222-3001222033333102-3201121101201113-1303223000202010-0120230110212212-1113101221303311-3220103301020330): complete subsection reference.

<a id="canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.aws_configuration_id_selector` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.aws_configuration_id_selector

<a id="canonical-3203022033232033-3300020132202320-1102030032203330-2003000231210111-3312100132300310-3123020102121330-3011301311010131-0202031002313021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws configuration ID selector.

Additional upstream details:

List of CloudFront distributions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ids")}
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
aws_configuration_id_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011302202020303-3231230010302202-3112310123122023-3330100303123232-0031311311132212-1103131321112002-0102010301302002-3102313003030203"></a>

### Direct properties for `cloudfront.aws_configuration_id_selector`

<a id="canonical-3303221233233132-0320333202103230-3223102230210232-1233010332032202-2100132030001101-1333323132221330-2013021033103023-1332331330212200"></a>

#### `cloudfront.aws_configuration_id_selector.ids` property

Type: `["list", "string"]`. Optional.

Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.aws_configuration_tag_selector` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.aws_configuration_tag_selector

<a id="canonical-3111200223233223-1200320210211330-0230003313032323-1021100122121321-3103023312000001-3333002130312022-1220101211330211-3303323330333200"></a>

Type: `"object"`. single nested block, Optional.

Distribution Tag List. CloudFront distribution tag list.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tags")}
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
aws_configuration_tag_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013130111331212-3311012232100021-1021301022032202-0132133133120130-1113322303132123-1133210112323203-0131223000033313-0222131032113201"></a>

### Direct properties for `cloudfront.aws_configuration_tag_selector`

<a id="canonical-3020023302330110-2111212012031200-1033132330322123-0311203023233002-0302330302031012-1002012103021021-2220302301233330-2330202313221023"></a>

#### `cloudfront.aws_configuration_tag_selector.tags` property

Type: `["map", "string"]`. Optional.

List contains the CloudFront distribution selection by tags key is an AWS tag name, and the value is
regular expression to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20,\"minProperties\":1},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.min_pairs\":\"1\",\"ves.io.schema.rules.map.values.string.max_len\":\"256\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\",\"ves.io.schema.rules.map.values.string.regex\":\"true\"},\"values\":{\"format\":\"regex\",\"maxLength\":256,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20,
      "minProperties": 1
    },
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
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.min_pairs": "1",
      "ves.io.schema.rules.map.values.string.max_len": "256",
      "ves.io.schema.rules.map.values.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.regex": "true"
    },
    "values": {
      "format": "regex",
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_aws_configuration` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.disable_aws_configuration

<a id="canonical-3010332231112002-0212022212200200-0102313303312133-3321122002011231-0313303023203122-3313003310233133-3302302230013120-0123113001011332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable aws configuration.

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
disable_aws_configuration = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_js_insert` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.disable_js_insert

<a id="canonical-1230102223233022-0213201032200202-2121230130310300-0201101222312303-3200023102330103-3110320331321110-2320112112030201-0211233121022103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021320220121330-1123001003212312-0103101010131121-2011033121210101-0313032101220000-3302332203012213-3111202202113020-2302220232120222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.disable_mobile_sdk

<a id="canonical-0023020223203211-3222322130010222-1303112203011322-0002030320131221-3231003133223110-1013210213202031-3320203303220120-2110212223330220"></a>

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
disable_mobile_sdk = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- CloudFront.js_insertion_rules

<a id="canonical-3210233031332321-1322302011001003-1221210302212332-2320120111323113-0031000121030112-2102222233000002-3200320002220101-0130303100011223"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030023123232121-1031322331020022-2231212210012123-1201213200231323-2321222302320303-3020021213211000-2011302101102022-0012011232103130"></a>

### Direct properties for `cloudfront.js_insertion_rules`

- [exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103): complete subsection reference.

<a id="canonical-2202113333201211-0202033333110130-3122112132112320-0102033330112232-0120131210132022-3323330202111303-0032112220023001-0230103022020231"></a>

<a id="canonical-0221200312013222-3011110100113033-1300002113020231-0223130122032113-3222200220233331-3031110212111321-0321230012010120-1030102300023031"></a>

#### `cloudfront.js_insertion_rules.javascript_location` property

Type: `"string"`. Optional.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT","JAVA_SCRIPT_LOCATION_UNDEFINED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JAVA_SCRIPT_LOCATION_UNDEFINED",
  "enum": [
    "JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130113100230033-3222100303210203-2033033033211020-2231210332003110-1021200220210211-3022213130222300-0011333102002313-3021131033130331"></a>

<a id="canonical-3110201222120110-1021002102002203-0012131211102012-2113100120101232-1113121023112013-0233111101332112-0230210131131102-2120000221132112"></a>

#### `cloudfront.js_insertion_rules.javascript_mode` property

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

<a id="canonical-1220121003202302-1221023020103302-3333133331011023-2113030202202010-3120020131211223-3110113030200310-3210233230210213-1321322232010110"></a>

<a id="canonical-2023100011023113-2301213212211013-2211230200313231-3202332111112312-2002110221200130-0013000113231022-2302332022212020-1112010020321111"></a>

#### `cloudfront.js_insertion_rules.js_download_path` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules](resources--protected_application--reference--group-003.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110): complete subsection reference.

<a id="canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- CloudFront.js_insertion_rules.exclude_list

<a id="canonical-1101033233220033-1113023000031000-0030302122212213-2022303303330232-2112200022103331-3100233310302212-1303323211233213-0031011213222123"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321210221010021-1110200202232310-1322233333221003-3122120213203332-0112002322032132-3003312112130311-1300303332200123-1130123130113313"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list`

- [any_domain](resources--protected_application--reference--group-002.md#canonical-0022030103303102-2202331303322220-3022121202010330-1031302103233020-2112220330332130-0010031203101232-2212333311021031-2113120120013133): complete subsection reference.

- [domain](resources--protected_application--reference--group-003.md#canonical-0223132131303030-1021321002133031-1322031020210023-3021012231212331-3323212220200021-3333310133301130-1033301220110133-0232301203331323): complete subsection reference.

- [metadata](resources--protected_application--reference--group-003.md#canonical-2010012310003231-3020100210200133-2210333023120033-1323211001220022-1201221122322313-0012223022121231-2311120202101323-3322032232313013): complete subsection reference.

- [path](resources--protected_application--reference--group-003.md#canonical-2021213021001211-1331122120210300-3310333130003330-2200033112003033-1101020130331102-3033121122222031-1213332313232000-1103013002303331): complete subsection reference.

<a id="canonical-0022030103303102-2202331303322220-3022121202010330-1031302103233020-2112220330332130-0010031203101232-2212333311021031-2113120120013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [CloudFront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- CloudFront.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3302113222033311-2330222021330030-1232232230111230-3132310211021313-0303002021030313-2302010301330002-3302331002201222-1100003210233000"></a>

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
