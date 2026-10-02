---
page_title: "items.object.spec.gc_spec.infra.hw_info.bios"
subcategory: ""
description: "Bios Data. BIOS information."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info bios"], "body_bytes": 2536, "body_sha256": "sha256:e7372f7194ef14d7e0fce49a20e921bd921d5898acb462fe0cfbc9fb4d161456", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:bios", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/bios/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1100223311131212-2332112211223212-3110022210000310-1232123202300033-3222321322111332-1103333113131100-1313311321000221-3200011230231200", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "bios"], "schema_version": 1, "sections": [{"aliases": ["date"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--bios--date", "description": "Information from /sys/class/dmi/ID/bios_date.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "bios", "date"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--bios--vendor", "description": "Information from /sys/class/dmi/ID/bios_vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "bios", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--bios--version", "description": "Information from /sys/class/dmi/ID/bios_version.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:bios", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "bios", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bios Data. BIOS information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.bios

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.bios

<a id="section"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--date"></a>

### date property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--bios--version"></a>

### version property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
