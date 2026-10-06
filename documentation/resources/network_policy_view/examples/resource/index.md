---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_policy_view."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1109, "body_sha256": "sha256:1eee369d650def7cac421a9bdfea8a61d13d55c79d254338ce16b123bf51d231", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320", "source_path": "examples/resources/xcsh_network_policy_view/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy_view:example:resource", "parent_id": "xcsh-docs:resources:network_policy_view:examples", "path": "documentation/resources/network_policy_view/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3221010333023301-0110103133312212-2033330230320132-2233213323033311-0012103003303123-3212202312020311-3112130203033032-3102033103110332", "registry_path": "docs/guides/resources--network_policy_view--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_network_policy_view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
