---
page_title: "items.system_metadata.initializers.pending"
subcategory: ""
description: "items.system_metadata.initializers.pending for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 1707, "body_sha256": "sha256:5d0b5fdc77d7d8d617ff91d529274cda515b8967e5067b749c30f573d7534eba", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:pending", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:pending", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--system_metadata--initializers--pending.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "system_metadata", "initializers", "pending"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/pending/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.system_metadata.initializers.pending for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata.initializers.pending

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [items.system_metadata](data-sources--site_registrations_by_state--properties--items--system_metadata.md)
- [items.system_metadata.initializers](data-sources--site_registrations_by_state--properties--items--system_metadata--initializers.md)
- items.system_metadata.initializers.pending

<a id="section"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

## Direct properties

<a id="schema-items--system_metadata--initializers--pending--name"></a>

### name property

Type: `"string"`. Computed.

Name of the service that is responsible for initializing this object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

## Next pages

- [items.system_metadata.initializers](data-sources--site_registrations_by_state--properties--items--system_metadata--initializers.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
