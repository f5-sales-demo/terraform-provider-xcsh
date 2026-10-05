---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_dnslb_health_checks."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1464, "body_sha256": "sha256:8efeadf8735ce861ff397f35af91a6298850ba0d266492ad4549b177f5099127", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_dnslb_health_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79", "source_path": "examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_dnslb_health_checks:example:data-source", "parent_id": "xcsh-docs:data-sources:network_dnslb_health_checks:examples", "path": "documentation/data-sources/network_dnslb_health_checks/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_dnslb_health_checks", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0321111320013123-0223332231031100-3000220203002123-2232221023212333-2331320120300320-0322302331131021-1022131010202012-1111322330310001", "registry_path": "docs/guides/data-sources--network_dnslb_health_checks--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_dnslb_health_checks/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_network_dnslb_health_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf`; digest `sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_dnslb_health_checks" "https_probe" {}

# Match this explicit ingress port to the monitored endpoint.
output "https_health_check_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_dnslb_health_checks.https_probe.cidr_blocks
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/examples/)
- [xcsh_network_dnslb_health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_dnslb_health_checks/)
