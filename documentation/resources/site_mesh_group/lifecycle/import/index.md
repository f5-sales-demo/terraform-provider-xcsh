---
page_title: "Import"
subcategory: "Infrastructure"
description: "Import for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["site mesh group"], "body_bytes": 370, "body_sha256": "sha256:d3e322d3247c3350c0b149c12d704f853a486ee72de0b4c8ea76caa37acbd61f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:site_mesh_group:fundamentals", "path": "documentation/resources/site_mesh_group/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1312210011120113-1303221000111003-1021230032103113-3203130011102120-1211003202133203-1120330320010200-2120311120331102-3303123223322310", "registry_path": "docs/guides/resources--site_mesh_group--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_site_mesh_group.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_site_mesh_group.example system/example
```
