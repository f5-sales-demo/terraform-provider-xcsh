---
page_title: "items.owner_view"
subcategory: ""
description: "items.owner_view for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 1823, "body_sha256": "sha256:31264f113edce7174e73580a198fb893838751ae1e2ac45c2608e726101660a7", "canonical_id": "xcsh-docs:data-sources:site_registrations:properties:items:owner_view", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:owner_view", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "docs/guides/data-sources--site_registrations--properties--items--owner_view.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "owner_view"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/owner_view/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.owner_view for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.owner_view

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md)
- [Property reference](data-sources--site_registrations--reference.md)
- [items](data-sources--site_registrations--properties--items.md)
- items.owner_view

<a id="section"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

## Direct properties

<a id="schema-items--owner_view--kind"></a>

### kind property

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="schema-items--owner_view--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--owner_view--namespace"></a>

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

<a id="schema-items--owner_view--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. UID of the view object.

## Next pages

- [items](data-sources--site_registrations--properties--items.md)
- [xcsh_site_registrations](../data-sources/site_registrations.md)
