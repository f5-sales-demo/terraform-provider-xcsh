---
page_title: "xcsh_rate_limiter_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy landing."
---

# xcsh_rate_limiter_policy landing

<a id="canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-308a34eb3722d5f4d48d00d1953fd64b7e860f2c7f4faee3dc30290548bce31d"></a>

## xcsh_rate_limiter_policy — xcsh_rate_limiter_policy / 4872ee914e6d / 2

Breadcrumbs:

- xcsh_rate_limiter_policy

Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create
specification. configuration.

<a id="canonical-de1c537934105572ee891945637c056a3869521cebce706ad1c0fd411eb71e40"></a>

## Prerequisites — xcsh_rate_limiter_policy / 4872ee914e6d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-69f6b650a41326b33ac5feb2bc36e439442d8a2ffc2cff3536c0f7f406902738"></a>

## Minimal configuration — xcsh_rate_limiter_policy / 4872ee914e6d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiterPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiterPolicy by name
data "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}

output "rate_limiter_policy_id" {
  value = data.xcsh_rate_limiter_policy.example.id
}
```

<a id="canonical-a3649a17364bf41bd81c4303da10298840aa918c945241cd06eb32d0c0c02b61"></a>

## Root configuration — xcsh_rate_limiter_policy / 4872ee914e6d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2457217aee095b5d68e8f830c4b7869be8cdb36fbbe75acc49395405f20bc1e7"></a>

## Next pages — xcsh_rate_limiter_policy / 4872ee914e6d / 6

- [Property reference](../guides/data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [Examples](../guides/data-sources--rate_limiter_policy--examples--group-001.md#canonical-9af59c23111ef40dd083e7310441f58a2c7e0e1ad5fb8e562709dd4ea7f2db9f)
