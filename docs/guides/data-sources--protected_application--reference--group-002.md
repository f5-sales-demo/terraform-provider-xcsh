---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-001.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- Cloudflare.protected_endpoints.web_client.continue

<a id="canonical-3033223111030212-1323010303013003-3200331221323221-1101033300312210-3302233103010013-2232211211310000-2001200132311213-1320131123112022"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

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

<a id="canonical-1030130012212212-3300220112332102-0300233033013033-0103213030030300-1001032210231200-0331123012230101-1323210101110130-1231221133301213"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client.continue`

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-2213203203331103-0010213030021002-0332213232103003-0032031231103300-2332212202213301-1210101311103020-1131112102012300-3133000302020121): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-1002012310211100-0032321321000012-0100330113302020-2020000201222130-1301013203020111-2131321133103320-1213310120130211-2133333110123312): complete subsection reference.

<a id="canonical-2213203203331103-0010213030021002-0332213232103003-0032031231103300-2332212202213301-1210101311103020-1131112102012300-3133000302020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-001.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- Cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-2302222211301220-3031310102131031-1103310221031332-1130212312222213-0333223112031333-0023310202113310-0223233013112121-0320132230303001"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002012310211100-0032321321000012-0100330113302020-2020000201222130-1301013203020111-2131321133103320-1213310120130211-2133333110123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.continue.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-001.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- Cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-3222300131021023-1031102203233311-3100323311330011-0330113023203022-2323130203313302-0033233220130131-2002201000313131-3222223120023013"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123313211310020-0123021022331131-2031002202001321-0330103123221201-0213010332000230-2332032020103022-1002323213011322-0223022223202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_client.redirect` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-001.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- Cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-3213203111103102-1320033023102031-3130012122012100-1301232221020200-2320012310200010-0213123101313001-3120013103121110-0122031033012130"></a>

Type: `"single"`. Computed.

Redirect. Redirect.

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

<a id="canonical-1201302122001233-1112013301322023-2210031010000313-3323331013202033-0323003233210131-2101113011100322-2023110323310201-1000311311032011"></a>

### Direct properties for `cloudflare.protected_endpoints.web_client.redirect`

<a id="canonical-3003312021332122-3331321302030003-3302303013222020-1011310110030003-2010023212102222-2313202123332022-2130023322111230-1131222332222120"></a>

#### `cloudflare.protected_endpoints.web_client.redirect.location` property

Type: `"string"`. Computed.

Location. URI location for redirect response.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0331233013321231-2230122321133001-0221100133213321-3331232020031102-2331322023100113-0031000133013333-1110003102001203-2030320213332032"></a>

<a id="canonical-2130133102012300-0220103031103222-2302312003032310-1302330312201012-2113302320320030-0331310222023212-0011031101120011-0310010102230210"></a>

#### `cloudflare.protected_endpoints.web_client.redirect.status` property

Type: `"string"`. Computed.

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

<a id="canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.web_mobile_client

<a id="canonical-2112112301200233-2231312101110100-3002121230330310-3230312132111031-0333120023323121-2332301222310030-0212132023212001-2221112130100311"></a>

Type: `"single"`. Computed.

Web and Mobile client configuration OPTIONS.

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

<a id="canonical-0212021201113320-0100300313320202-3000120111220020-3303231110121012-2133113312011300-3130222112100312-2312313113031212-3332113210000202"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client`

- [block_mobile](data-sources--protected_application--reference--group-002.md#canonical-3220122212111120-1300023312301010-1003202002102300-2312130233321233-1010133123111001-3321020311230133-1323213303332323-2132033133203311): complete subsection reference.

- [block_web](data-sources--protected_application--reference--group-002.md#canonical-3220320230212100-1332113211301131-2123313223101100-3231203230312031-0213000201330332-1300210203112322-0230020201311020-2012201300100101): complete subsection reference.

- [continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323): complete subsection reference.

- [continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020): complete subsection reference.

- [redirect_web](data-sources--protected_application--reference--group-002.md#canonical-3222312311333212-3303212132003231-3202021020011322-2012303002330010-2333223203012313-3201131232122332-0022132221221222-1230130310103133): complete subsection reference.

<a id="canonical-3220122212111120-1300023312301010-1003202002102300-2312130233321233-1010133123111001-3321020311230133-1323213303332323-2132033133203311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.block_mobile` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- Cloudflare.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-2001020110002110-3201232002303102-2011333110113121-1123301032220230-1330100230120122-3333001001110132-1120023322210212-2301020011023121"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0130103020112311-0322023011102323-3223103123021201-1123023020012300-2112032322223222-2330223323330101-2230330003220301-3021300031222320"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.block_mobile`

<a id="canonical-3322102211021001-3011132022130031-1320020313121333-2130020131121021-2030323310213112-0113010010233203-0022123103012212-3013100032123311"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` property

