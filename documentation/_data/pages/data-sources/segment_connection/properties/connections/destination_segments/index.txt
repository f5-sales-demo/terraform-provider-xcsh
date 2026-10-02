---
page_title: "connections.destination_segments"
subcategory: ""
description: "Configuration parameter for destination segments."
xcsh_docs: {"aliases": ["connections destination segments"], "body_bytes": 2210, "body_sha256": "sha256:277effec06a5b3280956431f500e2526ca53e50908972552ca109e5787227ab8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "parent_id": "xcsh-docs:data-sources:segment_connection:properties:connections", "path": "documentation/data-sources/segment_connection/properties/connections/destination_segments/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2111121120233300-2202200133023311-3232311111002301-0313312121011330-0231100012123312-3203113023122212-2233032301102213-1130310121021232", "registry_path": "docs/guides/data-sources--segment_connection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connections", "destination_segments"], "schema_version": 1, "sections": [{"aliases": ["kind"], "anchor": "schema-connections--destination_segments--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "destination_segments", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-connections--destination_segments--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "destination_segments", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-connections--destination_segments--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "destination_segments", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-connections--destination_segments--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "destination_segments", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["uid"], "anchor": "schema-connections--destination_segments--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "destination_segments", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/destination_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for destination segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connections.destination_segments

Breadcrumbs:

- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/)
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

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/)
- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
