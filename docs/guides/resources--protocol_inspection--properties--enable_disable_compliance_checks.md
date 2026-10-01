---
page_title: "enable_disable_compliance_checks"
subcategory: ""
description: "enable_disable_compliance_checks for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 1891, "body_sha256": "sha256:6aa65b79c64657b25e581f85919a3b00ecfb76212d63824651df5f8efafe970c", "canonical_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "child_ids": ["xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks"], "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "parent_id": "xcsh-docs:resources:protocol_inspection:reference", "path": "docs/guides/resources--protocol_inspection--properties--enable_disable_compliance_checks.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_disable_compliance_checks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_compliance_checks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_disable_compliance_checks for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_compliance_checks

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md)
- [Property reference](resources--protocol_inspection--reference.md)
- enable_disable_compliance_checks

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Enable Disable Compliance Checks Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_compliance_checks",
    "enable_compliance_checks")}
```

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

Terraform syntax:

```terraform
enable_disable_compliance_checks {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_compliance_checks](resources--protocol_inspection--properties--enable_disable_compliance_checks--disable_compliance_checks.md): complete subsection reference.

- [enable_compliance_checks](resources--protocol_inspection--properties--enable_disable_compliance_checks--enable_compliance_checks.md): complete subsection reference.

## Next pages

- [enable_disable_compliance_checks.disable_compliance_checks](resources--protocol_inspection--properties--enable_disable_compliance_checks--disable_compliance_checks.md)
- [enable_disable_compliance_checks.enable_compliance_checks](resources--protocol_inspection--properties--enable_disable_compliance_checks--enable_compliance_checks.md)
- [Property reference](resources--protocol_inspection--reference.md)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md)
