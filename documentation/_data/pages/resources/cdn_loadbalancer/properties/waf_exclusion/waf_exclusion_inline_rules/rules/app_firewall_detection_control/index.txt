---
page_title: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4165, "body_sha256": "sha256:e85cc66231b34df1bb5e9e93789a147268aede5da8cc30a7430300219810f1ea", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "path": "documentation/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/)
- [waf_exclusion.waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
