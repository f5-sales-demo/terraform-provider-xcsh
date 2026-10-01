---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-f3c77c2eec7a6583f85685443aa979fc6b543ce3ade7ee9bc29764ac20370f6b"></a>

## regex_value property — cloudflare.protected_endpoints.domain / 9101070ad78b / 5

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

<a id="canonical-3cb1822b2899b1a9bfc647ccd6f00f0851d91c35b386a9ceac3f644b4b6dcb66"></a>

<a id="canonical-35e83a80e5bdcec6b16194b56c77352e16ea1a6cb06a723bb0efad62739e3ef8"></a>

## suffix_value property — cloudflare.protected_endpoints.domain / 9101070ad78b / 6

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

<a id="canonical-0b27bf18ecf46d211ada3b7452b8e2da272db68dafa14da8eb4bd881ee3b93a6"></a>

## Next pages — cloudflare.protected_endpoints.domain / 9101070ad78b / 7

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-371da26b9e490b11558667c8d31d2c8e0426ba3d621105d0c366ad9fb3d40efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-383b3425e4d5e7841b23692c91013072208bc24eae294468a1442ad953aa30ee"></a>

## cloudflare.protected_endpoints.metadata — cloudflare.protected_endpoints.metadata / d789f0405e75 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.metadata

<a id="canonical-06b9a63e266e0610e31e6c5d7d6e25d5334956ba6865b14bcf04878722211a26"></a>

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

<a id="canonical-81391c077e02b66124dc7562d65fc6b67195f7dafefbdd2846eb23a611b39920"></a>

## Direct properties — cloudflare.protected_endpoints.metadata / d789f0405e75 / 3

<a id="canonical-e96b8d51d208c7d6e3849e558a7fe022fc9f81b6982b5638ace030fd543f68de"></a>

<a id="canonical-c6a9f9713b42c9b163863e4e76099e172cac3745ca211fd76fd5d60256cc124f"></a>

## description_spec property — cloudflare.protected_endpoints.metadata / d789f0405e75 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8cd7a3fb096e73444923d0b4c131273c410ec38b89d40ea823673961a0a295cb"></a>

<a id="canonical-e0e04cbf2c7b60fd1e83de633aa0ad5f590538c49e1dd09a55b59848ff7c3b61"></a>

## name property — cloudflare.protected_endpoints.metadata / d789f0405e75 / 5

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

<a id="canonical-4df86c8ab5c664943dcadc75e3ee4d98413370130314c7fc190a9f8d3896448b"></a>

## Next pages — cloudflare.protected_endpoints.metadata / d789f0405e75 / 6

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7950680c4839e19fb225bbefdade9558da50bce1be02364534e7fe1c0b7dd4f9"></a>

## cloudflare.protected_endpoints.mobile_client — cloudflare.protected_endpoints.mobile_client / 50124ca94266 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.mobile_client

<a id="canonical-9fbba17190ee8a82d4b592374f60b1a75a3818281a9d80d3ebcfb85bfe7e9b4e"></a>

Type: `"object"`. single nested block, Optional.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

Terraform syntax:

