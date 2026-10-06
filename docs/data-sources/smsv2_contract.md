---
page_title: "xcsh_smsv2_contract"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_contract."
---

# xcsh_smsv2_contract

<a id="canonical-3131310111322210-1120220333113231-0203221002232000-0033221002232031-3333312322003013-3120013232102311-0030302003133323-0212121103011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_smsv2_contract

Publishes the immutable clean-break SMSv2 AWS, Azure, and KVM capability contracts compiled into
this provider release.

<a id="canonical-1111102100021210-2010212132031011-1021331132031301-0101121303201021-2130211110211211-0001101232031023-3103101120333102-2202230230311311"></a>

### Prerequisites for `xcsh_smsv2_contract`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2012110021333203-2021203212211232-2000101100212301-0112103301132212-1010122220311213-3123211002221223-1200033223133103-2032231301101322"></a>

### Minimal configuration for `xcsh_smsv2_contract`

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

<a id="canonical-3001120213212031-3230202302320301-0300232102013112-0102233002010310-2130301302223200-1032103120302232-3331122203023303-3203023002223020"></a>

### Root configuration for `xcsh_smsv2_contract`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0211013211332311-3102211332101133-2133133003320312-3203311231001202-3021332033033301-1030133333232012-3333303122302320-0323012303233302"></a>

### Explore this collection for `xcsh_smsv2_contract`

- [Property reference](../guides/data-sources--smsv2_contract--reference--group-001.md#canonical-2232311231011322-0030001303101003-3113022301021321-2022203112032301-3321212110130112-3122331000100210-1110021311120232-0302332010100121)
- [Examples](../guides/data-sources--smsv2_contract--examples--group-001.md#canonical-3030331213122021-0110310112010102-3102312212320202-0330203101112213-2021110020232010-0103021002202230-0113302220300300-1231113210020030)
