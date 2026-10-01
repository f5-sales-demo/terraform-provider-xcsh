---
page_title: "xcsh_site_cloud_init landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_cloud_init landing."
---

# xcsh_site_cloud_init landing

<a id="canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff73dd2e2b915723f837c18ad331cec2467349323f1405a78018eba7adca7300"></a>

## xcsh_site_cloud_init — xcsh_site_cloud_init / d132062d6435 / 2

Breadcrumbs:

- xcsh_site_cloud_init

Retrieve Customer Edge cloud-init template.

<a id="canonical-ddd573fb3eff406fcdf9e69ceb086d1b07288f5a895a5210072813c8d3f46889"></a>

## Prerequisites — xcsh_site_cloud_init / d132062d6435 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1be823cd5e50310db4ea18accca77c0e729aa5f434d99198094188ff49e70a07"></a>

## Minimal configuration — xcsh_site_cloud_init / d132062d6435 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```

<a id="canonical-040d2c08e8d79886a453d660bceafb8eed1cadc3dd0edf4d6f817bf8c824186a"></a>

## Root configuration — xcsh_site_cloud_init / d132062d6435 / 5

Required root properties: `provider_ref`, `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-745c18569f13c7d67616027fd22f31604c4a5de88bf859ab35b3f7ed8f472e3b"></a>

## Next pages — xcsh_site_cloud_init / d132062d6435 / 6

- [Property reference](../guides/data-sources--site_cloud_init--reference--group-001.md#canonical-449146e9b12ccd06dd9f1b3726adeaaf418081765a16b4774522c5040c4e90da)
- [Examples](../guides/data-sources--site_cloud_init--examples--group-001.md#canonical-2af1fe52fef4663a7a40baa6d63bcf1d5b04630a18ac1e64cc5f4883fd0f6ca4)
