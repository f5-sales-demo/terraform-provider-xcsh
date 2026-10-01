---
page_title: "xcsh_global_log_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver examples."
---

# xcsh_global_log_receiver examples

<a id="canonical-5c26fc0d1d2807236ae02af52707999e5e595a8aae44687b256637a6e9951aaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a37af3b277fa1d3f04b424af1b0edf5e4cdfd848b3ec13c5bbd60db63b384698"></a>

## Examples — Examples / 3fa133e41ac6 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- Examples

<a id="canonical-6c4fc73ad78c619020f12ac66da6270b44ba59790a43a3e2dc7b350d09c54793"></a>

## Complete configurations — Examples / 3fa133e41ac6 / 3

- [Resource](resources--global_log_receiver--examples--group-001.md#canonical-729034fd2c4d33f4bc68c5d86a9a0b944d7f942af43426560cb0a3f765102419): valid configuration.

<a id="canonical-31489bfaec5d671bd288d165f045b38f710062e9ff4593d647bf3b9aa229c10f"></a>

## Next pages — Examples / 3fa133e41ac6 / 4

- [Resource](resources--global_log_receiver--examples--group-001.md#canonical-729034fd2c4d33f4bc68c5d86a9a0b944d7f942af43426560cb0a3f765102419)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)

<a id="canonical-729034fd2c4d33f4bc68c5d86a9a0b944d7f942af43426560cb0a3f765102419"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d6f8bde40f5e6514087116ca35d9cdc98df4b74b7436870d05e679e13f2ab63"></a>

## Resource — Resource / ea28125e5321 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
- [Examples](resources--global_log_receiver--examples--group-001.md#canonical-5c26fc0d1d2807236ae02af52707999e5e595a8aae44687b256637a6e9951aaf)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_global_log_receiver/resource.tf`; digest `sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802`.

```terraform
# GlobalLogReceiver Resource Example
# Manages new Global Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GlobalLogReceiver configuration
resource "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}
```

<a id="canonical-40a65d801933ee4e301711f925403f4e8615ec46819d312edc2b488af06c6a44"></a>

## Next pages — Resource / ea28125e5321 / 3

- [Examples](resources--global_log_receiver--examples--group-001.md#canonical-5c26fc0d1d2807236ae02af52707999e5e595a8aae44687b256637a6e9951aaf)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-1e880e7c7bdf1c11b455555a94902bc778fe398e7ddf9f5ea18b9a24baf3c5e5)
