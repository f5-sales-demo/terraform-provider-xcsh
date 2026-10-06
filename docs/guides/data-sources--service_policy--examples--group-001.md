---
page_title: "xcsh_service_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy examples."
---

# xcsh_service_policy examples

<a id="canonical-2310032000213002-0013023100320201-1122323331333212-3213302213103111-0003331221123320-2033213121331022-2322303123023213-0001233132023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- Examples

<a id="canonical-2011301001220202-3130220112223112-1100031002022231-1103233301131231-0123330103000221-2101033031111111-0010332200330201-0112121101221030"></a>

### Complete configurations for `xcsh_service_policy`

- [Data source](data-sources--service_policy--examples--group-001.md#canonical-3133021133101133-1331033230023211-3202032232200331-2023300021323301-0130231331111303-2301211203303113-2131112300200131-1331020312233013): valid configuration.

<a id="canonical-3133021133101133-1331033230023211-3202032232200331-2023300021323301-0130231331111303-2301211203303113-2131112300200131-1331020312233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Examples](data-sources--service_policy--examples--group-001.md#canonical-2310032000213002-0013023100320201-1122323331333212-3213302213103111-0003331221123320-2033213121331022-2322303123023213-0001233132023000)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy/data-source.tf`; digest `sha256:9189dd7384a926dd448e27affcd3a13f4fb5015e7cdfa0bb42501f00fe48f58d`.

```terraform
# ServicePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicy by name
data "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}

output "service_policy_id" {
  value = data.xcsh_service_policy.example.id
}
```
