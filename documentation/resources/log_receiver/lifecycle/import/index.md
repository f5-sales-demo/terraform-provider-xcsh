---
page_title: "Import"
subcategory: "Monitoring"
description: "Import for xcsh_log_receiver."
xcsh_docs: {"aliases": ["log receiver"], "body_bytes": 361, "body_sha256": "sha256:de737d20de4bb5689932326ef88a040708cbb225d99bd3627fa3c1957aa002bc", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:log_receiver:fundamentals", "path": "documentation/resources/log_receiver/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3123223113211002-2122000132323322-0202030313031023-2331001321000232-0120013032201311-1230233312003002-0122233323122202-3221300302102201", "registry_path": "docs/guides/resources--log_receiver--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/lifecycle/import/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Import for xcsh_log_receiver.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_log_receiver.example system/example
```
