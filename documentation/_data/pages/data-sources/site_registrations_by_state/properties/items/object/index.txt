---
page_title: "items.object"
subcategory: ""
description: "Registration object stores node registration and information regarding the node."
xcsh_docs: {"aliases": ["items object"], "body_bytes": 2320, "body_sha256": "sha256:f73670d6fa803ff73c5860db06863a45357c6c15fd3f62cc8c903fd2d2436450", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1223022213213010-2123312320000202-3212010203300330-2123121110030020-2121322322103003-3031303200312203-3323112222021022-0003011010032003", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object"], "schema_version": 1, "sections": [{"aliases": ["items object metadata"], "anchor": "section", "description": "ObjectMetaType is metadata(common attributes) of an object that all configuration objects will have. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec"], "anchor": "section", "description": "Configuration for registration.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object status"], "anchor": "section", "description": "Status Type. Most recent observer status of object.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "status"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata"], "anchor": "section", "description": "SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Registration object stores node registration and information regarding the node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [items.object.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/metadata/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/status/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
