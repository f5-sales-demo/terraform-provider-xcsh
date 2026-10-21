---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_protected_application."
xcsh_docs: {"aliases": ["protected application"], "body_bytes": 388, "body_sha256": "sha256:fa29caa275275d984658ed357d4a61439f1159d9b2a4ff531f83f7ef7716005b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:protected_application:fundamentals", "path": "documentation/resources/protected_application/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0112302200331003-2330231312231032-1202220233011001-2023010321003301-0312222323010120-2000220100001223-0223021100133221-2331002221210112", "registry_path": "docs/guides/resources--protected_application--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_protected_application.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
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
