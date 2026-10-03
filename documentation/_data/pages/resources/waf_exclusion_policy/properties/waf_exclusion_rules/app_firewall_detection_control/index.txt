---
page_title: "waf_exclusion_rules.app_firewall_detection_control"
subcategory: ""
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["waf exclusion rules app firewall detection control"], "body_bytes": 3429, "body_sha256": "sha256:84fe67777eaf90693a99e8683dc37298c9c4bcb82ae89f390fb172e417b0eb9d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules", "path": "documentation/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313", "registry_path": "docs/guides/resources--waf_exclusion_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion rules app firewall detection control exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_attack_type_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf exclusion rules app firewall detection control exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf exclusion rules app firewall detection control exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--signature_id", "enforcement": "provider-schema", "group": "waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts:RequiredListObjectAttributes:signature_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "type": "requires"}], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf exclusion rules app firewall detection control exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_violation_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- waf_exclusion_rules.app_firewall_detection_control

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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
