---
page_title: "Import"
subcategory: "Load Balancing"
description: "Import for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 376, "body_sha256": "sha256:ef6441d7cebd5c2ab019ef283f053d7ff41699bf91c21f318595c446c800b04d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:http_loadbalancer:fundamentals", "path": "documentation/resources/http_loadbalancer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1000101212023023-0103011103132001-3301102330230310-1230010102120212-3223112021023333-2200311133300111-1031101210100332-2211331001010311", "registry_path": "docs/guides/resources--http_loadbalancer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Import for xcsh_http_loadbalancer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_http_loadbalancer.example system/example
```
