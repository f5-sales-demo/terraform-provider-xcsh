---
page_title: "items.object.spec.gc_spec.infra.hw_info.storage"
subcategory: ""
description: "Storage. List of storage devices in server."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info storage"], "body_bytes": 2997, "body_sha256": "sha256:ada739cf375050508870f4f490cb8511ce15611cd2e4c654b890559b88719cc6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/storage/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2121231301220010-2202332120313122-2201123200113300-1201210032120001-2302021002123232-1113223110230301-1310030332301223-1001331232023120", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage"], "schema_version": 1, "sections": [{"aliases": ["driver"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--driver", "description": "Driver. Driver of device.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["model"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--model", "description": "Model. Model of device.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--name", "description": "Name. Name of device, eg. Nvme0n1.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["serial"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--serial", "description": "Serial Number. Serial of device.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["size gb"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--size_gb", "description": "Size(GB). Device size in GB.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "size_gb"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--storage--vendor", "description": "Vendor. Vendor of device.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "storage", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/storage/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Storage. List of storage devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.storage

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.storage

<a id="section"></a>

Type: `"list"`. Computed.

Storage. List of storage devices in server.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--model"></a>

### model property

Type: `"string"`. Computed.

Model. Model of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of device, eg. Nvme0n1.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--serial"></a>

### serial property

Type: `"string"`. Computed.

Serial Number. Serial of device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--size_gb"></a>

### size_gb property

Type: `"number"`. Computed.

Size(GB). Device size in GB.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--storage--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of device.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
