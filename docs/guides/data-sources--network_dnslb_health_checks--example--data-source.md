---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_dnslb_health_checks."
xcsh_docs: {"aliases": [], "body_bytes": 1258, "body_sha256": "sha256:3959fe0e54b6d76b111dc17bcfd4c9cca4183cf0f1fa049af253fa9d4dcddcde", "canonical_id": "xcsh-docs:data-sources:network_dnslb_health_checks:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_dnslb_health_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79", "source_path": "examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_dnslb_health_checks:example:data-source", "parent_id": "xcsh-docs:data-sources:network_dnslb_health_checks:examples", "path": "docs/guides/data-sources--network_dnslb_health_checks--example--data-source.md", "provider_name": "network_dnslb_health_checks", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_dnslb_health_checks/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_dnslb_health_checks.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md)
- [Examples](data-sources--network_dnslb_health_checks--examples.md)
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

- [Examples](data-sources--network_dnslb_health_checks--examples.md)
- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md)
