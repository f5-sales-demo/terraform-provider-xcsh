---
page_title: "routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4427, "body_sha256": "sha256:286ce0d1a7c82b999dcedb2054564022bc55e37798000808e3facb52744168e9", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection:javascript_tags:tag_attributes"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection:javascript_tags", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/javascript_tags/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "bot_defense_javascript_injection", "javascript_tags"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/javascript_tags/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

<a id="section"></a>

Type: `"list"`. Computed.

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot
Javascript should be added first.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-routes--simple_route--advanced_options--bot_defense_javascript_injection--javascript_tags--javascript_url"></a>

### javascript_url property

Type: `"string"`. Computed.

Please enter the full URL (include domain and path), or relative path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/javascript_tags/tag_attributes/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/javascript_tags/tag_attributes/)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
