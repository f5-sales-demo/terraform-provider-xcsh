---
page_title: "items.object.spec.gc_spec.infra.hw_info.cpu"
subcategory: ""
description: "CPU Information. CPU information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info cpu"], "body_bytes": 2882, "body_sha256": "sha256:e90bfd1945ea5e546dbdbe6c341f45ad1478d20dfa5fa58ed819db2df18e2c3d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/cpu/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2023203003032122-2330112103112102-3230201121330003-1231102133223113-0130321101323110-2133131200101022-1103331221130122-3021013103033232", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu"], "schema_version": 1, "sections": [{"aliases": ["cache"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--cache", "description": "Cache. CPU cache size in KB.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "cache"], "syntax": "attribute", "type": "number"}, {"aliases": ["cores"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--cores", "description": "Cores. Number of physical CPU cores.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "cores"], "syntax": "attribute", "type": "number"}, {"aliases": ["cpus"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--cpus", "description": "CPUs. Number of physical CPUs.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "cpus"], "syntax": "attribute", "type": "number"}, {"aliases": ["model"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--model", "description": "Model. CPU model", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--speed", "description": "Speed. CPU clock rate in MHz.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["threads"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--threads", "description": "Threads. Number of logical (HT) CPU cores.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "threads"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--cpu--vendor", "description": "Vendor. CPU vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "cpu", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/cpu/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CPU Information. CPU information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.cpu

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.cpu

<a id="section"></a>

Type: `"single"`. Computed.

CPU Information. CPU information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--cache"></a>

### cache property

Type: `"number"`. Computed.

Cache. CPU cache size in KB.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--cores"></a>

### cores property

Type: `"number"`. Computed.

Cores. Number of physical CPU cores.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--cpus"></a>

### cpus property

Type: `"number"`. Computed.

CPUs. Number of physical CPUs.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--model"></a>

### model property

Type: `"string"`. Computed.

Model. CPU model

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. CPU clock rate in MHz.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--threads"></a>

### threads property

Type: `"number"`. Computed.

Threads. Number of logical (HT) CPU cores.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--cpu--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. CPU vendor.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
