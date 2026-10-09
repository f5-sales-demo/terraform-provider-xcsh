---
page_title: "bot_defense.policy.js_insert_all_pages"
subcategory: "Load Balancing"
description: "Insert Bot Defense JavaScript in all pages."
xcsh_docs: {"aliases": ["bot defense policy js insert all pages"], "body_bytes": 1731, "body_sha256": "sha256:499e6677cfb3354dab091a0405de193395c0bdefd78758b242bf8b379314776a", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1003302212322301-2302122101132100-0211101131320001-2323202301302120-0223302322201203-3201011311303032-0102333133311303-2203212031001122", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insert_all_pages"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy js insert all pages javascript location"], "anchor": "schema-bot_defense--policy--js_insert_all_pages--javascript_location", "description": "All inside networks. Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages", "javascript_location"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Insert Bot Defense JavaScript in all pages.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insert_all_pages

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/)
- bot_defense.policy.js_insert_all_pages

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

<a id="schema-bot_defense--policy--js_insert_all_pages--javascript_location"></a>

### javascript_location property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

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
