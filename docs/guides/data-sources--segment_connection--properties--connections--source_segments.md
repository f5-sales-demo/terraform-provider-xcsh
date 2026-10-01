---
page_title: "connections.source_segments"
subcategory: ""
description: "connections.source_segments for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 1913, "body_sha256": "sha256:eb1f159ba49e4648b6d7f5fba6959770b55dfe6ff66be775dcd10bc70a8fdcc5", "canonical_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "child_ids": [], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "parent_id": "xcsh-docs:data-sources:segment_connection:properties:connections", "path": "docs/guides/data-sources--segment_connection--properties--connections--source_segments.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["connections", "source_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/source_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connections.source_segments for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connections.source_segments

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md)
- [Property reference](data-sources--segment_connection--reference.md)
- [connections](data-sources--segment_connection--properties--connections.md)
- connections.source_segments

<a id="section"></a>

Type: `"list"`. Computed.

Configuration parameter for source segments.

## Direct properties

<a id="schema-connections--source_segments--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-connections--source_segments--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-connections--source_segments--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-connections--source_segments--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-connections--source_segments--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [connections](data-sources--segment_connection--properties--connections.md)
- [xcsh_segment_connection](../data-sources/segment_connection.md)
