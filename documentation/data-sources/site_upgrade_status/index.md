---
page_title: "xcsh_site_upgrade_status"
subcategory: ""
description: "Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and operating-system targets to converge."
xcsh_docs: {"aliases": ["site upgrade status"], "body_bytes": 1619, "body_sha256": "sha256:edc66b639dbed178b1c9cb5d6cb27c14bc8c9f2a936673a5bec76b8f58d49a7a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_upgrade_status:reference", "xcsh-docs:data-sources:site_upgrade_status:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_upgrade_status:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_upgrade_status/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0311103030013201-2023013222321300-1123321212232123-1232102210103331-2001110222022222-0103301221312320-3101213121122310-1101333312101011", "registry_path": "docs/data-sources/site_upgrade_status.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and operating-system targets to converge.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_upgrade_status

Breadcrumbs:

- xcsh_site_upgrade_status

Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and
operating-system targets to converge.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Observe upgrade eligibility or wait for supplied software and OS targets to
# be installed with the site back ONLINE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 7.3.0"
    }
  }
}

data "xcsh_site_upgrade_status" "site" {
  site = "example-smsv2-site"

  expected_software_version = "crt-20260201-0179"
  expected_os_version       = "9.2026.17"
  wait                      = true
  timeout_seconds           = 7200
  poll_interval_seconds     = 30
}

output "upgrade_converged" {
  value = data.xcsh_site_upgrade_status.site.target_converged
}
```

## Root configuration

Required root properties: `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/examples/)
