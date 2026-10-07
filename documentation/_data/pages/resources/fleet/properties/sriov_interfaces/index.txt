---
page_title: "sriov_interfaces"
subcategory: ""
description: "List of all custom SR-IOV interfaces configuration."
xcsh_docs: {"aliases": ["sriov interfaces"], "body_bytes": 946, "body_sha256": "sha256:7ddd17d65ddb54775f0d4fb7ce2a866918dd5e8e9a5fcdd2eaa3f1571eaed650", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:sriov_interfaces", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/sriov_interfaces/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces"], "schema_version": 1, "sections": [{"aliases": ["sriov interfaces sriov interface"], "anchor": "section", "description": "Use custom SR-IOV interfaces Configuration.", "document_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "type": "requires"}, {"anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:sriov_interfaces:sriov_interface", "type": "requires"}], "schema_path": ["sriov_interfaces", "sriov_interface"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/sriov_interfaces/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "List of all custom SR-IOV interfaces configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- sriov_interfaces

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/sriov_interfaces/sriov_interface/): complete subsection reference.
