---
page_title: "pagination"
subcategory: ""
description: "Pagination for Request with number and size."
xcsh_docs: {"aliases": ["pagination"], "body_bytes": 994, "body_sha256": "sha256:6bc04115107a177e41383fa240279804e4d1db9b10adb1d58607cc985a86f659", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "documentation/data-sources/device_intelligence_device_history/properties/pagination/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1100211222301000-2011120232120333-2313010233103222-1211133021031310-1033002013001121-3022231203222231-2013022331323131-2010311231202002", "registry_path": "docs/guides/data-sources--device_intelligence_device_history--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["pagination"], "schema_version": 1, "sections": [{"aliases": ["pagination page number"], "anchor": "schema-pagination--page_number", "description": "Configuration parameter for page number.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pagination", "page_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["pagination page size"], "anchor": "schema-pagination--page_size", "description": "Page Size. Size or capacity specification", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pagination", "page_size"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/pagination/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Pagination for Request with number and size.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# pagination

Breadcrumbs:

- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/)
- pagination

<a id="section"></a>

Type: `"single"`. Optional.

Pagination for Request with number and size.

## Direct properties

<a id="schema-pagination--page_number"></a>

### page_number property

Type: `"number"`. Optional.

Configuration parameter for page number.

<a id="schema-pagination--page_size"></a>

### page_size property

Type: `"number"`. Optional.

Page Size. Size or capacity specification

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```
