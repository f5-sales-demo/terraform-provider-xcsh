---
page_title: "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts"
subcategory: ""
description: "Bot Names to be excluded for the defined match criteria."
xcsh_docs: {"aliases": ["waf exclusion rules app firewall detection control exclude bot name contexts"], "body_bytes": 3223, "body_sha256": "sha256:b481a66c1b463decd97be59919ff6120e6bc67317258d7c2e5f3607673d7ee7e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "path": "documentation/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2023312020310013-0223003320122301-3302023002030222-3221000023213303-1213233220000111-2003211211003301-0232203033320200-1310220320231132", "registry_path": "docs/guides/resources--waf_exclusion_policy--reference--group-001.md", "relationships": [{"anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_bot_name_contexts"], "schema_version": 1, "sections": [{"aliases": ["bot name"], "anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_bot_name_contexts", "bot_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Names to be excluded for the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/)
- waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name"></a>

### bot_name property

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
