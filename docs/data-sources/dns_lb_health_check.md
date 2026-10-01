---
page_title: "xcsh_dns_lb_health_check landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check landing."
---

# xcsh_dns_lb_health_check landing

<a id="canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032312001130133-0013023112003100-1313003203033123-1030001312010232-0330300322320102-1132113031303003-2321112232223113-3113002011011231"></a>

## xcsh_dns_lb_health_check — xcsh_dns_lb_health_check / 022002330010 / 2

Breadcrumbs:

- xcsh_dns_lb_health_check

Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a
error in F5 Distributed Cloud.

<a id="canonical-0312221301002310-3302003102010301-1202222001000211-2000223123313223-1000202100213211-2110212201013330-0122332203132010-0102031201332332"></a>

## Prerequisites — xcsh_dns_lb_health_check / 022002330010 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0010223313003013-1011000012320131-0210100311002332-0020200331013103-3212312110032220-2231010132201220-2130310321003323-0200022123202302"></a>

## Minimal configuration — xcsh_dns_lb_health_check / 022002330010 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1330131320220033-2123003323030221-1213001013310302-1003032033020132-3003002302212332-1213130221131213-3102211330310013-2033200032030103"></a>

## Root configuration — xcsh_dns_lb_health_check / 022002330010 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0323330213123023-2313333123212110-1302020030030102-0212223100210033-0010020200110300-0112331021122023-1133232111102233-0332032102032231"></a>

## Next pages — xcsh_dns_lb_health_check / 022002330010 / 6

- [Property reference](../guides/data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [Examples](../guides/data-sources--dns_lb_health_check--examples--group-001.md#canonical-1100310202032332-3100202312011332-3023002312200212-2123103002011211-3032130332210023-3033222223332222-3012210210300001-2201111013023133)
