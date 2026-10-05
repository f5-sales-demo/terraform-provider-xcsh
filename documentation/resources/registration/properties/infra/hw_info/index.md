---
page_title: "infra.hw_info"
subcategory: ""
description: "OsInfo holds information about host OS and HW."
xcsh_docs: {"aliases": ["infra hw info"], "body_bytes": 5469, "body_sha256": "sha256:c5453785b523ab3d91f6d49a9db237e6edd8d7aae0911b6b22c66bc624e94079", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:infra:hw_info:bios", "xcsh-docs:resources:registration:properties:infra:hw_info:board", "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "xcsh-docs:resources:registration:properties:infra:hw_info:cpu", "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "xcsh-docs:resources:registration:properties:infra:hw_info:network", "xcsh-docs:resources:registration:properties:infra:hw_info:os", "xcsh-docs:resources:registration:properties:infra:hw_info:product", "xcsh-docs:resources:registration:properties:infra:hw_info:storage", "xcsh-docs:resources:registration:properties:infra:hw_info:usb"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/hw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "sections": [{"aliases": ["infra hw info bios"], "anchor": "section", "description": "BIOS information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "bios"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info board"], "anchor": "section", "description": "Board information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:board", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "board"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info chassis"], "anchor": "section", "description": "Chassis information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "chassis"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info cpu"], "anchor": "section", "description": "CPU information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:cpu", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "cpu"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info gpu"], "anchor": "section", "description": "GPU information on server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "gpu"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info kernel"], "anchor": "section", "description": "Kernel information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "kernel"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info memory"], "anchor": "section", "description": "Memory information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "memory"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info network"], "anchor": "section", "description": "List of network devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:network", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "network"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info numa nodes"], "anchor": "schema-infra--hw_info--numa_nodes", "description": "Non-uniform memory access (NUMA) nodes count.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info os"], "anchor": "section", "description": "Details of Operating System.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:os", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "os"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info product"], "anchor": "section", "description": "Product information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:product", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "product"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info storage"], "anchor": "section", "description": "List of storage devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "storage"], "syntax": "block", "type": "object"}, {"aliases": ["infra hw info usb"], "anchor": "section", "description": "List of USB devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "usb"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "OsInfo holds information about host OS and HW.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- infra.hw_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
hw_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/bios/): complete subsection reference.

- [board](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/board/): complete subsection reference.

- [chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/chassis/): complete subsection reference.

- [cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/cpu/): complete subsection reference.

- [gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/): complete subsection reference.

- [kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/kernel/): complete subsection reference.

- [memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/memory/): complete subsection reference.

- [network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/network/): complete subsection reference.

<a id="schema-infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Optional.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

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

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/os/): complete subsection reference.

- [product](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/product/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/storage/): complete subsection reference.

- [usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/usb/): complete subsection reference.

## Next pages

- [infra.hw_info.bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/bios/)
- [infra.hw_info.board](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/board/)
- [infra.hw_info.chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/chassis/)
- [infra.hw_info.cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/cpu/)
- [infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/gpu/)
- [infra.hw_info.kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/kernel/)
- [infra.hw_info.memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/memory/)
- [infra.hw_info.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/network/)
- [infra.hw_info.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/os/)
- [infra.hw_info.product](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/product/)
- [infra.hw_info.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/storage/)
- [infra.hw_info.usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/usb/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
