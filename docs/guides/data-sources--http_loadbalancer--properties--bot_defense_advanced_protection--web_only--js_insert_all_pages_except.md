---
page_title: "bot_defense_advanced_protection.web_only.js_insert_all_pages_except"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only.js_insert_all_pages_except for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2502, "body_sha256": "sha256:bd4beaca516268916dcbfedb6cefe02cb167f24d9367f7b99e1713a865892c8f", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only.js_insert_all_pages_except for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insert_all_pages_except

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except

<a id="section"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages with the exceptions.

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

- [exclude_list](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list.md): complete subsection reference.

<a id="schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--javascript_location"></a>

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

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list.md)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
