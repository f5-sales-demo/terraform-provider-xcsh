---
page_title: "xcsh_smsv2_aws_runtime examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_aws_runtime examples."
---

# xcsh_smsv2_aws_runtime examples

<a id="canonical-0133013233122312-1321111213003103-0302002301312333-3233212133322132-1120031213220132-2300020203302100-0031003332102011-2112330331200013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-1120113232301113-0213023023132120-1120122010310202-0111213100202201-0123213030213220-2303201220111233-2220132332120003-3013101222121310)
- Examples

<a id="canonical-1320101123123231-2332112332101233-1330211013010320-0202013231120122-2111333211332223-0231201223120131-2011130330213220-3220112030303331"></a>

### Complete configurations for `xcsh_smsv2_aws_runtime`

- [Data source](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-1220032022202012-2231333113102233-0210321323030212-0110131031312222-2033303300200222-3221110300202013-0031301303112132-2100130123303312): valid configuration.

<a id="canonical-1220032022202012-2231333113102233-0210321323030212-0110131031312222-2033303300200222-3221110300202013-0031301303112132-2100130123303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-1120113232301113-0213023023132120-1120122010310202-0111213100202201-0123213030213220-2303201220111233-2220132332120003-3013101222121310)
- [Examples](data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-0133013233122312-1321111213003103-0302002301312333-3233212133322132-1120031213220132-2300020203302100-0031003332102011-2112330331200013)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_aws_runtime/data-source.tf`; digest `sha256:e299ec8ab27b0aebf0736cf9e2b336fdc7524df91d37a6ad7dac1a9e71242dd5`.

```terraform
# Correlate stable logical node keys and AWS-authoritative ENI MAC addresses
# with the SMSv2 interface configuration and runtime health observed by F5 XC.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_aws_runtime" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  nodes = {
    node_0_slo = {
      node = "node-0"
      role = "slo"
      mac  = "02:00:00:00:00:10"
    }
    node_0_sli = {
      node = "node-0"
      role = "sli"
      mac  = "02:00:00:00:00:11"
    }
  }
}

output "smsv2_interfaces" {
  value = data.xcsh_smsv2_aws_runtime.site.interfaces
}

output "smsv2_healthy" {
  value = data.xcsh_smsv2_aws_runtime.site.healthy
}
```
