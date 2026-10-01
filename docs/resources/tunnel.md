---
page_title: "xcsh_tunnel landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel landing."
---

# xcsh_tunnel landing

<a id="canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8da67c76d56dea336d3b72b9d1abd387c2566e23242d1babcfa3d85c21e24f02"></a>

## xcsh_tunnel — xcsh_tunnel / 62b486b97f2f / 2

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-4b5ba0ca2c6005f304e4fb88af492022b712f275d320f7564bd092733018a907"></a>

## Prerequisites — xcsh_tunnel / 62b486b97f2f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b9a4dc64efcf5e7aea14c0acc50892d030eb8be74ffc8124c2839321ff673fd3"></a>

## Minimal configuration — xcsh_tunnel / 62b486b97f2f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

<a id="canonical-09a4034357a2a64f47d6705ca7e7c8ccea2e0c1bffd8a32919d4133b733cb174"></a>

## Root configuration — xcsh_tunnel / 62b486b97f2f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9c70ed40e77e4149f9c8c7b8c3e20da09f7cf0f37076a4c9a393676726312dc3"></a>

## Next pages — xcsh_tunnel / 62b486b97f2f / 6

- [Property reference](../guides/resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [Examples](../guides/resources--tunnel--examples--group-001.md#canonical-b70f6360292460f47d62d7c611fcc6bc888c2d4745820fdb7eea37be934eda5d)
- [Import](../guides/resources--tunnel--lifecycle--group-001.md#canonical-4476f878e34235afd0c963eb24321e83bcdb97296f8fcdfec6e93cd17cd9ec04)
- [Timeouts](../guides/resources--tunnel--lifecycle--group-001.md#canonical-6c0c7853f8b86f9b02598dc790580451be606dac6f7f848b20dd6891f10d6d5d)
