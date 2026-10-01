---
page_title: "xcsh_code_base_integration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration landing."
---

# xcsh_code_base_integration landing

<a id="canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8529acce506a373f96f49eaad1bcaaf33673b95db2a38bf532175dbadd373820"></a>

## xcsh_code_base_integration — xcsh_code_base_integration / 81bcaf2beb45 / 2

Breadcrumbs:

- xcsh_code_base_integration

Manages integration details in F5 Distributed Cloud.

<a id="canonical-4aec17149e94f1a15e178ff8e9dc64480b2c9afbca5a3755c9a5d9f5b1558dc8"></a>

## Prerequisites — xcsh_code_base_integration / 81bcaf2beb45 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0fc5daab14222ae45b7652b9797ed5a1b828d3160732c5d41c30644e90fb9d0a"></a>

## Minimal configuration — xcsh_code_base_integration / 81bcaf2beb45 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

<a id="canonical-005e909bd37a2b692a4af812c1558c1b3259a9d873300380c15f4d60cd95f883"></a>

## Root configuration — xcsh_code_base_integration / 81bcaf2beb45 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-19f8077a713498df1ee0eea5192969e228072268db266652766095f9569bcc96"></a>

## Next pages — xcsh_code_base_integration / 81bcaf2beb45 / 6

- [Property reference](../guides/resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [Examples](../guides/resources--code_base_integration--examples--group-001.md#canonical-ebd765c530052f7dfefad865c74b821bdddd2ce9c13805dcfc9c9bdea1b87240)
- [Import](../guides/resources--code_base_integration--lifecycle--group-001.md#canonical-e9857ac1ed338d3232223089b0ed5225775440e3451a7d20239c3c190e4506a7)
- [Timeouts](../guides/resources--code_base_integration--lifecycle--group-001.md#canonical-2a574d849df264fa58f692225dd9561bc9e61888ae655e0bcb3d2ab2686d1468)
