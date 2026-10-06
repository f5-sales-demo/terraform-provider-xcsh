---
page_title: "xcsh_service_policy_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_set examples."
---

# xcsh_service_policy_set examples

<a id="canonical-1210220030301313-1112121100102103-1032120221121012-3133300310001112-1132200310032010-1032303010220113-3230331113010111-2023023333130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-1102322222212201-0323000120102213-0210100111122123-0013000010002101-1001031323000300-3302313333113330-3123230211030001-2322203121212312)
- Examples

<a id="canonical-3001223221032201-0332210301020302-0310110213102303-0230111021000022-0232231210300033-0300200331212112-3330032023131021-2000002011022232"></a>

### Complete configurations for `xcsh_service_policy_set`

- [Data source](data-sources--service_policy_set--examples--group-001.md#canonical-0330102310230122-1223203021122031-0203023200302031-0313123001221103-2231000301310011-2303023113023020-2021003211310330-1333032122312302): valid configuration.

<a id="canonical-0330102310230122-1223203021122031-0203023200302031-0313123001221103-2231000301310011-2303023113023020-2021003211310330-1333032122312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md#canonical-1102322222212201-0323000120102213-0210100111122123-0013000010002101-1001031323000300-3302313333113330-3123230211030001-2322203121212312)
- [Examples](data-sources--service_policy_set--examples--group-001.md#canonical-1210220030301313-1112121100102103-1032120221121012-3133300310001112-1132200310032010-1032303010220113-3230331113010111-2023023333130212)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_set/data-source.tf`; digest `sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a`.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```
