---
page_title: "xcsh_alert_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver examples."
---

# xcsh_alert_receiver examples

<a id="canonical-2322202131233023-2131223013300313-1311001032100001-2201231332310233-0302203023203330-0120223112233121-2021102111032332-0233300211000222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- Examples

<a id="canonical-3021032122212102-3021123200021002-1120020032100303-3211021103133121-1110230122030232-3011313012011310-3031020131012213-3202223331333222"></a>

### Complete configurations for `xcsh_alert_receiver`

- [Resource](resources--alert_receiver--examples--group-001.md#canonical-0302211312132003-2302010323020101-0131303320113013-1001122332121003-1012202002120331-1200011020020323-3122133110000031-2221032233231212): valid configuration.

<a id="canonical-0302211312132003-2302010323020101-0131303320113013-1001122332121003-1012202002120331-1200011020020323-3122133110000031-2221032233231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311)
- [Examples](resources--alert_receiver--examples--group-001.md#canonical-2322202131233023-2131223013300313-1311001032100001-2201231332310233-0302203023203330-0120223112233121-2021102111032332-0233300211000222)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_receiver/resource.tf`; digest `sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb`.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```
