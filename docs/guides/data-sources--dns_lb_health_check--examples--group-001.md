---
page_title: "xcsh_dns_lb_health_check examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check examples."
---

# xcsh_dns_lb_health_check examples

<a id="canonical-1100310202032332-3100202312011332-3023002312200212-2123103002011211-3032130332210023-3033222223332222-3012210210300001-2201111013023133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- Examples

<a id="canonical-1132212110023033-2103020333131032-1132332113302030-3122321202111322-1201221013031131-1332311030333231-0101211020202212-2333021313031231"></a>

### Complete configurations for `xcsh_dns_lb_health_check`

- [Data source](data-sources--dns_lb_health_check--examples--group-001.md#canonical-2003210233311222-2101232131020233-0133202322230232-2223022112312200-1313022310003123-3100013002122131-2120200330211101-0332020133301210): valid configuration.

<a id="canonical-2003210233311222-2101232131020233-0133202322230232-2223022112312200-1313022310003123-3100013002122131-2120200330211101-0332020133301210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Examples](data-sources--dns_lb_health_check--examples--group-001.md#canonical-1100310202032332-3100202312011332-3023002312200212-2123103002011211-3032130332210023-3033222223332222-3012210210300001-2201111013023133)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_lb_health_check/data-source.tf`; digest `sha256:687cdf9541bb7302cd31e410ec561a4c9e1516b1f24c1acc1a51bcdfad32aecc`.

```terraform
# DNSLBHealthCheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBHealthCheck by name
data "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}

output "dns_lb_health_check_id" {
  value = data.xcsh_dns_lb_health_check.example.id
}
```
