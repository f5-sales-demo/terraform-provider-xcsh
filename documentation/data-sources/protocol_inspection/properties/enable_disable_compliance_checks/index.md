---
page_title: "enable_disable_compliance_checks"
subcategory: ""
description: "enable_disable_compliance_checks for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1983, "body_sha256": "sha256:2bed90acfe9f5bd9ffaaa3fb75f4b1364c06e140488a8f0ef0913e00bdf110fc", "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks"], "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "documentation/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/index.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["enable_disable_compliance_checks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_disable_compliance_checks for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_compliance_checks

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/)
- enable_disable_compliance_checks

<a id="section"></a>

Type: `"single"`. Computed.

Enable Disable Compliance Checks Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compliance_check_choice": "[\"disable_compliance_checks\",\"enable_compliance_checks\"]"
}
```

## Direct properties

- [disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/): complete subsection reference.

- [enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/): complete subsection reference.

## Next pages

- [enable_disable_compliance_checks.disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/)
- [enable_disable_compliance_checks.enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
