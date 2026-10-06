---
page_title: "xcsh_alert_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver examples."
---

# xcsh_alert_receiver examples

<a id="canonical-2130120101331130-2120000310312101-0313131112010322-3013100020120003-0221123111021020-1300012221211231-3110312120331222-0223310311022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- Examples

<a id="canonical-1202213123332133-0332303221133023-2302120001232023-2011310103010202-0012103011110322-2232012330333310-3010123003102123-2121121303131220"></a>

### Complete configurations for `xcsh_alert_receiver`

- [Data source](data-sources--alert_receiver--examples--group-001.md#canonical-3221110320020111-3312110013000223-1332222210300322-1033301303313312-1122112222311311-1303332201220113-0031310322112310-2103120203003001): valid configuration.

<a id="canonical-3221110320020111-3312110013000223-1332222210300322-1033301303313312-1122112222311311-1303332201220113-0031310322112310-2103120203003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102)
- [Examples](data-sources--alert_receiver--examples--group-001.md#canonical-2130120101331130-2120000310312101-0313131112010322-3013100020120003-0221123111021020-1300012221211231-3110312120331222-0223310311022310)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_receiver/data-source.tf`; digest `sha256:96027faa6721d178ff8fb480c66120ec2d45036c783323915216275133e6a1d6`.

```terraform
# AlertReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertReceiver by name
data "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}

output "alert_receiver_id" {
  value = data.xcsh_alert_receiver.example.id
}
```
