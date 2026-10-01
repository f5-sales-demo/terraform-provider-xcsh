---
page_title: "xcsh_smsv2_contract landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_contract landing."
---

# xcsh_smsv2_contract landing

<a id="canonical-ddd15ea458a3f5ed23a42b800fa42b8dffdba0c7d81ee4b50cc837fb2665315d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-554902648499e34549f5e371116738499c9549650146e34bd3458fd2a2b2cd75"></a>

## xcsh_smsv2_contract — xcsh_smsv2_contract / 05febd154682 / 2

Breadcrumbs:

- xcsh_smsv2_contract

Publishes the immutable clean-break SMSv2 AWS, Azure, and KVM capability contracts compiled into
this provider release.

<a id="canonical-86509fe3898e696e804509b1164f17a6446a8d67db942a6b603eb7d38eb7147a"></a>

## Prerequisites — xcsh_smsv2_contract / 05febd154682 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c162798dec8b2e3130b921d612bc21349cc72ae04e4d8caefd6a32f3e32c2ac8"></a>

## Minimal configuration — xcsh_smsv2_contract / 05febd154682 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Read the immutable clean-break SMSv2 contract compiled into the provider.
# Required capabilities are checked during planning before any F5 API request.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_contract" "current" {
  required_capabilities = ["runtime_status"]
}

output "smsv2_contract" {
  value = {
    id                               = data.xcsh_smsv2_contract.current.contract_id
    version                          = data.xcsh_smsv2_contract.current.contract_version
    api_release                      = data.xcsh_smsv2_contract.current.api_release_tag
    telemetry_schema_id              = data.xcsh_smsv2_contract.current.telemetry_schema_id
    capabilities                     = data.xcsh_smsv2_contract.current.capabilities
    f5xc_authorities                 = data.xcsh_smsv2_contract.current.f5xc_authorities
    aws_authorities                  = data.xcsh_smsv2_contract.current.aws_authorities
    azure_route_server_ebgp_multihop = data.xcsh_smsv2_contract.current.azure_route_server_ebgp_multihop
  }
}
```

<a id="canonical-251e5fb5d297e45f9f7c3e36e3d6d062c9f8f3f14c7ffb86ffcdacb83b1b3bf2"></a>

## Root configuration — xcsh_smsv2_contract / 05febd154682 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-694e2d55ef8e7159096af27303c19b3b93d54c03ece52365d536a7d430ade7e5"></a>

## Next pages — xcsh_smsv2_contract / 05febd154682 / 6

- [Property reference](../guides/data-sources--smsv2_contract--reference--group-001.md#canonical-aed6d17a0c073443d72b12798a8d63b1f9994716daf404245427562e32f84419)
- [Examples](../guides/data-sources--smsv2_contract--examples--group-001.md#canonical-ccf6768914d16112d2da6e223c8d15a789508b84132428ac17ca8c306d5e420c)
