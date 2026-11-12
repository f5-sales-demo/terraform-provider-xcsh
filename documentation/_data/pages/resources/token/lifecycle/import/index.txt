---
page_title: "Import"
subcategory: "Identity"
description: "Import for xcsh_token."
xcsh_docs: {"aliases": ["token"], "body_bytes": 340, "body_sha256": "sha256:2e9b5a139a56a1d56a84eb1942dd1bdace88af25004ab7de789dd5958159c2fb", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "id": "xcsh-docs:resources:token:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:token:fundamentals", "path": "documentation/resources/token/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3131222232031130-1102111221001201-1310323230133202-1310100223300000-2132132202210232-0122021232113323-2003233122130213-1021331220333332", "registry_path": "docs/guides/resources--token--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_token.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["tokenCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
