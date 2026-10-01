---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0201031313232311-3220121303010230-2211133102231031-0222312301021132-0332220331002203-2103122101232200-3100332121110222-3220123332121131"></a>

## Cloudflare.protected_endpoints.metadata — metadata / 332323102122 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.metadata

<a id="canonical-3221022120313012-1303120330313233-0020321102321130-0111022310321011-2323201323110120-1012031131222321-3200101033123013-2031130333101212"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-2131213011033203-0123302122021220-2323112230203002-3223332013222310-3110311201313121-1230311230212123-1113231110321221-2313310101223211"></a>

## Direct properties — metadata / 332323102122 / 3

<a id="canonical-0232220100232332-2332111110331120-2101103203100001-3013213121133323-2000221001130302-2031231222203113-1030332010001200-2023231311000330"></a>

<a id="canonical-3033320203323002-1300303101302320-1103100100323010-2331321302122031-1100202232123020-3001222131101011-0223113120210102-2300112122303103"></a>

## description_spec property — metadata / 332323102122 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1322213010200031-0123000022122010-0321222011132203-3232100110301301-0323300111313021-2200100323133133-0113011023021310-1110013232013221"></a>

<a id="canonical-1021323320103230-1112130110033233-0212222213222123-0030021331010000-3132200022133320-3313330302310230-2321120001120202-3200103332020321"></a>

## name property — metadata / 332323102122 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-3110031201211110-1301201320121010-1310333001123230-2131322201130233-0322032111033201-2031213210320002-0131212300313211-2031133123102211"></a>

## Next pages — metadata / 332323102122 / 6

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010232302020023-1130032301030330-0120211122323333-0132212331012001-0001201221323033-0231311312333330-1000302101132010-3332100211231312"></a>

## Cloudflare.protected_endpoints.mobile_client — mobile_client / 310231012232 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.mobile_client

<a id="canonical-1010220333030321-3011112023023130-2000012003112221-3201230323313001-1023022010333321-1223132233320133-1302031320211301-0001321020130330"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

<a id="canonical-3330102311203003-0112021321020000-3220211320020312-2120120313103132-3121132120000230-1302221222322321-0301300133130213-0233222013201332"></a>

## Direct properties — mobile_client / 310231012232 / 3

- [block](data-sources--protected_application--reference--group-002.md#canonical-2100321001121021-2011131032233122-1001321032021333-1333022302220221-1003031023303100-2010013313300132-0120002103202232-3312121202113200): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000): complete subsection reference.

<a id="canonical-0120220011032103-1101123200222032-2031020121111022-3132111230002032-2212230330100201-0320233201210000-2201331330011030-0130233013121030"></a>

## Next pages — mobile_client / 310231012232 / 4

- [cloudflare.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-002.md#canonical-2100321001121021-2011131032233122-1001321032021333-1333022302220221-1003031023303100-2010013313300132-0120002103202232-3312121202113200)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2100321001121021-2011131032233122-1001321032021333-1333022302220221-1003031023303100-2010013313300132-0120002103202232-3312121202113200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031233310003033-1023112103202020-1122233303300300-3123210102212211-1223002101333211-2032033103031021-3131012010030132-0233000002133232"></a>

## Cloudflare.protected_endpoints.mobile_client.block — block / 112121020323 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- Cloudflare.protected_endpoints.mobile_client.block

<a id="canonical-0011030321232033-3112330300003222-2203313000102322-0002121322210020-2212321201102030-3211133011132213-3300013021011003-2333212102321001"></a>

Type: `"single"`. Computed.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

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

<a id="canonical-1101103033230331-0113311200132300-0200301323321231-1233302232203001-1111033221102200-0112123001233110-2021311222002132-1000103113313000"></a>

## Direct properties — block / 112121020323 / 3

<a id="canonical-3002212002121322-0331311301302111-0012010022113031-2230210221011222-3003200120222332-1322010330021320-3332030300032112-0303010102310223"></a>

<a id="canonical-0020321012132232-3203323031202200-1131231302133223-2031020232223022-3310122121020310-0100102031300223-1111223211022320-1222000003211013"></a>

## body property — block / 112121020323 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-1132100231013330-3023221030031000-2100323111010232-1011033311333212-3323321023312112-2220201332223303-3232002233220031-3000023120031022"></a>

<a id="canonical-1101002012013022-2130232120310010-2010303101000313-2211313332031001-1130310321331010-2323111320020121-3311232133210330-2020122323303310"></a>

## content_type property — block / 112121020323 / 5

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0330033333010320-3100002011320022-2130100230032300-1222332212012123-0033010303330201-3211233131301203-1000100310233232-1231111111020322"></a>

<a id="canonical-1110213220001100-3113310321003201-2231120121103123-2200010120211110-3023233200130011-1222101132313221-0321031322202222-1121233031301000"></a>

## status property — block / 112121020323 / 6

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

<a id="canonical-0302032111120330-0002222123301301-3122101303131021-1110313033100001-3233103112300210-2232132120022331-1103000223100303-3330102222323020"></a>

## Next pages — block / 112121020323 / 7

- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232020231311220-3031322002302202-1322302202021212-1001000322011203-2113032231122230-3300112303010210-0320011012211221-0131301101331302"></a>

## Cloudflare.protected_endpoints.mobile_client.continue — continue / 333333200311 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- Cloudflare.protected_endpoints.mobile_client.continue

<a id="canonical-1023223321310210-3331200330112313-2012330003311121-3331213123320200-1133031020111031-3003121231301002-1313100030131010-0232210303022310"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

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

<a id="canonical-0010110132201032-2102301110001303-1211323302101031-1031301022122132-0200213210003333-0010120030230123-1110223100012333-1313131122010300"></a>

## Direct properties — continue / 333333200311 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-3220100012023330-1323031030130212-1131231231310221-1212102033201120-2211211323133110-1122230313222030-1212231011013130-3110110322201230): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-1212133201112222-1300303201121101-0132232031112131-1310032212313312-0231332101100333-2012201200001023-0311333313212001-0120112013230311): complete subsection reference.

<a id="canonical-2231302231330200-0111223311003131-0033200110303030-2130213312303033-3323032311210023-3002233223120311-3021300122121023-3210022031033022"></a>

## Next pages — continue / 333333200311 / 4

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-3220100012023330-1323031030130212-1131231231310221-1212102033201120-2211211323133110-1122230313222030-1212231011013130-3110110322201230)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-1212133201112222-1300303201121101-0132232031112131-1310032212313312-0231332101100333-2012201200001023-0311333313212001-0120112013230311)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3220100012023330-1323031030130212-1131231231310221-1212102033201120-2211211323133110-1122230313222030-1212231011013130-3110110322201230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201122202001021-2131333030023310-2123322203323122-0032121002012312-2022023203211130-1322011332011330-3010030031202001-3200102011122030"></a>

## Cloudflare.protected_endpoints.mobile_client.continue.add_header — add_header / 333101021331 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000)
- Cloudflare.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-2233131220133101-3320211231301300-0012002002200121-1321030012023211-1311203032023033-3020032301233210-3302020022232133-3101023103230002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0200222112233033-0001000311221120-3132003001223010-3212003311321322-2230332210020332-3122211031232322-1300313022331121-0322330002332020"></a>

## Direct properties — add_header / 333101021331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113210113123331-0211330001333112-3113011223303020-0300302120121000-0223221010230301-3033322020033330-0110003203220103-1221321012130320"></a>

## Next pages — add_header / 333101021331 / 4

- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1212133201112222-1300303201121101-0132232031112131-1310032212313312-0231332101100333-2012201200001023-0311333313212001-0120112013230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330200020300212-0330210132112200-3003232301302103-2012302101123100-2311313031302130-0222103033030123-0132113303032131-0113211120121132"></a>

## Cloudflare.protected_endpoints.mobile_client.continue.no_header — no_header / 321002101032 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000)
- Cloudflare.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-3123322100331231-2200302121210320-0222112133223112-3011323032221213-2011211113231233-3021111323231201-2120221131321112-3330002100121031"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0103011310231113-0103303011333122-3023100232010132-3312200301132332-3010023131123021-3111101113023023-0212110310102203-0131000110020131"></a>

## Direct properties — no_header / 321002101032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123132222130233-3131132231203010-3121132222001022-0131101222030112-2021110330123300-0211131022230202-0303303133210123-2202300200032323"></a>

## Next pages — no_header / 321002101032 / 4

- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-0311320122200022-2213032123021112-1102223311223231-0210303203321020-0001233023123033-3002220212021321-3332311010210313-0313201122120000)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3232201300123231-0223201311223113-1102002220010310-1220002021030023-0202033222102033-0032122023100220-2313102223310321-3310323301001221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122322000110232-2013303323002211-0320010223111221-2013111132311312-0002201332232303-0203111223132102-1012120331022023-2003022011222112"></a>

## Cloudflare.protected_endpoints.path — path / 322331133310 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.path

<a id="canonical-1000132322033112-3231322312102012-1222330212022121-0300021310033210-3013021100123112-1303103332231231-2133103032313212-0030112230320003"></a>

Type: `"single"`. Computed.

Path. URI Path

Upstream description:

URI Path

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

<a id="canonical-2102120013223213-2310311301002031-3033021002333103-0330031300310012-1033233223003023-1031100331110013-3012022021221023-2210013132121312"></a>

