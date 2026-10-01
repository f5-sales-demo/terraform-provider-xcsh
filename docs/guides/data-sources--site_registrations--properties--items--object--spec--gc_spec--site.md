---
page_title: "items.object.spec.gc_spec.site"
subcategory: ""
description: "items.object.spec.gc_spec.site for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 2515, "body_sha256": "sha256:c838932a86045a3181885002b1d466f20f7875cb74f0bfcedd0a9100a48e11a2", "canonical_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:site", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:site", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "path": "docs/guides/data-sources--site_registrations--properties--items--object--spec--gc_spec--site.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.site for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.site

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md)
- [Property reference](data-sources--site_registrations--reference.md)
- [items](data-sources--site_registrations--properties--items.md)
- [items.object](data-sources--site_registrations--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations--properties--items--object--spec--gc_spec.md)
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

- [items.object.spec.gc_spec](data-sources--site_registrations--properties--items--object--spec--gc_spec.md)
- [xcsh_site_registrations](../data-sources/site_registrations.md)
