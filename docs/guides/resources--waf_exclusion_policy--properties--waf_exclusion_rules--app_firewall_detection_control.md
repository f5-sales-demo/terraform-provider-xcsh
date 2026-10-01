---
page_title: "waf_exclusion_rules.app_firewall_detection_control"
subcategory: ""
description: "waf_exclusion_rules.app_firewall_detection_control for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2780, "body_sha256": "sha256:b91cfb837de278686e81639a775cd61ce14cfe1bb9e3242bdbfbc06e2a997eec", "canonical_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "child_ids": ["xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules", "path": "docs/guides/resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control.md", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion_rules.app_firewall_detection_control for xcsh_waf_exclusion_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
- [Property reference](resources--waf_exclusion_policy--reference.md)
- [waf_exclusion_rules](resources--waf_exclusion_policy--properties--waf_exclusion_rules.md)
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

- [exclude_attack_type_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_attack_type_contexts.md): complete subsection reference.

- [exclude_bot_name_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts.md): complete subsection reference.

- [exclude_signature_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts.md): complete subsection reference.

- [exclude_violation_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_violation_contexts.md): complete subsection reference.

## Next pages

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_attack_type_contexts.md)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts.md)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts.md)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](resources--waf_exclusion_policy--properties--waf_exclusion_rules--app_firewall_detection_control--exclude_violation_contexts.md)
- [waf_exclusion_rules](resources--waf_exclusion_policy--properties--waf_exclusion_rules.md)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md)
