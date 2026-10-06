---
page_title: "xcsh_dns_compliance_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks examples."
---

# xcsh_dns_compliance_checks examples

<a id="canonical-1031133330112323-2333310330130012-0303023000001221-2303133311033202-2003032212223212-3300303330201300-1220311133311020-2332213023222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0031332033312303-1303103011111022-1022102332331120-2133012122303100-3003222020332210-1130031122302313-3111210203332032-1102020121010322)
- Examples

<a id="canonical-2232020232212132-1321121323130231-2303111331311310-0321220323333231-2031203313022113-2330202212312103-2001323120111133-3033321312301321"></a>

### Complete configurations for `xcsh_dns_compliance_checks`

- [Data source](data-sources--dns_compliance_checks--examples--group-001.md#canonical-3223301002123130-1102123320223113-1000323131213300-2020130331010232-3112022012022111-3203321303113121-1112002121312113-2331320022132211): valid configuration.

<a id="canonical-3223301002123130-1102123320223113-1000323131213300-2020130331010232-3112022012022111-3203321303113121-1112002121312113-2331320022132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0031332033312303-1303103011111022-1022102332331120-2133012122303100-3003222020332210-1130031122302313-3111210203332032-1102020121010322)
- [Examples](data-sources--dns_compliance_checks--examples--group-001.md#canonical-1031133330112323-2333310330130012-0303023000001221-2303133311033202-2003032212223212-3300303330201300-1220311133311020-2332213023222223)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_compliance_checks/data-source.tf`; digest `sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4`.

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
