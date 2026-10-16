---
page_title: "Import"
subcategory: "DNS"
description: "Import for xcsh_dns_zone."
xcsh_docs: {"aliases": ["dns zone"], "body_bytes": 349, "body_sha256": "sha256:7580d3c047cd84f0dfd2f8b25bf34f5e0be72779e31f2330d1b6f2caf21c1bc2", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:dns_zone:fundamentals", "path": "documentation/resources/dns_zone/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0030002023313301-2322000302203010-0323211321011000-2312301301301013-2012312311321222-3323013232021001-0322012302312020-1231013222302331", "registry_path": "docs/guides/resources--dns_zone--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/lifecycle/import/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Import for xcsh_dns_zone.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_dns_zone.example system/example
```
