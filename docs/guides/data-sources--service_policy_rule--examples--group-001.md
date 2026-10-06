---
page_title: "xcsh_service_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule examples."
---

# xcsh_service_policy_rule examples

<a id="canonical-1223112310330201-1133220313031233-2011110022130103-3103332022313100-1011012222033323-0131131222332313-3103032031112112-0002010031221231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- Examples

<a id="canonical-0320320001313302-3020221220233113-1110111331313200-1322000333013131-1123012133102113-3203012023230133-3200312200011313-0030123232122322"></a>

### Complete configurations for `xcsh_service_policy_rule`

- [Data source](data-sources--service_policy_rule--examples--group-001.md#canonical-3320121011320201-2020022320330013-1010121120101230-3200001120211013-2022023012201323-0333232213320213-2103310221122313-3011332222023310): valid configuration.

<a id="canonical-3320121011320201-2020022320330013-1010121120101230-3200001120211013-2022023012201323-0333232213320213-2103310221122313-3011332222023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Examples](data-sources--service_policy_rule--examples--group-001.md#canonical-1223112310330201-1133220313031233-2011110022130103-3103332022313100-1011012222033323-0131131222332313-3103032031112112-0002010031221231)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_rule/data-source.tf`; digest `sha256:ae57799b0f1696fd5e6a11587ed3b1a1bc30f50fe8953264b9b955e84a34c726`.

```terraform
# ServicePolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicyRule by name
data "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"
}

output "service_policy_rule_id" {
  value = data.xcsh_service_policy_rule.example.id
}
```
