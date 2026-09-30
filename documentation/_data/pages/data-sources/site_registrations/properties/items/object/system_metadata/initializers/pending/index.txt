---
page_title: "items.object.system_metadata.initializers.pending"
subcategory: ""
description: "items.object.system_metadata.initializers.pending for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 2076, "body_sha256": "sha256:30e96be6c0147435691b48066d0b7f504b960597f09660f14ef036ca9f215658", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers:pending", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers", "path": "documentation/data-sources/site_registrations/properties/items/object/system_metadata/initializers/pending/index.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers", "pending"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/system_metadata/initializers/pending/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.system_metadata.initializers.pending for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.system_metadata.initializers.pending

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/)
- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/initializers/)
- items.object.system_metadata.initializers.pending

<a id="section"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

## Direct properties

<a id="schema-items--object--system_metadata--initializers--pending--name"></a>

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

- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/initializers/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
