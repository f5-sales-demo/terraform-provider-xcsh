---
page_title: "infra.hw_info"
subcategory: ""
description: "OsInfo holds information about host OS and HW."
xcsh_docs: {"aliases": ["infra hw info"], "body_bytes": 3410, "body_sha256": "sha256:c233a136556a415600b506745049f12862e13a491b6819dfbd7c7f867ebc9103", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra:hw_info:bios", "xcsh-docs:data-sources:registration:properties:infra:hw_info:board", "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "xcsh-docs:data-sources:registration:properties:infra:hw_info:product", "xcsh-docs:data-sources:registration:properties:infra:hw_info:storage", "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "documentation/data-sources/registration/properties/infra/hw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "sections": [{"aliases": ["infra hw info bios"], "anchor": "section", "description": "BIOS information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "bios"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info board"], "anchor": "section", "description": "Board information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "board"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info chassis"], "anchor": "section", "description": "Chassis information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "chassis"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info cpu"], "anchor": "section", "description": "CPU information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "cpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info gpu"], "anchor": "section", "description": "GPU information on server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:gpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "gpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info kernel"], "anchor": "section", "description": "Kernel information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:kernel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "kernel"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info memory"], "anchor": "section", "description": "Memory information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:memory", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "memory"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info network"], "anchor": "section", "description": "List of network devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "network"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info numa nodes"], "anchor": "schema-infra--hw_info--numa_nodes", "description": "Non-uniform memory access (NUMA) nodes count.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info os"], "anchor": "section", "description": "Details of Operating System.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "os"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info product"], "anchor": "section", "description": "Product information.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:product", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "product"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info storage"], "anchor": "section", "description": "List of storage devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["infra hw info usb"], "anchor": "section", "description": "List of USB devices in server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "usb"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "OsInfo holds information about host OS and HW.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
