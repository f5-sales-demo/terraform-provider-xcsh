---
page_title: "bot_defense_advanced_protection.both_web_and_mobile"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.both_web_and_mobile for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4097, "body_sha256": "sha256:e05eb960e1ce9367942324ca803098c798793d12677fa050ce4b7608731adbaa", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_js_insert", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:disable_mobile_sdk", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages_except", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:web"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.both_web_and_mobile for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense_advanced_protection.both_web_and_mobile

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
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

- [disable_js_insert](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--disable_js_insert.md): complete subsection reference.

- [disable_mobile_sdk](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--disable_mobile_sdk.md): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insertion_rules.md): complete subsection reference.

- [mobile](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--mobile.md): complete subsection reference.

- [mobile_sdk_config](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--mobile_sdk_config.md): complete subsection reference.

- [web](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--web.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--disable_js_insert.md)
- [bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--disable_mobile_sdk.md)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages.md)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages_except.md)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--js_insertion_rules.md)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--mobile.md)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--mobile_sdk_config.md)
- [bot_defense_advanced_protection.both_web_and_mobile.web](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--both_web_and_mobile--web.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
