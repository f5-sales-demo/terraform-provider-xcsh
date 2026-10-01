---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 4567, "body_sha256": "sha256:c631198d6af981a647e6bae287effc72e4713511ac716124c71ae224172c8b50", "canonical_id": "xcsh-docs:data-sources:segment_connection:reference", "child_ids": ["xcsh-docs:data-sources:segment_connection:properties:connections"], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:reference", "parent_id": "xcsh-docs:data-sources:segment_connection:fundamentals", "path": "docs/guides/data-sources--segment_connection--reference.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [connections](data-sources--segment_connection--properties--connections.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the SegmentConnection to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the SegmentConnection.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--segment_connection--reference.md#schema-annotations) |
| `connections` | [connections](data-sources--segment_connection--properties--connections.md#section) |
| `connections.destination_segments` | [connections.destination_segments](data-sources--segment_connection--properties--connections--destination_segments.md#section) |
| `connections.destination_segments.kind` | [connections.destination_segments.kind](data-sources--segment_connection--properties--connections--destination_segments.md#schema-connections--destination_segments--kind) |
| `connections.destination_segments.name` | [connections.destination_segments.name](data-sources--segment_connection--properties--connections--destination_segments.md#schema-connections--destination_segments--name) |
| `connections.destination_segments.namespace` | [connections.destination_segments.namespace](data-sources--segment_connection--properties--connections--destination_segments.md#schema-connections--destination_segments--namespace) |
| `connections.destination_segments.tenant` | [connections.destination_segments.tenant](data-sources--segment_connection--properties--connections--destination_segments.md#schema-connections--destination_segments--tenant) |
| `connections.destination_segments.uid` | [connections.destination_segments.uid](data-sources--segment_connection--properties--connections--destination_segments.md#schema-connections--destination_segments--uid) |
| `connections.direct` | [connections.direct](data-sources--segment_connection--properties--connections--direct.md#section) |
| `connections.source_segments` | [connections.source_segments](data-sources--segment_connection--properties--connections--source_segments.md#section) |
| `connections.source_segments.kind` | [connections.source_segments.kind](data-sources--segment_connection--properties--connections--source_segments.md#schema-connections--source_segments--kind) |
| `connections.source_segments.name` | [connections.source_segments.name](data-sources--segment_connection--properties--connections--source_segments.md#schema-connections--source_segments--name) |
| `connections.source_segments.namespace` | [connections.source_segments.namespace](data-sources--segment_connection--properties--connections--source_segments.md#schema-connections--source_segments--namespace) |
| `connections.source_segments.tenant` | [connections.source_segments.tenant](data-sources--segment_connection--properties--connections--source_segments.md#schema-connections--source_segments--tenant) |
| `connections.source_segments.uid` | [connections.source_segments.uid](data-sources--segment_connection--properties--connections--source_segments.md#schema-connections--source_segments--uid) |
| `description` | [description](data-sources--segment_connection--reference.md#schema-description) |
| `id` | [id](data-sources--segment_connection--reference.md#schema-id) |
| `labels` | [labels](data-sources--segment_connection--reference.md#schema-labels) |
| `name` | [name](data-sources--segment_connection--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--segment_connection--reference.md#schema-namespace) |

## Next pages

- [connections](data-sources--segment_connection--properties--connections.md)
- [xcsh_segment_connection](../data-sources/segment_connection.md)
