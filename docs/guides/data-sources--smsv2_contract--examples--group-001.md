---
page_title: "xcsh_smsv2_contract examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_contract examples."
---

# xcsh_smsv2_contract examples

<a id="canonical-3030331213122021-0110310112010102-3102312212320202-0330203101112213-2021110020232010-0103021002202230-0113302220300300-1231113210020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-3131310111322210-1120220333113231-0203221002232000-0033221002232031-3333312322003013-3120013232102311-0030302003133323-0212121103011131)
- Examples

<a id="canonical-0000322323131210-3031203011322103-1223112323120122-0121030010131330-0211122213213102-1222223103210302-2023323321100311-3111013102131000"></a>

### Complete configurations for `xcsh_smsv2_contract`

- [Data source](data-sources--smsv2_contract--examples--group-001.md#canonical-2110122221312022-2232102303200000-0031102010330212-3302100323023133-1232123332011220-2332032003030211-2001011312231033-1333022133021312): valid configuration.

<a id="canonical-2110122221312022-2232102303200000-0031102010330212-3302100323023133-1232123332011220-2332032003030211-2001011312231033-1333022133021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_smsv2_contract](../data-sources/smsv2_contract.md#canonical-3131310111322210-1120220333113231-0203221002232000-0033221002232031-3333312322003013-3120013232102311-0030302003133323-0212121103011131)
- [Examples](data-sources--smsv2_contract--examples--group-001.md#canonical-3030331213122021-0110310112010102-3102312212320202-0330203101112213-2021110020232010-0103021002202230-0113302220300300-1231113210020030)
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
