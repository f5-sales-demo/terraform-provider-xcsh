---
page_title: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control"
subcategory: "Load Balancing"
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules rules app firewall detection control"], "body_bytes": 4182, "body_sha256": "sha256:970f4b278e6c520a8838d5f13db0265c7fbcc7f95df819fa7ad95dcf4bc8e624", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "path": "documentation/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_attack_type_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_signature_contexts--signature_id", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts:RequiredListObjectAttributes:signature_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "type": "requires"}], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_violation_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/)
- [waf_exclusion.waf_exclusion_inline_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
