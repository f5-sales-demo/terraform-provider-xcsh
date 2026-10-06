---
page_title: "items.get_spec.infra.hw_info.gpu.gpu_device"
subcategory: ""
description: "GPU devices. List of GPU devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info gpu gpu device"], "body_bytes": 1839, "body_sha256": "sha256:7236aa1a4ce40ab51e77d578dd8fc90a0effb0c771fb273ba205e48866129a3c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/gpu/gpu_device/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1233312223031333-3311132012133130-0330213321211120-2000200020120321-1021323323213202-1221312320320032-0030032103112321-1311032112001133", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info gpu gpu device id"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--id", "description": "GPU ID. GPU ID", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info gpu gpu device processes"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--processes", "description": "Processes. GPU Processes.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "processes"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info gpu gpu device product name"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--product_name", "description": "Product Name. GPU Product Name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "product_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/gpu/gpu_device/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "GPU devices. List of GPU devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.gpu.gpu_device

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- [items.get_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/gpu/)
- items.get_spec.infra.hw_info.gpu.gpu_device

<a id="section"></a>

Type: `"list"`. Computed.

GPU devices. List of GPU devices in server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--gpu--gpu_device--id"></a>

### id property

Type: `"string"`. Computed.

GPU ID. GPU ID

<a id="schema-items--get_spec--infra--hw_info--gpu--gpu_device--processes"></a>

### processes property

Type: `"string"`. Computed.

Processes. GPU Processes.

<a id="schema-items--get_spec--infra--hw_info--gpu--gpu_device--product_name"></a>

### product_name property

Type: `"string"`. Computed.

Product Name. GPU Product Name.
