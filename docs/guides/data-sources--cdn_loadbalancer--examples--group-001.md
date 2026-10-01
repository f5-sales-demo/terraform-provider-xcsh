---
page_title: "xcsh_cdn_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer examples."
---

# xcsh_cdn_loadbalancer examples

<a id="canonical-d465eef705caed1a30c11c7311fd6668fc60bad15c23d46b23c44c5afc61bab2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63eec822bcaaeddf307abdef8f30f35237c3df7b26ddabbdf9f36e86d635949f"></a>

## Examples — Examples / 2c008a6fc0fa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- Examples

<a id="canonical-853675f642a6590e4363e9281c4edd4a82663bcf3819eb9ec8da865ce3352c64"></a>

## Complete configurations — Examples / 2c008a6fc0fa / 3

- [Data source](data-sources--cdn_loadbalancer--examples--group-001.md#canonical-dbda87c790ac85ecd3b190760fbeb9e160d1a7945ee12a7c27eb04cf20813dc6): valid configuration.

<a id="canonical-85fe4d6a6990e808be1d7243bbd3902ede32bc0fe2f559039b7746891aeddd9f"></a>

## Next pages — Examples / 2c008a6fc0fa / 4

- [Data source](data-sources--cdn_loadbalancer--examples--group-001.md#canonical-dbda87c790ac85ecd3b190760fbeb9e160d1a7945ee12a7c27eb04cf20813dc6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dbda87c790ac85ecd3b190760fbeb9e160d1a7945ee12a7c27eb04cf20813dc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3466b82ad6bf5d5376994fa2dc4e472921cec49c986641024a527da4f80b99d"></a>

## Data source — Data source / f021a4719157 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Examples](data-sources--cdn_loadbalancer--examples--group-001.md#canonical-d465eef705caed1a30c11c7311fd6668fc60bad15c23d46b23c44c5afc61bab2)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_loadbalancer/data-source.tf`; digest `sha256:065d14809a577f5f36891bc47ab4a39d45652a5da6b5d83b3e19904703a77bc5`.

```terraform
# CDNLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNLoadBalancer by name
data "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"
}

output "cdn_loadbalancer_id" {
  value = data.xcsh_cdn_loadbalancer.example.id
}
```

<a id="canonical-af9c75be6b542b74392d183102936f339451a4f49907f0d0d8c8b4d90fcdf73b"></a>

## Next pages — Data source / f021a4719157 / 3

- [Examples](data-sources--cdn_loadbalancer--examples--group-001.md#canonical-d465eef705caed1a30c11c7311fd6668fc60bad15c23d46b23c44c5afc61bab2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
