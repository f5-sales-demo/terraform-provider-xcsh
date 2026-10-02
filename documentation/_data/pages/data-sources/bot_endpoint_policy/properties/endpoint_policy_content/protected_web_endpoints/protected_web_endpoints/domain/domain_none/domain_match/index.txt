---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_none.domain_match"
subcategory: ""
description: "Domain Matcher(s). Domain Matchers."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints domain domain none domain match"], "body_bytes": 3296, "body_sha256": "sha256:a2614bc9948110b6617e1a0311ee9955cafc358c148fd123b33f47db7d490acd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_none:domain_match", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_none", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_none/domain_match/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0223111133310031-0131012233233022-0220121310020122-1320130122023122-2112212331111223-1222333103030232-1032221301033103-3033201023110100", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_none", "domain_match"], "schema_version": 1, "sections": [{"aliases": ["negation"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--negation", "description": "Select from one of the Negation Operator. - NO: No - YES: Yes. Possible values are `NO`, `YES`. Defaults to `NO`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_none:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_none", "domain_match", "negation"], "syntax": "attribute", "type": "string"}, {"aliases": ["operator"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--operator", "description": "Select from one of the Comparison Operator. - EXACT: exact value - CONTAIN: contain value - START_WITH: start with value - END_WITH: end with value. Possible values are `EXACT`, `CONTAIN`, `START_WITH`, `END_WITH`. Defaults to `EXACT`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_none:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_none", "domain_match", "operator"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--value", "description": "Value. Domain Matcher Value.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_none:domain_match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_none", "domain_match", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_none/domain_match/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Domain Matcher(s). Domain Matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_none.domain_match

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_none/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_none.domain_match

<a id="section"></a>

Type: `"list"`. Computed.

Domain Matcher(s). Domain Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--negation"></a>

### negation property

Type: `"string"`. Computed.

\[Enum: NO|YES\] Select from one of the Negation Operator. - NO: No - YES: Yes. Possible values are
\`NO\`, \`YES\`. Defaults to \`NO\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--operator"></a>

### operator property

Type: `"string"`. Computed.

\[Enum: EXACT|CONTAIN|START\_WITH|END\_WITH\] Select from one of the Comparison Operator. - EXACT:
exact value - CONTAIN: contain value - START\_WITH: start with value - END\_WITH: end with value.
Possible values are \`EXACT\`, \`CONTAIN\`, \`START\_WITH\`, \`END\_WITH\`. Defaults to \`EXACT\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_none--domain_match--value"></a>

### value property

Type: `"string"`. Computed.

Value. Domain Matcher Value.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_none/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
