---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_protected_application."
xcsh_docs: {"aliases": ["protected application"], "body_bytes": 388, "body_sha256": "sha256:fa29caa275275d984658ed357d4a61439f1159d9b2a4ff531f83f7ef7716005b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:protected_application:fundamentals", "path": "documentation/resources/protected_application/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0112302200331003-2330231312231032-1202220233011001-2023010321003301-0312222323010120-2000220100001223-0223021100133221-2331002221210112", "registry_path": "docs/guides/resources--protected_application--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_protected_application.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_protected_application.example system/example
```
