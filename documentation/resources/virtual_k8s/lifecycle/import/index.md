---
page_title: "Import"
subcategory: "Container"
description: "Import for xcsh_virtual_k8s."
xcsh_docs: {"aliases": ["virtual k8s"], "body_bytes": 358, "body_sha256": "sha256:943732b2f27a8ef568dbdca58e96c51dff9f30b8793f34668fc14dcd4624d998", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:virtual_k8s:fundamentals", "path": "documentation/resources/virtual_k8s/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0210330301020310-1301021323100000-2131300013132312-3320113213110132-2212310131312000-1210333120230202-1012303221101223-0311213103022233", "registry_path": "docs/guides/resources--virtual_k8s--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_virtual_k8s.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_virtual_k8s.example system/example
```
