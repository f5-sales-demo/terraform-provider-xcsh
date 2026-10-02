---
page_title: "xcsh_network_policy"
subcategory: "Security"
description: "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["network policy"], "body_bytes": 1402, "body_sha256": "sha256:982ff3d79a7e1f92dc5d0f32c0b261953bad6653e2afb0f1f8ab534e19c76d23", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy:reference", "xcsh-docs:data-sources:network_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_policy/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033", "registry_path": "docs/data-sources/network_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_policy

Breadcrumbs:

- xcsh_network_policy

Manages new network policy with configured parameters in specified namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicy by name
data "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}

output "network_policy_id" {
  value = data.xcsh_network_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/examples/)
