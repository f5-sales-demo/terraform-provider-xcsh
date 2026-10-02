---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and.query_match.contain_value"
subcategory: ""
description: "Configuration parameter for contain value."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints query query and query match contain value"], "body_bytes": 3307, "body_sha256": "sha256:e3dc60b8a73ec0b6c2c65f7e7c350e4b96db552972d27745bd923b9643b30f7e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:query:query_and:query_match:contain_value", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:query:query_and:query_match", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/query_and/query_match/contain_value/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0313100130121130-1203003021101202-2102003233332131-1320222121020000-1223110023232303-0032020200302312-3011233230320221-0131230102111102", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "query", "query_and", "query_match", "contain_value"], "schema_version": 1, "sections": [{"aliases": ["case insensitive"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--case_insensitive", "description": "Case-Insensitive. Case insensitive checker.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:query:query_and:query_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "query", "query_and", "query_match", "contain_value", "case_insensitive"], "syntax": "attribute", "type": "bool"}, {"aliases": ["not"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--not", "description": "Not(!). Not checker.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:query:query_and:query_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "query", "query_and", "query_match", "contain_value", "not"], "syntax": "attribute", "type": "bool"}, {"aliases": ["value"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--value", "description": "Value. Query Matcher Value.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:query:query_and:query_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "query", "query_and", "query_match", "contain_value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/query_and/query_match/contain_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for contain value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and.query_match.contain_value

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/query_and/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and.query_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/query_and/query_match/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and.query_match.contain_value

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for contain value.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--case_insensitive"></a>

### case_insensitive property

Type: `"bool"`. Computed.

Case-Insensitive. Case insensitive checker.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--not"></a>

### not property

Type: `"bool"`. Computed.

Not(!). Not checker.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--query--query_and--query_match--contain_value--value"></a>

### value property

Type: `"string"`. Computed.

Value. Query Matcher Value.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.query.query_and.query_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/query/query_and/query_match/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
