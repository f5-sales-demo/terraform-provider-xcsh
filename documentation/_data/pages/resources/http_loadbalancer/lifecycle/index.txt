---
page_title: "Lifecycle"
subcategory: "Load Balancing"
description: "Lifecycle for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": ["http loadbalancer"], "body_bytes": 1474, "body_sha256": "sha256:d784342a954f141cc167fe56a1318e4c8c9e7aafa141411e279e6b8cd27fe0af", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:lifecycle", "parent_id": "xcsh-docs:resources:http_loadbalancer:fundamentals", "path": "documentation/resources/http_loadbalancer/lifecycle/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2201010321200222-2110320331131013-1013211030201130-0130323333132231-1030203101130301-0030312322132221-0233000232333000-3101312320200212", "registry_path": "docs/guides/resources--http_loadbalancer--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "lifecycle", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/lifecycle/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Lifecycle for xcsh_http_loadbalancer.", "tasks": ["lifecycle"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Lifecycle

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- Lifecycle

Changing the loadbalancer_type selection between
[http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/http/),
[https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/),
[https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/),
including an omitted selection, requires recreation and may interrupt service.
F5 Distributed Cloud cannot change this type selection in place.
Certificate rotation and other supported settings within the same selected type remain updates.

Terraform identifies the affected type blocks as **must be replaced**. Unknown block presence requires replacement when unchanged selection cannot be proven; unknown child settings alone do not. Recommendations do not establish API defaults.

Use `lifecycle { prevent_destroy = true }` to reject replacement before remote writes. Terraform controls replacement ordering: its default destroys before creating; `create_before_destroy` requests creation first, which may fail if XC requires a unique name. Plan a maintenance window or use a distinct name for a staged replacement.
