---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_service_policy."
xcsh_docs: {"aliases": ["service policy"], "body_bytes": 367, "body_sha256": "sha256:6c0e5914399c2a4240224fbc742791cc24f66539a34811e548c7b33a9710627f", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:service_policy:fundamentals", "path": "documentation/resources/service_policy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3001210010320130-2032113022230020-3233101011020330-3120033130332021-2132221202012233-0010012202130213-0110100100330330-1301210030303010", "registry_path": "docs/guides/resources--service_policy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_service_policy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["service_policyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_service_policy.example system/example
```
