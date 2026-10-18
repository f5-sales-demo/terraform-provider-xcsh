---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_workload_flavor."
xcsh_docs: {"aliases": ["workload flavor"], "body_bytes": 370, "body_sha256": "sha256:d11d192475aafcdcbf34a3251515c627f37a4512439cb382d9e602c99c621f56", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:workload_flavor:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload_flavor:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:workload_flavor:fundamentals", "path": "documentation/resources/workload_flavor/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2323021122210002-2212110111100100-0233321233232203-1322003201220002-3323130031221320-1012200000333031-0303302222301133-2111222131130223", "registry_path": "docs/guides/resources--workload_flavor--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload_flavor/lifecycle/import/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Import for xcsh_workload_flavor.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_workload_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload_flavor/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_workload_flavor.example system/example
```
