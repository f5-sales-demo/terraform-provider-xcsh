---
page_title: "items.get_spec.infra.hw_info.gpu.gpu_device"
subcategory: ""
description: "GPU devices. List of GPU devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info gpu gpu device"], "body_bytes": 1831, "body_sha256": "sha256:b56c3eb27cd84a81b66789890d224e3b1a3568e9f81c27bf2aaf687fdae3e8ca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:gpu", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/gpu/gpu_device/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3130232231111021-0111332031201120-3023322011203011-1312123012232112-3012203221313131-0320131131101311-3011320001200002-2310120222120301", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info gpu gpu device id"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--id", "description": "GPU ID. GPU ID", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info gpu gpu device processes"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--processes", "description": "Processes. GPU Processes.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "processes"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info gpu gpu device product name"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--gpu_device--product_name", "description": "Product Name. GPU Product Name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device", "product_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/gpu/gpu_device/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "GPU devices. List of GPU devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.gpu.gpu_device

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- [items.get_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/gpu/)
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
