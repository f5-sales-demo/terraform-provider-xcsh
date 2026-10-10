---
page_title: "items.object"
subcategory: ""
description: "Registration object stores node registration and information regarding the node."
xcsh_docs: {"aliases": ["items object"], "body_bytes": 1408, "body_sha256": "sha256:6eeec534d47fe934e78f9c3d082b05a93dee598e628d3c8d1d2a243ae113f84e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1223022213213010-2123312320000202-3212010203300330-2123121110030020-2121322322103003-3031303200312203-3323112222021022-0003011010032003", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object"], "schema_version": 1, "sections": [{"aliases": ["items object metadata"], "anchor": "section", "description": "ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec"], "anchor": "section", "description": "Configuration for registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object status"], "anchor": "section", "description": "Status Type. Most recent observer status of object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "status"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata"], "anchor": "section", "description": "SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Registration object stores node registration and information regarding the node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- items.object

<a id="section"></a>

Type: `"single"`. Computed.

Registration object stores node registration and information regarding the node.

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/): complete subsection reference.

- [status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/status/): complete subsection reference.

- [system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/): complete subsection reference.
