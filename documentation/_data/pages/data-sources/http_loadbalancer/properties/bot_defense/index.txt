---
page_title: "bot_defense"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Bot Defense Policy."
xcsh_docs: {"aliases": ["bot defense"], "body_bytes": 3644, "body_sha256": "sha256:10c1bf3dbb5618fbeac8567ee2abc2b936d4defc26394130946dde90aa534cac", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:disable_cors_support", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:enable_cors_support", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense"], "schema_version": 1, "sections": [{"aliases": ["bot defense disable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:disable_cors_support", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "disable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense enable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:enable_cors_support", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "enable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy"], "anchor": "section", "description": "This defines various configuration OPTIONS for Bot Defense policy.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense regional endpoint"], "anchor": "schema-bot_defense--regional_endpoint", "description": "Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "regional_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot defense timeout", "duration"], "anchor": "schema-bot_defense--timeout", "description": "The timeout for the inference check, in milliseconds.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This defines various configuration OPTIONS for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- bot_defense

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Additional upstream details:

This defines various configuration OPTIONS for Bot Defense Policy.

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

OneOf alternatives in this subsection:

- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/#section)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/#section)
- [disable_bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/disable_bot_defense/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/disable_cors_support/): complete subsection reference.

- [enable_cors_support](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/enable_cors_support/): complete subsection reference.

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/): complete subsection reference.

<a id="schema-bot_defense--regional_endpoint"></a>

### regional_endpoint property

Type: `"string"`. Computed.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

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

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
