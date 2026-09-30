---
page_title: "bot_defense.policy.js_insert_all_pages_except"
subcategory: "Load Balancing"
description: "bot_defense.policy.js_insert_all_pages_except for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2143, "body_sha256": "sha256:1a8479ee84b7353dcbbd6ec634cf5bc8c20897f1898cab8dda4d23a5c7180d93", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.js_insert_all_pages_except for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.js_insert_all_pages_except

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.js_insert_all_pages_except

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

- [exclude_list](data-sources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except--exclude_list.md): complete subsection reference.

<a id="schema-bot_defense--policy--js_insert_all_pages_except--javascript_location"></a>

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

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except--exclude_list.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
