---
page_title: "items"
subcategory: ""
description: "items for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 3387, "body_sha256": "sha256:94e4053aef95f3e8b677376d3dbca5dc9d0f97807cc520cc8b298e8d1cbe8b18", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:annotations", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:labels", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:owner_view", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:system_metadata"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:reference", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Items represents the collection in response.

## Direct properties

- [annotations](data-sources--site_registrations_by_site--properties--items--annotations.md): complete subsection reference.

<a id="schema-items--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

The description set for this registration.

<a id="schema-items--disabled"></a>

### disabled property

Type: `"bool"`. Computed.

Value of true indicates registration is administratively disabled.

- [get_spec](data-sources--site_registrations_by_site--properties--items--get_spec.md): complete subsection reference.

- [labels](data-sources--site_registrations_by_site--properties--items--labels.md): complete subsection reference.

- [metadata](data-sources--site_registrations_by_site--properties--items--metadata.md): complete subsection reference.

<a id="schema-items--name"></a>

### name property

Type: `"string"`. Computed.

Name. The name of this registration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--namespace"></a>

### namespace property

Type: `"string"`. Computed.

Namespace. The namespace this item belongs to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

- [object](data-sources--site_registrations_by_site--properties--items--object.md): complete subsection reference.

- [owner_view](data-sources--site_registrations_by_site--properties--items--owner_view.md): complete subsection reference.

- [system_metadata](data-sources--site_registrations_by_site--properties--items--system_metadata.md): complete subsection reference.

<a id="schema-items--tenant"></a>

### tenant property

Type: `"string"`. Computed.

Tenant. The tenant this item belongs to.

<a id="schema-items--uid"></a>

### uid property

Type: `"string"`. Computed.

UID. The unique uid of this registration.

## Next pages

- [items.annotations](data-sources--site_registrations_by_site--properties--items--annotations.md)
- [items.get_spec](data-sources--site_registrations_by_site--properties--items--get_spec.md)
- [items.labels](data-sources--site_registrations_by_site--properties--items--labels.md)
- [items.metadata](data-sources--site_registrations_by_site--properties--items--metadata.md)
- [items.object](data-sources--site_registrations_by_site--properties--items--object.md)
- [items.owner_view](data-sources--site_registrations_by_site--properties--items--owner_view.md)
- [items.system_metadata](data-sources--site_registrations_by_site--properties--items--system_metadata.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
