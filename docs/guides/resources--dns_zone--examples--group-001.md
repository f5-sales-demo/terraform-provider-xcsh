---
page_title: "xcsh_dns_zone examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone examples."
---

# xcsh_dns_zone examples

<a id="canonical-3202030322001231-1303203113122030-3132231122312030-2110213333302302-0210313030331003-2330021132233031-2332122210201220-1113202010312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- Examples

<a id="canonical-1202210022012010-2001103003010213-3131111232232320-3323122021113301-2123311130220232-0211333330033301-2312200002312233-3131121201032111"></a>

### Complete configurations for `xcsh_dns_zone`

- [Resource](resources--dns_zone--examples--group-001.md#canonical-3033303200211321-2113102023332300-3033203001220231-1023130201113300-1122313032231021-1210123232321120-0312031112110211-0311333033133100): valid configuration.

<a id="canonical-3033303200211321-2113102023332300-3033203001220231-1023130201113300-1122313032231021-1210123232321120-0312031112110211-0311333033133100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Examples](resources--dns_zone--examples--group-001.md#canonical-3202030322001231-1303203113122030-3132231122312030-2110213333302302-0210313030331003-2330021132233031-2332122210201220-1113202010312131)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_zone/resource.tf`; digest `sha256:a65c60c4a2a466885fcf90c9e747d8f8ba0052be1693d774ae212f433fd43631`.

```terraform
# DNSZone Resource Example
# Manages DNS Zone in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSZone configuration
resource "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"

  primary {
    allow_http_lb_managed_records = true
  }
}
```
