---
page_title: "xcsh_service_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule examples."
---

# xcsh_service_policy_rule examples

<a id="canonical-0023101333202233-2312130003023013-2031120302010112-1223202002333100-1103023303201230-3132030320300230-2102312130202330-0111300211010120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- Examples

<a id="canonical-1020220133030110-0333232310301323-1103221001202332-1100202110020203-2212212010132320-3323111233120213-2121031123122000-2302302300113020"></a>

### Complete configurations for `xcsh_service_policy_rule`

- [Resource](resources--service_policy_rule--examples--group-001.md#canonical-0031012330311210-2200131123310002-3303003331230202-0221301312323212-1332133113013300-3031023113223131-2121002220121312-2020130021122221): valid configuration.

<a id="canonical-0031012330311210-2200131123310002-3303003331230202-0221301312323212-1332133113013300-3031023113223131-2121002220121312-2020130021122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Examples](resources--service_policy_rule--examples--group-001.md#canonical-0023101333202233-2312130003023013-2031120302010112-1223202002333100-1103023303201230-3132030320300230-2102312130202330-0111300211010120)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy_rule/resource.tf`; digest `sha256:ddddb543e61078456915a99cb29c3a42a961001d81b7eb50761d3e637258e884`.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```
