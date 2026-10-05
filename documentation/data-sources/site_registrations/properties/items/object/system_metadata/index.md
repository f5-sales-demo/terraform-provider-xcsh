---
page_title: "items.object.system_metadata"
subcategory: ""
description: "SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users."
xcsh_docs: {"aliases": ["items object system metadata"], "body_bytes": 7501, "body_sha256": "sha256:2cd32c0cb47a9dd4ac99f5ed6aab4fc225c65b786ab22f7455d2a781a49eab15", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:labels", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:namespace", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:owner_view"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object", "path": "documentation/data-sources/site_registrations/properties/items/object/system_metadata/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0002122310011211-2113223023200301-2122000032130322-0120013233032121-3220102321300022-1201332312302233-0230300030303023-3100002110032210", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata creation timestamp"], "anchor": "schema-items--object--system_metadata--creation_timestamp", "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "creation_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata creator class"], "anchor": "schema-items--object--system_metadata--creator_class", "description": "Value identifying the class of the user or service which created this configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "creator_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata creator cookie"], "anchor": "schema-items--object--system_metadata--creator_cookie", "description": "Can used by the creator of the object for later audit for e.g. By storing the version identifying information of the object so at future it can be determined if version present at remote end is current or stale.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "creator_cookie"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata creator id"], "anchor": "schema-items--object--system_metadata--creator_id", "description": "Value identifying the exact user or service that created this configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "creator_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata deletion timestamp"], "anchor": "schema-items--object--system_metadata--deletion_timestamp", "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not..", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "deletion_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata direct ref hash"], "anchor": "schema-items--object--system_metadata--direct_ref_hash", "description": "Hash of the UIDs of direct references on this object. This can be used to determine if this object hash has had references become resolved/unresolved.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "direct_ref_hash"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata finalizers"], "anchor": "schema-items--object--system_metadata--finalizers", "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "finalizers"], "syntax": "attribute", "type": "list"}, {"aliases": ["items object system metadata initializers"], "anchor": "section", "description": "Initializers tracks the progress of initialization of a configuration object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the operator or software. Values here can be interpreted by software(backend or frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata modification timestamp"], "anchor": "schema-items--object--system_metadata--modification_timestamp", "description": "ModificationTimestamp is a timestamp representing the server time when this object was last modified.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "modification_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata namespace"], "anchor": "section", "description": "The namespace this object belongs to. This is populated by the service based on the metadata.namespace field when an object is created.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:namespace", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "system_metadata", "namespace"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata object index"], "anchor": "schema-items--object--system_metadata--object_index", "description": "Unique index for the object. Some objects need a unique integer index to be allocated for each object type. This field will be populated for all objects that need it and will be zero otherwise.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "object_index"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object system metadata owner view"], "anchor": "section", "description": "ViewRefType represents a reference to a view.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:owner_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "system_metadata", "owner_view"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object system metadata revision"], "anchor": "schema-items--object--system_metadata--revision", "description": "Revision number which always increases with each modification of the object in storage This doesn't necessarily increase sequentially, but should always increase. This will be 0 when first created, and before any modifications.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "revision"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata sre disable"], "anchor": "schema-items--object--system_metadata--sre_disable", "description": "Should be set to true If F5XC/SRE operator wants to suppress an object from being presented to business-logic of a daemon(e.g. Due to bad-form/issue-causing Object). This is meant only to be used in temporary situations for operational continuity till a fix is rolled out in business-logic.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "sre_disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["items object system metadata tenant"], "anchor": "schema-items--object--system_metadata--tenant", "description": "Tenant to which this configuration object belongs to. The value for this is found from presented credentials.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata trace info"], "anchor": "schema-items--object--system_metadata--trace_info", "description": "Trace_info holds information(<trace-ID>:<span-ID>:<parent-span-ID>) of the request doing the object modification. This can be used on the watch side to create subsequent spans. This information can be used to co-relate activities across services (modulo state compression) for a synchronous API.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "trace_info"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata uid", "succeeded", "success", "successful"], "anchor": "schema-items--object--system_metadata--uid", "description": "Uid is the unique in time and space value for this object. It is generated by the server on successful creation of an object and is not allowed to change on Replace API. The value of is taken from uid field of ObjectMetaType, if provided.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "uid"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata vtrp id"], "anchor": "schema-items--object--system_metadata--vtrp_id", "description": "VTRP ID. Indicate origin of this object.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "vtrp_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata vtrp stale"], "anchor": "schema-items--object--system_metadata--vtrp_stale", "description": "Indicate whether mars deems this object to be stale via graceful restart timer information.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "vtrp_stale"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/system_metadata/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- items.object.system_metadata

