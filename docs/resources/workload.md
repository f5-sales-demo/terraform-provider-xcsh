---
page_title: "xcsh_workload landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload landing."
---

# xcsh_workload landing

<a id="canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110312003112200-0030020121102132-2100010303033220-0033130331110231-0013031111133230-3032031211023100-3230111302312030-3030220100300223"></a>

## xcsh_workload — xcsh_workload / 011322030030 / 2

Breadcrumbs:

- xcsh_workload

Manages a Workload resource in F5 Distributed Cloud for workload. configuration.

<a id="canonical-0000033002332223-0020111332032131-1303233222223120-1303200210332213-1200003120232313-3321312021023220-3133321112102232-3110100212032331"></a>

## Prerequisites — xcsh_workload / 011322030030 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_k8s`.

- virtual_k8s: Namespace for workload deployment

<a id="canonical-2132100131320303-3310130121112120-0123232000013320-1332131202203330-0232201003202021-2223133121330020-2010322020221233-0101330222230313"></a>

## Minimal configuration — xcsh_workload / 011322030030 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Workload Resource Example
# Manages a Workload resource in F5 Distributed Cloud for workload.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Workload configuration
resource "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}
```

<a id="canonical-1330122233121011-2130221200032202-1112023203311330-0100301020322102-1323223333202223-0300020022031322-2101331120000201-1003213132333210"></a>

## Root configuration — xcsh_workload / 011322030030 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1331111121302031-0320101100032122-2221310331220201-2230131003212220-1222310020020121-2010331002302031-0111233311320211-2211210312012031"></a>

## Next pages — xcsh_workload / 011322030030 / 6

- [Property reference](../guides/resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [Examples](../guides/resources--workload--examples--group-001.md#canonical-3201002123330130-0130300023013311-3322300132010333-1032200112231302-3120201212103100-0310100313110131-1230202231003130-0112021013220320)
- [Import](../guides/resources--workload--lifecycle--group-001.md#canonical-3211222301121232-1020210312011003-0231102321331130-1110200130120010-3131300210013301-2331311113011020-3211110001033120-2021101202123230)
- [Timeouts](../guides/resources--workload--lifecycle--group-001.md#canonical-2000311130222102-3030213013223321-0231303002010223-1102020011301223-2301320010112010-0031301311003121-1330000201303030-2301100211222111)
