---
page_title: "Import"
subcategory: "Container"
description: "Import for xcsh_container_registry."
xcsh_docs: {"aliases": ["container registry"], "body_bytes": 379, "body_sha256": "sha256:89fbb9ccc311835d3ea55ffae504413bb067dd833b92c2419686e410f6b0c9fe", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:resources:container_registry:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:container_registry:fundamentals", "path": "documentation/resources/container_registry/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1232313310001120-0013101312033331-2100002211100101-3203130201321101-1203113000232330-0132211032022212-1300223300001310-0320122201310122", "registry_path": "docs/guides/resources--container_registry--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_container_registry.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["container_registryCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_container_registry.example system/example
```
