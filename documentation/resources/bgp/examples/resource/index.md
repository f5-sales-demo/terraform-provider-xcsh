---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1201, "body_sha256": "sha256:a09a12fea9bcd6b611a8be835557eab55f615756b94909c13c2fa2896ffe7dc7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7d85805409033bf09b03057cbf852697d9f9c443321c552ee56fb1009994fdd5", "source_path": "examples/resources/xcsh_bgp/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp:example:resource", "parent_id": "xcsh-docs:resources:bgp:examples", "path": "documentation/resources/bgp/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2310210011323101-1213010202332001-0300113302132302-2013031322133103-3330212310300001-3122322121201332-1203200131333113-1313200031001223", "registry_path": "docs/guides/resources--bgp--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_bgp.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp/resource.tf`; digest `sha256:7d85805409033bf09b03057cbf852697d9f9c443321c552ee56fb1009994fdd5`.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/examples/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
