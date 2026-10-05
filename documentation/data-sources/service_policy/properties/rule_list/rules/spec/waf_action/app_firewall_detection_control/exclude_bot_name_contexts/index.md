---
page_title: "rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts"
subcategory: "Security"
description: "Bot Names to be excluded for the defined match criteria."
xcsh_docs: {"aliases": ["rule list rules spec waf action app firewall detection control exclude bot name contexts"], "body_bytes": 3418, "body_sha256": "sha256:5da5159817b9c53663c8205753953553bb421db2cc1065b1d930a8e5106772f3", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030", "registry_path": "docs/guides/data-sources--service_policy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec waf action app firewall detection control exclude bot name contexts bot name"], "anchor": "schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts", "bot_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot Names to be excluded for the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="section"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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

## Direct properties

<a id="schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name"></a>

### bot_name property

Type: `"string"`. Computed.

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

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
