---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_segment."
xcsh_docs: {"aliases": ["segment"], "body_bytes": 346, "body_sha256": "sha256:9053b85b40b4f40a627923af5360599ca4b12088a92251c2352c1104acb93e16", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "id": "xcsh-docs:resources:segment:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:segment:fundamentals", "path": "documentation/resources/segment/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1001002102021311-1203023103330002-2021312223200031-3213233012332321-2031323020033021-2230320332123133-0002031022101301-2032211323110003", "registry_path": "docs/guides/resources--segment--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_segment.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["segmentCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_segment.example system/example
```
