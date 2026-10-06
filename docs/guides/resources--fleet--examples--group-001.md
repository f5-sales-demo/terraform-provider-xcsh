---
page_title: "xcsh_fleet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet examples."
---

# xcsh_fleet examples

<a id="canonical-3300202033332320-3213223220133221-1232320113200122-3233113313222131-1230323102323110-0100312311110323-1122133332121301-1012301011332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- Examples

<a id="canonical-1112123220320332-2120012332123233-2330301312131000-3220023310313300-0002230112022113-3210333012002223-3033030103032013-0312333102130312"></a>

### Complete configurations for `xcsh_fleet`

- [Resource](resources--fleet--examples--group-001.md#canonical-2212131230021011-2120022302100122-3213222103232310-2322313112333112-0111200130300320-2002101113202230-0202111303112333-1310122332322322): valid configuration.

<a id="canonical-2212131230021011-2120022302100122-3213222103232310-2322313112333112-0111200130300320-2002101113202230-0202111303112333-1310122332322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Examples](resources--fleet--examples--group-001.md#canonical-3300202033332320-3213223220133221-1232320113200122-3233113313222131-1230323102323110-0100312311110323-1122133332121301-1012301011332211)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fleet/resource.tf`; digest `sha256:4d32cab1b63bbbd4184a815cf7049069724ae120ad87141da3ed9def32e6d8f3`.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```
