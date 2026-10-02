---
page_title: "items.object.spec.gc_spec.infra.hugepages"
subcategory: ""
description: "Hugepage settings for CE on K8s SMV2 site."
xcsh_docs: {"aliases": ["items object spec gc spec infra hugepages"], "body_bytes": 2192, "body_sha256": "sha256:8a1a5813c024c6b4559a539b832dd81bf6c10b75d284b13b1439ac022fa768aa", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hugepages/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1130200130211011-1302113220332232-1213023013002122-1112231332000032-3203010111312211-2233100222102103-3123210112131033-0220231211201223", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages"], "schema_version": 1, "sections": [{"aliases": ["free"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--free", "description": "Free Hugepages. Total number of free hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "free"], "syntax": "attribute", "type": "number"}, {"aliases": ["page size"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--page_size", "description": "Hugepage Size. Size of each hugepage.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "page_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["total"], "anchor": "schema-items--object--spec--gc_spec--infra--hugepages--total", "description": "Total Hugepages. Total number of hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hugepages", "total"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/hugepages/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Hugepage settings for CE on K8s SMV2 site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hugepages

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
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

- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
