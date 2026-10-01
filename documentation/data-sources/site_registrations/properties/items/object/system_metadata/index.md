---
page_title: "items.object.system_metadata"
subcategory: ""
description: "items.object.system_metadata for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 7501, "body_sha256": "sha256:2cd32c0cb47a9dd4ac99f5ed6aab4fc225c65b786ab22f7455d2a781a49eab15", "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:initializers", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:labels", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:namespace", "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata:owner_view"], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:system_metadata", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object", "path": "documentation/data-sources/site_registrations/properties/items/object/system_metadata/index.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "object", "system_metadata"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/system_metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.system_metadata for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
