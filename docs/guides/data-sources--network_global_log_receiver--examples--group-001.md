---
page_title: "xcsh_network_global_log_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_log_receiver examples."
---

# xcsh_network_global_log_receiver examples

<a id="canonical-62c4b37833b41dd8f8d601a10f553355c910f426180fe77dc01774450591215d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e58492606ce7f8371529958cd4cb93e9075182e6ad7f26e90bb5df12a2036986"></a>

## Examples — Examples / daa120eb997d / 2

Breadcrumbs:

- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-251b976c87f6429c83905ed8e699743bc66f5d388ccb287b6ab1dc588898d9a5)
- Examples

<a id="canonical-38d044fef4ccf4b9aa4ba64f9f8c4794655884292b693fe096b32b35532be77f"></a>

## Complete configurations — Examples / daa120eb997d / 3

- [Data source](data-sources--network_global_log_receiver--examples--group-001.md#canonical-8bfb678c80cce10a2987106ffa2f56605a08f8d415e3f93bb7973a622e9a11d9): valid configuration.

<a id="canonical-1afa40cfb2cd0b365e11dd54ad10f2073dae5dc33aeeb9f13298cf895c17cabb"></a>

## Next pages — Examples / daa120eb997d / 4

- [Data source](data-sources--network_global_log_receiver--examples--group-001.md#canonical-8bfb678c80cce10a2987106ffa2f56605a08f8d415e3f93bb7973a622e9a11d9)
- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-251b976c87f6429c83905ed8e699743bc66f5d388ccb287b6ab1dc588898d9a5)

<a id="canonical-8bfb678c80cce10a2987106ffa2f56605a08f8d415e3f93bb7973a622e9a11d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b43b4a2289c7371e2cc7592b912a11cd2aaa59f8bcee165163f0832e74f2f27c"></a>

## Data source — Data source / 480c2f81118c / 2

Breadcrumbs:

- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-251b976c87f6429c83905ed8e699743bc66f5d388ccb287b6ab1dc588898d9a5)
- [Examples](data-sources--network_global_log_receiver--examples--group-001.md#canonical-62c4b37833b41dd8f8d601a10f553355c910f426180fe77dc01774450591215d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_log_receiver/data-source.tf`; digest `sha256:e6481bf54911d3f66c740203a948d54d5c759ec25e22c585e36cf6e6833c987c`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_log_receiver" "receivers" {}

# This example chooses TLS syslog on TCP 6514. The manifest supplies only
# destinations; choose the port required by the configured log receiver.
output "tls_syslog_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 6514
    destinations = data.xcsh_network_global_log_receiver.receivers.cidr_blocks
  }
}
```

<a id="canonical-8e45e5a8c0ff834bcaba5a3cb0c1010da1bae583b802f8f599482d2e2257a336"></a>

## Next pages — Data source / 480c2f81118c / 3

- [Examples](data-sources--network_global_log_receiver--examples--group-001.md#canonical-62c4b37833b41dd8f8d601a10f553355c910f426180fe77dc01774450591215d)
- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-251b976c87f6429c83905ed8e699743bc66f5d388ccb287b6ab1dc588898d9a5)
