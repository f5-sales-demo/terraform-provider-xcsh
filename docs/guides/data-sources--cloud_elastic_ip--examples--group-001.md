---
page_title: "xcsh_cloud_elastic_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip examples."
---

# xcsh_cloud_elastic_ip examples

<a id="canonical-0013120111021103-0230021112331013-0312230200221322-2231002130302222-0203222212300212-0002203313103020-0032030222303213-2110302301220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102)
- Examples

<a id="canonical-1020303111113012-2113330303303303-0110132100011130-2012312011111030-3000323003121103-3102122302100320-1313220212213211-2010131130121302"></a>

### Complete configurations for `xcsh_cloud_elastic_ip`

- [Data source](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-3203123010020303-2030133310010000-0012121012233011-0331031301032120-0222122233201013-1132201131121321-3123323320313312-1332210203331331): valid configuration.

<a id="canonical-3203123010020303-2030133310010000-0012121012233011-0331031301032120-0222122233201013-1132201131121321-3123323320313312-1332210203331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102)
- [Examples](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-0013120111021103-0230021112331013-0312230200221322-2231002130302222-0203222212300212-0002203313103020-0032030222303213-2110302301220132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_elastic_ip/data-source.tf`; digest `sha256:a222dbb0119e3d807c71bf21695db7199a7a0767160af1c0ed21e3f33eb7b545`.

```terraform
# CloudElasticIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudElasticIP by name
data "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"
}

output "cloud_elastic_ip_id" {
  value = data.xcsh_cloud_elastic_ip.example.id
}
```
