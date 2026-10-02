---
page_title: "bot_defense"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Bot Defense Policy."
xcsh_docs: {"aliases": ["bot defense"], "body_bytes": 4492, "body_sha256": "sha256:42c0436e8853326e5634ec7213bd5bcf04dd16379c3c8077a6e3a038650fb1e6", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:enable_cors_support", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense:ConflictingObjectAttributes:disable_cors_support,enable_cors_support", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense:ConflictingObjectAttributes:disable_cors_support,enable_cors_support", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:enable_cors_support", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense"], "schema_version": 1, "sections": [{"aliases": ["disable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "disable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:enable_cors_support", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "enable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy"], "anchor": "section", "description": "This defines various configuration OPTIONS for Bot Defense policy.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_mobile_sdk,mobile_sdk_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:ConflictingObjectAttributes:disable_mobile_sdk,mobile_sdk_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy:RequiredObjectAttributes:protected_app_endpoints", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "type": "requires"}], "schema_path": ["bot_defense", "policy"], "syntax": "block", "type": "object"}, {"aliases": ["regional endpoint"], "anchor": "schema-bot_defense--regional_endpoint", "description": "Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "regional_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeout"], "anchor": "schema-bot_defense--timeout", "description": "The timeout for the inference check, in milliseconds.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines various configuration OPTIONS for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- bot_defense

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_cors_support",
    "enable_cors_support")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/disable_cors_support/): complete subsection reference.

- [enable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/enable_cors_support/): complete subsection reference.

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/): complete subsection reference.

<a id="schema-bot_defense--regional_endpoint"></a>

### regional_endpoint property

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AUTO",
    "US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-bot_defense--timeout"></a>

### timeout property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

## Next pages

- [bot_defense.disable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/disable_cors_support/)
- [bot_defense.enable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/enable_cors_support/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
