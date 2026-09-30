---
page_title: "xcsh_route"
subcategory: ""
description: "xcsh_route for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1197, "body_sha256": "sha256:451f8587c6df5fc28450ac6cde50ecdedf27ddfac1e3971d21cfaeda7d8a2a01", "canonical_id": "xcsh-docs:data-sources:route:fundamentals", "child_ids": ["xcsh-docs:data-sources:route:reference", "xcsh-docs:data-sources:route:examples"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:fundamentals", "parent_id": null, "path": "docs/data-sources/route.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_route for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--route--reference.md)
- [Examples](../guides/data-sources--route--examples.md)
