---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_authorization_server."
xcsh_docs: {"aliases": ["authorization server"], "body_bytes": 385, "body_sha256": "sha256:23a473100b7e265618624e261f6e39a9aa77c897ba15977ad84c72592be635c9", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:resources:authorization_server:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:authorization_server:fundamentals", "path": "documentation/resources/authorization_server/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "authorization_server", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3311130103213132-0210202300312100-3032323010000201-2311222100103221-0100311121130303-1011113132100033-2320012131020013-1013321200310133", "registry_path": "docs/guides/resources--authorization_server--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authorization_server/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Import for xcsh_authorization_server.", "tasks": ["authentication", "import"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_authorization_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authorization_server/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_authorization_server.example system/example
```
