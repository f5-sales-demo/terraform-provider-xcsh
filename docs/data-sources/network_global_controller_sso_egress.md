---
page_title: "xcsh_network_global_controller_sso_egress landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_controller_sso_egress landing."
---

# xcsh_network_global_controller_sso_egress landing

<a id="canonical-36ec88256389fe8809b5331dd933d33252f115e7f05db401622b88b8bd7836c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d00fa75963efae52605968eaf6832da4db8dcecb6b6822809947574c9e7bd3e"></a>

## xcsh_network_global_controller_sso_egress — xcsh_network_global_controller_sso_egress / 441279dee30f / 2

Breadcrumbs:

- xcsh_network_global_controller_sso_egress

Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

<a id="canonical-c9b6f1fadbb2dec562208670df039ba8bd9c81e71e1cb3da70caaeb8e5c3352a"></a>

## Prerequisites — xcsh_network_global_controller_sso_egress / 441279dee30f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fae227649f3908ebe7dd5fd9441f430bd195c35206de0eca3af1a00b92872b94"></a>

## Minimal configuration — xcsh_network_global_controller_sso_egress / 441279dee30f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

<a id="canonical-b1ec8c4b2ddff4ddbf603796ddb5197ac38541c258dbc8a7b5a186a0d90a93ff"></a>

## Root configuration — xcsh_network_global_controller_sso_egress / 441279dee30f / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-590443283c3a279ee4e75ddefbbab8d143745216dbf60fb6832dad072d4b8dcf"></a>

## Next pages — xcsh_network_global_controller_sso_egress / 441279dee30f / 6

- [Property reference](../guides/data-sources--network_global_controller_sso_egress--reference--group-001.md#canonical-c3ebd65ceaf4103fa50b294a6ae35690e17f6d904c8f7e0a181ddcbe89007600)
- [Examples](../guides/data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-7c98b3c6732c61ea16a62233454bc840c3ea94a3f443551d55f740e0a4eff02e)
