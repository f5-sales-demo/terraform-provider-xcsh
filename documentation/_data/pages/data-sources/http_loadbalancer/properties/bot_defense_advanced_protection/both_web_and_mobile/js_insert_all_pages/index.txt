---
page_title: "bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages"
subcategory: "Load Balancing"
description: "Insert Bot Defense JavaScript in all pages."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insert all pages"], "body_bytes": 2443, "body_sha256": "sha256:a1b2df72a57075df24bbbd8c0b0f668f370f85f8ddf1e60500a7daa9fb006f32", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1201210022313210-2103230132310311-1203113220232331-2011213222111123-2031100030103232-1323210221200230-2213103212001111-0103100131110322", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insert_all_pages"], "schema_version": 1, "sections": [{"aliases": ["javascript location"], "anchor": "schema-bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages--javascript_location", "description": "All inside networks. Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insert_all_pages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insert_all_pages", "javascript_location"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insert_all_pages/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Insert Bot Defense JavaScript in all pages.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages

<a id="section"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages.

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

<a id="schema-bot_defense_advanced_protection--both_web_and_mobile--js_insert_all_pages--javascript_location"></a>

### javascript_location property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
