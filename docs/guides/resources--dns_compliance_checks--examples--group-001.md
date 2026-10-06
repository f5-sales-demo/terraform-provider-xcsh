---
page_title: "xcsh_dns_compliance_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks examples."
---

# xcsh_dns_compliance_checks examples

<a id="canonical-3223100300030001-0303112321333032-1230022123231032-0200122111033303-2032110213100222-3010330013000322-1120330313333123-1101230331011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)
- Examples

<a id="canonical-2220333011232223-0233313313020121-2200111102011220-2021112033210130-2213211210230130-1011211220033100-2320033202210123-0222111102322222"></a>

### Complete configurations for `xcsh_dns_compliance_checks`

- [Resource](resources--dns_compliance_checks--examples--group-001.md#canonical-0310000323310100-1023120330100322-0333100003022130-1231033132331112-0002303000233003-1030200222032233-3320101210031232-2022232323030301): valid configuration.

<a id="canonical-0310000323310100-1023120330100322-0333100003022130-1231033132331112-0002303000233003-1030200222032233-3320101210031232-2022232323030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210)
- [Examples](resources--dns_compliance_checks--examples--group-001.md#canonical-3223100300030001-0303112321333032-1230022123231032-0200122111033303-2032110213100222-3010330013000322-1120330313333123-1101230331011201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_compliance_checks/resource.tf`; digest `sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1`.

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
