---
page_title: "xcsh_nginx_service_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery landing."
---

# xcsh_nginx_service_discovery landing

<a id="canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34aa3ac8216e6940f27bbda6ed546a4e3ead6709d693d08795e0241839eab9e9"></a>

## xcsh_nginx_service_discovery — xcsh_nginx_service_discovery / 8b1a11a820c8 / 2

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

<a id="canonical-db91003af90acab99be7d4821987514a12294bb088c333bc989129376be3fe20"></a>

## Prerequisites — xcsh_nginx_service_discovery / 8b1a11a820c8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-247b455981ddadc3e588625b92c8d4d2bc929ac72746f9a52b928849152d00e7"></a>

## Minimal configuration — xcsh_nginx_service_discovery / 8b1a11a820c8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```

<a id="canonical-73c5d7acb0f74c8bab62d50c065c6b9c6be13bb13a8594608ef3dbd73f05be9c"></a>

## Root configuration — xcsh_nginx_service_discovery / 8b1a11a820c8 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c24e841476d74e6481ef02b9780cdf8bd8690e2102265d08ecf12699e344d768"></a>

## Next pages — xcsh_nginx_service_discovery / 8b1a11a820c8 / 6

- [Property reference](../guides/data-sources--nginx_service_discovery--reference--group-001.md#canonical-bd7faee5a6863c243227640fd4c42e49997c33eab641c67241f2503c658612e6)
- [Examples](../guides/data-sources--nginx_service_discovery--examples--group-001.md#canonical-766ba811139fdb5dc6fb3bb8d510b3535c22c310114d846a562ea2bab8392701)
