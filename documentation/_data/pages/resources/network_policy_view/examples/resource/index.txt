---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_policy_view."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1352, "body_sha256": "sha256:66f8587997c89a2c3d4c5aa9c31b75a5e1774c4304cd1b39117d6f25b37da8c5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320", "source_path": "examples/resources/xcsh_network_policy_view/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy_view:example:resource", "parent_id": "xcsh-docs:resources:network_policy_view:examples", "path": "documentation/resources/network_policy_view/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3221010333023301-0110103133312212-2033330230320132-2233213323033311-0012103003303123-3212202312020311-3112130203033032-3102033103110332", "registry_path": "docs/guides/resources--network_policy_view--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_policy_view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
