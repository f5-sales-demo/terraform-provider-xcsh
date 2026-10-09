---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_route."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 983, "body_sha256": "sha256:57ccd9abfc5fdec1de3871b0fe1ab97ea19c3e466c4a46179e4310d9c5d0b734", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f", "source_path": "examples/data-sources/xcsh_route/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:route:example:data-source", "parent_id": "xcsh-docs:data-sources:route:examples", "path": "documentation/data-sources/route/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2111001110122330-3313200121202320-1032123300123212-3030231013303020-2311021122201333-1201002113331113-2200322112201000-2203220100210013", "registry_path": "docs/guides/data-sources--route--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["routeCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_route/data-source.tf`; digest `sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f`.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```
