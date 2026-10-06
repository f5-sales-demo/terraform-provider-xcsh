---
page_title: "xcsh_network_global_log_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_log_receiver examples."
---

# xcsh_network_global_log_receiver examples

<a id="canonical-1202301023031320-0303231001313120-3320311200012201-0033111103031111-3021010033100212-0120003332131331-3000011313101011-0011210102011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-0211012321131230-2013331210022130-2003210011323120-3212212113100323-3012123311310320-2030302302201323-1222230131301120-2020212031212211)
- Examples

<a id="canonical-3211201021021200-1230321333200313-0111022121112030-3110302321033221-0013110120023212-2231133302123221-0023231131330102-2202000312212012"></a>

### Complete configurations for `xcsh_network_global_log_receiver`

- [Data source](data-sources--network_global_log_receiver--examples--group-001.md#canonical-2023332312132030-2000303032010022-0221201301001233-3322023311121200-1122002033203110-0111320333210323-2313211303221202-0232212201013121): valid configuration.

<a id="canonical-2023332312132030-2000303032010022-0221201301001233-3322023311121200-1122002033203110-0111320333210323-2313211303221202-0232212201013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_global_log_receiver](../data-sources/network_global_log_receiver.md#canonical-0211012321131230-2013331210022130-2003210011323120-3212212113100323-3012123311310320-2030302302201323-1222230131301120-2020212031212211)
- [Examples](data-sources--network_global_log_receiver--examples--group-001.md#canonical-1202301023031320-0303231001313120-3320311200012201-0033111103031111-3021010033100212-0120003332131331-3000011313101011-0011210102011131)
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
