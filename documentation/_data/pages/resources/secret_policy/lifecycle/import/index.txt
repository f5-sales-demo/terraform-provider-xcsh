---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_secret_policy."
xcsh_docs: {"aliases": ["secret policy"], "body_bytes": 364, "body_sha256": "sha256:569aff61a7b894c4b8dbdb08beeea85ee4829d61c6cd49afcc16bb57ec4cb451", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_policy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:secret_policy:fundamentals", "path": "documentation/resources/secret_policy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2111131021220031-2323002001201332-3203310300212013-2033202133131221-2323022112202023-3031312211011310-1303310232320032-3332003112313022", "registry_path": "docs/guides/resources--secret_policy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_secret_policy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["secret_policyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
