---
page_title: "xcsh_waf_exclusion_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy."
---

# xcsh_waf_exclusion_policy

<a id="canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_waf_exclusion_policy

Reads WAF Exclusion Policy information from F5 Distributed Cloud.

<a id="canonical-3000222023312201-2131130031003212-3022322001132032-1333231330023302-1023031100033002-1010310122113203-1020003320201313-3021332033312013"></a>

### Prerequisites for `xcsh_waf_exclusion_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1101233030333303-2011002122022222-1031112002213202-1013133122103311-1023132120302210-0232333213121000-2323122031010011-2123212201223303"></a>

### Minimal configuration for `xcsh_waf_exclusion_policy`

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

<a id="canonical-0332331233020122-1333012223232221-1013211132123003-3032132021030330-0021103133101311-2212122101102002-1130110302201030-0210101322120311"></a>

### Root configuration for `xcsh_waf_exclusion_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1033323223002102-1231223220102102-0213021031121031-3003132113200320-2123020223200010-0323331302301021-3212213001102000-0032031320331310"></a>

### Explore this collection for `xcsh_waf_exclusion_policy`

- [Property reference](../guides/data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [Examples](../guides/data-sources--waf_exclusion_policy--examples--group-001.md#canonical-3020013233303311-2211022210333012-2212312020331032-2103020301023332-1022211031311300-2010310120031311-1033232321032333-1210123000222333)
