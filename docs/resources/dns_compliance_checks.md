---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks."
---

# xcsh_dns_compliance_checks

<a id="canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_compliance_checks

Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-3320332123230130-1200311123102103-1000002311020223-0222332212001021-0100010212021330-3012300302011300-1230221130211003-3133030001333122"></a>

### Prerequisites for `xcsh_dns_compliance_checks`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0313013021031012-1113203110012103-1101320222133203-0123010100020023-1230131123111112-0030113113321003-2130032203330010-2110030130110000"></a>

### Minimal configuration for `xcsh_dns_compliance_checks`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Resource Example
# Manages DNS Compliance Checks Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSComplianceChecks configuration
resource "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"

  domain_denylist = ["example-value"]
}
```

<a id="canonical-2202013233103022-2302302122223133-0220322231033222-2330323130230310-2032320012233311-2131013033131122-0132302002110032-1213333211220120"></a>

### Root configuration for `xcsh_dns_compliance_checks`

Required root properties: `domain_denylist`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0122231133221123-3313122033122331-2010222300110112-1123301032020113-2103212030201201-3121111022322032-0002321013302332-0110020111031122"></a>

### Explore this collection for `xcsh_dns_compliance_checks`

- [Property reference](../guides/resources--dns_compliance_checks--reference--group-001.md#canonical-3232320123103113-0100311102312313-0121302010013132-2312320000130313-1222132032330322-2302213002320222-1233300132211231-3112210312300202)
- [Examples](../guides/resources--dns_compliance_checks--examples--group-001.md#canonical-3223100300030001-0303112321333032-1230022123231032-0200122111033303-2032110213100222-3010330013000322-1120330313333123-1101230331011201)
- [Import](../guides/resources--dns_compliance_checks--lifecycle--group-001.md#canonical-0332203010112020-0330310122032110-2023110300311131-1212322311002112-1201323103020202-3010210330100323-1022020332100223-0103033113122301)
- [Timeouts](../guides/resources--dns_compliance_checks--lifecycle--group-001.md#canonical-2312121020011203-3130101230331010-2311131120002133-2201023300033220-0211200200330233-1123302233212230-1120020010110233-3322112333300103)
