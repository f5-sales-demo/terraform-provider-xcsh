---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_crl."
xcsh_docs: {"aliases": ["crl"], "body_bytes": 334, "body_sha256": "sha256:b0b993a699d1c8c90d96ee4a848a86fb1a5acc9690477b013d71d57927bb0f7b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "id": "xcsh-docs:resources:crl:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:crl:fundamentals", "path": "documentation/resources/crl/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3311000203101232-2031203322230023-1030231121212302-1002211312332133-0101301200101300-0011223301313212-1032313213320100-1103210210212030", "registry_path": "docs/guides/resources--crl--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_crl.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_crl.example system/example
```