## Direct properties — path / 322331133310 / 3

<a id="canonical-0312232213302331-2320312320032131-1120002000222233-0313233322213231-1031223322233202-3030333132113130-1200231023301011-3120122300231301"></a>

<a id="canonical-3230130233202122-3313211210201010-2023021103021130-3023102021332021-0220021333121301-2120120021203210-0110132212121313-0033103022011020"></a>

## caseinsensitive property — path / 322331133310 / 4

Type: `"bool"`. Computed.

Should path be searched case insensitive;.

Upstream description:

Should path be searched case insensitive;

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

<a id="canonical-1330213122201030-0033211301103231-1310031230220323-1021210322203233-0123013212023301-2011320003110202-1103000031222100-2321330210213331"></a>

<a id="canonical-0102231301001303-0132100110302110-3310030222223111-0102212111300000-1000200332213110-1033223031101310-1331133201033212-1302303201232211"></a>

## path property — path / 322331133310 / 5

Type: `"string"`. Computed.

Path. URI Path

Upstream description:

URI Path

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

<a id="canonical-3032200311310233-0031020231033331-2232200101202233-3220233233030230-1231021003302301-2331002200222022-0333220002332113-3330200012311133"></a>

## Next pages — path / 322331133310 / 6

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011003211110332-1200223022021103-2112132012210323-3122111011322001-1232013102020130-0130221122031212-1123203320003211-1321322131210201"></a>

## Cloudflare.protected_endpoints.web_client — web_client / 322133022223 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.web_client

<a id="canonical-2123131313302132-3212102211121033-1132332332002121-0020011313131111-0212321310321313-1331132313301321-3123223312012210-2332101133222332"></a>

Type: `"single"`. Computed.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

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

<a id="canonical-2131023001233332-2301212021122310-3323223210200122-2010331331013231-2123030223321031-1032030333233001-3321112100131011-0010212310032131"></a>

## Direct properties — web_client / 322133022223 / 3

