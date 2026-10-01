---
page_title: "xcsh_oidc_oauth_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery landing."
---

# xcsh_oidc_oauth_discovery landing

<a id="canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a38869c9a3083d98ef2b66c0ff0e4f971236ffbb2b86550570d8208ab3a8bb88"></a>

## xcsh_oidc_oauth_discovery — xcsh_oidc_oauth_discovery / 93db28599dd6 / 2

Breadcrumbs:

- xcsh_oidc_oauth_discovery

Resource creation operation.

<a id="canonical-7b5dcdbc352e9b8a303823e64210c539cdd374a91abcbd6dcb1125a216c3e1a4"></a>

## Prerequisites — xcsh_oidc_oauth_discovery / 93db28599dd6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4a5b0af13868cd02c6c9a91c7bfc5a5c065ada0561deb8d946fdc29e07d8233c"></a>

## Minimal configuration — xcsh_oidc_oauth_discovery / 93db28599dd6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-6c0d8b583393d520138d3101a52885a6750b0f5c04622d45f6cbcfd171692a70"></a>

## Root configuration — xcsh_oidc_oauth_discovery / 93db28599dd6 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-862f7763957f014339110c378b97f2ca288b6fccc8192dd00e2ce72ef2a7b1dd"></a>

## Next pages — xcsh_oidc_oauth_discovery / 93db28599dd6 / 6

- [Property reference](../guides/data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a)
- [Examples](../guides/data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-d1c9c61b2b856005f03f81e500769e4378d5d66b010de9eb777ffee119e652a8)
