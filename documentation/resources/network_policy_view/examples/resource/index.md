---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:b69a95964f852d1ab5463a36dab8ec4dfde16e2e9b39dd73b33ced26790cd313", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320", "source_path": "examples/resources/xcsh_network_policy_view/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy_view:example:resource", "parent_id": "xcsh-docs:resources:network_policy_view:examples", "path": "documentation/resources/network_policy_view/examples/resource/index.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_view/resource.tf`; digest `sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/examples/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
