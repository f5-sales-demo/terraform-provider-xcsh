---
page_title: "Import"
subcategory: "Container"
description: "Import for xcsh_workload."
xcsh_docs: {"aliases": ["workload"], "body_bytes": 349, "body_sha256": "sha256:dc9fb83a908bc83a600f193cff1f6d8b2a314a93681c6fb0df2e7431a10f4d46", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:workload:fundamentals", "path": "documentation/resources/workload/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3211222301121232-1020210312011003-0231102321331130-1110200130120010-3131300210013301-2331311113011020-3211110001033120-2021101202123230", "registry_path": "docs/guides/resources--workload--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_workload.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["workloadCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_workload.example system/example
```
