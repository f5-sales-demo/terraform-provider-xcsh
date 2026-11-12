---
page_title: "xcsh_site_upgrade_sw"
subcategory: ""
description: "Request an in-place site software upgrade."
xcsh_docs: {"aliases": ["site upgrade sw"], "body_bytes": 1484, "body_sha256": "sha256:d6ef829bb7d960d7a4e37025c4e741a51219189d642406c550d1b89dfb1e301f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:site_upgrade_sw:reference", "xcsh-docs:actions:site_upgrade_sw:examples", "xcsh-docs:actions:site_upgrade_sw:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_upgrade_sw:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/site_upgrade_sw/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "actions", "registry_anchor": "canonical-0120130312312320-0223010000032113-2233123110100201-3130110122132200-2323212123211333-3120313220122233-0023002110221112-0320032112111211", "registry_path": "docs/actions/site_upgrade_sw.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Request an in-place site software upgrade.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_upgrade_sw

Breadcrumbs:

- xcsh_site_upgrade_sw

Request an in-place site software upgrade.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteUpgradeSw Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts the upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_sw" "example" {
  config {
    site             = "example-value"
    software_version = "example-value"
  }
}
```

## Root configuration

Required root properties: `site`, `software_version`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/lifecycle/)
