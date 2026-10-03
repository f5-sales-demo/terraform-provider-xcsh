---
page_title: "connections.source_segments"
subcategory: ""
description: "Configuration parameter for source segments."
xcsh_docs: {"aliases": ["connections source segments"], "body_bytes": 2170, "body_sha256": "sha256:dff59ae318aa076c68566465f78a4c2c644baaeeabda6531a6da440ebe85dba4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "parent_id": "xcsh-docs:data-sources:segment_connection:properties:connections", "path": "documentation/data-sources/segment_connection/properties/connections/source_segments/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1101020230333122-0002013101023002-3330321120220113-3200221231102220-0023223221323321-1021003033021323-0320310231031021-3111012333303301", "registry_path": "docs/guides/data-sources--segment_connection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connections", "source_segments"], "schema_version": 1, "sections": [{"aliases": ["connections source segments kind"], "anchor": "schema-connections--source_segments--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "source_segments", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["connections source segments name"], "anchor": "schema-connections--source_segments--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "source_segments", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["connections source segments namespace"], "anchor": "schema-connections--source_segments--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "source_segments", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["connections source segments tenant"], "anchor": "schema-connections--source_segments--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "source_segments", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["connections source segments uid"], "anchor": "schema-connections--source_segments--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connections", "source_segments", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/source_segments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for source segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
