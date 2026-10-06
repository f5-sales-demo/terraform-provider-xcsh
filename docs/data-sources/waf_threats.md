---
page_title: "xcsh_waf_threats"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threats."
---

# xcsh_waf_threats

<a id="canonical-2231102302300310-3000321032301121-1323113333312231-2132011213333010-3012303030303021-1322212103212101-0130001033031000-2100323021012200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_waf_threats

Reads WAF threats information from F5 Distributed Cloud.

<a id="canonical-2023101030003230-2331332001312232-1230020201331331-1131230000221231-2210011021010310-1323130031201310-0220013232111020-0102200013000120"></a>

### Prerequisites for `xcsh_waf_threats`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2203000002001210-3101032223133113-1132101102130332-1301103033112223-2323232123303320-1012111300010210-0231320222202112-1022130321001332"></a>

### Minimal configuration for `xcsh_waf_threats`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```

<a id="canonical-2003101100221110-1333003333213211-0103203312101121-0013011100330031-2201323313330233-0030103210012130-1313111110031010-2133212332333112"></a>

### Root configuration for `xcsh_waf_threats`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3200023303330300-1331301211330232-3011220003303301-1021210022222121-0331011202121102-3013301002113112-0223302012031231-1020231013121323"></a>

### Explore this collection for `xcsh_waf_threats`

- [Property reference](../guides/data-sources--waf_threats--reference--group-001.md#canonical-1221102221200320-3320211232000012-1012023131032222-0213323223221202-3331003100012130-1302101323323020-0023211201200003-3020110212312032)
- [Examples](../guides/data-sources--waf_threats--examples--group-001.md#canonical-3230103102203013-1010033330333121-2003113212233311-1300003030111202-3133333103013112-3030011121032033-2232210213030222-1301003232032012)
