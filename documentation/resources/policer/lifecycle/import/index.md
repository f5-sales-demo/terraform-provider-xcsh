---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_policer."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 346, "body_sha256": "sha256:4e41d449b2be85ef1527694e9fb9f78a0ab130720bf1abfc5bb00a2003bb01c2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:policer:fundamentals", "path": "documentation/resources/policer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1203212110220131-0233233101122320-0213000303323231-0320300032332300-1100131122030312-3033202030310222-0302021222001310-2322311101201010", "registry_path": "docs/guides/resources--policer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_policer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_policer.example system/example
```
