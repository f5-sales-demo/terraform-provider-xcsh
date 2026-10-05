---
page_title: "routes.bot_defense_javascript_injection"
subcategory: ""
description: "Bot Defense Javascript Injection Configuration for inline bot defense deployments."
xcsh_docs: {"aliases": ["routes bot defense javascript injection"], "body_bytes": 2812, "body_sha256": "sha256:2a5ae2a0f9cdbb9315185c5390971d072e9342fc630b05ab0810f6d9012ea406", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection:javascript_tags"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/bot_defense_javascript_injection/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3012131313022222-2001213223301312-2233030002013101-2003113330213301-2010231110200010-3301033100022212-2131321012232112-0103110012233022", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.bot_defense_javascript_injection:RequiredObjectAttributes:javascript_tags", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "bot_defense_javascript_injection"], "schema_version": 1, "sections": [{"aliases": ["routes bot defense javascript injection javascript location"], "anchor": "schema-routes--bot_defense_javascript_injection--javascript_location", "description": "All inside networks. Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes bot defense javascript injection javascript tags"], "anchor": "section", "description": "Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.", "document_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url", "enforcement": "provider-schema", "group": "routes.bot_defense_javascript_injection.javascript_tags:RequiredListObjectAttributes:javascript_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "type": "requires"}], "schema_path": ["routes", "bot_defense_javascript_injection", "javascript_tags"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/bot_defense_javascript_injection/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Bot Defense Javascript Injection Configuration for inline bot defense deployments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.bot_defense_javascript_injection

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- routes.bot_defense_javascript_injection

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Javascript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
```

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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--bot_defense_javascript_injection--javascript_location"></a>

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

- [javascript_tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection.javascript_tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
