---
page_title: "items.metadata"
subcategory: ""
description: "items.metadata for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 2732, "body_sha256": "sha256:446fdf1b23e516efd03893cd71391763e5fe82e22ddd0bca962fad8ba10137f7", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata:labels"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:metadata", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items", "path": "docs/guides/data-sources--site_registrations_by_site--properties--items--metadata.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "metadata"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.metadata for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.metadata

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
- [Property reference](data-sources--site_registrations_by_site--reference.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- items.metadata

<a id="section"></a>

Type: `"single"`. Computed.

ObjectGetMetaType is metadata that can be specified in GET/Create response of an object.

## Direct properties

<a id="schema-items--metadata--annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

<a id="schema-items--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Human readable description for the object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1200),
}
```

<a id="schema-items--metadata--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Value of true will administratively disable the object.

- [labels](data-sources--site_registrations_by_site--properties--items--metadata--labels.md): complete subsection reference.

<a id="schema-items--metadata--name"></a>

### name property

Type: `"string"`. Computed.

Name of configuration object. It has to be unique within the namespace. It can only be specified
during create API and cannot be changed during replace API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--metadata--namespace"></a>

### namespace property

Type: `"string"`. Computed.

Defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ''.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

## Next pages

- [items.metadata.labels](data-sources--site_registrations_by_site--properties--items--metadata--labels.md)
- [items](data-sources--site_registrations_by_site--properties--items.md)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md)