Type: `"string"`. Computed.

Body. Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-3003131003003022-0100012310020321-2123302232130010-1110321132131011-3022123213331321-2303301312003021-1212010122322110-0131210302220310"></a>

<a id="canonical-2330303110203010-0220031203102323-0233113230000122-2210111010130113-0010110230121301-0121020311223103-3132302132121210-1200200221010110"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` property

Type: `"string"`. Computed.

Content type to use in a block response.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1312131222031321-3310001230202232-3012101132320322-1332002330033212-0023200123002102-1201333331333303-3121300030331031-3021013300210301"></a>

<a id="canonical-3331211120302021-0011030310322311-2322112320202100-3013321133200131-3231102120023202-1303130221321221-2331032031003133-3032032102113312"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` property

Type: `"string"`. Computed.

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

<a id="canonical-3220320230212100-1332113211301131-2123313223101100-3231203230312031-0213000201330332-1300210203112322-0230020201311020-2012201300100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.block_web` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- Cloudflare.protected_endpoints.web_mobile_client.block_web

<a id="canonical-2113232300131302-2003103232210213-3013023012201131-1313102222300031-0201303011031022-3223232033332021-2221022110001312-0332231113300023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1203321103100102-0102012031000220-3013101022232112-2033131322123011-2223210102032323-0021000122333000-1222100122202132-2013321213232221"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.block_web`

<a id="canonical-2100100030101302-3220201313312030-1233110231033130-1113032010320000-1110211310131123-3112132003310132-3231010232203000-3231213223323300"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.body` property

Type: `"string"`. Computed.

Body. Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0023123233002320-3033113200013000-3330312310120330-1133230121010021-1231110310010011-0213200332201310-2013210202020101-0110303103221103"></a>

<a id="canonical-0312221222023123-1303302223032003-2211232000203220-0332200233031312-2302221301202000-1232103300012230-0232202202120122-2332112121321311"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` property

Type: `"string"`. Computed.

Content type to use in a block response.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1223111310332000-0222331210231123-1301010200232330-0212030300020320-3133323122302200-0120331211310223-2131130102122023-0102120212331103"></a>

<a id="canonical-2330112230010223-3220112230110211-1020011202302231-1133020120112323-3303130103133212-3120302121120002-0330330200330200-2302112323211110"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.block_web.status` property

Type: `"string"`. Computed.

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

<a id="canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-0221210010023111-2231022330022130-1101303212302322-3033231332230123-2100210110010111-2001100023011022-3121113032120002-3230200001223103"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

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

<a id="canonical-2113123320111232-3023221021300031-1311221020333001-2021213113222110-2212300323323000-1132233101103230-3103113222003131-3303003210202022"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.continue_mobile`

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-1233302023031331-3313322001003313-0223303201001030-0321102021222213-1220200011320203-3213101111102100-0200023113001131-3131311213320301): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-0131032231132202-0203313030032310-0301132220330310-1330103313321312-3232222300012032-0000011302312310-3122112212202231-1212313120131111): complete subsection reference.

<a id="canonical-1233302023031331-3313322001003313-0223303201001030-0321102021222213-1220200011320203-3213101111102100-0200023113001131-3131311213320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-0331302300003201-2010200122202231-0110030123230012-3332310023113300-3313020212223222-0022233113302023-1230331122310022-3312323333010112"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131032231132202-0203313030032310-0301132220330310-1330103313321312-3232222300012032-0000011302312310-3122112212202231-1212313120131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-0320302001212000-2003310030032222-0003112231202312-0030102103232320-3022001332322222-1121232202330002-0020100012312213-2133211103230020"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-1220113120203121-0321130233111033-0200011103020021-2231333210202112-0212000020112012-1300200223321030-2000210110000023-2322021010031131"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

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

<a id="canonical-1132103101222012-2221012220310312-3232212001210112-1023101103000231-3123203333300330-0331203221023230-0312212013133331-1021011333301131"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.continue_web`

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-0321031302023121-0330203232021023-0011313001011321-0123011332010311-2130333230112000-1120132330322323-2111020301010210-1031202222003322): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-0203121022120332-2120330133103212-3312002013111301-3033230300110210-1113130111031202-1231103123312031-1312322201301223-0121222312021200): complete subsection reference.

