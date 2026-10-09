---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.block"
subcategory: ""
description: "Web Client Block request and respond with custom content."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints block"], "body_bytes": 4702, "body_sha256": "sha256:3a52c4422fcb808b0191a8d749e5dc11d65a2b362fd9d3780d9fb331ad09b74c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block:name_value_pair"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1202133233203333-1232022313203232-0231211200101312-2103323122120030-0221131203021112-0223103300233210-1133111320221000-2231300303201031", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "block"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content protected web endpoints protected web endpoints block body"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--block--body", "description": "Body. Request or response body content", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "block", "body"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content protected web endpoints protected web endpoints block name value pair"], "anchor": "section", "description": "Response Header Name and Value Pair. Response Header Pair.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block:name_value_pair", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "block", "name_value_pair"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content protected web endpoints protected web endpoints block status"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--block--status", "description": "HTTP response status codes for block action EmptyStatusCode response codes means it is not specified Continue status code Switching Protocols status code Processing (WebDAV) status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content.. Possible values are", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "block", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/block/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Web Client Block request and respond with custom content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
