---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": ["k8s pod security admission"], "body_bytes": 403, "body_sha256": "sha256:55ecd58e178d1b23fa7eedd28fbd8f1362135c7dffe62d192d6c82b406d9fad4", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:k8s_pod_security_admission:fundamentals", "path": "documentation/resources/k8s_pod_security_admission/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1332322211301202-3112300032013110-3333012330223211-2030001100132313-1030222133313333-3012033302200111-0012121021232011-3133300211112002", "registry_path": "docs/guides/resources--k8s_pod_security_admission--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_k8s_pod_security_admission.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_admission/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_k8s_pod_security_admission.example system/example
```
