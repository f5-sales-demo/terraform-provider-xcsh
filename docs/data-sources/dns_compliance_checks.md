---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks."
---

# xcsh_dns_compliance_checks

<a id="canonical-0031332033312303-1303103011111022-1022102332331120-2133012122303100-3003222020332210-1130031122302313-3111210203332032-1102020121010322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_compliance_checks

Reads DNS Compliance Checks information from F5 Distributed Cloud.

<a id="canonical-2110103310300110-3300131230001133-1123232231321300-0102200030003100-2230102202320113-1111230202121300-3222101021322003-1132201222110202"></a>

### Prerequisites for `xcsh_dns_compliance_checks`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2331221302321121-0131200223203302-0030012012122331-2022312010301002-3100022121020223-0221103332002020-3122301110331011-0113013012033232"></a>

### Minimal configuration for `xcsh_dns_compliance_checks`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSComplianceChecks by name
data "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"
}

output "dns_compliance_checks_id" {
  value = data.xcsh_dns_compliance_checks.example.id
}
```

<a id="canonical-1210032010010131-3201010311130002-0303331313012130-2030030200121022-2200230200103020-0200231101310102-3330002232133233-1030102000133011"></a>

### Root configuration for `xcsh_dns_compliance_checks`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1312211010322210-1302312231330103-0210202322011030-2131312131031331-2032303021022102-1233010332000321-3203333033010301-2332323131201212"></a>

### Explore this collection for `xcsh_dns_compliance_checks`

- [Property reference](../guides/data-sources--dns_compliance_checks--reference--group-001.md#canonical-0330131002333110-1321122220010200-2131223002102021-1310111221323211-0220230233200220-1023112122322033-2333123220200320-2331221120111033)
- [Examples](../guides/data-sources--dns_compliance_checks--examples--group-001.md#canonical-1031133330112323-2333310330130012-0303023000001221-2303133311033202-2003032212223212-3300303330201300-1220311133311020-2332213023222223)
