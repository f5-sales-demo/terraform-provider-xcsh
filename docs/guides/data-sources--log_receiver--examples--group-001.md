---
page_title: "xcsh_log_receiver examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver examples."
---

# xcsh_log_receiver examples

<a id="canonical-ab16e03ee0393ef7f87f0a8f8a383f10816d2968a2662d94cd2abdabf08be7bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9dad40f51f003f9f7d151990a9a777bcc197aa0c1513001396dc0068ebdc708"></a>

## Examples — Examples / d6bc19e18d91 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- Examples

<a id="canonical-af89bab77d99eef7c629c6eb4bf3ef6d677e5d924a47ab0544826e489288f9cc"></a>

## Complete configurations — Examples / d6bc19e18d91 / 3

- [Data source](data-sources--log_receiver--examples--group-001.md#canonical-8ca968b2cf77fcaf736b23757449811f7f53888f72a2d8a45b009a4c35dc929e): valid configuration.

<a id="canonical-e4f842fb8bb377183621088e4075ef6afa4af40ab6cfeda98ac72a1e6e1b96a1"></a>

## Next pages — Examples / d6bc19e18d91 / 4

- [Data source](data-sources--log_receiver--examples--group-001.md#canonical-8ca968b2cf77fcaf736b23757449811f7f53888f72a2d8a45b009a4c35dc929e)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)

<a id="canonical-8ca968b2cf77fcaf736b23757449811f7f53888f72a2d8a45b009a4c35dc929e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-081bdd4633d69d3bc3449ee7e1f6e0ee4cb263a451f9258449f88d4a93e8ae95"></a>

## Data source — Data source / 29d3e88f9f2c / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
- [Examples](data-sources--log_receiver--examples--group-001.md#canonical-ab16e03ee0393ef7f87f0a8f8a383f10816d2968a2662d94cd2abdabf08be7bd)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_log_receiver/data-source.tf`; digest `sha256:e0ffa47f4ff4996df5615eeba45c356df21a5deecfea79894da1f56537fef3b2`.

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

<a id="canonical-bc8e45117914dee504a735ca6b4597b700c30a7619fc88ac6a9d17ec680fdfdb"></a>

## Next pages — Data source / 29d3e88f9f2c / 3

- [Examples](data-sources--log_receiver--examples--group-001.md#canonical-ab16e03ee0393ef7f87f0a8f8a383f10816d2968a2662d94cd2abdabf08be7bd)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-442ee10ad438a227b1ada43bcfe36c2313a2037062235cef8ab995cae280a352)
