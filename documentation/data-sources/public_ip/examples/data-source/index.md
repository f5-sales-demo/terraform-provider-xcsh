---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_public_ip."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1021, "body_sha256": "sha256:8ee59332c1bd403f055d2923473c9357ae479b784afb07e4436e1bc9b6dd6b95", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3249669279b5070f3336fdc2ea70d7926ac4af4c76c377a1b4478e268c3c5d4d", "source_path": "examples/data-sources/xcsh_public_ip/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:public_ip:example:data-source", "parent_id": "xcsh-docs:data-sources:public_ip:examples", "path": "documentation/data-sources/public_ip/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "public_ip", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2222001222003213-2122001223202112-0013211320323330-2122030312130011-1020202330211320-2122303320102033-2022122111200311-0013303122122100", "registry_path": "docs/guides/data-sources--public_ip--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_public_ip.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_public_ip/data-source.tf`; digest `sha256:3249669279b5070f3336fdc2ea70d7926ac4af4c76c377a1b4478e268c3c5d4d`.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```
