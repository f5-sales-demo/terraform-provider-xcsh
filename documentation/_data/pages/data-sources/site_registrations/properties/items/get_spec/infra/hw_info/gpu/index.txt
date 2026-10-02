---
page_title: "items.get_spec.infra.hw_info.gpu"
subcategory: ""
description: "GPU. GPU information on server."
xcsh_docs: {"aliases": ["items get spec infra hw info gpu"], "body_bytes": 2107, "body_sha256": "sha256:e787f2295f9b63eb0b22b0ecf0d01a9960aa2432ab2dc4a48b92ed7b10f0b4c1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu:gpu_device"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3232310111021330-2211123101011101-2323001102002223-1013203311032030-2313131310312103-3233032012310232-0200100133213021-2123111230113020", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu"], "schema_version": 1, "sections": [{"aliases": ["cuda version"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--cuda_version", "description": "Cuda Version. GPU Cuda Version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "cuda_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["driver version"], "anchor": "schema-items--get_spec--infra--hw_info--gpu--driver_version", "description": "Driver Version. GPU Driver Version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "driver_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["gpu device"], "anchor": "section", "description": "GPU devices. List of GPU devices in server.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:gpu:gpu_device", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "gpu", "gpu_device"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "GPU. GPU information on server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.gpu

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.gpu

<a id="section"></a>

Type: `"single"`. Computed.

GPU. GPU information on server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--gpu--cuda_version"></a>

### cuda_version property

Type: `"string"`. Computed.

Cuda Version. GPU Cuda Version.

<a id="schema-items--get_spec--infra--hw_info--gpu--driver_version"></a>

### driver_version property

Type: `"string"`. Computed.

Driver Version. GPU Driver Version.

- [gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/gpu_device/): complete subsection reference.

## Next pages

- [items.get_spec.infra.hw_info.gpu.gpu_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/gpu/gpu_device/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
