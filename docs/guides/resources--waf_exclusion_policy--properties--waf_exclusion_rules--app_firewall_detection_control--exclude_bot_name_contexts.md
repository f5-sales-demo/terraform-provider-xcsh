---
page_title: "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts"
subcategory: ""
description: "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2818, "body_sha256": "sha256:21d4b9c2655992d392c36b260ddcdc95f67f21fc97b46ee64805c3b111f72a05", "canonical_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "child_ids": [], "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "path": "docs/guides/resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts.md", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_bot_name_contexts"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts for xcsh_waf_exclusion_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
- [Property reference](resources--waf_exclusion_policy--reference.md)
- [waf_exclusion_rules](resources--waf_exclusion_policy--properties--waf_exclusion_rules.md)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control.md)
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

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control.md)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
