---
page_title: "xcsh_flow_anomaly examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_flow_anomaly examples."
---

# xcsh_flow_anomaly examples

<a id="canonical-0dcd935591bd3fbe55ba3d6048f4471096c0b1a5141006cd2d6134cd31c1b720"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f91e5912343e5886904ff413aedec09d02283153f744fa1f7056b95138d9538"></a>

## Examples — Examples / f9b788aba948 / 2

Breadcrumbs:

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-36b34809215da1235dc2f83cf637368d9b15f5ad28182217e93d829fbb145231)
- Examples

<a id="canonical-ab20432aaaa7f752ada1a992c31b6d910a24617b15e9e8c983b4b861019164e7"></a>

## Complete configurations — Examples / f9b788aba948 / 3

- [Data source](data-sources--flow_anomaly--examples--group-001.md#canonical-52058c0741e7e1aaf2c740a7262ec9997a23f0690aeeb15b1a2bab2e768fbf36): valid configuration.

<a id="canonical-cdcecf84048e7483ae0118a1b88255f609d3beca73608654d536de1bd5999f12"></a>

## Next pages — Examples / f9b788aba948 / 4

- [Data source](data-sources--flow_anomaly--examples--group-001.md#canonical-52058c0741e7e1aaf2c740a7262ec9997a23f0690aeeb15b1a2bab2e768fbf36)
- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-36b34809215da1235dc2f83cf637368d9b15f5ad28182217e93d829fbb145231)

<a id="canonical-52058c0741e7e1aaf2c740a7262ec9997a23f0690aeeb15b1a2bab2e768fbf36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c90b242e2f4bb2232a4f37cd8d90f7a28bba8afab91a2399907c35d1aa5cfd2"></a>

## Data source — Data source / 853d0e80ddb6 / 2

Breadcrumbs:

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-36b34809215da1235dc2f83cf637368d9b15f5ad28182217e93d829fbb145231)
- [Examples](data-sources--flow_anomaly--examples--group-001.md#canonical-0dcd935591bd3fbe55ba3d6048f4471096c0b1a5141006cd2d6134cd31c1b720)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_flow_anomaly/data-source.tf`; digest `sha256:d39bc93b81bc691cec5167bb0d331b58de96886d72cd9a1fb46d22490728b1ca`.

```terraform
# FlowAnomaly Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FlowAnomaly by name
data "xcsh_flow_anomaly" "example" {
  name      = "example-flow-anomaly"
  namespace = "staging"
}

output "flow_anomaly_id" {
  value = data.xcsh_flow_anomaly.example.id
}
```

<a id="canonical-481e2ca4c8366221bd393367584ad2756eaf9fc686d2c247ea186e505863bb71"></a>

## Next pages — Data source / 853d0e80ddb6 / 3

- [Examples](data-sources--flow_anomaly--examples--group-001.md#canonical-0dcd935591bd3fbe55ba3d6048f4471096c0b1a5141006cd2d6134cd31c1b720)
- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-36b34809215da1235dc2f83cf637368d9b15f5ad28182217e93d829fbb145231)
