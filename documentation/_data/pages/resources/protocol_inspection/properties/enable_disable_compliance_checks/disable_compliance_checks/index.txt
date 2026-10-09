---
page_title: "enable_disable_compliance_checks.disable_compliance_checks"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable disable compliance checks disable compliance checks"], "body_bytes": 1161, "body_sha256": "sha256:4d0c0f104ed3e7c22ed7b24e3f76e92ad9ca6326d77bad809b46054b421c75a5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks:disable_compliance_checks", "parent_id": "xcsh-docs:resources:protocol_inspection:properties:enable_disable_compliance_checks", "path": "documentation/resources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2002213210222313-1301203032330232-1130002101020131-1031220220200202-0211001113021203-2132333320232031-1300203131011000-1120103321023000", "registry_path": "docs/guides/resources--protocol_inspection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_disable_compliance_checks", "disable_compliance_checks"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_disable_compliance_checks.disable_compliance_checks

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/)
- [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/properties/enable_disable_compliance_checks/)
- enable_disable_compliance_checks.disable_compliance_checks

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable compliance checks.

Additional upstream details:

This can be used for messages where no values are needed.

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
disable_compliance_checks = {}
```

This is an empty object or choice marker. It has no direct properties.
