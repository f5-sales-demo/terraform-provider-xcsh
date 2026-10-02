---
page_title: "enable_disable_compliance_checks"
subcategory: ""
description: "Enable Disable Compliance Checks Choice."
xcsh_docs: {"aliases": ["enable disable compliance checks"], "body_bytes": 1983, "body_sha256": "sha256:2bed90acfe9f5bd9ffaaa3fb75f4b1364c06e140488a8f0ef0913e00bdf110fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks", "parent_id": "xcsh-docs:data-sources:protocol_inspection:reference", "path": "documentation/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1223020321120303-0303012222310120-1230131211300021-2121032111222012-3110102121333010-2132331121321033-0210312231020030-0323222302001133", "registry_path": "docs/guides/data-sources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_compliance_checks"], "schema_version": 1, "sections": [{"aliases": ["disable compliance checks"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_disable_compliance_checks", "disable_compliance_checks"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable compliance checks"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks:enable_compliance_checks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_disable_compliance_checks", "enable_compliance_checks"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Enable Disable Compliance Checks Choice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
