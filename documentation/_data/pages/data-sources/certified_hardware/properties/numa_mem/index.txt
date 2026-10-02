---
page_title: "numa_mem"
subcategory: ""
description: "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes."
xcsh_docs: {"aliases": ["numa mem"], "body_bytes": 1150, "body_sha256": "sha256:aff39fdfc64356fef5b03999c8bc9bce2018a95d193e618d78951e03fe463aad", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/numa_mem/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2330233232230001-2320121010311121-0321011030100311-3021331102302000-0130032123111310-0033312132220332-3000213332113201-0111301333202103", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["numa_mem"], "schema_version": 1, "sections": [{"aliases": ["memory"], "anchor": "schema-numa_mem--memory", "description": "The number of MB of instance memory to map to instance NUMA node N.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "memory"], "syntax": "attribute", "type": "number"}, {"aliases": ["node"], "anchor": "schema-numa_mem--node", "description": "Node. NUMA node instance with mapped memory.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "node"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/numa_mem/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
