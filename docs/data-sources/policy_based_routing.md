---
page_title: "xcsh_policy_based_routing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing landing."
---

# xcsh_policy_based_routing landing

<a id="canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cb03e5b205caa27e409c0654733a7e7853cdd342891b38915faedf805221b81"></a>

## xcsh_policy_based_routing — xcsh_policy_based_routing / dad85e9ab155 / 2

Breadcrumbs:

- xcsh_policy_based_routing

Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing
create specification. configuration.

<a id="canonical-14d724a01d32613610c0dd08f7a568d92fa4efdaaa707e2ecb7c8a5168368631"></a>

## Prerequisites — xcsh_policy_based_routing / dad85e9ab155 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-968340ca600f4466718d3caaa76d39980db39e2c907c13958e88551c7d0838c3"></a>

## Minimal configuration — xcsh_policy_based_routing / dad85e9ab155 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```

<a id="canonical-85e1515f13495be23bc35c931f7ad22287e260f9b2d3036fdc6d02b728615499"></a>

## Root configuration — xcsh_policy_based_routing / dad85e9ab155 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-692d2faded383211d58f2825e441832511dc45d708c344df71eb4529df6ecfe7"></a>

## Next pages — xcsh_policy_based_routing / dad85e9ab155 / 6

- [Property reference](../guides/data-sources--policy_based_routing--reference--group-001.md#canonical-c8a08487489a65e4943a96746f170ededb15005ed6493658c87fec8120c12e9d)
- [Examples](../guides/data-sources--policy_based_routing--examples--group-001.md#canonical-815b5acc1617054d1c4c6e24d1ebd4641bd0adb8f150bb7e456ec4289bb2c40d)
