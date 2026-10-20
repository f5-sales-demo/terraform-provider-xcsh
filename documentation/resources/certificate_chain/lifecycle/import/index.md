---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_certificate_chain."
xcsh_docs: {"aliases": ["cert", "certificate", "certificate chain", "existing certificates", "tls certificates"], "body_bytes": 376, "body_sha256": "sha256:43784253fabba6e6b4b3f1d04952ed888df757689a4354042ac7b972ddeea1ef", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate_chain:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate_chain:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:certificate_chain:fundamentals", "path": "documentation/resources/certificate_chain/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1331002123100010-3112000330321013-2321013022131223-0103323211201210-1000130023312232-3002103200202023-3202203202221023-2101203311102111", "registry_path": "docs/guides/resources--certificate_chain--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate_chain/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Import for xcsh_certificate_chain.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_certificate_chain.example system/example
```
