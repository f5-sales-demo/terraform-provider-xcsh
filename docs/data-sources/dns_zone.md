---
page_title: "xcsh_dns_zone"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone."
---

# xcsh_dns_zone

<a id="canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_zone

Reads an existing DNS zone in its namespace.

<a id="canonical-2223200121332322-1102120230011223-3212330230221302-2111230012231232-3100300022313321-3312112202323332-2001320301211312-1330010233112313"></a>

### Prerequisites for `xcsh_dns_zone`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

<a id="canonical-2232203022211023-1122003310021303-2022013033201320-0322203012113330-3232010211213013-3111233003330100-1312313230110322-3303133121122002"></a>

### Minimal configuration for `xcsh_dns_zone`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1121022113312033-3112212301312233-3320212110001211-2012003230332232-1301032201032103-3310003110101002-1201123231123130-1211312311033111"></a>

### Root configuration for `xcsh_dns_zone`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1311310132311313-3330131112310012-0330231120303132-3211312003303130-3021201302210320-3122202202233303-2222121122100202-3120210213311210"></a>

### Explore this collection for `xcsh_dns_zone`

- [Property reference](../guides/data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [Examples](../guides/data-sources--dns_zone--examples--group-001.md#canonical-0113332031001202-0203203323031122-1110313201121223-0211122311121012-1002330222223130-0033020333320010-2320222023212020-2311211332221300)
