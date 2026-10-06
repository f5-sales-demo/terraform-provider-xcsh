---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_certificate."
xcsh_docs: {"aliases": ["certificate"], "body_bytes": 358, "body_sha256": "sha256:d363ac05107cf6e845192eb5e51a5ce6940c811e55e2f9ddd4720f65aac23333", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:certificate:fundamentals", "path": "documentation/resources/certificate/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1101233031133031-3130000233121232-3102010323131202-3101303002132031-0012011222011220-2110202303020302-1303110123110020-3321102223002210", "registry_path": "docs/guides/resources--certificate--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_certificate.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_certificate.example system/example
```
