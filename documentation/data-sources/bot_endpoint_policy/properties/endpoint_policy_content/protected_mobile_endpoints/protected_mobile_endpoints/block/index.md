---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.block"
subcategory: ""
description: "Web Client Block request and respond with custom content."
xcsh_docs: {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints block"], "body_bytes": 5454, "body_sha256": "sha256:f2f28d088341b7a40d9321a4040e4394c152dae32f3f5ed4c3eed992b51997b4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block:name_value_pair"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/block/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2133331201212212-0112113200212102-0002030132022333-3101122002223210-0103121132011203-0030310002120310-2030303312001013-1223133230102002", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "block"], "schema_version": 1, "sections": [{"aliases": ["body"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block--body", "description": "Body. Request or response body content", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "block", "body"], "syntax": "attribute", "type": "string"}, {"aliases": ["name value pair"], "anchor": "section", "description": "Response Header Name and Value Pair. Response Header Pair.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block:name_value_pair", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "block", "name_value_pair"], "syntax": "attribute", "type": "object"}, {"aliases": ["status"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block--status", "description": "HTTP response status codes for block action EmptyStatusCode response codes means it is not specified Continue status code Switching Protocols status code Processing (WebDAV) status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content.. Possible values are", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "block", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/block/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Web Client Block request and respond with custom content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.block

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.block

<a id="section"></a>

Type: `"single"`. Computed.

Web Client Block request and respond with custom content.

## Direct properties

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block--body"></a>

### body property

Type: `"string"`. Computed.

Body. Request or response body content

- [name_value_pair](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/block/name_value_pair/): complete subsection reference.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block--status"></a>

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

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.block.name_value_pair](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/block/name_value_pair/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
