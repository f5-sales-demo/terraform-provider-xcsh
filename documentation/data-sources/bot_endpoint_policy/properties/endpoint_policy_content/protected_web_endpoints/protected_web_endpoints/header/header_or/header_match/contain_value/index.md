---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or.header_match.contain_value"
subcategory: ""
description: "Configuration parameter for contain value."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints header header or header match contain value"], "body_bytes": 3329, "body_sha256": "sha256:57bd73bf94a9f1f74756fa349f72839482598c621ae6e6e28970c3eeabbab819", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:header:header_or:header_match:contain_value", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:header:header_or:header_match", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/header_or/header_match/contain_value/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0011103001010332-1232212003201102-1312222233310211-1312030011321120-2220022023003112-2332323201311333-2201003013011320-1102333231022322", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "header", "header_or", "header_match", "contain_value"], "schema_version": 1, "sections": [{"aliases": ["case insensitive"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--case_insensitive", "description": "Case-Insensitive. Case insensitive checker.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:header:header_or:header_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "header", "header_or", "header_match", "contain_value", "case_insensitive"], "syntax": "attribute", "type": "bool"}, {"aliases": ["not"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--not", "description": "Not(!). Not checker.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:header:header_or:header_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "header", "header_or", "header_match", "contain_value", "not"], "syntax": "attribute", "type": "bool"}, {"aliases": ["value"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--value", "description": "Value. Query Matcher Value.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:header:header_or:header_match:contain_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "header", "header_or", "header_match", "contain_value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/header_or/header_match/contain_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for contain value.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or.header_match.contain_value

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/header_or/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or.header_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/header_or/header_match/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or.header_match.contain_value

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for contain value.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--case_insensitive"></a>

### case_insensitive property

Type: `"bool"`. Computed.

Case-Insensitive. Case insensitive checker.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--not"></a>

### not property

Type: `"bool"`. Computed.

Not(!). Not checker.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--header--header_or--header_match--contain_value--value"></a>

### value property

Type: `"string"`. Computed.

Value. Query Matcher Value.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.header.header_or.header_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/header/header_or/header_match/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
