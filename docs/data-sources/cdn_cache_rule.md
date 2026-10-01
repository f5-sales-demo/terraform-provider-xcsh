---
page_title: "xcsh_cdn_cache_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule landing."
---

# xcsh_cdn_cache_rule landing

<a id="canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd67cec200faa74c51023312257e505a0f156d1b6bd63d206665c15c41c795fc"></a>

## xcsh_cdn_cache_rule — xcsh_cdn_cache_rule / ce855d997dc5 / 2

Breadcrumbs:

- xcsh_cdn_cache_rule

Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.
configuration.

<a id="canonical-3840602217c06804d86c6b86acc1161fbfd2b11a175b47e955860b6ab4d6e778"></a>

## Prerequisites — xcsh_cdn_cache_rule / ce855d997dc5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-bfe16d2a38a9fab9a94f5f8f3e1237518d7f688e0a10d94213ebe91d5811201b"></a>

## Minimal configuration — xcsh_cdn_cache_rule / ce855d997dc5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNCacheRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNCacheRule by name
data "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}

output "cdn_cache_rule_id" {
  value = data.xcsh_cdn_cache_rule.example.id
}
```

<a id="canonical-dab367562aa2124443445f3a900154c52cb98d39cbfee0e5e8d6e2310ca63a50"></a>

## Root configuration — xcsh_cdn_cache_rule / ce855d997dc5 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-cbb62068f36be7488282a845219d0b1642f5606647cfeb16c86acfdd518bcc4f"></a>

## Next pages — xcsh_cdn_cache_rule / ce855d997dc5 / 6

- [Property reference](../guides/data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [Examples](../guides/data-sources--cdn_cache_rule--examples--group-001.md#canonical-5bb95f551d7b0fd1f82574637552008edccd3cf016caf1cf657dbfc5e5b9b9ef)
