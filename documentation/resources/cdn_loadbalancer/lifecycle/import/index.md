---
page_title: "Import"
subcategory: "Load Balancing"
description: "Import for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": ["cdn loadbalancer"], "body_bytes": 373, "body_sha256": "sha256:8cf29e3f36526ba8a59c0c66c4c5c7b2021a6b1b241c338484ba953d50f8bc9e", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:fundamentals", "path": "documentation/resources/cdn_loadbalancer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0113203320131311-2001120113302103-1202220232210031-1210210002111130-1222002311032311-2200223111013300-1112121120212322-0211212113231331", "registry_path": "docs/guides/resources--cdn_loadbalancer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_cdn_loadbalancer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_cdn_loadbalancer.example system/example
```