<a id="canonical-0321031302023121-0330203232021023-0011313001011321-0123011332010311-2130333230112000-1120132330322323-2111020301010210-1031202222003322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-1303313102200013-0222132122330330-2022330302031311-3111023022320113-0011021303030010-0323011332113210-0220202300303121-1213102112133200"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203121022120332-2120330133103212-3312002013111301-3033230300110210-1113130111031202-1231103123312031-1312322201301223-0121222312021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-2323311222112121-2113101131132201-1310312131330012-1033210130203302-0220103103320330-3023110113000232-3213023010230200-2010030011002130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222312311333212-3303212132003231-3202021020011322-2012303002330010-2333223203012313-3201131232122332-0022132221221222-1230130310103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.protected_endpoints.web_mobile_client.redirect_web` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- Cloudflare.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-3322310100322132-2011033011100300-2203103213030121-0002023312111320-1323131230112332-0222100030120332-2033002211123313-0130003003220032"></a>

Type: `"single"`. Computed.

Redirect. Redirect.

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

<a id="canonical-1222000001313202-3230102002311112-0113332120313201-1030033212300233-3002332200031122-0331203030132013-0031313302120221-1221232301303102"></a>

### Direct properties for `cloudflare.protected_endpoints.web_mobile_client.redirect_web`

<a id="canonical-0223103100301103-1130220121012220-3333123111231203-3130222111220233-3202000133122030-3201000031033221-3112021230231200-2102230313112110"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` property

Type: `"string"`. Computed.

Location. URI location for redirect response.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1321313331110232-2103303012201332-0232302211121023-2101131131223113-3223230021121011-0332220202320011-0220332011201133-3020210331211233"></a>

<a id="canonical-1302312230112300-1100000333220300-2332301300031113-1222323313131313-2231323001232222-0332310312023212-1312130102033030-3123323303010323"></a>

#### `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` property

Type: `"string"`. Computed.

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

<a id="canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.trusted_clients

<a id="canonical-0133102322111220-3311232001000011-0230021001031010-1022212123321122-1123130233203113-1323330210311020-3110323313222301-1120130113013233"></a>

Type: `"list"`. Computed.

Define your allowlists to skip Bot Defense inference processing.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0131102223001202-3123131222012002-0222232001003011-0233132233111300-1310031223320201-0102022331213012-2030300312013231-0033231100300321"></a>

### Direct properties for `cloudflare.trusted_clients`

- [http_header](data-sources--protected_application--reference--group-002.md#canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121): complete subsection reference.

<a id="canonical-3133011202313222-2231321120132330-2100330312012112-1320100332033113-1102130100303332-1212011011121320-1011331231023333-0320101202301023"></a>

<a id="canonical-1030012011022021-1033130320033101-3003333320113201-1030330100033322-2023131123121200-1220110210013101-2032313012023121-0321001211332000"></a>

#### `cloudflare.trusted_clients.ip_prefix` property

Type: `"string"`. Computed.

Exclusive with \[http\_header\] IP prefix string.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-2002233311030303-0021002120221301-1032331120101232-0032033033031330-1033132033312113-1033003102330001-3101121220133011-0323323122022220): complete subsection reference.

<a id="canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- Cloudflare.trusted_clients.http_header

<a id="canonical-3331130021002200-1121230312011231-0331100323123333-2031103323110322-1223013033300130-2001233130111002-1021201032222331-2001222332221103"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

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

<a id="canonical-0131203233212320-2221303232130012-3012012320113300-3110012103120123-3010331012230200-1122310130111221-2133130112021223-3313312123002310"></a>

