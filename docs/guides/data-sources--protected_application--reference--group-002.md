---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-c4bb220b5c3b133c1895aeff1e9bd18101869ecf2dd76ffc40c91784fe425b76"></a>

## cloudflare.protected_endpoints.mobile_client — cloudflare.protected_endpoints.mobile_client / 6f2132d2d1ae / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.mobile_client

<a id="canonical-44a3f339c558b2dc801835a9e1b3bdc14b284ff96b7afe1f7237897101e4873c"></a>

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

<a id="canonical-fc4b58c316279200e8978236986374ded979802c72a6aeb931c1f7272fa8787e"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client / 6f2132d2d1ae / 3

- [block](data-sources--protected_application--reference--group-002.md#canonical-90e416498574ebda41e4e27f7f2b2a294334bcd0841f7c1e180938aef66625e0): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600): complete subsection reference.

<a id="canonical-18a05393516e0a8e8d21954ade56c08ea6b3c42138be1900a1f7c14c1cbc764c"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client / 6f2132d2d1ae / 4

- [cloudflare.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-002.md#canonical-90e416498574ebda41e4e27f7f2b2a294334bcd0841f7c1e180938aef66625e0)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-90e416498574ebda41e4e27f7f2b2a294334bcd0841f7c1e180938aef66625e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dbf40cf4b5938885abf3c30db9129a56b091fe58e3d3349dd18431e2f0027ee"></a>

## cloudflare.protected_endpoints.mobile_client.block — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- cloudflare.protected_endpoints.mobile_client.block

<a id="canonical-05339b8fd6f300eaa3dc04ba0267a908a6e6148ce57c57a7f01c9143bf992e41"></a>

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

<a id="canonical-514cfb3d17d607b020c7be6d6fcae8c1553e94a0166c1bd489d6a09e404d7dc0"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 3

<a id="canonical-c298267a3dd71c950610a5cdac92916ac3818abe7a13c278fe33039633112d2b"></a>

<a id="canonical-08e467aee3ecd8a05db727eb8d22eacaf46992341048dc2b55ae52b86a003947"></a>

## body property — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 4

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

<a id="canonical-5e42d1fccba4c34090ed512e453f5fe6fbe4bd96a887eaf3ee0afa0dc02d834a"></a>

<a id="canonical-510861ca9cb98d0484cd1037a5dfe3415cd39f44bb578219f5b9f93c886bbcf4"></a>

## content_type property — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 5

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

<a id="canonical-3c3ff138d0085e0a9c42c3b06afa619b0f133f21e5bddc6340434bee6d55523a"></a>

<a id="canonical-549e8050d7d390e1ad6194dba0118954cbbe07056a45ede93937a8aa59bcdc40"></a>

## status property — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 6

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

<a id="canonical-3239563c02a9bc71da47374954dcf401ef4d6c24ae7982bd5302b433fc4aaec8"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.block / 4267a859923b / 7

- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee22dd68cde82ca27aca22664103a163973ad6acf05b3124381469691dc51f72"></a>

## cloudflare.protected_endpoints.mobile_client.continue — cloudflare.protected_endpoints.mobile_client.continue / 247e3efff835 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- cloudflare.protected_endpoints.mobile_client.continue

<a id="canonical-4baf9d24fd83c5b786f03d59fd9dbe205f34854dc366dc427740c7442e9332b4"></a>

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

<a id="canonical-0451e84e92c5407365ef244d4dc4a69e209e40ff0460cb1b54ad01bf7775a130"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue / 247e3efff835 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-e84062fc7b34c7265db6dd296648f858a597b7d45ab37a8c66b451dcd453a86c): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-667e15aa70ce16511eb8d59d743a6df62df9143f8686004b35ff798118587b35): complete subsection reference.

<a id="canonical-adcadf2015af50dd0f814ccc9c9f6ccffb3b590bc2beb635c9c1a64be428d3ca"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue / 247e3efff835 / 4

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-e84062fc7b34c7265db6dd296648f858a597b7d45ab37a8c66b451dcd453a86c)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-667e15aa70ce16511eb8d59d743a6df62df9143f8686004b35ff798118587b35)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e84062fc7b34c7265db6dd296648f858a597b7d45ab37a8c66b451dcd453a86c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-616a20499dfcc2f49bea3eda0e6421b68a2e395c7a17e17cc430d881e048568c"></a>

## cloudflare.protected_endpoints.mobile_client.continue.add_header — cloudflare.protected_endpoints.mobile_client.continue.add_header / 6879edfd127d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600)
- cloudflare.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-af7687d1f896dc7006082819793062e5758ce2cfc83b1be4f220ab9fd12d3b02"></a>

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

<a id="canonical-20a96bcf01035a58de0c1ac4e60f5e7aacfa423eda94dbba70dcaf593af02f88"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue.add_header / 6879edfd127d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d79176fd25f01fd6d716bcc830c986402ba44b31cfe883fc140e3a1369e46738"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue.add_header / 6879edfd127d / 4

- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-667e15aa70ce16511eb8d59d743a6df62df9143f8686004b35ff798118587b35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c808c263c91e5a0c3bb1c9386c916d0b5dcdc9c2a4cf31b1e5f339d1795865e"></a>

## cloudflare.protected_endpoints.mobile_client.continue.no_header — cloudflare.protected_endpoints.mobile_client.continue.no_header / a4eccae4244e / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600)
- cloudflare.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-dbe90f6da0c999382a59fad6c5ecea6785957b6fc957bb6198a5de56fc09064d"></a>

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

<a id="canonical-13174b5713cc5fdacb42e11ef68317bec42dd6c9d54572cb265344a31d01421d"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue.no_header / a4eccae4244e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b7aa72fdd7ad8c4d97aa04a1d46a3168953c6f02574ab2233cdf91ba2c203bb"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue.no_header / a4eccae4244e / 4

- [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-35e1a80aa739b25652af5aed24ce3e4801bcb6cfc2a26279fed449373785a600)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ee8706ed2b875ad7520a81346808930b223ea48f0e68b428b74abd39f4ef1069"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ae8052e87cfb0a53812b5698755ed760287ebb32356b7924663d28b83285a96"></a>

## cloudflare.protected_endpoints.path — cloudflare.protected_endpoints.path / 5280b7ebd7f4 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.path

<a id="canonical-407ba3d6edeb64866af26299302743e4c72506d6734feb6d9f4cede60c5ace03"></a>

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

<a id="canonical-92607ae7b4d7108dcf242fd33c370d064fbeb0cb4d43d507c6289a4ba41de676"></a>

## Direct properties — cloudflare.protected_endpoints.path / 5280b7ebd7f4 / 3

<a id="canonical-36ba7cbdb8db839d58080aaf37bfa9ed4dafabe2ccfde5dc60b4bc45d86b0b71"></a>

<a id="canonical-ec72f89af79648448b25325ccb489f892827f671986098e4147a66770f4ca148"></a>

## caseinsensitive property — cloudflare.protected_endpoints.path / 5280b7ebd7f4 / 4

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

<a id="canonical-7c9da84c0f9714ed7436ca3b4993a8ef1b1e62f185e035225300da90b9f249fd"></a>

<a id="canonical-12b710731e414c94f432aad512995c004083e9d44facd4747d7e13e672ce1ba5"></a>

## path property — cloudflare.protected_endpoints.path / 5280b7ebd7f4 / 5

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

<a id="canonical-ce835d2f0d22d3fdae8118afe8bef32c6d243cb1bd0a0a8a3fa02f97fc806d5f"></a>

## Next pages — cloudflare.protected_endpoints.path / 5280b7ebd7f4 / 6

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-850e553e60aca2539678693bda545e816e1d221c1ca5a3665b8f80e579e9d921"></a>

## cloudflare.protected_endpoints.web_client — cloudflare.protected_endpoints.web_client / d5dccde9f2ab / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.web_client

<a id="canonical-9b777c9ee64a564f5efbe0990817775526e74e777d7b7c79dbaf61a4be45fabe"></a>

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

<a id="canonical-9d2c1bfeb19896b4fbae481a84f7d1ed9b32be4d4e33fbc1f9590745049b439d"></a>

## Direct properties — cloudflare.protected_endpoints.web_client / d5dccde9f2ab / 3

