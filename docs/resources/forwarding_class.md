---
page_title: "xcsh_forwarding_class landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class landing."
---

# xcsh_forwarding_class landing

<a id="canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c76fd4ff6d86cafc395ec7287bd491733395df38a59af0d0303512af82728692"></a>

## xcsh_forwarding_class — xcsh_forwarding_class / 193e0abc93d8 / 2

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

<a id="canonical-251a8928a191a8449bba46213face721c9a1bd814e91c3a19d2bac014b5cb631"></a>

## Prerequisites — xcsh_forwarding_class / 193e0abc93d8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6822deb96949b2dcb0495406547cd3edbe5bb321739d2c98eb1f558fbf618e7b"></a>

## Minimal configuration — xcsh_forwarding_class / 193e0abc93d8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```

<a id="canonical-23cdb1859a0d8f127428b03282e94565888c49ac76be1401ee52013b44e3b5c3"></a>

## Root configuration — xcsh_forwarding_class / 193e0abc93d8 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4119b592a5774924db6937c97122675ef47c2fb96439f550c3f1d9f617783cc9"></a>

## Next pages — xcsh_forwarding_class / 193e0abc93d8 / 6

- [Property reference](../guides/resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [Examples](../guides/resources--forwarding_class--examples--group-001.md#canonical-5955f4f6761160e682dae3b90ae4fc0c70c41e30a8dcd6aaaefdab6915e072ea)
- [Import](../guides/resources--forwarding_class--lifecycle--group-001.md#canonical-6547cf28bafd92feb766ef14dac5e9fce04cf59db97ec6605fb2a23310534e1d)
- [Timeouts](../guides/resources--forwarding_class--lifecycle--group-001.md#canonical-f50c8b329fb289cd6b5ee18052e078467715dfe30a13fc6fd67666440e16fabb)
