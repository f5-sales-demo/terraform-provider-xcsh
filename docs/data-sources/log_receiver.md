---
page_title: "xcsh_log_receiver landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver landing."
---

# xcsh_log_receiver landing

<a id="canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3a6f08afb02db1c34a30efe79cd9fc74065b5610ad98d0d30d7653d4b520977"></a>

## xcsh_log_receiver — xcsh_log_receiver / eb42317e6cf9 / 2

Breadcrumbs:

- xcsh_log_receiver

Manages new Log Receiver object in F5 Distributed Cloud.

<a id="canonical-e425c4408934a2db3a504b34fa88582b9ee8a7c7d0d40732d17aadd0d8ca8624"></a>

## Prerequisites — xcsh_log_receiver / eb42317e6cf9 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-cbe0581ec73e0dea760dd8148166db31335a44de85fff7a9da55cee9d6f28cf3"></a>

## Minimal configuration — xcsh_log_receiver / eb42317e6cf9 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LogReceiver by name
data "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}

output "log_receiver_id" {
  value = data.xcsh_log_receiver.example.id
}
```

<a id="canonical-993e503f892d2de1d239620757f228c3409566c03793af0756a55b811f873257"></a>

## Root configuration — xcsh_log_receiver / eb42317e6cf9 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-035a0a1f6cddea60dd3600e5e78bc6c0bdf4a94a7131670bcb8856340f761b11"></a>

## Next pages — xcsh_log_receiver / eb42317e6cf9 / 6

- [Property reference](../guides/data-sources--log_receiver--reference--group-001.md#canonical-72ddf5904e939117996ebb0df85159b509473e049321d2b4345c22a8d61761af)
- [Examples](../guides/data-sources--log_receiver--examples--group-001.md#canonical-ab16e03ee0393ef7f87f0a8f8a383f10816d2968a2662d94cd2abdabf08be7bd)
