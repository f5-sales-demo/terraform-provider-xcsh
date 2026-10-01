---
page_title: "xcsh_network_policy_view"
subcategory: ""
description: "xcsh_network_policy_view for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1656, "body_sha256": "sha256:3e7c16faa77ccf34fc3ebdd28e037b60c11138478d81706cda0074bd3f7b6a89", "child_ids": ["xcsh-docs:resources:network_policy_view:reference", "xcsh-docs:resources:network_policy_view:examples", "xcsh-docs:resources:network_policy_view:import", "xcsh-docs:resources:network_policy_view:timeouts"], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:fundamentals", "parent_id": null, "path": "documentation/resources/network_policy_view/index.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_policy_view for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_policy_view

Breadcrumbs:

- xcsh_network_policy_view

Manages a Network Policy View resource in F5 Distributed Cloud for network policy view
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyView Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyView configuration
resource "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/lifecycle/timeouts/)
