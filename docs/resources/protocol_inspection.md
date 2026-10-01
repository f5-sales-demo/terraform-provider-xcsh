---
page_title: "xcsh_protocol_inspection landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection landing."
---

# xcsh_protocol_inspection landing

<a id="canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aed4749ae4457329e17438cc00f6a30d774d6dd355e88dd305ed920205f37ddf"></a>

## xcsh_protocol_inspection — xcsh_protocol_inspection / 44f022b8f9f6 / 2

Breadcrumbs:

- xcsh_protocol_inspection

Manages Protocol Inspection Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-4a36dc84a481f7f8e66acce548d7f37f1721c1c91d943fc01f5c2ba32bb99d97"></a>

## Prerequisites — xcsh_protocol_inspection / 44f022b8f9f6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c52f2f1ed72a8813a8923ebbcf6d166887d819291c4078b2d5274d2ac2a11475"></a>

## Minimal configuration — xcsh_protocol_inspection / 44f022b8f9f6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```

<a id="canonical-629c3ae00fe0454a4700a4f46a188c91767af2215a40e7ce7cdabd429f881270"></a>

## Root configuration — xcsh_protocol_inspection / 44f022b8f9f6 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e3feaf4fd874a207d3048b9a106097d408b0acf73ca952026bce4bb371bd0ec2"></a>

## Next pages — xcsh_protocol_inspection / 44f022b8f9f6 / 6

- [Property reference](../guides/resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [Examples](../guides/resources--protocol_inspection--examples--group-001.md#canonical-daa07a474b657dbd62339295ce1f9bb6e4e70833d976326d48460822d296ef4f)
- [Import](../guides/resources--protocol_inspection--lifecycle--group-001.md#canonical-49e74acc2efc92347f83c9fbfaa202d79c981fe9f204a191bef4ac39c40e649b)
- [Timeouts](../guides/resources--protocol_inspection--lifecycle--group-001.md#canonical-3cc7b5304d8da3de2835b6f2803fe8cd1d321bc60e7bd2279e7900fee659dd68)
