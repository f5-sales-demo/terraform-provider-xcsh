---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_authentication."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 367, "body_sha256": "sha256:cbd4ecc18e3b20563c28fae378db117a449096712e3c89df2b86f593eba291fc", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:authentication:fundamentals", "path": "documentation/resources/authentication/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3232010031111121-3102203230033011-2100313102221210-3023201111100300-2023031020230000-2220221213123332-1211111121312203-1213230233000232", "registry_path": "docs/guides/resources--authentication--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/lifecycle/import/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Import for xcsh_authentication.", "tasks": ["authentication", "import"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["authenticationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_authentication.example system/example
```
