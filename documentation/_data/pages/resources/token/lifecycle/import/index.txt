---
page_title: "Import"
subcategory: "Identity"
description: "Import for xcsh_token."
xcsh_docs: {"aliases": ["token"], "body_bytes": 340, "body_sha256": "sha256:2e9b5a139a56a1d56a84eb1942dd1bdace88af25004ab7de789dd5958159c2fb", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "id": "xcsh-docs:resources:token:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:token:fundamentals", "path": "documentation/resources/token/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3131222232031130-1102111221001201-1310323230133202-1310100223300000-2132132202210232-0122021232113323-2003233122130213-1021331220333332", "registry_path": "docs/guides/resources--token--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_token.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["tokenCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
