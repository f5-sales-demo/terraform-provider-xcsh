---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_route."
xcsh_docs: {"aliases": ["route"], "body_bytes": 340, "body_sha256": "sha256:cc512abc068805c4d62021c60c10adba04d44b3284ad02e53a9894545956dae1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:route:fundamentals", "path": "documentation/resources/route/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1001200300032303-3102230123232020-1333122123132203-0301001023230121-1023222222113310-1130203222113323-1233012033130310-1032112013100301", "registry_path": "docs/guides/resources--route--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Import for xcsh_route.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["routeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_route.example system/example
```
