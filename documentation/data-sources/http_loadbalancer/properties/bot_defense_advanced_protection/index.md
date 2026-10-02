---
page_title: "bot_defense_advanced_protection"
subcategory: "Load Balancing"
description: "Bot Defense Advanced Protection - replaces BotDefenseAdvancedType."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection"], "body_bytes": 2267, "body_sha256": "sha256:e631e46dcab19444cab2fdde9234a4064919e6e1b2261d20120358b2151c92d6", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection"], "schema_version": 1, "sections": [{"aliases": ["both web and mobile"], "anchor": "section", "description": "Both Web and Mobile configuration.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile"], "syntax": "attribute", "type": "object"}, {"aliases": ["mobile only"], "anchor": "section", "description": "Mobile only configuration.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "mobile_only"], "syntax": "attribute", "type": "object"}, {"aliases": ["web only"], "anchor": "section", "description": "Web only configuration.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- bot_defense_advanced_protection

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

## Direct properties

- [both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/): complete subsection reference.

- [mobile_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/mobile_only/): complete subsection reference.

- [web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [bot_defense_advanced_protection.mobile_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/mobile_only/)
- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
