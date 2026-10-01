---
page_title: "xcsh_waf_exclusion_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy landing."
---

# xcsh_waf_exclusion_policy landing

<a id="canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0a8bda19d70d0e6cae8178e7fb7c2f24b3503c244d1a5e3480f8877c9f8fd87"></a>

## xcsh_waf_exclusion_policy — xcsh_waf_exclusion_policy / f59c1bdfc10a / 2

Breadcrumbs:

- xcsh_waf_exclusion_policy

Manages WAF exclusion policy in F5 Distributed Cloud.

<a id="canonical-51bccff38509a2aa4d5829e2477da4f54b798ca42efe7640bb68d1059b9a1af3"></a>

## Prerequisites — xcsh_waf_exclusion_policy / f59c1bdfc10a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3ef6f21a7f1abba94795e6c3ce78933c094df475a66914825c53284c2447a635"></a>

## Minimal configuration — xcsh_waf_exclusion_policy / f59c1bdfc10a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFExclusionPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WAFExclusionPolicy by name
data "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}

output "waf_exclusion_policy_id" {
  value = data.xcsh_waf_exclusion_policy.example.id
}
```

<a id="canonical-4feeb0926dae84922724d64dc37978389b22b8043bf72c49e69c14800e378f74"></a>

## Root configuration — xcsh_waf_exclusion_policy / f59c1bdfc10a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-148dd10523c6c231b7c34c8806b7237f963e5d952b7962ec2e45209dc13b71b0"></a>

## Next pages — xcsh_waf_exclusion_policy / f59c1bdfc10a / 6

- [Property reference](../guides/data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [Examples](../guides/data-sources--waf_exclusion_policy--examples--group-001.md#canonical-c81efcf5a52a4fc6a6d88f4e932312fe4a94dd7084d183754fbb93bf646c0abf)
