---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block"
subcategory: ""
description: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 5382, "body_sha256": "sha256:a2e2002bf55833449fd972bfa4974805b0f8c297980340a80ac11e0b2bacf793", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block:name_value_pair"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/index.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block

<a id="section"></a>

Type: `"single"`. Computed.

Web Client Block request and respond with custom content.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--block--body"></a>

### body property

Type: `"string"`. Computed.

Body. Request or response body content

- [name_value_pair](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/name_value_pair/): complete subsection reference.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--block--status"></a>

### status property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|SwitchingProtocols|Processing|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|TeaPot|EnhanceYourCalm|UnprocessableEntity|Locked|FailedDependency|ReservedforWebDAV|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|NoResponse|RetryWith|Blockedby|UnavailableForLegalReasons|ClientClosedRequest|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|BandwidthLimitExceeded|NotExtended|NetworkAuthenticationRequired|NetworkReadRimeoutError|NetworkConnectTimeoutError\]
HTTP response status codes for block action EmptyStatusCode response codes means it is not specified
Continue status code Switching Protocols status code Processing (WebDAV) status code OK status code
Created status code Accepted status code Non Authoritative Information status code No Content..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`SwitchingProtocols\`, \`Processing\`,
\`OK\`, \`Created\`, \`Accepted\`, \`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`,
\`PartialContent\`, \`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`,
\`MovedPermanently\`, \`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`TeaPot\`, \`EnhanceYourCalm\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`ReservedforWebDAV\`, \`UpgradeRequired\`, \`PreconditionRequired\`,
\`TooManyRequests\`, \`RequestHeaderFieldsTooLarge\`, \`NoResponse\`, \`RetryWith\`, \`Blockedby\`,
\`UnavailableForLegalReasons\`, \`ClientClosedRequest\`, \`InternalServerError\`,
\`NotImplemented\`, \`BadGateway\`, \`ServiceUnavailable\`, \`GatewayTimeout\`,
\`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`, \`InsufficientStorage\`, \`LoopDetected\`,
\`BandwidthLimitExceeded\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`,
\`NetworkReadRimeoutError\`, \`NetworkConnectTimeoutError\`. Defaults to \`EmptyStatusCode\`.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block.name_value_pair](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/name_value_pair/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
