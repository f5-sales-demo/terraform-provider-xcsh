---
page_title: "items.object.system_metadata.initializers.pending"
subcategory: ""
description: "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients."
xcsh_docs: {"aliases": ["items object system metadata initializers pending"], "body_bytes": 2175, "body_sha256": "sha256:67e6177e351c3d881d513fe7de7c0517b9f2462c43ffabf14769e63f7098cef1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers:pending", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers", "path": "documentation/data-sources/site_registrations/properties/items/object/system_metadata/initializers/pending/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3210013023310313-2202310233200330-1310131023002130-0230213202010222-0220032111313002-1210330003203232-2313122033311212-1332212333310311", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers", "pending"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-items--object--system_metadata--initializers--pending--name", "description": "Name of the service that is responsible for initializing this object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers:pending", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "pending", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/system_metadata/initializers/pending/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
