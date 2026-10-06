---
page_title: "xcsh_log_receiver examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver examples."
---

# xcsh_log_receiver examples

<a id="canonical-2223011232000332-3200032103323313-3320133300222033-2022032003330100-2001123102211220-2202121202312110-3031022223312223-3300202332132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- Examples

<a id="canonical-3121312231100033-1101330000033321-3313310111012121-0022212213131323-3030012113222200-3001110103000001-0321123130000012-2032233130130020"></a>

### Complete configurations for `xcsh_log_receiver`

- [Data source](data-sources--log_receiver--examples--group-001.md#canonical-2030222112202302-3033131333302233-1303122302031311-1310102120010133-1333110320202033-1302220231202210-1123000021221030-0311313021022132): valid configuration.

<a id="canonical-2030222112202302-3033131333302233-1303122302031311-1310102120010133-1333110320202033-1302220231202210-1123000021221030-0311313021022132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Examples](data-sources--log_receiver--examples--group-001.md#canonical-2223011232000332-3200032103323313-3320133300222033-2022032003330100-2001123102211220-2202121202312110-3031022223312223-3300202332132331)
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
