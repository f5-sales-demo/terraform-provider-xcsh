---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1006, "body_sha256": "sha256:d514ed5f06431778f485b39e0d904a0be2302f6f86839965698e2abc65ed50f1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7d85805409033bf09b03057cbf852697d9f9c443321c552ee56fb1009994fdd5", "source_path": "examples/resources/xcsh_bgp/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp:example:resource", "parent_id": "xcsh-docs:resources:bgp:examples", "path": "documentation/resources/bgp/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2310210011323101-1213010202332001-0300113302132302-2013031322133103-3330212310300001-3122322121201332-1203200131333113-1313200031001223", "registry_path": "docs/guides/resources--bgp--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_bgp.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
