---
page_title: "infra.hw_info"
subcategory: ""
description: "OsInfo holds information about host OS and HW."
xcsh_docs: {"aliases": ["infra hw info"], "body_bytes": 3410, "body_sha256": "sha256:2a68ed8db7fd963047d4a9c8353e9dfb45b8fbb5baa1bbce169a71ee4e955feb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra:hw_info:bios", "xcsh-docs:data-sources:registration:properties:infra:hw_info:board", "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "xcsh-docs:data-sources:registration:properties:infra:hw_info:product", "xcsh-docs:data-sources:registration:properties:infra:hw_info:storage", "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "documentation/data-sources/registration/properties/infra/hw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "sections": [{"aliases": ["infra hw info bios"], "anchor": "section", "description": "BIOS information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:bios", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "bios"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info board"], "anchor": "section", "description": "Board information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:board", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "board"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info chassis"], "anchor": "section", "description": "Chassis information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "chassis"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info cpu"], "anchor": "section", "description": "CPU information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "cpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info gpu"], "anchor": "section", "description": "GPU information on server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "gpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info kernel"], "anchor": "section", "description": "Kernel information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "kernel"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info memory"], "anchor": "section", "description": "Memory information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "memory"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info network"], "anchor": "section", "description": "List of network devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "network"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info numa nodes"], "anchor": "schema-infra--hw_info--numa_nodes", "description": "Non-uniform memory access (NUMA) nodes count.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info os"], "anchor": "section", "description": "Details of Operating System.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "os"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info product"], "anchor": "section", "description": "Product information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:product", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "product"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info storage"], "anchor": "section", "description": "List of storage devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info usb"], "anchor": "section", "description": "List of USB devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "usb"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "OsInfo holds information about host OS and HW.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["registrationCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- infra.hw_info

<a id="section"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

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

## Direct properties

- [bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/bios/): complete subsection reference.

- [board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/board/): complete subsection reference.

- [chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/chassis/): complete subsection reference.

- [cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/cpu/): complete subsection reference.

- [gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/gpu/): complete subsection reference.

- [kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/kernel/): complete subsection reference.

- [memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/memory/): complete subsection reference.

- [network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/network/): complete subsection reference.

<a id="schema-infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  }
}
```

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/os/): complete subsection reference.

- [product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/product/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/storage/): complete subsection reference.

- [usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/usb/): complete subsection reference.
