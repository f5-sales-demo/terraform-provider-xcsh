---
page_title: "Import"
subcategory: "DNS"
description: "Import for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": ["dns load balancer"], "body_bytes": 376, "body_sha256": "sha256:b2b8c137050e33e6417742521b3bb1e567f5d8d70301ab2b278c9347d084f850", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "path": "documentation/resources/dns_load_balancer/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0130210010103222-1210113203203231-3123131212201010-0212231011303020-0311103213211033-1102302312321331-2031301213323123-1323130320030110", "registry_path": "docs/guides/resources--dns_load_balancer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/lifecycle/import/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Import for xcsh_dns_load_balancer.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_dns_load_balancer.example system/example
```
