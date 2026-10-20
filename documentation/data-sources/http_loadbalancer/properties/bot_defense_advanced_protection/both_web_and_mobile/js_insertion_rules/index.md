---
page_title: "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules"
subcategory: "Load Balancing"
description: "This defines custom JavaScript insertion rules for Bot Defense Policy."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insertion rules"], "body_bytes": 1649, "body_sha256": "sha256:a6d825b512a719bdd27cb924372434a3238d02cc0587829f0599b4c59d2949fd", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1211002210021201-2202322300130310-2023100220133211-1321333331210022-1231132311233203-2212123130220103-2311030203032003-1321023320321120", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insertion rules exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules", "exclude_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insertion rules rules"], "anchor": "section", "description": "Required list of pages to insert Bot Defense client JavaScript.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

## Direct properties

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/): complete subsection reference.

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/rules/): complete subsection reference.