### Direct properties for `cloudflare.trusted_clients.http_header`

- [headers](data-sources--protected_application--reference--group-002.md#canonical-3330100302031312-1203202323230320-1113013130203101-1211103132200221-3202331320032322-2113112202230200-3013112103121221-0003200020202011): complete subsection reference.

<a id="canonical-3330100302031312-1203202323230320-1113013130203101-1211103132200221-3202331320032322-2113112202230200-3013112103121221-0003200020202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121)
- Cloudflare.trusted_clients.http_header.headers

<a id="canonical-3201332011001112-3200213100211003-1320021103130310-0320110301030203-0313313100313313-2113212330310012-0013122201333103-1221233210200310"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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

<a id="canonical-3331202131333100-3232123312331303-3311011322301021-3103111023033300-0002222111332223-0212023010101220-1121232332033330-2123232321013302"></a>

### Direct properties for `cloudflare.trusted_clients.http_header.headers`

<a id="canonical-2031001310010111-2123222002333301-2222322302232033-2130020203033300-1231233123120211-2010032010310102-3330320201203010-1222223021122033"></a>

#### `cloudflare.trusted_clients.http_header.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Header value to match exactly.

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

<a id="canonical-1022303210231330-0300112130233100-1223323233212202-1231311101120110-2010230220033013-3032301100320230-2112211222200021-2123120320231131"></a>

<a id="canonical-0113022021310122-2010200230032012-2211102202332110-3332320312100213-1202031312233002-2120112003022233-0312333233013301-2302023121302103"></a>

#### `cloudflare.trusted_clients.http_header.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

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

<a id="canonical-1111103330312232-3023222333111313-3013222121220132-3232031200211323-0100023312312223-2303022201212313-0113323303021213-2202032023200231"></a>

<a id="canonical-3102023000231221-2112022223222013-1210220023301221-1110301213021232-2000312003133230-0121232323021032-0322111321210203-2301011130333021"></a>

#### `cloudflare.trusted_clients.http_header.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

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

<a id="canonical-2002233311030303-0021002120221301-1032331120101232-0032033033031330-1033132033312113-1033003102330001-3101121220133011-0323323122022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudflare.trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- Cloudflare.trusted_clients.metadata

<a id="canonical-0210013011001033-2103113201332303-1013210202313003-3213000112121213-3232331223112212-1312310233212232-0113002220312202-0020011031110311"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0322310333122212-0230330033032100-3023302111033220-3122323313030120-2330113010123201-0130111331321032-0330331332102333-2120222313011111"></a>

### Direct properties for `cloudflare.trusted_clients.metadata`

<a id="canonical-3102200112001100-0302320111131300-3001023202330231-2101332330202102-2020011030303312-2031202020320100-3201023210031300-3303233103133210"></a>

#### `cloudflare.trusted_clients.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3020101202300232-0103110132023222-0111132213330013-0132002333331313-2010031132201301-0032222003302211-1303130200300312-3002010032032001"></a>

<a id="canonical-2112023310323230-0023330020200330-3312131013032311-3330111230221302-1201113313033300-3101333322201100-2000003232102312-2133301310013033"></a>

#### `cloudflare.trusted_clients.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- CloudFront

<a id="canonical-1001230210132003-1312122030130213-3323113320202103-0212010010230110-2113003310332323-2200230323320021-0021000333232102-1313302122221002"></a>

Type: `"single"`. Computed.

Bot Defense policy configuration for AWS CloudFront.

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

<a id="canonical-0210121203020121-3323221000211201-1020122331112000-1102233333123312-3222112330223321-1002210312333210-0303112111332321-0313011320231100"></a>

### Direct properties for `cloudfront`

- [aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-3003122220113222-1131032101001130-3201033110130300-2110030102303022-0223211112012013-0323212101332322-1210200032213331-0333020213121322): complete subsection reference.

- [aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331): complete subsection reference.

<a id="canonical-1301031001003013-2221130322231013-3100001332203031-2221023330132113-2112310033313122-3203101010111211-0202200123001312-0212211011233021"></a>

<a id="canonical-2022331203023012-0222313132120013-2113131233010031-3210112132101002-0231212113131320-3011322300012231-2303303031230200-1123012101332112"></a>

#### `cloudfront.continue_mitigation_action_hdr` property

Type: `"string"`. Computed.

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1132102323231201-2202010203312032-2222333222221131-3222300022031011-0001232023301231-2101011201233201-0123331210120113-3112213113101110"></a>

<a id="canonical-2100203002113121-3132213020100301-0213111022323322-3221110112331233-3303233332202131-1020130122031120-3312133203122033-0103203002113203"></a>

#### `cloudfront.data_sample` property

Type: `"number"`. Computed.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-0113021203101100-0221102102222310-3021322330123313-0230130301101223-0221130202123110-3222011230311203-0112033023011310-1123112330333313): complete subsection reference.

