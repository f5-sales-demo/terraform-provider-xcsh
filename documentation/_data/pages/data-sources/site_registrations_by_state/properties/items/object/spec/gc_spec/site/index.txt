---
page_title: "items.object.spec.gc_spec.site"
subcategory: ""
description: "Site for this registration, assigned after registration is assigned to site."
xcsh_docs: {"aliases": ["items object spec gc spec site"], "body_bytes": 3006, "body_sha256": "sha256:1bc453344d4ba4d736ef918e6fba36e4a9310a68552408cc8fcd706f773b698e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/site/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0221221101033201-1020331112200021-1120033002221013-1033203123012222-1023313332113323-0023330303303010-3030120231213132-2233211312001103", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "site"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec site kind"], "anchor": "schema-items--object--spec--gc_spec--site--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec site name"], "anchor": "schema-items--object--spec--gc_spec--site--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec site namespace"], "anchor": "schema-items--object--spec--gc_spec--site--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec site tenant"], "anchor": "schema-items--object--spec--gc_spec--site--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec site uid"], "anchor": "schema-items--object--spec--gc_spec--site--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Site for this registration, assigned after registration is assigned to site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.site

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
- items.object.spec.gc_spec.site

<a id="section"></a>

Type: `"list"`. Computed.

Site for this registration, assigned after registration is assigned to site.

## Direct properties

<a id="schema-items--object--spec--gc_spec--site--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-items--object--spec--gc_spec--site--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-items--object--spec--gc_spec--site--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--object--spec--gc_spec--site--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-items--object--spec--gc_spec--site--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
