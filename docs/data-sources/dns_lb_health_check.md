---
page_title: "xcsh_dns_lb_health_check landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check landing."
---

# xcsh_dns_lb_health_check landing

<a id="canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ced8171f072d60d0770e33db4c07612e3cc3ae125e5cdcc3b95aead7d708516d"></a>

## xcsh_dns_lb_health_check — xcsh_dns_lb_health_check / decb95282f04 / 2

Breadcrumbs:

- xcsh_dns_lb_health_check

Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a
error in F5 Distributed Cloud.

<a id="canonical-36a710b4f20d213162a8102580adbdeb408909e5949a11fc1afa378412361fbe"></a>

## Prerequisites — xcsh_dns_lb_health_check / decb95282f04 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-04af70c745006e1d244350be0883d1d3e6d943a8ad11e8689cd390fb2029b8b2"></a>

## Minimal configuration — xcsh_dns_lb_health_check / decb95282f04 / 4

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

<a id="canonical-7c778a0f9b0fb32967047d324338f21ec30b29be67729767d297cd078f80e313"></a>

## Root configuration — xcsh_dns_lb_health_check / decb95282f04 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3bf276cbb7fdb9947220c31226ad090f0422053016f4968b5fb954af3e3923ad"></a>

## Next pages — xcsh_dns_lb_health_check / decb95282f04 / 6

- [Property reference](../guides/data-sources--dns_lb_health_check--reference--group-001.md#canonical-d86ea89b46c68a54168d9f9992eb27baa50e8e624dc918265e3478da3270f50c)
- [Examples](../guides/data-sources--dns_lb_health_check--examples--group-001.md#canonical-50d223bed08b617ecb0b68269b4c2165ce73e90bcfaabfaac6924c01a15472df)
