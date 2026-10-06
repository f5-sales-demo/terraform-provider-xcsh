---
page_title: "xcsh_log_receiver examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver examples."
---

# xcsh_log_receiver examples

<a id="canonical-0101332232122303-0002030101233212-3100332320223031-3313332101002120-3223320021112202-0120333120203012-2101332223200021-3100030120332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- Examples

<a id="canonical-2201203002212001-2313323033230000-2000300203233000-2302013320310033-3013133020101022-3212300113333103-2231221122002230-2131220311210311"></a>

### Complete configurations for `xcsh_log_receiver`

- [Resource](resources--log_receiver--examples--group-001.md#canonical-1203001122000212-1132303023100012-3031323130113333-1231323300332032-3101211023032323-2020121020301202-0331323333021131-1021011023231322): valid configuration.

<a id="canonical-1203001122000212-1132303023100012-3031323130113333-1231323300332032-3101211023032323-2020121020301202-0331323333021131-1021011023231322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003)
- [Examples](resources--log_receiver--examples--group-001.md#canonical-0101332232122303-0002030101233212-3100332320223031-3313332101002120-3223320021112202-0120333120203012-2101332223200021-3100030120332330)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_log_receiver/resource.tf`; digest `sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d`.

```terraform
# LogReceiver Resource Example
# Manages new Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic LogReceiver configuration
resource "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}
```
