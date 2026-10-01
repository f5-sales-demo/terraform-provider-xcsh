---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1230030310313333-1331133020002011-0332100333200230-1101323202002200-1203211321010230-3110132022030123-1331301100200223-3021102200301331"></a>

## status property — block / 121021311131 / 5

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

<a id="canonical-2020211212311022-1220302203013133-1231000331100230-1133300131201013-1230133322120103-3123311020301212-1010011103022230-2123302132103220"></a>

## Next pages — block / 121021311131 / 6

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000032111112200-3211003123231332-0011321123310302-3232010111231320-1130003210010012-0231222232030200-2313230233031010-3202121010233132"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag — flag / 031211313200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="canonical-1111201203311012-0220020201120300-0210331212113230-3323211121033211-2010003113333112-2331110331202213-1231322332312222-3223200232113003"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132320300103033-2022123212201201-1030133310232221-0010311002322200-1313312323222020-2202011222222032-3012013102333312-1333211133333213"></a>

## Direct properties — flag / 031211313200 / 3

- [append_headers](resources--http_loadbalancer--reference--group-013.md#canonical-1213312221221213-2031002222203202-2021211223321101-1200032013313110-0202003213321101-2000333232320323-0132231003111201-1022010330023110): complete subsection reference.

- [no_headers](resources--http_loadbalancer--reference--group-013.md#canonical-0211301100002311-2002221301313232-1321120122303213-0100132032331320-2302302022002311-1223322321301211-2101310303113030-3323212202102222): complete subsection reference.

<a id="canonical-2230023033202321-0001001002121203-1123011210203311-2102330312223220-0003201322000201-0130301212013310-2310011000213322-3112221020323021"></a>

## Next pages — flag / 031211313200 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](resources--http_loadbalancer--reference--group-013.md#canonical-1213312221221213-2031002222203202-2021211223321101-1200032013313110-0202003213321101-2000333232320323-0132231003111201-1022010330023110)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](resources--http_loadbalancer--reference--group-013.md#canonical-0211301100002311-2002221301313232-1321120122303213-0100132032331320-2302302022002311-1223322321301211-2101310303113030-3323212202102222)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1213312221221213-2031002222203202-2021211223321101-1200032013313110-0202003213321101-2000333232320323-0132231003111201-1022010330023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211120213012122-3311212332201221-2000220301210012-1223201031300201-1211130103132312-2331230103013031-3233110022212212-3200301121011111"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers — append_headers / 323213101323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers

<a id="canonical-1320212321122303-0330022202031233-1020221132012032-1011231220200322-2312112210101031-0012201323201212-1320312220022001-3001020322133020"></a>

Type: `"object"`. single nested block, Optional.

Append flag mitigation headers to forwarded request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auto_type_header_name",
    "inference_header_name")}
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
append_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221210312131232-1122300320311221-3212011202310102-1310323203213312-1110010020203121-3132012233020120-1211123301130302-2233122232230312"></a>

## Direct properties — append_headers / 323213101323 / 3

<a id="canonical-3300112202021111-1211202230111020-1121101112322023-3123320010310310-3012303112301021-2222132132112221-0031233321213002-3331110131122213"></a>

<a id="canonical-1102312221221211-1323213232121200-3102023133013211-0011002031032331-2010013012003122-3003313311123021-1121320100300313-0021222202221330"></a>

## auto_type_header_name property — append_headers / 323213101323 / 4

Type: `"string"`. Optional.

Automation Type Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2031131233231211-1212311312232120-1232120233001233-3200123312303020-3211132011010232-3122312001011120-2121302012113222-3012233333121010"></a>

<a id="canonical-1030232200102231-0203331332301102-0231130303202032-0030123320313323-3233121030323023-1330210232101220-1313011202010120-3220321020100130"></a>

## inference_header_name property — append_headers / 323213101323 / 5

Type: `"string"`. Optional.

Inference Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3323022231101110-0220311132310121-3313103233032122-2032321121120231-2033200030230311-2022312133222131-3311111213013010-3311102332102332"></a>

## Next pages — append_headers / 323213101323 / 6

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0211301100002311-2002221301313232-1321120122303213-0100132032331320-2302302022002311-1223322321301211-2101310303113030-3323212202102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231312330012131-1332100013302222-0020033033322211-0101301103110110-0103031000230023-0210201233032113-2033011330110121-1321032102033022"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers — no_headers / 323013300132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers

<a id="canonical-1301201003033201-3301312200102033-0011102313321223-0131201203100221-3102312101222230-1112313020311232-3100113113223101-3032312301302220"></a>

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
no_headers = {}
```

<a id="canonical-1233222133320201-3103131001231322-0310033231003221-3222012221223121-2331131233311133-1331110030311223-2232111030113310-3113211200121021"></a>

## Direct properties — no_headers / 323013300132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322030101303313-2031310303232101-3032103222103021-2210033033121303-2133300300311211-0323223222202320-0221031330002203-3013300030223210"></a>

## Next pages — no_headers / 323013300132 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213311020230121-2103112011112112-0231111001200101-2322233211001233-1300231102301121-3001131333013232-0033032100321300-1113312322131201"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.redirect — redirect / 302030133121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-2110331300223000-2231230223320130-1102203331301330-2133001311221300-0113033302223202-2202023202312233-1002120002203332-3230113301013022"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Upstream description:

Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
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

<a id="canonical-1200133200123122-2122111321313101-2323331110100232-0121210231101113-3113130032131111-0011210013332223-3112122301332132-2100311201120332"></a>

## Direct properties — redirect / 302030133121 / 3

<a id="canonical-0310212310321303-0122131332003130-1202301212103010-2311003302332332-0020302103232021-3131231320323301-2013101302022000-2300132202222133"></a>

<a id="canonical-2023200211230132-1100333033001020-1322031303000331-3222202311030331-1123313112311210-0211122220211202-3113312012332033-1210322303300130"></a>

## uri property — redirect / 302030133121 / 4

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-0301110300203301-2201001033223130-3001012330012201-0233233222303010-2132100120220103-3012322011213113-0003323331303323-2120212201202033"></a>

## Next pages — redirect / 302030133121 / 5

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0200310013012011-1313310020131023-0203321100301133-2101212302102010-3320231321202232-3031131303020130-0120030303011112-1220331131110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323233213311300-3131320210120121-1011002303121002-3213123020203133-2031122303212200-2020033021302012-3201111323302210-2201123210311131"></a>

## bot_defense.policy.protected_app_endpoints.mobile — mobile / 303310112332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-3313303032331200-3232300021021302-3201202210123022-2123122011112212-0013110302223232-3302012020203332-0332120013032022-0100111110021120"></a>

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
mobile = {}
```

<a id="canonical-2021232110112211-3131113020323032-3213323221122301-2133032321312110-1033300303000330-3220212101321200-1201012021331300-0311121033322030"></a>

## Direct properties — mobile / 303310112332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010021013001212-3313203010210031-0010331303320001-2223002011102222-1231023333111311-0031023320021322-1210222031032123-0233123210030330"></a>

## Next pages — mobile / 303310112332 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3132020021201300-2213301200112303-1020320300322230-3032200333333311-3311101002332123-1010201023012022-0122220101130023-1320121003331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013231102013202-3332003231130313-0122000122233110-0123303101210211-2123113003103232-1100113110221220-3301130310302213-3122231311301032"></a>

## bot_defense.policy.protected_app_endpoints.path — path / 322113101020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-2102102331201332-3031310220311310-0210310331230303-2131322131222330-1032213332201311-1031213020310130-1001231331121131-2133110323302003"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0331132321203223-1100303132111010-1321313330131102-0313002310222301-0203132230010021-2222211332030223-1213011230021130-1031331121210212"></a>

## Direct properties — path / 322113101020 / 3

<a id="canonical-0231120322232033-1222330301213110-2303203322232230-1230332311000201-0323312112213231-2220000033022011-1000300330302110-0301203233203302"></a>

<a id="canonical-0121103303012322-3331202020000022-2320012000220103-2001132211022313-1003030101123200-0010332020331030-1031002102132132-1323130000103111"></a>

## path property — path / 322113101020 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-0210320230130222-2132013203023330-1232033232213320-0020121321232113-2333013033303112-2303010110231122-2332003300132203-0013223303030010"></a>

<a id="canonical-1211102112103201-2013121003210333-3122022313330112-3120303232311322-0020323001211220-0333303322202212-1102111021230022-1111013012231211"></a>

