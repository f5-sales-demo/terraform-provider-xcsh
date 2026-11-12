---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_proxy."
xcsh_docs: {"aliases": ["proxy"], "body_bytes": 340, "body_sha256": "sha256:66ee1fcbf3459571c5629330c1f303a30d9fc65a9aac23caafa1ba2a40e706c2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:proxy:fundamentals", "path": "documentation/resources/proxy/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2233020331221110-3220321131302032-0111012300331101-3323011320113313-0213103313030023-3211320021321122-1013320201322113-2220013020312123", "registry_path": "docs/guides/resources--proxy--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/lifecycle/import/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Import for xcsh_proxy.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["proxyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_proxy.example system/example
```
