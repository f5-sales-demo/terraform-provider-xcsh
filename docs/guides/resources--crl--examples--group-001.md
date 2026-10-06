---
page_title: "xcsh_crl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl examples."
---

# xcsh_crl examples

<a id="canonical-3310232032200130-0113023102123231-0210011231122312-1033220100003200-0021120321310011-2301211233010312-1100231101320111-2312112221212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
- Examples

<a id="canonical-2011002303033020-1223311302123203-2012032113311000-1331123022131123-3011111203122312-2111303123021312-1301333302032321-0031303100132223"></a>

### Complete configurations for `xcsh_crl`

- [Resource](resources--crl--examples--group-001.md#canonical-1222233233230320-2223120312320112-2110210320101320-1002302222132120-0230233333021011-3312133333020310-0132322131300302-2012300022303313): valid configuration.

<a id="canonical-1222233233230320-2223120312320112-2110210320101320-1002302222132120-0230233333021011-3312133333020310-0132322131300302-2012300022303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
- [Examples](resources--crl--examples--group-001.md#canonical-3310232032200130-0113023102123231-0210011231122312-1033220100003200-0021120321310011-2301211233010312-1100231101320111-2312112221212132)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_crl/resource.tf`; digest `sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3`.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```