```terraform
mobile_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-925319ef603309d81cd5c645c3cfd91438a8e222b35f0d8c0f4277eafa1d56c8"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client / 50124ca94266 / 3

- [block](resources--protected_application--reference--group-002.md#canonical-856e00d90cca4a889a23e989e4a3549d2c58283523dc00a9762dff761d2f195d): complete subsection reference.

- [continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58): complete subsection reference.

<a id="canonical-5799787687d1d072d4699b2b4abdc916525bd3762293e9efc7c507e8ae6aa9f8"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client / 50124ca94266 / 4

- [cloudflare.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-002.md#canonical-856e00d90cca4a889a23e989e4a3549d2c58283523dc00a9762dff761d2f195d)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-856e00d90cca4a889a23e989e4a3549d2c58283523dc00a9762dff761d2f195d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3106a968724510c1e3cc75ae5468a3592a965d5b7f47d1a7b35926bd7dd6b71"></a>

## cloudflare.protected_endpoints.mobile_client.block — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- cloudflare.protected_endpoints.mobile_client.block

<a id="canonical-f75831ebdba17d982aaf85633595b628ce6edb2f449ea627d15aebeb20886761"></a>

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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-fead865e6a2094722366946847dcacd9e559a20d4acd8f75869a5e213e84937c"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 3

<a id="canonical-de0e156bd0d81b3ec1d06909ef4543c928ed015541fdd64eddaf89515f89ecce"></a>

<a id="canonical-69d6030afa4558ae0f3a449334ec5a4fc9ac15a54fa63a8ba4252288141c42f5"></a>

## body property — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 4

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

<a id="canonical-3fcd1efeecde9e4f9214a15be5f38b3bd4d22e8cad41631d0260d8d45f5b3fe7"></a>

<a id="canonical-7de773a8fede8f28e0d57e117f938141710a390bce3ba7801a8791879c7129c8"></a>

## content_type property — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 5

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

<a id="canonical-67da709b01321e010e6afa09ac20c2018a5a1e5486a3df9db346b250439d51ec"></a>

<a id="canonical-65d54d9bb22c79f7e40da5b1a520c96ef4fe40a589b97582f7893efb6b1ddb8d"></a>

## status property — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 6

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

<a id="canonical-521a76fc14e7d1be4c533ddd5952fa195fde8c7198d1f7451148b56884d9f0de"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.block / c1cb80c03df3 / 7

- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-180af1e7f82a34287a916581c4dfda2f24e9c49cd932600267d52e2c0baad692"></a>

## cloudflare.protected_endpoints.mobile_client.continue — cloudflare.protected_endpoints.mobile_client.continue / cc2c8f8356c7 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- cloudflare.protected_endpoints.mobile_client.continue

<a id="canonical-56e59604f0d05a45820bdc46d147e3f3ccc56d1e8c0ce5ff094083bbb81f6ffe"></a>

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

<a id="canonical-ba091ba53980c459156c751525fd2135f7ed324029876b9d337c221ced7b933b"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue / cc2c8f8356c7 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-648168249208c56e19d73c73daf98a2bdf5fae4dcae05cb6ccfa4dd14140cd20): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-8cb704e916b28fded9c280d8f43dd90832de1e54e0ec70e049a0019deccf913c): complete subsection reference.

<a id="canonical-f987ee850684002dd899600b9e6fc45951ea486d729e60e71166b945d2bed762"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue / cc2c8f8356c7 / 4

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-648168249208c56e19d73c73daf98a2bdf5fae4dcae05cb6ccfa4dd14140cd20)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-8cb704e916b28fded9c280d8f43dd90832de1e54e0ec70e049a0019deccf913c)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-648168249208c56e19d73c73daf98a2bdf5fae4dcae05cb6ccfa4dd14140cd20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e56800ff88808375cecdefb8aa17effce0b9de1b1276633fccb82e08f92bf57e"></a>

## cloudflare.protected_endpoints.mobile_client.continue.add_header — cloudflare.protected_endpoints.mobile_client.continue.add_header / 65226a325e7d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58)
- cloudflare.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-d56dc4ecf0f68eac81b58d74c7796af6f54d56cbc6b14e6f7e000194ac1ff9bd"></a>

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

<a id="canonical-22fcb9e647101658c814065abb53ebe751872e4a524f6d9465ca7726b5f92164"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue.add_header / 65226a325e7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77f86d5276195832aceb01d4b0617b699054b0355d03f23b9349a3b1f0b599f3"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue.add_header / 65226a325e7d / 4

- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-8cb704e916b28fded9c280d8f43dd90832de1e54e0ec70e049a0019deccf913c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73c6f93752ebee9be829831b57834bade4b0e1fc7188f81d331648bf9bd119cb"></a>

## cloudflare.protected_endpoints.mobile_client.continue.no_header — cloudflare.protected_endpoints.mobile_client.continue.no_header / ea301dfdf7bd / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58)
- cloudflare.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-7caf9af1ff22cda054dfd08f573f3710a83077cd8013bdc7abd544d4efb00905"></a>

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

<a id="canonical-638c4c2c65fd4b2cf7a123bc39fc5aaaa599f6812ef3cee1d5ee01bcf1a50254"></a>

## Direct properties — cloudflare.protected_endpoints.mobile_client.continue.no_header / ea301dfdf7bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9bb5e9f3a53bb51255bb1569a20d76c6b442dcf4a91128ef73a972a6d1e7086e"></a>

## Next pages — cloudflare.protected_endpoints.mobile_client.continue.no_header / ea301dfdf7bd / 4

- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-16c9f7d43af5c05095a237e903b7ecc3dbd3fc263bd41f840f5601917f75de58)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-41c264bde6f1257bd08480e52e48086ca57e3355de7ceceff66d8eca33a99e66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8891a48038bc3a7ced0686bd14da0a03074564642fbf6d7c8c93b7d9a60553f0"></a>

## cloudflare.protected_endpoints.path — cloudflare.protected_endpoints.path / 73030c082167 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.path

<a id="canonical-3f15db54066a22f881af4237f32bdf04a2c7c5ebe403e4be2511c09e6dfaab7d"></a>

Type: `"object"`. single nested block, Optional.

Path. URI Path

Upstream description:

URI Path

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-231f46723cf5c336922e8d4355ad05f0bf4ed118c5f2254148234a1185ea054a"></a>

## Direct properties — cloudflare.protected_endpoints.path / 73030c082167 / 3

<a id="canonical-3f39f639ad8695b6416adbdba62aabf5f931cf1c05435c78f58c9d2900550720"></a>

<a id="canonical-75c0adf42f4d91a5542de9da99bbaada6625541f7ea2a62670fb2dc1890084ed"></a>

## caseinsensitive property — cloudflare.protected_endpoints.path / 73030c082167 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-4db42f87134ba20eaff689d5fb4c02397ce14705860cda7792a2aec7997a0cc1"></a>

<a id="canonical-bf284e3548e05916c324b868c20d749edd439be72ba5e36bff64208e1dd93a1c"></a>

## path property — cloudflare.protected_endpoints.path / 73030c082167 / 5

Type: `"string"`. Optional.

Path. URI Path

Upstream description:

URI Path

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

<a id="canonical-0bf04816007beee854a50fbb149d09bef47efe9dd922a23997ed14470abe126b"></a>

## Next pages — cloudflare.protected_endpoints.path / 73030c082167 / 6

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4752d8f1f05f4f8588b0a6177dc59c704990e0fe0ca3ded7d611063bdf2df8ee"></a>

## cloudflare.protected_endpoints.web_client — cloudflare.protected_endpoints.web_client / 02ef02a20745 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.web_client

<a id="canonical-d9d2b4d531138187d850c7a8e2b6cf124bba186738d207ff5a851e443fb8c2aa"></a>

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

<a id="canonical-5c1ee7d03616e6746bb81e17525c499fea728679b34c8e49c304bd1dd465694b"></a>

## Direct properties — cloudflare.protected_endpoints.web_client / 02ef02a20745 / 3

- [block](resources--protected_application--reference--group-002.md#canonical-d61af53367fb74d8891dc1ae67c1bed50fbc7a6da752dd965a84a607e7c2bf5b): complete subsection reference.

- [continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8): complete subsection reference.

- [redirect](resources--protected_application--reference--group-002.md#canonical-3d78df9ae5251cd15c86faaba4a6aea0264d90510affde047fe96545555db47e): complete subsection reference.

<a id="canonical-ef6e8ca459c99acbe0c79b948f1560c3f513d22f4d06fce3a5211056a5e77a86"></a>

## Next pages — cloudflare.protected_endpoints.web_client / 02ef02a20745 / 4

- [cloudflare.protected_endpoints.web_client.block](resources--protected_application--reference--group-002.md#canonical-d61af53367fb74d8891dc1ae67c1bed50fbc7a6da752dd965a84a607e7c2bf5b)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8)
- [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-002.md#canonical-3d78df9ae5251cd15c86faaba4a6aea0264d90510affde047fe96545555db47e)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-d61af53367fb74d8891dc1ae67c1bed50fbc7a6da752dd965a84a607e7c2bf5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c98a1c9dd1da63305b4f726854c677fbcdcf905483e9cb6c533db0a6eee45c7"></a>

## cloudflare.protected_endpoints.web_client.block — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- cloudflare.protected_endpoints.web_client.block

<a id="canonical-d7b0dd9e47fcc53453851dd184216c26e93db3920f198eda9e1be94d5e18994b"></a>

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

<a id="canonical-d4145e0596147922b441a1f1f54dec258399d2e6f3444e69bc489203b6173461"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 3

<a id="canonical-40cd0c5877941547e3563dc784aeb4f130a69420c7475dc04c3b8e73e09a0681"></a>

<a id="canonical-e15f10e9f1aaf140841c347ff36d8cef807295e82915340f733fde706eaf124a"></a>

## body property — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 4

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

<a id="canonical-4f16cda1a30cf438eb57c4bc617e493d4f891627e7ed987590a57c34616a6e3f"></a>

<a id="canonical-62e4ded940c2eed901f2e8e35c41e4a4736c492c84985ee6c05680390761f921"></a>

## content_type property — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 5

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

<a id="canonical-a5e14c83d47648a6d583b7046131e3132819ac27ca9a4f7b73accfd2fb306b51"></a>

<a id="canonical-66ca6c79223291f420fcef6e29d4d9b12b215d511a2f64a6bfcf69bdb6dceafa"></a>

## status property — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 6

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

<a id="canonical-ab2f857d70de85b1cf20a45c84693c38edca8b48229b5307978d1556be4e2a9c"></a>

## Next pages — cloudflare.protected_endpoints.web_client.block / b69297e7689e / 7

- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-755c40fa583f7b52a69175bdff3a8e770b67d072e9587fdb6839b2bd6351ebb5"></a>

## cloudflare.protected_endpoints.web_client.continue — cloudflare.protected_endpoints.web_client.continue / e72788768bfc / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- cloudflare.protected_endpoints.web_client.continue

<a id="canonical-13c40c6e9289f06dd1112d121ee5b1f6af5cade9ba248276b863bc6ac48e1f92"></a>

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

<a id="canonical-a3aba088a3a7e2f365a5545befd826d12bf89fc2f8fe2c61c5d5c31ee44023a8"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue / e72788768bfc / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-6627ae33ec2849413c8678c11a1c673d86bd25fda62d17ba8944ce5aabe678b6): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-71661c5e545d04675fcf6f118c94aa9f123e9799417d1a965dd7fca8ec607333): complete subsection reference.

<a id="canonical-292fe67732aafb6da5d4b90674fd3392bae4d1e54eb1779eb57f90cfbfbec7b8"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue / e72788768bfc / 4

- [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-6627ae33ec2849413c8678c11a1c673d86bd25fda62d17ba8944ce5aabe678b6)
- [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-71661c5e545d04675fcf6f118c94aa9f123e9799417d1a965dd7fca8ec607333)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6627ae33ec2849413c8678c11a1c673d86bd25fda62d17ba8944ce5aabe678b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c65272c7192312b6d266c99a2ea10c36192519220cf4bf3cefcfb23e1a7cc55"></a>

## cloudflare.protected_endpoints.web_client.continue.add_header — cloudflare.protected_endpoints.web_client.continue.add_header / d610a9706a52 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8)
- cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-b9cf3159d2fd7e2212cf5ef268bbfe85aa740bcbfb32e90dacfac49c002603f0"></a>

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

<a id="canonical-e1cf68abc8ad8bf3c8c089c8c5eda3183a1d8622e458c47d322be1b78357f27b"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue.add_header / d610a9706a52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18d7bd44fc042a95dd9ba017892153bdd64808f3cd36bb89dbdbdd34e74b57b1"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue.add_header / d610a9706a52 / 4

- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-71661c5e545d04675fcf6f118c94aa9f123e9799417d1a965dd7fca8ec607333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af21699c0f965f8246eae16a4b8e861c52f29662311e8a43fd6dc6cb19a0edac"></a>

## cloudflare.protected_endpoints.web_client.continue.no_header — cloudflare.protected_endpoints.web_client.continue.no_header / fff03c2f172f / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8)
- cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-35b8ed51e1f4a7aa35b3f3ba30fa28f1a125c586f6825473db38b7c06d8af8cc"></a>

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

<a id="canonical-0405bd18f4d115341b265b2459c3132d64043915e5c23fcb2f531b93a13d562d"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.continue.no_header / fff03c2f172f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-057d01c0c414c5ff683392ee7d7ec55d8f929b4c3b489801704909d2326c779d"></a>

## Next pages — cloudflare.protected_endpoints.web_client.continue.no_header / fff03c2f172f / 4

- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-be047af0594892de8b833e435ebebda23f945fce79bfdf228110540ccbd964a8)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-3d78df9ae5251cd15c86faaba4a6aea0264d90510affde047fe96545555db47e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-239d82b4a1c9ec0b9fef7e4f998bbed23a36eb986248d883963532f7953abcf4"></a>

## cloudflare.protected_endpoints.web_client.redirect — cloudflare.protected_endpoints.web_client.redirect / 624be86fa454 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-bcf5b26eb495170ee86bfeeb66648457d84f50c3bac0e7e4851c3180bd4cc0e5"></a>

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

<a id="canonical-28ed32be3f1b5f446b6699f1dc21ef72429e6dacfb4d815dc6c741a3ad99219a"></a>

## Direct properties — cloudflare.protected_endpoints.web_client.redirect / 624be86fa454 / 3

<a id="canonical-f3e8b0e828e6b5d49ed4863bd2bcda858aa9e2c35c600326a0166c7172ef3d23"></a>

<a id="canonical-6300d791cdfd98e9f473a52c84f77b3192633f6534b3a185babf0734ce729f6b"></a>

## location property — cloudflare.protected_endpoints.web_client.redirect / 624be86fa454 / 4

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

<a id="canonical-e0c129b1ae365c6da4d790175f47ff57ebea807b3e07bfeaf7077038ae6df792"></a>

<a id="canonical-33f2350cd5aff1327c386750033a7b6e63fd0fb144d0235d72b53dc456a6c81a"></a>

## status property — cloudflare.protected_endpoints.web_client.redirect / 624be86fa454 / 5

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

<a id="canonical-1583e575e55b7105ca3e8b7f4a2a85a034842ec9f267ca976693272f20a93bdf"></a>

## Next pages — cloudflare.protected_endpoints.web_client.redirect / 624be86fa454 / 6

- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d45040a87b6860d8bdb955532ed9a3a23d11112cfa73af6af60f2f2ffced6265"></a>

## cloudflare.protected_endpoints.web_mobile_client — cloudflare.protected_endpoints.web_mobile_client / 283352f8740e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.web_mobile_client

<a id="canonical-9b748861d6d0c22a4062dd4599f06ff03a225ccdffa45ad2cc3be32a18b8d158"></a>

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

<a id="canonical-47635dfdc9b95a8b98a1a21644b7be5ac550d9176b24299e88118aec8a5bbf6b"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client / 283352f8740e / 3

- [block_mobile](resources--protected_application--reference--group-002.md#canonical-c568ac349ed82aa701accd8951e352de9f4dc6b4f5117e81f6ed2093fe2001f7): complete subsection reference.

- [block_web](resources--protected_application--reference--group-002.md#canonical-302130b0fc9b9a13c029d72fe28c799f4d4c4e630b9dfde99d67cbc6aeef0781): complete subsection reference.

- [continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d): complete subsection reference.

- [continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6): complete subsection reference.

- [redirect_web](resources--protected_application--reference--group-002.md#canonical-eb58ef2fe51ddbcef934c36f65057187c4acdba59bdbe5f5d02f39c5042684fd): complete subsection reference.

<a id="canonical-794e6111a729b16124cc4b7dbcad92403337c502437ea91393d63d07f523d595"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client / 283352f8740e / 4

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-002.md#canonical-c568ac349ed82aa701accd8951e352de9f4dc6b4f5117e81f6ed2093fe2001f7)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-002.md#canonical-302130b0fc9b9a13c029d72fe28c799f4d4c4e630b9dfde99d67cbc6aeef0781)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-002.md#canonical-eb58ef2fe51ddbcef934c36f65057187c4acdba59bdbe5f5d02f39c5042684fd)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-c568ac349ed82aa701accd8951e352de9f4dc6b4f5117e81f6ed2093fe2001f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a160a595c70b1c3107ce7b87d35b80a67f996b06b10963f42160421fd3f5e4e"></a>

## cloudflare.protected_endpoints.web_mobile_client.block_mobile — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- cloudflare.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-1f177f475b234766682e7580a454987b45c40cb5160999617d53d76bee5d4884"></a>

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

<a id="canonical-177d0f0593d2de929b29d97ca7f1281aa7333282c9793b03bc21b097ac9b0859"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 3

<a id="canonical-336877c6758134bd4f11cbc04f7d4c7ec252b740c1d2ef0c23ef0ef87c9c8fbf"></a>

<a id="canonical-9f6dfab44d10dfff1e3cf1efb30eaba57a2de3dedfffe2c38f42399c9f5d4b71"></a>

## body property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 4

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

<a id="canonical-03151bc852efb1572583eb5afff29edb51fdc9d093efa5d1132435dd8d9e3dd6"></a>

<a id="canonical-91770245f6b7da4d95cbfbb6f7f3f1d0e7801bafdd10a7d439c29cfd67e0a1b7"></a>

## content_type property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 5

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

<a id="canonical-554b2cee32fd0d2412ec1dadcd51aa9236b1f0aa13e758628a6b9359f9a2a61d"></a>

<a id="canonical-056755d946c8987513fc44062152bc41603c433c9486e768d65fcc2065c5667a"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 6

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

<a id="canonical-81c253ba608c427d191f61aa87a05c6d27f769752a13ccebac8f6c24b8f39684"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.block_mobile / b75077b76741 / 7

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-302130b0fc9b9a13c029d72fe28c799f4d4c4e630b9dfde99d67cbc6aeef0781"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22034f03d7dc397b95a0fb0044b7f6313323f261e02baaf4a7e7d2fff3122edf"></a>

## cloudflare.protected_endpoints.web_mobile_client.block_web — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- cloudflare.protected_endpoints.web_mobile_client.block_web

<a id="canonical-99701a6507543c569f93387d5fd0aa7140e79d33b56d9104a482a780ccb21b26"></a>

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

<a id="canonical-5b2ea567cfb031858709f1ce0322a3dbdcb580f862f4fdce7db2abb69e3be786"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 3

<a id="canonical-b25296dbfe972d13197c0381890b04dc823b337145256c5055ce5506b3c8d040"></a>

<a id="canonical-495fed4bed0b45fdb066bc972810c896f2c847438860c1614b1da8ebe5e5a1df"></a>

## body property — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 4

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

<a id="canonical-2cb2add3da283bf1100c7c474759cd7c9b2141b17e6ff88e9fb219a987b72ad9"></a>

<a id="canonical-96ccd71db8eabeec110771a29c3b4b0430d29d3297cb8e03af52d739b4452f67"></a>

## content_type property — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 5

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

<a id="canonical-dbf7a2200ce1ef20e07a0beb1cff986d3bbfeafac7743604c1b98c42b3a3efe3"></a>

<a id="canonical-fa214f35b6dfd543de018aca923b66742f88e4dc46071942ec0ba5da610b4a81"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 6

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

<a id="canonical-b8b5f675a4b11ab5c1d6d5df3ec278c1ecdc19577ec6b486d77856d6e1d92097"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.block_web / 9cba18f82388 / 7

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff7eae2f52e9fea009d33f41db8ae86cc847f3290b74d78325407edbf55529af"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 033d1ff2c86c / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-08d882d43e4a8e60ecc084cc5e6b19a12f9d911de32b20151b54e4f692611187"></a>

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

<a id="canonical-789dbcbb479a3e34d5a845044120e9614ab078fe9111e702c249a16a8c607c63"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 033d1ff2c86c / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-92f1682beed18cda2922d39923439ba8e1129a0bad2e18f4a84bbc4272fc6691): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-e57793a90bacabdb60842efe3c65890a77efaa0733c84dd31b3bc92485a5a471): complete subsection reference.

<a id="canonical-7e8f4bd592f6df7820ca29aa04acaecd9dddbf932f4cb49755b2cc55c84ab7b2"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile / 033d1ff2c86c / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-002.md#canonical-92f1682beed18cda2922d39923439ba8e1129a0bad2e18f4a84bbc4272fc6691)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-002.md#canonical-e57793a90bacabdb60842efe3c65890a77efaa0733c84dd31b3bc92485a5a471)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-92f1682beed18cda2922d39923439ba8e1129a0bad2e18f4a84bbc4272fc6691"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36e11e2228e526593f26edc634a63fa966042ff45c6686ed8c431f253b4c1189"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / 70ecb2fdfe8d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-622a48451b9f319f091e9f5f5cbf0df3c68d86864ade1b4ddfb25ba2a59bf5f3"></a>

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

<a id="canonical-0a84d8af7a37afeff730adf1e77e38701eb517ba22e2256670db43e23f9b50ea"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / 70ecb2fdfe8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55192a9508b0c7347dcbecb45a44e7e78399363f16fc9cac518e774838416b74"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header / 70ecb2fdfe8d / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-e57793a90bacabdb60842efe3c65890a77efaa0733c84dd31b3bc92485a5a471"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-176207a2cb860e5c949d7074f9ad44e22ba1fa2f7156699d5aad23b41bee7351"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / b054f5e275e1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-f3c9eb8a7d3ddd7299d5986a8c972a9bf26de7bf1d0e3f21c8f2a8623279ff33"></a>

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

<a id="canonical-eee2d53c1d8486c73964622b5a7511c62b9a1c8645aa365bb1feb6fe1b96e59a"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / b054f5e275e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8aa8b4cb4f3eb27f2a279f2713c41daab16301d40ef6efff9c0a1a6a92f03f29"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header / b054f5e275e1 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-69b435c50acdf93cbee50d574fe0ebbaea68c1e162a22b173e0ef815a74a815d)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-453c545a6f5cdd0a87bd4134d80a017db9bf3d4c93d70697503c46068d6bab5c"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web — cloudflare.protected_endpoints.web_mobile_client.continue_web / 2683ad67fa09 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-6d5cc2928083ab9ab9d9b201173e2f263c546f28339eab534421ef2992e29196"></a>

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

<a id="canonical-3ed06f40f2acdfdc6c24a986c7ffbbf03288f536979b245db567f2364042ec86"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web / 2683ad67fa09 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-ee5a5f5b2ee921d68515846c7eedd828f05c49b3ca110b1026a91680c25b60a0): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-9b5f0c9f1f4e802f7b6376acb364cbf3558e4e211f73de78fda530e6fb7b6a9f): complete subsection reference.

<a id="canonical-b1d7c8a53d5befd1ade237aadbaa9f5e838b2b82a4fae644237e4debc7bc5bf0"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web / 2683ad67fa09 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-002.md#canonical-ee5a5f5b2ee921d68515846c7eedd828f05c49b3ca110b1026a91680c25b60a0)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-002.md#canonical-9b5f0c9f1f4e802f7b6376acb364cbf3558e4e211f73de78fda530e6fb7b6a9f)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-ee5a5f5b2ee921d68515846c7eedd828f05c49b3ca110b1026a91680c25b60a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a53182aa68a2de6cfbc7c417845dbcccca3c486be978d60c03e5fa10f98a05f6"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / e133097f0e6e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6)
- cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-d66d1a7a2bcd24ccec8bc4c76f9fc8d244beedbe0168b483c992964bc1d74d12"></a>

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

<a id="canonical-77155a9b1899d7ed5eb6f1074c044829c98ed6005b980e2c3ddfa5ba4d16b10a"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / e133097f0e6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12addee01b20e8ddff7cae41dfd03ee19770f5996183493c2a7e68466e3166d6"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header / e133097f0e6e / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-9b5f0c9f1f4e802f7b6376acb364cbf3558e4e211f73de78fda530e6fb7b6a9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56b8cf683f5b5add7d226e28a067d6c471d2e4c6c4efbdf4e430bbe33cac894d"></a>

## cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 8dd81fb014fc / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6)
- cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-25c7e0c85fec933c04640822b41944dc009254c23d3b3b9fe8807b4a02ff211c"></a>

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

<a id="canonical-5428eee8381f740962b3f4012cc3c2fde1a6d6a7091f3cb51ddb1979c180a67c"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 8dd81fb014fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce71ebbd4ffee05b6f0fcf1a22b83ca48011479aa860f5a7398af2c3184c6d49"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header / 8dd81fb014fc / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1285363d7813752a3b40e6c6bf52c89e14a1c48420eace8a27975cfe5e81bdf6)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-eb58ef2fe51ddbcef934c36f65057187c4acdba59bdbe5f5d02f39c5042684fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5c6234ecb22718cf375d6f7490fb57858c7dc501c7fb0dd8aa0d0790911db06"></a>

## cloudflare.protected_endpoints.web_mobile_client.redirect_web — cloudflare.protected_endpoints.web_mobile_client.redirect_web / c2de60e7b431 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- cloudflare.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-f9647c23a9c2241c8165c0d3ab51ebd70c5fc926cfa6c145c9bd7357c146185b"></a>

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

<a id="canonical-ffae2221dc46f98ffe4cb73353c7229a043bcf0b7c5f453adf3b05251fef3f2a"></a>

## Direct properties — cloudflare.protected_endpoints.web_mobile_client.redirect_web / c2de60e7b431 / 3

<a id="canonical-4f4b771966560e369310acf738941121ad6aa9365c0a1ca84a1338784af418f6"></a>

<a id="canonical-71bed1042cfc731ef43df9cc9634d49ec82e8ab949656d3dad2930c667957d95"></a>

## location property — cloudflare.protected_endpoints.web_mobile_client.redirect_web / c2de60e7b431 / 4

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

<a id="canonical-6d9baf5782474f90c24c57532fdc1b85176c1dc9b1d5b3d4e098fadcb52f0981"></a>

<a id="canonical-9aefc9eea17dc2489f2c4d07eecb31b31e400c0243c30a000db525a0cca7e769"></a>

## status property — cloudflare.protected_endpoints.web_mobile_client.redirect_web / c2de60e7b431 / 5

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

<a id="canonical-a9f1c0613250f415e572f76d28c4ca31008dca3e3780bb5ff929196c0cc0198d"></a>

## Next pages — cloudflare.protected_endpoints.web_mobile_client.redirect_web / c2de60e7b431 / 6

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bb83e1ca3b57750b6b9de5c23cf2711f90402e1ba2e8ed3a79ce313f8f61c47"></a>

## cloudflare.trusted_clients — cloudflare.trusted_clients / 9f149f893f63 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.trusted_clients

<a id="canonical-5dfde5e2b1cf80ffb447f5765569159e2fe4cfff8466709bef2882a5e8900244"></a>

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

<a id="canonical-9585cd42b55ac51b29b2e1d51502bb7195ad3c5a39a144b4355e4ca89ecf5b12"></a>

## Direct properties — cloudflare.trusted_clients / 9f149f893f63 / 3

- [http_header](resources--protected_application--reference--group-002.md#canonical-3e7922815a8af62e0a1520afa156f62e2a3c2aaf914dc9aff466ae678c2c5573): complete subsection reference.

<a id="canonical-a1a8eb4b50557e803404a10112511757528039f712c76f0b6adf3663a952deb8"></a>

<a id="canonical-43d3ec1bacf1d5e44bad535582aad03f3b799f7379eaf3b2f3f1c82bea09b462"></a>

## ip_prefix property — cloudflare.trusted_clients / 9f149f893f63 / 4

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

- [metadata](resources--protected_application--reference--group-002.md#canonical-d4a0fffb3635d323fb42adaf0a58e5f5574edf129ff3192908abe6249c17f6f9): complete subsection reference.

<a id="canonical-1b1c08f3d8dbbac3810d62a6aca997017ce12114d3ccde7b9f6a52bbb1c9e535"></a>

## Next pages — cloudflare.trusted_clients / 9f149f893f63 / 5

- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-3e7922815a8af62e0a1520afa156f62e2a3c2aaf914dc9aff466ae678c2c5573)
- [cloudflare.trusted_clients.metadata](resources--protected_application--reference--group-002.md#canonical-d4a0fffb3635d323fb42adaf0a58e5f5574edf129ff3192908abe6249c17f6f9)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-3e7922815a8af62e0a1520afa156f62e2a3c2aaf914dc9aff466ae678c2c5573"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-070a5611ffc87f69814366065817220aceebd24dc015c5a2963f16c1883b8d28"></a>

## cloudflare.trusted_clients.http_header — cloudflare.trusted_clients.http_header / 28a4379ceac8 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- cloudflare.trusted_clients.http_header

<a id="canonical-a72319ecab94e3cf22e2709be43a2ab0e24e98dae777d0804684a413e02d14d7"></a>

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

<a id="canonical-20c29a0a3b7950a127432273995cc378de672a57322de9e231c2b0b48a1b8565"></a>

## Direct properties — cloudflare.trusted_clients.http_header / 28a4379ceac8 / 3

- [headers](resources--protected_application--reference--group-002.md#canonical-8c221172f7e177ef0cd56c32b0672308cc941dc6cbaf49b7cabe7b72d81c894d): complete subsection reference.

<a id="canonical-6b7e3db4a7fbb1a9bf3fffb0bbe22818ed241c49be96de6855bffe6d6ad24084"></a>

## Next pages — cloudflare.trusted_clients.http_header / 28a4379ceac8 / 4

- [cloudflare.trusted_clients.http_header.headers](resources--protected_application--reference--group-002.md#canonical-8c221172f7e177ef0cd56c32b0672308cc941dc6cbaf49b7cabe7b72d81c894d)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-8c221172f7e177ef0cd56c32b0672308cc941dc6cbaf49b7cabe7b72d81c894d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95f058ce854a1301eabc5d62428a3499b9b08795ca45b08cb0fe8afe3e9309f9"></a>

## cloudflare.trusted_clients.http_header.headers — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-3e7922815a8af62e0a1520afa156f62e2a3c2aaf914dc9aff466ae678c2c5573)
- cloudflare.trusted_clients.http_header.headers

<a id="canonical-fcb4ade7eb3c48cd378384820a2d4cc3329bedb497068064ef2d5bdb02765d34"></a>

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

<a id="canonical-ceb99d150c58dc3803b0638e148b5c938934e89e6f64c4c7d5325dccdcda99f7"></a>

## Direct properties — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 3

<a id="canonical-7fc26feef11b995253f93216e71771e86b8906107dad18c272ca45840a0bd6d5"></a>

<a id="canonical-66db7cb1917bc631b210f63bc8a14bc7e22886249d20834ec06c3c800bc62219"></a>

## exact property — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 4

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

<a id="canonical-b9ab3334a43295db8e451d2743f87d1808d987cdb3a8e7d1fd7c0f9c2b8503ad"></a>

<a id="canonical-901c090202132c7197e46b33806947c56c1aab767c0fb530f31923eff4259b4a"></a>

## name property — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 5

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

<a id="canonical-ac65196f0ef59e4f305d4df1199a58807a786c0ff81f4ef30a5f7bdfd58f5187"></a>

<a id="canonical-14e25fee04e7e111713c2e375884b236c8a31053115962894b9003650cd11beb"></a>

## regex property — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 6

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

<a id="canonical-dcf32d8e1d9d4bc08e2e51a5301eaa8c852195083001e2f769f3fd435fb3413e"></a>

## Next pages — cloudflare.trusted_clients.http_header.headers / 286cd1e05bdb / 7

- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-3e7922815a8af62e0a1520afa156f62e2a3c2aaf914dc9aff466ae678c2c5573)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-d4a0fffb3635d323fb42adaf0a58e5f5574edf129ff3192908abe6249c17f6f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d58fa92d59b66f64defddb5c57336b4547076e754f1bb5f7c8da9fe937fb1549"></a>

## cloudflare.trusted_clients.metadata — cloudflare.trusted_clients.metadata / f357b70a8d56 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- cloudflare.trusted_clients.metadata

<a id="canonical-a1ea7ad629d33e0d344e7752f4c75f25cc3e7212f7577ddd4fc06dfd8befeae3"></a>

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

<a id="canonical-13f8fa383f928af0a750e43e59bd04fc515b95d28ec0494cf3334dee4362d50f"></a>

## Direct properties — cloudflare.trusted_clients.metadata / f357b70a8d56 / 3

<a id="canonical-15970282e564761db1d7dacd86ad21c271e6eac70c4b94eed0a3cf319be22d39"></a>

<a id="canonical-8182697e0ffcfaf4019d97a55b5e9e9e82f3e2f12207c902da6545f556cb7154"></a>

## description_spec property — cloudflare.trusted_clients.metadata / f357b70a8d56 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-7b7d7fbc44c7804ee1b0be73deebaf34fe72fe8563e4be86a3d7faa10458a323"></a>

<a id="canonical-0d18dfaf8b93a1d674b4d56ae77e32de099b979ababa5c5ade53e43ad429750c"></a>

## name property — cloudflare.trusted_clients.metadata / f357b70a8d56 / 5

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

<a id="canonical-cb631bd55a1d435046333d34db884352afa02bfa6291c3d8e8e2036353b8267d"></a>

## Next pages — cloudflare.trusted_clients.metadata / f357b70a8d56 / 6

- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31ece6b44b19a09836cd28abfa94629e3952716d64e32dcd29a531c715d59d0c"></a>

## cloudfront — cloudfront / 3a6a600a4f47 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- cloudfront

<a id="canonical-3b163d79026cd25596af8a276de79a6ad6020516bd2678119c801b88b42eff3c"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for AWS Cloudfront.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-06ca292178a8d527dd2346df90aa9c355d539cf5114f64aeee921cecf6906036"></a>

## Direct properties — cloudfront / 3a6a600a4f47 / 3

- [aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-81d4bf9ebb98c9930afb0fffc325be34092f8d3dc005e9d97d9734e9b6649297): complete subsection reference.

- [aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-b10cf11e850ba2934e2f05ffe331449f6f0f991344ca35a37ee73d6a24b8a8ba): complete subsection reference.

<a id="canonical-cb4904b08850cb9ba80ce4401dd09f5bee5c6f326d37f46b35af7ce6a2202a7f"></a>

<a id="canonical-f3b97ebdfc900b9394fcc3a6c9313a1dcc469aebb7bcaaa81498ceadc373554c"></a>

## continue_mitigation_action_hdr property — cloudfront / 3a6a600a4f47 / 4

Type: `"string"`. Optional.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-9acc36ecd0fb6e1810881fbfcde54e76db191ba716628b7b833938ed30d2b1c4"></a>

<a id="canonical-3a8fd35b7c6df264b7b88bc0e50ad816cc25149d456240d2ef9b642c4661e02c"></a>

## data_sample property — cloudfront / 3a6a600a4f47 / 5

Type: `"number"`. Optional.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

Upstream description:

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte)

Provider validators and defaults (from schema source):

```go
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

