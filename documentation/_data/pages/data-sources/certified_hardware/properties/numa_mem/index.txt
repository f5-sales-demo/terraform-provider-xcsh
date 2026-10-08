---
page_title: "numa_mem"
subcategory: ""
description: "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes."
xcsh_docs: {"aliases": ["numa mem"], "body_bytes": 892, "body_sha256": "sha256:72a5b314304589c0c145e68d91d176c1480d49fab4890cae00413a02c3430a6c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/numa_mem/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2330233232230001-2320121010311121-0321011030100311-3021331102302000-0130032123111310-0033312132220332-3000213332113201-0111301333202103", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["numa_mem"], "schema_version": 1, "sections": [{"aliases": ["numa mem memory"], "anchor": "schema-numa_mem--memory", "description": "The number of MB of instance memory to map to instance NUMA node N.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "memory"], "syntax": "attribute", "type": "number"}, {"aliases": ["numa mem node"], "anchor": "schema-numa_mem--node", "description": "Node. NUMA node instance with mapped memory.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "node"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/numa_mem/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# numa_mem

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- numa_mem

<a id="section"></a>

Type: `"list"`. Computed.

List of Numa nodes with the number of MB of instance memory to map to node instance If not
specified, memory is evenly divided among available NUMA nodes.

## Direct properties

<a id="schema-numa_mem--memory"></a>

### memory property

Type: `"number"`. Computed.

The number of MB of instance memory to map to instance NUMA node N.

<a id="schema-numa_mem--node"></a>

### node property

Type: `"number"`. Computed.

Node. NUMA node instance with mapped memory.
