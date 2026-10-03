---
page_title: "items.object.system_metadata.initializers"
subcategory: ""
description: "Initializers tracks the progress of initialization of a configuration object."
xcsh_docs: {"aliases": ["items object system metadata initializers"], "body_bytes": 2219, "body_sha256": "sha256:2a3a9fa6ecb1cbc4551a23906f6104a578efdf1a4f9b841b63573dfb29cf61e9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:pending", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2121321020130223-3321220332111003-3033012023232233-3200300102203020-2220332102023033-1233321221111301-1100022123211132-1010220223313232", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata initializers pending"], "anchor": "section", "description": "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:pending", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "pending"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata initializers result"], "anchor": "section", "description": "Status is a return value for calls that don't return other objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Initializers tracks the progress of initialization of a configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.initializers

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- items.object.system_metadata.initializers

<a id="section"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

## Direct properties

- [pending](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/pending/): complete subsection reference.

- [result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/result/): complete subsection reference.

## Next pages

- [items.object.system_metadata.initializers.pending](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/pending/)
- [items.object.system_metadata.initializers.result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/result/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