## prefix property — path / 322113101020 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2000323232003103-1303122001032101-2110100011012123-0110010221202122-2211023002023201-2103332213031303-0131220101011021-0301133210133132"></a>

<a id="canonical-0321321231111132-0133221230313011-3133101221300212-2103032231132132-3310331112230120-0333310201221121-0013231023231133-2100010100020313"></a>

## regex property — path / 322113101020 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-2122103231002300-3200330002112213-2311230321212322-2231031310102310-2211331322223321-3130111121003020-0003213332130112-1233312001031030"></a>

## Next pages — path / 322113101020 / 7

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130130003012032-0000203233102331-3233320322103221-3231302301311320-1133302011012023-3002320211321213-1101330220001031-0321301233003211"></a>

## bot_defense.policy.protected_app_endpoints.query_params — query_params / 032322122002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-0011121301220020-0322113321302121-2233221320233011-0033221120220332-3200313111112303-3102132331232110-0200010100330023-1123201203103132"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    }
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002003133303221-0201223012132200-3111232310132013-3330002221023123-1133020220321313-1123323310002020-2221310232203321-1133233332121021"></a>

## Direct properties — query_params / 032322122002 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-013.md#canonical-2101332322221030-2100012222013201-1113313132003002-1331113222301200-1012332201201212-2130021021102202-1032321332312322-3012021031112223): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-013.md#canonical-2303010213110121-0113212220013123-0130010231301113-0132322301333233-0212030200331313-1212301133123000-2122013231332132-0020202330230001): complete subsection reference.

<a id="canonical-1220033123103110-3320010303031313-1201321012312322-3111330202112303-2032301102232211-2330222120011230-0121110222123322-2210310003213122"></a>

<a id="canonical-2001330231102012-3031113332330230-2010220023100111-3332212223302310-3301001201211211-2111030011200332-3030322211231323-2112332300033003"></a>

## invert_matcher property — query_params / 032322122002 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-013.md#canonical-3013113233301011-1230322322233002-2033320211013202-3320003010033202-3021232200131000-3310210332031313-1121113230030033-3130031023300023): complete subsection reference.

<a id="canonical-2122003301100110-3013202221322133-2331332032310320-0131221031001200-2323101322232101-3100211031011111-2313333112320002-1332021332223100"></a>

<a id="canonical-3203101010212333-1100012102331202-3333033303330310-0213100311121300-1112002112313331-0100312023111230-1023313100021231-3013102222000302"></a>

## key property — query_params / 032322122002 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3012223121322222-1023032002031211-3202332032100120-3013131220131230-1220232232123001-0121032302133203-2312030323333311-3103110001321022"></a>

## Next pages — query_params / 032322122002 / 6

