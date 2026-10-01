---
page_title: "xcsh_flow_anomaly landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_flow_anomaly landing."
---

# xcsh_flow_anomaly landing

<a id="canonical-0312230310200021-0201113122010203-1131300233200330-3312031303122031-2123011133112231-0220012002020113-3221033120022133-2323011011020301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301000113300302-2223001211312210-0210201323023321-3133123303103113-0033013200023313-2232221113331113-0212302100030301-0310102232310212"></a>

## xcsh_flow_anomaly — xcsh_flow_anomaly / 312230211100 / 2

Breadcrumbs:

- xcsh_flow_anomaly

Manages a Flow Anomaly resource in F5 Distributed Cloud for flow anomaly specification.
configuration. (read-only data source)

<a id="canonical-3311001323312010-3111202201123211-1220022002032220-2001222112233323-3010003203002223-3011310010021211-2232130101333233-3233210013013221"></a>

## Prerequisites — xcsh_flow_anomaly / 312230211100 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1010010202223102-0002322330222012-0321202212233321-0330303102032202-0032101013231332-1333131113031012-2331212211021210-0330231120013100"></a>

## Minimal configuration — xcsh_flow_anomaly / 312230211100 / 4

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

<a id="canonical-2222331120010102-3201221130221121-0102320230311320-2323203022202233-2302022122300022-0332123230203320-2331122103230133-3110210100201032"></a>

## Root configuration — xcsh_flow_anomaly / 312230211100 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2222323133301001-1032301131332213-1220130010130232-0022010300003203-3200013223213231-3103001201131121-2102323220022231-2110233210202100"></a>

## Next pages — xcsh_flow_anomaly / 312230211100 / 6

- [Property reference](../guides/data-sources--flow_anomaly--reference--group-001.md#canonical-1001030133000121-1133022331020113-1013220120013233-1103310331100231-0302200102302313-0033012100301122-3101322032301330-3222201320301020)
- [Examples](../guides/data-sources--flow_anomaly--examples--group-001.md#canonical-0031303121031111-2101233103332332-1111232203311200-1020331010130100-2112300023012211-0110010000123031-0231120103103031-0301300123130200)
