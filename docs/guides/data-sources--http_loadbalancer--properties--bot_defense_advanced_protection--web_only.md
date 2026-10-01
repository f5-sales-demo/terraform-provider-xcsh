---
page_title: "bot_defense_advanced_protection.web_only"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2805, "body_sha256": "sha256:f0bda9e515ff15b4623579ac24bead79fdf5c95c914d364d3d63463032ca34f8", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:disable_js_insert", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:web"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
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

- [disable_js_insert](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--disable_js_insert.md): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md): complete subsection reference.

- [web](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--web.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.web_only.disable_js_insert](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--disable_js_insert.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md)
- [bot_defense_advanced_protection.web_only.web](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--web.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
