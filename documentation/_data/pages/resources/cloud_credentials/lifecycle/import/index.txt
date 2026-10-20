---
page_title: "Import"
subcategory: "Infrastructure"
description: "Import for xcsh_cloud_credentials."
xcsh_docs: {"aliases": ["authentication", "cloud credentials", "credential setup", "credentials"], "body_bytes": 376, "body_sha256": "sha256:bff855ccd8534ef3cf26bcabe2baabb4cfede692e130ed4a23072b48ff0591c1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:cloud_credentials:fundamentals", "path": "documentation/resources/cloud_credentials/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1100112101112012-0311100211012110-3221001133101211-2122001021022023-0232233203111021-3323211210312032-0320303120222020-0322212111313330", "registry_path": "docs/guides/resources--cloud_credentials--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_cloud_credentials.", "tasks": ["authentication", "import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_cloud_credentials.example system/example
```
