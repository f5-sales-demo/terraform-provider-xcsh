---
page_title: "items.object.spec.gc_spec.infra.hugepages"
subcategory: ""
description: "Hugepage settings for CE on K8s SMV2 site."
xcsh_docs: {"aliases": ["items object spec gc spec infra hugepages"], "body_bytes": 2203, "body_sha256": "sha256:5f8dc087ddc2b9486b1cc806e43b7ef42186a9955024edba198de0a278e1202d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hugepages/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1010130201023122-0331201310011320-0200321032112302-2331132201131103-3100322223211002-1010233200132100-2330210312110131-3222123030231103", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages"], "schema_version": 1, "sections": [{"aliases": ["free"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--free", "description": "Free Hugepages. Total number of free hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "free"], "syntax": "attribute", "type": "number"}, {"aliases": ["page size"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--page_size", "description": "Hugepage Size. Size of each hugepage.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "page_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["total"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--total", "description": "Total Hugepages. Total number of hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "total"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hugepages/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Hugepage settings for CE on K8s SMV2 site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hugepages

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/)
- items.object.spec.gc_spec.infra.hugepages

<a id="section"></a>

Type: `"list"`. Computed.

Hugepage settings for CE on K8s SMV2 site.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hugepages--free"></a>

### free property

Type: `"number"`. Computed.

Free Hugepages. Total number of free hugepages present.

<a id="schema-items--object--spec--gc_spec--infra--hugepages--page_size"></a>

### page_size property

Type: `"number"`. Computed.

Hugepage Size. Size of each hugepage.

<a id="schema-items--object--spec--gc_spec--infra--hugepages--total"></a>

### total property

Type: `"number"`. Computed.

Total Hugepages. Total number of hugepages present.

## Next pages

- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
