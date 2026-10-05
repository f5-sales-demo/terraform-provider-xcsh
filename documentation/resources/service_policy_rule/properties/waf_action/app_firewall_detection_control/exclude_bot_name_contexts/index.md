---
page_title: "waf_action.app_firewall_detection_control.exclude_bot_name_contexts"
subcategory: ""
description: "Bot Names to be excluded for the defined match criteria."
xcsh_docs: {"aliases": ["waf action app firewall detection control exclude bot name contexts"], "body_bytes": 3134, "body_sha256": "sha256:80d5c20bd4d76d2c542bf50e763381bb4cb714696062627548c840ad2455e4d2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "path": "documentation/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "waf_action.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "schema_version": 1, "sections": [{"aliases": ["waf action app firewall detection control exclude bot name contexts bot name"], "anchor": "schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts", "bot_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot Names to be excluded for the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.app_firewall_detection_control.exclude_bot_name_contexts

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/)
- [waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
