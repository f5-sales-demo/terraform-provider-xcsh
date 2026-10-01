---
page_title: "xcsh_rate_limiter landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter landing."
---

# xcsh_rate_limiter landing

<a id="canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf154b272576abc126a809ee3520c9c923bd936abb9982332ebc85fc80287e0b"></a>

## xcsh_rate_limiter — xcsh_rate_limiter / 04c1ce5906db / 2

Breadcrumbs:

- xcsh_rate_limiter

Manages rate\_limiter creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-f988b2bc2fcb05a181f0b8185914d3b485f1fbe210b9d509111b7b5c39c90f30"></a>

## Prerequisites — xcsh_rate_limiter / 04c1ce5906db / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `rate_limiter_policy`.

- rate_limiter_policy: Detailed rate limiting rules

<a id="canonical-b96e8f0db9bcb1bbef98f37927c7de501525600f71b28bdc22cb694df7694f80"></a>

## Minimal configuration — xcsh_rate_limiter / 04c1ce5906db / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiter Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiter by name
data "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}

output "rate_limiter_id" {
  value = data.xcsh_rate_limiter.example.id
}
```

<a id="canonical-f879789b27d79d279413ece9ce07499bd6242d7b9a5c54508c90ab1f5bbc7391"></a>

## Root configuration — xcsh_rate_limiter / 04c1ce5906db / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dd1512812a76063997a6a818c44a439c744c35eb4eee1b697119572dd23a41f3"></a>

## Next pages — xcsh_rate_limiter / 04c1ce5906db / 6

- [Property reference](../guides/data-sources--rate_limiter--reference--group-001.md#canonical-675f099fbfef5fab17eb7caf586a5b843f100d7f7b90b906635610c96b9d19eb)
- [Examples](../guides/data-sources--rate_limiter--examples--group-001.md#canonical-d2aa5143f978329a2cb5d3186781df77d0e69f043c28f96ae10250235db9d56e)
