---
page_title: "xcsh_flow_anomaly examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_flow_anomaly examples."
---

# xcsh_flow_anomaly examples

<a id="canonical-0031303121031111-2101233103332332-1111232203311200-1020331010130100-2112300023012211-0110010000123031-0231120103103031-0301300123130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-0312230310200021-0201113122010203-1131300233200330-3312031303122031-2123011133112231-0220012002020113-3221033120022133-2323011011020301)
- Examples

<a id="canonical-1333210132112101-0203100332112020-1221001033331001-0322323132300021-3100020220030111-0333131010332201-3313001112232111-0103203121110320"></a>

### Complete configurations for `xcsh_flow_anomaly`

- [Data source](data-sources--flow_anomaly--examples--group-001.md#canonical-1102001120300013-1001321332012222-3302301310002213-0212023230212121-1322020333001221-0022323223011123-0122022322230232-1312203323330312): valid configuration.

<a id="canonical-1102001120300013-1001321332012222-3302301310002213-0212023230212121-1322020333001221-0022323223011123-0122022322230232-1312203323330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md#canonical-0312230310200021-0201113122010203-1131300233200330-3312031303122031-2123011133112231-0220012002020113-3221033120022133-2323011011020301)
- [Examples](data-sources--flow_anomaly--examples--group-001.md#canonical-0031303121031111-2101233103332332-1111232203311200-1020331010130100-2112300023012211-0110010000123031-0231120103103031-0301300123130200)
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
