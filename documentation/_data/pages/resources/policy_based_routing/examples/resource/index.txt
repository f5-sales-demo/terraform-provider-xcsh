---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_policy_based_routing."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1381, "body_sha256": "sha256:fdf33dc197d18e625f99f759df0a828ad1b0ee9ebca617c9ffc2c1b926e28ea8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968", "source_path": "examples/resources/xcsh_policy_based_routing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:policy_based_routing:example:resource", "parent_id": "xcsh-docs:resources:policy_based_routing:examples", "path": "documentation/resources/policy_based_routing/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3310100020110201-2313301030311313-2123101220300301-3311301012232032-0302121032101013-3003300102331202-1220111310000022-2200130313332003", "registry_path": "docs/guides/resources--policy_based_routing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_policy_based_routing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policy_based_routing/resource.tf`; digest `sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968`.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/examples/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
