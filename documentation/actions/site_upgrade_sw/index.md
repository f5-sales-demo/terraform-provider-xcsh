---
page_title: "xcsh_site_upgrade_sw"
subcategory: ""
description: "Request an in-place site software upgrade."
xcsh_docs: {"aliases": ["site upgrade sw"], "body_bytes": 1471, "body_sha256": "sha256:dcaf02c87f7f3d2fa918874d502bf175b86ae856d8ea9150b98ce5afc39d4c67", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:site_upgrade_sw:reference", "xcsh-docs:actions:site_upgrade_sw:examples", "xcsh-docs:actions:site_upgrade_sw:lifecycle"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_upgrade_sw:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/site_upgrade_sw/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-0120130312312320-0223010000032113-2233123110100201-3130110122132200-2323212123211333-3120313220122233-0023002110221112-0320032112111211", "registry_path": "docs/actions/site_upgrade_sw.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Request an in-place site software upgrade.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/lifecycle/)
