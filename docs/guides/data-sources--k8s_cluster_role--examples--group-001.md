---
page_title: "xcsh_k8s_cluster_role examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role examples."
---

# xcsh_k8s_cluster_role examples

<a id="canonical-2020200223133133-1213011132112100-3201302032210003-2331332220121210-0033322213031201-3101003002030200-1103001321300103-2312010013201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- Examples

<a id="canonical-0011330323023210-0030230032203211-0203222323320331-3300322012332003-3013313302100321-2030311322223131-0102002022210033-0220202310021210"></a>

### Complete configurations for `xcsh_k8s_cluster_role`

- [Data source](data-sources--k8s_cluster_role--examples--group-001.md#canonical-1212100012103223-0023010010321332-0031131100331110-2013103323033232-3211320211320212-1220111002313301-0232010321000002-0001231221323331): valid configuration.

<a id="canonical-1212100012103223-0023010010321332-0031131100331110-2013103323033232-3211320211320212-1220111002313301-0232010321000002-0001231221323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122)
- [Examples](data-sources--k8s_cluster_role--examples--group-001.md#canonical-2020200223133133-1213011132112100-3201302032210003-2331332220121210-0033322213031201-3101003002030200-1103001321300103-2312010013201210)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster_role/data-source.tf`; digest `sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112`.

```terraform
# K8SClusterRole Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRole by name
data "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.xcsh_k8s_cluster_role.example.id
}
```
