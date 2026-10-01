---
page_title: "xcsh_flow_anomaly landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_flow_anomaly landing."
---

# xcsh_flow_anomaly landing

<a id="canonical-36b34809215da1235dc2f83cf637368d9b15f5ad28182217e93d829fbb145231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1017c32ab065da42487b2f9df6f34d70f1e02f7aea57f5726c90331344aed26"></a>

## xcsh_flow_anomaly — xcsh_flow_anomaly / f372d3dac950 / 2

Breadcrumbs:

- xcsh_flow_anomaly

Manages a Flow Anomaly resource in F5 Distributed Cloud for flow anomaly specification.
configuration. (read-only data source)

<a id="canonical-f507bd84d58a16e5682823a881a96bfbc40e30abc5d04265ae711fefef9071e9"></a>

## Prerequisites — xcsh_flow_anomaly / f372d3dac950 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-44122ad202ebca86398a6bf93ccd23a20e447b7e7f757346bd9a52643cb581d0"></a>

## Minimal configuration — xcsh_flow_anomaly / f372d3dac950 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-aaf58112e1a5ca5912e2cd78bb8ca8afb229ac0a3e6ec8f8bd693b1fd491084e"></a>

## Root configuration — xcsh_flow_anomaly / f372d3dac950 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aaedfc414ec5dfa76870472e0a1300e3e01eb9edd306175992ee82ad94be4890"></a>

## Next pages — xcsh_flow_anomaly / f372d3dac950 / 6

- [Property reference](../guides/data-sources--flow_anomaly--reference--group-001.md#canonical-4131f0195f2bd21747a181ef53d3d42d32812cb70f190c5ad1e8ec7cea878c48)
- [Examples](../guides/data-sources--flow_anomaly--examples--group-001.md#canonical-0dcd935591bd3fbe55ba3d6048f4471096c0b1a5141006cd2d6134cd31c1b720)
