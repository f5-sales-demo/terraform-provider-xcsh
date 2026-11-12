---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_k8s_cluster."
xcsh_docs: {"aliases": ["k8s cluster"], "body_bytes": 358, "body_sha256": "sha256:3fd801e8deb4a6f56aab7460766ba6ec56613cea9db051af91a6f02e1c049b0d", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:k8s_cluster:fundamentals", "path": "documentation/resources/k8s_cluster/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2302311103303033-2313331022021113-1133010032311330-2313311203133131-3220203023311112-0011320110212120-0112222032010100-2223332022012120", "registry_path": "docs/guides/resources--k8s_cluster--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_k8s_cluster.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_k8s_cluster.example system/example
```
