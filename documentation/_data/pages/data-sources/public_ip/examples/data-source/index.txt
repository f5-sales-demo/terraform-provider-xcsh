---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_public_ip."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1240, "body_sha256": "sha256:87871bd73e9869eae3a8a66ee3c2295994bb2e926e31d79b1fdac91c687c86da", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3249669279b5070f3336fdc2ea70d7926ac4af4c76c377a1b4478e268c3c5d4d", "source_path": "examples/data-sources/xcsh_public_ip/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:public_ip:example:data-source", "parent_id": "xcsh-docs:data-sources:public_ip:examples", "path": "documentation/data-sources/public_ip/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "public_ip", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2222001222003213-2122001223202112-0013211320323330-2122030312130011-1020202330211320-2122303320102033-2022122111200311-0013303122122100", "registry_path": "docs/guides/data-sources--public_ip--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_public_ip.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/examples/)
- [xcsh_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/)
