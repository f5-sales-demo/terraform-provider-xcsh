---
page_title: "connections"
subcategory: ""
description: "Configure a Segment Connector to allow network traffic between Segments."
xcsh_docs: {"aliases": ["connections"], "body_bytes": 1067, "body_sha256": "sha256:ed0e3e3ca70ef2c944acd2b6ad411b0da1ffb77da6787388308718433a8b0446", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "xcsh-docs:data-sources:segment_connection:properties:connections:direct", "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections", "parent_id": "xcsh-docs:data-sources:segment_connection:reference", "path": "documentation/data-sources/segment_connection/properties/connections/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1231031300300121-2311310111302032-0202012003300032-2323223320112011-0021200220023011-3221022101102002-2131111311012323-3032300303212311", "registry_path": "docs/guides/data-sources--segment_connection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connections"], "schema_version": 1, "sections": [{"aliases": ["connections destination segments"], "anchor": "section", "description": "Configuration parameter for destination segments.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["connections", "destination_segments"], "syntax": "attribute", "type": "object"}, {"aliases": ["connections direct"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:direct", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["connections", "direct"], "syntax": "attribute", "type": "object"}, {"aliases": ["connections source segments"], "anchor": "section", "description": "Configuration parameter for source segments.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["connections", "source_segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configure a Segment Connector to allow network traffic between Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connections

Breadcrumbs:

- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- connections

<a id="section"></a>

Type: `"list"`. Computed.

Configure a Segment Connector to allow network traffic between Segments.

## Direct properties

- [destination_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/destination_segments/): complete subsection reference.

- [direct](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/direct/): complete subsection reference.

- [source_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/source_segments/): complete subsection reference.
