---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_cloud_user_account."
xcsh_docs: {"aliases": ["cloud user account"], "body_bytes": 379, "body_sha256": "sha256:f3d7c697a99ce27e9dd5eb72c6e605d77742b5ba2656ce842fc4e646373d5939", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:cloud_user_account:fundamentals", "path": "documentation/resources/cloud_user_account/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3201120320023322-0133231032231320-1222233130132021-2211030001331102-3101023332001003-2022030322301001-2200000132022010-1133230210131223", "registry_path": "docs/guides/resources--cloud_user_account--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_cloud_user_account.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
