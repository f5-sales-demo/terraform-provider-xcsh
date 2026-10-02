---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.redirect"
subcategory: ""
description: "Web Client Redirect request to a custom URI."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints redirect"], "body_bytes": 4382, "body_sha256": "sha256:510976ba8f70006a5e48315c780cba6d9b3803fd07281efa939c832ed11b3ed4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:redirect", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/redirect/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3233313030203230-2021122100310233-3233331123102203-2020322333010032-0312301131101032-1110011131202022-1020221231103102-0130202222330211", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "redirect"], "schema_version": 1, "sections": [{"aliases": ["status"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--redirect--status", "description": "HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status.. Possible values are", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "redirect", "status"], "syntax": "attribute", "type": "string"}, {"aliases": ["uri"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--redirect--uri", "description": "URI location for redirect may be relative or absolute.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "redirect", "uri"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/redirect/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Web Client Redirect request to a custom URI.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.redirect

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.redirect

<a id="section"></a>

Type: `"single"`. Computed.

Web Client Redirect request to a custom URI.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--redirect--status"></a>

### status property

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

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--redirect--uri"></a>

### uri property

Type: `"string"`. Computed.

URI location for redirect may be relative or absolute.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
