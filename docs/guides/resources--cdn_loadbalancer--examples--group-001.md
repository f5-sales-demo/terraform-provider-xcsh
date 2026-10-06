---
page_title: "xcsh_cdn_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer examples."
---

# xcsh_cdn_loadbalancer examples

<a id="canonical-2223333103031111-2331013121000300-2231102333110312-1231112211103303-1003012200322202-1023032201331100-3322113303302113-1213200102031112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- Examples

<a id="canonical-3023213333301020-3213012211020200-1232030101002020-3223330321333121-2020022122021002-3012131022003331-0233231211021020-3132130110230113"></a>

### Complete configurations for `xcsh_cdn_loadbalancer`

- [Resource](resources--cdn_loadbalancer--examples--group-001.md#canonical-2233222001213003-1100302333323121-2230023101203200-1030021131132032-2210331310200012-1001001122001323-3312010220131211-0113301100112130): valid configuration.

<a id="canonical-2233222001213003-1100302333323121-2230023101203200-1030021131132032-2210331310200012-1001001122001323-3312010220131211-0113301100112130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Examples](resources--cdn_loadbalancer--examples--group-001.md#canonical-2223333103031111-2331013121000300-2231102333110312-1231112211103303-1003012200322202-1023032201331100-3322113303302113-1213200102031112)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_loadbalancer/resource.tf`; digest `sha256:9dbe7d10d66d8f67145599515660dfb082a890fef657898fc4e2002f5c72f6dd`.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```
