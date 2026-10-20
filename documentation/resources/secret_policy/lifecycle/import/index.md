---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_secret_policy."
xcsh_docs: {"aliases": ["secret policy"], "body_bytes": 364, "body_sha256": "sha256:569aff61a7b894c4b8dbdb08beeea85ee4829d61c6cd49afcc16bb57ec4cb451", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_policy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:secret_policy:fundamentals", "path": "documentation/resources/secret_policy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2111131021220031-2323002001201332-3203310300212013-2033202133131221-2323022112202023-3031312211011310-1303310232320032-3332003112313022", "registry_path": "docs/guides/resources--secret_policy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_secret_policy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["secret_policyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_secret_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_secret_policy.example system/example
```