- [block](data-sources--protected_application--reference--group-002.md#canonical-2365159aa8b1d69abde0d172ce5ad65808060bd244c501b43436a04303128fba): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c): complete subsection reference.

- [redirect](data-sources--protected_application--reference--group-002.md#canonical-1bde5d081b24af5d8d0a20793c4dba612713e02cbe3884ca42ee717a2b2ab8b7): complete subsection reference.

<a id="canonical-30b11a6184469db4043c57d49e3c40c76202bb5036b64533b0ba6321539ed6c6"></a>

## Next pages — cloudflare.protected_endpoints.web_client / d5dccde9f2ab / 4

- [cloudflare.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-002.md#canonical-2365159aa8b1d69abde0d172ce5ad65808060bd244c501b43436a04303128fba)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c)
- [cloudflare.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-002.md#canonical-1bde5d081b24af5d8d0a20793c4dba612713e02cbe3884ca42ee717a2b2ab8b7)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2365159aa8b1d69abde0d172ce5ad65808060bd244c501b43436a04303128fba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6862af315aab2202443f3ecfb21b0c566870aebd041a517e2210d5a4874fe747"></a>

## cloudflare.protected_endpoints.web_client.block — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- cloudflare.protected_endpoints.web_client.block

<a id="canonical-14e52a69c3b183602ac8194aff8bba90ea78e7d239263ac4caa3d9009db31e23"></a>

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

<a id="canonical-17a54391e66f81d0ba0854a9c3558263672b43436731133fb1e579ad2fb09d4a"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 3

<a id="canonical-550670b1f51e061ae78e83465e3bf694dac25465de8a6f31c33c893065de5b84"></a>

<a id="canonical-4b814a28efe6d66adc7a1675a41136d40dca14c1cd3734dc488b1c04d56c7ac6"></a>

## body property — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 4

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

<a id="canonical-4700cce8fea3916e54f86e77a48feb73c3f4a61524469c686e6808e888573162"></a>

<a id="canonical-fd6aae51310874849dfd36aad8a8d4a82cd5ff8e7643e2635cb3f714b8cc4b95"></a>

## content_type property — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 5

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

<a id="canonical-bbf240d50533d61745e35bdaa16d9a8d1d718ef3277636a21e8707561db9fa34"></a>

<a id="canonical-04edd3535eb4166be5cc46ff5df2d8bb1420f03696359fe5e17c18e78881d110"></a>

## status property — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 6

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

<a id="canonical-9d04985d9c2f9ddea22338709585fed488ab348a2c95d166fe4149ce31396687"></a>

## Next pages — cloudflare.protected_endpoints.web_client.block / 42c51c350744 / 7

- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c7069a6f0a16f9230bcf1cf139cc330413a4b603d6c6b117b91151c6da5fc67"></a>

## cloudflare.protected_endpoints.web_client.continue — cloudflare.protected_endpoints.web_client.continue / b2dde1db9861 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- cloudflare.protected_endpoints.web_client.continue

<a id="canonical-cfad53267b1331c3e0f69ee9513f0da4f2bd3107ae965d008181ed677875b58a"></a>

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

<a id="canonical-cb86f5ee25dd4c17fd0b91b5f8e95c0d3207228fa73e436dc2918b0db8508910"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue / b2dde1db9861 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-a78e3f53049cc2423e9ee4c30e36d4f0be9a29f1644754c85d5921b0df032219): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-421b49500ee7900610f17c8888021a9c711e32159de5f4f867d187259ffd46f6): complete subsection reference.

<a id="canonical-4e7dd3a378cef6d0b5cc085597b6f4f4ea045b49c5c2d955d95efbfdd0c58269"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue / b2dde1db9861 / 4

- [cloudflare.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-a78e3f53049cc2423e9ee4c30e36d4f0be9a29f1644754c85d5921b0df032219)
- [cloudflare.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-421b49500ee7900610f17c8888021a9c711e32159de5f4f867d187259ffd46f6)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-a78e3f53049cc2423e9ee4c30e36d4f0be9a29f1644754c85d5921b0df032219"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1adbf789a36496adb86d40bda63e7329ee87b1126019034d307e8ac37e3df9b"></a>

## cloudflare.protected_endpoints.web_client.continue.add_header — cloudflare.protected_endpoints.web_client.continue.add_header / 111fa25939ed / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c)
- cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-b2aa5c68cdd1274d53d2937e5c9b6aa73fad637f0bd225f42bbc7599387accc1"></a>

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

<a id="canonical-c0c23c6bb79cb35ec9f1a9b37642d6bad069931faa96d5b57652914c1647db71"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue.add_header / 111fa25939ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29c8f44dcec0910e99aab0c746dcd8fdc245a8eadb4e3d675304e7fb96ebb990"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue.add_header / 111fa25939ed / 4

- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-421b49500ee7900610f17c8888021a9c711e32159de5f4f867d187259ffd46f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdf1f6438137cec9f70e265e03b1eac665dd56b45597ec398ad1abfd3d4a2b9f"></a>

## cloudflare.protected_endpoints.web_client.continue.no_header — cloudflare.protected_endpoints.web_client.continue.no_header / 92fd4fa4cc76 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c)
- cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-eac1d24b4d4a3bf5d0ef5f053c5cb8cabb723df20fbe871d82840dddeaad82c7"></a>

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

<a id="canonical-28442122bf1e87a429cfba7aee486554bb265514f64b131c34bcef27770026b5"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue.no_header / 92fd4fa4cc76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-acd7712a3543fa1a2ab52a727245f0b217f0833f2945800d4e6179d081572709"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue.no_header / 92fd4fa4cc76 / 4

- [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-e7876ae4297620fda853364f95cd58bf90e2ab3aea700e79785155e3359adf4c)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-1bde5d081b24af5d8d0a20793c4dba612713e02cbe3884ca42ee717a2b2ab8b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61c9a06f561f1e8ba4344037fbf4788f3b0ef91d915c543a8b53bd2140d75385"></a>

## cloudflare.protected_endpoints.web_client.redirect — cloudflare.protected_endpoints.web_client.redirect / 7aacf5795c8d / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-e78d54d2783cb48ddc19a19071ba9220b81b4804276d1dc1d81d36541a34f19c"></a>

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

<a id="canonical-9c7d21b0284cd4eab2d833b472f3684697cb8e0c3dd2a2e60535160534112b24"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.redirect / 7aacf5795c8d / 3

<a id="canonical-c3d89f9afde72303f2cc7a8845d14303842e64aab789bf8a9c2fa56c5dabea98"></a>

<a id="canonical-6b8da96aabc2823ef9a72e132bce0623c29d0e8de34193022bde7d9b65ff17f7"></a>

## location property — cloudflare.protected_endpoints.web_client.redirect / 7aacf5795c8d / 4

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

<a id="canonical-3dbc7e6dac6b97c12941f9f9fdb88352bde8b4170d01f1ff540d20638ce27f8e"></a>

<a id="canonical-7aeeb54faa1923749de2c358ef19cfc1b1a8c3eadafd4cb49a9f8a627a0bbdf7"></a>

## status property — cloudflare.protected_endpoints.web_client.redirect / 7aacf5795c8d / 5

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

<a id="canonical-e35a027d188c08b6a5b73624c179f8dee5ee2fd7a3b4bb4c19f7d32065fee1bb"></a>

## Next pages — cloudflare.protected_endpoints.web_client.redirect / 7aacf5795c8d / 6

- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-262615f810c37e22c0615a08f3b546469f5f6170dca96436b6dd7366fe5e4022"></a>

## cloudflare.protected_endpoints.web_mobile_client — cloudflare.protected_endpoints.web_mobile_client / 1614d4bda7b7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.web_mobile_client

<a id="canonical-965b182fadd91510c266cf34ecd9e54d3f60bed9bec6ad0c2678b981a959c435"></a>

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

<a id="canonical-d2cfb433b8c96155c2422aebf0a206ff150f1082cb35ee34d6852e7c372855aa"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client / 1614d4bda7b7 / 3

- [block_mobile](data-sources--protected_application--reference--group-002.md#canonical-e86a6558702f6c44438824b0b672fe6f447db541f9235b1f7b9f3fbb9e3df8f5): complete subsection reference.

- [block_web](data-sources--protected_application--reference--group-002.md#canonical-e8e2c9907e5e5c5d9bdeb450ed8ecd8d27021f3e709235ba2c221d4886870411): complete subsection reference.

- [continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb): complete subsection reference.

- [continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248): complete subsection reference.

- [redirect_web](data-sources--protected_application--reference--group-002.md#canonical-eadb5fe6f399e0ede224817a86cc2f04bfae31b7e176e6be0a7a9a6a6c7344df): complete subsection reference.

<a id="canonical-39b01d978e01f50078e3e63ed227b7b4fb24f71ec5375563d6811b2e3b534040"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client / 1614d4bda7b7 / 4

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-002.md#canonical-e86a6558702f6c44438824b0b672fe6f447db541f9235b1f7b9f3fbb9e3df8f5)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-002.md#canonical-e8e2c9907e5e5c5d9bdeb450ed8ecd8d27021f3e709235ba2c221d4886870411)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-002.md#canonical-eadb5fe6f399e0ede224817a86cc2f04bfae31b7e176e6be0a7a9a6a6c7344df)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e86a6558702f6c44438824b0b672fe6f447db541f9235b1f7b9f3fbb9e3df8f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c4c85b53a2c54bbeb4db2615b2c81b0963baaeabcafbf11acf03a31c9c0dab8"></a>

## cloudflare.protected_endpoints.web_mobile_client.block_mobile — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- cloudflare.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-81214094e1b82cd285fd45d95bc4ea2c7c42c61aff04151e582fa926b12052d9"></a>

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

<a id="canonical-bccd48c4283634bb2f5ec01aa45447170452c67119235ad3dec9e66460829114"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 3

<a id="canonical-fa4a5241c578a70d7823767f9c21d6498cef49d617104be30a6d31a6c740e6f5"></a>

<a id="canonical-fd958c8905334eb5ba5b8890c7e5f81ded4982e273729e69bd38d0dfce3925f6"></a>

## body property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 4

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

<a id="canonical-c37430ca101b42399bcae70454e5e745ca6e7f79b3c760c96611ae941d932a34"></a>

<a id="canonical-e0b7170129ab93ef69ff97f44a03b76ac89bc21f0e66b97ba42f39f81eb7eabe"></a>

## content_type property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 5

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

<a id="canonical-7676a379f406c8aec645ee3a7e0bc3e60b81b09261ffdff3d9c0cf4dc91f0931"></a>

<a id="canonical-16db57d994db3cae00d7d89e69d74aabf7c283a68b52ac24245682a0fd083005"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 6

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

<a id="canonical-cfebc7d8fd0c335ed8ae997fe41df94f6b69a79a80ff43064b520543654b23ae"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.block_mobile / cff10dfbe3b5 / 7

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e8e2c9907e5e5c5d9bdeb450ed8ecd8d27021f3e709235ba2c221d4886870411"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e534121218d028c744ab968f77a6c5ab9123bb0901afc06a41a89e87e67ba9"></a>

## cloudflare.protected_endpoints.web_mobile_client.block_web — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- cloudflare.protected_endpoints.web_mobile_client.block_web

<a id="canonical-97bb0772834ee927c72c685d774aac0d21cc534aebb8ff89a92940763eb57c0b"></a>

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

<a id="canonical-36a6a2db73cab383a5b808e83e82f376b2a718806e4f01ac2e8a261abe599e75"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 3

<a id="canonical-9040c472e8877d8c6f52d3dc57384e005497475bd6783d1eed12e8c0ed9ebef0"></a>

<a id="canonical-bc5ac12be85ac52548162cad5f2185bbf37137e6d8c996023cf20f20b25bb954"></a>

## body property — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 4

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

<a id="canonical-0b6ef0b8cf5e01c0fcdb463c5fb191096d5341052783e8748792221114cd3a53"></a>

<a id="canonical-5780ebca57b2354b2021a06606c128b06904e533c9a08299c3b644f57dad7b01"></a>

## content_type property — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 5

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

<a id="canonical-6b574f802af64b5b71120bbc26330238dfedaca018f65d2b9d71268b12626f53"></a>

<a id="canonical-867736c1638aab17e746504fc9d1d1ee4644da3bad8e423fb9496f7fbfc4d66e"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 6

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

<a id="canonical-1f63b00692fcea330aff75160eda6d72cd71224292d6146402b7640d7ddbecf8"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.block_web / 2835784403ae / 7

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-976f856ecba49c0d75a48fc1899d7a94a6c3bec05ebd14ecd35ea0ddf30e488a"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 89a136ec1bd8 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-299042d5ad2bc29c51ce6cbacfb7eb1b909141158140b14ad95ce602ec801ad3"></a>

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

<a id="canonical-5b358f2e83319c3e67446f717eb4f1d5c10081b8a4574b2872fac4e280e42b2f"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 89a136ec1bd8 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-6fc8b37df7e810f72bce104c39489aa768805e23e7455490202d705dddd67e31): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-1d3ad7a223dcc3b4317a8f347c4f7e76eeab018e00172db4da5a68ad66dd8755): complete subsection reference.

<a id="canonical-770bbcdc6ec4f57a417e66a65961c307552cbccfb8d5030301096782445bc39a"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 89a136ec1bd8 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-002.md#canonical-6fc8b37df7e810f72bce104c39489aa768805e23e7455490202d705dddd67e31)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-002.md#canonical-1d3ad7a223dcc3b4317a8f347c4f7e76eeab018e00172db4da5a68ad66dd8755)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-6fc8b37df7e810f72bce104c39489aa768805e23e7455490202d705dddd67e31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08248741f9360f199eb927af5537636c28afbe2ed45c9a70800677ca0d50af6b"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / c6ade62147b8 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-3dcb00e18481a8ad1431bb06fed0b5f0f7226aea0abd7c8b6cf5ad0af6eff116"></a>

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

<a id="canonical-3848f92dd4617d0e5ab5ee8df2584c0e42cbabd087d3a1f825dd2bc0dc26ecc0"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / c6ade62147b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa19cbbd33d4bb12a86d1fcf576dca9f954c1caae0cb713ea38155162a2a637a"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / c6ade62147b8 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-1d3ad7a223dcc3b4317a8f347c4f7e76eeab018e00172db4da5a68ad66dd8755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8663b0b5e1cf77c003d3ad488d30b02c9cdb7520e68c3705145f0b4fe5f1160"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / f47aa7491cb9 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-38c8198083d0c3aa035ad8b60c493bb8ca07eeaa59ba2f0208406da79f953b08"></a>

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

<a id="canonical-1cc56ce5a33e86c916da87e8cf51b6d04a05802c673c7a3da0e992fe291c7223"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / f47aa7491cb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30878c54660e7f5c637ccd1061d27b87498eb7870be3f84b35dc5ebb911da13a"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / f47aa7491cb9 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-ef78b2012cc09c4316b41dcc43bfba123fdaddd71c7379202ee990c5ecdadbfb)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e4d1a86a91a8d36ee9819164b45302ddb8ffc3c3d8e92ec369877fd4917fc5d"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web — cloudflare.protected_endpoints.web_mobile_client.continue_web / 6721ae647660 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-685d88d93972f54f20153209adfe4896260085867082be4c8091400bba24435d"></a>

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

<a id="canonical-9f120e7f86aaf77374c98d3aeaeba90b5a424dcc0d4857daadc2a4bb26829fe1"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web / 6721ae647660 / 3

- [add_header](data-sources--protected_application--reference--group-002.md#canonical-393722d93c8ee24b05dc11791b17e1359cfec580587bcebb952311244d8aa0fa): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-002.md#canonical-2364a63e98f1f4e6f6087571cfb30524577153626d4dbd8d76ea1c6b19ab6260): complete subsection reference.

<a id="canonical-b50dc42db629634913f6ab1a8e6bcd7006f0b5da3fa9244a2c708ea943c8d0e7"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web / 6721ae647660 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-002.md#canonical-393722d93c8ee24b05dc11791b17e1359cfec580587bcebb952311244d8aa0fa)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-002.md#canonical-2364a63e98f1f4e6f6087571cfb30524577153626d4dbd8d76ea1c6b19ab6260)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-393722d93c8ee24b05dc11791b17e1359cfec580587bcebb952311244d8aa0fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b3374c0dd303454ce9a4271a6e11f469d7279bfc508651c06d2b825916ab26f"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / 8491efb8408c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248)
- cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-73dd28072a79af3c8af32375d52cae17052733043b17e5e4288b0cd9674967e0"></a>

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

<a id="canonical-a3dc948dacdb123779d46db58db62a30a405d30e7009620725dff8b47155568b"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / 8491efb8408c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02b4a864767d4240c6cb9335126d574f84a94790c46b88f794702917077098fb"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / 8491efb8408c / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2364a63e98f1f4e6f6087571cfb30524577153626d4dbd8d76ea1c6b19ab6260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c11168d62766347515ae204e342f59f665c2b7f746931c9808b8c4ed650232d"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 4b56b4ae49e0 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248)
- cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-bbd6a5999745d7a174d9df064f91c8f2284d3e3ccb51702ee72c4b208430509c"></a>

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

<a id="canonical-c79d712e1abd26f452493b560f593bb451fcf365e75f5f1685e96c1bafd2268f"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 4b56b4ae49e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f670e1f754e1ff6924709e5253438909df1feaa67dce2bb3297d0a1b874d28e0"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 4b56b4ae49e0 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1661985fb2361bfd88e0bc9727b40385069ec1b58d04a8633f8390ac47f8b248)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-eadb5fe6f399e0ede224817a86cc2f04bfae31b7e176e6be0a7a9a6a6c7344df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a001de2ec482d5617f98de14c3e6c2fc2fa035a3d8cc7870ddf262969bb1cd2"></a>

## cloudflare.protected_endpoints.web_mobile_client.redirect_web — cloudflare.protected_endpoints.web_mobile_client.redirect_web / 480ffd411961 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- cloudflare.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-fad10e9e853c5430a34e7319022f65787b76c5be2a40c63e8f0a56f71c0c3a0e"></a>

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

<a id="canonical-72dac5b05003fa30bec703576aef7777adec1baa3ed362e6767123ccdbef313b"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.redirect_web / 480ffd411961 / 3

<a id="canonical-2b4d0c535ca191a8ff6d5b63dca95a2fe201f68ce100d3e9d626cb6092b37594"></a>

<a id="canonical-a44b4fa5d5d07dc7c4bfe7ceadc0249b3b4e6e7c4f74f58c5a13ecd885639393"></a>

## location property — cloudflare.protected_endpoints.web_mobile_client.redirect_web / 480ffd411961 / 4

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

<a id="canonical-79dfd52e93cc687e2eca564b9175dad7ebb096453ea22e0528f8585fc893d96f"></a>

<a id="canonical-875e28183f4e3fb67e8882b32a2d50c6379f88654f8245e72303ea0eb7621d3d"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.redirect_web / 480ffd411961 / 5

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

<a id="canonical-a1a1aad0c81a7ae51ee115947d5ce008af4674e460fd3bd22ab33a16f0e89185"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.redirect_web / 480ffd411961 / 6

- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d4ab062db76a1822ab810c52f7af5707436be21122bd9c68cc361ed0fb50c39"></a>

## cloudflare.trusted_clients — cloudflare.trusted_clients / 51b9c846e874 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.trusted_clients

<a id="canonical-1f4ba568f5b810052c2413444a99be5a5b72f8d77bf24d48d4ef7ab1587171ef"></a>

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

<a id="canonical-4c1852894f7383d1c3ff85e14cf103fa8b75b660685241d18edc62d939065f80"></a>

## Direct properties — cloudflare.trusted_clients / 51b9c846e874 / 3

- [http_header](data-sources--protected_application--reference--group-002.md#canonical-e184649946894941af06d28bed59bcdf11b5209896ef70a7ea00f3e597519459): complete subsection reference.

<a id="canonical-df162deaade587bc90f361967843e3d752710cfe6614567845f6d2ff38462c4b"></a>

<a id="canonical-b3b04b69284340f0600c803fdc80d06e14bbf42105a0db39ed5592a108d85b2f"></a>

## ip_prefix property — cloudflare.trusted_clients / 51b9c846e874 / 4

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

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-82bf533309098a714ef5846e0e3cf37c4f78fd974f0d2f01d16687c53beda2a8): complete subsection reference.

<a id="canonical-5bb62ca646a3951b165baadbfcd1705b637c46c3e3bb480019193b957e4e0930"></a>

## Next pages — cloudflare.trusted_clients / 51b9c846e874 / 5

- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-e184649946894941af06d28bed59bcdf11b5209896ef70a7ea00f3e597519459)
- [cloudflare.trusted_clients.metadata](data-sources--protected_application--reference--group-002.md#canonical-82bf533309098a714ef5846e0e3cf37c4f78fd974f0d2f01d16687c53beda2a8)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e184649946894941af06d28bed59bcdf11b5209896ef70a7ea00f3e597519459"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d8ef9b8a9cee706c61b85f0d419361bc4f46b205ad1c5699f71626bf7d9b0b4"></a>

## cloudflare.trusted_clients.http_header — cloudflare.trusted_clients.http_header / 3dec49ef7db9 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- cloudflare.trusted_clients.http_header

<a id="canonical-fd7090a059b3616d3d43b6ff8d4fb53a6b1cfc1c81bdc5424984eabd81abea53"></a>

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

<a id="canonical-f329d1f6b63868ec43fa123351d376a91f4e598678e25690ce895235e163359d"></a>

## Direct properties — cloudflare.trusted_clients.http_header / 3dec49ef7db9 / 3

- [headers](data-sources--protected_application--reference--group-002.md#canonical-fc432376638bbb38571dc8d1654de829e2f783ba975a2b20c759366903808885): complete subsection reference.

<a id="canonical-ccd4d9a2a965664899c78f7b3c8f8efb354a75390f174ccf2f4d99630b573507"></a>

## Next pages — cloudflare.trusted_clients.http_header / 3dec49ef7db9 / 4

- [cloudflare.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-002.md#canonical-fc432376638bbb38571dc8d1654de829e2f783ba975a2b20c759366903808885)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-fc432376638bbb38571dc8d1654de829e2f783ba975a2b20c759366903808885"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd89dfd0ee6f6f73f517ac49d354b3f002a95fab262c446859bbe3fc9bbb91f2"></a>

## cloudflare.trusted_clients.http_header.headers — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-e184649946894941af06d28bed59bcdf11b5209896ef70a7ea00f3e597519459)
- cloudflare.trusted_clients.http_header.headers

<a id="canonical-e1f85056e09d0943782537343853132337dd0df7979bcd06076a1fd369be4834"></a>

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

<a id="canonical-17289d1a8482c386a54a2f94fee3642762376bc2985832af36fef1f1b22d9c93"></a>

## Direct properties — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 3

<a id="canonical-8d0741159ba82ff1aaeb2b8f9c2233f06dbdb62584384d12fce218c46aac968f"></a>

<a id="canonical-d22c0b69962aba8764a0bc6954c6726e80d837ec19bbb24e3a579923b115cfc9"></a>

## exact property — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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

<a id="canonical-4ace4b7c3059cbd06beef9a26dd5161484b283c7cec50e2c9696a8099b638b5d"></a>

<a id="canonical-2afacc2882a2bba818bb57ff0279412e6288a211094e7db59c3bb946c983bc75"></a>

## name property — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 5

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

<a id="canonical-554fcdaecbabf577c7a99a1eee36097b102f6dabb32a19b717ef3267a238b82d"></a>

<a id="canonical-2f2d4280902f2467ec789282d78d30ffc6df025d8f91265dcdd5a74a1be1c487"></a>

## regex property — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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

<a id="canonical-56e7e6a28b0c97c10ad616069081cdda9276462a5821676af0385ba9074f1105"></a>

## Next pages — cloudflare.trusted_clients.http_header.headers / 8da2ef865af3 / 7

- [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-e184649946894941af06d28bed59bcdf11b5209896ef70a7ea00f3e597519459)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-82bf533309098a714ef5846e0e3cf37c4f78fd974f0d2f01d16687c53beda2a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ad3f6a62cf0f390cbc953e8daef7318bc5c46e11c57de4e3cf7e4bf98ab7155"></a>

## cloudflare.trusted_clients.metadata — cloudflare.trusted_clients.metadata / 048f67fc63ad / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- cloudflare.trusted_clients.metadata

<a id="canonical-241c504f935e1fb347922dc3e7016667eef6b5a676d2f9ae170a8da20814d535"></a>

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

<a id="canonical-962f4eec0bf0883cf67473b5fc56ca72615f73f0d1ffa850800ee4b69fc741cf"></a>

## Direct properties — cloudflare.trusted_clients.metadata / 048f67fc63ad / 3

<a id="canonical-d281605032e15770c12e2f2d91fbc8928814ccf68d888e10e12e4370f3bd37e4"></a>

<a id="canonical-21872c11d863c0f2d6d7a3809869e4fd888b60735554f4333201cc9cd3135516"></a>

## description_spec property — cloudflare.trusted_clients.metadata / 048f67fc63ad / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-c8462c2e1351e2ea157a7f071e0bff778435e8710ea83ca573720c36c210e381"></a>

<a id="canonical-f1de45edba054b853dbb681f9111e114a48e4254bd599feb65281eaf2ecaf955"></a>

## name property — cloudflare.trusted_clients.metadata / 048f67fc63ad / 5

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

<a id="canonical-95c42e83f6da0e1abc06d01f782d143b4a843d61b79a6a1344aa72122cb67235"></a>

## Next pages — cloudflare.trusted_clients.metadata / 048f67fc63ad / 6

- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24663219fba40961486bd58052bff6f6ea5bcaf942936fe433595fb937178b50"></a>

## cloudfront — cloudfront / 4bbdb7e077a0 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- cloudfront

<a id="canonical-41b247837668c727fb5f889326104b14970f4fbba0b3be090903fb9277c9aa42"></a>

Type: `"single"`. Computed.

Bot Defense policy configuration for AWS Cloudfront.

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

<a id="canonical-8af632c62adde6079776f10de459e4422d997778c5eb01adb3ccdb205b191f96"></a>

## Direct properties — cloudfront / 4bbdb7e077a0 / 3

- [aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-c36a85ea5d39105ce13d473094312cca2b9561873b991fba6480e9fd3f22767a): complete subsection reference.

- [aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-a4a729753c11376ce56cff36eb4091c35e7764fa5429ab331cbfe66b6d43fd7d): complete subsection reference.

<a id="canonical-713410c7a973ab47d007e8cda92fc79796d0fddae34445652281b07626945bc9"></a>

<a id="canonical-908c25d9de9c84312754aefae9516f6ff3bfe89d4871a358f67e368f138c25e3"></a>

## continue_mitigation_action_hdr property — cloudfront / 4bbdb7e077a0 / 4

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

<a id="canonical-5e4bbb61a2123d8eaafeaa5deac0a34501b8bc6d91161be11bf64617d69d7454"></a>

<a id="canonical-6e80323bf3de2e1ff2b30427745b661219393b83972b7fd856b3f73201a9a649"></a>

## data_sample property — cloudfront / 4bbdb7e077a0 / 5

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

- [disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-1726345029492ab4c9ebc6f72c73146b297226d4ea16cd63163cb1745b5bcff7): complete subsection reference.

- [disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-db631d35eae4903696cc487a003fd24f41f038e41ca8789c489ee644ddea7baf): complete subsection reference.

- [disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-4707d07fdea3b9a915148c99f771232f176fece1ec69d5d33edf84492dc0766c): complete subsection reference.

- [js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0): complete subsection reference.

<a id="canonical-a137552ac01ee7d9899aa4e4fd078ba03a62b5663650fc1c6ef7c4ccd64aebc5"></a>

<a id="canonical-00a898048521927f9c7e0be9192f6f5bcf40c6f2ec91769b29daadb1c3a67d1b"></a>

## loglevel property — cloudfront / 4bbdb7e077a0 / 6

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

- [manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-db9e12a0440c850d9f05981f532646eaeb8c7862d32075a72efd124dfe3bd2f4): complete subsection reference.

- [mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017): complete subsection reference.

- [protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0): complete subsection reference.

<a id="canonical-71b05a4526574b2d94f6e8635f1401b5389ce9105efbc2e299f0127a034b5507"></a>

<a id="canonical-5be0bad77e59f4a12595fb04b8beeeb36117341f41383dd8a4dc30746dddd3e0"></a>

## timeout property — cloudfront / 4bbdb7e077a0 / 7

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

- [trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-dd0c1b728601671aee10a15c8aaaa84a60bbd5a8cab02d5dc4a3cec65ae4fcca): complete subsection reference.

<a id="canonical-15ad908c6488d2b0076da7bfe85199f88db64962cf2530256bb77616d8bf6d19"></a>

## Next pages — cloudfront / 4bbdb7e077a0 / 8

- [cloudfront.aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-c36a85ea5d39105ce13d473094312cca2b9561873b991fba6480e9fd3f22767a)
- [cloudfront.aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-a4a729753c11376ce56cff36eb4091c35e7764fa5429ab331cbfe66b6d43fd7d)
- [cloudfront.disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-1726345029492ab4c9ebc6f72c73146b297226d4ea16cd63163cb1745b5bcff7)
- [cloudfront.disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-db631d35eae4903696cc487a003fd24f41f038e41ca8789c489ee644ddea7baf)
- [cloudfront.disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-4707d07fdea3b9a915148c99f771232f176fece1ec69d5d33edf84492dc0766c)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-db9e12a0440c850d9f05981f532646eaeb8c7862d32075a72efd124dfe3bd2f4)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [cloudfront.trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-dd0c1b728601671aee10a15c8aaaa84a60bbd5a8cab02d5dc4a3cec65ae4fcca)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-c36a85ea5d39105ce13d473094312cca2b9561873b991fba6480e9fd3f22767a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f49843e2df9ee7bb3c22811b722ce0025b20b5d095931eff7a4cff01522f71ac"></a>

## cloudfront.aws_configuration_id_selector — cloudfront.aws_configuration_id_selector / f3627ace5558 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.aws_configuration_id_selector

<a id="canonical-281482fa84436ea871a92c64da4fc88ca3419ce012cf56ba6884d8549ef99d97"></a>

Type: `"single"`. Computed.

Configuration parameter for aws configuration id selector.

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

<a id="canonical-c9458f148954b88e1e106b6ec32c2580b95e2f8317017c1776100ad8232313b8"></a>

## Direct properties — cloudfront.aws_configuration_id_selector / f3627ace5558 / 3

<a id="canonical-267370a3bc2ae2d297fa45d49848460d8a53afcb496b14ba721d7de1b82659e0"></a>

<a id="canonical-1a2e97fc2f9a23b00561637ebc02d15302dc22edb92001f54087610be8ac08b3"></a>

## ids property — cloudfront.aws_configuration_id_selector / f3627ace5558 / 4

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

<a id="canonical-b24f2c84571022801193d0e998fd3d4ad7f1796a90c64b436649175da6c87fc4"></a>

## Next pages — cloudfront.aws_configuration_id_selector / f3627ace5558 / 5

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-a4a729753c11376ce56cff36eb4091c35e7764fa5429ab331cbfe66b6d43fd7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-062b7c3c37cccfdc4f032e7ccc1b02db42a146d45b5bfb98526f29981a96e8b8"></a>

## cloudfront.aws_configuration_tag_selector — cloudfront.aws_configuration_tag_selector / d6fb4945d3cc / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.aws_configuration_tag_selector

<a id="canonical-d999ac44092e4dca866dd6341fb01d9a72d1e14544cceefb582299f38d5e622c"></a>

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

<a id="canonical-d4580adbc9e23641836fdf0c23988b7190004344826798a4cac6ce00048943f6"></a>

## Direct properties — cloudfront.aws_configuration_tag_selector / d6fb4945d3cc / 3

<a id="canonical-3460d7ebd7f21fb37d2de3bb8ba771986d3d2611b5ac18da5cadbd2261db1640"></a>

<a id="canonical-003a9b339eef8d69f504ed6d5765ca81faa6f17633e94259da1d6f0558ac6dd3"></a>

## tags property — cloudfront.aws_configuration_tag_selector / d6fb4945d3cc / 4

Type: `["map", "string"]`. Computed.

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Upstream description:

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
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

<a id="canonical-2af0e2191e8eeffb3eeb14a5e8ec2fb21fc262bb7e9f9d16bb379a19fc77ada6"></a>

## Next pages — cloudfront.aws_configuration_tag_selector / d6fb4945d3cc / 5

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-1726345029492ab4c9ebc6f72c73146b297226d4ea16cd63163cb1745b5bcff7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1017f154958c06bd8df847c2bf0978ed852a06452745fbfd7f102258ac58e772"></a>

## cloudfront.disable_aws_configuration — cloudfront.disable_aws_configuration / 085eee232577 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.disable_aws_configuration

<a id="canonical-a7a5404e40bc394a545f752ff1115e216544eab82b2047af9ffc2a51c04e7415"></a>

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

<a id="canonical-1435480755b19ca094473ac20a222a01e39c58c30632d757e7915d819edaa29e"></a>

## Direct properties — cloudfront.disable_aws_configuration / 085eee232577 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5d2770879a78f590781f72d6746d8aa6f5e687a54dca400d488f5727ddae7f0"></a>

## Next pages — cloudfront.disable_aws_configuration / 085eee232577 / 4

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-db631d35eae4903696cc487a003fd24f41f038e41ca8789c489ee644ddea7baf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46782154ba211552e50d126bde6a81d60fead5ae29ce6c0fc842b6d013f944d0"></a>

## cloudfront.disable_js_insert — cloudfront.disable_js_insert / 983112777390 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.disable_js_insert

<a id="canonical-d466dec7844bf89f6b0fb4014585c22aa9de343484f1224bda2c010c706665bb"></a>

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

<a id="canonical-7f461a8f2e225372cf9e561d2ccfe847141338402d8c7e0a4723ac3e2a517d80"></a>

## Direct properties — cloudfront.disable_js_insert / 983112777390 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9eb3150f4c4d1075bfc3bdb5f450de680e5bf9425a5c4641433129ee810d53a4"></a>

## Next pages — cloudfront.disable_js_insert / 983112777390 / 4

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-4707d07fdea3b9a915148c99f771232f176fece1ec69d5d33edf84492dc0766c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17edd7f86d8b531eca5db8714e89cf3d1edbc0518cd844e173ce55702c104439"></a>

## cloudfront.disable_mobile_sdk — cloudfront.disable_mobile_sdk / 016fd74f6bde / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.disable_mobile_sdk

<a id="canonical-627efefb2af39b1220864f7c68d2eed00d00d106fd05cce7caeed9a4515e07f2"></a>

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

<a id="canonical-4a95053831b3a871821532ff44f90de100d461a9563429aa30a7abd6f4aa9bea"></a>

## Direct properties — cloudfront.disable_mobile_sdk / 016fd74f6bde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ded5d7a657a535c9560e40b13d1267377de218919b1ab343a9de6f67a1bc371"></a>

## Next pages — cloudfront.disable_mobile_sdk / 016fd74f6bde / 4

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-febade161f51d7082b48ab1e8fd1de284f61361111118870b4ddd1a4f212206b"></a>

## cloudfront.js_insertion_rules — cloudfront.js_insertion_rules / 55ed90a1fce5 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.js_insertion_rules

<a id="canonical-aaa791313e2b8d45bd9b9f914d117d86e4bd94b409b8b2783562d1dc675554d7"></a>

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

<a id="canonical-66139d753b9a86b0ee9a8924cacee03dd3c0f4b4f7900c0c7cfa8e2c0cb04240"></a>

## Direct properties — cloudfront.js_insertion_rules / 55ed90a1fce5 / 3

- [exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d): complete subsection reference.

<a id="canonical-a79cdb66038cc502a8ddd958af4012cfc64d04be01b31219c027c4be829372ea"></a>

<a id="canonical-81effc630a62131743f3ca3660be7dff56ba68ee9690422b97528ab897cffce5"></a>

## javascript_location property — cloudfront.js_insertion_rules / 55ed90a1fce5 / 4

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

<a id="canonical-008c89b19f62052e90ef19ab0ef245fa07e8739ad451c1d00551562bdda6de78"></a>

<a id="canonical-37b5527528773752e4b8fa21f6c9b7affbbde5da49f39d2147e9dd26f60486ca"></a>

## javascript_mode property — cloudfront.js_insertion_rules / 55ed90a1fce5 / 5

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

<a id="canonical-e5937f84a4f9200c04ad58d2b142d4dea108d4095d82694b31b551df6f2cbb84"></a>

<a id="canonical-1c5992980c0924d1acc712563b28c8531995fdbe2752eb72ae69d32851ca42c4"></a>

## js_download_path property — cloudfront.js_insertion_rules / 55ed90a1fce5 / 6

Type: `"string"`. Computed.

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

- [rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a): complete subsection reference.

<a id="canonical-218bfeff8d6f2229e96074657d18a2c4caf3bc77185109be20abe44a3e5d295a"></a>

## Next pages — cloudfront.js_insertion_rules / 55ed90a1fce5 / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-466391d86c489ae93f5b5d410a8ed73d873625de06ce75e8d7edc845a122eba9"></a>

## cloudfront.js_insertion_rules.exclude_list — cloudfront.js_insertion_rules.exclude_list / 295962602059 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- cloudfront.js_insertion_rules.exclude_list

<a id="canonical-4a512b977af7af8241c94d710e66a36606bd68d50e7cd67e5855dc6f0dde1dc3"></a>

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

<a id="canonical-930dc0c52d4cf70f6df4513a6b01bb2cccf6b161533a19a1b8ee1660c9329df9"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list / 295962602059 / 3

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-e99013c9e7ed652433fb4a70c410da0179c0e53c9058b142689ed5895fc91cf4): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-eb79de1c4d7630657cad76cf7725ec81263a35fd0ef9515719042b01f10bb48a): complete subsection reference.

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-e1d40a8492513c4e34bf265c4c2d4cf77ad679d46d10f60d7b1685ba3c91160e): complete subsection reference.

- [path](data-sources--protected_application--reference--group-002.md#canonical-edbce4ed073bc4947b0404e40cabe24e251b6a0e95682863aecacb6d311185ad): complete subsection reference.

<a id="canonical-e4ef79a2778470bb83ed58995d408948a458522453facd11235ebcd19f396d4e"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list / 295962602059 / 4

- [cloudfront.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-002.md#canonical-e99013c9e7ed652433fb4a70c410da0179c0e53c9058b142689ed5895fc91cf4)
- [cloudfront.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-002.md#canonical-eb79de1c4d7630657cad76cf7725ec81263a35fd0ef9515719042b01f10bb48a)
- [cloudfront.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-002.md#canonical-e1d40a8492513c4e34bf265c4c2d4cf77ad679d46d10f60d7b1685ba3c91160e)
- [cloudfront.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-002.md#canonical-edbce4ed073bc4947b0404e40cabe24e251b6a0e95682863aecacb6d311185ad)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e99013c9e7ed652433fb4a70c410da0179c0e53c9058b142689ed5895fc91cf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a04fc0af3c78350d27e5c8d7c465c58589ef8a24dec40183bda83845f0f81a13"></a>

## cloudfront.js_insertion_rules.exclude_list.any_domain — cloudfront.js_insertion_rules.exclude_list.any_domain / 3d6ee5093737 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- cloudfront.js_insertion_rules.exclude_list.any_domain

<a id="canonical-28b101c07f751195be7101f3edd2c87a01306c8d08c96dd8c89934e7f2a528ae"></a>

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

<a id="canonical-3459bd06b6b8de1819e72e4a69fed2ede8d35c11517cfa88416036b84125c78f"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.any_domain / 3d6ee5093737 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da605070dc7b94521c6fea6ecd4aaad8991ee258e549dc2af2bf5405e64a22c4"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.any_domain / 3d6ee5093737 / 4

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-eb79de1c4d7630657cad76cf7725ec81263a35fd0ef9515719042b01f10bb48a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cab04ee0fc6b0f0a3fb2a381400e11956e3feb8c5bf7eada67c2f618f7acad73"></a>

## cloudfront.js_insertion_rules.exclude_list.domain — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- cloudfront.js_insertion_rules.exclude_list.domain

<a id="canonical-b41624e898361b588d542e9891927e462336f9140bd2a805fc66db14385c019c"></a>

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

<a id="canonical-7882bb432d247a17f80588edaef634318645ebbafb50ccf1e64f451f995425c3"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 3

<a id="canonical-760c42599d7af5a3717f5f37721406e36d48982b655ebdca14d6f78e38b36643"></a>

<a id="canonical-57a2ae2b21d59492369424b3368d0a45b4bd8cdd33d457322ab3b98045337623"></a>

## exact_value property — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-9627afbc9d50255f9d8e0ce6a6ac8036cb2d65afe86aa9c43e86de699a993c1f"></a>

<a id="canonical-22bd5098d3af4b745bae114feee57f98e16dd1518a5e10b4d39b2c162b9145c8"></a>

## regex_value property — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 5

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

<a id="canonical-e24d19077aaf13739506ecdb1c421525d571df9f95a24e644c0769dd30ed8fe7"></a>

<a id="canonical-54a7517d48f31824294368ff4b433020c5839df224b8daa6e906ee3ef9d8c3a7"></a>

## suffix_value property — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-23c11638b0bf3438b9802b9dca664d712a5cfb841b1ac9caa3c7072c732651af"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.domain / dd1011fd9528 / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-e1d40a8492513c4e34bf265c4c2d4cf77ad679d46d10f60d7b1685ba3c91160e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b2c8333431a02e550816fd9492501ff84e13bff1a90df568a307ee2f39758e1"></a>

## cloudfront.js_insertion_rules.exclude_list.metadata — cloudfront.js_insertion_rules.exclude_list.metadata / 5b6b6c02bb8a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- cloudfront.js_insertion_rules.exclude_list.metadata

<a id="canonical-154919c36c5d536d22949ae7daeb800e957f368d0b11fc652cf9a15888cec202"></a>

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

<a id="canonical-57683490a143e3ffe96a83c672c6c98094990c65dca7f0979e4f3db27cfec094"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.metadata / 5b6b6c02bb8a / 3

<a id="canonical-66e97ee1ff7605a9b0caa9dbad244b024f06c40f3af3770c62064bc07509c3e4"></a>

<a id="canonical-19f6d2dca7a356e1a26b7707bcc09e3ecd16c9767262eb95f49ba892b1be5e8c"></a>

## description_spec property — cloudfront.js_insertion_rules.exclude_list.metadata / 5b6b6c02bb8a / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-25599158161290e25d755509bc26b032bb178237d06c3df0abe5ceeb9b745aea"></a>

<a id="canonical-2a77197447ed650a3cb09af1c205cc0efae2660fec8b25661eed00bd5a786008"></a>

## name property — cloudfront.js_insertion_rules.exclude_list.metadata / 5b6b6c02bb8a / 5

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

<a id="canonical-1488ece7b4803434b70bef41d0543a617fc2d4432e7c91d5aef7a751f6ab6d81"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.metadata / 5b6b6c02bb8a / 6

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-edbce4ed073bc4947b0404e40cabe24e251b6a0e95682863aecacb6d311185ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1904ec77e4674b4e08158b1c5f0eb97be4bf6c62aa95444ea9fdc59e427eaa90"></a>

## cloudfront.js_insertion_rules.exclude_list.path — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- cloudfront.js_insertion_rules.exclude_list.path

<a id="canonical-d982962766ecddb6187b35fae8b9dfe3385cfc3ffc8d83ddac01e90f4975e77c"></a>

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

<a id="canonical-1b976b9cba9092588a35954baebf1eae5a08adca91cd047668840a56a8e9e69e"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 3

<a id="canonical-2f00889807c9f8e73cb4a2c245389ae1f9e7c2541ac0c05bcedc48ad1520cfdd"></a>

<a id="canonical-9eee86763fa192cc01736da3830b86d7ab4016943b8ee4bf12761154e10b7991"></a>

## path property — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-8264e3e265a8bbb66b1bde98d84736995f5300a19e67fd95c02b46d1f7c80c03"></a>

<a id="canonical-10d38504b3bf5e4e88a1a915a8f51d4d44e4e280670fea1355743824ddf52b08"></a>

## prefix property — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-84376cc2e551bfdc04c071ef075dcf492cf3f48f7646773f7bfea84283a42776"></a>

<a id="canonical-8ce0645d70fec38042822a03151fa9e4f0c41911c80ee50ffe975b0216cb1dab"></a>

## regex property — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 6

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

<a id="canonical-768b2a7886e79d2154bf14e447a618bb383ddaafb71716c88cd3074141d792d7"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.path / 259882ad299a / 7

- [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-f8f7eaa6c9927645f02b5859fc2f9a8ff41291204d79d2975b1f2685f6bc203d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b706045fc2e7d6b1d40db092fef473100e39925a6db4362f018401e19c7bbb2"></a>

## cloudfront.js_insertion_rules.rules — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- cloudfront.js_insertion_rules.rules

<a id="canonical-d697b8066bebbad9694874a0ad7d1f49b004ebff2d63c4e46693ec2846d772b9"></a>

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

<a id="canonical-0dc320972acde9ba9410d8341b453c6c99be11f50086919301d2d8cb1590f159"></a>

## Direct properties — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 3

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-f209a045f270978d01fc53f830e88ee4b70546369736e264562ddf6ad6f19c6f): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-64140a35774fb49412b06a4922a418d80c047155f6a5eac0749358844574388e): complete subsection reference.

<a id="canonical-bc2ac7a5e90a05e05075b8b42443c34629589f7e8546c9fa99b8f6753e57ea8a"></a>

<a id="canonical-eace6b7a70a7cab2f636a1f6e298d29e8ef73f72c70d1268e6cce92fc980a105"></a>

## exact_path property — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 4

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

<a id="canonical-270d24f12d7f3e7404ac0a100d3cfb73344717ff15c209f9fdd99606ed1b3a59"></a>

<a id="canonical-398b12eb9432a03f53a86ce2817190a0529087bbb7c5e09c98c52cebada47f2a"></a>

## glob property — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 5

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

- [metadata](data-sources--protected_application--reference--group-002.md#canonical-d701a04314016c7539e314bc5ba794b6d617f31ca80b43b1c6b1cc08a6b32749): complete subsection reference.

<a id="canonical-525ba94c541e770863e22cbd15d507a2852f20cc944a43eacad1633218b2292c"></a>

<a id="canonical-cb3baf8db34e5d7a64b008de1a50d6d79aa0987db545b7724c2e2a91b45722c2"></a>

## prefix property — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 6

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

<a id="canonical-f249ebd80c9d0a74705ff46a294fd33cc8212716a5eb6f8088df03d61655442a"></a>

## Next pages — cloudfront.js_insertion_rules.rules / f5729c4b1dba / 7

- [cloudfront.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-002.md#canonical-f209a045f270978d01fc53f830e88ee4b70546369736e264562ddf6ad6f19c6f)
- [cloudfront.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-002.md#canonical-64140a35774fb49412b06a4922a418d80c047155f6a5eac0749358844574388e)
- [cloudfront.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-002.md#canonical-d701a04314016c7539e314bc5ba794b6d617f31ca80b43b1c6b1cc08a6b32749)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-f209a045f270978d01fc53f830e88ee4b70546369736e264562ddf6ad6f19c6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec1ee0e4c982781d54bc3286a22bb09d8cce2d39b4ae50ce3fcc0c50c55d28e1"></a>

## cloudfront.js_insertion_rules.rules.any_domain — cloudfront.js_insertion_rules.rules.any_domain / b50804ad1b96 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- cloudfront.js_insertion_rules.rules.any_domain

<a id="canonical-2a0ea93be25528e8529786273f5766e507591f17ab3e44edd6a9aad684184447"></a>

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

<a id="canonical-35491f9e616cb48ab253996500cd685cc26ee118391b328c12417ef69095e881"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.any_domain / b50804ad1b96 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-137eb4f1c2d9d7e4a7b9d0dd80bfd0c2cca8b3073f650cb5f54ba297c44a26b8"></a>

## Next pages — cloudfront.js_insertion_rules.rules.any_domain / b50804ad1b96 / 4

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-64140a35774fb49412b06a4922a418d80c047155f6a5eac0749358844574388e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9592ec18b645a72c81125e97d2a0dfa06293e5d87f159e7e30ceed5b63c7d72"></a>

## cloudfront.js_insertion_rules.rules.domain — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- cloudfront.js_insertion_rules.rules.domain

<a id="canonical-fda66c1f6b9733f38473a25a126867748588329b07d6ba4e124afe2c9e9a8d74"></a>

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

<a id="canonical-de4354ac4fa733a8ef4c888e74e31d73e5408b7d4d4f7f428929ac796e6e4641"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 3

<a id="canonical-9bb3c970022d1b988669b77827feb196bd43cbdc279dc3c41ad016862ec0259d"></a>

<a id="canonical-b939ba6edd2a23e322d00686d5ec537fa30a5af76eef39cb96da3310d3f18e47"></a>

## exact_value property — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-757b6d535f98a937925fe45f5af508c6dd7e693b2668aa2ab3834cc8e90cf022"></a>

<a id="canonical-e54e08442b6a4b4bb34642e5310992b5d204b72bc8b1b097f9cc70a9204ac4d0"></a>

## regex_value property — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 5

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

<a id="canonical-e7d1a35bf8ddc34c5f7a8a2e4b8790d9353b0128ae34f8fe66c3842c9bfb5c8e"></a>

<a id="canonical-48215faa1411091835713016c3bd92d6a187c3e1cd045cfa315ef894861d7609"></a>

## suffix_value property — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-1e1d70d325853eff1ab79275ef5ad7e6aba95ea30941f66767020e1c615f9dab"></a>

## Next pages — cloudfront.js_insertion_rules.rules.domain / 6e5d9aee9458 / 7

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-d701a04314016c7539e314bc5ba794b6d617f31ca80b43b1c6b1cc08a6b32749"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51b76719cc45e68698f1c9b723fb0b9cee735cfe1f82bfe97bd50ee6f586488b"></a>

## cloudfront.js_insertion_rules.rules.metadata — cloudfront.js_insertion_rules.rules.metadata / ee53c65c0fc2 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-9c4450e9c1256dd0ee8d4326a9ef5f8570e2b8d89d3a1ca5545ba0c5ab275ba0)
- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- cloudfront.js_insertion_rules.rules.metadata

<a id="canonical-0ad324cd3db5233fd8e6ad7608743f55a63fbc7772ace366bd70a8087af47d51"></a>

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

<a id="canonical-6b73a21fbf157ba187d3bd91b39d0453629d54b8dfb0968d66ef29e19cff6b43"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.metadata / ee53c65c0fc2 / 3

<a id="canonical-c1b76d5976b41e0903dd36e63cd965336d0f945df1b3357f6d31869e9d3a8ec8"></a>

<a id="canonical-28d6338aac5c9b3dc0e970ea5141c7a8fdae364b8e5688084116e9f3333853f2"></a>

## description_spec property — cloudfront.js_insertion_rules.rules.metadata / ee53c65c0fc2 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-9b7c699fb48a8d6f27b0bcb9e1c39c07e48f6954d65e4589733265cb8c5232e5"></a>

<a id="canonical-c7480d63138408f93dea8fe908796a8c774612a050db119426b72455ad821881"></a>

## name property — cloudfront.js_insertion_rules.rules.metadata / ee53c65c0fc2 / 5

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

<a id="canonical-2ffc08b31f1975817a4e0f93be678040a9b237143c4c898862f78bdccec723fb"></a>

## Next pages — cloudfront.js_insertion_rules.rules.metadata / ee53c65c0fc2 / 6

- [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-7bd11ba5ca33cf68e99cc4d5646f0489c0b4e6ad45faa6f633ee062b6386b91a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-db9e12a0440c850d9f05981f532646eaeb8c7862d32075a72efd124dfe3bd2f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11c021b70f665c0ea26fd0a710c351b9ce75ad08272e8df8dfdd2fdd6bbed8f0"></a>

## cloudfront.manual_js_insert — cloudfront.manual_js_insert / 22f293d994b3 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.manual_js_insert

<a id="canonical-6d15f12bdab3a2a9c9a83bfd91d555c1b73cfd789122b57fa76df88a26b6299e"></a>

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

<a id="canonical-cc9bea5c0529951f6b6a9fc26441fd8a9bc5197284cd525fbdb154909f9bdc0f"></a>

## Direct properties — cloudfront.manual_js_insert / 22f293d994b3 / 3

<a id="canonical-8dadc1b8422e83d95bce96e093e4d387eb4d1718ec96ede38c8a80f68521f76e"></a>

<a id="canonical-f0bdf1339c44906827eea062ef47d90f8fd846d6dc8c49cb456461467d91d174"></a>

## javascript_mode property — cloudfront.manual_js_insert / 22f293d994b3 / 4

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

<a id="canonical-37f4c7e64e57b3b31024e528ee55634f1d1632023c7ad19bec98a51821cc108b"></a>

<a id="canonical-28730e21982d74d43b67652aad686dc2fdcc1ee83239c27ef86c05db638e0b85"></a>

## js_download_path property — cloudfront.manual_js_insert / 22f293d994b3 / 5

Type: `"string"`. Computed.

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

<a id="canonical-c1b2997c27780059163521509a41ca5a4493faf4b725fee159ef0f8a1172b33e"></a>

## Next pages — cloudfront.manual_js_insert / 22f293d994b3 / 6

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9192a98f7474a6503084e6bbe3e34414d03ea1598298103be589e2f310d33cd2"></a>

## cloudfront.mobile_sdk_config — cloudfront.mobile_sdk_config / cd5a20f17867 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.mobile_sdk_config

<a id="canonical-762646863dc0921bac8fd7e3697d69ccb149a9e92a7546d4f9978932f6097e1e"></a>

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

<a id="canonical-212dfb6dc4b34d392a6643851c847639173ab6aba58bde45692eaab6dde9eee7"></a>

## Direct properties — cloudfront.mobile_sdk_config / cd5a20f17867 / 3

- [mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-d52a00e2e73ebb7ac13320ad49f22a150ac3374fd947da670d6cd678c8f8b08d): complete subsection reference.

<a id="canonical-a6d49602e69afb24969ef8bee225fb508c89d5864624b201cfd80876290527c5"></a>

## Next pages — cloudfront.mobile_sdk_config / cd5a20f17867 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-d52a00e2e73ebb7ac13320ad49f22a150ac3374fd947da670d6cd678c8f8b08d)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-d52a00e2e73ebb7ac13320ad49f22a150ac3374fd947da670d6cd678c8f8b08d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da858cfefafc6dbb9809a66a54dfd9a960dd2a13dac88e16a0b0045b2303293c"></a>

## cloudfront.mobile_sdk_config.mobile_identifier — cloudfront.mobile_sdk_config.mobile_identifier / 0f115d17ac49 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017)
- cloudfront.mobile_sdk_config.mobile_identifier

<a id="canonical-7c3bbb9078731bacdba2a8b1b38ed7556c2e18a1d80700b4d91aebd768527152"></a>

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

<a id="canonical-46a1a7074672b493ed68d6710f5971565a07c5b891e05aac1eb72ca53c86079a"></a>

## Direct properties — cloudfront.mobile_sdk_config.mobile_identifier / 0f115d17ac49 / 3

- [headers](data-sources--protected_application--reference--group-002.md#canonical-ea1b848fbaad553ab8aa4e912617ac9dc8783eec7c7b904ff3fce8973def720f): complete subsection reference.

<a id="canonical-8820831e4159526389df81598eac72340508f7aba9f91021e2e0446c05c6b56e"></a>

## Next pages — cloudfront.mobile_sdk_config.mobile_identifier / 0f115d17ac49 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-002.md#canonical-ea1b848fbaad553ab8aa4e912617ac9dc8783eec7c7b904ff3fce8973def720f)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ea1b848fbaad553ab8aa4e912617ac9dc8783eec7c7b904ff3fce8973def720f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f0cd0bf1fd0098019850978f38a975c58eb8a2099b1461e879df71e6ca7172a"></a>

## cloudfront.mobile_sdk_config.mobile_identifier.headers — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-bb7cbfe571f4dadc3bc691ffc9fdf5be827d927a793c64e3adcd2cecf956c017)
- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-d52a00e2e73ebb7ac13320ad49f22a150ac3374fd947da670d6cd678c8f8b08d)
- cloudfront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-7ad94d8fb817f77c9b710c033f24b9a2bfdbcd93d8ef08998b49df1cb05accaf"></a>

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

<a id="canonical-26687891a09f4dc1af867e389d20d8cdec915b10b27dfee6c3d5049f28289a9a"></a>

## Direct properties — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 3

<a id="canonical-baacc2690f779e2d3dac529ad26922169813269b498246bbce0d3f4617ee749e"></a>

<a id="canonical-c52f33c9351a41c0b54eb2362e0b48841b61ad9e4d0c492086ee70ea1711b00d"></a>

## exact property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 4

Type: `"string"`. Computed.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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

<a id="canonical-12b541bc13d3795287bca29f18d1550ebc1d804ade6b671f64970563d94f1516"></a>

<a id="canonical-698c46c69b3bde5cf64b91a2b6bfaccfa1c6ec3b623e93731dfa0bc953275c30"></a>

## name property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 5

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

<a id="canonical-3ef5f4f2c67efd2f3e65179a6104ca8c0635c6ff8c16e5dcddba126be8d03392"></a>

<a id="canonical-5f9c7e2411b914f86b41a817da1b2244f27c5118bc26ae1debaae73338d70fdb"></a>

## regex property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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

<a id="canonical-5f326f0440a8ca8ef2c598355ab9dd41535663567b54b819b888147230011173"></a>

## Next pages — cloudfront.mobile_sdk_config.mobile_identifier.headers / 9ca0c221fdae / 7

- [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-d52a00e2e73ebb7ac13320ad49f22a150ac3374fd947da670d6cd678c8f8b08d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba065e002aa1985ef032ae933128cf5e3a0b8d285a7220fe115728e48e9f14f9"></a>

## cloudfront.protected_endpoints — cloudfront.protected_endpoints / 6218e407f94b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- cloudfront.protected_endpoints

<a id="canonical-2e9c89bb4831a4dead28d748a15fdc3315c79f219a235fbcf3393c6b81c28894"></a>

Type: `"list"`. Computed.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

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

<a id="canonical-7a5417b7838ae5801a894b79e26df368c8d102c7b9ba6003bbe97a0c7a1a197a"></a>

## Direct properties — cloudfront.protected_endpoints / 6218e407f94b / 3

- [any_domain](data-sources--protected_application--reference--group-002.md#canonical-2676d43f941401198dd468e36b780cba3ffdbc419cd1746e1cc7fc526e39956d): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-002.md#canonical-6d26aad627091dcf6f6cb416b19172d09eec33ebf4702b09f2dbe2c4b028c7d3): complete subsection reference.

- [flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170): complete subsection reference.

<a id="canonical-f38f68cd18ce9af63d327e3c7cd1a9d52e9cd5bc9201ac7c7945093ec017a7ee"></a>

<a id="canonical-5b9a8b3fcec0cfe80fa4f7a45e2a77c00d33f480c8e4d6eb76beb571abcfed7a"></a>

## http_methods property — cloudfront.protected_endpoints / 6218e407f94b / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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

- [metadata](data-sources--protected_application--reference--group-003.md#canonical-89b33e808ca3e0c9d8b1aa72cca8cfdca797e6766a6bfc54e8c606d221d7b6ca): complete subsection reference.

- [mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a): complete subsection reference.

<a id="canonical-9a1919d7a9860cd29a148be229f9fa3d24b63e419b4f0a5ddbe87f64275557be"></a>

<a id="canonical-5e189077065d82348f64f724e6468fd5776e22b0e7547396097e759432f6eddb"></a>

## path property — cloudfront.protected_endpoints / 6218e407f94b / 5

Type: `"string"`. Computed.

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

<a id="canonical-cebdb8166a09e3883df2a7707d5e0987dfc6fe580cefcd2456382c962d0ee851"></a>

<a id="canonical-918c8826aea4d927df7f1f72d4be09ae116a37238c57be01cf58f8fd0b72e430"></a>

## query property — cloudfront.protected_endpoints / 6218e407f94b / 6

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

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

- [undefined_flow_label](data-sources--protected_application--reference--group-003.md#canonical-9cefac8a26e21e93a205a05fd4b81f09b05eb4e939dd5edcfcef6f2fc77d8cda): complete subsection reference.

- [web_client](data-sources--protected_application--reference--group-003.md#canonical-09769448fb2289c7e9c1998fcacb5f88119778cb1f52a1322920154605743268): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-69aa4be319e7ca4fe944058eefbf39d0b7041554f72b2be6c3d7ade860486f98): complete subsection reference.

<a id="canonical-f521311e9db09723bfd27c2dc7fa066617f3824cb85dcc8d8c77a48242c2be5d"></a>

## Next pages — cloudfront.protected_endpoints / 6218e407f94b / 7

- [cloudfront.protected_endpoints.any_domain](data-sources--protected_application--reference--group-002.md#canonical-2676d43f941401198dd468e36b780cba3ffdbc419cd1746e1cc7fc526e39956d)
- [cloudfront.protected_endpoints.domain](data-sources--protected_application--reference--group-002.md#canonical-6d26aad627091dcf6f6cb416b19172d09eec33ebf4702b09f2dbe2c4b028c7d3)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-ce018662954f11ce098a0d1e212b4244577a48ad670ab6012305548a43902170)
- [cloudfront.protected_endpoints.metadata](data-sources--protected_application--reference--group-003.md#canonical-89b33e808ca3e0c9d8b1aa72cca8cfdca797e6766a6bfc54e8c606d221d7b6ca)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-3041f54a787ce2d720d9ad74bcd8214f1b51726f556bd6613a34596b7b59f37a)
- [cloudfront.protected_endpoints.undefined_flow_label](data-sources--protected_application--reference--group-003.md#canonical-9cefac8a26e21e93a205a05fd4b81f09b05eb4e939dd5edcfcef6f2fc77d8cda)
- [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-003.md#canonical-09769448fb2289c7e9c1998fcacb5f88119778cb1f52a1322920154605743268)
- [cloudfront.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-69aa4be319e7ca4fe944058eefbf39d0b7041554f72b2be6c3d7ade860486f98)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2676d43f941401198dd468e36b780cba3ffdbc419cd1746e1cc7fc526e39956d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51b68679ce32525ca62566a548c77df7a637d6ef83c20fd2ca0f563b412a00b4"></a>

## cloudfront.protected_endpoints.any_domain — cloudfront.protected_endpoints.any_domain / 7a3998ba0c7e / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.any_domain

<a id="canonical-3269411c8d4d34403208a9080ab224ebf73dfcd1771c88c4629cb55b7b0c50e5"></a>

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

<a id="canonical-1905612c293e31314b76b43fae6939cff1441aa31ea612f0ab4e9e977aaaddcb"></a>

## Direct properties — cloudfront.protected_endpoints.any_domain / 7a3998ba0c7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c32acd236988dc598a46e91130ef2fe3cd511785fa123f310af4504a9812c170"></a>

## Next pages — cloudfront.protected_endpoints.any_domain / 7a3998ba0c7e / 4

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-6d26aad627091dcf6f6cb416b19172d09eec33ebf4702b09f2dbe2c4b028c7d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9169f1dba67196ce2aaa2d5004e7c35b7709ad3e0eeb659856bcd4838eb465af"></a>

## cloudfront.protected_endpoints.domain — cloudfront.protected_endpoints.domain / 31d3eb088934 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-63ccd5887227947c4027402da326499fc8dd87b5a738da40e7c7ef5f74e724a0)
- cloudfront.protected_endpoints.domain

<a id="canonical-b2b612f4b7c9002adb7a18a5375f6bb5ffa3a1e305808990249e1fa822b1bf93"></a>

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
