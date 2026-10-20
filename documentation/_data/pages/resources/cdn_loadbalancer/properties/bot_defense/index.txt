---
page_title: "bot_defense"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Bot Defense Policy."
xcsh_docs: {"aliases": ["bot defense"], "body_bytes": 2965, "body_sha256": "sha256:a8e731ba72bad9eb12650b479c31f3dab76cc37f2530e94f337ae036d87e4e24", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:enable_cors_support", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense"], "schema_version": 1, "sections": [{"aliases": ["bot defense disable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "disable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense enable cors support"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:enable_cors_support", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "enable_cors_support"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy"], "anchor": "section", "description": "This defines various configuration OPTIONS for Bot Defense policy.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense regional endpoint"], "anchor": "schema-bot_defense--regional_endpoint", "description": "Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "regional_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot defense timeout", "duration"], "anchor": "schema-bot_defense--timeout", "description": "The timeout for the inference check, in milliseconds.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "timeout"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This defines various configuration OPTIONS for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
