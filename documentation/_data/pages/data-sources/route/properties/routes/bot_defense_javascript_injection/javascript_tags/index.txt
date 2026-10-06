---
page_title: "routes.bot_defense_javascript_injection.javascript_tags"
subcategory: ""
description: "Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first."
xcsh_docs: {"aliases": ["routes bot defense javascript injection javascript tags"], "body_bytes": 3268, "body_sha256": "sha256:0217fd3ebeb2e6c8e86c8fc30b998be609918f9942cd38aabce31ee9a971434e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "parent_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection", "path": "documentation/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0032020112201122-1012130210333101-0230310321112302-3001301011031320-1330333031033033-3002212331120302-3202100220110013-3210300021012002", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags"], "schema_version": 1, "sections": [{"aliases": ["routes bot defense javascript injection javascript tags javascript url"], "anchor": "schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url", "description": "Please enter the full URL (include domain and path), or relative path.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "javascript_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes bot defense javascript injection javascript tags tag attributes"], "anchor": "section", "description": "Add the tag attributes you want to include in your Javascript tag.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "tag_attributes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
