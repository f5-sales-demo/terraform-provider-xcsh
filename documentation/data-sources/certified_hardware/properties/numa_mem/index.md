---
page_title: "numa_mem"
subcategory: ""
description: "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes."
xcsh_docs: {"aliases": ["numa mem"], "body_bytes": 1150, "body_sha256": "sha256:aff39fdfc64356fef5b03999c8bc9bce2018a95d193e618d78951e03fe463aad", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/numa_mem/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2330233232230001-2320121010311121-0321011030100311-3021331102302000-0130032123111310-0033312132220332-3000213332113201-0111301333202103", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["numa_mem"], "schema_version": 1, "sections": [{"aliases": ["memory"], "anchor": "schema-numa_mem--memory", "description": "The number of MB of instance memory to map to instance NUMA node N.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "memory"], "syntax": "attribute", "type": "number"}, {"aliases": ["node"], "anchor": "schema-numa_mem--node", "description": "Node. NUMA node instance with mapped memory.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_mem", "node"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/numa_mem/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