- [disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-3123120301310311-3222321021000312-2112303010201322-0000033331021033-1001330003203210-0130222013202130-1020213232121010-3131322213232233): complete subsection reference.

- [disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-1013001331001333-3132220323212221-0111011020302121-3313130102030233-0113123332303201-3230122131113103-0332313320101021-0231300013121230): complete subsection reference.

- [js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200): complete subsection reference.

<a id="canonical-2201031311110222-3000013232133121-2021212222103210-3331001320232200-0322120223111212-0312110033300130-1232331330103030-3112102232233011"></a>

<a id="canonical-1232200003020323-3303313202320133-3302230300100213-1310112312120102-0121032103232003-2113022313333120-1112230333130302-0001222122121021"></a>

#### `cloudfront.loglevel` property

Type: `"string"`. Computed.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

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

- [manual_js_insert](data-sources--protected_application--reference--group-003.md#canonical-3123213201022200-1010003020110031-2133001121200133-1103021210123222-3223203013201202-3103020013112213-0232333101021031-3332032331023310): complete subsection reference.

- [mobile_sdk_config](data-sources--protected_application--reference--group-003.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113): complete subsection reference.

- [protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200): complete subsection reference.

<a id="canonical-1301230011221011-0212111310230231-2110331232201203-1133011000012311-0320213032210100-1132332330023202-2121330001021322-0003102311110013"></a>

<a id="canonical-0000222021200010-2011020121021333-2130133200233221-0121023312331123-3033100030123302-3230210113122123-0221312222312301-3003221213310123"></a>

#### `cloudfront.timeout` property

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-3131003001231302-2012000112130122-3232010022011130-2022222222201022-1200232331112220-3022230002311131-3010220330323012-1122321033303022): complete subsection reference.

<a id="canonical-3003122220113222-1131032101001130-3201033110130300-2110030102303022-0223211112012013-0323212101332322-1210200032213331-0333020213121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.aws_configuration_id_selector` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.aws_configuration_id_selector

<a id="canonical-0220011020023322-2010100312322220-1301222102301210-3122103330202030-2203100121303200-0102303311122322-1220201031201110-2132332121312113"></a>

Type: `"single"`. Computed.

Configuration parameter for aws configuration ID selector.

Additional upstream details:

List of CloudFront distributions.

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

<a id="canonical-3310212010033202-3133213232132323-0330020220010123-1302023032000002-1123020023113100-2111210301323333-1322103033330001-1102023313012230"></a>

### Direct properties for `cloudfront.aws_configuration_id_selector`

<a id="canonical-0212130313002203-2330022232023102-2113332210113110-2120102010120031-2022110322333023-1021122301102322-1302013113313201-2320021211213200"></a>

#### `cloudfront.aws_configuration_id_selector.ids` property

Type: `["list", "string"]`. Computed.

Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.

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

<a id="canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.aws_configuration_tag_selector` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.aws_configuration_tag_selector

<a id="canonical-3121212122301010-0021023210313022-2012123131120310-0133230001312122-1302310132011011-1010303032323323-1120020221213303-2031113212020230"></a>

Type: `"single"`. Computed.

Distribution Tag List. CloudFront distribution tag list.

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

<a id="canonical-0012022313300330-0313303030333130-1033000302321330-3030012300023123-1002220110123110-1123112333232120-1102123302212120-0122211232202320"></a>

### Direct properties for `cloudfront.aws_configuration_tag_selector`

<a id="canonical-0310120031133223-3113330201332303-1331023132032323-2023221313012120-1231033102120101-2311223001203122-1130223123310202-1201312301121000"></a>

