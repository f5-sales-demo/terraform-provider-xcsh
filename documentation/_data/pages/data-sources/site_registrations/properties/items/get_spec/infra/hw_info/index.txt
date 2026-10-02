---
page_title: "items.get_spec.infra.hw_info"
subcategory: ""
description: "OsInfo holds information about host OS and HW."
xcsh_docs: {"aliases": ["items get spec infra hw info"], "body_bytes": 5762, "body_sha256": "sha256:5225fe6f22be43a3de902cf257a951df3154bf3cd2c113fa6cfd3fd4e3836a9c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:chassis", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:kernel", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:memory", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:network", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:product", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:storage", "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:usb"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2011322120111013-0031321000222033-0200101230110001-3200033020300003-3323212211131002-2210312323000123-0313032203322200-3000010122123130", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info"], "schema_version": 1, "sections": [{"aliases": ["bios"], "anchor": "section", "description": "Bios Data. BIOS information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "bios"], "syntax": "attribute", "type": "object"}, {"aliases": ["board"], "anchor": "section", "description": "Board Details. Board information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:board", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "board"], "syntax": "attribute", "type": "object"}, {"aliases": ["chassis"], "anchor": "section", "description": "Chassis Details. Chassis information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "chassis"], "syntax": "attribute", "type": "object"}, {"aliases": ["cpu"], "anchor": "section", "description": "CPU Information. CPU information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["gpu"], "anchor": "section", "description": "GPU. GPU information on server.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu"], "syntax": "attribute", "type": "object"}, {"aliases": ["kernel"], "anchor": "section", "description": "Kernel. Kernel information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:kernel", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "kernel"], "syntax": "attribute", "type": "object"}, {"aliases": ["memory"], "anchor": "section", "description": "Memory Information. Memory information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:memory", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "memory"], "syntax": "attribute", "type": "object"}, {"aliases": ["network"], "anchor": "section", "description": "Network. List of network devices in server.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network"], "syntax": "attribute", "type": "object"}, {"aliases": ["numa nodes"], "anchor": "schema-items--get_spec--infra--hw_info--numa_nodes", "description": "Non-uniform memory access (NUMA) nodes count.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["os"], "anchor": "section", "description": "OS. Details of Operating System.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "os"], "syntax": "attribute", "type": "object"}, {"aliases": ["product"], "anchor": "section", "description": "Product Information. Product information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:product", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "product"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage"], "anchor": "section", "description": "Storage. List of storage devices in server.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["usb"], "anchor": "section", "description": "USB devices. List of USB devices in server.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "usb"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OsInfo holds information about host OS and HW.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- items.get_spec.infra.hw_info

<a id="section"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

## Direct properties

- [bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/bios/): complete subsection reference.

- [board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/board/): complete subsection reference.

- [chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/chassis/): complete subsection reference.

- [cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/cpu/): complete subsection reference.

- [gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/): complete subsection reference.

- [kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/kernel/): complete subsection reference.

- [memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/memory/): complete subsection reference.

- [network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/network/): complete subsection reference.

<a id="schema-items--get_spec--infra--hw_info--numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/os/): complete subsection reference.

- [product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/product/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/storage/): complete subsection reference.

- [usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/usb/): complete subsection reference.

## Next pages

- [items.get_spec.infra.hw_info.bios](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/bios/)
- [items.get_spec.infra.hw_info.board](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/board/)
- [items.get_spec.infra.hw_info.chassis](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/chassis/)
- [items.get_spec.infra.hw_info.cpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/cpu/)
- [items.get_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/)
- [items.get_spec.infra.hw_info.kernel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/kernel/)
- [items.get_spec.infra.hw_info.memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/memory/)
- [items.get_spec.infra.hw_info.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/network/)
- [items.get_spec.infra.hw_info.os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/os/)
- [items.get_spec.infra.hw_info.product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/product/)
- [items.get_spec.infra.hw_info.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/storage/)
- [items.get_spec.infra.hw_info.usb](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/usb/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
