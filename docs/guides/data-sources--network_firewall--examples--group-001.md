---
page_title: "xcsh_network_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall examples."
---

# xcsh_network_firewall examples

<a id="canonical-66ef622e203ba224649c17f5e45fd3d55d87a3a0afb4f33e4afc95a5d0d63fdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd9465ead070712eb12a07e9dc2fba272b64322a28ec2ad1b1f979fcf02a9277"></a>

## Examples — Examples / f68bbd94110e / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- Examples

<a id="canonical-b4976572d5a02e52c1d0c201762da367f2e85d4879b8aabe9d9908e6f0fe4841"></a>

## Complete configurations — Examples / f68bbd94110e / 3

- [Data source](data-sources--network_firewall--examples--group-001.md#canonical-f0f939ff749d0b5b89e930648af76287bedc4cde6ca9cf1b9913e92c35e69a57): valid configuration.

<a id="canonical-820e567a7bbebf73f086c2b20693a63b267e8a03becc77c56f9102ef634c2f64"></a>

## Next pages — Examples / f68bbd94110e / 4

- [Data source](data-sources--network_firewall--examples--group-001.md#canonical-f0f939ff749d0b5b89e930648af76287bedc4cde6ca9cf1b9913e92c35e69a57)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)

<a id="canonical-f0f939ff749d0b5b89e930648af76287bedc4cde6ca9cf1b9913e92c35e69a57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-942872fa858bb81a5d4451417a16805b4e5f762716ab7d74c582768676316c6b"></a>

## Data source — Data source / c5d1d801d849 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
- [Examples](data-sources--network_firewall--examples--group-001.md#canonical-66ef622e203ba224649c17f5e45fd3d55d87a3a0afb4f33e4afc95a5d0d63fdd)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_firewall/data-source.tf`; digest `sha256:5641721d230eef959a292217d55745f5f8be26fc7ce1feff6b83cbca8e44625c`.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```

<a id="canonical-790618309030c21c09af066bf161fd0699c2c6b408d73f6c9e1019974395fd80"></a>

## Next pages — Data source / c5d1d801d849 / 3

- [Examples](data-sources--network_firewall--examples--group-001.md#canonical-66ef622e203ba224649c17f5e45fd3d55d87a3a0afb4f33e4afc95a5d0d63fdd)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-9b62b76cf7e299087b5f4928cbe980ffbde2f654430a3985cbcf52e5b4f58a56)
