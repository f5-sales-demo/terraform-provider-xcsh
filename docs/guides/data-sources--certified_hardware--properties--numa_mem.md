---
page_title: "numa_mem"
subcategory: ""
description: "numa_mem for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 942, "body_sha256": "sha256:a3222c7aaf97a1b4adf5d7242251454f8e45fe05ae1765a88c35ec2fef2dc336", "canonical_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "docs/guides/data-sources--certified_hardware--properties--numa_mem.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["numa_mem"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/numa_mem/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "numa_mem for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# numa_mem

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
- [Property reference](data-sources--certified_hardware--reference.md)
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

- [Property reference](data-sources--certified_hardware--reference.md)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
