---
page_title: "connections.source_segments"
subcategory: ""
description: "connections.source_segments for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 2071, "body_sha256": "sha256:d1a015f85c65b7eb8ce591bee8dec6657e243dcd61f9ddaf5d793dabdc2c67d2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "parent_id": "xcsh-docs:data-sources:segment_connection:properties:connections", "path": "documentation/data-sources/segment_connection/properties/connections/source_segments/index.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["connections", "source_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/source_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "connections.source_segments for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# connections.source_segments

Breadcrumbs:

- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/)
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

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/)
- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
