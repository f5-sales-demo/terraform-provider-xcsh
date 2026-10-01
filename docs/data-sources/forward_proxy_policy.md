---
page_title: "xcsh_forward_proxy_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy landing."
---

# xcsh_forward_proxy_policy landing

<a id="canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73f0fbec649f18b356ec0319fec7a6f154f2524397b2a0a66e9978651dcab040"></a>

## xcsh_forward_proxy_policy — xcsh_forward_proxy_policy / 6d877fac4e21 / 2

Breadcrumbs:

- xcsh_forward_proxy_policy

Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy
specification. configuration.

<a id="canonical-c86ebcec5b41ee012c16bf6739fc42ac2db3370a55e540c5adb09b8809d0f7a3"></a>

## Prerequisites — xcsh_forward_proxy_policy / 6d877fac4e21 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-714d3e4dce0ce65d84d1acff261994ce83757b8b9d790037a43398c05e401bcf"></a>

## Minimal configuration — xcsh_forward_proxy_policy / 6d877fac4e21 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardProxyPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardProxyPolicy by name
data "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}

output "forward_proxy_policy_id" {
  value = data.xcsh_forward_proxy_policy.example.id
}
```

<a id="canonical-9b02c52dc25d7f6f7da4befb075de196362626ca8b11a7283f13e07e4ed809a8"></a>

## Root configuration — xcsh_forward_proxy_policy / 6d877fac4e21 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-07d8982fb8748536763286e98681297a96048b57ab0b7343e92c1b9896892d00"></a>

## Next pages — xcsh_forward_proxy_policy / 6d877fac4e21 / 6

- [Property reference](../guides/data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [Examples](../guides/data-sources--forward_proxy_policy--examples--group-001.md#canonical-f108bba99b58a1f0aa3abba2accf4a02569830c4a7fc769e98e7fe3de1c6625e)
