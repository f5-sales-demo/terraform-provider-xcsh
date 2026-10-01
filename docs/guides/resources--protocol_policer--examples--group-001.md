---
page_title: "xcsh_protocol_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer examples."
---

# xcsh_protocol_policer examples

<a id="canonical-d8eecb4e2fce187ac1c664867ddead2a9474a00809319f7a26614fa8502ffd78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71aafec049a9021d33c3841925d89ac4052814b8a6973118f2133e915f7a8d5e"></a>

## Examples — Examples / b7937d374741 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- Examples

<a id="canonical-ec82348e56e02c9070156539aeefdebc1ea605e6480af73f97f2784441193e00"></a>

## Complete configurations — Examples / b7937d374741 / 3

- [Resource](resources--protocol_policer--examples--group-001.md#canonical-9ab057120bca3d890e43f924a39e414c2ae4e2dd79299bef1707f63366d4d9be): valid configuration.

<a id="canonical-da97efa6edec55b81d706ce3f16845ffeab6b796525d321af2dafef6af919a23"></a>

## Next pages — Examples / b7937d374741 / 4

- [Resource](resources--protocol_policer--examples--group-001.md#canonical-9ab057120bca3d890e43f924a39e414c2ae4e2dd79299bef1707f63366d4d9be)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-9ab057120bca3d890e43f924a39e414c2ae4e2dd79299bef1707f63366d4d9be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1789c8f1362e790cf49506cb929d62a2e9cb7a106c961e28ba3bf2c1e7285ed"></a>

## Resource — Resource / 6644a47e2fd3 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Examples](resources--protocol_policer--examples--group-001.md#canonical-d8eecb4e2fce187ac1c664867ddead2a9474a00809319f7a26614fa8502ffd78)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_policer/resource.tf`; digest `sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c`.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```

<a id="canonical-0485618db71af29ca052ab9ef8219b1156f5832ab52e55dcf4b9e61606c46f82"></a>

## Next pages — Resource / 6644a47e2fd3 / 3

- [Examples](resources--protocol_policer--examples--group-001.md#canonical-d8eecb4e2fce187ac1c664867ddead2a9474a00809319f7a26614fa8502ffd78)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
