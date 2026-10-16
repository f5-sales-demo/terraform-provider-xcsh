---
page_title: "routes.bot_defense_javascript_injection.javascript_tags.tag_attributes"
subcategory: ""
description: "Add the tag attributes you want to include in your Javascript tag."
xcsh_docs: {"aliases": ["routes bot defense javascript injection javascript tags tag attributes"], "body_bytes": 3683, "body_sha256": "sha256:e69d0fb181a6a8d3b5192098af336440d1df300d04e8581fdd4b88c257fe2919", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes", "parent_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "path": "documentation/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2131210232233001-0332212010223020-3103022313003020-1331311221202031-3330101313013002-3231212212132323-0330203211012311-0321003012023313", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "tag_attributes"], "schema_version": 1, "sections": [{"aliases": ["routes bot defense javascript injection javascript tags tag attributes javascript tag"], "anchor": "schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--javascript_tag", "description": "Select from one of the predefined tag attributes.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "tag_attributes", "javascript_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes bot defense javascript injection javascript tags tag attributes tag value"], "anchor": "schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--tag_value", "description": "Add the tag attribute value.", "document_id": "xcsh-docs:data-sources:route:properties:routes:bot_defense_javascript_injection:javascript_tags:tag_attributes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags", "tag_attributes", "tag_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Add the tag attributes you want to include in your Javascript tag.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["routeCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/)
- [routes.bot_defense_javascript_injection.javascript_tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/)
- routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="section"></a>

Type: `"list"`. Computed.

Add the tag attributes you want to include in your Javascript tag.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--javascript_tag"></a>

### javascript_tag property

Type: `"string"`. Computed.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--tag_value"></a>

### tag_value property

Type: `"string"`. Computed.

Value. Add the tag attribute value.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```
