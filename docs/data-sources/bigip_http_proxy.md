---
page_title: "xcsh_bigip_http_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy landing."
---

# xcsh_bigip_http_proxy landing

<a id="canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4251202174fc36d643734bea21c561beccf6c7501668706f6d98edf959894f47"></a>

## xcsh_bigip_http_proxy — xcsh_bigip_http_proxy / bddb4ff5f084 / 2

Breadcrumbs:

- xcsh_bigip_http_proxy

Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5
Distributed Cloud.

<a id="canonical-2c5cf37be00d9beb26085c41e0f4caaff87be96cc413d4b213bd876debdab991"></a>

## Prerequisites — xcsh_bigip_http_proxy / bddb4ff5f084 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1eac443d27660743577480a5bc6c43c9634cd10c075161035e5aefc6a14d43de"></a>

## Minimal configuration — xcsh_bigip_http_proxy / bddb4ff5f084 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPHTTPProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPHTTPProxy by name
data "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}

output "bigip_http_proxy_id" {
  value = data.xcsh_bigip_http_proxy.example.id
}
```

<a id="canonical-bfb0c1bc6cec75d2b542f803a7cd88026bc5844de27b4761904a058f04ca3779"></a>

## Root configuration — xcsh_bigip_http_proxy / bddb4ff5f084 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0738b3d4027e46464445af12012d7f0566b2876aef214d3ec945d3c674ca3a99"></a>

## Next pages — xcsh_bigip_http_proxy / bddb4ff5f084 / 6

- [Property reference](../guides/data-sources--bigip_http_proxy--reference--group-001.md#canonical-0eccd4618e2bc28472c640ae95271eb83457e947f382e2a5bf2e4c1264ae2d42)
- [Examples](../guides/data-sources--bigip_http_proxy--examples--group-001.md#canonical-66a4638e347f7e6208eb33c9e5df7fa32d2a01474868f7b8331496cc994da90c)
