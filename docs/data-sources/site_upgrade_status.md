---
page_title: "xcsh_site_upgrade_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_status landing."
---

# xcsh_site_upgrade_status landing

<a id="canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f19be87ef3e1738c5319c498cddfc89e09b5e3ac268138e2541eec00e71087df"></a>

## xcsh_site_upgrade_status — xcsh_site_upgrade_status / 0ec5f9456efa / 2

Breadcrumbs:

- xcsh_site_upgrade_status

Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and
operating-system targets to converge.

<a id="canonical-61e59d3767970490baaaebb5284bd57af6faf5adf7143546d5491aeb6bb718ee"></a>

## Prerequisites — xcsh_site_upgrade_status / 0ec5f9456efa / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-62ac7cd9f8b758496cc87f5fdd74f0fb136163e1d0a4f08a39c13f3bc3aa4fb9"></a>

## Minimal configuration — xcsh_site_upgrade_status / 0ec5f9456efa / 4

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

<a id="canonical-1df8bc63f8456ddf2a1b4f4ddd75b9def7da5d156053703ea513f03c85e6a098"></a>

## Root configuration — xcsh_site_upgrade_status / 0ec5f9456efa / 5

Required root properties: `site`. Full root flags and choices appear in the property reference.

<a id="canonical-a9c8d277f4c6b8caeb9a5bc010651575a1f2332f0890c6d68a14879d4a8b523e"></a>

## Next pages — xcsh_site_upgrade_status / 0ec5f9456efa / 6

- [Property reference](../guides/data-sources--site_upgrade_status--reference--group-001.md#canonical-026e1aecd2a970e096dc9476e95faa6d85dfc2ee51f7bdce20b6431348fda126)
- [Examples](../guides/data-sources--site_upgrade_status--examples--group-001.md#canonical-eb57b52c1bb429fc2286edc14d3638f601a51538ea5c7af286ba952b82116410)
