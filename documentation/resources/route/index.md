---
page_title: "xcsh_route"
subcategory: ""
description: "xcsh_route for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1561, "body_sha256": "sha256:f86c083ded35991e3dc4b9ddc9d7fae0060d83a705757088f206d1e52e76ce82", "child_ids": ["xcsh-docs:resources:route:reference", "xcsh-docs:resources:route:examples", "xcsh-docs:resources:route:import", "xcsh-docs:resources:route:timeouts"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:fundamentals", "parent_id": null, "path": "documentation/resources/route/index.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_route for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_route

Breadcrumbs:

- xcsh_route

Manages route object in a given namespace. Route object is list of route rules. Each rule has match
condition to match incoming requests and actions to take on matching requests in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/lifecycle/timeouts/)
