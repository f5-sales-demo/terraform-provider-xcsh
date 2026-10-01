---
page_title: "xcsh_forward_proxy_policy"
subcategory: "Security"
description: "xcsh_forward_proxy_policy for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1718, "body_sha256": "sha256:1e6f812af93f1e4c55755e917eb545957b518b159d786138783455219c8b0a2a", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:reference", "xcsh-docs:resources:forward_proxy_policy:examples", "xcsh-docs:resources:forward_proxy_policy:import", "xcsh-docs:resources:forward_proxy_policy:timeouts"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:fundamentals", "parent_id": null, "path": "documentation/resources/forward_proxy_policy/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_forward_proxy_policy for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_forward_proxy_policy

Breadcrumbs:

- xcsh_forward_proxy_policy

Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/lifecycle/timeouts/)
