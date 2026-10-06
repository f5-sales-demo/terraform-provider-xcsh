---
page_title: "xcsh_dns_zone examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone examples."
---

# xcsh_dns_zone examples

<a id="canonical-0113332031001202-0203203323031122-1110313201121223-0211122311121012-1002330222223130-0033020333320010-2320222023212020-2311211332221300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- Examples

<a id="canonical-1322002332001110-2331103113031021-2120330300302231-2131332301113303-2022011030030013-3330320011033322-1300111110010011-3121211230220003"></a>

### Complete configurations for `xcsh_dns_zone`

- [Data source](data-sources--dns_zone--examples--group-001.md#canonical-1022231110303301-2111303011121232-1120312000231330-3303002003013002-2303310130300221-1313130113032001-2020202123303123-0122221332001130): valid configuration.

<a id="canonical-1022231110303301-2111303011121232-1120312000231330-3303002003013002-2303310130300221-1313130113032001-2020202123303123-0122221332001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Examples](data-sources--dns_zone--examples--group-001.md#canonical-0113332031001202-0203203323031122-1110313201121223-0211122311121012-1002330222223130-0033020333320010-2320222023212020-2311211332221300)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone/data-source.tf`; digest `sha256:45c5f9a8c56f2ed91ab404e89f23a1d58d7ad35539ecaaccb097a0e0066d41b9`.

```terraform
# DNSZone Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSZone by name
data "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"
}

# Fail closed when this stack depends on an externally owned zone.
resource "terraform_data" "require_managed_records" {
  lifecycle {
    precondition {
      condition = try(
        data.xcsh_dns_zone.example.primary.allow_http_lb_managed_records,
        false
      )
      error_message = "The selected DNS zone must enable HTTP LB managed records."
    }
  }
}

output "dns_zone_id" {
  value = data.xcsh_dns_zone.example.id
}
```
