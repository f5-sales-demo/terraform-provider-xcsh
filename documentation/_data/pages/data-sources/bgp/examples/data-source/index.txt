---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bgp."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 963, "body_sha256": "sha256:62436dfb20d970b0c0df475cf5604d61b14d159b94bb50a2da525701ccddf1f3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b", "source_path": "examples/data-sources/xcsh_bgp/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bgp:example:data-source", "parent_id": "xcsh-docs:data-sources:bgp:examples", "path": "documentation/data-sources/bgp/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0321223223233121-0312330200233031-1130101030022030-2323111212200121-0203112130013111-3111001012023220-2322313032020103-0201102003110302", "registry_path": "docs/guides/data-sources--bgp--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_bgp.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["bgpCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp/data-source.tf`; digest `sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b`.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```
