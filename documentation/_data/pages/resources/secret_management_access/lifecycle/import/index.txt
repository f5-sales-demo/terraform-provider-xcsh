---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["secret management access"], "body_bytes": 397, "body_sha256": "sha256:7ff84de8d4949da38128d53e7cd971ee3da82b02efd47a649cf1138f3242a767", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:secret_management_access:fundamentals", "path": "documentation/resources/secret_management_access/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1311023322202000-3031202311101120-0012032203312013-3302002203322322-0331232113331102-1310131321133112-2030200011222103-3101131100333230", "registry_path": "docs/guides/resources--secret_management_access--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_secret_management_access.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_secret_management_access.example system/example
```
