---
page_title: "routes.bot_defense_javascript_injection.javascript_tags"
subcategory: ""
description: "Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first."
xcsh_docs: {"aliases": ["routes bot defense javascript injection javascript tags"], "body_bytes": 3782, "body_sha256": "sha256:472c00b57643c048bff6416db13d7804468905c47df63e30436310b069252cfc", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "parent_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection", "path": "documentation/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0032020112201122-1012130210333101-0230310321112302-3001301011031320-1330333031033033-3002212331120302-3202100220110013-3210300021012002", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags"], "schema_version": 1, "sections": [{"aliases": ["javascript url"], "anchor": "schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url", "description": "Please enter the full URL (include domain and path), or relative path.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "javascript_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tag attributes"], "anchor": "section", "description": "Add the tag attributes you want to include in your Javascript tag.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "tag_attributes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.bot_defense_javascript_injection.javascript_tags

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/)
- routes.bot_defense_javascript_injection.javascript_tags

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

<a id="schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url"></a>

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

- [tag_attributes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/)
- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
