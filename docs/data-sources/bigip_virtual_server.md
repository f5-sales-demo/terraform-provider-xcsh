---
page_title: "xcsh_bigip_virtual_server landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_virtual_server landing."
---

# xcsh_bigip_virtual_server landing

<a id="canonical-5de83ca3dfb9cf73d57be7ed86f617e30ad06fc103f2f9653fda1240ecf86f88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24de24625eee055743109c5e39f3ceb57a92779c4ad2aa256c9a6fb04c0eca89"></a>

## xcsh_bigip_virtual_server — xcsh_bigip_virtual_server / f63bc12fcef6 / 2

Breadcrumbs:

- xcsh_bigip_virtual_server

Manages a BIG-IP Virtual Server resource in F5 Distributed Cloud for big-ip virtual server
specification. configuration. (read-only data source)

<a id="canonical-28ea7014f5ec6621d982eb4188228c1f4b6579c3c7049f844bd2adbdd74f239a"></a>

## Prerequisites — xcsh_bigip_virtual_server / f63bc12fcef6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-82c5c4eb9eb35320530f3150e829304b20edada6eadbfc2a01b4a7be32187e50"></a>

## Minimal configuration — xcsh_bigip_virtual_server / f63bc12fcef6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

<a id="canonical-e343a6382cc7ba4fd36968b521a6bdca0fe3a4791eb0a76d785d8de8e68c7c19"></a>

## Root configuration — xcsh_bigip_virtual_server / f63bc12fcef6 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9ef672cc3a7b71af2d2adf0492eb81c92150e831429369513dec1ba29ac47557"></a>

## Next pages — xcsh_bigip_virtual_server / f63bc12fcef6 / 6

- [Property reference](../guides/data-sources--bigip_virtual_server--reference--group-001.md#canonical-3c4f044f14de66769dee96461aaa25d45d7892c0fc761f5b76498185321c197f)
- [Examples](../guides/data-sources--bigip_virtual_server--examples--group-001.md#canonical-803e305ddcc14d26305861439a5a6225a5a00bb0c19fe90bd14bec7fbdf53bd9)
