---
page_title: "bot_defense_advanced_protection.web_only"
subcategory: "Load Balancing"
description: "Web only configuration."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only"], "body_bytes": 3552, "body_sha256": "sha256:2ef8c4546113ac69b5221c45e9f4d436c940b5124607256ac089b4a7e4e3078f", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:disable_js_insert", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:web"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only"], "schema_version": 1, "sections": [{"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:disable_js_insert", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages except"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages with the exceptions.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only web"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:web", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "web"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Web only configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- bot_defense_advanced_protection.web_only

<a id="section"></a>

Type: `"single"`. Computed.

Web. Web only configuration.

Upstream description:

Web only configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

## Direct properties

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/disable_js_insert/): complete subsection reference.

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/): complete subsection reference.

- [web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/web/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.web_only.disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/disable_js_insert/)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages/)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/)
- [bot_defense_advanced_protection.web_only.web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/web/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