<a id="section"></a>

Type: `"single"`. Computed.

SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and
cannot be updated directly by users.

## Direct properties

<a id="schema-items--object--system_metadata--creation_timestamp"></a>

### creation_timestamp property

Type: `"string"`. Computed.

CreationTimestamp is a timestamp representing the server time when this object was created. It is
not guaranteed to be set in happens-before order across separate operations. Clients may not set
this value.

<a id="schema-items--object--system_metadata--creator_class"></a>

### creator_class property

Type: `"string"`. Computed.

Value identifying the class of the user or service which created this configuration object.

<a id="schema-items--object--system_metadata--creator_cookie"></a>

### creator_cookie property

Type: `"string"`. Computed.

Can used by the creator of the object for later audit for e.g. By storing the version identifying
information of the object so at future it can be determined if version present at remote end is
current or stale.

<a id="schema-items--object--system_metadata--creator_id"></a>

### creator_id property

Type: `"string"`. Computed.

Value identifying the exact user or service that created this configuration object.

<a id="schema-items--object--system_metadata--deletion_timestamp"></a>

### deletion_timestamp property

Type: `"string"`. Computed.

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is
set by the server when a graceful deletion is requested by the user, and is not directly settable by
a client. The resource is expected to be deleted (no longer visible from resource lists, and not..

<a id="schema-items--object--system_metadata--direct_ref_hash"></a>

### direct_ref_hash property

Type: `"string"`. Computed.

Hash of the UIDs of direct references on this object. This can be used to determine if this object
hash has had references become resolved/unresolved.

<a id="schema-items--object--system_metadata--finalizers"></a>

### finalizers property

Type: `["list", "string"]`. Computed.

Must be empty before the object is deleted from the registry. Each entry is an identifier for the
responsible component that will remove the entry from the list. If the deletionTimestamp of the
object is non-nil, entries in this list can only be removed.

- [initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/initializers/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/labels/): complete subsection reference.

<a id="schema-items--object--system_metadata--modification_timestamp"></a>

### modification_timestamp property

Type: `"string"`. Computed.

ModificationTimestamp is a timestamp representing the server time when this object was last
modified.

- [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/namespace/): complete subsection reference.

<a id="schema-items--object--system_metadata--object_index"></a>

### object_index property

Type: `"number"`. Computed.

Unique index for the object. Some objects need a unique integer index to be allocated for each
object type. This field will be populated for all objects that need it and will be zero otherwise.

- [owner_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/owner_view/): complete subsection reference.

<a id="schema-items--object--system_metadata--revision"></a>

### revision property

Type: `"string"`. Computed.

Revision number which always increases with each modification of the object in storage This doesn't
necessarily increase sequentially, but should always increase. This will be 0 when first created,
and before any modifications.

<a id="schema-items--object--system_metadata--sre_disable"></a>

### sre_disable property

Type: `"bool"`. Computed.

Should be set to true If F5XC/SRE operator wants to suppress an object from being presented to
business-logic of a daemon(e.g. Due to bad-form/issue-causing Object). This is meant only to be used
in temporary situations for operational continuity till a fix is rolled out in business-logic.

<a id="schema-items--object--system_metadata--tenant"></a>

### tenant property

Type: `"string"`. Computed.

Tenant to which this configuration object belongs to. The value for this is found from presented
credentials.

<a id="schema-items--object--system_metadata--trace_info"></a>

### trace_info property

Type: `"string"`. Computed.

Trace\_info holds information(&lt;trace-ID&gt;:&lt;span-ID&gt;:&lt;parent-span-ID&gt;) of the
request doing the object modification. This can be used on the watch side to create subsequent
spans. This information can be used to co-relate activities across services (modulo state
compression) for a synchronous API.

<a id="schema-items--object--system_metadata--uid"></a>

### uid property

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. It is generated by the server on
successful creation of an object and is not allowed to change on Replace API. The value of is taken
from uid field of ObjectMetaType, if provided.

<a id="schema-items--object--system_metadata--vtrp_id"></a>

### vtrp_id property

Type: `"string"`. Computed.

VTRP ID. Indicate origin of this object.

<a id="schema-items--object--system_metadata--vtrp_stale"></a>

### vtrp_stale property

Type: `"bool"`. Computed.

Indicate whether mars deems this object to be stale via graceful restart timer information.

## Next pages

- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/initializers/)
- [items.object.system_metadata.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/labels/)
- [items.object.system_metadata.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/namespace/)
- [items.object.system_metadata.owner_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/system_metadata/owner_view/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
