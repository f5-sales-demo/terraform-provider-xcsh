---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["trusted ca list"], "body_bytes": 370, "body_sha256": "sha256:c523ffed6263fd46b6da6d24b4d6098d51be5a41c362a0277450c0f7549f75dd", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "id": "xcsh-docs:resources:trusted_ca_list:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:trusted_ca_list:fundamentals", "path": "documentation/resources/trusted_ca_list/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2100031021303333-2312102102232223-1130221210320233-0033003021231131-1111033121121320-1122131302001000-3121332113222202-3222212211031113", "registry_path": "docs/guides/resources--trusted_ca_list--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_trusted_ca_list.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_trusted_ca_list.example system/example
```
