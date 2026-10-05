---
page_title: "items.system_metadata.initializers"
subcategory: ""
description: "Initializers tracks the progress of initialization of a configuration object."
xcsh_docs: {"aliases": ["items system metadata initializers"], "body_bytes": 1995, "body_sha256": "sha256:de736f98c4f54beb153a229657a19322e51d7f6fb36f4b169f4892f55897ff22", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:pending", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:result"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata", "path": "documentation/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1330013113012011-1311330023231332-3131000221130223-2321013000333320-2233213022221332-2222030203222020-2301003003012312-0323202333030032", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "system_metadata", "initializers"], "schema_version": 1, "sections": [{"aliases": ["items system metadata initializers pending"], "anchor": "section", "description": "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:pending", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "system_metadata", "initializers", "pending"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata initializers result"], "anchor": "section", "description": "Status is a return value for calls that don't return other objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers:result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata", "initializers", "result"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Initializers tracks the progress of initialization of a configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata.initializers

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/)
- items.system_metadata.initializers

<a id="section"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

## Direct properties

- [pending](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/pending/): complete subsection reference.

- [result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/result/): complete subsection reference.

## Next pages

- [items.system_metadata.initializers.pending](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/pending/)
- [items.system_metadata.initializers.result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/initializers/result/)
- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/system_metadata/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