- [bot_defense.policy.protected_app_endpoints.query_params.check_not_present](resources--http_loadbalancer--reference--group-013.md#canonical-2101332322221030-2100012222013201-1113313132003002-1331113222301200-1012332201201212-2130021021102202-1032321332312322-3012021031112223)
- [bot_defense.policy.protected_app_endpoints.query_params.check_present](resources--http_loadbalancer--reference--group-013.md#canonical-2303010213110121-0113212220013123-0130010231301113-0132322301333233-0212030200331313-1212301133123000-2122013231332132-0020202330230001)
- [bot_defense.policy.protected_app_endpoints.query_params.item](resources--http_loadbalancer--reference--group-013.md#canonical-3013113233301011-1230322322233002-2033320211013202-3320003010033202-3021232200131000-3310210332031313-1121113230030033-3130031023300023)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101332322221030-2100012222013201-1113313132003002-1331113222301200-1012332201201212-2130021021102202-1032321332312322-3012021031112223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211312303132222-1221102020332111-3302322210001013-3232110300330003-3010131210131102-1013200230001223-0030233210202200-1321131133330102"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_not_present — check_not_present / 300113121212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-1132101020022110-1121013301110201-2322322022132001-2013011301323312-3021320131203101-1030010003103333-2211003101311102-1333133220311011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1322022022031302-0032200323111020-1031013303322212-1131203112222210-1023320102221000-0322223231202331-1311023013301020-2122031311332000"></a>

## Direct properties — check_not_present / 300113121212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100201130022220-3010213211303011-2232203203213113-2202333212331011-0310002031101121-0101102211201203-0101022121330203-2301122113113231"></a>

## Next pages — check_not_present / 300113121212 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2303010213110121-0113212220013123-0130010231301113-0132322301333233-0212030200331313-1212301133123000-2122013231332132-0020202330230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320132212010032-2323120113221003-2231233013003232-0013312133212231-0201010210231320-2113323310013320-1001112300322110-0132300031222122"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_present — check_present / 010002332011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-2112033302211131-2212211113311013-3320101223033202-1013101013030232-0311320331130311-1230223100000300-1123211132132332-1230013200103313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3021213332232221-3212301231202302-3301101032321021-2203231130320133-2223333333032231-2101012220301231-1200011112121311-0212023310003031"></a>

## Direct properties — check_present / 010002332011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122022021011202-0012101220311012-1222111102211301-2013112233221313-0220202100300300-2300221101033032-3330313023331102-1302031010132332"></a>

## Next pages — check_present / 010002332011 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3013113233301011-1230322322233002-2033320211013202-3320003010033202-3021232200131000-3310210332031313-1121113230030033-3130031023300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012210020030130-1201303100130022-1013132132113310-1031001002203110-2030010321310331-0001101231210311-2211333212022310-2123011302221222"></a>

## bot_defense.policy.protected_app_endpoints.query_params.item — item / 102103101020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-0013301000320112-3131203021212303-2221033132301211-0033331323232201-2101033310100213-2033201130203312-3101100030323002-2101013213023223"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123021013211021-2112331030022212-1202200003023002-2231032100302112-3002012010122111-0322011012102102-3113023110113123-1103331000120030"></a>

## Direct properties — item / 102103101020 / 3

<a id="canonical-1012120213310332-1220000301320103-0103322321033031-0322213211130003-0303102132333213-0210202212210301-1302012103120231-1013003200002101"></a>

<a id="canonical-2110102110322101-0331032021012020-3312010132203131-1023211113300111-1010112200201013-2320230332312220-2301033023003220-2200211100110123"></a>

## exact_values property — item / 102103101020 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0100211110013221-3001123030100203-0200323110113012-0220302300030132-1132133210013230-2001132010103003-3111210213122301-1330321311231231"></a>

<a id="canonical-1301303320330001-2031021122130301-1210000201220220-0210033131001311-1001003122000100-2200203301213321-2211110330003311-1233103132030213"></a>

## regex_values property — item / 102103101020 / 5

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0130120321022212-1022230333231321-1330022213030110-0222202231001111-0330231320213320-2300203122202230-3021131232031313-0111033202133220"></a>

<a id="canonical-3033330220000113-1331022010010312-0323131200230222-1213321313333200-0120322323133003-1212000220103102-3123230130330333-2232302230321313"></a>

## transformers property — item / 102103101020 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1022121202300222-1321002313232022-1302111023222112-3011320010002310-1110003030132230-1113003012300003-1310113011200102-3023021011311301"></a>

## Next pages — item / 102103101020 / 7

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1203210321222120-0033120203013313-2131212022012032-1233321301313330-3120311213033232-3302302000221200-3023302011011311-0232031130132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102213321223200-3201122212331220-3201001312301313-0331203133112231-0112230223332003-3112320201302013-2022002021101223-2223320223323103"></a>

## bot_defense.policy.protected_app_endpoints.undefined_flow_label — undefined_flow_label / 001101100120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-3231333310002230-2111320230302220-3003230031133120-0233301332231012-0203123211201132-0001312013032023-2123110201223011-3303230001201011"></a>

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
undefined_flow_label = {}
```

<a id="canonical-2212001311302023-3310322031002330-0132222000133100-0212020301103321-3330310023031010-1002101231313311-3030000320010232-0323121132132211"></a>

## Direct properties — undefined_flow_label / 001101100120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233030001133222-2330230331133332-3202230211212202-3331020323013233-0001230022102222-0231122122221302-3331213010311033-0332331323303133"></a>

## Next pages — undefined_flow_label / 001101100120 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130020123113200-0230312201110222-1333301202301130-3302211302301012-2231221131222131-3321323201300333-0101100233311213-2020002023103320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223101112200222-2133320013130301-1100222321020312-3310211201112120-0121211133020312-2321322021103233-3222323113013322-2320233000133221"></a>

## bot_defense.policy.protected_app_endpoints.web — web / 211300313302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-3332033113333131-0102010031313110-2003301203222232-0213332102233133-0303201321103330-2013200333310310-3223301221322331-0210112122320232"></a>

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
web = {}
```

<a id="canonical-2023011131110221-0221013220102100-0000231331020132-2102132200321201-3221212233121120-2223112323320303-2201021021002220-3311231102130302"></a>

## Direct properties — web / 211300313302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222111221032302-0101010123232103-3321010301031022-3321010222003202-2111112212322203-3220233010002303-1101100001033021-1132123112003203"></a>

## Next pages — web / 211300313302 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101313121221202-1301032013003023-1220301222330001-2201122130031032-2313210123222031-3023312033232202-2101213232211113-0202330221311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300120032121133-1302303030010101-1032323031022330-2132322010201131-3000130120331321-3313303323111023-3320322132003301-2201123001302011"></a>

## bot_defense.policy.protected_app_endpoints.web_mobile — web_mobile / 202132132203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-2310331132311233-3121223003002301-1030010212011311-3321313131330120-2333031302301020-0331101012213133-2331223011010231-1131310202231202"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

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
web_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011032323111030-1233113101320313-3312211101020220-1231102010210001-2212303130023321-3213030203313203-0230221203202111-3210101121101223"></a>

## Direct properties — web_mobile / 202132132203 / 3

<a id="canonical-3002103320313313-3000321033103112-3132331002323023-0122013332133113-0213011211331201-0313313131133003-2001332333023103-0013320113110220"></a>

<a id="canonical-1000131112230013-1302130302011130-3213213022030111-1112120311122213-1023330111322100-3023220333310313-1103221220103220-0223203300101300"></a>

## mobile_identifier property — web_mobile / 202132132203 / 4

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311112212303010-3321021202033300-3232030010303303-1001212301232201-2130223323213130-0330020133033130-1222023202202113-3202301111030130"></a>

## Next pages — web_mobile / 202132132203 / 5

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223121030001112-0102000002332110-0102101031231100-2130122203112003-0013020002312232-0323333130123133-1330100202131000-3212231222111001"></a>

## bot_defense_advanced_protection — bot_defense_advanced_protection / 311232332113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- bot_defense_advanced_protection

<a id="canonical-1331023323322330-0120310010111332-1123202023223302-3231330202031222-2112233201102033-2111320023203321-1120030130231113-3212030313203022"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("both_web_and_mobile",
    "mobile_only"),
  validators.ConflictingObjectAttributes("both_web_and_mobile",
    "web_only"),
  validators.ConflictingObjectAttributes("mobile_only",
    "web_only")}
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
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

Terraform syntax:

```terraform
bot_defense_advanced_protection {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201033132120331-1100101030102310-3002310202233201-3300120300030223-1330322111310133-0110332213121301-3210112232232222-1323031222131110"></a>

## Direct properties — bot_defense_advanced_protection / 311232332113 / 3

- [both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233): complete subsection reference.

- [mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033): complete subsection reference.

- [web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212): complete subsection reference.

<a id="canonical-3112332020222301-1233023022102032-0022013110310212-0333013313111220-3300302321033001-1131322201130301-3012100120103333-0002022113121011"></a>

## Next pages — bot_defense_advanced_protection / 311232332113 / 4

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103232201111222-2222132213210221-1121200303013312-1000303210311022-0131102000020012-0001221301010133-1003012011200211-3312032013113213"></a>

## bot_defense_advanced_protection.both_web_and_mobile — both_web_and_mobile / 300323320122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.both_web_and_mobile

<a id="canonical-0011111001010222-3031000001230011-0331010111323231-1132312320133201-3200030003211130-2133003121022102-2313033122133020-1031222202012023"></a>

Type: `"object"`. single nested block, Optional.

Both Web &amp; Mobile. Both Web and Mobile configuration.

Upstream description:

Both Web and Mobile configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
both_web_and_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003032001022213-0323313322322000-2000300232002110-1032022210200310-3010012110013001-1303221000101113-1222333132113332-3213010322230000"></a>

## Direct properties — both_web_and_mobile / 300323320122 / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-013.md#canonical-0313310300303030-0332030023231201-3100303131122331-2231221230211102-2222331033131311-2022032021031222-3233003002000020-1323231023031312): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--reference--group-013.md#canonical-1003101001321320-0233032302330132-0311303011302022-0111210200301000-0311213323331133-0020211303331000-2001332210101131-1033222210001110): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-013.md#canonical-0120102310320113-3323113113120310-3311231020303300-0030331100311200-1230011032000020-2011301210131032-0100213103123202-0022322122133201): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110): complete subsection reference.

- [mobile](resources--http_loadbalancer--reference--group-013.md#canonical-3113113321230312-1011132221021022-2222200122232321-1121313322131023-3010032012223210-0112301112312333-2003213221232012-3022332221010201): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-013.md#canonical-3200000311231133-3323200203131332-0220323231120332-1111202111132333-0131021232301123-1131022130213101-1010110003233102-3213003220120002): complete subsection reference.

<a id="canonical-1101000323000332-0200311022231221-1022101031303123-1013312321333202-2201213323110202-1001303311132210-0201130010220220-0023200220020102"></a>

## Next pages — both_web_and_mobile / 300323320122 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert](resources--http_loadbalancer--reference--group-013.md#canonical-0313310300303030-0332030023231201-3100303131122331-2231221230211102-2222331033131311-2022032021031222-3233003002000020-1323231023031312)
- [bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk](resources--http_loadbalancer--reference--group-013.md#canonical-1003101001321320-0233032302330132-0311303011302022-0111210200301000-0311213323331133-0020211303331000-2001332210101131-1033222210001110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages](resources--http_loadbalancer--reference--group-013.md#canonical-0120102310320113-3323113113120310-3311231020303300-0030331100311200-1230011032000020-2011301210131032-0100213103123202-0022322122133201)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile](resources--http_loadbalancer--reference--group-013.md#canonical-3113113321230312-1011132221021022-2222200122232321-1121313322131023-3010032012223210-0112301112312333-2003213221232012-3022332221010201)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.web](resources--http_loadbalancer--reference--group-013.md#canonical-3200000311231133-3323200203131332-0220323231120332-1111202111132333-0131021232301123-1131022130213101-1010110003233102-3213003220120002)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0313310300303030-0332030023231201-3100303131122331-2231221230211102-2222331033131311-2022032021031222-3233003002000020-1323231023031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122213112110113-2222330131001002-1133320000131113-0121103211200001-1302220001002311-0001110110220323-2000003013220012-2111321032223022"></a>

## bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert — disable_js_insert / 112303001031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert

<a id="canonical-2100230222030030-1330131300220203-1033122213122100-1013102230222211-1222231002320301-1130202231323011-3110112100302330-0102121302130333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-1012120212203232-3220111332200310-2011121231333100-2033213000103002-0230230121210201-3121011011030323-3123032113233221-0213213001203120"></a>

## Direct properties — disable_js_insert / 112303001031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022100132212232-3311213201031002-3120313113320112-0303300102303223-2212121112100212-3032131313222113-0030220312100221-0023200131310132"></a>

## Next pages — disable_js_insert / 112303001031 / 4

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1003101001321320-0233032302330132-0311303011302022-0111210200301000-0311213323331133-0020211303331000-2001332210101131-1033222210001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020232213233030-2122033202110230-1333011202220300-2310300103102032-0122332332000202-3230230330111332-0113100113111233-0303013032003213"></a>

## bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk — disable_mobile_sdk / 332200133021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk

<a id="canonical-0333311120232000-0331332212030002-0001102000000320-2030121210010211-1303133002201220-3132213223310311-3233302121102000-0330213212200200"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-2333303101203311-0223211200102022-0113130300100201-3130132130022221-3201123031123133-1312200222330303-1111131000331101-2120202303231210"></a>

## Direct properties — disable_mobile_sdk / 332200133021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223100233222313-3120322131033221-2012210030232301-2130122133011320-3320212230331120-2103203123330002-1203322111122330-2222221122030010"></a>

## Next pages — disable_mobile_sdk / 332200133021 / 4

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0120102310320113-3323113113120310-3311231020303300-0030331100311200-1230011032000020-2011301210131032-0100213103123202-0022322122133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132333022132102-0121023310233223-0100000211003201-2220001121111033-2202101310030220-1110132021330303-1223031123133121-1221320332332312"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages — js_insert_all_pages / 332002333323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages

<a id="canonical-3022023332201123-3011011102310101-1311131033002301-1020222013113302-3002223002300020-3230033013113330-3010132313103130-2313102230200101"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320001223100122-2020222301103132-2012313000003133-3022121230232302-0021033021101312-3211023100330003-2220220113301223-2113103202313330"></a>

## Direct properties — js_insert_all_pages / 332002333323 / 3

<a id="canonical-2000131323110012-2202131212011000-2013301230200312-2101130213221101-2031301320211333-2001021133013131-3111112002313110-2130123132122012"></a>

<a id="canonical-3211102010031030-1311032302032102-3032101203122211-2023132202123111-3003131033222321-0220122320022200-0103033131232130-3222212012300210"></a>

## javascript_location property — js_insert_all_pages / 332002333323 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

<a id="canonical-0230032113132322-0002110130013021-2102110233301010-2030210232031221-2000210033202122-2330201300012232-3312001202133201-0311012320023311"></a>

## Next pages — js_insert_all_pages / 332002333323 / 5

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232322223132312-1232233102103013-1313003030310012-2330121212231101-3031230201230230-1320030232311000-2112020013103232-1013200022311233"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except — js_insert_all_pages_except / 331110310302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except

<a id="canonical-2310302220102221-0221212010022122-0211302010200230-1313330303002323-2031131010322012-0023000300002123-3011103212033213-1222231010300333"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311300203301021-3101113100002312-1303301021210103-2333000221011031-0033223111031023-0022031001312121-1100203331301230-1022113102330101"></a>

## Direct properties — js_insert_all_pages_except / 331110310302 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123): complete subsection reference.

<a id="canonical-0233300213210101-3303221222222130-2300103032013201-1212332313101021-3100301320331012-3303102011003033-1030213311020002-0000131003331101"></a>

<a id="canonical-2233202031302230-1121111010100230-0221112330210233-2113020030003010-2030012331330232-0302001313122333-0012223133131300-2130230020011121"></a>

## javascript_location property — js_insert_all_pages_except / 331110310302 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

<a id="canonical-3312210012211301-3111010000011100-0202020221112320-0331111331301000-0011303303310230-3030011111200322-1023002311013222-2302032320303221"></a>

## Next pages — js_insert_all_pages_except / 331110310302 / 5

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032100113011311-0103001230210322-0333232301123112-3221033132333032-1112100202330202-3331023210230010-2013111113130320-0211320211200123"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list — exclude_list / 232012032202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list

<a id="canonical-3300211223123133-2310330102320301-0230213223031313-0231223332232332-1110130231232001-3200113221330100-2131120033011311-1311023200202323"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2333331221010111-1031203030102010-1320032032333012-2210033013210223-2131002332011221-3300321110032033-0013233213123231-0032123213330332"></a>

## Direct properties — exclude_list / 232012032202 / 3

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3331232331321231-1221000220032021-1303320301323033-3103230220000310-1111233003133033-3331021131033013-3202030001330303-1232100322202313): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-1002310310313002-1313312000211033-1133131121233023-3002320023122312-3221110013120120-0110120211231033-3211313223021322-1300100131321201): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-3310223023102312-0113032103112001-3230233230132310-2323022123331023-2220120001200003-1323213303100030-0011322010230312-2020101203300032): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-3223123302031303-2121331213002022-3030100231303021-0111222033122202-3100231200220022-0023302321101222-1031231113110012-0202121330000121): complete subsection reference.

<a id="canonical-0213203021121100-0013003302102230-2012321010012211-0032313033013130-1323302021332310-2130323302323211-3323021322013012-3101232131112301"></a>

## Next pages — exclude_list / 232012032202 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3331232331321231-1221000220032021-1303320301323033-3103230220000310-1111233003133033-3331021131033013-3202030001330303-1232100322202313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-013.md#canonical-1002310310313002-1313312000211033-1133131121233023-3002320023122312-3221110013120120-0110120211231033-3211313223021322-1300100131321201)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-013.md#canonical-3310223023102312-0113032103112001-3230233230132310-2323022123331023-2220120001200003-1323213303100030-0011322010230312-2020101203300032)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-013.md#canonical-3223123302031303-2121331213002022-3030100231303021-0111222033122202-3100231200220022-0023302321101222-1031231113110012-0202121330000121)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3331232331321231-1221000220032021-1303320301323033-3103230220000310-1111233003133033-3331021131033013-3202030001330303-1232100322202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232313301102122-3000101300303301-2300132301112011-3133323221303122-2101230333201302-3023003222013232-3133133203233311-3122122301132201"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 121220302221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1123102122021033-3001221033221033-3110113111333001-0032303023121021-1320203312302300-0033110010231230-1332011030210002-0200333011130201"></a>

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

<a id="canonical-3011012323323120-2000301300322012-0211202331101210-2211233320003220-3212113011202001-3231020123211120-0102000000333330-2023031113011022"></a>

## Direct properties — any_domain / 121220302221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022222033301113-0302003111031312-2011030302333001-3222203302010122-3003030201013113-0223300211302030-2032301133120201-1020212010310330"></a>

## Next pages — any_domain / 121220302221 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002310310313002-1313312000211033-1133131121233023-3002320023122312-3221110013120120-0110120211231033-3211313223021322-1300100131321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122022232322211-1223023113012123-0110232323011030-2120231210011223-1130311001202002-1021100022221013-1120302020330130-0001102300031333"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain — domain / 300020013311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-1302200000121101-1203001023210313-0213232303132101-2132032223220212-3122212012032132-1331311121002013-3010110200333322-0003010002112130"></a>

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

<a id="canonical-2301131030223031-1323231223101011-2032333102222201-1332322121310011-1312313131101321-0322012321012122-0121101112003303-1213202013312010"></a>

## Direct properties — domain / 300020013311 / 3

<a id="canonical-1310110101011331-2331001133333232-3100303112232130-1102302320000132-2212000012032010-1310100101210000-3002022232233210-3020202003123031"></a>

<a id="canonical-0221201101303233-1313011201331302-3302200300112032-0320011333231033-3322122321100203-2212300122131331-3222133001200200-3323200131012330"></a>

## exact_value property — domain / 300020013311 / 4

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

<a id="canonical-0021021201131212-3221220211120113-2211210120223221-0013031202301311-3301013011303000-3000130310200001-0323112321110203-0312112020020100"></a>

<a id="canonical-2112221310331000-1000011230022310-1333212111221020-2111123310313013-0312123311320030-0022121013102212-2102030313002012-0112231230121123"></a>

## regex_value property — domain / 300020013311 / 5

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

<a id="canonical-2230121301003100-0033230201321232-1130131202133311-2123312021122010-3032230203033202-1322301103123310-2002301200031222-3330031112221032"></a>

<a id="canonical-0220301123210321-0223130302000210-2203022301023333-2232101300202033-0311132322110000-0130132212133111-1310231021010222-1213313011233120"></a>

## suffix_value property — domain / 300020013311 / 6

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

<a id="canonical-2121213300132221-3202103113022101-3212331031000331-3202200023222312-2120112221033300-1002133010203203-0201123020131330-0033113303200330"></a>

## Next pages — domain / 300020013311 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3310223023102312-0113032103112001-3230233230132310-2323022123331023-2220120001200003-1323213303100030-0011322010230312-2020101203300032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033021121301201-1121201322223330-2201133021123231-3213003222021123-2011110332323120-3011112313320000-2103101002001201-2033222323130321"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata — metadata / 313121130021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-3012113233000222-1222201100001220-0210201103021012-3022332000210232-1112112310222122-3110132332022233-2233023101101010-0023233212000032"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1320311231123031-1300113212332211-2302222311232331-2202233202031301-1013101303300200-1010111000313000-1230000032230123-3013123200302023"></a>

## Direct properties — metadata / 313121130021 / 3

<a id="canonical-2101022300032300-0223133002123111-2310332012322230-1103230033222310-0320001320332110-0322213100112003-0322322232031211-2132222111200332"></a>

<a id="canonical-0030232100300020-3213330321111231-2132110302003022-2213011221313002-0112301211131013-0012013323112122-2011000301300231-0211030321330211"></a>

## description_spec property — metadata / 313121130021 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3031102032131232-2031212303222133-0302001121310201-1323022110101132-3022222100102232-3311200322002200-1010133301312032-1333022211321200"></a>

<a id="canonical-2023013021221123-2110213113212120-3003202110221003-3321230111222103-3303213231102221-3000323020300123-1213301230131230-1330311210231320"></a>

## name property — metadata / 313121130021 / 5

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

<a id="canonical-2123323101123212-2211210101222211-3200033201012110-3310112213303130-1132030322231110-3130230223233312-0202310212012132-1221300312113200"></a>

## Next pages — metadata / 313121130021 / 6

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3223123302031303-2121331213002022-3030100231303021-0111222033122202-3100231200220022-0023302321101222-1031231113110012-0202121330000121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001121231303101-2113122131221031-2133021131122001-1212031021232102-0123010112212101-2122003103131103-1111023330001022-0332302130310000"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path — path / 102323310110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-013.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0213232031122322-1002003110020101-0030102311330213-2030133122212313-0302031301321103-3003310010220112-2032001302310330-2032012003113100"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2302031222323301-3300302210203311-2010331312302330-3032331020303010-1312303102233323-3230102010323213-0132111322133331-0020200132123221"></a>

## Direct properties — path / 102323310110 / 3

<a id="canonical-1231013221223200-0331220022103123-0003213000212211-3121302102022000-2232303203012011-3002101033220101-2323312021023330-0311122210013102"></a>

<a id="canonical-0310312121132233-0213032222133021-1203303013221010-1131222130102232-1313302110100301-1102323031301221-3022211323010311-2233133110321331"></a>

## path property — path / 102323310110 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3000210323120332-0033232132001100-1032012103032222-1331231002131011-3010011233200011-2113231212002320-2231000020332221-3122031301120312"></a>

<a id="canonical-0002313303331223-0230030211332113-0232021302330231-1313332320313203-2102310122131223-0020101010033020-1121122210221232-1232120202202202"></a>

## prefix property — path / 102323310110 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1212123302030320-1313003303131300-0321132021220331-1233222200332020-2223203311223132-0002131022331322-3203311323203000-3031112010103101"></a>

<a id="canonical-1231303001322131-0200320310303100-1110213022103020-1100013321210001-2111033210020010-2212212110310202-1323321332310220-0130302303130031"></a>

## regex property — path / 102323310110 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-3302222231320023-3110030200211320-3121023330112313-2233300133012303-0322211110130101-3112032212200203-2331312322201103-1221231011311200"></a>

## Next pages — path / 102323310110 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101331131000210-0222213203330033-1102313133001301-2123231022223233-0202011110222112-0200322031210223-3212322000023011-0221233223132220"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules — js_insertion_rules / 002001002301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

<a id="canonical-1322323020010130-0212021132010313-0203031102023321-1311322102122130-0233010103232221-2332123332312232-0001122312012112-0001012211033233"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2222013222113323-1132121202213221-0203212330022011-1100322200323211-0033202023222122-1122030121331301-1203113120311113-2222231100111000"></a>

## Direct properties — js_insertion_rules / 002001002301 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130): complete subsection reference.

<a id="canonical-1201023312113102-2013211231201122-1202010103233330-3303133032322120-1033130032312231-2100333130100213-3313331323120333-0122101201321120"></a>

## Next pages — js_insertion_rules / 002001002301 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101303033202122-3323020300211022-3031332221011301-0330101222310102-0021013013120033-2031003101133223-3130303313011301-3210321200300012"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list — exclude_list / 101333123001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list

<a id="canonical-1202330301101212-1220320202301032-0002330121110100-2102333200001021-0112231103230013-0021031031011312-1320023100130202-3032011323201211"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2012200133311122-3201021231332312-0200333321012321-2022303121332102-2312030330223312-3313013223022220-0111203031300023-1011312220121202"></a>

## Direct properties — exclude_list / 101333123001 / 3

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3203220101033033-1221321022123231-2123112331212230-2200032312313001-1212020211212032-3101123211021002-0033313231120103-3330310313322023): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-2210202110212101-2031110200321121-2223102312220223-3120230131303210-3100001322021321-0230003233113232-3222101210322133-0210211332231000): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-0133330222010013-0230302203212302-2233003031002313-3222111301010011-0232212012030221-2113101311122233-0031332232100011-2222300211322211): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-2301312002033012-2233003003110011-0110200300122230-1301130332320001-2012333232213321-1010131331333022-1213332320321302-0333133232103312): complete subsection reference.

<a id="canonical-2111220333302210-2001323020311001-2023122233221022-0133213132311331-0321213122031310-1122132223100122-0112102230023002-2122230002131122"></a>

## Next pages — exclude_list / 101333123001 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3203220101033033-1221321022123231-2123112331212230-2200032312313001-1212020211212032-3101123211021002-0033313231120103-3330310313322023)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-013.md#canonical-2210202110212101-2031110200321121-2223102312220223-3120230131303210-3100001322021321-0230003233113232-3222101210322133-0210211332231000)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-013.md#canonical-0133330222010013-0230302203212302-2233003031002313-3222111301010011-0232212012030221-2113101311122233-0031332232100011-2222300211322211)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-013.md#canonical-2301312002033012-2233003003110011-0110200300122230-1301130332320001-2012333232213321-1010131331333022-1213332320321302-0333133232103312)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3203220101033033-1221321022123231-2123112331212230-2200032312313001-1212020211212032-3101123211021002-0033313231120103-3330310313322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221213021200031-1011031123100232-0112033002022210-3001321323211223-0020132302111033-2310201220102332-1230021232213123-1032232120031133"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain — any_domain / 013131132303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3210012210212101-0300212232101021-0301231121200321-2233101220330020-1300003212021322-1100113321211301-2032221130122320-0101022003332032"></a>

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

<a id="canonical-2322301023033000-3222123120312000-0212310330312031-1302033103233123-2132331023010320-1321101313213002-0023130011202332-0132230111320022"></a>

## Direct properties — any_domain / 013131132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310330321311022-2000331221011111-3122210310312033-1022203031121120-3232233311103030-2202202100003022-3301012013320123-2213201132121031"></a>

## Next pages — any_domain / 013131132303 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2210202110212101-2031110200321121-2223102312220223-3120230131303210-3100001322021321-0230003233113232-3222101210322133-0210211332231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000300113110302-0021201201321112-0113302132010210-1010031012222032-3013320303131003-3313002301110033-2320012333320310-2311201312332101"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain — domain / 111302332101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain

<a id="canonical-1330333303213002-3223022111323322-0321031321331232-1321310302003200-1303333113210001-1232013001012302-2222021022132221-3311300100321330"></a>

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

<a id="canonical-3133032032201231-0210222021223303-2331112323001202-0010030133100332-3322330132201123-3300322213211223-3100312223331032-2003002223121020"></a>

## Direct properties — domain / 111302332101 / 3

<a id="canonical-0121030313110330-1020002122300300-1222012211020323-2221102203222322-1221021111332210-2213030022013223-0321113011010000-2012031313221032"></a>

<a id="canonical-0233033200110333-1100221213020131-0211003023131031-3032032311302121-1010112023030010-0102200002331001-0220300300031312-3012300233233211"></a>

## exact_value property — domain / 111302332101 / 4

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

<a id="canonical-3313211103312222-3010303131230033-2030112132323330-0223303333010223-3213112213220310-2212010133322210-1003321200330121-0100002020003112"></a>

<a id="canonical-1202223111321133-3300010011130212-3323123120230301-2021233102211323-0032310302330112-1200303220303200-1031303133331011-3110312133332030"></a>

## regex_value property — domain / 111302332101 / 5

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

<a id="canonical-0201300113101211-2321212333223131-3112231100120133-2110331231332233-3101010001122020-3100231032303213-1021323001033031-2032231002212230"></a>

<a id="canonical-0222202003323110-3111213011103213-1120330223113003-0002001233333000-3110110113233223-2000311112321323-3000031002223001-3133301221301113"></a>

## suffix_value property — domain / 111302332101 / 6

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

<a id="canonical-1221013013133322-0021022010120200-3020222233312311-3332222203132133-2133311312223300-2021223011201122-2201003113110223-1021103321332001"></a>

## Next pages — domain / 111302332101 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0133330222010013-0230302203212302-2233003031002313-3222111301010011-0232212012030221-2113101311122233-0031332232100011-2222300211322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301301130222221-0031120032311233-0023001102033113-0101001003013331-3320131301002200-0331331221031012-1302122010312203-0310222322200101"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata — metadata / 232231313330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata

<a id="canonical-1021222212202100-1022100213010001-1110013301120122-2211132030303331-2120300020112230-2112322113131123-3120213313132123-3321211112011300"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2222002122322031-0320133333321110-2322031330111310-1132103213110210-2022322332310010-2322131002020330-0213020102201220-3301111212011312"></a>

## Direct properties — metadata / 232231313330 / 3

<a id="canonical-0332010132033220-1033011232203031-2103100330300110-3201312032221001-2131133023002131-3300131132311233-2133231221102112-0131200130233132"></a>

<a id="canonical-2110102102030002-1232303003303301-3232112023311202-1202121033211110-1220102213301303-1123200132331000-2011100332113102-3031120200302331"></a>

## description_spec property — metadata / 232231313330 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3002023322012021-2002022112232202-3020222231133211-0002021231313020-1001202031301223-3111101333013033-1200230322333103-2113021103022032"></a>

<a id="canonical-2332123203121102-2223030210010203-1222130132001031-1303302211313331-0310101310231031-2303203022301033-1210002220210131-2311100031310101"></a>

## name property — metadata / 232231313330 / 5

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

<a id="canonical-2132121101010332-2000023321010113-0231303020210013-1003131322302122-3222033232333010-3312301120220231-0100103020222231-3012030221320233"></a>

## Next pages — metadata / 232231313330 / 6

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2301312002033012-2233003003110011-0110200300122230-1301130332320001-2012333232213321-1010131331333022-1213332320321302-0333133232103312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023330310300202-0202200210001323-3111323033100032-2100333303213112-0212333021003103-0123210011121233-0321001333230210-1221233201322231"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path — path / 330313020212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path

<a id="canonical-2110010211221331-1121132120122313-0231123320211210-0122130320012113-3122110031301202-2330101202122230-2003200110111322-1113121322103321"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300011310133020-1213220302100321-0202203230002231-3202211301302123-1112322110332100-1330031303301101-2230123202021332-1021323102321011"></a>

## Direct properties — path / 330313020212 / 3

<a id="canonical-3221130022221113-2230001301232330-3013332130313200-2232201310230221-2101233202020001-2212122111300112-0022210220202320-1301103301122013"></a>

<a id="canonical-0312103320012101-3021031301311310-3332221112232333-0032313301332131-2233123120111301-1030113200210201-1130301201323113-3002001022221320"></a>

## path property — path / 330313020212 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-2332303112312010-1332202301303001-2001212232031201-2111320012332012-0030323101020221-3101123133100221-0310333223033013-2000201312321023"></a>

<a id="canonical-3012300010313230-2301000023013021-3200032220031110-1120203211002211-0303331212203100-3002220331330210-0021200031302321-0013021310020123"></a>

## prefix property — path / 330313020212 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2012130221111232-2220212320323022-0201111301031233-1310102101320012-2003130022132223-1232120202000110-3320112202001331-3122200310333133"></a>

<a id="canonical-3231133221122010-0231003000122330-0310322003023213-3323332011121233-1233222012331031-0121201323103210-2131033012200223-2032122233223321"></a>

## regex property — path / 330313020212 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-1332113301003220-2122333312010013-0122233133303213-3221201032220001-2101202032320132-1113022010231123-3111331100131120-0101100321320201"></a>

## Next pages — path / 330313020212 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130010030230120-0201033123300012-0202202302112212-0120313001021133-0230220310223131-1002011133221130-0121101322002223-1310322131332220"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules — rules / 210102003221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules

<a id="canonical-2111232211320012-0121011021031310-3003001111013212-0233323122133211-1313233003100322-3210022221232021-1332322333012331-2302213031303002"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302230221130312-2321002110302230-0333230121003233-0302211022103103-3233332132231220-3201032312113230-1203030122232300-0210022110330212"></a>

## Direct properties — rules / 210102003221 / 3

- [any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3302323211110130-0110203013301300-3010220320201231-0331123130023032-3201333011023011-2000220011201021-1111210110312033-2201010102300002): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-013.md#canonical-3210233113300130-0110011020100103-3022032333031122-2030033023200330-1132020102012221-0300221030131132-2302312020000213-2302130212021310): complete subsection reference.

<a id="canonical-2332011103221221-0221011211120120-3221021112332030-0003110301323122-3320331330332030-2322100031212013-3230012130103103-2022001333112112"></a>

<a id="canonical-0220110102233301-1011202032210230-2310310322220221-2003112101102213-1303303202211331-3311022031132302-0003320311011132-1101312103111211"></a>

## javascript_location property — rules / 210102003221 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

- [metadata](resources--http_loadbalancer--reference--group-013.md#canonical-0203223320013121-3331011111301320-0022103222030001-2121021302203113-2233320132221203-1211113232203231-0023203103000333-0130021320133222): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-0222233300033223-0331203312001231-1021001121003333-1013320231220321-2323211102211312-2321101311030333-0213323310233321-0022010023322101): complete subsection reference.

<a id="canonical-1033322033103000-1033112112201033-1332131010232213-0022311312330120-2030210232130302-3223020132232032-3032313331233111-0102032022030101"></a>

## Next pages — rules / 210102003221 / 5

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-013.md#canonical-3302323211110130-0110203013301300-3010220320201231-0331123130023032-3201333011023011-2000220011201021-1111210110312033-2201010102300002)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-013.md#canonical-3210233113300130-0110011020100103-3022032333031122-2030033023200330-1132020102012221-0300221030131132-2302312020000213-2302130212021310)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-013.md#canonical-0203223320013121-3331011111301320-0022103222030001-2121021302203113-2233320132221203-1211113232203231-0023203103000333-0130021320133222)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-013.md#canonical-0222233300033223-0331203312001231-1021001121003333-1013320231220321-2323211102211312-2321101311030333-0213323310233321-0022010023322101)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3302323211110130-0110203013301300-3010220320201231-0331123130023032-3201333011023011-2000220011201021-1111210110312033-2201010102300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230321210033333-3200012232222021-2032303221210102-2313113213010231-2210323203213113-2003332013030011-1102031030333302-2030020212123313"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain — any_domain / 110133212033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain

<a id="canonical-1321133220010122-2112123203100123-0321321322231113-1330111201310000-0133211313201123-3322210230320330-2332202322101002-2000301113122113"></a>

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

<a id="canonical-3100010022002312-2322000222321000-1312010000130021-2313302200031120-3232331130021320-3030123200100321-0303300001023331-3322222010030222"></a>

## Direct properties — any_domain / 110133212033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313001022133101-3133123102100000-3023003032312221-2022323131211000-2301011003232111-3131120323302213-3310233313320232-0212332010232111"></a>

## Next pages — any_domain / 110133212033 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210233113300130-0110011020100103-3022032333031122-2030033023200330-1132020102012221-0300221030131132-2302312020000213-2302130212021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220202120001011-3203313002211322-2232011033100313-3233130331011221-0321020332011011-2113010012300102-3000330033032310-2320203330002321"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain — domain / 231023201312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain

<a id="canonical-2030030301022001-1220112212022232-3332330322101001-3332010030021010-2333302131303013-1110230021110320-2000121010012312-2221201333133002"></a>

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

<a id="canonical-2320230021131320-3312203021000011-3311003303000320-3001230112020320-1221021232001222-1010000203222301-2132333033011200-1300123001210312"></a>

## Direct properties — domain / 231023201312 / 3

<a id="canonical-2321322123203311-2331310131213011-0232323122121201-0010221332123111-1201031212233023-2333323122221102-1313122312330303-0203120001122320"></a>

<a id="canonical-3231221023111130-2111300322110302-3323301202330012-3203312332312021-1023310220100032-1203303123311232-1112232122131232-2330002201233231"></a>

## exact_value property — domain / 231023201312 / 4

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

<a id="canonical-2321023313333310-3333100000333201-2222010113123102-1323302120100021-2333010212303331-1002120022221133-3100103322020301-2312200203333020"></a>

<a id="canonical-0102312302310201-1100122222311322-3203322002322003-0311323231302012-3333222123123232-2002233301013010-0010321011011323-3002130103132203"></a>

## regex_value property — domain / 231023201312 / 5

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

<a id="canonical-1020013022103212-3223332102230101-0323000312110113-1001003220211030-2102001300031012-0132303001231311-3033120001030130-1302332210322201"></a>

<a id="canonical-1230231232011223-3212322203031332-3112330311301121-0123111322213332-1021000233302221-2033221211333233-2121231002310031-0102021102102221"></a>

## suffix_value property — domain / 231023201312 / 6

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

<a id="canonical-2010232120023132-1023001002030130-1111330303001120-2213030102000232-3322231001300102-3032310022112210-3301121322102220-3321331331203003"></a>

## Next pages — domain / 231023201312 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203223320013121-3331011111301320-0022103222030001-2121021302203113-2233320132221203-1211113232203231-0023203103000333-0130021320133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003303013031003-1010032322330323-0112101011201223-2111322103200112-1330101212012003-0212020200221332-0100122113021011-2112233201103133"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata — metadata / 011111023200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata

<a id="canonical-1033311100102013-3313221313033310-2231312112200301-1111212020112313-2000332100330201-3022110303121330-2000211012313221-2011221013233103"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2222323203032111-1301011231131231-0323110313032222-0200133111120112-3312031321032331-3033010103303310-2223022030113013-2221030021320111"></a>

## Direct properties — metadata / 011111023200 / 3

<a id="canonical-3232030033201202-0310002331011203-0002233213032032-3013120220222003-3212021023303312-2102103023332133-0131323201223123-0032323123002101"></a>

<a id="canonical-1300030302021010-0320311021133022-0112103020313230-1331101122331313-3031003013011112-1022221322332221-3302001103333130-3032010233310221"></a>

## description_spec property — metadata / 011111023200 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3010233320133002-0312023131021012-3112222312320120-3313232121030200-2021313302201322-3031123133101112-1333100000300302-3132122001001220"></a>

<a id="canonical-1302002320312301-2333023121010323-1233322122301002-2201002302010330-0300203322320110-1201133312113213-2011100213222013-0211031120123031"></a>

## name property — metadata / 011111023200 / 5

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

<a id="canonical-0110020331131100-3201332312023201-2011103322021220-1131131303330030-1103212313310213-3112022332030210-0332022201133030-3121331010123220"></a>

## Next pages — metadata / 011111023200 / 6

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222233300033223-0331203312001231-1021001121003333-1013320231220321-2323211102211312-2321101311030333-0213323310233321-0022010023322101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111122233202121-3131332103320331-1322023320313312-1033333223013221-0102010101002110-1121302021122132-3220023203301312-2010011223113100"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path — path / 320032023311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path

<a id="canonical-1221030210321211-0331302030330331-3300102113321001-1321310312022211-0301320132203031-3123210332302132-3002231123331332-1020323212103330"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3111302101130200-3333332033301020-1311331232022201-3223111001002231-3031103101130021-0213012102233221-2121300001102120-0301203302022023"></a>

## Direct properties — path / 320032023311 / 3

<a id="canonical-3202113103023002-2332202323030323-3110112332231303-1200030121001201-2011310131030321-0302321012121213-3301130132311320-1022230200030200"></a>

<a id="canonical-2130232202131133-1111112221202110-3303022131112000-3000310111223322-0000123201313223-2330300102303001-3212130220202123-0102210123022300"></a>

## path property — path / 320032023311 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3303202330102231-2032103230311303-2212130231231100-1212210332310100-1133022300000020-2312311013101301-0132121133230113-0202033232112220"></a>

<a id="canonical-1211033301322013-0032330331331220-3130013333301303-0122021122220013-3132222122310301-2313121231001332-0113030002201303-0330110001313030"></a>

## prefix property — path / 320032023311 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1312031212002001-3301330231230111-1021222303120321-0002303311102331-2002210121202222-0312133333022230-2010010001330221-2102330022302021"></a>

<a id="canonical-0300200103230010-2101331033221000-0122000232002302-1100330333222313-1303233023221021-0003310123121021-2112212000103121-1312211020031203"></a>

## regex property — path / 320032023311 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-2301221221132223-3310233122211330-2122330110110121-0022101133112101-0110002222022013-1221032102210030-3210332103102230-2122201230103210"></a>

## Next pages — path / 320032023311 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-013.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3113113321230312-1011132221021022-2222200122232321-1121313322131023-3010032012223210-0112301112312333-2003213221232012-3022332221010201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012133030103323-0222123320333322-1002020111112000-1201103010033312-2121203000201012-0320230331121302-3011112323230033-0311011111101100"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile — mobile / 012101111121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.mobile

<a id="canonical-0103132030310130-0213012332313021-1120313223301212-1223332013211033-0310310010021311-0033102202320013-0321131020320302-2021022333103322"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011023302322323-3113112012031131-3133310321302232-0013100010113002-3113111230031210-3212222123221210-3131321300020120-2221233302133102"></a>

## Direct properties — mobile / 012101111121 / 3

<a id="canonical-2033002331311212-3102321010021111-1011133002302310-3021213102311132-1112312210112220-0030330112000003-3113122230121021-2113112020301232"></a>

<a id="canonical-2131103100331003-0131310033122102-0123121032221000-2033031310210123-1321123310133300-1133203110221011-2311122001121330-0022203010222220"></a>

## name property — mobile / 012101111121 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1232113220032222-1331003211210033-3330231222233212-0033112123202331-0110031133131302-3310323131123312-3300101233003200-0210031032131130"></a>

<a id="canonical-2222200211110322-1133032220200332-3023302123110223-3023230011031021-2223231220233120-0301102122122012-3332121011313201-2102023210111022"></a>

## namespace property — mobile / 012101111121 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3031000220333222-2032302201313313-2321302001030020-1123302113222231-1322321200032200-1112321000213103-1021000032000011-3203103223313302"></a>

<a id="canonical-1202332123312300-1323213202322232-2303021320010130-0313212020311112-3313021212110123-1030212011223033-0020013330023123-3120112020303030"></a>

## tenant property — mobile / 012101111121 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0113112030222220-1331310322121131-0021132122320001-1221112330110221-3202201121130030-2010203012022320-2212232223101033-1020013202322331"></a>

## Next pages — mobile / 012101111121 / 7

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332001120133120-3120103102320013-1031101211110202-0001131010222120-3102320111233302-0210222230201210-3301301222130231-0211220210233330"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config — mobile_sdk_config / 220320210110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config

<a id="canonical-0010120112031102-3233302001013301-2313122201220120-1322233310310120-2230331302012223-2011202111101313-3030333121333111-1232233121200312"></a>

Type: `"object"`. single nested block, Optional.

Mobile Request Identifier Headers. Mobile Request Identifier Headers.

Upstream description:

Mobile Request Identifier Headers.

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

<a id="canonical-2213113131030000-2302313331113212-1020230333113311-3102221011120103-3101313202003003-1311022032303223-3131230321100010-3213113012212202"></a>

## Direct properties — mobile_sdk_config / 220320210110 / 3

- [mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022): complete subsection reference.

<a id="canonical-0320323131300130-3230302220211201-1030331020133103-3313201123303320-1032222011001213-1121310133230301-1233013131233310-0122322011311001"></a>

## Next pages — mobile_sdk_config / 220320210110 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300221030120300-0230330321123011-1303330020100201-2132201320102022-2012321310231112-3222331321331300-2001032110312210-3331020211210312"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier — mobile_identifier / 203132030310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier

<a id="canonical-2101113021320121-3110022033123233-0032111223023232-2210320113012310-2020332331210230-0123000003203012-3312233113101323-3133321203103102"></a>

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

<a id="canonical-0220133321322230-1303212200210303-1130002012200003-1211122012231333-1312330330231310-3311310322200103-1333313312103002-3110023213122300"></a>

## Direct properties — mobile_identifier / 203132030310 / 3

- [headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120): complete subsection reference.

<a id="canonical-0322003131203131-2233323011131020-3203210102230033-2330033300200011-1100132313331213-3022220202123000-3200121222001103-2113110122110011"></a>

## Next pages — mobile_identifier / 203132030310 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001313203033303-1032223122222311-0113320103023111-1331313132220111-0230221111333111-3020122333003223-0000120300202233-1331110112220022"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers — headers / 031210310010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1022121230111131-2323023221232310-0030100322210120-1232031123322231-0112230002101231-2222312231000301-2222132021011230-1003121131001221"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333222231330212-2301311221203311-0320310330322122-0012131111120211-1032113120203311-0313131103122002-3230113203031223-1102311022132200"></a>

## Direct properties — headers / 031210310010 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-013.md#canonical-3131222020202133-1120203030323000-1033223001030230-3013311033123100-0030131120203303-3131313120131012-1003222231012233-1220112131132230): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-013.md#canonical-3032222213131332-3322000020230202-0323332232103122-1111121320101031-3020033200313011-1011031210211310-3310211100330130-1203112102123111): complete subsection reference.

- [item](resources--http_loadbalancer--reference--group-013.md#canonical-2031320001212222-2121021321103321-1203322320003022-2232221122330301-2301113302230013-1000202333210022-1002213013222023-2133003201221102): complete subsection reference.

<a id="canonical-3032322300122221-0311021301110211-3133013020103120-0202202133310312-3102120322012201-2111330233230300-3311313012120030-0321311033022202"></a>

<a id="canonical-0023102022222231-2111230310230203-1112113101231033-3012330223203232-0022220313120230-2010302312010323-1113103202232211-1222010010213322"></a>

## name property — headers / 031210310010 / 4

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0232300110200122-3131220120131113-2301033132323322-3012013011021102-1212022302312321-0231002231131231-0033323001023000-1110220230213321"></a>

## Next pages — headers / 031210310010 / 5

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present](resources--http_loadbalancer--reference--group-013.md#canonical-3131222020202133-1120203030323000-1033223001030230-3013311033123100-0030131120203303-3131313120131012-1003222231012233-1220112131132230)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present](resources--http_loadbalancer--reference--group-013.md#canonical-3032222213131332-3322000020230202-0323332232103122-1111121320101031-3020033200313011-1011031210211310-3310211100330130-1203112102123111)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item](resources--http_loadbalancer--reference--group-013.md#canonical-2031320001212222-2121021321103321-1203322320003022-2232221122330301-2301113302230013-1000202333210022-1002213013222023-2133003201221102)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3131222020202133-1120203030323000-1033223001030230-3013311033123100-0030131120203303-3131313120131012-1003222231012233-1220112131132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102031210102211-0033010120200302-1002233101010311-1221110303221113-2032201022332031-2213223003323113-3133300331020212-3321220002321131"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present — check_not_present / 021200031100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-2301100120031121-3211210313003122-3332122123102113-3231202120100013-1031212230013122-2021012330021030-3233323123111320-0331133111221313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3012012001313210-3220321001231023-0230003213302200-3133221300332230-1131011332202201-1010023312210201-1330113013323012-0323113333330232"></a>

## Direct properties — check_not_present / 021200031100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321112210102101-1132302220112122-2101332000011113-3031200032232211-0002010310221220-3120213221010310-2321020120001320-3110132223003103"></a>

## Next pages — check_not_present / 021200031100 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3032222213131332-3322000020230202-0323332232103122-1111121320101031-3020033200313011-1011031210211310-3310211100330130-1203112102123111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011210333322131-1022211003132330-0003023010013103-3210033233001220-2031020033320312-0330111113323130-2122321333003322-3222203101313033"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present — check_present / 100112200033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-1121230300112211-2002311113201330-1200100203303103-3020112301120213-3001233303222221-2011303203221323-0322311220120002-2333113221012230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0002233133331110-0023332232022002-3312123002320130-3110221020010021-1221211122102002-0012313321323201-1313003000030031-2313100210233331"></a>

## Direct properties — check_present / 100112200033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131003132113011-3022312320112212-3133322202231110-2312301221310203-0021031001130011-1122112121303212-3231303102310130-3121220233033000"></a>

## Next pages — check_present / 100112200033 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2031320001212222-2121021321103321-1203322320003022-2232221122330301-2301113302230013-1000202333210022-1002213013222023-2133003201221102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330302201023213-0320112232111130-3233330220222103-1203222223301011-2130302301000011-1130303021110320-3233010211203003-2230330102211033"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item — item / 122100212122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-013.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-013.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-1103003211132012-0330120213213331-0322000322320020-3123233203222001-3312102112030030-0223323101112221-1023310313301310-0131012220300233"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211131100210101-3031032032032003-0223330200030203-2102312230121013-0133112323202023-2223022230310133-2200200033032130-3232322323010221"></a>

## Direct properties — item / 122100212122 / 3

<a id="canonical-3200033033202021-3013212300121233-1111202232013010-1221230200022230-3201230220313213-1230022300123031-2102022223302111-2231103322022112"></a>

<a id="canonical-2220030003133312-0010231230233033-2030203032223021-1213222122033211-1033311010000310-1100032223201021-1011010113203320-0323331011321011"></a>

## exact_values property — item / 122100212122 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3230123320230333-0313031012232200-0313322121332003-0120200031123110-2313210100033301-1111331123021013-3022130332013001-3202330301221033"></a>

<a id="canonical-2113103323022313-2223311112313332-2011331113011120-1000313223001002-2330101212102312-3320133231022010-3202211013100311-2332222103001221"></a>

## regex_values property — item / 122100212122 / 5

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333031112030320-2313101133320031-1000010123013310-0220030010123211-2321212031230123-0101000203232020-1313102201122210-0023300213122313"></a>

<a id="canonical-3311231222232333-2233031322303230-3011333212031330-3122130030123032-2122220123331303-1213101001330011-0233322222101120-3313102320000221"></a>

## transformers property — item / 122100212122 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2113321000023303-1313313233311113-1131302111302002-2003321120323203-0020321201012121-2130211301310230-1031023221312323-1231113012022020"></a>

## Next pages — item / 122100212122 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-013.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200000311231133-3323200203131332-0220323231120332-1111202111132333-0131021232301123-1131022130213101-1010110003233102-3213003220120002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122032020103303-1321101122012132-0202001000032322-0000222323212333-0113212032100130-3222131323101113-1321320231111222-2112303122110013"></a>

## bot_defense_advanced_protection.both_web_and_mobile.web — web / 201012002323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.web

<a id="canonical-1323222133101321-1201302222222323-0320011321302213-0223020201232322-3331323001321030-0033222020100003-0122002113001313-3103000023312331"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
web {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310121030200102-1002312000302210-2110021013112133-2013021310111031-0230303113020201-3123323322100201-1311331212211010-3131233003010030"></a>

## Direct properties — web / 201012002323 / 3

<a id="canonical-0003330231232031-0021200100322001-2000203023310301-0000011103102122-1230231013221322-3311000023031311-0233323011132333-3210112102033331"></a>

<a id="canonical-0002001033121100-1332002102313320-3322223311301021-1221110023222001-0333221020023100-1313130301221030-1310001320222231-3311122313233110"></a>

## name property — web / 201012002323 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2301321103020023-3122002030230222-2111223012111220-2312100032131130-2331220131302323-1022312312221312-2132010010203033-0122321011231013"></a>

<a id="canonical-1211313202101320-3112120201202102-1010130313302221-2002232030331202-3320130223023103-3001013012023103-1330132223012133-0120021111012310"></a>

## namespace property — web / 201012002323 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1212323323313233-3130313020110121-0220021120121030-0130230310202123-1200003121200110-0100220123120021-0132113103331113-0200010102221223"></a>

<a id="canonical-0233333332303112-2232323331212011-2032310201100011-0000030012212122-0300112333122022-0133130303331222-1202222113101130-2233102011112222"></a>

## tenant property — web / 201012002323 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0221121020023230-0100303332321031-1302220131333020-3333332203112233-3020322011113310-3022202312010220-2121311233331212-3322211202100113"></a>

## Next pages — web / 201012002323 / 7

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031020132310000-1021320301033000-1020013133001111-2321112233221030-3111332023032000-2300030203220231-2112133121111213-2310300102222013"></a>

## bot_defense_advanced_protection.mobile_only — mobile_only / 130101130313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.mobile_only

<a id="canonical-1212113101121233-0310101112200203-1332201203302303-0132333133221113-3303211010210110-1002213112031200-1131010100131313-1331310321312020"></a>

Type: `"object"`. single nested block, Optional.

Mobile. Mobile only configuration.

Upstream description:

Mobile only configuration.

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
mobile_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001211310201333-1232321202130020-1130102032120022-3030221000221222-0020132101233000-2130032111203301-1321330021002112-3212120123210211"></a>

## Direct properties — mobile_only / 130101130313 / 3

- [mobile](resources--http_loadbalancer--reference--group-013.md#canonical-2302013123321210-2310223331012130-2033001202111121-2211013221120030-0200121213132111-0120112120311120-0233312220213230-0032133101312331): complete subsection reference.

<a id="canonical-0113313111332121-0303201111021100-2330211110110021-1120303331002131-3321000011213122-0033322320130111-0030321012233010-3211210220211001"></a>

## Next pages — mobile_only / 130101130313 / 4

- [bot_defense_advanced_protection.mobile_only.mobile](resources--http_loadbalancer--reference--group-013.md#canonical-2302013123321210-2310223331012130-2033001202111121-2211013221120030-0200121213132111-0120112120311120-0233312220213230-0032133101312331)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2302013123321210-2310223331012130-2033001202111121-2211013221120030-0200121213132111-0120112120311120-0233312220213230-0032133101312331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
