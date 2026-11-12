---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": ["dns lb pool"], "body_bytes": 358, "body_sha256": "sha256:711ddc6a078de42274bc3eebb6c6e16e0604baf046ddd09713894c8b0b327deb", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:dns_lb_pool:fundamentals", "path": "documentation/resources/dns_lb_pool/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1002022001013301-1011001213213310-1330111101332012-2031213323300312-1300202313100221-1203332221113112-2201022010133213-1003022230011103", "registry_path": "docs/guides/resources--dns_lb_pool--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_dns_lb_pool.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_dns_lb_pool.example system/example
```
