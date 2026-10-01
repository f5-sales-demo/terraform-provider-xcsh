---
page_title: "xcsh_virtual_k8s landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s landing."
---

# xcsh_virtual_k8s landing

<a id="canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032233012031332-3302100020002203-0310210001113111-0111320313213110-3133133230100010-1020311132210103-1130012003121201-1200013003101331"></a>

## xcsh_virtual_k8s — xcsh_virtual_k8s / 100210230220 / 2

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-3333012223130010-2013301101311120-2010013012123231-0102130022331231-3230300222310021-3000310131101023-0013222213121031-1022033312001010"></a>

## Prerequisites — xcsh_virtual_k8s / 100210230220 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

<a id="canonical-2120202332222200-2020031120332310-2023301203022213-2120121100122100-1210122021101030-0001002131013333-3330223331023020-0132323213033301"></a>

## Minimal configuration — xcsh_virtual_k8s / 100210230220 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```

<a id="canonical-1111003221121310-3120323120103313-2103313003132033-0212200033200012-1103322032230333-0332330030210003-0211300302123112-1003111100131011"></a>

## Root configuration — xcsh_virtual_k8s / 100210230220 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2213321102030013-3133321002103020-3330100121133233-3110013110323222-3321303303202101-2323333000302120-2213201113323201-0201122030323221"></a>

## Next pages — xcsh_virtual_k8s / 100210230220 / 6

- [Property reference](../guides/resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- [Examples](../guides/resources--virtual_k8s--examples--group-001.md#canonical-2321313332101022-1203232122231303-0133012333202120-0131022023203202-1233113033322022-0320211011333031-0111110210103001-1010200302220301)
- [Import](../guides/resources--virtual_k8s--lifecycle--group-001.md#canonical-0210330301020310-1301021323100000-2131300013132312-3320113213110132-2212310131312000-1210333120230202-1012303221101223-0311213103022233)
- [Timeouts](../guides/resources--virtual_k8s--lifecycle--group-001.md#canonical-2230113013220122-2012332000102030-2001013212230203-1000122000103321-2022111002102331-3213221222220312-3303322020213331-0102211012333233)
