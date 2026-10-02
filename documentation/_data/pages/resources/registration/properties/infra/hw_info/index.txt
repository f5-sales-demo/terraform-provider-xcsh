---
page_title: "infra.hw_info"
subcategory: ""
description: "OsInfo holds information about host OS and HW."
xcsh_docs: {"aliases": ["infra hw info"], "body_bytes": 5469, "body_sha256": "sha256:e9ac093b200d16f97d5374512d2b52d589932a39d4b90a72620b5ff209f4c0a3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:infra:hw_info:bios", "xcsh-docs:resources:registration:properties:infra:hw_info:board", "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "xcsh-docs:resources:registration:properties:infra:hw_info:cpu", "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "xcsh-docs:resources:registration:properties:infra:hw_info:network", "xcsh-docs:resources:registration:properties:infra:hw_info:os", "xcsh-docs:resources:registration:properties:infra:hw_info:product", "xcsh-docs:resources:registration:properties:infra:hw_info:storage", "xcsh-docs:resources:registration:properties:infra:hw_info:usb"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/hw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info"], "schema_version": 1, "sections": [{"aliases": ["bios"], "anchor": "section", "description": "BIOS information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "bios"], "syntax": "block", "type": "object"}, {"aliases": ["board"], "anchor": "section", "description": "Board information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:board", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "board"], "syntax": "block", "type": "object"}, {"aliases": ["chassis"], "anchor": "section", "description": "Chassis information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "chassis"], "syntax": "block", "type": "object"}, {"aliases": ["cpu"], "anchor": "section", "description": "CPU information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:cpu", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "cpu"], "syntax": "block", "type": "object"}, {"aliases": ["gpu"], "anchor": "section", "description": "GPU information on server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:gpu", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "gpu"], "syntax": "block", "type": "object"}, {"aliases": ["kernel"], "anchor": "section", "description": "Kernel information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:kernel", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "kernel"], "syntax": "block", "type": "object"}, {"aliases": ["memory"], "anchor": "section", "description": "Memory information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:memory", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "memory"], "syntax": "block", "type": "object"}, {"aliases": ["network"], "anchor": "section", "description": "List of network devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:network", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "network"], "syntax": "block", "type": "object"}, {"aliases": ["numa nodes"], "anchor": "schema-infra--hw_info--numa_nodes", "description": "Non-uniform memory access (NUMA) nodes count.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["os"], "anchor": "section", "description": "Details of Operating System.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:os", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "os"], "syntax": "block", "type": "object"}, {"aliases": ["product"], "anchor": "section", "description": "Product information.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:product", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info", "product"], "syntax": "block", "type": "object"}, {"aliases": ["storage"], "anchor": "section", "description": "List of storage devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "storage"], "syntax": "block", "type": "object"}, {"aliases": ["usb"], "anchor": "section", "description": "List of USB devices in server.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hw_info", "usb"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OsInfo holds information about host OS and HW.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
