---
page_title: "bot_defense_advanced_protection.both_web_and_mobile"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.both_web_and_mobile for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5237, "body_sha256": "sha256:0057be82e27a17430ae5d833eb48a1d8677c8c2fe0623e84eedc4a832b760803", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_js_insert", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_mobile_sdk", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages_except", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:web"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.both_web_and_mobile for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- bot_defense_advanced_protection.both_web_and_mobile

<a id="section"></a>

Type: `"single"`. Computed.

Both Web &amp; Mobile. Both Web and Mobile configuration.

Upstream description:

Both Web and Mobile configuration.

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

## Direct properties

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_js_insert/): complete subsection reference.

- [disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_mobile_sdk/): complete subsection reference.

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/): complete subsection reference.

- [mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile/): complete subsection reference.

- [mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/): complete subsection reference.

- [web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/web/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_js_insert/)
- [bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/disable_mobile_sdk/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages_except/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/)
- [bot_defense_advanced_protection.both_web_and_mobile.web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/web/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
