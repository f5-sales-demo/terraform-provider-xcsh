---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_virtual_host."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1051, "body_sha256": "sha256:05d8c71e57307b9f4ce1a12f8bbac441bd4c9f0591612d16bbd3a766b45e6e86", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72", "source_path": "examples/data-sources/xcsh_virtual_host/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_host:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_host:examples", "path": "documentation/data-sources/virtual_host/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3310330001110212-1232132132312200-0232230121322300-3003102003212202-3022012012310032-2323033202222022-2133023021331231-1003112013212303", "registry_path": "docs/guides/data-sources--virtual_host--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_virtual_host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_host/data-source.tf`; digest `sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72`.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```
