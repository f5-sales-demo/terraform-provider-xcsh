---
page_title: "xcsh_site_upgrade_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_status examples."
---

# xcsh_site_upgrade_status examples

<a id="canonical-eb57b52c1bb429fc2286edc14d3638f601a51538ea5c7af286ba952b82116410"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bd1714aba84fc617441807d604f659c96f6f8d5638a39b2bec0f065618d0964"></a>

## Examples — Examples / dacb244f01f5 / 2

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)
- Examples

<a id="canonical-376d7a49234a87bc81f59c71d524a4d028754450c70ced45f1e24bd5a1a45e34"></a>

## Complete configurations — Examples / dacb244f01f5 / 3

- [Data source](data-sources--site_upgrade_status--examples--group-001.md#canonical-b0a93d669a4b865842ba601b16c4fd75e523d0911bb7d1fc3aef788ac0c858bb): valid configuration.

<a id="canonical-2877ab84e6a366779326f39cb7c5e8cdd4dbcb5a79b8c56270943baea64794c8"></a>

## Next pages — Examples / dacb244f01f5 / 4

- [Data source](data-sources--site_upgrade_status--examples--group-001.md#canonical-b0a93d669a4b865842ba601b16c4fd75e523d0911bb7d1fc3aef788ac0c858bb)
- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)

<a id="canonical-b0a93d669a4b865842ba601b16c4fd75e523d0911bb7d1fc3aef788ac0c858bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d99ad04e5ad51e1d57be12367058941ac2c445bb12134bec820d4855d3b7654e"></a>

## Data source — Data source / 38eea9616398 / 2

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)
- [Examples](data-sources--site_upgrade_status--examples--group-001.md#canonical-eb57b52c1bb429fc2286edc14d3638f601a51538ea5c7af286ba952b82116410)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_upgrade_status/data-source.tf`; digest `sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683`.

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

<a id="canonical-952f48ca24e975aa18f1cae225199f66e1a275cd5204f2bca21ace3120fcb120"></a>

## Next pages — Data source / 38eea9616398 / 3

- [Examples](data-sources--site_upgrade_status--examples--group-001.md#canonical-eb57b52c1bb429fc2286edc14d3638f601a51538ea5c7af286ba952b82116410)
- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md#canonical-354cc1e18b1eae705be66b9b6e4a44fd8152a2aa13c69db8d19d96b451ff6445)
