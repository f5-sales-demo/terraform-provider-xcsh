---
page_title: "Import"
subcategory: "Identity"
description: "Import for xcsh_token."
xcsh_docs: {"aliases": ["token"], "body_bytes": 340, "body_sha256": "sha256:2e9b5a139a56a1d56a84eb1942dd1bdace88af25004ab7de789dd5958159c2fb", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "id": "xcsh-docs:resources:token:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:token:fundamentals", "path": "documentation/resources/token/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3131222232031130-1102111221001201-1310323230133202-1310100223300000-2132132202210232-0122021232113323-2003233122130213-1021331220333332", "registry_path": "docs/guides/resources--token--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/lifecycle/import/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Import for xcsh_token.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["tokenCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_token.example system/example
```
