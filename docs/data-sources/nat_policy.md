---
page_title: "xcsh_nat_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy landing."
---

# xcsh_nat_policy landing

<a id="canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79be24f873cf068cc4c44dc6c09fa2d83fb8a89a8fe56b29159c52271e1f6cb4"></a>

## xcsh_nat_policy — xcsh_nat_policy / 3ff1933d5caf / 2

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

<a id="canonical-9a894880b344d75d3fa37fa34ab219559429dcd5aba05b8b3479bd29c1397cff"></a>

## Prerequisites — xcsh_nat_policy / 3ff1933d5caf / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-303000cc35151dc233580df4ed3b18b797a944cfcb92b95de922ee7adf375c20"></a>

## Minimal configuration — xcsh_nat_policy / 3ff1933d5caf / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NATPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NATPolicy by name
data "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}

output "nat_policy_id" {
  value = data.xcsh_nat_policy.example.id
}
```

<a id="canonical-1a31ed2c037826616589410b7aad7b689128c0704e8d6dd004978061bf965ac4"></a>

## Root configuration — xcsh_nat_policy / 3ff1933d5caf / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-59898b853de5b079a32c1e1fb5c79e20a45c88ff6c81ccced61e2b10f05287d4"></a>

## Next pages — xcsh_nat_policy / 3ff1933d5caf / 6

- [Property reference](../guides/data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [Examples](../guides/data-sources--nat_policy--examples--group-001.md#canonical-29569badf4ac6533a49f84111b8f371cb2a24d15313f0d6ee8386e1a348410c6)
