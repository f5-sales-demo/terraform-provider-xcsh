---
page_title: "xcsh_app_security_evidence"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_security_evidence."
---

# xcsh_app_security_evidence

<a id="canonical-3103302000231201-2113100232313221-0102211103022300-2213011330310113-3301220213112210-1312012211213332-0301203231301211-3001330131022123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_app_security_evidence

Reads App Security Evidence information from F5 Distributed Cloud.

<a id="canonical-0203012011101233-1013032310320321-2023000233112301-1332123022101102-1301303032011012-2323130111333113-0210122301002002-0023200301133002"></a>

### Prerequisites for `xcsh_app_security_evidence`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3020213212231000-0331033013201023-1103231130013133-1212321011201330-0221112210031100-1300122012322201-3312211111230010-0201211301013312"></a>

### Minimal configuration for `xcsh_app_security_evidence`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSecurityEvidence DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_app_security_evidence" "example" {
  namespace = "example-value"
}

output "app_security_evidence_result" {
  value = data.xcsh_app_security_evidence.example
}
```

<a id="canonical-0003210321331102-1310201301031313-2220030101133112-2301231133013330-3201133221123013-3110220300131210-3212102330302310-3031030301003223"></a>

### Root configuration for `xcsh_app_security_evidence`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1103211331201011-1033330302300213-3231111020101331-1231013021013310-2302222101211302-3111202111212223-3330230211102102-3333020001322000"></a>

### Explore this collection for `xcsh_app_security_evidence`

- [Property reference](../guides/data-sources--app_security_evidence--reference--group-001.md#canonical-0112003311231232-2332331022230311-2301110100313003-2332031131201310-1301022100202113-0011223203211232-1031003322010031-0310200130000020)
- [Examples](../guides/data-sources--app_security_evidence--examples--group-001.md#canonical-1111331131013003-3213311303211110-2313013321313320-3311311023003023-2000200332111220-3022002302310032-1122310003121131-0331331311233212)
