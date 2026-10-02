---
page_title: "xcsh_public_ip_binding landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip_binding landing."
---

# xcsh_public_ip_binding landing

<a id="canonical-0203231102121100-0021222023322132-1033002220322133-0210023131031121-3331122113020013-3010312110213323-0102222031221012-2102213113003031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301021211201000-0332022133001032-0212303120101133-2212001100210120-1030123020303331-1111301101100030-2212321312313103-3101012301233332"></a>

## xcsh_public_ip_binding — xcsh_public_ip_binding / 002210212031 / 2

Breadcrumbs:

- xcsh_public_ip_binding

Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the
binding; deletion restores its captured original binding. This resource never allocates, deallocates
or deletes the public IP.

<a id="canonical-3003313133322131-0010332001330211-0320012032102211-0311321133321213-0112233320203221-3313232223200003-0123320200220313-3310133211212131"></a>

## Prerequisites — xcsh_public_ip_binding / 002210212031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2102310302301130-2121302020232213-1330203031232311-1113302121221212-2201011131331023-3132233212320013-0132000030033110-3110201302133232"></a>

## Minimal configuration — xcsh_public_ip_binding / 002210212031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PublicIPBinding Resource Example
# Manage the regional virtual-site binding of an already allocated public IP.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PublicIPBinding configuration
resource "xcsh_public_ip_binding" "example" {
  name      = "example-public-ip-binding"
  namespace = "staging"

  expected_ip            = "example-value"
  virtual_site           = "example-value"
  virtual_site_namespace = "example-value"
}
```

<a id="canonical-3230313230322311-0112021333300012-0031303131231333-0013311221121001-1033111301030103-2112122023222110-2221002230012021-1302013001312331"></a>

## Root configuration — xcsh_public_ip_binding / 002210212031 / 5

Required root properties: `expected_ip`, `name`, `namespace`, `virtual_site`, `virtual_site_namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1313112313122111-1321022002010132-0110130321311233-3122103300222232-1311133001212321-1300320322201201-2103011103101310-1031310212130131"></a>

## Next pages — xcsh_public_ip_binding / 002210212031 / 6

- [Property reference](../guides/resources--public_ip_binding--reference--group-001.md#canonical-0000123303013230-2022101212032211-3303121103322023-1113210013222113-2332132123021221-0201100212311212-3322201001001113-1203020333320300)
- [Examples](../guides/resources--public_ip_binding--examples--group-001.md#canonical-3200320122123303-1130300212320023-3232122301302102-2221120230202213-0322231312213330-3311300112112302-3000101330021310-3012112203332131)
