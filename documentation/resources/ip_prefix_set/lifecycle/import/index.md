---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": ["ip prefix set"], "body_bytes": 364, "body_sha256": "sha256:e0111f92d737f1018621aa6b44a91cc766a0aeecc205da97be79e30e539eacec", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:ip_prefix_set:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:ip_prefix_set:fundamentals", "path": "documentation/resources/ip_prefix_set/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0213010321223311-2321231203103123-1011313301132112-1021202301302033-1002123211303101-1100313122301010-1230020032132330-3320212231030012", "registry_path": "docs/guides/resources--ip_prefix_set--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_ip_prefix_set.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
