---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-b8446fbdc56f31ed526d94c812d316f9d8c828e1772b689cc60ee19c364f29c1"></a>

## status property — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 6

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

<a id="canonical-a057d3aad22e957f28a5ab0b798b1a0e9bde2db003d62cc7360b6201b5ea0818"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 7

- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aae7586985291df9859dbdbf2218fac3afc62f63e3c7224e4b28b0958d2b212f"></a>

## cloudfront.protected_endpoints.mobile_client.continue — cloudfront.protected_endpoints.mobile_client.continue / b8e01e627b66 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- cloudfront.protected_endpoints.mobile_client.continue

<a id="canonical-1a4962871cef4997128adf0e4ea5a58a161f329e7850cf385ba1910ac9ea3ce2"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-4dab887650f001d67cf7fe1509a646db618cca882f77034cbecd75ed7329f59f"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue / b8e01e627b66 / 3

- [add_header](resources--protected_application--reference--group-004.md#canonical-2392d86d20f3da7f2f891609fd67dd01edb5f3e219a33111927a8482cc8fe408): complete subsection reference.

- [no_header](resources--protected_application--reference--group-004.md#canonical-524f420d1d0b9e38cad9baa81c4886de21fe968bede025e620b537c876738530): complete subsection reference.

<a id="canonical-c55cf38769010ac9fb45ae99727f3aaac99b55da48317425ee238820bb9f0481"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue / b8e01e627b66 / 4

- [cloudfront.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-2392d86d20f3da7f2f891609fd67dd01edb5f3e219a33111927a8482cc8fe408)
- [cloudfront.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-524f420d1d0b9e38cad9baa81c4886de21fe968bede025e620b537c876738530)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2392d86d20f3da7f2f891609fd67dd01edb5f3e219a33111927a8482cc8fe408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cd62c6ef56c4a5445e38a804c657da2734dbffad7c8429c1ada039b84b23d7f"></a>

## cloudfront.protected_endpoints.mobile_client.continue.add_header — cloudfront.protected_endpoints.mobile_client.continue.add_header / a920d8a90325 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699)
- cloudfront.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-98542123d7004c5fdd6bce13686a595f5157ebb9d16faadb35792315366c7037"></a>

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
add_header = {}
```

<a id="canonical-b6a8aff6a7a6e802f42a132387f705e375b12afdc214dcbc15365b147c82a9fd"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue.add_header / a920d8a90325 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db44bb6a0276431c00627f1e780adb93f609ad455b1b2855462cd325bb9d5676"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue.add_header / a920d8a90325 / 4

- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-524f420d1d0b9e38cad9baa81c4886de21fe968bede025e620b537c876738530"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7776152ae3b72fed7756f8bfd96ca899a2f7e5b7c9701ebc4ef307c3a8525742"></a>

## cloudfront.protected_endpoints.mobile_client.continue.no_header — cloudfront.protected_endpoints.mobile_client.continue.no_header / 602552ff18c0 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699)
- cloudfront.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-222b671b76d4141a0ef64b956c6d835a9f39d00ccbb700c8a8cbebeb729ce63f"></a>

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
no_header = {}
```

<a id="canonical-5c524deab3f4f0fb5a36b446b735ab99618f79f8ecc53e63019e05a857b75b48"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.continue.no_header / 602552ff18c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fdecbd93dd9ab4bcaa090d83c091fdd2eb3bebc6968764e10e01a548e4f1a564"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client.continue.no_header / 602552ff18c0 / 4

- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1d2d3979782ee14890ffb72d01eafe38fd61faa7ab83a4a97646f1c687e2c16a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7c6ee3434205591a854667a4adaf9ee2b0df5735e0a830cacd97d8f6df7e323"></a>

## cloudfront.protected_endpoints.undefined_flow_label — cloudfront.protected_endpoints.undefined_flow_label / 312ddcc201fb / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.undefined_flow_label

<a id="canonical-f4177e4a9b9942cc7a3a0ea3aa3e25c2a7da6c84bb03d9616518d8a99799ffe3"></a>

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

<a id="canonical-e172d8ce655cf2c0c8c91746c28ec16c83d0435d287409f0627fdec9cfb88e89"></a>

## Direct properties — cloudfront.protected_endpoints.undefined_flow_label / 312ddcc201fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e027ea9c27458674993e45b6278f6560af240b4bd215e9a630439f2391361411"></a>

## Next pages — cloudfront.protected_endpoints.undefined_flow_label / 312ddcc201fb / 4

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0cf3da6885d6d6021311c3ed4e3ca366190ed107685bb4b4e898abae3838edd"></a>

## cloudfront.protected_endpoints.web_client — cloudfront.protected_endpoints.web_client / df0b170622e9 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.web_client

<a id="canonical-0d38b0aff2bb4d072a105a35662d782c723c056242112a0250a9a091f3602048"></a>

Type: `"object"`. single nested block, Optional.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-ce239838f6abc8b7db9610a9f08285606d2ff6b6883423055a6fc66e2f33215d"></a>

## Direct properties — cloudfront.protected_endpoints.web_client / df0b170622e9 / 3

- [block](resources--protected_application--reference--group-004.md#canonical-2be883d71488040e8831cc894830ed14d442e36a5919c5317df359628f5b3ec0): complete subsection reference.

- [continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb): complete subsection reference.

- [redirect](resources--protected_application--reference--group-004.md#canonical-6bb06030054e51e9ada861f15c38c7825f7cc91908c9a6fd51469ff523c0b6c8): complete subsection reference.

<a id="canonical-a98535da4da9ea73af1e1982e382b5fe8ffb1a015c8be7830b39594b18087fba"></a>

## Next pages — cloudfront.protected_endpoints.web_client / df0b170622e9 / 4

- [cloudfront.protected_endpoints.web_client.block](resources--protected_application--reference--group-004.md#canonical-2be883d71488040e8831cc894830ed14d442e36a5919c5317df359628f5b3ec0)
- [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb)
- [cloudfront.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-004.md#canonical-6bb06030054e51e9ada861f15c38c7825f7cc91908c9a6fd51469ff523c0b6c8)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2be883d71488040e8831cc894830ed14d442e36a5919c5317df359628f5b3ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec63eb4a4be491256ecd807b907d55c8ca25c3acdfbab4aa823af2876525a964"></a>

## cloudfront.protected_endpoints.web_client.block — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- cloudfront.protected_endpoints.web_client.block

<a id="canonical-c92e7edb1e3ea7dc79b72699bfd0a855a1844bfc59d5a9b80862fa2163c472b4"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-10cdab9c467420bd9dd2eacad6b2ae4df6d7903647495daf5017b6fec5e95a02"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 3

<a id="canonical-147c0ec59d1b8f190dd63d865bdb6a32a2adf9462c865b26d6a0301fe4d912cd"></a>

<a id="canonical-78c2297a4895610842e43f53a0c0b65ee3ae076805eec6061640bf0548d9e0c8"></a>

## body property — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-ae0c6ea6180bff2950bb0ea033e69f77412f1041d44ae6b2bea105f8a4750906"></a>

<a id="canonical-eed21fcc390ab705fb44319531d590c2cdfb08fcb28db65125b496294459797a"></a>

## content_type property — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-a1bb8c6c0c338365c2dff0ddc1fdeb83e70dd63cb1aca853d7795804775f8876"></a>

<a id="canonical-09789401d2efd61439ddfe69280ec0e896ddca97c023efadec841cd2d5490e39"></a>

## status property — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 6

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

<a id="canonical-3f09067f5eabcb686543acc647a61f6a344df0b7cb780c1ec6e6966e5da9b300"></a>

## Next pages — cloudfront.protected_endpoints.web_client.block / a83a8cd91d02 / 7

- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-784c3df02369796242de9a4e328de08156ed669daf9555b590b7b693b64e871b"></a>

## cloudfront.protected_endpoints.web_client.continue — cloudfront.protected_endpoints.web_client.continue / 435b79176c35 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- cloudfront.protected_endpoints.web_client.continue

<a id="canonical-b54cec048df2da1e56d767276c7fd2567be771c967bc854f13db8b00b77364f4"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-ba5f2210765f68a2183288ecf2abeff293c9f3b247e539c58eadd0b41823a6c4"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.continue / 435b79176c35 / 3

- [add_header](resources--protected_application--reference--group-004.md#canonical-2e42dcb71ebe1130786014f4c912652c593c9514dd06b0b4eb8ddf4e428163ae): complete subsection reference.

- [no_header](resources--protected_application--reference--group-004.md#canonical-008cd2a05d276c4115f46186999077beb0a4cc6f2ab9a8f448a25e2b864922aa): complete subsection reference.

<a id="canonical-fa2203ed25610eaf48482738a52635fe1a995ae66fe9339773d1b355d3461d29"></a>

## Next pages — cloudfront.protected_endpoints.web_client.continue / 435b79176c35 / 4

- [cloudfront.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-2e42dcb71ebe1130786014f4c912652c593c9514dd06b0b4eb8ddf4e428163ae)
- [cloudfront.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-008cd2a05d276c4115f46186999077beb0a4cc6f2ab9a8f448a25e2b864922aa)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2e42dcb71ebe1130786014f4c912652c593c9514dd06b0b4eb8ddf4e428163ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f783e3a280c2925203cb9dcd57850779dc87bd6608b40a47851b7f90dd7ee6a5"></a>

## cloudfront.protected_endpoints.web_client.continue.add_header — cloudfront.protected_endpoints.web_client.continue.add_header / de0c47d15009 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb)
- cloudfront.protected_endpoints.web_client.continue.add_header

<a id="canonical-8d5ce116bfa7dd86223e04bf5720e3b97e9388b0e42f0156a5b1f185ad76c25d"></a>

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
add_header = {}
```

<a id="canonical-8ab13a46f17f54b9b664c300bf492f8b2eca6ede6e64324798c240ed6809e4e0"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.continue.add_header / de0c47d15009 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e53860d9aabf5d0c45363fb8dbe2dd28edacca95088996da98601b4c42c18ea8"></a>

## Next pages — cloudfront.protected_endpoints.web_client.continue.add_header / de0c47d15009 / 4

- [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-008cd2a05d276c4115f46186999077beb0a4cc6f2ab9a8f448a25e2b864922aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2693f80a66b2d346480b4e8772160c0b8b9ea10829478a946edf5e25eaea02d"></a>

## cloudfront.protected_endpoints.web_client.continue.no_header — cloudfront.protected_endpoints.web_client.continue.no_header / ba765865c79e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb)
- cloudfront.protected_endpoints.web_client.continue.no_header

<a id="canonical-2aa36aeb3554cc761a6a9e600d9ab6304a3fdb9204cafc30281899f7c22fc737"></a>

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
no_header = {}
```

<a id="canonical-c5bbe45e133b50b95c7b9caef6886e6832145fd49c9582f862327a2fc08ee5d3"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.continue.no_header / ba765865c79e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-867d1ce85e4a7db8af76d2298ded794c74dd113f56a313db734038518efd0286"></a>

## Next pages — cloudfront.protected_endpoints.web_client.continue.no_header / ba765865c79e / 4

- [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-3d6c1409ac7884a2e853bbbc45ba52fc6e3af4459063d376e89f2694ef2ba1cb)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6bb06030054e51e9ada861f15c38c7825f7cc91908c9a6fd51469ff523c0b6c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-446c917243b103fa30aaec9aa7c636b61f35d3e133611db455620330a46f56fa"></a>

## cloudfront.protected_endpoints.web_client.redirect — cloudfront.protected_endpoints.web_client.redirect / e1c501a65a00 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- cloudfront.protected_endpoints.web_client.redirect

<a id="canonical-5dab7753c34966e046df80b3aeb3eadd7feda8d8ccbd766553c4ed8646ae2d8c"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Upstream description:

Redirect.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-b0b6ba9acfe06af2bbffd054b6891b3c96052064bec57f8ab8d0ec1491c29142"></a>

## Direct properties — cloudfront.protected_endpoints.web_client.redirect / e1c501a65a00 / 3

<a id="canonical-59e6ea4596e9a6568b1f1eb456b5cea1d9728a1e25c4a19d838702030fff3060"></a>

<a id="canonical-bfc570b023b8b2d267cc8e2d726d638f7cf078374a117c574b97b4fe24eedcfa"></a>

## location property — cloudfront.protected_endpoints.web_client.redirect / e1c501a65a00 / 4

Type: `"string"`. Optional.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-4d73f65dccac6a0febe32c364a8348585543c3607a0b15c53ceb9aec090ef2ec"></a>

<a id="canonical-ef369a41b5c38ba66faa37d703c0367e6d811ee73a67ec74d548db1c6c93e333"></a>

## status property — cloudfront.protected_endpoints.web_client.redirect / e1c501a65a00 / 5

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

<a id="canonical-97c66247207cdc7d2726009e0bb03e9d2efed8e654cf612f699e77b6e9c871c1"></a>

## Next pages — cloudfront.protected_endpoints.web_client.redirect / e1c501a65a00 / 6

- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-330edbc275f8c09ba49d57dabfead6705c958abd63ba0d57d16be3aa26560e83"></a>

## cloudfront.protected_endpoints.web_mobile_client — cloudfront.protected_endpoints.web_mobile_client / 3676e1dd8221 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.web_mobile_client

<a id="canonical-78e9983ea813aa1fbc01c557eb90cd5cac0201c60bdc9d62c95e57f989388d61"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-e4e43455ed3ae5b478d3cdb7284b0b9da4e5beba6f597186c90b4ac470155801"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client / 3676e1dd8221 / 3

- [block_mobile](resources--protected_application--reference--group-004.md#canonical-12f6f9d5cc0367f1db34c00935ac162fecdf16544e8e1fc03b43cc90a5c82c6a): complete subsection reference.

- [block_web](resources--protected_application--reference--group-004.md#canonical-bd806e3cf4ea0a981a182a13bad8b5953697cca8d10247f771852e30411a0c92): complete subsection reference.

- [continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376): complete subsection reference.

- [continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea): complete subsection reference.

- [redirect_web](resources--protected_application--reference--group-004.md#canonical-21e6abce8075f446e7482bbafc9494d73fad30f9a9ff8bbf46c27afddf845d54): complete subsection reference.

<a id="canonical-530e39ea1d365bb4da3e467e33a69587c3f6dd4d9ef766f956b03824f1b1fcc1"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client / 3676e1dd8221 / 4

- [cloudfront.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-004.md#canonical-12f6f9d5cc0367f1db34c00935ac162fecdf16544e8e1fc03b43cc90a5c82c6a)
- [cloudfront.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-004.md#canonical-bd806e3cf4ea0a981a182a13bad8b5953697cca8d10247f771852e30411a0c92)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea)
- [cloudfront.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-004.md#canonical-21e6abce8075f446e7482bbafc9494d73fad30f9a9ff8bbf46c27afddf845d54)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-12f6f9d5cc0367f1db34c00935ac162fecdf16544e8e1fc03b43cc90a5c82c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d2b1cf309b1f8dd2bb6e058caed1b8387555818b3886a489483ee9a9594da1a"></a>

## cloudfront.protected_endpoints.web_mobile_client.block_mobile — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- cloudfront.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-355c7d4a112835c13cc51cf72abfd0b138f7aaf1504ac56705e4ea0fe242805c"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
block_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7fd1d8510698befbb2376f2dbb2362b351d635f03d4e0e20c5d87617a7581d15"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 3

<a id="canonical-c881e7c7a4d6d8ff6a19a8c73f569f1945e3d8b4ece08bab1975b24a5c85609d"></a>

<a id="canonical-7d7d1230ef76a837bf014c2cf431e08915048acdb2b6ebcb5dd7a7027ad4dfc0"></a>

## body property — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3430efca028dc8a8b0994449a73032f0f4c4444a3f77832395b795bbd35567a0"></a>

<a id="canonical-0a5546fad2410e8ea35ffcc57172cae11a70953b1cd0f10c16935c489b17a191"></a>

## content_type property — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-87fe4d5bf1c041f04e0edffc9c4b7d8e6e04af47b251d41887d6bbbec78e1e5f"></a>

<a id="canonical-e6a364b9287c5f21c26c85b6cd237f2191c936dd90acd3ff96f6bf28cd7e5834"></a>

## status property — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 6

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

<a id="canonical-8339b520ec33a6d82d2c70650b6b567d792badde67e63cd5f0d948a5696872b6"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.block_mobile / 8b6bdfb2d130 / 7

- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-bd806e3cf4ea0a981a182a13bad8b5953697cca8d10247f771852e30411a0c92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a31bac81ee3a054641d711f5ed4283ac4227110131e2af974b61a27dc81b33df"></a>

## cloudfront.protected_endpoints.web_mobile_client.block_web — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- cloudfront.protected_endpoints.web_mobile_client.block_web

<a id="canonical-9756ef73f8d4ccdc6d6ffa7990fb0f72525ab5b58203dcae0e0d251cbafa1da7"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
block_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f48e5aa2c58a562a5cae52356383add08cef991346f3d7ee30f6d189ab17743"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 3

<a id="canonical-b8d2c5cbf66fba6c867f9a43da8a0ea6e2ecff736179de5ab6645df7d42a7bfc"></a>

<a id="canonical-6ec4bb936f4048dc34d2a286eb5c112e820b29bf6545b272192013b83de36c95"></a>

## body property — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-70c901cc0271f936a8052d4abf3c35653936143ecced290b537d5837f0d17b47"></a>

<a id="canonical-600adf20e6590d26d4bcaff877601b4a639bfbb4ec91e10e18486c92128a40ff"></a>

## content_type property — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-c670a6cc6fae35c57f2dc989a7d9ad35ed96ac91f048529c44aa3b7881a37e86"></a>

<a id="canonical-7f8781df817f9cfa7529e0784b689b1de288aa3f78ca906dbe340abdab3a02c7"></a>

## status property — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 6

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

<a id="canonical-9537cac2795ce823ae6bf16c490b60fb7ed6d9537ff7a5cc551764f8cb3c4ec3"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.block_web / 14b91f2cc1f1 / 7

- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb911251f213edc6dfe0a718d92a46641c2d98d4bf22ac4e75ae4ecdaea708db"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_mobile — cloudfront.protected_endpoints.web_mobile_client.continue_mobile / c0bbe52bedcd / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- cloudfront.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-f0b09b05329b881c2abb55d9fa562424632fab19948b410b07c67d70f8d7ddae"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-c6ec3aa450ec4539ca077a8c531bdae774e1bec640cc0e31dad56a5fbb3bc853"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_mobile / c0bbe52bedcd / 3

- [add_header](resources--protected_application--reference--group-004.md#canonical-e90612a5d1d5f29c7b8d4b7390c2d59c789cefb3e6a9c6e2bc64d1f6e64af660): complete subsection reference.

- [no_header](resources--protected_application--reference--group-004.md#canonical-febcee16289704f2f474078179bcf6ed76406175d6e01eb04aa483327c4d080f): complete subsection reference.

<a id="canonical-577a37ae288b8bc3ab934d7ec935afe1c0f48c861bcabe483447d7ab6d9f3151"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_mobile / c0bbe52bedcd / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-004.md#canonical-e90612a5d1d5f29c7b8d4b7390c2d59c789cefb3e6a9c6e2bc64d1f6e64af660)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-004.md#canonical-febcee16289704f2f474078179bcf6ed76406175d6e01eb04aa483327c4d080f)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-e90612a5d1d5f29c7b8d4b7390c2d59c789cefb3e6a9c6e2bc64d1f6e64af660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab3981cb0ffb9f09180561c1882c3d9c727903d1fd2d9ed273948f8b5702887d"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header / bdfae897b74b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376)
- cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-cf833703d156a0b9f7d659fd6a197dc5edbf6169864257626a28a47671c821c7"></a>

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
add_header = {}
```

<a id="canonical-1c199a7ecd8662b61d3265cd66a50e7793d57c2066b3f6cb8d4b0e5bce2dfd65"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header / bdfae897b74b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-487740bdbc067a61d329d05e1b22e1f57e84328b76c3ca512c434a569a9d9ba3"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header / bdfae897b74b / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-febcee16289704f2f474078179bcf6ed76406175d6e01eb04aa483327c4d080f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-227050d95c8b5066487401339701f1e7f6fb01fdcdc24b157869d6539a99481f"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header / 96eab9d3c227 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376)
- cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-b05d688663a64d0dcfec0c1b6e82b821843dcea1bc095a19de84dadee386e79b"></a>

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
no_header = {}
```

<a id="canonical-901c44831762119b7c51b0f07a088b190347dc0a51aa6f6c255237a7d4b5bf96"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header / 96eab9d3c227 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66cd721c9ad22b405551316163ae7564c0d47abdbc35be4d170bc64016bd2d39"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header / 96eab9d3c227 / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-dbe903e65ffcb97c62bc0a9df6e4ce124448f9f3dc624d85d14c26d67aa7a376)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6933c92301090106d82beee703e0d207a0e4ff859f42ed786b9675637c492ae6"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_web — cloudfront.protected_endpoints.web_mobile_client.continue_web / 8726a7bb0062 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- cloudfront.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-f2db42c531c317702b6ab42fcc30b6e7d2df4559da1adb1dbea0c55a49f01e87"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-c90d244fe73bc80d33e0b4574c18153741f5851ee1c8202c0fdfdb585f22d6c9"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_web / 8726a7bb0062 / 3

- [add_header](resources--protected_application--reference--group-004.md#canonical-51f820e0342e270e76e66c43481c2e24ffec09a19e60aebd1ff61afbdd5811dc): complete subsection reference.

- [no_header](resources--protected_application--reference--group-004.md#canonical-322939eff5eb17de42084fe2df8059c55b1d6018b93d4b68feecac45e9185e34): complete subsection reference.

<a id="canonical-abb3a65ff2f7c40ac0957dca47e2132f0ef67b1da38d071a466297db03eb3df6"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_web / 8726a7bb0062 / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-004.md#canonical-51f820e0342e270e76e66c43481c2e24ffec09a19e60aebd1ff61afbdd5811dc)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-004.md#canonical-322939eff5eb17de42084fe2df8059c55b1d6018b93d4b68feecac45e9185e34)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-51f820e0342e270e76e66c43481c2e24ffec09a19e60aebd1ff61afbdd5811dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7522318166953109ffbd5a3a7e67641723203971ee52c730974165450697ea2b"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header — cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header / 6928d8b4a3ff / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea)
- cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-dc7d707d3fc01e7159ae14954236a310fa52eb9762ed69a5d7d6586bb8693998"></a>

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
add_header = {}
```

<a id="canonical-7f7f39e362d1807bd73fd0c899fe960151399d897a641e5c541c41781096f8b0"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header / 6928d8b4a3ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed89422f8eaa7ac62377a10e5d62e6b11a9d0500ac75ed5a957f23d1b91fba6f"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header / 6928d8b4a3ff / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-322939eff5eb17de42084fe2df8059c55b1d6018b93d4b68feecac45e9185e34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47bd0905b4585c7be81f9b508c068dec6125570ba18fa30fa3c0aedd89b78de5"></a>

## cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header — cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header / 8c53c20ab7ad / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea)
- cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-40974e0c4df7379efecec45dfd8c0668d481303024a8b3ffdffb470235c91544"></a>

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
no_header = {}
```

<a id="canonical-3c292de6f03b962695c7834978d3f35bdbbbd5dd93c85fb4a05c2513f4d9e85d"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header / 8c53c20ab7ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a99bc9e81fc9ad3cb05ca563b32e42b9bc6266f924568150a33db6a6d9ef9794"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header / 8c53c20ab7ad / 4

- [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-ac70cdeac099be15fe7d86c602f2441b758025ec6242e2637a93f65c8d10c9ea)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-21e6abce8075f446e7482bbafc9494d73fad30f9a9ff8bbf46c27afddf845d54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f075feca697cfd8ea02d8916a66bf1f08d2b1396e2101e7a230e389319572503"></a>

## cloudfront.protected_endpoints.web_mobile_client.redirect_web — cloudfront.protected_endpoints.web_mobile_client.redirect_web / 63efe2cc8cb7 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- cloudfront.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-93b9a5c554524e95f155fa2a5c9abadec4a8b8424fb1a27e3b2bea0a623ea3ae"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Upstream description:

Redirect.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3a79905dba22dca6cada8d36563ed076ff0a95eaaa679c0c6bce779da6988f18"></a>

## Direct properties — cloudfront.protected_endpoints.web_mobile_client.redirect_web / 63efe2cc8cb7 / 3

<a id="canonical-567106317eae5b703adbfb1f479d0f7882c85f0d3b9293b6ddfb765170fc7476"></a>

<a id="canonical-f4337a34e8d5bcd729d10c58285fc50052e976c772e52b219044952364b18202"></a>

## location property — cloudfront.protected_endpoints.web_mobile_client.redirect_web / 63efe2cc8cb7 / 4

Type: `"string"`. Optional.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-831fe9fb1f88b7bdb61ab5bc614506d8625a2ae76de26741c2a37bff3a71fffa"></a>

<a id="canonical-31819575ae3ae109641507332216f2c9f51d3718b0a2dbe2fbf45b5ee1ae8da9"></a>

## status property — cloudfront.protected_endpoints.web_mobile_client.redirect_web / 63efe2cc8cb7 / 5

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

<a id="canonical-674a8a53350554986eb348043cf02eb8d20d006eff8301f59b5b68897c809c88"></a>

## Next pages — cloudfront.protected_endpoints.web_mobile_client.redirect_web / 63efe2cc8cb7 / 6

- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4831992d05b52f69cbdc619f2ef96f8ded92929cb99fe43c5522954b6d07a1a"></a>

## cloudfront.trusted_clients — cloudfront.trusted_clients / 6aa98eeb78eb / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.trusted_clients

<a id="canonical-5e6288851f7fb1da6f26e7a1d20eefbd704751e5bc0364c966365561d58b1cbd"></a>

Type: `"object"`. list nested block, Optional.

Define your allowlists to skip Bot Defense inference processing.

Provider validators and defaults (from schema source):

```go
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

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e2f87bd66a5f9acce437d9e34a926e9b3ee69f40b88f309dfea445f424ddf88"></a>

## Direct properties — cloudfront.trusted_clients / 6aa98eeb78eb / 3

- [http_header](resources--protected_application--reference--group-004.md#canonical-1f5d2e77cbf9d7f14642ebdde41974aa083bcc4722475022aa0de4156f6eaa3b): complete subsection reference.

<a id="canonical-2b114f1806faa6414877b8ce6332ea856e43463c3451338697d53443365aa663"></a>

<a id="canonical-ea4b6ca2dc839c3497cd381867b1e51ffccaacf7cdcf7fcb828f624b410d55a3"></a>

## ip_prefix property — cloudfront.trusted_clients / 6aa98eeb78eb / 4

Type: `"string"`. Optional.

Exclusive with \[http\_header\] IP prefix string.

Upstream description:

Exclusive with \[http\_header\] IP prefix string.

Provider validators and defaults (from schema source):

```go
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

- [metadata](resources--protected_application--reference--group-004.md#canonical-8fd86bf9354e3a143faf333ec487d342b732420fef8eada89b1da55828892d82): complete subsection reference.

<a id="canonical-347bc23ab6065792f45f7abbe876d014c450e3d5ccdfa7a026d22e5aa71c340f"></a>

## Next pages — cloudfront.trusted_clients / 6aa98eeb78eb / 5

- [cloudfront.trusted_clients.http_header](resources--protected_application--reference--group-004.md#canonical-1f5d2e77cbf9d7f14642ebdde41974aa083bcc4722475022aa0de4156f6eaa3b)
- [cloudfront.trusted_clients.metadata](resources--protected_application--reference--group-004.md#canonical-8fd86bf9354e3a143faf333ec487d342b732420fef8eada89b1da55828892d82)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1f5d2e77cbf9d7f14642ebdde41974aa083bcc4722475022aa0de4156f6eaa3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a1354838edca2af3dadd94521d14e2224885ee15d43369e42e7bae068dd0361"></a>

## cloudfront.trusted_clients.http_header — cloudfront.trusted_clients.http_header / 0ed1bf331a0d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- cloudfront.trusted_clients.http_header

<a id="canonical-a56ae78fb57a1ba1089d033cc7cb3d0a4a8973513f1d9de3ca2f476f4df13d3d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-58cf0c19a10c6b366e10ecbb6d79fdafbb215ff4871c3cf1acdd1bf67580b730"></a>

## Direct properties — cloudfront.trusted_clients.http_header / 0ed1bf331a0d / 3

- [headers](resources--protected_application--reference--group-004.md#canonical-8cd97a665c701244efd51da3bc9d3924c04e74ea92aa10690098510ac5870f40): complete subsection reference.

<a id="canonical-4df391482427d1d54834953a52ef247a35f8e59ce50bd6fcaa513bec092336ed"></a>

## Next pages — cloudfront.trusted_clients.http_header / 0ed1bf331a0d / 4

- [cloudfront.trusted_clients.http_header.headers](resources--protected_application--reference--group-004.md#canonical-8cd97a665c701244efd51da3bc9d3924c04e74ea92aa10690098510ac5870f40)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-8cd97a665c701244efd51da3bc9d3924c04e74ea92aa10690098510ac5870f40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c90c7406565bf88356d01bb832c3190657011259dad74a79569f50a6e7cc3ee"></a>

## cloudfront.trusted_clients.http_header.headers — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- [cloudfront.trusted_clients.http_header](resources--protected_application--reference--group-004.md#canonical-1f5d2e77cbf9d7f14642ebdde41974aa083bcc4722475022aa0de4156f6eaa3b)
- cloudfront.trusted_clients.http_header.headers

<a id="canonical-337d92f8b6862f67d84de9ddc22b82dbe29ac39ea0de81879db26a8a2945b4a4"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

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

<a id="canonical-fde5efe956ca125c75f02b039d9cceb0c01f6dc94ebb5a33fc2b8ca27c372003"></a>

## Direct properties — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 3

<a id="canonical-fb61a8222aae52320769c67c25648188c12708300661b8f1f1aa422269bb6872"></a>

<a id="canonical-9c7cecd37048ad39c3c2d06e413f6551e563ac0800d1f79293465ef81012fbe2"></a>

## exact property — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 4

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

<a id="canonical-3fde174f2baa92798f502549930cde2a9812b3efe1e86e6eb9f5a706c5fea0a1"></a>

<a id="canonical-46346a6c7b2a4e3f1a68630039977e9b8aa307194eedaa5bebdbd05173d9c4da"></a>

## name property — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 5

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

<a id="canonical-1ab2eb65bc622782eb0f92ad6a2c042f2be1d9ab8044094ba420d4d13a41dfe0"></a>

<a id="canonical-4c6b04c26e7d3f38c78e39c685bc77475dd4d8572618a5928c7d465003f9f706"></a>

## regex property — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 6

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

<a id="canonical-faa1b192cff894d0e2541b05ed87bfc332799c5816a89b02d75189a18001cf3b"></a>

## Next pages — cloudfront.trusted_clients.http_header.headers / 7292956214b6 / 7

- [cloudfront.trusted_clients.http_header](resources--protected_application--reference--group-004.md#canonical-1f5d2e77cbf9d7f14642ebdde41974aa083bcc4722475022aa0de4156f6eaa3b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-8fd86bf9354e3a143faf333ec487d342b732420fef8eada89b1da55828892d82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f6c33c721e6bf83d1b7afd84a0e97caff0d2801b59b04d4d0a095b92ed5073d"></a>

## cloudfront.trusted_clients.metadata — cloudfront.trusted_clients.metadata / bcc02b4af891 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- cloudfront.trusted_clients.metadata

<a id="canonical-3eadabc7bb13231be0c8f4a7cc39932338d7836a73fe444937adf3f39c937127"></a>

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

<a id="canonical-35d7f3d7eb31e649e2fcd8ea550068c25a53b6e7ff3a9f22f62e323ff909422b"></a>

## Direct properties — cloudfront.trusted_clients.metadata / bcc02b4af891 / 3

<a id="canonical-b4206a2f865ae2ede2df301e976ef86a86921d58ec9ba17fa4b718649bc8b442"></a>

<a id="canonical-13ece9fb01e9f9d99321c2396d3f3870d6694648aa9d697af57239c02de1715e"></a>

## description_spec property — cloudfront.trusted_clients.metadata / bcc02b4af891 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-5da5ff0e4c6ed1c20ea7dec93fe2d3987c5642abd0d9cb0c97e17cc2005c92be"></a>

<a id="canonical-fdbe14f4fa956c8a89c74227f83540cbccf338c122b513f45cc68a8d0ce6ae6f"></a>

## name property — cloudfront.trusted_clients.metadata / bcc02b4af891 / 5

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

<a id="canonical-b643734634beabb44c2266954e1dc13260e6c83495c0226e8587c7728695e24b"></a>

## Next pages — cloudfront.trusted_clients.metadata / bcc02b4af891 / 6

- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-015119427834384a4021619b34741d838615efc7b8a44226c1650dcff1fc786e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91c3229c83bbc86ee12fd2096083b57d70a80700078be4bf9c5a6773cf334c31"></a>

## custom_connector — custom_connector / b30e8e0caa13 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- custom_connector

<a id="canonical-7aa10585b9f37614e34760a4e1a8a5437bfa0cb3f8c89cc28da45f215976a824"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom connector.

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
custom_connector = {}
```

<a id="canonical-08512a350bf6e0a780db8945cf0dbb4df8a0b045b0d09d7aada642601929f5e0"></a>

## Direct properties — custom_connector / b30e8e0caa13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d126ed27152f398d8fe6d860bff59729f64ee0fb2c029349a5990d769357b01"></a>

## Next pages — custom_connector / b30e8e0caa13 / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-09aaf1806353d50ca14d6e132e2cb3694e4b277d578cb3e7f1afd534500fa0d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2d662f66d911f2577d58a6aa1907c922783ba8ceb70cbad8a7b2a6fe969dfe"></a>

## f5_big_ip — f5_big_ip / 53e2a54e80ed / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- f5_big_ip

<a id="canonical-08cf5923e93d3f704c81fba4e5beb3f6d8d9fa99aeb9e828733fbbf75e44bf1e"></a>

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
f5_big_ip = {}
```

<a id="canonical-732e1d74af9f9febe016f7208b18e775f191a2790081de357cd5911fa8913e97"></a>

## Direct properties — f5_big_ip / 53e2a54e80ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be936407bb0c19cbaa9b432078bc6ec6b5abecad5ce630410208e2b5978cebc6"></a>

## Next pages — f5_big_ip / 53e2a54e80ed / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-be3e87bf4eaaa0be35a3e894bbcd2e36a0721ae85f61c76da9372a3145122e64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59bb66d69c947fb6d3a08e482a5be059bfef01e7c01f0e50bb210370eb61a2f4"></a>

## salesforce_commerce_connector — salesforce_commerce_connector / a52802ffc930 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- salesforce_commerce_connector

<a id="canonical-3d84769a6a83f8cdb4493bbe723fe45f7247bdbb72d85de52ba9827eb3ed0bb5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for salesforce commerce connector.

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
salesforce_commerce_connector = {}
```

<a id="canonical-a755ea22117fcc7789b1a1d3c3983970c5318ceff5caf45fc6356ba0064a8b3b"></a>

## Direct properties — salesforce_commerce_connector / a52802ffc930 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2eb8d2f1631672f73e8ad225b8826b377b0f58c137361130e042607d0761491a"></a>

## Next pages — salesforce_commerce_connector / a52802ffc930 / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-795c21823eea76f40b96ac6f5e098873ab1262b372f3bba87d702754c6a10784"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ceeb400331c1376acb0daf0e6b83de629410e1605fdf03781a9c5044219ce34"></a>

## timeouts — timeouts / 67e3cdfc0a78 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- timeouts

<a id="canonical-7058d92adc2561c625a2c5a22bb04e060c8b8e68e6b0c1f6ed1cb748ae27499e"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3fb5287c82a57eb484e41fa9c76f8743c63a86e2c49ca6f1b85d490cd5105aac"></a>

## Direct properties — timeouts / 67e3cdfc0a78 / 3

<a id="canonical-9580f11f96e011c089aea5f5cf0674403107c6f0f1973410bc39cf0e6d179ce5"></a>

<a id="canonical-749ec39bdf9b4184a8886ac544e1f38e969e8c2ea42512da952862497df305f6"></a>

## create property — timeouts / 67e3cdfc0a78 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1b7ce3533ae0424b633140459ca5a2e5659e35ea2246c911e379abae84987f60"></a>

<a id="canonical-0b3d6e1af57e6cc088f19718297e45bfab50aa585c5e6b0bb0e0449830f91252"></a>

## delete property — timeouts / 67e3cdfc0a78 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-5c10c88110fbaede537d5184a594e583f93765af4cd899809f9c99e08c0185a1"></a>

<a id="canonical-0df96801655fa7728b4e145119dcc8a46e0383214b274b6a699de65f590a4c0c"></a>

## read property — timeouts / 67e3cdfc0a78 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9091fba4ccdaac05861adc5a958998ae608393d30389e90adc737fa6af1e9910"></a>

<a id="canonical-e17a0995e8d9b764db2c0821f0226251413c7da34a62a0245993d9440e3d1444"></a>

## update property — timeouts / 67e3cdfc0a78 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-db949df79933384d889e6629679e83e7446b4b5c45d739bc0db46570a6243b96"></a>

## Next pages — timeouts / 67e3cdfc0a78 / 8

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
