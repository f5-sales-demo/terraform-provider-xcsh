---
page_title: "xcsh_tcp_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer examples."
---

# xcsh_tcp_loadbalancer examples

<a id="canonical-65b66562cddaab262db2cae975842adac58ad22257dcba28c2897932fad6d3c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55008a51f0f9ce45a42b782e8b855c0208b656fd7596b1889df98e933d935f49"></a>

## Examples — Examples / 1568d769bf29 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- Examples

<a id="canonical-300d247e18c5d753e2ad34a01d9cbb5ac80c5de275a8232640cffeceb7ffd9fb"></a>

## Complete configurations — Examples / 1568d769bf29 / 3

- [Data source](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-2f8636178cea39e319aa0a01505f73025ed517f42308c0d3e91fe855ad401104): valid configuration.

<a id="canonical-7776cb0657a7380fed0ceccb6c4a3f3dc3550a7f5964faf5351b1c6a5e84712f"></a>

## Next pages — Examples / 1568d769bf29 / 4

- [Data source](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-2f8636178cea39e319aa0a01505f73025ed517f42308c0d3e91fe855ad401104)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-2f8636178cea39e319aa0a01505f73025ed517f42308c0d3e91fe855ad401104"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-113342be4d9a83adaef168ca5255a4bd30ce3edaa3565a4e3d424df0a376a793"></a>

## Data source — Data source / ec6c5e8c2f8f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Examples](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-65b66562cddaab262db2cae975842adac58ad22257dcba28c2897932fad6d3c7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tcp_loadbalancer/data-source.tf`; digest `sha256:8aadc1f29b777e99791395cd3d6a0a6d69101be8986adffcec19d0c93ccfd1f9`.

```terraform
# TCPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TCPLoadBalancer by name
data "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}

output "tcp_loadbalancer_id" {
  value = data.xcsh_tcp_loadbalancer.example.id
}
```

<a id="canonical-9d179fe48ac67478a0e25760d68708da9113f3234a0afdd2ee40a7904350ccf5"></a>

## Next pages — Data source / ec6c5e8c2f8f / 3

- [Examples](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-65b66562cddaab262db2cae975842adac58ad22257dcba28c2897932fad6d3c7)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
