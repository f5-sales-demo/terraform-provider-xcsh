---
page_title: "items.get_spec.infra.hugepages"
subcategory: ""
description: "Hugepage settings for CE on K8s SMV2 site."
xcsh_docs: {"aliases": ["items get spec infra hugepages"], "body_bytes": 1777, "body_sha256": "sha256:d20ed29c512893d9bd0136bde12051d8e76b571fdf2b4e2ea3d40ad936d2c872", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hugepages", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hugepages/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3211310302111100-1113113022123333-0113032133311102-1003103020000300-2232013021033120-0101012133110303-0133031021301031-1330101211110000", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hugepages"], "schema_version": 1, "sections": [{"aliases": ["free"], "anchor": "schema-items--get_spec--infra--hugepages--free", "description": "Free Hugepages. Total number of free hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hugepages", "free"], "syntax": "attribute", "type": "number"}, {"aliases": ["page size"], "anchor": "schema-items--get_spec--infra--hugepages--page_size", "description": "Hugepage Size. Size of each hugepage.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hugepages", "page_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["total"], "anchor": "schema-items--get_spec--infra--hugepages--total", "description": "Total Hugepages. Total number of hugepages present.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hugepages", "total"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hugepages/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Hugepage settings for CE on K8s SMV2 site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hugepages

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- items.get_spec.infra.hugepages

<a id="section"></a>

Type: `"list"`. Computed.

Hugepage settings for CE on K8s SMV2 site.

## Direct properties

<a id="schema-items--get_spec--infra--hugepages--free"></a>

### free property

Type: `"number"`. Computed.

Free Hugepages. Total number of free hugepages present.

<a id="schema-items--get_spec--infra--hugepages--page_size"></a>

### page_size property

Type: `"number"`. Computed.

Hugepage Size. Size of each hugepage.

<a id="schema-items--get_spec--infra--hugepages--total"></a>

### total property

Type: `"number"`. Computed.

Total Hugepages. Total number of hugepages present.

## Next pages

- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
