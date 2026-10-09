---
page_title: "items.object"
subcategory: ""
description: "Registration object stores node registration and information regarding the node."
xcsh_docs: {"aliases": ["items object"], "body_bytes": 1336, "body_sha256": "sha256:0f342da14c4f7662b5370b538fc9c68e704472c23664a92dd66bc6eef082d814", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:object:metadata", "xcsh-docs:data-sources:site_registrations:properties:items:object:spec", "xcsh-docs:data-sources:site_registrations:properties:items:object:status", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "documentation/data-sources/site_registrations/properties/items/object/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3232231033322100-3201310333020303-3010031030002223-1023233122220331-2222333223031331-0202221320012112-0331222300131210-0220102020130021", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object"], "schema_version": 1, "sections": [{"aliases": ["items object metadata"], "anchor": "section", "description": "ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec"], "anchor": "section", "description": "Configuration for registration.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object status"], "anchor": "section", "description": "Status Type. Most recent observer status of object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:status", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "status"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata"], "anchor": "section", "description": "SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Registration object stores node registration and information regarding the node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- items.object

<a id="section"></a>

Type: `"single"`. Computed.

Registration object stores node registration and information regarding the node.

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/): complete subsection reference.

- [status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/status/): complete subsection reference.

- [system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/): complete subsection reference.
