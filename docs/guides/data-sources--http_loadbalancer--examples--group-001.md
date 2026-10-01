---
page_title: "xcsh_http_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer examples."
---

# xcsh_http_loadbalancer examples

<a id="canonical-864faaea9ccf4e93d6c60b5343dae9fe756da6f5fa5779fee9023c494681a6d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79430b8b42dcf3381d4a4125eb5f5e81e6dd3fc8793b57e58e920612d28dd83c"></a>

## Examples — Examples / baf3dd8b634f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- Examples

<a id="canonical-55abfa476139a9ab2258a8afe47db86d4a02d596b8bfa5545946e23b51fb73f9"></a>

## Complete configurations — Examples / baf3dd8b634f / 3

- [Data source](data-sources--http_loadbalancer--examples--group-001.md#canonical-dcc1fc365a9c8cab61bfb58aae2b409f2475dead84302ab8b875d098ec8beddc): valid configuration.

<a id="canonical-06d885d5556ea79086f2f753de80d9b49fdc58f9bb076d322cc40b94b5a11223"></a>

## Next pages — Examples / baf3dd8b634f / 4

- [Data source](data-sources--http_loadbalancer--examples--group-001.md#canonical-dcc1fc365a9c8cab61bfb58aae2b409f2475dead84302ab8b875d098ec8beddc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dcc1fc365a9c8cab61bfb58aae2b409f2475dead84302ab8b875d098ec8beddc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbf093cbe01c1094d453c9a2800f2bf3b74910ec02fbeeb119575ecdf910b551"></a>

## Data source — Data source / 7eb5d7ebbf6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Examples](data-sources--http_loadbalancer--examples--group-001.md#canonical-864faaea9ccf4e93d6c60b5343dae9fe756da6f5fa5779fee9023c494681a6d4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_http_loadbalancer/data-source.tf`; digest `sha256:f7d3fa4bb60b803988578a72b5dac4a76f304bac2925ae4fdfe075100bed37ad`.

```terraform
# HTTPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing HTTPLoadBalancer by name
data "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"
}

output "http_loadbalancer_id" {
  value = data.xcsh_http_loadbalancer.example.id
}
```

<a id="canonical-06e1e6ea775c5d888d9ef0e383ebf634ef458edd1bda08dee6d8b6b0d8601037"></a>

## Next pages — Data source / 7eb5d7ebbf6c / 3

- [Examples](data-sources--http_loadbalancer--examples--group-001.md#canonical-864faaea9ccf4e93d6c60b5343dae9fe756da6f5fa5779fee9023c494681a6d4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