#### `cloudfront.aws_configuration_tag_selector.tags` property

Type: `["map", "string"]`. Computed.

List contains the CloudFront distribution selection by tags key is an AWS tag name, and the value is
regular expression to match.

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

<a id="canonical-0113021203101100-0221102102222310-3021322330123313-0230130301101223-0221130202123110-3222011230311203-0112033023011310-1123112330333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_aws_configuration` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_aws_configuration

<a id="canonical-2213221110001032-1000233003211022-1110113313110233-3301010111320201-1211101032222320-0223020010132233-2133333002221101-3000103213100111"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123120301310311-3222321021000312-2112303010201322-0000033331021033-1001330003203210-0130222013202130-1020213232121010-3131322213232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_js_insert` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_js_insert

<a id="canonical-3110121231323013-2010102333202133-1223003323100001-1011201130020222-2221313203100310-2010330102021023-3122023000010030-1300121212112323"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013001331001333-3132220323212221-0111011020302121-3313130102030233-0113123332303201-3230122131113103-0332313320101021-0231300013121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_mobile_sdk

<a id="canonical-1202133233323323-0222330321230102-0200201210331330-1220310232323100-0031000031010012-3331001130303213-3022323231212210-1101113200133302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.js_insertion_rules

<a id="canonical-2222221321010301-0332022320311011-2331212321332101-1031010113312012-3210233121102310-0021232023021320-0311120231013130-1213111111103113"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

<a id="canonical-3332232231320112-0133110131130020-0223102022230132-2033310131320220-1033120103120101-0101010120201300-2310313131012210-3302010202001223"></a>

### Direct properties for `cloudfront.js_insertion_rules`

- [exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331): complete subsection reference.

<a id="canonical-2213213031231212-0003203030110002-2220313131211120-2233100001023033-3012103100102332-0001230301020121-3000021330102332-2002210313023222"></a>

<a id="canonical-1212010321311311-0323212220122300-3232212220210210-3022303232000331-3103300033102310-3313210000300030-1330332220320230-0030230010021000"></a>

#### `cloudfront.js_insertion_rules.javascript_location` property

Type: `"string"`. Computed.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

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

<a id="canonical-0000203020212301-2133120200110232-2100323301212223-0032330210113322-0013322013032122-3110110130013100-0011110111120223-3131221231321320"></a>

<a id="canonical-2001323333301203-0022120201030113-1003330330220312-1200233213313333-1112232212203232-2112210010020223-2113110220222320-2113303333303211"></a>

#### `cloudfront.js_insertion_rules.javascript_mode` property

Type: `"string"`. Computed.

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

<a id="canonical-3211210313332010-2210332102000030-0010223111203102-2301100231103132-2201002031100021-1131200212211023-0301231111013133-1233023023232010"></a>

<a id="canonical-0313231111021311-0220131303131102-3210232033220201-3312302123132233-3323233132113122-1021330321310201-1013322131310212-3312001020123022"></a>

#### `cloudfront.js_insertion_rules.js_download_path` property

Type: `"string"`. Computed.

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

- [rules](data-sources--protected_application--reference--group-003.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122): complete subsection reference.

<a id="canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- CloudFront.js_insertion_rules.exclude_list

<a id="canonical-1022110102232113-1322331322332002-1001302110311301-0032121222031212-0012233112203111-0032133031121332-1120111131301233-0031313201313003"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1012120321013120-1230102021223221-0333112311311001-0022203231130331-2013031202113132-0012303213113220-3113323130201011-2201020232232221"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list`

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-3221210001033021-3213323112110210-0303332310221300-3010010031220001-1321300032110330-2100112023011002-1220213231112021-1133302101303310): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-3223132131320130-1031131203001211-1330223113123033-1313021132302001-0212032203113331-0032332111011113-0121001002230001-3301002323102022): complete subsection reference.

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-3201311000222010-2102110103301032-0310233302121130-1030023110303313-1322311213213110-1231010033120031-1323011220112322-0330210101120032): complete subsection reference.

