---
page_title: "xcsh_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver landing."
---

# xcsh_global_log_receiver landing

<a id="canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33fc6bdd09fdf34427f300fe58242323c3655efbc1967dfe257d3e6e2c687fd1"></a>

## xcsh_global_log_receiver — xcsh_global_log_receiver / 2e6d53a8cd43 / 2

Breadcrumbs:

- xcsh_global_log_receiver

Manages new Global Log Receiver object in F5 Distributed Cloud.

<a id="canonical-fc6d3dc1a7027a2ac0ea86ec8587c3587c82a81fa7ecc44a0969603c2c196f80"></a>

## Prerequisites — xcsh_global_log_receiver / 2e6d53a8cd43 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-885d1b00e565f65fd2237bc68b5dbbabcc0acf2112e57e63386886e8bd342698"></a>

## Minimal configuration — xcsh_global_log_receiver / 2e6d53a8cd43 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GlobalLogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GlobalLogReceiver by name
data "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}

output "global_log_receiver_id" {
  value = data.xcsh_global_log_receiver.example.id
}
```

<a id="canonical-52cb1f39ddfec97a7e26c6fab8de31c518aa2921d36bbe3a61a8940a9db43d51"></a>

## Root configuration — xcsh_global_log_receiver / 2e6d53a8cd43 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-093512a44859e26683562daafbce0c9e25fa5d75446c3d57b410993a6517059a"></a>

## Next pages — xcsh_global_log_receiver / 2e6d53a8cd43 / 6

- [Property reference](../guides/data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [Examples](../guides/data-sources--global_log_receiver--examples--group-001.md#canonical-cbd6e8041fc1a2b6d29872815c4d2cbae71d5823e5f69ef6270a7d868231f463)
