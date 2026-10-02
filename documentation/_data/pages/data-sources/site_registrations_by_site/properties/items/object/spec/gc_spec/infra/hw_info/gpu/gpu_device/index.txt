---
page_title: "items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device"
subcategory: ""
description: "GPU devices. List of GPU devices in server."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info gpu gpu device"], "body_bytes": 2619, "body_sha256": "sha256:718f7dbfc00c0fa2de9511ed02acaea103aead0d93f6aafe4d57b0f3f35dbccb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu:gpu_device", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/gpu/gpu_device/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3223011200211202-0022130101123332-3120302123201110-1131310121301112-3321303001331132-2301203000231001-2121130020113001-3133101222130010", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "gpu", "gpu_device"], "schema_version": 1, "sections": [{"aliases": ["id"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--id", "description": "GPU ID. GPU ID", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu:gpu_device", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "gpu", "gpu_device", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["processes"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--processes", "description": "Processes. GPU Processes.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu:gpu_device", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "gpu", "gpu_device", "processes"], "syntax": "attribute", "type": "string"}, {"aliases": ["product name"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--product_name", "description": "Product Name. GPU Product Name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:gpu:gpu_device", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "gpu", "gpu_device", "product_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/gpu/gpu_device/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "GPU devices. List of GPU devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- [items.object.spec.gc_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/gpu/)
- items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device

<a id="section"></a>

Type: `"list"`. Computed.

GPU devices. List of GPU devices in server.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--id"></a>

### id property

Type: `"string"`. Computed.

GPU ID. GPU ID

<a id="schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--processes"></a>

### processes property

Type: `"string"`. Computed.

Processes. GPU Processes.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--gpu--gpu_device--product_name"></a>

### product_name property

Type: `"string"`. Computed.

Product Name. GPU Product Name.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info.gpu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/gpu/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
