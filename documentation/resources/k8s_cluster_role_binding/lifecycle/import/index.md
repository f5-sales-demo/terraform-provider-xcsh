---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": ["k8s cluster role binding"], "body_bytes": 397, "body_sha256": "sha256:f22f82dbb33cf0ab391e522f31b0f46174af95f02923b90d863bb13830471ea7", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:k8s_cluster_role_binding:fundamentals", "path": "documentation/resources/k8s_cluster_role_binding/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2200110211313011-0101011230210131-3310310000320012-3302230000112323-3223020122121020-3013120000003322-3213003213103332-1010101321132211", "registry_path": "docs/guides/resources--k8s_cluster_role_binding--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/lifecycle/import/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Import for xcsh_k8s_cluster_role_binding.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_k8s_cluster_role_binding.example system/example
```
