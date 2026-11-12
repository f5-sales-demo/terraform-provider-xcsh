---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_cloud_user_account."
xcsh_docs: {"aliases": ["cloud user account"], "body_bytes": 379, "body_sha256": "sha256:f3d7c697a99ce27e9dd5eb72c6e605d77742b5ba2656ce842fc4e646373d5939", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:cloud_user_account:fundamentals", "path": "documentation/resources/cloud_user_account/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3201120320023322-0133231032231320-1222233130132021-2211030001331102-3101023332001003-2022030322301001-2200000132022010-1133230210131223", "registry_path": "docs/guides/resources--cloud_user_account--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_cloud_user_account.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_cloud_user_account.example system/example
```
