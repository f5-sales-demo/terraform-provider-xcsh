---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": ["ip prefix set"], "body_bytes": 364, "body_sha256": "sha256:e0111f92d737f1018621aa6b44a91cc766a0aeecc205da97be79e30e539eacec", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:ip_prefix_set:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:ip_prefix_set:fundamentals", "path": "documentation/resources/ip_prefix_set/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0213010321223311-2321231203103123-1011313301132112-1021202301302033-1002123211303101-1100313122301010-1230020032132330-3320212231030012", "registry_path": "docs/guides/resources--ip_prefix_set--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/lifecycle/import/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Import for xcsh_ip_prefix_set.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_ip_prefix_set.example system/example
```
