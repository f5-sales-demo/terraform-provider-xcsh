---
page_title: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3452, "body_sha256": "sha256:a28cd9e52197cd304447be3bbc35f6d01bbc3ec5786a2b23a57af7d6c6bf33a0", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "path": "docs/guides/resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [waf_exclusion](resources--http_loadbalancer--properties--waf_exclusion.md)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md)
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

- [exclude_attack_type_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_attack_type_contexts.md): complete subsection reference.

- [exclude_bot_name_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_bot_name_contexts.md): complete subsection reference.

- [exclude_signature_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_signature_contexts.md): complete subsection reference.

- [exclude_violation_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_violation_contexts.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_attack_type_contexts.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_bot_name_contexts.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_signature_contexts.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control--exclude_violation_contexts.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
