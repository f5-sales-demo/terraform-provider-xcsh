---
page_title: "bot_defense_advanced_protection.both_web_and_mobile"
subcategory: "Load Balancing"
description: "Both Web and Mobile configuration."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile"], "body_bytes": 3045, "body_sha256": "sha256:bfac929dbc10e84ca04f98ea88d5138f090c298992f7238e48b7ec0a27f8f7ad", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_mobile_sdk", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:web"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile"], "schema_version": 1, "sections": [{"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_js_insert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile disable mobile sdk"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_mobile_sdk", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "disable_mobile_sdk"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insert all pages"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insert_all_pages"], "syntax": "block", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insert all pages except"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages with the exceptions.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages_except", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insert_all_pages_except"], "syntax": "block", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules"], "syntax": "block", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile mobile"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "mobile"], "syntax": "block", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile mobile sdk config"], "anchor": "section", "description": "Mobile Request Identifier Headers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "mobile_sdk_config"], "syntax": "block", "type": "object"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile web"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:web", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "web"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Both Web and Mobile configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- bot_defense_advanced_protection.both_web_and_mobile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Both Web &amp; Mobile. Both Web and Mobile configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
both_web_and_mobile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_js_insert/): complete subsection reference.

- [disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_mobile_sdk/): complete subsection reference.

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/): complete subsection reference.

- [mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile/): complete subsection reference.

- [mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/): complete subsection reference.

- [web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/web/): complete subsection reference.
