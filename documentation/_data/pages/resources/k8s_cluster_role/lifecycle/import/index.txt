---
page_title: "Import"
subcategory: "Container"
description: "Import for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": ["k8s cluster role"], "body_bytes": 373, "body_sha256": "sha256:aa40806b01638be2436a42d04c1ef015586d17f1699ff8bc275e613da6a10753", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:k8s_cluster_role:fundamentals", "path": "documentation/resources/k8s_cluster_role/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1311100021231112-1103201000203120-1230302223123300-2033030131200232-0101221011113101-2030130202312213-1323300030210202-2230301011323023", "registry_path": "docs/guides/resources--k8s_cluster_role--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_k8s_cluster_role.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_k8s_cluster_role.example system/example
```
