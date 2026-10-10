---
page_title: "items.system_metadata"
subcategory: ""
description: "SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users."
xcsh_docs: {"aliases": ["items system metadata"], "body_bytes": 3911, "body_sha256": "sha256:001bb240f8cca8f4ff42fb6f1aa7fa6743104dfb83c0e95d92e823ea5da3dfa4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers", "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:labels", "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items", "path": "documentation/data-sources/site_registrations/properties/items/system_metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0322000312200113-1132021031313322-3203023002102221-3300330123332101-0111312300233330-3100121121000211-3030031131013210-3032323020221203", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "system_metadata"], "schema_version": 1, "sections": [{"aliases": ["items system metadata creation timestamp"], "anchor": "schema-items--system_metadata--creation_timestamp", "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "creation_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata creator class"], "anchor": "schema-items--system_metadata--creator_class", "description": "Value identifying the class of the user or service which created this configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "creator_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata creator id"], "anchor": "schema-items--system_metadata--creator_id", "description": "Value identifying the exact user or service that created this configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "creator_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata deletion timestamp"], "anchor": "schema-items--system_metadata--deletion_timestamp", "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not..", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "deletion_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata finalizers"], "anchor": "schema-items--system_metadata--finalizers", "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "finalizers"], "syntax": "attribute", "type": "list"}, {"aliases": ["items system metadata initializers"], "anchor": "section", "description": "Initializers tracks the progress of initialization of a configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata", "initializers"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the operator or software. Values here can be interpreted by software(backend or frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata modification timestamp"], "anchor": "schema-items--system_metadata--modification_timestamp", "description": "ModificationTimestamp is a timestamp representing the server time when this object was last modified.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "modification_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata object index"], "anchor": "schema-items--system_metadata--object_index", "description": "Unique index for the object. Some objects need a unique integer index to be allocated for each object type. This field will be populated for all objects that need it and will be zero otherwise.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "object_index"], "syntax": "attribute", "type": "number"}, {"aliases": ["items system metadata owner view"], "anchor": "section", "description": "ViewRefType represents a reference to a view.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:owner_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "system_metadata", "owner_view"], "syntax": "attribute", "type": "object"}, {"aliases": ["items system metadata tenant"], "anchor": "schema-items--system_metadata--tenant", "description": "Tenant to which this configuration object belongs to. The value for this is found from presented credentials.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items system metadata uid", "succeeded", "success", "successful"], "anchor": "schema-items--system_metadata--uid", "description": "Uid is the unique in time and space value for this object. It is generated by the server on successful creation of an object and is not allowed to change on Replace API. The value of is taken from uid field of ObjectMetaType, if provided.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/system_metadata/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- items.system_metadata

<a id="section"></a>

Type: `"single"`. Computed.

SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects
and cannot be updated directly by users.

## Direct properties

<a id="schema-items--system_metadata--creation_timestamp"></a>

### creation_timestamp property

Type: `"string"`. Computed.

CreationTimestamp is a timestamp representing the server time when this object was created. It is
not guaranteed to be set in happens-before order across separate operations. Clients may not set
this value.

<a id="schema-items--system_metadata--creator_class"></a>

### creator_class property

Type: `"string"`. Computed.

Value identifying the class of the user or service which created this configuration object.

<a id="schema-items--system_metadata--creator_id"></a>

### creator_id property

Type: `"string"`. Computed.

Value identifying the exact user or service that created this configuration object.

<a id="schema-items--system_metadata--deletion_timestamp"></a>

### deletion_timestamp property

Type: `"string"`. Computed.

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is
set by the server when a graceful deletion is requested by the user, and is not directly settable by
a client. The resource is expected to be deleted (no longer visible from resource lists, and not..

<a id="schema-items--system_metadata--finalizers"></a>

### finalizers property

Type: `["list", "string"]`. Computed.

Must be empty before the object is deleted from the registry. Each entry is an identifier for the
responsible component that will remove the entry from the list. If the deletionTimestamp of the
object is non-nil, entries in this list can only be removed.

- [initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/initializers/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/labels/): complete subsection reference.

<a id="schema-items--system_metadata--modification_timestamp"></a>

### modification_timestamp property

Type: `"string"`. Computed.

ModificationTimestamp is a timestamp representing the server time when this object was last
modified.

<a id="schema-items--system_metadata--object_index"></a>

### object_index property

Type: `"number"`. Computed.

Unique index for the object. Some objects need a unique integer index to be allocated for each
object type. This field will be populated for all objects that need it and will be zero otherwise.

- [owner_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/owner_view/): complete subsection reference.

<a id="schema-items--system_metadata--tenant"></a>

### tenant property

Type: `"string"`. Computed.

Tenant to which this configuration object belongs to. The value for this is found from presented
credentials.

<a id="schema-items--system_metadata--uid"></a>

### uid property

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. It is generated by the server on
successful creation of an object and is not allowed to change on Replace API. The value of is taken
from uid field of ObjectMetaType, if provided.
