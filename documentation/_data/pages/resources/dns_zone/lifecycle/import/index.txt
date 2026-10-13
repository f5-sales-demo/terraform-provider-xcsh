---
page_title: "Import"
subcategory: "DNS"
description: "Import for xcsh_dns_zone."
xcsh_docs: {"aliases": ["dns zone"], "body_bytes": 349, "body_sha256": "sha256:7580d3c047cd84f0dfd2f8b25bf34f5e0be72779e31f2330d1b6f2caf21c1bc2", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:dns_zone:fundamentals", "path": "documentation/resources/dns_zone/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "resources", "registry_anchor": "canonical-0030002023313301-2322000302203010-0323211321011000-2312301301301013-2012312311321222-3323013232021001-0322012302312020-1231013222302331", "registry_path": "docs/guides/resources--dns_zone--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/lifecycle/import/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Import for xcsh_dns_zone.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
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
