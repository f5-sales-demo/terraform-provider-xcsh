---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_tunnel."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 343, "body_sha256": "sha256:73e2301b31de8b8f3b4c35a4032d9eb8fad066afe665db54f6a64d98739da0d3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:tunnel:fundamentals", "path": "documentation/resources/tunnel/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1010131233201320-3203100203112233-3100302112033223-0210030201322003-2330312321130221-1233203330313332-3012322103303101-1330312132300010", "registry_path": "docs/guides/resources--tunnel--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_tunnel.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_tunnel.example system/example
```
