---
page_title: "connections"
subcategory: ""
description: "Configure a Segment Connector to allow network traffic between Segments."
xcsh_docs: {"aliases": ["connections"], "body_bytes": 1800, "body_sha256": "sha256:776212bc49111398f346f52a362a4a46270cd536197df79807c728727a72f814", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "xcsh-docs:data-sources:segment_connection:properties:connections:direct", "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:properties:connections", "parent_id": "xcsh-docs:data-sources:segment_connection:reference", "path": "documentation/data-sources/segment_connection/properties/connections/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1231031300300121-2311310111302032-0202012003300032-2323223320112011-0021200220023011-3221022101102002-2131111311012323-3032300303212311", "registry_path": "docs/guides/data-sources--segment_connection--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connections"], "schema_version": 1, "sections": [{"aliases": ["destination segments"], "anchor": "section", "description": "Configuration parameter for destination segments.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:destination_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["connections", "destination_segments"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:direct", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["connections", "direct"], "syntax": "attribute", "type": "object"}, {"aliases": ["source segments"], "anchor": "section", "description": "Configuration parameter for source segments.", "document_id": "xcsh-docs:data-sources:segment_connection:properties:connections:source_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["connections", "source_segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/properties/connections/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure a Segment Connector to allow network traffic between Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [connections.destination_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/destination_segments/)
- [connections.direct](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/direct/)
- [connections.source_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/connections/source_segments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
