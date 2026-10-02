---
page_title: "waf_exclusion_rules.app_firewall_detection_control"
subcategory: ""
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["waf exclusion rules app firewall detection control"], "body_bytes": 3329, "body_sha256": "sha256:e41eb2d1adcfca3ad0860cc595848a1147247ba57aab8c9f35940670bbe121cd", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "path": "documentation/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112", "registry_path": "docs/guides/data-sources--waf_exclusion_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_attack_type_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_violation_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- waf_exclusion_rules.app_firewall_detection_control

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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
