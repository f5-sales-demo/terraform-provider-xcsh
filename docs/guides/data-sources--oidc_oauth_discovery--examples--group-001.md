---
page_title: "xcsh_oidc_oauth_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery examples."
---

# xcsh_oidc_oauth_discovery examples

<a id="canonical-d1c9c61b2b856005f03f81e500769e4378d5d66b010de9eb777ffee119e652a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3183c6780c3f7118278ad34b6e5d83b44b6d6632fd2d9124093affb32041781c"></a>

## Examples — Examples / ce32b6ff4f83 / 2

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
- Examples

<a id="canonical-b086a744a11c241713fd97d3595f6fb027d650a682fe3bcd1ed3901cf5ec318e"></a>

## Complete configurations — Examples / ce32b6ff4f83 / 3

- [Data source](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-0bcfa869517295be43d3a7737493cfcd1a82c022082101f0ecb1634d33bb1d5f): valid configuration.

<a id="canonical-bd86006a3f1c74fca0517d0318f48b9e004fc7f85f5ab35e8030c59f1dcb56ab"></a>

## Next pages — Examples / ce32b6ff4f83 / 4

- [Data source](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-0bcfa869517295be43d3a7737493cfcd1a82c022082101f0ecb1634d33bb1d5f)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)

<a id="canonical-0bcfa869517295be43d3a7737493cfcd1a82c022082101f0ecb1634d33bb1d5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b7480dfb4f14b29f65d050d7aa4c39fb7606a665d49a941cf3c266e4b9377f0"></a>

## Data source — Data source / fd19f6382efe / 2

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
- [Examples](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-d1c9c61b2b856005f03f81e500769e4378d5d66b010de9eb777ffee119e652a8)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf`; digest `sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a`.

```terraform
# OIDCOauthDiscovery DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_oidc_oauth_discovery" "example" {
  namespace = "example-value"
}

output "oidc_oauth_discovery_result" {
  value = data.xcsh_oidc_oauth_discovery.example
}
```

<a id="canonical-da13b34ea21ad53ba7d5147f3d6ffdf65fbe25a60e93b8161cefa988c6df2852"></a>

## Next pages — Data source / fd19f6382efe / 3

- [Examples](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-d1c9c61b2b856005f03f81e500769e4378d5d66b010de9eb777ffee119e652a8)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
