---
page_title: "items.object.spec.gc_spec.infra.hw_info.memory"
subcategory: ""
description: "Memory Information. Memory information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info memory"], "body_bytes": 2361, "body_sha256": "sha256:1016f6aa0ef35cb32e5fb1bb1185d7bb60aaa87abe974d64702160b5c84ea2c2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:memory", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/memory/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1302122202012200-0101011202123231-2100210031212103-3110013323232220-2002111000113010-1321031031123333-3231010102230120-3312022021133101", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra hw info memory size mb"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--size_mb", "description": "RAM. RAM size in MB.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:memory", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "size_mb"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object spec gc spec infra hw info memory speed"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--speed", "description": "Speed. RAM data rate in MT/s.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:memory", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object spec gc spec infra hw info memory type"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--memory--type", "description": "Type. Type of memory, eg. DDR4.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hw_info:memory", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "memory", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/memory/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Memory Information. Memory information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.memory

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.memory

<a id="section"></a>

Type: `"single"`. Computed.

Memory Information. Memory information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--size_mb"></a>

### size_mb property

Type: `"number"`. Computed.

RAM. RAM size in MB.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. RAM data rate in MT/s.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--memory--type"></a>

### type property

Type: `"string"`. Computed.

Type. Type of memory, eg. DDR4.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
