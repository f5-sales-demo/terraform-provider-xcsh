---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_policer."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 346, "body_sha256": "sha256:4e41d449b2be85ef1527694e9fb9f78a0ab130720bf1abfc5bb00a2003bb01c2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:policer:fundamentals", "path": "documentation/resources/policer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1203212110220131-0233233101122320-0213000303323231-0320300032332300-1100131122030312-3033202030310222-0302021222001310-2322311101201010", "registry_path": "docs/guides/resources--policer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Import for xcsh_policer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["policerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
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