- [path](data-sources--protected_application--reference--group-002.md#canonical-3231233032103231-0013032330102110-1323001000103210-0030222332021032-0211012312220032-2111122002201203-2232302230231231-0301010120112231): complete subsection reference.

<a id="canonical-3221210001033021-3213323112110210-0303332310221300-3010010031220001-1321300032110330-2100112023011002-1220213231112021-1133302101303310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- CloudFront.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0220230100013000-1333131101012111-2332130100013303-3231310230201322-0001030012302031-0020302112313120-3020212103103213-3302221102202232"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223132131320130-1031131203001211-1330223113123033-1313021132302001-0212032203113331-0032332111011113-0121001002230001-3301002323102022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- CloudFront.js_insertion_rules.exclude_list.domain

<a id="canonical-2310011202103220-2120031201231120-2031111002322120-2101210213321012-0203031233210110-0023310222200011-3330121231230110-0320113000012130"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

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

<a id="canonical-3022230010323200-3330122300330022-0333230222032001-1000003201012111-1232033332232030-1123331332223122-1213300233120120-3313223022311303"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.domain`

<a id="canonical-1312003010021121-2131132233112203-1301133311330313-1302011000123203-1231102021200223-1211113223313022-0110311233132032-0320230312121003"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2112021322332330-2131110002111133-2131203200303212-2212223020000312-3023023112112233-3220122222213010-0332201231321221-2122212103300133"></a>

<a id="canonical-1320200223231003-0231021013220113-3320001120203231-2232331203100301-2012101132232322-3323110030303301-3212103310110133-2121111002113003"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-3202103101210013-1322223301031303-2111001232303123-0130100201110211-3111130131332133-2111220210321210-1030001312213131-0300323120333213"></a>

<a id="canonical-1113220222320223-0201311121102102-0312211002102303-0312203100221011-2310233120303131-0303311011130302-0222230323212000-1011030313120203"></a>

#### `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-3201311000222010-2102110103301032-0310233302121130-1030023110303313-1322311213213110-1231010033120031-1323011220112322-0330210101120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- CloudFront.js_insertion_rules.exclude_list.metadata

<a id="canonical-0111102101213003-1230113111031231-0202211021223213-3122322320000032-2111133303122031-0023010133301211-0230332122011120-2020303230020002"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-1023023020030303-1003012200023211-1100200112333121-1021021100013333-2010320103233333-0122210031331112-2022030013323202-3303211311203201"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1212322113323201-3333131200112221-2300302222213123-2231021010230002-1033001230100033-0322330313130030-1202001210233000-1311002130033210"></a>

#### `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0211112121011120-0112010221003202-1131131111110021-2330021223000302-2323011320020313-3100123003313300-2223321130323223-2123131011223222"></a>

<a id="canonical-1113122003102100-2201100332033333-3221122220033012-1302301230212000-2110212100301211-3130221333002113-2132103303312302-1330333230002110"></a>

#### `cloudfront.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3231233032103231-0013032330102110-1323001000103210-0030222332021032-0211012312220032-2111122002201203-2232302230231231-0301010120112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cloudfront.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- CloudFront.js_insertion_rules.exclude_list.path

<a id="canonical-3121200221120213-1212323031312312-0120132303113322-3220232131333203-0320113033300333-3330203120033131-2230000132210033-1021131132131330"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-0121001032301313-3210121310231032-0020011120230130-1133003223211323-3210233312301202-2222211110101032-2221333130112132-1002133222222100"></a>

### Direct properties for `cloudfront.js_insertion_rules.exclude_list.path`

<a id="canonical-0233000020202120-0013302133203213-0330231022023002-1011032021223201-3321321330021110-0122300030001123-3032313010202231-0111020030333131"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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

<a id="canonical-2002121032033202-1211222023232312-1223012331322120-3120101303122121-1133110300002201-2132121333312111-3000022310123101-3313302000300003"></a>

<a id="canonical-0123211312232130-2322210021021120-2022031121111023-2232233301322232-1122002022313022-2101303100101312-1220201000221112-2220322132122132"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2010031312303002-3211110123333130-0010300013013233-0013113130331021-0230330333102033-1312101213130333-1323333222201002-2003221002131312"></a>

<a id="canonical-2132323220121312-0333220121023030-0001130312312203-2003002320123113-2223100001122110-0323203232102333-0102131201011110-3201002313212101"></a>

#### `cloudfront.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