- [disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-84b405f9de26fd3a6eecf69e856dcc71435072499475bf7b2258d2161f6d6810): complete subsection reference.

- [disable_js_insert](resources--protected_application--reference--group-002.md#canonical-899eeaaf7cc70d1a484815b3b04255e23ba771c64f92712b5af2fc706a323a22): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-09e2867c5b0439b613444759853d991137391a00f2fa31a7d58a25c8b2a2e62a): complete subsection reference.

- [js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4): complete subsection reference.

<a id="canonical-209ea247869b3bd007a8fb169b4eeccdfe39e98fa7d2c0c21b974478be4ac759"></a>

<a id="canonical-777d2a87c312235d78cb40bb333706c0aeedc34f3b342f2b0c7bfd661900300c"></a>

## loglevel property — cloudfront / 3a6a600a4f47 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
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

- [manual_js_insert](resources--protected_application--reference--group-002.md#canonical-f74be01baadf5de507626f93e13653bfbe7bfe1b5a6ea09bbb14b9df523d76cd): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209): complete subsection reference.

- [protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae): complete subsection reference.

<a id="canonical-490a991dff05e269cd674b1ef43bc6cb95b6bd31151ce3c1d57a1029842bd68c"></a>

<a id="canonical-876dbc30e08b0e6a9286c11d07e8b167e332bd9f68389ae9d888e209e7865238"></a>

## timeout property — cloudfront / 3a6a600a4f47 / 7

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
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

- [trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c): complete subsection reference.

<a id="canonical-3222af76095bbaecc3bf1e212e279d623de3cabf203f9419a43228591ab2e368"></a>

## Next pages — cloudfront / 3a6a600a4f47 / 8

- [cloudfront.aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-81d4bf9ebb98c9930afb0fffc325be34092f8d3dc005e9d97d9734e9b6649297)
- [cloudfront.aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-b10cf11e850ba2934e2f05ffe331449f6f0f991344ca35a37ee73d6a24b8a8ba)
- [cloudfront.disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-84b405f9de26fd3a6eecf69e856dcc71435072499475bf7b2258d2161f6d6810)
- [cloudfront.disable_js_insert](resources--protected_application--reference--group-002.md#canonical-899eeaaf7cc70d1a484815b3b04255e23ba771c64f92712b5af2fc706a323a22)
- [cloudfront.disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-09e2867c5b0439b613444759853d991137391a00f2fa31a7d58a25c8b2a2e62a)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.manual_js_insert](resources--protected_application--reference--group-002.md#canonical-f74be01baadf5de507626f93e13653bfbe7bfe1b5a6ea09bbb14b9df523d76cd)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-7ddf39e75a6ab5aac1a8ffd2e165185773ac088418b149a657469cf5e84f123c)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-81d4bf9ebb98c9930afb0fffc325be34092f8d3dc005e9d97d9734e9b6649297"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85ca2233edb04ca2d6d1b68bfc4336ee0dd757a65377958212131c82d2dc3323"></a>

## cloudfront.aws_configuration_id_selector — cloudfront.aws_configuration_id_selector / ef510fb6c223 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.aws_configuration_id_selector

<a id="canonical-e328fb8ff021e8b85230e8fc8302d915f641ec34db21267cc5c7511d22342dc9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws configuration id selector.

Upstream description:

List of CloudFront distributions.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-c6a955c142b0c055929ad1491dab49e2f9be226557f4c8ed115ecf7d2f18e7af"></a>

## Direct properties — cloudfront.aws_configuration_id_selector / ef510fb6c223 / 3

<a id="canonical-f3a6fbde38fe24eceb4ac92e6f13e3a29078c0517fedea7c8724f4cb7ef7c9a0"></a>

<a id="canonical-27be3d3a573b4fe8ab347a48a6df94e25e267d8d0b0672f30eb77d54269d7085"></a>

## ids property — cloudfront.aws_configuration_id_selector / ef510fb6c223 / 4

Type: `["list", "string"]`. Optional.

Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-86adb8167e8029058bd1f3bcdd5377432c29b44896331ef9fdd6930372a69801"></a>

## Next pages — cloudfront.aws_configuration_id_selector / ef510fb6c223 / 5

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-b10cf11e850ba2934e2f05ffe331449f6f0f991344ca35a37ee73d6a24b8a8ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87715f66f51ae40949c4a3a21e7df61c57eb379b5f916ee31dac03f72a74e5e1"></a>

## cloudfront.aws_configuration_tag_selector — cloudfront.aws_configuration_tag_selector / bc9a6043048e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.aws_configuration_tag_selector

<a id="canonical-d582bbeb60e2497c2c0f73bb4941a679d32f6001ff09cd8a68465f25f3efcfe0"></a>

Type: `"object"`. single nested block, Optional.

Distribution Tag List. CloudFront distribution tag list.

Upstream description:

CloudFront distribution tag list.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-60514da1a75f3878670638febcc21b14fa68af201748228d4123b83b2c6e90e3"></a>

## Direct properties — cloudfront.aws_configuration_tag_selector / bc9a6043048e / 3

<a id="canonical-c82f2f14959863604f7bce9b358cbbc232f3234642193249a8cb1bfcbc8b7a4b"></a>

<a id="canonical-0d5bf74e358fa5ad291af2006aeb122feca39784d80eaa52f54e27d9c91f0dfa"></a>

## tags property — cloudfront.aws_configuration_tag_selector / bc9a6043048e / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-a5bc2a1b40b8267740199bc2344fd67e183fe38114984c155f7b7ca6d5c428e8"></a>

## Next pages — cloudfront.aws_configuration_tag_selector / bc9a6043048e / 5

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-84b405f9de26fd3a6eecf69e856dcc71435072499475bf7b2258d2161f6d6810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75df44210fe573edc84ff5c7e49853140da11b11aba038d3561b5ad609617f4a"></a>

## cloudfront.disable_aws_configuration — cloudfront.disable_aws_configuration / f57f58b2c7c1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.disable_aws_configuration

<a id="canonical-c4fad582262a682012df3d9ff968216d37ccb8daf70f4bdff2cac1d81b5c117e"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_aws_configuration = {}
```

<a id="canonical-78824825823d623c4a6d0332d3391bf9fde6f1bab48729a7c578ef1b726d2817"></a>

## Direct properties — cloudfront.disable_aws_configuration / f57f58b2c7c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0d941432c7510b3612f38b042959a1bb01a06e776c4a8ec964640b2110c5eac"></a>

## Next pages — cloudfront.disable_aws_configuration / f57f58b2c7c1 / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-899eeaaf7cc70d1a484815b3b04255e23ba771c64f92712b5af2fc706a323a22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375bca29bd692b60cc61bc003dfc5e48e09eba6ae16307229357232378bdd51b"></a>

## cloudfront.disable_js_insert — cloudfront.disable_js_insert / d7fa32972df5 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.disable_js_insert

<a id="canonical-6c4abbca2784e82299b1cd302146adb3e02d2f13d4e3de54b859632125bd9293"></a>

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

<a id="canonical-69875b5b49d955fc40878064313142f03d43115ba194d90f7547edf917baa7d9"></a>

## Direct properties — cloudfront.disable_js_insert / d7fa32972df5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c73f229d11d026ceceafb2696850296f8d6a4e6531e13299d7256a6c347d2f95"></a>

## Next pages — cloudfront.disable_js_insert / d7fa32972df5 / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-09e2867c5b0439b613444759853d991137391a00f2fa31a7d58a25c8b2a2e62a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e375739aa08656674bca14622c4ff52a70d822d15840f91e0ff4f3a15f63a1"></a>

## cloudfront.disable_mobile_sdk — cloudfront.disable_mobile_sdk / af706ebbe48e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.disable_mobile_sdk

<a id="canonical-0b22b8e5eae9c12a735a317a02338769ed0dfad44792788df88f3a18949abf28"></a>

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

<a id="canonical-54ef1c4572fb8dfdc33311413158f32316dbdd7514a295e2b3f1e4e06cc85b31"></a>

## Direct properties — cloudfront.disable_mobile_sdk / af706ebbe48e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c124268bf634394c793c9ef86e5c438742ce3216d60fdaa052f5b7cb81cdc37"></a>

## Next pages — cloudfront.disable_mobile_sdk / af706ebbe48e / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c2dbb994debd20aad9a419b619e0b7bb9ab2e33c826794085c9148a0616e4dc"></a>

## cloudfront.js_insertion_rules — cloudfront.js_insertion_rules / d4bac3f45714 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.js_insertion_rules

<a id="canonical-e4bcdfb97ac85043699329beb8615ed70d01931692aaf002e0e02a111ccd016b"></a>

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

<a id="canonical-298361eac55105cf7009722d2b71a397ea828bfdcd52657939b061184c4b02cd"></a>

## Direct properties — cloudfront.js_insertion_rules / d4bac3f45714 / 3

- [exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93): complete subsection reference.

<a id="canonical-a25ff865223ff51cda59e5b8123fc5ae1876478afbf225730e5a82c12c4ca22d"></a>

<a id="canonical-d486a614490920a3067654869741846e5764b5872f551f962c91d75298029796"></a>

## javascript_location property — cloudfront.js_insertion_rules / d4bac3f45714 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-5c5d0b0fea4339238f3cf948ad93e0d449828925ca9dcab005fd20b7c974f73d"></a>

<a id="canonical-8b4052d7b19e6947a5b20dede2f955b68252981c07017b4ab2f8a98856108e55"></a>

## javascript_mode property — cloudfront.js_insertion_rules / d4bac3f45714 / 5

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

<a id="canonical-686438b2692c84f2ff7fd14b97322884d821d96bd45cc834e4bec92779eae114"></a>

<a id="canonical-da4e8a72bbcc243f6d1f26a339759c69b88dfca23daadfc6733f1be2f303f2db"></a>

## js_download_path property — cloudfront.js_insertion_rules / d4bac3f45714 / 6

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

- [rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754): complete subsection reference.

<a id="canonical-e87930c0bb162d884d24857962b2900c5bcea4fbad02188019a046153ed4cbbb"></a>

## Next pages — cloudfront.js_insertion_rules / d4bac3f45714 / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f992910954822bb47abffa43da6278fe160ba39ec3d9673570cfe81b5c6dc5f7"></a>

## cloudfront.js_insertion_rules.exclude_list — cloudfront.js_insertion_rules.exclude_list / 8759da03d4f3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- cloudfront.js_insertion_rules.exclude_list

<a id="canonical-513efa0f572c03400cc9a9a78acf3f2e9680a4fdd0bf4ca673ee5be70d167a9b"></a>

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

<a id="canonical-dd0fd7033eeb0e284f770cab5116daf6f2bb2fc094700487672d3087ae61f0cb"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list / 8759da03d4f3 / 3

- [any_domain](resources--protected_application--reference--group-002.md#canonical-0a313cd2a2f73ea8ca66213c4dc93bc896a3cf9c0436346ea6ff524d976181df): complete subsection reference.

- [domain](resources--protected_application--reference--group-002.md#canonical-2b79dccc49e427cd7a34890bc91ad9bdfb9a8809ffd1fc5c4fc6851f2ec63f7b): complete subsection reference.

- [metadata](resources--protected_application--reference--group-002.md#canonical-841b40edc842481fa4fcb60f7b941a0a61a5aeb706aca66db562247bfa3aedc7): complete subsection reference.

- [path](resources--protected_application--reference--group-002.md#canonical-899c90657d698930f4fdc0fca03d60cf5121cf52cf65aa8d67fb7b80531c2cfd): complete subsection reference.

<a id="canonical-7deca6b6cb25d33452d3d97f99b8066b0c43dc7568b1722b489c8fe1c33b7eec"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list / 8759da03d4f3 / 4

- [cloudfront.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-002.md#canonical-0a313cd2a2f73ea8ca66213c4dc93bc896a3cf9c0436346ea6ff524d976181df)
- [cloudfront.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-002.md#canonical-2b79dccc49e427cd7a34890bc91ad9bdfb9a8809ffd1fc5c4fc6851f2ec63f7b)
- [cloudfront.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-002.md#canonical-841b40edc842481fa4fcb60f7b941a0a61a5aeb706aca66db562247bfa3aedc7)
- [cloudfront.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-002.md#canonical-899c90657d698930f4fdc0fca03d60cf5121cf52cf65aa8d67fb7b80531c2cfd)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-0a313cd2a2f73ea8ca66213c4dc93bc896a3cf9c0436346ea6ff524d976181df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4dcba07d73e212c93b0f6e0f0b6f8b02a0d78f0fd80822a458bdd29c5e6dcd6"></a>

## cloudfront.js_insertion_rules.exclude_list.any_domain — cloudfront.js_insertion_rules.exclude_list.any_domain / 2b22d1c9a4c6 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- cloudfront.js_insertion_rules.exclude_list.any_domain

<a id="canonical-f25ea3f5bca89f0c6ebac56cded2527733089337b2131f02f2f4286a500e4bc0"></a>

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

<a id="canonical-1c574f39ff1dc50649c8be9b0c6af2a26cb620e7823ebebef4f68c3e4406a377"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.any_domain / 2b22d1c9a4c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50d01622868a825dd8df28ee335c439513421116a08b7efa6f12b2bb5657a293"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.any_domain / 2b22d1c9a4c6 / 4

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2b79dccc49e427cd7a34890bc91ad9bdfb9a8809ffd1fc5c4fc6851f2ec63f7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b94c1c9faf4a1d32a7c2b91871493e4206f671f1164b085309a359862b9ad517"></a>

## cloudfront.js_insertion_rules.exclude_list.domain — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- cloudfront.js_insertion_rules.exclude_list.domain

<a id="canonical-142d157c900d1864e8931dd6da19be1da307d8c300682248665880090e150a8d"></a>

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

<a id="canonical-61a1b67c71261df885a0cf5cfc5b6b7bd5454d0df337d94d1664acc11d920721"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 3

<a id="canonical-ff98dfa46f0621a4401a4176258dedf14c3998b122fadc3824aa50279cdb3238"></a>

<a id="canonical-29c43e8fb415b7cdf3751395f3afb8126701d6d7c16c341d9c575c9b77d01eaf"></a>

## exact_value property — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 4

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

<a id="canonical-66d873d7f651efb02de1f01606eba923c8616b02d6aaf9eeafb908f93db53fda"></a>

<a id="canonical-9a10f04fc66296a060dc0da12ae27184e3329b3479407f22b69aefee7189ddec"></a>

## regex_value property — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 5

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

<a id="canonical-c4d88e39c99c6eaccd2cfbddb9f067260563a46b6e6b564e769e51f54fb93809"></a>

<a id="canonical-11e3598cf70492ba9dfcf2e4820179ca1b7c19ad4970761d1a0f3a7e560d7495"></a>

## suffix_value property — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 6

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

<a id="canonical-30fbe3b001f5fc7d37ca031145c1a88368b8b7752b86caa593c70696ca599cf0"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.domain / 78237bf5c2a3 / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-841b40edc842481fa4fcb60f7b941a0a61a5aeb706aca66db562247bfa3aedc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5485526c01d09a599c62fa577811d357d5e27b79d31fa21fd0e5d2f660a43527"></a>

## cloudfront.js_insertion_rules.exclude_list.metadata — cloudfront.js_insertion_rules.exclude_list.metadata / 28f12fe049a3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- cloudfront.js_insertion_rules.exclude_list.metadata

<a id="canonical-3d743622f969dc2198f396b1d3431b8200af63d2d9c28a0356f798cf50db4b5f"></a>

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

<a id="canonical-ce8f4c7c6c63b9e5e805e43c851c0f15f2ec4e5d7377913205e6042bff0f738e"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.metadata / 28f12fe049a3 / 3

<a id="canonical-9001d0f68a54c800d56cb30b80ed40c9dec1871acee561052c54805d7fe2faa4"></a>

<a id="canonical-2ec7872c9ff1fc985eb5c08b57ad169ab9e4037a93670bc80713bb716d7a5267"></a>

## description_spec property — cloudfront.js_insertion_rules.exclude_list.metadata / 28f12fe049a3 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-ad57ea524a551dce4b56c2962f495647fcd720c626e63973dea8671aa3e45de6"></a>

<a id="canonical-a9257b617cfdbee6c16c790b782e42fec706e2292115121e03a721dd50c58b15"></a>

## name property — cloudfront.js_insertion_rules.exclude_list.metadata / 28f12fe049a3 / 5

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

<a id="canonical-d109a3d5600fe7201dacc82daf6da4715505190c618bd9491f44a092724ca98e"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.metadata / 28f12fe049a3 / 6

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-899c90657d698930f4fdc0fca03d60cf5121cf52cf65aa8d67fb7b80531c2cfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c22746c398900bc025e3504a7a2779e3ffcf6acec1e7060662ef421f51d5a3a"></a>

## cloudfront.js_insertion_rules.exclude_list.path — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- cloudfront.js_insertion_rules.exclude_list.path

<a id="canonical-66fe972d83c54877aefed9fa4b067c9b71161971a89d5979ea73f0e0e94c43bb"></a>

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

<a id="canonical-2382c4b1e3b37838cf3c137f49002f9bf8c32a7f44d0bd51ef187bc71591a237"></a>

## Direct properties — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 3

<a id="canonical-7eb00272f4d133184b00ae7ef484b14592c535540f59a017d7a303cd3a56a608"></a>

<a id="canonical-75df17d5862f6b12dbbb937f7eb4bc7bf3612fd489a923eb20f2a931495ec1a7"></a>

## path property — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 4

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

<a id="canonical-36254fe61f2c6729f7707f002b12647dd5091e39ac8c3b79fa158226238b23a8"></a>

<a id="canonical-af25033b4fbd15047ba916f83a0848a8119f838ec510e03880f0d4203ea28a87"></a>

## prefix property — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 5

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

<a id="canonical-7c9020a8e03f56f777f1cf2688df44c845f3c997d82cd512525a5d4a2f11e037"></a>

<a id="canonical-196b38291b0db80e1d1851088c169bb2218307af5a2c9e928f56368ceaf77687"></a>

## regex property — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 6

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

<a id="canonical-fee28ff0632272a3978136c5a32a7bfb6a7f61f74edcfae7efe2c9513122443b"></a>

## Next pages — cloudfront.js_insertion_rules.exclude_list.path / 1af8e94cf5fc / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-cc63671df9a1aad2027f0325e747bafbe714b30f514edae88729bd8792ab8b93)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fea07da7afb6b96d4c227c1a878752da18adced7daa18a300b0e2aa8cf19a872"></a>

## cloudfront.js_insertion_rules.rules — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- cloudfront.js_insertion_rules.rules

<a id="canonical-aec65af9e99ebc34af80f0920491c90d525b65be7a05fab4f2425f3b8197f7bd"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-4d03e01e0530925745b1624ab4916c777e1d3f6841acd348003420b38e5c25e0"></a>

## Direct properties — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 3

- [any_domain](resources--protected_application--reference--group-002.md#canonical-d77b35e977ea50a46172b0f4a32c7b7aa62d81774c3a2157e15056618cbf2ef3): complete subsection reference.

- [domain](resources--protected_application--reference--group-002.md#canonical-0b9b8ea32766e816a825a9ebebdbc37ea273e2ed48b86816ca1e34df7306fe74): complete subsection reference.

<a id="canonical-3bd4732ad653bf3b2f2fefe9bdc03ab7c7b70e5fa552da20d314be3640993448"></a>

<a id="canonical-fc758cab291d19b04661f2cdf54d56e9a5d1e31b7462e9d873ff85c7a049847a"></a>

## exact_path property — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 4

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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

<a id="canonical-8b20ee2e59fd61580deb396841cfd2a25fc3d35abc8a972d9b46f956469d0170"></a>

<a id="canonical-4111dc6a5e81661157bc5ebfd569872546ac8daf1ec571621f97104b7bfd5f3c"></a>

## glob property — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

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

- [metadata](resources--protected_application--reference--group-002.md#canonical-6b2abd602d9f5e41c3c8b5a3eea1515fc56d345b3f02dcffcb3542a88086b71a): complete subsection reference.

<a id="canonical-eb363bce831a18a6ae7d01efe26dcba2bcdc788047942742b26b93ae3566ae94"></a>

<a id="canonical-7c7eb033a00ec10f029a89ff9472b8d6128a84289c9c00a17dfc659ff2ada497"></a>

## prefix property — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-df674ce2dd231fe93f759c96cb6fd2afe02fc12a20dc116ccb269a578de3e14a"></a>

## Next pages — cloudfront.js_insertion_rules.rules / cbe207c5d58f / 7

- [cloudfront.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-002.md#canonical-d77b35e977ea50a46172b0f4a32c7b7aa62d81774c3a2157e15056618cbf2ef3)
- [cloudfront.js_insertion_rules.rules.domain](resources--protected_application--reference--group-002.md#canonical-0b9b8ea32766e816a825a9ebebdbc37ea273e2ed48b86816ca1e34df7306fe74)
- [cloudfront.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-002.md#canonical-6b2abd602d9f5e41c3c8b5a3eea1515fc56d345b3f02dcffcb3542a88086b71a)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-d77b35e977ea50a46172b0f4a32c7b7aa62d81774c3a2157e15056618cbf2ef3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0501dd6c1b08478d5e95d27f8b78285da5deaa6b538e07b3a46731c38d48464a"></a>

## cloudfront.js_insertion_rules.rules.any_domain — cloudfront.js_insertion_rules.rules.any_domain / 548ce67b9c50 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- cloudfront.js_insertion_rules.rules.any_domain

<a id="canonical-9ed4c5145b9015ac58df873ec7168017cc18345d59033d4997f7d8fffafcc53e"></a>

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

<a id="canonical-65bdc8f61a39e60c8986b0efc50db16a6bfac1ccd7551531d52598d73671a009"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.any_domain / 548ce67b9c50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a9eb67611f7e9fdfb54078e2c09a013de843ece0a197cc2293c0945ede3e4a3"></a>

## Next pages — cloudfront.js_insertion_rules.rules.any_domain / 548ce67b9c50 / 4

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-0b9b8ea32766e816a825a9ebebdbc37ea273e2ed48b86816ca1e34df7306fe74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae629721a097c766736239be40280251c4373a344287c986ce3e1f527eca378d"></a>

## cloudfront.js_insertion_rules.rules.domain — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- cloudfront.js_insertion_rules.rules.domain

<a id="canonical-0740a08332ffd2e98195fe85cf195b171a76df40b27f6f7064031f13160151b3"></a>

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

<a id="canonical-e06247e3083c3b89d1381fce7f6be500551611aa48520938ab4e4aabc91c38e2"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 3

<a id="canonical-374ea9e9b8e32426631fab5663189d9d488ef9efc260f4039ccc74e574e17899"></a>

<a id="canonical-cf1ba2320892c0f0b488168880eb76ba6c1ca6d21207287e95c703ed53611694"></a>

## exact_value property — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 4

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

<a id="canonical-60be98993a2be1af5bac84faeb50932e575db617963213c4f4954f9ddf639b29"></a>

<a id="canonical-ce05a215b5128531f30fe6f8d75a622d3fd22103ab843ebfeaec17d775dff3d9"></a>

## regex_value property — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 5

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

<a id="canonical-e4e0470128574b7bc07f42ea1f4fb19e5aeccea04c19ec84486adcccbb90e352"></a>

<a id="canonical-fa89c19ab34727e81cc9e957acda43f173883b66fa26abb10c821b1a298221c0"></a>

## suffix_value property — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 6

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

<a id="canonical-ed0c454e95ea2be32bc78251b45fdb483863cafb7cceea77ad2bd146a8495d75"></a>

## Next pages — cloudfront.js_insertion_rules.rules.domain / c973d438ab51 / 7

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6b2abd602d9f5e41c3c8b5a3eea1515fc56d345b3f02dcffcb3542a88086b71a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1440ae4085bde291cc0eca737fdf4403ed64585c3ecd108b5de6457158e0ad0"></a>

## cloudfront.js_insertion_rules.rules.metadata — cloudfront.js_insertion_rules.rules.metadata / ae18eeae24f2 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-2df3289afff14b842fce07283169bf5ea0842fe088a96f647726d5a2c29ec1a4)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- cloudfront.js_insertion_rules.rules.metadata

<a id="canonical-4d11e7564ae09b9ddd9fa38c5af7cfc69a5fe0d6a8fdda220a696055f37c6f72"></a>

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

<a id="canonical-3c30e6f517056e5991c9f623e181e4499000430bebb880b91f7381e94389eb48"></a>

## Direct properties — cloudfront.js_insertion_rules.rules.metadata / ae18eeae24f2 / 3

<a id="canonical-46e95e2e243dda92d7812bfce9584dba5bbf2dfe6005123dad2b3d81917018a1"></a>

<a id="canonical-f947f32e4ca4af5d4c111dc7ccbee4b205811b6f2a05cc7603f401af1e1e2dbb"></a>

## description_spec property — cloudfront.js_insertion_rules.rules.metadata / ae18eeae24f2 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-51abe93bb3c9c92c170d3b9d3cac8322e0a6b4f5f323403ee6fc894d29d62d3b"></a>

<a id="canonical-1ca26ec1e33a7fc16d2f2e848361cf0aafbcc5bd1c430a708e0c1e7f9de6a03d"></a>

## name property — cloudfront.js_insertion_rules.rules.metadata / ae18eeae24f2 / 5

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

<a id="canonical-5698d1694af4c901b89acc74ad82ba2cd9dd27ec2c2d9c17f207cba93f0251b9"></a>

## Next pages — cloudfront.js_insertion_rules.rules.metadata / ae18eeae24f2 / 6

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-5562b19f44bae585023eb899a280383f3caf486f14167ca31455bbe22e81b754)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-f74be01baadf5de507626f93e13653bfbe7bfe1b5a6ea09bbb14b9df523d76cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f5f056452869a2903ddeaf530a0063cfaff085faa46c1eb7eaf40ef9e23892"></a>

## cloudfront.manual_js_insert — cloudfront.manual_js_insert / 0d91b8d3aca1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.manual_js_insert

<a id="canonical-d5d40e8c0a4002482e6e5f8ee9d336d892f567a2ddf0df8bbe8ce1a3920726ac"></a>

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

<a id="canonical-b49382576fe6147d61d07ecdca3bda62cff28ca969f125d67fdc28da9043688a"></a>

## Direct properties — cloudfront.manual_js_insert / 0d91b8d3aca1 / 3

<a id="canonical-0c068ab569fea3f3df56aa2a55c5a2b2cb6f602db17de20425955fd51b0cadad"></a>

<a id="canonical-6dc65ca0810d5345a94972dc2ca64e80dc30e3bc6de9cab86d490809b3856d47"></a>

## javascript_mode property — cloudfront.manual_js_insert / 0d91b8d3aca1 / 4

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

<a id="canonical-050f1a9274bd8da4191c2dec0ba0c80f4a6c84ce0b640dd9c9a847dcf258e7ff"></a>

<a id="canonical-d55748254ee0eae002c47bf7ffd5f5f82488e878b2d6264a77523a62b6725177"></a>

## js_download_path property — cloudfront.manual_js_insert / 0d91b8d3aca1 / 5

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

<a id="canonical-5b707983b3edcd2882857d802433edc479e3a142470eec354fede68ff4814b61"></a>

## Next pages — cloudfront.manual_js_insert / 0d91b8d3aca1 / 6

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e37092c2e7ed72f352d04cf4364d9ba0cbcb2b4fb6377842cd7cb05297bc49"></a>

## cloudfront.mobile_sdk_config — cloudfront.mobile_sdk_config / 10b2d178b698 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.mobile_sdk_config

<a id="canonical-034acbfa5b732cb6741535d55072adf9ceae85dc60b62d644bd712cabc1d028b"></a>

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

<a id="canonical-449d6555bb03ed79805475c33efd9222270e29dd533331c819bec554af4166ad"></a>

## Direct properties — cloudfront.mobile_sdk_config / 10b2d178b698 / 3

- [mobile_identifier](resources--protected_application--reference--group-002.md#canonical-27af8a569a472b2a77c5ea06a6c1ebdbcd31bf63c37fe6714d01b6ef4f90e0db): complete subsection reference.

<a id="canonical-707a7ba56b44e1d6e48d518700970ab2ea625572ec693a44de57e54544525e32"></a>

## Next pages — cloudfront.mobile_sdk_config / 10b2d178b698 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-002.md#canonical-27af8a569a472b2a77c5ea06a6c1ebdbcd31bf63c37fe6714d01b6ef4f90e0db)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-27af8a569a472b2a77c5ea06a6c1ebdbcd31bf63c37fe6714d01b6ef4f90e0db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-961fd1cc9407ed9ed855c170275792ae487d9f6b49ca2980ea2a748001a06ec5"></a>

## cloudfront.mobile_sdk_config.mobile_identifier — cloudfront.mobile_sdk_config.mobile_identifier / a6f68b10d581 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209)
- cloudfront.mobile_sdk_config.mobile_identifier

<a id="canonical-4bb491bbc9a218e4e6db240e6273455e23c1d540f024b712373d12452552e839"></a>

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

<a id="canonical-f5342ee30ad8482f80e87fa6379381481c702e82eede9890365269cf0d054bac"></a>

## Direct properties — cloudfront.mobile_sdk_config.mobile_identifier / a6f68b10d581 / 3

- [headers](resources--protected_application--reference--group-002.md#canonical-eaf35f2cf5cb33880f8c8c2c1d3d0bac2abafe872044f85aac52121cc6b330ce): complete subsection reference.

<a id="canonical-fb910e0dd4852ce5f97f73ad053f7f5070faae429c667264246e9862c30b7501"></a>

## Next pages — cloudfront.mobile_sdk_config.mobile_identifier / a6f68b10d581 / 4

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-002.md#canonical-eaf35f2cf5cb33880f8c8c2c1d3d0bac2abafe872044f85aac52121cc6b330ce)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-eaf35f2cf5cb33880f8c8c2c1d3d0bac2abafe872044f85aac52121cc6b330ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
