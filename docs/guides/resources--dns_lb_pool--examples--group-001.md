---
page_title: "xcsh_dns_lb_pool examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool examples."
---

# xcsh_dns_lb_pool examples

<a id="canonical-3021310113213213-0120010022310331-0010123011100200-0333131131321313-1322121133200223-1032031202330333-3023111233300100-2320221233303000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- Examples

<a id="canonical-0310003003031122-3202030022212333-0013233223013312-3022322213221330-1333111312012211-0121100212231231-2113320000211111-0122111310011123"></a>

### Complete configurations for `xcsh_dns_lb_pool`

- [Resource](resources--dns_lb_pool--examples--group-001.md#canonical-0131031232300121-2212232023211202-2122022301200203-0330133212020000-3310112012123003-2232233201322301-1220111133231000-0121033120100332): valid configuration.

<a id="canonical-0131031232300121-2212232023211202-2122022301200203-0330133212020000-3310112012123003-2232233201322301-1220111133231000-0121033120100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332)
- [Examples](resources--dns_lb_pool--examples--group-001.md#canonical-3021310113213213-0120010022310331-0010123011100200-0333131131321313-1322121133200223-1032031202330333-3023111233300100-2320221233303000)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_pool/resource.tf`; digest `sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763`.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```
