---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_api_discovery."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1061, "body_sha256": "sha256:94d8e0b4ca6e56e316e3b8fb20183bfa79333923d5b273b0c061b521d8cdb189", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f37cea7bc8746642eef8cfd7a0d2f33975184479df2f41465a9bdda95de57bb2", "source_path": "examples/data-sources/xcsh_api_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:api_discovery:examples", "path": "documentation/data-sources/api_discovery/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-1020010203113302-2202232221321101-2020332001232022-3210013310130031-1120320311223300-3012133011202200-0132102332301033-2320210212012310", "registry_path": "docs/guides/data-sources--api_discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Data source for xcsh_api_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_discovery/data-source.tf`; digest `sha256:f37cea7bc8746642eef8cfd7a0d2f33975184479df2f41465a9bdda95de57bb2`.

```terraform
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```
