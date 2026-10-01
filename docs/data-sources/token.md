---
page_title: "xcsh_token landing"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token landing."
---

# xcsh_token landing

<a id="canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a43789e8a93e168a2d293493a0fac597c787a2ee4ef47ba0eaebbc2081a9cecf"></a>

## xcsh_token — xcsh_token / 1b0d4da1c5a3 / 2

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

<a id="canonical-057aa478537dd9c277d470b1d1f07cb7c3cc48c03c2993b3a8b6b20b69546ade"></a>

## Prerequisites — xcsh_token / 1b0d4da1c5a3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-f92b0a2734b555f7bd1351041835f4454a4a7a1c76e3fe4a5ad87702ff98b6a1"></a>

## Minimal configuration — xcsh_token / 1b0d4da1c5a3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```

<a id="canonical-50fad55f8b8fc8d5579a31437f6557cf8c4b5dfded1934a72e1e0281c9eff561"></a>

## Root configuration — xcsh_token / 1b0d4da1c5a3 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2df63e046f0e99e9e12a051271e5423370f06d4f269639a96e8dc9bd7877b0b3"></a>

## Next pages — xcsh_token / 1b0d4da1c5a3 / 6

- [Property reference](../guides/data-sources--token--reference--group-001.md#canonical-722430f94f35b0d2e9c17930c6fab80d7093f146d88e1f310120119163a9f05a)
- [Examples](../guides/data-sources--token--examples--group-001.md#canonical-48e55ab18f6d9f5efd96fc6d6200ed21e196eeca9c0b4d74d36b64f40d067778)
