---
page_title: "xcsh_smsv2_contract examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_contract examples."
---

# xcsh_smsv2_contract examples

<a id="canonical-ccf6768914d16112d2da6e223c8d15a789508b84132428ac17ca8c306d5e420c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00ebb764cd8c5e936b5bb61a1930477c256a79d26aad39328bef9435d51d2740"></a>

## Examples — Examples / 3ef1979b2142 / 2

Breadcrumbs:

- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-ddd15ea458a3f5ed23a42b800fa42b8dffdba0c7d81ee4b50cc837fb2665315d)
- Examples

<a id="canonical-e9ce351486f189bcdc9c600200eed227156bf7bd5f85924add4fb097eb12ffda"></a>

## Complete configurations — Examples / 3ef1979b2142 / 3

- [Data source](data-sources--smsv2_contract--examples--group-001.md#canonical-946a9d8aae4b38000d484f26f243b2df6e6fe168be38332581176b4f7f29f276): valid configuration.

<a id="canonical-249d71d926cd0aa86caa45caa782a40151fa2fd13c0f9bb3dc427b67679b0e1d"></a>

## Next pages — Examples / 3ef1979b2142 / 4

- [Data source](data-sources--smsv2_contract--examples--group-001.md#canonical-946a9d8aae4b38000d484f26f243b2df6e6fe168be38332581176b4f7f29f276)
- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-ddd15ea458a3f5ed23a42b800fa42b8dffdba0c7d81ee4b50cc837fb2665315d)

<a id="canonical-946a9d8aae4b38000d484f26f243b2df6e6fe168be38332581176b4f7f29f276"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11ca0bc3d840e195fec1c82252154606fb67e793250d04664e76c926cddae800"></a>

## Data source — Data source / 51581c59f382 / 2

Breadcrumbs:

- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-ddd15ea458a3f5ed23a42b800fa42b8dffdba0c7d81ee4b50cc837fb2665315d)
- [Examples](data-sources--smsv2_contract--examples--group-001.md#canonical-ccf6768914d16112d2da6e223c8d15a789508b84132428ac17ca8c306d5e420c)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_contract/data-source.tf`; digest `sha256:e80a034ec67b555714059852f638089936add60d6d1f3b5f77d3c2d3ae050f0f`.

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

<a id="canonical-7a18f8f5c1f90972a5afbb55a65801e388918bbca4aefb2fde619e0574fbf0f4"></a>

## Next pages — Data source / 51581c59f382 / 3

- [Examples](data-sources--smsv2_contract--examples--group-001.md#canonical-ccf6768914d16112d2da6e223c8d15a789508b84132428ac17ca8c306d5e420c)
- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-ddd15ea458a3f5ed23a42b800fa42b8dffdba0c7d81ee4b50cc837fb2665315d)