- [block](data-sources--protected_application--reference--group-002.md#canonical-0203121101112122-2220230131122122-2331320031011302-3032112231121120-0020001200233102-1010301100012310-0310031222001003-0003010220332322): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030): complete subsection reference.

- [redirect](data-sources--protected_application--reference--group-002.md#canonical-0123313211310020-0123021022331131-2031002202001321-0330103123221201-0213010332000230-2332032020103022-1002323213011322-0223022223202313): complete subsection reference.

<a id="canonical-0300230101221201-2010101221312310-0010033011133110-2132033010003013-1202000223231100-0312231210110303-2300232212030201-1103213231123012"></a>

## Next pages — web_client / 322133022223 / 4

- [cloudflare.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-002.md#canonical-0203121101112122-2220230131122122-2331320031011302-3032112231121120-0020001200233102-1010301100012310-0310031222001003-0003010220332322)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- [cloudflare.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-002.md#canonical-0123313211310020-0123021022331131-2031002202001321-0330103123221201-0213010332000230-2332032020103022-1002323213011322-0223022223202313)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0203121101112122-2220230131122122-2331320031011302-3032112231121120-0020001200233102-1010301100012310-0310031222001003-0003010220332322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220120222330301-1122222302020002-1010033303323033-2302012300301112-1220130022322331-0010012211011332-0202010031112210-2013103332131013"></a>

## Cloudflare.protected_endpoints.web_client.block — block / 031100131010 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- Cloudflare.protected_endpoints.web_client.block

<a id="canonical-0110321102221221-3003230120031200-0222302001211022-3333202323222100-3222132032133102-0321021203223010-3022220331210000-2131230301320203"></a>

Type: `"single"`. Computed.

Block Response. Block Response.

Upstream description:

Block Response.

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

<a id="canonical-0113221110032101-3212123320013100-2322002011102221-3003111120021203-1213022310031003-1213030101030333-2301321113212231-0233230021311022"></a>

## Direct properties — block / 031100131010 / 3

<a id="canonical-1111001213002301-3311013200120122-3213203220031012-1132032333122110-3122300211101211-3132202212330301-3003033020210300-1211313211232010"></a>

<a id="canonical-1023200110220220-3233321231121222-3130132201121311-2210010103123110-0031302201103001-3031031303103130-1020202301300010-3111123013223012"></a>

## body property — block / 031100131010 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-1013000030303220-3332220321011232-1110332012321313-2210203332231303-3003331022120111-0210101221301220-1232122000203220-2020111303011202"></a>

<a id="canonical-3331122222321101-0301002013102010-2131333103122222-3120222031102220-0230311133332032-1312100332021203-1130230333130110-2320303010232111"></a>

## content_type property — block / 031100131010 / 5

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2323330210003111-0011030331120113-1011320311233122-2201123121222031-0131130120323303-0213131203122202-0132201300131112-0131232133220310"></a>

<a id="canonical-0010323131031103-1132231001121223-3211303010123333-1131330231202323-0110020033000312-2112031121333211-3201133001203213-2020200131010100"></a>

## status property — block / 031100131010 / 6

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

<a id="canonical-2131001021201131-2130023321313132-2202020303201300-2111201133323110-2020222303102022-0230211131011212-3332100110213032-0301032112122013"></a>

## Next pages — block / 031100131010 / 7

- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030130012212212-3300220112332102-0300233033013033-0103213030030300-1001032210231200-0331123012230101-1323210101110130-1231221133301213"></a>

## Cloudflare.protected_endpoints.web_client.continue — continue / 312321201201 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- Cloudflare.protected_endpoints.web_client.continue

<a id="canonical-3033223111030212-1323010303013003-3200331221323221-1101033300312210-3302233103010013-2232211211310000-2001200132311213-1320131123112022"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

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

<a id="canonical-3023201233113232-0211313110300113-3331002321012311-3320322111300031-0302001302022033-2213033210031231-3002210120230031-2320110020210100"></a>

## Direct properties — continue / 312321201201 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-2213203203331103-0010213030021002-0332213232103003-0032031231103300-2332212202213301-1210101311103020-1131112102012300-3133000302020121): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-1002012310211100-0032321321000012-0100330113302020-2020000201222130-1301013203020111-2131321133103320-1213310120130211-2133333110123312): complete subsection reference.

<a id="canonical-1032133131032203-1320303233123100-2311303000201111-2113231233103310-3222001011231021-3011300231211111-3121113233233331-3100301120021221"></a>

## Next pages — continue / 312321201201 / 4

- [cloudflare.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-2213203203331103-0010213030021002-0332213232103003-0032031231103300-2332212202213301-1210101311103020-1131112102012300-3133000302020121)
- [cloudflare.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-1002012310211100-0032321321000012-0100330113302020-2020000201222130-1301013203020111-2131321133103320-1213310120130211-2133333110123312)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2213203203331103-0010213030021002-0332213232103003-0032031231103300-2332212202213301-1210101311103020-1131112102012300-3133000302020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201223123331320-2122031210211222-3123201231100023-3122120332130302-2132322013230101-0212000121000310-3103001332202230-0313320331332123"></a>

## Cloudflare.protected_endpoints.web_client.continue.add_header — add_header / 112103213231 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- Cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-2302222211301220-3031310102131031-1103310221031332-1130212312222213-0333223112031333-0023310202113310-0223233013112121-0320132230303001"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3000300203301223-2313213023031132-3021330122212303-1312100231122322-3100122121030133-2222211231112311-1312110221011030-0112101331231301"></a>

## Direct properties — add_header / 112103213231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221302033101031-3032300021010032-2121222223003013-1012313031203331-3002101122203222-3123103203311213-1103001032133323-2112322323212100"></a>

## Next pages — add_header / 112103213231 / 4

- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1002012310211100-0032321321000012-0100330113302020-2020000201222130-1301013203020111-2131321133103320-1213310120130211-2133333110123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031330133121003-2001031330323021-3313003202121132-0003230132223012-1211313111122310-1111211332300321-2022310122233331-0331102202232133"></a>

## Cloudflare.protected_endpoints.web_client.continue.no_header — no_header / 221030301312 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- Cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-3222300131021023-1031102203233311-3100323311330011-0330113023203022-2323130203313302-0033233220130131-2002201000313131-3222223120023013"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0220101002010202-2333013220132210-0221303323221322-3232102012111110-2323021211110110-3312102301030130-0310233032330213-1313000002122311"></a>

## Direct properties — no_header / 221030301312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230311313010222-0311100333220122-0222231102221302-1302101133002302-0113330020030333-0221101120000031-1032120113213100-2001111302130021"></a>

## Next pages — no_header / 221030301312 / 4

- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3213201312223210-0221131202003331-2220110303121033-2111303111202333-2100320222230322-3222130000321321-1320110111113203-0311212231331030)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0123313211310020-0123021022331131-2031002202001321-0330103123221201-0213010332000230-2332032020103022-1002323213011322-0223022223202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201302122001233-1112013301322023-2210031010000313-3323331013202033-0323003233210131-2101113011100322-2023110323310201-1000311311032011"></a>

## Cloudflare.protected_endpoints.web_client.redirect — redirect / 132111302031 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- Cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-3213203111103102-1320033023102031-3130012122012100-1301232221020200-2320012310200010-0213123101313001-3120013103121110-0122031033012130"></a>

Type: `"single"`. Computed.

Redirect. Redirect.

Upstream description:

Redirect.

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

<a id="canonical-2130133102012300-0220103031103222-2302312003032310-1302330312201012-2113302320320030-0331310222023212-0011031101120011-0310010102230210"></a>

## Direct properties — redirect / 132111302031 / 3

<a id="canonical-3003312021332122-3331321302030003-3302303013222020-1011310110030003-2010023212102222-2313202123332022-2130023322111230-1131222332222120"></a>

<a id="canonical-1223203122211222-2223300220020332-3321221302320103-0223303200120203-3002213100322031-3203100121030002-0223313213312123-1211333301133313"></a>

## location property — redirect / 132111302031 / 4

Type: `"string"`. Computed.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

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

<a id="canonical-1322323223111033-2222012102031310-2131320230031120-3233012130333001-2301222030033222-3122333110302310-2122213320221202-1322002323313313"></a>

## status property — redirect / 132111302031 / 5

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

<a id="canonical-3203112200021331-0120203000202312-2211231303120210-3001132133203132-3211323202333113-2203231023231030-0121331331030200-1211333232012323"></a>

## Next pages — redirect / 132111302031 / 6

- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212021201113320-0100300313320202-3000120111220020-3303231110121012-2133113312011300-3130222112100312-2312313113031212-3332113210000202"></a>

## Cloudflare.protected_endpoints.web_mobile_client — web_mobile_client / 233122132313 / 2

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

<a id="canonical-3102303323100303-2320302112011111-3002100202223223-3300220200123333-0111003301002002-3023031132320310-3112201102321330-0313022011112222"></a>

## Direct properties — web_mobile_client / 233122132313 / 3

- [block_mobile](data-sources--protected_application--reference--group-002.md#canonical-3220122212111120-1300023312301010-1003202002102300-2312130233321233-1010133123111001-3321020311230133-1323213303332323-2132033133203311): complete subsection reference.

- [block_web](data-sources--protected_application--reference--group-002.md#canonical-3220320230212100-1332113211301131-2123313223101100-3231203230312031-0213000201330332-1300210203112322-0230020201311020-2012201300100101): complete subsection reference.

- [continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323): complete subsection reference.

- [continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020): complete subsection reference.

- [redirect_web](data-sources--protected_application--reference--group-002.md#canonical-3222312311333212-3303212132003231-3202021020011322-2012303002330010-2333223203012313-3201131232122332-0022132221221222-1230130310103133): complete subsection reference.

<a id="canonical-0321230001312113-2032000133110000-1320320332120332-3102021323132310-3323021033130132-3011031311111203-3112200101230232-0323110310001000"></a>

## Next pages — web_mobile_client / 233122132313 / 4

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-002.md#canonical-3220122212111120-1300023312301010-1003202002102300-2312130233321233-1010133123111001-3321020311230133-1323213303332323-2132033133203311)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-002.md#canonical-3220320230212100-1332113211301131-2123313223101100-3231203230312031-0213000201330332-1300210203112322-0230020201311020-2012201300100101)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-002.md#canonical-3222312311333212-3303212132003231-3202021020011322-2012303002330010-2333223203012313-3201131232122332-0022132221221222-1230130310103133)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3220122212111120-1300023312301010-1003202002102300-2312130233321233-1010133123111001-3321020311230133-1323213303332323-2132033133203311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130103020112311-0322023011102323-3223103123021201-1123023020012300-2112032322223222-2330223323330101-2230330003220301-3021300031222320"></a>

## Cloudflare.protected_endpoints.web_mobile_client.block_mobile — block_mobile / 332332032311 / 2

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

Upstream description:

Block Response.

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

<a id="canonical-2330303110203010-0220031203102323-0233113230000122-2210111010130113-0010110230121301-0121020311223103-3132302132121210-1200200221010110"></a>

## Direct properties — block_mobile / 332332032311 / 3

<a id="canonical-3322102211021001-3011132022130031-1320020313121333-2130020131121021-2030323310213112-0113010010233203-0022123103012212-3013100032123311"></a>

<a id="canonical-3331211120302021-0011030310322311-2322112320202100-3013321133200131-3231102120023202-1303130221321221-2331032031003133-3032032102113312"></a>

## body property — block_mobile / 332332032311 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-3003131003003022-0100012310020321-2123302232130010-1110321132131011-3022123213331321-2303301312003021-1212010122322110-0131210302220310"></a>

<a id="canonical-3200231301130001-0221222321033233-1221333321133310-1022000323131222-3020212330020133-0032121223211323-2210023303213320-0132231332222332"></a>

## content_type property — block_mobile / 332332032311 / 5

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1312131222031321-3310001230202232-3012101132320322-1332002330033212-0023200123002102-1201333331333303-3121300030331031-3021013300210301"></a>

<a id="canonical-0112312311133121-2110312303302232-0000311331202132-1221311310222223-3313300220032212-2023110222300210-0210111220022200-3331002003000011"></a>

## status property — block_mobile / 332332032311 / 6

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

<a id="canonical-3033322330133120-3331003003031132-3120223221211333-3210013133211033-1223122122132122-2000333310030012-1023110200111003-1211102302032232"></a>

## Next pages — block_mobile / 332332032311 / 7

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3220320230212100-1332113211301131-2123313223101100-3231203230312031-0213000201330332-1300210203112322-0230020201311020-2012201300100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203321103100102-0102012031000220-3013101022232112-2033131322123011-2223210102032323-0021000122333000-1222100122202132-2013321213232221"></a>

## Cloudflare.protected_endpoints.web_mobile_client.block_web — block_web / 101000032232 / 2

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

Upstream description:

Block Response.

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

<a id="canonical-0312221222023123-1303302223032003-2211232000203220-0332200233031312-2302221301202000-1232103300012230-0232202202120122-2332112121321311"></a>

## Direct properties — block_web / 101000032232 / 3

<a id="canonical-2100100030101302-3220201313312030-1233110231033130-1113032010320000-1110211310131123-3112132003310132-3231010232203000-3231213223323300"></a>

<a id="canonical-2330112230010223-3220112230110211-1020011202302231-1133020120112323-3303130103133212-3120302121120002-0330330200330200-2302112323211110"></a>

## body property — block_web / 101000032232 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0023123233002320-3033113200013000-3330312310120330-1133230121010021-1231110310010011-0213200332201310-2013210202020101-0110303103221103"></a>

<a id="canonical-1113200032233022-1113230203111023-0200020122001212-0012300102202300-1221001032110303-3021220020022121-3003231210103311-1331223113230001"></a>

## content_type property — block_web / 101000032232 / 5

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1223111310332000-0222331210231123-1301010200232330-0212030300020320-3133323122302200-0120331211310223-2131130102122023-0102120212331103"></a>

<a id="canonical-2012131303123001-1203202222230113-3213101211001033-3021310131013232-1012101031220323-2231203210020333-2321102112331333-2333301031121232"></a>

## status property — block_web / 101000032232 / 6

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

<a id="canonical-0133120323000012-2102333032220303-0022333313110112-0032312212311302-3031130102021002-2102311201101210-0002231312100031-1331312332303320"></a>

## Next pages — block_web / 101000032232 / 7

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113123320111232-3023221021300031-1311221020333001-2021213113222110-2212300323323000-1132233101103230-3103113222003131-3303003210202022"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile — continue_mobile / 323001233120 / 2

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

Upstream description:

Continue mitigation action.

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

<a id="canonical-1123031120330232-2003030121300332-1213101012331301-1332231033013111-3001000020012320-2210111310230220-1302332230103202-2000321002230233"></a>

## Direct properties — continue_mobile / 323001233120 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-1233302023031331-3313322001003313-0223303201001030-0321102021222213-1220200011320203-3213101111102100-0200023113001131-3131311213320301): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-0131032231132202-0203313030032310-0301132220330310-1330103313321312-3232222300012032-0000011302312310-3122112212202231-1212313120131111): complete subsection reference.

<a id="canonical-1313002323303130-1232301033111322-1001133212122212-1121120130030013-1111023023303033-2320311100030003-0001002112132002-1010112330032122"></a>

## Next pages — continue_mobile / 323001233120 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-002.md#canonical-1233302023031331-3313322001003313-0223303201001030-0321102021222213-1220200011320203-3213101111102100-0200023113001131-3131311213320301)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-002.md#canonical-0131032231132202-0203313030032310-0301132220330310-1330103313321312-3232222300012032-0000011302312310-3122112212202231-1212313120131111)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1233302023031331-3313322001003313-0223303201001030-0321102021222213-1220200011320203-3213101111102100-0200023113001131-3131311213320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020021020131001-3321031200330121-2132232102132233-1111031312031230-0220223323320232-3110113021221300-2000001213133022-0031110022331223"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header — add_header / 020110132320 / 2

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

<a id="canonical-0320102033210231-3110120113310032-1122231132322031-3302112010300032-1002302322233100-2013310322013320-0211313102233000-3130021232303000"></a>

## Direct properties — add_header / 020110132320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222012130232331-0303311023230102-2220123101333033-1113123130222133-2111103001302222-3200302313010332-2203200111110112-0222022212031322"></a>

## Next pages — add_header / 020110132320 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0131032231132202-0203313030032310-0301132220330310-1330103313321312-3232222300012032-0000011302312310-3122112212202231-1212313120131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320121203230023-1132013033131330-0000033103223110-2020310300230002-3021303123131102-0032122030031300-1101101133002310-3332113301011200"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header — no_header / 102101302321 / 2

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

<a id="canonical-0130301112303211-2203033220123021-0112312220133220-3033110123123100-1022001120000230-1213033013220331-2200322121023332-0221013013020203"></a>

## Direct properties — no_header / 102101302321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300201320301110-1212003213331130-1203133030310100-1201310213232013-1021203223132013-0023320333201023-0311313011322323-2101013122010322"></a>

## Next pages — no_header / 102101302321 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-3233132023020001-0230300021301003-0112231001313030-1003233323220102-0333312231313113-0130130313210200-0232322121003011-3230312231233323)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132103101222012-2221012220310312-3232212001210112-1023101103000231-3123203333300330-0331203221023230-0312212013133331-1021011333301131"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web — continue_web / 121013121200 / 2

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

Upstream description:

Continue mitigation action.

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

<a id="canonical-2133010200321333-2012222233131303-1310302120310322-3222322322210023-1122100210313030-0031102011133122-2231300222102323-0212200221333201"></a>

## Direct properties — continue_web / 121013121200 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-0321031302023121-0330203232021023-0011313001011321-0123011332010311-2130333230112000-1120132330322323-2111020301010210-1031202222003322): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-0203121022120332-2120330133103212-3312002013111301-3033230300110210-1113130111031202-1231103123312031-1312322201301223-0121222312021200): complete subsection reference.

<a id="canonical-2311003130100231-2312022112031021-0103331222230122-2032122330311300-0012330023113122-0333222102101022-0230130020322221-1003302031003213"></a>

## Next pages — continue_web / 121013121200 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-002.md#canonical-0321031302023121-0330203232021023-0011313001011321-0123011332010311-2130333230112000-1120132330322323-2111020301010210-1031202222003322)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-002.md#canonical-0203121022120332-2120330133103212-3312002013111301-3033230300110210-1113130111031202-1231103123312031-1312322201301223-0121222312021200)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0321031302023121-0330203232021023-0011313001011321-0123011332010311-2130333230112000-1120132330322323-2111020301010210-1031202222003322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223030313103000-3131030003101110-3032212210021301-2212320101331012-2131130213212333-3011002012110130-0012310223200211-2101122223021233"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header — add_header / 232010002030 / 2

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

<a id="canonical-2203313021102031-2230312301020313-1321311012312311-2031231202220300-2210001131030032-1300002112020013-0211313333202310-1301111111122023"></a>

## Direct properties — add_header / 232010002030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002231022201210-1312133110021000-3012302321030311-0102123111131033-2010222110132100-3010122320203313-2110130002210113-0013130021203323"></a>

## Next pages — add_header / 232010002030 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0203121022120332-2120330133103212-3312002013111301-3033230300110210-1113130111031202-1231103123312031-1312322201301223-0121222312021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130010101122031-1202131212031013-1101112232020010-3203100233112133-1212113002231333-1310122103013021-2000202320301032-3112110002030231"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header — no_header / 223210213200 / 2

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

<a id="canonical-3013213113010232-0122233102123310-1102102103231112-0033112103232310-1101333033031211-3213113311330112-2011322112300123-2233310202122033"></a>

## Direct properties — no_header / 223210213200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312130032013313-1110320133331221-0210130021321102-1103100320210021-3133013332222212-1331303202232303-0221133100220123-2013103102203200"></a>

## Next pages — no_header / 223210213200 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-0112120121201133-2302031201233331-2020320023302113-0213231000032011-0012213230012311-2031001022201203-0333200321002230-1013332023021020)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3222312311333212-3303212132003231-3202021020011322-2012303002330010-2333223203012313-3201131232122332-0022132221221222-1230130310103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222000001313202-3230102002311112-0113332120313201-1030033212300233-3002332200031122-0331203030132013-0031313302120221-1221232301303102"></a>

## Cloudflare.protected_endpoints.web_mobile_client.redirect_web — redirect_web / 100101211201 / 2

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

Upstream description:

Redirect.

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

<a id="canonical-1302312230112300-1100000333220300-2332301300031113-1222323313131313-2231323001232222-0332310312023212-1312130102033030-3123323303010323"></a>

## Direct properties — redirect_web / 100101211201 / 3

<a id="canonical-0223103100301103-1130220121012220-3333123111231203-3130222111220233-3202000133122030-3201000031033221-3112021230231200-2102230313112110"></a>

<a id="canonical-2210102310332211-3111310013313013-3010233332133032-2231300002102123-0323103212321330-1033131033112030-1122010332303120-2011120321032103"></a>

## location property — redirect_web / 100101211201 / 4

Type: `"string"`. Computed.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

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

<a id="canonical-2013113202200120-0333103203332312-1332202020022303-0222023111003012-0313213320201211-1033200210113213-0203000332220032-2313120201310331"></a>

## status property — redirect_web / 100101211201 / 5

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

<a id="canonical-2201220122223100-3020012213223211-0132320101112110-1331113032000020-2233101213103210-1200333103233102-0222230303220112-3300322021012011"></a>

## Next pages — redirect_web / 100101211201 / 6

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131102223001202-3123131222012002-0222232001003011-0233132233111300-1310031223320201-0102022331213012-2030300312013231-0033231100300321"></a>

## Cloudflare.trusted_clients — trusted_clients / 101232201310 / 2

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1030012011022021-1033130320033101-3003333320113201-1030330100033322-2023131123121200-1220110210013101-2032313012023121-0321001211332000"></a>

## Direct properties — trusted_clients / 101232201310 / 3

- [http_header](data-sources--protected_application--reference--group-002.md#canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121): complete subsection reference.

<a id="canonical-3133011202313222-2231321120132330-2100330312012112-1320100332033113-1102130100303332-1212011011121320-1011331231023333-0320101202301023"></a>

<a id="canonical-2303230010231221-0220100310003300-1200003020000333-3130200031001232-0110232333100201-0011220031230321-3231111121022201-0020312011230233"></a>

## ip_prefix property — trusted_clients / 101232201310 / 4

Type: `"string"`. Computed.

Exclusive with \[http\_header\] IP prefix string.

Upstream description:

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-2002233311030303-0021002120221301-1032331120101232-0032033033031330-1033132033312113-1033003102330001-3101121220133011-0323323122022220): complete subsection reference.

<a id="canonical-1123231202302212-1012220321110123-0112112322223123-3330310113001123-1203133010123003-3203232310200000-0121012103232111-1332103200210300"></a>

## Next pages — trusted_clients / 101232201310 / 5

- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121)
- [cloudflare.trusted_clients.metadata](data-sources--protected_application--reference--group-002.md#canonical-2002233311030303-0021002120221301-1032331120101232-0032033033031330-1033132033312113-1033003102330001-3101121220133011-0323323122022220)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131203233212320-2221303232130012-3012012320113300-3110012103120123-3010331012230200-1122310130111221-2133130112021223-3313312123002310"></a>

## Cloudflare.trusted_clients.http_header — http_header / 323313312321 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- Cloudflare.trusted_clients.http_header

<a id="canonical-3331130021002200-1121230312011231-0331100323123333-2031103323110322-1223013033300130-2001233130111002-1021201032222331-2001222332221103"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

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

<a id="canonical-3303022131013312-2312032012203230-1003332201020303-1101310313122221-0133103211212012-1320320211122100-3032202111020311-3201120303112131"></a>

## Direct properties — http_header / 323313312321 / 3

- [headers](data-sources--protected_application--reference--group-002.md#canonical-3330100302031312-1203202323230320-1113013130203101-1211103132200221-3202331320032322-2113112202230200-3013112103121221-0003200020202011): complete subsection reference.

<a id="canonical-3030311031212202-2221121112121020-2121301320331323-0330203320323323-0311102213110321-0033011310303033-0233103121211203-0023111303110013"></a>

## Next pages — http_header / 323313312321 / 4

- [cloudflare.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-002.md#canonical-3330100302031312-1203202323230320-1113013130203101-1211103132200221-3202331320032322-2113112202230200-3013112103121221-0003200020202011)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3330100302031312-1203202323230320-1113013130203101-1211103132200221-3202331320032322-2113112202230200-3013112103121221-0003200020202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331202131333100-3232123312331303-3311011322301021-3103111023033300-0002222111332223-0212023010101220-1121232332033330-2123232321013302"></a>

## Cloudflare.trusted_clients.http_header.headers — headers / 201211223303 / 2

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0113022021310122-2010200230032012-2211102202332110-3332320312100213-1202031312233002-2120112003022233-0312333233013301-2302023121302103"></a>

## Direct properties — headers / 201211223303 / 3

<a id="canonical-2031001310010111-2123222002333301-2222322302232033-2130020203033300-1231233123120211-2010032010310102-3330320201203010-1222223021122033"></a>

<a id="canonical-3102023000231221-2112022223222013-1210220023301221-1110301213021232-2000312003133230-0121232323021032-0322111321210203-2301011130333021"></a>

## exact property — headers / 201211223303 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Header value to match exactly.

Upstream description:

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

<a id="canonical-1022303210231330-0300112130233100-1223323233212202-1231311101120110-2010230220033013-3032301100320230-2112211222200021-2123120320231131"></a>

<a id="canonical-0222332230300220-2002220223232220-0120232311133333-0002132110010232-1202202022020101-0021103213312311-2130032323211012-3021200323301311"></a>

## name property — headers / 201211223303 / 5

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-1111103330312232-3023222333111313-3013222121220132-3232031200211323-0100023312312223-2303022201212313-0113323303021213-2202032023200231"></a>

<a id="canonical-0233023110022000-2100023302101213-3230132021022002-3113203103003333-3012313300021131-2033210102121131-3031311122131022-0123320130102013"></a>

## regular expression property — headers / 201211223303 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

Upstream description:

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

<a id="canonical-1112321332122202-2023003021133001-0022311201120012-2100200130313122-2102131210120222-1120020112131222-3300032011232221-0013103301010011"></a>

## Next pages — headers / 201211223303 / 7

- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2002233311030303-0021002120221301-1032331120101232-0032033033031330-1033132033312113-1033003102330001-3101121220133011-0323323122022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322310333122212-0230330033032100-3023302111033220-3122323313030120-2330113010123201-0130111331321032-0330331332102333-2120222313011111"></a>

## Cloudflare.trusted_clients.metadata — metadata / 333012032231 / 2

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
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-2112023310323230-0023330020200330-3312131013032311-3330111230221302-1201113313033300-3101333322201100-2000003232102312-2133301310013033"></a>

## Direct properties — metadata / 333012032231 / 3

<a id="canonical-3102200112001100-0302320111131300-3001023202330231-2101332330202102-2020011030303312-2031202020320100-3201023210031300-3303233103133210"></a>

<a id="canonical-0201201302300101-3120120330003302-3112311322032000-2120122132103331-2020202312001303-1111111033100303-0302000130302130-3103010311110112"></a>

## description_spec property — metadata / 333012032231 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3020101202300232-0103110132023222-0111132213330013-0132002333331313-2010031132201301-0032222003302211-1303130200300312-3002010032032001"></a>

<a id="canonical-3301313210113231-2322001110232011-0331232312200133-2101010132010110-2210203210021110-2331112121333223-1211022001322233-0232302233211111"></a>

## name property — metadata / 333012032231 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-2111301002322003-3312312200320122-2330001231000133-1320023101100323-1022201003311201-2313212212220103-1010222213020102-0230231213020311"></a>

## Next pages — metadata / 333012032231 / 6

- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210121203020121-3323221000211201-1020122331112000-1102233333123312-3222112330223321-1002210312333210-0303112111332321-0313011320231100"></a>

## CloudFront — CloudFront / 320013132200 / 2

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

<a id="canonical-2022331203023012-0222313132120013-2113131233010031-3210112132101002-0231212113131320-3011322300012231-2303303031230200-1123012101332112"></a>

## Direct properties — CloudFront / 320013132200 / 3

- [aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-3003122220113222-1131032101001130-3201033110130300-2110030102303022-0223211112012013-0323212101332322-1210200032213331-0333020213121322): complete subsection reference.

- [aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331): complete subsection reference.

<a id="canonical-1301031001003013-2221130322231013-3100001332203031-2221023330132113-2112310033313122-3203101010111211-0202200123001312-0212211011233021"></a>

<a id="canonical-2100203002113121-3132213020100301-0213111022323322-3221110112331233-3303233332202131-1020130122031120-3312133203122033-0103203002113203"></a>

## continue_mitigation_action_hdr property — CloudFront / 320013132200 / 4

Type: `"string"`. Computed.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

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

<a id="canonical-1232200003020323-3303313202320133-3302230300100213-1310112312120102-0121032103232003-2113022313333120-1112230333130302-0001222122121021"></a>

## data_sample property — CloudFront / 320013132200 / 5

Type: `"number"`. Computed.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

Upstream description:

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte)

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0000222021200010-2011020121021333-2130133200233221-0121023312331123-3033100030123302-3230210113122123-0221312222312301-3003221213310123"></a>

## loglevel property — CloudFront / 320013132200 / 6

Type: `"string"`. Computed.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Upstream description:

Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and
Informational)

&#8203;- LOG\_UNDEFINED: Undefined

&#8203;- LOG\_ERROR: Error

Log only errors &#8203;- LOG\_WARNING: Warning

Log malicious requests &#8203;- LOG\_INFO: Info

Log all requests &#8203;- LOG\_DEBUG: Debug

Log debugging data.

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

- [manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-3123213201022200-1010003020110031-2133001121200133-1103021210123222-3223203013201202-3103020013112213-0232333101021031-3332032331023310): complete subsection reference.

- [mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113): complete subsection reference.

- [protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200): complete subsection reference.

<a id="canonical-1301230011221011-0212111310230231-2110331232201203-1133011000012311-0320213032210100-1132332330023202-2121330001021322-0003102311110013"></a>

<a id="canonical-1123320023223113-1332112133102201-0211211133230010-2320233232322303-1201011303100133-1001032003313120-2210313003001310-1231313131033200"></a>

## timeout property — CloudFront / 320013132200 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0111223121002030-1210202031022300-0013123122132333-3220110121213320-2031231210211202-3033021103000211-1223231313120112-3120233312310121"></a>

## Next pages — CloudFront / 320013132200 / 8

- [cloudfront.aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-3003122220113222-1131032101001130-3201033110130300-2110030102303022-0223211112012013-0323212101332322-1210200032213331-0333020213121322)
- [cloudfront.aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331)
- [cloudfront.disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-0113021203101100-0221102102222310-3021322330123313-0230130301101223-0221130202123110-3222011230311203-0112033023011310-1123112330333313)
- [cloudfront.disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-3123120301310311-3222321021000312-2112303010201322-0000033331021033-1001330003203210-0130222013202130-1020213232121010-3131322213232233)
- [cloudfront.disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-1013001331001333-3132220323212221-0111011020302121-3313130102030233-0113123332303201-3230122131113103-0332313320101021-0231300013121230)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-3123213201022200-1010003020110031-2133001121200133-1103021210123222-3223203013201202-3103020013112213-0232333101021031-3332032331023310)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-3131003001231302-2012000112130122-3232010022011130-2022222222201022-1200232331112220-3022230002311131-3010220330323012-1122321033303022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3003122220113222-1131032101001130-3201033110130300-2110030102303022-0223211112012013-0323212101332322-1210200032213331-0333020213121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310212010033202-3133213232132323-0330020220010123-1302023032000002-1123020023113100-2111210301323333-1322103033330001-1102023313012230"></a>

## CloudFront.aws_configuration_id_selector — aws_configuration_id_selector / 303211111120 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.aws_configuration_id_selector

<a id="canonical-0220011020023322-2010100312322220-1301222102301210-3122103330202030-2203100121303200-0102303311122322-1220201031201110-2132332121312113"></a>

Type: `"single"`. Computed.

Configuration parameter for aws configuration ID selector.

Upstream description:

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

<a id="canonical-3021101120330110-2021111023202032-0132010012231232-3003023002112000-2321113202332003-0113000113300113-1312010000223120-0203020301032320"></a>

## Direct properties — aws_configuration_id_selector / 303211111120 / 3

<a id="canonical-0212130313002203-2330022232023102-2113332210113110-2120102010120031-2022110322333023-1021122301102322-1302013113313201-2320021211213200"></a>

<a id="canonical-0122023221133330-0233212202032300-0011120112031332-2330000231011103-0002313002023231-2321020000013311-1000201312010023-3220223000202303"></a>

## IDs property — aws_configuration_id_selector / 303211111120 / 4

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

<a id="canonical-2302103302302010-1113010002022000-0101210331003221-2120333103311022-3113330113211222-2100301210231003-1212102101131131-2212302013333010"></a>

## Next pages — aws_configuration_id_selector / 303211111120 / 5

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012022313300330-0313303030333130-1033000302321330-3030012300023123-1002220110123110-1123112333232120-1102123302212120-0122211232202320"></a>

## CloudFront.aws_configuration_tag_selector — aws_configuration_tag_selector / 101131033030 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.aws_configuration_tag_selector

<a id="canonical-3121212122301010-0021023210313022-2012123131120310-0133230001312122-1302310132011011-1010303032323323-1120020221213303-2031113212020230"></a>

Type: `"single"`. Computed.

Distribution Tag List. CloudFront distribution tag list.

Upstream description:

CloudFront distribution tag list.

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

<a id="canonical-3110112000223123-3021320203121001-2003123331330030-0203212020231301-2100000010031010-2002121321202210-3022301230320000-0010202110033312"></a>

## Direct properties — aws_configuration_tag_selector / 101131033030 / 3

<a id="canonical-0310120031133223-3113330201332303-1331023132032323-2023221313012120-1231033102120101-2311223001203122-1130223123310202-1201312301121000"></a>

<a id="canonical-0000032221230303-2132323320311221-3311001032311231-1113121130222001-3322221233011312-0303322110021121-3122013112330011-1120223012313103"></a>

## tags property — aws_configuration_tag_selector / 101131033030 / 4

Type: `["map", "string"]`. Computed.

List contains the CloudFront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Upstream description:

List contains the CloudFront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0222330032020121-0132203232333323-0332322301102211-3220323002332302-0133300212022323-1332213321310112-2323031321220121-3330131322312212"></a>

## Next pages — aws_configuration_tag_selector / 101131033030 / 5

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0113021203101100-0221102102222310-3021322330123313-0230130301101223-0221130202123110-3222011230311203-0112033023011310-1123112330333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100011333011110-2111203000122331-2031332010133002-2333002113203231-2011022200121011-0213101133233331-1333010002021120-2230112032131302"></a>

## CloudFront.disable_aws_configuration — disable_aws_configuration / 020302111313 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_aws_configuration

<a id="canonical-2213221110001032-1000233003211022-1110113313110233-3301010111320201-1211101032222320-0223020010132233-2133333002221101-3000103213100111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable aws configuration.

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

<a id="canonical-0110031110200013-1111230121302200-2110101303223002-0022020202220001-3203213011203003-0012030231131113-3213210111312001-2132312222022132"></a>

## Direct properties — disable_aws_configuration / 020302111313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211310213130020-1321221320331121-0013200133130231-1213101231202222-1233113212201322-1110313022100000-3110202033111302-1331312232133300"></a>

## Next pages — disable_aws_configuration / 020302111313 / 4

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3123120301310311-3222321021000312-2112303010201322-0000033331021033-1001330003203210-0130222013202130-1020213232121010-3131322213232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012132002011110-2322020101111102-3211003101021223-3132122220013112-0033322231112232-0221303212300033-3020100223123100-0103332110103100"></a>

## CloudFront.disable_js_insert — disable_js_insert / 131313032100 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_js_insert

<a id="canonical-3110121231323013-2010102333202133-1223003323100001-1011201130020222-2221313203100310-2010330102021023-3122023000010030-1300121212112323"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1333101201222033-0232020211031302-3033213211120131-0230303332201013-0110010303201000-0231203013320022-1013020322300332-0222110113312000"></a>

## Direct properties — disable_js_insert / 131313032100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132230301110033-1030103101001311-2333300323312311-3310110031321220-0032112333211002-1122113010121001-1003030102213232-2001003111032210"></a>

## Next pages — disable_js_insert / 131313032100 / 4

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1013001331001333-3132220323212221-0111011020302121-3313130102030233-0113123332303201-3230122131113103-0332313320101021-0231300013121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113323131133320-1231202311030132-3022113123201301-1032202130330331-0132312330001101-2030312010103201-1303303211111300-0230010010100321"></a>

## CloudFront.disable_mobile_sdk — disable_mobile_sdk / 103312233132 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.disable_mobile_sdk

<a id="canonical-1202133233323323-0222330321230102-0200201210331330-1220310232323100-0031000031010012-3331001130303213-3022323231212210-1101113200133302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1022211100110320-0301230322201301-2002011103023333-1010332100313201-0000311012012221-1112031002212222-0300221322233112-3310222221233222"></a>

## Direct properties — disable_mobile_sdk / 103312233132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131323111311322-1211132211031130-2111120032100023-0103310102121303-1313313202012021-0121230122230310-0322213132123312-1322012330031301"></a>

## Next pages — disable_mobile_sdk / 103312233132 / 4

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332232231320112-0133110131130020-0223102022230132-2033310131320220-1033120103120101-0101010120201300-2310313131012210-3302010202001223"></a>

## CloudFront.js_insertion_rules — js_insertion_rules / 220133303211 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.js_insertion_rules

<a id="canonical-2222221321010301-0332022320311011-2331212321332101-1031010113312012-3210233121102310-0021232023021320-0311120231013130-1213111111103113"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

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

<a id="canonical-1212010321311311-0323212220122300-3232212220210210-3022303232000331-3103300033102310-3313210000300030-1330332220320230-0030230010021000"></a>

## Direct properties — js_insertion_rules / 220133303211 / 3

- [exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331): complete subsection reference.

<a id="canonical-2213213031231212-0003203030110002-2220313131211120-2233100001023033-3012103100102332-0001230301020121-3000021330102332-2002210313023222"></a>

<a id="canonical-2001323333301203-0022120201030113-1003330330220312-1200233213313333-1112232212203232-2112210010020223-2113110220222320-2113303333303211"></a>

## javascript_location property — js_insertion_rules / 220133303211 / 4

Type: `"string"`. Computed.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Upstream description:

All inside networks.

&#8203;- JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED

Undefined Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag.
Insert JavaScript before first tag.

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

<a id="canonical-0313231111021311-0220131303131102-3210232033220201-3312302123132233-3323233132113122-1021330321310201-1013322131310212-3312001020123022"></a>

## javascript_mode property — js_insertion_rules / 220133303211 / 5

Type: `"string"`. Computed.

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

<a id="canonical-0130112121022120-0030002102103101-2230301301021112-0323022030201103-0121211133312332-0213110232231302-2232122131030220-1101302210023010"></a>

## js_download_path property — js_insertion_rules / 220133303211 / 6

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/CommonJS’.

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

- [rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122): complete subsection reference.

<a id="canonical-0201202333323333-2031123302020221-3221120013101211-1331012022023010-3022330323301313-0120110100212332-0200222332101022-0332113102211122"></a>

## Next pages — js_insertion_rules / 220133303211 / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012120321013120-1230102021223221-0333112311311001-0022203231130331-2013031202113132-0012303213113220-3113323130201011-2201020232232221"></a>

## CloudFront.js_insertion_rules.exclude_list — exclude_list / 120002001121 / 2

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

<a id="canonical-2103003130003011-0231103033130033-1231331011010322-1223000123230230-3030331223011201-1103032201212201-2320323201121200-3021030221313321"></a>

## Direct properties — exclude_list / 120002001121 / 3

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-3221210001033021-3213323112110210-0303332310221300-3010010031220001-1321300032110330-2100112023011002-1220213231112021-1133302101303310): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-3223132131320130-1031131203001211-1330223113123033-1313021132302001-0212032203113331-0032332111011113-0121001002230001-3301002323102022): complete subsection reference.

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-3201311000222010-2102110103301032-0310233302121130-1030023110303313-1322311213213110-1231010033120031-1323011220112322-0330210101120032): complete subsection reference.

- [path](data-sources--protected_application--reference--group-002.md#canonical-3231233032103231-0013032330102110-1323001000103210-0030222332021032-0211012312220032-2111122002201203-2232302230231231-0301010120112231): complete subsection reference.

<a id="canonical-3210323313212202-1313201013002323-2003323111202121-1131100020211020-2210112011020210-1103332230310101-0203113223303101-2133032112311032"></a>

## Next pages — exclude_list / 120002001121 / 4

- [cloudfront.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-002.md#canonical-3221210001033021-3213323112110210-0303332310221300-3010010031220001-1321300032110330-2100112023011002-1220213231112021-1133302101303310)
- [cloudfront.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-002.md#canonical-3223132131320130-1031131203001211-1330223113123033-1313021132302001-0212032203113331-0032332111011113-0121001002230001-3301002323102022)
- [cloudfront.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-002.md#canonical-3201311000222010-2102110103301032-0310233302121130-1030023110303313-1322311213213110-1231010033120031-1323011220112322-0330210101120032)
- [cloudfront.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-002.md#canonical-3231233032103231-0013032330102110-1323001000103210-0030222332021032-0211012312220032-2111122002201203-2232302230231231-0301010120112231)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3221210001033021-3213323112110210-0303332310221300-3010010031220001-1321300032110330-2100112023011002-1220213231112021-1133302101303310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200103330002233-0330132003110031-0213321130203113-3010121130112011-2021323320220210-3132301000012003-2331222003201011-3300332001220103"></a>

## CloudFront.js_insertion_rules.exclude_list.any_domain — any_domain / 002103130313 / 2

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

<a id="canonical-0310112123310012-2312232031320120-0121321302321022-1221333231023231-3220310311300101-1101133033222020-1001120003122320-1001021130132033"></a>

## Direct properties — any_domain / 002103130313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122120011001300-3130132321101102-0130123332221232-3031102222223120-2121013232021120-3211102131300222-3302233311100011-3212102202023010"></a>

## Next pages — any_domain / 002103130313 / 4

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3223132131320130-1031131203001211-1330223113123033-1313021132302001-0212032203113331-0032332111011113-0121001002230001-3301002323102022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022230010323200-3330122300330022-0333230222032001-1000003201012111-1232033332232030-1123331332223122-1213300233120120-3313223022311303"></a>

## CloudFront.js_insertion_rules.exclude_list.domain — domain / 333121110220 / 2

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

Upstream description:

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

<a id="canonical-1320200223231003-0231021013220113-3320001120203231-2232331203100301-2012101132232322-3323110030303301-3212103310110133-2121111002113003"></a>

## Direct properties — domain / 333121110220 / 3

<a id="canonical-1312003010021121-2131132233112203-1301133311330313-1302011000123203-1231102021200223-1211113223313022-0110311233132032-0320230312121003"></a>

<a id="canonical-1113220222320223-0201311121102102-0312211002102303-0312203100221011-2310233120303131-0303311011130302-0222230323212000-1011030313120203"></a>

## exact_value property — domain / 333121110220 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-2112021322332330-2131110002111133-2131203200303212-2212223020000312-3023023112112233-3220122222213010-0332201231321221-2122212103300133"></a>

<a id="canonical-0202233111002120-3103223310231310-1123223201011033-3232321113332120-3201123131011101-2022113201002310-3103212302300112-0223210110113020"></a>

## regex_value property — domain / 333121110220 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

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

<a id="canonical-3202103101210013-1322223301031303-2111001232303123-0130100201110211-3111130131332133-2111220210321210-1030001312213131-0300323120333213"></a>

<a id="canonical-1110221311011331-1020330301200210-0221100312203333-1023100303000200-3011200321313302-0210232031222212-3221001232320332-3321312030032213"></a>

## suffix_value property — domain / 333121110220 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-0203300101120320-2300233303100320-2321200002232131-3022121210311301-0222113033232010-0123012230213022-2203301300130230-1303021211012233"></a>

## Next pages — domain / 333121110220 / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3201311000222010-2102110103301032-0310233302121130-1030023110303313-1322311213213110-1231010033120031-1323011220112322-0330210101120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023023020030303-1003012200023211-1100200112333121-1021021100013333-2010320103233333-0122210031331112-2022030013323202-3303211311203201"></a>

## CloudFront.js_insertion_rules.exclude_list.metadata — metadata / 000223232022 / 2

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
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-1113122003102100-2201100332033333-3221122220033012-1302301230212000-2110212100301211-3130221333002113-2132103303312302-1330333230002110"></a>

## Direct properties — metadata / 000223232022 / 3

<a id="canonical-1212322113323201-3333131200112221-2300302222213123-2231021010230002-1033001230100033-0322330313130030-1202001210233000-1311002130033210"></a>

<a id="canonical-0121331231023130-2213220311123201-2202122313130013-2330300021320332-3031011230211312-1302120232232111-3310212322202102-2301233211322030"></a>

## description_spec property — metadata / 000223232022 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0211112121011120-0112010221003202-1131131111110021-2330021223000302-2323011320020313-3100123003313300-2223321130323223-2123131011223222"></a>

<a id="canonical-0222131301211310-1013323112110022-0330230021223301-3002001130300032-3322320212120033-3230202302111212-0132323100002331-1122132012000020"></a>

## name property — metadata / 000223232022 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-0110202032303213-2310200003100310-2313002332331001-3100111003221201-1333300231101003-0232133021013111-2232331322131101-3312222312312001"></a>

## Next pages — metadata / 000223232022 / 6

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3231233032103231-0013032330102110-1323001000103210-0030222332021032-0211012312220032-2111122002201203-2232302230231231-0301010120112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121001032301313-3210121310231032-0020011120230130-1133003223211323-3210233312301202-2222211110101032-2221333130112132-1002133222222100"></a>

## CloudFront.js_insertion_rules.exclude_list.path — path / 223102212122 / 2

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

<a id="canonical-0123211312232130-2322210021021120-2022031121111023-2232233301322232-1122002022313022-2101303100101312-1220201000221112-2220322132122132"></a>

## Direct properties — path / 223102212122 / 3

<a id="canonical-0233000020202120-0013302133203213-0330231022023002-1011032021223201-3321321330021110-0122300030001123-3032313010202231-0111020030333131"></a>

<a id="canonical-2132323220121312-0333220121023030-0001130312312203-2003002320123113-2223100001122110-0323203232102333-0102131201011110-3201002313212101"></a>

## path property — path / 223102212122 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

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

<a id="canonical-2002121032033202-1211222023232312-1223012331322120-3120101303122121-1133110300002201-2132121333312111-3000022310123101-3313302000300003"></a>

<a id="canonical-0100310320110010-2303233311321032-2020220122210111-2220331101311031-1010321032022000-1213003332220103-1111131003200210-3131331102230020"></a>

## prefix property — path / 223102212122 / 5

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

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

<a id="canonical-2010031312303002-3211110123333130-0010300013013233-0013113130331021-0230330333102033-1312101213130333-1323333222201002-2003221002131312"></a>

<a id="canonical-2030320012101131-1300333230032000-1002200202220003-0111013322213210-3300301001210101-3020003232110033-3332211311230002-0112302301312223"></a>

## regular expression property — path / 223102212122 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-1312202302221320-2012321321310201-1110233301103210-1013221201202323-0320033131222233-2313011301123020-2030310300131001-1001311321023113"></a>

## Next pages — path / 223102212122 / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-3320331332222212-3021210213121011-3300022311201121-3330023321222033-3310010221010200-1031132131022113-1123013302122011-3312233002000331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323130012001011-3330023213311223-0131100031230021-0233323310130301-0000320321210211-2212312310031202-3300012010000132-0121301323232302"></a>

## CloudFront.js_insertion_rules.rules — rules / 102301312322 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- CloudFront.js_insertion_rules.rules

<a id="canonical-3112211323200012-1223322323223121-1221102013102200-2231133101331021-2300001032233333-0231120330103210-1212210332300220-1012311313022321"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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

<a id="canonical-0031300302002113-0222303132212322-2110010031200310-0123101103301230-2121233201013311-0000201221012103-0001310231203023-0111210033011121"></a>

## Direct properties — rules / 102301312322 / 3

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-3302002122001011-3302130021132031-0001333011033320-0300322020323210-2313001110120312-2113031232021210-1112023131331222-3112330121301233): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-1210011000220311-1313103323102110-0102230012221021-0202221001203120-0030001013011111-3312221132223000-1310210311202010-1011131003202032): complete subsection reference.

<a id="canonical-2330022230132211-3221002200113200-1100131123202310-0210100330031012-0221112021331332-2011101230213322-2121232033121311-0332111332222022"></a>

<a id="canonical-3222303212231322-1300221330222302-3312031222013312-3202212031022132-2032331303331302-3013003101021220-3212303032210233-3021200022010011"></a>

## exact_path property — rules / 102301312322 / 4

Type: `"string"`. Computed.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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

<a id="canonical-0213003102103301-0231133303321310-0010223000220100-0031033033231303-0310101301133333-0111300200213321-3331312121120012-3231012303221121"></a>

<a id="canonical-0321202301023223-2110030222000333-1103222012303202-2001130121002200-1102210020132323-2313301132002130-2120301102303223-2231221013330222"></a>

## glob property — rules / 102301312322 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

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

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-3113000122001003-0110000112301311-0321320301102330-1123221321102312-3112011333030130-2220002310032301-3012230130300020-2212230302131021): complete subsection reference.

<a id="canonical-1102112322211030-1110013213130020-1203320202302331-0111311100132202-2011023302003030-2110102210033222-3022310112030302-0120230202210230"></a>

<a id="canonical-3023032322332031-2303103211311322-1210230000203132-0122110031123113-2122220021201331-2311101123131302-1030023202222101-2310111302023002"></a>

## prefix property — rules / 102301312322 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-3302102132233120-0030213100221310-1300113333101222-0221103331030330-3020020102130112-2211322312332000-2020313300033112-0112111110100222"></a>

## Next pages — rules / 102301312322 / 7

- [cloudfront.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-002.md#canonical-3302002122001011-3302130021132031-0001333011033320-0300322020323210-2313001110120312-2113031232021210-1112023131331222-3112330121301233)
- [cloudfront.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-002.md#canonical-1210011000220311-1313103323102110-0102230012221021-0202221001203120-0030001013011111-3312221132223000-1310210311202010-1011131003202032)
- [cloudfront.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-002.md#canonical-3113000122001003-0110000112301311-0321320301102330-1123221321102312-3112011333030130-2220002310032301-3012230130300020-2212230302131021)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3302002122001011-3302130021132031-0001333011033320-0300322020323210-2313001110120312-2113031232021210-1112023131331222-3112330121301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230013232003210-3021200213200131-1110233003022012-2202022323002131-2030303202310321-2310223211003032-0333303000301100-3011113102203201"></a>

## CloudFront.js_insertion_rules.rules.any_domain — any_domain / 223101232112 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.any_domain

<a id="canonical-0222003222210323-3202111102203220-1102211320120213-0333111312123211-0013112101330113-2223033210103231-3112222122223112-2010012010101013"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0311102101332132-1201123023102022-2302110321211211-0000303112201130-3002123232010120-0321012303022030-0102100113323312-2100211132202001"></a>

## Direct properties — any_domain / 223101232112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103133223103301-3002312131133210-2213232131003131-2000233331003002-3030222023030013-0333121100302311-3311102322022113-3010102202122320"></a>

## Next pages — any_domain / 223101232112 / 4

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1210011000220311-1313103323102110-0102230012221021-0202221001203120-0030001013011111-3312221132223000-1310210311202010-1011131003202032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321112102323001-2023121011221302-3020010102113221-1331022200313322-0012022103321131-2013330111213213-3203003032323111-2312033013311302"></a>

## CloudFront.js_insertion_rules.rules.domain — domain / 323221101120 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.domain

<a id="canonical-3331221212300133-1223211303033303-2010130322021122-0102122012131310-2011202003022123-0013311223221032-0102102233320230-2132212220311310"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

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

<a id="canonical-3132100311102230-1033221303032220-3233103020202032-1310320301311303-3211100020231331-1031103313331002-2021022122301321-1232123210121001"></a>

## Direct properties — domain / 323221101120 / 3

<a id="canonical-2123230330211300-0002023101232120-2012122123131320-0213333223012112-2331100330233130-0213213130033010-0122310001122012-0232300002112131"></a>

<a id="canonical-2321032123221232-3131022202033203-0202310000122012-3111323011031333-2203002211223313-1232323303213023-2112312203030100-3103330120321013"></a>

## exact_value property — domain / 323221101120 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-1311132312311103-1133212022210313-2102113332101133-1122331100203012-3131133212210323-0212122022220222-2303200310303020-3221003033000202"></a>

<a id="canonical-3211103200201010-0223122210231023-2303101210023211-0301002121022311-3102001023130223-3020230123002113-3321303013002221-0200102230103100"></a>

## regex_value property — domain / 323221101120 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

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

<a id="canonical-3213310122031123-3320313130031030-1133132220220232-1023201321003121-0311032300010220-2232031033203332-1212300320100230-2123332311302032"></a>

<a id="canonical-1020020111332222-0110010100210120-0311130103000112-3003233121023112-2201201330033201-3031001011303322-0301113233202110-2012013113120021"></a>

## suffix_value property — domain / 323221101120 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-0132013113003103-0211201103323333-0122231321021311-3233112231133212-2223222111322203-0021100133121213-1213000200320130-1201113321312223"></a>

## Next pages — domain / 323221101120 / 7

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3113000122001003-0110000112301311-0321320301102330-1123221321102312-3112011333030130-2220002310032301-3012230130300020-2212230302131021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101231312130121-3030101132122012-2120330130212313-0203332300232130-3232130311303332-0133200223333221-1323311100323212-3311201210202023"></a>

## CloudFront.js_insertion_rules.rules.metadata — metadata / 113000333002 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- CloudFront.js_insertion_rules.rules.metadata

<a id="canonical-0022310302103031-0331231102030333-3120321222311312-0020131003331111-2212033323301313-1302223032031212-2331130022200020-1322331013311101"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-1223130322020133-2333011113232201-2013310323312101-2303213100101103-1202213111102320-3133230021122031-1212323302213201-2130333312231003"></a>

## Direct properties — metadata / 113000333002 / 3

<a id="canonical-3001231312311121-1312231001320021-0003313103123212-0330312112110303-1231003321101131-3301230303111333-1231030120122132-2131032220323020"></a>

<a id="canonical-0220311203032022-2230113021230331-3000322113003222-1101100130132220-3331223203121023-2032111220200020-1001011232213303-0303032011033302"></a>

## description_spec property — metadata / 113000333002 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2123133012212133-2310202220311233-0213230023302321-3201300321300013-3210203312211110-3112113210112021-1303030212113023-2030110203023211"></a>

<a id="canonical-3013102000311203-0103201000203321-0331322220333221-0020132112222030-1313101201022200-1100312301012110-0212231302101111-2231200201202001"></a>

## name property — metadata / 113000333002 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-0233333000202303-0133012113112001-1322103200332103-2332121320001000-2221230203130110-0330103020212020-1202331320233130-3032301302033323"></a>

## Next pages — metadata / 113000333002 / 6

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-1323310101232211-3022030330331220-3221213030103111-1210123300102021-3000231032122231-1011332222123312-0303323200120223-1203201223210122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3123213201022200-1010003020110031-2133001121200133-1103021210123222-3223203013201202-3103020013112213-0232333101021031-3332032331023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101300002012313-0033121211300032-2202123331002213-0100300311012321-3032131122310020-0213023220313320-3133313102333131-1223233231203300"></a>

## CloudFront.manual_js_insert — manual_js_insert / 312121102303 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.manual_js_insert

<a id="canonical-1231011133010223-3122230322022221-3021222003233331-2101311111113001-2313033033311320-2101020223111333-2213123133202022-0212231202212132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3030212332221130-0011022121110133-1223122221333002-1210100133312022-2123301101211302-2010303111021133-2331230111102100-2133212331300033"></a>

## Direct properties — manual_js_insert / 312121102303 / 3

<a id="canonical-2031223130012320-1002023220033121-1123303221123200-2103321031032013-3223103101130120-3230211232313203-2030202220003312-2011020133131232"></a>

<a id="canonical-3300233133010303-2130101021001220-0213323222001202-3233101331210033-2033312010123112-3130203010213023-1011121012011012-1331210131011310"></a>

## javascript_mode property — manual_js_insert / 312121102303 / 4

Type: `"string"`. Computed.

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

<a id="canonical-0313331030133212-1032111323032303-0100021032110220-3232111112031033-0131011203020002-0330132231012123-3230212022110120-0201303001002023"></a>

<a id="canonical-0220130300320201-2120023113103110-0323121312110222-2231122012313002-3331303001323220-0302032130021332-3320123000113123-1203203200232011"></a>

## js_download_path property — manual_js_insert / 312121102303 / 5

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/CommonJS’.

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

<a id="canonical-3001230221211330-0213132000001121-0112031102011100-2122100130221122-1010210333223310-2313021133323201-1121323300332022-0101130223030332"></a>

## Next pages — manual_js_insert / 312121102303 / 6

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101210222212033-1310131022121100-0300201032122323-3203320310100110-3100033222011121-2002212001000323-3211202132023303-0100310303303102"></a>

## CloudFront.mobile_sdk_config — mobile_sdk_config / 330113201213 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- CloudFront.mobile_sdk_config

<a id="canonical-1312021210122012-0331300021020123-2230203331133203-1221133112213030-2301102122213221-0222131110123110-3321211320210302-3312002113320132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0201023133231231-3010230310310321-0222121210032011-0130201013120321-0113032223122223-2211202331321011-1221023222222312-3131322132323213"></a>

## Direct properties — mobile_sdk_config / 330113201213 / 3

- [mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031): complete subsection reference.

<a id="canonical-2212311021120002-3212212233230210-2112213233202332-3202021133231100-2030202131112012-1012021023020001-3033312000201312-0221001102133011"></a>

## Next pages — mobile_sdk_config / 330113201213 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122201120303332-3322333012312323-2120002122121222-1110313331212221-1200313102220103-3122302020320112-2200230000101123-0203000302210330"></a>

## CloudFront.mobile_sdk_config.mobile_identifier — mobile_identifier / 011322301021 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- CloudFront.mobile_sdk_config.mobile_identifier

<a id="canonical-1330032323232100-1320130301232230-3123220222202301-2303203231131111-1230023201202201-3120001300002310-3121012232233113-1220110213011102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1012220122130013-1012130223102103-3231122031121301-0033112113011112-1122001330112320-2101320011222230-0132231302302211-0330201200132122"></a>

## Direct properties — mobile_identifier / 011322301021 / 3

- [headers](data-sources--protected_application--reference--group-002.md#canonical-3222012320102033-2322223111110322-2320222210322101-0212011322302131-3020132003323230-1330132321001033-3303333032202113-0331323313020033): complete subsection reference.

<a id="canonical-2020020020030132-1001112111021203-2021313320011121-2032223013020310-0011002033132223-2221332101000201-3202320010101230-0011301223111232"></a>

## Next pages — mobile_identifier / 011322301021 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-002.md#canonical-3222012320102033-2322223111110322-2320222210322101-0212011322302131-3020132003323230-1330132321001033-3303333032202113-0331323313020033)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3222012320102033-2322223111110322-2320222210322101-0212011322302131-3020132003323230-1330132321001033-3303333032202113-0331323313020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133003031002333-0133310000212000-0121201100211320-3303202221131130-1120322320220200-2121230110120132-2013213133130132-1230221301130222"></a>

## CloudFront.mobile_sdk_config.mobile_identifier.headers — headers / 020133312232 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113)
- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031)
- CloudFront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1322312110312033-2320011333131330-2123130100300003-0333021023212202-2333312330312103-3120323300202121-2023102131330130-2300112230302233"></a>

Type: `"list"`. Computed.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

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

<a id="canonical-0212122013202101-2200213310313001-2233201213320320-2131020031203031-3230210111230100-2302133133323212-3003311100102133-0220022021222122"></a>

## Direct properties — headers / 020133312232 / 3

<a id="canonical-2322223030021221-0033131321320231-0331223011022122-3102122102020112-2120010302122123-1021200210122323-3032003103331012-0113323213102132"></a>

<a id="canonical-3011023303033021-0311012210013000-2311103223020312-0232002310202010-0123120122312132-1031003010210200-2012323213003222-0113010123000031"></a>

## exact property — headers / 020133312232 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Header value to match exactly.

Upstream description:

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

<a id="canonical-0102231110012330-0103310313211102-2013233022022133-0120310111110032-2330013120001022-3132122312130133-1210211300111203-3121103301110112"></a>

<a id="canonical-1221203010123012-2123032331321130-3312102321012202-2312233322303033-2201301232300323-1202033221031303-0131332200233021-1103021311300300"></a>

## name property — headers / 020133312232 / 5

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-0332331133103302-3012133233310233-0332121101132122-1201001030222030-0012031130123333-2030011232113130-3131232201021223-3220310003032102"></a>

<a id="canonical-1133213013320210-0101232101103320-1223100122200113-3122012302021010-3302133011010120-2330021222320131-3223222232130303-0320311300333123"></a>

## regular expression property — headers / 020133312232 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

Upstream description:

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

<a id="canonical-1133030212330010-1000222030222032-3302301121200311-1122232131311001-1103111212031112-1323111023200121-2320202001101302-0300000101011303"></a>

## Next pages — headers / 020133312232 / 7

- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
