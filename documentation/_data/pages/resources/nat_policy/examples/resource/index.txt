---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nat_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1288, "body_sha256": "sha256:f5c263912436794a60348b853b2728a3864386c9e4e10663d43f7b378736a37f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08", "source_path": "examples/resources/xcsh_nat_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nat_policy:example:resource", "parent_id": "xcsh-docs:resources:nat_policy:examples", "path": "documentation/resources/nat_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2302320100223323-1103033101202233-2011332102122012-2321313122020033-1303211313033103-1112322132223331-1321113120223303-0002311000301120", "registry_path": "docs/guides/resources--nat_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_nat_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nat_policy/resource.tf`; digest `sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08`.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/examples/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
