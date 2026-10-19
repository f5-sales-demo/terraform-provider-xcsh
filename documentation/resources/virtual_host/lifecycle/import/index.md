---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_virtual_host."
xcsh_docs: {"aliases": ["virtual host"], "body_bytes": 361, "body_sha256": "sha256:298f4647170688e9846e38ac118d82b86585c20abcd45a31f39f6be4452b0cb3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:virtual_host:fundamentals", "path": "documentation/resources/virtual_host/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3010122311030023-3301131312110032-1003122023332231-3031220211222311-0010113103131003-3110200330233031-1023022301030303-3300223123110301", "registry_path": "docs/guides/resources--virtual_host--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_virtual_host.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_virtual_host.example system/example
```
