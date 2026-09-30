---
page_title: "items.system_metadata"
subcategory: ""
description: "items.system_metadata for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 4138, "body_sha256": "sha256:d95adcca2db628874ddbb1ae57d5877ff1dd2625d74f3c80ba934d67f34bfac2", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:initializers", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:labels", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata:owner_view"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:system_metadata", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--system_metadata.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "system_metadata"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/system_metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.system_metadata for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.system_metadata

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
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

- [initializers](data-sources--site_registrations_by_state--properties--items--system_metadata--initializers.md): complete subsection reference.

- [labels](data-sources--site_registrations_by_state--properties--items--system_metadata--labels.md): complete subsection reference.

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

- [owner_view](data-sources--site_registrations_by_state--properties--items--system_metadata--owner_view.md): complete subsection reference.

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

## Next pages

- [items.system_metadata.initializers](data-sources--site_registrations_by_state--properties--items--system_metadata--initializers.md)
- [items.system_metadata.labels](data-sources--site_registrations_by_state--properties--items--system_metadata--labels.md)
- [items.system_metadata.owner_view](data-sources--site_registrations_by_state--properties--items--system_metadata--owner_view.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
