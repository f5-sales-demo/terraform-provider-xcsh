---
page_title: "xcsh_virtual_host landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host landing."
---

# xcsh_virtual_host landing

<a id="canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab0af4543bbc46570b780fb5f6e40dc6981b83117508ad134bc6545994f735f"></a>

## xcsh_virtual_host — xcsh_virtual_host / f5109a52e8f3 / 2

Breadcrumbs:

- xcsh_virtual_host

Manages virtual host in a given namespace in F5 Distributed Cloud.

<a id="canonical-d143caa652a3d88ba5653db41e25083e9bc8d09efbcd2bd4a79b79417fc0c797"></a>

## Prerequisites — xcsh_virtual_host / f5109a52e8f3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-451510beb3a020ad4bcd779b3d699f0928132b6405c4537232628af2fbb53fad"></a>

## Minimal configuration — xcsh_virtual_host / f5109a52e8f3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```

<a id="canonical-a0bcbf64dab9852f13f1614a9895f3241f189e7b6887ba43b4b1e24c30eae5ec"></a>

## Root configuration — xcsh_virtual_host / f5109a52e8f3 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-589e714adc5d7d7629931cff171bab63a59ee5910fb6ea25b6c0362e6af7c4c2"></a>

## Next pages — xcsh_virtual_host / f5109a52e8f3 / 6

- [Property reference](../guides/data-sources--virtual_host--reference--group-001.md#canonical-7d0b64e615839ee4869f74f3a3ecb15ca62c06c786a0aed14e53a7a714f35a5d)
- [Examples](../guides/data-sources--virtual_host--examples--group-001.md#canonical-1007d5b46c405c4dac828dcf0244f383fa0b2d87d812c2533acf250ab93f0856)
