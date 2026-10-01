---
page_title: "connections.destination_segments"
subcategory: ""
description: "connections.destination_segments for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 1953, "body_sha256": "sha256:816389a6a40dc9125b50acc18cb5243d43f48404d5d21aac1285b13480a3e30a", "canonical_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "child_ids": [], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "parent_id": "xcsh-docs:data-sources:segment_connection:properties:connections", "path": "docs/guides/data-sources--segment_connection--properties--connections--destination_segments.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connections", "destination_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/destination_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connections.destination_segments for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connections.destination_segments

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md)
- [Property reference](data-sources--segment_connection--reference.md)
- [connections](data-sources--segment_connection--properties--connections.md)
- connections.destination_segments

<a id="section"></a>

Type: `"list"`. Computed.

Configuration parameter for destination segments.

## Direct properties

<a id="schema-connections--destination_segments--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-connections--destination_segments--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-connections--destination_segments--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-connections--destination_segments--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-connections--destination_segments--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [connections](data-sources--segment_connection--properties--connections.md)
- [xcsh_segment_connection](../data-sources/segment_connection.md)
