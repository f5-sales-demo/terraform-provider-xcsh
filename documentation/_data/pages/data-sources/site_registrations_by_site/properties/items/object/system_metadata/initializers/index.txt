---
page_title: "items.object.system_metadata.initializers"
subcategory: ""
description: "Initializers tracks the progress of initialization of a configuration object."
xcsh_docs: {"aliases": ["items object system metadata initializers"], "body_bytes": 1461, "body_sha256": "sha256:2f9c2146d6d2f0bcac42674e08c82fc3711e940e7748464199f192016b8016cf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:pending", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0002213333320022-1330230231002000-2201123330223120-1302312313003333-2311300321110020-2233301301210233-1110002323232033-0210230202203311", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata initializers pending"], "anchor": "section", "description": "Pending is a list of initializers that must execute in order before this object is initialized. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:pending", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "pending"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata initializers result"], "anchor": "section", "description": "Status is a return value for calls that don't return other objects.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Initializers tracks the progress of initialization of a configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.initializers

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/)
- items.object.system_metadata.initializers

<a id="section"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

## Direct properties

- [pending](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/pending/): complete subsection reference.

- [result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/result/): complete subsection reference.
