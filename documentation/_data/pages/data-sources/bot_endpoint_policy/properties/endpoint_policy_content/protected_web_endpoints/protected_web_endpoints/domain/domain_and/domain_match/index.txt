---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_and.domain_match"
subcategory: ""
description: "Domain Matcher(s). Domain Matchers."
xcsh_docs: {"aliases": ["endpoint policy content protected web endpoints protected web endpoints domain domain and domain match"], "body_bytes": 2865, "body_sha256": "sha256:2618c35b92ead4e9d32d081be734c414d57897728e02b67b0ddf76d425d94ab8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_and:domain_match", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_and", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_and/domain_match/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2203132002111221-0003323032323302-2223110200210222-3300220012122213-3121031022301013-2222103102213001-3002103111133121-2012211021000221", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_and", "domain_match"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content protected web endpoints protected web endpoints domain domain and domain match negation"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--negation", "description": "Select from one of the Negation Operator. - NO: No - YES: Yes. Possible values are `NO`, `YES`. Defaults to `NO`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_and:domain_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_and", "domain_match", "negation"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content protected web endpoints protected web endpoints domain domain and domain match operator"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--operator", "description": "Select from one of the Comparison Operator. - EXACT: exact value - CONTAIN: contain value - START_WITH: start with value - END_WITH: end with value. Possible values are `EXACT`, `CONTAIN`, `START_WITH`, `END_WITH`. Defaults to `EXACT`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_and:domain_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_and", "domain_match", "operator"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content protected web endpoints protected web endpoints domain domain and domain match value"], "anchor": "schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--value", "description": "Value. Domain Matcher Value.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:domain:domain_and:domain_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "domain", "domain_and", "domain_match", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_and/domain_match/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Domain Matcher(s). Domain Matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_and.domain_match

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/domain/domain_and/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.domain.domain_and.domain_match

<a id="section"></a>

Type: `"list"`. Computed.

Domain Matcher(s). Domain Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--negation"></a>

### negation property

Type: `"string"`. Computed.

\[Enum: NO|YES\] Select from one of the Negation Operator. - NO: No - YES: Yes. Possible values are
\`NO\`, \`YES\`. Defaults to \`NO\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--operator"></a>

### operator property

Type: `"string"`. Computed.

\[Enum: EXACT|CONTAIN|START\_WITH|END\_WITH\] Select from one of the Comparison Operator. - EXACT:
exact value - CONTAIN: contain value - START\_WITH: start with value - END\_WITH: end with value.
Possible values are \`EXACT\`, \`CONTAIN\`, \`START\_WITH\`, \`END\_WITH\`. Defaults to \`EXACT\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--domain--domain_and--domain_match--value"></a>

### value property

Type: `"string"`. Computed.

Value. Domain Matcher Value.
