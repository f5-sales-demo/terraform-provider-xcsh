---
page_title: "xcsh_forward_proxy_policy"
subcategory: "Security"
description: "xcsh_forward_proxy_policy for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1378, "body_sha256": "sha256:4eadd45ecfc2a473f796e8fe22f7f498de005c622358c9d6f4c49f64bb0f0c46", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:reference", "xcsh-docs:data-sources:forward_proxy_policy:examples"], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:fundamentals", "parent_id": null, "path": "documentation/data-sources/forward_proxy_policy/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_forward_proxy_policy for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# ForwardProxyPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardProxyPolicy by name
data "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}

output "forward_proxy_policy_id" {
  value = data.xcsh_forward_proxy_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/examples/)
