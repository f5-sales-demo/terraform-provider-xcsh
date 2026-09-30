---
page_title: "numa_mem"
subcategory: ""
description: "numa_mem for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 843, "body_sha256": "sha256:dd6a4b28b6f71d21a2775381ff39e14c5478fe0bc105fcd35e0a7d12e285a6a9", "canonical_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "docs/guides/data-sources--certified_hardware--properties--numa_mem.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["numa_mem"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/numa_mem/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "numa_mem for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
