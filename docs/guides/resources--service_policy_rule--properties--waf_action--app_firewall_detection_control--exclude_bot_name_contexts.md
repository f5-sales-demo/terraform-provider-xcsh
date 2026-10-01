---
page_title: "waf_action.app_firewall_detection_control.exclude_bot_name_contexts"
subcategory: ""
description: "waf_action.app_firewall_detection_control.exclude_bot_name_contexts for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2828, "body_sha256": "sha256:dac0138e46ff1c05507951f4254bb4cbc02436c7b9f46d9fad789c0d830dd436", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "path": "docs/guides/resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action.app_firewall_detection_control.exclude_bot_name_contexts for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.app_firewall_detection_control.exclude_bot_name_contexts

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [waf_action](resources--service_policy_rule--properties--waf_action.md)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

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

<a id="schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name"></a>

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

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
