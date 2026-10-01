---
page_title: "xcsh_cloud_elastic_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip examples."
---

# xcsh_cloud_elastic_ip examples

<a id="canonical-076152532c256f4736b20a7aad09ccaa23aa6c26028f74c80e32ace794cb1a1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48cd55c697f33cf31479015c86d8554cc0ec3653d26b243877a269e58475c672"></a>

## Examples — Examples / cdd75a35f0ff / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
- Examples

<a id="canonical-60d9e9eec9245aa479f6ede1348ea21722e8fb09b882eb1eb2611fbf598e4f34"></a>

## Complete configurations — Examples / cdd75a35f0ff / 3

- [Data source](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-e36c42338c7f410006646bc53d3713982a6af8475e85d679dbef8df67e923f7d): valid configuration.

<a id="canonical-6d15d8bd0377c0edaba156b2e96df80959ffd111daa97fe7cb477964648cd21f"></a>

## Next pages — Examples / cdd75a35f0ff / 4

- [Data source](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-e36c42338c7f410006646bc53d3713982a6af8475e85d679dbef8df67e923f7d)
- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)

<a id="canonical-e36c42338c7f410006646bc53d3713982a6af8475e85d679dbef8df67e923f7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9649fcf09a25b8343ac09cc1b4607605745982d71c092aa20dfa669b4e0a4596"></a>

## Data source — Data source / c8c5182da2d5 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
- [Examples](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-076152532c256f4736b20a7aad09ccaa23aa6c26028f74c80e32ace794cb1a1e)
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

<a id="canonical-d37a3b2afee8d13b69cfab339d64b26672da29739245444f426a2c6ea5cc5c9f"></a>

## Next pages — Data source / c8c5182da2d5 / 3

- [Examples](data-sources--cloud_elastic_ip--examples--group-001.md#canonical-076152532c256f4736b20a7aad09ccaa23aa6c26028f74c80e32ace794cb1a1e)
- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292)
