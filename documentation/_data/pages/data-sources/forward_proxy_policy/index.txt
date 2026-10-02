---
page_title: "xcsh_forward_proxy_policy"
subcategory: "Security"
description: "Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification. configuration."
xcsh_docs: {"aliases": ["forward proxy policy"], "body_bytes": 1477, "body_sha256": "sha256:dae95c8ee5743442bfc91c79c797c1149e0e38a6b8fdb9f122719e45f783505a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:reference", "xcsh-docs:data-sources:forward_proxy_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/forward_proxy_policy/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203", "registry_path": "docs/data-sources/forward_proxy_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
