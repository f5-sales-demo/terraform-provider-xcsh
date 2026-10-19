---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_dnslb_health_checks."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1191, "body_sha256": "sha256:4a44e51a07abdd4a5bad8d05da15a3522806b85cd11c9ffced5054bf9f7434a3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_dnslb_health_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79", "source_path": "examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_dnslb_health_checks:example:data-source", "parent_id": "xcsh-docs:data-sources:network_dnslb_health_checks:examples", "path": "documentation/data-sources/network_dnslb_health_checks/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_dnslb_health_checks", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0321111320013123-0223332231031100-3000220203002123-2232221023212333-2331320120300320-0322302331131021-1022131010202012-1111322330310001", "registry_path": "docs/guides/data-sources--network_dnslb_health_checks--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_dnslb_health_checks/examples/data-source/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Data source for xcsh_network_dnslb_health_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": [], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
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
