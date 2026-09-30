---
page_title: "bot_defense_advanced_protection.web_only.js_insert_all_pages"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only.js_insert_all_pages for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2561, "body_sha256": "sha256:8ce3501d8eee7a39e4334fd4f1cf0e6b75de89bf3616e34473e57d3d08e4d13e", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only.js_insert_all_pages for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense_advanced_protection.web_only.js_insert_all_pages

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- bot_defense_advanced_protection.web_only.js_insert_all_pages

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_defense_advanced_protection--web_only--js_insert_all_pages--javascript_location"></a>

### javascript_location property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
