---
page_title: "xcsh_voltstack_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site landing."
---

# xcsh_voltstack_site landing

<a id="canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320202121222033-1230313131031320-0122303020310230-0023023113000122-0122222201020123-0130221232123321-2120022022100321-0300200203233233"></a>

## xcsh_voltstack_site — xcsh_voltstack_site / 123133332223 / 2

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

<a id="canonical-3102131112033310-3100200132320131-1132120221221013-0312003031120030-1313122211002103-2031123032020211-1221311321323312-1102002220000022"></a>

## Prerequisites — xcsh_voltstack_site / 123133332223 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1031211211213221-2300033032132223-0323130310330131-3200013130103102-2032110210130113-1032120012133003-1313131002223201-0002012312210012"></a>

## Minimal configuration — xcsh_voltstack_site / 123133332223 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

<a id="canonical-1132010312302120-0331023312300003-1332130003303013-2113112201301301-3200210202310233-3031223102313201-2030123210213011-2100232133212132"></a>

## Root configuration — xcsh_voltstack_site / 123133332223 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0321112122312232-3230133311133113-2202012033332002-1100331202231100-3111320310232100-1123231020232331-1011220033332023-0211200002000301"></a>

## Next pages — xcsh_voltstack_site / 123133332223 / 6

- [Property reference](../guides/data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [Examples](../guides/data-sources--voltstack_site--examples--group-001.md#canonical-1303010302133333-0003310111003013-3020311221331011-2131123320112100-1121002131100132-3023311100320300-2032310133322301-1310303333231132)
