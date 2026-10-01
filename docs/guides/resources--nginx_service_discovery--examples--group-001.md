---
page_title: "xcsh_nginx_service_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery examples."
---

# xcsh_nginx_service_discovery examples

<a id="canonical-c18b13d10d0226163fc8cf1360174bc98d4ef06d5d4daa9e87bea575065b2b6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53846a5fa4c4229a981fc627ec4b4a8a6b05362a74f3cc145f43728ebf7a997f"></a>

## Examples — Examples / b65faed422c8 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- Examples

<a id="canonical-8291e2d5ad40a19d1064a7a05c45c05d3c7612d69e85ce29a76dac15d1f4a659"></a>

## Complete configurations — Examples / b65faed422c8 / 3

- [Resource](resources--nginx_service_discovery--examples--group-001.md#canonical-23c2226dc5353a550b8ae567aca1d92e7169de7f7311a5e2b5f531759a269df5): valid configuration.

<a id="canonical-fcc10e64437aa25a89a926fcc3ba3496d6b2bf8b92ede31cee168b4f6b5961b9"></a>

## Next pages — Examples / b65faed422c8 / 4

- [Resource](resources--nginx_service_discovery--examples--group-001.md#canonical-23c2226dc5353a550b8ae567aca1d92e7169de7f7311a5e2b5f531759a269df5)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-23c2226dc5353a550b8ae567aca1d92e7169de7f7311a5e2b5f531759a269df5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-818a4c889fb96f9f37753345e0946aef8c0a181ece33a5ba0f1450d0f6fc774f"></a>

## Resource — Resource / e68bb7501c17 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Examples](resources--nginx_service_discovery--examples--group-001.md#canonical-c18b13d10d0226163fc8cf1360174bc98d4ef06d5d4daa9e87bea575065b2b6a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nginx_service_discovery/resource.tf`; digest `sha256:dd43ba5ab8c8a702c201e30638f18d6c0eafcf265b5180bf65bbcaa03289e5e6`.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```

<a id="canonical-cc5e260faed8f94ba72fa0ba58bad04a85b0b20124206ec940cfc152c4c021dd"></a>

## Next pages — Resource / e68bb7501c17 / 3

- [Examples](resources--nginx_service_discovery--examples--group-001.md#canonical-c18b13d10d0226163fc8cf1360174bc98d4ef06d5d4daa9e87bea575065b2b6a)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
