---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1282, "body_sha256": "sha256:4bb91c6512789e7ad3a8a3725b90d9420491215cd05eb91a809ab0f37443578c", "child_ids": [], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3fff12bb997c990a591800c80ace62a7790c01337e27ba68116856af1dfaf950", "source_path": "examples/data-sources/xcsh_forward_proxy_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:forward_proxy_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:examples", "path": "documentation/data-sources/forward_proxy_policy/examples/data-source/index.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forward_proxy_policy/data-source.tf`; digest `sha256:3fff12bb997c990a591800c80ace62a7790c01337e27ba68116856af1dfaf950`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/examples/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
