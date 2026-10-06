---
page_title: "xcsh_virtual_k8s examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s examples."
---

# xcsh_virtual_k8s examples

<a id="canonical-3003212211022232-3322110222230213-0231321321013322-1203302321122120-1333320121200133-0312320223001013-0332223203022313-1331301003223120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- Examples

<a id="canonical-3320302002212310-1300120131122201-3102200012312213-3323120333120201-0212112002121213-2203301302003201-2022010201111022-1201023210301102"></a>

### Complete configurations for `xcsh_virtual_k8s`

- [Data source](data-sources--virtual_k8s--examples--group-001.md#canonical-1220312122123020-3330103200030130-2312131021103011-3010313313132313-1102122321123330-0230231102220120-2220001102200031-1013300102120123): valid configuration.

<a id="canonical-1220312122123020-3330103200030130-2312131021103011-3010313313132313-1102122321123330-0230231102220120-2220001102200031-1013300102120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- [Examples](data-sources--virtual_k8s--examples--group-001.md#canonical-3003212211022232-3322110222230213-0231321321013322-1203302321122120-1333320121200133-0312320223001013-0332223203022313-1331301003223120)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_k8s/data-source.tf`; digest `sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d`.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```
