---
page_title: "enable_disable_compliance_checks"
subcategory: ""
description: "Enable Disable Compliance Checks Choice."
xcsh_docs: {"aliases": ["enable disable compliance checks"], "body_bytes": 2299, "body_sha256": "sha256:b1fa97a0fb0257e8a3228902f1be8cd345f4b19aafda56190d3ec3254f7d1e3f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "parent_id": "xcsh-docs:resources:protocol_inspection:reference", "path": "documentation/resources/protocol_inspection/properties/enable_disable_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3032303120120322-2021100000002111-2310220332010231-1233101110001211-1130323103302010-2213213331133112-1031321210203310-1002133331133202", "registry_path": "docs/guides/resources--protocol_inspection--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_compliance_checks:ConflictingObjectAttributes:disable_compliance_checks,enable_compliance_checks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_disable_compliance_checks:ConflictingObjectAttributes:disable_compliance_checks,enable_compliance_checks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_compliance_checks"], "schema_version": 1, "sections": [{"aliases": ["disable compliance checks"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_compliance_checks", "disable_compliance_checks"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable compliance checks"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_disable_compliance_checks--enable_compliance_checks--name", "enforcement": "provider-schema", "group": "enable_disable_compliance_checks.enable_compliance_checks:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks", "type": "requires"}], "schema_path": ["enable_disable_compliance_checks", "enable_compliance_checks"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_compliance_checks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Enable Disable Compliance Checks Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_compliance_checks

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
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

- [disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/): complete subsection reference.

- [enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/): complete subsection reference.

## Next pages

- [enable_disable_compliance_checks.disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/)
- [enable_disable_compliance_checks.enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
