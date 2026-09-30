---
page_title: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3989, "body_sha256": "sha256:55ce0a7b291b073161a7a55e051777884446c308b9b39d78c43caec96b3f03e3", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "path": "documentation/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/)
- [waf_exclusion.waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
